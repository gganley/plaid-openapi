# LiabilityOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the liability object, either &#x60;credit&#x60; or &#x60;student&#x60;. | 
**PurchaseApr** | **float32** | The purchase APR percentage value. For simplicity, this is the only interest rate used to calculate interest charges. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**CashApr** | **float32** | The cash APR percentage value. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**BalanceTransferApr** | **float32** | The balance transfer APR percentage value. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**SpecialApr** | **float32** | The special APR percentage value. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**LastPaymentAmount** | **float32** | Override the &#x60;last_payment_amount&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**LastStatementBalance** | **float32** | Override the &#x60;last_statement_balance&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**MinimumPaymentAmount** | **float32** | Override the &#x60;minimum_payment_amount&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;credit&#x60;. | 
**IsOverdue** | **bool** | Override the &#x60;is_overdue&#x60; field | 
**OriginationDate** | **string** | The date on which the loan was initially lent, in ISO 8601 (YYYY-MM-DD) format. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**Principal** | **float32** | The original loan principal. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**NominalApr** | **float32** | The interest rate on the loan as a percentage. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**InterestCapitalizationGracePeriodMonths** | **float32** | If set, interest capitalization begins at the given number of months after loan origination. By default interest is never capitalized. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**RepaymentModel** | [**StudentLoanRepaymentModel**](StudentLoanRepaymentModel.md) |  | 
**ExpectedPayoffDate** | **string** | Override the &#x60;expected_payoff_date&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**Guarantor** | **string** | Override the &#x60;guarantor&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**IsFederal** | **bool** | Override the &#x60;is_federal&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**LoanName** | **string** | Override the &#x60;loan_name&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**LoanStatus** | **string** | Override the &#x60;loan_status&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**PaymentReferenceNumber** | **string** | Override the &#x60;payment_reference_number&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**PslfStatus** | **string** | Override the &#x60;pslf_status&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**RepaymentPlanDescription** | **string** | Override the &#x60;repayment_plan.description&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**RepaymentPlanType** | **string** | Override the &#x60;repayment_plan.type&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. Possible values are: &#x60;\&quot;extended graduated\&quot;&#x60;, &#x60;\&quot;extended standard\&quot;&#x60;, &#x60;\&quot;graduated\&quot;&#x60;, &#x60;\&quot;income-contingent repayment\&quot;&#x60;, &#x60;\&quot;income-based repayment\&quot;&#x60;, &#x60;\&quot;interest only\&quot;&#x60;, &#x60;\&quot;other\&quot;&#x60;, &#x60;\&quot;pay as you earn\&quot;&#x60;, &#x60;\&quot;revised pay as you earn\&quot;&#x60;, or &#x60;\&quot;standard\&quot;&#x60;. | 
**SequenceNumber** | **string** | Override the &#x60;sequence_number&#x60; field. Can only be set if &#x60;type&#x60; is &#x60;student&#x60;. | 
**ServicerAddress** | [**NullableAddress**](Address.md) |  | 

## Methods

### NewLiabilityOverride

`func NewLiabilityOverride(type_ string, purchaseApr float32, cashApr float32, balanceTransferApr float32, specialApr float32, lastPaymentAmount float32, lastStatementBalance float32, minimumPaymentAmount float32, isOverdue bool, originationDate string, principal float32, nominalApr float32, interestCapitalizationGracePeriodMonths float32, repaymentModel StudentLoanRepaymentModel, expectedPayoffDate string, guarantor string, isFederal bool, loanName string, loanStatus string, paymentReferenceNumber string, pslfStatus string, repaymentPlanDescription string, repaymentPlanType string, sequenceNumber string, servicerAddress NullableAddress, ) *LiabilityOverride`

NewLiabilityOverride instantiates a new LiabilityOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiabilityOverrideWithDefaults

`func NewLiabilityOverrideWithDefaults() *LiabilityOverride`

