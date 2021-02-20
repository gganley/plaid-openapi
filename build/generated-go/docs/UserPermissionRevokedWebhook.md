# UserPermissionRevokedWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;ITEM&#x60; | 
**WebhookCode** | **string** | &#x60;USER_PERMISSION_REVOKED&#x60; | 
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 

## Methods

### NewUserPermissionRevokedWebhook

`func NewUserPermissionRevokedWebhook(webhookType string, webhookCode string, itemId string, ) *UserPermissionRevokedWebhook`

NewUserPermissionRevokedWebhook instantiates a new UserPermissionRevokedWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPermissionRevokedWebhookWithDefaults

`func NewUserPermissionRevokedWebhookWithDefaults() *UserPermissionRevokedWebhook`

NewUserPermissionRevokedWebhookWithDefaults instantiates a new UserPermissionRevokedWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *UserPermissionRevokedWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *UserPermissionRevokedWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *UserPermissionRevokedWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *UserPermissionRevokedWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *UserPermissionRevokedWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *UserPermissionRevokedWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetItemId

`func (o *UserPermissionRevokedWebhook) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *UserPermissionRevokedWebhook) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *UserPermissionRevokedWebhook) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetError

`func (o *UserPermissionRevokedWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *UserPermissionRevokedWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *UserPermissionRevokedWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *UserPermissionRevokedWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *UserPermissionRevokedWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *UserPermissionRevokedWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


