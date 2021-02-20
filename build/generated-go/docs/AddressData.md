# AddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The full city name | 
**Region** | Pointer to **NullableString** | The region or state Example: &#x60;\&quot;NC\&quot;&#x60; | [optional] 
**Street** | **string** | The full street address Example: &#x60;\&quot;564 Main Street, APT 15\&quot;&#x60; | 
**PostalCode** | Pointer to **NullableString** | The postal code | [optional] 
**Country** | **string** | The ISO 3166-1 alpha-2 country code | 

## Methods

### NewAddressData

`func NewAddressData(city string, street string, country string, ) *AddressData`

NewAddressData instantiates a new AddressData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDataWithDefaults

`func NewAddressDataWithDefaults() *AddressData`

NewAddressDataWithDefaults instantiates a new AddressData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *AddressData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressData) SetCity(v string)`

SetCity sets City field to given value.


### GetRegion

`func (o *AddressData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddressData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *AddressData) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AddressData) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStreet

`func (o *AddressData) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressData) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressData) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetPostalCode

`func (o *AddressData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *AddressData) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *AddressData) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *AddressData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressData) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