NewLiabilityOverrideWithDefaults instantiates a new LiabilityOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LiabilityOverride) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LiabilityOverride) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LiabilityOverride) SetType(v string)`

SetType sets Type field to given value.


### GetPurchaseApr

`func (o *LiabilityOverride) GetPurchaseApr() float32`

GetPurchaseApr returns the PurchaseApr field if non-nil, zero value otherwise.

### GetPurchaseAprOk

`func (o *LiabilityOverride) GetPurchaseAprOk() (*float32, bool)`

GetPurchaseAprOk returns a tuple with the PurchaseApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseApr

`func (o *LiabilityOverride) SetPurchaseApr(v float32)`

SetPurchaseApr sets PurchaseApr field to given value.


### GetCashApr

`func (o *LiabilityOverride) GetCashApr() float32`

GetCashApr returns the CashApr field if non-nil, zero value otherwise.

### GetCashAprOk

`func (o *LiabilityOverride) GetCashAprOk() (*float32, bool)`

GetCashAprOk returns a tuple with the CashApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashApr

`func (o *LiabilityOverride) SetCashApr(v float32)`

SetCashApr sets CashApr field to given value.


### GetBalanceTransferApr

`func (o *LiabilityOverride) GetBalanceTransferApr() float32`

GetBalanceTransferApr returns the BalanceTransferApr field if non-nil, zero value otherwise.

### GetBalanceTransferAprOk

`func (o *LiabilityOverride) GetBalanceTransferAprOk() (*float32, bool)`

GetBalanceTransferAprOk returns a tuple with the BalanceTransferApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceTransferApr

`func (o *LiabilityOverride) SetBalanceTransferApr(v float32)`

SetBalanceTransferApr sets BalanceTransferApr field to given value.


### GetSpecialApr

`func (o *LiabilityOverride) GetSpecialApr() float32`

GetSpecialApr returns the SpecialApr field if non-nil, zero value otherwise.

### GetSpecialAprOk

`func (o *LiabilityOverride) GetSpecialAprOk() (*float32, bool)`

GetSpecialAprOk returns a tuple with the SpecialApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialApr

`func (o *LiabilityOverride) SetSpecialApr(v float32)`

SetSpecialApr sets SpecialApr field to given value.


### GetLastPaymentAmount

`func (o *LiabilityOverride) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *LiabilityOverride) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *LiabilityOverride) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.


### GetLastStatementBalance

`func (o *LiabilityOverride) GetLastStatementBalance() float32`

GetLastStatementBalance returns the LastStatementBalance field if non-nil, zero value otherwise.

### GetLastStatementBalanceOk

`func (o *LiabilityOverride) GetLastStatementBalanceOk() (*float32, bool)`

GetLastStatementBalanceOk returns a tuple with the LastStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementBalance

`func (o *LiabilityOverride) SetLastStatementBalance(v float32)`

SetLastStatementBalance sets LastStatementBalance field to given value.


### GetMinimumPaymentAmount

`func (o *LiabilityOverride) GetMinimumPaymentAmount() float32`

GetMinimumPaymentAmount returns the MinimumPaymentAmount field if non-nil, zero value otherwise.

### GetMinimumPaymentAmountOk

`func (o *LiabilityOverride) GetMinimumPaymentAmountOk() (*float32, bool)`

GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentAmount

`func (o *LiabilityOverride) SetMinimumPaymentAmount(v float32)`

SetMinimumPaymentAmount sets MinimumPaymentAmount field to given value.


### GetIsOverdue

`func (o *LiabilityOverride) GetIsOverdue() bool`

GetIsOverdue returns the IsOverdue field if non-nil, zero value otherwise.

### GetIsOverdueOk

`func (o *LiabilityOverride) GetIsOverdueOk() (*bool, bool)`

GetIsOverdueOk returns a tuple with the IsOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdue

`func (o *LiabilityOverride) SetIsOverdue(v bool)`

SetIsOverdue sets IsOverdue field to given value.


### GetOriginationDate

`func (o *LiabilityOverride) GetOriginationDate() string`

GetOriginationDate returns the OriginationDate field if non-nil, zero value otherwise.

### GetOriginationDateOk

`func (o *LiabilityOverride) GetOriginationDateOk() (*string, bool)`

GetOriginationDateOk returns a tuple with the OriginationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationDate

