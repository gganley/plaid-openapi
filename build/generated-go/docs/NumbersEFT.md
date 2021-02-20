# NumbersEFT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid account ID associated with the account numbers | 
**Account** | **string** | The EFT account number for the account | 
**Institution** | **string** | The EFT institution number for the account | 
**Branch** | **string** | The EFT branch number for the account | 

## Methods

### NewNumbersEFT

`func NewNumbersEFT(accountId string, account string, institution string, branch string, ) *NumbersEFT`

NewNumbersEFT instantiates a new NumbersEFT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersEFTWithDefaults

`func NewNumbersEFTWithDefaults() *NumbersEFT`

NewNumbersEFTWithDefaults instantiates a new NumbersEFT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NumbersEFT) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NumbersEFT) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NumbersEFT) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccount

`func (o *NumbersEFT) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NumbersEFT) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NumbersEFT) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetInstitution

`func (o *NumbersEFT) GetInstitution() string`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *NumbersEFT) GetInstitutionOk() (*string, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *NumbersEFT) SetInstitution(v string)`

SetInstitution sets Institution field to given value.


### GetBranch

`func (o *NumbersEFT) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *NumbersEFT) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *NumbersEFT) SetBranch(v string)`

SetBranch sets Branch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


