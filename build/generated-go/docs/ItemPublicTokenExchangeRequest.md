# ItemPublicTokenExchangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**PublicToken** | **string** | Your &#x60;public_token&#x60;, obtained from the Link &#x60;onSuccess&#x60; callback or &#x60;/sandbox/item/public_token/create&#x60;. | 

## Methods

### NewItemPublicTokenExchangeRequest

`func NewItemPublicTokenExchangeRequest(publicToken string, ) *ItemPublicTokenExchangeRequest`

NewItemPublicTokenExchangeRequest instantiates a new ItemPublicTokenExchangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPublicTokenExchangeRequestWithDefaults

`func NewItemPublicTokenExchangeRequestWithDefaults() *ItemPublicTokenExchangeRequest`

NewItemPublicTokenExchangeRequestWithDefaults instantiates a new ItemPublicTokenExchangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ItemPublicTokenExchangeRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ItemPublicTokenExchangeRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ItemPublicTokenExchangeRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ItemPublicTokenExchangeRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ItemPublicTokenExchangeRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ItemPublicTokenExchangeRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ItemPublicTokenExchangeRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ItemPublicTokenExchangeRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPublicToken

`func (o *ItemPublicTokenExchangeRequest) GetPublicToken() string`

GetPublicToken returns the PublicToken field if non-nil, zero value otherwise.

### GetPublicTokenOk

`func (o *ItemPublicTokenExchangeRequest) GetPublicTokenOk() (*string, bool)`

GetPublicTokenOk returns a tuple with the PublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicToken

`func (o *ItemPublicTokenExchangeRequest) SetPublicToken(v string)`

SetPublicToken sets PublicToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


