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

// LinkTokenCreateRequestIncomeVerification Specifies options for initializing Link for use with the Income Verification (beta) product. This field is required if `income_verification` is included in the `products` array.
type LinkTokenCreateRequestIncomeVerification struct {
	// The `income_verification_id` of the verification instance, as provided by `/income/verification/create`.
	IncomeVerificationId string `json:"income_verification_id"`
	// The `asset_report_id` of an asset report associated with the user, as provided by `/asset_report/create`. Providing an `asset_report_id` is optional and can be used to verify the user through a streamlined flow. If provided, the bank linking flow will be skipped.
	AssetReportId        NullableString `json:"asset_report_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkTokenCreateRequestIncomeVerification LinkTokenCreateRequestIncomeVerification

// NewLinkTokenCreateRequestIncomeVerification instantiates a new LinkTokenCreateRequestIncomeVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestIncomeVerification(incomeVerificationId string) *LinkTokenCreateRequestIncomeVerification {
	this := LinkTokenCreateRequestIncomeVerification{}
	this.IncomeVerificationId = incomeVerificationId
	return &this
}

// NewLinkTokenCreateRequestIncomeVerificationWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestIncomeVerificationWithDefaults() *LinkTokenCreateRequestIncomeVerification {
	this := LinkTokenCreateRequestIncomeVerification{}
	return &this
}

// GetIncomeVerificationId returns the IncomeVerificationId field value
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IncomeVerificationId
}

// GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncomeVerificationId, true
}

// SetIncomeVerificationId sets field value
func (o *LinkTokenCreateRequestIncomeVerification) SetIncomeVerificationId(v string) {
	o.IncomeVerificationId = v
}

// GetAssetReportId returns the AssetReportId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportId() string {
	if o == nil || o.AssetReportId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AssetReportId.Get()
}

// GetAssetReportIdOk returns a tuple with the AssetReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssetReportId.Get(), o.AssetReportId.IsSet()
}

// HasAssetReportId returns a boolean if a field has been set.
func (o *LinkTokenCreateRequestIncomeVerification) HasAssetReportId() bool {
	if o != nil && o.AssetReportId.IsSet() {
		return true
	}

	return false
}

// SetAssetReportId gets a reference to the given NullableString and assigns it to the AssetReportId field.
func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportId(v string) {
	o.AssetReportId.Set(&v)
}

// SetAssetReportIdNil sets the value for AssetReportId to be an explicit nil
func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportIdNil() {
	o.AssetReportId.Set(nil)
}

// UnsetAssetReportId ensures that no value is present for AssetReportId, not even an explicit nil
func (o *LinkTokenCreateRequestIncomeVerification) UnsetAssetReportId() {
	o.AssetReportId.Unset()
}

func (o LinkTokenCreateRequestIncomeVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["income_verification_id"] = o.IncomeVerificationId
	}
	if o.AssetReportId.IsSet() {
		toSerialize["asset_report_id"] = o.AssetReportId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *LinkTokenCreateRequestIncomeVerification) UnmarshalJSON(bytes []byte) (err error) {
	varLinkTokenCreateRequestIncomeVerification := _LinkTokenCreateRequestIncomeVerification{}

	if err = json.Unmarshal(bytes, &varLinkTokenCreateRequestIncomeVerification); err == nil {
		*o = LinkTokenCreateRequestIncomeVerification(varLinkTokenCreateRequestIncomeVerification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "income_verification_id")
		delete(additionalProperties, "asset_report_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkTokenCreateRequestIncomeVerification struct {
	value *LinkTokenCreateRequestIncomeVerification
	isSet bool
}

func (v NullableLinkTokenCreateRequestIncomeVerification) Get() *LinkTokenCreateRequestIncomeVerification {
	return v.value
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) Set(val *LinkTokenCreateRequestIncomeVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestIncomeVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestIncomeVerification(val *LinkTokenCreateRequestIncomeVerification) *NullableLinkTokenCreateRequestIncomeVerification {
	return &NullableLinkTokenCreateRequestIncomeVerification{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestIncomeVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestIncomeVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
