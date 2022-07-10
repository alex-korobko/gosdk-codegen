// Package ordersV0 provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package ordersV0

// Address defines model for Address.
type Address struct {

	// The street address.
	AddressLine1 *string `json:"AddressLine1,omitempty"`

	// Additional street address information, if required.
	AddressLine2 *string `json:"AddressLine2,omitempty"`

	// Additional street address information, if required.
	AddressLine3 *string `json:"AddressLine3,omitempty"`

	// The address type of the shipping address.
	AddressType *string `json:"AddressType,omitempty"`

	// The city
	City *string `json:"City,omitempty"`

	// The country code. A two-character country code, in ISO 3166-1 alpha-2 format.
	CountryCode *string `json:"CountryCode,omitempty"`

	// The county.
	County *string `json:"County,omitempty"`

	// The district.
	District *string `json:"District,omitempty"`

	// The municipality.
	Municipality *string `json:"Municipality,omitempty"`

	// The name.
	Name string `json:"Name"`

	// The phone number. Not returned for Fulfillment by Amazon (FBA) orders.
	Phone *string `json:"Phone,omitempty"`

	// The postal code.
	PostalCode *string `json:"PostalCode,omitempty"`

	// The state or region.
	StateOrRegion *string `json:"StateOrRegion,omitempty"`
}

// AutomatedShippingSettings defines model for AutomatedShippingSettings.
type AutomatedShippingSettings struct {

	// Auto-generated carrier for SSA orders.
	AutomatedCarrier *string `json:"AutomatedCarrier,omitempty"`

	// Auto-generated ship method for SSA orders.
	AutomatedShipMethod *string `json:"AutomatedShipMethod,omitempty"`

	// If true, this order has automated shipping settings generated by Amazon. This order could be identified as an SSA order.
	HasAutomatedShippingSettings *bool `json:"HasAutomatedShippingSettings,omitempty"`
}

// BuyerCustomizedInfoDetail defines model for BuyerCustomizedInfoDetail.
type BuyerCustomizedInfoDetail struct {

	// The location of a zip file containing Amazon Custom data.
	CustomizedURL *string `json:"CustomizedURL,omitempty"`
}

// BuyerInfo defines model for BuyerInfo.
type BuyerInfo struct {

	// The county of the buyer.
	BuyerCounty *string `json:"BuyerCounty,omitempty"`

	// The anonymized email address of the buyer.
	BuyerEmail *string `json:"BuyerEmail,omitempty"`

	// The name of the buyer.
	BuyerName *string `json:"BuyerName,omitempty"`

	// Tax information about the buyer.
	BuyerTaxInfo *BuyerTaxInfo `json:"BuyerTaxInfo,omitempty"`

	// The purchase order (PO) number entered by the buyer at checkout. Returned only for orders where the buyer entered a PO number at checkout.
	PurchaseOrderNumber *string `json:"PurchaseOrderNumber,omitempty"`
}

// BuyerRequestedCancel defines model for BuyerRequestedCancel.
type BuyerRequestedCancel struct {

	// The reason that the buyer requested cancellation.
	BuyerCancelReason *string `json:"BuyerCancelReason,omitempty"`

	// When true, the buyer has requested cancellation.
	IsBuyerRequestedCancel *bool `json:"IsBuyerRequestedCancel,omitempty"`
}

// BuyerTaxInfo defines model for BuyerTaxInfo.
type BuyerTaxInfo struct {

	// The legal name of the company.
	CompanyLegalName *string `json:"CompanyLegalName,omitempty"`

	// A list of tax classifications that apply to the order.
	TaxClassifications *[]TaxClassification `json:"TaxClassifications,omitempty"`

	// The country or region imposing the tax.
	TaxingRegion *string `json:"TaxingRegion,omitempty"`
}

// BuyerTaxInformation defines model for BuyerTaxInformation.
type BuyerTaxInformation struct {

	// Business buyer's address.
	BuyerBusinessAddress *string `json:"BuyerBusinessAddress,omitempty"`

	// Business buyer's company legal name.
	BuyerLegalCompanyName *string `json:"BuyerLegalCompanyName,omitempty"`

	// Business buyer's tax office.
	BuyerTaxOffice *string `json:"BuyerTaxOffice,omitempty"`

	// Business buyer's tax registration ID.
	BuyerTaxRegistrationId *string `json:"BuyerTaxRegistrationId,omitempty"`
}

// Error defines model for Error.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional details that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message string `json:"message"`
}

// ErrorList defines model for ErrorList.
type ErrorList []Error

// FulfillmentInstruction defines model for FulfillmentInstruction.
type FulfillmentInstruction struct {

	// Denotes the recommended sourceId where the order should be fulfilled from.
	FulfillmentSupplySourceId *string `json:"FulfillmentSupplySourceId,omitempty"`
}

