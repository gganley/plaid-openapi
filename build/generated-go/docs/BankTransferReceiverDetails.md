# BankTransferReceiverDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableBalance** | Pointer to **NullableString** | The sign of the available balance for the receiver bank account associated with the receiver event at the time the matching transaction was found. Can be &#x60;positive&#x60;, &#x60;negative&#x60;, or null if the balance was not available at the time. | [optional] 

## Methods

### NewBankTransferReceiverDetails

`func NewBankTransferReceiverDetails() *BankTransferReceiverDetails`

NewBankTransferReceiverDetails instantiates a new BankTransferReceiverDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferReceiverDetailsWithDefaults

`func NewBankTransferReceiverDetailsWithDefaults() *BankTransferReceiverDetails`

NewBankTransferReceiverDetailsWithDefaults instantiates a new BankTransferReceiverDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableBalance

`func (o *BankTransferReceiverDetails) GetAvailableBalance() string`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *BankTransferReceiverDetails) GetAvailableBalanceOk() (*string, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *BankTransferReceiverDetails) SetAvailableBalance(v string)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *BankTransferReceiverDetails) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *BankTransferReceiverDetails) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *BankTransferReceiverDetails) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


