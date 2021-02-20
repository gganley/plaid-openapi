# InstitutionsGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]Products**](Products.md) | Filter the Institutions based on which products they support.  | [optional] 
**RoutingNumbers** | Pointer to **[]string** | Specify an array of routing numbers to filter institutions. | [optional] 
**Oauth** | Pointer to **bool** | Limit results to institutions with or without OAuth login flows. This is primarily relevant to institutions with European country codes. | [optional] 
**IncludeOptionalMetadata** | Pointer to **bool** | When &#x60;true&#x60;, return the institution&#39;s homepage URL, logo and primary brand color.  Note that Plaid does not own any of the logos shared by the API, and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos. | [optional] 

## Methods

### NewInstitutionsGetRequestOptions

`func NewInstitutionsGetRequestOptions() *InstitutionsGetRequestOptions`

NewInstitutionsGetRequestOptions instantiates a new InstitutionsGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetRequestOptionsWithDefaults

`func NewInstitutionsGetRequestOptionsWithDefaults() *InstitutionsGetRequestOptions`

NewInstitutionsGetRequestOptionsWithDefaults instantiates a new InstitutionsGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InstitutionsGetRequestOptions) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InstitutionsGetRequestOptions) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InstitutionsGetRequestOptions) SetProducts(v []Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InstitutionsGetRequestOptions) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetRoutingNumbers

`func (o *InstitutionsGetRequestOptions) GetRoutingNumbers() []string`

GetRoutingNumbers returns the RoutingNumbers field if non-nil, zero value otherwise.

### GetRoutingNumbersOk

`func (o *InstitutionsGetRequestOptions) GetRoutingNumbersOk() (*[]string, bool)`

GetRoutingNumbersOk returns a tuple with the RoutingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumbers

`func (o *InstitutionsGetRequestOptions) SetRoutingNumbers(v []string)`

SetRoutingNumbers sets RoutingNumbers field to given value.

### HasRoutingNumbers

`func (o *InstitutionsGetRequestOptions) HasRoutingNumbers() bool`

HasRoutingNumbers returns a boolean if a field has been set.

### GetOauth

`func (o *InstitutionsGetRequestOptions) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InstitutionsGetRequestOptions) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InstitutionsGetRequestOptions) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InstitutionsGetRequestOptions) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadata() bool`

GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field if non-nil, zero value otherwise.

### GetIncludeOptionalMetadataOk

`func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool)`

GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) SetIncludeOptionalMetadata(v bool)`

SetIncludeOptionalMetadata sets IncludeOptionalMetadata field to given value.

### HasIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) HasIncludeOptionalMetadata() bool`

HasIncludeOptionalMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


