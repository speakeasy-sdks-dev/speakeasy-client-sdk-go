// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetWorkspaceAccessRequest struct {
	// Unique identifier of the generation target.
	GenLockID *string `queryParam:"style=form,explode=true,name=genLockId"`
	// Skip side-effects like incrementing metrics.
	Passive *bool `queryParam:"style=form,explode=true,name=passive"`
	// The type of the generated target.
	TargetType *string `queryParam:"style=form,explode=true,name=targetType"`
}

func (o *GetWorkspaceAccessRequest) GetGenLockID() *string {
	if o == nil {
		return nil
	}
	return o.GenLockID
}

func (o *GetWorkspaceAccessRequest) GetPassive() *bool {
	if o == nil {
		return nil
	}
	return o.Passive
}

func (o *GetWorkspaceAccessRequest) GetTargetType() *string {
	if o == nil {
		return nil
	}
	return o.TargetType
}

type GetWorkspaceAccessResponse struct {
	// OK
	AccessDetails *shared.AccessDetails
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetWorkspaceAccessResponse) GetAccessDetails() *shared.AccessDetails {
	if o == nil {
		return nil
	}
	return o.AccessDetails
}

func (o *GetWorkspaceAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkspaceAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkspaceAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}