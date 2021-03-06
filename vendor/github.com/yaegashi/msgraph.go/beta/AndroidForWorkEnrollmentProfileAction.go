// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkEnrollmentProfileRevokeTokenRequestParameter undocumented
type AndroidForWorkEnrollmentProfileRevokeTokenRequestParameter struct {
}

// AndroidForWorkEnrollmentProfileCreateTokenRequestParameter undocumented
type AndroidForWorkEnrollmentProfileCreateTokenRequestParameter struct {
	// TokenValidityInSeconds undocumented
	TokenValidityInSeconds *int `json:"tokenValidityInSeconds,omitempty"`
}

//
type AndroidForWorkEnrollmentProfileRevokeTokenRequestBuilder struct{ BaseRequestBuilder }

// RevokeToken action undocumented
func (b *AndroidForWorkEnrollmentProfileRequestBuilder) RevokeToken(reqObj *AndroidForWorkEnrollmentProfileRevokeTokenRequestParameter) *AndroidForWorkEnrollmentProfileRevokeTokenRequestBuilder {
	bb := &AndroidForWorkEnrollmentProfileRevokeTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/revokeToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkEnrollmentProfileRevokeTokenRequest struct{ BaseRequest }

//
func (b *AndroidForWorkEnrollmentProfileRevokeTokenRequestBuilder) Request() *AndroidForWorkEnrollmentProfileRevokeTokenRequest {
	return &AndroidForWorkEnrollmentProfileRevokeTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkEnrollmentProfileRevokeTokenRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type AndroidForWorkEnrollmentProfileCreateTokenRequestBuilder struct{ BaseRequestBuilder }

// CreateToken action undocumented
func (b *AndroidForWorkEnrollmentProfileRequestBuilder) CreateToken(reqObj *AndroidForWorkEnrollmentProfileCreateTokenRequestParameter) *AndroidForWorkEnrollmentProfileCreateTokenRequestBuilder {
	bb := &AndroidForWorkEnrollmentProfileCreateTokenRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/createToken"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkEnrollmentProfileCreateTokenRequest struct{ BaseRequest }

//
func (b *AndroidForWorkEnrollmentProfileCreateTokenRequestBuilder) Request() *AndroidForWorkEnrollmentProfileCreateTokenRequest {
	return &AndroidForWorkEnrollmentProfileCreateTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkEnrollmentProfileCreateTokenRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
