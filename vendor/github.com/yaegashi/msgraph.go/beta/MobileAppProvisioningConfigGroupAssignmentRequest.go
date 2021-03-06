// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MobileAppProvisioningConfigGroupAssignmentRequestBuilder is request builder for MobileAppProvisioningConfigGroupAssignment
type MobileAppProvisioningConfigGroupAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppProvisioningConfigGroupAssignmentRequest
func (b *MobileAppProvisioningConfigGroupAssignmentRequestBuilder) Request() *MobileAppProvisioningConfigGroupAssignmentRequest {
	return &MobileAppProvisioningConfigGroupAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppProvisioningConfigGroupAssignmentRequest is request for MobileAppProvisioningConfigGroupAssignment
type MobileAppProvisioningConfigGroupAssignmentRequest struct{ BaseRequest }

// Get performs GET request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Get(ctx context.Context) (resObj *MobileAppProvisioningConfigGroupAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Update(ctx context.Context, reqObj *MobileAppProvisioningConfigGroupAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
