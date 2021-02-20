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

// DepositSwitchGetRequest DepositSwitchGetRequest defines the request schema for `/deposit_switch/get`
type DepositSwitchGetRequest struct {
	// Your Plaid API `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`.
	Secret *string `json:"secret,omitempty"`
	// The ID of the deposit switch
	DepositSwitchId string `json:"deposit_switch_id"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchGetRequest DepositSwitchGetRequest

// NewDepositSwitchGetRequest instantiates a new DepositSwitchGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchGetRequest(depositSwitchId string, ) *DepositSwitchGetRequest {
	this := DepositSwitchGetRequest{}
	this.DepositSwitchId = depositSwitchId
	return &this
}

// NewDepositSwitchGetRequestWithDefaults instantiates a new DepositSwitchGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchGetRequestWithDefaults() *DepositSwitchGetRequest {
	this := DepositSwitchGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DepositSwitchGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DepositSwitchGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DepositSwitchGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *DepositSwitchGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *DepositSwitchGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *DepositSwitchGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *DepositSwitchGetRequest) GetDepositSwitchId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetRequest) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *DepositSwitchGetRequest) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

func (o DepositSwitchGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchGetRequest := _DepositSwitchGetRequest{}

	if err = json.Unmarshal(bytes, &varDepositSwitchGetRequest); err == nil {
		*o = DepositSwitchGetRequest(varDepositSwitchGetRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "deposit_switch_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchGetRequest struct {
	value *DepositSwitchGetRequest
	isSet bool
}

func (v NullableDepositSwitchGetRequest) Get() *DepositSwitchGetRequest {
	return v.value
}

func (v *NullableDepositSwitchGetRequest) Set(val *DepositSwitchGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchGetRequest(val *DepositSwitchGetRequest) *NullableDepositSwitchGetRequest {
	return &NullableDepositSwitchGetRequest{value: val, isSet: true}
}

func (v NullableDepositSwitchGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

