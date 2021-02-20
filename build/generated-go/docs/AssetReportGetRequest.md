# AssetReportGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AssetReportToken** | **string** | A token that can be provided to endpoints such as &#x60;/asset_report/get&#x60; or &#x60;/asset_report/pdf/get&#x60; to fetch or update an Asset Report. | 
**IncludeInsights** | Pointer to **bool** | &#x60;true&#x60; if you would like to retrieve the Asset Report with Insights, &#x60;false&#x60; otherwise. This field defaults to &#x60;false&#x60; if omitted. [Contact Plaid Support](https://dashboard.plaid.com/support) to get access to this feature. | [optional] 

## Methods

### NewAssetReportGetRequest

`func NewAssetReportGetRequest(assetReportToken string, ) *AssetReportGetRequest`

NewAssetReportGetRequest instantiates a new AssetReportGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportGetRequestWithDefaults

`func NewAssetReportGetRequestWithDefaults() *AssetReportGetRequest`

NewAssetReportGetRequestWithDefaults instantiates a new AssetReportGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReportGetRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportGetRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportGetRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetIncludeInsights

`func (o *AssetReportGetRequest) GetIncludeInsights() bool`

GetIncludeInsights returns the IncludeInsights field if non-nil, zero value otherwise.

### GetIncludeInsightsOk

`func (o *AssetReportGetRequest) GetIncludeInsightsOk() (*bool, bool)`

GetIncludeInsightsOk returns a tuple with the IncludeInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInsights

`func (o *AssetReportGetRequest) SetIncludeInsights(v bool)`

SetIncludeInsights sets IncludeInsights field to given value.

### HasIncludeInsights

`func (o *AssetReportGetRequest) HasIncludeInsights() bool`

HasIncludeInsights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


