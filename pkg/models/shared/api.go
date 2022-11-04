package shared

import (
	"time"
)

// API
// An Api is representation of a API (a collection of API Endpoints) within the Speakeasy Platform.
type API struct {
	APIID       string              `json:"api_id"`
	CreatedAt   time.Time           `json:"created_at"`
	Description string              `json:"description"`
	Matched     *bool               `json:"matched,omitempty"`
	MetaData    map[string][]string `json:"meta_data,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at"`
	VersionID   string              `json:"version_id"`
	WorkspaceID string              `json:"workspace_id"`
}
