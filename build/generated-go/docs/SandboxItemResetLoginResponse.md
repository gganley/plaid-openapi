# SandboxItemResetLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResetLogin** | **bool** | &#x60;true&#x60; if the call succeeded | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewSandboxItemResetLoginResponse

`func NewSandboxItemResetLoginResponse(resetLogin bool, requestId string, ) *SandboxItemResetLoginResponse`

NewSandboxItemResetLoginResponse instantiates a new SandboxItemResetLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxItemResetLoginResponseWithDefaults

`func NewSandboxItemResetLoginResponseWithDefaults() *SandboxItemResetLoginResponse`

NewSandboxItemResetLoginResponseWithDefaults instantiates a new SandboxItemResetLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResetLogin

`func (o *SandboxItemResetLoginResponse) GetResetLogin() bool`

GetResetLogin returns the ResetLogin field if non-nil, zero value otherwise.

### GetResetLoginOk

`func (o *SandboxItemResetLoginResponse) GetResetLoginOk() (*bool, bool)`

GetResetLoginOk returns a tuple with the ResetLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetLogin

`func (o *SandboxItemResetLoginResponse) SetResetLogin(v bool)`

SetResetLogin sets ResetLogin field to given value.


### GetRequestId

`func (o *SandboxItemResetLoginResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SandboxItemResetLoginResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SandboxItemResetLoginResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


