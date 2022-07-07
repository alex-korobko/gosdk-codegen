// Package catalog provides primitives to interact the openapi HTTP API.
//
// Code generated by go-sdk-codegen DO NOT EDIT.
package catalog

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// BrandRefinement defines model for BrandRefinement.
type BrandRefinement struct {

	// Brand name. For display and can be used as a search refinement.
	BrandName string `json:"brandName"`

	// The estimated number of results that would still be returned if refinement key applied.
	NumberOfResults int `json:"numberOfResults"`
}

// ClassificationRefinement defines model for ClassificationRefinement.
type ClassificationRefinement struct {

	// Identifier for the classification that can be used for search refinement purposes.
	ClassificationId string `json:"classificationId"`

	// Display name for the classification.
	DisplayName string `json:"displayName"`

	// The estimated number of results that would still be returned if refinement key applied.
	NumberOfResults int `json:"numberOfResults"`
}

// Dimension defines model for Dimension.
type Dimension struct {

	// Measurement unit of the dimension value.
	Unit *string `json:"unit,omitempty"`

	// Numeric dimension value.
	Value *float32 `json:"value,omitempty"`
}

// Dimensions defines model for Dimensions.
type Dimensions struct {

	// Individual dimension value of an Amazon catalog item or item package.
	Height *Dimension `json:"height,omitempty"`

	// Individual dimension value of an Amazon catalog item or item package.
	Length *Dimension `json:"length,omitempty"`

	// Individual dimension value of an Amazon catalog item or item package.
	Weight *Dimension `json:"weight,omitempty"`

	// Individual dimension value of an Amazon catalog item or item package.
	Width *Dimension `json:"width,omitempty"`
}

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

// Item defines model for Item.
type Item struct {

	// Amazon Standard Identification Number (ASIN) is the unique identifier for an item in the Amazon catalog.
	Asin ItemAsin `json:"asin"`

	// A JSON object that contains structured item attribute data keyed by attribute name. Catalog item attributes conform to the related product type definitions available in the Selling Partner API for Product Type Definitions.
	Attributes *ItemAttributes `json:"attributes,omitempty"`

	// Array of dimensions associated with the item in the Amazon catalog by Amazon marketplace.
	Dimensions *ItemDimensions `json:"dimensions,omitempty"`

	// Identifiers associated with the item in the Amazon catalog, such as UPC and EAN identifiers.
	Identifiers *ItemIdentifiers `json:"identifiers,omitempty"`

	// Images for an item in the Amazon catalog.
	Images *ItemImages `json:"images,omitempty"`

	// Product types associated with the Amazon catalog item.
	ProductTypes *ItemProductTypes `json:"productTypes,omitempty"`

	// Relationships by marketplace for an Amazon catalog item (for example, variations).
	Relationships *ItemRelationships `json:"relationships,omitempty"`

	// Sales ranks of an Amazon catalog item.
	SalesRanks *ItemSalesRanks `json:"salesRanks,omitempty"`

	// Summary details of an Amazon catalog item.
	Summaries *ItemSummaries `json:"summaries,omitempty"`

	// Vendor details associated with an Amazon catalog item. Vendor details are available to vendors only.
	VendorDetails *ItemVendorDetails `json:"vendorDetails,omitempty"`
}

// ItemAsin defines model for ItemAsin.
type ItemAsin string

// ItemAttributes defines model for ItemAttributes.
type ItemAttributes struct {
	AdditionalProperties map[string]interface{} `json:"-"`
}

// ItemBrowseClassification defines model for ItemBrowseClassification.
type ItemBrowseClassification struct {

	// Identifier of the classification (browse node identifier).
	ClassificationId string `json:"classificationId"`

	// Display name for the classification.
	DisplayName string `json:"displayName"`
}

// ItemClassificationSalesRank defines model for ItemClassificationSalesRank.
type ItemClassificationSalesRank struct {

	// Identifier of the classification associated with the sales rank.
	ClassificationId string `json:"classificationId"`

	// Corresponding Amazon retail website link, or URL, for the sales rank.
	Link *string `json:"link,omitempty"`

	// Sales rank value.
	Rank int `json:"rank"`

	// Title, or name, of the sales rank.
	Title string `json:"title"`
}

