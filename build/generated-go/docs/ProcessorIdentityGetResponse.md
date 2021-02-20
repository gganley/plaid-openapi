# ProcessorIdentityGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountIdentity**](AccountIdentity.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewProcessorIdentityGetResponse

`func NewProcessorIdentityGetResponse(account AccountIdentity, requestId string, ) *ProcessorIdentityGetResponse`

NewProcessorIdentityGetResponse instantiates a new ProcessorIdentityGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorIdentityGetResponseWithDefaults

`func NewProcessorIdentityGetResponseWithDefaults() *ProcessorIdentityGetResponse`

NewProcessorIdentityGetResponseWithDefaults instantiates a new ProcessorIdentityGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProcessorIdentityGetResponse) GetAccount() AccountIdentity`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProcessorIdentityGetResponse) GetAccountOk() (*AccountIdentity, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProcessorIdentityGetResponse) SetAccount(v AccountIdentity)`

SetAccount sets Account field to given value.


### GetRequestId

`func (o *ProcessorIdentityGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorIdentityGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorIdentityGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


