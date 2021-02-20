# LinkTokenGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**LinkToken** | **string** | A &#x60;link_token&#x60; from a previous invocation of &#x60;/link/token/create&#x60; | 

## Methods

### NewLinkTokenGetRequest

`func NewLinkTokenGetRequest(linkToken string, ) *LinkTokenGetRequest`

NewLinkTokenGetRequest instantiates a new LinkTokenGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenGetRequestWithDefaults

`func NewLinkTokenGetRequestWithDefaults() *LinkTokenGetRequest`

NewLinkTokenGetRequestWithDefaults instantiates a new LinkTokenGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *LinkTokenGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *LinkTokenGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *LinkTokenGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *LinkTokenGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *LinkTokenGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *LinkTokenGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *LinkTokenGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *LinkTokenGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetLinkToken

`func (o *LinkTokenGetRequest) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *LinkTokenGetRequest) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *LinkTokenGetRequest) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


