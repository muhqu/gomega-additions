// Unifieddiff Package provides GomegaMatcher that give improved failure messages using `git diff -u --color`.
package unifieddiff

import "github.com/onsi/gomega/types"

func MatchJSON(actual interface{}) types.GomegaMatcher {
	return &jsonMatcher{
		JSONToMatch: actual,
	}
}
