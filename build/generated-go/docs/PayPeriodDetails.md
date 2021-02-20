# PayPeriodDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **NullableString** | The pay period start date, in ISO 8601 format: \&quot;yyyy-mm-dd\&quot;. | [optional] 
**EndDate** | Pointer to **NullableString** | The pay period end date, in ISO 8601 format: \&quot;yyyy-mm-dd\&quot;. | [optional] 
**PayDay** | Pointer to **NullableString** | The date on which the paystub was issued, in ISO 8601 format (\&quot;yyyy-mm-dd\&quot;). | [optional] 
**GrossEarnings** | Pointer to **NullableFloat32** | Total earnings before tax. | [optional] 
**CheckAmount** | Pointer to **NullableFloat32** | The net amount of the paycheck. | [optional] 

## Methods

### NewPayPeriodDetails

`func NewPayPeriodDetails() *PayPeriodDetails`

NewPayPeriodDetails instantiates a new PayPeriodDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayPeriodDetailsWithDefaults

`func NewPayPeriodDetailsWithDefaults() *PayPeriodDetails`

NewPayPeriodDetailsWithDefaults instantiates a new PayPeriodDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *PayPeriodDetails) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PayPeriodDetails) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PayPeriodDetails) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *PayPeriodDetails) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *PayPeriodDetails) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *PayPeriodDetails) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *PayPeriodDetails) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PayPeriodDetails) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PayPeriodDetails) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PayPeriodDetails) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *PayPeriodDetails) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *PayPeriodDetails) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPayDay

`func (o *PayPeriodDetails) GetPayDay() string`

GetPayDay returns the PayDay field if non-nil, zero value otherwise.

### GetPayDayOk

`func (o *PayPeriodDetails) GetPayDayOk() (*string, bool)`

GetPayDayOk returns a tuple with the PayDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayDay

`func (o *PayPeriodDetails) SetPayDay(v string)`

SetPayDay sets PayDay field to given value.

### HasPayDay

`func (o *PayPeriodDetails) HasPayDay() bool`

HasPayDay returns a boolean if a field has been set.

### SetPayDayNil

`func (o *PayPeriodDetails) SetPayDayNil(b bool)`

 SetPayDayNil sets the value for PayDay to be an explicit nil

### UnsetPayDay
`func (o *PayPeriodDetails) UnsetPayDay()`

UnsetPayDay ensures that no value is present for PayDay, not even an explicit nil
### GetGrossEarnings

`func (o *PayPeriodDetails) GetGrossEarnings() float32`

GetGrossEarnings returns the GrossEarnings field if non-nil, zero value otherwise.

### GetGrossEarningsOk

`func (o *PayPeriodDetails) GetGrossEarningsOk() (*float32, bool)`

GetGrossEarningsOk returns a tuple with the GrossEarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossEarnings

`func (o *PayPeriodDetails) SetGrossEarnings(v float32)`

SetGrossEarnings sets GrossEarnings field to given value.

### HasGrossEarnings

`func (o *PayPeriodDetails) HasGrossEarnings() bool`

HasGrossEarnings returns a boolean if a field has been set.

### SetGrossEarningsNil

`func (o *PayPeriodDetails) SetGrossEarningsNil(b bool)`

 SetGrossEarningsNil sets the value for GrossEarnings to be an explicit nil

### UnsetGrossEarnings
`func (o *PayPeriodDetails) UnsetGrossEarnings()`

UnsetGrossEarnings ensures that no value is present for GrossEarnings, not even an explicit nil
### GetCheckAmount

`func (o *PayPeriodDetails) GetCheckAmount() float32`

GetCheckAmount returns the CheckAmount field if non-nil, zero value otherwise.

### GetCheckAmountOk

`func (o *PayPeriodDetails) GetCheckAmountOk() (*float32, bool)`

GetCheckAmountOk returns a tuple with the CheckAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAmount

`func (o *PayPeriodDetails) SetCheckAmount(v float32)`

SetCheckAmount sets CheckAmount field to given value.

### HasCheckAmount

`func (o *PayPeriodDetails) HasCheckAmount() bool`

HasCheckAmount returns a boolean if a field has been set.

### SetCheckAmountNil

`func (o *PayPeriodDetails) SetCheckAmountNil(b bool)`

 SetCheckAmountNil sets the value for CheckAmount to be an explicit nil

### UnsetCheckAmount
`func (o *PayPeriodDetails) UnsetCheckAmount()`

UnsetCheckAmount ensures that no value is present for CheckAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


