# ItemStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Investments** | Pointer to [**NullableItemStatusInvestments**](ItemStatus_investments.md) |  | [optional] 
**Transactions** | Pointer to [**NullableItemStatusTransactions**](ItemStatus_transactions.md) |  | [optional] 
**LastWebhook** | Pointer to [**NullableItemStatusLastWebhook**](ItemStatus_last_webhook.md) |  | [optional] 

## Methods

### NewItemStatus

`func NewItemStatus() *ItemStatus`

NewItemStatus instantiates a new ItemStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStatusWithDefaults

`func NewItemStatusWithDefaults() *ItemStatus`

NewItemStatusWithDefaults instantiates a new ItemStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestments

`func (o *ItemStatus) GetInvestments() ItemStatusInvestments`

GetInvestments returns the Investments field if non-nil, zero value otherwise.

### GetInvestmentsOk

`func (o *ItemStatus) GetInvestmentsOk() (*ItemStatusInvestments, bool)`

GetInvestmentsOk returns a tuple with the Investments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestments

`func (o *ItemStatus) SetInvestments(v ItemStatusInvestments)`

SetInvestments sets Investments field to given value.

### HasInvestments

`func (o *ItemStatus) HasInvestments() bool`

HasInvestments returns a boolean if a field has been set.

### SetInvestmentsNil

`func (o *ItemStatus) SetInvestmentsNil(b bool)`

 SetInvestmentsNil sets the value for Investments to be an explicit nil

### UnsetInvestments
`func (o *ItemStatus) UnsetInvestments()`

UnsetInvestments ensures that no value is present for Investments, not even an explicit nil
### GetTransactions

`func (o *ItemStatus) GetTransactions() ItemStatusTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *ItemStatus) GetTransactionsOk() (*ItemStatusTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *ItemStatus) SetTransactions(v ItemStatusTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *ItemStatus) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *ItemStatus) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *ItemStatus) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
### GetLastWebhook

`func (o *ItemStatus) GetLastWebhook() ItemStatusLastWebhook`

GetLastWebhook returns the LastWebhook field if non-nil, zero value otherwise.

### GetLastWebhookOk

`func (o *ItemStatus) GetLastWebhookOk() (*ItemStatusLastWebhook, bool)`

GetLastWebhookOk returns a tuple with the LastWebhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWebhook

`func (o *ItemStatus) SetLastWebhook(v ItemStatusLastWebhook)`

SetLastWebhook sets LastWebhook field to given value.

### HasLastWebhook

`func (o *ItemStatus) HasLastWebhook() bool`

HasLastWebhook returns a boolean if a field has been set.

### SetLastWebhookNil

`func (o *ItemStatus) SetLastWebhookNil(b bool)`

 SetLastWebhookNil sets the value for LastWebhook to be an explicit nil

### UnsetLastWebhook
`func (o *ItemStatus) UnsetLastWebhook()`

UnsetLastWebhook ensures that no value is present for LastWebhook, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


