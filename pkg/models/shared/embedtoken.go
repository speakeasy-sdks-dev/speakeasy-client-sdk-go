// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// EmbedToken - A representation of an embed token granted for working with Speakeasy components.
type EmbedToken struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The ID of the user that created this token.
	CreatedBy string `json:"created_by"`
	// A detailed description of the EmbedToken.
	Description string `json:"description"`
	// The time this token expires.
	ExpiresAt time.Time `json:"expires_at"`
	// The filters applied to this token.
	Filters string `json:"filters"`
	// The ID of this EmbedToken.
	ID string `json:"id"`
	// The last time this token was used.
	LastUsed *time.Time `json:"last_used,omitempty"`
	// The time this token was revoked.
	RevokedAt *time.Time `json:"revoked_at,omitempty"`
	// The ID of the user that revoked this token.
	RevokedBy *string `json:"revoked_by,omitempty"`
	// The workspace ID this token belongs to.
	WorkspaceID string `json:"workspace_id"`
}

func (o *EmbedToken) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *EmbedToken) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *EmbedToken) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *EmbedToken) GetExpiresAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ExpiresAt
}

func (o *EmbedToken) GetFilters() string {
	if o == nil {
		return ""
	}
	return o.Filters
}

func (o *EmbedToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EmbedToken) GetLastUsed() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsed
}

func (o *EmbedToken) GetRevokedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RevokedAt
}

func (o *EmbedToken) GetRevokedBy() *string {
	if o == nil {
		return nil
	}
	return o.RevokedBy
}

func (o *EmbedToken) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
