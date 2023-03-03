package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
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
	StatusCode  int
	RawResponse *http.Response
}
