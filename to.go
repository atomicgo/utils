package utils

import (
	"atomicgo.dev/constraints"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ToJSON converts the given value to a JSON string.
func ToJSON(v any) (string, error) {
	r, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func ToPrettyJSON(v any, indent ...string) (string, error) {
	if len(indent) == 0 {
		indent = append(indent, "  ")
	}

	r, err := json.MarshalIndent(v, "", indent[0])
	if err != nil {
		return "", err
	}
	return string(r), nil
}

// ToString converts the given value to a string.
func ToString(v any) string {
	return fmt.Sprint(v)
}

// ToInt converts the given value to an int.
// If the value is a float, it will be rounded to the nearest integer. (Rounds up if the decimal is 0.5 or higher)
func ToInt[T string | constraints.Number](v T) int {
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
		v += 0.5
		return int(v)
	case float64:
		v += 0.5
		return int(v)
	case string:
		var i int

		if !strings.Contains(v, ".") {
			i, _ = strconv.Atoi(v)
		} else {
			f, _ := strconv.ParseFloat(v, 64)
			f += 0.5 // Round up
			i = int(f)
		}

		return i
	default:
		return 0
	}
}
