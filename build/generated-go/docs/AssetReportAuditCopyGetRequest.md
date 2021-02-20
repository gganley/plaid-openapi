# AssetReportAuditCopyGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AuditCopyToken** | **string** | The &#x60;audit_copy_token&#x60; granting access to the Audit Copy you would like to get. | 

## Methods

### NewAssetReportAuditCopyGetRequest

`func NewAssetReportAuditCopyGetRequest(auditCopyToken string, ) *AssetReportAuditCopyGetRequest`

NewAssetReportAuditCopyGetRequest instantiates a new AssetReportAuditCopyGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAuditCopyGetRequestWithDefaults

`func NewAssetReportAuditCopyGetRequestWithDefaults() *AssetReportAuditCopyGetRequest`

NewAssetReportAuditCopyGetRequestWithDefaults instantiates a new AssetReportAuditCopyGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AssetReportAuditCopyGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AssetReportAuditCopyGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AssetReportAuditCopyGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AssetReportAuditCopyGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *AssetReportAuditCopyGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AssetReportAuditCopyGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AssetReportAuditCopyGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AssetReportAuditCopyGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAuditCopyToken

`func (o *AssetReportAuditCopyGetRequest) GetAuditCopyToken() string`

GetAuditCopyToken returns the AuditCopyToken field if non-nil, zero value otherwise.

### GetAuditCopyTokenOk

`func (o *AssetReportAuditCopyGetRequest) GetAuditCopyTokenOk() (*string, bool)`

GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditCopyToken

`func (o *AssetReportAuditCopyGetRequest) SetAuditCopyToken(v string)`

SetAuditCopyToken sets AuditCopyToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


