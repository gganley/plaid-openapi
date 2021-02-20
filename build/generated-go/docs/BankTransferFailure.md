# BankTransferFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchReturnCode** | Pointer to **NullableString** | The ACH return code, e.g. &#x60;R01&#x60;.  A return code will be provided if and only if the transfer status is &#x60;reversed&#x60;. For a full listing of ACH return codes, see [Bank Transfers errors](/docs/errors/bank-transfers/#ach-return-codes). | [optional] 
**Description** | Pointer to **string** | A human-readable description of the reason for the failure or reversal. | [optional] 

## Methods

### NewBankTransferFailure

`func NewBankTransferFailure() *BankTransferFailure`

NewBankTransferFailure instantiates a new BankTransferFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferFailureWithDefaults

`func NewBankTransferFailureWithDefaults() *BankTransferFailure`

NewBankTransferFailureWithDefaults instantiates a new BankTransferFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchReturnCode

`func (o *BankTransferFailure) GetAchReturnCode() string`

GetAchReturnCode returns the AchReturnCode field if non-nil, zero value otherwise.

### GetAchReturnCodeOk

`func (o *BankTransferFailure) GetAchReturnCodeOk() (*string, bool)`

GetAchReturnCodeOk returns a tuple with the AchReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchReturnCode

`func (o *BankTransferFailure) SetAchReturnCode(v string)`

SetAchReturnCode sets AchReturnCode field to given value.

### HasAchReturnCode

`func (o *BankTransferFailure) HasAchReturnCode() bool`

HasAchReturnCode returns a boolean if a field has been set.

### SetAchReturnCodeNil

`func (o *BankTransferFailure) SetAchReturnCodeNil(b bool)`

 SetAchReturnCodeNil sets the value for AchReturnCode to be an explicit nil

### UnsetAchReturnCode
`func (o *BankTransferFailure) UnsetAchReturnCode()`

UnsetAchReturnCode ensures that no value is present for AchReturnCode, not even an explicit nil
### GetDescription

`func (o *BankTransferFailure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransferFailure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransferFailure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTransferFailure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


