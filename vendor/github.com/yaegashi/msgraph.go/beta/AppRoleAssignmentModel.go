// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AppRoleAssignment undocumented
type AppRoleAssignment struct {
	// Entity is the base model of AppRoleAssignment
	Entity
	// AppRoleID undocumented
	AppRoleID *UUID `json:"appRoleId,omitempty"`
	// CreationTimestamp undocumented
	CreationTimestamp *time.Time `json:"creationTimestamp,omitempty"`
	// PrincipalDisplayName undocumented
	PrincipalDisplayName *string `json:"principalDisplayName,omitempty"`
	// PrincipalID undocumented
	PrincipalID *UUID `json:"principalId,omitempty"`
	// PrincipalType undocumented
	PrincipalType *string `json:"principalType,omitempty"`
	// ResourceDisplayName undocumented
	ResourceDisplayName *string `json:"resourceDisplayName,omitempty"`
	// ResourceID undocumented
	ResourceID *UUID `json:"resourceId,omitempty"`
}