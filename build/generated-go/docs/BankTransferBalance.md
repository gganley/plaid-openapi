# BankTransferBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **string** | The total available balance - the sum of all successful debit transfer amounts minus all credit transfer amounts. | 
**Transactable** | **string** | The transactable balance shows the amount in your account that you are able to use for transfers, and is essentially your available balance minus your minimum balance. | 

## Methods

### NewBankTransferBalance

`func NewBankTransferBalance(available string, transactable string, ) *BankTransferBalance`

NewBankTransferBalance instantiates a new BankTransferBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferBalanceWithDefaults

`func NewBankTransferBalanceWithDefaults() *BankTransferBalance`

NewBankTransferBalanceWithDefaults instantiates a new BankTransferBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *BankTransferBalance) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *BankTransferBalance) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *BankTransferBalance) SetAvailable(v string)`

SetAvailable sets Available field to given value.


### GetTransactable

`func (o *BankTransferBalance) GetTransactable() string`

GetTransactable returns the Transactable field if non-nil, zero value otherwise.

### GetTransactableOk

`func (o *BankTransferBalance) GetTransactableOk() (*string, bool)`

GetTransactableOk returns a tuple with the Transactable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactable

`func (o *BankTransferBalance) SetTransactable(v string)`

SetTransactable sets Transactable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


