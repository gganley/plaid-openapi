# StandaloneInvestmentTransactionSubtype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountFee** | Pointer to **string** | Fees paid for account maintenance | [optional] 
**Assignment** | Pointer to **string** | Assignment of short option holding | [optional] 
**Buy** | Pointer to **string** | Purchase to open or increase a position | [optional] 
**BuyToCover** | Pointer to **string** | Purchase to close a short position | [optional] 
**Contribution** | Pointer to **string** | Inflow of assets into a tax-advantaged account | [optional] 
**Deposit** | Pointer to **string** | Inflow of cash into an account | [optional] 
**Distribution** | Pointer to **string** | Outflow of assets from a tax-advantaged account | [optional] 
**Dividend** | Pointer to **string** | Inflow of cash from a dividend | [optional] 
**DividendReinvestment** | Pointer to **string** | Purchase using proceeds from a cash dividend | [optional] 
**Exercise** | Pointer to **string** | Exercise of an option or warrant contract | [optional] 
**Expire** | Pointer to **string** | Expiration of an option or warrant contract  | [optional] 
**FundFee** | Pointer to **string** | Fees paid for administration of a mutual fund or other pooled investment vehicle | [optional] 
**Interest** | Pointer to **string** | Inflow of cash from interest | [optional] 
**InterestReceivable** | Pointer to **string** | Inflow of cash from interest receivable | [optional] 
**InterestReinvestment** | Pointer to **string** | Purchase using proceeds from a cash interest payment | [optional] 
**LegalFee** | Pointer to **string** | Fees paid for legal charges or services | [optional] 
**LoanPayment** | Pointer to **string** | Inflow of cash related to payment on a loan | [optional] 
**LongTermCapitalGain** | Pointer to **string** | Long-term capital gain received as cash | [optional] 
**LongTermCapitalGainReinvestment** | Pointer to **string** | Purchase using long-term capital gain cash proceeds | [optional] 
**ManagementFee** | Pointer to **string** | Fees paid for investment management of a mutual fund or other pooled investment vehicle | [optional] 
**MarginExpense** | Pointer to **string** | Fees paid for maintaining margin debt | [optional] 
**Merger** | Pointer to **string** | Stock exchanged at a pre-defined ratio as part of a merger between companies | [optional] 
**MiscellaneousFee** | Pointer to **string** | Fee associated with various account or holding actions | [optional] 
**NonQualifiedDividend** | Pointer to **string** | Inflow of cash from a non-qualified dividend | [optional] 
**NonResidentTax** | Pointer to **string** | Taxes paid on behalf of the investor for non-residency in investment jurisdiction | [optional] 
**PendingCredit** | Pointer to **string** | Pending inflow of cash | [optional] 
**PendingDebit** | Pointer to **string** | Pending outflow of cash | [optional] 
**QualifiedDividend** | Pointer to **string** | Inflow of cash from a qualified dividend | [optional] 
**Rebalance** | Pointer to **string** | Rebalancing transaction (buy or sell) with no net impact to value in the account | [optional] 
**ReturnOfPrincipal** | Pointer to **string** | Repayment of loan principal | [optional] 
**Sell** | Pointer to **string** | Sell to close or decrease an existing holding | [optional] 
**SellShort** | Pointer to **string** | Sell to open a short position | [optional] 
**ShortTermCapitalGain** | Pointer to **string** | Short-term capital gain received as cash | [optional] 
**ShortTermCapitalGainReinvestment** | Pointer to **string** | Purchase using short-term capital gain cash proceeds | [optional] 
**SpinOff** | Pointer to **string** | Inflow of stock from spin-off transaction of an existing holding | [optional] 
**Split** | Pointer to **string** | Inflow of stock from a forward split of an existing holding | [optional] 
**StockDistribution** | Pointer to **string** | Inflow of stock from a distribution | [optional] 
**Tax** | Pointer to **string** | Taxes paid on behalf of the investor | [optional] 
**TaxWithheld** | Pointer to **string** | Taxes withheld on behalf of the customer | [optional] 
**Transfer** | Pointer to **string** | Movement of assets into or out of an account | [optional] 
**TransferFee** | Pointer to **string** | Fees incurred for transfer of a holding or account | [optional] 
**TrustFee** | Pointer to **string** | Fees related to adminstration of a trust account | [optional] 
**UnqualifiedGain** | Pointer to **string** | Unqualified capital gain received as cash | [optional] 
**Withdrawal** | Pointer to **string** | Outflow of cash from an account | [optional] 

