package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type UpsertPluginResponse struct {
	ContentType string
	Error       *shared.Error
	Plugin      *shared.Plugin
	StatusCode  int
	RawResponse *http.Response
}
