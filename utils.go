package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func parseStructTag(tagKey string, field reflect.StructField) map[string]string {
	tag := field.Tag.Get(tagKey)
	if tag == "" {
		return nil
	}

	values := map[string]string{}

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")

		switch len(parts) {
		case 1:
			// flag option
			parts = append(parts, "true")
		case 2:
			// key=value option
			break
		default:
			// invalid option
			continue
		}

		values[parts[0]] = parts[1]
	}

	return values
}

func unmarshalJsonFromResponseBody(body io.Reader, out interface{}) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("error unmarshalling json response body: %w", err)
	}

	return nil
}
