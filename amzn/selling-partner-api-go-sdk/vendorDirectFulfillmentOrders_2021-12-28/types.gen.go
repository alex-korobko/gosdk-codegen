// Package vendorOrders provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package vendorOrders

import (
	"time"
)

// AcknowledgementStatus defines model for AcknowledgementStatus.
type AcknowledgementStatus struct {

	// Acknowledgement code is a unique two digit value which indicates the status of the acknowledgement. For a list of acknowledgement codes that Amazon supports, see the Vendor Direct Fulfillment APIs Use Case Guide.
	Code *string `json:"code,omitempty"`

	// Reason for the acknowledgement code.
	Description *string `json:"description,omitempty"`
}

// Address defines model for Address.
type Address struct {

	// First line of the address.
	AddressLine1 string `json:"addressLine1"`

	// Additional address information, if required.
	AddressLine2 *string `json:"addressLine2,omitempty"`

	// Additional address information, if required.
	AddressLine3 *string `json:"addressLine3,omitempty"`

	// The attention name of the person at that address.
	Attention *string `json:"attention,omitempty"`

	// The city where the person, business or institution is located.
	City *string `json:"city,omitempty"`

	// The two digit country code. In ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode"`

	// The county where person, business or institution is located.
	County *string `json:"county,omitempty"`

	// The district where person, business or institution is located.
	District *string `json:"district,omitempty"`

	// The name of the person, business or institution at that address.
	Name string `json:"name"`

	// The phone number of the person, business or institution located at that address.
	Phone *string `json:"phone,omitempty"`

	// The postal code of that address. It conatins a series of letters or digits or both, sometimes including spaces or punctuation.
	PostalCode *string `json:"postalCode,omitempty"`

	// The state or region where person, business or institution is located.
	StateOrRegion string `json:"stateOrRegion"`
}

// Decimal defines model for Decimal.
type Decimal string

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
type ErrorList struct {
	Errors []Error `json:"errors"`
}

// GiftDetails defines model for GiftDetails.
type GiftDetails struct {

	// Gift message to be printed in shipment.
	GiftMessage *string `json:"giftMessage,omitempty"`

	// Gift wrap identifier for the gift wrapping, if any.
	GiftWrapId *string `json:"giftWrapId,omitempty"`
}

// ItemQuantity defines model for ItemQuantity.
type ItemQuantity struct {

	// Acknowledged quantity. This value should not be zero.
	Amount *int `json:"amount,omitempty"`

	// Unit of measure for the acknowledged quantity.
	UnitOfMeasure *string `json:"unitOfMeasure,omitempty"`
}

// Money defines model for Money.
type Money struct {

	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
	Amount *Decimal `json:"amount,omitempty"`

	// Three digit currency code in ISO 4217 format. String of length 3.
	CurrencyCode *string `json:"currencyCode,omitempty"`
}

// Order defines model for Order.
type Order struct {

	// Details of an order.
	OrderDetails *OrderDetails `json:"orderDetails,omitempty"`

	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	PurchaseOrderNumber string `json:"purchaseOrderNumber"`
}

// OrderAcknowledgementItem defines model for OrderAcknowledgementItem.
type OrderAcknowledgementItem struct {

	// The date and time when the order is acknowledged, in ISO-8601 date/time format. For example: 2018-07-16T23:00:00Z / 2018-07-16T23:00:00-05:00 / 2018-07-16T23:00:00-08:00.
	AcknowledgementDate time.Time `json:"acknowledgementDate"`

	// Status of acknowledgement.
	AcknowledgementStatus AcknowledgementStatus `json:"acknowledgementStatus"`

	// Item details including acknowledged quantity.
	ItemAcknowledgements []OrderItemAcknowledgement `json:"itemAcknowledgements"`

	// The purchase order number for this order. Formatting Notes: alpha-numeric code.
	PurchaseOrderNumber string              `json:"purchaseOrderNumber"`
	SellingParty        PartyIdentification `json:"sellingParty"`
	ShipFromParty       PartyIdentification `json:"shipFromParty"`

	// The vendor's order number for this order.
	VendorOrderNumber string `json:"vendorOrderNumber"`
}

