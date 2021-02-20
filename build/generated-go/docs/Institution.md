# Institution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstitutionId** | **string** | Unique identifier for the institution | 
**Name** | **string** | The official name of the institution | 
**Products** | [**[]Products**](Products.md) | A list of the Plaid products supported by the institution | 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | A list of the country codes supported by the institution. | 
**Url** | Pointer to **NullableString** | The URL for the institution&#39;s website | [optional] 
**PrimaryColor** | Pointer to **NullableString** | Hexadecimal representation of the primary color used by the institution | [optional] 
**Logo** | Pointer to **NullableString** | Base64 encoded representation of the institution&#39;s logo | [optional] 
**RoutingNumbers** | Pointer to **[]string** | A partial list of routing numbers associated with the institution. This list is provided for the purpose of looking up institutions by routing number. It is not comprehensive and should never be used as a complete list of routing numbers for an institution. | [optional] 
**Oauth** | **bool** | Indicates that the institution has an OAuth login flow. This is primarily relevant to institutions with European country codes. | 
**Status** | Pointer to [**InstitutionStatus**](InstitutionStatus.md) |  | [optional] 

## Methods

### NewInstitution

`func NewInstitution(institutionId string, name string, products []Products, countryCodes []CountryCode, oauth bool, ) *Institution`

NewInstitution instantiates a new Institution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionWithDefaults

`func NewInstitutionWithDefaults() *Institution`

NewInstitutionWithDefaults instantiates a new Institution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstitutionId

`func (o *Institution) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *Institution) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *Institution) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.


### GetName

`func (o *Institution) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Institution) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Institution) SetName(v string)`

SetName sets Name field to given value.


### GetProducts

`func (o *Institution) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *Institution) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *Institution) SetProducts(v []Products)`

SetProducts sets Products field to given value.


### GetCountryCodes

`func (o *Institution) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *Institution) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *Institution) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.


### GetUrl

`func (o *Institution) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Institution) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Institution) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Institution) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Institution) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Institution) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetPrimaryColor

`func (o *Institution) GetPrimaryColor() string`

GetPrimaryColor returns the PrimaryColor field if non-nil, zero value otherwise.

### GetPrimaryColorOk

`func (o *Institution) GetPrimaryColorOk() (*string, bool)`

GetPrimaryColorOk returns a tuple with the PrimaryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryColor

`func (o *Institution) SetPrimaryColor(v string)`

SetPrimaryColor sets PrimaryColor field to given value.

### HasPrimaryColor

`func (o *Institution) HasPrimaryColor() bool`

HasPrimaryColor returns a boolean if a field has been set.

### SetPrimaryColorNil

`func (o *Institution) SetPrimaryColorNil(b bool)`

 SetPrimaryColorNil sets the value for PrimaryColor to be an explicit nil

### UnsetPrimaryColor
`func (o *Institution) UnsetPrimaryColor()`

UnsetPrimaryColor ensures that no value is present for PrimaryColor, not even an explicit nil
### GetLogo

`func (o *Institution) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Institution) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Institution) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *Institution) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *Institution) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *Institution) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetRoutingNumbers

`func (o *Institution) GetRoutingNumbers() []string`

GetRoutingNumbers returns the RoutingNumbers field if non-nil, zero value otherwise.

### GetRoutingNumbersOk

`func (o *Institution) GetRoutingNumbersOk() (*[]string, bool)`

GetRoutingNumbersOk returns a tuple with the RoutingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumbers

`func (o *Institution) SetRoutingNumbers(v []string)`

SetRoutingNumbers sets RoutingNumbers field to given value.

### HasRoutingNumbers

`func (o *Institution) HasRoutingNumbers() bool`

HasRoutingNumbers returns a boolean if a field has been set.

### SetRoutingNumbersNil

`func (o *Institution) SetRoutingNumbersNil(b bool)`

 SetRoutingNumbersNil sets the value for RoutingNumbers to be an explicit nil

### UnsetRoutingNumbers
`func (o *Institution) UnsetRoutingNumbers()`

UnsetRoutingNumbers ensures that no value is present for RoutingNumbers, not even an explicit nil
### GetOauth

`func (o *Institution) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *Institution) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *Institution) SetOauth(v bool)`

SetOauth sets Oauth field to given value.


### GetStatus

`func (o *Institution) GetStatus() InstitutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Institution) GetStatusOk() (*InstitutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Institution) SetStatus(v InstitutionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Institution) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


