# StandaloneCurrencyCodeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCurrencyCode** | **string** | Plaid supports all ISO 4217 currency codes. | 
**UnofficialCurrencyCode** | **string** | List of unofficial currency codes | 

## Methods

### NewStandaloneCurrencyCodeList

`func NewStandaloneCurrencyCodeList(isoCurrencyCode string, unofficialCurrencyCode string, ) *StandaloneCurrencyCodeList`

NewStandaloneCurrencyCodeList instantiates a new StandaloneCurrencyCodeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandaloneCurrencyCodeListWithDefaults

`func NewStandaloneCurrencyCodeListWithDefaults() *StandaloneCurrencyCodeList`

NewStandaloneCurrencyCodeListWithDefaults instantiates a new StandaloneCurrencyCodeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCurrencyCode

`func (o *StandaloneCurrencyCodeList) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *StandaloneCurrencyCodeList) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *StandaloneCurrencyCodeList) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetUnofficialCurrencyCode

`func (o *StandaloneCurrencyCodeList) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *StandaloneCurrencyCodeList) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *StandaloneCurrencyCodeList) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


