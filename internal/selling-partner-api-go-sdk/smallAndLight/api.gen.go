// Package smallAndLight provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package smallAndLight

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetSmallAndLightEligibilityBySellerSKU request
	GetSmallAndLightEligibilityBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteSmallAndLightEnrollmentBySellerSKU request
	DeleteSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSmallAndLightEnrollmentBySellerSKU request
	GetSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutSmallAndLightEnrollmentBySellerSKU request
	PutSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSmallAndLightFeePreview request  with any body
	GetSmallAndLightFeePreviewWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	GetSmallAndLightFeePreview(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetSmallAndLightEligibilityBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSmallAndLightEligibilityBySellerSKURequest(c.Server, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteSmallAndLightEnrollmentBySellerSKURequest(c.Server, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSmallAndLightEnrollmentBySellerSKURequest(c.Server, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutSmallAndLightEnrollmentBySellerSKU(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutSmallAndLightEnrollmentBySellerSKURequest(c.Server, sellerSKU, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSmallAndLightFeePreviewWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSmallAndLightFeePreviewRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSmallAndLightFeePreview(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSmallAndLightFeePreviewRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetSmallAndLightEligibilityBySellerSKURequest generates requests for GetSmallAndLightEligibilityBySellerSKU
func NewGetSmallAndLightEligibilityBySellerSKURequest(server string, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "sellerSKU", runtime.ParamLocationPath, sellerSKU)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/fba/smallAndLight/v1/eligibilities/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marketplaceIds", runtime.ParamLocationQuery, params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteSmallAndLightEnrollmentBySellerSKURequest generates requests for DeleteSmallAndLightEnrollmentBySellerSKU
func NewDeleteSmallAndLightEnrollmentBySellerSKURequest(server string, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "sellerSKU", runtime.ParamLocationPath, sellerSKU)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marketplaceIds", runtime.ParamLocationQuery, params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSmallAndLightEnrollmentBySellerSKURequest generates requests for GetSmallAndLightEnrollmentBySellerSKU
func NewGetSmallAndLightEnrollmentBySellerSKURequest(server string, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "sellerSKU", runtime.ParamLocationPath, sellerSKU)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marketplaceIds", runtime.ParamLocationQuery, params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutSmallAndLightEnrollmentBySellerSKURequest generates requests for PutSmallAndLightEnrollmentBySellerSKU
func NewPutSmallAndLightEnrollmentBySellerSKURequest(server string, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "sellerSKU", runtime.ParamLocationPath, sellerSKU)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/fba/smallAndLight/v1/enrollments/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	queryValues := queryURL.Query()

	if queryFrag, err := runtime.StyleParamWithLocation("form", true, "marketplaceIds", runtime.ParamLocationQuery, params.MarketplaceIds); err != nil {
		return nil, err
	} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
		return nil, err
	} else {
		for k, v := range parsed {
			for _, v2 := range v {
				queryValues.Add(k, v2)
			}
		}
	}

	queryURL.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("PUT", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSmallAndLightFeePreviewRequest calls the generic GetSmallAndLightFeePreview builder with application/json body
func NewGetSmallAndLightFeePreviewRequest(server string, body GetSmallAndLightFeePreviewJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewGetSmallAndLightFeePreviewRequestWithBody(server, "application/json", bodyReader)
}

// NewGetSmallAndLightFeePreviewRequestWithBody generates requests for GetSmallAndLightFeePreview with any type of body
func NewGetSmallAndLightFeePreviewRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/fba/smallAndLight/v1/feePreviews")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetSmallAndLightEligibilityBySellerSKU request
	GetSmallAndLightEligibilityBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams, reqEditors ...RequestEditorFn) (*GetSmallAndLightEligibilityBySellerSKUResponse, error)

	// DeleteSmallAndLightEnrollmentBySellerSKU request
	DeleteSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*DeleteSmallAndLightEnrollmentBySellerSKUResponse, error)

	// GetSmallAndLightEnrollmentBySellerSKU request
	GetSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*GetSmallAndLightEnrollmentBySellerSKUResponse, error)

	// PutSmallAndLightEnrollmentBySellerSKU request
	PutSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*PutSmallAndLightEnrollmentBySellerSKUResponse, error)

	// GetSmallAndLightFeePreview request  with any body
	GetSmallAndLightFeePreviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetSmallAndLightFeePreviewResponse, error)

	GetSmallAndLightFeePreviewWithResponse(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody, reqEditors ...RequestEditorFn) (*GetSmallAndLightFeePreviewResponse, error)
}

type GetSmallAndLightEligibilityBySellerSKUResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEligibility
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightEligibilityBySellerSKUResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightEligibilityBySellerSKUResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteSmallAndLightEnrollmentBySellerSKUResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r DeleteSmallAndLightEnrollmentBySellerSKUResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSmallAndLightEnrollmentBySellerSKUResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSmallAndLightEnrollmentBySellerSKUResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEnrollment
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightEnrollmentBySellerSKUResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightEnrollmentBySellerSKUResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutSmallAndLightEnrollmentBySellerSKUResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightEnrollment
	JSON400      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON413      *ErrorList
	JSON415      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r PutSmallAndLightEnrollmentBySellerSKUResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutSmallAndLightEnrollmentBySellerSKUResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSmallAndLightFeePreviewResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *SmallAndLightFeePreviews
	JSON400      *ErrorList
	JSON401      *ErrorList
	JSON403      *ErrorList
	JSON404      *ErrorList
	JSON429      *ErrorList
	JSON500      *ErrorList
	JSON503      *ErrorList
}

