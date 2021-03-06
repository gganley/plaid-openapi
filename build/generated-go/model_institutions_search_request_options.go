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

// InstitutionsSearchRequestOptions An optional object to filter `/institutions/search` results.
type InstitutionsSearchRequestOptions struct {
	// Limit results to institutions with or without OAuth login flows. This is primarily relevant to institutions with European country codes
	Oauth *bool `json:"oauth,omitempty"`
	// When true, return the institution's homepage URL, logo and primary brand color. Learn more
	IncludeOptionalMetadata *bool                            `json:"include_optional_metadata,omitempty"`
	AccountFilter           *InstitutionsSearchAccountFilter `json:"account_filter,omitempty"`
	AdditionalProperties    map[string]interface{}
}

type _InstitutionsSearchRequestOptions InstitutionsSearchRequestOptions

// NewInstitutionsSearchRequestOptions instantiates a new InstitutionsSearchRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchRequestOptions() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	return &this
}

// NewInstitutionsSearchRequestOptionsWithDefaults instantiates a new InstitutionsSearchRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchRequestOptionsWithDefaults() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	return &this
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetOauth() bool {
	if o == nil || o.Oauth == nil {
		var ret bool
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetOauthOk() (*bool, bool) {
	if o == nil || o.Oauth == nil {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasOauth() bool {
	if o != nil && o.Oauth != nil {
		return true
	}

	return false
}

// SetOauth gets a reference to the given bool and assigns it to the Oauth field.
func (o *InstitutionsSearchRequestOptions) SetOauth(v bool) {
	o.Oauth = &v
}

// GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadata() bool {
	if o == nil || o.IncludeOptionalMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptionalMetadata
}

// GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeOptionalMetadata == nil {
		return nil, false
	}
	return o.IncludeOptionalMetadata, true
}

// HasIncludeOptionalMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludeOptionalMetadata() bool {
	if o != nil && o.IncludeOptionalMetadata != nil {
		return true
	}

	return false
}

// SetIncludeOptionalMetadata gets a reference to the given bool and assigns it to the IncludeOptionalMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludeOptionalMetadata(v bool) {
	o.IncludeOptionalMetadata = &v
}

// GetAccountFilter returns the AccountFilter field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetAccountFilter() InstitutionsSearchAccountFilter {
	if o == nil || o.AccountFilter == nil {
		var ret InstitutionsSearchAccountFilter
		return ret
	}
	return *o.AccountFilter
}

// GetAccountFilterOk returns a tuple with the AccountFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetAccountFilterOk() (*InstitutionsSearchAccountFilter, bool) {
	if o == nil || o.AccountFilter == nil {
		return nil, false
	}
	return o.AccountFilter, true
}

// HasAccountFilter returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasAccountFilter() bool {
	if o != nil && o.AccountFilter != nil {
		return true
	}

	return false
}

// SetAccountFilter gets a reference to the given InstitutionsSearchAccountFilter and assigns it to the AccountFilter field.
func (o *InstitutionsSearchRequestOptions) SetAccountFilter(v InstitutionsSearchAccountFilter) {
	o.AccountFilter = &v
}

func (o InstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Oauth != nil {
		toSerialize["oauth"] = o.Oauth
	}
	if o.IncludeOptionalMetadata != nil {
		toSerialize["include_optional_metadata"] = o.IncludeOptionalMetadata
	}
	if o.AccountFilter != nil {
		toSerialize["account_filter"] = o.AccountFilter
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InstitutionsSearchRequestOptions) UnmarshalJSON(bytes []byte) (err error) {
	varInstitutionsSearchRequestOptions := _InstitutionsSearchRequestOptions{}

	if err = json.Unmarshal(bytes, &varInstitutionsSearchRequestOptions); err == nil {
		*o = InstitutionsSearchRequestOptions(varInstitutionsSearchRequestOptions)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "oauth")
		delete(additionalProperties, "include_optional_metadata")
		delete(additionalProperties, "account_filter")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstitutionsSearchRequestOptions struct {
	value *InstitutionsSearchRequestOptions
	isSet bool
}

func (v NullableInstitutionsSearchRequestOptions) Get() *InstitutionsSearchRequestOptions {
	return v.value
}

func (v *NullableInstitutionsSearchRequestOptions) Set(val *InstitutionsSearchRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchRequestOptions(val *InstitutionsSearchRequestOptions) *NullableInstitutionsSearchRequestOptions {
	return &NullableInstitutionsSearchRequestOptions{value: val, isSet: true}
}

func (v NullableInstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
