# SandboxProcessorTokenCreateRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverrideUsername** | Pointer to **NullableString** | Test username to use for the creation of the Sandbox Item. Default value is &#x60;user_good&#x60;. | [optional] [default to "user_good"]
**OverridePassword** | Pointer to **NullableString** | Test password to use for the creation of the Sandbox Item. Default value is &#x60;pass_good&#x60;. | [optional] [default to "pass_good"]

## Methods

### NewSandboxProcessorTokenCreateRequestOptions

`func NewSandboxProcessorTokenCreateRequestOptions() *SandboxProcessorTokenCreateRequestOptions`

NewSandboxProcessorTokenCreateRequestOptions instantiates a new SandboxProcessorTokenCreateRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxProcessorTokenCreateRequestOptionsWithDefaults

`func NewSandboxProcessorTokenCreateRequestOptionsWithDefaults() *SandboxProcessorTokenCreateRequestOptions`

NewSandboxProcessorTokenCreateRequestOptionsWithDefaults instantiates a new SandboxProcessorTokenCreateRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverrideUsername

`func (o *SandboxProcessorTokenCreateRequestOptions) GetOverrideUsername() string`

GetOverrideUsername returns the OverrideUsername field if non-nil, zero value otherwise.

### GetOverrideUsernameOk

`func (o *SandboxProcessorTokenCreateRequestOptions) GetOverrideUsernameOk() (*string, bool)`

GetOverrideUsernameOk returns a tuple with the OverrideUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUsername

`func (o *SandboxProcessorTokenCreateRequestOptions) SetOverrideUsername(v string)`

SetOverrideUsername sets OverrideUsername field to given value.

### HasOverrideUsername

`func (o *SandboxProcessorTokenCreateRequestOptions) HasOverrideUsername() bool`

HasOverrideUsername returns a boolean if a field has been set.

### SetOverrideUsernameNil

`func (o *SandboxProcessorTokenCreateRequestOptions) SetOverrideUsernameNil(b bool)`

 SetOverrideUsernameNil sets the value for OverrideUsername to be an explicit nil

### UnsetOverrideUsername
`func (o *SandboxProcessorTokenCreateRequestOptions) UnsetOverrideUsername()`

UnsetOverrideUsername ensures that no value is present for OverrideUsername, not even an explicit nil
### GetOverridePassword

`func (o *SandboxProcessorTokenCreateRequestOptions) GetOverridePassword() string`

GetOverridePassword returns the OverridePassword field if non-nil, zero value otherwise.

### GetOverridePasswordOk

`func (o *SandboxProcessorTokenCreateRequestOptions) GetOverridePasswordOk() (*string, bool)`

GetOverridePasswordOk returns a tuple with the OverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePassword

`func (o *SandboxProcessorTokenCreateRequestOptions) SetOverridePassword(v string)`

SetOverridePassword sets OverridePassword field to given value.

### HasOverridePassword

`func (o *SandboxProcessorTokenCreateRequestOptions) HasOverridePassword() bool`

HasOverridePassword returns a boolean if a field has been set.

### SetOverridePasswordNil

`func (o *SandboxProcessorTokenCreateRequestOptions) SetOverridePasswordNil(b bool)`

 SetOverridePasswordNil sets the value for OverridePassword to be an explicit nil

### UnsetOverridePassword
`func (o *SandboxProcessorTokenCreateRequestOptions) UnsetOverridePassword()`

UnsetOverridePassword ensures that no value is present for OverridePassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


