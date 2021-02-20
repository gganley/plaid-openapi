# ProcessorBalanceGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountBase**](AccountBase.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewProcessorBalanceGetResponse

`func NewProcessorBalanceGetResponse(account AccountBase, requestId string, ) *ProcessorBalanceGetResponse`

NewProcessorBalanceGetResponse instantiates a new ProcessorBalanceGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorBalanceGetResponseWithDefaults

`func NewProcessorBalanceGetResponseWithDefaults() *ProcessorBalanceGetResponse`

NewProcessorBalanceGetResponseWithDefaults instantiates a new ProcessorBalanceGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProcessorBalanceGetResponse) GetAccount() AccountBase`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProcessorBalanceGetResponse) GetAccountOk() (*AccountBase, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProcessorBalanceGetResponse) SetAccount(v AccountBase)`

SetAccount sets Account field to given value.


### GetRequestId

`func (o *ProcessorBalanceGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorBalanceGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorBalanceGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


