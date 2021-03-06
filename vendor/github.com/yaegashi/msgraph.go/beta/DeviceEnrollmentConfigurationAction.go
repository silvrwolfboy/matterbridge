// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter undocumented
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

// DeviceEnrollmentConfigurationSetPriorityRequestParameter undocumented
type DeviceEnrollmentConfigurationSetPriorityRequestParameter struct {
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}

// DeviceEnrollmentConfigurationAssignRequestParameter undocumented
type DeviceEnrollmentConfigurationAssignRequestParameter struct {
	// EnrollmentConfigurationAssignments undocumented
	EnrollmentConfigurationAssignments []EnrollmentConfigurationAssignment `json:"enrollmentConfigurationAssignments,omitempty"`
}

//
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceManagementDeviceEnrollmentConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// HasPayloadLinks action undocumented
func (b *UserDeviceEnrollmentConfigurationsCollectionRequestBuilder) HasPayloadLinks(reqObj *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestParameter) *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder {
	bb := &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequestBuilder) Request() *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest {
	return &DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

//
func (r *DeviceEnrollmentConfigurationCollectionHasPayloadLinksRequest) Get(ctx context.Context) ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

//
type DeviceEnrollmentConfigurationSetPriorityRequestBuilder struct{ BaseRequestBuilder }

// SetPriority action undocumented
func (b *DeviceEnrollmentConfigurationRequestBuilder) SetPriority(reqObj *DeviceEnrollmentConfigurationSetPriorityRequestParameter) *DeviceEnrollmentConfigurationSetPriorityRequestBuilder {
	bb := &DeviceEnrollmentConfigurationSetPriorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setPriority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationSetPriorityRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationSetPriorityRequestBuilder) Request() *DeviceEnrollmentConfigurationSetPriorityRequest {
	return &DeviceEnrollmentConfigurationSetPriorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationSetPriorityRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type DeviceEnrollmentConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceEnrollmentConfigurationRequestBuilder) Assign(reqObj *DeviceEnrollmentConfigurationAssignRequestParameter) *DeviceEnrollmentConfigurationAssignRequestBuilder {
	bb := &DeviceEnrollmentConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceEnrollmentConfigurationAssignRequest struct{ BaseRequest }

//
func (b *DeviceEnrollmentConfigurationAssignRequestBuilder) Request() *DeviceEnrollmentConfigurationAssignRequest {
	return &DeviceEnrollmentConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceEnrollmentConfigurationAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}
