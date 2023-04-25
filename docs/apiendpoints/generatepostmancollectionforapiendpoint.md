# GeneratePostmanCollectionForAPIEndpoint
Available in: `APIEndpoints`

Generates a postman collection that allows the endpoint to be called from postman variables produced for any path/query/header parameters included in the OpenAPI document.

## Example Usage
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