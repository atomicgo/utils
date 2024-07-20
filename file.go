package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ReadFile reads the given file and returns its content.
func ReadFile(path string) (string, error) {
	res, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not read file: %w", err)
	}

	return string(res), nil
}

// WriteFile writes the given content to the given file.
// Accepts a string or a byte slice.
func WriteFile[T string | []byte](path string, content T) error {
	err := os.WriteFile(path, []byte(content), 0o600)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}

// AppendToFile appends the given content to the given file.
// Accepts a string or a byte slice.
func AppendToFile[T string | []byte](path string, content T) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	if _, err = file.Write([]byte(content)); err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}

	return nil
}

// FileExists returns true if the given file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// DownloadFile downloads the given URL to the given path.
// If the file already exists, it will be overwritten.
func DownloadFile(url, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer out.Close()

	resp, err := http.Get(url) //nolint:gosec // wrapper function
	if err != nil {
		return fmt.Errorf("could not download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status) //nolint: err113
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}
