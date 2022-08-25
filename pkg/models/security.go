package models

type Scheme1 struct {
	APIKey string `security:"name=x-api-key,type=apiKey,in=header"`
}

type Security struct {
	Scheme1 Scheme1 `security:"scheme"`
}
