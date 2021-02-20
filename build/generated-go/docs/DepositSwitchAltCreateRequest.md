# DepositSwitchAltCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**TargetAccount** | [**DepositSwitchTargetAccount**](DepositSwitchTargetAccount.md) |  | 
**TargetUser** | [**DepositSwitchTargetUser**](DepositSwitchTargetUser.md) |  | 

## Methods

### NewDepositSwitchAltCreateRequest

`func NewDepositSwitchAltCreateRequest(targetAccount DepositSwitchTargetAccount, targetUser DepositSwitchTargetUser, ) *DepositSwitchAltCreateRequest`

NewDepositSwitchAltCreateRequest instantiates a new DepositSwitchAltCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchAltCreateRequestWithDefaults

`func NewDepositSwitchAltCreateRequestWithDefaults() *DepositSwitchAltCreateRequest`

NewDepositSwitchAltCreateRequestWithDefaults instantiates a new DepositSwitchAltCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchAltCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchAltCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchAltCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchAltCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchAltCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchAltCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchAltCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchAltCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetTargetAccount

`func (o *DepositSwitchAltCreateRequest) GetTargetAccount() DepositSwitchTargetAccount`

GetTargetAccount returns the TargetAccount field if non-nil, zero value otherwise.

### GetTargetAccountOk

`func (o *DepositSwitchAltCreateRequest) GetTargetAccountOk() (*DepositSwitchTargetAccount, bool)`

GetTargetAccountOk returns a tuple with the TargetAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccount

`func (o *DepositSwitchAltCreateRequest) SetTargetAccount(v DepositSwitchTargetAccount)`

SetTargetAccount sets TargetAccount field to given value.


### GetTargetUser

`func (o *DepositSwitchAltCreateRequest) GetTargetUser() DepositSwitchTargetUser`

GetTargetUser returns the TargetUser field if non-nil, zero value otherwise.

### GetTargetUserOk

`func (o *DepositSwitchAltCreateRequest) GetTargetUserOk() (*DepositSwitchTargetUser, bool)`

GetTargetUserOk returns a tuple with the TargetUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUser

`func (o *DepositSwitchAltCreateRequest) SetTargetUser(v DepositSwitchTargetUser)`

SetTargetUser sets TargetUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


