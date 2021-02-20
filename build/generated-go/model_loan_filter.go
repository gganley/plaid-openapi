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

// LoanFilter A filter to apply to `loan`-type accounts
type LoanFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](/docs/api/accounts#accounts-schema). 
	AccountSubtypes []AccountSubtype `json:"account_subtypes"`
	AdditionalProperties map[string]interface{}
}

type _LoanFilter LoanFilter

// NewLoanFilter instantiates a new LoanFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoanFilter(accountSubtypes []AccountSubtype, ) *LoanFilter {
	this := LoanFilter{}
	this.AccountSubtypes = accountSubtypes
	return &this
}

// NewLoanFilterWithDefaults instantiates a new LoanFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoanFilterWithDefaults() *LoanFilter {
	this := LoanFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value
func (o *LoanFilter) GetAccountSubtypes() []AccountSubtype {
	if o == nil  {
		var ret []AccountSubtype
		return ret
	}

	return o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value
// and a boolean to check if the value has been set.
func (o *LoanFilter) GetAccountSubtypesOk() (*[]AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountSubtypes, true
}

// SetAccountSubtypes sets field value
func (o *LoanFilter) SetAccountSubtypes(v []AccountSubtype) {
	o.AccountSubtypes = v
}

func (o LoanFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LoanFilter) UnmarshalJSON(bytes []byte) (err error) {
	varLoanFilter := _LoanFilter{}

	if err = json.Unmarshal(bytes, &varLoanFilter); err == nil {
		*o = LoanFilter(varLoanFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLoanFilter struct {
	value *LoanFilter
	isSet bool
}

func (v NullableLoanFilter) Get() *LoanFilter {
	return v.value
}

func (v *NullableLoanFilter) Set(val *LoanFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLoanFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLoanFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoanFilter(val *LoanFilter) *NullableLoanFilter {
	return &NullableLoanFilter{value: val, isSet: true}
}

func (v NullableLoanFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoanFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


