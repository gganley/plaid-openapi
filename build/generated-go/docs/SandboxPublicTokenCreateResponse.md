# SandboxPublicTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicToken** | **string** | A public token that can be exchanged for an access token using &#x60;/item/public_token/exchange&#x60; | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewSandboxPublicTokenCreateResponse

`func NewSandboxPublicTokenCreateResponse(publicToken string, requestId string, ) *SandboxPublicTokenCreateResponse`

NewSandboxPublicTokenCreateResponse instantiates a new SandboxPublicTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPublicTokenCreateResponseWithDefaults

`func NewSandboxPublicTokenCreateResponseWithDefaults() *SandboxPublicTokenCreateResponse`

NewSandboxPublicTokenCreateResponseWithDefaults instantiates a new SandboxPublicTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicToken

`func (o *SandboxPublicTokenCreateResponse) GetPublicToken() string`

GetPublicToken returns the PublicToken field if non-nil, zero value otherwise.

### GetPublicTokenOk

`func (o *SandboxPublicTokenCreateResponse) GetPublicTokenOk() (*string, bool)`

GetPublicTokenOk returns a tuple with the PublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicToken

`func (o *SandboxPublicTokenCreateResponse) SetPublicToken(v string)`

SetPublicToken sets PublicToken field to given value.


### GetRequestId

`func (o *SandboxPublicTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SandboxPublicTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SandboxPublicTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


