# DepositSwitchTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**DepositSwitchId** | **string** | The ID of the deposit switch | 

## Methods

### NewDepositSwitchTokenCreateRequest

`func NewDepositSwitchTokenCreateRequest(depositSwitchId string, ) *DepositSwitchTokenCreateRequest`

NewDepositSwitchTokenCreateRequest instantiates a new DepositSwitchTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchTokenCreateRequestWithDefaults

`func NewDepositSwitchTokenCreateRequestWithDefaults() *DepositSwitchTokenCreateRequest`

NewDepositSwitchTokenCreateRequestWithDefaults instantiates a new DepositSwitchTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *DepositSwitchTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *DepositSwitchTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *DepositSwitchTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *DepositSwitchTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *DepositSwitchTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *DepositSwitchTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *DepositSwitchTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *DepositSwitchTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDepositSwitchId

`func (o *DepositSwitchTokenCreateRequest) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchTokenCreateRequest) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchTokenCreateRequest) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


