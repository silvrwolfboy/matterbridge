// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkforceIntegrationEncryption undocumented
type WorkforceIntegrationEncryption struct {
	// Object is the base model of WorkforceIntegrationEncryption
	Object
	// Protocol undocumented
	Protocol *WorkforceIntegrationEncryptionProtocol `json:"protocol,omitempty"`
	// Secret undocumented
	Secret *string `json:"secret,omitempty"`
}
