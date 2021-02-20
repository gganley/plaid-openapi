# BankTransferCancelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**BankTransferId** | **string** | Plaid’s unique identifier for a bank transfer. | 

## Methods

### NewBankTransferCancelRequest

`func NewBankTransferCancelRequest(bankTransferId string, ) *BankTransferCancelRequest`

NewBankTransferCancelRequest instantiates a new BankTransferCancelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferCancelRequestWithDefaults

`func NewBankTransferCancelRequestWithDefaults() *BankTransferCancelRequest`

NewBankTransferCancelRequestWithDefaults instantiates a new BankTransferCancelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferCancelRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferCancelRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferCancelRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferCancelRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferCancelRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferCancelRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferCancelRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferCancelRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetBankTransferId

`func (o *BankTransferCancelRequest) GetBankTransferId() string`

GetBankTransferId returns the BankTransferId field if non-nil, zero value otherwise.

### GetBankTransferIdOk

`func (o *BankTransferCancelRequest) GetBankTransferIdOk() (*string, bool)`

GetBankTransferIdOk returns a tuple with the BankTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferId

`func (o *BankTransferCancelRequest) SetBankTransferId(v string)`

SetBankTransferId sets BankTransferId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


