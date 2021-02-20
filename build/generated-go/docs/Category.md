# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **string** | An identifying number for the category. &#x60;category_id&#x60; is a Plaid-specific identifier and does not necessarily correspond to merchant category codes. | 
**Group** | **string** | &#x60;place&#x60; for physical transactions or &#x60;special&#x60; for other transactions such as bank charges. | 
**Hierarchy** | **[]string** | A hierarchical array of the categories to which this &#x60;category_id&#x60; belongs. | 

## Methods

### NewCategory

`func NewCategory(categoryId string, group string, hierarchy []string, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *Category) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Category) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Category) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.


### GetGroup

`func (o *Category) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Category) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Category) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetHierarchy

`func (o *Category) GetHierarchy() []string`

GetHierarchy returns the Hierarchy field if non-nil, zero value otherwise.

### GetHierarchyOk

`func (o *Category) GetHierarchyOk() (*[]string, bool)`

GetHierarchyOk returns a tuple with the Hierarchy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHierarchy

`func (o *Category) SetHierarchy(v []string)`

SetHierarchy sets Hierarchy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


