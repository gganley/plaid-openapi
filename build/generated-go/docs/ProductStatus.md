# ProductStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | &#x60;HEALTHY&#x60;: the majority of requests are successful &#x60;DEGRADED&#x60;: only some requests are successful &#x60;DOWN&#x60;: all requests are failing | 
**LastStatusChange** | **string** | ISO 8601 formatted timestamp of the last status change for the institution. | 
**Breakdown** | [**ProductStatusBreakdown**](ProductStatusBreakdown.md) |  | 

## Methods

### NewProductStatus

`func NewProductStatus(status string, lastStatusChange string, breakdown ProductStatusBreakdown, ) *ProductStatus`

NewProductStatus instantiates a new ProductStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStatusWithDefaults

`func NewProductStatusWithDefaults() *ProductStatus`

NewProductStatusWithDefaults instantiates a new ProductStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ProductStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProductStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProductStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastStatusChange

`func (o *ProductStatus) GetLastStatusChange() string`

GetLastStatusChange returns the LastStatusChange field if non-nil, zero value otherwise.

### GetLastStatusChangeOk

`func (o *ProductStatus) GetLastStatusChangeOk() (*string, bool)`

GetLastStatusChangeOk returns a tuple with the LastStatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusChange

`func (o *ProductStatus) SetLastStatusChange(v string)`

SetLastStatusChange sets LastStatusChange field to given value.


### GetBreakdown

`func (o *ProductStatus) GetBreakdown() ProductStatusBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *ProductStatus) GetBreakdownOk() (*ProductStatusBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *ProductStatus) SetBreakdown(v ProductStatusBreakdown)`

SetBreakdown sets Breakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


