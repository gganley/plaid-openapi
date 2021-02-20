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

// PaymentInitiationAddress The optional address of the payment recipient. This object is not currently required to make payments from UK institutions and should not be populated, though may be necessary for future European expansion.
type PaymentInitiationAddress struct {
	// An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters.
	Street *[]string `json:"street,omitempty"`
	// The city where the recipient is located. Maximum of 35 characters.
	City *string `json:"city,omitempty"`
	// The postal code where the recipient is located. Maximum of 16 characters.
	PostalCode *string `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code where the recipient is located.
	Country              *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationAddress PaymentInitiationAddress

// NewPaymentInitiationAddress instantiates a new PaymentInitiationAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationAddress() *PaymentInitiationAddress {
	this := PaymentInitiationAddress{}
	return &this
}

// NewPaymentInitiationAddressWithDefaults instantiates a new PaymentInitiationAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationAddressWithDefaults() *PaymentInitiationAddress {
	this := PaymentInitiationAddress{}
	return &this
}

// GetStreet returns the Street field value if set, zero value otherwise.
func (o *PaymentInitiationAddress) GetStreet() []string {
	if o == nil || o.Street == nil {
		var ret []string
		return ret
	}
	return *o.Street
}

// GetStreetOk returns a tuple with the Street field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetStreetOk() (*[]string, bool) {
	if o == nil || o.Street == nil {
		return nil, false
	}
	return o.Street, true
}

// HasStreet returns a boolean if a field has been set.
func (o *PaymentInitiationAddress) HasStreet() bool {
	if o != nil && o.Street != nil {
		return true
	}

	return false
}

// SetStreet gets a reference to the given []string and assigns it to the Street field.
func (o *PaymentInitiationAddress) SetStreet(v []string) {
	o.Street = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *PaymentInitiationAddress) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *PaymentInitiationAddress) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *PaymentInitiationAddress) SetCity(v string) {
	o.City = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *PaymentInitiationAddress) GetPostalCode() string {
	if o == nil || o.PostalCode == nil {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetPostalCodeOk() (*string, bool) {
	if o == nil || o.PostalCode == nil {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *PaymentInitiationAddress) HasPostalCode() bool {
	if o != nil && o.PostalCode != nil {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *PaymentInitiationAddress) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *PaymentInitiationAddress) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentInitiationAddress) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentInitiationAddress) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *PaymentInitiationAddress) SetCountry(v string) {
	o.Country = &v
}

func (o PaymentInitiationAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Street != nil {
		toSerialize["street"] = o.Street
	}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.PostalCode != nil {
		toSerialize["postal_code"] = o.PostalCode
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationAddress) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationAddress := _PaymentInitiationAddress{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationAddress); err == nil {
		*o = PaymentInitiationAddress(varPaymentInitiationAddress)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "street")
		delete(additionalProperties, "city")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationAddress struct {
	value *PaymentInitiationAddress
	isSet bool
}

func (v NullablePaymentInitiationAddress) Get() *PaymentInitiationAddress {
	return v.value
}

func (v *NullablePaymentInitiationAddress) Set(val *PaymentInitiationAddress) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationAddress) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationAddress(val *PaymentInitiationAddress) *NullablePaymentInitiationAddress {
	return &NullablePaymentInitiationAddress{value: val, isSet: true}
}

func (v NullablePaymentInitiationAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
