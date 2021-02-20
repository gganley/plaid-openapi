# DepositSwitchGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**DepositSwitchId** | **string** | The ID of the deposit switch | 

## Methods

### NewDepositSwitchGetRequest

`func NewDepositSwitchGetRequest(depositSwitchId string, ) *DepositSwitchGetRequest`

NewDepositSwitchGetRequest instantiates a new DepositSwitchGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchGetRequestWithDefaults

`func NewDepositSwitchGetRequestWithDefaults() *DepositSwitchGetRequest`

NewDepositSwitchGetRequestWithDefaults instantiates a new DepositSwitchGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDepositSwitchId

`func (o *DepositSwitchGetRequest) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchGetRequest) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchGetRequest) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


