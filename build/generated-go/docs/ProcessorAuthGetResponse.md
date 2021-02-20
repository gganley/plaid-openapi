# ProcessorAuthGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**Numbers** | [**ProcessorNumber**](ProcessorNumber.md) |  | 
**Account** | [**AccountBase**](AccountBase.md) |  | 

## Methods

### NewProcessorAuthGetResponse

`func NewProcessorAuthGetResponse(requestId string, numbers ProcessorNumber, account AccountBase, ) *ProcessorAuthGetResponse`

NewProcessorAuthGetResponse instantiates a new ProcessorAuthGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorAuthGetResponseWithDefaults

`func NewProcessorAuthGetResponseWithDefaults() *ProcessorAuthGetResponse`

NewProcessorAuthGetResponseWithDefaults instantiates a new ProcessorAuthGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ProcessorAuthGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorAuthGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorAuthGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetNumbers

`func (o *ProcessorAuthGetResponse) GetNumbers() ProcessorNumber`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *ProcessorAuthGetResponse) GetNumbersOk() (*ProcessorNumber, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *ProcessorAuthGetResponse) SetNumbers(v ProcessorNumber)`

SetNumbers sets Numbers field to given value.


### GetAccount

`func (o *ProcessorAuthGetResponse) GetAccount() AccountBase`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProcessorAuthGetResponse) GetAccountOk() (*AccountBase, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProcessorAuthGetResponse) SetAccount(v AccountBase)`

SetAccount sets Account field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


