# MortgageInterestRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | Pointer to **NullableFloat32** | Percentage value (interest rate of current mortgage, not APR) of interest payable on a loan. | [optional] 
**Type** | Pointer to **NullableString** | The type of interest charged (fixed or variable). | [optional] 

## Methods

### NewMortgageInterestRate

`func NewMortgageInterestRate() *MortgageInterestRate`

NewMortgageInterestRate instantiates a new MortgageInterestRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMortgageInterestRateWithDefaults

`func NewMortgageInterestRateWithDefaults() *MortgageInterestRate`

NewMortgageInterestRateWithDefaults instantiates a new MortgageInterestRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *MortgageInterestRate) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *MortgageInterestRate) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *MortgageInterestRate) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *MortgageInterestRate) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *MortgageInterestRate) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *MortgageInterestRate) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetType

`func (o *MortgageInterestRate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MortgageInterestRate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MortgageInterestRate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MortgageInterestRate) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MortgageInterestRate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MortgageInterestRate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


