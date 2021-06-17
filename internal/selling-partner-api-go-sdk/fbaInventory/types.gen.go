// Package fbaInventory provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.7.1 DO NOT EDIT.
package fbaInventory

import (
	"time"
)

// Defines values for GranularityGranularityType.
const (
	GranularityGranularityTypeMarketplace GranularityGranularityType = "Marketplace"
)

// Defines values for ResearchingQuantityEntryName.
const (
	ResearchingQuantityEntryNameResearchingQuantityInLongTerm ResearchingQuantityEntryName = "researchingQuantityInLongTerm"

	ResearchingQuantityEntryNameResearchingQuantityInMidTerm ResearchingQuantityEntryName = "researchingQuantityInMidTerm"

	ResearchingQuantityEntryNameResearchingQuantityInShortTerm ResearchingQuantityEntryName = "researchingQuantityInShortTerm"
)

// An error response returned when the request is unsuccessful.
type Error struct {

	// An error code that identifies the type of error that occurred.
	Code string `json:"code"`

	// Additional information that can help the caller understand or fix the issue.
	Details *string `json:"details,omitempty"`

	// A message that describes the error condition in a human-readable form.
	Message *string `json:"message,omitempty"`
}

// A list of error responses returned when a request is unsuccessful.
type ErrorList []Error

// The Response schema.
type GetInventorySummariesResponse struct {

	// A list of error responses returned when a request is unsuccessful.
	Errors *ErrorList `json:"errors,omitempty"`

	// The process of returning the results to a request in batches of a defined size called pages. This is done to exercise some control over result size and overall throughput. It's a form of traffic management.
	Pagination *Pagination `json:"pagination,omitempty"`

	// The payload schema for the getInventorySummaries operation.
	Payload *GetInventorySummariesResult `json:"payload,omitempty"`
}

// The payload schema for the getInventorySummaries operation.
type GetInventorySummariesResult struct {

	// Describes a granularity at which inventory data can be aggregated. For example, if you use Marketplace granularity, the fulfillable quantity will reflect inventory that could be fulfilled in the given marketplace.
	Granularity Granularity `json:"granularity"`

	// A list of inventory summaries.
	InventorySummaries InventorySummaries `json:"inventorySummaries"`
}

// Describes a granularity at which inventory data can be aggregated. For example, if you use Marketplace granularity, the fulfillable quantity will reflect inventory that could be fulfilled in the given marketplace.
type Granularity struct {

	// The granularity ID for the specified granularity type. When granularityType is Marketplace, specify the marketplaceId.
	GranularityId *string `json:"granularityId,omitempty"`

	// The granularity type for the inventory aggregation level.
	GranularityType *GranularityGranularityType `json:"granularityType,omitempty"`
}

// The granularity type for the inventory aggregation level.
type GranularityGranularityType string

// Summarized inventory details. This object will not appear if the details parameter in the request is false.
type InventoryDetails struct {

	// The item quantity that can be picked, packed, and shipped.
	FulfillableQuantity *int `json:"fulfillableQuantity,omitempty"`

	// The number of units that have not yet been received at an Amazon fulfillment center for processing, but are part of an inbound shipment with some units that have already been received and processed.
	InboundReceivingQuantity *int `json:"inboundReceivingQuantity,omitempty"`

	// The number of units in an inbound shipment that you have notified Amazon about and have provided a tracking number.
	InboundShippedQuantity *int `json:"inboundShippedQuantity,omitempty"`

	// The number of units in an inbound shipment for which you have notified Amazon.
	InboundWorkingQuantity *int `json:"inboundWorkingQuantity,omitempty"`

	// The number of misplaced or warehouse damaged units that are actively being confirmed at our fulfillment centers.
	ResearchingQuantity *ResearchingQuantity `json:"researchingQuantity,omitempty"`

	// The quantity of reserved inventory.
	ReservedQuantity *ReservedQuantity `json:"reservedQuantity,omitempty"`

	// The quantity of unfulfillable inventory.
	UnfulfillableQuantity *UnfulfillableQuantity `json:"unfulfillableQuantity,omitempty"`
}

// A list of inventory summaries.
type InventorySummaries []InventorySummary

// Inventory summary for a specific item.
type InventorySummary struct {

	// The Amazon Standard Identification Number (ASIN) of an item.
	Asin *string `json:"asin,omitempty"`

	// The condition of the item as described by the seller (for example, New Item).
	Condition *string `json:"condition,omitempty"`

	// Amazon's fulfillment network SKU identifier.
	FnSku *string `json:"fnSku,omitempty"`

	// Summarized inventory details. This object will not appear if the details parameter in the request is false.
	InventoryDetails *InventoryDetails `json:"inventoryDetails,omitempty"`

	// The date and time that any quantity was last updated.
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty"`

	// The localized language product title of the item within the specific marketplace.
	ProductName *string `json:"productName,omitempty"`

	// The seller SKU of the item.
	SellerSku *string `json:"sellerSku,omitempty"`

	// The total number of units in an inbound shipment or in Amazon fulfillment centers.
	TotalQuantity *int `json:"totalQuantity,omitempty"`
}