## Methods

### NewStandaloneInvestmentTransactionSubtype

`func NewStandaloneInvestmentTransactionSubtype() *StandaloneInvestmentTransactionSubtype`

NewStandaloneInvestmentTransactionSubtype instantiates a new StandaloneInvestmentTransactionSubtype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneInvestmentTransactionSubtypeWithDefaults

`func NewStandaloneInvestmentTransactionSubtypeWithDefaults() *StandaloneInvestmentTransactionSubtype`

NewStandaloneInvestmentTransactionSubtypeWithDefaults instantiates a new StandaloneInvestmentTransactionSubtype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountFee

`func (o *StandaloneInvestmentTransactionSubtype) GetAccountFee() string`

GetAccountFee returns the AccountFee field if non-nil, zero value otherwise.

### GetAccountFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetAccountFeeOk() (*string, bool)`

GetAccountFeeOk returns a tuple with the AccountFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFee

`func (o *StandaloneInvestmentTransactionSubtype) SetAccountFee(v string)`

SetAccountFee sets AccountFee field to given value.

### HasAccountFee

`func (o *StandaloneInvestmentTransactionSubtype) HasAccountFee() bool`

HasAccountFee returns a boolean if a field has been set.

### GetAssignment

`func (o *StandaloneInvestmentTransactionSubtype) GetAssignment() string`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetAssignmentOk() (*string, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *StandaloneInvestmentTransactionSubtype) SetAssignment(v string)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *StandaloneInvestmentTransactionSubtype) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### GetBuy

`func (o *StandaloneInvestmentTransactionSubtype) GetBuy() string`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *StandaloneInvestmentTransactionSubtype) GetBuyOk() (*string, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *StandaloneInvestmentTransactionSubtype) SetBuy(v string)`

SetBuy sets Buy field to given value.

### HasBuy

`func (o *StandaloneInvestmentTransactionSubtype) HasBuy() bool`

HasBuy returns a boolean if a field has been set.

### GetBuyToCover

`func (o *StandaloneInvestmentTransactionSubtype) GetBuyToCover() string`

GetBuyToCover returns the BuyToCover field if non-nil, zero value otherwise.

### GetBuyToCoverOk

`func (o *StandaloneInvestmentTransactionSubtype) GetBuyToCoverOk() (*string, bool)`

GetBuyToCoverOk returns a tuple with the BuyToCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyToCover

`func (o *StandaloneInvestmentTransactionSubtype) SetBuyToCover(v string)`

SetBuyToCover sets BuyToCover field to given value.

### HasBuyToCover

`func (o *StandaloneInvestmentTransactionSubtype) HasBuyToCover() bool`

HasBuyToCover returns a boolean if a field has been set.

### GetContribution

`func (o *StandaloneInvestmentTransactionSubtype) GetContribution() string`

GetContribution returns the Contribution field if non-nil, zero value otherwise.

### GetContributionOk

`func (o *StandaloneInvestmentTransactionSubtype) GetContributionOk() (*string, bool)`

GetContributionOk returns a tuple with the Contribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContribution

`func (o *StandaloneInvestmentTransactionSubtype) SetContribution(v string)`

SetContribution sets Contribution field to given value.

### HasContribution

`func (o *StandaloneInvestmentTransactionSubtype) HasContribution() bool`

HasContribution returns a boolean if a field has been set.

### GetDeposit

`func (o *StandaloneInvestmentTransactionSubtype) GetDeposit() string`

GetDeposit returns the Deposit field if non-nil, zero value otherwise.

### GetDepositOk

`func (o *StandaloneInvestmentTransactionSubtype) GetDepositOk() (*string, bool)`

GetDepositOk returns a tuple with the Deposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeposit

`func (o *StandaloneInvestmentTransactionSubtype) SetDeposit(v string)`

SetDeposit sets Deposit field to given value.

### HasDeposit

`func (o *StandaloneInvestmentTransactionSubtype) HasDeposit() bool`

HasDeposit returns a boolean if a field has been set.

### GetDistribution

`func (o *StandaloneInvestmentTransactionSubtype) GetDistribution() string`

GetDistribution returns the Distribution field if non-nil, zero value otherwise.

### GetDistributionOk

`func (o *StandaloneInvestmentTransactionSubtype) GetDistributionOk() (*string, bool)`

GetDistributionOk returns a tuple with the Distribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution

`func (o *StandaloneInvestmentTransactionSubtype) SetDistribution(v string)`

SetDistribution sets Distribution field to given value.

### HasDistribution

`func (o *StandaloneInvestmentTransactionSubtype) HasDistribution() bool`

HasDistribution returns a boolean if a field has been set.

### GetDividend

`func (o *StandaloneInvestmentTransactionSubtype) GetDividend() string`

GetDividend returns the Dividend field if non-nil, zero value otherwise.

### GetDividendOk

`func (o *StandaloneInvestmentTransactionSubtype) GetDividendOk() (*string, bool)`

GetDividendOk returns a tuple with the Dividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividend

`func (o *StandaloneInvestmentTransactionSubtype) SetDividend(v string)`

SetDividend sets Dividend field to given value.

### HasDividend

`func (o *StandaloneInvestmentTransactionSubtype) HasDividend() bool`

HasDividend returns a boolean if a field has been set.

### GetDividendReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) GetDividendReinvestment() string`

