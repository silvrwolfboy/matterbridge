// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ThreatAssessmentStatus undocumented
type ThreatAssessmentStatus int

const (
	// ThreatAssessmentStatusVPending undocumented
	ThreatAssessmentStatusVPending ThreatAssessmentStatus = 1
	// ThreatAssessmentStatusVCompleted undocumented
	ThreatAssessmentStatusVCompleted ThreatAssessmentStatus = 2
)

// ThreatAssessmentStatusPPending returns a pointer to ThreatAssessmentStatusVPending
func ThreatAssessmentStatusPPending() *ThreatAssessmentStatus {
	v := ThreatAssessmentStatusVPending
	return &v
}

// ThreatAssessmentStatusPCompleted returns a pointer to ThreatAssessmentStatusVCompleted
func ThreatAssessmentStatusPCompleted() *ThreatAssessmentStatus {
	v := ThreatAssessmentStatusVCompleted
	return &v
}
