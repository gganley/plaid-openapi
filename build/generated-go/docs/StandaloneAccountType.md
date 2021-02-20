# StandaloneAccountType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depository** | **string** | An account type holding cash, in which funds are deposited. Supported products for &#x60;depository&#x60; accounts are: Auth, Balance, Transactions, Identity, Payment Initiation, and Assets. | 
**Credit** | **string** | A credit card type account. Supported products for &#x60;credit&#x60; accounts are: Balance, Transactions, Identity, and Liabilities. | 
**Loan** | **string** | A loan type account. Supported products for &#x60;loan&#x60; accounts are: Balance, Liabilities, and Transactions. | 
**Investment** | **string** | An investment account. Supported products for &#x60;investment&#x60; accounts are: Balance and Investments. | 
**Other** | **string** | Other or unknown account type. Supported products for &#x60;other&#x60; accounts are: Balance, Transactions, Identity, and Assets. | 

## Methods

### NewStandaloneAccountType

`func NewStandaloneAccountType(depository string, credit string, loan string, investment string, other string, ) *StandaloneAccountType`

NewStandaloneAccountType instantiates a new StandaloneAccountType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneAccountTypeWithDefaults

`func NewStandaloneAccountTypeWithDefaults() *StandaloneAccountType`

NewStandaloneAccountTypeWithDefaults instantiates a new StandaloneAccountType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepository

`func (o *StandaloneAccountType) GetDepository() string`

GetDepository returns the Depository field if non-nil, zero value otherwise.

### GetDepositoryOk

`func (o *StandaloneAccountType) GetDepositoryOk() (*string, bool)`

GetDepositoryOk returns a tuple with the Depository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepository

`func (o *StandaloneAccountType) SetDepository(v string)`

SetDepository sets Depository field to given value.


### GetCredit

`func (o *StandaloneAccountType) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *StandaloneAccountType) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *StandaloneAccountType) SetCredit(v string)`

SetCredit sets Credit field to given value.


### GetLoan

`func (o *StandaloneAccountType) GetLoan() string`

GetLoan returns the Loan field if non-nil, zero value otherwise.

### GetLoanOk

`func (o *StandaloneAccountType) GetLoanOk() (*string, bool)`

GetLoanOk returns a tuple with the Loan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoan

`func (o *StandaloneAccountType) SetLoan(v string)`

SetLoan sets Loan field to given value.


### GetInvestment

`func (o *StandaloneAccountType) GetInvestment() string`

GetInvestment returns the Investment field if non-nil, zero value otherwise.

### GetInvestmentOk

`func (o *StandaloneAccountType) GetInvestmentOk() (*string, bool)`

GetInvestmentOk returns a tuple with the Investment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestment

`func (o *StandaloneAccountType) SetInvestment(v string)`

SetInvestment sets Investment field to given value.


### GetOther

`func (o *StandaloneAccountType) GetOther() string`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *StandaloneAccountType) GetOtherOk() (*string, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *StandaloneAccountType) SetOther(v string)`

SetOther sets Other field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


