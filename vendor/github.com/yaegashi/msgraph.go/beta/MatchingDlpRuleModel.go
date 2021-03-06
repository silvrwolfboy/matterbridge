// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MatchingDlpRule undocumented
type MatchingDlpRule struct {
	// Object is the base model of MatchingDlpRule
	Object
	// RuleID undocumented
	RuleID *string `json:"ruleId,omitempty"`
	// RuleName undocumented
	RuleName *string `json:"ruleName,omitempty"`
	// PolicyID undocumented
	PolicyID *string `json:"policyId,omitempty"`
	// PolicyName undocumented
	PolicyName *string `json:"policyName,omitempty"`
	// IsMostRestrictive undocumented
	IsMostRestrictive *bool `json:"isMostRestrictive,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
	// Actions undocumented
	Actions []DlpActionInfo `json:"actions,omitempty"`
	// RuleMode undocumented
	RuleMode *RuleMode `json:"ruleMode,omitempty"`
}
