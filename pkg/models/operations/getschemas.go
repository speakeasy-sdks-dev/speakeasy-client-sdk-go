package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetSchemasPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemasRequest struct {
	PathParams GetSchemasPathParams
}

type GetSchemasResponse struct {
	ContentType string
	Error       *shared.Error
	Schemata    []shared.Schema
	StatusCode  int
	RawResponse *http.Response
}
