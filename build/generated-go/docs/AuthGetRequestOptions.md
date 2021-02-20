# AuthGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | A list of &#x60;account_ids&#x60; to retrieve for the Item. Note: An error will be returned if a provided &#x60;account_id&#x60; is not associated with the Item. | [optional] 

## Methods

### NewAuthGetRequestOptions

`func NewAuthGetRequestOptions() *AuthGetRequestOptions`

NewAuthGetRequestOptions instantiates a new AuthGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthGetRequestOptionsWithDefaults

`func NewAuthGetRequestOptionsWithDefaults() *AuthGetRequestOptions`

NewAuthGetRequestOptionsWithDefaults instantiates a new AuthGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIds

`func (o *AuthGetRequestOptions) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *AuthGetRequestOptions) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *AuthGetRequestOptions) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *AuthGetRequestOptions) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


