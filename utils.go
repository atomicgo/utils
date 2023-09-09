package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Ternary is a ternary operator.
// It returns a if the condition is true, otherwise it returns b.
func Ternary[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

// region ToXxx

// ToJSON converts the given value to a JSON string.
func ToJSON(v any) (string, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

// ToString converts the given value to a string.
func ToString(v any) string {
	return fmt.Sprint(v)
}

func ToInt[T string | number](v T) int {
	switch v := any(v).(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case uint:
		return int(v)
	case uint8:
		return int(v)
	case uint16:
		return int(v)
	case uint32:
		return int(v)
	case uint64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

//endregion

//region FileXxx

// FileRead reads the given file and returns its content.
func FileRead(path string) (string, error) {
	res, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// FileWrite writes the given content to the given file.
func FileWrite(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// FileAppend appends the given content to the given file.
func FileAppend(path, content string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		return err
	}

	return nil
}

// FileExists returns true if the given file exists.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileDownload downloads the given URL to the given path.
// If the file already exists, it will be overwritten.
func FileDownload(url, path string) error {
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

//endregion
