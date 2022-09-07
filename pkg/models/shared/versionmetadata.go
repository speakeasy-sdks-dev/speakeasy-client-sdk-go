package shared

import (
	"time"
)

type VersionMetadata struct {
	APIID       string    `json:"api_id"`
	CreatedAt   time.Time `json:"created_at"`
	MetaKey     string    `json:"meta_key"`
	MetaValue   string    `json:"meta_value"`
	VersionID   string    `json:"version_id"`
	WorkspaceID string    `json:"workspace_id"`
}