// GetOrderAddressResponse defines model for GetOrderAddressResponse.
type GetOrderAddressResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The shipping address for the order.
	Payload *OrderAddress `json:"payload,omitempty"`
}

// GetOrderBuyerInfoResponse defines model for GetOrderBuyerInfoResponse.
type GetOrderBuyerInfoResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// Buyer information for an order.
	Payload *OrderBuyerInfo `json:"payload,omitempty"`
}

// GetOrderItemsBuyerInfoResponse defines model for GetOrderItemsBuyerInfoResponse.
type GetOrderItemsBuyerInfoResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A single order item's buyer information list with the order ID.
	Payload *OrderItemsBuyerInfoList `json:"payload,omitempty"`
}

// GetOrderItemsResponse defines model for GetOrderItemsResponse.
type GetOrderItemsResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The order items list along with the order ID.
	Payload *OrderItemsList `json:"payload,omitempty"`
}

// GetOrderRegulatedInfoResponse defines model for GetOrderRegulatedInfoResponse.
type GetOrderRegulatedInfoResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The order's regulated information along with its verification status.
	Payload *OrderRegulatedInfo `json:"payload,omitempty"`
}

// GetOrderResponse defines model for GetOrderResponse.
type GetOrderResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// Order information.
	Payload *Order `json:"payload,omitempty"`
}

// GetOrdersResponse defines model for GetOrdersResponse.
type GetOrdersResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// A list of orders along with additional information to make subsequent API calls.
	Payload *OrdersList `json:"payload,omitempty"`
}

// ItemBuyerInfo defines model for ItemBuyerInfo.
type ItemBuyerInfo struct {

	// Buyer information for custom orders from the Amazon Custom program.
	BuyerCustomizedInfo *BuyerCustomizedInfoDetail `json:"BuyerCustomizedInfo,omitempty"`

	// A gift message provided by the buyer.
	GiftMessageText *string `json:"GiftMessageText,omitempty"`

	// The gift wrap level specified by the buyer.
	GiftWrapLevel *string `json:"GiftWrapLevel,omitempty"`

	// The monetary value of the order.
	GiftWrapPrice *Money `json:"GiftWrapPrice,omitempty"`

	// The monetary value of the order.
	GiftWrapTax *Money `json:"GiftWrapTax,omitempty"`
}

// MarketplaceId defines model for MarketplaceId.
type MarketplaceId string

// MarketplaceTaxInfo defines model for MarketplaceTaxInfo.
type MarketplaceTaxInfo struct {

	// A list of tax classifications that apply to the order.
	TaxClassifications *[]TaxClassification `json:"TaxClassifications,omitempty"`
}

// Money defines model for Money.
type Money struct {

	// The currency amount.
	Amount *string `json:"Amount,omitempty"`

	// The three-digit currency code. In ISO 4217 format.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
}

