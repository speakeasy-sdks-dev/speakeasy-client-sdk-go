# InsertVersionMetadata
Available in: `Metadata`

Insert metadata for a particular apiID and versionID.

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
    req := operations.InsertVersionMetadataRequest{
        VersionMetadataInput: shared.VersionMetadataInput{
            MetaKey: "architecto",
            MetaValue: "mollitia",
        },
        APIID: "dolorem",
        VersionID: "culpa",
    }

    res, err := s.Metadata.InsertVersionMetadata(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```