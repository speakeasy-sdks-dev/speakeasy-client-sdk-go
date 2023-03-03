package shared

type Security struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=x-api-key"`
}