// OrderDetails defines model for OrderDetails.
type OrderDetails struct {
	BillToParty PartyIdentification `json:"billToParty"`

	// The customer order number.
	CustomerOrderNumber string `json:"customerOrderNumber"`

	// A list of items in this purchase order.
	Items []OrderItem `json:"items"`

	// The date the order was placed. This  field is expected to be in ISO-8601 date/time format, for example:2018-07-16T23:00:00Z/ 2018-07-16T23:00:00-05:00 /2018-07-16T23:00:00-08:00. If no time zone is specified, UTC should be assumed.
	OrderDate time.Time `json:"orderDate"`

	// Current status of the order.
	OrderStatus   *string             `json:"orderStatus,omitempty"`
	SellingParty  PartyIdentification `json:"sellingParty"`
	ShipFromParty PartyIdentification `json:"shipFromParty"`

	// Address of the party.
	ShipToParty Address `json:"shipToParty"`

	// Shipment details required for the shipment.
	ShipmentDetails ShipmentDetails `json:"shipmentDetails"`

	// Total tax details for the line item.
	TaxTotal *TaxItemDetails `json:"taxTotal,omitempty"`
}

// OrderItem defines model for OrderItem.
type OrderItem struct {

	// Buyer's standard identification number (ASIN) of an item.
	BuyerProductIdentifier *string `json:"buyerProductIdentifier,omitempty"`

	// Gift details for the item.
	GiftDetails *GiftDetails `json:"giftDetails,omitempty"`

	// Numbering of the item on the purchase order. The first item will be 1, the second 2, and so on.
	ItemSequenceNumber string `json:"itemSequenceNumber"`

	// An amount of money, including units in the form of currency.
	NetPrice Money `json:"netPrice"`

	// Details of quantity ordered.
	OrderedQuantity ItemQuantity `json:"orderedQuantity"`

	// Dates for the scheduled delivery shipments.
	ScheduledDeliveryShipment *ScheduledDeliveryShipment `json:"scheduledDeliveryShipment,omitempty"`

	// Total tax details for the line item.
	TaxDetails *TaxItemDetails `json:"taxDetails,omitempty"`

	// Title for the item.
	Title *string `json:"title,omitempty"`

	// An amount of money, including units in the form of currency.
	TotalPrice *Money `json:"totalPrice,omitempty"`

	// The vendor selected product identification of the item.
	VendorProductIdentifier *string `json:"vendorProductIdentifier,omitempty"`
}

// OrderItemAcknowledgement defines model for OrderItemAcknowledgement.
type OrderItemAcknowledgement struct {

	// Details of quantity ordered.
	AcknowledgedQuantity ItemQuantity `json:"acknowledgedQuantity"`

	// Buyer's standard identification number (ASIN) of an item.
	BuyerProductIdentifier *string `json:"buyerProductIdentifier,omitempty"`

	// Line item sequence number for the item.
	ItemSequenceNumber string `json:"itemSequenceNumber"`

	// The vendor selected product identification of the item. Should be the same as was provided in the purchase order.
	VendorProductIdentifier *string `json:"vendorProductIdentifier,omitempty"`
}

