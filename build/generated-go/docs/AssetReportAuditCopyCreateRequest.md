# AssetReportAuditCopyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AssetReportToken** | **string** | A token that can be provided to endpoints such as &#x60;/asset_report/get&#x60; or &#x60;/asset_report/pdf/get&#x60; to fetch or update an Asset Report. | 
**AuditorId** | **string** | The &#x60;auditor_id&#x60; of the third party with whom you would like to share the Asset Report. | 

## Methods

### NewAssetReportAuditCopyCreateRequest

`func NewAssetReportAuditCopyCreateRequest(assetReportToken string, auditorId string, ) *AssetReportAuditCopyCreateRequest`

NewAssetReportAuditCopyCreateRequest instantiates a new AssetReportAuditCopyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAuditCopyCreateRequestWithDefaults

`func NewAssetReportAuditCopyCreateRequestWithDefaults() *AssetReportAuditCopyCreateRequest`

NewAssetReportAuditCopyCreateRequestWithDefaults instantiates a new AssetReportAuditCopyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportAuditCopyCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportAuditCopyCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportAuditCopyCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportAuditCopyCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportAuditCopyCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportAuditCopyCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportAuditCopyCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportAuditCopyCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAssetReportToken

`func (o *AssetReportAuditCopyCreateRequest) GetAssetReportToken() string`

GetAssetReportToken returns the AssetReportToken field if non-nil, zero value otherwise.

### GetAssetReportTokenOk

`func (o *AssetReportAuditCopyCreateRequest) GetAssetReportTokenOk() (*string, bool)`

GetAssetReportTokenOk returns a tuple with the AssetReportToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportToken

`func (o *AssetReportAuditCopyCreateRequest) SetAssetReportToken(v string)`

SetAssetReportToken sets AssetReportToken field to given value.


### GetAuditorId

`func (o *AssetReportAuditCopyCreateRequest) GetAuditorId() string`

GetAuditorId returns the AuditorId field if non-nil, zero value otherwise.

### GetAuditorIdOk

`func (o *AssetReportAuditCopyCreateRequest) GetAuditorIdOk() (*string, bool)`

GetAuditorIdOk returns a tuple with the AuditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditorId

`func (o *AssetReportAuditCopyCreateRequest) SetAuditorId(v string)`

SetAuditorId sets AuditorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


