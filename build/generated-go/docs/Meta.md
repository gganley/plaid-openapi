# Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The account&#39;s name | 
**OfficialName** | **string** | The account&#39;s official name | 
**Limit** | **float32** | The account&#39;s limit | 

## Methods

### NewMeta

`func NewMeta(name string, officialName string, limit float32, ) *Meta`

NewMeta instantiates a new Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetaWithDefaults

`func NewMetaWithDefaults() *Meta`

NewMetaWithDefaults instantiates a new Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Meta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Meta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Meta) SetName(v string)`

SetName sets Name field to given value.


### GetOfficialName

`func (o *Meta) GetOfficialName() string`

GetOfficialName returns the OfficialName field if non-nil, zero value otherwise.

### GetOfficialNameOk

`func (o *Meta) GetOfficialNameOk() (*string, bool)`

GetOfficialNameOk returns a tuple with the OfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficialName

`func (o *Meta) SetOfficialName(v string)`

SetOfficialName sets OfficialName field to given value.


### GetLimit

`func (o *Meta) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Meta) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Meta) SetLimit(v float32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


