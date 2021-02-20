# StudentLoanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **NullableString** | The date until which the loan will be in its current status. Dates are returned in an ISO 8601 format (YYYY-MM-DD).  | [optional] 
**Type** | Pointer to **NullableString** | The status type of the student loan | [optional] 

## Methods

### NewStudentLoanStatus

`func NewStudentLoanStatus() *StudentLoanStatus`

NewStudentLoanStatus instantiates a new StudentLoanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStudentLoanStatusWithDefaults

`func NewStudentLoanStatusWithDefaults() *StudentLoanStatus`

NewStudentLoanStatusWithDefaults instantiates a new StudentLoanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *StudentLoanStatus) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *StudentLoanStatus) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *StudentLoanStatus) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *StudentLoanStatus) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *StudentLoanStatus) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *StudentLoanStatus) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetType

`func (o *StudentLoanStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StudentLoanStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StudentLoanStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StudentLoanStatus) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *StudentLoanStatus) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *StudentLoanStatus) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


