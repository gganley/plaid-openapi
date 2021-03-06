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

// JWKPublicKey A JSON Web Key (JWK) that can be used in conjunction with [JWT libraries](https://jwt.io/#libraries-io) to verify Plaid webhooks
type JWKPublicKey struct {
	// The alg member identifies the cryptographic algorithm family used with the key.
	Alg *string `json:"alg,omitempty"`
	// The crv member identifies the cryptographic curve used with the key.
	Crv *string `json:"crv,omitempty"`
	// The kid (Key ID) member can be used to match a specific key. This can be used, for instance, to choose among a set of keys within the JWK during key rollover.
	Kid *string `json:"kid,omitempty"`
	// The kty (key type) parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC.
	Kty *string `json:"kty,omitempty"`
	// The use (public key use) parameter identifies the intended use of the public key.
	Use *string `json:"use,omitempty"`
	// The x member contains the x coordinate for the elliptic curve point.
	X *string `json:"x,omitempty"`
	// The y member contains the y coordinate for the elliptic curve point.
	Y                    *string       `json:"y,omitempty"`
	CreatedAt            *int32        `json:"created_at,omitempty"`
	ExpiredAt            NullableInt32 `json:"expired_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JWKPublicKey JWKPublicKey

// NewJWKPublicKey instantiates a new JWKPublicKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJWKPublicKey() *JWKPublicKey {
	this := JWKPublicKey{}
	return &this
}

// NewJWKPublicKeyWithDefaults instantiates a new JWKPublicKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJWKPublicKeyWithDefaults() *JWKPublicKey {
	this := JWKPublicKey{}
	return &this
}

// GetAlg returns the Alg field value if set, zero value otherwise.
func (o *JWKPublicKey) GetAlg() string {
	if o == nil || o.Alg == nil {
		var ret string
		return ret
	}
	return *o.Alg
}

// GetAlgOk returns a tuple with the Alg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetAlgOk() (*string, bool) {
	if o == nil || o.Alg == nil {
		return nil, false
	}
	return o.Alg, true
}

// HasAlg returns a boolean if a field has been set.
func (o *JWKPublicKey) HasAlg() bool {
	if o != nil && o.Alg != nil {
		return true
	}

	return false
}

// SetAlg gets a reference to the given string and assigns it to the Alg field.
func (o *JWKPublicKey) SetAlg(v string) {
	o.Alg = &v
}

// GetCrv returns the Crv field value if set, zero value otherwise.
func (o *JWKPublicKey) GetCrv() string {
	if o == nil || o.Crv == nil {
		var ret string
		return ret
	}
	return *o.Crv
}

// GetCrvOk returns a tuple with the Crv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetCrvOk() (*string, bool) {
	if o == nil || o.Crv == nil {
		return nil, false
	}
	return o.Crv, true
}

// HasCrv returns a boolean if a field has been set.
func (o *JWKPublicKey) HasCrv() bool {
	if o != nil && o.Crv != nil {
		return true
	}

	return false
}

// SetCrv gets a reference to the given string and assigns it to the Crv field.
func (o *JWKPublicKey) SetCrv(v string) {
	o.Crv = &v
}

// GetKid returns the Kid field value if set, zero value otherwise.
func (o *JWKPublicKey) GetKid() string {
	if o == nil || o.Kid == nil {
		var ret string
		return ret
	}
	return *o.Kid
}

// GetKidOk returns a tuple with the Kid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetKidOk() (*string, bool) {
	if o == nil || o.Kid == nil {
		return nil, false
	}
	return o.Kid, true
}

// HasKid returns a boolean if a field has been set.
func (o *JWKPublicKey) HasKid() bool {
	if o != nil && o.Kid != nil {
		return true
	}

	return false
}

// SetKid gets a reference to the given string and assigns it to the Kid field.
func (o *JWKPublicKey) SetKid(v string) {
	o.Kid = &v
}

// GetKty returns the Kty field value if set, zero value otherwise.
func (o *JWKPublicKey) GetKty() string {
	if o == nil || o.Kty == nil {
		var ret string
		return ret
	}
	return *o.Kty
}

// GetKtyOk returns a tuple with the Kty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetKtyOk() (*string, bool) {
	if o == nil || o.Kty == nil {
		return nil, false
	}
	return o.Kty, true
}

// HasKty returns a boolean if a field has been set.
func (o *JWKPublicKey) HasKty() bool {
	if o != nil && o.Kty != nil {
		return true
	}

	return false
}

// SetKty gets a reference to the given string and assigns it to the Kty field.
func (o *JWKPublicKey) SetKty(v string) {
	o.Kty = &v
}

// GetUse returns the Use field value if set, zero value otherwise.
func (o *JWKPublicKey) GetUse() string {
	if o == nil || o.Use == nil {
		var ret string
		return ret
	}
	return *o.Use
}

// GetUseOk returns a tuple with the Use field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetUseOk() (*string, bool) {
	if o == nil || o.Use == nil {
		return nil, false
	}
	return o.Use, true
}

// HasUse returns a boolean if a field has been set.
func (o *JWKPublicKey) HasUse() bool {
	if o != nil && o.Use != nil {
		return true
	}

	return false
}

// SetUse gets a reference to the given string and assigns it to the Use field.
func (o *JWKPublicKey) SetUse(v string) {
	o.Use = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *JWKPublicKey) GetX() string {
	if o == nil || o.X == nil {
		var ret string
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetXOk() (*string, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *JWKPublicKey) HasX() bool {
	if o != nil && o.X != nil {
		return true
	}

	return false
}

// SetX gets a reference to the given string and assigns it to the X field.
func (o *JWKPublicKey) SetX(v string) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *JWKPublicKey) GetY() string {
	if o == nil || o.Y == nil {
		var ret string
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetYOk() (*string, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *JWKPublicKey) HasY() bool {
	if o != nil && o.Y != nil {
		return true
	}

	return false
}

// SetY gets a reference to the given string and assigns it to the Y field.
func (o *JWKPublicKey) SetY(v string) {
	o.Y = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *JWKPublicKey) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JWKPublicKey) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *JWKPublicKey) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *JWKPublicKey) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetExpiredAt returns the ExpiredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JWKPublicKey) GetExpiredAt() int32 {
	if o == nil || o.ExpiredAt.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ExpiredAt.Get()
}

// GetExpiredAtOk returns a tuple with the ExpiredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JWKPublicKey) GetExpiredAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiredAt.Get(), o.ExpiredAt.IsSet()
}

// HasExpiredAt returns a boolean if a field has been set.
func (o *JWKPublicKey) HasExpiredAt() bool {
	if o != nil && o.ExpiredAt.IsSet() {
		return true
	}

	return false
}

// SetExpiredAt gets a reference to the given NullableInt32 and assigns it to the ExpiredAt field.
func (o *JWKPublicKey) SetExpiredAt(v int32) {
	o.ExpiredAt.Set(&v)
}

// SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil
func (o *JWKPublicKey) SetExpiredAtNil() {
	o.ExpiredAt.Set(nil)
}

// UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
func (o *JWKPublicKey) UnsetExpiredAt() {
	o.ExpiredAt.Unset()
}

func (o JWKPublicKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alg != nil {
		toSerialize["alg"] = o.Alg
	}
	if o.Crv != nil {
		toSerialize["crv"] = o.Crv
	}
	if o.Kid != nil {
		toSerialize["kid"] = o.Kid
	}
	if o.Kty != nil {
		toSerialize["kty"] = o.Kty
	}
	if o.Use != nil {
		toSerialize["use"] = o.Use
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiredAt.IsSet() {
		toSerialize["expired_at"] = o.ExpiredAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *JWKPublicKey) UnmarshalJSON(bytes []byte) (err error) {
	varJWKPublicKey := _JWKPublicKey{}

	if err = json.Unmarshal(bytes, &varJWKPublicKey); err == nil {
		*o = JWKPublicKey(varJWKPublicKey)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "alg")
		delete(additionalProperties, "crv")
		delete(additionalProperties, "kid")
		delete(additionalProperties, "kty")
		delete(additionalProperties, "use")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expired_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJWKPublicKey struct {
	value *JWKPublicKey
	isSet bool
}

func (v NullableJWKPublicKey) Get() *JWKPublicKey {
	return v.value
}

func (v *NullableJWKPublicKey) Set(val *JWKPublicKey) {
	v.value = val
	v.isSet = true
}

func (v NullableJWKPublicKey) IsSet() bool {
	return v.isSet
}

func (v *NullableJWKPublicKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJWKPublicKey(val *JWKPublicKey) *NullableJWKPublicKey {
	return &NullableJWKPublicKey{value: val, isSet: true}
}

func (v NullableJWKPublicKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJWKPublicKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
