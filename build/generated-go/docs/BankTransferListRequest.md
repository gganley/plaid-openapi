# BankTransferListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**StartDate** | Pointer to **NullableTime** | The start datetime of bank transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**EndDate** | Pointer to **NullableTime** | The end datetime of bank transfers to list. This should be in RFC 3339 format (i.e. &#x60;2019-12-06T22:35:49Z&#x60;) | [optional] 
**Count** | Pointer to **int32** | The maximum number of bank transfers to return. | [optional] [default to 25]
**Offset** | Pointer to **int32** | The number of bank transfers to skip before returning results. | [optional] [default to 0]
**OriginationAccountId** | Pointer to **NullableString** | Filter bank transfers to only those originated through the specified origination account. | [optional] 
**Direction** | Pointer to [**NullableBankTransferDirection**](BankTransferDirection.md) |  | [optional] 

## Methods

### NewBankTransferListRequest

`func NewBankTransferListRequest() *BankTransferListRequest`

NewBankTransferListRequest instantiates a new BankTransferListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferListRequestWithDefaults

`func NewBankTransferListRequestWithDefaults() *BankTransferListRequest`

NewBankTransferListRequestWithDefaults instantiates a new BankTransferListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStartDate

`func (o *BankTransferListRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *BankTransferListRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *BankTransferListRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *BankTransferListRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *BankTransferListRequest) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *BankTransferListRequest) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *BankTransferListRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *BankTransferListRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *BankTransferListRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *BankTransferListRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *BankTransferListRequest) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *BankTransferListRequest) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetCount

`func (o *BankTransferListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BankTransferListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BankTransferListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BankTransferListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOffset

`func (o *BankTransferListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BankTransferListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BankTransferListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BankTransferListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOriginationAccountId

`func (o *BankTransferListRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransferListRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransferListRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *BankTransferListRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *BankTransferListRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *BankTransferListRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetDirection

`func (o *BankTransferListRequest) GetDirection() BankTransferDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BankTransferListRequest) GetDirectionOk() (*BankTransferDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BankTransferListRequest) SetDirection(v BankTransferDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BankTransferListRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *BankTransferListRequest) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *BankTransferListRequest) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


