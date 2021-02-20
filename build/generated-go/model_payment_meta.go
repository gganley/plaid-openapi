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

// PaymentMeta Transaction information specific to inter-bank transfers. If the transaction was not an inter-bank transfer, all fields will be `null`.  If the `transaction` object was returned by a Transactions endpoint such as `/transactions/get`, the `payment_meta` key will always appear, but no data elements are guaranteed. If the `transaction` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
type PaymentMeta struct {
	// The transaction reference number supplied by the financial institution.
	ReferenceNumber NullableString `json:"reference_number,omitempty"`
	// The ACH PPD ID for the payer.
	PpdId NullableString `json:"ppd_id,omitempty"`
	// For transfers, the party that is receiving the transaction.
	Payee NullableString `json:"payee,omitempty"`
	// The party initiating a wire transfer. Will be `null` if the transaction is not a wire transfer.
	ByOrderOf NullableString `json:"by_order_of,omitempty"`
	// For transfers, the party that is paying the transaction.
	Payer NullableString `json:"payer,omitempty"`
	// The type of transfer, e.g. 'ACH'
	PaymentMethod NullableString `json:"payment_method,omitempty"`
	// The name of the payment processor
	PaymentProcessor NullableString `json:"payment_processor,omitempty"`
	// The payer-supplied description of the transfer.
	Reason               NullableString `json:"reason,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMeta PaymentMeta

// NewPaymentMeta instantiates a new PaymentMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMeta() *PaymentMeta {
	this := PaymentMeta{}
	return &this
}

// NewPaymentMetaWithDefaults instantiates a new PaymentMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMetaWithDefaults() *PaymentMeta {
	this := PaymentMeta{}
	return &this
}

// GetReferenceNumber returns the ReferenceNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetReferenceNumber() string {
	if o == nil || o.ReferenceNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReferenceNumber.Get()
}

// GetReferenceNumberOk returns a tuple with the ReferenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetReferenceNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceNumber.Get(), o.ReferenceNumber.IsSet()
}

// HasReferenceNumber returns a boolean if a field has been set.
func (o *PaymentMeta) HasReferenceNumber() bool {
	if o != nil && o.ReferenceNumber.IsSet() {
		return true
	}

	return false
}

// SetReferenceNumber gets a reference to the given NullableString and assigns it to the ReferenceNumber field.
func (o *PaymentMeta) SetReferenceNumber(v string) {
	o.ReferenceNumber.Set(&v)
}

// SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil
func (o *PaymentMeta) SetReferenceNumberNil() {
	o.ReferenceNumber.Set(nil)
}

// UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
func (o *PaymentMeta) UnsetReferenceNumber() {
	o.ReferenceNumber.Unset()
}

// GetPpdId returns the PpdId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetPpdId() string {
	if o == nil || o.PpdId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PpdId.Get()
}

// GetPpdIdOk returns a tuple with the PpdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPpdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PpdId.Get(), o.PpdId.IsSet()
}

// HasPpdId returns a boolean if a field has been set.
func (o *PaymentMeta) HasPpdId() bool {
	if o != nil && o.PpdId.IsSet() {
		return true
	}

	return false
}

// SetPpdId gets a reference to the given NullableString and assigns it to the PpdId field.
func (o *PaymentMeta) SetPpdId(v string) {
	o.PpdId.Set(&v)
}

// SetPpdIdNil sets the value for PpdId to be an explicit nil
func (o *PaymentMeta) SetPpdIdNil() {
	o.PpdId.Set(nil)
}

// UnsetPpdId ensures that no value is present for PpdId, not even an explicit nil
func (o *PaymentMeta) UnsetPpdId() {
	o.PpdId.Unset()
}

// GetPayee returns the Payee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetPayee() string {
	if o == nil || o.Payee.Get() == nil {
		var ret string
		return ret
	}
	return *o.Payee.Get()
}

// GetPayeeOk returns a tuple with the Payee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPayeeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payee.Get(), o.Payee.IsSet()
}

// HasPayee returns a boolean if a field has been set.
func (o *PaymentMeta) HasPayee() bool {
	if o != nil && o.Payee.IsSet() {
		return true
	}

	return false
}

// SetPayee gets a reference to the given NullableString and assigns it to the Payee field.
func (o *PaymentMeta) SetPayee(v string) {
	o.Payee.Set(&v)
}

// SetPayeeNil sets the value for Payee to be an explicit nil
func (o *PaymentMeta) SetPayeeNil() {
	o.Payee.Set(nil)
}

// UnsetPayee ensures that no value is present for Payee, not even an explicit nil
func (o *PaymentMeta) UnsetPayee() {
	o.Payee.Unset()
}

// GetByOrderOf returns the ByOrderOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetByOrderOf() string {
	if o == nil || o.ByOrderOf.Get() == nil {
		var ret string
		return ret
	}
	return *o.ByOrderOf.Get()
}

// GetByOrderOfOk returns a tuple with the ByOrderOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetByOrderOfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ByOrderOf.Get(), o.ByOrderOf.IsSet()
}

// HasByOrderOf returns a boolean if a field has been set.
func (o *PaymentMeta) HasByOrderOf() bool {
	if o != nil && o.ByOrderOf.IsSet() {
		return true
	}

	return false
}

// SetByOrderOf gets a reference to the given NullableString and assigns it to the ByOrderOf field.
func (o *PaymentMeta) SetByOrderOf(v string) {
	o.ByOrderOf.Set(&v)
}

// SetByOrderOfNil sets the value for ByOrderOf to be an explicit nil
func (o *PaymentMeta) SetByOrderOfNil() {
	o.ByOrderOf.Set(nil)
}

// UnsetByOrderOf ensures that no value is present for ByOrderOf, not even an explicit nil
func (o *PaymentMeta) UnsetByOrderOf() {
	o.ByOrderOf.Unset()
}

// GetPayer returns the Payer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetPayer() string {
	if o == nil || o.Payer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Payer.Get()
}

// GetPayerOk returns a tuple with the Payer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPayerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payer.Get(), o.Payer.IsSet()
}

// HasPayer returns a boolean if a field has been set.
func (o *PaymentMeta) HasPayer() bool {
	if o != nil && o.Payer.IsSet() {
		return true
	}

	return false
}

// SetPayer gets a reference to the given NullableString and assigns it to the Payer field.
func (o *PaymentMeta) SetPayer(v string) {
	o.Payer.Set(&v)
}

// SetPayerNil sets the value for Payer to be an explicit nil
func (o *PaymentMeta) SetPayerNil() {
	o.Payer.Set(nil)
}

// UnsetPayer ensures that no value is present for Payer, not even an explicit nil
func (o *PaymentMeta) UnsetPayer() {
	o.Payer.Unset()
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetPaymentMethod() string {
	if o == nil || o.PaymentMethod.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentMethod.Get()
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPaymentMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentMethod.Get(), o.PaymentMethod.IsSet()
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *PaymentMeta) HasPaymentMethod() bool {
	if o != nil && o.PaymentMethod.IsSet() {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given NullableString and assigns it to the PaymentMethod field.
func (o *PaymentMeta) SetPaymentMethod(v string) {
	o.PaymentMethod.Set(&v)
}

// SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil
func (o *PaymentMeta) SetPaymentMethodNil() {
	o.PaymentMethod.Set(nil)
}

// UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
func (o *PaymentMeta) UnsetPaymentMethod() {
	o.PaymentMethod.Unset()
}

// GetPaymentProcessor returns the PaymentProcessor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetPaymentProcessor() string {
	if o == nil || o.PaymentProcessor.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentProcessor.Get()
}

// GetPaymentProcessorOk returns a tuple with the PaymentProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetPaymentProcessorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PaymentProcessor.Get(), o.PaymentProcessor.IsSet()
}

// HasPaymentProcessor returns a boolean if a field has been set.
func (o *PaymentMeta) HasPaymentProcessor() bool {
	if o != nil && o.PaymentProcessor.IsSet() {
		return true
	}

	return false
}

// SetPaymentProcessor gets a reference to the given NullableString and assigns it to the PaymentProcessor field.
func (o *PaymentMeta) SetPaymentProcessor(v string) {
	o.PaymentProcessor.Set(&v)
}

// SetPaymentProcessorNil sets the value for PaymentProcessor to be an explicit nil
func (o *PaymentMeta) SetPaymentProcessorNil() {
	o.PaymentProcessor.Set(nil)
}

// UnsetPaymentProcessor ensures that no value is present for PaymentProcessor, not even an explicit nil
func (o *PaymentMeta) UnsetPaymentProcessor() {
	o.PaymentProcessor.Unset()
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMeta) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMeta) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *PaymentMeta) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *PaymentMeta) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil
func (o *PaymentMeta) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *PaymentMeta) UnsetReason() {
	o.Reason.Unset()
}

func (o PaymentMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceNumber.IsSet() {
		toSerialize["reference_number"] = o.ReferenceNumber.Get()
	}
	if o.PpdId.IsSet() {
		toSerialize["ppd_id"] = o.PpdId.Get()
	}
	if o.Payee.IsSet() {
		toSerialize["payee"] = o.Payee.Get()
	}
	if o.ByOrderOf.IsSet() {
		toSerialize["by_order_of"] = o.ByOrderOf.Get()
	}
	if o.Payer.IsSet() {
		toSerialize["payer"] = o.Payer.Get()
	}
	if o.PaymentMethod.IsSet() {
		toSerialize["payment_method"] = o.PaymentMethod.Get()
	}
	if o.PaymentProcessor.IsSet() {
		toSerialize["payment_processor"] = o.PaymentProcessor.Get()
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentMeta) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentMeta := _PaymentMeta{}

	if err = json.Unmarshal(bytes, &varPaymentMeta); err == nil {
		*o = PaymentMeta(varPaymentMeta)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reference_number")
		delete(additionalProperties, "ppd_id")
		delete(additionalProperties, "payee")
		delete(additionalProperties, "by_order_of")
		delete(additionalProperties, "payer")
		delete(additionalProperties, "payment_method")
		delete(additionalProperties, "payment_processor")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMeta struct {
	value *PaymentMeta
	isSet bool
}

func (v NullablePaymentMeta) Get() *PaymentMeta {
	return v.value
}

func (v *NullablePaymentMeta) Set(val *PaymentMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMeta(val *PaymentMeta) *NullablePaymentMeta {
	return &NullablePaymentMeta{value: val, isSet: true}
}

func (v NullablePaymentMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
