# StandaloneInvestmentTransactionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buy** | **string** | Buying an investment | 
**Sell** | **string** | Selling an investment | 
**Cancel** | **string** | A cancellation of a pending transaction | 
**Cash** | **string** | Activity that modifies a cash position | 
**Fee** | **string** | Fees on the account, e.g. commission, bookkeeping, options-related. | 
**Transfer** | **string** | Activity which modifies a position, but not through buy/sell activity e.g. options exercise, portfolio transfer | 

## Methods

### NewStandaloneInvestmentTransactionType

`func NewStandaloneInvestmentTransactionType(buy string, sell string, cancel string, cash string, fee string, transfer string, ) *StandaloneInvestmentTransactionType`

NewStandaloneInvestmentTransactionType instantiates a new StandaloneInvestmentTransactionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneInvestmentTransactionTypeWithDefaults

`func NewStandaloneInvestmentTransactionTypeWithDefaults() *StandaloneInvestmentTransactionType`

NewStandaloneInvestmentTransactionTypeWithDefaults instantiates a new StandaloneInvestmentTransactionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuy

`func (o *StandaloneInvestmentTransactionType) GetBuy() string`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *StandaloneInvestmentTransactionType) GetBuyOk() (*string, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *StandaloneInvestmentTransactionType) SetBuy(v string)`

SetBuy sets Buy field to given value.


### GetSell

`func (o *StandaloneInvestmentTransactionType) GetSell() string`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *StandaloneInvestmentTransactionType) GetSellOk() (*string, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *StandaloneInvestmentTransactionType) SetSell(v string)`

SetSell sets Sell field to given value.


### GetCancel

`func (o *StandaloneInvestmentTransactionType) GetCancel() string`

GetCancel returns the Cancel field if non-nil, zero value otherwise.

### GetCancelOk

`func (o *StandaloneInvestmentTransactionType) GetCancelOk() (*string, bool)`

GetCancelOk returns a tuple with the Cancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancel

`func (o *StandaloneInvestmentTransactionType) SetCancel(v string)`

SetCancel sets Cancel field to given value.


### GetCash

`func (o *StandaloneInvestmentTransactionType) GetCash() string`

GetCash returns the Cash field if non-nil, zero value otherwise.

### GetCashOk

`func (o *StandaloneInvestmentTransactionType) GetCashOk() (*string, bool)`

GetCashOk returns a tuple with the Cash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCash

`func (o *StandaloneInvestmentTransactionType) SetCash(v string)`

SetCash sets Cash field to given value.


### GetFee

`func (o *StandaloneInvestmentTransactionType) GetFee() string`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *StandaloneInvestmentTransactionType) GetFeeOk() (*string, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *StandaloneInvestmentTransactionType) SetFee(v string)`

SetFee sets Fee field to given value.


### GetTransfer

`func (o *StandaloneInvestmentTransactionType) GetTransfer() string`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *StandaloneInvestmentTransactionType) GetTransferOk() (*string, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *StandaloneInvestmentTransactionType) SetTransfer(v string)`

SetTransfer sets Transfer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


