# CreditCardLiability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** | The ID of the account that this liability belongs to. | [optional] 
**Aprs** | [**[]APR**](APR.md) | The various interest rates that apply to the account. | 
**IsOverdue** | Pointer to **NullableBool** | true if a payment is currently overdue. Availability for this field is limited. | [optional] 
**LastPaymentAmount** | **float32** | The amount of the last payment. | 
**LastPaymentDate** | **string** | The date of the last payment. Dates are returned in an ISO 8601 format (YYYY-MM-DD). Availability for this field is limited. | 
**LastStatementBalance** | **float32** | The outstanding balance on the last statement. Availability for this field is limited. | 
**LastStatementIssueDate** | **string** | The date of the last statement. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | 
**MinimumPaymentAmount** | **float32** | The minimum payment due for the next billing cycle. | 
**NextPaymentDueDate** | **string** | The due date for the next payment. The due date is &#x60;null&#x60; if a payment is not expected. Dates are returned in an ISO 8601 format (YYYY-MM-DD). | 

## Methods

### NewCreditCardLiability

`func NewCreditCardLiability(aprs []APR, lastPaymentAmount float32, lastPaymentDate string, lastStatementBalance float32, lastStatementIssueDate string, minimumPaymentAmount float32, nextPaymentDueDate string, ) *CreditCardLiability`

NewCreditCardLiability instantiates a new CreditCardLiability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardLiabilityWithDefaults

`func NewCreditCardLiabilityWithDefaults() *CreditCardLiability`

NewCreditCardLiabilityWithDefaults instantiates a new CreditCardLiability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreditCardLiability) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreditCardLiability) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreditCardLiability) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreditCardLiability) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *CreditCardLiability) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *CreditCardLiability) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAprs

`func (o *CreditCardLiability) GetAprs() []APR`

GetAprs returns the Aprs field if non-nil, zero value otherwise.

### GetAprsOk

`func (o *CreditCardLiability) GetAprsOk() (*[]APR, bool)`

GetAprsOk returns a tuple with the Aprs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAprs

`func (o *CreditCardLiability) SetAprs(v []APR)`

SetAprs sets Aprs field to given value.


### GetIsOverdue

`func (o *CreditCardLiability) GetIsOverdue() bool`

GetIsOverdue returns the IsOverdue field if non-nil, zero value otherwise.

### GetIsOverdueOk

`func (o *CreditCardLiability) GetIsOverdueOk() (*bool, bool)`

GetIsOverdueOk returns a tuple with the IsOverdue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverdue

`func (o *CreditCardLiability) SetIsOverdue(v bool)`

SetIsOverdue sets IsOverdue field to given value.

### HasIsOverdue

`func (o *CreditCardLiability) HasIsOverdue() bool`

HasIsOverdue returns a boolean if a field has been set.

### SetIsOverdueNil

`func (o *CreditCardLiability) SetIsOverdueNil(b bool)`

 SetIsOverdueNil sets the value for IsOverdue to be an explicit nil

### UnsetIsOverdue
`func (o *CreditCardLiability) UnsetIsOverdue()`

UnsetIsOverdue ensures that no value is present for IsOverdue, not even an explicit nil
### GetLastPaymentAmount

`func (o *CreditCardLiability) GetLastPaymentAmount() float32`

GetLastPaymentAmount returns the LastPaymentAmount field if non-nil, zero value otherwise.

### GetLastPaymentAmountOk

`func (o *CreditCardLiability) GetLastPaymentAmountOk() (*float32, bool)`

GetLastPaymentAmountOk returns a tuple with the LastPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentAmount

`func (o *CreditCardLiability) SetLastPaymentAmount(v float32)`

SetLastPaymentAmount sets LastPaymentAmount field to given value.


### GetLastPaymentDate

`func (o *CreditCardLiability) GetLastPaymentDate() string`

GetLastPaymentDate returns the LastPaymentDate field if non-nil, zero value otherwise.

### GetLastPaymentDateOk

`func (o *CreditCardLiability) GetLastPaymentDateOk() (*string, bool)`

GetLastPaymentDateOk returns a tuple with the LastPaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentDate

`func (o *CreditCardLiability) SetLastPaymentDate(v string)`

SetLastPaymentDate sets LastPaymentDate field to given value.


### GetLastStatementBalance

`func (o *CreditCardLiability) GetLastStatementBalance() float32`

GetLastStatementBalance returns the LastStatementBalance field if non-nil, zero value otherwise.

### GetLastStatementBalanceOk

`func (o *CreditCardLiability) GetLastStatementBalanceOk() (*float32, bool)`

GetLastStatementBalanceOk returns a tuple with the LastStatementBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementBalance

`func (o *CreditCardLiability) SetLastStatementBalance(v float32)`

SetLastStatementBalance sets LastStatementBalance field to given value.


### GetLastStatementIssueDate

`func (o *CreditCardLiability) GetLastStatementIssueDate() string`

GetLastStatementIssueDate returns the LastStatementIssueDate field if non-nil, zero value otherwise.

### GetLastStatementIssueDateOk

`func (o *CreditCardLiability) GetLastStatementIssueDateOk() (*string, bool)`

GetLastStatementIssueDateOk returns a tuple with the LastStatementIssueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatementIssueDate

`func (o *CreditCardLiability) SetLastStatementIssueDate(v string)`

SetLastStatementIssueDate sets LastStatementIssueDate field to given value.


### GetMinimumPaymentAmount

`func (o *CreditCardLiability) GetMinimumPaymentAmount() float32`

GetMinimumPaymentAmount returns the MinimumPaymentAmount field if non-nil, zero value otherwise.

### GetMinimumPaymentAmountOk

`func (o *CreditCardLiability) GetMinimumPaymentAmountOk() (*float32, bool)`

GetMinimumPaymentAmountOk returns a tuple with the MinimumPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPaymentAmount

`func (o *CreditCardLiability) SetMinimumPaymentAmount(v float32)`

SetMinimumPaymentAmount sets MinimumPaymentAmount field to given value.


### GetNextPaymentDueDate

`func (o *CreditCardLiability) GetNextPaymentDueDate() string`

GetNextPaymentDueDate returns the NextPaymentDueDate field if non-nil, zero value otherwise.

### GetNextPaymentDueDateOk

`func (o *CreditCardLiability) GetNextPaymentDueDateOk() (*string, bool)`

GetNextPaymentDueDateOk returns a tuple with the NextPaymentDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentDueDate

`func (o *CreditCardLiability) SetNextPaymentDueDate(v string)`

SetNextPaymentDueDate sets NextPaymentDueDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


