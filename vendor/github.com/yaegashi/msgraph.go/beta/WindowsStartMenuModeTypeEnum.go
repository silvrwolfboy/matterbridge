// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsStartMenuModeType undocumented
type WindowsStartMenuModeType int

const (
	// WindowsStartMenuModeTypeVUserDefined undocumented
	WindowsStartMenuModeTypeVUserDefined WindowsStartMenuModeType = 0
	// WindowsStartMenuModeTypeVFullScreen undocumented
	WindowsStartMenuModeTypeVFullScreen WindowsStartMenuModeType = 1
	// WindowsStartMenuModeTypeVNonFullScreen undocumented
	WindowsStartMenuModeTypeVNonFullScreen WindowsStartMenuModeType = 2
)

// WindowsStartMenuModeTypePUserDefined returns a pointer to WindowsStartMenuModeTypeVUserDefined
func WindowsStartMenuModeTypePUserDefined() *WindowsStartMenuModeType {
	v := WindowsStartMenuModeTypeVUserDefined
	return &v
}

// WindowsStartMenuModeTypePFullScreen returns a pointer to WindowsStartMenuModeTypeVFullScreen
func WindowsStartMenuModeTypePFullScreen() *WindowsStartMenuModeType {
	v := WindowsStartMenuModeTypeVFullScreen
	return &v
}

// WindowsStartMenuModeTypePNonFullScreen returns a pointer to WindowsStartMenuModeTypeVNonFullScreen
func WindowsStartMenuModeTypePNonFullScreen() *WindowsStartMenuModeType {
	v := WindowsStartMenuModeTypeVNonFullScreen
	return &v
}
