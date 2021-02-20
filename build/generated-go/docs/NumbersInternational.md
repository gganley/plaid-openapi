# NumbersInternational

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid account ID associated with the account numbers | 
**Iban** | **string** | The International Bank Account Number (IBAN) for the account | 
**Bic** | **string** | The Bank Identifier Code (BIC) for the account | 

## Methods

### NewNumbersInternational

`func NewNumbersInternational(accountId string, iban string, bic string, ) *NumbersInternational`

NewNumbersInternational instantiates a new NumbersInternational object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersInternationalWithDefaults

`func NewNumbersInternationalWithDefaults() *NumbersInternational`

NewNumbersInternationalWithDefaults instantiates a new NumbersInternational object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NumbersInternational) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NumbersInternational) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NumbersInternational) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetIban

`func (o *NumbersInternational) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *NumbersInternational) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *NumbersInternational) SetIban(v string)`

SetIban sets Iban field to given value.


### GetBic

`func (o *NumbersInternational) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *NumbersInternational) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *NumbersInternational) SetBic(v string)`

SetBic sets Bic field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


