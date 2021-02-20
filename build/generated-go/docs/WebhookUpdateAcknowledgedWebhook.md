# WebhookUpdateAcknowledgedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;ITEM&#x60; | 
**WebhookCode** | **string** | &#x60;WEBHOOK_UPDATE_ACKNOWLEDGED&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**NewWebhookUrl** | **string** | The new webhook URL | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 

## Methods

### NewWebhookUpdateAcknowledgedWebhook

`func NewWebhookUpdateAcknowledgedWebhook(webhookType string, webhookCode string, itemId string, newWebhookUrl string, ) *WebhookUpdateAcknowledgedWebhook`

NewWebhookUpdateAcknowledgedWebhook instantiates a new WebhookUpdateAcknowledgedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookUpdateAcknowledgedWebhookWithDefaults

`func NewWebhookUpdateAcknowledgedWebhookWithDefaults() *WebhookUpdateAcknowledgedWebhook`

NewWebhookUpdateAcknowledgedWebhookWithDefaults instantiates a new WebhookUpdateAcknowledgedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *WebhookUpdateAcknowledgedWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *WebhookUpdateAcknowledgedWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *WebhookUpdateAcknowledgedWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *WebhookUpdateAcknowledgedWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *WebhookUpdateAcknowledgedWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *WebhookUpdateAcknowledgedWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetNewWebhookUrl

`func (o *WebhookUpdateAcknowledgedWebhook) GetNewWebhookUrl() string`

GetNewWebhookUrl returns the NewWebhookUrl field if non-nil, zero value otherwise.

### GetNewWebhookUrlOk

`func (o *WebhookUpdateAcknowledgedWebhook) GetNewWebhookUrlOk() (*string, bool)`

GetNewWebhookUrlOk returns a tuple with the NewWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewWebhookUrl

`func (o *WebhookUpdateAcknowledgedWebhook) SetNewWebhookUrl(v string)`

SetNewWebhookUrl sets NewWebhookUrl field to given value.


### GetError

`func (o *WebhookUpdateAcknowledgedWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookUpdateAcknowledgedWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookUpdateAcknowledgedWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookUpdateAcknowledgedWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WebhookUpdateAcknowledgedWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WebhookUpdateAcknowledgedWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


