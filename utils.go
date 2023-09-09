package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

func ToJSON(v any) (string, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(r), nil
}

// Ternary is a ternary operator.
// It returns a if the condition is true, otherwise it returns b.
func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}

	return b
}

type file struct{}

var File = file{}

func (file) Read(path string) (string, error) {
	r, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(r), nil
}

func (file) Write(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

type net struct {
	Request request
}

var Net = net{
	Request: request{},
}

type request struct{}

func (request) Get(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (request) Post(url, contentType, body string) (string, error) {
	r, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
