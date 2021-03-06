// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OMASettingInteger undocumented
type OMASettingInteger struct {
	// OMASetting is the base model of OMASettingInteger
	OMASetting
	// Value Value.
	Value *int `json:"value,omitempty"`
	// IsReadOnly By setting to true, the CSP (configuration service provider) specified in the OMA-URI will perform a get, instead of set
	IsReadOnly *bool `json:"isReadOnly,omitempty"`
}
