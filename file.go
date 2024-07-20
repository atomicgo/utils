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
		return "", err
	}

	return string(res), nil
}

// WriteFile writes the given content to the given file.
// Accepts a string or a byte slice.
func WriteFile[T string | []byte](path string, content T) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// AppendToFile appends the given content to the given file.
// Accepts a string or a byte slice.
func AppendToFile[T string | []byte](path string, content T) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write([]byte(content)); err != nil {
		return err
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
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
