# DefaultUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;TRANSACTIONS&#x60; | 
**WebhookCode** | **string** | &#x60;DEFAULT_UPDATE&#x60; | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**NewTransactions** | **float32** | The number of new transactions detected since the last time this webhook was fired. | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item the webhook relates to. | 

## Methods

### NewDefaultUpdateWebhook

`func NewDefaultUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string, ) *DefaultUpdateWebhook`

NewDefaultUpdateWebhook instantiates a new DefaultUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultUpdateWebhookWithDefaults

`func NewDefaultUpdateWebhookWithDefaults() *DefaultUpdateWebhook`

NewDefaultUpdateWebhookWithDefaults instantiates a new DefaultUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *DefaultUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *DefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *DefaultUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *DefaultUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *DefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *DefaultUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetError

`func (o *DefaultUpdateWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DefaultUpdateWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DefaultUpdateWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *DefaultUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *DefaultUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *DefaultUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNewTransactions

`func (o *DefaultUpdateWebhook) GetNewTransactions() float32`

GetNewTransactions returns the NewTransactions field if non-nil, zero value otherwise.

### GetNewTransactionsOk

`func (o *DefaultUpdateWebhook) GetNewTransactionsOk() (*float32, bool)`

GetNewTransactionsOk returns a tuple with the NewTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTransactions

`func (o *DefaultUpdateWebhook) SetNewTransactions(v float32)`

SetNewTransactions sets NewTransactions field to given value.


### GetItemId

`func (o *DefaultUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *DefaultUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *DefaultUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


