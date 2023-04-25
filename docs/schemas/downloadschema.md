# DownloadSchema
Available in: `Schemas`

Download the latest schema for a particular apiID.

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
    req := operations.DownloadSchemaRequest{
        APIID: "cum",
        VersionID: "perferendis",
    }

    res, err := s.Schemas.DownloadSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```