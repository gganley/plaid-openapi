# LoanFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtypes** | [**[]AccountSubtype**](AccountSubtype.md) | An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](/docs/api/accounts#accounts-schema).  | 

## Methods

### NewLoanFilter

`func NewLoanFilter(accountSubtypes []AccountSubtype, ) *LoanFilter`

NewLoanFilter instantiates a new LoanFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoanFilterWithDefaults

`func NewLoanFilterWithDefaults() *LoanFilter`

NewLoanFilterWithDefaults instantiates a new LoanFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtypes

`func (o *LoanFilter) GetAccountSubtypes() []AccountSubtype`

GetAccountSubtypes returns the AccountSubtypes field if non-nil, zero value otherwise.

### GetAccountSubtypesOk

`func (o *LoanFilter) GetAccountSubtypesOk() (*[]AccountSubtype, bool)`

GetAccountSubtypesOk returns a tuple with the AccountSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypes

`func (o *LoanFilter) SetAccountSubtypes(v []AccountSubtype)`

SetAccountSubtypes sets AccountSubtypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


