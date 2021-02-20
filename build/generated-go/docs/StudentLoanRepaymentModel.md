# StudentLoanRepaymentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The only currently supported value for this field is &#x60;standard&#x60;. | 
**NonRepaymentMonths** | **float32** | Configures the number of months before repayment starts. | 
**RepaymentMonths** | **float32** | Configures the number of months of repayments before the loan is paid off. | 

## Methods

### NewStudentLoanRepaymentModel

`func NewStudentLoanRepaymentModel(type_ string, nonRepaymentMonths float32, repaymentMonths float32, ) *StudentLoanRepaymentModel`

NewStudentLoanRepaymentModel instantiates a new StudentLoanRepaymentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentLoanRepaymentModelWithDefaults

`func NewStudentLoanRepaymentModelWithDefaults() *StudentLoanRepaymentModel`

NewStudentLoanRepaymentModelWithDefaults instantiates a new StudentLoanRepaymentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *StudentLoanRepaymentModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StudentLoanRepaymentModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StudentLoanRepaymentModel) SetType(v string)`

SetType sets Type field to given value.


### GetNonRepaymentMonths

`func (o *StudentLoanRepaymentModel) GetNonRepaymentMonths() float32`

GetNonRepaymentMonths returns the NonRepaymentMonths field if non-nil, zero value otherwise.

### GetNonRepaymentMonthsOk

`func (o *StudentLoanRepaymentModel) GetNonRepaymentMonthsOk() (*float32, bool)`

GetNonRepaymentMonthsOk returns a tuple with the NonRepaymentMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonRepaymentMonths

`func (o *StudentLoanRepaymentModel) SetNonRepaymentMonths(v float32)`

SetNonRepaymentMonths sets NonRepaymentMonths field to given value.


### GetRepaymentMonths

`func (o *StudentLoanRepaymentModel) GetRepaymentMonths() float32`

GetRepaymentMonths returns the RepaymentMonths field if non-nil, zero value otherwise.

### GetRepaymentMonthsOk

`func (o *StudentLoanRepaymentModel) GetRepaymentMonthsOk() (*float32, bool)`

GetRepaymentMonthsOk returns a tuple with the RepaymentMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentMonths

`func (o *StudentLoanRepaymentModel) SetRepaymentMonths(v float32)`

SetRepaymentMonths sets RepaymentMonths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


