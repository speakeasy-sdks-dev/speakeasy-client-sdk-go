package shared

type APIKey struct {
	APIKey string `security:"name=x-api-key"`
}

type Security struct {
	APIKey APIKey `security:"scheme,type=apiKey,subtype=header"`
}
