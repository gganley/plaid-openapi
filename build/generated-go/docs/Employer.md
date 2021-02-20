# Employer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerId** | **string** | Plaid&#39;s unique identifier for the employer. | 
**Name** | **string** | The name of the employer | 
**Address** | Pointer to [**NullableAddressData**](AddressData.md) |  | [optional] 
**ConfidenceScore** | Pointer to **NullableFloat32** | A number from 0 to 1 indicating Plaid&#39;s level of confidence in the pairing between the employer and the institution (not yet implemented). | [optional] 

## Methods

### NewEmployer

`func NewEmployer(employerId string, name string, ) *Employer`

NewEmployer instantiates a new Employer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerWithDefaults

`func NewEmployerWithDefaults() *Employer`

NewEmployerWithDefaults instantiates a new Employer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerId

`func (o *Employer) GetEmployerId() string`

GetEmployerId returns the EmployerId field if non-nil, zero value otherwise.

### GetEmployerIdOk

`func (o *Employer) GetEmployerIdOk() (*string, bool)`

GetEmployerIdOk returns a tuple with the EmployerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerId

`func (o *Employer) SetEmployerId(v string)`

SetEmployerId sets EmployerId field to given value.


### GetName

`func (o *Employer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Employer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Employer) SetName(v string)`

SetName sets Name field to given value.


### GetAddress

`func (o *Employer) GetAddress() AddressData`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Employer) GetAddressOk() (*AddressData, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Employer) SetAddress(v AddressData)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Employer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *Employer) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Employer) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetConfidenceScore

`func (o *Employer) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *Employer) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *Employer) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.

### HasConfidenceScore

`func (o *Employer) HasConfidenceScore() bool`

HasConfidenceScore returns a boolean if a field has been set.

### SetConfidenceScoreNil

`func (o *Employer) SetConfidenceScoreNil(b bool)`

 SetConfidenceScoreNil sets the value for ConfidenceScore to be an explicit nil

### UnsetConfidenceScore
`func (o *Employer) UnsetConfidenceScore()`

UnsetConfidenceScore ensures that no value is present for ConfidenceScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


