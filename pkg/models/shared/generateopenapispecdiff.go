package shared

type GenerateOpenAPISpecDiff struct {
	CurrentSchema string `json:"current_schema"`
	NewSchema     string `json:"new_schema"`
}
