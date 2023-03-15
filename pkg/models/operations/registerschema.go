package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type RegisterSchemaRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type RegisterSchemaRequestBody struct {
	File RegisterSchemaRequestBodyFile `multipartForm:"file"`
}

type RegisterSchemaRequest struct {
	RequestBody RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
	APIID       string                    `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID   string                    `pathParam:"style=simple,explode=false,name=versionID"`
}

type RegisterSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