// Order defines model for Order.
type Order struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// Contains information regarding the Shipping Settings Automation program, such as whether the order's shipping settings were generated automatically, and what those settings are.
	AutomatedShippingSettings *AutomatedShippingSettings `json:"AutomatedShippingSettings,omitempty"`

	// Buyer information
	BuyerInfo *BuyerInfo `json:"BuyerInfo,omitempty"`

	// The buyer's invoicing preference. Available only in the TR marketplace.
	BuyerInvoicePreference *string `json:"BuyerInvoicePreference,omitempty"`

	// Contains the business invoice tax information. Available only in the TR marketplace.
	BuyerTaxInformation *BuyerTaxInformation `json:"BuyerTaxInformation,omitempty"`

	// Custom ship label for Checkout by Amazon (CBA).
	CbaDisplayableShippingLabel *string `json:"CbaDisplayableShippingLabel,omitempty"`

	// The shipping address for the order.
	DefaultShipFromLocationAddress *Address `json:"DefaultShipFromLocationAddress,omitempty"`

	// The start of the time period within which you have committed to fulfill the order. In ISO 8601 date time format. Returned only for seller-fulfilled orders.
	EarliestDeliveryDate *string `json:"EarliestDeliveryDate,omitempty"`

	// The start of the time period within which you have committed to ship the order. In ISO 8601 date time format. Returned only for seller-fulfilled orders.
	//
	// Note: EarliestShipDate might not be returned for orders placed before February 1, 2013.
	EarliestShipDate *string `json:"EarliestShipDate,omitempty"`

	// The status of the Amazon Easy Ship order. This property is included only for Amazon Easy Ship orders.
	//
	// Possible values: PendingPickUp, LabelCanceled, PickedUp, OutForDelivery, Damaged, Delivered, RejectedByBuyer, Undeliverable, ReturnedToSeller, ReturningToSeller.
	EasyShipShipmentStatus *string `json:"EasyShipShipmentStatus,omitempty"`

	// Whether the order was fulfilled by Amazon (AFN) or by the seller (MFN).
	FulfillmentChannel *string `json:"FulfillmentChannel,omitempty"`

	// Contains the instructions about the fulfillment like where should it be fulfilled from.
	FulfillmentInstruction *FulfillmentInstruction `json:"FulfillmentInstruction,omitempty"`

	// Whether the order contains regulated items which may require additional approval steps before being fulfilled.
	HasRegulatedItems *bool `json:"HasRegulatedItems,omitempty"`

	// When true, this order is marked to be delivered to an Access Point. The access location is chosen by the customer. Access Points include Amazon Hub Lockers, Amazon Hub Counters, and pickup points operated by carriers.
	IsAccessPointOrder *bool `json:"IsAccessPointOrder,omitempty"`

	// When true, the order is an Amazon Business order. An Amazon Business order is an order where the buyer is a Verified Business Buyer.
	IsBusinessOrder *bool `json:"IsBusinessOrder,omitempty"`

	// When true, the estimated ship date is set for the order. Returned only for Sourcing on Demand orders.
	IsEstimatedShipDateSet *bool `json:"IsEstimatedShipDateSet,omitempty"`

	// When true, the order is a GlobalExpress order.
	IsGlobalExpressEnabled *bool `json:"IsGlobalExpressEnabled,omitempty"`

	// When true, the item within this order was bought and re-sold by Amazon Business EU SARL (ABEU). By buying and instantly re-selling your items, ABEU becomes the seller of record, making your inventory available for sale to customers who would not otherwise purchase from a third-party seller.
	IsIBA *bool `json:"IsIBA,omitempty"`

	// When true, this order is marked to be picked up from a store rather than delivered.
	IsISPU *bool `json:"IsISPU,omitempty"`

	// When true, the order has a Premium Shipping Service Level Agreement. For more information about Premium Shipping orders, see "Premium Shipping Options" in the Seller Central Help for your marketplace.
	IsPremiumOrder *bool `json:"IsPremiumOrder,omitempty"`

	// When true, the order is a seller-fulfilled Amazon Prime order.
	IsPrime *bool `json:"IsPrime,omitempty"`

	// When true, this is a replacement order.
	IsReplacementOrder *bool `json:"IsReplacementOrder,omitempty"`

	// When true, the item within this order was bought and re-sold by Amazon Business EU SARL (ABEU). By buying and instantly re-selling your items, ABEU becomes the seller of record, making your inventory available for sale to customers who would not otherwise purchase from a third-party seller.
	IsSoldByAB *bool `json:"IsSoldByAB,omitempty"`

	// The date when the order was last updated.
	//
	// Note: LastUpdateDate is returned with an incorrect date for orders that were last updated before 2009-04-01.
	LastUpdateDate string `json:"LastUpdateDate"`

	// The end of the time period within which you have committed to fulfill the order. In ISO 8601 date time format. Returned only for seller-fulfilled orders that do not have a PendingAvailability, Pending, or Canceled status.
	LatestDeliveryDate *string `json:"LatestDeliveryDate,omitempty"`

	// The end of the time period within which you have committed to ship the order. In ISO 8601 date time format. Returned only for seller-fulfilled orders.
	//
	// Note: LatestShipDate might not be returned for orders placed before February 1, 2013.
	LatestShipDate *string `json:"LatestShipDate,omitempty"`

	// The identifier for the marketplace where the order was placed.
	MarketplaceId *string `json:"MarketplaceId,omitempty"`

	// Tax information about the marketplace.
	MarketplaceTaxInfo *MarketplaceTaxInfo `json:"MarketplaceTaxInfo,omitempty"`

	// The number of items shipped.
	NumberOfItemsShipped *int `json:"NumberOfItemsShipped,omitempty"`

	// The number of items unshipped.
	NumberOfItemsUnshipped *int `json:"NumberOfItemsUnshipped,omitempty"`

	// The order channel of the first item in the order.
	OrderChannel *string `json:"OrderChannel,omitempty"`

	// The current order status.
	OrderStatus string `json:"OrderStatus"`

	// The monetary value of the order.
	OrderTotal *Money `json:"OrderTotal,omitempty"`

	// The type of the order.
	OrderType *string `json:"OrderType,omitempty"`

	// A list of payment execution detail items.
	PaymentExecutionDetail *PaymentExecutionDetailItemList `json:"PaymentExecutionDetail,omitempty"`

	// The payment method for the order. This property is limited to Cash On Delivery (COD) and Convenience Store (CVS) payment methods. Unless you need the specific COD payment information provided by the PaymentExecutionDetailItem object, we recommend using the PaymentMethodDetails property to get payment method information.
	PaymentMethod *string `json:"PaymentMethod,omitempty"`

	// A list of payment method detail items.
	PaymentMethodDetails *PaymentMethodDetailItemList `json:"PaymentMethodDetails,omitempty"`

	// Indicates the date by which the seller must respond to the buyer with an estimated ship date. Returned only for Sourcing on Demand orders.
	PromiseResponseDueDate *string `json:"PromiseResponseDueDate,omitempty"`

	// The date when the order was created.
	PurchaseDate string `json:"PurchaseDate"`

	// The order ID value for the order that is being replaced. Returned only if IsReplacementOrder = true.
	ReplacedOrderId *string `json:"ReplacedOrderId,omitempty"`

	// The sales channel of the first item in the order.
	SalesChannel *string `json:"SalesChannel,omitempty"`

	// The seller’s friendly name registered in the marketplace.
	SellerDisplayName *string `json:"SellerDisplayName,omitempty"`

	// A seller-defined order identifier.
	SellerOrderId *string `json:"SellerOrderId,omitempty"`

	// The shipment service level of the order.
	ShipServiceLevel *string `json:"ShipServiceLevel,omitempty"`

	// The shipment service level category of the order.
	//
	// Possible values: Expedited, FreeEconomy, NextDay, SameDay, SecondDay, Scheduled, Standard.
	ShipmentServiceLevelCategory *string `json:"ShipmentServiceLevelCategory,omitempty"`

	// The shipping address for the order.
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`
}

// OrderAddress defines model for OrderAddress.
type OrderAddress struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// The shipping address for the order.
	ShippingAddress *Address `json:"ShippingAddress,omitempty"`
}

// OrderBuyerInfo defines model for OrderBuyerInfo.
type OrderBuyerInfo struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// The county of the buyer.
	BuyerCounty *string `json:"BuyerCounty,omitempty"`

	// The anonymized email address of the buyer.
	BuyerEmail *string `json:"BuyerEmail,omitempty"`

	// The name of the buyer.
	BuyerName *string `json:"BuyerName,omitempty"`

	// Tax information about the buyer.
	BuyerTaxInfo *BuyerTaxInfo `json:"BuyerTaxInfo,omitempty"`

	// The purchase order (PO) number entered by the buyer at checkout. Returned only for orders where the buyer entered a PO number at checkout.
	PurchaseOrderNumber *string `json:"PurchaseOrderNumber,omitempty"`
}

// OrderItem defines model for OrderItem.
type OrderItem struct {

	// The Amazon Standard Identification Number (ASIN) of the item.
	ASIN string `json:"ASIN"`

	// A single item's buyer information.
	BuyerInfo *ItemBuyerInfo `json:"BuyerInfo,omitempty"`

	// Information about whether or not a buyer requested cancellation.
	BuyerRequestedCancel *BuyerRequestedCancel `json:"BuyerRequestedCancel,omitempty"`

	// The monetary value of the order.
	CODFee *Money `json:"CODFee,omitempty"`

	// The monetary value of the order.
	CODFeeDiscount *Money `json:"CODFeeDiscount,omitempty"`

	// The condition of the item.
	//
	// Possible values: New, Used, Collectible, Refurbished, Preorder, Club.
	ConditionId *string `json:"ConditionId,omitempty"`

	// The condition of the item as described by the seller.
	ConditionNote *string `json:"ConditionNote,omitempty"`

	// The subcondition of the item.
	//
	// Possible values: New, Mint, Very Good, Good, Acceptable, Poor, Club, OEM, Warranty, Refurbished Warranty, Refurbished, Open Box, Any, Other.
	ConditionSubtypeId *string `json:"ConditionSubtypeId,omitempty"`

	// The category of deemed reseller. This applies to selling partners that are not based in the EU and is used to help them meet the VAT Deemed Reseller tax laws in the EU and UK.
	DeemedResellerCategory *string `json:"DeemedResellerCategory,omitempty"`

	// The IOSS number for the marketplace. Sellers shipping to the European Union (EU) from outside of the EU must provide this IOSS number to their carrier when Amazon has collected the VAT on the sale.
	IossNumber *string `json:"IossNumber,omitempty"`

	// When true, the item is a gift.
	IsGift *bool `json:"IsGift,omitempty"`

	// When true, transparency codes are required.
	IsTransparency *bool `json:"IsTransparency,omitempty"`

	// The monetary value of the order.
	ItemPrice *Money `json:"ItemPrice,omitempty"`

	// The monetary value of the order.
	ItemTax *Money `json:"ItemTax,omitempty"`

	// An Amazon-defined order item identifier.
	OrderItemId string `json:"OrderItemId"`

	// The number of Amazon Points offered with the purchase of an item, and their monetary value.
	PointsGranted *PointsGrantedDetail `json:"PointsGranted,omitempty"`

	// Indicates that the selling price is a special price that is available only for Amazon Business orders. For more information about the Amazon Business Seller Program, see the [Amazon Business website](https://www.amazon.com/b2b/info/amazon-business).
	//
	// Possible values: BusinessPrice - A special price that is available only for Amazon Business orders.
	PriceDesignation *string `json:"PriceDesignation,omitempty"`

	// Product information on the number of items.
	ProductInfo *ProductInfoDetail `json:"ProductInfo,omitempty"`

	// The monetary value of the order.
	PromotionDiscount *Money `json:"PromotionDiscount,omitempty"`

	// The monetary value of the order.
	PromotionDiscountTax *Money `json:"PromotionDiscountTax,omitempty"`

	// A list of promotion identifiers provided by the seller when the promotions were created.
	PromotionIds *PromotionIdList `json:"PromotionIds,omitempty"`

	// The number of items in the order.
	QuantityOrdered int `json:"QuantityOrdered"`

	// The number of items shipped.
	QuantityShipped *int `json:"QuantityShipped,omitempty"`

	// The end date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date time format.
	ScheduledDeliveryEndDate *string `json:"ScheduledDeliveryEndDate,omitempty"`

	// The start date of the scheduled delivery window in the time zone of the order destination. In ISO 8601 date time format.
	ScheduledDeliveryStartDate *string `json:"ScheduledDeliveryStartDate,omitempty"`

	// The seller stock keeping unit (SKU) of the item.
	SellerSKU *string `json:"SellerSKU,omitempty"`

	// When true, the product type for this item has a serial number.
	//
	// Returned only for Amazon Easy Ship orders.
	SerialNumberRequired *bool `json:"SerialNumberRequired,omitempty"`

	// The monetary value of the order.
	ShippingDiscount *Money `json:"ShippingDiscount,omitempty"`

	// The monetary value of the order.
	ShippingDiscountTax *Money `json:"ShippingDiscountTax,omitempty"`

	// The monetary value of the order.
	ShippingPrice *Money `json:"ShippingPrice,omitempty"`

	// The monetary value of the order.
	ShippingTax *Money `json:"ShippingTax,omitempty"`

	// The store chain store identifier. Linked to a specific store in a store chain.
	StoreChainStoreId *string `json:"StoreChainStoreId,omitempty"`

	// Information about withheld taxes.
	TaxCollection *TaxCollection `json:"TaxCollection,omitempty"`

	// The name of the item.
	Title *string `json:"Title,omitempty"`
}

// OrderItemBuyerInfo defines model for OrderItemBuyerInfo.
type OrderItemBuyerInfo struct {

	// Buyer information for custom orders from the Amazon Custom program.
	BuyerCustomizedInfo *BuyerCustomizedInfoDetail `json:"BuyerCustomizedInfo,omitempty"`

	// A gift message provided by the buyer.
	GiftMessageText *string `json:"GiftMessageText,omitempty"`

	// The gift wrap level specified by the buyer.
	GiftWrapLevel *string `json:"GiftWrapLevel,omitempty"`

	// The monetary value of the order.
	GiftWrapPrice *Money `json:"GiftWrapPrice,omitempty"`

	// The monetary value of the order.
	GiftWrapTax *Money `json:"GiftWrapTax,omitempty"`

	// An Amazon-defined order item identifier.
	OrderItemId string `json:"OrderItemId"`
}

// OrderItemBuyerInfoList defines model for OrderItemBuyerInfoList.
type OrderItemBuyerInfoList []OrderItemBuyerInfo

// OrderItemList defines model for OrderItemList.
type OrderItemList []OrderItem

// OrderItems defines model for OrderItems.
type OrderItems []struct {

	// the unique identifier for the order item
	OrderItemId *string `json:"orderItemId,omitempty"`

	// the quantity of items that needs an update of the shipment status
	Quantity *int `json:"quantity,omitempty"`
}

// OrderItemsBuyerInfoList defines model for OrderItemsBuyerInfoList.
type OrderItemsBuyerInfoList struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`

	// A single order item's buyer information list.
	OrderItems OrderItemBuyerInfoList `json:"OrderItems"`
}

