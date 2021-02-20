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

// AccountIdentity struct for AccountIdentity
type AccountIdentity struct {
	// Plaid’s unique identifier for the account. This value will not change unless Plaid can't reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new `account_id` will be assigned to the account.  The `account_id` can also change if the `access_token` is deleted and the same credentials that were used to generate that `access_token` are used to generate a new `access_token` on a later date. In that case, the new `account_id` will be different from the old `account_id`.  Like all Plaid identifiers, the `account_id` is case sensitive.
	AccountId string `json:"account_id"`
	Balances AccountBalance `json:"balances"`
	// The last 2-4 alphanumeric characters of an account's official account number. Note that the mask may be non-unique between an Item's accounts, and it may also not match the mask that the bank displays to the user.
	Mask NullableString `json:"mask,omitempty"`
	// The name of the account, either assigned by the user or by the financial institution itself
	Name string `json:"name"`
	// The official name of the account as given by the financial institution
	OfficialName NullableString `json:"official_name,omitempty"`
	Type AccountType `json:"type"`
	Subtype NullableAccountSubtype `json:"subtype"`
	// The current verification status of an Auth Item initiated through Automated or Manual micro-deposits.  Returned for Auth Items only.  `pending_automatic_verification`: The Item is pending automatic verification  `pending_manual_verification`: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the two amounts.  `automatically_verified`: The Item has successfully been automatically verified   `manually_verified`: The Item has successfully been manually verified  `verification_expired`: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  `verification_failed`: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.   
	VerificationStatus NullableString `json:"verification_status,omitempty"`
	// Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. Multiple owners on a single account will be represented in the same `owner` object, not in multiple owner objects within the array.
	Owners []Owner `json:"owners"`
	AdditionalProperties map[string]interface{}
}

type _AccountIdentity AccountIdentity

// NewAccountIdentity instantiates a new AccountIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountIdentity(accountId string, balances AccountBalance, name string, type_ AccountType, subtype NullableAccountSubtype, owners []Owner, ) *AccountIdentity {
	this := AccountIdentity{}
	this.AccountId = accountId
	this.Balances = balances
	this.Name = name
	this.Type = type_
	this.Subtype = subtype
	this.Owners = owners
	return &this
}

// NewAccountIdentityWithDefaults instantiates a new AccountIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountIdentityWithDefaults() *AccountIdentity {
	this := AccountIdentity{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccountIdentity) GetAccountId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccountIdentity) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccountIdentity) SetAccountId(v string) {
	o.AccountId = v
}

// GetBalances returns the Balances field value
func (o *AccountIdentity) GetBalances() AccountBalance {
	if o == nil  {
		var ret AccountBalance
		return ret
	}

	return o.Balances
}

// GetBalancesOk returns a tuple with the Balances field value
// and a boolean to check if the value has been set.
func (o *AccountIdentity) GetBalancesOk() (*AccountBalance, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Balances, true
}

// SetBalances sets field value
func (o *AccountIdentity) SetBalances(v AccountBalance) {
	o.Balances = v
}

// GetMask returns the Mask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentity) GetMask() string {
	if o == nil || o.Mask.Get() == nil {
		var ret string
		return ret
	}
	return *o.Mask.Get()
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentity) GetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Mask.Get(), o.Mask.IsSet()
}

// HasMask returns a boolean if a field has been set.
func (o *AccountIdentity) HasMask() bool {
	if o != nil && o.Mask.IsSet() {
		return true
	}

	return false
}

// SetMask gets a reference to the given NullableString and assigns it to the Mask field.
func (o *AccountIdentity) SetMask(v string) {
	o.Mask.Set(&v)
}
// SetMaskNil sets the value for Mask to be an explicit nil
func (o *AccountIdentity) SetMaskNil() {
	o.Mask.Set(nil)
}

// UnsetMask ensures that no value is present for Mask, not even an explicit nil
func (o *AccountIdentity) UnsetMask() {
	o.Mask.Unset()
}

// GetName returns the Name field value
func (o *AccountIdentity) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountIdentity) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountIdentity) SetName(v string) {
	o.Name = v
}

// GetOfficialName returns the OfficialName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentity) GetOfficialName() string {
	if o == nil || o.OfficialName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OfficialName.Get()
}

// GetOfficialNameOk returns a tuple with the OfficialName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentity) GetOfficialNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OfficialName.Get(), o.OfficialName.IsSet()
}

