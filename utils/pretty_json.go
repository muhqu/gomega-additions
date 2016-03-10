package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func PrettyJson(value interface{}) (string, error) {
	var str string
	switch value := value.(type) {
	case string:
		str = value
	case fmt.Stringer:
		str = value.String()
	default:
		return "", errors.New("expected json, but it's not even a string")
	}

	bytes := new(bytes.Buffer)
	err := json.Indent(bytes, []byte(str), "", "  ")
	if err != nil {
		return "", errors.New("expected json, but json is invalid")
	}
	return bytes.String(), nil
}
