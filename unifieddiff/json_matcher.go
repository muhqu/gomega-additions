package unifieddiff

import (
	"github.com/muhqu/gomega-additions/utils"
)

type jsonMatcher struct {
	JSONToMatch    interface{}
	PrettyExpected string
	PrettyActual   string
}

func (m *jsonMatcher) Match(actual interface{}) (success bool, err error) {
	if m.PrettyActual, err = utils.PrettyJson(actual); err != nil {
		return
	}
	if m.PrettyExpected, err = utils.PrettyJson(m.JSONToMatch); err != nil {
		return
	}
	return m.PrettyActual == m.PrettyExpected, nil
}

func (m *jsonMatcher) FailureMessage(actual interface{}) (message string) {
	coloredDiff, err := utils.GitDiff(m.PrettyActual, m.PrettyExpected, "-u", "--color")
	if err != nil {
		return "Expected to match the provided JSON. however, there was an issue creating the diff:\n" + err.Error()
	}
	return "Expected to match the provided JSON, here are the differences:\n" + coloredDiff
}

func (m *jsonMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Expected to NOT match the provided JSON, but it does: \n" +
		"\x1b[0m" + m.PrettyActual + "\x1b[0m"
}
