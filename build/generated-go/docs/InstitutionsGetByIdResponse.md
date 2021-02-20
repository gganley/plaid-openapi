# InstitutionsGetByIdResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Institution** | [**Institution**](Institution.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewInstitutionsGetByIdResponse

`func NewInstitutionsGetByIdResponse(institution Institution, requestId string, ) *InstitutionsGetByIdResponse`

NewInstitutionsGetByIdResponse instantiates a new InstitutionsGetByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetByIdResponseWithDefaults

`func NewInstitutionsGetByIdResponseWithDefaults() *InstitutionsGetByIdResponse`

NewInstitutionsGetByIdResponseWithDefaults instantiates a new InstitutionsGetByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitution

`func (o *InstitutionsGetByIdResponse) GetInstitution() Institution`

GetInstitution returns the Institution field if non-nil, zero value otherwise.

### GetInstitutionOk

`func (o *InstitutionsGetByIdResponse) GetInstitutionOk() (*Institution, bool)`

GetInstitutionOk returns a tuple with the Institution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitution

`func (o *InstitutionsGetByIdResponse) SetInstitution(v Institution)`

SetInstitution sets Institution field to given value.


### GetRequestId

`func (o *InstitutionsGetByIdResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *InstitutionsGetByIdResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *InstitutionsGetByIdResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


