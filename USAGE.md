<!-- Start SDK Example Usage -->
```go
package main

import (
    "github.com/speakeasy-api/speakeasy-client-sdk-go"
    "github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
    "github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    opts := []sdk.SDKOption{
        sdk.WithSecurity(
            shared.Security{
                APIKey: shared.SchemeAPIKey{
                    APIKey: "YOUR_API_KEY_HERE",
                },
            }
        ),
    }

    s := sdk.New(opts...)
    
    req := operations.GetApisRequest{
        QueryParams: operations.GetApisQueryParams{
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
        },
    }
    
    res, err := s.Apis.GetApis(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->