# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NullableAddressData**](AddressData.md) |  | 
**Primary** | Pointer to **NullableBool** | When &#x60;true&#x60;, identifies the address as the primary address on an account. | [optional] 

## Methods

### NewAddress

`func NewAddress(data NullableAddressData, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Address) GetData() AddressData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Address) GetDataOk() (*AddressData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Address) SetData(v AddressData)`

SetData sets Data field to given value.


### SetDataNil

`func (o *Address) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Address) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetPrimary

`func (o *Address) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Address) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Address) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *Address) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *Address) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *Address) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


