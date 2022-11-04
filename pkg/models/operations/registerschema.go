package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type RegisterSchemaPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type RegisterSchemaRequest struct {
	PathParams RegisterSchemaPathParams
	Request    RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
}

type RegisterSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}

type RegisterSchemaRequestBody struct {
	File RegisterSchemaRequestBodyFile `multipartForm:"file"`
}

type RegisterSchemaRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}
