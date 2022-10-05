package shared

import (
	"time"
)

type EmbedToken struct {
	CreatedAt   time.Time  `json:"created_at"`
	CreatedBy   string     `json:"created_by"`
	Description string     `json:"description"`
	ExpiresAt   string     `json:"expires_at"`
	Filters     string     `json:"filters"`
	ID          string     `json:"id"`
	LastUsed    *time.Time `json:"last_used"`
	RevokedAt   *time.Time `json:"revoked_at"`
	RevokedBy   *string    `json:"revoked_by"`
	WorkspaceID string     `json:"workspace_id"`
}
