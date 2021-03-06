// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AgreementAcceptanceRequestBuilder is request builder for AgreementAcceptance
type AgreementAcceptanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgreementAcceptanceRequest
func (b *AgreementAcceptanceRequestBuilder) Request() *AgreementAcceptanceRequest {
	return &AgreementAcceptanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgreementAcceptanceRequest is request for AgreementAcceptance
type AgreementAcceptanceRequest struct{ BaseRequest }

// Get performs GET request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Get(ctx context.Context) (resObj *AgreementAcceptance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Update(ctx context.Context, reqObj *AgreementAcceptance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AgreementAcceptance
func (r *AgreementAcceptanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
