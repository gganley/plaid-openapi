# Employee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the employee. | [optional] 
**Address** | Pointer to [**NullableAddressData**](AddressData.md) |  | [optional] 
**SsnMasked** | Pointer to **NullableString** | The SSN of the employee, with all but the last 4 digits masked out. For example: \&quot;XXX-XX-1111\&quot;. | [optional] 

## Methods

### NewEmployee

`func NewEmployee() *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Employee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Employee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Employee) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Employee) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Employee) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Employee) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetAddress

`func (o *Employee) GetAddress() AddressData`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Employee) GetAddressOk() (*AddressData, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Employee) SetAddress(v AddressData)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Employee) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Employee) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Employee) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetSsnMasked

`func (o *Employee) GetSsnMasked() string`

GetSsnMasked returns the SsnMasked field if non-nil, zero value otherwise.

### GetSsnMaskedOk

`func (o *Employee) GetSsnMaskedOk() (*string, bool)`

GetSsnMaskedOk returns a tuple with the SsnMasked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsnMasked

`func (o *Employee) SetSsnMasked(v string)`

SetSsnMasked sets SsnMasked field to given value.

### HasSsnMasked

`func (o *Employee) HasSsnMasked() bool`

HasSsnMasked returns a boolean if a field has been set.

### SetSsnMaskedNil

`func (o *Employee) SetSsnMaskedNil(b bool)`

 SetSsnMaskedNil sets the value for SsnMasked to be an explicit nil

### UnsetSsnMasked
`func (o *Employee) UnsetSsnMasked()`

UnsetSsnMasked ensures that no value is present for SsnMasked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


