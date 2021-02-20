# TransactionOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionDate** | **string** | The date of the transaction, in ISO8601 (YYYY-MM-DD) format. Transaction dates in the past or present will result in posted transactions; transaction dates in the future will result in pending transactions. Transactions in Sandbox will move from pending to posted once their transaction date has been reached. | 
**PostedDate** | **string** | The date the transaction posted, in ISO8601 (YYYY-MM-DD) format | 
**Amount** | **float32** | The transaction amount. Can be negative. | 
**Description** | **string** | The transaction description. | 
**Currency** | Pointer to **string** | The ISO-4217 format currency code for the transaction. | [optional] 

## Methods

### NewTransactionOverride

`func NewTransactionOverride(transactionDate string, postedDate string, amount float32, description string, ) *TransactionOverride`

NewTransactionOverride instantiates a new TransactionOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOverrideWithDefaults

`func NewTransactionOverrideWithDefaults() *TransactionOverride`

NewTransactionOverrideWithDefaults instantiates a new TransactionOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionDate

`func (o *TransactionOverride) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *TransactionOverride) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *TransactionOverride) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.


### GetPostedDate

`func (o *TransactionOverride) GetPostedDate() string`

GetPostedDate returns the PostedDate field if non-nil, zero value otherwise.

### GetPostedDateOk

`func (o *TransactionOverride) GetPostedDateOk() (*string, bool)`

GetPostedDateOk returns a tuple with the PostedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostedDate

`func (o *TransactionOverride) SetPostedDate(v string)`

SetPostedDate sets PostedDate field to given value.


### GetAmount

`func (o *TransactionOverride) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionOverride) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionOverride) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *TransactionOverride) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionOverride) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionOverride) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCurrency

`func (o *TransactionOverride) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *TransactionOverride) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *TransactionOverride) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *TransactionOverride) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


