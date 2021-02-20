# LiabilitiesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountBase**](AccountBase.md) | An array of accounts associated with the Item | 
**Item** | [**Item**](Item.md) |  | 
**Liabilities** | [**LiabilitiesObject**](LiabilitiesObject.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewLiabilitiesGetResponse

`func NewLiabilitiesGetResponse(accounts []AccountBase, item Item, liabilities LiabilitiesObject, requestId string, ) *LiabilitiesGetResponse`

NewLiabilitiesGetResponse instantiates a new LiabilitiesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiabilitiesGetResponseWithDefaults

`func NewLiabilitiesGetResponseWithDefaults() *LiabilitiesGetResponse`

NewLiabilitiesGetResponseWithDefaults instantiates a new LiabilitiesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *LiabilitiesGetResponse) GetAccounts() []AccountBase`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *LiabilitiesGetResponse) GetAccountsOk() (*[]AccountBase, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *LiabilitiesGetResponse) SetAccounts(v []AccountBase)`

SetAccounts sets Accounts field to given value.


### GetItem

`func (o *LiabilitiesGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *LiabilitiesGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *LiabilitiesGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetLiabilities

`func (o *LiabilitiesGetResponse) GetLiabilities() LiabilitiesObject`

GetLiabilities returns the Liabilities field if non-nil, zero value otherwise.

### GetLiabilitiesOk

`func (o *LiabilitiesGetResponse) GetLiabilitiesOk() (*LiabilitiesObject, bool)`

GetLiabilitiesOk returns a tuple with the Liabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiabilities

`func (o *LiabilitiesGetResponse) SetLiabilities(v LiabilitiesObject)`

SetLiabilities sets Liabilities field to given value.


### GetRequestId

`func (o *LiabilitiesGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LiabilitiesGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LiabilitiesGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


