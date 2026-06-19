package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type VideoService struct{}

func NewVideoService() *VideoService {
	return &VideoService{}
}

// UploadToStorage simulates saving to S3/CDN or local disk
func (s *VideoService) UploadToStorage(file io.Reader, filename string) (string, error) {
	// Clean the filename to prevent directory traversal attacks
	cleanName := filepath.Base(filename)
	targetPath := filepath.Join("./storage", fmt.Sprintf("%d_%s", time.Now().Unix(), cleanName))

	dst, err := os.Create(targetPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Stream the data chunks smoothly from the upload body into the destination file
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	// In real life, you would push this file to AWS S3 / Cloudflare R2 and get a CDN URL.
	// For now, we return a mock URL pointing to our media asset pipeline.
	mockCDNURL := fmt.Sprintf("https://cdn.myapp.com/storage/%s", filepath.Base(targetPath))
	return mockCDNURL, nil
}
