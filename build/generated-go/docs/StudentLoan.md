# StudentLoan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | The ID of the account that this liability belongs to. | [optional] 
**AccountNumber** | Pointer to **NullableString** | The account number of the loan. | [optional] 
**DisbursementDates** | Pointer to **[]string** | The dates on which loaned funds were disbursed or will be disbursed. These are often in the past. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**ExpectedPayoffDate** | Pointer to **NullableString** | The date when the student loan is expected to be paid off. Availability for this field is limited. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**Guarantor** | Pointer to **NullableString** | The guarantor of the student loan. | [optional] 
**InterestRatePercentage** | **float32** | The interest rate on the loan as a percentage. | 
**IsOverdue** | Pointer to **NullableBool** | &#x60;true&#x60; if a payment is currently overdue. Availability for this field is limited. | [optional] 
**LastPaymentAmount** | Pointer to **NullableFloat32** | The amount of the last payment. | [optional] 
**LastPaymentDate** | Pointer to **NullableString** | The date of the last payment. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**LastStatementBalance** | Pointer to **NullableFloat32** | The outstanding balance on the last statement. This field could also be interpreted as the next payment due. Availability for this field is limited.  | [optional] 
**LastStatementIssueDate** | Pointer to **NullableString** | The date of the last statement. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**LoanName** | Pointer to **NullableString** | The type of loan, e.g., \&quot;Consolidation Loans\&quot;. | [optional] 
**LoanStatus** | Pointer to [**StudentLoanStatus**](StudentLoanStatus.md) |  | [optional] 
**MinimumPaymentAmount** | Pointer to **NullableFloat32** | The minimum payment due for the next billing cycle. There are some exceptions: Some institutions require a minimum payment across all loans associated with an account number. Our API presents that same minimum payment amount on each loan. The institutions that do this are: Great Lakes ( &#x60;ins_116861&#x60;), Firstmark (&#x60;ins_116295&#x60;), Commonbond Firstmark Services (&#x60;ins_116950&#x60;), Nelnet (&#x60;ins_116528&#x60;), EdFinancial Services (&#x60;ins_116304&#x60;), Granite State (&#x60;ins_116308&#x60;), and Oklahoma Student Loan Authority (&#x60;ins_116945&#x60;). Firstmark (&#x60;ins_116295&#x60; ) will display as $0 if there is an autopay program in effect. | [optional] 
**NextPaymentDueDate** | Pointer to **NullableString** | The due date for the next payment. The due date is &#x60;null&#x60; if a payment is not expected. A payment is not expected if &#x60;loan_status.type&#x60; is &#x60;deferment&#x60;, &#x60;in_school&#x60;, &#x60;consolidated&#x60;, &#x60;paid in full&#x60;, or &#x60;transferred&#x60;. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | [optional] 
**OriginationDate** | Pointer to **NullableString** | The date on which the loan was initially lent. Dates are returned in an ISO 8601 format (YYYY-MM-DD).  | [optional] 
**OriginationPrincipalAmount** | Pointer to **NullableFloat32** | The original principal balance of the loan. | [optional] 
**OutstandingInterestAmount** | Pointer to **NullableFloat32** | The total dollar amount of the accrued interest balance. For Sallie Mae ( &#x60;ins_116944&#x60;), this amount is included in the current balance of the loan, so this field will return as &#x60;null&#x60;. | [optional] 
**PaymentReferenceNumber** | Pointer to **NullableString** | The relevant account number that should be used to reference this loan for payments. In the majority of cases, &#x60;payment_reference_number&#x60; will match a&#x60;ccount_number,&#x60; but in some institutions, such as Great Lakes (&#x60;ins_116861&#x60;), it will be different. | [optional] 
**PslfStatus** | Pointer to [**PSLFStatus**](PSLFStatus.md) |  | [optional] 
**RepaymentPlan** | Pointer to [**StudentRepaymentPlan**](StudentRepaymentPlan.md) |  | [optional] 
**SequenceNumber** | Pointer to **NullableString** | The sequence number of the student loan. Heartland ECSI (&#x60;ins_116948&#x60;) does not make this field available. | [optional] 
**ServicerAddress** | Pointer to [**ServicerAddressData**](ServicerAddressData.md) |  | [optional] 
**YtdInterestPaid** | Pointer to **NullableFloat32** | The year to date (YTD) interest paid. Availability for this field is limited. | [optional] 
**YtdPrincipalPaid** | Pointer to **NullableFloat32** | The year to date (YTD) principal paid. Availability for this field is limited. | [optional] 

## Methods

### NewStudentLoan

`func NewStudentLoan(interestRatePercentage float32, ) *StudentLoan`

NewStudentLoan instantiates a new StudentLoan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentLoanWithDefaults

`func NewStudentLoanWithDefaults() *StudentLoan`

NewStudentLoanWithDefaults instantiates a new StudentLoan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *StudentLoan) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StudentLoan) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StudentLoan) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StudentLoan) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *StudentLoan) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *StudentLoan) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountNumber

