# GetSchema
Available in: `Schemas`

Returns information about the last uploaded schema for a particular API version. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    req := operations.GetSchemaRequest{
        APIID: "maiores",
        VersionID: "dicta",
    }

    res, err := s.Schemas.GetSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```