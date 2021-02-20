# InvestmentTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvestmentTransactionId** | **string** | The ID of the Investment transaction, unique across all Plaid transactions. Like all Plaid identifiers, the &#x60;investment_transaction_id&#x60; is case sensitive. | 
**CancelTransactionId** | Pointer to **NullableString** |  | [optional] 
**AccountId** | **string** | The &#x60;account_id&#x60; of the account against which this transaction posted. | 
**SecurityId** | Pointer to **NullableString** | The &#x60;security_id&#x60; to which this transaction is related. | [optional] 
**Date** | **string** | The ISO-8601 posting date for the transaction, or transacted date for pending transactions. | 
**Name** | **string** | The institution’s description of the transaction. | 
**Quantity** | **float32** | The number of units of the security involved in this transactions | 
**Amount** | **float32** | The complete value of the transaction. Positive values when cash is debited, e.g. purchases of stock; negative values when cash is credited, e.g. sales of stock. Treatment remains the same for cash-only movements unassociated with securities. | 
**Price** | **float32** | The price of the security at which this transaction occurred. | 
**Fees** | Pointer to **NullableFloat32** | The combined value of all fees applied to this transaction | [optional] 
**Type** | **string** | Value is one of the following: &#x60;buy&#x60;: Buying an investment &#x60;sell&#x60;: Selling an investment &#x60;cancel&#x60;: A cancellation of a pending transaction  &#x60;cash&#x60;: Activity that modifies a cash position &#x60;fee&#x60;: A fee on the account &#x60;transfer&#x60;: Activity which modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer | 
**Subtype** | **string** | For descriptions of possible transaction subtypes, see [Investment transaction subtypes schema](/docs/api/accounts/#investment-transaction-subtypes-schema). | 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the transaction. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-&#x60;null&#x60;. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the holding. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 

## Methods

### NewInvestmentTransaction

`func NewInvestmentTransaction(investmentTransactionId string, accountId string, date string, name string, quantity float32, amount float32, price float32, type_ string, subtype string, ) *InvestmentTransaction`

NewInvestmentTransaction instantiates a new InvestmentTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentTransactionWithDefaults

`func NewInvestmentTransactionWithDefaults() *InvestmentTransaction`

NewInvestmentTransactionWithDefaults instantiates a new InvestmentTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvestmentTransactionId

`func (o *InvestmentTransaction) GetInvestmentTransactionId() string`

GetInvestmentTransactionId returns the InvestmentTransactionId field if non-nil, zero value otherwise.

### GetInvestmentTransactionIdOk

`func (o *InvestmentTransaction) GetInvestmentTransactionIdOk() (*string, bool)`

GetInvestmentTransactionIdOk returns a tuple with the InvestmentTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentTransactionId

`func (o *InvestmentTransaction) SetInvestmentTransactionId(v string)`

SetInvestmentTransactionId sets InvestmentTransactionId field to given value.


### GetCancelTransactionId

`func (o *InvestmentTransaction) GetCancelTransactionId() string`

GetCancelTransactionId returns the CancelTransactionId field if non-nil, zero value otherwise.

### GetCancelTransactionIdOk

`func (o *InvestmentTransaction) GetCancelTransactionIdOk() (*string, bool)`

GetCancelTransactionIdOk returns a tuple with the CancelTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTransactionId

`func (o *InvestmentTransaction) SetCancelTransactionId(v string)`

SetCancelTransactionId sets CancelTransactionId field to given value.

### HasCancelTransactionId

`func (o *InvestmentTransaction) HasCancelTransactionId() bool`

HasCancelTransactionId returns a boolean if a field has been set.

### SetCancelTransactionIdNil

`func (o *InvestmentTransaction) SetCancelTransactionIdNil(b bool)`

 SetCancelTransactionIdNil sets the value for CancelTransactionId to be an explicit nil

### UnsetCancelTransactionId
`func (o *InvestmentTransaction) UnsetCancelTransactionId()`

UnsetCancelTransactionId ensures that no value is present for CancelTransactionId, not even an explicit nil
### GetAccountId

`func (o *InvestmentTransaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InvestmentTransaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InvestmentTransaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSecurityId

`func (o *InvestmentTransaction) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *InvestmentTransaction) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *InvestmentTransaction) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *InvestmentTransaction) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### SetSecurityIdNil

`func (o *InvestmentTransaction) SetSecurityIdNil(b bool)`

 SetSecurityIdNil sets the value for SecurityId to be an explicit nil

### UnsetSecurityId
`func (o *InvestmentTransaction) UnsetSecurityId()`

UnsetSecurityId ensures that no value is present for SecurityId, not even an explicit nil
### GetDate

`func (o *InvestmentTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvestmentTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvestmentTransaction) SetDate(v string)`

SetDate sets Date field to given value.


### GetName

`func (o *InvestmentTransaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvestmentTransaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvestmentTransaction) SetName(v string)`

SetName sets Name field to given value.


### GetQuantity

`func (o *InvestmentTransaction) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvestmentTransaction) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvestmentTransaction) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetAmount

`func (o *InvestmentTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvestmentTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvestmentTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetPrice

`func (o *InvestmentTransaction) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvestmentTransaction) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvestmentTransaction) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetFees

`func (o *InvestmentTransaction) GetFees() float32`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *InvestmentTransaction) GetFeesOk() (*float32, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *InvestmentTransaction) SetFees(v float32)`

SetFees sets Fees field to given value.

### HasFees

`func (o *InvestmentTransaction) HasFees() bool`

HasFees returns a boolean if a field has been set.

### SetFeesNil

`func (o *InvestmentTransaction) SetFeesNil(b bool)`

 SetFeesNil sets the value for Fees to be an explicit nil

### UnsetFees
`func (o *InvestmentTransaction) UnsetFees()`

UnsetFees ensures that no value is present for Fees, not even an explicit nil
### GetType

`func (o *InvestmentTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvestmentTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvestmentTransaction) SetType(v string)`

SetType sets Type field to given value.


### GetSubtype

`func (o *InvestmentTransaction) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *InvestmentTransaction) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *InvestmentTransaction) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.


### GetIsoCurrencyCode

`func (o *InvestmentTransaction) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *InvestmentTransaction) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *InvestmentTransaction) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *InvestmentTransaction) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *InvestmentTransaction) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *InvestmentTransaction) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *InvestmentTransaction) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *InvestmentTransaction) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *InvestmentTransaction) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *InvestmentTransaction) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *InvestmentTransaction) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *InvestmentTransaction) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