`func (o *StudentLoan) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *StudentLoan) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *StudentLoan) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *StudentLoan) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *StudentLoan) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *StudentLoan) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetDisbursementDates

`func (o *StudentLoan) GetDisbursementDates() []string`

GetDisbursementDates returns the DisbursementDates field if non-nil, zero value otherwise.

### GetDisbursementDatesOk

`func (o *StudentLoan) GetDisbursementDatesOk() (*[]string, bool)`

GetDisbursementDatesOk returns a tuple with the DisbursementDates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisbursementDates

`func (o *StudentLoan) SetDisbursementDates(v []string)`

SetDisbursementDates sets DisbursementDates field to given value.

### HasDisbursementDates

`func (o *StudentLoan) HasDisbursementDates() bool`

HasDisbursementDates returns a boolean if a field has been set.

### SetDisbursementDatesNil

`func (o *StudentLoan) SetDisbursementDatesNil(b bool)`

 SetDisbursementDatesNil sets the value for DisbursementDates to be an explicit nil

### UnsetDisbursementDates
`func (o *StudentLoan) UnsetDisbursementDates()`

UnsetDisbursementDates ensures that no value is present for DisbursementDates, not even an explicit nil
### GetExpectedPayoffDate

`func (o *StudentLoan) GetExpectedPayoffDate() string`

GetExpectedPayoffDate returns the ExpectedPayoffDate field if non-nil, zero value otherwise.

### GetExpectedPayoffDateOk

`func (o *StudentLoan) GetExpectedPayoffDateOk() (*string, bool)`

GetExpectedPayoffDateOk returns a tuple with the ExpectedPayoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedPayoffDate

`func (o *StudentLoan) SetExpectedPayoffDate(v string)`

SetExpectedPayoffDate sets ExpectedPayoffDate field to given value.

### HasExpectedPayoffDate

`func (o *StudentLoan) HasExpectedPayoffDate() bool`

HasExpectedPayoffDate returns a boolean if a field has been set.

### SetExpectedPayoffDateNil

`func (o *StudentLoan) SetExpectedPayoffDateNil(b bool)`

 SetExpectedPayoffDateNil sets the value for ExpectedPayoffDate to be an explicit nil

### UnsetExpectedPayoffDate
`func (o *StudentLoan) UnsetExpectedPayoffDate()`

UnsetExpectedPayoffDate ensures that no value is present for ExpectedPayoffDate, not even an explicit nil
### GetGuarantor

`func (o *StudentLoan) GetGuarantor() string`

GetGuarantor returns the Guarantor field if non-nil, zero value otherwise.

### GetGuarantorOk

`func (o *StudentLoan) GetGuarantorOk() (*string, bool)`

GetGuarantorOk returns a tuple with the Guarantor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuarantor

`func (o *StudentLoan) SetGuarantor(v string)`

SetGuarantor sets Guarantor field to given value.

### HasGuarantor

`func (o *StudentLoan) HasGuarantor() bool`

HasGuarantor returns a boolean if a field has been set.

### SetGuarantorNil

`func (o *StudentLoan) SetGuarantorNil(b bool)`

 SetGuarantorNil sets the value for Guarantor to be an explicit nil

### UnsetGuarantor
`func (o *StudentLoan) UnsetGuarantor()`

UnsetGuarantor ensures that no value is present for Guarantor, not even an explicit nil
### GetInterestRatePercentage

`func (o *StudentLoan) GetInterestRatePercentage() float32`

GetInterestRatePercentage returns the InterestRatePercentage field if non-nil, zero value otherwise.

### GetInterestRatePercentageOk

`func (o *StudentLoan) GetInterestRatePercentageOk() (*float32, bool)`

GetInterestRatePercentageOk returns a tuple with the InterestRatePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestRatePercentage

`func (o *StudentLoan) SetInterestRatePercentage(v float32)`

SetInterestRatePercentage sets InterestRatePercentage field to given value.


### GetIsOverdue

`func (o *StudentLoan) GetIsOverdue() bool`

GetIsOverdue returns the IsOverdue field if non-nil, zero value otherwise.

### GetIsOverdueOk

`func (o *StudentLoan) GetIsOverdueOk() (*bool, bool)`

GetIsOverdueOk returns a tuple with the IsOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdue

`func (o *StudentLoan) SetIsOverdue(v bool)`

SetIsOverdue sets IsOverdue field to given value.

### HasIsOverdue

`func (o *StudentLoan) HasIsOverdue() bool`

HasIsOverdue returns a boolean if a field has been set.

### SetIsOverdueNil

`func (o *StudentLoan) SetIsOverdueNil(b bool)`

 SetIsOverdueNil sets the value for IsOverdue to be an explicit nil

### UnsetIsOverdue
`func (o *StudentLoan) UnsetIsOverdue()`

UnsetIsOverdue ensures that no value is present for IsOverdue, not even an explicit nil
### GetLastPaymentAmount

`func (o *StudentLoan) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *StudentLoan) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *StudentLoan) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.

