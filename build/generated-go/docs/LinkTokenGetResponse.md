# LinkTokenGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkToken** | Pointer to **string** | A &#x60;link_token&#x60;, which can be supplied to Link in order to initialize it and receive a &#x60;public_token&#x60;, which can be exchanged for an &#x60;access_token&#x60;. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | The creation timestamp for the &#x60;link_token&#x60;, in ISO 8601 format. | [optional] 
**Expiration** | Pointer to **NullableTime** | The expiration timestamp for the &#x60;link_token&#x60;, in ISO 8601 format. | [optional] 
**Metadata** | Pointer to [**LinkTokenGetMetadataResponse**](LinkTokenGetMetadataResponse.md) |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewLinkTokenGetResponse

`func NewLinkTokenGetResponse(requestId string, ) *LinkTokenGetResponse`

NewLinkTokenGetResponse instantiates a new LinkTokenGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenGetResponseWithDefaults

`func NewLinkTokenGetResponseWithDefaults() *LinkTokenGetResponse`

NewLinkTokenGetResponseWithDefaults instantiates a new LinkTokenGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkToken

`func (o *LinkTokenGetResponse) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *LinkTokenGetResponse) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *LinkTokenGetResponse) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.

### HasLinkToken

`func (o *LinkTokenGetResponse) HasLinkToken() bool`

HasLinkToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LinkTokenGetResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LinkTokenGetResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LinkTokenGetResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LinkTokenGetResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LinkTokenGetResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LinkTokenGetResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetExpiration

`func (o *LinkTokenGetResponse) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LinkTokenGetResponse) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LinkTokenGetResponse) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *LinkTokenGetResponse) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### SetExpirationNil

`func (o *LinkTokenGetResponse) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *LinkTokenGetResponse) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetMetadata

`func (o *LinkTokenGetResponse) GetMetadata() LinkTokenGetMetadataResponse`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LinkTokenGetResponse) GetMetadataOk() (*LinkTokenGetMetadataResponse, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LinkTokenGetResponse) SetMetadata(v LinkTokenGetMetadataResponse)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LinkTokenGetResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRequestId

`func (o *LinkTokenGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LinkTokenGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LinkTokenGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


