# AccountIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Plaid’s unique identifier for the account. This value will not change unless Plaid can&#39;t reconcile the account with the data returned by the financial institution. This may occur, for example, when the name of the account changes. If this happens a new &#x60;account_id&#x60; will be assigned to the account.  The &#x60;account_id&#x60; can also change if the &#x60;access_token&#x60; is deleted and the same credentials that were used to generate that &#x60;access_token&#x60; are used to generate a new &#x60;access_token&#x60; on a later date. In that case, the new &#x60;account_id&#x60; will be different from the old &#x60;account_id&#x60;.  Like all Plaid identifiers, the &#x60;account_id&#x60; is case sensitive. | 
**Balances** | [**AccountBalance**](AccountBalance.md) |  | 
**Mask** | Pointer to **NullableString** | The last 2-4 alphanumeric characters of an account&#39;s official account number. Note that the mask may be non-unique between an Item&#39;s accounts, and it may also not match the mask that the bank displays to the user. | [optional] 
**Name** | **string** | The name of the account, either assigned by the user or by the financial institution itself | 
**OfficialName** | Pointer to **NullableString** | The official name of the account as given by the financial institution | [optional] 
**Type** | [**AccountType**](AccountType.md) |  | 
**Subtype** | [**NullableAccountSubtype**](AccountSubtype.md) |  | 
**VerificationStatus** | Pointer to **NullableString** | The current verification status of an Auth Item initiated through Automated or Manual micro-deposits.  Returned for Auth Items only.  &#x60;pending_automatic_verification&#x60;: The Item is pending automatic verification  &#x60;pending_manual_verification&#x60;: The Item is pending manual micro-deposit verification. Items remain in this state until the user successfully verifies the two amounts.  &#x60;automatically_verified&#x60;: The Item has successfully been automatically verified   &#x60;manually_verified&#x60;: The Item has successfully been manually verified  &#x60;verification_expired&#x60;: Plaid was unable to automatically verify the deposit within 7 calendar days and will no longer attempt to validate the Item. Users may retry by submitting their information again through Link.  &#x60;verification_failed&#x60;: The Item failed manual micro-deposit verification because the user exhausted all 3 verification attempts. Users may retry by submitting their information again through Link.    | [optional] 
**Owners** | [**[]Owner**](Owner.md) | Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. Multiple owners on a single account will be represented in the same &#x60;owner&#x60; object, not in multiple owner objects within the array. | 

## Methods

### NewAccountIdentity

`func NewAccountIdentity(accountId string, balances AccountBalance, name string, type_ AccountType, subtype NullableAccountSubtype, owners []Owner, ) *AccountIdentity`

NewAccountIdentity instantiates a new AccountIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdentityWithDefaults

`func NewAccountIdentityWithDefaults() *AccountIdentity`

NewAccountIdentityWithDefaults instantiates a new AccountIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountIdentity) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountIdentity) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountIdentity) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBalances

`func (o *AccountIdentity) GetBalances() AccountBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountIdentity) GetBalancesOk() (*AccountBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountIdentity) SetBalances(v AccountBalance)`

SetBalances sets Balances field to given value.


### GetMask

`func (o *AccountIdentity) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *AccountIdentity) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *AccountIdentity) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *AccountIdentity) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *AccountIdentity) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *AccountIdentity) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetName

`func (o *AccountIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountIdentity) SetName(v string)`

SetName sets Name field to given value.


### GetOfficialName

`func (o *AccountIdentity) GetOfficialName() string`

GetOfficialName returns the OfficialName field if non-nil, zero value otherwise.

### GetOfficialNameOk

`func (o *AccountIdentity) GetOfficialNameOk() (*string, bool)`

GetOfficialNameOk returns a tuple with the OfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialName

`func (o *AccountIdentity) SetOfficialName(v string)`

SetOfficialName sets OfficialName field to given value.

### HasOfficialName

`func (o *AccountIdentity) HasOfficialName() bool`

HasOfficialName returns a boolean if a field has been set.

### SetOfficialNameNil

`func (o *AccountIdentity) SetOfficialNameNil(b bool)`

 SetOfficialNameNil sets the value for OfficialName to be an explicit nil

### UnsetOfficialName
`func (o *AccountIdentity) UnsetOfficialName()`

UnsetOfficialName ensures that no value is present for OfficialName, not even an explicit nil
### GetType

`func (o *AccountIdentity) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountIdentity) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountIdentity) SetType(v AccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *AccountIdentity) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AccountIdentity) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AccountIdentity) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.


### SetSubtypeNil

`func (o *AccountIdentity) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *AccountIdentity) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetVerificationStatus

`func (o *AccountIdentity) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AccountIdentity) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AccountIdentity) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AccountIdentity) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### SetVerificationStatusNil

`func (o *AccountIdentity) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *AccountIdentity) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
### GetOwners

`func (o *AccountIdentity) GetOwners() []Owner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AccountIdentity) GetOwnersOk() (*[]Owner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AccountIdentity) SetOwners(v []Owner)`

SetOwners sets Owners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


