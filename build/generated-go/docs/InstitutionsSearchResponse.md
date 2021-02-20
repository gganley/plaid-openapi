# InstitutionsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Institutions** | [**[]Institution**](Institution.md) | An array of institutions matching the search criteria | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewInstitutionsSearchResponse

`func NewInstitutionsSearchResponse(institutions []Institution, requestId string, ) *InstitutionsSearchResponse`

NewInstitutionsSearchResponse instantiates a new InstitutionsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsSearchResponseWithDefaults

`func NewInstitutionsSearchResponseWithDefaults() *InstitutionsSearchResponse`

NewInstitutionsSearchResponseWithDefaults instantiates a new InstitutionsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutions

`func (o *InstitutionsSearchResponse) GetInstitutions() []Institution`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *InstitutionsSearchResponse) GetInstitutionsOk() (*[]Institution, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *InstitutionsSearchResponse) SetInstitutions(v []Institution)`

SetInstitutions sets Institutions field to given value.


### GetRequestId

`func (o *InstitutionsSearchResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InstitutionsSearchResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InstitutionsSearchResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


