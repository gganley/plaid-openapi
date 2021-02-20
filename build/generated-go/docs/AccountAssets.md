# AccountAssets

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
**DaysAvailable** | Pointer to **NullableFloat32** | The duration of transaction history available for this Item, typically defined as the time since the date of the earliest transaction in that account. Only returned by Assets endpoints. | [optional] 
**Transactions** | Pointer to [**[]AssetReportTransaction**](AssetReportTransaction.md) | Transaction history associated with the account. Only returned by Assets endpoints. Transaction history returned by endpoints such as &#x60;/transactions/get&#x60; or &#x60;/investments/transactions/get&#x60; will be returned in the top-level &#x60;transactions&#x60; field instead. | [optional] 
**Owners** | [**[]Owner**](Owner.md) | Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. Multiple owners on a single account will be represented in the same &#x60;owner&#x60; object, not in multiple owner objects within the array. | 
**HistoricalBalances** | Pointer to [**[]HistoricalBalance**](HistoricalBalance.md) | Calculated data about the historical balances on the account. Only returned by Assets endpoints. | [optional] 

## Methods

### NewAccountAssets

`func NewAccountAssets(accountId string, balances AccountBalance, name string, type_ AccountType, subtype NullableAccountSubtype, owners []Owner, ) *AccountAssets`

NewAccountAssets instantiates a new AccountAssets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAssetsWithDefaults

`func NewAccountAssetsWithDefaults() *AccountAssets`

NewAccountAssetsWithDefaults instantiates a new AccountAssets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AccountAssets) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AccountAssets) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AccountAssets) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBalances

`func (o *AccountAssets) GetBalances() AccountBalance`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### GetBalancesOk

`func (o *AccountAssets) GetBalancesOk() (*AccountBalance, bool)`

GetBalancesOk returns a tuple with the Balances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalances

`func (o *AccountAssets) SetBalances(v AccountBalance)`

SetBalances sets Balances field to given value.


### GetMask

`func (o *AccountAssets) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *AccountAssets) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *AccountAssets) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *AccountAssets) HasMask() bool`

HasMask returns a boolean if a field has been set.

### SetMaskNil

`func (o *AccountAssets) SetMaskNil(b bool)`

 SetMaskNil sets the value for Mask to be an explicit nil

### UnsetMask
`func (o *AccountAssets) UnsetMask()`

UnsetMask ensures that no value is present for Mask, not even an explicit nil
### GetName

`func (o *AccountAssets) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountAssets) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountAssets) SetName(v string)`

SetName sets Name field to given value.


### GetOfficialName

`func (o *AccountAssets) GetOfficialName() string`

GetOfficialName returns the OfficialName field if non-nil, zero value otherwise.

### GetOfficialNameOk

`func (o *AccountAssets) GetOfficialNameOk() (*string, bool)`

GetOfficialNameOk returns a tuple with the OfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialName

`func (o *AccountAssets) SetOfficialName(v string)`

SetOfficialName sets OfficialName field to given value.

### HasOfficialName

`func (o *AccountAssets) HasOfficialName() bool`

HasOfficialName returns a boolean if a field has been set.

### SetOfficialNameNil

`func (o *AccountAssets) SetOfficialNameNil(b bool)`

 SetOfficialNameNil sets the value for OfficialName to be an explicit nil

### UnsetOfficialName
`func (o *AccountAssets) UnsetOfficialName()`

UnsetOfficialName ensures that no value is present for OfficialName, not even an explicit nil
### GetType

`func (o *AccountAssets) GetType() AccountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountAssets) GetTypeOk() (*AccountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountAssets) SetType(v AccountType)`

SetType sets Type field to given value.


### GetSubtype

`func (o *AccountAssets) GetSubtype() AccountSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *AccountAssets) GetSubtypeOk() (*AccountSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *AccountAssets) SetSubtype(v AccountSubtype)`

