# GetSchemaDiff
Available in: `Schemas`

Get a diff of two schema revisions for an Api.

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
    req := operations.GetSchemaDiffRequest{
        APIID: "corporis",
        BaseRevisionID: "dolore",
        TargetRevisionID: "iusto",
        VersionID: "dicta",
    }

    res, err := s.Schemas.GetSchemaDiff(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaDiff != nil {
        // handle response
    }
}
```