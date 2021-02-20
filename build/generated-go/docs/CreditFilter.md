# CreditFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSubtypes** | [**[]AccountSubtype**](AccountSubtype.md) | An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](/docs/api/accounts#accounts-schema).  | 

## Methods

### NewCreditFilter

`func NewCreditFilter(accountSubtypes []AccountSubtype, ) *CreditFilter`

NewCreditFilter instantiates a new CreditFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditFilterWithDefaults

`func NewCreditFilterWithDefaults() *CreditFilter`

NewCreditFilterWithDefaults instantiates a new CreditFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSubtypes

`func (o *CreditFilter) GetAccountSubtypes() []AccountSubtype`

GetAccountSubtypes returns the AccountSubtypes field if non-nil, zero value otherwise.

### GetAccountSubtypesOk

`func (o *CreditFilter) GetAccountSubtypesOk() (*[]AccountSubtype, bool)`

GetAccountSubtypesOk returns a tuple with the AccountSubtypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtypes

`func (o *CreditFilter) SetAccountSubtypes(v []AccountSubtype)`

SetAccountSubtypes sets AccountSubtypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


