# speakeasy-client-sdk-go

## Example usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    ctx := context.Background()

	sdk := sdk.New()
	sdk.ConfigureSecurity(shared.Security{
		APIKey: shared.SchemeAPIKey{
			APIKey: "YOUR_API_KEY", // Replace with your API key from your Speakeasy Workspace
		},
	})

	res, err := sdk.GetApis(ctx, operations.GetApisRequest{
		QueryParams: operations.GetApisQueryParams{
			Metadata: map[string][]string{
				"label": {"1"},
			},
			Op: &operations.GetApisOp{
				And: true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Apis)
	}
}
```