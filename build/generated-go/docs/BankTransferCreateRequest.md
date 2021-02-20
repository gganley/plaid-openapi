# BankTransferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**IdempotencyKey** | **string** | A random key provided by the client, per unique bank transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a bank transfer fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single bank transfer is created. | 
**AccessToken** | **string** | The Plaid &#x60;access_token&#x60; for the account that will be debited or credited. | 
**AccountId** | **string** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. | 
**Type** | [**BankTransferType**](BankTransferType.md) |  | 
**Network** | [**BankTransferNetwork**](BankTransferNetwork.md) |  | 
**Amount** | **string** | The transfer amount (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**IsoCurrencyCode** | **string** | The currency of the transfer amount – should be set to \&quot;USD\&quot;. | 
**Description** | **string** | The transfer description. Maximum of 10 characters. | 
**AchClass** | Pointer to [**ACHClass**](ACHClass.md) |  | [optional] 
**User** | [**BankTransferUser**](BankTransferUser.md) |  | 
**CustomTag** | Pointer to **NullableString** | An arbitrary string provided by the client for storage with the bank transfer. Will be returned in all &#x60;BankTransfer&#x60; objects. May be up to 100 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | [optional] 
**OriginationAccountId** | Pointer to **NullableString** | Plaid’s unique identifier for the origination account for this transfer. If you have more than one origination account, this value must be specified. | [optional] 

## Methods

### NewBankTransferCreateRequest

`func NewBankTransferCreateRequest(idempotencyKey string, accessToken string, accountId string, type_ BankTransferType, network BankTransferNetwork, amount string, isoCurrencyCode string, description string, user BankTransferUser, ) *BankTransferCreateRequest`

NewBankTransferCreateRequest instantiates a new BankTransferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferCreateRequestWithDefaults

`func NewBankTransferCreateRequestWithDefaults() *BankTransferCreateRequest`

NewBankTransferCreateRequestWithDefaults instantiates a new BankTransferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *BankTransferCreateRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *BankTransferCreateRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *BankTransferCreateRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetAccessToken

`func (o *BankTransferCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BankTransferCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BankTransferCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *BankTransferCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankTransferCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankTransferCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *BankTransferCreateRequest) GetType() BankTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankTransferCreateRequest) GetTypeOk() (*BankTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankTransferCreateRequest) SetType(v BankTransferType)`

SetType sets Type field to given value.


### GetNetwork

`func (o *BankTransferCreateRequest) GetNetwork() BankTransferNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BankTransferCreateRequest) GetNetworkOk() (*BankTransferNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BankTransferCreateRequest) SetNetwork(v BankTransferNetwork)`

SetNetwork sets Network field to given value.


### GetAmount

`func (o *BankTransferCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransferCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransferCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *BankTransferCreateRequest) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *BankTransferCreateRequest) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *BankTransferCreateRequest) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetDescription

`func (o *BankTransferCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransferCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransferCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAchClass

`func (o *BankTransferCreateRequest) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *BankTransferCreateRequest) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *BankTransferCreateRequest) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.

### HasAchClass

`func (o *BankTransferCreateRequest) HasAchClass() bool`

HasAchClass returns a boolean if a field has been set.

### GetUser

`func (o *BankTransferCreateRequest) GetUser() BankTransferUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BankTransferCreateRequest) GetUserOk() (*BankTransferUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BankTransferCreateRequest) SetUser(v BankTransferUser)`

SetUser sets User field to given value.


### GetCustomTag

`func (o *BankTransferCreateRequest) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *BankTransferCreateRequest) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *BankTransferCreateRequest) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *BankTransferCreateRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### SetCustomTagNil

`func (o *BankTransferCreateRequest) SetCustomTagNil(b bool)`

 SetCustomTagNil sets the value for CustomTag to be an explicit nil

### UnsetCustomTag
`func (o *BankTransferCreateRequest) UnsetCustomTag()`

UnsetCustomTag ensures that no value is present for CustomTag, not even an explicit nil
### GetMetadata

`func (o *BankTransferCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BankTransferCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BankTransferCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BankTransferCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *BankTransferCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *BankTransferCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginationAccountId

`func (o *BankTransferCreateRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransferCreateRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransferCreateRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *BankTransferCreateRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *BankTransferCreateRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *BankTransferCreateRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


