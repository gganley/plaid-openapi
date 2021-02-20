# SandboxPublicTokenCreateRequestOptionsTransactions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **string** | The earliest date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD. | [optional] 
**EndDate** | Pointer to **string** | The most recent date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD. | [optional] 

## Methods

### NewSandboxPublicTokenCreateRequestOptionsTransactions

`func NewSandboxPublicTokenCreateRequestOptionsTransactions() *SandboxPublicTokenCreateRequestOptionsTransactions`

NewSandboxPublicTokenCreateRequestOptionsTransactions instantiates a new SandboxPublicTokenCreateRequestOptionsTransactions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPublicTokenCreateRequestOptionsTransactionsWithDefaults

`func NewSandboxPublicTokenCreateRequestOptionsTransactionsWithDefaults() *SandboxPublicTokenCreateRequestOptionsTransactions`

NewSandboxPublicTokenCreateRequestOptionsTransactionsWithDefaults instantiates a new SandboxPublicTokenCreateRequestOptionsTransactions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *SandboxPublicTokenCreateRequestOptionsTransactions) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


