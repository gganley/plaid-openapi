# TransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the transaction. | 
**Amount** | **float32** | The amount of the transaction. | 
**Date** | **string** | The date of the transaction, in ISO 8601 format (\&quot;yyyy-mm-dd\&quot;). | 
**AccountId** | **string** | A unique identifier for the end user&#39;s account. | 
**TransactionId** | **string** | A unique identifier for the transaction. | 

## Methods

### NewTransactionData

`func NewTransactionData(description string, amount float32, date string, accountId string, transactionId string, ) *TransactionData`

NewTransactionData instantiates a new TransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionDataWithDefaults

`func NewTransactionDataWithDefaults() *TransactionData`

NewTransactionDataWithDefaults instantiates a new TransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TransactionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionData) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAmount

`func (o *TransactionData) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionData) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionData) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetDate

`func (o *TransactionData) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TransactionData) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TransactionData) SetDate(v string)`

SetDate sets Date field to given value.


### GetAccountId

`func (o *TransactionData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransactionData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransactionData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTransactionId

`func (o *TransactionData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *TransactionData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *TransactionData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


