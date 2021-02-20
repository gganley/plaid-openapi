# LinkTokenCreateRequestIncomeVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeVerificationId** | **string** | The &#x60;income_verification_id&#x60; of the verification instance, as provided by &#x60;/income/verification/create&#x60;. | 
**AssetReportId** | Pointer to **NullableString** | The &#x60;asset_report_id&#x60; of an asset report associated with the user, as provided by &#x60;/asset_report/create&#x60;. Providing an &#x60;asset_report_id&#x60; is optional and can be used to verify the user through a streamlined flow. If provided, the bank linking flow will be skipped. | [optional] 

## Methods

### NewLinkTokenCreateRequestIncomeVerification

`func NewLinkTokenCreateRequestIncomeVerification(incomeVerificationId string, ) *LinkTokenCreateRequestIncomeVerification`

NewLinkTokenCreateRequestIncomeVerification instantiates a new LinkTokenCreateRequestIncomeVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestIncomeVerificationWithDefaults

`func NewLinkTokenCreateRequestIncomeVerificationWithDefaults() *LinkTokenCreateRequestIncomeVerification`

NewLinkTokenCreateRequestIncomeVerificationWithDefaults instantiates a new LinkTokenCreateRequestIncomeVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeVerificationId

`func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *LinkTokenCreateRequestIncomeVerification) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.


### GetAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *LinkTokenCreateRequestIncomeVerification) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.

### HasAssetReportId

`func (o *LinkTokenCreateRequestIncomeVerification) HasAssetReportId() bool`

HasAssetReportId returns a boolean if a field has been set.

### SetAssetReportIdNil

`func (o *LinkTokenCreateRequestIncomeVerification) SetAssetReportIdNil(b bool)`

 SetAssetReportIdNil sets the value for AssetReportId to be an explicit nil

### UnsetAssetReportId
`func (o *LinkTokenCreateRequestIncomeVerification) UnsetAssetReportId()`

UnsetAssetReportId ensures that no value is present for AssetReportId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


