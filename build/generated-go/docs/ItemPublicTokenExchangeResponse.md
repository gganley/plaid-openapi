# ItemPublicTokenExchangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**ItemId** | **string** | The &#x60;item_id&#x60; value of the Item associated with the returned &#x60;access_token&#x60; | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewItemPublicTokenExchangeResponse

`func NewItemPublicTokenExchangeResponse(accessToken string, itemId string, requestId string, ) *ItemPublicTokenExchangeResponse`

NewItemPublicTokenExchangeResponse instantiates a new ItemPublicTokenExchangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPublicTokenExchangeResponseWithDefaults

`func NewItemPublicTokenExchangeResponseWithDefaults() *ItemPublicTokenExchangeResponse`

NewItemPublicTokenExchangeResponseWithDefaults instantiates a new ItemPublicTokenExchangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ItemPublicTokenExchangeResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ItemPublicTokenExchangeResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ItemPublicTokenExchangeResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetItemId

`func (o *ItemPublicTokenExchangeResponse) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemPublicTokenExchangeResponse) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemPublicTokenExchangeResponse) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetRequestId

`func (o *ItemPublicTokenExchangeResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemPublicTokenExchangeResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemPublicTokenExchangeResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


