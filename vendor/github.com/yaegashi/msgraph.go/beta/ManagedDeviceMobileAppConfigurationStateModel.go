// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceMobileAppConfigurationState Managed Device Mobile App Configuration State for a given device.
type ManagedDeviceMobileAppConfigurationState struct {
	// Entity is the base model of ManagedDeviceMobileAppConfigurationState
	Entity
	// DisplayName The name of the policy for this policyBase
	DisplayName *string `json:"displayName,omitempty"`
	// Version The version of the policy
	Version *int `json:"version,omitempty"`
	// PlatformType Platform type that the policy applies to
	PlatformType *PolicyPlatformType `json:"platformType,omitempty"`
	// State The compliance state of the policy
	State *ComplianceStatus `json:"state,omitempty"`
	// SettingCount Count of how many setting a policy holds
	SettingCount *int `json:"settingCount,omitempty"`
	// UserID User unique identifier, must be Guid
	UserID *string `json:"userId,omitempty"`
	// UserPrincipalName User Principal Name
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
