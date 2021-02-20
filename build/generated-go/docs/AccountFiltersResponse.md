# AccountFiltersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depository** | Pointer to [**DepositoryFilter**](DepositoryFilter.md) |  | [optional] 
**Credit** | Pointer to [**CreditFilter**](CreditFilter.md) |  | [optional] 
**Loan** | Pointer to [**LoanFilter**](LoanFilter.md) |  | [optional] 
**Investment** | Pointer to [**InvestmentFilter**](InvestmentFilter.md) |  | [optional] 

## Methods

### NewAccountFiltersResponse

`func NewAccountFiltersResponse() *AccountFiltersResponse`

NewAccountFiltersResponse instantiates a new AccountFiltersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFiltersResponseWithDefaults

`func NewAccountFiltersResponseWithDefaults() *AccountFiltersResponse`

NewAccountFiltersResponseWithDefaults instantiates a new AccountFiltersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepository

`func (o *AccountFiltersResponse) GetDepository() DepositoryFilter`

GetDepository returns the Depository field if non-nil, zero value otherwise.

### GetDepositoryOk

`func (o *AccountFiltersResponse) GetDepositoryOk() (*DepositoryFilter, bool)`

GetDepositoryOk returns a tuple with the Depository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepository

`func (o *AccountFiltersResponse) SetDepository(v DepositoryFilter)`

SetDepository sets Depository field to given value.

### HasDepository

`func (o *AccountFiltersResponse) HasDepository() bool`

HasDepository returns a boolean if a field has been set.

### GetCredit

`func (o *AccountFiltersResponse) GetCredit() CreditFilter`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *AccountFiltersResponse) GetCreditOk() (*CreditFilter, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *AccountFiltersResponse) SetCredit(v CreditFilter)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *AccountFiltersResponse) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetLoan

`func (o *AccountFiltersResponse) GetLoan() LoanFilter`

GetLoan returns the Loan field if non-nil, zero value otherwise.

### GetLoanOk

`func (o *AccountFiltersResponse) GetLoanOk() (*LoanFilter, bool)`

GetLoanOk returns a tuple with the Loan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoan

`func (o *AccountFiltersResponse) SetLoan(v LoanFilter)`

SetLoan sets Loan field to given value.

### HasLoan

`func (o *AccountFiltersResponse) HasLoan() bool`

HasLoan returns a boolean if a field has been set.

### GetInvestment

`func (o *AccountFiltersResponse) GetInvestment() InvestmentFilter`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *AccountFiltersResponse) GetInvestmentOk() (*InvestmentFilter, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *AccountFiltersResponse) SetInvestment(v InvestmentFilter)`

SetInvestment sets Investment field to given value.

### HasInvestment

`func (o *AccountFiltersResponse) HasInvestment() bool`

HasInvestment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


