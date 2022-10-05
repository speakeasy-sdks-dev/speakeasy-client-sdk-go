package sdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/internal/utils"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"io"
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

func (s *SDK) DeleteAPI(ctx context.Context, request operations.DeleteAPIRequest) (*operations.DeleteAPIResponse, error) {
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

	res := &operations.DeleteAPIResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteAPIResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteAPIResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteAPIResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteAPIEndpoint(ctx context.Context, request operations.DeleteAPIEndpointRequest) (*operations.DeleteAPIEndpointResponse, error) {
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

	res := &operations.DeleteAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteSchema(ctx context.Context, request operations.DeleteSchemaRequest) (*operations.DeleteSchemaResponse, error) {
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

	res := &operations.DeleteSchemaResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteSchemaResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteSchemaResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteSchemaResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DeleteVersionMetadata(ctx context.Context, request operations.DeleteVersionMetadataRequest) (*operations.DeleteVersionMetadataResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata/{metaKey}/{metaValue}", request.PathParams)

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

	res := &operations.DeleteVersionMetadataResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DeleteVersionMetadataResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DeleteVersionMetadataResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DeleteVersionMetadataResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchema(ctx context.Context, request operations.DownloadSchemaRequest) (*operations.DownloadSchemaResponse, error) {
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

	res := &operations.DownloadSchemaResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DownloadSchemaResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DownloadSchemaResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []byte
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaResponses{
				Schema: out,
			}
		case utils.MatchContentType(contentType, `application/x-yaml`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaResponses{
				Schema: data,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) DownloadSchemaRevision(ctx context.Context, request operations.DownloadSchemaRevisionRequest) (*operations.DownloadSchemaRevisionResponse, error) {
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

	res := &operations.DownloadSchemaRevisionResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.DownloadSchemaRevisionResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.DownloadSchemaRevisionResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []byte
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaRevisionResponses{
				Schema: out,
			}
		case utils.MatchContentType(contentType, `application/x-yaml`):
			data, err := io.ReadAll(httpRes.Body)
			if err != nil {
				return nil, fmt.Errorf("error reading response body: %w", err)
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaRevisionResponses{
				Schema: data,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.DownloadSchemaRevisionResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) FindAPIEndpoint(ctx context.Context, request operations.FindAPIEndpointRequest) (*operations.FindAPIEndpointResponse, error) {
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

	res := &operations.FindAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.FindAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.FindAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.FindAPIEndpointResponses{
				APIEndpoint: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.FindAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GenerateOpenAPISpec(ctx context.Context, request operations.GenerateOpenAPISpecRequest) (*operations.GenerateOpenAPISpecResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/generate/openapi", request.PathParams)

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

	res := &operations.GenerateOpenAPISpecResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GenerateOpenAPISpecResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GenerateOpenAPISpecResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.GenerateOpenAPISpecDiff
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateOpenAPISpecResponses{
				GenerateOpenAPISpecDiff: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateOpenAPISpecResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GenerateOpenAPISpecForAPIEndpoint(ctx context.Context, request operations.GenerateOpenAPISpecForAPIEndpointRequest) (*operations.GenerateOpenAPISpecForAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/openapi", request.PathParams)

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

	res := &operations.GenerateOpenAPISpecForAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GenerateOpenAPISpecForAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GenerateOpenAPISpecForAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.GenerateOpenAPISpecDiff
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateOpenAPISpecForAPIEndpointResponses{
				GenerateOpenAPISpecDiff: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateOpenAPISpecForAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GeneratePostmanCollection(ctx context.Context, request operations.GeneratePostmanCollectionRequest) (*operations.GeneratePostmanCollectionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/generate/postman", request.PathParams)

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

	res := &operations.GeneratePostmanCollectionResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GeneratePostmanCollectionResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GeneratePostmanCollectionResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []byte
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GeneratePostmanCollectionResponses{
				PostmanCollection: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GeneratePostmanCollectionResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GeneratePostmanCollectionForAPIEndpoint(ctx context.Context, request operations.GeneratePostmanCollectionForAPIEndpointRequest) (*operations.GeneratePostmanCollectionForAPIEndpointResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID}/generate/postman", request.PathParams)

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

	res := &operations.GeneratePostmanCollectionForAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GeneratePostmanCollectionForAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GeneratePostmanCollectionForAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []byte
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GeneratePostmanCollectionForAPIEndpointResponses{
				PostmanCollection: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GeneratePostmanCollectionForAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GenerateRequestPostmanCollection(ctx context.Context, request operations.GenerateRequestPostmanCollectionRequest) (*operations.GenerateRequestPostmanCollectionResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/eventlog/{requestID}/generate/postman", request.PathParams)

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

	res := &operations.GenerateRequestPostmanCollectionResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GenerateRequestPostmanCollectionResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GenerateRequestPostmanCollectionResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []byte
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateRequestPostmanCollectionResponses{
				PostmanCollection: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GenerateRequestPostmanCollectionResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIEndpoints(ctx context.Context, request operations.GetAllAPIEndpointsRequest) (*operations.GetAllAPIEndpointsResponse, error) {
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

	res := &operations.GetAllAPIEndpointsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllAPIEndpointsResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllAPIEndpointsResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIEndpointsResponses{
				APIEndpoints: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIEndpointsResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllAPIVersions(ctx context.Context, request operations.GetAllAPIVersionsRequest) (*operations.GetAllAPIVersionsResponse, error) {
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

	res := &operations.GetAllAPIVersionsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllAPIVersionsResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllAPIVersionsResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIVersionsResponses{
				Apis: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllAPIVersionsResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAllForVersionAPIEndpoints(ctx context.Context, request operations.GetAllForVersionAPIEndpointsRequest) (*operations.GetAllForVersionAPIEndpointsResponse, error) {
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

	res := &operations.GetAllForVersionAPIEndpointsResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAllForVersionAPIEndpointsResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAllForVersionAPIEndpointsResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllForVersionAPIEndpointsResponses{
				APIEndpoints: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAllForVersionAPIEndpointsResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetAPIEndpoint(ctx context.Context, request operations.GetAPIEndpointRequest) (*operations.GetAPIEndpointResponse, error) {
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

	res := &operations.GetAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAPIEndpointResponses{
				APIEndpoint: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetApis(ctx context.Context, request operations.GetApisRequest) (*operations.GetApisResponse, error) {
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

	res := &operations.GetApisResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetApisResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetApisResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetApisResponses{
				Apis: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetApisResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetEmbedAccessToken(ctx context.Context, request operations.GetEmbedAccessTokenRequest) (*operations.GetEmbedAccessTokenResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/workspace/embed-access-token"

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

	res := &operations.GetEmbedAccessTokenResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetEmbedAccessTokenResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetEmbedAccessTokenResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.EmbedAccessTokenResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetEmbedAccessTokenResponses{
				EmbedAccessTokenResponse: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetEmbedAccessTokenResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetRequestFromEventLog(ctx context.Context, request operations.GetRequestFromEventLogRequest) (*operations.GetRequestFromEventLogResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/eventlog/{requestID}", request.PathParams)

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

	res := &operations.GetRequestFromEventLogResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetRequestFromEventLogResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetRequestFromEventLogResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.UnboundedRequest
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetRequestFromEventLogResponses{
				UnboundedRequest: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetRequestFromEventLogResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchema(ctx context.Context, request operations.GetSchemaRequest) (*operations.GetSchemaResponse, error) {
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

	res := &operations.GetSchemaResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaResponses{
				Schema: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaDiff(ctx context.Context, request operations.GetSchemaDiffRequest) (*operations.GetSchemaDiffResponse, error) {
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

	res := &operations.GetSchemaDiffResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaDiffResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaDiffResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SchemaDiff
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaDiffResponses{
				SchemaDiff: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaDiffResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemaRevision(ctx context.Context, request operations.GetSchemaRevisionRequest) (*operations.GetSchemaRevisionResponse, error) {
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

	res := &operations.GetSchemaRevisionResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemaRevisionResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemaRevisionResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaRevisionResponses{
				Schema: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemaRevisionResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetSchemas(ctx context.Context, request operations.GetSchemasRequest) (*operations.GetSchemasResponse, error) {
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

	res := &operations.GetSchemasResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetSchemasResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetSchemasResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.Schema
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemasResponses{
				Schemata: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetSchemasResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetValidEmbedAccessTokens(ctx context.Context) (*operations.GetValidEmbedAccessTokensResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/workspace/embed-access-tokens/valid"

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

	res := &operations.GetValidEmbedAccessTokensResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetValidEmbedAccessTokensResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetValidEmbedAccessTokensResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.EmbedToken
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetValidEmbedAccessTokensResponses{
				EmbedTokens: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetValidEmbedAccessTokensResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) GetVersionMetadata(ctx context.Context, request operations.GetVersionMetadataRequest) (*operations.GetVersionMetadataResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata", request.PathParams)

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

	res := &operations.GetVersionMetadataResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.GetVersionMetadataResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.GetVersionMetadataResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetVersionMetadataResponses{
				VersionMetadata: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.GetVersionMetadataResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) InsertVersionMetadata(ctx context.Context, request operations.InsertVersionMetadataRequest) (*operations.InsertVersionMetadataResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/apis/{apiID}/version/{versionID}/metadata", request.PathParams)

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

	res := &operations.InsertVersionMetadataResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.InsertVersionMetadataResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.InsertVersionMetadataResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.VersionMetadata
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.InsertVersionMetadataResponses{
				VersionMetadata: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.InsertVersionMetadataResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) QueryEventLog(ctx context.Context, request operations.QueryEventLogRequest) (*operations.QueryEventLogResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/eventlog/query"

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

	res := &operations.QueryEventLogResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.QueryEventLogResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.QueryEventLogResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.BoundedRequest
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.QueryEventLogResponses{
				BoundedRequests: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.QueryEventLogResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) RegisterSchema(ctx context.Context, request operations.RegisterSchemaRequest) (*operations.RegisterSchemaResponse, error) {
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

	res := &operations.RegisterSchemaResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.RegisterSchemaResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.RegisterSchemaResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.RegisterSchemaResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) RevokeEmbedAccessToken(ctx context.Context, request operations.RevokeEmbedAccessTokenRequest) (*operations.RevokeEmbedAccessTokenResponse, error) {
	baseURL := s.serverURL
	url := utils.GenerateURL(ctx, baseURL, "/v1/workspace/embed-access-tokens/{tokenID}", request.PathParams)

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

	res := &operations.RevokeEmbedAccessTokenResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.RevokeEmbedAccessTokenResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.RevokeEmbedAccessTokenResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.RevokeEmbedAccessTokenResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPI(ctx context.Context, request operations.UpsertAPIRequest) (*operations.UpsertAPIResponse, error) {
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

	res := &operations.UpsertAPIResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.UpsertAPIResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.UpsertAPIResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.API
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIResponses{
				API: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIResponses{
				Error: out,
			}
		}
	}

	return res, nil
}

func (s *SDK) UpsertAPIEndpoint(ctx context.Context, request operations.UpsertAPIEndpointRequest) (*operations.UpsertAPIEndpointResponse, error) {
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

	res := &operations.UpsertAPIEndpointResponse{
		StatusCode:  int64(httpRes.StatusCode),
		ContentType: contentType,
		Responses:   make(map[int64]map[string]operations.UpsertAPIEndpointResponses),
	}

	if _, ok := res.Responses[int64(httpRes.StatusCode)]; !ok {
		res.Responses[int64(httpRes.StatusCode)] = make(map[string]operations.UpsertAPIEndpointResponses)
	}

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.APIEndpoint
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIEndpointResponses{
				APIEndpoint: out,
			}
		}
	default:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.Error
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.Responses[int64(httpRes.StatusCode)][contentType] = operations.UpsertAPIEndpointResponses{
				Error: out,
			}
		}
	}

	return res, nil
}
