package shared

import (
	"time"
)

// Schema
// A Schema represents an API schema for a particular Api and Version.
type Schema struct {
	APIID       string    `json:"api_id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	RevisionID  string    `json:"revision_id"`
	VersionID   string    `json:"version_id"`
	WorkspaceID string    `json:"workspace_id"`
}
