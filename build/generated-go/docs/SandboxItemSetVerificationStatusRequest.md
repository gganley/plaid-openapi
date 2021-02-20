# SandboxItemSetVerificationStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**AccountId** | **string** | The &#x60;account_id&#x60; of the account whose verification status is to be modified | 
**VerificationStatus** | **string** | The verification status to set the account to. | 

## Methods

### NewSandboxItemSetVerificationStatusRequest

`func NewSandboxItemSetVerificationStatusRequest(accessToken string, accountId string, verificationStatus string, ) *SandboxItemSetVerificationStatusRequest`

NewSandboxItemSetVerificationStatusRequest instantiates a new SandboxItemSetVerificationStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxItemSetVerificationStatusRequestWithDefaults

`func NewSandboxItemSetVerificationStatusRequestWithDefaults() *SandboxItemSetVerificationStatusRequest`

NewSandboxItemSetVerificationStatusRequestWithDefaults instantiates a new SandboxItemSetVerificationStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxItemSetVerificationStatusRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxItemSetVerificationStatusRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxItemSetVerificationStatusRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxItemSetVerificationStatusRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxItemSetVerificationStatusRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxItemSetVerificationStatusRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxItemSetVerificationStatusRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxItemSetVerificationStatusRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *SandboxItemSetVerificationStatusRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *SandboxItemSetVerificationStatusRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *SandboxItemSetVerificationStatusRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *SandboxItemSetVerificationStatusRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SandboxItemSetVerificationStatusRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SandboxItemSetVerificationStatusRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetVerificationStatus

`func (o *SandboxItemSetVerificationStatusRequest) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *SandboxItemSetVerificationStatusRequest) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *SandboxItemSetVerificationStatusRequest) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


