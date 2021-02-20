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

// TransactionCode An identifier classifying the transaction type.  This field is only populated for European institutions. For institutions in the US and Canada, this field is set to `null`.  `adjustment:` Bank adjustment  `atm:` Cash deposit or withdrawal via an automated teller machine  `bank charge:` Charge or fee levied by the institution  `bill payment`: Payment of a bill  `cash:` Cash deposit or withdrawal  `cashback:` Cash withdrawal while making a debit card purchase  `cheque:` Document ordering the payment of money to another person or organization  `direct debit:` Automatic withdrawal of funds initiated by a third party at a regular interval  `interest:` Interest earned or incurred  `purchase:` Purchase made with a debit or credit card  `standing order:` Payment instructed by the account holder to a third party at a regular interval  `transfer:` Transfer of money between accounts
type TransactionCode string

// List of TransactionCode
const (
	TRANSACTIONCODE_ADJUSTMENT     TransactionCode = "adjustment"
	TRANSACTIONCODE_ATM            TransactionCode = "atm"
	TRANSACTIONCODE_BANK_CHARGE    TransactionCode = "bank charge"
	TRANSACTIONCODE_BILL_PAYMENT   TransactionCode = "bill payment"
	TRANSACTIONCODE_CASH           TransactionCode = "cash"
	TRANSACTIONCODE_CASHBACK       TransactionCode = "cashback"
	TRANSACTIONCODE_CHEQUE         TransactionCode = "cheque"
	TRANSACTIONCODE_DIRECT_DEBIT   TransactionCode = "direct debit"
	TRANSACTIONCODE_INTEREST       TransactionCode = "interest"
	TRANSACTIONCODE_PURCHASE       TransactionCode = "purchase"
	TRANSACTIONCODE_STANDING_ORDER TransactionCode = "standing order"
	TRANSACTIONCODE_TRANSFER       TransactionCode = "transfer"
	TRANSACTIONCODE_NULL           TransactionCode = "null"
)

func (v *TransactionCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TransactionCode(value)
	for _, existing := range []TransactionCode{"adjustment", "atm", "bank charge", "bill payment", "cash", "cashback", "cheque", "direct debit", "interest", "purchase", "standing order", "transfer", "null"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TransactionCode", value)
}

// Ptr returns reference to TransactionCode value
func (v TransactionCode) Ptr() *TransactionCode {
	return &v
}

type NullableTransactionCode struct {
	value *TransactionCode
	isSet bool
}

func (v NullableTransactionCode) Get() *TransactionCode {
	return v.value
}

func (v *NullableTransactionCode) Set(val *TransactionCode) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCode) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCode(val *TransactionCode) *NullableTransactionCode {
	return &NullableTransactionCode{value: val, isSet: true}
}

func (v NullableTransactionCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
