package shared

type SchemaDiffModificationsValueChange struct {
	From string `json:"From"`
	To   string `json:"To"`
}

type SchemaDiff struct {
	Additions     []string                                      `json:"additions"`
	Deletions     []string                                      `json:"deletions"`
	Modifications map[string]SchemaDiffModificationsValueChange `json:"modifications"`
}
