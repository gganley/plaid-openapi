# BankTransferEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **int32** | Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers. | 
**Timestamp** | **string** | The datetime when this event occurred. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60;. | 
**EventType** | [**BankTransferEventType**](BankTransferEventType.md) |  | 
**AccountId** | **string** | The account ID associated with the bank transfer. | 
**BankTransferId** | **string** | Plaid’s unique identifier for a bank transfer. | 
**OriginationAccountId** | Pointer to **NullableString** | The ID of the origination account that this balance belongs to. | [optional] 
**BankTransferType** | [**BankTransferType**](BankTransferType.md) |  | 
**BankTransferAmount** | **string** | The bank transfer amount. | 
**BankTransferIsoCurrencyCode** | **string** | The currency of the bank transfer amount. | 
**FailureReason** | [**NullableBankTransferFailure**](BankTransferFailure.md) |  | 
**Direction** | [**NullableBankTransferDirection**](BankTransferDirection.md) |  | 
**ReceiverDetails** | [**NullableBankTransferReceiverDetails**](BankTransferReceiverDetails.md) |  | 

## Methods

### NewBankTransferEvent

`func NewBankTransferEvent(eventId int32, timestamp string, eventType BankTransferEventType, accountId string, bankTransferId string, bankTransferType BankTransferType, bankTransferAmount string, bankTransferIsoCurrencyCode string, failureReason NullableBankTransferFailure, direction NullableBankTransferDirection, receiverDetails NullableBankTransferReceiverDetails, ) *BankTransferEvent`

NewBankTransferEvent instantiates a new BankTransferEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferEventWithDefaults

`func NewBankTransferEventWithDefaults() *BankTransferEvent`

NewBankTransferEventWithDefaults instantiates a new BankTransferEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *BankTransferEvent) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *BankTransferEvent) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *BankTransferEvent) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetTimestamp

`func (o *BankTransferEvent) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BankTransferEvent) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BankTransferEvent) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetEventType

`func (o *BankTransferEvent) GetEventType() BankTransferEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *BankTransferEvent) GetEventTypeOk() (*BankTransferEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *BankTransferEvent) SetEventType(v BankTransferEventType)`

SetEventType sets EventType field to given value.


### GetAccountId

`func (o *BankTransferEvent) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankTransferEvent) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankTransferEvent) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBankTransferId

`func (o *BankTransferEvent) GetBankTransferId() string`

GetBankTransferId returns the BankTransferId field if non-nil, zero value otherwise.

### GetBankTransferIdOk

`func (o *BankTransferEvent) GetBankTransferIdOk() (*string, bool)`

GetBankTransferIdOk returns a tuple with the BankTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferId

`func (o *BankTransferEvent) SetBankTransferId(v string)`

SetBankTransferId sets BankTransferId field to given value.


### GetOriginationAccountId

`func (o *BankTransferEvent) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransferEvent) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransferEvent) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *BankTransferEvent) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *BankTransferEvent) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *BankTransferEvent) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetBankTransferType

`func (o *BankTransferEvent) GetBankTransferType() BankTransferType`

GetBankTransferType returns the BankTransferType field if non-nil, zero value otherwise.

### GetBankTransferTypeOk

`func (o *BankTransferEvent) GetBankTransferTypeOk() (*BankTransferType, bool)`

GetBankTransferTypeOk returns a tuple with the BankTransferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferType

`func (o *BankTransferEvent) SetBankTransferType(v BankTransferType)`

SetBankTransferType sets BankTransferType field to given value.


### GetBankTransferAmount

`func (o *BankTransferEvent) GetBankTransferAmount() string`

GetBankTransferAmount returns the BankTransferAmount field if non-nil, zero value otherwise.

### GetBankTransferAmountOk

`func (o *BankTransferEvent) GetBankTransferAmountOk() (*string, bool)`

GetBankTransferAmountOk returns a tuple with the BankTransferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferAmount

`func (o *BankTransferEvent) SetBankTransferAmount(v string)`

SetBankTransferAmount sets BankTransferAmount field to given value.


### GetBankTransferIsoCurrencyCode

`func (o *BankTransferEvent) GetBankTransferIsoCurrencyCode() string`

GetBankTransferIsoCurrencyCode returns the BankTransferIsoCurrencyCode field if non-nil, zero value otherwise.

### GetBankTransferIsoCurrencyCodeOk

`func (o *BankTransferEvent) GetBankTransferIsoCurrencyCodeOk() (*string, bool)`

GetBankTransferIsoCurrencyCodeOk returns a tuple with the BankTransferIsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferIsoCurrencyCode

`func (o *BankTransferEvent) SetBankTransferIsoCurrencyCode(v string)`

SetBankTransferIsoCurrencyCode sets BankTransferIsoCurrencyCode field to given value.


### GetFailureReason

`func (o *BankTransferEvent) GetFailureReason() BankTransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *BankTransferEvent) GetFailureReasonOk() (*BankTransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *BankTransferEvent) SetFailureReason(v BankTransferFailure)`

SetFailureReason sets FailureReason field to given value.


### SetFailureReasonNil

`func (o *BankTransferEvent) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *BankTransferEvent) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetDirection

`func (o *BankTransferEvent) GetDirection() BankTransferDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BankTransferEvent) GetDirectionOk() (*BankTransferDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BankTransferEvent) SetDirection(v BankTransferDirection)`

SetDirection sets Direction field to given value.


### SetDirectionNil

`func (o *BankTransferEvent) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *BankTransferEvent) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetReceiverDetails

`func (o *BankTransferEvent) GetReceiverDetails() BankTransferReceiverDetails`

GetReceiverDetails returns the ReceiverDetails field if non-nil, zero value otherwise.

### GetReceiverDetailsOk

`func (o *BankTransferEvent) GetReceiverDetailsOk() (*BankTransferReceiverDetails, bool)`

GetReceiverDetailsOk returns a tuple with the ReceiverDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverDetails

`func (o *BankTransferEvent) SetReceiverDetails(v BankTransferReceiverDetails)`

SetReceiverDetails sets ReceiverDetails field to given value.


### SetReceiverDetailsNil

`func (o *BankTransferEvent) SetReceiverDetailsNil(b bool)`

 SetReceiverDetailsNil sets the value for ReceiverDetails to be an explicit nil

### UnsetReceiverDetails
`func (o *BankTransferEvent) UnsetReceiverDetails()`

UnsetReceiverDetails ensures that no value is present for ReceiverDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


