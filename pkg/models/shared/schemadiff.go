package shared

// SchemaDiff
// A SchemaDiff represents a diff of two Schemas.
type SchemaDiff struct {
	Additions     []string                         `json:"additions"`
	Deletions     []string                         `json:"deletions"`
	Modifications map[string]SchemaDiffValueChange `json:"modifications"`
}

type SchemaDiffValueChange struct {
	From string `json:"From"`
	To   string `json:"To"`
}
