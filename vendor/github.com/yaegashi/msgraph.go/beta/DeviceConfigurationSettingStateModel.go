// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationSettingState undocumented
type DeviceConfigurationSettingState struct {
	// Object is the base model of DeviceConfigurationSettingState
	Object
	// Setting The setting that is being reported
	Setting *string `json:"setting,omitempty"`
	// SettingName Localized/user friendly setting name that is being reported
	SettingName *string `json:"settingName,omitempty"`
	// InstanceDisplayName Name of setting instance that is being reported.
	InstanceDisplayName *string `json:"instanceDisplayName,omitempty"`
	// State The compliance state of the setting
	State *ComplianceStatus `json:"state,omitempty"`
	// ErrorCode Error code for the setting
	ErrorCode *int `json:"errorCode,omitempty"`
	// ErrorDescription Error description
	ErrorDescription *string `json:"errorDescription,omitempty"`
	// UserID UserId
	UserID *string `json:"userId,omitempty"`
	// UserName UserName
	UserName *string `json:"userName,omitempty"`
	// UserEmail UserEmail
	UserEmail *string `json:"userEmail,omitempty"`
	// UserPrincipalName UserPrincipalName.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// Sources Contributing policies
	Sources []SettingSource `json:"sources,omitempty"`
	// CurrentValue Current value of setting on device
	CurrentValue *string `json:"currentValue,omitempty"`
}
