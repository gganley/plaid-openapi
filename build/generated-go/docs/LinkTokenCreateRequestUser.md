# LinkTokenCreateRequestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUserId** | **string** | A unique ID representing the end user. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the &#x60;client_user_id&#x60;. | 
**LegalName** | Pointer to **string** | The user&#39;s full legal name. This is an optional field used in the [returning user experience](/docs/link/returning-user) to associate Items to the user. | [optional] 
**PhoneNumber** | Pointer to **string** | The user&#39;s phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. This field is optional, but required to enable the [returning user experience](/docs/link/returning-user). | [optional] 
**PhoneNumberVerifiedTime** | Pointer to **string** | The date and time the phone number was verified in ISO 8601 format (&#x60;YYYY-MM-DDThh:mm:ssZ&#x60;). This field is optional, but required to enable any [returning user experience](/docs/link/returning-user).   Only pass a verification time for a phone number that you have verified. If you have performed verification but don’t have the time, you may supply a signal value of the start of the UNIX epoch.   Example: &#x60;2020-01-01T00:00:00Z&#x60;  | [optional] 
**EmailAddress** | Pointer to **string** | The user&#39;s email address. This field is optional, but required to enable the [pre-authenticated returning user flow](/docs/link/returning-user/#enabling-the-returning-user-experience). | [optional] 
**EmailAddressVerifiedTime** | Pointer to **string** | The date and time the email address was verified in ISO 8601 format (&#x60;YYYY-MM-DDThh:mm:ssZ&#x60;). This is an optional field used in the [returning user experience](/docs/link/returning-user).   Only pass a verification time for an email address that you have verified. If you have performed verification but don’t have the time, you may supply a signal value of the start of the UNIX epoch.   Example: &#x60;2020-01-01T00:00:00Z&#x60; | [optional] 
**Ssn** | Pointer to **string** | To be provided in the format \&quot;ddd-dd-dddd\&quot;. This field is optional and will support not-yet-implemented functionality for new products. | [optional] 
**DateOfBirth** | Pointer to **string** | To be provided in the format \&quot;yyyy-mm-dd\&quot;. This field is optional and will support not-yet-implemented functionality for new products. | [optional] 

## Methods

### NewLinkTokenCreateRequestUser

`func NewLinkTokenCreateRequestUser(clientUserId string, ) *LinkTokenCreateRequestUser`

NewLinkTokenCreateRequestUser instantiates a new LinkTokenCreateRequestUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestUserWithDefaults

`func NewLinkTokenCreateRequestUserWithDefaults() *LinkTokenCreateRequestUser`

NewLinkTokenCreateRequestUserWithDefaults instantiates a new LinkTokenCreateRequestUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientUserId

`func (o *LinkTokenCreateRequestUser) GetClientUserId() string`

GetClientUserId returns the ClientUserId field if non-nil, zero value otherwise.

### GetClientUserIdOk

`func (o *LinkTokenCreateRequestUser) GetClientUserIdOk() (*string, bool)`

GetClientUserIdOk returns a tuple with the ClientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserId

`func (o *LinkTokenCreateRequestUser) SetClientUserId(v string)`

SetClientUserId sets ClientUserId field to given value.


### GetLegalName

`func (o *LinkTokenCreateRequestUser) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *LinkTokenCreateRequestUser) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *LinkTokenCreateRequestUser) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *LinkTokenCreateRequestUser) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *LinkTokenCreateRequestUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *LinkTokenCreateRequestUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *LinkTokenCreateRequestUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *LinkTokenCreateRequestUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPhoneNumberVerifiedTime

`func (o *LinkTokenCreateRequestUser) GetPhoneNumberVerifiedTime() string`

GetPhoneNumberVerifiedTime returns the PhoneNumberVerifiedTime field if non-nil, zero value otherwise.

### GetPhoneNumberVerifiedTimeOk

`func (o *LinkTokenCreateRequestUser) GetPhoneNumberVerifiedTimeOk() (*string, bool)`

GetPhoneNumberVerifiedTimeOk returns a tuple with the PhoneNumberVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberVerifiedTime

`func (o *LinkTokenCreateRequestUser) SetPhoneNumberVerifiedTime(v string)`

SetPhoneNumberVerifiedTime sets PhoneNumberVerifiedTime field to given value.

### HasPhoneNumberVerifiedTime

`func (o *LinkTokenCreateRequestUser) HasPhoneNumberVerifiedTime() bool`

HasPhoneNumberVerifiedTime returns a boolean if a field has been set.

### GetEmailAddress

`func (o *LinkTokenCreateRequestUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *LinkTokenCreateRequestUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *LinkTokenCreateRequestUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *LinkTokenCreateRequestUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAddressVerifiedTime

`func (o *LinkTokenCreateRequestUser) GetEmailAddressVerifiedTime() string`

GetEmailAddressVerifiedTime returns the EmailAddressVerifiedTime field if non-nil, zero value otherwise.

### GetEmailAddressVerifiedTimeOk

`func (o *LinkTokenCreateRequestUser) GetEmailAddressVerifiedTimeOk() (*string, bool)`

GetEmailAddressVerifiedTimeOk returns a tuple with the EmailAddressVerifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressVerifiedTime

`func (o *LinkTokenCreateRequestUser) SetEmailAddressVerifiedTime(v string)`

SetEmailAddressVerifiedTime sets EmailAddressVerifiedTime field to given value.

### HasEmailAddressVerifiedTime

`func (o *LinkTokenCreateRequestUser) HasEmailAddressVerifiedTime() bool`

HasEmailAddressVerifiedTime returns a boolean if a field has been set.

### GetSsn

`func (o *LinkTokenCreateRequestUser) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *LinkTokenCreateRequestUser) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *LinkTokenCreateRequestUser) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *LinkTokenCreateRequestUser) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *LinkTokenCreateRequestUser) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *LinkTokenCreateRequestUser) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *LinkTokenCreateRequestUser) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *LinkTokenCreateRequestUser) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


