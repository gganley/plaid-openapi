# AssetsProductReadyWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;ASSETS&#x60; | 
**WebhookCode** | **string** | &#x60;PRODUCT_READY&#x60; | 
**AssetReportId** | **string** | The &#x60;asset_report_id&#x60; that can be provided to &#x60;/asset_report/get&#x60; to retrieve the Asset Report. | 

## Methods

### NewAssetsProductReadyWebhook

`func NewAssetsProductReadyWebhook(webhookType string, webhookCode string, assetReportId string, ) *AssetsProductReadyWebhook`

NewAssetsProductReadyWebhook instantiates a new AssetsProductReadyWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsProductReadyWebhookWithDefaults

`func NewAssetsProductReadyWebhookWithDefaults() *AssetsProductReadyWebhook`

NewAssetsProductReadyWebhookWithDefaults instantiates a new AssetsProductReadyWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *AssetsProductReadyWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *AssetsProductReadyWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *AssetsProductReadyWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *AssetsProductReadyWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *AssetsProductReadyWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *AssetsProductReadyWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetAssetReportId

`func (o *AssetsProductReadyWebhook) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *AssetsProductReadyWebhook) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *AssetsProductReadyWebhook) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


