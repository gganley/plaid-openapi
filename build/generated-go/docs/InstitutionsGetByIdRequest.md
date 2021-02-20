# InstitutionsGetByIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**InstitutionId** | **string** | The ID of the institution to get details about | 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard.  | 
**Options** | Pointer to [**InstitutionsGetByIdRequestOptions**](InstitutionsGetByIdRequestOptions.md) |  | [optional] 

## Methods

### NewInstitutionsGetByIdRequest

`func NewInstitutionsGetByIdRequest(institutionId string, countryCodes []CountryCode, ) *InstitutionsGetByIdRequest`

NewInstitutionsGetByIdRequest instantiates a new InstitutionsGetByIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetByIdRequestWithDefaults

`func NewInstitutionsGetByIdRequestWithDefaults() *InstitutionsGetByIdRequest`

NewInstitutionsGetByIdRequestWithDefaults instantiates a new InstitutionsGetByIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InstitutionsGetByIdRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InstitutionsGetByIdRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InstitutionsGetByIdRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InstitutionsGetByIdRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *InstitutionsGetByIdRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InstitutionsGetByIdRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InstitutionsGetByIdRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InstitutionsGetByIdRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetInstitutionId

`func (o *InstitutionsGetByIdRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *InstitutionsGetByIdRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *InstitutionsGetByIdRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetCountryCodes

`func (o *InstitutionsGetByIdRequest) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *InstitutionsGetByIdRequest) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *InstitutionsGetByIdRequest) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.


### GetOptions

`func (o *InstitutionsGetByIdRequest) GetOptions() InstitutionsGetByIdRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstitutionsGetByIdRequest) GetOptionsOk() (*InstitutionsGetByIdRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstitutionsGetByIdRequest) SetOptions(v InstitutionsGetByIdRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstitutionsGetByIdRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


