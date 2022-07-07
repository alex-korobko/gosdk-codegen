// Package vendorTransaction provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package vendorTransaction

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// GetTransactionResponse defines model for GetTransactionResponse.
type GetTransactionResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The payload for the getTransactionStatus operation.
	Payload *TransactionStatus `json:"payload,omitempty"`
}

// Transaction defines model for Transaction.
type Transaction struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// Current processing status of the transaction.
	Status string `json:"status"`

	// The unique identifier sent in the 'transactionId' field in response to the post request of a specific transaction.
	TransactionId string `json:"transactionId"`
}

// TransactionStatus defines model for TransactionStatus.
type TransactionStatus struct {

	// The transaction status details.
	TransactionStatus *Transaction `json:"transactionStatus,omitempty"`
}
