# PSLFStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EstimatedEligibilityDate** | Pointer to **NullableString** | The estimated date borrower will have completed 120 qualifying monthly payments. Returned in ISO 8601 format (YYYY-MM-DD). | [optional] 
**PaymentsMade** | Pointer to **NullableFloat32** | The number of qualifying payments that have been made. | [optional] 
**PaymentsRemaining** | Pointer to **NullableFloat32** | The number of qualifying payments remaining. | [optional] 

## Methods

### NewPSLFStatus

`func NewPSLFStatus() *PSLFStatus`

NewPSLFStatus instantiates a new PSLFStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPSLFStatusWithDefaults

`func NewPSLFStatusWithDefaults() *PSLFStatus`

NewPSLFStatusWithDefaults instantiates a new PSLFStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimatedEligibilityDate

`func (o *PSLFStatus) GetEstimatedEligibilityDate() string`

GetEstimatedEligibilityDate returns the EstimatedEligibilityDate field if non-nil, zero value otherwise.

### GetEstimatedEligibilityDateOk

`func (o *PSLFStatus) GetEstimatedEligibilityDateOk() (*string, bool)`

GetEstimatedEligibilityDateOk returns a tuple with the EstimatedEligibilityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEligibilityDate

`func (o *PSLFStatus) SetEstimatedEligibilityDate(v string)`

SetEstimatedEligibilityDate sets EstimatedEligibilityDate field to given value.

### HasEstimatedEligibilityDate

`func (o *PSLFStatus) HasEstimatedEligibilityDate() bool`

HasEstimatedEligibilityDate returns a boolean if a field has been set.

### SetEstimatedEligibilityDateNil

`func (o *PSLFStatus) SetEstimatedEligibilityDateNil(b bool)`

 SetEstimatedEligibilityDateNil sets the value for EstimatedEligibilityDate to be an explicit nil

### UnsetEstimatedEligibilityDate
`func (o *PSLFStatus) UnsetEstimatedEligibilityDate()`

UnsetEstimatedEligibilityDate ensures that no value is present for EstimatedEligibilityDate, not even an explicit nil
### GetPaymentsMade

`func (o *PSLFStatus) GetPaymentsMade() float32`

GetPaymentsMade returns the PaymentsMade field if non-nil, zero value otherwise.

### GetPaymentsMadeOk

`func (o *PSLFStatus) GetPaymentsMadeOk() (*float32, bool)`

GetPaymentsMadeOk returns a tuple with the PaymentsMade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsMade

`func (o *PSLFStatus) SetPaymentsMade(v float32)`

SetPaymentsMade sets PaymentsMade field to given value.

### HasPaymentsMade

`func (o *PSLFStatus) HasPaymentsMade() bool`

HasPaymentsMade returns a boolean if a field has been set.

### SetPaymentsMadeNil

`func (o *PSLFStatus) SetPaymentsMadeNil(b bool)`

 SetPaymentsMadeNil sets the value for PaymentsMade to be an explicit nil

### UnsetPaymentsMade
`func (o *PSLFStatus) UnsetPaymentsMade()`

UnsetPaymentsMade ensures that no value is present for PaymentsMade, not even an explicit nil
### GetPaymentsRemaining

`func (o *PSLFStatus) GetPaymentsRemaining() float32`

GetPaymentsRemaining returns the PaymentsRemaining field if non-nil, zero value otherwise.

### GetPaymentsRemainingOk

`func (o *PSLFStatus) GetPaymentsRemainingOk() (*float32, bool)`

GetPaymentsRemainingOk returns a tuple with the PaymentsRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsRemaining

`func (o *PSLFStatus) SetPaymentsRemaining(v float32)`

SetPaymentsRemaining sets PaymentsRemaining field to given value.

### HasPaymentsRemaining

`func (o *PSLFStatus) HasPaymentsRemaining() bool`

HasPaymentsRemaining returns a boolean if a field has been set.

### SetPaymentsRemainingNil

`func (o *PSLFStatus) SetPaymentsRemainingNil(b bool)`

 SetPaymentsRemainingNil sets the value for PaymentsRemaining to be an explicit nil

### UnsetPaymentsRemaining
`func (o *PSLFStatus) UnsetPaymentsRemaining()`

UnsetPaymentsRemaining ensures that no value is present for PaymentsRemaining, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


