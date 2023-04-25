# GetVersionMetadata
Available in: `Metadata`

Get all metadata for a particular apiID and versionID.

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
    req := operations.GetVersionMetadataRequest{
        APIID: "doloribus",
        VersionID: "sapiente",
    }

    res, err := s.Metadata.GetVersionMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```