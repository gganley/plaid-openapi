# OwnerOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | A list of names associated with the account by the financial institution. These should always be the names of individuals, even for business accounts. Note that the same name data will be used for all accounts associated with an Item. | 
**PhoneNumbers** | [**[]PhoneNumber**](PhoneNumber.md) | A list of phone numbers associated with the account. | 
**Emails** | [**[]Email**](Email.md) | A list of email addresses associated with the account. | 
**Addresses** | [**[]Address**](Address.md) | Data about the various addresses associated with the account. | 

## Methods

### NewOwnerOverride

`func NewOwnerOverride(names []string, phoneNumbers []PhoneNumber, emails []Email, addresses []Address, ) *OwnerOverride`

NewOwnerOverride instantiates a new OwnerOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerOverrideWithDefaults

`func NewOwnerOverrideWithDefaults() *OwnerOverride`

NewOwnerOverrideWithDefaults instantiates a new OwnerOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *OwnerOverride) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *OwnerOverride) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *OwnerOverride) SetNames(v []string)`

SetNames sets Names field to given value.


### GetPhoneNumbers

`func (o *OwnerOverride) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *OwnerOverride) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *OwnerOverride) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetEmails

`func (o *OwnerOverride) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *OwnerOverride) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *OwnerOverride) SetEmails(v []Email)`

SetEmails sets Emails field to given value.


### GetAddresses

`func (o *OwnerOverride) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *OwnerOverride) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *OwnerOverride) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


