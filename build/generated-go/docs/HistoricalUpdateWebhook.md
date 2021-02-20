# HistoricalUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;TRANSACTIONS&#x60; | 
**WebhookCode** | **string** | &#x60;HISTORICAL_UPDATE&#x60; | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**NewTransactions** | **float32** | The number of new, unfetched transactions available | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 

## Methods

### NewHistoricalUpdateWebhook

`func NewHistoricalUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string, ) *HistoricalUpdateWebhook`

NewHistoricalUpdateWebhook instantiates a new HistoricalUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalUpdateWebhookWithDefaults

`func NewHistoricalUpdateWebhookWithDefaults() *HistoricalUpdateWebhook`

NewHistoricalUpdateWebhookWithDefaults instantiates a new HistoricalUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *HistoricalUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *HistoricalUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *HistoricalUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *HistoricalUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *HistoricalUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *HistoricalUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetError

`func (o *HistoricalUpdateWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HistoricalUpdateWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HistoricalUpdateWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *HistoricalUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HistoricalUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HistoricalUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNewTransactions

`func (o *HistoricalUpdateWebhook) GetNewTransactions() float32`

GetNewTransactions returns the NewTransactions field if non-nil, zero value otherwise.

### GetNewTransactionsOk

`func (o *HistoricalUpdateWebhook) GetNewTransactionsOk() (*float32, bool)`

GetNewTransactionsOk returns a tuple with the NewTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTransactions

`func (o *HistoricalUpdateWebhook) SetNewTransactions(v float32)`

SetNewTransactions sets NewTransactions field to given value.


### GetItemId

`func (o *HistoricalUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *HistoricalUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *HistoricalUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


