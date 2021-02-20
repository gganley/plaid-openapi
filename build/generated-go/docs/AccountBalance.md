# AccountBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **NullableFloat32** | The amount of funds available to be withdrawn from the account, as determined by the financial institution.  For &#x60;credit&#x60;-type accounts, the &#x60;available&#x60; balance typically equals the &#x60;limit&#x60; less the &#x60;current&#x60; balance, less any pending outflows plus any pending inflows.  For &#x60;depository&#x60;-type accounts, the &#x60;available&#x60; balance typically equals the &#x60;current&#x60; balance less any pending outflows plus any pending inflows. For &#x60;depository&#x60;-type accounts, the &#x60;available&#x60; balance does not include the overdraft limit.  For &#x60;investment&#x60;-type accounts, the &#x60;available&#x60; balance is the total cash available to withdraw as presented by the institution.  Note that not all institutions calculate the &#x60;available&#x60;  balance. In the event that &#x60;available&#x60; balance is unavailable, Plaid will return an &#x60;available&#x60; balance value of &#x60;null&#x60;.  Available balance may be cached and is not guaranteed to be up-to-date in realtime unless the value was returned by &#x60;/accounts/balance/get&#x60;. | [optional] 
**Current** | **float32** | The total amount of funds in or owed by the account.  For &#x60;credit&#x60;-type accounts, a positive balance indicates the amount owed; a negative amount indicates the lender owing the account holder.  For &#x60;loan&#x60;-type accounts, the current balance is the principal remaining on the loan, except in the case of student loan accounts at Sallie Mae (&#x60;ins_116944&#x60;). For Sallie Mae student loans, the account&#39;s balance includes both principal and any outstanding interest.  For &#x60;investment&#x60;-type accounts, the current balance is the total value of assets as presented by the institution.  Note that balance information may be cached unless the value was returned by &#x60;/accounts/balance/get&#x60;, and current balance information is typically not updated intra-day. If you require realtime balance information, use the &#x60;available&#x60; balance as provided by &#x60;/accounts/balance/get&#x60;. | 
**Limit** | Pointer to **NullableFloat32** | For &#x60;credit&#x60;-type accounts, this represents the credit limit.  For &#x60;depository&#x60;-type accounts, this represents the pre-arranged overdraft limit, which is common for current (checking) accounts in Europe.  In North America, this field is typically only available for &#x60;credit&#x60;-type accounts. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the balance. Always null if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the balance. Always null if &#x60;iso_currency_code&#x60; is non-null. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;unofficial_currency_code&#x60;s. | [optional] 

## Methods

### NewAccountBalance

`func NewAccountBalance(current float32, ) *AccountBalance`

NewAccountBalance instantiates a new AccountBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBalanceWithDefaults

`func NewAccountBalanceWithDefaults() *AccountBalance`

NewAccountBalanceWithDefaults instantiates a new AccountBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AccountBalance) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountBalance) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountBalance) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AccountBalance) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *AccountBalance) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *AccountBalance) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCurrent

`func (o *AccountBalance) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *AccountBalance) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *AccountBalance) SetCurrent(v float32)`

SetCurrent sets Current field to given value.


### GetLimit

`func (o *AccountBalance) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *AccountBalance) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *AccountBalance) SetLimit(v float32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *AccountBalance) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimitNil

`func (o *AccountBalance) SetLimitNil(b bool)`

 SetLimitNil sets the value for Limit to be an explicit nil

### UnsetLimit
`func (o *AccountBalance) UnsetLimit()`

UnsetLimit ensures that no value is present for Limit, not even an explicit nil
### GetIsoCurrencyCode

`func (o *AccountBalance) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *AccountBalance) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *AccountBalance) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *AccountBalance) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *AccountBalance) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *AccountBalance) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *AccountBalance) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *AccountBalance) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *AccountBalance) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *AccountBalance) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *AccountBalance) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *AccountBalance) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


