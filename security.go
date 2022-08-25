package sdk

import (
	"net/http"
	"reflect"
	"strings"
)

const (
	securityTagKey = "security"
)

type securityTag struct {
	Scheme bool
	Name   string
	Type   string
	In     string
}

type SecurityClient struct {
	client  *http.Client
	headers map[string]string
}

var _ HTTPClient = &SecurityClient{}

func newSecurityClient() *SecurityClient {
	return &SecurityClient{
		client:  http.DefaultClient,
		headers: make(map[string]string),
	}
}

func (c *SecurityClient) Do(req *http.Request) (*http.Response, error) {
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	return c.client.Do(req)
}

func createSecurityClient(security interface{}) *SecurityClient {
	client := parseSecurityStruct(security)
	if client != nil {
		return client
	}

	return newSecurityClient()
}

func parseSecurityStruct(security interface{}) *SecurityClient {
	securityStructType := reflect.TypeOf(security)
	securityValType := reflect.ValueOf(security)

	for i := 0; i < securityStructType.NumField(); i++ {
		fieldType := securityStructType.Field(i)
		valType := securityValType.Field(i)

		if fieldType.Type.Kind() == reflect.Pointer && valType.Elem().IsNil() {
			continue
		}

		secTag := parseSecurityTag(fieldType)
		if secTag != nil && secTag.Scheme {
			return parseSecurityScheme(valType.Interface())
		}
	}

	return nil
}

func parseSecurityScheme(scheme interface{}) *SecurityClient {
	schemeStructType := reflect.TypeOf(scheme)
	schemeValType := reflect.ValueOf(scheme)

	client := newSecurityClient()

	for i := 0; i < schemeStructType.NumField(); i++ {
		fieldType := schemeStructType.Field(i)
		valType := schemeValType.Field(i)

		secTag := parseSecurityTag(fieldType)
		if secTag == nil || secTag.Scheme {
			continue
		}

		// TODO: deal with other types
		if secTag.Type == "apiKey" {
			// ToDo: deal with other in values
			switch secTag.In {
			case "header":
				client.headers[secTag.Name] = valType.String()
			}
		}
	}

	return client
}

func parseSecurityTag(field reflect.StructField) *securityTag {
	tag := field.Tag.Get(securityTagKey)
	if tag == "" {
		return nil
	}

	scheme := false
	name := ""
	securityType := ""
	securityIn := ""

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")
		if len(parts) < 1 || len(parts) > 2 {
			continue
		}

		if len(parts) == 1 && parts[0] == "scheme" {
			scheme = true
			continue
		}

		switch parts[0] {
		case "name":
			name = parts[1]
		case "type":
			securityType = parts[1]
		case "in":
			securityIn = parts[1]
		}
	}

	if !scheme && (name == "" || securityType == "" || securityIn == "") {
		return nil
	}

	return &securityTag{
		Scheme: scheme,
		Name:   name,
		Type:   securityType,
		In:     securityIn,
	}
}
