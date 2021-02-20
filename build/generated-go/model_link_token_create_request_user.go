/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.5.3
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// LinkTokenCreateRequestUser An object specifying information about the end user who will be linking their account.
type LinkTokenCreateRequestUser struct {
	// A unique ID representing the end user. Typically this will be a user ID number from your application. Personally identifiable information, such as an email address or phone number, should not be used in the `client_user_id`.
	ClientUserId string `json:"client_user_id"`
	// The user's full legal name. This is an optional field used in the [returning user experience](/docs/link/returning-user) to associate Items to the user.
	LegalName *string `json:"legal_name,omitempty"`
	// The user's phone number in [E.164](https://en.wikipedia.org/wiki/E.164) format. This field is optional, but required to enable the [returning user experience](/docs/link/returning-user).
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The date and time the phone number was verified in ISO 8601 format (`YYYY-MM-DDThh:mm:ssZ`). This field is optional, but required to enable any [returning user experience](/docs/link/returning-user).   Only pass a verification time for a phone number that you have verified. If you have performed verification but don’t have the time, you may supply a signal value of the start of the UNIX epoch.   Example: `2020-01-01T00:00:00Z` 
	PhoneNumberVerifiedTime *string `json:"phone_number_verified_time,omitempty"`
	// The user's email address. This field is optional, but required to enable the [pre-authenticated returning user flow](/docs/link/returning-user/#enabling-the-returning-user-experience).
	EmailAddress *string `json:"email_address,omitempty"`
	// The date and time the email address was verified in ISO 8601 format (`YYYY-MM-DDThh:mm:ssZ`). This is an optional field used in the [returning user experience](/docs/link/returning-user).   Only pass a verification time for an email address that you have verified. If you have performed verification but don’t have the time, you may supply a signal value of the start of the UNIX epoch.   Example: `2020-01-01T00:00:00Z`
	EmailAddressVerifiedTime *string `json:"email_address_verified_time,omitempty"`
	// To be provided in the format \"ddd-dd-dddd\". This field is optional and will support not-yet-implemented functionality for new products.
	Ssn *string `json:"ssn,omitempty"`
	// To be provided in the format \"yyyy-mm-dd\". This field is optional and will support not-yet-implemented functionality for new products.
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateRequestUser LinkTokenCreateRequestUser

// NewLinkTokenCreateRequestUser instantiates a new LinkTokenCreateRequestUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestUser(clientUserId string, ) *LinkTokenCreateRequestUser {
	this := LinkTokenCreateRequestUser{}
	this.ClientUserId = clientUserId
	return &this
}

// NewLinkTokenCreateRequestUserWithDefaults instantiates a new LinkTokenCreateRequestUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestUserWithDefaults() *LinkTokenCreateRequestUser {
	this := LinkTokenCreateRequestUser{}
	return &this
}

// GetClientUserId returns the ClientUserId field value
func (o *LinkTokenCreateRequestUser) GetClientUserId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetClientUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientUserId, true
}

// SetClientUserId sets field value
func (o *LinkTokenCreateRequestUser) SetClientUserId(v string) {
	o.ClientUserId = v
}

// GetLegalName returns the LegalName field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetLegalName() string {
	if o == nil || o.LegalName == nil {
		var ret string
		return ret
	}
	return *o.LegalName
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetLegalNameOk() (*string, bool) {
	if o == nil || o.LegalName == nil {
		return nil, false
	}
	return o.LegalName, true
}

// HasLegalName returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasLegalName() bool {
	if o != nil && o.LegalName != nil {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given string and assigns it to the LegalName field.
func (o *LinkTokenCreateRequestUser) SetLegalName(v string) {
	o.LegalName = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *LinkTokenCreateRequestUser) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetPhoneNumberVerifiedTime returns the PhoneNumberVerifiedTime field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetPhoneNumberVerifiedTime() string {
	if o == nil || o.PhoneNumberVerifiedTime == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumberVerifiedTime
}

// GetPhoneNumberVerifiedTimeOk returns a tuple with the PhoneNumberVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetPhoneNumberVerifiedTimeOk() (*string, bool) {
	if o == nil || o.PhoneNumberVerifiedTime == nil {
		return nil, false
	}
	return o.PhoneNumberVerifiedTime, true
}

// HasPhoneNumberVerifiedTime returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasPhoneNumberVerifiedTime() bool {
	if o != nil && o.PhoneNumberVerifiedTime != nil {
		return true
	}

	return false
}

