# AccountsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**Options** | Pointer to [**AccountsGetRequestOptions**](AccountsGetRequestOptions.md) |  | [optional] 

## Methods

### NewAccountsGetRequest

`func NewAccountsGetRequest(accessToken string, ) *AccountsGetRequest`

NewAccountsGetRequest instantiates a new AccountsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsGetRequestWithDefaults

`func NewAccountsGetRequestWithDefaults() *AccountsGetRequest`

NewAccountsGetRequestWithDefaults instantiates a new AccountsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AccountsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AccountsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AccountsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AccountsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AccountsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AccountsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AccountsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AccountsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *AccountsGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccountsGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccountsGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetOptions

`func (o *AccountsGetRequest) GetOptions() AccountsGetRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AccountsGetRequest) GetOptionsOk() (*AccountsGetRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AccountsGetRequest) SetOptions(v AccountsGetRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AccountsGetRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