`func (o *LiabilityOverride) SetOriginationDate(v string)`

SetOriginationDate sets OriginationDate field to given value.


### GetPrincipal

`func (o *LiabilityOverride) GetPrincipal() float32`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *LiabilityOverride) GetPrincipalOk() (*float32, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *LiabilityOverride) SetPrincipal(v float32)`

SetPrincipal sets Principal field to given value.


### GetNominalApr

`func (o *LiabilityOverride) GetNominalApr() float32`

GetNominalApr returns the NominalApr field if non-nil, zero value otherwise.

### GetNominalAprOk

`func (o *LiabilityOverride) GetNominalAprOk() (*float32, bool)`

GetNominalAprOk returns a tuple with the NominalApr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominalApr

`func (o *LiabilityOverride) SetNominalApr(v float32)`

SetNominalApr sets NominalApr field to given value.


### GetInterestCapitalizationGracePeriodMonths

`func (o *LiabilityOverride) GetInterestCapitalizationGracePeriodMonths() float32`

GetInterestCapitalizationGracePeriodMonths returns the InterestCapitalizationGracePeriodMonths field if non-nil, zero value otherwise.

### GetInterestCapitalizationGracePeriodMonthsOk

`func (o *LiabilityOverride) GetInterestCapitalizationGracePeriodMonthsOk() (*float32, bool)`

GetInterestCapitalizationGracePeriodMonthsOk returns a tuple with the InterestCapitalizationGracePeriodMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestCapitalizationGracePeriodMonths

`func (o *LiabilityOverride) SetInterestCapitalizationGracePeriodMonths(v float32)`

SetInterestCapitalizationGracePeriodMonths sets InterestCapitalizationGracePeriodMonths field to given value.


### GetRepaymentModel

`func (o *LiabilityOverride) GetRepaymentModel() StudentLoanRepaymentModel`

GetRepaymentModel returns the RepaymentModel field if non-nil, zero value otherwise.

### GetRepaymentModelOk

`func (o *LiabilityOverride) GetRepaymentModelOk() (*StudentLoanRepaymentModel, bool)`

GetRepaymentModelOk returns a tuple with the RepaymentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentModel

`func (o *LiabilityOverride) SetRepaymentModel(v StudentLoanRepaymentModel)`

SetRepaymentModel sets RepaymentModel field to given value.


### GetExpectedPayoffDate

`func (o *LiabilityOverride) GetExpectedPayoffDate() string`

GetExpectedPayoffDate returns the ExpectedPayoffDate field if non-nil, zero value otherwise.

### GetExpectedPayoffDateOk

`func (o *LiabilityOverride) GetExpectedPayoffDateOk() (*string, bool)`

GetExpectedPayoffDateOk returns a tuple with the ExpectedPayoffDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedPayoffDate

`func (o *LiabilityOverride) SetExpectedPayoffDate(v string)`

SetExpectedPayoffDate sets ExpectedPayoffDate field to given value.


### GetGuarantor

`func (o *LiabilityOverride) GetGuarantor() string`

GetGuarantor returns the Guarantor field if non-nil, zero value otherwise.

### GetGuarantorOk

`func (o *LiabilityOverride) GetGuarantorOk() (*string, bool)`

GetGuarantorOk returns a tuple with the Guarantor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuarantor

`func (o *LiabilityOverride) SetGuarantor(v string)`

SetGuarantor sets Guarantor field to given value.


### GetIsFederal

`func (o *LiabilityOverride) GetIsFederal() bool`

GetIsFederal returns the IsFederal field if non-nil, zero value otherwise.

### GetIsFederalOk

`func (o *LiabilityOverride) GetIsFederalOk() (*bool, bool)`

GetIsFederalOk returns a tuple with the IsFederal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFederal

`func (o *LiabilityOverride) SetIsFederal(v bool)`

SetIsFederal sets IsFederal field to given value.


### GetLoanName

`func (o *LiabilityOverride) GetLoanName() string`

GetLoanName returns the LoanName field if non-nil, zero value otherwise.

### GetLoanNameOk

`func (o *LiabilityOverride) GetLoanNameOk() (*string, bool)`

GetLoanNameOk returns a tuple with the LoanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanName

`func (o *LiabilityOverride) SetLoanName(v string)`

SetLoanName sets LoanName field to given value.


### GetLoanStatus

`func (o *LiabilityOverride) GetLoanStatus() string`

GetLoanStatus returns the LoanStatus field if non-nil, zero value otherwise.

### GetLoanStatusOk

`func (o *LiabilityOverride) GetLoanStatusOk() (*string, bool)`

GetLoanStatusOk returns a tuple with the LoanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanStatus

`func (o *LiabilityOverride) SetLoanStatus(v string)`

SetLoanStatus sets LoanStatus field to given value.


### GetPaymentReferenceNumber

`func (o *LiabilityOverride) GetPaymentReferenceNumber() string`

GetPaymentReferenceNumber returns the PaymentReferenceNumber field if non-nil, zero value otherwise.

### GetPaymentReferenceNumberOk

`func (o *LiabilityOverride) GetPaymentReferenceNumberOk() (*string, bool)`

GetPaymentReferenceNumberOk returns a tuple with the PaymentReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReferenceNumber

`func (o *LiabilityOverride) SetPaymentReferenceNumber(v string)`

SetPaymentReferenceNumber sets PaymentReferenceNumber field to given value.


### GetPslfStatus

`func (o *LiabilityOverride) GetPslfStatus() string`

GetPslfStatus returns the PslfStatus field if non-nil, zero value otherwise.

### GetPslfStatusOk

`func (o *LiabilityOverride) GetPslfStatusOk() (*string, bool)`

GetPslfStatusOk returns a tuple with the PslfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPslfStatus

`func (o *LiabilityOverride) SetPslfStatus(v string)`

SetPslfStatus sets PslfStatus field to given value.


### GetRepaymentPlanDescription

`func (o *LiabilityOverride) GetRepaymentPlanDescription() string`

GetRepaymentPlanDescription returns the RepaymentPlanDescription field if non-nil, zero value otherwise.

### GetRepaymentPlanDescriptionOk

`func (o *LiabilityOverride) GetRepaymentPlanDescriptionOk() (*string, bool)`

GetRepaymentPlanDescriptionOk returns a tuple with the RepaymentPlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentPlanDescription

`func (o *LiabilityOverride) SetRepaymentPlanDescription(v string)`

SetRepaymentPlanDescription sets RepaymentPlanDescription field to given value.


### GetRepaymentPlanType

`func (o *LiabilityOverride) GetRepaymentPlanType() string`

GetRepaymentPlanType returns the RepaymentPlanType field if non-nil, zero value otherwise.

### GetRepaymentPlanTypeOk

`func (o *LiabilityOverride) GetRepaymentPlanTypeOk() (*string, bool)`

GetRepaymentPlanTypeOk returns a tuple with the RepaymentPlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentPlanType

`func (o *LiabilityOverride) SetRepaymentPlanType(v string)`

SetRepaymentPlanType sets RepaymentPlanType field to given value.


### GetSequenceNumber

`func (o *LiabilityOverride) GetSequenceNumber() string`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *LiabilityOverride) GetSequenceNumberOk() (*string, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *LiabilityOverride) SetSequenceNumber(v string)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetServicerAddress

`func (o *LiabilityOverride) GetServicerAddress() Address`

GetServicerAddress returns the ServicerAddress field if non-nil, zero value otherwise.

### GetServicerAddressOk

`func (o *LiabilityOverride) GetServicerAddressOk() (*Address, bool)`

GetServicerAddressOk returns a tuple with the ServicerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicerAddress

`func (o *LiabilityOverride) SetServicerAddress(v Address)`

SetServicerAddress sets ServicerAddress field to given value.


### SetServicerAddressNil

`func (o *LiabilityOverride) SetServicerAddressNil(b bool)`

 SetServicerAddressNil sets the value for ServicerAddress to be an explicit nil

### UnsetServicerAddress
`func (o *LiabilityOverride) UnsetServicerAddress()`

UnsetServicerAddress ensures that no value is present for ServicerAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


