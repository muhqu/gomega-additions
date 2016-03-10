package worddiff

type GomegaMatcher interface {
	Match(actual interface{}) (success bool, err error)
	FailureMessage(actual interface{}) (message string)
	NegatedFailureMessage(actual interface{}) (message string)
}

func MatchJSON(actual interface{}) GomegaMatcher {
	return &jsonMatcher{
		JSONToMatch: actual,
	}
}
