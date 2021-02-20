# HistoricalBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | The date of the calculated historical balance, in an ISO 8601 format (YYYY-MM-DD) | 
**Current** | **float32** | The total amount of funds in the account, calculated from the &#x60;current&#x60; balance in the &#x60;balance&#x60; object by subtracting inflows and adding back outflows according to the posted date of each transaction.  If the account has any pending transactions, historical balance amounts on or after the date of the earliest pending transaction may differ if retrieved in subsequent Asset Reports as a result of those pending transactions posting. | 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the balance. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-&#x60;null&#x60;. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the balance. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 

## Methods

### NewHistoricalBalance

`func NewHistoricalBalance(date string, current float32, ) *HistoricalBalance`

NewHistoricalBalance instantiates a new HistoricalBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalBalanceWithDefaults

`func NewHistoricalBalanceWithDefaults() *HistoricalBalance`

NewHistoricalBalanceWithDefaults instantiates a new HistoricalBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *HistoricalBalance) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *HistoricalBalance) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *HistoricalBalance) SetDate(v string)`

SetDate sets Date field to given value.


### GetCurrent

`func (o *HistoricalBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *HistoricalBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *HistoricalBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.


### GetIsoCurrencyCode

`func (o *HistoricalBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *HistoricalBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *HistoricalBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *HistoricalBalance) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *HistoricalBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *HistoricalBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *HistoricalBalance) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *HistoricalBalance) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *HistoricalBalance) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *HistoricalBalance) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *HistoricalBalance) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *HistoricalBalance) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


