# InstitutionsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Count** | **int32** | The total number of Institutions to return. | 
**Offset** | **int32** | The number of Institutions to skip. | 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | Specify an array of Plaid-supported country codes this institution supports, using the ISO-3166-1 alpha-2 country code standard.  | 
**Options** | Pointer to [**InstitutionsGetRequestOptions**](InstitutionsGetRequestOptions.md) |  | [optional] 

## Methods

### NewInstitutionsGetRequest

`func NewInstitutionsGetRequest(count int32, offset int32, countryCodes []CountryCode, ) *InstitutionsGetRequest`

NewInstitutionsGetRequest instantiates a new InstitutionsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetRequestWithDefaults

`func NewInstitutionsGetRequestWithDefaults() *InstitutionsGetRequest`

NewInstitutionsGetRequestWithDefaults instantiates a new InstitutionsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InstitutionsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InstitutionsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InstitutionsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InstitutionsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *InstitutionsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InstitutionsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InstitutionsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InstitutionsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCount

`func (o *InstitutionsGetRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *InstitutionsGetRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *InstitutionsGetRequest) SetCount(v int32)`

SetCount sets Count field to given value.


### GetOffset

`func (o *InstitutionsGetRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InstitutionsGetRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InstitutionsGetRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetCountryCodes

`func (o *InstitutionsGetRequest) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *InstitutionsGetRequest) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *InstitutionsGetRequest) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.


### GetOptions

`func (o *InstitutionsGetRequest) GetOptions() InstitutionsGetRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InstitutionsGetRequest) GetOptionsOk() (*InstitutionsGetRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InstitutionsGetRequest) SetOptions(v InstitutionsGetRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InstitutionsGetRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


