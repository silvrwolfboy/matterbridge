// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationExternalSource undocumented
type EducationExternalSource int

const (
	// EducationExternalSourceVSis undocumented
	EducationExternalSourceVSis EducationExternalSource = 0
	// EducationExternalSourceVManual undocumented
	EducationExternalSourceVManual EducationExternalSource = 1
	// EducationExternalSourceVUnknownFutureValue undocumented
	EducationExternalSourceVUnknownFutureValue EducationExternalSource = 2
)

// EducationExternalSourcePSis returns a pointer to EducationExternalSourceVSis
func EducationExternalSourcePSis() *EducationExternalSource {
	v := EducationExternalSourceVSis
	return &v
}

// EducationExternalSourcePManual returns a pointer to EducationExternalSourceVManual
func EducationExternalSourcePManual() *EducationExternalSource {
	v := EducationExternalSourceVManual
	return &v
}

// EducationExternalSourcePUnknownFutureValue returns a pointer to EducationExternalSourceVUnknownFutureValue
func EducationExternalSourcePUnknownFutureValue() *EducationExternalSource {
	v := EducationExternalSourceVUnknownFutureValue
	return &v
}