// OrderList defines model for OrderList.
type OrderList struct {
	Orders     *[]Order    `json:"orders,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {

	// A generated string used to pass information to your next request. If NextToken is returned, pass the value of NextToken to the next request. If NextToken is not returned, there are no more order items to return.
	NextToken *string `json:"nextToken,omitempty"`
}

// PartyIdentification defines model for PartyIdentification.
type PartyIdentification struct {

	// Address of the party.
	Address *Address `json:"address,omitempty"`

	// Assigned identification for the party. For example, warehouse code or vendor code. Please refer to specific party for more details.
	PartyId string `json:"partyId"`

	// Tax registration details of the entity.
	TaxInfo *TaxRegistrationDetails `json:"taxInfo,omitempty"`
}

// ScheduledDeliveryShipment defines model for ScheduledDeliveryShipment.
type ScheduledDeliveryShipment struct {

	// Earliest nominated delivery date for the scheduled delivery.
	EarliestNominatedDeliveryDate *time.Time `json:"earliestNominatedDeliveryDate,omitempty"`

	// Latest nominated delivery date for the scheduled delivery.
	LatestNominatedDeliveryDate *time.Time `json:"latestNominatedDeliveryDate,omitempty"`

	// Scheduled delivery service type.
	ScheduledDeliveryServiceType *string `json:"scheduledDeliveryServiceType,omitempty"`
}

// ShipmentDates defines model for ShipmentDates.
type ShipmentDates struct {

	// Delivery date promised to the Amazon customer.
	PromisedDeliveryDate *time.Time `json:"promisedDeliveryDate,omitempty"`

	// Time by which the vendor is required to ship the order.
	RequiredShipDate time.Time `json:"requiredShipDate"`
}

// ShipmentDetails defines model for ShipmentDetails.
type ShipmentDetails struct {

	// When true, the order contain a gift. Include the gift message and gift wrap information.
	IsGift *bool `json:"isGift,omitempty"`

	// When true, this is a priority shipment.
	IsPriorityShipment bool `json:"isPriorityShipment"`

	// When true, a packing slip is required to be sent to the customer.
	IsPslipRequired bool `json:"isPslipRequired"`

	// When true, this order is part of a scheduled delivery program.
	IsScheduledDeliveryShipment *bool `json:"isScheduledDeliveryShipment,omitempty"`

	// Message to customer for order status.
	MessageToCustomer string `json:"messageToCustomer"`

	// Ship method to be used for shipping the order. Amazon defines ship method codes indicating the shipping carrier and shipment service level. To see the full list of ship methods in use, including both the code and the friendly name, search the 'Help' section on Vendor Central for 'ship methods'.
	ShipMethod string `json:"shipMethod"`

	// Shipment dates.
	ShipmentDates ShipmentDates `json:"shipmentDates"`
}

// SubmitAcknowledgementRequest defines model for SubmitAcknowledgementRequest.
type SubmitAcknowledgementRequest struct {

	// A list of one or more purchase orders.
	OrderAcknowledgements *[]OrderAcknowledgementItem `json:"orderAcknowledgements,omitempty"`
}

// SubmitAcknowledgementResponse defines model for SubmitAcknowledgementResponse.
type SubmitAcknowledgementResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors  *ErrorList     `json:"errors,omitempty"`
	Payload *TransactionId `json:"payload,omitempty"`
}

// TaxDetails defines model for TaxDetails.
type TaxDetails struct {

	// An amount of money, including units in the form of currency.
	TaxAmount Money `json:"taxAmount"`

	// A decimal number with no loss of precision. Useful when precision loss is unacceptable, as with currencies. Follows RFC7159 for number representation.
	TaxRate *Decimal `json:"taxRate,omitempty"`

	// An amount of money, including units in the form of currency.
	TaxableAmount *Money `json:"taxableAmount,omitempty"`

	// Tax type.
	Type *string `json:"type,omitempty"`
}

// TaxItemDetails defines model for TaxItemDetails.
type TaxItemDetails struct {

	// A list of tax line items.
	TaxLineItem *TaxLineItem `json:"taxLineItem,omitempty"`
}

// TaxLineItem defines model for TaxLineItem.
type TaxLineItem []TaxDetails

// TaxRegistrationDetails defines model for TaxRegistrationDetails.
type TaxRegistrationDetails struct {

	// Address of the party.
	TaxRegistrationAddress *Address `json:"taxRegistrationAddress,omitempty"`

	// Tax registration message that can be used for additional tax related details.
	TaxRegistrationMessages *string `json:"taxRegistrationMessages,omitempty"`

	// Tax registration number for the party. For example, VAT ID.
	TaxRegistrationNumber string `json:"taxRegistrationNumber"`

	// Tax registration type for the entity.
	TaxRegistrationType *string `json:"taxRegistrationType,omitempty"`
}

// TransactionId defines model for TransactionId.
type TransactionId struct {

	// GUID assigned by Amazon to identify this transaction. This value can be used with the Transaction Status API to return the status of this transaction.
	TransactionId *string `json:"transactionId,omitempty"`
}

// SubmitAcknowledgementJSONBody defines parameters for SubmitAcknowledgement.
type SubmitAcknowledgementJSONBody SubmitAcknowledgementRequest

// GetOrdersParams defines parameters for GetOrders.
type GetOrdersParams struct {

	// The vendor warehouse identifier for the fulfillment warehouse. If not specified, the result will contain orders for all warehouses.
	ShipFromPartyId *string `json:"shipFromPartyId,omitempty"`

	// Returns only the purchase orders that match the specified status. If not specified, the result will contain orders that match any status.
	Status *string `json:"status,omitempty"`

	// The limit to the number of purchase orders returned.
	Limit *int64 `json:"limit,omitempty"`

	// Purchase orders that became available after this date and time will be included in the result. Must be in ISO-8601 date/time format.
	CreatedAfter time.Time `json:"createdAfter"`

	// Purchase orders that became available before this date and time will be included in the result. Must be in ISO-8601 date/time format.
	CreatedBefore time.Time `json:"createdBefore"`

	// Sort the list in ascending or descending order by order creation date.
	SortOrder *string `json:"sortOrder,omitempty"`

	// Used for pagination when there are more orders than the specified result size limit. The token value is returned in the previous API call.
	NextToken *string `json:"nextToken,omitempty"`

	// When true, returns the complete purchase order details. Otherwise, only purchase order numbers are returned.
	IncludeDetails *string `json:"includeDetails,omitempty"`
}

// SubmitAcknowledgementRequestBody defines body for SubmitAcknowledgement for application/json ContentType.
type SubmitAcknowledgementJSONRequestBody SubmitAcknowledgementJSONBody