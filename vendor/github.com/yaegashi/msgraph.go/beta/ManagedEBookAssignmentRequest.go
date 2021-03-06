// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedEBookAssignmentRequestBuilder is request builder for ManagedEBookAssignment
type ManagedEBookAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookAssignmentRequest
func (b *ManagedEBookAssignmentRequestBuilder) Request() *ManagedEBookAssignmentRequest {
	return &ManagedEBookAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookAssignmentRequest is request for ManagedEBookAssignment
type ManagedEBookAssignmentRequest struct{ BaseRequest }

// Get performs GET request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Get(ctx context.Context) (resObj *ManagedEBookAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Update(ctx context.Context, reqObj *ManagedEBookAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedEBookAssignment
func (r *ManagedEBookAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
