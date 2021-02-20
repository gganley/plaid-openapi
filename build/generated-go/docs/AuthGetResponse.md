# AuthGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountBase**](AccountBase.md) | The &#x60;accounts&#x60; for which numbers are being retrieved. | 
**Numbers** | [**AuthGetNumbers**](AuthGetNumbers.md) |  | 
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewAuthGetResponse

`func NewAuthGetResponse(accounts []AccountBase, numbers AuthGetNumbers, item Item, requestId string, ) *AuthGetResponse`

NewAuthGetResponse instantiates a new AuthGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetResponseWithDefaults

`func NewAuthGetResponseWithDefaults() *AuthGetResponse`

NewAuthGetResponseWithDefaults instantiates a new AuthGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AuthGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AuthGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AuthGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetNumbers

`func (o *AuthGetResponse) GetNumbers() AuthGetNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *AuthGetResponse) GetNumbersOk() (*AuthGetNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *AuthGetResponse) SetNumbers(v AuthGetNumbers)`

SetNumbers sets Numbers field to given value.


### GetItem

`func (o *AuthGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AuthGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AuthGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *AuthGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AuthGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AuthGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


