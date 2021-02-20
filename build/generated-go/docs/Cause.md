# Cause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**Error** | [**NullableError**](Error.md) |  | 

## Methods

### NewCause

`func NewCause(itemId string, error_ NullableError, ) *Cause`

NewCause instantiates a new Cause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCauseWithDefaults

`func NewCauseWithDefaults() *Cause`

NewCauseWithDefaults instantiates a new Cause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *Cause) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Cause) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Cause) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetError

`func (o *Cause) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Cause) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Cause) SetError(v Error)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *Cause) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Cause) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


