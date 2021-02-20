# DepositSwitchGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositSwitchId** | **string** | The ID of the deposit switch | 
**TargetAccountId** | **NullableString** | The ID of the bank account the direct deposit was switched to | 
**TargetItemId** | **NullableString** | The ID of the Item the direct deposit was switched to. | 
**State** | **string** | The state of the deposit switch. | 
**AccountHasMultipleAllocations** | **NullableBool** | When &#x60;true&#x60;, user’s direct deposit goes to multiple banks. When false, user’s direct deposit only goes to the target account. Always &#x60;null&#x60; if the deposit switch has not been completed. | 
**IsAllocatedRemainder** | **NullableBool** | When &#x60;true&#x60;, the target account is allocated the remainder of direct deposit after all other allocations have been deducted. When &#x60;false&#x60;, user’s direct deposit is allocated as a percent or amount. Always &#x60;null&#x60; if the deposit switch has not been completed. | 
**PercentAllocated** | **NullableFloat32** | The percentage of direct deposit allocated to the target account. Always &#x60;null&#x60; if the target account is not allocated a percentage or if the deposit switch has not been completed or if &#x60;is_allocated_remainder&#x60; is true. | 
**AmountAllocated** | **NullableFloat32** | The dollar amount of direct deposit allocated to the target account. Always &#x60;null&#x60; if the target account is not allocated an amount or if the deposit switch has not been completed. | 
**DateCreated** | **string** | ISO8601 date the deposit switch was created. | 
**DateCompleted** | **NullableString** | ISO8601 date the deposit switch was completed. Always &#x60;null&#x60; if the deposit switch has not been completed. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewDepositSwitchGetResponse

`func NewDepositSwitchGetResponse(depositSwitchId string, targetAccountId NullableString, targetItemId NullableString, state string, accountHasMultipleAllocations NullableBool, isAllocatedRemainder NullableBool, percentAllocated NullableFloat32, amountAllocated NullableFloat32, dateCreated string, dateCompleted NullableString, requestId string, ) *DepositSwitchGetResponse`

NewDepositSwitchGetResponse instantiates a new DepositSwitchGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchGetResponseWithDefaults

`func NewDepositSwitchGetResponseWithDefaults() *DepositSwitchGetResponse`

NewDepositSwitchGetResponseWithDefaults instantiates a new DepositSwitchGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositSwitchId

