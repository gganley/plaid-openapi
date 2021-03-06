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

// SandboxItemFireWebhookRequest SandboxItemFireWebhookRequest defines the request schema for `/sandbox/item/fire_webhook`
type SandboxItemFireWebhookRequest struct {
	// Your Plaid API `client_id`.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	// The following values for `webhook_code` are supported:  * `DEFAULT_UPDATE`
	WebhookCode          *string `json:"webhook_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SandboxItemFireWebhookRequest SandboxItemFireWebhookRequest

// NewSandboxItemFireWebhookRequest instantiates a new SandboxItemFireWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxItemFireWebhookRequest(accessToken string) *SandboxItemFireWebhookRequest {
	this := SandboxItemFireWebhookRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewSandboxItemFireWebhookRequestWithDefaults instantiates a new SandboxItemFireWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxItemFireWebhookRequestWithDefaults() *SandboxItemFireWebhookRequest {
	this := SandboxItemFireWebhookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxItemFireWebhookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxItemFireWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *SandboxItemFireWebhookRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *SandboxItemFireWebhookRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetWebhookCode returns the WebhookCode field value if set, zero value otherwise.
func (o *SandboxItemFireWebhookRequest) GetWebhookCode() string {
	if o == nil || o.WebhookCode == nil {
		var ret string
		return ret
	}
	return *o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxItemFireWebhookRequest) GetWebhookCodeOk() (*string, bool) {
	if o == nil || o.WebhookCode == nil {
		return nil, false
	}
	return o.WebhookCode, true
}

// HasWebhookCode returns a boolean if a field has been set.
func (o *SandboxItemFireWebhookRequest) HasWebhookCode() bool {
	if o != nil && o.WebhookCode != nil {
		return true
	}

	return false
}

// SetWebhookCode gets a reference to the given string and assigns it to the WebhookCode field.
func (o *SandboxItemFireWebhookRequest) SetWebhookCode(v string) {
	o.WebhookCode = &v
}

func (o SandboxItemFireWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.WebhookCode != nil {
		toSerialize["webhook_code"] = o.WebhookCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxItemFireWebhookRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxItemFireWebhookRequest := _SandboxItemFireWebhookRequest{}

	if err = json.Unmarshal(bytes, &varSandboxItemFireWebhookRequest); err == nil {
		*o = SandboxItemFireWebhookRequest(varSandboxItemFireWebhookRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_id")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "webhook_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxItemFireWebhookRequest struct {
	value *SandboxItemFireWebhookRequest
	isSet bool
}

func (v NullableSandboxItemFireWebhookRequest) Get() *SandboxItemFireWebhookRequest {
	return v.value
}

func (v *NullableSandboxItemFireWebhookRequest) Set(val *SandboxItemFireWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxItemFireWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxItemFireWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxItemFireWebhookRequest(val *SandboxItemFireWebhookRequest) *NullableSandboxItemFireWebhookRequest {
	return &NullableSandboxItemFireWebhookRequest{value: val, isSet: true}
}

func (v NullableSandboxItemFireWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxItemFireWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
