<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-api/speakeasy-client-sdk-go"
    "github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
    "github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasy.New(
        WithSecurity(        shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    req := operations.GetApisRequest{
        Metadata: map[string][]string{
            "deserunt": []string{
                "nulla",
                "id",
                "vero",
            },
            "perspiciatis": []string{
                "nihil",
                "fuga",
                "facilis",
                "eum",
            },
            "iusto": []string{
                "saepe",
                "inventore",
            },
        },
        Op: &operations.GetApisOp{
            And: false,
        },
    }

    ctx := context.Background()
    res, err := s.Apis.GetApis(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->