GetDividendReinvestment returns the DividendReinvestment field if non-nil, zero value otherwise.

### GetDividendReinvestmentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetDividendReinvestmentOk() (*string, bool)`

GetDividendReinvestmentOk returns a tuple with the DividendReinvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) SetDividendReinvestment(v string)`

SetDividendReinvestment sets DividendReinvestment field to given value.

### HasDividendReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) HasDividendReinvestment() bool`

HasDividendReinvestment returns a boolean if a field has been set.

### GetExercise

`func (o *StandaloneInvestmentTransactionSubtype) GetExercise() string`

GetExercise returns the Exercise field if non-nil, zero value otherwise.

### GetExerciseOk

`func (o *StandaloneInvestmentTransactionSubtype) GetExerciseOk() (*string, bool)`

GetExerciseOk returns a tuple with the Exercise field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExercise

`func (o *StandaloneInvestmentTransactionSubtype) SetExercise(v string)`

SetExercise sets Exercise field to given value.

### HasExercise

`func (o *StandaloneInvestmentTransactionSubtype) HasExercise() bool`

HasExercise returns a boolean if a field has been set.

### GetExpire

`func (o *StandaloneInvestmentTransactionSubtype) GetExpire() string`

GetExpire returns the Expire field if non-nil, zero value otherwise.

### GetExpireOk

`func (o *StandaloneInvestmentTransactionSubtype) GetExpireOk() (*string, bool)`

GetExpireOk returns a tuple with the Expire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpire

`func (o *StandaloneInvestmentTransactionSubtype) SetExpire(v string)`

SetExpire sets Expire field to given value.

### HasExpire

`func (o *StandaloneInvestmentTransactionSubtype) HasExpire() bool`

HasExpire returns a boolean if a field has been set.

### GetFundFee

`func (o *StandaloneInvestmentTransactionSubtype) GetFundFee() string`

GetFundFee returns the FundFee field if non-nil, zero value otherwise.

### GetFundFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetFundFeeOk() (*string, bool)`

GetFundFeeOk returns a tuple with the FundFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFee

`func (o *StandaloneInvestmentTransactionSubtype) SetFundFee(v string)`

SetFundFee sets FundFee field to given value.

### HasFundFee

`func (o *StandaloneInvestmentTransactionSubtype) HasFundFee() bool`

HasFundFee returns a boolean if a field has been set.

### GetInterest

`func (o *StandaloneInvestmentTransactionSubtype) GetInterest() string`

GetInterest returns the Interest field if non-nil, zero value otherwise.

### GetInterestOk

`func (o *StandaloneInvestmentTransactionSubtype) GetInterestOk() (*string, bool)`

GetInterestOk returns a tuple with the Interest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterest

`func (o *StandaloneInvestmentTransactionSubtype) SetInterest(v string)`

SetInterest sets Interest field to given value.

