// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AccessPackageResourceRoleRequestBuilder is request builder for AccessPackageResourceRole
type AccessPackageResourceRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageResourceRoleRequest
func (b *AccessPackageResourceRoleRequestBuilder) Request() *AccessPackageResourceRoleRequest {
	return &AccessPackageResourceRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageResourceRoleRequest is request for AccessPackageResourceRole
type AccessPackageResourceRoleRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageResourceRole
func (r *AccessPackageResourceRoleRequest) Get(ctx context.Context) (resObj *AccessPackageResourceRole, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageResourceRole
func (r *AccessPackageResourceRoleRequest) Update(ctx context.Context, reqObj *AccessPackageResourceRole) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageResourceRole
func (r *AccessPackageResourceRoleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageResource is navigation property
func (b *AccessPackageResourceRoleRequestBuilder) AccessPackageResource() *AccessPackageResourceRequestBuilder {
	bb := &AccessPackageResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResource"
	return bb
}
