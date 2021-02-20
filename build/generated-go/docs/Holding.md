# Holding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid &#x60;account_id&#x60; associated with the holding. | 
**SecurityId** | **string** | The Plaid &#x60;security_id&#x60; associated with the holding. | 
**InstitutionPrice** | **float32** | The last price given by the institution for this security. | 
**InstitutionPriceAsOf** | Pointer to **NullableString** | The date at which &#x60;institution_price&#x60; was current. | [optional] 
**InstitutionValue** | **float32** | The value of the holding, as reported by the institution. | 
**CostBasis** | Pointer to **NullableFloat32** | The cost basis of the holding. | [optional] 
**Quantity** | **float32** | The total quantity of the asset held, as reported by the financial institution. | 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the holding. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-&#x60;null&#x60;. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the holding. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s.  | [optional] 

## Methods

### NewHolding

`func NewHolding(accountId string, securityId string, institutionPrice float32, institutionValue float32, quantity float32, ) *Holding`

NewHolding instantiates a new Holding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHoldingWithDefaults

`func NewHoldingWithDefaults() *Holding`

NewHoldingWithDefaults instantiates a new Holding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Holding) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Holding) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Holding) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSecurityId

`func (o *Holding) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *Holding) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *Holding) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.


### GetInstitutionPrice

`func (o *Holding) GetInstitutionPrice() float32`

GetInstitutionPrice returns the InstitutionPrice field if non-nil, zero value otherwise.

### GetInstitutionPriceOk

`func (o *Holding) GetInstitutionPriceOk() (*float32, bool)`

GetInstitutionPriceOk returns a tuple with the InstitutionPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionPrice

`func (o *Holding) SetInstitutionPrice(v float32)`

SetInstitutionPrice sets InstitutionPrice field to given value.


### GetInstitutionPriceAsOf

`func (o *Holding) GetInstitutionPriceAsOf() string`

GetInstitutionPriceAsOf returns the InstitutionPriceAsOf field if non-nil, zero value otherwise.

### GetInstitutionPriceAsOfOk

`func (o *Holding) GetInstitutionPriceAsOfOk() (*string, bool)`

GetInstitutionPriceAsOfOk returns a tuple with the InstitutionPriceAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionPriceAsOf

`func (o *Holding) SetInstitutionPriceAsOf(v string)`

SetInstitutionPriceAsOf sets InstitutionPriceAsOf field to given value.

### HasInstitutionPriceAsOf

`func (o *Holding) HasInstitutionPriceAsOf() bool`

HasInstitutionPriceAsOf returns a boolean if a field has been set.

### SetInstitutionPriceAsOfNil

`func (o *Holding) SetInstitutionPriceAsOfNil(b bool)`

 SetInstitutionPriceAsOfNil sets the value for InstitutionPriceAsOf to be an explicit nil

### UnsetInstitutionPriceAsOf
`func (o *Holding) UnsetInstitutionPriceAsOf()`

UnsetInstitutionPriceAsOf ensures that no value is present for InstitutionPriceAsOf, not even an explicit nil
### GetInstitutionValue

`func (o *Holding) GetInstitutionValue() float32`

GetInstitutionValue returns the InstitutionValue field if non-nil, zero value otherwise.

### GetInstitutionValueOk

`func (o *Holding) GetInstitutionValueOk() (*float32, bool)`

GetInstitutionValueOk returns a tuple with the InstitutionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionValue

`func (o *Holding) SetInstitutionValue(v float32)`

SetInstitutionValue sets InstitutionValue field to given value.


### GetCostBasis

`func (o *Holding) GetCostBasis() float32`

GetCostBasis returns the CostBasis field if non-nil, zero value otherwise.

### GetCostBasisOk

`func (o *Holding) GetCostBasisOk() (*float32, bool)`

GetCostBasisOk returns a tuple with the CostBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasis

`func (o *Holding) SetCostBasis(v float32)`

SetCostBasis sets CostBasis field to given value.

### HasCostBasis

`func (o *Holding) HasCostBasis() bool`

HasCostBasis returns a boolean if a field has been set.

### SetCostBasisNil

`func (o *Holding) SetCostBasisNil(b bool)`

 SetCostBasisNil sets the value for CostBasis to be an explicit nil

### UnsetCostBasis
`func (o *Holding) UnsetCostBasis()`

UnsetCostBasis ensures that no value is present for CostBasis, not even an explicit nil
### GetQuantity

`func (o *Holding) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Holding) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Holding) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetIsoCurrencyCode

`func (o *Holding) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *Holding) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *Holding) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *Holding) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *Holding) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *Holding) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *Holding) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *Holding) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *Holding) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *Holding) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *Holding) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *Holding) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