// ItemDimensions defines model for ItemDimensions.
type ItemDimensions []ItemDimensionsByMarketplace

// ItemDimensionsByMarketplace defines model for ItemDimensionsByMarketplace.
type ItemDimensionsByMarketplace struct {

	// Dimensions of an Amazon catalog item or item in its packaging.
	Item *Dimensions `json:"item,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Dimensions of an Amazon catalog item or item in its packaging.
	Package *Dimensions `json:"package,omitempty"`
}

// ItemDisplayGroupSalesRank defines model for ItemDisplayGroupSalesRank.
type ItemDisplayGroupSalesRank struct {

	// Corresponding Amazon retail website link, or URL, for the sales rank.
	Link *string `json:"link,omitempty"`

	// Sales rank value.
	Rank int `json:"rank"`

	// Title, or name, of the sales rank.
	Title string `json:"title"`

	// Name of the website display group associated with the sales rank
	WebsiteDisplayGroup string `json:"websiteDisplayGroup"`
}

// ItemIdentifier defines model for ItemIdentifier.
type ItemIdentifier struct {

	// Identifier.
	Identifier string `json:"identifier"`

	// Type of identifier, such as UPC, EAN, or ISBN.
	IdentifierType string `json:"identifierType"`
}

// ItemIdentifiers defines model for ItemIdentifiers.
type ItemIdentifiers []ItemIdentifiersByMarketplace

// ItemIdentifiersByMarketplace defines model for ItemIdentifiersByMarketplace.
type ItemIdentifiersByMarketplace struct {

	// Identifiers associated with the item in the Amazon catalog for the indicated Amazon marketplace.
	Identifiers []ItemIdentifier `json:"identifiers"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

// ItemImage defines model for ItemImage.
type ItemImage struct {

	// Height of the image in pixels.
	Height int `json:"height"`

	// Link, or URL, for the image.
	Link string `json:"link"`

	// Variant of the image, such as `MAIN` or `PT01`.
	Variant string `json:"variant"`

	// Width of the image in pixels.
	Width int `json:"width"`
}

// ItemImages defines model for ItemImages.
type ItemImages []ItemImagesByMarketplace

// ItemImagesByMarketplace defines model for ItemImagesByMarketplace.
type ItemImagesByMarketplace struct {

	// Images for an item in the Amazon catalog for the indicated Amazon marketplace.
	Images []ItemImage `json:"images"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

// ItemProductTypeByMarketplace defines model for ItemProductTypeByMarketplace.
type ItemProductTypeByMarketplace struct {

	// Amazon marketplace identifier.
	MarketplaceId *string `json:"marketplaceId,omitempty"`

	// Name of the product type associated with the Amazon catalog item.
	ProductType *string `json:"productType,omitempty"`
}

// ItemProductTypes defines model for ItemProductTypes.
type ItemProductTypes []ItemProductTypeByMarketplace

// ItemRelationship defines model for ItemRelationship.
type ItemRelationship struct {

	// Identifiers (ASINs) of the related items that are children of this item.
	ChildAsins *[]string `json:"childAsins,omitempty"`

	// Identifiers (ASINs) of the related items that are parents of this item.
	ParentAsins *[]string `json:"parentAsins,omitempty"`

	// Type of relationship.
	Type string `json:"type"`

	// Variation theme indicating the combination of Amazon item catalog attributes that define the variation family.
	VariationTheme *ItemVariationTheme `json:"variationTheme,omitempty"`
}

// ItemRelationships defines model for ItemRelationships.
type ItemRelationships []ItemRelationshipsByMarketplace

// ItemRelationshipsByMarketplace defines model for ItemRelationshipsByMarketplace.
type ItemRelationshipsByMarketplace struct {

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Relationships for the item.
	Relationships []ItemRelationship `json:"relationships"`
}

// ItemSalesRanks defines model for ItemSalesRanks.
type ItemSalesRanks []ItemSalesRanksByMarketplace

// ItemSalesRanksByMarketplace defines model for ItemSalesRanksByMarketplace.
type ItemSalesRanksByMarketplace struct {

	// Sales ranks of an Amazon catalog item for an Amazon marketplace by classification.
	ClassificationRanks *[]ItemClassificationSalesRank `json:"classificationRanks,omitempty"`

	// Sales ranks of an Amazon catalog item for an Amazon marketplace by website display group.
	DisplayGroupRanks *[]ItemDisplayGroupSalesRank `json:"displayGroupRanks,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`
}

// ItemSearchResults defines model for ItemSearchResults.
type ItemSearchResults struct {

	// A list of items from the Amazon catalog.
	Items []Item `json:"items"`

	// For `identifiers`-based searches, the total number of Amazon catalog items found. For `keywords`-based searches, the estimated total number of Amazon catalog items matched by the search query (only results up to the page count limit will be returned per request regardless of the number found).
	//
	// Note: The maximum number of items (ASINs) that can be returned and paged through is 1000.
	NumberOfResults int `json:"numberOfResults"`

	// When a request produces a response that exceeds the `pageSize`, pagination occurs. This means the response is divided into individual pages. To retrieve the next page or the previous page, you must pass the `nextToken` value or the `previousToken` value as the `pageToken` parameter in the next request. When you receive the last page, there will be no `nextToken` key in the pagination object.
	Pagination Pagination `json:"pagination"`

	// Search refinements.
	Refinements Refinements `json:"refinements"`
}

// ItemSummaries defines model for ItemSummaries.
type ItemSummaries []ItemSummaryByMarketplace

// ItemSummaryByMarketplace defines model for ItemSummaryByMarketplace.
type ItemSummaryByMarketplace struct {

	// Name of the brand associated with an Amazon catalog item.
	Brand *string `json:"brand,omitempty"`

	// Classification (browse node) associated with an Amazon catalog item.
	BrowseClassification *ItemBrowseClassification `json:"browseClassification,omitempty"`

	// Name of the color associated with an Amazon catalog item.
	Color *string `json:"color,omitempty"`

	// Classification type associated with the Amazon catalog item.
	ItemClassification *string `json:"itemClassification,omitempty"`

	// Name, or title, associated with an Amazon catalog item.
	ItemName *string `json:"itemName,omitempty"`

	// Name of the manufacturer associated with an Amazon catalog item.
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Model number associated with an Amazon catalog item.
	ModelNumber *string `json:"modelNumber,omitempty"`

	// Quantity of an Amazon catalog item in one package.
	PackageQuantity *int `json:"packageQuantity,omitempty"`

	// Part number associated with an Amazon catalog item.
	PartNumber *string `json:"partNumber,omitempty"`

	// Name of the size associated with an Amazon catalog item.
	Size *string `json:"size,omitempty"`

	// Name of the style associated with an Amazon catalog item.
	Style *string `json:"style,omitempty"`

	// Name of the website display group associated with an Amazon catalog item.
	WebsiteDisplayGroup *string `json:"websiteDisplayGroup,omitempty"`
}

// ItemVariationTheme defines model for ItemVariationTheme.
type ItemVariationTheme struct {

	// Names of the Amazon catalog item attributes associated with the variation theme.
	Attributes *[]string `json:"attributes,omitempty"`

	// Variation theme indicating the combination of Amazon item catalog attributes that define the variation family.
	Theme *string `json:"theme,omitempty"`
}

// ItemVendorDetails defines model for ItemVendorDetails.
type ItemVendorDetails []ItemVendorDetailsByMarketplace

// ItemVendorDetailsByMarketplace defines model for ItemVendorDetailsByMarketplace.
type ItemVendorDetailsByMarketplace struct {

	// Brand code associated with an Amazon catalog item.
	BrandCode *string `json:"brandCode,omitempty"`

	// Manufacturer code associated with an Amazon catalog item.
	ManufacturerCode *string `json:"manufacturerCode,omitempty"`

	// Parent vendor code of the manufacturer code.
	ManufacturerCodeParent *string `json:"manufacturerCodeParent,omitempty"`

	// Amazon marketplace identifier.
	MarketplaceId string `json:"marketplaceId"`

	// Product category or subcategory associated with an Amazon catalog item.
	ProductCategory *ItemVendorDetailsCategory `json:"productCategory,omitempty"`

	// Product group associated with an Amazon catalog item.
	ProductGroup *string `json:"productGroup,omitempty"`

	// Product category or subcategory associated with an Amazon catalog item.
	ProductSubcategory *ItemVendorDetailsCategory `json:"productSubcategory,omitempty"`

	// Replenishment category associated with an Amazon catalog item.
	ReplenishmentCategory *string `json:"replenishmentCategory,omitempty"`
}

// ItemVendorDetailsCategory defines model for ItemVendorDetailsCategory.
type ItemVendorDetailsCategory struct {

	// Display name of the product category or subcategory
	DisplayName *string `json:"displayName,omitempty"`

	// Value (code) of the product category or subcategory.
	Value *string `json:"value,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {

	// A token that can be used to fetch the next page.
	NextToken *string `json:"nextToken,omitempty"`

	// A token that can be used to fetch the previous page.
	PreviousToken *string `json:"previousToken,omitempty"`
}

// Refinements defines model for Refinements.
type Refinements struct {

	// Brand search refinements.
	Brands []BrandRefinement `json:"brands"`

	// Classification search refinements.
	Classifications []ClassificationRefinement `json:"classifications"`
}

// SearchCatalogItemsParams defines parameters for SearchCatalogItems.
type SearchCatalogItemsParams struct {

	// A comma-delimited list of product identifiers to search the Amazon catalog for. **Note:** Cannot be used with `keywords`.
	Identifiers *[]string `json:"identifiers,omitempty"`

	// Type of product identifiers to search the Amazon catalog for. **Note:** Required when `identifiers` are provided.
	IdentifiersType *string `json:"identifiersType,omitempty"`

	// A comma-delimited list of Amazon marketplace identifiers for the request.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: `summaries`.
	IncludedData *[]string `json:"includedData,omitempty"`

	// Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace.
	Locale *string `json:"locale,omitempty"`

	// A selling partner identifier, such as a seller account or vendor code. **Note:** Required when `identifiersType` is `SKU`.
	SellerId *string `json:"sellerId,omitempty"`

	// A comma-delimited list of words to search the Amazon catalog for. **Note:** Cannot be used with `identifiers`.
	Keywords *[]string `json:"keywords,omitempty"`

	// A comma-delimited list of brand names to limit the search for `keywords`-based queries. **Note:** Cannot be used with `identifiers`.
	BrandNames *[]string `json:"brandNames,omitempty"`

	// A comma-delimited list of classification identifiers to limit the search for `keywords`-based queries. **Note:** Cannot be used with `identifiers`.
	ClassificationIds *[]string `json:"classificationIds,omitempty"`

	// Number of results to be returned per page.
	PageSize *int `json:"pageSize,omitempty"`

	// A token to fetch a certain page when there are multiple pages worth of results.
	PageToken *string `json:"pageToken,omitempty"`

	// The language of the keywords provided for `keywords`-based queries. Defaults to the primary locale of the marketplace. **Note:** Cannot be used with `identifiers`.
	KeywordsLocale *string `json:"keywordsLocale,omitempty"`
}

// GetCatalogItemParams defines parameters for GetCatalogItem.
type GetCatalogItemParams struct {

	// A comma-delimited list of Amazon marketplace identifiers. Data sets in the response contain data only for the specified marketplaces.
	MarketplaceIds []string `json:"marketplaceIds"`

	// A comma-delimited list of data sets to include in the response. Default: `summaries`.
	IncludedData *[]string `json:"includedData,omitempty"`

	// Locale for retrieving localized summaries. Defaults to the primary locale of the marketplace.
	Locale *string `json:"locale,omitempty"`
}

// Getter for additional properties for ItemAttributes. Returns the specified
// element and whether it was found
func (a ItemAttributes) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ItemAttributes
func (a *ItemAttributes) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a *ItemAttributes) UnmarshalJSON(b []byte) error {
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

// Override default JSON handling for ItemAttributes to handle AdditionalProperties
func (a ItemAttributes) MarshalJSON() ([]byte, error) {
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
