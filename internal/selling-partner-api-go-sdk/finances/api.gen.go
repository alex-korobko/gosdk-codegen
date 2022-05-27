// Package finances provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package finances

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	runt "runtime"
	"strings"

	"github.com/alex-korobko/selling-partner-api-sdk/pkg/runtime"
)

// RequestBeforeFn  is the function signature for the RequestBefore callback function
type RequestBeforeFn func(ctx context.Context, req *http.Request) error

// ResponseAfterFn  is the function signature for the ResponseAfter callback function
type ResponseAfterFn func(ctx context.Context, rsp *http.Response) error

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
	Endpoint string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestBefore RequestBeforeFn

	// A callback for modifying response which are generated before sending over
	// the network.
	ResponseAfter ResponseAfterFn

	// The user agent header identifies your application, its version number, and the platform and programming language you are using.
	// You must include a user agent header in each request submitted to the sales partner API.
	UserAgent string
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(endpoint string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Endpoint: endpoint,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the endpoint URL always has a trailing slash
	if !strings.HasSuffix(client.Endpoint, "/") {
		client.Endpoint += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	// setting the default useragent
	if client.UserAgent == "" {
		client.UserAgent = fmt.Sprintf("selling-partner-api-sdk/v1.0 (Language=%s; Platform=%s-%s)", strings.Replace(runt.Version(), "go", "go/", -1), runt.GOOS, runt.GOARCH)
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

// WithUserAgent set up useragent
// add user agent to every request automatically
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) error {
		c.UserAgent = userAgent
		return nil
	}
}

// WithRequestBefore allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestBefore(fn RequestBeforeFn) ClientOption {
	return func(c *Client) error {
		c.RequestBefore = fn
		return nil
	}
}

// WithResponseAfter allows setting up a callback function, which will be
// called right after get response the request. This can be used to log.
func WithResponseAfter(fn ResponseAfterFn) ClientOption {
	return func(c *Client) error {
		c.ResponseAfter = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListFinancialEventGroups request
	ListFinancialEventGroups(ctx context.Context, params *ListFinancialEventGroupsParams) (*http.Response, error)

	// ListFinancialEventsByGroupId request
	ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*http.Response, error)

	// ListFinancialEvents request
	ListFinancialEvents(ctx context.Context, params *ListFinancialEventsParams) (*http.Response, error)

	// ListFinancialEventsByOrderId request
	ListFinancialEventsByOrderId(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams) (*http.Response, error)
}