// OrderItemsList defines model for OrderItemsList.
type OrderItemsList struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`

	// A list of order items.
	OrderItems OrderItemList `json:"OrderItems"`
}

// OrderList defines model for OrderList.
type OrderList []Order

// OrderRegulatedInfo defines model for OrderRegulatedInfo.
type OrderRegulatedInfo struct {

	// An Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderId string `json:"AmazonOrderId"`

	// The regulated information collected during purchase and used to verify the order.
	RegulatedInformation RegulatedInformation `json:"RegulatedInformation"`

	// The verification status of the order along with associated approval or rejection metadata.
	RegulatedOrderVerificationStatus RegulatedOrderVerificationStatus `json:"RegulatedOrderVerificationStatus"`

	// Whether the order requires attaching a dosage information label when shipped.
	RequiresDosageLabel bool `json:"RequiresDosageLabel"`
}

// OrdersList defines model for OrdersList.
type OrdersList struct {

	// A date used for selecting orders created before (or at) a specified time. Only orders placed before the specified time are returned. The date must be in ISO 8601 format.
	CreatedBefore *string `json:"CreatedBefore,omitempty"`

	// A date used for selecting orders that were last updated before (or at) a specified time. An update is defined as any change in order status, including the creation of a new order. Includes updates made by Amazon and by the seller. All dates must be in ISO 8601 format.
	LastUpdatedBefore *string `json:"LastUpdatedBefore,omitempty"`

	// When present and not empty, pass this string token in the next request to return the next response page.
	NextToken *string `json:"NextToken,omitempty"`

	// A list of orders.
	Orders OrderList `json:"Orders"`
}

// PaymentExecutionDetailItem defines model for PaymentExecutionDetailItem.
type PaymentExecutionDetailItem struct {

	// The monetary value of the order.
	Payment Money `json:"Payment"`

	// A sub-payment method for a COD order.
	//
	// Possible values:
	//
	// * COD - Cash On Delivery.
	//
	// * GC - Gift Card.
	//
	// * PointsAccount - Amazon Points.
	PaymentMethod string `json:"PaymentMethod"`
}

// PaymentExecutionDetailItemList defines model for PaymentExecutionDetailItemList.
type PaymentExecutionDetailItemList []PaymentExecutionDetailItem

// PaymentMethodDetailItemList defines model for PaymentMethodDetailItemList.
type PaymentMethodDetailItemList []string

// PointsGrantedDetail defines model for PointsGrantedDetail.
type PointsGrantedDetail struct {

	// The monetary value of the order.
	PointsMonetaryValue *Money `json:"PointsMonetaryValue,omitempty"`

	// The number of Amazon Points granted with the purchase of an item.
	PointsNumber *int `json:"PointsNumber,omitempty"`
}

// ProductInfoDetail defines model for ProductInfoDetail.
type ProductInfoDetail struct {

	// The total number of items that are included in the ASIN.
	NumberOfItems *int `json:"NumberOfItems,omitempty"`
}

// PromotionIdList defines model for PromotionIdList.
type PromotionIdList []string

// RegulatedInformation defines model for RegulatedInformation.
type RegulatedInformation struct {

	// A list of regulated information fields as collected from the regulatory form.
	Fields []RegulatedInformationField `json:"Fields"`
}

// RegulatedInformationField defines model for RegulatedInformationField.
type RegulatedInformationField struct {

	// The unique identifier for the field.
	FieldId string `json:"FieldId"`

	// The human-readable name for the field.
	FieldLabel string `json:"FieldLabel"`

	// The type of field the field.
	FieldType string `json:"FieldType"`

	// The content of the field as collected in regulatory form. Note that FileAttachment type fields will contain an URL to download the attachment here.
	FieldValue string `json:"FieldValue"`
}

// RegulatedOrderVerificationStatus defines model for RegulatedOrderVerificationStatus.
type RegulatedOrderVerificationStatus struct {

	// The identifier for the order's regulated information reviewer.
	ExternalReviewerId *string `json:"ExternalReviewerId,omitempty"`

	// The reason for rejecting the order's regulated information. Not present if the order isn't rejected.
	RejectionReason *RejectionReason `json:"RejectionReason,omitempty"`

	// Whether the regulated information provided in the order requires a review by the merchant.
	RequiresMerchantAction bool `json:"RequiresMerchantAction"`

	// The date the order was reviewed. In ISO 8601 date time format.
	ReviewDate *string `json:"ReviewDate,omitempty"`

	// The verification status of the order.
	Status VerificationStatus `json:"Status"`

	// A list of valid rejection reasons that may be used to reject the order's regulated information.
	ValidRejectionReasons []RejectionReason `json:"ValidRejectionReasons"`
}

// RejectionReason defines model for RejectionReason.
type RejectionReason struct {

	// The human-readable description of this rejection reason.
	RejectionReasonDescription string `json:"RejectionReasonDescription"`

	// The unique identifier for the rejection reason.
	RejectionReasonId string `json:"RejectionReasonId"`
}

// ShipmentStatus defines model for ShipmentStatus.
type ShipmentStatus string

// List of ShipmentStatus
const (
	ShipmentStatus_PickedUp       ShipmentStatus = "PickedUp"
	ShipmentStatus_ReadyForPickup ShipmentStatus = "ReadyForPickup"
	ShipmentStatus_RefusedPickup  ShipmentStatus = "RefusedPickup"
)

// TaxClassification defines model for TaxClassification.
type TaxClassification struct {

	// The type of tax.
	Name *string `json:"Name,omitempty"`

	// The buyer's tax identifier.
	Value *string `json:"Value,omitempty"`
}

// TaxCollection defines model for TaxCollection.
type TaxCollection struct {

	// The tax collection model applied to the item.
	Model *string `json:"Model,omitempty"`

	// The party responsible for withholding the taxes and remitting them to the taxing authority.
	ResponsibleParty *string `json:"ResponsibleParty,omitempty"`
}

// UpdateShipmentStatusErrorResponse defines model for UpdateShipmentStatusErrorResponse.
type UpdateShipmentStatusErrorResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// UpdateShipmentStatusRequest defines model for UpdateShipmentStatusRequest.
type UpdateShipmentStatusRequest struct {

	// the unobfuscated marketplace ID
	MarketplaceId MarketplaceId `json:"marketplaceId"`

	// the list of order items and quantities when the seller wants to partially update the shipment status of the order
	OrderItems *OrderItems `json:"orderItems,omitempty"`

	// the status of the shipment of the order to be updated
	ShipmentStatus ShipmentStatus `json:"shipmentStatus"`
}

// UpdateVerificationStatusErrorResponse defines model for UpdateVerificationStatusErrorResponse.
type UpdateVerificationStatusErrorResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`
}

