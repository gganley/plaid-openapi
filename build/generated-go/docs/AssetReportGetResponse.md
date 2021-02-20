# AssetReportGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Report** | [**AssetReport**](AssetReport.md) |  | 
**Warnings** | [**[]Warning**](Warning.md) | If the Asset Report generation was successful but identity information cannot be returned, this array will contain information about the errors causing identity information to be missing | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewAssetReportGetResponse

`func NewAssetReportGetResponse(report AssetReport, warnings []Warning, requestId string, ) *AssetReportGetResponse`

NewAssetReportGetResponse instantiates a new AssetReportGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportGetResponseWithDefaults

`func NewAssetReportGetResponseWithDefaults() *AssetReportGetResponse`

NewAssetReportGetResponseWithDefaults instantiates a new AssetReportGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReport

`func (o *AssetReportGetResponse) GetReport() AssetReport`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *AssetReportGetResponse) GetReportOk() (*AssetReport, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *AssetReportGetResponse) SetReport(v AssetReport)`

SetReport sets Report field to given value.


### GetWarnings

`func (o *AssetReportGetResponse) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AssetReportGetResponse) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AssetReportGetResponse) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.


### GetRequestId

`func (o *AssetReportGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AssetReportGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AssetReportGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


