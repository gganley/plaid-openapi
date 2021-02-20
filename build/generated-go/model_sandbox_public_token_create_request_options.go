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

// SandboxPublicTokenCreateRequestOptions An optional set of options to be used when configuring the Item. If specified, must not be `null`.
type SandboxPublicTokenCreateRequestOptions struct {
	// Specify a webhook to associate with the new Item.
	Webhook *string `json:"webhook,omitempty"`
	// Test username to use for the creation of the Sandbox Item. Default value is `user_good`.
	OverrideUsername NullableString `json:"override_username,omitempty"`
	// Test password to use for the creation of the Sandbox Item. Default value is `pass_good`.
	OverridePassword NullableString `json:"override_password,omitempty"`
	Transactions *SandboxPublicTokenCreateRequestOptionsTransactions `json:"transactions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SandboxPublicTokenCreateRequestOptions SandboxPublicTokenCreateRequestOptions

// NewSandboxPublicTokenCreateRequestOptions instantiates a new SandboxPublicTokenCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxPublicTokenCreateRequestOptions() *SandboxPublicTokenCreateRequestOptions {
	this := SandboxPublicTokenCreateRequestOptions{}
	var overrideUsername string = "user_good"
	this.OverrideUsername = *NewNullableString(&overrideUsername)
	var overridePassword string = "pass_good"
	this.OverridePassword = *NewNullableString(&overridePassword)
	return &this
}

// NewSandboxPublicTokenCreateRequestOptionsWithDefaults instantiates a new SandboxPublicTokenCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxPublicTokenCreateRequestOptionsWithDefaults() *SandboxPublicTokenCreateRequestOptions {
	this := SandboxPublicTokenCreateRequestOptions{}
	var overrideUsername string = "user_good"
	this.OverrideUsername = *NewNullableString(&overrideUsername)
	var overridePassword string = "pass_good"
	this.OverridePassword = *NewNullableString(&overridePassword)
	return &this
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *SandboxPublicTokenCreateRequestOptions) SetWebhook(v string) {
	o.Webhook = &v
}

// GetOverrideUsername returns the OverrideUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxPublicTokenCreateRequestOptions) GetOverrideUsername() string {
	if o == nil || o.OverrideUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverrideUsername.Get()
}

// GetOverrideUsernameOk returns a tuple with the OverrideUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxPublicTokenCreateRequestOptions) GetOverrideUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverrideUsername.Get(), o.OverrideUsername.IsSet()
}

// HasOverrideUsername returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptions) HasOverrideUsername() bool {
	if o != nil && o.OverrideUsername.IsSet() {
		return true
	}

	return false
}

// SetOverrideUsername gets a reference to the given NullableString and assigns it to the OverrideUsername field.
func (o *SandboxPublicTokenCreateRequestOptions) SetOverrideUsername(v string) {
	o.OverrideUsername.Set(&v)
}
// SetOverrideUsernameNil sets the value for OverrideUsername to be an explicit nil
func (o *SandboxPublicTokenCreateRequestOptions) SetOverrideUsernameNil() {
	o.OverrideUsername.Set(nil)
}

// UnsetOverrideUsername ensures that no value is present for OverrideUsername, not even an explicit nil
func (o *SandboxPublicTokenCreateRequestOptions) UnsetOverrideUsername() {
	o.OverrideUsername.Unset()
}

// GetOverridePassword returns the OverridePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SandboxPublicTokenCreateRequestOptions) GetOverridePassword() string {
	if o == nil || o.OverridePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.OverridePassword.Get()
}

// GetOverridePasswordOk returns a tuple with the OverridePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SandboxPublicTokenCreateRequestOptions) GetOverridePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverridePassword.Get(), o.OverridePassword.IsSet()
}

// HasOverridePassword returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptions) HasOverridePassword() bool {
	if o != nil && o.OverridePassword.IsSet() {
		return true
	}

	return false
}

// SetOverridePassword gets a reference to the given NullableString and assigns it to the OverridePassword field.
func (o *SandboxPublicTokenCreateRequestOptions) SetOverridePassword(v string) {
	o.OverridePassword.Set(&v)
}
// SetOverridePasswordNil sets the value for OverridePassword to be an explicit nil
func (o *SandboxPublicTokenCreateRequestOptions) SetOverridePasswordNil() {
	o.OverridePassword.Set(nil)
}

// UnsetOverridePassword ensures that no value is present for OverridePassword, not even an explicit nil
func (o *SandboxPublicTokenCreateRequestOptions) UnsetOverridePassword() {
	o.OverridePassword.Unset()
}

// GetTransactions returns the Transactions field value if set, zero value otherwise.
func (o *SandboxPublicTokenCreateRequestOptions) GetTransactions() SandboxPublicTokenCreateRequestOptionsTransactions {
	if o == nil || o.Transactions == nil {
		var ret SandboxPublicTokenCreateRequestOptionsTransactions
		return ret
	}
	return *o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxPublicTokenCreateRequestOptions) GetTransactionsOk() (*SandboxPublicTokenCreateRequestOptionsTransactions, bool) {
	if o == nil || o.Transactions == nil {
		return nil, false
	}
	return o.Transactions, true
}

// HasTransactions returns a boolean if a field has been set.
func (o *SandboxPublicTokenCreateRequestOptions) HasTransactions() bool {
	if o != nil && o.Transactions != nil {
		return true
	}

	return false
}

// SetTransactions gets a reference to the given SandboxPublicTokenCreateRequestOptionsTransactions and assigns it to the Transactions field.
func (o *SandboxPublicTokenCreateRequestOptions) SetTransactions(v SandboxPublicTokenCreateRequestOptionsTransactions) {
	o.Transactions = &v
}

func (o SandboxPublicTokenCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.OverrideUsername.IsSet() {
		toSerialize["override_username"] = o.OverrideUsername.Get()
	}
	if o.OverridePassword.IsSet() {
		toSerialize["override_password"] = o.OverridePassword.Get()
	}
	if o.Transactions != nil {
		toSerialize["transactions"] = o.Transactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxPublicTokenCreateRequestOptions) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxPublicTokenCreateRequestOptions := _SandboxPublicTokenCreateRequestOptions{}

	if err = json.Unmarshal(bytes, &varSandboxPublicTokenCreateRequestOptions); err == nil {
		*o = SandboxPublicTokenCreateRequestOptions(varSandboxPublicTokenCreateRequestOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook")
		delete(additionalProperties, "override_username")
		delete(additionalProperties, "override_password")
		delete(additionalProperties, "transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxPublicTokenCreateRequestOptions struct {
	value *SandboxPublicTokenCreateRequestOptions
	isSet bool
}

func (v NullableSandboxPublicTokenCreateRequestOptions) Get() *SandboxPublicTokenCreateRequestOptions {
	return v.value
}

func (v *NullableSandboxPublicTokenCreateRequestOptions) Set(val *SandboxPublicTokenCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxPublicTokenCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxPublicTokenCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxPublicTokenCreateRequestOptions(val *SandboxPublicTokenCreateRequestOptions) *NullableSandboxPublicTokenCreateRequestOptions {
	return &NullableSandboxPublicTokenCreateRequestOptions{value: val, isSet: true}
}

func (v NullableSandboxPublicTokenCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxPublicTokenCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


