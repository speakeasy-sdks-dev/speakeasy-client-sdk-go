# APIEndpoints

## Overview

REST APIs for managing ApiEndpoint entities

### Available Operations

* [DeleteAPIEndpoint](#deleteapiendpoint) - Delete an ApiEndpoint.
* [FindAPIEndpoint](#findapiendpoint) - Find an ApiEndpoint via its displayName.
* [GenerateOpenAPISpecForAPIEndpoint](#generateopenapispecforapiendpoint) - Generate an OpenAPI specification for a particular ApiEndpoint.
* [GeneratePostmanCollectionForAPIEndpoint](#generatepostmancollectionforapiendpoint) - Generate a Postman collection for a particular ApiEndpoint.
* [GetAllAPIEndpoints](#getallapiendpoints) - Get all Api endpoints for a particular apiID.
* [GetAllForVersionAPIEndpoints](#getallforversionapiendpoints) - Get all ApiEndpoints for a particular apiID and versionID.
* [GetAPIEndpoint](#getapiendpoint) - Get an ApiEndpoint.
* [UpsertAPIEndpoint](#upsertapiendpoint) - Upsert an ApiEndpoint.

## DeleteAPIEndpoint

Delete an ApiEndpoint. This will also delete all associated Request Logs (if using a Postgres datastore).

### Example Usage

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
    req := operations.DeleteAPIEndpointRequest{
        APIEndpointID: "delectus",
        APIID: "tempora",
        VersionID: "suscipit",
    }

    res, err := s.APIEndpoints.DeleteAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## FindAPIEndpoint

Find an ApiEndpoint via its displayName (set by operationId from a registered OpenAPI schema).
This is useful for finding the ID of an ApiEndpoint to use in the /v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID} endpoints.

### Example Usage

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
    req := operations.FindAPIEndpointRequest{
        APIID: "molestiae",
        DisplayName: "minus",
        VersionID: "placeat",
    }

    res, err := s.APIEndpoints.FindAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```

## GenerateOpenAPISpecForAPIEndpoint

This endpoint will generate a new operation in any registered OpenAPI document if the operation does not already exist in the document.
Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.

### Example Usage

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
    req := operations.GenerateOpenAPISpecForAPIEndpointRequest{
        APIEndpointID: "voluptatum",
        APIID: "iusto",
        VersionID: "excepturi",
    }

    res, err := s.APIEndpoints.GenerateOpenAPISpecForAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateOpenAPISpecDiff != nil {
        // handle response
    }
}
```

## GeneratePostmanCollectionForAPIEndpoint

Generates a postman collection that allows the endpoint to be called from postman variables produced for any path/query/header parameters included in the OpenAPI document.

### Example Usage

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
    req := operations.GeneratePostmanCollectionForAPIEndpointRequest{
        APIEndpointID: "nisi",
        APIID: "recusandae",
        VersionID: "temporibus",
    }

    res, err := s.APIEndpoints.GeneratePostmanCollectionForAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostmanCollection != nil {
        // handle response
    }
}
```

## GetAllAPIEndpoints

Get all Api endpoints for a particular apiID.

### Example Usage

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
    req := operations.GetAllAPIEndpointsRequest{
        APIID: "ab",
    }

    res, err := s.APIEndpoints.GetAllAPIEndpoints(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoints != nil {
        // handle response
    }
}
```

## GetAllForVersionAPIEndpoints

Get all ApiEndpoints for a particular apiID and versionID.

### Example Usage

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
    req := operations.GetAllForVersionAPIEndpointsRequest{
        APIID: "quis",
        VersionID: "veritatis",
    }

    res, err := s.APIEndpoints.GetAllForVersionAPIEndpoints(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoints != nil {
        // handle response
    }
}
```

## GetAPIEndpoint

Get an ApiEndpoint.

### Example Usage

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
    req := operations.GetAPIEndpointRequest{
        APIEndpointID: "deserunt",
        APIID: "perferendis",
        VersionID: "ipsam",
    }

    res, err := s.APIEndpoints.GetAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```

## UpsertAPIEndpoint

Upsert an ApiEndpoint. If the ApiEndpoint does not exist it will be created, otherwise it will be updated.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    s := speakeasy.New(
        speakeasy.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()    
    req := operations.UpsertAPIEndpointRequest{
        APIEndpointInput: shared.APIEndpointInput{
            APIEndpointID: "repellendus",
            Description: "sapiente",
            DisplayName: "quo",
            Method: "odit",
            Path: "at",
            VersionID: "at",
        },
        APIEndpointID: "maiores",
        APIID: "molestiae",
        VersionID: "quod",
    }

    res, err := s.APIEndpoints.UpsertAPIEndpoint(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```