// SetPhoneNumberVerifiedTime gets a reference to the given string and assigns it to the PhoneNumberVerifiedTime field.
func (o *LinkTokenCreateRequestUser) SetPhoneNumberVerifiedTime(v string) {
	o.PhoneNumberVerifiedTime = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *LinkTokenCreateRequestUser) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetEmailAddressVerifiedTime returns the EmailAddressVerifiedTime field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetEmailAddressVerifiedTime() string {
	if o == nil || o.EmailAddressVerifiedTime == nil {
		var ret string
		return ret
	}
	return *o.EmailAddressVerifiedTime
}

// GetEmailAddressVerifiedTimeOk returns a tuple with the EmailAddressVerifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetEmailAddressVerifiedTimeOk() (*string, bool) {
	if o == nil || o.EmailAddressVerifiedTime == nil {
		return nil, false
	}
	return o.EmailAddressVerifiedTime, true
}

// HasEmailAddressVerifiedTime returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasEmailAddressVerifiedTime() bool {
	if o != nil && o.EmailAddressVerifiedTime != nil {
		return true
	}

	return false
}

// SetEmailAddressVerifiedTime gets a reference to the given string and assigns it to the EmailAddressVerifiedTime field.
func (o *LinkTokenCreateRequestUser) SetEmailAddressVerifiedTime(v string) {
	o.EmailAddressVerifiedTime = &v
}

// GetSsn returns the Ssn field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetSsn() string {
	if o == nil || o.Ssn == nil {
		var ret string
		return ret
	}
	return *o.Ssn
}

// GetSsnOk returns a tuple with the Ssn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetSsnOk() (*string, bool) {
	if o == nil || o.Ssn == nil {
		return nil, false
	}
	return o.Ssn, true
}

// HasSsn returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasSsn() bool {
	if o != nil && o.Ssn != nil {
		return true
	}

	return false
}

// SetSsn gets a reference to the given string and assigns it to the Ssn field.
func (o *LinkTokenCreateRequestUser) SetSsn(v string) {
	o.Ssn = &v
}

// GetDateOfBirth returns the DateOfBirth field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestUser) GetDateOfBirth() string {
	if o == nil || o.DateOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DateOfBirth
}

// GetDateOfBirthOk returns a tuple with the DateOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestUser) GetDateOfBirthOk() (*string, bool) {
	if o == nil || o.DateOfBirth == nil {
		return nil, false
	}
	return o.DateOfBirth, true
}

// HasDateOfBirth returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestUser) HasDateOfBirth() bool {
	if o != nil && o.DateOfBirth != nil {
		return true
	}

	return false
}

// SetDateOfBirth gets a reference to the given string and assigns it to the DateOfBirth field.
func (o *LinkTokenCreateRequestUser) SetDateOfBirth(v string) {
	o.DateOfBirth = &v
}

func (o LinkTokenCreateRequestUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_user_id"] = o.ClientUserId
	}
	if o.LegalName != nil {
		toSerialize["legal_name"] = o.LegalName
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if o.PhoneNumberVerifiedTime != nil {
		toSerialize["phone_number_verified_time"] = o.PhoneNumberVerifiedTime
	}
	if o.EmailAddress != nil {
		toSerialize["email_address"] = o.EmailAddress
	}
	if o.EmailAddressVerifiedTime != nil {
		toSerialize["email_address_verified_time"] = o.EmailAddressVerifiedTime
	}
	if o.Ssn != nil {
		toSerialize["ssn"] = o.Ssn
	}
	if o.DateOfBirth != nil {
		toSerialize["date_of_birth"] = o.DateOfBirth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateRequestUser) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateRequestUser := _LinkTokenCreateRequestUser{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateRequestUser); err == nil {
		*o = LinkTokenCreateRequestUser(varLinkTokenCreateRequestUser)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_user_id")
		delete(additionalProperties, "legal_name")
		delete(additionalProperties, "phone_number")
		delete(additionalProperties, "phone_number_verified_time")
		delete(additionalProperties, "email_address")
		delete(additionalProperties, "email_address_verified_time")
		delete(additionalProperties, "ssn")
		delete(additionalProperties, "date_of_birth")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateRequestUser struct {
	value *LinkTokenCreateRequestUser
	isSet bool
}

func (v NullableLinkTokenCreateRequestUser) Get() *LinkTokenCreateRequestUser {
	return v.value
}

func (v *NullableLinkTokenCreateRequestUser) Set(val *LinkTokenCreateRequestUser) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestUser) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestUser(val *LinkTokenCreateRequestUser) *NullableLinkTokenCreateRequestUser {
	return &NullableLinkTokenCreateRequestUser{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


