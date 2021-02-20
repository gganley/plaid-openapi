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

// AddressData Data about the components comprising an address.
type AddressData struct {
	// The full city name
	City string `json:"city"`
	// The region or state Example: `\"NC\"`
	Region NullableString `json:"region,omitempty"`
	// The full street address Example: `\"564 Main Street, APT 15\"`
	Street string `json:"street"`
	// The postal code
	PostalCode NullableString `json:"postal_code,omitempty"`
	// The ISO 3166-1 alpha-2 country code
	Country              string `json:"country"`
	AdditionalProperties map[string]interface{}
}

type _AddressData AddressData

// NewAddressData instantiates a new AddressData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressData(city string, street string, country string) *AddressData {
	this := AddressData{}
	this.City = city
	this.Street = street
	this.Country = country
	return &this
}

// NewAddressDataWithDefaults instantiates a new AddressData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressDataWithDefaults() *AddressData {
	this := AddressData{}
	return &this
}

// GetCity returns the City field value
func (o *AddressData) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *AddressData) SetCity(v string) {
	o.City = v
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressData) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressData) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *AddressData) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *AddressData) SetRegion(v string) {
	o.Region.Set(&v)
}

// SetRegionNil sets the value for Region to be an explicit nil
func (o *AddressData) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *AddressData) UnsetRegion() {
	o.Region.Unset()
}

// GetStreet returns the Street field value
func (o *AddressData) GetStreet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Street
}

// GetStreetOk returns a tuple with the Street field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetStreetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Street, true
}

// SetStreet sets field value
func (o *AddressData) SetStreet(v string) {
	o.Street = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddressData) GetPostalCode() string {
	if o == nil || o.PostalCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.PostalCode.Get()
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddressData) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostalCode.Get(), o.PostalCode.IsSet()
}

// HasPostalCode returns a boolean if a field has been set.
func (o *AddressData) HasPostalCode() bool {
	if o != nil && o.PostalCode.IsSet() {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given NullableString and assigns it to the PostalCode field.
func (o *AddressData) SetPostalCode(v string) {
	o.PostalCode.Set(&v)
}

// SetPostalCodeNil sets the value for PostalCode to be an explicit nil
func (o *AddressData) SetPostalCodeNil() {
	o.PostalCode.Set(nil)
}

// UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
func (o *AddressData) UnsetPostalCode() {
	o.PostalCode.Unset()
}

// GetCountry returns the Country field value
func (o *AddressData) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *AddressData) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *AddressData) SetCountry(v string) {
	o.Country = v
}

func (o AddressData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["city"] = o.City
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	if true {
		toSerialize["street"] = o.Street
	}
	if o.PostalCode.IsSet() {
		toSerialize["postal_code"] = o.PostalCode.Get()
	}
	if true {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AddressData) UnmarshalJSON(bytes []byte) (err error) {
	varAddressData := _AddressData{}

	if err = json.Unmarshal(bytes, &varAddressData); err == nil {
		*o = AddressData(varAddressData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "region")
		delete(additionalProperties, "street")
		delete(additionalProperties, "postal_code")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAddressData struct {
	value *AddressData
	isSet bool
}

func (v NullableAddressData) Get() *AddressData {
	return v.value
}

func (v *NullableAddressData) Set(val *AddressData) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressData) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressData(val *AddressData) *NullableAddressData {
	return &NullableAddressData{value: val, isSet: true}
}

func (v NullableAddressData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
