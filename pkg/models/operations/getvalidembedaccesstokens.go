package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetValidEmbedAccessTokensResponse struct {
	ContentType string
	EmbedTokens []shared.EmbedToken
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
