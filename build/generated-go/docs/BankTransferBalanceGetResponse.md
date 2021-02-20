# BankTransferBalanceGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | [**BankTransferBalance**](BankTransferBalance.md) |  | 
**OriginationAccountId** | Pointer to **NullableString** | The ID of the origination account that this balance belongs to. | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferBalanceGetResponse

`func NewBankTransferBalanceGetResponse(balance BankTransferBalance, requestId string, ) *BankTransferBalanceGetResponse`

NewBankTransferBalanceGetResponse instantiates a new BankTransferBalanceGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferBalanceGetResponseWithDefaults

`func NewBankTransferBalanceGetResponseWithDefaults() *BankTransferBalanceGetResponse`

NewBankTransferBalanceGetResponseWithDefaults instantiates a new BankTransferBalanceGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *BankTransferBalanceGetResponse) GetBalance() BankTransferBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BankTransferBalanceGetResponse) GetBalanceOk() (*BankTransferBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BankTransferBalanceGetResponse) SetBalance(v BankTransferBalance)`

SetBalance sets Balance field to given value.


### GetOriginationAccountId

`func (o *BankTransferBalanceGetResponse) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *BankTransferBalanceGetResponse) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *BankTransferBalanceGetResponse) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *BankTransferBalanceGetResponse) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *BankTransferBalanceGetResponse) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *BankTransferBalanceGetResponse) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetRequestId

`func (o *BankTransferBalanceGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferBalanceGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferBalanceGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


