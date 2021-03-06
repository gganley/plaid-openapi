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

// IncomeSummary The verified fields from a paystub verification. All fields are provided as reported on the paystub.
type IncomeSummary struct {
	EmployerName         *EmployerIncomeSummaryFieldString  `json:"employer_name,omitempty"`
	EmployeeName         *EmployeeIncomeSummaryFieldString  `json:"employee_name,omitempty"`
	YtdGrossIncome       *YTDGrossIncomeSummaryFieldNumber  `json:"ytd_gross_income,omitempty"`
	YtdNetIncome         *YTDNetIncomeSummaryFieldNumber    `json:"ytd_net_income,omitempty"`
	PayFrequency         NullablePayFrequency               `json:"pay_frequency,omitempty"`
	ProjectedWage        *ProjectedIncomeSummaryFieldNumber `json:"projected_wage,omitempty"`
	VerifiedTransaction  NullableTransactionData            `json:"verified_transaction,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IncomeSummary IncomeSummary

// NewIncomeSummary instantiates a new IncomeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeSummary() *IncomeSummary {
	this := IncomeSummary{}
	return &this
}

// NewIncomeSummaryWithDefaults instantiates a new IncomeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeSummaryWithDefaults() *IncomeSummary {
	this := IncomeSummary{}
	return &this
}

// GetEmployerName returns the EmployerName field value if set, zero value otherwise.
func (o *IncomeSummary) GetEmployerName() EmployerIncomeSummaryFieldString {
	if o == nil || o.EmployerName == nil {
		var ret EmployerIncomeSummaryFieldString
		return ret
	}
	return *o.EmployerName
}

// GetEmployerNameOk returns a tuple with the EmployerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetEmployerNameOk() (*EmployerIncomeSummaryFieldString, bool) {
	if o == nil || o.EmployerName == nil {
		return nil, false
	}
	return o.EmployerName, true
}

// HasEmployerName returns a boolean if a field has been set.
func (o *IncomeSummary) HasEmployerName() bool {
	if o != nil && o.EmployerName != nil {
		return true
	}

	return false
}

// SetEmployerName gets a reference to the given EmployerIncomeSummaryFieldString and assigns it to the EmployerName field.
func (o *IncomeSummary) SetEmployerName(v EmployerIncomeSummaryFieldString) {
	o.EmployerName = &v
}

// GetEmployeeName returns the EmployeeName field value if set, zero value otherwise.
func (o *IncomeSummary) GetEmployeeName() EmployeeIncomeSummaryFieldString {
	if o == nil || o.EmployeeName == nil {
		var ret EmployeeIncomeSummaryFieldString
		return ret
	}
	return *o.EmployeeName
}

// GetEmployeeNameOk returns a tuple with the EmployeeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetEmployeeNameOk() (*EmployeeIncomeSummaryFieldString, bool) {
	if o == nil || o.EmployeeName == nil {
		return nil, false
	}
	return o.EmployeeName, true
}

// HasEmployeeName returns a boolean if a field has been set.
func (o *IncomeSummary) HasEmployeeName() bool {
	if o != nil && o.EmployeeName != nil {
		return true
	}

	return false
}

// SetEmployeeName gets a reference to the given EmployeeIncomeSummaryFieldString and assigns it to the EmployeeName field.
func (o *IncomeSummary) SetEmployeeName(v EmployeeIncomeSummaryFieldString) {
	o.EmployeeName = &v
}

// GetYtdGrossIncome returns the YtdGrossIncome field value if set, zero value otherwise.
func (o *IncomeSummary) GetYtdGrossIncome() YTDGrossIncomeSummaryFieldNumber {
	if o == nil || o.YtdGrossIncome == nil {
		var ret YTDGrossIncomeSummaryFieldNumber
		return ret
	}
	return *o.YtdGrossIncome
}

// GetYtdGrossIncomeOk returns a tuple with the YtdGrossIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetYtdGrossIncomeOk() (*YTDGrossIncomeSummaryFieldNumber, bool) {
	if o == nil || o.YtdGrossIncome == nil {
		return nil, false
	}
	return o.YtdGrossIncome, true
}

// HasYtdGrossIncome returns a boolean if a field has been set.
func (o *IncomeSummary) HasYtdGrossIncome() bool {
	if o != nil && o.YtdGrossIncome != nil {
		return true
	}

	return false
}

// SetYtdGrossIncome gets a reference to the given YTDGrossIncomeSummaryFieldNumber and assigns it to the YtdGrossIncome field.
func (o *IncomeSummary) SetYtdGrossIncome(v YTDGrossIncomeSummaryFieldNumber) {
	o.YtdGrossIncome = &v
}

// GetYtdNetIncome returns the YtdNetIncome field value if set, zero value otherwise.
func (o *IncomeSummary) GetYtdNetIncome() YTDNetIncomeSummaryFieldNumber {
	if o == nil || o.YtdNetIncome == nil {
		var ret YTDNetIncomeSummaryFieldNumber
		return ret
	}
	return *o.YtdNetIncome
}

// GetYtdNetIncomeOk returns a tuple with the YtdNetIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetYtdNetIncomeOk() (*YTDNetIncomeSummaryFieldNumber, bool) {
	if o == nil || o.YtdNetIncome == nil {
		return nil, false
	}
	return o.YtdNetIncome, true
}

// HasYtdNetIncome returns a boolean if a field has been set.
func (o *IncomeSummary) HasYtdNetIncome() bool {
	if o != nil && o.YtdNetIncome != nil {
		return true
	}

	return false
}

// SetYtdNetIncome gets a reference to the given YTDNetIncomeSummaryFieldNumber and assigns it to the YtdNetIncome field.
func (o *IncomeSummary) SetYtdNetIncome(v YTDNetIncomeSummaryFieldNumber) {
	o.YtdNetIncome = &v
}

// GetPayFrequency returns the PayFrequency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeSummary) GetPayFrequency() PayFrequency {
	if o == nil || o.PayFrequency.Get() == nil {
		var ret PayFrequency
		return ret
	}
	return *o.PayFrequency.Get()
}

// GetPayFrequencyOk returns a tuple with the PayFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeSummary) GetPayFrequencyOk() (*PayFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return o.PayFrequency.Get(), o.PayFrequency.IsSet()
}

// HasPayFrequency returns a boolean if a field has been set.
func (o *IncomeSummary) HasPayFrequency() bool {
	if o != nil && o.PayFrequency.IsSet() {
		return true
	}

	return false
}

// SetPayFrequency gets a reference to the given NullablePayFrequency and assigns it to the PayFrequency field.
func (o *IncomeSummary) SetPayFrequency(v PayFrequency) {
	o.PayFrequency.Set(&v)
}

// SetPayFrequencyNil sets the value for PayFrequency to be an explicit nil
func (o *IncomeSummary) SetPayFrequencyNil() {
	o.PayFrequency.Set(nil)
}

// UnsetPayFrequency ensures that no value is present for PayFrequency, not even an explicit nil
func (o *IncomeSummary) UnsetPayFrequency() {
	o.PayFrequency.Unset()
}

// GetProjectedWage returns the ProjectedWage field value if set, zero value otherwise.
func (o *IncomeSummary) GetProjectedWage() ProjectedIncomeSummaryFieldNumber {
	if o == nil || o.ProjectedWage == nil {
		var ret ProjectedIncomeSummaryFieldNumber
		return ret
	}
	return *o.ProjectedWage
}

// GetProjectedWageOk returns a tuple with the ProjectedWage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeSummary) GetProjectedWageOk() (*ProjectedIncomeSummaryFieldNumber, bool) {
	if o == nil || o.ProjectedWage == nil {
		return nil, false
	}
	return o.ProjectedWage, true
}

// HasProjectedWage returns a boolean if a field has been set.
func (o *IncomeSummary) HasProjectedWage() bool {
	if o != nil && o.ProjectedWage != nil {
		return true
	}

	return false
}

// SetProjectedWage gets a reference to the given ProjectedIncomeSummaryFieldNumber and assigns it to the ProjectedWage field.
func (o *IncomeSummary) SetProjectedWage(v ProjectedIncomeSummaryFieldNumber) {
	o.ProjectedWage = &v
}

// GetVerifiedTransaction returns the VerifiedTransaction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IncomeSummary) GetVerifiedTransaction() TransactionData {
	if o == nil || o.VerifiedTransaction.Get() == nil {
		var ret TransactionData
		return ret
	}
	return *o.VerifiedTransaction.Get()
}

// GetVerifiedTransactionOk returns a tuple with the VerifiedTransaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IncomeSummary) GetVerifiedTransactionOk() (*TransactionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.VerifiedTransaction.Get(), o.VerifiedTransaction.IsSet()
}

// HasVerifiedTransaction returns a boolean if a field has been set.
func (o *IncomeSummary) HasVerifiedTransaction() bool {
	if o != nil && o.VerifiedTransaction.IsSet() {
		return true
	}

	return false
}

// SetVerifiedTransaction gets a reference to the given NullableTransactionData and assigns it to the VerifiedTransaction field.
func (o *IncomeSummary) SetVerifiedTransaction(v TransactionData) {
	o.VerifiedTransaction.Set(&v)
}

// SetVerifiedTransactionNil sets the value for VerifiedTransaction to be an explicit nil
func (o *IncomeSummary) SetVerifiedTransactionNil() {
	o.VerifiedTransaction.Set(nil)
}

// UnsetVerifiedTransaction ensures that no value is present for VerifiedTransaction, not even an explicit nil
func (o *IncomeSummary) UnsetVerifiedTransaction() {
	o.VerifiedTransaction.Unset()
}

func (o IncomeSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmployerName != nil {
		toSerialize["employer_name"] = o.EmployerName
	}
	if o.EmployeeName != nil {
		toSerialize["employee_name"] = o.EmployeeName
	}
	if o.YtdGrossIncome != nil {
		toSerialize["ytd_gross_income"] = o.YtdGrossIncome
	}
	if o.YtdNetIncome != nil {
		toSerialize["ytd_net_income"] = o.YtdNetIncome
	}
	if o.PayFrequency.IsSet() {
		toSerialize["pay_frequency"] = o.PayFrequency.Get()
	}
	if o.ProjectedWage != nil {
		toSerialize["projected_wage"] = o.ProjectedWage
	}
	if o.VerifiedTransaction.IsSet() {
		toSerialize["verified_transaction"] = o.VerifiedTransaction.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IncomeSummary) UnmarshalJSON(bytes []byte) (err error) {
	varIncomeSummary := _IncomeSummary{}

	if err = json.Unmarshal(bytes, &varIncomeSummary); err == nil {
		*o = IncomeSummary(varIncomeSummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "employer_name")
		delete(additionalProperties, "employee_name")
		delete(additionalProperties, "ytd_gross_income")
		delete(additionalProperties, "ytd_net_income")
		delete(additionalProperties, "pay_frequency")
		delete(additionalProperties, "projected_wage")
		delete(additionalProperties, "verified_transaction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIncomeSummary struct {
	value *IncomeSummary
	isSet bool
}

func (v NullableIncomeSummary) Get() *IncomeSummary {
	return v.value
}

func (v *NullableIncomeSummary) Set(val *IncomeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeSummary(val *IncomeSummary) *NullableIncomeSummary {
	return &NullableIncomeSummary{value: val, isSet: true}
}

func (v NullableIncomeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
