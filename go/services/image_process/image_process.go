package service

import (
	"io"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func GetImageBytes(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}
func ParseImagePathToText(path string) (string, error) {
	imgBytes, err := GetImageBytes(path)
	if err != nil {
		return "", err
	}
	result, error := ParseImageToText(imgBytes)
	if error != nil {
		return "", error
	}
	return result, nil
}
func ParseImageToText(data []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImageFromBytes(data)
	client.SetLanguage("eng", "tha")
	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}
