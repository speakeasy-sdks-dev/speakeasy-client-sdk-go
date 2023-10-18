# Plugins
(*Plugins*)

## Overview

REST APIs for managing and running plugins

### Available Operations

* [GetPlugins](#getplugins) - Get all plugins for the current workspace.
* [RunPlugin](#runplugin) - Run a plugin
* [UpsertPlugin](#upsertplugin) - Upsert a plugin

## GetPlugins

Get all plugins for the current workspace.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Plugins.GetPlugins(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugins != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetPluginsResponse](../../models/operations/getpluginsresponse.md), error**


## RunPlugin

Run a plugin

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Plugins.RunPlugin(ctx, operations.RunPluginRequest{
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "<key>",
                    Operator: "Designer",
                    Value: "Pickup",
                },
            },
            Limit: 455139,
            Offset: 173644,
            Operator: "Indiana",
        },
        PluginID: "Chrysler",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BoundedRequests != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.RunPluginRequest](../../models/operations/runpluginrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.RunPluginResponse](../../models/operations/runpluginresponse.md), error**


## UpsertPlugin

Upsert a plugin

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Plugins.UpsertPlugin(ctx, shared.Plugin{
        Code: "Fish",
        PluginID: "Rock",
        Title: "North",
        WorkspaceID: "exuding",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Plugin != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Plugin](../../models/shared/plugin.md)        | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpsertPluginResponse](../../models/operations/upsertpluginresponse.md), error**

