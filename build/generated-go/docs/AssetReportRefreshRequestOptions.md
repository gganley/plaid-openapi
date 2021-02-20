# AssetReportRefreshRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientReportId** | Pointer to **string** | Client-generated identifier, which can be used by lenders to track loan applications. | [optional] 
**Webhook** | Pointer to **string** | URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready. | [optional] 
**User** | Pointer to [**AssetReportUser**](AssetReportUser.md) |  | [optional] 

## Methods

### NewAssetReportRefreshRequestOptions

`func NewAssetReportRefreshRequestOptions() *AssetReportRefreshRequestOptions`

NewAssetReportRefreshRequestOptions instantiates a new AssetReportRefreshRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportRefreshRequestOptionsWithDefaults

`func NewAssetReportRefreshRequestOptionsWithDefaults() *AssetReportRefreshRequestOptions`

NewAssetReportRefreshRequestOptionsWithDefaults instantiates a new AssetReportRefreshRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientReportId

`func (o *AssetReportRefreshRequestOptions) GetClientReportId() string`

GetClientReportId returns the ClientReportId field if non-nil, zero value otherwise.

### GetClientReportIdOk

`func (o *AssetReportRefreshRequestOptions) GetClientReportIdOk() (*string, bool)`

GetClientReportIdOk returns a tuple with the ClientReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReportId

`func (o *AssetReportRefreshRequestOptions) SetClientReportId(v string)`

SetClientReportId sets ClientReportId field to given value.

### HasClientReportId

`func (o *AssetReportRefreshRequestOptions) HasClientReportId() bool`

HasClientReportId returns a boolean if a field has been set.

### GetWebhook

`func (o *AssetReportRefreshRequestOptions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AssetReportRefreshRequestOptions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AssetReportRefreshRequestOptions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *AssetReportRefreshRequestOptions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetUser

`func (o *AssetReportRefreshRequestOptions) GetUser() AssetReportUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssetReportRefreshRequestOptions) GetUserOk() (*AssetReportUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssetReportRefreshRequestOptions) SetUser(v AssetReportUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AssetReportRefreshRequestOptions) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


