# AssetReportCreateRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientReportId** | Pointer to **string** | Client-generated identifier, which can be used by lenders to track loan applications. | [optional] 
**Webhook** | Pointer to **string** | URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready. | [optional] 
**User** | Pointer to [**AssetReportUser**](AssetReportUser.md) |  | [optional] 

## Methods

### NewAssetReportCreateRequestOptions

`func NewAssetReportCreateRequestOptions() *AssetReportCreateRequestOptions`

NewAssetReportCreateRequestOptions instantiates a new AssetReportCreateRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportCreateRequestOptionsWithDefaults

`func NewAssetReportCreateRequestOptionsWithDefaults() *AssetReportCreateRequestOptions`

NewAssetReportCreateRequestOptionsWithDefaults instantiates a new AssetReportCreateRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientReportId

`func (o *AssetReportCreateRequestOptions) GetClientReportId() string`

GetClientReportId returns the ClientReportId field if non-nil, zero value otherwise.

### GetClientReportIdOk

`func (o *AssetReportCreateRequestOptions) GetClientReportIdOk() (*string, bool)`

GetClientReportIdOk returns a tuple with the ClientReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReportId

`func (o *AssetReportCreateRequestOptions) SetClientReportId(v string)`

SetClientReportId sets ClientReportId field to given value.

### HasClientReportId

`func (o *AssetReportCreateRequestOptions) HasClientReportId() bool`

HasClientReportId returns a boolean if a field has been set.

### GetWebhook

`func (o *AssetReportCreateRequestOptions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AssetReportCreateRequestOptions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AssetReportCreateRequestOptions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *AssetReportCreateRequestOptions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetUser

`func (o *AssetReportCreateRequestOptions) GetUser() AssetReportUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssetReportCreateRequestOptions) GetUserOk() (*AssetReportUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssetReportCreateRequestOptions) SetUser(v AssetReportUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AssetReportCreateRequestOptions) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


