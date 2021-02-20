# InitialUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;TRANSACTIONS&#x60; | 
**WebhookCode** | **string** | &#x60;INITIAL_UPDATE&#x60; | 
**Error** | Pointer to **NullableString** | The error code associated with the webhook. | [optional] 
**NewTransactions** | **float32** | The number of new, unfetched transactions available. | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 

## Methods

### NewInitialUpdateWebhook

`func NewInitialUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string, ) *InitialUpdateWebhook`

NewInitialUpdateWebhook instantiates a new InitialUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitialUpdateWebhookWithDefaults

`func NewInitialUpdateWebhookWithDefaults() *InitialUpdateWebhook`

NewInitialUpdateWebhookWithDefaults instantiates a new InitialUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *InitialUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *InitialUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *InitialUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *InitialUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *InitialUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *InitialUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetError

`func (o *InitialUpdateWebhook) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InitialUpdateWebhook) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InitialUpdateWebhook) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *InitialUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *InitialUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *InitialUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNewTransactions

`func (o *InitialUpdateWebhook) GetNewTransactions() float32`

GetNewTransactions returns the NewTransactions field if non-nil, zero value otherwise.

### GetNewTransactionsOk

`func (o *InitialUpdateWebhook) GetNewTransactionsOk() (*float32, bool)`

GetNewTransactionsOk returns a tuple with the NewTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTransactions

`func (o *InitialUpdateWebhook) SetNewTransactions(v float32)`

SetNewTransactions sets NewTransactions field to given value.


### GetItemId

`func (o *InitialUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InitialUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InitialUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


