# ProcessorApexProcessorTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**AccountId** | **string** | The &#x60;account_id&#x60; value obtained from the &#x60;onSuccess&#x60; callback in Link | 

## Methods

### NewProcessorApexProcessorTokenCreateRequest

`func NewProcessorApexProcessorTokenCreateRequest(accessToken string, accountId string, ) *ProcessorApexProcessorTokenCreateRequest`

NewProcessorApexProcessorTokenCreateRequest instantiates a new ProcessorApexProcessorTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorApexProcessorTokenCreateRequestWithDefaults

`func NewProcessorApexProcessorTokenCreateRequestWithDefaults() *ProcessorApexProcessorTokenCreateRequest`

NewProcessorApexProcessorTokenCreateRequestWithDefaults instantiates a new ProcessorApexProcessorTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ProcessorApexProcessorTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProcessorApexProcessorTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProcessorApexProcessorTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProcessorApexProcessorTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ProcessorApexProcessorTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProcessorApexProcessorTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProcessorApexProcessorTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProcessorApexProcessorTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *ProcessorApexProcessorTokenCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ProcessorApexProcessorTokenCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ProcessorApexProcessorTokenCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *ProcessorApexProcessorTokenCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ProcessorApexProcessorTokenCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ProcessorApexProcessorTokenCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