### HasInterest

`func (o *StandaloneInvestmentTransactionSubtype) HasInterest() bool`

HasInterest returns a boolean if a field has been set.

### GetInterestReceivable

`func (o *StandaloneInvestmentTransactionSubtype) GetInterestReceivable() string`

GetInterestReceivable returns the InterestReceivable field if non-nil, zero value otherwise.

### GetInterestReceivableOk

`func (o *StandaloneInvestmentTransactionSubtype) GetInterestReceivableOk() (*string, bool)`

GetInterestReceivableOk returns a tuple with the InterestReceivable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestReceivable

`func (o *StandaloneInvestmentTransactionSubtype) SetInterestReceivable(v string)`

SetInterestReceivable sets InterestReceivable field to given value.

### HasInterestReceivable

`func (o *StandaloneInvestmentTransactionSubtype) HasInterestReceivable() bool`

HasInterestReceivable returns a boolean if a field has been set.

### GetInterestReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) GetInterestReinvestment() string`

GetInterestReinvestment returns the InterestReinvestment field if non-nil, zero value otherwise.

### GetInterestReinvestmentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetInterestReinvestmentOk() (*string, bool)`

GetInterestReinvestmentOk returns a tuple with the InterestReinvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterestReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) SetInterestReinvestment(v string)`

SetInterestReinvestment sets InterestReinvestment field to given value.

### HasInterestReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) HasInterestReinvestment() bool`

HasInterestReinvestment returns a boolean if a field has been set.

### GetLegalFee

`func (o *StandaloneInvestmentTransactionSubtype) GetLegalFee() string`

GetLegalFee returns the LegalFee field if non-nil, zero value otherwise.

### GetLegalFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetLegalFeeOk() (*string, bool)`

GetLegalFeeOk returns a tuple with the LegalFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalFee

`func (o *StandaloneInvestmentTransactionSubtype) SetLegalFee(v string)`

SetLegalFee sets LegalFee field to given value.

### HasLegalFee

`func (o *StandaloneInvestmentTransactionSubtype) HasLegalFee() bool`

HasLegalFee returns a boolean if a field has been set.

### GetLoanPayment

`func (o *StandaloneInvestmentTransactionSubtype) GetLoanPayment() string`

GetLoanPayment returns the LoanPayment field if non-nil, zero value otherwise.

### GetLoanPaymentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetLoanPaymentOk() (*string, bool)`

GetLoanPaymentOk returns a tuple with the LoanPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoanPayment

`func (o *StandaloneInvestmentTransactionSubtype) SetLoanPayment(v string)`

SetLoanPayment sets LoanPayment field to given value.

### HasLoanPayment

`func (o *StandaloneInvestmentTransactionSubtype) HasLoanPayment() bool`

HasLoanPayment returns a boolean if a field has been set.

### GetLongTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) GetLongTermCapitalGain() string`

GetLongTermCapitalGain returns the LongTermCapitalGain field if non-nil, zero value otherwise.

### GetLongTermCapitalGainOk

`func (o *StandaloneInvestmentTransactionSubtype) GetLongTermCapitalGainOk() (*string, bool)`

GetLongTermCapitalGainOk returns a tuple with the LongTermCapitalGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) SetLongTermCapitalGain(v string)`

SetLongTermCapitalGain sets LongTermCapitalGain field to given value.

### HasLongTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) HasLongTermCapitalGain() bool`

HasLongTermCapitalGain returns a boolean if a field has been set.

### GetLongTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) GetLongTermCapitalGainReinvestment() string`

GetLongTermCapitalGainReinvestment returns the LongTermCapitalGainReinvestment field if non-nil, zero value otherwise.

### GetLongTermCapitalGainReinvestmentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetLongTermCapitalGainReinvestmentOk() (*string, bool)`

GetLongTermCapitalGainReinvestmentOk returns a tuple with the LongTermCapitalGainReinvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) SetLongTermCapitalGainReinvestment(v string)`

SetLongTermCapitalGainReinvestment sets LongTermCapitalGainReinvestment field to given value.

### HasLongTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) HasLongTermCapitalGainReinvestment() bool`

HasLongTermCapitalGainReinvestment returns a boolean if a field has been set.

### GetManagementFee

