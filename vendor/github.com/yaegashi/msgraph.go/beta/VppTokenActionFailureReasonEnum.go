// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VppTokenActionFailureReason undocumented
type VppTokenActionFailureReason int

const (
	// VppTokenActionFailureReasonVNone undocumented
	VppTokenActionFailureReasonVNone VppTokenActionFailureReason = 0
	// VppTokenActionFailureReasonVAppleFailure undocumented
	VppTokenActionFailureReasonVAppleFailure VppTokenActionFailureReason = 1
	// VppTokenActionFailureReasonVInternalError undocumented
	VppTokenActionFailureReasonVInternalError VppTokenActionFailureReason = 2
	// VppTokenActionFailureReasonVExpiredVppToken undocumented
	VppTokenActionFailureReasonVExpiredVppToken VppTokenActionFailureReason = 3
	// VppTokenActionFailureReasonVExpiredApplePushNotificationCertificate undocumented
	VppTokenActionFailureReasonVExpiredApplePushNotificationCertificate VppTokenActionFailureReason = 4
)

// VppTokenActionFailureReasonPNone returns a pointer to VppTokenActionFailureReasonVNone
func VppTokenActionFailureReasonPNone() *VppTokenActionFailureReason {
	v := VppTokenActionFailureReasonVNone
	return &v
}

// VppTokenActionFailureReasonPAppleFailure returns a pointer to VppTokenActionFailureReasonVAppleFailure
func VppTokenActionFailureReasonPAppleFailure() *VppTokenActionFailureReason {
	v := VppTokenActionFailureReasonVAppleFailure
	return &v
}

// VppTokenActionFailureReasonPInternalError returns a pointer to VppTokenActionFailureReasonVInternalError
func VppTokenActionFailureReasonPInternalError() *VppTokenActionFailureReason {
	v := VppTokenActionFailureReasonVInternalError
	return &v
}

// VppTokenActionFailureReasonPExpiredVppToken returns a pointer to VppTokenActionFailureReasonVExpiredVppToken
func VppTokenActionFailureReasonPExpiredVppToken() *VppTokenActionFailureReason {
	v := VppTokenActionFailureReasonVExpiredVppToken
	return &v
}

// VppTokenActionFailureReasonPExpiredApplePushNotificationCertificate returns a pointer to VppTokenActionFailureReasonVExpiredApplePushNotificationCertificate
func VppTokenActionFailureReasonPExpiredApplePushNotificationCertificate() *VppTokenActionFailureReason {
	v := VppTokenActionFailureReasonVExpiredApplePushNotificationCertificate
	return &v
}
