# IncidentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The content of the update. | [optional] 
**Status** | Pointer to **string** | The status of the incident. | [optional] 
**UpdatedDate** | Pointer to **string** | The date when the update was published, in ISO 8601 format, e.g. &#x60;\&quot;2020-10-30T15:26:48Z\&quot;&#x60;. | [optional] 

## Methods

### NewIncidentUpdate

`func NewIncidentUpdate() *IncidentUpdate`

NewIncidentUpdate instantiates a new IncidentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentUpdateWithDefaults

`func NewIncidentUpdateWithDefaults() *IncidentUpdate`

NewIncidentUpdateWithDefaults instantiates a new IncidentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IncidentUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IncidentUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IncidentUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IncidentUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *IncidentUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IncidentUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IncidentUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IncidentUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *IncidentUpdate) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *IncidentUpdate) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *IncidentUpdate) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *IncidentUpdate) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