`func (o *StandaloneInvestmentTransactionSubtype) GetManagementFee() string`

GetManagementFee returns the ManagementFee field if non-nil, zero value otherwise.

### GetManagementFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetManagementFeeOk() (*string, bool)`

GetManagementFeeOk returns a tuple with the ManagementFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementFee

`func (o *StandaloneInvestmentTransactionSubtype) SetManagementFee(v string)`

SetManagementFee sets ManagementFee field to given value.

### HasManagementFee

`func (o *StandaloneInvestmentTransactionSubtype) HasManagementFee() bool`

HasManagementFee returns a boolean if a field has been set.

### GetMarginExpense

`func (o *StandaloneInvestmentTransactionSubtype) GetMarginExpense() string`

GetMarginExpense returns the MarginExpense field if non-nil, zero value otherwise.

### GetMarginExpenseOk

`func (o *StandaloneInvestmentTransactionSubtype) GetMarginExpenseOk() (*string, bool)`

GetMarginExpenseOk returns a tuple with the MarginExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginExpense

`func (o *StandaloneInvestmentTransactionSubtype) SetMarginExpense(v string)`

SetMarginExpense sets MarginExpense field to given value.

### HasMarginExpense

`func (o *StandaloneInvestmentTransactionSubtype) HasMarginExpense() bool`

HasMarginExpense returns a boolean if a field has been set.

### GetMerger

`func (o *StandaloneInvestmentTransactionSubtype) GetMerger() string`

GetMerger returns the Merger field if non-nil, zero value otherwise.

### GetMergerOk

`func (o *StandaloneInvestmentTransactionSubtype) GetMergerOk() (*string, bool)`

GetMergerOk returns a tuple with the Merger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerger

`func (o *StandaloneInvestmentTransactionSubtype) SetMerger(v string)`

SetMerger sets Merger field to given value.

### HasMerger

`func (o *StandaloneInvestmentTransactionSubtype) HasMerger() bool`

HasMerger returns a boolean if a field has been set.

### GetMiscellaneousFee

`func (o *StandaloneInvestmentTransactionSubtype) GetMiscellaneousFee() string`

GetMiscellaneousFee returns the MiscellaneousFee field if non-nil, zero value otherwise.

### GetMiscellaneousFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetMiscellaneousFeeOk() (*string, bool)`

GetMiscellaneousFeeOk returns a tuple with the MiscellaneousFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiscellaneousFee

`func (o *StandaloneInvestmentTransactionSubtype) SetMiscellaneousFee(v string)`

SetMiscellaneousFee sets MiscellaneousFee field to given value.

### HasMiscellaneousFee

`func (o *StandaloneInvestmentTransactionSubtype) HasMiscellaneousFee() bool`

HasMiscellaneousFee returns a boolean if a field has been set.

### GetNonQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) GetNonQualifiedDividend() string`

GetNonQualifiedDividend returns the NonQualifiedDividend field if non-nil, zero value otherwise.

### GetNonQualifiedDividendOk

`func (o *StandaloneInvestmentTransactionSubtype) GetNonQualifiedDividendOk() (*string, bool)`

GetNonQualifiedDividendOk returns a tuple with the NonQualifiedDividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) SetNonQualifiedDividend(v string)`

SetNonQualifiedDividend sets NonQualifiedDividend field to given value.

### HasNonQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) HasNonQualifiedDividend() bool`

HasNonQualifiedDividend returns a boolean if a field has been set.

### GetNonResidentTax

`func (o *StandaloneInvestmentTransactionSubtype) GetNonResidentTax() string`

GetNonResidentTax returns the NonResidentTax field if non-nil, zero value otherwise.

### GetNonResidentTaxOk

`func (o *StandaloneInvestmentTransactionSubtype) GetNonResidentTaxOk() (*string, bool)`

GetNonResidentTaxOk returns a tuple with the NonResidentTax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResidentTax

`func (o *StandaloneInvestmentTransactionSubtype) SetNonResidentTax(v string)`

SetNonResidentTax sets NonResidentTax field to given value.

### HasNonResidentTax

`func (o *StandaloneInvestmentTransactionSubtype) HasNonResidentTax() bool`

HasNonResidentTax returns a boolean if a field has been set.