func (c *Client) ListFinancialEventGroups(ctx context.Context, params *ListFinancialEventGroupsParams) (*http.Response, error) {
	req, err := NewListFinancialEventGroupsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) ListFinancialEventsByGroupId(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*http.Response, error) {
	req, err := NewListFinancialEventsByGroupIdRequest(c.Endpoint, eventGroupId, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) ListFinancialEvents(ctx context.Context, params *ListFinancialEventsParams) (*http.Response, error) {
	req, err := NewListFinancialEventsRequest(c.Endpoint, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

func (c *Client) ListFinancialEventsByOrderId(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams) (*http.Response, error) {
	req, err := NewListFinancialEventsByOrderIdRequest(c.Endpoint, orderId, params)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.UserAgent)
	if c.RequestBefore != nil {
		err = c.RequestBefore(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if c.ResponseAfter != nil {
		err = c.ResponseAfter(ctx, rsp)
		if err != nil {
			return nil, err
		}
	}
	return rsp, nil
}

// NewListFinancialEventGroupsRequest generates requests for ListFinancialEventGroups
func NewListFinancialEventGroupsRequest(endpoint string, params *ListFinancialEventGroupsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/finances/v0/financialEventGroups")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "MaxResultsPerPage", *params.MaxResultsPerPage); err != nil {
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

	}

	if params.FinancialEventGroupStartedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "FinancialEventGroupStartedBefore", *params.FinancialEventGroupStartedBefore); err != nil {
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

	}

	if params.FinancialEventGroupStartedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "FinancialEventGroupStartedAfter", *params.FinancialEventGroupStartedAfter); err != nil {
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

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "NextToken", *params.NextToken); err != nil {
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

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsByGroupIdRequest generates requests for ListFinancialEventsByGroupId
func NewListFinancialEventsByGroupIdRequest(endpoint string, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "eventGroupId", eventGroupId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/finances/v0/financialEventGroups/%s/financialEvents", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "MaxResultsPerPage", *params.MaxResultsPerPage); err != nil {
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

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "NextToken", *params.NextToken); err != nil {
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

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsRequest generates requests for ListFinancialEvents
func NewListFinancialEventsRequest(endpoint string, params *ListFinancialEventsParams) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/finances/v0/financialEvents")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "MaxResultsPerPage", *params.MaxResultsPerPage); err != nil {
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

	}

	if params.PostedAfter != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "PostedAfter", *params.PostedAfter); err != nil {
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

	}

	if params.PostedBefore != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "PostedBefore", *params.PostedBefore); err != nil {
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

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "NextToken", *params.NextToken); err != nil {
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

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListFinancialEventsByOrderIdRequest generates requests for ListFinancialEventsByOrderId
func NewListFinancialEventsByOrderIdRequest(endpoint string, orderId string, params *ListFinancialEventsByOrderIdParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "orderId", orderId)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/finances/v0/orders/%s/financialEvents", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	queryValues := queryUrl.Query()

	if params.MaxResultsPerPage != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "MaxResultsPerPage", *params.MaxResultsPerPage); err != nil {
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

	}

	if params.NextToken != nil {

		if queryFrag, err := runtime.StyleParam("form", true, "NextToken", *params.NextToken); err != nil {
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

	}

	queryUrl.RawQuery = queryValues.Encode()

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(endpoint string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(endpoint, opts...)
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
		c.Endpoint = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListFinancialEventGroups request
	ListFinancialEventGroupsWithResponse(ctx context.Context, params *ListFinancialEventGroupsParams) (*ListFinancialEventGroupsResp, error)

	// ListFinancialEventsByGroupId request
	ListFinancialEventsByGroupIdWithResponse(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*ListFinancialEventsByGroupIdResp, error)

	// ListFinancialEvents request
	ListFinancialEventsWithResponse(ctx context.Context, params *ListFinancialEventsParams) (*ListFinancialEventsResp, error)

	// ListFinancialEventsByOrderId request
	ListFinancialEventsByOrderIdWithResponse(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams) (*ListFinancialEventsByOrderIdResp, error)
}

type ListFinancialEventGroupsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventGroupsResponse
	JSON400      *ListFinancialEventGroupsResponse
	JSON403      *ListFinancialEventGroupsResponse
	JSON404      *ListFinancialEventGroupsResponse
	JSON429      *ListFinancialEventGroupsResponse
	JSON500      *ListFinancialEventGroupsResponse
	JSON503      *ListFinancialEventGroupsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventGroupsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventGroupsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsByGroupIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsByGroupIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsByGroupIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListFinancialEventsByOrderIdResp struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ListFinancialEventsResponse
	JSON400      *ListFinancialEventsResponse
	JSON403      *ListFinancialEventsResponse
	JSON404      *ListFinancialEventsResponse
	JSON429      *ListFinancialEventsResponse
	JSON500      *ListFinancialEventsResponse
	JSON503      *ListFinancialEventsResponse
}

// Status returns HTTPResponse.Status
func (r ListFinancialEventsByOrderIdResp) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListFinancialEventsByOrderIdResp) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListFinancialEventGroupsWithResponse request returning *ListFinancialEventGroupsResponse
func (c *ClientWithResponses) ListFinancialEventGroupsWithResponse(ctx context.Context, params *ListFinancialEventGroupsParams) (*ListFinancialEventGroupsResp, error) {
	rsp, err := c.ListFinancialEventGroups(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventGroupsResp(rsp)
}

// ListFinancialEventsByGroupIdWithResponse request returning *ListFinancialEventsByGroupIdResponse
func (c *ClientWithResponses) ListFinancialEventsByGroupIdWithResponse(ctx context.Context, eventGroupId string, params *ListFinancialEventsByGroupIdParams) (*ListFinancialEventsByGroupIdResp, error) {
	rsp, err := c.ListFinancialEventsByGroupId(ctx, eventGroupId, params)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsByGroupIdResp(rsp)
}

// ListFinancialEventsWithResponse request returning *ListFinancialEventsResponse
func (c *ClientWithResponses) ListFinancialEventsWithResponse(ctx context.Context, params *ListFinancialEventsParams) (*ListFinancialEventsResp, error) {
	rsp, err := c.ListFinancialEvents(ctx, params)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsResp(rsp)
}

// ListFinancialEventsByOrderIdWithResponse request returning *ListFinancialEventsByOrderIdResponse
func (c *ClientWithResponses) ListFinancialEventsByOrderIdWithResponse(ctx context.Context, orderId string, params *ListFinancialEventsByOrderIdParams) (*ListFinancialEventsByOrderIdResp, error) {
	rsp, err := c.ListFinancialEventsByOrderId(ctx, orderId, params)
	if err != nil {
		return nil, err
	}
	return ParseListFinancialEventsByOrderIdResp(rsp)
}

// ParseListFinancialEventGroupsResp parses an HTTP response from a ListFinancialEventGroupsWithResponse call
func ParseListFinancialEventGroupsResp(rsp *http.Response) (*ListFinancialEventGroupsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventGroupsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventGroupsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsByGroupIdResp parses an HTTP response from a ListFinancialEventsByGroupIdWithResponse call
func ParseListFinancialEventsByGroupIdResp(rsp *http.Response) (*ListFinancialEventsByGroupIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsByGroupIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsResp parses an HTTP response from a ListFinancialEventsWithResponse call
func ParseListFinancialEventsResp(rsp *http.Response) (*ListFinancialEventsResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}

// ParseListFinancialEventsByOrderIdResp parses an HTTP response from a ListFinancialEventsByOrderIdWithResponse call
func ParseListFinancialEventsByOrderIdResp(rsp *http.Response) (*ListFinancialEventsByOrderIdResp, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ListFinancialEventsByOrderIdResp{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 403:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON403 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest ListFinancialEventsResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	return response, nil
}
