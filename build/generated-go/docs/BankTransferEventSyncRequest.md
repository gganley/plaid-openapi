# BankTransferEventSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AfterId** | **int32** | The latest (largest) &#x60;event_id&#x60; fetched via the sync endpoint, or 0 initially. | 
**Count** | Pointer to **NullableInt32** | The maximum number of bank transfer events to return. | [optional] [default to 25]

## Methods

### NewBankTransferEventSyncRequest

`func NewBankTransferEventSyncRequest(afterId int32, ) *BankTransferEventSyncRequest`

NewBankTransferEventSyncRequest instantiates a new BankTransferEventSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferEventSyncRequestWithDefaults

`func NewBankTransferEventSyncRequestWithDefaults() *BankTransferEventSyncRequest`

NewBankTransferEventSyncRequestWithDefaults instantiates a new BankTransferEventSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferEventSyncRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferEventSyncRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferEventSyncRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferEventSyncRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferEventSyncRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferEventSyncRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferEventSyncRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferEventSyncRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAfterId

`func (o *BankTransferEventSyncRequest) GetAfterId() int32`

GetAfterId returns the AfterId field if non-nil, zero value otherwise.

### GetAfterIdOk

`func (o *BankTransferEventSyncRequest) GetAfterIdOk() (*int32, bool)`

GetAfterIdOk returns a tuple with the AfterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterId

`func (o *BankTransferEventSyncRequest) SetAfterId(v int32)`

SetAfterId sets AfterId field to given value.


### GetCount

`func (o *BankTransferEventSyncRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BankTransferEventSyncRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BankTransferEventSyncRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BankTransferEventSyncRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BankTransferEventSyncRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BankTransferEventSyncRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


