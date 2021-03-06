// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookRangeFillClearRequestParameter undocumented
type WorkbookRangeFillClearRequestParameter struct {
}

//
type WorkbookRangeFillClearRequestBuilder struct{ BaseRequestBuilder }

// Clear action undocumented
func (b *WorkbookRangeFillRequestBuilder) Clear(reqObj *WorkbookRangeFillClearRequestParameter) *WorkbookRangeFillClearRequestBuilder {
	bb := &WorkbookRangeFillClearRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clear"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookRangeFillClearRequest struct{ BaseRequest }

//
func (b *WorkbookRangeFillClearRequestBuilder) Request() *WorkbookRangeFillClearRequest {
	return &WorkbookRangeFillClearRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookRangeFillClearRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
