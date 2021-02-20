# ProcessorBalanceGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**ProcessorToken** | **string** | The processor token obtained from the Plaid integration partner. Processor tokens are in the format: &#x60;processor-&lt;environment&gt;-&lt;identifier&gt;&#x60; | 

## Methods

### NewProcessorBalanceGetRequest

`func NewProcessorBalanceGetRequest(processorToken string, ) *ProcessorBalanceGetRequest`

NewProcessorBalanceGetRequest instantiates a new ProcessorBalanceGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorBalanceGetRequestWithDefaults

`func NewProcessorBalanceGetRequestWithDefaults() *ProcessorBalanceGetRequest`

NewProcessorBalanceGetRequestWithDefaults instantiates a new ProcessorBalanceGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ProcessorBalanceGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProcessorBalanceGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProcessorBalanceGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProcessorBalanceGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ProcessorBalanceGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProcessorBalanceGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProcessorBalanceGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProcessorBalanceGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetProcessorToken

`func (o *ProcessorBalanceGetRequest) GetProcessorToken() string`

GetProcessorToken returns the ProcessorToken field if non-nil, zero value otherwise.

### GetProcessorTokenOk

`func (o *ProcessorBalanceGetRequest) GetProcessorTokenOk() (*string, bool)`

GetProcessorTokenOk returns a tuple with the ProcessorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorToken

`func (o *ProcessorBalanceGetRequest) SetProcessorToken(v string)`

SetProcessorToken sets ProcessorToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


