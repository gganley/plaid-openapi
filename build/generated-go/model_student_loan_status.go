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

// StudentLoanStatus An object representing the status of the student loan
type StudentLoanStatus struct {
	// The date until which the loan will be in its current status. Dates are returned in an ISO 8601 format (YYYY-MM-DD). 
	EndDate NullableString `json:"end_date,omitempty"`
	// The status type of the student loan
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StudentLoanStatus StudentLoanStatus

// NewStudentLoanStatus instantiates a new StudentLoanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudentLoanStatus() *StudentLoanStatus {
	this := StudentLoanStatus{}
	return &this
}

// NewStudentLoanStatusWithDefaults instantiates a new StudentLoanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudentLoanStatusWithDefaults() *StudentLoanStatus {
	this := StudentLoanStatus{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StudentLoanStatus) GetEndDate() string {
	if o == nil || o.EndDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.EndDate.Get()
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoanStatus) GetEndDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDate.Get(), o.EndDate.IsSet()
}

// HasEndDate returns a boolean if a field has been set.
func (o *StudentLoanStatus) HasEndDate() bool {
	if o != nil && o.EndDate.IsSet() {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given NullableString and assigns it to the EndDate field.
func (o *StudentLoanStatus) SetEndDate(v string) {
	o.EndDate.Set(&v)
}
// SetEndDateNil sets the value for EndDate to be an explicit nil
func (o *StudentLoanStatus) SetEndDateNil() {
	o.EndDate.Set(nil)
}

// UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
func (o *StudentLoanStatus) UnsetEndDate() {
	o.EndDate.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StudentLoanStatus) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StudentLoanStatus) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *StudentLoanStatus) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *StudentLoanStatus) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *StudentLoanStatus) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *StudentLoanStatus) UnsetType() {
	o.Type.Unset()
}

func (o StudentLoanStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EndDate.IsSet() {
		toSerialize["end_date"] = o.EndDate.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StudentLoanStatus) UnmarshalJSON(bytes []byte) (err error) {
	varStudentLoanStatus := _StudentLoanStatus{}

	if err = json.Unmarshal(bytes, &varStudentLoanStatus); err == nil {
		*o = StudentLoanStatus(varStudentLoanStatus)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "end_date")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStudentLoanStatus struct {
	value *StudentLoanStatus
	isSet bool
}

func (v NullableStudentLoanStatus) Get() *StudentLoanStatus {
	return v.value
}

func (v *NullableStudentLoanStatus) Set(val *StudentLoanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStudentLoanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStudentLoanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudentLoanStatus(val *StudentLoanStatus) *NullableStudentLoanStatus {
	return &NullableStudentLoanStatus{value: val, isSet: true}
}

func (v NullableStudentLoanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudentLoanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


