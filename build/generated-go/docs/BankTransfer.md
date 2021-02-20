# BankTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Plaid’s unique identifier for a bank transfer. | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**AccountId** | **string** | The account ID that should be credited/debited for this bank transfer. | 
**Type** | [**BankTransferType**](BankTransferType.md) |  | 
**User** | [**BankTransferUser**](BankTransferUser.md) |  | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. “10.00”). | 
**IsoCurrencyCode** | **string** | The currency of the transfer amount, e.g. \&quot;USD\&quot; | 
**Description** | **string** | The description of the transfer. | 
**Created** | **string** | The datetime when this bank transfer was created. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60; | 
**Status** | [**BankTransferStatus**](BankTransferStatus.md) |  | 
**Network** | [**BankTransferNetwork**](BankTransferNetwork.md) |  | 
**Cancellable** | **bool** | When &#x60;true&#x60;, you can still cancel this bank transfer. | 
**FailureReason** | Pointer to [**NullableBankTransferFailure**](BankTransferFailure.md) |  | [optional] 
**CustomTag** | Pointer to **NullableString** | A string containing the custom tag provided by the client in the create request. Will be null if not provided. | [optional] 
**Metadata** | **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | 
**OriginationAccountId** | **string** | Plaid’s unique identifier for the origination account that was used for this transfer. | 
**Direction** | [**NullableBankTransferDirection**](BankTransferDirection.md) |  | 

## Methods

### NewBankTransfer

`func NewBankTransfer(id string, achClass ACHClass, accountId string, type_ BankTransferType, user BankTransferUser, amount string, isoCurrencyCode string, description string, created string, status BankTransferStatus, network BankTransferNetwork, cancellable bool, metadata map[string]string, originationAccountId string, direction NullableBankTransferDirection, ) *BankTransfer`

NewBankTransfer instantiates a new BankTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferWithDefaults

`func NewBankTransferWithDefaults() *BankTransfer`

NewBankTransferWithDefaults instantiates a new BankTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransfer) SetId(v string)`

SetId sets Id field to given value.


### GetAchClass

`func (o *BankTransfer) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *BankTransfer) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *BankTransfer) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetAccountId

`func (o *BankTransfer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankTransfer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankTransfer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *BankTransfer) GetType() BankTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankTransfer) GetTypeOk() (*BankTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankTransfer) SetType(v BankTransferType)`

SetType sets Type field to given value.


### GetUser

`func (o *BankTransfer) GetUser() BankTransferUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BankTransfer) GetUserOk() (*BankTransferUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BankTransfer) SetUser(v BankTransferUser)`

SetUser sets User field to given value.


### GetAmount

`func (o *BankTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *BankTransfer) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *BankTransfer) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *BankTransfer) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetDescription

`func (o *BankTransfer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransfer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransfer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreated

`func (o *BankTransfer) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BankTransfer) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BankTransfer) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetStatus

`func (o *BankTransfer) GetStatus() BankTransferStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankTransfer) GetStatusOk() (*BankTransferStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankTransfer) SetStatus(v BankTransferStatus)`

SetStatus sets Status field to given value.


### GetNetwork

`func (o *BankTransfer) GetNetwork() BankTransferNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BankTransfer) GetNetworkOk() (*BankTransferNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BankTransfer) SetNetwork(v BankTransferNetwork)`

SetNetwork sets Network field to given value.


### GetCancellable

`func (o *BankTransfer) GetCancellable() bool`

GetCancellable returns the Cancellable field if non-nil, zero value otherwise.

### GetCancellableOk

`func (o *BankTransfer) GetCancellableOk() (*bool, bool)`

GetCancellableOk returns a tuple with the Cancellable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellable

`func (o *BankTransfer) SetCancellable(v bool)`

SetCancellable sets Cancellable field to given value.


### GetFailureReason

`func (o *BankTransfer) GetFailureReason() BankTransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BankTransfer) GetFailureReasonOk() (*BankTransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BankTransfer) SetFailureReason(v BankTransferFailure)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *BankTransfer) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *BankTransfer) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *BankTransfer) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetCustomTag

`func (o *BankTransfer) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *BankTransfer) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *BankTransfer) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *BankTransfer) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### SetCustomTagNil

`func (o *BankTransfer) SetCustomTagNil(b bool)`

 SetCustomTagNil sets the value for CustomTag to be an explicit nil

### UnsetCustomTag
`func (o *BankTransfer) UnsetCustomTag()`

UnsetCustomTag ensures that no value is present for CustomTag, not even an explicit nil
### GetMetadata

`func (o *BankTransfer) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BankTransfer) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BankTransfer) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *BankTransfer) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BankTransfer) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginationAccountId

`func (o *BankTransfer) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransfer) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransfer) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### GetDirection

`func (o *BankTransfer) GetDirection() BankTransferDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BankTransfer) GetDirectionOk() (*BankTransferDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BankTransfer) SetDirection(v BankTransferDirection)`

SetDirection sets Direction field to given value.


### SetDirectionNil

`func (o *BankTransfer) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *BankTransfer) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


