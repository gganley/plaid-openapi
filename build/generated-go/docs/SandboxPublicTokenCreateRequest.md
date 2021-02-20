# SandboxPublicTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**InstitutionId** | **string** | The ID of the institution the Item will be associated with | 
**InitialProducts** | [**[]Products**](Products.md) | The products to initially pull for the Item. May be any products that the specified &#x60;institution_id&#x60;  supports. This array may not be empty. | 
**Options** | Pointer to [**SandboxPublicTokenCreateRequestOptions**](SandboxPublicTokenCreateRequestOptions.md) |  | [optional] 

## Methods

### NewSandboxPublicTokenCreateRequest

`func NewSandboxPublicTokenCreateRequest(institutionId string, initialProducts []Products, ) *SandboxPublicTokenCreateRequest`

NewSandboxPublicTokenCreateRequest instantiates a new SandboxPublicTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxPublicTokenCreateRequestWithDefaults

`func NewSandboxPublicTokenCreateRequestWithDefaults() *SandboxPublicTokenCreateRequest`

NewSandboxPublicTokenCreateRequestWithDefaults instantiates a new SandboxPublicTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxPublicTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxPublicTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxPublicTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxPublicTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxPublicTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxPublicTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxPublicTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxPublicTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInstitutionId

`func (o *SandboxPublicTokenCreateRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *SandboxPublicTokenCreateRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *SandboxPublicTokenCreateRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetInitialProducts

`func (o *SandboxPublicTokenCreateRequest) GetInitialProducts() []Products`

GetInitialProducts returns the InitialProducts field if non-nil, zero value otherwise.

### GetInitialProductsOk

`func (o *SandboxPublicTokenCreateRequest) GetInitialProductsOk() (*[]Products, bool)`

GetInitialProductsOk returns a tuple with the InitialProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialProducts

`func (o *SandboxPublicTokenCreateRequest) SetInitialProducts(v []Products)`

SetInitialProducts sets InitialProducts field to given value.


### GetOptions

`func (o *SandboxPublicTokenCreateRequest) GetOptions() SandboxPublicTokenCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SandboxPublicTokenCreateRequest) GetOptionsOk() (*SandboxPublicTokenCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SandboxPublicTokenCreateRequest) SetOptions(v SandboxPublicTokenCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SandboxPublicTokenCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


