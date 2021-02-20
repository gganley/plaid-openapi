# ProcessorTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**AccountId** | **string** | The &#x60;account_id&#x60; value obtained from the &#x60;onSuccess&#x60; callback in Link | 
**Processor** | **string** | The processor you are integrating with. Valid values are &#x60;\&quot;achq\&quot;&#x60;, &#x60;\&quot;check\&quot;&#x60;, &#x60;\&quot;checkbook\&quot;&#x60;, &#x60;\&quot;circle\&quot;&#x60;, &#x60;\&quot;drivewealth\&quot;&#x60;, &#x60;\&quot;dwolla\&quot;&#x60;, &#x60;\&quot;galileo\&quot;&#x60;, \&quot;&#x60;interactive_brokers&#x60;\&quot;, &#x60;\&quot;modern_treasury\&quot;&#x60;, &#x60;\&quot;ocrolus\&quot;&#x60;, &#x60;\&quot;prime_trust\&quot;&#x60;, &#x60;\&quot;rize\&quot;&#x60;, &#x60;\&quot;sila_money\&quot;&#x60;, &#x60;\&quot;unit\&quot;&#x60;, &#x60;\&quot;velox\&quot;&#x60;, &#x60;\&quot;vesta\&quot;&#x60;, &#x60;\&quot;vopay\&quot;&#x60;, &#x60;\&quot;wyre\&quot;&#x60; | 

## Methods

### NewProcessorTokenCreateRequest

`func NewProcessorTokenCreateRequest(accessToken string, accountId string, processor string, ) *ProcessorTokenCreateRequest`

NewProcessorTokenCreateRequest instantiates a new ProcessorTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorTokenCreateRequestWithDefaults

`func NewProcessorTokenCreateRequestWithDefaults() *ProcessorTokenCreateRequest`

NewProcessorTokenCreateRequestWithDefaults instantiates a new ProcessorTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ProcessorTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProcessorTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProcessorTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProcessorTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ProcessorTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProcessorTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProcessorTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProcessorTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *ProcessorTokenCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ProcessorTokenCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ProcessorTokenCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *ProcessorTokenCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ProcessorTokenCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ProcessorTokenCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProcessor

`func (o *ProcessorTokenCreateRequest) GetProcessor() string`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *ProcessorTokenCreateRequest) GetProcessorOk() (*string, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *ProcessorTokenCreateRequest) SetProcessor(v string)`

SetProcessor sets Processor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