SetSubtype sets Subtype field to given value.


### SetSubtypeNil

`func (o *AccountAssets) SetSubtypeNil(b bool)`

 SetSubtypeNil sets the value for Subtype to be an explicit nil

### UnsetSubtype
`func (o *AccountAssets) UnsetSubtype()`

UnsetSubtype ensures that no value is present for Subtype, not even an explicit nil
### GetVerificationStatus

`func (o *AccountAssets) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *AccountAssets) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *AccountAssets) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *AccountAssets) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.

### SetVerificationStatusNil

`func (o *AccountAssets) SetVerificationStatusNil(b bool)`

 SetVerificationStatusNil sets the value for VerificationStatus to be an explicit nil

### UnsetVerificationStatus
`func (o *AccountAssets) UnsetVerificationStatus()`

UnsetVerificationStatus ensures that no value is present for VerificationStatus, not even an explicit nil
### GetDaysAvailable

`func (o *AccountAssets) GetDaysAvailable() float32`

GetDaysAvailable returns the DaysAvailable field if non-nil, zero value otherwise.

### GetDaysAvailableOk

`func (o *AccountAssets) GetDaysAvailableOk() (*float32, bool)`

GetDaysAvailableOk returns a tuple with the DaysAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysAvailable

`func (o *AccountAssets) SetDaysAvailable(v float32)`

SetDaysAvailable sets DaysAvailable field to given value.

### HasDaysAvailable

`func (o *AccountAssets) HasDaysAvailable() bool`

HasDaysAvailable returns a boolean if a field has been set.

### SetDaysAvailableNil

`func (o *AccountAssets) SetDaysAvailableNil(b bool)`

 SetDaysAvailableNil sets the value for DaysAvailable to be an explicit nil

### UnsetDaysAvailable
`func (o *AccountAssets) UnsetDaysAvailable()`

UnsetDaysAvailable ensures that no value is present for DaysAvailable, not even an explicit nil
### GetTransactions

`func (o *AccountAssets) GetTransactions() []AssetReportTransaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *AccountAssets) GetTransactionsOk() (*[]AssetReportTransaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *AccountAssets) SetTransactions(v []AssetReportTransaction)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *AccountAssets) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### SetTransactionsNil

`func (o *AccountAssets) SetTransactionsNil(b bool)`

 SetTransactionsNil sets the value for Transactions to be an explicit nil

### UnsetTransactions
`func (o *AccountAssets) UnsetTransactions()`

UnsetTransactions ensures that no value is present for Transactions, not even an explicit nil
### GetOwners

`func (o *AccountAssets) GetOwners() []Owner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AccountAssets) GetOwnersOk() (*[]Owner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AccountAssets) SetOwners(v []Owner)`

SetOwners sets Owners field to given value.


### GetHistoricalBalances

`func (o *AccountAssets) GetHistoricalBalances() []HistoricalBalance`

GetHistoricalBalances returns the HistoricalBalances field if non-nil, zero value otherwise.

### GetHistoricalBalancesOk

`func (o *AccountAssets) GetHistoricalBalancesOk() (*[]HistoricalBalance, bool)`

GetHistoricalBalancesOk returns a tuple with the HistoricalBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalBalances

`func (o *AccountAssets) SetHistoricalBalances(v []HistoricalBalance)`

SetHistoricalBalances sets HistoricalBalances field to given value.

### HasHistoricalBalances

`func (o *AccountAssets) HasHistoricalBalances() bool`

HasHistoricalBalances returns a boolean if a field has been set.

### SetHistoricalBalancesNil

`func (o *AccountAssets) SetHistoricalBalancesNil(b bool)`

 SetHistoricalBalancesNil sets the value for HistoricalBalances to be an explicit nil

### UnsetHistoricalBalances
`func (o *AccountAssets) UnsetHistoricalBalances()`

UnsetHistoricalBalances ensures that no value is present for HistoricalBalances, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


