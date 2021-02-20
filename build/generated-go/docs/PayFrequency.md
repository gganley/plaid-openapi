# PayFrequency

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The frequency of the pay period. | 
**VerificationStatus** | [**VerificationStatus**](VerificationStatus.md) |  | 

## Methods

### NewPayFrequency

`func NewPayFrequency(value string, verificationStatus VerificationStatus, ) *PayFrequency`

NewPayFrequency instantiates a new PayFrequency object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayFrequencyWithDefaults

`func NewPayFrequencyWithDefaults() *PayFrequency`

NewPayFrequencyWithDefaults instantiates a new PayFrequency object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *PayFrequency) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PayFrequency) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PayFrequency) SetValue(v string)`

SetValue sets Value field to given value.


### GetVerificationStatus

`func (o *PayFrequency) GetVerificationStatus() VerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *PayFrequency) GetVerificationStatusOk() (*VerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *PayFrequency) SetVerificationStatus(v VerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


