# UpsertAPI
Available in: `Apis`

Upsert an Api. If the Api does not exist, it will be created.
If the Api exists, it will be updated.

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
    req := operations.UpsertAPIRequest{
        APIInput: shared.APIInput{
            APIID: "dolor",
            Description: "natus",
            MetaData: map[string][]string{
                "hic": []string{
                    "fuga",
                    "in",
                    "corporis",
                    "iste",
                },
                "iure": []string{
                    "quidem",
                    "architecto",
                    "ipsa",
                    "reiciendis",
                },
            },
            VersionID: "est",
        },
        APIID: "mollitia",
    }

    res, err := s.Apis.UpsertAPI(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.API != nil {
        // handle response
    }
}
```