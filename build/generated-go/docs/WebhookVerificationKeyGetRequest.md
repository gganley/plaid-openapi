# WebhookVerificationKeyGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**KeyId** | **string** | The key ID ( &#x60;kid&#x60; ) from the JWT header. | 

## Methods

### NewWebhookVerificationKeyGetRequest

`func NewWebhookVerificationKeyGetRequest(keyId string, ) *WebhookVerificationKeyGetRequest`

NewWebhookVerificationKeyGetRequest instantiates a new WebhookVerificationKeyGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookVerificationKeyGetRequestWithDefaults

`func NewWebhookVerificationKeyGetRequestWithDefaults() *WebhookVerificationKeyGetRequest`

NewWebhookVerificationKeyGetRequestWithDefaults instantiates a new WebhookVerificationKeyGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *WebhookVerificationKeyGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WebhookVerificationKeyGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WebhookVerificationKeyGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *WebhookVerificationKeyGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *WebhookVerificationKeyGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *WebhookVerificationKeyGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *WebhookVerificationKeyGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *WebhookVerificationKeyGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetKeyId

`func (o *WebhookVerificationKeyGetRequest) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *WebhookVerificationKeyGetRequest) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *WebhookVerificationKeyGetRequest) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


