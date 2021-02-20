# ServicerAddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **NullableString** | The full city name | [optional] 
**Region** | Pointer to **NullableString** | The region or state Example: &#x60;\&quot;NC\&quot;&#x60; | [optional] 
**Street** | Pointer to **NullableString** | The full street address Example: &#x60;\&quot;564 Main Street, APT 15\&quot;&#x60; | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code | [optional] 

## Methods

### NewServicerAddressData

`func NewServicerAddressData() *ServicerAddressData`

NewServicerAddressData instantiates a new ServicerAddressData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicerAddressDataWithDefaults

`func NewServicerAddressDataWithDefaults() *ServicerAddressData`

NewServicerAddressDataWithDefaults instantiates a new ServicerAddressData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *ServicerAddressData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ServicerAddressData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ServicerAddressData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ServicerAddressData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *ServicerAddressData) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *ServicerAddressData) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *ServicerAddressData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ServicerAddressData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ServicerAddressData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ServicerAddressData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *ServicerAddressData) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *ServicerAddressData) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStreet

`func (o *ServicerAddressData) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *ServicerAddressData) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *ServicerAddressData) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *ServicerAddressData) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### SetStreetNil

`func (o *ServicerAddressData) SetStreetNil(b bool)`

 SetStreetNil sets the value for Street to be an explicit nil

### UnsetStreet
`func (o *ServicerAddressData) UnsetStreet()`

UnsetStreet ensures that no value is present for Street, not even an explicit nil
### GetPostalCode

`func (o *ServicerAddressData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ServicerAddressData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ServicerAddressData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *ServicerAddressData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *ServicerAddressData) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *ServicerAddressData) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *ServicerAddressData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ServicerAddressData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ServicerAddressData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ServicerAddressData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *ServicerAddressData) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *ServicerAddressData) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


