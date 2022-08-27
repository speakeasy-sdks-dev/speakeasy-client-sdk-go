package models

import "time"

type API struct {
	APIID       string              `json:"api_id"`
	CreatedAt   time.Time           `json:"created_at"`
	Description string              `json:"description"`
	Matched     *bool               `json:"matched"`
	MetaData    map[string][]string `json:"meta_data"`
	UpdatedAt   time.Time           `json:"updated_at"`
	VersionID   string              `json:"version_id"`
	WorkspaceID string              `json:"workspace_id"`
}
