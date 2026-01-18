package ocr

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"sync"

	"github.com/otiai10/gosseract/v2"
)

type OCRService struct {
	client *gosseract.Client
	mu     sync.Mutex
	queue  chan struct{}
}

func NewOCRService(maxQueue int) *OCRService {
	client := gosseract.NewClient()
	return &OCRService{
		client: client,
		queue:  make(chan struct{}, maxQueue),
	}
}

func (s *OCRService) ProcessImages(reader *multipart.Reader) ([]string, error) {
	select {
	case s.queue <- struct{}{}:
		defer func() { <-s.queue }()
	default:
		return nil, fmt.Errorf("queue full")
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	var results []string
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if part.FormName() != "images" {
			continue
		}
		if err != nil {
			return nil, err
		}

		tmpFile, err := os.CreateTemp("", "ocr-*.png")
		if err != nil {
			return nil, err
		}
		tmpPath := tmpFile.Name()
		_, err = io.Copy(tmpFile, part)
		tmpFile.Close()
		s.client.SetImage(tmpPath)
		text, _ := s.client.Text()
		results = append(results, text)
		os.Remove(tmpFile.Name())
	}
	return results, nil
}

func (s *OCRService) Close() {
	s.client.Close()
}
