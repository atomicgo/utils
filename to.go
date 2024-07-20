package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"atomicgo.dev/constraints"
)

// ToJSON converts the given value to a JSON string.
func ToJSON(v any) (string, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("could not marshal json: %w", err)
	}

	return string(r), nil
}

func ToPrettyJSON(v any, indent ...string) (string, error) {
	if len(indent) == 0 {
		indent = append(indent, "  ")
	}

	r, err := json.MarshalIndent(v, "", indent[0])
	if err != nil {
		return "", fmt.Errorf("could not marshal json: %w", err)
	}

	return string(r), nil
}

// ToString converts the given value to a string.
func ToString(v any) string {
	return fmt.Sprint(v)
}

// ToInt converts the given value to an int.
// If the value is a float, it will be rounded to the nearest integer. (Rounds up if the decimal is 0.5 or higher)
func ToInt[T string | constraints.Number](value T) int {
	switch value := any(value).(type) {
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		return int(value)
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		return int(value)
	case uint64:
		return int(value)
	case float32:
		value += 0.5
		return int(value)
	case float64:
		value += 0.5
		return int(value)
	case string:
		var i int

		if !strings.Contains(value, ".") {
			i, _ = strconv.Atoi(value)
		} else {
			f, _ := strconv.ParseFloat(value, 64)
			f += 0.5 // Round up
			i = int(f)
		}

		return i

	default:
		return 0
	}
}
