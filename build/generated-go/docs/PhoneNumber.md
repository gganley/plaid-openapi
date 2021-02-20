# PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The phone number. | 
**Primary** | Pointer to **NullableBool** | When &#x60;true&#x60;, identifies the phone number as the primary number on an account. | [optional] 
**Type** | Pointer to **NullableString** | The type of phone number. | [optional] 

## Methods

### NewPhoneNumber

`func NewPhoneNumber(data string, ) *PhoneNumber`

NewPhoneNumber instantiates a new PhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithDefaults

`func NewPhoneNumberWithDefaults() *PhoneNumber`

NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PhoneNumber) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PhoneNumber) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PhoneNumber) SetData(v string)`

SetData sets Data field to given value.


### GetPrimary

`func (o *PhoneNumber) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PhoneNumber) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PhoneNumber) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PhoneNumber) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *PhoneNumber) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *PhoneNumber) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetType

`func (o *PhoneNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneNumber) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhoneNumber) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *PhoneNumber) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *PhoneNumber) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


