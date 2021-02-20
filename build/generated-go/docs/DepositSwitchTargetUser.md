# DepositSwitchTargetUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** | The given name (first name) of the user. | 
**FamilyName** | **string** | The family name (last name) of the user. | 
**Phone** | **string** | The phone number of the user. The endpoint can accept a variety of phone number formats, including E.164. | 
**Email** | **string** | The email address of the user. | 
**Address** | Pointer to [**DepositSwitchAddressData**](DepositSwitchAddressData.md) |  | [optional] 
**TaxPayerId** | Pointer to **string** | The taxpayer ID of the user, generally their SSN, EIN, or TIN. | [optional] 

## Methods

### NewDepositSwitchTargetUser

`func NewDepositSwitchTargetUser(givenName string, familyName string, phone string, email string, ) *DepositSwitchTargetUser`

NewDepositSwitchTargetUser instantiates a new DepositSwitchTargetUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchTargetUserWithDefaults

`func NewDepositSwitchTargetUserWithDefaults() *DepositSwitchTargetUser`

NewDepositSwitchTargetUserWithDefaults instantiates a new DepositSwitchTargetUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *DepositSwitchTargetUser) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *DepositSwitchTargetUser) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *DepositSwitchTargetUser) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### GetFamilyName

`func (o *DepositSwitchTargetUser) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *DepositSwitchTargetUser) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *DepositSwitchTargetUser) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### GetPhone

`func (o *DepositSwitchTargetUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *DepositSwitchTargetUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *DepositSwitchTargetUser) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetEmail

`func (o *DepositSwitchTargetUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DepositSwitchTargetUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DepositSwitchTargetUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetAddress

`func (o *DepositSwitchTargetUser) GetAddress() DepositSwitchAddressData`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DepositSwitchTargetUser) GetAddressOk() (*DepositSwitchAddressData, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DepositSwitchTargetUser) SetAddress(v DepositSwitchAddressData)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DepositSwitchTargetUser) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetTaxPayerId

`func (o *DepositSwitchTargetUser) GetTaxPayerId() string`

GetTaxPayerId returns the TaxPayerId field if non-nil, zero value otherwise.

### GetTaxPayerIdOk

`func (o *DepositSwitchTargetUser) GetTaxPayerIdOk() (*string, bool)`

GetTaxPayerIdOk returns a tuple with the TaxPayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxPayerId

`func (o *DepositSwitchTargetUser) SetTaxPayerId(v string)`

SetTaxPayerId sets TaxPayerId field to given value.

### HasTaxPayerId

`func (o *DepositSwitchTargetUser) HasTaxPayerId() bool`

HasTaxPayerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


