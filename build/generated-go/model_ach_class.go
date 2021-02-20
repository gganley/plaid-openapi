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
	"fmt"
)

// ACHClass Specifies the use case of the transfer.  Required for transfers on an ACH network.  `\"arc\"` - Accounts Receivable Entry  `\"cbr`\" - Cross Border Entry  `\"ccd\"` - Corporate Credit or Debit - fund transfer between two corporate bank accounts  `\"cie\"` - Customer Initiated Entry  `\"cor\"` - Automated Notification of Change  `\"ctx\"` - Corporate Trade Exchange  `\"iat\"` - International  `\"mte\"` - Machine Transfer Entry  `\"pbr\"` - Cross Border Entry  `\"pop`\" - Point-of-Purchase Entry  `\"pos\"` - Point-of-Sale Entry  `\"ppd\"` - Prearranged Payment or Deposit - the transfer is part of a pre-existing relationship with a consumer, eg. bill payment  `\"rck\"` - Re-presented Check Entry  `\"tel\"` - Telephone-Initiated Entry  `\"web\"` - Internet-Initiated Entry - debits from a consumer’s account where their authorization is obtained over the Internet
type ACHClass string

// List of ACHClass
const (
	ACHCLASS_ARC ACHClass = "arc"
	ACHCLASS_CBR ACHClass = "cbr"
	ACHCLASS_CCD ACHClass = "ccd"
	ACHCLASS_CIE ACHClass = "cie"
	ACHCLASS_COR ACHClass = "cor"
	ACHCLASS_CTX ACHClass = "ctx"
	ACHCLASS_IAT ACHClass = "iat"
	ACHCLASS_MTE ACHClass = "mte"
	ACHCLASS_PBR ACHClass = "pbr"
	ACHCLASS_POP ACHClass = "pop"
	ACHCLASS_POS ACHClass = "pos"
	ACHCLASS_PPD ACHClass = "ppd"
	ACHCLASS_RCK ACHClass = "rck"
	ACHCLASS_TEL ACHClass = "tel"
	ACHCLASS_WEB ACHClass = "web"
)

func (v *ACHClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ACHClass(value)
	for _, existing := range []ACHClass{"arc", "cbr", "ccd", "cie", "cor", "ctx", "iat", "mte", "pbr", "pop", "pos", "ppd", "rck", "tel", "web"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ACHClass", value)
}

// Ptr returns reference to ACHClass value
func (v ACHClass) Ptr() *ACHClass {
	return &v
}

type NullableACHClass struct {
	value *ACHClass
	isSet bool
}

func (v NullableACHClass) Get() *ACHClass {
	return v.value
}

func (v *NullableACHClass) Set(val *ACHClass) {
	v.value = val
	v.isSet = true
}

func (v NullableACHClass) IsSet() bool {
	return v.isSet
}

func (v *NullableACHClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACHClass(val *ACHClass) *NullableACHClass {
	return &NullableACHClass{value: val, isSet: true}
}

func (v NullableACHClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACHClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
