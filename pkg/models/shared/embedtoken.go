package shared

import (
	"time"
)

// EmbedToken
// A representation of an embed token granted for working with Speakeasy components.
type EmbedToken struct {
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	Description string     `json:"description"`
	ExpiresAt   time.Time  `json:"expires_at"`
	Filters     string     `json:"filters"`
	ID          string     `json:"id"`
	LastUsed    *time.Time `json:"last_used,omitempty"`
	RevokedAt   *time.Time `json:"revoked_at,omitempty"`
	RevokedBy   *string    `json:"revoked_by,omitempty"`
	WorkspaceID string     `json:"workspace_id"`
}
