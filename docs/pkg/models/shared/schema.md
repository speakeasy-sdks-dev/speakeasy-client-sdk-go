# Schema

A Schema represents an API schema for a particular Api and Version.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `APIID`                                                   | *string*                                                  | :heavy_check_mark:                                        | The ID of the Api this Schema belongs to.                 |
| `CreatedAt`                                               | [time.Time](https://pkg.go.dev/time#Time)                 | :heavy_check_mark:                                        | Creation timestamp.                                       |
| `Description`                                             | *string*                                                  | :heavy_check_mark:                                        | A detailed description of the Schema.                     |
| `RevisionID`                                              | *string*                                                  | :heavy_check_mark:                                        | An ID referencing this particular revision of the Schema. |
| `VersionID`                                               | *string*                                                  | :heavy_check_mark:                                        | The version ID of the Api this Schema belongs to.         |
| `WorkspaceID`                                             | *string*                                                  | :heavy_check_mark:                                        | The workspace ID this Schema belongs to.                  |