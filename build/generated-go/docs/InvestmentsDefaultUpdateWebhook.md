# InvestmentsDefaultUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;INVESTMENTS_TRANSACTIONS&#x60; | 
**WebhookCode** | **string** | &#x60;DEFAULT_UPDATE&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**NewInvestmentsTransactions** | **float32** | The number of new transactions reported since the last time this webhook was fired. | 
**CanceledInvestmentsTransactions** | **float32** | The number of canceled transactions reported since the last time this webhook was fired. | 

## Methods

### NewInvestmentsDefaultUpdateWebhook

`func NewInvestmentsDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, newInvestmentsTransactions float32, canceledInvestmentsTransactions float32, ) *InvestmentsDefaultUpdateWebhook`

NewInvestmentsDefaultUpdateWebhook instantiates a new InvestmentsDefaultUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentsDefaultUpdateWebhookWithDefaults

`func NewInvestmentsDefaultUpdateWebhookWithDefaults() *InvestmentsDefaultUpdateWebhook`

NewInvestmentsDefaultUpdateWebhookWithDefaults instantiates a new InvestmentsDefaultUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *InvestmentsDefaultUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *InvestmentsDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *InvestmentsDefaultUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *InvestmentsDefaultUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *InvestmentsDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *InvestmentsDefaultUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *InvestmentsDefaultUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *InvestmentsDefaultUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *InvestmentsDefaultUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetError

`func (o *InvestmentsDefaultUpdateWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *InvestmentsDefaultUpdateWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *InvestmentsDefaultUpdateWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *InvestmentsDefaultUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *InvestmentsDefaultUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *InvestmentsDefaultUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNewInvestmentsTransactions

`func (o *InvestmentsDefaultUpdateWebhook) GetNewInvestmentsTransactions() float32`

GetNewInvestmentsTransactions returns the NewInvestmentsTransactions field if non-nil, zero value otherwise.

### GetNewInvestmentsTransactionsOk

`func (o *InvestmentsDefaultUpdateWebhook) GetNewInvestmentsTransactionsOk() (*float32, bool)`

GetNewInvestmentsTransactionsOk returns a tuple with the NewInvestmentsTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewInvestmentsTransactions

`func (o *InvestmentsDefaultUpdateWebhook) SetNewInvestmentsTransactions(v float32)`

SetNewInvestmentsTransactions sets NewInvestmentsTransactions field to given value.


### GetCanceledInvestmentsTransactions

`func (o *InvestmentsDefaultUpdateWebhook) GetCanceledInvestmentsTransactions() float32`

GetCanceledInvestmentsTransactions returns the CanceledInvestmentsTransactions field if non-nil, zero value otherwise.

### GetCanceledInvestmentsTransactionsOk

`func (o *InvestmentsDefaultUpdateWebhook) GetCanceledInvestmentsTransactionsOk() (*float32, bool)`

GetCanceledInvestmentsTransactionsOk returns a tuple with the CanceledInvestmentsTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledInvestmentsTransactions

`func (o *InvestmentsDefaultUpdateWebhook) SetCanceledInvestmentsTransactions(v float32)`

SetCanceledInvestmentsTransactions sets CanceledInvestmentsTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


