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

// DepositSwitchGetResponse DepositSwitchGetResponse defines the response schema for `/deposit_switch/get`
type DepositSwitchGetResponse struct {
	// The ID of the deposit switch
	DepositSwitchId string `json:"deposit_switch_id"`
	// The ID of the bank account the direct deposit was switched to
	TargetAccountId NullableString `json:"target_account_id"`
	// The ID of the Item the direct deposit was switched to.
	TargetItemId NullableString `json:"target_item_id"`
	// The state of the deposit switch.
	State string `json:"state"`
	// When `true`, user’s direct deposit goes to multiple banks. When false, user’s direct deposit only goes to the target account. Always `null` if the deposit switch has not been completed.
	AccountHasMultipleAllocations NullableBool `json:"account_has_multiple_allocations"`
	// When `true`, the target account is allocated the remainder of direct deposit after all other allocations have been deducted. When `false`, user’s direct deposit is allocated as a percent or amount. Always `null` if the deposit switch has not been completed.
	IsAllocatedRemainder NullableBool `json:"is_allocated_remainder"`
	// The percentage of direct deposit allocated to the target account. Always `null` if the target account is not allocated a percentage or if the deposit switch has not been completed or if `is_allocated_remainder` is true.
	PercentAllocated NullableFloat32 `json:"percent_allocated"`
	// The dollar amount of direct deposit allocated to the target account. Always `null` if the target account is not allocated an amount or if the deposit switch has not been completed.
	AmountAllocated NullableFloat32 `json:"amount_allocated"`
	// ISO8601 date the deposit switch was created.
	DateCreated string `json:"date_created"`
	// ISO8601 date the deposit switch was completed. Always `null` if the deposit switch has not been completed.
	DateCompleted NullableString `json:"date_completed"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId            string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _DepositSwitchGetResponse DepositSwitchGetResponse

// NewDepositSwitchGetResponse instantiates a new DepositSwitchGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchGetResponse(depositSwitchId string, targetAccountId NullableString, targetItemId NullableString, state string, accountHasMultipleAllocations NullableBool, isAllocatedRemainder NullableBool, percentAllocated NullableFloat32, amountAllocated NullableFloat32, dateCreated string, dateCompleted NullableString, requestId string) *DepositSwitchGetResponse {
	this := DepositSwitchGetResponse{}
	this.DepositSwitchId = depositSwitchId
	this.TargetAccountId = targetAccountId
	this.TargetItemId = targetItemId
	this.State = state
	this.AccountHasMultipleAllocations = accountHasMultipleAllocations
	this.IsAllocatedRemainder = isAllocatedRemainder
	this.PercentAllocated = percentAllocated
	this.AmountAllocated = amountAllocated
	this.DateCreated = dateCreated
	this.DateCompleted = dateCompleted
	this.RequestId = requestId
	return &this
}

// NewDepositSwitchGetResponseWithDefaults instantiates a new DepositSwitchGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchGetResponseWithDefaults() *DepositSwitchGetResponse {
	this := DepositSwitchGetResponse{}
	return &this
}

// GetDepositSwitchId returns the DepositSwitchId field value
func (o *DepositSwitchGetResponse) GetDepositSwitchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositSwitchId
}

// GetDepositSwitchIdOk returns a tuple with the DepositSwitchId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetDepositSwitchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositSwitchId, true
}

// SetDepositSwitchId sets field value
func (o *DepositSwitchGetResponse) SetDepositSwitchId(v string) {
	o.DepositSwitchId = v
}

// GetTargetAccountId returns the TargetAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetTargetAccountId() string {
	if o == nil || o.TargetAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetAccountId.Get()
}

// GetTargetAccountIdOk returns a tuple with the TargetAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetTargetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetAccountId.Get(), o.TargetAccountId.IsSet()
}

// SetTargetAccountId sets field value
func (o *DepositSwitchGetResponse) SetTargetAccountId(v string) {
	o.TargetAccountId.Set(&v)
}

// GetTargetItemId returns the TargetItemId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetTargetItemId() string {
	if o == nil || o.TargetItemId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TargetItemId.Get()
}

// GetTargetItemIdOk returns a tuple with the TargetItemId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetTargetItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetItemId.Get(), o.TargetItemId.IsSet()
}

// SetTargetItemId sets field value
func (o *DepositSwitchGetResponse) SetTargetItemId(v string) {
	o.TargetItemId.Set(&v)
}

// GetState returns the State field value
func (o *DepositSwitchGetResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DepositSwitchGetResponse) SetState(v string) {
	o.State = v
}

// GetAccountHasMultipleAllocations returns the AccountHasMultipleAllocations field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocations() bool {
	if o == nil || o.AccountHasMultipleAllocations.Get() == nil {
		var ret bool
		return ret
	}

	return *o.AccountHasMultipleAllocations.Get()
}

// GetAccountHasMultipleAllocationsOk returns a tuple with the AccountHasMultipleAllocations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetAccountHasMultipleAllocationsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountHasMultipleAllocations.Get(), o.AccountHasMultipleAllocations.IsSet()
}

// SetAccountHasMultipleAllocations sets field value
func (o *DepositSwitchGetResponse) SetAccountHasMultipleAllocations(v bool) {
	o.AccountHasMultipleAllocations.Set(&v)
}

// GetIsAllocatedRemainder returns the IsAllocatedRemainder field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *DepositSwitchGetResponse) GetIsAllocatedRemainder() bool {
	if o == nil || o.IsAllocatedRemainder.Get() == nil {
		var ret bool
		return ret
	}

	return *o.IsAllocatedRemainder.Get()
}

// GetIsAllocatedRemainderOk returns a tuple with the IsAllocatedRemainder field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetIsAllocatedRemainderOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAllocatedRemainder.Get(), o.IsAllocatedRemainder.IsSet()
}

// SetIsAllocatedRemainder sets field value
func (o *DepositSwitchGetResponse) SetIsAllocatedRemainder(v bool) {
	o.IsAllocatedRemainder.Set(&v)
}

// GetPercentAllocated returns the PercentAllocated field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DepositSwitchGetResponse) GetPercentAllocated() float32 {
	if o == nil || o.PercentAllocated.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PercentAllocated.Get()
}

// GetPercentAllocatedOk returns a tuple with the PercentAllocated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetPercentAllocatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PercentAllocated.Get(), o.PercentAllocated.IsSet()
}

// SetPercentAllocated sets field value
func (o *DepositSwitchGetResponse) SetPercentAllocated(v float32) {
	o.PercentAllocated.Set(&v)
}

// GetAmountAllocated returns the AmountAllocated field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *DepositSwitchGetResponse) GetAmountAllocated() float32 {
	if o == nil || o.AmountAllocated.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AmountAllocated.Get()
}

// GetAmountAllocatedOk returns a tuple with the AmountAllocated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetAmountAllocatedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AmountAllocated.Get(), o.AmountAllocated.IsSet()
}

// SetAmountAllocated sets field value
func (o *DepositSwitchGetResponse) SetAmountAllocated(v float32) {
	o.AmountAllocated.Set(&v)
}

// GetDateCreated returns the DateCreated field value
func (o *DepositSwitchGetResponse) GetDateCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetDateCreatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DateCreated, true
}

// SetDateCreated sets field value
func (o *DepositSwitchGetResponse) SetDateCreated(v string) {
	o.DateCreated = v
}

// GetDateCompleted returns the DateCompleted field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DepositSwitchGetResponse) GetDateCompleted() string {
	if o == nil || o.DateCompleted.Get() == nil {
		var ret string
		return ret
	}

	return *o.DateCompleted.Get()
}

// GetDateCompletedOk returns a tuple with the DateCompleted field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchGetResponse) GetDateCompletedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCompleted.Get(), o.DateCompleted.IsSet()
}

// SetDateCompleted sets field value
func (o *DepositSwitchGetResponse) SetDateCompleted(v string) {
	o.DateCompleted.Set(&v)
}

// GetRequestId returns the RequestId field value
func (o *DepositSwitchGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *DepositSwitchGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o DepositSwitchGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deposit_switch_id"] = o.DepositSwitchId
	}
	if true {
		toSerialize["target_account_id"] = o.TargetAccountId.Get()
	}
	if true {
		toSerialize["target_item_id"] = o.TargetItemId.Get()
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["account_has_multiple_allocations"] = o.AccountHasMultipleAllocations.Get()
	}
	if true {
		toSerialize["is_allocated_remainder"] = o.IsAllocatedRemainder.Get()
	}
	if true {
		toSerialize["percent_allocated"] = o.PercentAllocated.Get()
	}
	if true {
		toSerialize["amount_allocated"] = o.AmountAllocated.Get()
	}
	if true {
		toSerialize["date_created"] = o.DateCreated
	}
	if true {
		toSerialize["date_completed"] = o.DateCompleted.Get()
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DepositSwitchGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDepositSwitchGetResponse := _DepositSwitchGetResponse{}

	if err = json.Unmarshal(bytes, &varDepositSwitchGetResponse); err == nil {
		*o = DepositSwitchGetResponse(varDepositSwitchGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "deposit_switch_id")
		delete(additionalProperties, "target_account_id")
		delete(additionalProperties, "target_item_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "account_has_multiple_allocations")
		delete(additionalProperties, "is_allocated_remainder")
		delete(additionalProperties, "percent_allocated")
		delete(additionalProperties, "amount_allocated")
		delete(additionalProperties, "date_created")
		delete(additionalProperties, "date_completed")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDepositSwitchGetResponse struct {
	value *DepositSwitchGetResponse
	isSet bool
}

func (v NullableDepositSwitchGetResponse) Get() *DepositSwitchGetResponse {
	return v.value
}

func (v *NullableDepositSwitchGetResponse) Set(val *DepositSwitchGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchGetResponse(val *DepositSwitchGetResponse) *NullableDepositSwitchGetResponse {
	return &NullableDepositSwitchGetResponse{value: val, isSet: true}
}

func (v NullableDepositSwitchGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
