# ItemImportRequestUserAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Opaque user identifier | 
**AuthToken** | **string** | Authorization token Plaid will use to aggregate this user’s accounts | 

## Methods

### NewItemImportRequestUserAuth

`func NewItemImportRequestUserAuth(userId string, authToken string, ) *ItemImportRequestUserAuth`

NewItemImportRequestUserAuth instantiates a new ItemImportRequestUserAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImportRequestUserAuthWithDefaults

`func NewItemImportRequestUserAuthWithDefaults() *ItemImportRequestUserAuth`

NewItemImportRequestUserAuthWithDefaults instantiates a new ItemImportRequestUserAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ItemImportRequestUserAuth) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ItemImportRequestUserAuth) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ItemImportRequestUserAuth) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetAuthToken

`func (o *ItemImportRequestUserAuth) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *ItemImportRequestUserAuth) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *ItemImportRequestUserAuth) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