// UpdateVerificationStatusRequest defines model for UpdateVerificationStatusRequest.
type UpdateVerificationStatusRequest struct {

	// The updated values of the VerificationStatus field.
	RegulatedOrderVerificationStatus UpdateVerificationStatusRequestBody `json:"regulatedOrderVerificationStatus"`
}

// UpdateVerificationStatusRequestBody defines model for UpdateVerificationStatusRequestBody.
type UpdateVerificationStatusRequestBody struct {

	// The identifier for the order's regulated information reviewer.
	ExternalReviewerId string `json:"externalReviewerId"`

	// The unique identifier for the rejection reason used for rejecting the order's regulated information. Only required if the new status is rejected.
	RejectionReasonId *string `json:"rejectionReasonId,omitempty"`

	// The verification status of the order.
	Status VerificationStatus `json:"status"`
}

// VerificationStatus defines model for VerificationStatus.
type VerificationStatus string

// List of VerificationStatus
const (
	VerificationStatus_Approved  VerificationStatus = "Approved"
	VerificationStatus_Cancelled VerificationStatus = "Cancelled"
	VerificationStatus_Expired   VerificationStatus = "Expired"
	VerificationStatus_Pending   VerificationStatus = "Pending"
	VerificationStatus_Rejected  VerificationStatus = "Rejected"
)

