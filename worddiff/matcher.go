// Worddiff Package provides GomegaMatcher that give improved failure messages using `git diff --word-diff=color`.
package worddiff

import "github.com/onsi/gomega/types"

func MatchJSON(actual interface{}) types.GomegaMatcher {
	return &jsonMatcher{
		JSONToMatch: actual,
	}
}
