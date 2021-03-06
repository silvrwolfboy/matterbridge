// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ApplicationAddKeyRequestParameter undocumented
type ApplicationAddKeyRequestParameter struct {
	// KeyCredential undocumented
	KeyCredential *KeyCredential `json:"keyCredential,omitempty"`
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationAddPasswordRequestParameter undocumented
type ApplicationAddPasswordRequestParameter struct {
	// PasswordCredential undocumented
	PasswordCredential *PasswordCredential `json:"passwordCredential,omitempty"`
}

// ApplicationRemoveKeyRequestParameter undocumented
type ApplicationRemoveKeyRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// Proof undocumented
	Proof *string `json:"proof,omitempty"`
}

// ApplicationRemovePasswordRequestParameter undocumented
type ApplicationRemovePasswordRequestParameter struct {
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
}

//
type ApplicationAddKeyRequestBuilder struct{ BaseRequestBuilder }

// AddKey action undocumented
func (b *ApplicationRequestBuilder) AddKey(reqObj *ApplicationAddKeyRequestParameter) *ApplicationAddKeyRequestBuilder {
	bb := &ApplicationAddKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationAddKeyRequest struct{ BaseRequest }

//
func (b *ApplicationAddKeyRequestBuilder) Request() *ApplicationAddKeyRequest {
	return &ApplicationAddKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationAddKeyRequest) Post(ctx context.Context) (resObj *KeyCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ApplicationAddPasswordRequestBuilder struct{ BaseRequestBuilder }

// AddPassword action undocumented
func (b *ApplicationRequestBuilder) AddPassword(reqObj *ApplicationAddPasswordRequestParameter) *ApplicationAddPasswordRequestBuilder {
	bb := &ApplicationAddPasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/addPassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationAddPasswordRequest struct{ BaseRequest }

//
func (b *ApplicationAddPasswordRequestBuilder) Request() *ApplicationAddPasswordRequest {
	return &ApplicationAddPasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationAddPasswordRequest) Post(ctx context.Context) (resObj *PasswordCredential, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}

//
type ApplicationRemoveKeyRequestBuilder struct{ BaseRequestBuilder }

// RemoveKey action undocumented
func (b *ApplicationRequestBuilder) RemoveKey(reqObj *ApplicationRemoveKeyRequestParameter) *ApplicationRemoveKeyRequestBuilder {
	bb := &ApplicationRemoveKeyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removeKey"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationRemoveKeyRequest struct{ BaseRequest }

//
func (b *ApplicationRemoveKeyRequestBuilder) Request() *ApplicationRemoveKeyRequest {
	return &ApplicationRemoveKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationRemoveKeyRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type ApplicationRemovePasswordRequestBuilder struct{ BaseRequestBuilder }

// RemovePassword action undocumented
func (b *ApplicationRequestBuilder) RemovePassword(reqObj *ApplicationRemovePasswordRequestParameter) *ApplicationRemovePasswordRequestBuilder {
	bb := &ApplicationRemovePasswordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/removePassword"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ApplicationRemovePasswordRequest struct{ BaseRequest }

//
func (b *ApplicationRemovePasswordRequestBuilder) Request() *ApplicationRemovePasswordRequest {
	return &ApplicationRemovePasswordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ApplicationRemovePasswordRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
