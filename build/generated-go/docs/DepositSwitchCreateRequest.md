# DepositSwitchCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**TargetAccessToken** | **string** | Access token for the target Item, typically provided in the Import Item response.  | 
**TargetAccountId** | **string** | Plaid Account ID that specifies the target bank account. This account will become the recipient for a user&#39;s direct deposit. | 

## Methods

### NewDepositSwitchCreateRequest

`func NewDepositSwitchCreateRequest(targetAccessToken string, targetAccountId string, ) *DepositSwitchCreateRequest`

NewDepositSwitchCreateRequest instantiates a new DepositSwitchCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchCreateRequestWithDefaults

`func NewDepositSwitchCreateRequestWithDefaults() *DepositSwitchCreateRequest`

NewDepositSwitchCreateRequestWithDefaults instantiates a new DepositSwitchCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTargetAccessToken

`func (o *DepositSwitchCreateRequest) GetTargetAccessToken() string`

GetTargetAccessToken returns the TargetAccessToken field if non-nil, zero value otherwise.

### GetTargetAccessTokenOk

`func (o *DepositSwitchCreateRequest) GetTargetAccessTokenOk() (*string, bool)`

GetTargetAccessTokenOk returns a tuple with the TargetAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccessToken

`func (o *DepositSwitchCreateRequest) SetTargetAccessToken(v string)`

SetTargetAccessToken sets TargetAccessToken field to given value.


### GetTargetAccountId

`func (o *DepositSwitchCreateRequest) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *DepositSwitchCreateRequest) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *DepositSwitchCreateRequest) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


