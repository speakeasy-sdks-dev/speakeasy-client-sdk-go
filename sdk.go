package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/internal/utils"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
	"strings"
)

var Servers = []string{
	"http://api.prod.speakeasyapi.dev",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SDK struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
}

type SDKOption func(*SDK)

func WithServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SDK) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.serverURL = serverURL
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SDK) {
		sdk.securityClient = utils.CreateSecurityClient(security)
	}
}

func New(opts ...SDKOption) *SDK {
	sdk := &SDK{
		defaultClient:  http.DefaultClient,
		securityClient: http.DefaultClient,
	}
	for _, opt := range opts {
		opt(sdk)
	}
	if sdk.serverURL == "" {
		sdk.serverURL = Servers[0]
	}

	return sdk
}

func (s *SDK) DeleteAPIEndpointV1(ctx context.Context, request operations.DeleteAPIEndpointV1Request) (*operations.DeleteAPIEndpointV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteAPIEndpointV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteAPIEndpointV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteAPIV1(ctx context.Context, request operations.DeleteAPIV1Request) (*operations.DeleteAPIV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteAPIV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteAPIV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteAPIV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteAPIV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteSchemaV1(ctx context.Context, request operations.DeleteSchemaV1Request) (*operations.DeleteSchemaV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteSchemaV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteSchemaV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteVersionMetadataV1(ctx context.Context, request operations.DeleteVersionMetadataV1Request) (*operations.DeleteVersionMetadataV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/versions/{versionID}/metadata/{metaKey}/{metaValue}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteVersionMetadataV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteVersionMetadataV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchemaRevisionV1(ctx context.Context, request operations.DownloadSchemaRevisionV1Request) (*operations.DownloadSchemaRevisionV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}/download", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DownloadSchemaRevisionV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DownloadSchemaRevisionV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DownloadSchemaRevisionV1Responses)
	}

	switch {
	case httpRes.StatusCode == 302:
		res.Headers = httpRes.Header

	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaRevisionV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchemaV1(ctx context.Context, request operations.DownloadSchemaV1Request) (*operations.DownloadSchemaV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema/download", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DownloadSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DownloadSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DownloadSchemaV1Responses)
	}

	switch {
	case httpRes.StatusCode == 302:
		res.Headers = httpRes.Header

	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) FindAPIEndpointV1(ctx context.Context, request operations.FindAPIEndpointV1Request) (*operations.FindAPIEndpointV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/find/{displayName}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.FindAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.FindAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.FindAPIEndpointV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.FindAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.FindAPIEndpointV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIEndpointsV1(ctx context.Context, request operations.GetAllAPIEndpointsV1Request) (*operations.GetAllAPIEndpointsV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/api_endpoints", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllAPIEndpointsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllAPIEndpointsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllAPIEndpointsV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIEndpointsV1Responses{
				APIEndpoints: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIEndpointsV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIVersionsV1(ctx context.Context, request operations.GetAllAPIVersionsV1Request) (*operations.GetAllAPIVersionsV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllAPIVersionsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllAPIVersionsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllAPIVersionsV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIVersionsV1Responses{
				Apis: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIVersionsV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllForVersionAPIEndpointsV1(ctx context.Context, request operations.GetAllForVersionAPIEndpointsV1Request) (*operations.GetAllForVersionAPIEndpointsV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAllForVersionAPIEndpointsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllForVersionAPIEndpointsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllForVersionAPIEndpointsV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllForVersionAPIEndpointsV1Responses{
				APIEndpoints: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllForVersionAPIEndpointsV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAPIEndpointV1(ctx context.Context, request operations.GetAPIEndpointV1Request) (*operations.GetAPIEndpointV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAPIEndpointV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAPIEndpointV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetApisV1(ctx context.Context, request operations.GetApisV1Request) (*operations.GetApisV1Response, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/apis"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetApisV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetApisV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetApisV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetApisV1Responses{
				Apis: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetApisV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaDiffV1(ctx context.Context, request operations.GetSchemaDiffV1Request) (*operations.GetSchemaDiffV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema/{baseRevisionID}/diff/{targetRevisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemaDiffV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaDiffV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaDiffV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.SchemaDiff
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaDiffV1Responses{
				SchemaDiff: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaDiffV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaRevisionV1(ctx context.Context, request operations.GetSchemaRevisionV1Request) (*operations.GetSchemaRevisionV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema/{revisionID}", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemaRevisionV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaRevisionV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaRevisionV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaRevisionV1Responses{
				Schema: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaRevisionV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaV1(ctx context.Context, request operations.GetSchemaV1Request) (*operations.GetSchemaV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaV1Responses{
				Schema: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemasV1(ctx context.Context, request operations.GetSchemasV1Request) (*operations.GetSchemasV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schemas", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSchemasV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemasV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemasV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemasV1Responses{
				Schemata: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemasV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetUsageMetricsV1(ctx context.Context, request operations.GetUsageMetricsV1Request) (*operations.GetUsageMetricsV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/workspace/{workspaceID}/metrics", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	utils.PopulateQueryParams(ctx, req, request.QueryParams)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetUsageMetricsV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetUsageMetricsV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetUsageMetricsV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.UsageMetric
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetUsageMetricsV1Responses{
				UsageMetrics: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetUsageMetricsV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetVersionMetadataV1(ctx context.Context, request operations.GetVersionMetadataV1Request) (*operations.GetVersionMetadataV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/versions/{versionID}/metadata", request.PathParams)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetVersionMetadataV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out []shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetVersionMetadataV1Responses{
				VersionMetadata: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetVersionMetadataV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) InsertVersionMetadataV1(ctx context.Context, request operations.InsertVersionMetadataV1Request) (*operations.InsertVersionMetadataV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/versions/{versionID}/metadata", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
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

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.InsertVersionMetadataV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.InsertVersionMetadataV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.InsertVersionMetadataV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.InsertVersionMetadataV1Responses{
				VersionMetadata: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.InsertVersionMetadataV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) RegisterSchemaV1(ctx context.Context, request operations.RegisterSchemaV1Request) (*operations.RegisterSchemaV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/schema", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
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

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.RegisterSchemaV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.RegisterSchemaV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.RegisterSchemaV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.RegisterSchemaV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPIEndpointV1(ctx context.Context, request operations.UpsertAPIEndpointV1Request) (*operations.UpsertAPIEndpointV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
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

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpsertAPIEndpointV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.UpsertAPIEndpointV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.UpsertAPIEndpointV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIEndpointV1Responses{
				APIEndpoint: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIEndpointV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPIV1(ctx context.Context, request operations.UpsertAPIV1Request) (*operations.UpsertAPIV1Response, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}", request.PathParams)

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request)
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

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpsertAPIV1Response{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.UpsertAPIV1Responses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.UpsertAPIV1Responses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIV1Responses{
				API: out,
			}
		}
	default:
		switch contentType {
		case `application/json; charset=UTF-8`:
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIV1Responses{
				Error: out,
			}
		}
	}

	return res, nil
}
