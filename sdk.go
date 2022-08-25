package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/models"
	"io"
	"mime"
	"net/http"
	"strings"
)

const (
	ServerURL = "http://api.prod.speakeasyapi.dev"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SDK struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
}

func New() *SDK {
	return &SDK{
		defaultClient:  http.DefaultClient,
		securityClient: http.DefaultClient,
		serverURL:      ServerURL,
	}
}

func (s *SDK) ConfigureSecurity(security models.Security) {
	s.securityClient = createSecurityClient(security)
}

func (s *SDK) GetApisV1(ctx context.Context, request models.GetApisV1Request) (*models.GetApisV1Response, error) {
	url := strings.TrimSuffix(s.serverURL, "/") + "/v1/apis"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	populateQueryParams(ctx, req, request.QueryParams)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetApisV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetApisV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetApisV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.API
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetApisV1Responses{
				API: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetApisV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetApisV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetApisV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIVersionsV1(ctx context.Context, request models.GetAllAPIVersionsV1Request) (*models.GetAllAPIVersionsV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	populateQueryParams(ctx, req, request.QueryParams)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetAllAPIVersionsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetAllAPIVersionsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetAllAPIVersionsV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.API
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllAPIVersionsV1Responses{
				API: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllAPIVersionsV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllAPIVersionsV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllAPIVersionsV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPIV1(ctx context.Context, request models.UpsertAPIV1Request) (*models.UpsertAPIV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}", request.PathParams)

	bodyReader, reqContentType, err := serializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", reqContentType)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.UpsertAPIV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.UpsertAPIV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.UpsertAPIV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.API
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.UpsertAPIV1Responses{
				API: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.UpsertAPIV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.UpsertAPIV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.UpsertAPIV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIEndpointsV1(ctx context.Context, request models.GetAllAPIEndpointsV1Request) (*models.GetAllAPIEndpointsV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/api_endpoints", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetAllAPIEndpointsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetAllAPIEndpointsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetAllAPIEndpointsV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.APIEndpoint
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllAPIEndpointsV1Responses{
				APIEndpoint: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllAPIEndpointsV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllAPIEndpointsV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllAPIEndpointsV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteAPIV1(ctx context.Context, request models.DeleteAPIV1Request) (*models.DeleteAPIV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DeleteAPIV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DeleteAPIV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DeleteAPIV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DeleteAPIV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DeleteAPIV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllForVersionAPIEndpointsV1(ctx context.Context, request models.GetAllForVersionAPIEndpointsV1Request) (*models.GetAllForVersionAPIEndpointsV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetAllForVersionAPIEndpointsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetAllForVersionAPIEndpointsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetAllForVersionAPIEndpointsV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.APIEndpoint
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllForVersionAPIEndpointsV1Responses{
				APIEndpoint: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllForVersionAPIEndpointsV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAllForVersionAPIEndpointsV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAllForVersionAPIEndpointsV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) FindAPIEndpointV1(ctx context.Context, request models.FindAPIEndpointV1Request) (*models.FindAPIEndpointV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/find/{displayName}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.FindAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.FindAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.FindAPIEndpointV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.APIEndpoint
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.FindAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.FindAPIEndpointV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.FindAPIEndpointV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.FindAPIEndpointV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteAPIEndpointV1(ctx context.Context, request models.DeleteAPIEndpointV1Request) (*models.DeleteAPIEndpointV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DeleteAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DeleteAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DeleteAPIEndpointV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DeleteAPIEndpointV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DeleteAPIEndpointV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAPIEndpointV1(ctx context.Context, request models.GetAPIEndpointV1Request) (*models.GetAPIEndpointV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetAPIEndpointV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.APIEndpoint
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAPIEndpointV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetAPIEndpointV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetAPIEndpointV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPIEndpointV1(ctx context.Context, request models.UpsertAPIEndpointV1Request) (*models.UpsertAPIEndpointV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	bodyReader, reqContentType, err := serializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", reqContentType)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.UpsertAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.UpsertAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.UpsertAPIEndpointV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.APIEndpoint
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.UpsertAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.UpsertAPIEndpointV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.UpsertAPIEndpointV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.UpsertAPIEndpointV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaV1(ctx context.Context, request models.GetSchemaV1Request) (*models.GetSchemaV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetSchemaV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.Schema
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaV1Responses{
				Schema: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) RegisterSchemaV1(ctx context.Context, request models.RegisterSchemaV1Request) (*models.RegisterSchemaV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema", request.PathParams)

	bodyReader, reqContentType, err := serializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", reqContentType)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.RegisterSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.RegisterSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.RegisterSchemaV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.RegisterSchemaV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.RegisterSchemaV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchemaV1(ctx context.Context, request models.DownloadSchemaV1Request) (*models.DownloadSchemaV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema/download", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DownloadSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DownloadSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DownloadSchemaV1Responses)
	}

	switch httpRes.StatusCode {
	case 302:
		res.Headers = httpRes.Header

		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DownloadSchemaV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DownloadSchemaV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaDiffV1(ctx context.Context, request models.GetSchemaDiffV1Request) (*models.GetSchemaDiffV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema/{baseRevisionID}/diff/{targetRevisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetSchemaDiffV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetSchemaDiffV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetSchemaDiffV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.SchemaDiff
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaDiffV1Responses{
				SchemaDiff: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaDiffV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaDiffV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaDiffV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteSchemaV1(ctx context.Context, request models.DeleteSchemaV1Request) (*models.DeleteSchemaV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DeleteSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DeleteSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DeleteSchemaV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DeleteSchemaV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DeleteSchemaV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaRevisionV1(ctx context.Context, request models.GetSchemaRevisionV1Request) (*models.GetSchemaRevisionV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetSchemaRevisionV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetSchemaRevisionV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetSchemaRevisionV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.Schema
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaRevisionV1Responses{
				Schema: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaRevisionV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemaRevisionV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemaRevisionV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchemaRevisionV1(ctx context.Context, request models.DownloadSchemaRevisionV1Request) (*models.DownloadSchemaRevisionV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}/download", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DownloadSchemaRevisionV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DownloadSchemaRevisionV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DownloadSchemaRevisionV1Responses)
	}

	switch httpRes.StatusCode {
	case 302:
		res.Headers = httpRes.Header

		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DownloadSchemaRevisionV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DownloadSchemaRevisionV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemasV1(ctx context.Context, request models.GetSchemasV1Request) (*models.GetSchemasV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/version/{versionID}/schemas", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetSchemasV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetSchemasV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetSchemasV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.Schema
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemasV1Responses{
				Schema: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemasV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetSchemasV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetSchemasV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetVersionMetadataV1(ctx context.Context, request models.GetVersionMetadataV1Request) (*models.GetVersionMetadataV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/versions/{versionID}/metadata", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.GetVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.GetVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.GetVersionMetadataV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out []models.VersionMetadata
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetVersionMetadataV1Responses{
				VersionMetadata: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetVersionMetadataV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.GetVersionMetadataV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.GetVersionMetadataV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) InsertVersionMetadataV1(ctx context.Context, request models.InsertVersionMetadataV1Request) (*models.InsertVersionMetadataV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/versions/{versionID}/metadata", request.PathParams)

	bodyReader, reqContentType, err := serializeRequestBody(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", reqContentType)

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.InsertVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.InsertVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.InsertVersionMetadataV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		switch contentType {
		case "application/json":
			var out *models.VersionMetadata
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.InsertVersionMetadataV1Responses{
				VersionMetadata: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.InsertVersionMetadataV1Responses{
				RawResponse: data,
			}
		}

	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.InsertVersionMetadataV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.InsertVersionMetadataV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteVersionMetadataV1(ctx context.Context, request models.DeleteVersionMetadataV1Request) (*models.DeleteVersionMetadataV1Response, error) {
	url := generateURL(ctx, s.serverURL, "/v1/apis/{apiID}/versions/{versionID}/metadata/{metaKey}/{metaValue}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	httpRes, err := s.securityClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentTypeHeader := httpRes.Header.Get("Content-Type")
	contentType, _, err := mime.ParseMediaType(contentTypeHeader)
	if err != nil {
		contentType = "unknown"
	}

	res := &models.DeleteVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]models.DeleteVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]models.DeleteVersionMetadataV1Responses)
	}

	switch httpRes.StatusCode {
	case 200:
		break
	default:
		switch contentType {
		case "application/json":
			var out *models.Error
			if err := unmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)]["application/json"] = models.DeleteVersionMetadataV1Responses{
				Error: out,
			}
		default:
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = models.DeleteVersionMetadataV1Responses{
				RawResponse: data,
			}
		}
	}

	return res, nil
}
