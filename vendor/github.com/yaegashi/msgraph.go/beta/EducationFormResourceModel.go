// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationFormResource undocumented
type EducationFormResource struct {
	// EducationResource is the base model of EducationFormResource
	EducationResource
	// OriginalFormID undocumented
	OriginalFormID *string `json:"originalFormId,omitempty"`
	// FormID undocumented
	FormID *string `json:"formId,omitempty"`
	// IsGroupForm undocumented
	IsGroupForm *bool `json:"isGroupForm,omitempty"`
	// ViewURL undocumented
	ViewURL *string `json:"viewUrl,omitempty"`
	// EditURL undocumented
	EditURL *string `json:"editUrl,omitempty"`
}
