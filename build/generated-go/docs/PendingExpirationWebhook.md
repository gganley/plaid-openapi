# PendingExpirationWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;ITEM&#x60; | 
**WebhookCode** | **string** | &#x60;PENDING_EXPIRATION&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**ConsentExpirationTime** | **string** | The date and time at which the Item&#39;s access consent will expire, in ISO 8601 format | 

## Methods

### NewPendingExpirationWebhook

`func NewPendingExpirationWebhook(webhookType string, webhookCode string, itemId string, consentExpirationTime string, ) *PendingExpirationWebhook`

NewPendingExpirationWebhook instantiates a new PendingExpirationWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPendingExpirationWebhookWithDefaults

`func NewPendingExpirationWebhookWithDefaults() *PendingExpirationWebhook`

NewPendingExpirationWebhookWithDefaults instantiates a new PendingExpirationWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *PendingExpirationWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *PendingExpirationWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *PendingExpirationWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *PendingExpirationWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *PendingExpirationWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *PendingExpirationWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *PendingExpirationWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *PendingExpirationWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *PendingExpirationWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetConsentExpirationTime

`func (o *PendingExpirationWebhook) GetConsentExpirationTime() string`

GetConsentExpirationTime returns the ConsentExpirationTime field if non-nil, zero value otherwise.

### GetConsentExpirationTimeOk

`func (o *PendingExpirationWebhook) GetConsentExpirationTimeOk() (*string, bool)`

GetConsentExpirationTimeOk returns a tuple with the ConsentExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentExpirationTime

`func (o *PendingExpirationWebhook) SetConsentExpirationTime(v string)`

SetConsentExpirationTime sets ConsentExpirationTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


