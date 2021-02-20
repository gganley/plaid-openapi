# Paystub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaystubId** | **string** | The unique identifier for this paystub. | 
**AccountId** | Pointer to **NullableString** | The account identifier for the account associated with this paystub. | [optional] 
**Employer** | [**Employer**](Employer.md) |  | 
**Employee** | [**Employee**](Employee.md) |  | 
**PayPeriodDetails** | [**PayPeriodDetails**](PayPeriodDetails.md) |  | 
**IncomeBreakdown** | [**IncomeBreakdown**](IncomeBreakdown.md) |  | 
**YtdEarnings** | [**PaystubYTDDetails**](PaystubYTDDetails.md) |  | 

## Methods

### NewPaystub

`func NewPaystub(paystubId string, employer Employer, employee Employee, payPeriodDetails PayPeriodDetails, incomeBreakdown IncomeBreakdown, ytdEarnings PaystubYTDDetails, ) *Paystub`

NewPaystub instantiates a new Paystub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubWithDefaults

`func NewPaystubWithDefaults() *Paystub`

NewPaystubWithDefaults instantiates a new Paystub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaystubId

`func (o *Paystub) GetPaystubId() string`

GetPaystubId returns the PaystubId field if non-nil, zero value otherwise.

### GetPaystubIdOk

`func (o *Paystub) GetPaystubIdOk() (*string, bool)`

GetPaystubIdOk returns a tuple with the PaystubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystubId

`func (o *Paystub) SetPaystubId(v string)`

SetPaystubId sets PaystubId field to given value.


### GetAccountId

`func (o *Paystub) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Paystub) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Paystub) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Paystub) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Paystub) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Paystub) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetEmployer

`func (o *Paystub) GetEmployer() Employer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *Paystub) GetEmployerOk() (*Employer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *Paystub) SetEmployer(v Employer)`

SetEmployer sets Employer field to given value.


### GetEmployee

`func (o *Paystub) GetEmployee() Employee`

GetEmployee returns the Employee field if non-nil, zero value otherwise.

### GetEmployeeOk

`func (o *Paystub) GetEmployeeOk() (*Employee, bool)`

GetEmployeeOk returns a tuple with the Employee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployee

`func (o *Paystub) SetEmployee(v Employee)`

SetEmployee sets Employee field to given value.


### GetPayPeriodDetails

`func (o *Paystub) GetPayPeriodDetails() PayPeriodDetails`

GetPayPeriodDetails returns the PayPeriodDetails field if non-nil, zero value otherwise.

### GetPayPeriodDetailsOk

`func (o *Paystub) GetPayPeriodDetailsOk() (*PayPeriodDetails, bool)`

GetPayPeriodDetailsOk returns a tuple with the PayPeriodDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayPeriodDetails

`func (o *Paystub) SetPayPeriodDetails(v PayPeriodDetails)`

SetPayPeriodDetails sets PayPeriodDetails field to given value.


### GetIncomeBreakdown

`func (o *Paystub) GetIncomeBreakdown() IncomeBreakdown`

GetIncomeBreakdown returns the IncomeBreakdown field if non-nil, zero value otherwise.

### GetIncomeBreakdownOk

`func (o *Paystub) GetIncomeBreakdownOk() (*IncomeBreakdown, bool)`

GetIncomeBreakdownOk returns a tuple with the IncomeBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeBreakdown

`func (o *Paystub) SetIncomeBreakdown(v IncomeBreakdown)`

SetIncomeBreakdown sets IncomeBreakdown field to given value.


### GetYtdEarnings

`func (o *Paystub) GetYtdEarnings() PaystubYTDDetails`

GetYtdEarnings returns the YtdEarnings field if non-nil, zero value otherwise.

### GetYtdEarningsOk

`func (o *Paystub) GetYtdEarningsOk() (*PaystubYTDDetails, bool)`

GetYtdEarningsOk returns a tuple with the YtdEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdEarnings

`func (o *Paystub) SetYtdEarnings(v PaystubYTDDetails)`

SetYtdEarnings sets YtdEarnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


