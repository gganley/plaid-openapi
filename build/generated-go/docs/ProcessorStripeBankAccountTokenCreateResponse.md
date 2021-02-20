# ProcessorStripeBankAccountTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StripeBankAccountToken** | **string** | A token that can be sent to Stripe for use in making API calls to Plaid | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewProcessorStripeBankAccountTokenCreateResponse

`func NewProcessorStripeBankAccountTokenCreateResponse(stripeBankAccountToken string, requestId string, ) *ProcessorStripeBankAccountTokenCreateResponse`

NewProcessorStripeBankAccountTokenCreateResponse instantiates a new ProcessorStripeBankAccountTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorStripeBankAccountTokenCreateResponseWithDefaults

`func NewProcessorStripeBankAccountTokenCreateResponseWithDefaults() *ProcessorStripeBankAccountTokenCreateResponse`

NewProcessorStripeBankAccountTokenCreateResponseWithDefaults instantiates a new ProcessorStripeBankAccountTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStripeBankAccountToken

`func (o *ProcessorStripeBankAccountTokenCreateResponse) GetStripeBankAccountToken() string`

GetStripeBankAccountToken returns the StripeBankAccountToken field if non-nil, zero value otherwise.

### GetStripeBankAccountTokenOk

`func (o *ProcessorStripeBankAccountTokenCreateResponse) GetStripeBankAccountTokenOk() (*string, bool)`

GetStripeBankAccountTokenOk returns a tuple with the StripeBankAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeBankAccountToken

`func (o *ProcessorStripeBankAccountTokenCreateResponse) SetStripeBankAccountToken(v string)`

SetStripeBankAccountToken sets StripeBankAccountToken field to given value.


### GetRequestId

`func (o *ProcessorStripeBankAccountTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorStripeBankAccountTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorStripeBankAccountTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


