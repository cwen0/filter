package proxy


import (
	"log"
	"strings"
)

// Rule is rule
// pct(5)->delay(1000)|pct(1)->timeout will convert to
type Rule struct {
	Side       string
	Pct        string
	Action     string
	ActionArgs string
}

// pct(5)->delay(1000)|pct(1)->timeout will convert to rule{}
func getRulesFromRuleStr(rulesStr string) []*Rule {
	var rules []*Rule
	items := strings.Split(rulesStr, "|")
	for _, item := range items {
		rule := stringToRule(item)
		rules = append(rules, rule)
	}
	return rules
}

func stringToRule(ruleStr string) *Rule {
	rule := &Rule{}
	items := strings.Split(ruleStr, "->")
	for _, item := range items {
		if strings.HasPrefix(item, "pct(") {
			arr := strings.SplitN(item, "(", 2)
			pct := strings.Split(arr[1], ")")[0]
			rule.Pct = pct
		} else {
			arr := strings.SplitN(item, "(", 2)
			rule.Action = arr[0]
			log.Print(arr[1])
			rule.ActionArgs = strings.Split(arr[1], ")")[0]
		}
	}
	rule.Side = "right"
	return rule
}