### HasLastPaymentAmount

`func (o *StudentLoan) HasLastPaymentAmount() bool`

HasLastPaymentAmount returns a boolean if a field has been set.

### SetLastPaymentAmountNil

`func (o *StudentLoan) SetLastPaymentAmountNil(b bool)`

 SetLastPaymentAmountNil sets the value for LastPaymentAmount to be an explicit nil

### UnsetLastPaymentAmount
`func (o *StudentLoan) UnsetLastPaymentAmount()`

UnsetLastPaymentAmount ensures that no value is present for LastPaymentAmount, not even an explicit nil
### GetLastPaymentDate

`func (o *StudentLoan) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *StudentLoan) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *StudentLoan) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.

### HasLastPaymentDate

`func (o *StudentLoan) HasLastPaymentDate() bool`

HasLastPaymentDate returns a boolean if a field has been set.

### SetLastPaymentDateNil

`func (o *StudentLoan) SetLastPaymentDateNil(b bool)`

 SetLastPaymentDateNil sets the value for LastPaymentDate to be an explicit nil

### UnsetLastPaymentDate
`func (o *StudentLoan) UnsetLastPaymentDate()`

UnsetLastPaymentDate ensures that no value is present for LastPaymentDate, not even an explicit nil
### GetLastStatementBalance

`func (o *StudentLoan) GetLastStatementBalance() float32`

GetLastStatementBalance returns the LastStatementBalance field if non-nil, zero value otherwise.

### GetLastStatementBalanceOk

`func (o *StudentLoan) GetLastStatementBalanceOk() (*float32, bool)`

GetLastStatementBalanceOk returns a tuple with the LastStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementBalance

`func (o *StudentLoan) SetLastStatementBalance(v float32)`

SetLastStatementBalance sets LastStatementBalance field to given value.

### HasLastStatementBalance

`func (o *StudentLoan) HasLastStatementBalance() bool`

HasLastStatementBalance returns a boolean if a field has been set.

### SetLastStatementBalanceNil

`func (o *StudentLoan) SetLastStatementBalanceNil(b bool)`

 SetLastStatementBalanceNil sets the value for LastStatementBalance to be an explicit nil

### UnsetLastStatementBalance
`func (o *StudentLoan) UnsetLastStatementBalance()`

UnsetLastStatementBalance ensures that no value is present for LastStatementBalance, not even an explicit nil
### GetLastStatementIssueDate

`func (o *StudentLoan) GetLastStatementIssueDate() string`

GetLastStatementIssueDate returns the LastStatementIssueDate field if non-nil, zero value otherwise.

### GetLastStatementIssueDateOk

`func (o *StudentLoan) GetLastStatementIssueDateOk() (*string, bool)`

GetLastStatementIssueDateOk returns a tuple with the LastStatementIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementIssueDate

`func (o *StudentLoan) SetLastStatementIssueDate(v string)`

SetLastStatementIssueDate sets LastStatementIssueDate field to given value.

### HasLastStatementIssueDate

`func (o *StudentLoan) HasLastStatementIssueDate() bool`

HasLastStatementIssueDate returns a boolean if a field has been set.

### SetLastStatementIssueDateNil

`func (o *StudentLoan) SetLastStatementIssueDateNil(b bool)`

 SetLastStatementIssueDateNil sets the value for LastStatementIssueDate to be an explicit nil

### UnsetLastStatementIssueDate
`func (o *StudentLoan) UnsetLastStatementIssueDate()`

UnsetLastStatementIssueDate ensures that no value is present for LastStatementIssueDate, not even an explicit nil
### GetLoanName

`func (o *StudentLoan) GetLoanName() string`

GetLoanName returns the LoanName field if non-nil, zero value otherwise.

### GetLoanNameOk

`func (o *StudentLoan) GetLoanNameOk() (*string, bool)`

GetLoanNameOk returns a tuple with the LoanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanName

`func (o *StudentLoan) SetLoanName(v string)`

SetLoanName sets LoanName field to given value.

### HasLoanName

`func (o *StudentLoan) HasLoanName() bool`

HasLoanName returns a boolean if a field has been set.

### SetLoanNameNil

`func (o *StudentLoan) SetLoanNameNil(b bool)`

 SetLoanNameNil sets the value for LoanName to be an explicit nil

### UnsetLoanName
`func (o *StudentLoan) UnsetLoanName()`

UnsetLoanName ensures that no value is present for LoanName, not even an explicit nil
### GetLoanStatus

`func (o *StudentLoan) GetLoanStatus() StudentLoanStatus`

GetLoanStatus returns the LoanStatus field if non-nil, zero value otherwise.

### GetLoanStatusOk

`func (o *StudentLoan) GetLoanStatusOk() (*StudentLoanStatus, bool)`

GetLoanStatusOk returns a tuple with the LoanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanStatus

`func (o *StudentLoan) SetLoanStatus(v StudentLoanStatus)`

SetLoanStatus sets LoanStatus field to given value.

### HasLoanStatus

`func (o *StudentLoan) HasLoanStatus() bool`

HasLoanStatus returns a boolean if a field has been set.

### GetMinimumPaymentAmount

`func (o *StudentLoan) GetMinimumPaymentAmount() float32`

GetMinimumPaymentAmount returns the MinimumPaymentAmount field if non-nil, zero value otherwise.

### GetMinimumPaymentAmountOk

`func (o *StudentLoan) GetMinimumPaymentAmountOk() (*float32, bool)`

GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentAmount

`func (o *StudentLoan) SetMinimumPaymentAmount(v float32)`

SetMinimumPaymentAmount sets MinimumPaymentAmount field to given value.

### HasMinimumPaymentAmount

`func (o *StudentLoan) HasMinimumPaymentAmount() bool`

HasMinimumPaymentAmount returns a boolean if a field has been set.

### SetMinimumPaymentAmountNil

`func (o *StudentLoan) SetMinimumPaymentAmountNil(b bool)`

 SetMinimumPaymentAmountNil sets the value for MinimumPaymentAmount to be an explicit nil

### UnsetMinimumPaymentAmount
`func (o *StudentLoan) UnsetMinimumPaymentAmount()`

UnsetMinimumPaymentAmount ensures that no value is present for MinimumPaymentAmount, not even an explicit nil
### GetNextPaymentDueDate

`func (o *StudentLoan) GetNextPaymentDueDate() string`

GetNextPaymentDueDate returns the NextPaymentDueDate field if non-nil, zero value otherwise.

### GetNextPaymentDueDateOk

`func (o *StudentLoan) GetNextPaymentDueDateOk() (*string, bool)`

GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDueDate

`func (o *StudentLoan) SetNextPaymentDueDate(v string)`

SetNextPaymentDueDate sets NextPaymentDueDate field to given value.

### HasNextPaymentDueDate

`func (o *StudentLoan) HasNextPaymentDueDate() bool`

HasNextPaymentDueDate returns a boolean if a field has been set.

### SetNextPaymentDueDateNil

`func (o *StudentLoan) SetNextPaymentDueDateNil(b bool)`

 SetNextPaymentDueDateNil sets the value for NextPaymentDueDate to be an explicit nil

### UnsetNextPaymentDueDate
`func (o *StudentLoan) UnsetNextPaymentDueDate()`

UnsetNextPaymentDueDate ensures that no value is present for NextPaymentDueDate, not even an explicit nil
### GetOriginationDate

`func (o *StudentLoan) GetOriginationDate() string`

GetOriginationDate returns the OriginationDate field if non-nil, zero value otherwise.

### GetOriginationDateOk

`func (o *StudentLoan) GetOriginationDateOk() (*string, bool)`

GetOriginationDateOk returns a tuple with the OriginationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationDate

`func (o *StudentLoan) SetOriginationDate(v string)`

SetOriginationDate sets OriginationDate field to given value.

### HasOriginationDate

`func (o *StudentLoan) HasOriginationDate() bool`

HasOriginationDate returns a boolean if a field has been set.

### SetOriginationDateNil

`func (o *StudentLoan) SetOriginationDateNil(b bool)`

 SetOriginationDateNil sets the value for OriginationDate to be an explicit nil

### UnsetOriginationDate
`func (o *StudentLoan) UnsetOriginationDate()`

UnsetOriginationDate ensures that no value is present for OriginationDate, not even an explicit nil
### GetOriginationPrincipalAmount

`func (o *StudentLoan) GetOriginationPrincipalAmount() float32`

GetOriginationPrincipalAmount returns the OriginationPrincipalAmount field if non-nil, zero value otherwise.

### GetOriginationPrincipalAmountOk

`func (o *StudentLoan) GetOriginationPrincipalAmountOk() (*float32, bool)`

GetOriginationPrincipalAmountOk returns a tuple with the OriginationPrincipalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationPrincipalAmount

`func (o *StudentLoan) SetOriginationPrincipalAmount(v float32)`

SetOriginationPrincipalAmount sets OriginationPrincipalAmount field to given value.

### HasOriginationPrincipalAmount

`func (o *StudentLoan) HasOriginationPrincipalAmount() bool`

HasOriginationPrincipalAmount returns a boolean if a field has been set.

### SetOriginationPrincipalAmountNil

`func (o *StudentLoan) SetOriginationPrincipalAmountNil(b bool)`

 SetOriginationPrincipalAmountNil sets the value for OriginationPrincipalAmount to be an explicit nil

### UnsetOriginationPrincipalAmount
`func (o *StudentLoan) UnsetOriginationPrincipalAmount()`

UnsetOriginationPrincipalAmount ensures that no value is present for OriginationPrincipalAmount, not even an explicit nil
### GetOutstandingInterestAmount

`func (o *StudentLoan) GetOutstandingInterestAmount() float32`

GetOutstandingInterestAmount returns the OutstandingInterestAmount field if non-nil, zero value otherwise.

### GetOutstandingInterestAmountOk

`func (o *StudentLoan) GetOutstandingInterestAmountOk() (*float32, bool)`

GetOutstandingInterestAmountOk returns a tuple with the OutstandingInterestAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingInterestAmount

`func (o *StudentLoan) SetOutstandingInterestAmount(v float32)`

SetOutstandingInterestAmount sets OutstandingInterestAmount field to given value.

### HasOutstandingInterestAmount

`func (o *StudentLoan) HasOutstandingInterestAmount() bool`

HasOutstandingInterestAmount returns a boolean if a field has been set.

### SetOutstandingInterestAmountNil

`func (o *StudentLoan) SetOutstandingInterestAmountNil(b bool)`

 SetOutstandingInterestAmountNil sets the value for OutstandingInterestAmount to be an explicit nil

### UnsetOutstandingInterestAmount
`func (o *StudentLoan) UnsetOutstandingInterestAmount()`

UnsetOutstandingInterestAmount ensures that no value is present for OutstandingInterestAmount, not even an explicit nil
### GetPaymentReferenceNumber

`func (o *StudentLoan) GetPaymentReferenceNumber() string`

GetPaymentReferenceNumber returns the PaymentReferenceNumber field if non-nil, zero value otherwise.

### GetPaymentReferenceNumberOk

`func (o *StudentLoan) GetPaymentReferenceNumberOk() (*string, bool)`

GetPaymentReferenceNumberOk returns a tuple with the PaymentReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReferenceNumber

`func (o *StudentLoan) SetPaymentReferenceNumber(v string)`

SetPaymentReferenceNumber sets PaymentReferenceNumber field to given value.

### HasPaymentReferenceNumber

`func (o *StudentLoan) HasPaymentReferenceNumber() bool`

HasPaymentReferenceNumber returns a boolean if a field has been set.

### SetPaymentReferenceNumberNil

`func (o *StudentLoan) SetPaymentReferenceNumberNil(b bool)`

 SetPaymentReferenceNumberNil sets the value for PaymentReferenceNumber to be an explicit nil

### UnsetPaymentReferenceNumber
`func (o *StudentLoan) UnsetPaymentReferenceNumber()`

UnsetPaymentReferenceNumber ensures that no value is present for PaymentReferenceNumber, not even an explicit nil
### GetPslfStatus

`func (o *StudentLoan) GetPslfStatus() PSLFStatus`

GetPslfStatus returns the PslfStatus field if non-nil, zero value otherwise.

### GetPslfStatusOk

`func (o *StudentLoan) GetPslfStatusOk() (*PSLFStatus, bool)`

GetPslfStatusOk returns a tuple with the PslfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPslfStatus

`func (o *StudentLoan) SetPslfStatus(v PSLFStatus)`

SetPslfStatus sets PslfStatus field to given value.

### HasPslfStatus

`func (o *StudentLoan) HasPslfStatus() bool`

HasPslfStatus returns a boolean if a field has been set.

### GetRepaymentPlan

`func (o *StudentLoan) GetRepaymentPlan() StudentRepaymentPlan`

GetRepaymentPlan returns the RepaymentPlan field if non-nil, zero value otherwise.

### GetRepaymentPlanOk

`func (o *StudentLoan) GetRepaymentPlanOk() (*StudentRepaymentPlan, bool)`

GetRepaymentPlanOk returns a tuple with the RepaymentPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentPlan

`func (o *StudentLoan) SetRepaymentPlan(v StudentRepaymentPlan)`

SetRepaymentPlan sets RepaymentPlan field to given value.

### HasRepaymentPlan

`func (o *StudentLoan) HasRepaymentPlan() bool`

HasRepaymentPlan returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *StudentLoan) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *StudentLoan) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *StudentLoan) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *StudentLoan) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *StudentLoan) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *StudentLoan) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetServicerAddress

`func (o *StudentLoan) GetServicerAddress() ServicerAddressData`

GetServicerAddress returns the ServicerAddress field if non-nil, zero value otherwise.

### GetServicerAddressOk

`func (o *StudentLoan) GetServicerAddressOk() (*ServicerAddressData, bool)`

GetServicerAddressOk returns a tuple with the ServicerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicerAddress

`func (o *StudentLoan) SetServicerAddress(v ServicerAddressData)`

SetServicerAddress sets ServicerAddress field to given value.

### HasServicerAddress

`func (o *StudentLoan) HasServicerAddress() bool`

HasServicerAddress returns a boolean if a field has been set.

### GetYtdInterestPaid

`func (o *StudentLoan) GetYtdInterestPaid() float32`

GetYtdInterestPaid returns the YtdInterestPaid field if non-nil, zero value otherwise.

### GetYtdInterestPaidOk

`func (o *StudentLoan) GetYtdInterestPaidOk() (*float32, bool)`

GetYtdInterestPaidOk returns a tuple with the YtdInterestPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdInterestPaid

`func (o *StudentLoan) SetYtdInterestPaid(v float32)`

SetYtdInterestPaid sets YtdInterestPaid field to given value.

### HasYtdInterestPaid

`func (o *StudentLoan) HasYtdInterestPaid() bool`

HasYtdInterestPaid returns a boolean if a field has been set.

### SetYtdInterestPaidNil

`func (o *StudentLoan) SetYtdInterestPaidNil(b bool)`

 SetYtdInterestPaidNil sets the value for YtdInterestPaid to be an explicit nil

### UnsetYtdInterestPaid
`func (o *StudentLoan) UnsetYtdInterestPaid()`

UnsetYtdInterestPaid ensures that no value is present for YtdInterestPaid, not even an explicit nil
### GetYtdPrincipalPaid

`func (o *StudentLoan) GetYtdPrincipalPaid() float32`

GetYtdPrincipalPaid returns the YtdPrincipalPaid field if non-nil, zero value otherwise.

### GetYtdPrincipalPaidOk

`func (o *StudentLoan) GetYtdPrincipalPaidOk() (*float32, bool)`

GetYtdPrincipalPaidOk returns a tuple with the YtdPrincipalPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYtdPrincipalPaid

`func (o *StudentLoan) SetYtdPrincipalPaid(v float32)`

SetYtdPrincipalPaid sets YtdPrincipalPaid field to given value.

### HasYtdPrincipalPaid

`func (o *StudentLoan) HasYtdPrincipalPaid() bool`

HasYtdPrincipalPaid returns a boolean if a field has been set.

### SetYtdPrincipalPaidNil

`func (o *StudentLoan) SetYtdPrincipalPaidNil(b bool)`

 SetYtdPrincipalPaidNil sets the value for YtdPrincipalPaid to be an explicit nil

### UnsetYtdPrincipalPaid
`func (o *StudentLoan) UnsetYtdPrincipalPaid()`

UnsetYtdPrincipalPaid ensures that no value is present for YtdPrincipalPaid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


