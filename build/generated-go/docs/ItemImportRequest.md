# ItemImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Products** | [**[]Products**](Products.md) | Array of product strings | 
**UserAuth** | [**ItemImportRequestUserAuth**](ItemImportRequestUserAuth.md) |  | 
**Options** | Pointer to [**ItemImportRequestOptions**](ItemImportRequestOptions.md) |  | [optional] 

## Methods

### NewItemImportRequest

`func NewItemImportRequest(products []Products, userAuth ItemImportRequestUserAuth, ) *ItemImportRequest`

NewItemImportRequest instantiates a new ItemImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImportRequestWithDefaults

`func NewItemImportRequestWithDefaults() *ItemImportRequest`

NewItemImportRequestWithDefaults instantiates a new ItemImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ItemImportRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ItemImportRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ItemImportRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ItemImportRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ItemImportRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ItemImportRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ItemImportRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ItemImportRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetProducts

`func (o *ItemImportRequest) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ItemImportRequest) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ItemImportRequest) SetProducts(v []Products)`

SetProducts sets Products field to given value.


### GetUserAuth

`func (o *ItemImportRequest) GetUserAuth() ItemImportRequestUserAuth`

GetUserAuth returns the UserAuth field if non-nil, zero value otherwise.

### GetUserAuthOk

`func (o *ItemImportRequest) GetUserAuthOk() (*ItemImportRequestUserAuth, bool)`

GetUserAuthOk returns a tuple with the UserAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAuth

`func (o *ItemImportRequest) SetUserAuth(v ItemImportRequestUserAuth)`

SetUserAuth sets UserAuth field to given value.


### GetOptions

`func (o *ItemImportRequest) GetOptions() ItemImportRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ItemImportRequest) GetOptionsOk() (*ItemImportRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ItemImportRequest) SetOptions(v ItemImportRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ItemImportRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


