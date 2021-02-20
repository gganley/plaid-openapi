# ProcessorTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessorToken** | **string** | The &#x60;processor_token&#x60; that can then be used by the Plaid partner to make API requests | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewProcessorTokenCreateResponse

`func NewProcessorTokenCreateResponse(processorToken string, requestId string, ) *ProcessorTokenCreateResponse`

NewProcessorTokenCreateResponse instantiates a new ProcessorTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorTokenCreateResponseWithDefaults

`func NewProcessorTokenCreateResponseWithDefaults() *ProcessorTokenCreateResponse`

NewProcessorTokenCreateResponseWithDefaults instantiates a new ProcessorTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessorToken

`func (o *ProcessorTokenCreateResponse) GetProcessorToken() string`

GetProcessorToken returns the ProcessorToken field if non-nil, zero value otherwise.

### GetProcessorTokenOk

`func (o *ProcessorTokenCreateResponse) GetProcessorTokenOk() (*string, bool)`

GetProcessorTokenOk returns a tuple with the ProcessorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorToken

`func (o *ProcessorTokenCreateResponse) SetProcessorToken(v string)`

SetProcessorToken sets ProcessorToken field to given value.


### GetRequestId

`func (o *ProcessorTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


