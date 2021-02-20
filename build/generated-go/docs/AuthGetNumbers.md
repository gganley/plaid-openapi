# AuthGetNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ach** | Pointer to [**[]NumbersACH**](NumbersACH.md) | An array of ACH numbers identifying accounts. | [optional] 
**Eft** | Pointer to [**[]NumbersEFT**](NumbersEFT.md) | An array of EFT numbers identifying accounts. | [optional] 
**International** | Pointer to [**[]NumbersInternational**](NumbersInternational.md) | An array of IBAN numbers identifying accounts. | [optional] 
**Bacs** | Pointer to [**[]NumbersBACS**](NumbersBACS.md) | An array of BACS numbers identifying accounts. | [optional] 

## Methods

### NewAuthGetNumbers

`func NewAuthGetNumbers() *AuthGetNumbers`

NewAuthGetNumbers instantiates a new AuthGetNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetNumbersWithDefaults

`func NewAuthGetNumbersWithDefaults() *AuthGetNumbers`

NewAuthGetNumbersWithDefaults instantiates a new AuthGetNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAch

`func (o *AuthGetNumbers) GetAch() []NumbersACH`

GetAch returns the Ach field if non-nil, zero value otherwise.

### GetAchOk

`func (o *AuthGetNumbers) GetAchOk() (*[]NumbersACH, bool)`

GetAchOk returns a tuple with the Ach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAch

`func (o *AuthGetNumbers) SetAch(v []NumbersACH)`

SetAch sets Ach field to given value.

### HasAch

`func (o *AuthGetNumbers) HasAch() bool`

HasAch returns a boolean if a field has been set.

### GetEft

`func (o *AuthGetNumbers) GetEft() []NumbersEFT`

GetEft returns the Eft field if non-nil, zero value otherwise.

### GetEftOk

`func (o *AuthGetNumbers) GetEftOk() (*[]NumbersEFT, bool)`

GetEftOk returns a tuple with the Eft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEft

`func (o *AuthGetNumbers) SetEft(v []NumbersEFT)`

SetEft sets Eft field to given value.

### HasEft

`func (o *AuthGetNumbers) HasEft() bool`

HasEft returns a boolean if a field has been set.

### GetInternational

`func (o *AuthGetNumbers) GetInternational() []NumbersInternational`

GetInternational returns the International field if non-nil, zero value otherwise.

### GetInternationalOk

`func (o *AuthGetNumbers) GetInternationalOk() (*[]NumbersInternational, bool)`

GetInternationalOk returns a tuple with the International field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternational

`func (o *AuthGetNumbers) SetInternational(v []NumbersInternational)`

SetInternational sets International field to given value.

### HasInternational

`func (o *AuthGetNumbers) HasInternational() bool`

HasInternational returns a boolean if a field has been set.

### GetBacs

`func (o *AuthGetNumbers) GetBacs() []NumbersBACS`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *AuthGetNumbers) GetBacsOk() (*[]NumbersBACS, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *AuthGetNumbers) SetBacs(v []NumbersBACS)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *AuthGetNumbers) HasBacs() bool`

HasBacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


