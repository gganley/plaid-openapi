# ProductStatusBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **float32** | The percentage of login attempts that are successful, expressed as a decimal. | 
**ErrorPlaid** | **float32** | The percentage of logins that are failing due to an internal Plaid issue, expressed as a decimal.  | 
**ErrorInstitution** | **float32** | The percentage of logins that are failing due to an issue in the institution&#39;s system, expressed as a decimal. | 
**RefreshInterval** | Pointer to **NullableString** | The &#x60;refresh_interval&#x60; may be &#x60;DELAYED&#x60; or &#x60;STOPPED&#x60; even when the success rate is high. This value is only returned for Transactions status breakdowns. | [optional] 

## Methods

### NewProductStatusBreakdown

`func NewProductStatusBreakdown(success float32, errorPlaid float32, errorInstitution float32, ) *ProductStatusBreakdown`

NewProductStatusBreakdown instantiates a new ProductStatusBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductStatusBreakdownWithDefaults

`func NewProductStatusBreakdownWithDefaults() *ProductStatusBreakdown`

NewProductStatusBreakdownWithDefaults instantiates a new ProductStatusBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProductStatusBreakdown) GetSuccess() float32`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProductStatusBreakdown) GetSuccessOk() (*float32, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProductStatusBreakdown) SetSuccess(v float32)`

SetSuccess sets Success field to given value.


### GetErrorPlaid

`func (o *ProductStatusBreakdown) GetErrorPlaid() float32`

GetErrorPlaid returns the ErrorPlaid field if non-nil, zero value otherwise.

### GetErrorPlaidOk

`func (o *ProductStatusBreakdown) GetErrorPlaidOk() (*float32, bool)`

GetErrorPlaidOk returns a tuple with the ErrorPlaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorPlaid

`func (o *ProductStatusBreakdown) SetErrorPlaid(v float32)`

SetErrorPlaid sets ErrorPlaid field to given value.


### GetErrorInstitution

`func (o *ProductStatusBreakdown) GetErrorInstitution() float32`

GetErrorInstitution returns the ErrorInstitution field if non-nil, zero value otherwise.

### GetErrorInstitutionOk

`func (o *ProductStatusBreakdown) GetErrorInstitutionOk() (*float32, bool)`

GetErrorInstitutionOk returns a tuple with the ErrorInstitution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorInstitution

`func (o *ProductStatusBreakdown) SetErrorInstitution(v float32)`

SetErrorInstitution sets ErrorInstitution field to given value.


### GetRefreshInterval

`func (o *ProductStatusBreakdown) GetRefreshInterval() string`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *ProductStatusBreakdown) GetRefreshIntervalOk() (*string, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *ProductStatusBreakdown) SetRefreshInterval(v string)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *ProductStatusBreakdown) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### SetRefreshIntervalNil

`func (o *ProductStatusBreakdown) SetRefreshIntervalNil(b bool)`

 SetRefreshIntervalNil sets the value for RefreshInterval to be an explicit nil

### UnsetRefreshInterval
`func (o *ProductStatusBreakdown) UnsetRefreshInterval()`

UnsetRefreshInterval ensures that no value is present for RefreshInterval, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


