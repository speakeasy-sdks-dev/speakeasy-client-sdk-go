# DeleteVersionMetadata
Available in: `Metadata`

Delete metadata for a particular apiID and versionID.

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
    req := operations.DeleteVersionMetadataRequest{
        APIID: "excepturi",
        MetaKey: "accusantium",
        MetaValue: "iure",
        VersionID: "culpa",
    }

    res, err := s.Metadata.DeleteVersionMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```