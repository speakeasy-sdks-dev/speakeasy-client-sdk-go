package shared

type SchemaDiffValueChange struct {
	From string `json:"From"`
	To   string `json:"To"`
}

type SchemaDiff struct {
	Additions     []string                         `json:"additions"`
	Deletions     []string                         `json:"deletions"`
	Modifications map[string]SchemaDiffValueChange `json:"modifications"`
}
