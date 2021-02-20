# ItemPublicTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicToken** | **string** | A &#x60;public_token&#x60; for the particular Item corresponding to the specified &#x60;access_token&#x60; | 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewItemPublicTokenCreateResponse

`func NewItemPublicTokenCreateResponse(publicToken string, requestId string, ) *ItemPublicTokenCreateResponse`

NewItemPublicTokenCreateResponse instantiates a new ItemPublicTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemPublicTokenCreateResponseWithDefaults

`func NewItemPublicTokenCreateResponseWithDefaults() *ItemPublicTokenCreateResponse`

NewItemPublicTokenCreateResponseWithDefaults instantiates a new ItemPublicTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicToken

`func (o *ItemPublicTokenCreateResponse) GetPublicToken() string`

GetPublicToken returns the PublicToken field if non-nil, zero value otherwise.

### GetPublicTokenOk

`func (o *ItemPublicTokenCreateResponse) GetPublicTokenOk() (*string, bool)`

GetPublicTokenOk returns a tuple with the PublicToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicToken

`func (o *ItemPublicTokenCreateResponse) SetPublicToken(v string)`

SetPublicToken sets PublicToken field to given value.


### GetExpiration

`func (o *ItemPublicTokenCreateResponse) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ItemPublicTokenCreateResponse) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ItemPublicTokenCreateResponse) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ItemPublicTokenCreateResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRequestId

`func (o *ItemPublicTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemPublicTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemPublicTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


