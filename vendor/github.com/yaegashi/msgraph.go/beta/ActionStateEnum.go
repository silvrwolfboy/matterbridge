// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ActionState undocumented
type ActionState int

const (
	// ActionStateVNone undocumented
	ActionStateVNone ActionState = 0
	// ActionStateVPending undocumented
	ActionStateVPending ActionState = 1
	// ActionStateVCanceled undocumented
	ActionStateVCanceled ActionState = 2
	// ActionStateVActive undocumented
	ActionStateVActive ActionState = 3
	// ActionStateVDone undocumented
	ActionStateVDone ActionState = 4
	// ActionStateVFailed undocumented
	ActionStateVFailed ActionState = 5
	// ActionStateVNotSupported undocumented
	ActionStateVNotSupported ActionState = 6
)

// ActionStatePNone returns a pointer to ActionStateVNone
func ActionStatePNone() *ActionState {
	v := ActionStateVNone
	return &v
}

// ActionStatePPending returns a pointer to ActionStateVPending
func ActionStatePPending() *ActionState {
	v := ActionStateVPending
	return &v
}

// ActionStatePCanceled returns a pointer to ActionStateVCanceled
func ActionStatePCanceled() *ActionState {
	v := ActionStateVCanceled
	return &v
}

// ActionStatePActive returns a pointer to ActionStateVActive
func ActionStatePActive() *ActionState {
	v := ActionStateVActive
	return &v
}

// ActionStatePDone returns a pointer to ActionStateVDone
func ActionStatePDone() *ActionState {
	v := ActionStateVDone
	return &v
}

// ActionStatePFailed returns a pointer to ActionStateVFailed
func ActionStatePFailed() *ActionState {
	v := ActionStateVFailed
	return &v
}

// ActionStatePNotSupported returns a pointer to ActionStateVNotSupported
func ActionStatePNotSupported() *ActionState {
	v := ActionStateVNotSupported
	return &v
}
