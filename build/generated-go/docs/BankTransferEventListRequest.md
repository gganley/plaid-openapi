# BankTransferEventListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**StartDate** | Pointer to **NullableTime** | The start datetime of bank transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**EndDate** | Pointer to **NullableTime** | The end datetime of bank transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**BankTransferId** | Pointer to **NullableString** | Plaid’s unique identifier for a bank transfer. | [optional] 
**AccountId** | Pointer to **NullableString** | The account ID to get events for all transactions to/from an account. | [optional] 
**BankTransferType** | Pointer to **NullableString** | The type of bank transfer. This will be either &#x60;debit&#x60; or &#x60;credit&#x60;.  A &#x60;debit&#x60; indicates a transfer of money into your origination account; a &#x60;credit&#x60; indicates a transfer of money out of your origination account. | [optional] 
**EventTypes** | Pointer to [**[]BankTransferEventType**](BankTransferEventType.md) | Filter events by event type. | [optional] 
**Count** | Pointer to **NullableInt32** | The maximum number of bank transfer events to return. If the number of events matching the above parameters is greater than &#x60;count&#x60;, the most recent events will be returned. | [optional] [default to 25]
**Offset** | Pointer to **NullableInt32** | The offset into the list of bank transfer events. When &#x60;count&#x60;&#x3D;25 and &#x60;offset&#x60;&#x3D;0, the first 25 events will be returned. When &#x60;count&#x60;&#x3D;25 and &#x60;offset&#x60;&#x3D;25, the next 25 bank transfer events will be returned. | [optional] [default to 0]
**OriginationAccountId** | Pointer to **NullableString** | The origination account ID to get events for transfers from a specific origination account. | [optional] 
**Direction** | Pointer to **NullableString** | Indicates the direction of the transfer: &#x60;outbound&#x60; for API-initiated transfers, or &#x60;inbound&#x60; for payments received by the FBO account. | [optional] 

## Methods

### NewBankTransferEventListRequest

`func NewBankTransferEventListRequest() *BankTransferEventListRequest`

NewBankTransferEventListRequest instantiates a new BankTransferEventListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferEventListRequestWithDefaults

`func NewBankTransferEventListRequestWithDefaults() *BankTransferEventListRequest`

NewBankTransferEventListRequestWithDefaults instantiates a new BankTransferEventListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferEventListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferEventListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferEventListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferEventListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferEventListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferEventListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferEventListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferEventListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStartDate

`func (o *BankTransferEventListRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankTransferEventListRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankTransferEventListRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankTransferEventListRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BankTransferEventListRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BankTransferEventListRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BankTransferEventListRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankTransferEventListRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankTransferEventListRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankTransferEventListRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BankTransferEventListRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BankTransferEventListRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetBankTransferId

`func (o *BankTransferEventListRequest) GetBankTransferId() string`

GetBankTransferId returns the BankTransferId field if non-nil, zero value otherwise.

### GetBankTransferIdOk

`func (o *BankTransferEventListRequest) GetBankTransferIdOk() (*string, bool)`

GetBankTransferIdOk returns a tuple with the BankTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferId

`func (o *BankTransferEventListRequest) SetBankTransferId(v string)`

SetBankTransferId sets BankTransferId field to given value.

### HasBankTransferId

`func (o *BankTransferEventListRequest) HasBankTransferId() bool`

HasBankTransferId returns a boolean if a field has been set.

### SetBankTransferIdNil

`func (o *BankTransferEventListRequest) SetBankTransferIdNil(b bool)`

 SetBankTransferIdNil sets the value for BankTransferId to be an explicit nil

### UnsetBankTransferId
`func (o *BankTransferEventListRequest) UnsetBankTransferId()`

UnsetBankTransferId ensures that no value is present for BankTransferId, not even an explicit nil
### GetAccountId

`func (o *BankTransferEventListRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankTransferEventListRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankTransferEventListRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *BankTransferEventListRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *BankTransferEventListRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BankTransferEventListRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetBankTransferType

`func (o *BankTransferEventListRequest) GetBankTransferType() string`

GetBankTransferType returns the BankTransferType field if non-nil, zero value otherwise.

### GetBankTransferTypeOk

`func (o *BankTransferEventListRequest) GetBankTransferTypeOk() (*string, bool)`

GetBankTransferTypeOk returns a tuple with the BankTransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferType

`func (o *BankTransferEventListRequest) SetBankTransferType(v string)`

SetBankTransferType sets BankTransferType field to given value.

### HasBankTransferType

`func (o *BankTransferEventListRequest) HasBankTransferType() bool`

HasBankTransferType returns a boolean if a field has been set.

### SetBankTransferTypeNil

`func (o *BankTransferEventListRequest) SetBankTransferTypeNil(b bool)`

 SetBankTransferTypeNil sets the value for BankTransferType to be an explicit nil

### UnsetBankTransferType
`func (o *BankTransferEventListRequest) UnsetBankTransferType()`

UnsetBankTransferType ensures that no value is present for BankTransferType, not even an explicit nil
### GetEventTypes

`func (o *BankTransferEventListRequest) GetEventTypes() []BankTransferEventType`

GetEventTypes returns the EventTypes field if non-nil, zero value otherwise.

### GetEventTypesOk

`func (o *BankTransferEventListRequest) GetEventTypesOk() (*[]BankTransferEventType, bool)`

GetEventTypesOk returns a tuple with the EventTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypes

`func (o *BankTransferEventListRequest) SetEventTypes(v []BankTransferEventType)`

SetEventTypes sets EventTypes field to given value.

### HasEventTypes

`func (o *BankTransferEventListRequest) HasEventTypes() bool`

HasEventTypes returns a boolean if a field has been set.

### GetCount

`func (o *BankTransferEventListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BankTransferEventListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BankTransferEventListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BankTransferEventListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *BankTransferEventListRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *BankTransferEventListRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOffset

`func (o *BankTransferEventListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BankTransferEventListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BankTransferEventListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BankTransferEventListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *BankTransferEventListRequest) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *BankTransferEventListRequest) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetOriginationAccountId

`func (o *BankTransferEventListRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransferEventListRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransferEventListRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *BankTransferEventListRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *BankTransferEventListRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *BankTransferEventListRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetDirection

`func (o *BankTransferEventListRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BankTransferEventListRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BankTransferEventListRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BankTransferEventListRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *BankTransferEventListRequest) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *BankTransferEventListRequest) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


