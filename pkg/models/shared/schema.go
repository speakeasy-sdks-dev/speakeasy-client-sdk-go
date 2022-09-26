package shared

import (
	"time"
)

type Schema struct {
	APIID       string    `json:"api_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	RevisionID  string    `json:"revision_id"`
	VersionID   string    `json:"version_id"`
	WorkspaceID string    `json:"workspace_id"`
}
