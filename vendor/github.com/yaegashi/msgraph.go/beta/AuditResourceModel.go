// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AuditResource undocumented
type AuditResource struct {
	// Object is the base model of AuditResource
	Object
	// DisplayName Display name.
	DisplayName *string `json:"displayName,omitempty"`
	// ModifiedProperties List of modified properties.
	ModifiedProperties []AuditProperty `json:"modifiedProperties,omitempty"`
	// Type Audit resource's type.
	Type *string `json:"type,omitempty"`
	// ResourceID Audit resource's Id.
	ResourceID *string `json:"resourceId,omitempty"`
}
