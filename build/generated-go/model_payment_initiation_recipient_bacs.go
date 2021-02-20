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

// PaymentInitiationRecipientBacs struct for PaymentInitiationRecipientBacs
type PaymentInitiationRecipientBacs struct {
	Account string `json:"account"`
	SortCode string `json:"sort_code"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRecipientBacs PaymentInitiationRecipientBacs

// NewPaymentInitiationRecipientBacs instantiates a new PaymentInitiationRecipientBacs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipientBacs(account string, sortCode string, ) *PaymentInitiationRecipientBacs {
	this := PaymentInitiationRecipientBacs{}
	this.Account = account
	this.SortCode = sortCode
	return &this
}

// NewPaymentInitiationRecipientBacsWithDefaults instantiates a new PaymentInitiationRecipientBacs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientBacsWithDefaults() *PaymentInitiationRecipientBacs {
	this := PaymentInitiationRecipientBacs{}
	return &this
}

// GetAccount returns the Account field value
func (o *PaymentInitiationRecipientBacs) GetAccount() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientBacs) GetAccountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *PaymentInitiationRecipientBacs) SetAccount(v string) {
	o.Account = v
}

// GetSortCode returns the SortCode field value
func (o *PaymentInitiationRecipientBacs) GetSortCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SortCode
}

// GetSortCodeOk returns a tuple with the SortCode field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipientBacs) GetSortCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SortCode, true
}

// SetSortCode sets field value
func (o *PaymentInitiationRecipientBacs) SetSortCode(v string) {
	o.SortCode = v
}

func (o PaymentInitiationRecipientBacs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["sort_code"] = o.SortCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRecipientBacs) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRecipientBacs := _PaymentInitiationRecipientBacs{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRecipientBacs); err == nil {
		*o = PaymentInitiationRecipientBacs(varPaymentInitiationRecipientBacs)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account")
		delete(additionalProperties, "sort_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRecipientBacs struct {
	value *PaymentInitiationRecipientBacs
	isSet bool
}

func (v NullablePaymentInitiationRecipientBacs) Get() *PaymentInitiationRecipientBacs {
	return v.value
}

func (v *NullablePaymentInitiationRecipientBacs) Set(val *PaymentInitiationRecipientBacs) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipientBacs) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipientBacs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipientBacs(val *PaymentInitiationRecipientBacs) *NullablePaymentInitiationRecipientBacs {
	return &NullablePaymentInitiationRecipientBacs{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipientBacs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipientBacs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


