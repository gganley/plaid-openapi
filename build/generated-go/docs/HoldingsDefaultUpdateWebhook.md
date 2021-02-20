# HoldingsDefaultUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;HOLDINGS&#x60; | 
**WebhookCode** | **string** | &#x60;DEFAULT_UPDATE&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**NewHoldings** | **float32** | The number of new holdings reported since the last time this webhook was fired. | 
**UpdatedHoldings** | **float32** | The number of updated holdings reported since the last time this webhook was fired. | 

## Methods

### NewHoldingsDefaultUpdateWebhook

`func NewHoldingsDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, newHoldings float32, updatedHoldings float32, ) *HoldingsDefaultUpdateWebhook`

NewHoldingsDefaultUpdateWebhook instantiates a new HoldingsDefaultUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingsDefaultUpdateWebhookWithDefaults

`func NewHoldingsDefaultUpdateWebhookWithDefaults() *HoldingsDefaultUpdateWebhook`

NewHoldingsDefaultUpdateWebhookWithDefaults instantiates a new HoldingsDefaultUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *HoldingsDefaultUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *HoldingsDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *HoldingsDefaultUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *HoldingsDefaultUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *HoldingsDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *HoldingsDefaultUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *HoldingsDefaultUpdateWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *HoldingsDefaultUpdateWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *HoldingsDefaultUpdateWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetError

`func (o *HoldingsDefaultUpdateWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *HoldingsDefaultUpdateWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *HoldingsDefaultUpdateWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *HoldingsDefaultUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *HoldingsDefaultUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *HoldingsDefaultUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetNewHoldings

`func (o *HoldingsDefaultUpdateWebhook) GetNewHoldings() float32`

GetNewHoldings returns the NewHoldings field if non-nil, zero value otherwise.

### GetNewHoldingsOk

`func (o *HoldingsDefaultUpdateWebhook) GetNewHoldingsOk() (*float32, bool)`

GetNewHoldingsOk returns a tuple with the NewHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewHoldings

`func (o *HoldingsDefaultUpdateWebhook) SetNewHoldings(v float32)`

SetNewHoldings sets NewHoldings field to given value.


### GetUpdatedHoldings

`func (o *HoldingsDefaultUpdateWebhook) GetUpdatedHoldings() float32`

GetUpdatedHoldings returns the UpdatedHoldings field if non-nil, zero value otherwise.

### GetUpdatedHoldingsOk

`func (o *HoldingsDefaultUpdateWebhook) GetUpdatedHoldingsOk() (*float32, bool)`

GetUpdatedHoldingsOk returns a tuple with the UpdatedHoldings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedHoldings

`func (o *HoldingsDefaultUpdateWebhook) SetUpdatedHoldings(v float32)`

SetUpdatedHoldings sets UpdatedHoldings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


