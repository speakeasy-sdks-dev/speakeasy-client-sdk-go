// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetPluginsResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	Plugins     []shared.Plugin
	StatusCode  int
	RawResponse *http.Response
}
