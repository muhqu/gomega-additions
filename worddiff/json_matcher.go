package worddiff

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type jsonMatcher struct {
	JSONToMatch    interface{}
	PrettyExpected string
	PrettyActual   string
}

func prettyJson(value interface{}) (string, error) {
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

func (m *jsonMatcher) Match(actual interface{}) (success bool, err error) {
	m.PrettyActual, err = prettyJson(actual)
	if err != nil {
		return
	}
	m.PrettyExpected, err = prettyJson(m.JSONToMatch)
	if err != nil {
		return
	}

	return m.PrettyActual == m.PrettyExpected, nil
}

func (m *jsonMatcher) FailureMessage(actual interface{}) (message string) {
	coloredDiff, err := wordDiff(m.PrettyActual, m.PrettyExpected)
	if err != nil {
		return "Expected to match the provided JSON. however, there was an issue creating the diff:\n" + err.Error()
	}
	return "Expected to match the provided JSON, here are the differences:\n" + coloredDiff
}

func (m *jsonMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Expected to NOT match the provided JSON, but it does: \n" +
		"\x1b[0m" + m.PrettyActual + "\x1b[0m"
}
