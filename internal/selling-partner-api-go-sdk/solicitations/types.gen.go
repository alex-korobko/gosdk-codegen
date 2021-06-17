// Package solicitations provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package solicitations

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// The response schema for the createProductReviewAndSellerFeedbackSolicitation operation.
type CreateProductReviewAndSellerFeedbackSolicitationResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
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

// GetSchemaResponse defines model for GetSchemaResponse.
type GetSchemaResponse struct {
	Links *struct {

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A JSON schema document describing the expected payload of the action. This object can be validated against <a href=http://json-schema.org/draft-04/schema>http://json-schema.org/draft-04/schema</a>.
	Payload *Schema `json:"payload,omitempty"`
}

// Describes a solicitation action that can be taken for an order. Provides a JSON Hypertext Application Language (HAL) link to the JSON schema document that describes the expected input.
type GetSolicitationActionResponse struct {
	Embedded *struct {
		Schema *GetSchemaResponse `json:"schema,omitempty"`
	} `json:"_embedded,omitempty"`
	Links *struct {

		// A Link object.
		Schema LinkObject `json:"schema"`

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A simple object containing the name of the template.
	Payload *SolicitationsAction `json:"payload,omitempty"`
}

// The response schema for the getSolicitationActionsForOrder operation.
type GetSolicitationActionsForOrderResponse struct {
	Embedded *struct {
		Actions []GetSolicitationActionResponse `json:"actions"`
	} `json:"_embedded,omitempty"`
	Links *struct {

		// Eligible actions for the specified amazonOrderId.
		Actions []LinkObject `json:"actions"`

		// A Link object.
		Self LinkObject `json:"self"`
	} `json:"_links,omitempty"`

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// A Link object.
type LinkObject struct {

	// A URI for this object.
	Href string `json:"href"`

	// An identifier for this object.
	Name *string `json:"name,omitempty"`
}

// A JSON schema document describing the expected payload of the action. This object can be validated against <a href=http://json-schema.org/draft-04/schema>http://json-schema.org/draft-04/schema</a>.
type Schema struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// A simple object containing the name of the template.
type SolicitationsAction struct {
	Name string `json:"name"`
}

// GetSolicitationActionsForOrderParams defines parameters for GetSolicitationActionsForOrder.
type GetSolicitationActionsForOrderParams struct {

	// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// CreateProductReviewAndSellerFeedbackSolicitationParams defines parameters for CreateProductReviewAndSellerFeedbackSolicitation.
type CreateProductReviewAndSellerFeedbackSolicitationParams struct {

	// A marketplace identifier. This specifies the marketplace in which the order was placed. Only one marketplace can be specified.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// Getter for additional properties for Schema. Returns the specified
// element and whether it was found
func (a Schema) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for Schema
func (a *Schema) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for Schema to handle AdditionalProperties
func (a *Schema) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("error unmarshaling field %s", fieldName))
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for Schema to handle AdditionalProperties
func (a Schema) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("error marshaling '%s'", fieldName))
		}
	}
	return json.Marshal(object)
}
