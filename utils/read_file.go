package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type File struct {
	Name string
	Meta map[string]string
	File io.Reader
}

func NewFileFromImage(data []byte) (file *File, err error) {
	fileName := fmt.Sprintf("%s.png", RandomString(12))
	tempDir := os.TempDir()
	tempFilePath := filepath.Join(tempDir, fileName)
	err = os.WriteFile(tempFilePath, data, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %v", err)
	}

	reader, _ := os.Open(tempFilePath)
	metadata := make(map[string]string)

	return &File{
		Name: fileName,
		Meta: metadata,
		File: reader,
	}, nil
}
