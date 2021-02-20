# TransactionsRemovedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;TRANSACTIONS&#x60; | 
**WebhookCode** | **string** | &#x60;TRANSACTIONS_REMOVED&#x60; | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**RemovedTransactions** | **[]string** | An array of &#x60;transaction_ids&#x60; corresponding to the removed transactions | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 

## Methods

### NewTransactionsRemovedWebhook

`func NewTransactionsRemovedWebhook(webhookType string, webhookCode string, removedTransactions []string, itemId string, ) *TransactionsRemovedWebhook`

NewTransactionsRemovedWebhook instantiates a new TransactionsRemovedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsRemovedWebhookWithDefaults

`func NewTransactionsRemovedWebhookWithDefaults() *TransactionsRemovedWebhook`

NewTransactionsRemovedWebhookWithDefaults instantiates a new TransactionsRemovedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *TransactionsRemovedWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *TransactionsRemovedWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *TransactionsRemovedWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *TransactionsRemovedWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *TransactionsRemovedWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *TransactionsRemovedWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetError

`func (o *TransactionsRemovedWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TransactionsRemovedWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TransactionsRemovedWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *TransactionsRemovedWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *TransactionsRemovedWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *TransactionsRemovedWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetRemovedTransactions

`func (o *TransactionsRemovedWebhook) GetRemovedTransactions() []string`

GetRemovedTransactions returns the RemovedTransactions field if non-nil, zero value otherwise.

### GetRemovedTransactionsOk

`func (o *TransactionsRemovedWebhook) GetRemovedTransactionsOk() (*[]string, bool)`

GetRemovedTransactionsOk returns a tuple with the RemovedTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedTransactions

`func (o *TransactionsRemovedWebhook) SetRemovedTransactions(v []string)`

SetRemovedTransactions sets RemovedTransactions field to given value.


### GetItemId

`func (o *TransactionsRemovedWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *TransactionsRemovedWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *TransactionsRemovedWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


