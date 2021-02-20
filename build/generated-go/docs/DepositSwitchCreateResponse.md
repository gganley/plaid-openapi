# DepositSwitchCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositSwitchId** | **string** | ID of the deposit switch. This ID is persisted throughout the lifetime of the deposit switch. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewDepositSwitchCreateResponse

`func NewDepositSwitchCreateResponse(depositSwitchId string, requestId string, ) *DepositSwitchCreateResponse`

NewDepositSwitchCreateResponse instantiates a new DepositSwitchCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchCreateResponseWithDefaults

`func NewDepositSwitchCreateResponseWithDefaults() *DepositSwitchCreateResponse`

NewDepositSwitchCreateResponseWithDefaults instantiates a new DepositSwitchCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositSwitchId

`func (o *DepositSwitchCreateResponse) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchCreateResponse) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchCreateResponse) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.


### GetRequestId

`func (o *DepositSwitchCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DepositSwitchCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DepositSwitchCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


