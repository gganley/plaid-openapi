# DepositSwitchAltCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositSwitchId** | **string** | ID of the deposit switch. This ID is persisted throughout the lifetime of the deposit switch. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewDepositSwitchAltCreateResponse

`func NewDepositSwitchAltCreateResponse(depositSwitchId string, requestId string, ) *DepositSwitchAltCreateResponse`

NewDepositSwitchAltCreateResponse instantiates a new DepositSwitchAltCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchAltCreateResponseWithDefaults

`func NewDepositSwitchAltCreateResponseWithDefaults() *DepositSwitchAltCreateResponse`

NewDepositSwitchAltCreateResponseWithDefaults instantiates a new DepositSwitchAltCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositSwitchId

`func (o *DepositSwitchAltCreateResponse) GetDepositSwitchId() string`

GetDepositSwitchId returns the DepositSwitchId field if non-nil, zero value otherwise.

### GetDepositSwitchIdOk

`func (o *DepositSwitchAltCreateResponse) GetDepositSwitchIdOk() (*string, bool)`

GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchId

`func (o *DepositSwitchAltCreateResponse) SetDepositSwitchId(v string)`

SetDepositSwitchId sets DepositSwitchId field to given value.


### GetRequestId

`func (o *DepositSwitchAltCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DepositSwitchAltCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DepositSwitchAltCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


