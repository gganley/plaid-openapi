# AssetReportAuditCopyCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditCopyToken** | **string** | A token that can be shared with a third party auditor to allow them to obtain access to the Asset Report. This token should be stored securely. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewAssetReportAuditCopyCreateResponse

`func NewAssetReportAuditCopyCreateResponse(auditCopyToken string, requestId string, ) *AssetReportAuditCopyCreateResponse`

NewAssetReportAuditCopyCreateResponse instantiates a new AssetReportAuditCopyCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportAuditCopyCreateResponseWithDefaults

`func NewAssetReportAuditCopyCreateResponseWithDefaults() *AssetReportAuditCopyCreateResponse`

NewAssetReportAuditCopyCreateResponseWithDefaults instantiates a new AssetReportAuditCopyCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditCopyToken

`func (o *AssetReportAuditCopyCreateResponse) GetAuditCopyToken() string`

GetAuditCopyToken returns the AuditCopyToken field if non-nil, zero value otherwise.

### GetAuditCopyTokenOk

`func (o *AssetReportAuditCopyCreateResponse) GetAuditCopyTokenOk() (*string, bool)`

GetAuditCopyTokenOk returns a tuple with the AuditCopyToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditCopyToken

`func (o *AssetReportAuditCopyCreateResponse) SetAuditCopyToken(v string)`

SetAuditCopyToken sets AuditCopyToken field to given value.


### GetRequestId

`func (o *AssetReportAuditCopyCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssetReportAuditCopyCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssetReportAuditCopyCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