// GetOrdersParams defines parameters for GetOrders.
type GetOrdersParams struct {

	// A date used for selecting orders created after (or at) a specified time. Only orders placed after the specified time are returned. Either the CreatedAfter parameter or the LastUpdatedAfter parameter is required. Both cannot be empty. The date must be in ISO 8601 format.
	CreatedAfter *string `json:"CreatedAfter,omitempty"`

	// A date used for selecting orders created before (or at) a specified time. Only orders placed before the specified time are returned. The date must be in ISO 8601 format.
	CreatedBefore *string `json:"CreatedBefore,omitempty"`

	// A date used for selecting orders that were last updated after (or at) a specified time. An update is defined as any change in order status, including the creation of a new order. Includes updates made by Amazon and by the seller. The date must be in ISO 8601 format.
	LastUpdatedAfter *string `json:"LastUpdatedAfter,omitempty"`

	// A date used for selecting orders that were last updated before (or at) a specified time. An update is defined as any change in order status, including the creation of a new order. Includes updates made by Amazon and by the seller. The date must be in ISO 8601 format.
	LastUpdatedBefore *string `json:"LastUpdatedBefore,omitempty"`

	// A list of OrderStatus values used to filter the results. Possible values: PendingAvailability (This status is available for pre-orders only. The order has been placed, payment has not been authorized, and the release date of the item is in the future.); Pending (The order has been placed but payment has not been authorized); Unshipped (Payment has been authorized and the order is ready for shipment, but no items in the order have been shipped); PartiallyShipped (One or more, but not all, items in the order have been shipped); Shipped (All items in the order have been shipped); InvoiceUnconfirmed (All items in the order have been shipped. The seller has not yet given confirmation to Amazon that the invoice has been shipped to the buyer.); Canceled (The order has been canceled); and Unfulfillable (The order cannot be fulfilled. This state applies only to Multi-Channel Fulfillment orders.).
	OrderStatuses *[]string `json:"OrderStatuses,omitempty"`

	// A list of MarketplaceId values. Used to select orders that were placed in the specified marketplaces.
	//
	// See the [Selling Partner API Developer Guide](doc:marketplace-ids) for a complete list of marketplaceId values.
	MarketplaceIds []string `json:"MarketplaceIds"`

	// A list that indicates how an order was fulfilled. Filters the results by fulfillment channel. Possible values: AFN (Fulfillment by Amazon); MFN (Fulfilled by the seller).
	FulfillmentChannels *[]string `json:"FulfillmentChannels,omitempty"`

	// A list of payment method values. Used to select orders paid using the specified payment methods. Possible values: COD (Cash on delivery); CVS (Convenience store payment); Other (Any payment method other than COD or CVS).
	PaymentMethods *[]string `json:"PaymentMethods,omitempty"`

	// The email address of a buyer. Used to select orders that contain the specified email address.
	BuyerEmail *string `json:"BuyerEmail,omitempty"`

	// An order identifier that is specified by the seller. Used to select only the orders that match the order identifier. If SellerOrderId is specified, then FulfillmentChannels, OrderStatuses, PaymentMethod, LastUpdatedAfter, LastUpdatedBefore, and BuyerEmail cannot be specified.
	SellerOrderId *string `json:"SellerOrderId,omitempty"`

	// A number that indicates the maximum number of orders that can be returned per page. Value must be 1 - 100. Default 100.
	MaxResultsPerPage *int `json:"MaxResultsPerPage,omitempty"`

	// A list of EasyShipShipmentStatus values. Used to select Easy Ship orders with statuses that match the specified  values. If EasyShipShipmentStatus is specified, only Amazon Easy Ship orders are returned.Possible values: PendingPickUp (Amazon has not yet picked up the package from the seller). LabelCanceled (The seller canceled the pickup). PickedUp (Amazon has picked up the package from the seller). AtOriginFC (The packaged is at the origin fulfillment center). AtDestinationFC (The package is at the destination fulfillment center). OutForDelivery (The package is out for delivery). Damaged (The package was damaged by the carrier). Delivered (The package has been delivered to the buyer). RejectedByBuyer (The package has been rejected by the buyer). Undeliverable (The package cannot be delivered). ReturnedToSeller (The package was not delivered to the buyer and was returned to the seller). ReturningToSeller (The package was not delivered to the buyer and is being returned to the seller).
	EasyShipShipmentStatuses *[]string `json:"EasyShipShipmentStatuses,omitempty"`

	// A string token returned in the response of your previous request.
	NextToken *string `json:"NextToken,omitempty"`

	// A list of AmazonOrderId values. An AmazonOrderId is an Amazon-defined order identifier, in 3-7-7 format.
	AmazonOrderIds *[]string `json:"AmazonOrderIds,omitempty"`

	// Denotes the recommended sourceId where the order should be fulfilled from.
	ActualFulfillmentSupplySourceId *string `json:"ActualFulfillmentSupplySourceId,omitempty"`

	// When true, this order is marked to be picked up from a store rather than delivered.
	IsISPU *bool `json:"IsISPU,omitempty"`

	// The store chain store identifier. Linked to a specific store in a store chain.
	StoreChainStoreId *string `json:"StoreChainStoreId,omitempty"`
}

// GetOrderItemsParams defines parameters for GetOrderItems.
type GetOrderItemsParams struct {

	// A string token returned in the response of your previous request.
	NextToken *string `json:"NextToken,omitempty"`
}

// GetOrderItemsBuyerInfoParams defines parameters for GetOrderItemsBuyerInfo.
type GetOrderItemsBuyerInfoParams struct {

	// A string token returned in the response of your previous request.
	NextToken *string `json:"NextToken,omitempty"`
}

// UpdateVerificationStatusJSONBody defines parameters for UpdateVerificationStatus.
type UpdateVerificationStatusJSONBody UpdateVerificationStatusRequest

// UpdateShipmentStatusJSONBody defines parameters for UpdateShipmentStatus.
type UpdateShipmentStatusJSONBody UpdateShipmentStatusRequest

// UpdateVerificationStatusRequestBody defines body for UpdateVerificationStatus for application/json ContentType.
type UpdateVerificationStatusJSONRequestBody UpdateVerificationStatusJSONBody

// UpdateShipmentStatusRequestBody defines body for UpdateShipmentStatus for application/json ContentType.
type UpdateShipmentStatusJSONRequestBody UpdateShipmentStatusJSONBody
