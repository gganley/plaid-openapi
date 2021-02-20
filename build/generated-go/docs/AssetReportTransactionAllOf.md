# AssetReportTransactionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTransacted** | Pointer to **NullableString** | The date on which the transaction took place, in IS0 8601 format. | [optional] 
**OriginalDescription** | **string** | The string returned by the financial institution to describe the transaction | 

## Methods

### NewAssetReportTransactionAllOf

`func NewAssetReportTransactionAllOf(originalDescription string, ) *AssetReportTransactionAllOf`

NewAssetReportTransactionAllOf instantiates a new AssetReportTransactionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportTransactionAllOfWithDefaults

`func NewAssetReportTransactionAllOfWithDefaults() *AssetReportTransactionAllOf`

NewAssetReportTransactionAllOfWithDefaults instantiates a new AssetReportTransactionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTransacted

`func (o *AssetReportTransactionAllOf) GetDateTransacted() string`

GetDateTransacted returns the DateTransacted field if non-nil, zero value otherwise.

### GetDateTransactedOk

`func (o *AssetReportTransactionAllOf) GetDateTransactedOk() (*string, bool)`

GetDateTransactedOk returns a tuple with the DateTransacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTransacted

`func (o *AssetReportTransactionAllOf) SetDateTransacted(v string)`

SetDateTransacted sets DateTransacted field to given value.

### HasDateTransacted

`func (o *AssetReportTransactionAllOf) HasDateTransacted() bool`

HasDateTransacted returns a boolean if a field has been set.

### SetDateTransactedNil

`func (o *AssetReportTransactionAllOf) SetDateTransactedNil(b bool)`

 SetDateTransactedNil sets the value for DateTransacted to be an explicit nil

### UnsetDateTransacted
`func (o *AssetReportTransactionAllOf) UnsetDateTransacted()`

UnsetDateTransacted ensures that no value is present for DateTransacted, not even an explicit nil
### GetOriginalDescription

`func (o *AssetReportTransactionAllOf) GetOriginalDescription() string`

GetOriginalDescription returns the OriginalDescription field if non-nil, zero value otherwise.

### GetOriginalDescriptionOk

`func (o *AssetReportTransactionAllOf) GetOriginalDescriptionOk() (*string, bool)`

GetOriginalDescriptionOk returns a tuple with the OriginalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalDescription

`func (o *AssetReportTransactionAllOf) SetOriginalDescription(v string)`

SetOriginalDescription sets OriginalDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


