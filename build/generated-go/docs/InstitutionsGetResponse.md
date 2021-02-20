# InstitutionsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Institutions** | [**[]Institution**](Institution.md) | A list of Plaid Institution | 
**Total** | **int32** | The number of institutions returned | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewInstitutionsGetResponse

`func NewInstitutionsGetResponse(institutions []Institution, total int32, requestId string, ) *InstitutionsGetResponse`

NewInstitutionsGetResponse instantiates a new InstitutionsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetResponseWithDefaults

`func NewInstitutionsGetResponseWithDefaults() *InstitutionsGetResponse`

NewInstitutionsGetResponseWithDefaults instantiates a new InstitutionsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutions

`func (o *InstitutionsGetResponse) GetInstitutions() []Institution`

GetInstitutions returns the Institutions field if non-nil, zero value otherwise.

### GetInstitutionsOk

`func (o *InstitutionsGetResponse) GetInstitutionsOk() (*[]Institution, bool)`

GetInstitutionsOk returns a tuple with the Institutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutions

`func (o *InstitutionsGetResponse) SetInstitutions(v []Institution)`

SetInstitutions sets Institutions field to given value.


### GetTotal

`func (o *InstitutionsGetResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InstitutionsGetResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InstitutionsGetResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetRequestId

`func (o *InstitutionsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InstitutionsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InstitutionsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


