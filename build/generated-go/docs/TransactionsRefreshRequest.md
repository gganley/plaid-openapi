# TransactionsRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**AccessToken** | Pointer to **string** | The access token associated with the Item data is being requested for. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 

## Methods

### NewTransactionsRefreshRequest

`func NewTransactionsRefreshRequest() *TransactionsRefreshRequest`

NewTransactionsRefreshRequest instantiates a new TransactionsRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsRefreshRequestWithDefaults

`func NewTransactionsRefreshRequestWithDefaults() *TransactionsRefreshRequest`

NewTransactionsRefreshRequestWithDefaults instantiates a new TransactionsRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransactionsRefreshRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransactionsRefreshRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransactionsRefreshRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransactionsRefreshRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransactionsRefreshRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransactionsRefreshRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransactionsRefreshRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *TransactionsRefreshRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetSecret

`func (o *TransactionsRefreshRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransactionsRefreshRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransactionsRefreshRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransactionsRefreshRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


