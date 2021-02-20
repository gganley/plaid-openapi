# InstitutionsSearchRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oauth** | Pointer to **bool** | Limit results to institutions with or without OAuth login flows. This is primarily relevant to institutions with European country codes | [optional] 
**IncludeOptionalMetadata** | Pointer to **bool** | When true, return the institution&#39;s homepage URL, logo and primary brand color. Learn more | [optional] 
**AccountFilter** | Pointer to [**InstitutionsSearchAccountFilter**](InstitutionsSearchAccountFilter.md) |  | [optional] 

## Methods

### NewInstitutionsSearchRequestOptions

`func NewInstitutionsSearchRequestOptions() *InstitutionsSearchRequestOptions`

NewInstitutionsSearchRequestOptions instantiates a new InstitutionsSearchRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsSearchRequestOptionsWithDefaults

`func NewInstitutionsSearchRequestOptionsWithDefaults() *InstitutionsSearchRequestOptions`

NewInstitutionsSearchRequestOptionsWithDefaults instantiates a new InstitutionsSearchRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauth

`func (o *InstitutionsSearchRequestOptions) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InstitutionsSearchRequestOptions) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InstitutionsSearchRequestOptions) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InstitutionsSearchRequestOptions) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### GetIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadata() bool`

GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field if non-nil, zero value otherwise.

### GetIncludeOptionalMetadataOk

`func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool)`

GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) SetIncludeOptionalMetadata(v bool)`

SetIncludeOptionalMetadata sets IncludeOptionalMetadata field to given value.

### HasIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) HasIncludeOptionalMetadata() bool`

HasIncludeOptionalMetadata returns a boolean if a field has been set.

### GetAccountFilter

`func (o *InstitutionsSearchRequestOptions) GetAccountFilter() InstitutionsSearchAccountFilter`

GetAccountFilter returns the AccountFilter field if non-nil, zero value otherwise.

### GetAccountFilterOk

`func (o *InstitutionsSearchRequestOptions) GetAccountFilterOk() (*InstitutionsSearchAccountFilter, bool)`

GetAccountFilterOk returns a tuple with the AccountFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilter

`func (o *InstitutionsSearchRequestOptions) SetAccountFilter(v InstitutionsSearchAccountFilter)`

SetAccountFilter sets AccountFilter field to given value.

### HasAccountFilter

`func (o *InstitutionsSearchRequestOptions) HasAccountFilter() bool`

HasAccountFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


