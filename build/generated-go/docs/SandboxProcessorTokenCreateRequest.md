# SandboxProcessorTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**InstitutionId** | **string** | The ID of the institution the Item will be associated with | 
**Options** | Pointer to [**SandboxProcessorTokenCreateRequestOptions**](SandboxProcessorTokenCreateRequestOptions.md) |  | [optional] 

## Methods

### NewSandboxProcessorTokenCreateRequest

`func NewSandboxProcessorTokenCreateRequest(institutionId string, ) *SandboxProcessorTokenCreateRequest`

NewSandboxProcessorTokenCreateRequest instantiates a new SandboxProcessorTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxProcessorTokenCreateRequestWithDefaults

`func NewSandboxProcessorTokenCreateRequestWithDefaults() *SandboxProcessorTokenCreateRequest`

NewSandboxProcessorTokenCreateRequestWithDefaults instantiates a new SandboxProcessorTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxProcessorTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxProcessorTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxProcessorTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxProcessorTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxProcessorTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxProcessorTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxProcessorTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxProcessorTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInstitutionId

`func (o *SandboxProcessorTokenCreateRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *SandboxProcessorTokenCreateRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *SandboxProcessorTokenCreateRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetOptions

`func (o *SandboxProcessorTokenCreateRequest) GetOptions() SandboxProcessorTokenCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SandboxProcessorTokenCreateRequest) GetOptionsOk() (*SandboxProcessorTokenCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SandboxProcessorTokenCreateRequest) SetOptions(v SandboxProcessorTokenCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SandboxProcessorTokenCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


