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

// InvestmentsHoldingsGetResponse InvestmentsHoldingsGetResponse defines the response schema for `/investments/holdings/get`
type InvestmentsHoldingsGetResponse struct {
	// The accounts associated with the Item
	Accounts []AccountBase `json:"accounts"`
	// The holdings belonging to investment accounts associated with the Item. Details of the securities in the holdings are provided in the `securities` field. 
	Holdings []Holding `json:"holdings"`
	// Objects describing the securities held in the accounts associated with the Item. 
	Securities []Security `json:"securities"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsHoldingsGetResponse InvestmentsHoldingsGetResponse

// NewInvestmentsHoldingsGetResponse instantiates a new InvestmentsHoldingsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsHoldingsGetResponse(accounts []AccountBase, holdings []Holding, securities []Security, item Item, requestId string, ) *InvestmentsHoldingsGetResponse {
	this := InvestmentsHoldingsGetResponse{}
	this.Accounts = accounts
	this.Holdings = holdings
	this.Securities = securities
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewInvestmentsHoldingsGetResponseWithDefaults instantiates a new InvestmentsHoldingsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsHoldingsGetResponseWithDefaults() *InvestmentsHoldingsGetResponse {
	this := InvestmentsHoldingsGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *InvestmentsHoldingsGetResponse) GetAccounts() []AccountBase {
	if o == nil  {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *InvestmentsHoldingsGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetHoldings returns the Holdings field value
func (o *InvestmentsHoldingsGetResponse) GetHoldings() []Holding {
	if o == nil  {
		var ret []Holding
		return ret
	}

	return o.Holdings
}

// GetHoldingsOk returns a tuple with the Holdings field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetResponse) GetHoldingsOk() (*[]Holding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Holdings, true
}

// SetHoldings sets field value
func (o *InvestmentsHoldingsGetResponse) SetHoldings(v []Holding) {
	o.Holdings = v
}

// GetSecurities returns the Securities field value
func (o *InvestmentsHoldingsGetResponse) GetSecurities() []Security {
	if o == nil  {
		var ret []Security
		return ret
	}

	return o.Securities
}

// GetSecuritiesOk returns a tuple with the Securities field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetResponse) GetSecuritiesOk() (*[]Security, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Securities, true
}

// SetSecurities sets field value
func (o *InvestmentsHoldingsGetResponse) SetSecurities(v []Security) {
	o.Securities = v
}

// GetItem returns the Item field value
func (o *InvestmentsHoldingsGetResponse) GetItem() Item {
	if o == nil  {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *InvestmentsHoldingsGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *InvestmentsHoldingsGetResponse) GetRequestId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *InvestmentsHoldingsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o InvestmentsHoldingsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["holdings"] = o.Holdings
	}
	if true {
		toSerialize["securities"] = o.Securities
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsHoldingsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsHoldingsGetResponse := _InvestmentsHoldingsGetResponse{}

	if err = json.Unmarshal(bytes, &varInvestmentsHoldingsGetResponse); err == nil {
		*o = InvestmentsHoldingsGetResponse(varInvestmentsHoldingsGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "holdings")
		delete(additionalProperties, "securities")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsHoldingsGetResponse struct {
	value *InvestmentsHoldingsGetResponse
	isSet bool
}

func (v NullableInvestmentsHoldingsGetResponse) Get() *InvestmentsHoldingsGetResponse {
	return v.value
}

func (v *NullableInvestmentsHoldingsGetResponse) Set(val *InvestmentsHoldingsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsHoldingsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsHoldingsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsHoldingsGetResponse(val *InvestmentsHoldingsGetResponse) *NullableInvestmentsHoldingsGetResponse {
	return &NullableInvestmentsHoldingsGetResponse{value: val, isSet: true}
}

func (v NullableInvestmentsHoldingsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsHoldingsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

