# AssetReportFilterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AssetReportToken** | **string** | A token that can be provided to endpoints such as &#x60;/asset_report/get&#x60; or &#x60;/asset_report/pdf/get&#x60; to fetch or update an Asset Report. | 
**AccountIdsToExclude** | **[]string** | The accounts to exclude from the Asset Report, identified by &#x60;account_id&#x60;. | 

## Methods

### NewAssetReportFilterRequest

`func NewAssetReportFilterRequest(assetReportToken string, accountIdsToExclude []string, ) *AssetReportFilterRequest`

NewAssetReportFilterRequest instantiates a new AssetReportFilterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportFilterRequestWithDefaults

`func NewAssetReportFilterRequestWithDefaults() *AssetReportFilterRequest`

NewAssetReportFilterRequestWithDefaults instantiates a new AssetReportFilterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportFilterRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportFilterRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportFilterRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportFilterRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportFilterRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportFilterRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportFilterRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportFilterRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReportFilterRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportFilterRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportFilterRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetAccountIdsToExclude

`func (o *AssetReportFilterRequest) GetAccountIdsToExclude() []string`

GetAccountIdsToExclude returns the AccountIdsToExclude field if non-nil, zero value otherwise.

### GetAccountIdsToExcludeOk

`func (o *AssetReportFilterRequest) GetAccountIdsToExcludeOk() (*[]string, bool)`

GetAccountIdsToExcludeOk returns a tuple with the AccountIdsToExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdsToExclude

`func (o *AssetReportFilterRequest) SetAccountIdsToExclude(v []string)`

SetAccountIdsToExclude sets AccountIdsToExclude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


