# CategoriesGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]Category**](Category.md) | An array of all of the transaction categories used by Plaid. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewCategoriesGetResponse

`func NewCategoriesGetResponse(categories []Category, requestId string, ) *CategoriesGetResponse`

NewCategoriesGetResponse instantiates a new CategoriesGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoriesGetResponseWithDefaults

`func NewCategoriesGetResponseWithDefaults() *CategoriesGetResponse`

NewCategoriesGetResponseWithDefaults instantiates a new CategoriesGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *CategoriesGetResponse) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CategoriesGetResponse) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CategoriesGetResponse) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetRequestId

`func (o *CategoriesGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CategoriesGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CategoriesGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


