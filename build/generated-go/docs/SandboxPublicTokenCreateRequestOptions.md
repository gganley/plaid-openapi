# SandboxPublicTokenCreateRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | Pointer to **string** | Specify a webhook to associate with the new Item. | [optional] 
**OverrideUsername** | Pointer to **NullableString** | Test username to use for the creation of the Sandbox Item. Default value is &#x60;user_good&#x60;. | [optional] [default to "user_good"]
**OverridePassword** | Pointer to **NullableString** | Test password to use for the creation of the Sandbox Item. Default value is &#x60;pass_good&#x60;. | [optional] [default to "pass_good"]
**Transactions** | Pointer to [**SandboxPublicTokenCreateRequestOptionsTransactions**](SandboxPublicTokenCreateRequestOptionsTransactions.md) |  | [optional] 

## Methods

### NewSandboxPublicTokenCreateRequestOptions

`func NewSandboxPublicTokenCreateRequestOptions() *SandboxPublicTokenCreateRequestOptions`

NewSandboxPublicTokenCreateRequestOptions instantiates a new SandboxPublicTokenCreateRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPublicTokenCreateRequestOptionsWithDefaults

`func NewSandboxPublicTokenCreateRequestOptionsWithDefaults() *SandboxPublicTokenCreateRequestOptions`

NewSandboxPublicTokenCreateRequestOptionsWithDefaults instantiates a new SandboxPublicTokenCreateRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhook

`func (o *SandboxPublicTokenCreateRequestOptions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *SandboxPublicTokenCreateRequestOptions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *SandboxPublicTokenCreateRequestOptions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *SandboxPublicTokenCreateRequestOptions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetOverrideUsername

`func (o *SandboxPublicTokenCreateRequestOptions) GetOverrideUsername() string`

GetOverrideUsername returns the OverrideUsername field if non-nil, zero value otherwise.

### GetOverrideUsernameOk

`func (o *SandboxPublicTokenCreateRequestOptions) GetOverrideUsernameOk() (*string, bool)`

GetOverrideUsernameOk returns a tuple with the OverrideUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideUsername

`func (o *SandboxPublicTokenCreateRequestOptions) SetOverrideUsername(v string)`

SetOverrideUsername sets OverrideUsername field to given value.

### HasOverrideUsername

`func (o *SandboxPublicTokenCreateRequestOptions) HasOverrideUsername() bool`

HasOverrideUsername returns a boolean if a field has been set.

### SetOverrideUsernameNil

`func (o *SandboxPublicTokenCreateRequestOptions) SetOverrideUsernameNil(b bool)`

 SetOverrideUsernameNil sets the value for OverrideUsername to be an explicit nil

### UnsetOverrideUsername
`func (o *SandboxPublicTokenCreateRequestOptions) UnsetOverrideUsername()`

UnsetOverrideUsername ensures that no value is present for OverrideUsername, not even an explicit nil
### GetOverridePassword

`func (o *SandboxPublicTokenCreateRequestOptions) GetOverridePassword() string`

GetOverridePassword returns the OverridePassword field if non-nil, zero value otherwise.

### GetOverridePasswordOk

`func (o *SandboxPublicTokenCreateRequestOptions) GetOverridePasswordOk() (*string, bool)`

GetOverridePasswordOk returns a tuple with the OverridePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePassword

`func (o *SandboxPublicTokenCreateRequestOptions) SetOverridePassword(v string)`

SetOverridePassword sets OverridePassword field to given value.

### HasOverridePassword

`func (o *SandboxPublicTokenCreateRequestOptions) HasOverridePassword() bool`

HasOverridePassword returns a boolean if a field has been set.

### SetOverridePasswordNil

`func (o *SandboxPublicTokenCreateRequestOptions) SetOverridePasswordNil(b bool)`

 SetOverridePasswordNil sets the value for OverridePassword to be an explicit nil

### UnsetOverridePassword
`func (o *SandboxPublicTokenCreateRequestOptions) UnsetOverridePassword()`

UnsetOverridePassword ensures that no value is present for OverridePassword, not even an explicit nil
### GetTransactions

`func (o *SandboxPublicTokenCreateRequestOptions) GetTransactions() SandboxPublicTokenCreateRequestOptionsTransactions`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *SandboxPublicTokenCreateRequestOptions) GetTransactionsOk() (*SandboxPublicTokenCreateRequestOptionsTransactions, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *SandboxPublicTokenCreateRequestOptions) SetTransactions(v SandboxPublicTokenCreateRequestOptionsTransactions)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *SandboxPublicTokenCreateRequestOptions) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


