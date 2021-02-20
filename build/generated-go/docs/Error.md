# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | A broad categorization of the error. Safe for programatic use. | 
**ErrorCode** | **string** | The particular error code. Safe for programmatic use. | 
**ErrorMessage** | **string** | A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use. | 
**DisplayMessage** | Pointer to **NullableString** | A user-friendly representation of the error code. &#x60;null&#x60; if the error is not related to user action.  This may change over time and is not safe for programmatic use. | [optional] 
**RequestId** | **string** | A unique identifying the request, to be used for troubleshooting purposes. This field will be omitted in errors provided by webhooks. | 
**Causes** | Pointer to **[]interface{}** | In the Assets product, a request can pertain to more than one Item. If an error is returned for such a request, &#x60;causes&#x60; will return an array of errors containing a breakdown of these errors on the individual Item level, if any can be identified.  &#x60;causes&#x60; will only be provided for the &#x60;error_type&#x60; &#x60;ASSET_REPORT_ERROR&#x60;. | [optional] 
**Status** | Pointer to **NullableFloat32** | The HTTP status code associated with the error. This will only be returned in the response body when the error information is provided via a webhook. | [optional] 
**DocumentationUrl** | Pointer to **NullableString** | The URL of a Plaid documentation page with more information about the error | [optional] 
**SuggestedAction** | Pointer to **NullableString** | Suggested steps for resolving the error | [optional] 

## Methods

### NewError

`func NewError(errorType string, errorCode string, errorMessage string, requestId string, ) *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *Error) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *Error) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *Error) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorCode

`func (o *Error) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Error) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Error) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetErrorMessage

`func (o *Error) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *Error) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *Error) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### GetDisplayMessage

`func (o *Error) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *Error) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *Error) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.

### HasDisplayMessage

`func (o *Error) HasDisplayMessage() bool`

HasDisplayMessage returns a boolean if a field has been set.

### SetDisplayMessageNil

`func (o *Error) SetDisplayMessageNil(b bool)`

 SetDisplayMessageNil sets the value for DisplayMessage to be an explicit nil

### UnsetDisplayMessage
`func (o *Error) UnsetDisplayMessage()`

UnsetDisplayMessage ensures that no value is present for DisplayMessage, not even an explicit nil
### GetRequestId

`func (o *Error) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Error) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Error) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetCauses

`func (o *Error) GetCauses() []interface{}`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *Error) GetCausesOk() (*[]interface{}, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *Error) SetCauses(v []interface{})`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *Error) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetStatus

`func (o *Error) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Error) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Error) SetStatus(v float32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Error) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Error) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Error) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDocumentationUrl

`func (o *Error) GetDocumentationUrl() string`

GetDocumentationUrl returns the DocumentationUrl field if non-nil, zero value otherwise.

### GetDocumentationUrlOk

`func (o *Error) GetDocumentationUrlOk() (*string, bool)`

GetDocumentationUrlOk returns a tuple with the DocumentationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationUrl

`func (o *Error) SetDocumentationUrl(v string)`

SetDocumentationUrl sets DocumentationUrl field to given value.

### HasDocumentationUrl

`func (o *Error) HasDocumentationUrl() bool`

HasDocumentationUrl returns a boolean if a field has been set.

### SetDocumentationUrlNil

`func (o *Error) SetDocumentationUrlNil(b bool)`

 SetDocumentationUrlNil sets the value for DocumentationUrl to be an explicit nil

### UnsetDocumentationUrl
`func (o *Error) UnsetDocumentationUrl()`

UnsetDocumentationUrl ensures that no value is present for DocumentationUrl, not even an explicit nil
### GetSuggestedAction

`func (o *Error) GetSuggestedAction() string`

GetSuggestedAction returns the SuggestedAction field if non-nil, zero value otherwise.

### GetSuggestedActionOk

`func (o *Error) GetSuggestedActionOk() (*string, bool)`

GetSuggestedActionOk returns a tuple with the SuggestedAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedAction

`func (o *Error) SetSuggestedAction(v string)`

SetSuggestedAction sets SuggestedAction field to given value.

### HasSuggestedAction

`func (o *Error) HasSuggestedAction() bool`

HasSuggestedAction returns a boolean if a field has been set.

### SetSuggestedActionNil

`func (o *Error) SetSuggestedActionNil(b bool)`

 SetSuggestedActionNil sets the value for SuggestedAction to be an explicit nil

### UnsetSuggestedAction
`func (o *Error) UnsetSuggestedAction()`

UnsetSuggestedAction ensures that no value is present for SuggestedAction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


