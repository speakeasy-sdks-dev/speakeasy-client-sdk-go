package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"reflect"
)

const (
	requestTagKey       = "request"
	requestFieldName    = "Request"
	multipartFormTagKey = "multipartForm"
)

func SerializeRequestBody(ctx context.Context, request interface{}) (*bytes.Buffer, string, error) {
	requestStructType := reflect.TypeOf(request)
	requestValType := reflect.ValueOf(request)

	if requestStructType.Kind() == reflect.Pointer {
		if requestValType.IsNil() {
			return nil, "", nil
		}

		requestStructType = requestStructType.Elem()
		requestValType = requestValType.Elem()
	}

	requestField, ok := requestStructType.FieldByName(requestFieldName)
	if !ok {
		return nil, "", fmt.Errorf("request body not found")
	}

	tag := getRequestTag(requestField)
	if tag != nil {
		// Single Request Object at First Level
		requestVal := requestValType.FieldByName(requestFieldName)
		if requestField.Type.Kind() == reflect.Pointer && requestVal.IsNil() {
			return nil, "", nil
		}

		return serializeContentType(tag.MediaType, requestVal)
	}

	// Multi Request Object at First Level
	requestStructType = requestField.Type
	requestValType = requestValType.FieldByName(requestFieldName)

	for i := 0; i < requestStructType.NumField(); i++ {
		fieldType := requestStructType.Field(i)
		valType := requestValType.Field(i)

		if fieldType.Type.Kind() == reflect.Pointer {
			if valType.IsNil() {
				continue
			}
		}

		tag := getRequestTag(fieldType)
		if tag == nil {
			return nil, "", fmt.Errorf("missing request tag on request body field %s", fieldType.Name)
		}

		return serializeContentType(tag.MediaType, valType)
	}

	return nil, "", nil
}

func serializeContentType(mediaType string, val reflect.Value) (*bytes.Buffer, string, error) {
	buf := &bytes.Buffer{}

	switch mediaType {
	case "application/json":
		if err := json.NewEncoder(buf).Encode(val.Interface()); err != nil {
			return nil, "", err
		}
	case "multipart/form-data", "multipart/mixed":
		var err error
		mediaType, err = encodeMultipartFormData(buf, val.Interface())
		if err != nil {
			return nil, "", err
		}
	default:
		if val.Kind() != reflect.Slice && val.Kind() != reflect.Array && val.Type() != reflect.TypeOf([]byte(nil)) {
			return nil, "", fmt.Errorf("invalid request body type %s for mediaType %s", val.Type(), mediaType)
		}

		if _, err := buf.Write(val.Bytes()); err != nil {
			return nil, "", err
		}
	}

	return buf, mediaType, nil
}

func encodeMultipartFormData(w io.Writer, data interface{}) (string, error) {
	requestStructType := reflect.TypeOf(data)
	requestValType := reflect.ValueOf(data)

	if requestStructType.Kind() == reflect.Pointer {
		requestStructType = requestStructType.Elem()
		requestValType = requestValType.Elem()
	}

	writer := multipart.NewWriter(w)

	for i := 0; i < requestStructType.NumField(); i++ {
		field := requestStructType.Field(i)
		fieldType := field.Type
		valType := requestValType.Field(i)

		if fieldType.Kind() == reflect.Pointer {
			if valType.IsNil() {
				continue
			}

			fieldType = fieldType.Elem()
			valType = valType.Elem()
		}

		tag := parseMultipartFormTag(field)
		if tag.File {
			if err := encodeMultipartFormDataFile(writer, fieldType, valType); err != nil {
				writer.Close()
				return "", err
			}
		} else if tag.JSON {
			jw, err := writer.CreateFormField(tag.Name)
			if err != nil {
				writer.Close()
				return "", err
			}
			if err := json.NewEncoder(jw).Encode(valType.Interface()); err != nil {
				writer.Close()
				return "", err
			}
		} else {
			switch fieldType.Kind() {
			case reflect.Slice, reflect.Array:
				values := parseFormStyleArray(true, valType)
				for _, v := range values {
					if err := writer.WriteField(tag.Name+"[]", v); err != nil {
						writer.Close()
						return "", err
					}
				}
			default:
				if err := writer.WriteField(tag.Name, fmt.Sprintf("%v", valType.Interface())); err != nil {
					writer.Close()
					return "", err
				}
			}
		}
	}

	if err := writer.Close(); err != nil {
		return "", err
	}

	return writer.FormDataContentType(), nil
}

func encodeMultipartFormDataFile(w *multipart.Writer, fieldType reflect.Type, valType reflect.Value) error {
	if fieldType.Kind() != reflect.Struct {
		return fmt.Errorf("invalid type %s for multipart/form-data file", valType.Type())
	}

	var fieldName string
	var fileName string
	var content []byte

	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)
		val := valType.Field(i)

		tag := parseMultipartFormTag(field)
		if !tag.Content && tag.Name == "" {
			continue
		}

		if tag.Content {
			content = val.Bytes()
		} else {
			fieldName = tag.Name
			fileName = val.String()
		}
	}

	if fieldName == "" || fileName == "" || content == nil {
		return fmt.Errorf("invalid multipart/form-data file")
	}

	fw, err := w.CreateFormFile(fieldName, fileName)
	if err != nil {
		return err
	}
	if _, err := fw.Write(content); err != nil {
		return err
	}

	return nil
}

type requestTag struct {
	MediaType string
}

func getRequestTag(field reflect.StructField) *requestTag {
	// example `request:"mediaType=multipart/form-data"`
	values := parseStructTag(requestTagKey, field)
	if values == nil {
		return nil
	}

	tag := &requestTag{
		MediaType: "application/octet-stream",
	}

	for k, v := range values {
		switch k {
		case "mediaType":
			tag.MediaType = v
		}
	}

	return tag
}

type multipartFormTag struct {
	File    bool
	Content bool
	JSON    bool
	Name    string
}

func parseMultipartFormTag(field reflect.StructField) *multipartFormTag {
	// example `multipartForm:"name=file"`
	values := parseStructTag(multipartFormTagKey, field)

	tag := &multipartFormTag{}

	for k, v := range values {
		switch k {
		case "file":
			tag.File = v == "true"
		case "content":
			tag.Content = v == "true"
		case "name":
			tag.Name = v
		case "json":
			tag.JSON = v == "true"
		}
	}

	return tag
}
