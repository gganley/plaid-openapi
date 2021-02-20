# TransactionsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountBase**](AccountBase.md) | An array containing the &#x60;accounts&#x60; associated with the Item for which transactions are being returned. Each transaction can be mapped to its corresponding account via the &#x60;account_id&#x60; field. | 
**Transactions** | [**[]Transaction**](Transaction.md) | An array containing transactions from the account. Transactions are returned in reverse chronological order, with the most recent at the beginning of the array. The maximum number of transactions returned is determined by the &#x60;count&#x60; parameter. | 
**TotalTransactions** | **int32** | The total number of transactions available within the date range specified. If &#x60;total_transactions&#x60; is larger than the size of the &#x60;transactions&#x60; array, more transactions are available and can be fetched via manipulating the &#x60;offset&#x60; parameter. | 
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransactionsGetResponse

`func NewTransactionsGetResponse(accounts []AccountBase, transactions []Transaction, totalTransactions int32, item Item, requestId string, ) *TransactionsGetResponse`

NewTransactionsGetResponse instantiates a new TransactionsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsGetResponseWithDefaults

`func NewTransactionsGetResponseWithDefaults() *TransactionsGetResponse`

NewTransactionsGetResponseWithDefaults instantiates a new TransactionsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *TransactionsGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *TransactionsGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *TransactionsGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetTransactions

`func (o *TransactionsGetResponse) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionsGetResponse) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionsGetResponse) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetTotalTransactions

`func (o *TransactionsGetResponse) GetTotalTransactions() int32`

GetTotalTransactions returns the TotalTransactions field if non-nil, zero value otherwise.

### GetTotalTransactionsOk

`func (o *TransactionsGetResponse) GetTotalTransactionsOk() (*int32, bool)`

GetTotalTransactionsOk returns a tuple with the TotalTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTransactions

`func (o *TransactionsGetResponse) SetTotalTransactions(v int32)`

SetTotalTransactions sets TotalTransactions field to given value.


### GetItem

`func (o *TransactionsGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *TransactionsGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *TransactionsGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *TransactionsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


