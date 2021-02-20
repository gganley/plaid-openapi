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

// LinkTokenCreateRequestAccountSubtypesCredit A filter to apply to `credit`-type accounts
type LinkTokenCreateRequestAccountSubtypesCredit struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](/docs/api/accounts#accounts-schema). 
	AccountSubtypes *[]AccountSubtype `json:"account_subtypes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateRequestAccountSubtypesCredit LinkTokenCreateRequestAccountSubtypesCredit

// NewLinkTokenCreateRequestAccountSubtypesCredit instantiates a new LinkTokenCreateRequestAccountSubtypesCredit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestAccountSubtypesCredit() *LinkTokenCreateRequestAccountSubtypesCredit {
	this := LinkTokenCreateRequestAccountSubtypesCredit{}
	return &this
}

// NewLinkTokenCreateRequestAccountSubtypesCreditWithDefaults instantiates a new LinkTokenCreateRequestAccountSubtypesCredit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestAccountSubtypesCreditWithDefaults() *LinkTokenCreateRequestAccountSubtypesCredit {
	this := LinkTokenCreateRequestAccountSubtypesCredit{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value if set, zero value otherwise.
func (o *LinkTokenCreateRequestAccountSubtypesCredit) GetAccountSubtypes() []AccountSubtype {
	if o == nil || o.AccountSubtypes == nil {
		var ret []AccountSubtype
		return ret
	}
	return *o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestAccountSubtypesCredit) GetAccountSubtypesOk() (*[]AccountSubtype, bool) {
	if o == nil || o.AccountSubtypes == nil {
		return nil, false
	}
	return o.AccountSubtypes, true
}

// HasAccountSubtypes returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestAccountSubtypesCredit) HasAccountSubtypes() bool {
	if o != nil && o.AccountSubtypes != nil {
		return true
	}

	return false
}

// SetAccountSubtypes gets a reference to the given []AccountSubtype and assigns it to the AccountSubtypes field.
func (o *LinkTokenCreateRequestAccountSubtypesCredit) SetAccountSubtypes(v []AccountSubtype) {
	o.AccountSubtypes = &v
}

func (o LinkTokenCreateRequestAccountSubtypesCredit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountSubtypes != nil {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateRequestAccountSubtypesCredit) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateRequestAccountSubtypesCredit := _LinkTokenCreateRequestAccountSubtypesCredit{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateRequestAccountSubtypesCredit); err == nil {
		*o = LinkTokenCreateRequestAccountSubtypesCredit(varLinkTokenCreateRequestAccountSubtypesCredit)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateRequestAccountSubtypesCredit struct {
	value *LinkTokenCreateRequestAccountSubtypesCredit
	isSet bool
}

func (v NullableLinkTokenCreateRequestAccountSubtypesCredit) Get() *LinkTokenCreateRequestAccountSubtypesCredit {
	return v.value
}

func (v *NullableLinkTokenCreateRequestAccountSubtypesCredit) Set(val *LinkTokenCreateRequestAccountSubtypesCredit) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestAccountSubtypesCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestAccountSubtypesCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestAccountSubtypesCredit(val *LinkTokenCreateRequestAccountSubtypesCredit) *NullableLinkTokenCreateRequestAccountSubtypesCredit {
	return &NullableLinkTokenCreateRequestAccountSubtypesCredit{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestAccountSubtypesCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestAccountSubtypesCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

