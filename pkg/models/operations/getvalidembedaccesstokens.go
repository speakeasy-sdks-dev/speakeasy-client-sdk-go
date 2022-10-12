package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetValidEmbedAccessTokensResponse struct {
	ContentType string
	EmbedTokens []shared.EmbedToken
	Error       *shared.Error
	StatusCode  int64
}
