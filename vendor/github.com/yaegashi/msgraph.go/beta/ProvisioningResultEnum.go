// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProvisioningResult undocumented
type ProvisioningResult int

const (
	// ProvisioningResultVSuccess undocumented
	ProvisioningResultVSuccess ProvisioningResult = 0
	// ProvisioningResultVFailure undocumented
	ProvisioningResultVFailure ProvisioningResult = 1
	// ProvisioningResultVSkipped undocumented
	ProvisioningResultVSkipped ProvisioningResult = 2
	// ProvisioningResultVUnknownFutureValue undocumented
	ProvisioningResultVUnknownFutureValue ProvisioningResult = 3
)

// ProvisioningResultPSuccess returns a pointer to ProvisioningResultVSuccess
func ProvisioningResultPSuccess() *ProvisioningResult {
	v := ProvisioningResultVSuccess
	return &v
}

// ProvisioningResultPFailure returns a pointer to ProvisioningResultVFailure
func ProvisioningResultPFailure() *ProvisioningResult {
	v := ProvisioningResultVFailure
	return &v
}

// ProvisioningResultPSkipped returns a pointer to ProvisioningResultVSkipped
func ProvisioningResultPSkipped() *ProvisioningResult {
	v := ProvisioningResultVSkipped
	return &v
}

// ProvisioningResultPUnknownFutureValue returns a pointer to ProvisioningResultVUnknownFutureValue
func ProvisioningResultPUnknownFutureValue() *ProvisioningResult {
	v := ProvisioningResultVUnknownFutureValue
	return &v
}
