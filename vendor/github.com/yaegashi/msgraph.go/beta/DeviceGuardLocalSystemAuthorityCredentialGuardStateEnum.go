// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceGuardLocalSystemAuthorityCredentialGuardState undocumented
type DeviceGuardLocalSystemAuthorityCredentialGuardState int

const (
	// DeviceGuardLocalSystemAuthorityCredentialGuardStateVRunning undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardStateVRunning DeviceGuardLocalSystemAuthorityCredentialGuardState = 0
	// DeviceGuardLocalSystemAuthorityCredentialGuardStateVRebootRequired undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardStateVRebootRequired DeviceGuardLocalSystemAuthorityCredentialGuardState = 1
	// DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotLicensed undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotLicensed DeviceGuardLocalSystemAuthorityCredentialGuardState = 2
	// DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotConfigured undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotConfigured DeviceGuardLocalSystemAuthorityCredentialGuardState = 3
	// DeviceGuardLocalSystemAuthorityCredentialGuardStateVVirtualizationBasedSecurityNotRunning undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardStateVVirtualizationBasedSecurityNotRunning DeviceGuardLocalSystemAuthorityCredentialGuardState = 4
)

// DeviceGuardLocalSystemAuthorityCredentialGuardStatePRunning returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardStateVRunning
func DeviceGuardLocalSystemAuthorityCredentialGuardStatePRunning() *DeviceGuardLocalSystemAuthorityCredentialGuardState {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardStateVRunning
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardStatePRebootRequired returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardStateVRebootRequired
func DeviceGuardLocalSystemAuthorityCredentialGuardStatePRebootRequired() *DeviceGuardLocalSystemAuthorityCredentialGuardState {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardStateVRebootRequired
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardStatePNotLicensed returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotLicensed
func DeviceGuardLocalSystemAuthorityCredentialGuardStatePNotLicensed() *DeviceGuardLocalSystemAuthorityCredentialGuardState {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotLicensed
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardStatePNotConfigured returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotConfigured
func DeviceGuardLocalSystemAuthorityCredentialGuardStatePNotConfigured() *DeviceGuardLocalSystemAuthorityCredentialGuardState {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardStateVNotConfigured
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardStatePVirtualizationBasedSecurityNotRunning returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardStateVVirtualizationBasedSecurityNotRunning
func DeviceGuardLocalSystemAuthorityCredentialGuardStatePVirtualizationBasedSecurityNotRunning() *DeviceGuardLocalSystemAuthorityCredentialGuardState {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardStateVVirtualizationBasedSecurityNotRunning
	return &v
}
