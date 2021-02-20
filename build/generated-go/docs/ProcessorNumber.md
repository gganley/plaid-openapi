# ProcessorNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ach** | Pointer to [**NullableNumbersACH**](NumbersACH.md) |  | [optional] 
**Eft** | Pointer to [**NullableNumbersEFT**](NumbersEFT.md) |  | [optional] 
**International** | Pointer to [**NullableNumbersInternational**](NumbersInternational.md) |  | [optional] 
**Bacs** | Pointer to [**NullableNumbersBACS**](NumbersBACS.md) |  | [optional] 

## Methods

### NewProcessorNumber

`func NewProcessorNumber() *ProcessorNumber`

NewProcessorNumber instantiates a new ProcessorNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorNumberWithDefaults

`func NewProcessorNumberWithDefaults() *ProcessorNumber`

NewProcessorNumberWithDefaults instantiates a new ProcessorNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAch

`func (o *ProcessorNumber) GetAch() NumbersACH`

GetAch returns the Ach field if non-nil, zero value otherwise.

### GetAchOk

`func (o *ProcessorNumber) GetAchOk() (*NumbersACH, bool)`

GetAchOk returns a tuple with the Ach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAch

`func (o *ProcessorNumber) SetAch(v NumbersACH)`

SetAch sets Ach field to given value.

### HasAch

`func (o *ProcessorNumber) HasAch() bool`

HasAch returns a boolean if a field has been set.

### SetAchNil

`func (o *ProcessorNumber) SetAchNil(b bool)`

 SetAchNil sets the value for Ach to be an explicit nil

### UnsetAch
`func (o *ProcessorNumber) UnsetAch()`

UnsetAch ensures that no value is present for Ach, not even an explicit nil
### GetEft

`func (o *ProcessorNumber) GetEft() NumbersEFT`

GetEft returns the Eft field if non-nil, zero value otherwise.

### GetEftOk

`func (o *ProcessorNumber) GetEftOk() (*NumbersEFT, bool)`

GetEftOk returns a tuple with the Eft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEft

`func (o *ProcessorNumber) SetEft(v NumbersEFT)`

SetEft sets Eft field to given value.

### HasEft

`func (o *ProcessorNumber) HasEft() bool`

HasEft returns a boolean if a field has been set.

### SetEftNil

`func (o *ProcessorNumber) SetEftNil(b bool)`

 SetEftNil sets the value for Eft to be an explicit nil

### UnsetEft
`func (o *ProcessorNumber) UnsetEft()`

UnsetEft ensures that no value is present for Eft, not even an explicit nil
### GetInternational

`func (o *ProcessorNumber) GetInternational() NumbersInternational`

GetInternational returns the International field if non-nil, zero value otherwise.

### GetInternationalOk

`func (o *ProcessorNumber) GetInternationalOk() (*NumbersInternational, bool)`

GetInternationalOk returns a tuple with the International field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternational

`func (o *ProcessorNumber) SetInternational(v NumbersInternational)`

SetInternational sets International field to given value.

### HasInternational

`func (o *ProcessorNumber) HasInternational() bool`

HasInternational returns a boolean if a field has been set.

### SetInternationalNil

`func (o *ProcessorNumber) SetInternationalNil(b bool)`

 SetInternationalNil sets the value for International to be an explicit nil

### UnsetInternational
`func (o *ProcessorNumber) UnsetInternational()`

UnsetInternational ensures that no value is present for International, not even an explicit nil
### GetBacs

`func (o *ProcessorNumber) GetBacs() NumbersBACS`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *ProcessorNumber) GetBacsOk() (*NumbersBACS, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *ProcessorNumber) SetBacs(v NumbersBACS)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *ProcessorNumber) HasBacs() bool`

HasBacs returns a boolean if a field has been set.

### SetBacsNil

`func (o *ProcessorNumber) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *ProcessorNumber) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


