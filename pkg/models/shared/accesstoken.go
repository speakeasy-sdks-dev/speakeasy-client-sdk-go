// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type Claims struct {
}

type User struct {
	Admin         *bool      `json:"admin,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	DisplayName   *string    `json:"display_name,omitempty"`
	Email         *string    `json:"email,omitempty"`
	EmailVerified *bool      `json:"email_verified,omitempty"`
	ID            *string    `json:"id,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *User) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *User) GetAdmin() *bool {
	if o == nil {
		return nil
	}
	return o.Admin
}

func (o *User) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *User) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *User) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *User) GetEmailVerified() *bool {
	if o == nil {
		return nil
	}
	return o.EmailVerified
}

func (o *User) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type Workspaces struct {
}

// An AccessToken is a token that can be used to authenticate with the Speakeasy API.
type AccessToken struct {
	AccessToken string       `json:"access_token"`
	Claims      Claims       `json:"claims"`
	User        User         `json:"user"`
	Workspaces  []Workspaces `json:"workspaces,omitempty"`
}

func (o *AccessToken) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *AccessToken) GetClaims() Claims {
	if o == nil {
		return Claims{}
	}
	return o.Claims
}

func (o *AccessToken) GetUser() User {
	if o == nil {
		return User{}
	}
	return o.User
}

func (o *AccessToken) GetWorkspaces() []Workspaces {
	if o == nil {
		return nil
	}
	return o.Workspaces
}