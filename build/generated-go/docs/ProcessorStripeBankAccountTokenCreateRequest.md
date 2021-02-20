# ProcessorStripeBankAccountTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**AccountId** | **string** | The &#x60;account_id&#x60; value obtained from the &#x60;onSuccess&#x60; callback in Link | 

## Methods

### NewProcessorStripeBankAccountTokenCreateRequest

`func NewProcessorStripeBankAccountTokenCreateRequest(accessToken string, accountId string, ) *ProcessorStripeBankAccountTokenCreateRequest`

NewProcessorStripeBankAccountTokenCreateRequest instantiates a new ProcessorStripeBankAccountTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStripeBankAccountTokenCreateRequestWithDefaults

`func NewProcessorStripeBankAccountTokenCreateRequestWithDefaults() *ProcessorStripeBankAccountTokenCreateRequest`

NewProcessorStripeBankAccountTokenCreateRequestWithDefaults instantiates a new ProcessorStripeBankAccountTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProcessorStripeBankAccountTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProcessorStripeBankAccountTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProcessorStripeBankAccountTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProcessorStripeBankAccountTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ProcessorStripeBankAccountTokenCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ProcessorStripeBankAccountTokenCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ProcessorStripeBankAccountTokenCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


