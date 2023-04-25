# DownloadSchemaRevision
Available in: `Schemas`

Download a particular schema revision for an Api.

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
    req := operations.DownloadSchemaRevisionRequest{
        APIID: "doloremque",
        RevisionID: "reprehenderit",
        VersionID: "ut",
    }

    res, err := s.Schemas.DownloadSchemaRevision(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```