package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type RevokeEmbedAccessTokenPathParams struct {
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

type RevokeEmbedAccessTokenRequest struct {
	PathParams RevokeEmbedAccessTokenPathParams
}

type RevokeEmbedAccessTokenResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
