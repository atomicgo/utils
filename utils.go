package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Ternary is a ternary operator.
// It returns a if the condition is true, otherwise it returns b.
func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// PrettyJSON returns a pretty-printed JSON string.
// If indent is not provided, it defaults to "  " (two spaces).
func PrettyJSON(inputJSON string, indent ...string) (string, error) {
	var out bytes.Buffer
	if len(indent) == 0 {
		indent = append(indent, "  ")
	}

	err := json.Indent(&out, []byte(inputJSON), "", indent[0])
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// Fetch returns the body of a GET request to the given URL.
func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
