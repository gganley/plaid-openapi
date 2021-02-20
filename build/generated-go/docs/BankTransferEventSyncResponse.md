# BankTransferEventSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankTransferEvents** | [**[]BankTransferEvent**](BankTransferEvent.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferEventSyncResponse

`func NewBankTransferEventSyncResponse(bankTransferEvents []BankTransferEvent, requestId string, ) *BankTransferEventSyncResponse`

NewBankTransferEventSyncResponse instantiates a new BankTransferEventSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferEventSyncResponseWithDefaults

`func NewBankTransferEventSyncResponseWithDefaults() *BankTransferEventSyncResponse`

NewBankTransferEventSyncResponseWithDefaults instantiates a new BankTransferEventSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankTransferEvents

`func (o *BankTransferEventSyncResponse) GetBankTransferEvents() []BankTransferEvent`

GetBankTransferEvents returns the BankTransferEvents field if non-nil, zero value otherwise.

### GetBankTransferEventsOk

`func (o *BankTransferEventSyncResponse) GetBankTransferEventsOk() (*[]BankTransferEvent, bool)`

GetBankTransferEventsOk returns a tuple with the BankTransferEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferEvents

`func (o *BankTransferEventSyncResponse) SetBankTransferEvents(v []BankTransferEvent)`

SetBankTransferEvents sets BankTransferEvents field to given value.


### GetRequestId

`func (o *BankTransferEventSyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferEventSyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferEventSyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


