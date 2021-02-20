# WebhookVerificationKeyGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**JWKPublicKey**](JWKPublicKey.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewWebhookVerificationKeyGetResponse

`func NewWebhookVerificationKeyGetResponse(key JWKPublicKey, requestId string, ) *WebhookVerificationKeyGetResponse`

NewWebhookVerificationKeyGetResponse instantiates a new WebhookVerificationKeyGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookVerificationKeyGetResponseWithDefaults

`func NewWebhookVerificationKeyGetResponseWithDefaults() *WebhookVerificationKeyGetResponse`

NewWebhookVerificationKeyGetResponseWithDefaults instantiates a new WebhookVerificationKeyGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *WebhookVerificationKeyGetResponse) GetKey() JWKPublicKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WebhookVerificationKeyGetResponse) GetKeyOk() (*JWKPublicKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WebhookVerificationKeyGetResponse) SetKey(v JWKPublicKey)`

SetKey sets Key field to given value.


### GetRequestId

`func (o *WebhookVerificationKeyGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WebhookVerificationKeyGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WebhookVerificationKeyGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


