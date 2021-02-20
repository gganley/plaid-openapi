# AccountsBalanceGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Options** | Pointer to [**AccountsBalanceGetRequestOptions**](AccountsBalanceGetRequestOptions.md) |  | [optional] 

## Methods

### NewAccountsBalanceGetRequest

`func NewAccountsBalanceGetRequest(accessToken string, ) *AccountsBalanceGetRequest`

NewAccountsBalanceGetRequest instantiates a new AccountsBalanceGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsBalanceGetRequestWithDefaults

`func NewAccountsBalanceGetRequestWithDefaults() *AccountsBalanceGetRequest`

NewAccountsBalanceGetRequestWithDefaults instantiates a new AccountsBalanceGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccountsBalanceGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccountsBalanceGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccountsBalanceGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetSecret

`func (o *AccountsBalanceGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AccountsBalanceGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AccountsBalanceGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AccountsBalanceGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetClientId

`func (o *AccountsBalanceGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountsBalanceGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountsBalanceGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AccountsBalanceGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOptions

`func (o *AccountsBalanceGetRequest) GetOptions() AccountsBalanceGetRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AccountsBalanceGetRequest) GetOptionsOk() (*AccountsBalanceGetRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AccountsBalanceGetRequest) SetOptions(v AccountsBalanceGetRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AccountsBalanceGetRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


