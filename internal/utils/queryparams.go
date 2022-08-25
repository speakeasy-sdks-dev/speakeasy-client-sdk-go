package utils

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	queryParamTagKey = "queryParam"
)

func PopulateQueryParams(ctx context.Context, req *http.Request, queryParams interface{}) {
	queryParamsStructType := reflect.TypeOf(queryParams)
	queryParamsValType := reflect.ValueOf(queryParams)

	for i := 0; i < queryParamsStructType.NumField(); i++ {
		fieldType := queryParamsStructType.Field(i)
		valType := queryParamsValType.Field(i)

		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			continue
		}

		// TODO: support other styles
		switch qpTag.Style {
		case "deepObject":
			populateDeepObjectParams(ctx, req, qpTag.ParamName, fieldType.Type, valType)
		}
	}
}

func populateDeepObjectParams(ctx context.Context, req *http.Request, parentName string, objType reflect.Type, objValue reflect.Value) {
	if objType.Kind() == reflect.Pointer {
		if objValue.IsNil() {
			return
		}
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	queryParams := req.URL.Query()

	switch objType.Kind() {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldType := objType.Field(i)
			valType := objValue.Field(i)

			qpTag := parseQueryParamTag(fieldType)
			if qpTag == nil {
				continue
			}

			queryParams.Add(fmt.Sprintf("%s[%s]", parentName, qpTag.ParamName), fmt.Sprintf("%v", valType.Interface()))
		}
	case reflect.Map:
		iter := objValue.MapRange()
		for iter.Next() {
			// TODO support other types
			switch iter.Value().Kind() {
			case reflect.Array, reflect.Slice:
				// TODO how deep do we want to go?
				for i := 0; i < iter.Value().Len(); i++ {
					queryParams.Add(fmt.Sprintf("%s[%s]", parentName, iter.Key().String()), fmt.Sprintf("%v", iter.Value().Index(i).Interface()))
				}
			default:
				queryParams.Add(fmt.Sprintf("%s[%s]", parentName, iter.Key().String()), fmt.Sprintf("%v", iter.Value().Interface()))
			}
		}
	default:
		// TODO not currently supported. Print warning?
	}

	req.URL.RawQuery = queryParams.Encode()
}

type paramTag struct {
	Style     string
	Explode   bool
	ParamName string
}

func parseQueryParamTag(field reflect.StructField) *paramTag {
	// example `queryParam:"style=deepObject,explode=true,name=op"`
	values := parseStructTag(queryParamTagKey, field)

	tag := &paramTag{
		Style:     "form",
		Explode:   true,
		ParamName: strings.ToLower(field.Name),
	}

	for k, v := range values {
		switch k {
		case "style":
			tag.Style = v
		case "explode":
			tag.Explode = v == "true"
		case "name":
			tag.ParamName = v
		}
	}

	return tag
}
