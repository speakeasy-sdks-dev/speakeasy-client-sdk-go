// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

// API - An Api is representation of a API (a collection of API Endpoints) within the Speakeasy Platform.
type API struct {
	// The ID of this Api. This is a human-readable name (subject to change).
	APIID string `json:"api_id"`
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// A detailed description of the Api.
	Description string `json:"description"`
	// Determines if all the endpoints within the Api are found in the OpenAPI spec associated with the Api.
	Matched *bool `json:"matched,omitempty"`
	// A set of values associated with a meta_data key. This field is only set on get requests.
	MetaData map[string][]string `json:"meta_data,omitempty"`
	// Last update timestamp.
	UpdatedAt time.Time `json:"updated_at"`
	// The version ID of this Api. This is semantic version identifier.
	VersionID string `json:"version_id"`
	// The workspace ID this Api belongs to.
	WorkspaceID string `json:"workspace_id"`
}

func (a API) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *API) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *API) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *API) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *API) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *API) GetMatched() *bool {
	if o == nil {
		return nil
	}
	return o.Matched
}

func (o *API) GetMetaData() map[string][]string {
	if o == nil {
		return nil
	}
	return o.MetaData
}

func (o *API) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *API) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

func (o *API) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
