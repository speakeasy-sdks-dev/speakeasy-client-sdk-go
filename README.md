# github.com/speakeasy-api/speakeasy-client-sdk-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-api/speakeasy-client-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.GetApisRequest{
        Metadata: map[string][]string{
            "provident": []string{
                "quibusdam",
                "unde",
                "nulla",
            },
            "corrupti": []string{
                "vel",
                "error",
                "deserunt",
                "suscipit",
            },
            "iure": []string{
                "debitis",
                "ipsa",
            },
        },
        Op: &operations.GetApisOp{
            And: false,
        },
    }

    res, err := s.Apis.GetApis(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations

### [Speakeasy SDK](docs/speakeasy/README.md)

* [ValidateAPIKey](docs/speakeasy/validateapikey.md) - Validate the current api key.

### [APIEndpoints](docs/apiendpoints/README.md)

* [DeleteAPIEndpoint](docs/apiendpoints/deleteapiendpoint.md) - Delete an ApiEndpoint.
* [FindAPIEndpoint](docs/apiendpoints/findapiendpoint.md) - Find an ApiEndpoint via its displayName.
* [GenerateOpenAPISpecForAPIEndpoint](docs/apiendpoints/generateopenapispecforapiendpoint.md) - Generate an OpenAPI specification for a particular ApiEndpoint.
* [GeneratePostmanCollectionForAPIEndpoint](docs/apiendpoints/generatepostmancollectionforapiendpoint.md) - Generate a Postman collection for a particular ApiEndpoint.
* [GetAllAPIEndpoints](docs/apiendpoints/getallapiendpoints.md) - Get all Api endpoints for a particular apiID.
* [GetAllForVersionAPIEndpoints](docs/apiendpoints/getallforversionapiendpoints.md) - Get all ApiEndpoints for a particular apiID and versionID.
* [GetAPIEndpoint](docs/apiendpoints/getapiendpoint.md) - Get an ApiEndpoint.
* [UpsertAPIEndpoint](docs/apiendpoints/upsertapiendpoint.md) - Upsert an ApiEndpoint.

### [Apis](docs/apis/README.md)

* [DeleteAPI](docs/apis/deleteapi.md) - Delete an Api.
* [GenerateOpenAPISpec](docs/apis/generateopenapispec.md) - Generate an OpenAPI specification for a particular Api.
* [GeneratePostmanCollection](docs/apis/generatepostmancollection.md) - Generate a Postman collection for a particular Api.
* [GetAllAPIVersions](docs/apis/getallapiversions.md) - Get all Api versions for a particular ApiEndpoint.
* [GetApis](docs/apis/getapis.md) - Get a list of Apis for a given workspace
* [UpsertAPI](docs/apis/upsertapi.md) - Upsert an Api

### [Embeds](docs/embeds/README.md)

* [GetEmbedAccessToken](docs/embeds/getembedaccesstoken.md) - Get an embed access token for the current workspace.
* [GetValidEmbedAccessTokens](docs/embeds/getvalidembedaccesstokens.md) - Get all valid embed access tokens for the current workspace.
* [RevokeEmbedAccessToken](docs/embeds/revokeembedaccesstoken.md) - Revoke an embed access EmbedToken.

### [Metadata](docs/metadata/README.md)

* [DeleteVersionMetadata](docs/metadata/deleteversionmetadata.md) - Delete metadata for a particular apiID and versionID.
* [GetVersionMetadata](docs/metadata/getversionmetadata.md) - Get all metadata for a particular apiID and versionID.
* [InsertVersionMetadata](docs/metadata/insertversionmetadata.md) - Insert metadata for a particular apiID and versionID.

### [Plugins](docs/plugins/README.md)

* [GetPlugins](docs/plugins/getplugins.md) - Get all plugins for the current workspace.
* [RunPlugin](docs/plugins/runplugin.md) - Run a plugin
* [UpsertPlugin](docs/plugins/upsertplugin.md) - Upsert a plugin

### [Requests](docs/requests/README.md)

* [GenerateRequestPostmanCollection](docs/requests/generaterequestpostmancollection.md) - Generate a Postman collection for a particular request.
* [GetRequestFromEventLog](docs/requests/getrequestfromeventlog.md) - Get information about a particular request.
* [QueryEventLog](docs/requests/queryeventlog.md) - Query the event log to retrieve a list of requests.

### [Schemas](docs/schemas/README.md)

* [DeleteSchema](docs/schemas/deleteschema.md) - Delete a particular schema revision for an Api.
* [DownloadSchema](docs/schemas/downloadschema.md) - Download the latest schema for a particular apiID.
* [DownloadSchemaRevision](docs/schemas/downloadschemarevision.md) - Download a particular schema revision for an Api.
* [GetSchema](docs/schemas/getschema.md) - Get information about the latest schema.
* [GetSchemaDiff](docs/schemas/getschemadiff.md) - Get a diff of two schema revisions for an Api.
* [GetSchemaRevision](docs/schemas/getschemarevision.md) - Get information about a particular schema revision for an Api.
* [GetSchemas](docs/schemas/getschemas.md) - Get information about all schemas associated with a particular apiID.
* [RegisterSchema](docs/schemas/registerschema.md) - Register a schema.
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
