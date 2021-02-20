# LinkTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkToken** | **string** | A &#x60;link_token&#x60;, which can be supplied to Link in order to initialize it and receive a &#x60;public_token&#x60;, which can be exchanged for an &#x60;access_token&#x60;. | 
**Expiration** | **time.Time** | The expiration date for the &#x60;link_token&#x60;, in ISO 8601 format. A &#x60;link_token&#x60; created to generate a &#x60;public_token&#x60; that will be exchanged for a new &#x60;access_token&#x60; expires after 24 hours. A &#x60;link_token&#x60; created for an existing Item (such as when updating an existing &#x60;access_token&#x60; by launching Link in update mode) expires after 30 minutes. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewLinkTokenCreateResponse

`func NewLinkTokenCreateResponse(linkToken string, expiration time.Time, requestId string, ) *LinkTokenCreateResponse`

NewLinkTokenCreateResponse instantiates a new LinkTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateResponseWithDefaults

`func NewLinkTokenCreateResponseWithDefaults() *LinkTokenCreateResponse`

NewLinkTokenCreateResponseWithDefaults instantiates a new LinkTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkToken

`func (o *LinkTokenCreateResponse) GetLinkToken() string`

GetLinkToken returns the LinkToken field if non-nil, zero value otherwise.

### GetLinkTokenOk

`func (o *LinkTokenCreateResponse) GetLinkTokenOk() (*string, bool)`

GetLinkTokenOk returns a tuple with the LinkToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkToken

`func (o *LinkTokenCreateResponse) SetLinkToken(v string)`

SetLinkToken sets LinkToken field to given value.


### GetExpiration

`func (o *LinkTokenCreateResponse) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *LinkTokenCreateResponse) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *LinkTokenCreateResponse) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### GetRequestId

`func (o *LinkTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *LinkTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *LinkTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


