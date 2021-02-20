# AutomaticallyVerifiedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;AUTH&#x60; | 
**WebhookCode** | **string** | &#x60;AUTOMATICALLY_VERIFIED&#x60; | 
**AccountId** | **string** | The &#x60;account_id&#x60; of the account associated with the webhook | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 

## Methods

### NewAutomaticallyVerifiedWebhook

`func NewAutomaticallyVerifiedWebhook(webhookType string, webhookCode string, accountId string, itemId string, ) *AutomaticallyVerifiedWebhook`

NewAutomaticallyVerifiedWebhook instantiates a new AutomaticallyVerifiedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticallyVerifiedWebhookWithDefaults

`func NewAutomaticallyVerifiedWebhookWithDefaults() *AutomaticallyVerifiedWebhook`

NewAutomaticallyVerifiedWebhookWithDefaults instantiates a new AutomaticallyVerifiedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *AutomaticallyVerifiedWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *AutomaticallyVerifiedWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *AutomaticallyVerifiedWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *AutomaticallyVerifiedWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *AutomaticallyVerifiedWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *AutomaticallyVerifiedWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetAccountId

`func (o *AutomaticallyVerifiedWebhook) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutomaticallyVerifiedWebhook) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutomaticallyVerifiedWebhook) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetItemId

`func (o *AutomaticallyVerifiedWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AutomaticallyVerifiedWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AutomaticallyVerifiedWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


