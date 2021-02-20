# AssetReportRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AssetReportToken** | **string** | The &#x60;asset_report_token&#x60; returned by the original call to &#x60;/asset_report/create&#x60; | 
**DaysRequested** | Pointer to **int32** | The maximum number of days of history to include in the Asset Report. Must be an integer. If not specified, the value from the original call to &#x60;/asset_report/create&#x60; will be used. | [optional] 
**Options** | Pointer to [**AssetReportRefreshRequestOptions**](AssetReportRefreshRequestOptions.md) |  | [optional] 

## Methods

### NewAssetReportRefreshRequest

`func NewAssetReportRefreshRequest(assetReportToken string, ) *AssetReportRefreshRequest`

NewAssetReportRefreshRequest instantiates a new AssetReportRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportRefreshRequestWithDefaults

`func NewAssetReportRefreshRequestWithDefaults() *AssetReportRefreshRequest`

NewAssetReportRefreshRequestWithDefaults instantiates a new AssetReportRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportRefreshRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportRefreshRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportRefreshRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportRefreshRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportRefreshRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportRefreshRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportRefreshRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportRefreshRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReportRefreshRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportRefreshRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportRefreshRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetDaysRequested

`func (o *AssetReportRefreshRequest) GetDaysRequested() int32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *AssetReportRefreshRequest) GetDaysRequestedOk() (*int32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *AssetReportRefreshRequest) SetDaysRequested(v int32)`

SetDaysRequested sets DaysRequested field to given value.

### HasDaysRequested

`func (o *AssetReportRefreshRequest) HasDaysRequested() bool`

HasDaysRequested returns a boolean if a field has been set.

### GetOptions

`func (o *AssetReportRefreshRequest) GetOptions() AssetReportRefreshRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AssetReportRefreshRequest) GetOptionsOk() (*AssetReportRefreshRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AssetReportRefreshRequest) SetOptions(v AssetReportRefreshRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AssetReportRefreshRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


