// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetManifestRequest struct {
	NamespaceName     string `pathParam:"style=simple,explode=false,name=namespace_name"`
	OrganizationSlug  string `pathParam:"style=simple,explode=false,name=organization_slug"`
	RevisionReference string `pathParam:"style=simple,explode=false,name=revision_reference"`
	WorkspaceSlug     string `pathParam:"style=simple,explode=false,name=workspace_slug"`
}

func (o *GetManifestRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *GetManifestRequest) GetOrganizationSlug() string {
	if o == nil {
		return ""
	}
	return o.OrganizationSlug
}

func (o *GetManifestRequest) GetRevisionReference() string {
	if o == nil {
		return ""
	}
	return o.RevisionReference
}

func (o *GetManifestRequest) GetWorkspaceSlug() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceSlug
}

type GetManifestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// OK
	Manifest *shared.Manifest
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetManifestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetManifestResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetManifestResponse) GetManifest() *shared.Manifest {
	if o == nil {
		return nil
	}
	return o.Manifest
}

func (o *GetManifestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetManifestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}