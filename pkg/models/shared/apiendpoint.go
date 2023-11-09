// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v2/pkg/utils"
	"time"
)

// APIEndpoint - An ApiEndpoint is a description of an Endpoint for an API.
type APIEndpoint struct {
	// The ID of this ApiEndpoint. This is a hash of the method and path.
	APIEndpointID string `json:"api_endpoint_id"`
	// The ID of the Api this ApiEndpoint belongs to.
	APIID string `json:"api_id"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// A detailed description of the ApiEndpoint.
	Description string `json:"description"`
	// A human-readable name for the ApiEndpoint.
	DisplayName string `json:"display_name"`
	// Determines if the endpoint was found in the OpenAPI spec associated with the parent Api.
	Matched *bool `json:"matched,omitempty"`
	// HTTP verb.
	Method string `json:"method"`
	// Path that handles this Api.
	Path string `json:"path"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at"`
	// The version ID of the Api this ApiEndpoint belongs to.
	VersionID string `json:"version_id"`
	// The workspace ID this ApiEndpoint belongs to.
	WorkspaceID string `json:"workspace_id"`
}

func (a APIEndpoint) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIEndpoint) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIEndpoint) GetAPIEndpointID() string {
	if o == nil {
		return ""
	}
	return o.APIEndpointID
}

func (o *APIEndpoint) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *APIEndpoint) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *APIEndpoint) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *APIEndpoint) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *APIEndpoint) GetMatched() *bool {
	if o == nil {
		return nil
	}
	return o.Matched
}

func (o *APIEndpoint) GetMethod() string {
	if o == nil {
		return ""
	}
	return o.Method
}

func (o *APIEndpoint) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *APIEndpoint) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *APIEndpoint) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

func (o *APIEndpoint) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
