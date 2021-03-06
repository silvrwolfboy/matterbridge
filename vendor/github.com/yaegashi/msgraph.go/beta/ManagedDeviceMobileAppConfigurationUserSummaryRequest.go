// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationUserSummary
type ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationUserSummaryRequest
func (b *ManagedDeviceMobileAppConfigurationUserSummaryRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationUserSummaryRequest {
	return &ManagedDeviceMobileAppConfigurationUserSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationUserSummaryRequest is request for ManagedDeviceMobileAppConfigurationUserSummary
type ManagedDeviceMobileAppConfigurationUserSummaryRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Get(ctx context.Context) (resObj *ManagedDeviceMobileAppConfigurationUserSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Update(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationUserSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationUserSummary
func (r *ManagedDeviceMobileAppConfigurationUserSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
