package speakeasy

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
	"net/http"
	"strings"
	"time"
)

const (
	ServerProd string = "prod"
)

var ServerList = map[string]string{
	ServerProd: "https://api.prod.speakeasyapi.dev",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK Documentation: The Speakeasy API allows teams to manage common operations with their APIs
// https://docs.speakeasyapi.dev - The Speakeasy Platform Documentation
type Speakeasy struct {
	APIEndpoints *apiEndpoints
	Apis         *apis
	Embeds       *embeds
	Metadata     *metadata
	Plugins      *plugins
	Requests     *requests
	Schemas      *schemas

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Speakeasy)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Speakeasy) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Speakeasy) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithServer(server string) SDKOption {
	return func(sdk *Speakeasy) {
		serverURL, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		WithServerURL(serverURL)(sdk)
	}
}

func WithTemplatedServer(server string, params map[string]string) SDKOption {
	return func(sdk *Speakeasy) {
		serverURL, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		WithTemplatedServerURL(serverURL, params)(sdk)
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Speakeasy) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Speakeasy) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Speakeasy {
	sdk := &Speakeasy{
		_language:   "go",
		_sdkVersion: "1.8.2",
		_genVersion: "1.8.4",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[ServerProd]
	}

	sdk.APIEndpoints = newAPIEndpoints(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Apis = newApis(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Embeds = newEmbeds(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Metadata = newMetadata(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Plugins = newPlugins(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Requests = newRequests(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Schemas = newSchemas(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}

// ValidateAPIKey - Validate the current api key.
func (s *Speakeasy) ValidateAPIKey(ctx context.Context) (*operations.ValidateAPIKeyResponse, error) {
	baseURL := s._serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/v1/auth/validate"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s._language, s._sdkVersion, s._genVersion))

	client := s._securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ValidateAPIKeyResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
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

			res.Error = out
		}
	}

	return res, nil
}