### GetPendingCredit

`func (o *StandaloneInvestmentTransactionSubtype) GetPendingCredit() string`

GetPendingCredit returns the PendingCredit field if non-nil, zero value otherwise.

### GetPendingCreditOk

`func (o *StandaloneInvestmentTransactionSubtype) GetPendingCreditOk() (*string, bool)`

GetPendingCreditOk returns a tuple with the PendingCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCredit

`func (o *StandaloneInvestmentTransactionSubtype) SetPendingCredit(v string)`

SetPendingCredit sets PendingCredit field to given value.

### HasPendingCredit

`func (o *StandaloneInvestmentTransactionSubtype) HasPendingCredit() bool`

HasPendingCredit returns a boolean if a field has been set.

### GetPendingDebit

`func (o *StandaloneInvestmentTransactionSubtype) GetPendingDebit() string`

GetPendingDebit returns the PendingDebit field if non-nil, zero value otherwise.

### GetPendingDebitOk

`func (o *StandaloneInvestmentTransactionSubtype) GetPendingDebitOk() (*string, bool)`

GetPendingDebitOk returns a tuple with the PendingDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDebit

`func (o *StandaloneInvestmentTransactionSubtype) SetPendingDebit(v string)`

SetPendingDebit sets PendingDebit field to given value.

### HasPendingDebit

`func (o *StandaloneInvestmentTransactionSubtype) HasPendingDebit() bool`

HasPendingDebit returns a boolean if a field has been set.

### GetQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) GetQualifiedDividend() string`

GetQualifiedDividend returns the QualifiedDividend field if non-nil, zero value otherwise.

### GetQualifiedDividendOk

`func (o *StandaloneInvestmentTransactionSubtype) GetQualifiedDividendOk() (*string, bool)`

GetQualifiedDividendOk returns a tuple with the QualifiedDividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) SetQualifiedDividend(v string)`

SetQualifiedDividend sets QualifiedDividend field to given value.

### HasQualifiedDividend

`func (o *StandaloneInvestmentTransactionSubtype) HasQualifiedDividend() bool`

HasQualifiedDividend returns a boolean if a field has been set.

### GetRebalance

`func (o *StandaloneInvestmentTransactionSubtype) GetRebalance() string`

GetRebalance returns the Rebalance field if non-nil, zero value otherwise.

### GetRebalanceOk

`func (o *StandaloneInvestmentTransactionSubtype) GetRebalanceOk() (*string, bool)`

GetRebalanceOk returns a tuple with the Rebalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebalance

`func (o *StandaloneInvestmentTransactionSubtype) SetRebalance(v string)`

SetRebalance sets Rebalance field to given value.

### HasRebalance

`func (o *StandaloneInvestmentTransactionSubtype) HasRebalance() bool`

HasRebalance returns a boolean if a field has been set.

### GetReturnOfPrincipal

`func (o *StandaloneInvestmentTransactionSubtype) GetReturnOfPrincipal() string`

GetReturnOfPrincipal returns the ReturnOfPrincipal field if non-nil, zero value otherwise.

### GetReturnOfPrincipalOk

`func (o *StandaloneInvestmentTransactionSubtype) GetReturnOfPrincipalOk() (*string, bool)`

GetReturnOfPrincipalOk returns a tuple with the ReturnOfPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnOfPrincipal

`func (o *StandaloneInvestmentTransactionSubtype) SetReturnOfPrincipal(v string)`

SetReturnOfPrincipal sets ReturnOfPrincipal field to given value.

### HasReturnOfPrincipal

`func (o *StandaloneInvestmentTransactionSubtype) HasReturnOfPrincipal() bool`

HasReturnOfPrincipal returns a boolean if a field has been set.

### GetSell

`func (o *StandaloneInvestmentTransactionSubtype) GetSell() string`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *StandaloneInvestmentTransactionSubtype) GetSellOk() (*string, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *StandaloneInvestmentTransactionSubtype) SetSell(v string)`

SetSell sets Sell field to given value.

### HasSell

`func (o *StandaloneInvestmentTransactionSubtype) HasSell() bool`

HasSell returns a boolean if a field has been set.

### GetSellShort

`func (o *StandaloneInvestmentTransactionSubtype) GetSellShort() string`

