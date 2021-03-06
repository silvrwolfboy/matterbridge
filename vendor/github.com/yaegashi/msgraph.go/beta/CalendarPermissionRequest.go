// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// CalendarPermissionRequestBuilder is request builder for CalendarPermission
type CalendarPermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarPermissionRequest
func (b *CalendarPermissionRequestBuilder) Request() *CalendarPermissionRequest {
	return &CalendarPermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarPermissionRequest is request for CalendarPermission
type CalendarPermissionRequest struct{ BaseRequest }

// Get performs GET request for CalendarPermission
func (r *CalendarPermissionRequest) Get(ctx context.Context) (resObj *CalendarPermission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarPermission
func (r *CalendarPermissionRequest) Update(ctx context.Context, reqObj *CalendarPermission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarPermission
func (r *CalendarPermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
