# ItemGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**Item**](Item.md) |  | 
**Status** | Pointer to [**NullableItemStatus**](ItemStatus.md) |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**AccessToken** | Pointer to [**NullableNullableAccessToken**](NullableAccessToken.md) |  | [optional] 

## Methods

### NewItemGetResponse

`func NewItemGetResponse(item Item, requestId string, ) *ItemGetResponse`

NewItemGetResponse instantiates a new ItemGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGetResponseWithDefaults

`func NewItemGetResponseWithDefaults() *ItemGetResponse`

NewItemGetResponseWithDefaults instantiates a new ItemGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *ItemGetResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemGetResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemGetResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetStatus

`func (o *ItemGetResponse) GetStatus() ItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ItemGetResponse) GetStatusOk() (*ItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ItemGetResponse) SetStatus(v ItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ItemGetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ItemGetResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ItemGetResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetRequestId

`func (o *ItemGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetAccessToken

`func (o *ItemGetResponse) GetAccessToken() NullableAccessToken`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ItemGetResponse) GetAccessTokenOk() (*NullableAccessToken, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ItemGetResponse) SetAccessToken(v NullableAccessToken)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *ItemGetResponse) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *ItemGetResponse) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *ItemGetResponse) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


