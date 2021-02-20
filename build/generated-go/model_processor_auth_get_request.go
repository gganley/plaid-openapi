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

// ProcessorAuthGetRequest ProcessorAuthGetRequest defines the request schema for `/processor/auth/get`
type ProcessorAuthGetRequest struct {
	// Your Plaid API `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`.
	Secret *string `json:"secret,omitempty"`
	// The processor token obtained from the Plaid integration partner. Processor tokens are in the format: `processor-<environment>-<identifier>`
	ProcessorToken string `json:"processor_token"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorAuthGetRequest ProcessorAuthGetRequest

// NewProcessorAuthGetRequest instantiates a new ProcessorAuthGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorAuthGetRequest(processorToken string, ) *ProcessorAuthGetRequest {
	this := ProcessorAuthGetRequest{}
	this.ProcessorToken = processorToken
	return &this
}

// NewProcessorAuthGetRequestWithDefaults instantiates a new ProcessorAuthGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorAuthGetRequestWithDefaults() *ProcessorAuthGetRequest {
	this := ProcessorAuthGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ProcessorAuthGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ProcessorAuthGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ProcessorAuthGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ProcessorAuthGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ProcessorAuthGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ProcessorAuthGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorAuthGetRequest) GetProcessorToken() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetRequest) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorAuthGetRequest) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

func (o ProcessorAuthGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorAuthGetRequest) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorAuthGetRequest := _ProcessorAuthGetRequest{}

	if err = json.Unmarshal(bytes, &varProcessorAuthGetRequest); err == nil {
		*o = ProcessorAuthGetRequest(varProcessorAuthGetRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "processor_token")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorAuthGetRequest struct {
	value *ProcessorAuthGetRequest
	isSet bool
}

func (v NullableProcessorAuthGetRequest) Get() *ProcessorAuthGetRequest {
	return v.value
}

func (v *NullableProcessorAuthGetRequest) Set(val *ProcessorAuthGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorAuthGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorAuthGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorAuthGetRequest(val *ProcessorAuthGetRequest) *NullableProcessorAuthGetRequest {
	return &NullableProcessorAuthGetRequest{value: val, isSet: true}
}

func (v NullableProcessorAuthGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorAuthGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


