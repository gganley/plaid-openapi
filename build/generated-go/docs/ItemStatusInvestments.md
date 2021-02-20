# ItemStatusInvestments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSuccessfulUpdate** | Pointer to **NullableString** | ISO 8601 timestamp of the last successful investments update for the Item. The status will update each time Plaid successfully connects with the institution, regardless of whether any new data is available in the update. | [optional] 
**LastFailedUpdate** | Pointer to **NullableString** | ISO 8601 timestamp of the last failed investments update for the Item. The status will update each time Plaid fails an attempt to connect with the institution, regardless of whether any new data is available in the update. | [optional] 

## Methods

### NewItemStatusInvestments

`func NewItemStatusInvestments() *ItemStatusInvestments`

NewItemStatusInvestments instantiates a new ItemStatusInvestments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStatusInvestmentsWithDefaults

`func NewItemStatusInvestmentsWithDefaults() *ItemStatusInvestments`

NewItemStatusInvestmentsWithDefaults instantiates a new ItemStatusInvestments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSuccessfulUpdate

`func (o *ItemStatusInvestments) GetLastSuccessfulUpdate() string`

GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateOk

`func (o *ItemStatusInvestments) GetLastSuccessfulUpdateOk() (*string, bool)`

GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdate

`func (o *ItemStatusInvestments) SetLastSuccessfulUpdate(v string)`

SetLastSuccessfulUpdate sets LastSuccessfulUpdate field to given value.

### HasLastSuccessfulUpdate

`func (o *ItemStatusInvestments) HasLastSuccessfulUpdate() bool`

HasLastSuccessfulUpdate returns a boolean if a field has been set.

### SetLastSuccessfulUpdateNil

`func (o *ItemStatusInvestments) SetLastSuccessfulUpdateNil(b bool)`

 SetLastSuccessfulUpdateNil sets the value for LastSuccessfulUpdate to be an explicit nil

### UnsetLastSuccessfulUpdate
`func (o *ItemStatusInvestments) UnsetLastSuccessfulUpdate()`

UnsetLastSuccessfulUpdate ensures that no value is present for LastSuccessfulUpdate, not even an explicit nil
### GetLastFailedUpdate

`func (o *ItemStatusInvestments) GetLastFailedUpdate() string`

GetLastFailedUpdate returns the LastFailedUpdate field if non-nil, zero value otherwise.

### GetLastFailedUpdateOk

`func (o *ItemStatusInvestments) GetLastFailedUpdateOk() (*string, bool)`

GetLastFailedUpdateOk returns a tuple with the LastFailedUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedUpdate

`func (o *ItemStatusInvestments) SetLastFailedUpdate(v string)`

SetLastFailedUpdate sets LastFailedUpdate field to given value.

### HasLastFailedUpdate

`func (o *ItemStatusInvestments) HasLastFailedUpdate() bool`

HasLastFailedUpdate returns a boolean if a field has been set.

### SetLastFailedUpdateNil

`func (o *ItemStatusInvestments) SetLastFailedUpdateNil(b bool)`

 SetLastFailedUpdateNil sets the value for LastFailedUpdate to be an explicit nil

### UnsetLastFailedUpdate
`func (o *ItemStatusInvestments) UnsetLastFailedUpdate()`

UnsetLastFailedUpdate ensures that no value is present for LastFailedUpdate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


