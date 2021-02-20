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

// ProcessorTokenCreateResponse ProcessorTokenCreateResponse defines the response schema for `/processor/token/create` and `/processor/apex/processor_token/create`
type ProcessorTokenCreateResponse struct {
	// The `processor_token` that can then be used by the Plaid partner to make API requests
	ProcessorToken string `json:"processor_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId            string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorTokenCreateResponse ProcessorTokenCreateResponse

// NewProcessorTokenCreateResponse instantiates a new ProcessorTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorTokenCreateResponse(processorToken string, requestId string) *ProcessorTokenCreateResponse {
	this := ProcessorTokenCreateResponse{}
	this.ProcessorToken = processorToken
	this.RequestId = requestId
	return &this
}

// NewProcessorTokenCreateResponseWithDefaults instantiates a new ProcessorTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorTokenCreateResponseWithDefaults() *ProcessorTokenCreateResponse {
	this := ProcessorTokenCreateResponse{}
	return &this
}

// GetProcessorToken returns the ProcessorToken field value
func (o *ProcessorTokenCreateResponse) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateResponse) GetProcessorTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *ProcessorTokenCreateResponse) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetRequestId returns the RequestId field value
func (o *ProcessorTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ProcessorTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorTokenCreateResponse := _ProcessorTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varProcessorTokenCreateResponse); err == nil {
		*o = ProcessorTokenCreateResponse(varProcessorTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "processor_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorTokenCreateResponse struct {
	value *ProcessorTokenCreateResponse
	isSet bool
}

func (v NullableProcessorTokenCreateResponse) Get() *ProcessorTokenCreateResponse {
	return v.value
}

func (v *NullableProcessorTokenCreateResponse) Set(val *ProcessorTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorTokenCreateResponse(val *ProcessorTokenCreateResponse) *NullableProcessorTokenCreateResponse {
	return &NullableProcessorTokenCreateResponse{value: val, isSet: true}
}

func (v NullableProcessorTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
