package ocr

import (
	"io"
	"mime/multipart"
	"os"
	"sync"

	"github.com/labstack/echo/v4"
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

func (s *OCRService) ProcessImages(c echo.Context, reader *multipart.Reader) ([]string, error) {
	var results []string
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		res, err := func() (string, error) {
			defer part.Close()

			if part.FormName() != "images" {
				return "", nil
			}
			tmpFile, err := os.CreateTemp("", "ocr-*.png")
			if err != nil {
				return "", err
			}
			tmpPath := tmpFile.Name()

			defer os.Remove(tmpPath)
			defer tmpFile.Close()

			if _, err = io.Copy(tmpFile, part); err != nil {
				return "", err
			}
			tmpFile.Close()
			s.mu.Lock()
			defer s.mu.Unlock()
			select {
			case <-c.Request().Context().Done():
				return "", c.Request().Context().Err()
			default:
			}

			if err := s.client.SetImage(tmpPath); err != nil {
				return "", err
			}
			return s.client.Text()
		}()

		if err != nil {
			return nil, err
		}
		if res != "" {
			results = append(results, res)
		}
	}
	return results, nil
}

func (s *OCRService) Close() {
	s.client.Close()
}
