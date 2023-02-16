package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type ValidateAPIKeyResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
}
