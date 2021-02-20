# Owner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | A list of names associated with the account by the financial institution. These should always be the names of individuals, even for business accounts. If the name of a business is reported, please contact Plaid Support.  In the case of a joint account, the names of all account holders will be reported.  If an Item contains multiple accounts with different owner names, some institutions will report all names associated with the Item in each account&#39;s &#x60;names&#x60; array. | 
**PhoneNumbers** | [**[]PhoneNumber**](PhoneNumber.md) | A list of phone numbers associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution. | 
**Emails** | [**[]Email**](Email.md) | A list of email addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution. | 
**Addresses** | [**[]Address**](Address.md) | Data about the various addresses associated with the account by the financial institution. May be an empty array if no relevant information is returned from the financial institution. | 

## Methods

### NewOwner

`func NewOwner(names []string, phoneNumbers []PhoneNumber, emails []Email, addresses []Address, ) *Owner`

NewOwner instantiates a new Owner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnerWithDefaults

`func NewOwnerWithDefaults() *Owner`

NewOwnerWithDefaults instantiates a new Owner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *Owner) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *Owner) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *Owner) SetNames(v []string)`

SetNames sets Names field to given value.


### GetPhoneNumbers

`func (o *Owner) GetPhoneNumbers() []PhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *Owner) GetPhoneNumbersOk() (*[]PhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *Owner) SetPhoneNumbers(v []PhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetEmails

`func (o *Owner) GetEmails() []Email`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Owner) GetEmailsOk() (*[]Email, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Owner) SetEmails(v []Email)`

SetEmails sets Emails field to given value.


### GetAddresses

`func (o *Owner) GetAddresses() []Address`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *Owner) GetAddressesOk() (*[]Address, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *Owner) SetAddresses(v []Address)`

SetAddresses sets Addresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


