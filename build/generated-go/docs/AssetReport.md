# AssetReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetReportId** | **string** | A unique ID identifying an Asset Report. Like all Plaid identifiers, this ID is case sensitive. | 
**ClientReportId** | **string** | An identifier you determine and submit for the Asset Report. | 
**DateGenerated** | **string** | The date and time when the Asset Report was created, in ISO 8601 format (e.g. \&quot;2018-04-12T03:32:11Z\&quot;). | 
**DaysRequested** | **float32** | The duration of transaction history you requested | 
**User** | [**AssetReportUser**](AssetReportUser.md) |  | 
**Items** | [**[]AssetReportItem**](AssetReportItem.md) | Data returned by Plaid about each of the Items included in the Asset Report. | 

## Methods

### NewAssetReport

`func NewAssetReport(assetReportId string, clientReportId string, dateGenerated string, daysRequested float32, user AssetReportUser, items []AssetReportItem, ) *AssetReport`

NewAssetReport instantiates a new AssetReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportWithDefaults

`func NewAssetReportWithDefaults() *AssetReport`

NewAssetReportWithDefaults instantiates a new AssetReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetReportId

`func (o *AssetReport) GetAssetReportId() string`

GetAssetReportId returns the AssetReportId field if non-nil, zero value otherwise.

### GetAssetReportIdOk

`func (o *AssetReport) GetAssetReportIdOk() (*string, bool)`

GetAssetReportIdOk returns a tuple with the AssetReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetReportId

`func (o *AssetReport) SetAssetReportId(v string)`

SetAssetReportId sets AssetReportId field to given value.


### GetClientReportId

`func (o *AssetReport) GetClientReportId() string`

GetClientReportId returns the ClientReportId field if non-nil, zero value otherwise.

### GetClientReportIdOk

`func (o *AssetReport) GetClientReportIdOk() (*string, bool)`

GetClientReportIdOk returns a tuple with the ClientReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReportId

`func (o *AssetReport) SetClientReportId(v string)`

SetClientReportId sets ClientReportId field to given value.


### GetDateGenerated

`func (o *AssetReport) GetDateGenerated() string`

GetDateGenerated returns the DateGenerated field if non-nil, zero value otherwise.

### GetDateGeneratedOk

`func (o *AssetReport) GetDateGeneratedOk() (*string, bool)`

GetDateGeneratedOk returns a tuple with the DateGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateGenerated

`func (o *AssetReport) SetDateGenerated(v string)`

SetDateGenerated sets DateGenerated field to given value.


### GetDaysRequested

`func (o *AssetReport) GetDaysRequested() float32`

GetDaysRequested returns the DaysRequested field if non-nil, zero value otherwise.

### GetDaysRequestedOk

`func (o *AssetReport) GetDaysRequestedOk() (*float32, bool)`

GetDaysRequestedOk returns a tuple with the DaysRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRequested

`func (o *AssetReport) SetDaysRequested(v float32)`

SetDaysRequested sets DaysRequested field to given value.


### GetUser

`func (o *AssetReport) GetUser() AssetReportUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssetReport) GetUserOk() (*AssetReportUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssetReport) SetUser(v AssetReportUser)`

SetUser sets User field to given value.


### GetItems

`func (o *AssetReport) GetItems() []AssetReportItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AssetReport) GetItemsOk() (*[]AssetReportItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AssetReport) SetItems(v []AssetReportItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


