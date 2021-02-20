# HealthIncident

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to **NullableString** | The start date of the incident, in ISO 8601 format, e.g. &#x60;\&quot;2020-10-30T15:26:48Z\&quot;&#x60;. | [optional] 
**EndDate** | Pointer to **NullableString** | The end date of the incident, in ISO 8601 format, e.g. &#x60;\&quot;2020-10-30T15:26:48Z\&quot;&#x60;. | [optional] 
**Title** | Pointer to **string** | The title of the incident | [optional] 
**IncidentUpdates** | Pointer to [**[]IncidentUpdate**](IncidentUpdate.md) | Updates on the health incident. | [optional] 

## Methods

### NewHealthIncident

`func NewHealthIncident() *HealthIncident`

NewHealthIncident instantiates a new HealthIncident object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthIncidentWithDefaults

`func NewHealthIncidentWithDefaults() *HealthIncident`

NewHealthIncidentWithDefaults instantiates a new HealthIncident object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *HealthIncident) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *HealthIncident) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *HealthIncident) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *HealthIncident) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDateNil

`func (o *HealthIncident) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *HealthIncident) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *HealthIncident) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *HealthIncident) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *HealthIncident) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *HealthIncident) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *HealthIncident) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *HealthIncident) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetTitle

`func (o *HealthIncident) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HealthIncident) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HealthIncident) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HealthIncident) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIncidentUpdates

`func (o *HealthIncident) GetIncidentUpdates() []IncidentUpdate`

GetIncidentUpdates returns the IncidentUpdates field if non-nil, zero value otherwise.

### GetIncidentUpdatesOk

`func (o *HealthIncident) GetIncidentUpdatesOk() (*[]IncidentUpdate, bool)`

GetIncidentUpdatesOk returns a tuple with the IncidentUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentUpdates

`func (o *HealthIncident) SetIncidentUpdates(v []IncidentUpdate)`

SetIncidentUpdates sets IncidentUpdates field to given value.

### HasIncidentUpdates

`func (o *HealthIncident) HasIncidentUpdates() bool`

HasIncidentUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


