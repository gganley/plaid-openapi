# AssetReportUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientUserId** | Pointer to **NullableString** | An identifier you determine and submit for the user. | [optional] 
**FirstName** | Pointer to **NullableString** | The user&#39;s first name. Required for the Fannie Mae Day 1 Certainty™ program. | [optional] 
**MiddleName** | Pointer to **NullableString** | The user&#39;s middle name | [optional] 
**LastName** | Pointer to **NullableString** | The user&#39;s last name.  Required for the Fannie Mae Day 1 Certainty™ program. | [optional] 
**Ssn** | Pointer to **NullableString** | The user&#39;s Social Security Number. Required for the Fannie Mae Day 1 Certainty™ program.  Format: \&quot;ddd-dd-dddd\&quot; | [optional] 
**PhoneNumber** | Pointer to **NullableString** | The user&#39;s phone number, in E.164 format: +{countrycode}{number}. For example: \&quot;+14151234567\&quot; | [optional] 
**Email** | Pointer to **NullableString** | The user&#39;s email address. | [optional] 

## Methods

### NewAssetReportUser

`func NewAssetReportUser() *AssetReportUser`

NewAssetReportUser instantiates a new AssetReportUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetReportUserWithDefaults

`func NewAssetReportUserWithDefaults() *AssetReportUser`

NewAssetReportUserWithDefaults instantiates a new AssetReportUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientUserId

`func (o *AssetReportUser) GetClientUserId() string`

GetClientUserId returns the ClientUserId field if non-nil, zero value otherwise.

### GetClientUserIdOk

`func (o *AssetReportUser) GetClientUserIdOk() (*string, bool)`

GetClientUserIdOk returns a tuple with the ClientUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUserId

`func (o *AssetReportUser) SetClientUserId(v string)`

SetClientUserId sets ClientUserId field to given value.

### HasClientUserId

`func (o *AssetReportUser) HasClientUserId() bool`

HasClientUserId returns a boolean if a field has been set.

### SetClientUserIdNil

`func (o *AssetReportUser) SetClientUserIdNil(b bool)`

 SetClientUserIdNil sets the value for ClientUserId to be an explicit nil

### UnsetClientUserId
`func (o *AssetReportUser) UnsetClientUserId()`

UnsetClientUserId ensures that no value is present for ClientUserId, not even an explicit nil
### GetFirstName

`func (o *AssetReportUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AssetReportUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AssetReportUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AssetReportUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *AssetReportUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *AssetReportUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetMiddleName

`func (o *AssetReportUser) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *AssetReportUser) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *AssetReportUser) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *AssetReportUser) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *AssetReportUser) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *AssetReportUser) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetLastName

`func (o *AssetReportUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AssetReportUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AssetReportUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AssetReportUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *AssetReportUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *AssetReportUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetSsn

`func (o *AssetReportUser) GetSsn() string`

GetSsn returns the Ssn field if non-nil, zero value otherwise.

### GetSsnOk

`func (o *AssetReportUser) GetSsnOk() (*string, bool)`

GetSsnOk returns a tuple with the Ssn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsn

`func (o *AssetReportUser) SetSsn(v string)`

SetSsn sets Ssn field to given value.

### HasSsn

`func (o *AssetReportUser) HasSsn() bool`

HasSsn returns a boolean if a field has been set.

### SetSsnNil

`func (o *AssetReportUser) SetSsnNil(b bool)`

 SetSsnNil sets the value for Ssn to be an explicit nil

### UnsetSsn
`func (o *AssetReportUser) UnsetSsn()`

UnsetSsn ensures that no value is present for Ssn, not even an explicit nil
### GetPhoneNumber

`func (o *AssetReportUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AssetReportUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AssetReportUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AssetReportUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *AssetReportUser) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *AssetReportUser) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetEmail

`func (o *AssetReportUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AssetReportUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AssetReportUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AssetReportUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *AssetReportUser) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AssetReportUser) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


