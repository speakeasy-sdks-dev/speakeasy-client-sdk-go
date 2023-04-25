# UpsertAPIEndpoint
Available in: `APIEndpoints`

Upsert an ApiEndpoint. If the ApiEndpoint does not exist it will be created, otherwise it will be updated.

## Example Usage
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