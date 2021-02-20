# AssetReportItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | The &#x60;item_id&#x60; of the Item associated with this webhook, warning, or error | 
**InstitutionName** | **string** | The full financial institution name associated with the Item. | 
**InstitutionId** | **string** | The id of the financial institution associated with the Item. | 
**DateLastUpdated** | **string** | The date and time when this Item’s data was last retrieved from the financial institution, in ISO 8601 format. | 
**Accounts** | [**[]AccountAssets**](AccountAssets.md) | Data about each of the accounts open on the Item. | 

## Methods

### NewAssetReportItem

`func NewAssetReportItem(itemId string, institutionName string, institutionId string, dateLastUpdated string, accounts []AccountAssets, ) *AssetReportItem`

NewAssetReportItem instantiates a new AssetReportItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportItemWithDefaults

`func NewAssetReportItemWithDefaults() *AssetReportItem`

NewAssetReportItemWithDefaults instantiates a new AssetReportItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *AssetReportItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *AssetReportItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *AssetReportItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetInstitutionName

`func (o *AssetReportItem) GetInstitutionName() string`

GetInstitutionName returns the InstitutionName field if non-nil, zero value otherwise.

### GetInstitutionNameOk

`func (o *AssetReportItem) GetInstitutionNameOk() (*string, bool)`

GetInstitutionNameOk returns a tuple with the InstitutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionName

`func (o *AssetReportItem) SetInstitutionName(v string)`

SetInstitutionName sets InstitutionName field to given value.


### GetInstitutionId

`func (o *AssetReportItem) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *AssetReportItem) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *AssetReportItem) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetDateLastUpdated

`func (o *AssetReportItem) GetDateLastUpdated() string`

GetDateLastUpdated returns the DateLastUpdated field if non-nil, zero value otherwise.

### GetDateLastUpdatedOk

`func (o *AssetReportItem) GetDateLastUpdatedOk() (*string, bool)`

GetDateLastUpdatedOk returns a tuple with the DateLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateLastUpdated

`func (o *AssetReportItem) SetDateLastUpdated(v string)`

SetDateLastUpdated sets DateLastUpdated field to given value.


### GetAccounts

`func (o *AssetReportItem) GetAccounts() []AccountAssets`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AssetReportItem) GetAccountsOk() (*[]AccountAssets, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AssetReportItem) SetAccounts(v []AccountAssets)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


