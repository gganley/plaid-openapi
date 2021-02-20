# BankTransferListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankTransfers** | [**[]BankTransfer**](BankTransfer.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferListResponse

`func NewBankTransferListResponse(bankTransfers []BankTransfer, requestId string, ) *BankTransferListResponse`

NewBankTransferListResponse instantiates a new BankTransferListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferListResponseWithDefaults

`func NewBankTransferListResponseWithDefaults() *BankTransferListResponse`

NewBankTransferListResponseWithDefaults instantiates a new BankTransferListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankTransfers

`func (o *BankTransferListResponse) GetBankTransfers() []BankTransfer`

GetBankTransfers returns the BankTransfers field if non-nil, zero value otherwise.

### GetBankTransfersOk

`func (o *BankTransferListResponse) GetBankTransfersOk() (*[]BankTransfer, bool)`

GetBankTransfersOk returns a tuple with the BankTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransfers

`func (o *BankTransferListResponse) SetBankTransfers(v []BankTransfer)`

SetBankTransfers sets BankTransfers field to given value.


### GetRequestId

`func (o *BankTransferListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