`func (o *DepositSwitchGetResponse) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchGetResponse) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchGetResponse) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.


### GetTargetAccountId

`func (o *DepositSwitchGetResponse) GetTargetAccountId() string`

GetTargetAccountId returns the TargetAccountId field if non-nil, zero value otherwise.

### GetTargetAccountIdOk

`func (o *DepositSwitchGetResponse) GetTargetAccountIdOk() (*string, bool)`

GetTargetAccountIdOk returns a tuple with the TargetAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAccountId

`func (o *DepositSwitchGetResponse) SetTargetAccountId(v string)`

SetTargetAccountId sets TargetAccountId field to given value.


### SetTargetAccountIdNil

`func (o *DepositSwitchGetResponse) SetTargetAccountIdNil(b bool)`

 SetTargetAccountIdNil sets the value for TargetAccountId to be an explicit nil

### UnsetTargetAccountId
`func (o *DepositSwitchGetResponse) UnsetTargetAccountId()`

UnsetTargetAccountId ensures that no value is present for TargetAccountId, not even an explicit nil
### GetTargetItemId

`func (o *DepositSwitchGetResponse) GetTargetItemId() string`

GetTargetItemId returns the TargetItemId field if non-nil, zero value otherwise.

### GetTargetItemIdOk

`func (o *DepositSwitchGetResponse) GetTargetItemIdOk() (*string, bool)`

GetTargetItemIdOk returns a tuple with the TargetItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetItemId

`func (o *DepositSwitchGetResponse) SetTargetItemId(v string)`

SetTargetItemId sets TargetItemId field to given value.


### SetTargetItemIdNil

`func (o *DepositSwitchGetResponse) SetTargetItemIdNil(b bool)`

 SetTargetItemIdNil sets the value for TargetItemId to be an explicit nil

### UnsetTargetItemId
`func (o *DepositSwitchGetResponse) UnsetTargetItemId()`

UnsetTargetItemId ensures that no value is present for TargetItemId, not even an explicit nil
### GetState

`func (o *DepositSwitchGetResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DepositSwitchGetResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DepositSwitchGetResponse) SetState(v string)`

SetState sets State field to given value.


### GetAccountHasMultipleAllocations

`func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocations() bool`

GetAccountHasMultipleAllocations returns the AccountHasMultipleAllocations field if non-nil, zero value otherwise.

### GetAccountHasMultipleAllocationsOk

`func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocationsOk() (*bool, bool)`

GetAccountHasMultipleAllocationsOk returns a tuple with the AccountHasMultipleAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHasMultipleAllocations

`func (o *DepositSwitchGetResponse) SetAccountHasMultipleAllocations(v bool)`

SetAccountHasMultipleAllocations sets AccountHasMultipleAllocations field to given value.


### SetAccountHasMultipleAllocationsNil

`func (o *DepositSwitchGetResponse) SetAccountHasMultipleAllocationsNil(b bool)`

 SetAccountHasMultipleAllocationsNil sets the value for AccountHasMultipleAllocations to be an explicit nil

### UnsetAccountHasMultipleAllocations
`func (o *DepositSwitchGetResponse) UnsetAccountHasMultipleAllocations()`

UnsetAccountHasMultipleAllocations ensures that no value is present for AccountHasMultipleAllocations, not even an explicit nil
### GetIsAllocatedRemainder

`func (o *DepositSwitchGetResponse) GetIsAllocatedRemainder() bool`

GetIsAllocatedRemainder returns the IsAllocatedRemainder field if non-nil, zero value otherwise.

### GetIsAllocatedRemainderOk

`func (o *DepositSwitchGetResponse) GetIsAllocatedRemainderOk() (*bool, bool)`

GetIsAllocatedRemainderOk returns a tuple with the IsAllocatedRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAllocatedRemainder

`func (o *DepositSwitchGetResponse) SetIsAllocatedRemainder(v bool)`

SetIsAllocatedRemainder sets IsAllocatedRemainder field to given value.


### SetIsAllocatedRemainderNil

`func (o *DepositSwitchGetResponse) SetIsAllocatedRemainderNil(b bool)`

 SetIsAllocatedRemainderNil sets the value for IsAllocatedRemainder to be an explicit nil

### UnsetIsAllocatedRemainder
`func (o *DepositSwitchGetResponse) UnsetIsAllocatedRemainder()`

UnsetIsAllocatedRemainder ensures that no value is present for IsAllocatedRemainder, not even an explicit nil
### GetPercentAllocated

`func (o *DepositSwitchGetResponse) GetPercentAllocated() float32`

GetPercentAllocated returns the PercentAllocated field if non-nil, zero value otherwise.

### GetPercentAllocatedOk

`func (o *DepositSwitchGetResponse) GetPercentAllocatedOk() (*float32, bool)`

GetPercentAllocatedOk returns a tuple with the PercentAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentAllocated

`func (o *DepositSwitchGetResponse) SetPercentAllocated(v float32)`

SetPercentAllocated sets PercentAllocated field to given value.


### SetPercentAllocatedNil

`func (o *DepositSwitchGetResponse) SetPercentAllocatedNil(b bool)`

 SetPercentAllocatedNil sets the value for PercentAllocated to be an explicit nil

### UnsetPercentAllocated
`func (o *DepositSwitchGetResponse) UnsetPercentAllocated()`

UnsetPercentAllocated ensures that no value is present for PercentAllocated, not even an explicit nil
### GetAmountAllocated

`func (o *DepositSwitchGetResponse) GetAmountAllocated() float32`

GetAmountAllocated returns the AmountAllocated field if non-nil, zero value otherwise.

### GetAmountAllocatedOk

`func (o *DepositSwitchGetResponse) GetAmountAllocatedOk() (*float32, bool)`

GetAmountAllocatedOk returns a tuple with the AmountAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountAllocated

`func (o *DepositSwitchGetResponse) SetAmountAllocated(v float32)`

SetAmountAllocated sets AmountAllocated field to given value.


### SetAmountAllocatedNil

`func (o *DepositSwitchGetResponse) SetAmountAllocatedNil(b bool)`

 SetAmountAllocatedNil sets the value for AmountAllocated to be an explicit nil

### UnsetAmountAllocated
`func (o *DepositSwitchGetResponse) UnsetAmountAllocated()`

UnsetAmountAllocated ensures that no value is present for AmountAllocated, not even an explicit nil
### GetDateCreated

`func (o *DepositSwitchGetResponse) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *DepositSwitchGetResponse) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *DepositSwitchGetResponse) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.


### GetDateCompleted

`func (o *DepositSwitchGetResponse) GetDateCompleted() string`

GetDateCompleted returns the DateCompleted field if non-nil, zero value otherwise.

### GetDateCompletedOk

`func (o *DepositSwitchGetResponse) GetDateCompletedOk() (*string, bool)`

GetDateCompletedOk returns a tuple with the DateCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCompleted

`func (o *DepositSwitchGetResponse) SetDateCompleted(v string)`

SetDateCompleted sets DateCompleted field to given value.


### SetDateCompletedNil

`func (o *DepositSwitchGetResponse) SetDateCompletedNil(b bool)`

 SetDateCompletedNil sets the value for DateCompleted to be an explicit nil

### UnsetDateCompleted
`func (o *DepositSwitchGetResponse) UnsetDateCompleted()`

UnsetDateCompleted ensures that no value is present for DateCompleted, not even an explicit nil
### GetRequestId

`func (o *DepositSwitchGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DepositSwitchGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DepositSwitchGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


