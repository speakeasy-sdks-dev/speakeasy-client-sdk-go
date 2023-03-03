package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetPluginsResponse struct {
	ContentType string
	Error       *shared.Error
	Plugins     []shared.Plugin
	StatusCode  int
	RawResponse *http.Response
}
