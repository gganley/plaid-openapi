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

// BankTransferMigrateAccountRequest BankTransferMigrateAccountRequest defines the request schema for `/bank_transfer/migrate_account`
type BankTransferMigrateAccountRequest struct {
	// Your Plaid API `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`.
	Secret *string `json:"secret,omitempty"`
	// The user's account number.
	AccountNumber string `json:"account_number"`
	// The user's routing number.
	RoutingNumber string `json:"routing_number"`
	// The type of the bank account (`checking` or `savings`).
	AccountType          string `json:"account_type"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferMigrateAccountRequest BankTransferMigrateAccountRequest

// NewBankTransferMigrateAccountRequest instantiates a new BankTransferMigrateAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferMigrateAccountRequest(accountNumber string, routingNumber string, accountType string) *BankTransferMigrateAccountRequest {
	this := BankTransferMigrateAccountRequest{}
	this.AccountNumber = accountNumber
	this.RoutingNumber = routingNumber
	this.AccountType = accountType
	return &this
}

// NewBankTransferMigrateAccountRequestWithDefaults instantiates a new BankTransferMigrateAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferMigrateAccountRequestWithDefaults() *BankTransferMigrateAccountRequest {
	this := BankTransferMigrateAccountRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BankTransferMigrateAccountRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BankTransferMigrateAccountRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BankTransferMigrateAccountRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BankTransferMigrateAccountRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BankTransferMigrateAccountRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BankTransferMigrateAccountRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccountNumber returns the AccountNumber field value
func (o *BankTransferMigrateAccountRequest) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountRequest) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *BankTransferMigrateAccountRequest) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *BankTransferMigrateAccountRequest) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountRequest) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *BankTransferMigrateAccountRequest) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetAccountType returns the AccountType field value
func (o *BankTransferMigrateAccountRequest) GetAccountType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value
// and a boolean to check if the value has been set.
func (o *BankTransferMigrateAccountRequest) GetAccountTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountType, true
}

// SetAccountType sets field value
func (o *BankTransferMigrateAccountRequest) SetAccountType(v string) {
	o.AccountType = v
}

func (o BankTransferMigrateAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["account_number"] = o.AccountNumber
	}
	if true {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	if true {
		toSerialize["account_type"] = o.AccountType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferMigrateAccountRequest) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferMigrateAccountRequest := _BankTransferMigrateAccountRequest{}

	if err = json.Unmarshal(bytes, &varBankTransferMigrateAccountRequest); err == nil {
		*o = BankTransferMigrateAccountRequest(varBankTransferMigrateAccountRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "account_number")
		delete(additionalProperties, "routing_number")
		delete(additionalProperties, "account_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferMigrateAccountRequest struct {
	value *BankTransferMigrateAccountRequest
	isSet bool
}

func (v NullableBankTransferMigrateAccountRequest) Get() *BankTransferMigrateAccountRequest {
	return v.value
}

func (v *NullableBankTransferMigrateAccountRequest) Set(val *BankTransferMigrateAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferMigrateAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferMigrateAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferMigrateAccountRequest(val *BankTransferMigrateAccountRequest) *NullableBankTransferMigrateAccountRequest {
	return &NullableBankTransferMigrateAccountRequest{value: val, isSet: true}
}

func (v NullableBankTransferMigrateAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferMigrateAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
