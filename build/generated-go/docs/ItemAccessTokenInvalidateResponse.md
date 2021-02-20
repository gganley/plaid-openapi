# ItemAccessTokenInvalidateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewItemAccessTokenInvalidateResponse

`func NewItemAccessTokenInvalidateResponse(newAccessToken string, requestId string, ) *ItemAccessTokenInvalidateResponse`

NewItemAccessTokenInvalidateResponse instantiates a new ItemAccessTokenInvalidateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemAccessTokenInvalidateResponseWithDefaults

`func NewItemAccessTokenInvalidateResponseWithDefaults() *ItemAccessTokenInvalidateResponse`

NewItemAccessTokenInvalidateResponseWithDefaults instantiates a new ItemAccessTokenInvalidateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAccessToken

`func (o *ItemAccessTokenInvalidateResponse) GetNewAccessToken() string`

GetNewAccessToken returns the NewAccessToken field if non-nil, zero value otherwise.

### GetNewAccessTokenOk

`func (o *ItemAccessTokenInvalidateResponse) GetNewAccessTokenOk() (*string, bool)`

GetNewAccessTokenOk returns a tuple with the NewAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAccessToken

`func (o *ItemAccessTokenInvalidateResponse) SetNewAccessToken(v string)`

SetNewAccessToken sets NewAccessToken field to given value.


### GetRequestId

`func (o *ItemAccessTokenInvalidateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemAccessTokenInvalidateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemAccessTokenInvalidateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