GetSellShort returns the SellShort field if non-nil, zero value otherwise.

### GetSellShortOk

`func (o *StandaloneInvestmentTransactionSubtype) GetSellShortOk() (*string, bool)`

GetSellShortOk returns a tuple with the SellShort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellShort

`func (o *StandaloneInvestmentTransactionSubtype) SetSellShort(v string)`

SetSellShort sets SellShort field to given value.

### HasSellShort

`func (o *StandaloneInvestmentTransactionSubtype) HasSellShort() bool`

HasSellShort returns a boolean if a field has been set.

### GetShortTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) GetShortTermCapitalGain() string`

GetShortTermCapitalGain returns the ShortTermCapitalGain field if non-nil, zero value otherwise.

### GetShortTermCapitalGainOk

`func (o *StandaloneInvestmentTransactionSubtype) GetShortTermCapitalGainOk() (*string, bool)`

GetShortTermCapitalGainOk returns a tuple with the ShortTermCapitalGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) SetShortTermCapitalGain(v string)`

SetShortTermCapitalGain sets ShortTermCapitalGain field to given value.

### HasShortTermCapitalGain

`func (o *StandaloneInvestmentTransactionSubtype) HasShortTermCapitalGain() bool`

HasShortTermCapitalGain returns a boolean if a field has been set.

### GetShortTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) GetShortTermCapitalGainReinvestment() string`

GetShortTermCapitalGainReinvestment returns the ShortTermCapitalGainReinvestment field if non-nil, zero value otherwise.

### GetShortTermCapitalGainReinvestmentOk

`func (o *StandaloneInvestmentTransactionSubtype) GetShortTermCapitalGainReinvestmentOk() (*string, bool)`

GetShortTermCapitalGainReinvestmentOk returns a tuple with the ShortTermCapitalGainReinvestment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) SetShortTermCapitalGainReinvestment(v string)`

SetShortTermCapitalGainReinvestment sets ShortTermCapitalGainReinvestment field to given value.

### HasShortTermCapitalGainReinvestment

`func (o *StandaloneInvestmentTransactionSubtype) HasShortTermCapitalGainReinvestment() bool`

HasShortTermCapitalGainReinvestment returns a boolean if a field has been set.

### GetSpinOff

`func (o *StandaloneInvestmentTransactionSubtype) GetSpinOff() string`

GetSpinOff returns the SpinOff field if non-nil, zero value otherwise.

### GetSpinOffOk

`func (o *StandaloneInvestmentTransactionSubtype) GetSpinOffOk() (*string, bool)`

GetSpinOffOk returns a tuple with the SpinOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpinOff

`func (o *StandaloneInvestmentTransactionSubtype) SetSpinOff(v string)`

SetSpinOff sets SpinOff field to given value.

### HasSpinOff

`func (o *StandaloneInvestmentTransactionSubtype) HasSpinOff() bool`

HasSpinOff returns a boolean if a field has been set.

### GetSplit

`func (o *StandaloneInvestmentTransactionSubtype) GetSplit() string`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *StandaloneInvestmentTransactionSubtype) GetSplitOk() (*string, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *StandaloneInvestmentTransactionSubtype) SetSplit(v string)`

SetSplit sets Split field to given value.

### HasSplit

`func (o *StandaloneInvestmentTransactionSubtype) HasSplit() bool`

HasSplit returns a boolean if a field has been set.

### GetStockDistribution

`func (o *StandaloneInvestmentTransactionSubtype) GetStockDistribution() string`

GetStockDistribution returns the StockDistribution field if non-nil, zero value otherwise.

### GetStockDistributionOk

`func (o *StandaloneInvestmentTransactionSubtype) GetStockDistributionOk() (*string, bool)`

GetStockDistributionOk returns a tuple with the StockDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockDistribution

`func (o *StandaloneInvestmentTransactionSubtype) SetStockDistribution(v string)`

SetStockDistribution sets StockDistribution field to given value.

### HasStockDistribution

`func (o *StandaloneInvestmentTransactionSubtype) HasStockDistribution() bool`

HasStockDistribution returns a boolean if a field has been set.

### GetTax

