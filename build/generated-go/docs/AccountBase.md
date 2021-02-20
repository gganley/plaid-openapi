# AccountBase

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

## Methods

### NewAccountBase

`func NewAccountBase(accountId string, balances AccountBalance, name string, type_ AccountType, subtype NullableAccountSubtype, ) *AccountBase`

NewAccountBase instantiates a new AccountBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountBaseWithDefaults

`func NewAccountBaseWithDefaults() *AccountBase`

NewAccountBaseWithDefaults instantiates a new AccountBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountBase) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountBase) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountBase) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBalances

`func (o *AccountBase) GetBalances() AccountBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountBase) GetBalancesOk() (*AccountBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountBase) SetBalances(v AccountBalance)`

SetBalances sets Balances field to given value.


### GetMask

`func (o *AccountBase) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *AccountBase) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *AccountBase) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *AccountBase) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *AccountBase) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *AccountBase) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetName

`func (o *AccountBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountBase) SetName(v string)`

SetName sets Name field to given value.


### GetOfficialName

`func (o *AccountBase) GetOfficialName() string`

GetOfficialName returns the OfficialName field if non-nil, zero value otherwise.

### GetOfficialNameOk

`func (o *AccountBase) GetOfficialNameOk() (*string, bool)`

GetOfficialNameOk returns a tuple with the OfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialName

`func (o *AccountBase) SetOfficialName(v string)`

SetOfficialName sets OfficialName field to given value.

### HasOfficialName

`func (o *AccountBase) HasOfficialName() bool`

HasOfficialName returns a boolean if a field has been set.

### SetOfficialNameNil

`func (o *AccountBase) SetOfficialNameNil(b bool)`

 SetOfficialNameNil sets the value for OfficialName to be an explicit nil

### UnsetOfficialName
`func (o *AccountBase) UnsetOfficialName()`

UnsetOfficialName ensures that no value is present for OfficialName, not even an explicit nil
### GetType

`func (o *AccountBase) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountBase) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountBase) SetType(v AccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *AccountBase) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AccountBase) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AccountBase) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.


### SetSubtypeNil

`func (o *AccountBase) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *AccountBase) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetVerificationStatus

`func (o *AccountBase) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AccountBase) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AccountBase) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AccountBase) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### SetVerificationStatusNil

`func (o *AccountBase) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *AccountBase) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


