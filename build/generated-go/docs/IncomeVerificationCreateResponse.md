# IncomeVerificationCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeVerificationId** | **string** | ID of the verification. This ID is persisted throughout the lifetime of the verification. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewIncomeVerificationCreateResponse

`func NewIncomeVerificationCreateResponse(incomeVerificationId string, requestId string, ) *IncomeVerificationCreateResponse`

NewIncomeVerificationCreateResponse instantiates a new IncomeVerificationCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationCreateResponseWithDefaults

`func NewIncomeVerificationCreateResponseWithDefaults() *IncomeVerificationCreateResponse`

NewIncomeVerificationCreateResponseWithDefaults instantiates a new IncomeVerificationCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeVerificationId

`func (o *IncomeVerificationCreateResponse) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationCreateResponse) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationCreateResponse) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.


### GetRequestId

`func (o *IncomeVerificationCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


