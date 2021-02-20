# VerificationExpiredWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;AUTH&#x60; | 
**WebhookCode** | **string** | &#x60;VERIFICATION_EXPIRED&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**AccountId** | **string** | The &#x60;account_id&#x60; of the account associated with the webhook | 

## Methods

### NewVerificationExpiredWebhook

`func NewVerificationExpiredWebhook(webhookType string, webhookCode string, itemId string, accountId string, ) *VerificationExpiredWebhook`

NewVerificationExpiredWebhook instantiates a new VerificationExpiredWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationExpiredWebhookWithDefaults

`func NewVerificationExpiredWebhookWithDefaults() *VerificationExpiredWebhook`

NewVerificationExpiredWebhookWithDefaults instantiates a new VerificationExpiredWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *VerificationExpiredWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *VerificationExpiredWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *VerificationExpiredWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *VerificationExpiredWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *VerificationExpiredWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *VerificationExpiredWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *VerificationExpiredWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *VerificationExpiredWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *VerificationExpiredWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetAccountId

`func (o *VerificationExpiredWebhook) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VerificationExpiredWebhook) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VerificationExpiredWebhook) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


