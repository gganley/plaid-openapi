# Amount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The ISO-4217 currency code of the payment. For standing orders, &#x60;\&quot;GBP\&quot;&#x60; must be used. | 
**Value** | **float32** | The amount of the payment. Must contain at most two digits of precision e.g. &#x60;1.23&#x60;. Minimum accepted value is &#x60;1&#x60;. | 

## Methods

### NewAmount

`func NewAmount(currency string, value float32, ) *Amount`

NewAmount instantiates a new Amount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmountWithDefaults

`func NewAmountWithDefaults() *Amount`

NewAmountWithDefaults instantiates a new Amount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *Amount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Amount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Amount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetValue

`func (o *Amount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Amount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Amount) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


