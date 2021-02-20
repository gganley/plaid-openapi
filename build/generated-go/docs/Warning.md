# Warning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WarningType** | **string** | The warning type, which will always be &#x60;ASSET_REPORT_WARNING&#x60; | 
**WarningCode** | **string** | The warning code identifies a specific kind of warning. Currently, the only possible warning code is &#x60;OWNERS_UNAVAILABLE&#x60;, which indicates that account-owner information is not available. | 
**Cause** | [**Cause**](Cause.md) |  | 

## Methods

### NewWarning

`func NewWarning(warningType string, warningCode string, cause Cause, ) *Warning`

NewWarning instantiates a new Warning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningWithDefaults

`func NewWarningWithDefaults() *Warning`

NewWarningWithDefaults instantiates a new Warning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarningType

`func (o *Warning) GetWarningType() string`

GetWarningType returns the WarningType field if non-nil, zero value otherwise.

### GetWarningTypeOk

`func (o *Warning) GetWarningTypeOk() (*string, bool)`

GetWarningTypeOk returns a tuple with the WarningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningType

`func (o *Warning) SetWarningType(v string)`

SetWarningType sets WarningType field to given value.


### GetWarningCode

`func (o *Warning) GetWarningCode() string`

GetWarningCode returns the WarningCode field if non-nil, zero value otherwise.

### GetWarningCodeOk

`func (o *Warning) GetWarningCodeOk() (*string, bool)`

GetWarningCodeOk returns a tuple with the WarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningCode

`func (o *Warning) SetWarningCode(v string)`

SetWarningCode sets WarningCode field to given value.


### GetCause

`func (o *Warning) GetCause() Cause`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *Warning) GetCauseOk() (*Cause, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *Warning) SetCause(v Cause)`

SetCause sets Cause field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


