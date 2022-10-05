package shared

import (
	"time"
)

type APIEndpoint struct {
	APIEndpointID string    `json:"api_endpoint_id"`
	APIID         string    `json:"api_id"`
	CreatedAt     time.Time `json:"created_at"`
	Description   string    `json:"description"`
	DisplayName   string    `json:"display_name"`
	Matched       *bool     `json:"matched"`
	Method        string    `json:"method"`
	Path          string    `json:"path"`
	UpdatedAt     time.Time `json:"updated_at"`
	VersionID     string    `json:"version_id"`
	WorkspaceID   string    `json:"workspace_id"`
}
