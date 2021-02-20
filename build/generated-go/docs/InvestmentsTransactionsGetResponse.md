# InvestmentsTransactionsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**Item**](Item.md) |  | 
**Accounts** | [**[]AccountBase**](AccountBase.md) | The accounts for which transaction history is being fetched. | 
**Securities** | [**[]Security**](Security.md) | All securities for which there is a corresponding transaction being fetched. | 
**InvestmentTransactions** | [**[]InvestmentTransaction**](InvestmentTransaction.md) | The transactions being fetched | 
**TotalInvestmentTransactions** | **int32** | The total number of transactions available within the date range specified. If &#x60;total_investment_transactions&#x60; is larger than the size of the &#x60;transactions&#x60; array, more transactions are available and can be fetched via manipulating the &#x60;offset&#x60; parameter.&#39; | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewInvestmentsTransactionsGetResponse

`func NewInvestmentsTransactionsGetResponse(item Item, accounts []AccountBase, securities []Security, investmentTransactions []InvestmentTransaction, totalInvestmentTransactions int32, requestId string, ) *InvestmentsTransactionsGetResponse`

NewInvestmentsTransactionsGetResponse instantiates a new InvestmentsTransactionsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentsTransactionsGetResponseWithDefaults

`func NewInvestmentsTransactionsGetResponseWithDefaults() *InvestmentsTransactionsGetResponse`

NewInvestmentsTransactionsGetResponseWithDefaults instantiates a new InvestmentsTransactionsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *InvestmentsTransactionsGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *InvestmentsTransactionsGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *InvestmentsTransactionsGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetAccounts

`func (o *InvestmentsTransactionsGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *InvestmentsTransactionsGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *InvestmentsTransactionsGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetSecurities

`func (o *InvestmentsTransactionsGetResponse) GetSecurities() []Security`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *InvestmentsTransactionsGetResponse) GetSecuritiesOk() (*[]Security, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *InvestmentsTransactionsGetResponse) SetSecurities(v []Security)`

SetSecurities sets Securities field to given value.


### GetInvestmentTransactions

`func (o *InvestmentsTransactionsGetResponse) GetInvestmentTransactions() []InvestmentTransaction`

GetInvestmentTransactions returns the InvestmentTransactions field if non-nil, zero value otherwise.

### GetInvestmentTransactionsOk

`func (o *InvestmentsTransactionsGetResponse) GetInvestmentTransactionsOk() (*[]InvestmentTransaction, bool)`

GetInvestmentTransactionsOk returns a tuple with the InvestmentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentTransactions

`func (o *InvestmentsTransactionsGetResponse) SetInvestmentTransactions(v []InvestmentTransaction)`

SetInvestmentTransactions sets InvestmentTransactions field to given value.


### GetTotalInvestmentTransactions

`func (o *InvestmentsTransactionsGetResponse) GetTotalInvestmentTransactions() int32`

GetTotalInvestmentTransactions returns the TotalInvestmentTransactions field if non-nil, zero value otherwise.

### GetTotalInvestmentTransactionsOk

`func (o *InvestmentsTransactionsGetResponse) GetTotalInvestmentTransactionsOk() (*int32, bool)`

GetTotalInvestmentTransactionsOk returns a tuple with the TotalInvestmentTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInvestmentTransactions

`func (o *InvestmentsTransactionsGetResponse) SetTotalInvestmentTransactions(v int32)`

SetTotalInvestmentTransactions sets TotalInvestmentTransactions field to given value.


### GetRequestId

`func (o *InvestmentsTransactionsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InvestmentsTransactionsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InvestmentsTransactionsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


