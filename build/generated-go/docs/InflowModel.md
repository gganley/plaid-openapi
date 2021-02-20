# InflowModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Inflow model. One of the following:  &#x60;none&#x60;: No income &#x60;monthly-income&#x60;: Income occurs once per month &#x60;monthly-balance-payment&#x60;: Pays off the balance on a liability account at the given statement day of month &#x60;monthly-interest-only-payment&#x60;: Makes an interest-only payment on a liability account at the given statement day of month. Note that account types supported by Liabilities will accrue interest in the Sandbox. The types are account type &#x60;credit&#x60; with subtype &#x60;credit&#x60; or &#x60;paypal&#x60;, and account type &#x60;loan&#x60; with subtype &#x60;student&#x60; or &#x60;mortgage&#x60;. | 
**IncomeAmount** | **float32** | Amount of income per month. This value is required if &#x60;type&#x60; is &#x60;monthly-income&#x60;. | 
**PaymentDayOfMonth** | **float32** | Number between 1 and 28, or &#x60;last&#x60; meaning the last day of the month. The day of the month on which the income transaction will appear. The name of the income transaction. This field is required if &#x60;type&#x60; is &#x60;monthly-income&#x60;, &#x60;monthly-balance-payment&#x60; or &#x60;monthly-interest-only-payment&#x60;. | 
**TransactionName** | **string** | The name of the income transaction. This field is required if &#x60;type&#x60; is &#x60;monthly-income&#x60;, &#x60;monthly-balance-payment&#x60; or &#x60;monthly-interest-only-payment&#x60;. | 
**StatementDayOfMonth** | **string** | Number between 1 and 28, or &#x60;last&#x60; meaning the last day of the month. The day of the month on which the balance is calculated for the next payment. The name of the income transaction. This field is required if &#x60;type&#x60; is &#x60;monthly-balance-payment&#x60; or &#x60;monthly-interest-only-payment&#x60;. | 

## Methods

### NewInflowModel

`func NewInflowModel(type_ string, incomeAmount float32, paymentDayOfMonth float32, transactionName string, statementDayOfMonth string, ) *InflowModel`

NewInflowModel instantiates a new InflowModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInflowModelWithDefaults

`func NewInflowModelWithDefaults() *InflowModel`

NewInflowModelWithDefaults instantiates a new InflowModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InflowModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InflowModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InflowModel) SetType(v string)`

SetType sets Type field to given value.


### GetIncomeAmount

`func (o *InflowModel) GetIncomeAmount() float32`

GetIncomeAmount returns the IncomeAmount field if non-nil, zero value otherwise.

### GetIncomeAmountOk

`func (o *InflowModel) GetIncomeAmountOk() (*float32, bool)`

GetIncomeAmountOk returns a tuple with the IncomeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeAmount

`func (o *InflowModel) SetIncomeAmount(v float32)`

SetIncomeAmount sets IncomeAmount field to given value.


### GetPaymentDayOfMonth

`func (o *InflowModel) GetPaymentDayOfMonth() float32`

GetPaymentDayOfMonth returns the PaymentDayOfMonth field if non-nil, zero value otherwise.

### GetPaymentDayOfMonthOk

`func (o *InflowModel) GetPaymentDayOfMonthOk() (*float32, bool)`

GetPaymentDayOfMonthOk returns a tuple with the PaymentDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDayOfMonth

`func (o *InflowModel) SetPaymentDayOfMonth(v float32)`

SetPaymentDayOfMonth sets PaymentDayOfMonth field to given value.


### GetTransactionName

`func (o *InflowModel) GetTransactionName() string`

GetTransactionName returns the TransactionName field if non-nil, zero value otherwise.

### GetTransactionNameOk

`func (o *InflowModel) GetTransactionNameOk() (*string, bool)`

GetTransactionNameOk returns a tuple with the TransactionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionName

`func (o *InflowModel) SetTransactionName(v string)`

SetTransactionName sets TransactionName field to given value.


### GetStatementDayOfMonth

`func (o *InflowModel) GetStatementDayOfMonth() string`

GetStatementDayOfMonth returns the StatementDayOfMonth field if non-nil, zero value otherwise.

### GetStatementDayOfMonthOk

`func (o *InflowModel) GetStatementDayOfMonthOk() (*string, bool)`

GetStatementDayOfMonthOk returns a tuple with the StatementDayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementDayOfMonth

`func (o *InflowModel) SetStatementDayOfMonth(v string)`

SetStatementDayOfMonth sets StatementDayOfMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


