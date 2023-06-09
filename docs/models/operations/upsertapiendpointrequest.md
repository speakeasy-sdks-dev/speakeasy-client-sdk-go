# UpsertAPIEndpointRequest


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `APIEndpointInput`                                                 | [shared.APIEndpointInput](../../models/shared/apiendpointinput.md) | :heavy_check_mark:                                                 | A JSON representation of the ApiEndpoint to upsert.                |
| `APIEndpointID`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The ID of the ApiEndpoint to upsert.                               |
| `APIID`                                                            | *string*                                                           | :heavy_check_mark:                                                 | The ID of the Api the ApiEndpoint belongs to.                      |
| `VersionID`                                                        | *string*                                                           | :heavy_check_mark:                                                 | The version ID of the Api the ApiEndpoint belongs to.              |