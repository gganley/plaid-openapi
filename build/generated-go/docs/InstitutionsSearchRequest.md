# InstitutionsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Query** | **string** | The search query. Institutions with names matching the query are returned | 
**Products** | [**[]Products**](Products.md) | Filter the Institutions based on whether they support all products listed in products. Provide &#x60;null&#x60; to get institutions regardless of supported products | 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard.  | 
**Options** | Pointer to [**InstitutionsSearchRequestOptions**](InstitutionsSearchRequestOptions.md) |  | [optional] 

## Methods

### NewInstitutionsSearchRequest

`func NewInstitutionsSearchRequest(query string, products []Products, countryCodes []CountryCode, ) *InstitutionsSearchRequest`

NewInstitutionsSearchRequest instantiates a new InstitutionsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsSearchRequestWithDefaults

`func NewInstitutionsSearchRequestWithDefaults() *InstitutionsSearchRequest`

NewInstitutionsSearchRequestWithDefaults instantiates a new InstitutionsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InstitutionsSearchRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InstitutionsSearchRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InstitutionsSearchRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InstitutionsSearchRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *InstitutionsSearchRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InstitutionsSearchRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InstitutionsSearchRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InstitutionsSearchRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetQuery

`func (o *InstitutionsSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *InstitutionsSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *InstitutionsSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProducts

`func (o *InstitutionsSearchRequest) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InstitutionsSearchRequest) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InstitutionsSearchRequest) SetProducts(v []Products)`

SetProducts sets Products field to given value.


### GetCountryCodes

`func (o *InstitutionsSearchRequest) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *InstitutionsSearchRequest) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *InstitutionsSearchRequest) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.


### GetOptions

`func (o *InstitutionsSearchRequest) GetOptions() InstitutionsSearchRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstitutionsSearchRequest) GetOptionsOk() (*InstitutionsSearchRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstitutionsSearchRequest) SetOptions(v InstitutionsSearchRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstitutionsSearchRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


