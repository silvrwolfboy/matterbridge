// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Filter undocumented
type Filter struct {
	// Object is the base model of Filter
	Object
	// Groups undocumented
	Groups []FilterGroup `json:"groups,omitempty"`
	// InputFilterGroups undocumented
	InputFilterGroups []FilterGroup `json:"inputFilterGroups,omitempty"`
	// CategoryFilterGroups undocumented
	CategoryFilterGroups []FilterGroup `json:"categoryFilterGroups,omitempty"`
}