// Status returns HTTPResponse.Status
func (r GetSmallAndLightFeePreviewResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSmallAndLightFeePreviewResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetSmallAndLightEligibilityBySellerSKUWithResponse request returning *GetSmallAndLightEligibilityBySellerSKUResponse
func (c *ClientWithResponses) GetSmallAndLightEligibilityBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEligibilityBySellerSKUParams, reqEditors ...RequestEditorFn) (*GetSmallAndLightEligibilityBySellerSKUResponse, error) {
	rsp, err := c.GetSmallAndLightEligibilityBySellerSKU(ctx, sellerSKU, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightEligibilityBySellerSKUResponse(rsp)
}

// DeleteSmallAndLightEnrollmentBySellerSKUWithResponse request returning *DeleteSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) DeleteSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *DeleteSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*DeleteSmallAndLightEnrollmentBySellerSKUResponse, error) {
	rsp, err := c.DeleteSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSmallAndLightEnrollmentBySellerSKUResponse(rsp)
}

// GetSmallAndLightEnrollmentBySellerSKUWithResponse request returning *GetSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) GetSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *GetSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*GetSmallAndLightEnrollmentBySellerSKUResponse, error) {
	rsp, err := c.GetSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightEnrollmentBySellerSKUResponse(rsp)
}

// PutSmallAndLightEnrollmentBySellerSKUWithResponse request returning *PutSmallAndLightEnrollmentBySellerSKUResponse
func (c *ClientWithResponses) PutSmallAndLightEnrollmentBySellerSKUWithResponse(ctx context.Context, sellerSKU string, params *PutSmallAndLightEnrollmentBySellerSKUParams, reqEditors ...RequestEditorFn) (*PutSmallAndLightEnrollmentBySellerSKUResponse, error) {
	rsp, err := c.PutSmallAndLightEnrollmentBySellerSKU(ctx, sellerSKU, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutSmallAndLightEnrollmentBySellerSKUResponse(rsp)
}

// GetSmallAndLightFeePreviewWithBodyWithResponse request with arbitrary body returning *GetSmallAndLightFeePreviewResponse
func (c *ClientWithResponses) GetSmallAndLightFeePreviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*GetSmallAndLightFeePreviewResponse, error) {
	rsp, err := c.GetSmallAndLightFeePreviewWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightFeePreviewResponse(rsp)
}

func (c *ClientWithResponses) GetSmallAndLightFeePreviewWithResponse(ctx context.Context, body GetSmallAndLightFeePreviewJSONRequestBody, reqEditors ...RequestEditorFn) (*GetSmallAndLightFeePreviewResponse, error) {
	rsp, err := c.GetSmallAndLightFeePreview(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSmallAndLightFeePreviewResponse(rsp)
}

// ParseGetSmallAndLightEligibilityBySellerSKUResponse parses an HTTP response from a GetSmallAndLightEligibilityBySellerSKUWithResponse call
func ParseGetSmallAndLightEligibilityBySellerSKUResponse(rsp *http.Response) (*GetSmallAndLightEligibilityBySellerSKUResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightEligibilityBySellerSKUResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEligibility
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseDeleteSmallAndLightEnrollmentBySellerSKUResponse parses an HTTP response from a DeleteSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParseDeleteSmallAndLightEnrollmentBySellerSKUResponse(rsp *http.Response) (*DeleteSmallAndLightEnrollmentBySellerSKUResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteSmallAndLightEnrollmentBySellerSKUResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetSmallAndLightEnrollmentBySellerSKUResponse parses an HTTP response from a GetSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParseGetSmallAndLightEnrollmentBySellerSKUResponse(rsp *http.Response) (*GetSmallAndLightEnrollmentBySellerSKUResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightEnrollmentBySellerSKUResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEnrollment
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParsePutSmallAndLightEnrollmentBySellerSKUResponse parses an HTTP response from a PutSmallAndLightEnrollmentBySellerSKUWithResponse call
func ParsePutSmallAndLightEnrollmentBySellerSKUResponse(rsp *http.Response) (*PutSmallAndLightEnrollmentBySellerSKUResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PutSmallAndLightEnrollmentBySellerSKUResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightEnrollment
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 413:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON413 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 415:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON415 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseGetSmallAndLightFeePreviewResponse parses an HTTP response from a GetSmallAndLightFeePreviewWithResponse call
func ParseGetSmallAndLightFeePreviewResponse(rsp *http.Response) (*GetSmallAndLightFeePreviewResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetSmallAndLightFeePreviewResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SmallAndLightFeePreviews
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ErrorList
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
