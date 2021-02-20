# IncomeVerificationPaystubGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paystub** | Pointer to [**Paystub**](Paystub.md) |  | [optional] 
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | [optional] 

## Methods

### NewIncomeVerificationPaystubGetResponse

`func NewIncomeVerificationPaystubGetResponse() *IncomeVerificationPaystubGetResponse`

NewIncomeVerificationPaystubGetResponse instantiates a new IncomeVerificationPaystubGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPaystubGetResponseWithDefaults

`func NewIncomeVerificationPaystubGetResponseWithDefaults() *IncomeVerificationPaystubGetResponse`

NewIncomeVerificationPaystubGetResponseWithDefaults instantiates a new IncomeVerificationPaystubGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaystub

`func (o *IncomeVerificationPaystubGetResponse) GetPaystub() Paystub`

GetPaystub returns the Paystub field if non-nil, zero value otherwise.

### GetPaystubOk

`func (o *IncomeVerificationPaystubGetResponse) GetPaystubOk() (*Paystub, bool)`

GetPaystubOk returns a tuple with the Paystub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystub

`func (o *IncomeVerificationPaystubGetResponse) SetPaystub(v Paystub)`

SetPaystub sets Paystub field to given value.

### HasPaystub

`func (o *IncomeVerificationPaystubGetResponse) HasPaystub() bool`

HasPaystub returns a boolean if a field has been set.

### GetRequestId

`func (o *IncomeVerificationPaystubGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationPaystubGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationPaystubGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *IncomeVerificationPaystubGetResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


