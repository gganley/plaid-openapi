# AssetReportRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetReportId** | **string** | A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive. | 
**AssetReportToken** | **string** | A token that can be provided to endpoints such as &#x60;/asset_report/get&#x60; or &#x60;/asset_report/pdf/get&#x60; to fetch or update an Asset Report. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewAssetReportRefreshResponse

`func NewAssetReportRefreshResponse(assetReportId string, assetReportToken string, requestId string, ) *AssetReportRefreshResponse`

NewAssetReportRefreshResponse instantiates a new AssetReportRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportRefreshResponseWithDefaults

`func NewAssetReportRefreshResponseWithDefaults() *AssetReportRefreshResponse`

NewAssetReportRefreshResponseWithDefaults instantiates a new AssetReportRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetReportId

`func (o *AssetReportRefreshResponse) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *AssetReportRefreshResponse) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *AssetReportRefreshResponse) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.


### GetAssetReportToken

`func (o *AssetReportRefreshResponse) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportRefreshResponse) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportRefreshResponse) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetRequestId

`func (o *AssetReportRefreshResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssetReportRefreshResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssetReportRefreshResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


