# EmployersSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Employers** | [**[]Employer**](Employer.md) | A list of employers matching the search criteria. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewEmployersSearchResponse

`func NewEmployersSearchResponse(employers []Employer, requestId string, ) *EmployersSearchResponse`

NewEmployersSearchResponse instantiates a new EmployersSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersSearchResponseWithDefaults

`func NewEmployersSearchResponseWithDefaults() *EmployersSearchResponse`

NewEmployersSearchResponseWithDefaults instantiates a new EmployersSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployers

`func (o *EmployersSearchResponse) GetEmployers() []Employer`

GetEmployers returns the Employers field if non-nil, zero value otherwise.

### GetEmployersOk

`func (o *EmployersSearchResponse) GetEmployersOk() (*[]Employer, bool)`

GetEmployersOk returns a tuple with the Employers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployers

`func (o *EmployersSearchResponse) SetEmployers(v []Employer)`

SetEmployers sets Employers field to given value.


### GetRequestId

`func (o *EmployersSearchResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EmployersSearchResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EmployersSearchResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


