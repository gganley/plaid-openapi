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

// AccountsBalanceGetRequestOptions An optional object to filter `/accounts/balance/get` results.
type AccountsBalanceGetRequestOptions struct {
	// A list of `account_ids` to retrieve for the Item. The default value is `null`.  Note: An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds           *[]string `json:"account_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountsBalanceGetRequestOptions AccountsBalanceGetRequestOptions

// NewAccountsBalanceGetRequestOptions instantiates a new AccountsBalanceGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsBalanceGetRequestOptions() *AccountsBalanceGetRequestOptions {
	this := AccountsBalanceGetRequestOptions{}
	return &this
}

// NewAccountsBalanceGetRequestOptionsWithDefaults instantiates a new AccountsBalanceGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsBalanceGetRequestOptionsWithDefaults() *AccountsBalanceGetRequestOptions {
	this := AccountsBalanceGetRequestOptions{}
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *AccountsBalanceGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountsBalanceGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *AccountsBalanceGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *AccountsBalanceGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

func (o AccountsBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountsBalanceGetRequestOptions) UnmarshalJSON(bytes []byte) (err error) {
	varAccountsBalanceGetRequestOptions := _AccountsBalanceGetRequestOptions{}

	if err = json.Unmarshal(bytes, &varAccountsBalanceGetRequestOptions); err == nil {
		*o = AccountsBalanceGetRequestOptions(varAccountsBalanceGetRequestOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsBalanceGetRequestOptions struct {
	value *AccountsBalanceGetRequestOptions
	isSet bool
}

func (v NullableAccountsBalanceGetRequestOptions) Get() *AccountsBalanceGetRequestOptions {
	return v.value
}

func (v *NullableAccountsBalanceGetRequestOptions) Set(val *AccountsBalanceGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsBalanceGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsBalanceGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsBalanceGetRequestOptions(val *AccountsBalanceGetRequestOptions) *NullableAccountsBalanceGetRequestOptions {
	return &NullableAccountsBalanceGetRequestOptions{value: val, isSet: true}
}

func (v NullableAccountsBalanceGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsBalanceGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
