# GetSchemas
Available in: `Schemas`

Returns information the schemas associated with a particular apiID. 
This won't include the schemas themselves, they can be retrieved via the downloadSchema operation.

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
    req := operations.GetSchemasRequest{
        APIID: "commodi",
        VersionID: "repudiandae",
    }

    res, err := s.Schemas.GetSchemas(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schemata != nil {
        // handle response
    }
}
```