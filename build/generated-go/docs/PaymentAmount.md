# PaymentAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | **string** | The ISO-4217 currency code of the payment amount | 
**Value** | **float32** | The payment amount | 

## Methods

### NewPaymentAmount

`func NewPaymentAmount(currency string, value float32, ) *PaymentAmount`

NewPaymentAmount instantiates a new PaymentAmount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentAmountWithDefaults

`func NewPaymentAmountWithDefaults() *PaymentAmount`

NewPaymentAmountWithDefaults instantiates a new PaymentAmount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *PaymentAmount) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *PaymentAmount) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *PaymentAmount) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetValue

`func (o *PaymentAmount) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaymentAmount) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaymentAmount) SetValue(v float32)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


