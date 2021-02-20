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

// NullableAccessToken struct for NullableAccessToken
type NullableAccessToken struct {
	AdditionalProperties map[string]interface{}
}

type _NullableAccessToken NullableAccessToken

// NewNullableAccessToken instantiates a new NullableAccessToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNullableAccessToken() *NullableAccessToken {
	this := NullableAccessToken{}
	return &this
}

// NewNullableAccessTokenWithDefaults instantiates a new NullableAccessToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNullableAccessTokenWithDefaults() *NullableAccessToken {
	this := NullableAccessToken{}
	return &this
}

func (o NullableAccessToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *NullableAccessToken) UnmarshalJSON(bytes []byte) (err error) {
	varNullableAccessToken := _NullableAccessToken{}

	if err = json.Unmarshal(bytes, &varNullableAccessToken); err == nil {
		*o = NullableAccessToken(varNullableAccessToken)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNullableAccessToken struct {
	value *NullableAccessToken
	isSet bool
}

func (v NullableNullableAccessToken) Get() *NullableAccessToken {
	return v.value
}

func (v *NullableNullableAccessToken) Set(val *NullableAccessToken) {
	v.value = val
	v.isSet = true
}

func (v NullableNullableAccessToken) IsSet() bool {
	return v.isSet
}

func (v *NullableNullableAccessToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNullableAccessToken(val *NullableAccessToken) *NullableNullableAccessToken {
	return &NullableNullableAccessToken{value: val, isSet: true}
}

func (v NullableNullableAccessToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNullableAccessToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


