# AssetReportCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessTokens** | **[]string** | An array of access tokens corresponding to the Items that will be included in the report. The &#x60;assets&#x60; product must have been initialized for the Items during link; the Assets product cannot be added after initialization. | 
**DaysRequested** | **int32** | The maximum integer number of days of history to include in the Asset Report. If using Fannie Mae Day 1 Certainty, &#x60;days_requested&#x60; must be at least 61 for new originations or at least 31 for refinancings. | 
**Options** | Pointer to [**AssetReportCreateRequestOptions**](AssetReportCreateRequestOptions.md) |  | [optional] 

## Methods

### NewAssetReportCreateRequest

`func NewAssetReportCreateRequest(accessTokens []string, daysRequested int32, ) *AssetReportCreateRequest`

NewAssetReportCreateRequest instantiates a new AssetReportCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportCreateRequestWithDefaults

`func NewAssetReportCreateRequestWithDefaults() *AssetReportCreateRequest`

NewAssetReportCreateRequestWithDefaults instantiates a new AssetReportCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessTokens

`func (o *AssetReportCreateRequest) GetAccessTokens() []string`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *AssetReportCreateRequest) GetAccessTokensOk() (*[]string, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *AssetReportCreateRequest) SetAccessTokens(v []string)`

SetAccessTokens sets AccessTokens field to given value.


### GetDaysRequested

`func (o *AssetReportCreateRequest) GetDaysRequested() int32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *AssetReportCreateRequest) GetDaysRequestedOk() (*int32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *AssetReportCreateRequest) SetDaysRequested(v int32)`

SetDaysRequested sets DaysRequested field to given value.


### GetOptions

`func (o *AssetReportCreateRequest) GetOptions() AssetReportCreateRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AssetReportCreateRequest) GetOptionsOk() (*AssetReportCreateRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AssetReportCreateRequest) SetOptions(v AssetReportCreateRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AssetReportCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


