# BankTransferUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegalName** | **string** | The account holder’s full legal name. | 
**EmailAddress** | Pointer to **NullableString** | The account holder’s email. | [optional] 
**RoutingNumber** | Pointer to **string** | The account holder&#39;s routing number. This field is only used for response data. Do not provide this field when making requests. | [optional] [readonly] 

## Methods

### NewBankTransferUser

`func NewBankTransferUser(legalName string, ) *BankTransferUser`

NewBankTransferUser instantiates a new BankTransferUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferUserWithDefaults

`func NewBankTransferUserWithDefaults() *BankTransferUser`

NewBankTransferUserWithDefaults instantiates a new BankTransferUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegalName

`func (o *BankTransferUser) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *BankTransferUser) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *BankTransferUser) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.


### GetEmailAddress

`func (o *BankTransferUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *BankTransferUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *BankTransferUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *BankTransferUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *BankTransferUser) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *BankTransferUser) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetRoutingNumber

`func (o *BankTransferUser) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankTransferUser) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankTransferUser) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.

### HasRoutingNumber

`func (o *BankTransferUser) HasRoutingNumber() bool`

HasRoutingNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


