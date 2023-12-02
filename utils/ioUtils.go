package utils

import (
	"io"
	"os"
)

func LoadTextFile(path string) (string, error) {
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