// HasOfficialName returns a boolean if a field has been set.
func (o *AccountIdentity) HasOfficialName() bool {
	if o != nil && o.OfficialName.IsSet() {
		return true
	}

	return false
}

// SetOfficialName gets a reference to the given NullableString and assigns it to the OfficialName field.
func (o *AccountIdentity) SetOfficialName(v string) {
	o.OfficialName.Set(&v)
}
// SetOfficialNameNil sets the value for OfficialName to be an explicit nil
func (o *AccountIdentity) SetOfficialNameNil() {
	o.OfficialName.Set(nil)
}

// UnsetOfficialName ensures that no value is present for OfficialName, not even an explicit nil
func (o *AccountIdentity) UnsetOfficialName() {
	o.OfficialName.Unset()
}

// GetType returns the Type field value
func (o *AccountIdentity) GetType() AccountType {
	if o == nil  {
		var ret AccountType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AccountIdentity) GetTypeOk() (*AccountType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AccountIdentity) SetType(v AccountType) {
	o.Type = v
}

// GetSubtype returns the Subtype field value
// If the value is explicit nil, the zero value for AccountSubtype will be returned
func (o *AccountIdentity) GetSubtype() AccountSubtype {
	if o == nil || o.Subtype.Get() == nil {
		var ret AccountSubtype
		return ret
	}

	return *o.Subtype.Get()
}

// GetSubtypeOk returns a tuple with the Subtype field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentity) GetSubtypeOk() (*AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Subtype.Get(), o.Subtype.IsSet()
}

// SetSubtype sets field value
func (o *AccountIdentity) SetSubtype(v AccountSubtype) {
	o.Subtype.Set(&v)
}

// GetVerificationStatus returns the VerificationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AccountIdentity) GetVerificationStatus() string {
	if o == nil || o.VerificationStatus.Get() == nil {
		var ret string
		return ret
	}
	return *o.VerificationStatus.Get()
}

// GetVerificationStatusOk returns a tuple with the VerificationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountIdentity) GetVerificationStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VerificationStatus.Get(), o.VerificationStatus.IsSet()
}

// HasVerificationStatus returns a boolean if a field has been set.
func (o *AccountIdentity) HasVerificationStatus() bool {
	if o != nil && o.VerificationStatus.IsSet() {
		return true
	}

	return false
}

// SetVerificationStatus gets a reference to the given NullableString and assigns it to the VerificationStatus field.
func (o *AccountIdentity) SetVerificationStatus(v string) {
	o.VerificationStatus.Set(&v)
}
// SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil
func (o *AccountIdentity) SetVerificationStatusNil() {
	o.VerificationStatus.Set(nil)
}

// UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
func (o *AccountIdentity) UnsetVerificationStatus() {
	o.VerificationStatus.Unset()
}

// GetOwners returns the Owners field value
func (o *AccountIdentity) GetOwners() []Owner {
	if o == nil  {
		var ret []Owner
		return ret
	}

	return o.Owners
}

// GetOwnersOk returns a tuple with the Owners field value
// and a boolean to check if the value has been set.
func (o *AccountIdentity) GetOwnersOk() (*[]Owner, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Owners, true
}

// SetOwners sets field value
func (o *AccountIdentity) SetOwners(v []Owner) {
	o.Owners = v
}

func (o AccountIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["balances"] = o.Balances
	}
	if o.Mask.IsSet() {
		toSerialize["mask"] = o.Mask.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.OfficialName.IsSet() {
		toSerialize["official_name"] = o.OfficialName.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["subtype"] = o.Subtype.Get()
	}
	if o.VerificationStatus.IsSet() {
		toSerialize["verification_status"] = o.VerificationStatus.Get()
	}
	if true {
		toSerialize["owners"] = o.Owners
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varAccountIdentity := _AccountIdentity{}

	if err = json.Unmarshal(bytes, &varAccountIdentity); err == nil {
		*o = AccountIdentity(varAccountIdentity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "balances")
		delete(additionalProperties, "mask")
		delete(additionalProperties, "name")
		delete(additionalProperties, "official_name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "subtype")
		delete(additionalProperties, "verification_status")
		delete(additionalProperties, "owners")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountIdentity struct {
	value *AccountIdentity
	isSet bool
}

func (v NullableAccountIdentity) Get() *AccountIdentity {
	return v.value
}

func (v *NullableAccountIdentity) Set(val *AccountIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountIdentity(val *AccountIdentity) *NullableAccountIdentity {
	return &NullableAccountIdentity{value: val, isSet: true}
}

func (v NullableAccountIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