`func (o *StandaloneInvestmentTransactionSubtype) GetTax() string`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *StandaloneInvestmentTransactionSubtype) GetTaxOk() (*string, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *StandaloneInvestmentTransactionSubtype) SetTax(v string)`

SetTax sets Tax field to given value.

### HasTax

`func (o *StandaloneInvestmentTransactionSubtype) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxWithheld

`func (o *StandaloneInvestmentTransactionSubtype) GetTaxWithheld() string`

GetTaxWithheld returns the TaxWithheld field if non-nil, zero value otherwise.

### GetTaxWithheldOk

`func (o *StandaloneInvestmentTransactionSubtype) GetTaxWithheldOk() (*string, bool)`

GetTaxWithheldOk returns a tuple with the TaxWithheld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxWithheld

`func (o *StandaloneInvestmentTransactionSubtype) SetTaxWithheld(v string)`

SetTaxWithheld sets TaxWithheld field to given value.

### HasTaxWithheld

`func (o *StandaloneInvestmentTransactionSubtype) HasTaxWithheld() bool`

HasTaxWithheld returns a boolean if a field has been set.

### GetTransfer

`func (o *StandaloneInvestmentTransactionSubtype) GetTransfer() string`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *StandaloneInvestmentTransactionSubtype) GetTransferOk() (*string, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *StandaloneInvestmentTransactionSubtype) SetTransfer(v string)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *StandaloneInvestmentTransactionSubtype) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetTransferFee

`func (o *StandaloneInvestmentTransactionSubtype) GetTransferFee() string`

GetTransferFee returns the TransferFee field if non-nil, zero value otherwise.

### GetTransferFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetTransferFeeOk() (*string, bool)`

GetTransferFeeOk returns a tuple with the TransferFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferFee

`func (o *StandaloneInvestmentTransactionSubtype) SetTransferFee(v string)`

SetTransferFee sets TransferFee field to given value.

### HasTransferFee

`func (o *StandaloneInvestmentTransactionSubtype) HasTransferFee() bool`

HasTransferFee returns a boolean if a field has been set.

### GetTrustFee

`func (o *StandaloneInvestmentTransactionSubtype) GetTrustFee() string`

GetTrustFee returns the TrustFee field if non-nil, zero value otherwise.

### GetTrustFeeOk

`func (o *StandaloneInvestmentTransactionSubtype) GetTrustFeeOk() (*string, bool)`

GetTrustFeeOk returns a tuple with the TrustFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustFee

`func (o *StandaloneInvestmentTransactionSubtype) SetTrustFee(v string)`

SetTrustFee sets TrustFee field to given value.

### HasTrustFee

`func (o *StandaloneInvestmentTransactionSubtype) HasTrustFee() bool`

HasTrustFee returns a boolean if a field has been set.

### GetUnqualifiedGain

`func (o *StandaloneInvestmentTransactionSubtype) GetUnqualifiedGain() string`

GetUnqualifiedGain returns the UnqualifiedGain field if non-nil, zero value otherwise.

### GetUnqualifiedGainOk

`func (o *StandaloneInvestmentTransactionSubtype) GetUnqualifiedGainOk() (*string, bool)`

GetUnqualifiedGainOk returns a tuple with the UnqualifiedGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnqualifiedGain

`func (o *StandaloneInvestmentTransactionSubtype) SetUnqualifiedGain(v string)`

SetUnqualifiedGain sets UnqualifiedGain field to given value.

### HasUnqualifiedGain

`func (o *StandaloneInvestmentTransactionSubtype) HasUnqualifiedGain() bool`

HasUnqualifiedGain returns a boolean if a field has been set.

### GetWithdrawal

`func (o *StandaloneInvestmentTransactionSubtype) GetWithdrawal() string`

GetWithdrawal returns the Withdrawal field if non-nil, zero value otherwise.

### GetWithdrawalOk

`func (o *StandaloneInvestmentTransactionSubtype) GetWithdrawalOk() (*string, bool)`

GetWithdrawalOk returns a tuple with the Withdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawal

`func (o *StandaloneInvestmentTransactionSubtype) SetWithdrawal(v string)`

SetWithdrawal sets Withdrawal field to given value.

### HasWithdrawal

`func (o *StandaloneInvestmentTransactionSubtype) HasWithdrawal() bool`

HasWithdrawal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