// The process of returning the results to a request in batches of a defined size called pages. This is done to exercise some control over result size and overall throughput. It's a form of traffic management.
type Pagination struct {

	// A generated string used to retrieve the next page of the result. If nextToken is returned, pass the value of nextToken to the next request. If nextToken is not returned, there are no more items to return.
	NextToken *string `json:"nextToken,omitempty"`
}

// The number of misplaced or warehouse damaged units that are actively being confirmed at our fulfillment centers.
type ResearchingQuantity struct {

	// A list of quantity details for items currently being researched.
	ResearchingQuantityBreakdown *[]ResearchingQuantityEntry `json:"researchingQuantityBreakdown,omitempty"`

	// The total number of units currently being researched in Amazon's fulfillment network.
	TotalResearchingQuantity *int `json:"totalResearchingQuantity,omitempty"`
}

// The misplaced or warehouse damaged inventory that is actively being confirmed at our fulfillment centers.
type ResearchingQuantityEntry struct {

	// The duration of the research.
	Name ResearchingQuantityEntryName `json:"name"`

	// The number of units.
	Quantity int `json:"quantity"`
}

// The duration of the research.
type ResearchingQuantityEntryName string

// The quantity of reserved inventory.
type ReservedQuantity struct {

	// The number of units that have been sidelined at the fulfillment center for additional processing.
	FcProcessingQuantity *int `json:"fcProcessingQuantity,omitempty"`

	// The number of units reserved for customer orders.
	PendingCustomerOrderQuantity *int `json:"pendingCustomerOrderQuantity,omitempty"`

	// The number of units being transferred from one fulfillment center to another.
	PendingTransshipmentQuantity *int `json:"pendingTransshipmentQuantity,omitempty"`

	// The total number of units in Amazon's fulfillment network that are currently being picked, packed, and shipped; or are sidelined for measurement, sampling, or other internal processes.
	TotalReservedQuantity *int `json:"totalReservedQuantity,omitempty"`
}

// The quantity of unfulfillable inventory.
type UnfulfillableQuantity struct {

	// The number of units in carrier damaged disposition.
	CarrierDamagedQuantity *int `json:"carrierDamagedQuantity,omitempty"`

	// The number of units in customer damaged disposition.
	CustomerDamagedQuantity *int `json:"customerDamagedQuantity,omitempty"`

	// The number of units in defective disposition.
	DefectiveQuantity *int `json:"defectiveQuantity,omitempty"`

	// The number of units in distributor damaged disposition.
	DistributorDamagedQuantity *int `json:"distributorDamagedQuantity,omitempty"`

	// The number of units in expired disposition.
	ExpiredQuantity *int `json:"expiredQuantity,omitempty"`

	// The total number of units in Amazon's fulfillment network in unsellable condition.
	TotalUnfulfillableQuantity *int `json:"totalUnfulfillableQuantity,omitempty"`

	// The number of units in warehouse damaged disposition.
	WarehouseDamagedQuantity *int `json:"warehouseDamagedQuantity,omitempty"`
}

// GetInventorySummariesParams defines parameters for GetInventorySummaries.
type GetInventorySummariesParams struct {

	// true to return inventory summaries with additional summarized inventory details and quantities. Otherwise, returns inventory summaries only (default value).
	Details *bool `json:"details,omitempty"`

	// The granularity type for the inventory aggregation level.
	GranularityType GetInventorySummariesParamsGranularityType `json:"granularityType"`

	// The granularity ID for the inventory aggregation level.
	GranularityId string `json:"granularityId"`

	// A start date and time in ISO8601 format. If specified, all inventory summaries that have changed since then are returned. You must specify a date and time that is no earlier than 18 months prior to the date and time when you call the API. Note: Changes in inboundWorkingQuantity, inboundShippedQuantity and inboundReceivingQuantity are not detected.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`

	// A list of seller SKUs for which to return inventory summaries. You may specify up to 50 SKUs.
	SellerSkus *[]string `json:"sellerSkus,omitempty"`

	// String token returned in the response of your previous request.
	NextToken *string `json:"nextToken,omitempty"`

	// The marketplace ID for the marketplace for which to return inventory summaries.
	MarketplaceIds []string `json:"marketplaceIds"`
}

// GetInventorySummariesParamsGranularityType defines parameters for GetInventorySummaries.
type GetInventorySummariesParamsGranularityType string
