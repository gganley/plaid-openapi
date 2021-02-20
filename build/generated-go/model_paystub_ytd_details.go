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

// PaystubYTDDetails struct for PaystubYTDDetails
type PaystubYTDDetails struct {
	// Year-to-date gross earnings.
	GrossEarnings float32 `json:"gross_earnings"`
	// Year-to-date net (take home) earnings.
	NetEarnings float32 `json:"net_earnings"`
	AdditionalProperties map[string]interface{}
}

type _PaystubYTDDetails PaystubYTDDetails

// NewPaystubYTDDetails instantiates a new PaystubYTDDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubYTDDetails(grossEarnings float32, netEarnings float32, ) *PaystubYTDDetails {
	this := PaystubYTDDetails{}
	this.GrossEarnings = grossEarnings
	this.NetEarnings = netEarnings
	return &this
}

// NewPaystubYTDDetailsWithDefaults instantiates a new PaystubYTDDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubYTDDetailsWithDefaults() *PaystubYTDDetails {
	this := PaystubYTDDetails{}
	return &this
}

// GetGrossEarnings returns the GrossEarnings field value
func (o *PaystubYTDDetails) GetGrossEarnings() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.GrossEarnings
}

// GetGrossEarningsOk returns a tuple with the GrossEarnings field value
// and a boolean to check if the value has been set.
func (o *PaystubYTDDetails) GetGrossEarningsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.GrossEarnings, true
}

// SetGrossEarnings sets field value
func (o *PaystubYTDDetails) SetGrossEarnings(v float32) {
	o.GrossEarnings = v
}

// GetNetEarnings returns the NetEarnings field value
func (o *PaystubYTDDetails) GetNetEarnings() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.NetEarnings
}

// GetNetEarningsOk returns a tuple with the NetEarnings field value
// and a boolean to check if the value has been set.
func (o *PaystubYTDDetails) GetNetEarningsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetEarnings, true
}

// SetNetEarnings sets field value
func (o *PaystubYTDDetails) SetNetEarnings(v float32) {
	o.NetEarnings = v
}

func (o PaystubYTDDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["gross_earnings"] = o.GrossEarnings
	}
	if true {
		toSerialize["net_earnings"] = o.NetEarnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubYTDDetails) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubYTDDetails := _PaystubYTDDetails{}

	if err = json.Unmarshal(bytes, &varPaystubYTDDetails); err == nil {
		*o = PaystubYTDDetails(varPaystubYTDDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "gross_earnings")
		delete(additionalProperties, "net_earnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubYTDDetails struct {
	value *PaystubYTDDetails
	isSet bool
}

func (v NullablePaystubYTDDetails) Get() *PaystubYTDDetails {
	return v.value
}

func (v *NullablePaystubYTDDetails) Set(val *PaystubYTDDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubYTDDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubYTDDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubYTDDetails(val *PaystubYTDDetails) *NullablePaystubYTDDetails {
	return &NullablePaystubYTDDetails{value: val, isSet: true}
}

func (v NullablePaystubYTDDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubYTDDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


