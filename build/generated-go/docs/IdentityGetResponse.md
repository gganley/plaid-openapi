# IdentityGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]AccountIdentity**](AccountIdentity.md) | The accounts for which Identity data has been requested | 
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewIdentityGetResponse

`func NewIdentityGetResponse(accounts []AccountIdentity, item Item, requestId string, ) *IdentityGetResponse`

NewIdentityGetResponse instantiates a new IdentityGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityGetResponseWithDefaults

`func NewIdentityGetResponseWithDefaults() *IdentityGetResponse`

NewIdentityGetResponseWithDefaults instantiates a new IdentityGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *IdentityGetResponse) GetAccounts() []AccountIdentity`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *IdentityGetResponse) GetAccountsOk() (*[]AccountIdentity, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *IdentityGetResponse) SetAccounts(v []AccountIdentity)`

SetAccounts sets Accounts field to given value.


### GetItem

`func (o *IdentityGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *IdentityGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *IdentityGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *IdentityGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IdentityGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IdentityGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


