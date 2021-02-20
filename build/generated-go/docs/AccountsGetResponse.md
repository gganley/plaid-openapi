# AccountsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountBase**](AccountBase.md) | An array of financial institution accounts associated with the Item. If &#x60;/accounts/balance/get&#x60; was called, each account will include real-time balance information. | 
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewAccountsGetResponse

`func NewAccountsGetResponse(accounts []AccountBase, item Item, requestId string, ) *AccountsGetResponse`

NewAccountsGetResponse instantiates a new AccountsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsGetResponseWithDefaults

`func NewAccountsGetResponseWithDefaults() *AccountsGetResponse`

NewAccountsGetResponseWithDefaults instantiates a new AccountsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetItem

`func (o *AccountsGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AccountsGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AccountsGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *AccountsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AccountsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AccountsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


