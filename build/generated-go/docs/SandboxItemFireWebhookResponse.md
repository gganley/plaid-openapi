# SandboxItemFireWebhookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookFired** | **bool** | Value is &#x60;true&#x60;  if the test&#x60; webhook_code&#x60;  was successfully fired. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewSandboxItemFireWebhookResponse

`func NewSandboxItemFireWebhookResponse(webhookFired bool, requestId string, ) *SandboxItemFireWebhookResponse`

NewSandboxItemFireWebhookResponse instantiates a new SandboxItemFireWebhookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxItemFireWebhookResponseWithDefaults

`func NewSandboxItemFireWebhookResponseWithDefaults() *SandboxItemFireWebhookResponse`

NewSandboxItemFireWebhookResponseWithDefaults instantiates a new SandboxItemFireWebhookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookFired

`func (o *SandboxItemFireWebhookResponse) GetWebhookFired() bool`

GetWebhookFired returns the WebhookFired field if non-nil, zero value otherwise.

### GetWebhookFiredOk

`func (o *SandboxItemFireWebhookResponse) GetWebhookFiredOk() (*bool, bool)`

GetWebhookFiredOk returns a tuple with the WebhookFired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFired

`func (o *SandboxItemFireWebhookResponse) SetWebhookFired(v bool)`

SetWebhookFired sets WebhookFired field to given value.


### GetRequestId

`func (o *SandboxItemFireWebhookResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SandboxItemFireWebhookResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SandboxItemFireWebhookResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


