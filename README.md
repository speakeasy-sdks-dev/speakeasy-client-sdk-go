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
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models"
)

func main() {
    ctx := context.Background()

	sdk := sdk.New()
	sdk.ConfigureSecurity(models.Security{
		Scheme1: models.Scheme1{
			APIKey: "YOUR_API_KEY", // Replace with your API key from your Speakeasy Workspace
		},
	})

	getApisRes, err := sdk.GetApisV1(ctx, models.GetApisV1Request{
		QueryParams: models.GetApisV1QueryParams{
			Metadata: map[string][]string{
				"label": {"1"},
			},
			Op: &models.GetApisV1Op{
				And: true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if getApisRes.StatusCode != 200 {
		fmt.Println(getApisRes.Responses[getApisRes.StatusCode][getApisRes.ContentType].Error)
	} else {
		fmt.Println(getApisRes.Responses[getApisRes.StatusCode][getApisRes.ContentType].API)
	}
}
```