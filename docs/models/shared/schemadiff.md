# SchemaDiff

A SchemaDiff represents a diff of two Schemas.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Additions`                                                                      | []*string*                                                                       | :heavy_check_mark:                                                               | Holds every addition change in the diff.                                         |
| `Deletions`                                                                      | []*string*                                                                       | :heavy_check_mark:                                                               | Holds every deletion change in the diff.                                         |
| `Modifications`                                                                  | map[string][SchemaDiffValueChange](../../models/shared/schemadiffvaluechange.md) | :heavy_check_mark:                                                               | Holds every modification change in the diff.                                     |