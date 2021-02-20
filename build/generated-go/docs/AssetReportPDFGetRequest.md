# AssetReportPDFGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AssetReportToken** | **string** | A token that can be provided to endpoints such as &#x60;/asset_report/get&#x60; or &#x60;/asset_report/pdf/get&#x60; to fetch or update an Asset Report. | 

## Methods

### NewAssetReportPDFGetRequest

`func NewAssetReportPDFGetRequest(assetReportToken string, ) *AssetReportPDFGetRequest`

NewAssetReportPDFGetRequest instantiates a new AssetReportPDFGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportPDFGetRequestWithDefaults

`func NewAssetReportPDFGetRequestWithDefaults() *AssetReportPDFGetRequest`

NewAssetReportPDFGetRequestWithDefaults instantiates a new AssetReportPDFGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportPDFGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportPDFGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportPDFGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportPDFGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportPDFGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportPDFGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportPDFGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportPDFGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReportPDFGetRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportPDFGetRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportPDFGetRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


