# VersionMetadata

A set of keys and associated values, attached to a particular version of an Api.


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `APIID`                                             | *string*                                            | :heavy_check_mark:                                  | The ID of the Api this Metadata belongs to.         |
| `CreatedAt`                                         | [time.Time](https://pkg.go.dev/time#Time)           | :heavy_check_mark:                                  | Creation timestamp.                                 |
| `MetaKey`                                           | *string*                                            | :heavy_check_mark:                                  | The key for this metadata.                          |
| `MetaValue`                                         | *string*                                            | :heavy_check_mark:                                  | One of the values for this metadata.                |
| `VersionID`                                         | *string*                                            | :heavy_check_mark:                                  | The version ID of the Api this Metadata belongs to. |
| `WorkspaceID`                                       | *string*                                            | :heavy_check_mark:                                  | The workspace ID this Metadata belongs to.          |