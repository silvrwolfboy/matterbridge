// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DocumentCommentRequestBuilder is request builder for DocumentComment
type DocumentCommentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DocumentCommentRequest
func (b *DocumentCommentRequestBuilder) Request() *DocumentCommentRequest {
	return &DocumentCommentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DocumentCommentRequest is request for DocumentComment
type DocumentCommentRequest struct{ BaseRequest }

// Get performs GET request for DocumentComment
func (r *DocumentCommentRequest) Get(ctx context.Context) (resObj *DocumentComment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DocumentComment
func (r *DocumentCommentRequest) Update(ctx context.Context, reqObj *DocumentComment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DocumentComment
func (r *DocumentCommentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Replies returns request builder for DocumentCommentReply collection
func (b *DocumentCommentRequestBuilder) Replies() *DocumentCommentRepliesCollectionRequestBuilder {
	bb := &DocumentCommentRepliesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/replies"
	return bb
}

// DocumentCommentRepliesCollectionRequestBuilder is request builder for DocumentCommentReply collection
type DocumentCommentRepliesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DocumentCommentReply collection
func (b *DocumentCommentRepliesCollectionRequestBuilder) Request() *DocumentCommentRepliesCollectionRequest {
	return &DocumentCommentRepliesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DocumentCommentReply item
func (b *DocumentCommentRepliesCollectionRequestBuilder) ID(id string) *DocumentCommentReplyRequestBuilder {
	bb := &DocumentCommentReplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DocumentCommentRepliesCollectionRequest is request for DocumentCommentReply collection
type DocumentCommentRepliesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DocumentCommentReply, error) {
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
	var values []DocumentCommentReply
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
			value  []DocumentCommentReply
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

// Get performs GET request for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Get(ctx context.Context) ([]DocumentCommentReply, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DocumentCommentReply collection
func (r *DocumentCommentRepliesCollectionRequest) Add(ctx context.Context, reqObj *DocumentCommentReply) (resObj *DocumentCommentReply, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
