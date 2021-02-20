# PaystubDeduction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **NullableString** | The description of the deduction, as provided on the paystub. For example: &#x60;\&quot;401(k)\&quot;&#x60;, &#x60;\&quot;FICA MED TAX\&quot;&#x60;. | [optional] 
**IsPretax** | Pointer to **NullableBool** | &#x60;true&#x60; if the deduction is pre-tax; &#x60;false&#x60; otherwise. | [optional] 
**Total** | Pointer to **NullableFloat32** | The amount of the deduction. | [optional] 

## Methods

### NewPaystubDeduction

`func NewPaystubDeduction() *PaystubDeduction`

NewPaystubDeduction instantiates a new PaystubDeduction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubDeductionWithDefaults

`func NewPaystubDeductionWithDefaults() *PaystubDeduction`

NewPaystubDeductionWithDefaults instantiates a new PaystubDeduction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaystubDeduction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaystubDeduction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaystubDeduction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaystubDeduction) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PaystubDeduction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PaystubDeduction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetIsPretax

`func (o *PaystubDeduction) GetIsPretax() bool`

GetIsPretax returns the IsPretax field if non-nil, zero value otherwise.

### GetIsPretaxOk

`func (o *PaystubDeduction) GetIsPretaxOk() (*bool, bool)`

GetIsPretaxOk returns a tuple with the IsPretax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPretax

`func (o *PaystubDeduction) SetIsPretax(v bool)`

SetIsPretax sets IsPretax field to given value.

### HasIsPretax

`func (o *PaystubDeduction) HasIsPretax() bool`

HasIsPretax returns a boolean if a field has been set.

### SetIsPretaxNil

`func (o *PaystubDeduction) SetIsPretaxNil(b bool)`

 SetIsPretaxNil sets the value for IsPretax to be an explicit nil

### UnsetIsPretax
`func (o *PaystubDeduction) UnsetIsPretax()`

UnsetIsPretax ensures that no value is present for IsPretax, not even an explicit nil
### GetTotal

`func (o *PaystubDeduction) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaystubDeduction) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaystubDeduction) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaystubDeduction) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *PaystubDeduction) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *PaystubDeduction) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


