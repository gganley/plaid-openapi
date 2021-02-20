# SandboxItemFireWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**WebhookCode** | Pointer to **string** | The following values for &#x60;webhook_code&#x60; are supported:  * &#x60;DEFAULT_UPDATE&#x60; | [optional] 

## Methods

### NewSandboxItemFireWebhookRequest

`func NewSandboxItemFireWebhookRequest(accessToken string, ) *SandboxItemFireWebhookRequest`

NewSandboxItemFireWebhookRequest instantiates a new SandboxItemFireWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxItemFireWebhookRequestWithDefaults

`func NewSandboxItemFireWebhookRequestWithDefaults() *SandboxItemFireWebhookRequest`

NewSandboxItemFireWebhookRequestWithDefaults instantiates a new SandboxItemFireWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxItemFireWebhookRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxItemFireWebhookRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxItemFireWebhookRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxItemFireWebhookRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxItemFireWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxItemFireWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxItemFireWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxItemFireWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *SandboxItemFireWebhookRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SandboxItemFireWebhookRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SandboxItemFireWebhookRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetWebhookCode

`func (o *SandboxItemFireWebhookRequest) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *SandboxItemFireWebhookRequest) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *SandboxItemFireWebhookRequest) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.

### HasWebhookCode

`func (o *SandboxItemFireWebhookRequest) HasWebhookCode() bool`

HasWebhookCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


