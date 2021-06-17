// Package authorization provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package authorization

// A Login with Amazon (LWA) authorization code.
type AuthorizationCode struct {

	// A Login with Amazon (LWA) authorization code that can be exchanged for a refresh token and access token that authorize you to make calls to a Selling Partner API.
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
}

// Error response returned when the request is unsuccessful.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// A list of error responses returned when a request is unsuccessful.
type ErrorList []Error

// The response schema for the GetAuthorizationCode operation.
type GetAuthorizationCodeResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A Login with Amazon (LWA) authorization code.
	Payload *AuthorizationCode `json:"payload,omitempty"`
}

// GetAuthorizationCodeParams defines parameters for GetAuthorizationCode.
type GetAuthorizationCodeParams struct {

	// The seller ID of the seller for whom you are requesting Selling Partner API authorization. This must be the seller ID of the seller who authorized your application on the Marketplace Appstore.
	SellingPartnerId string `json:"sellingPartnerId"`

	// Your developer ID. This must be one of the developer ID values that you provided when you registered your application in Developer Central.
	DeveloperId string `json:"developerId"`

	// The MWS Auth Token that was generated when the seller authorized your application on the Marketplace Appstore.
	MwsAuthToken string `json:"mwsAuthToken"`
}
