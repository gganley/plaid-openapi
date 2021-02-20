# InvestmentsHoldingsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountBase**](AccountBase.md) | The accounts associated with the Item | 
**Holdings** | [**[]Holding**](Holding.md) | The holdings belonging to investment accounts associated with the Item. Details of the securities in the holdings are provided in the &#x60;securities&#x60; field.  | 
**Securities** | [**[]Security**](Security.md) | Objects describing the securities held in the accounts associated with the Item.  | 
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewInvestmentsHoldingsGetResponse

`func NewInvestmentsHoldingsGetResponse(accounts []AccountBase, holdings []Holding, securities []Security, item Item, requestId string, ) *InvestmentsHoldingsGetResponse`

NewInvestmentsHoldingsGetResponse instantiates a new InvestmentsHoldingsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentsHoldingsGetResponseWithDefaults

`func NewInvestmentsHoldingsGetResponseWithDefaults() *InvestmentsHoldingsGetResponse`

NewInvestmentsHoldingsGetResponseWithDefaults instantiates a new InvestmentsHoldingsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *InvestmentsHoldingsGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *InvestmentsHoldingsGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *InvestmentsHoldingsGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetHoldings

`func (o *InvestmentsHoldingsGetResponse) GetHoldings() []Holding`

GetHoldings returns the Holdings field if non-nil, zero value otherwise.

### GetHoldingsOk

`func (o *InvestmentsHoldingsGetResponse) GetHoldingsOk() (*[]Holding, bool)`

GetHoldingsOk returns a tuple with the Holdings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldings

`func (o *InvestmentsHoldingsGetResponse) SetHoldings(v []Holding)`

SetHoldings sets Holdings field to given value.


### GetSecurities

`func (o *InvestmentsHoldingsGetResponse) GetSecurities() []Security`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *InvestmentsHoldingsGetResponse) GetSecuritiesOk() (*[]Security, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *InvestmentsHoldingsGetResponse) SetSecurities(v []Security)`

SetSecurities sets Securities field to given value.


### GetItem

`func (o *InvestmentsHoldingsGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *InvestmentsHoldingsGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *InvestmentsHoldingsGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *InvestmentsHoldingsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InvestmentsHoldingsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InvestmentsHoldingsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


