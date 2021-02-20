# UserCustomPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **NullableString** | The version of the password schema to use, possible values are 1 or 2. The default value is 2. You should only specify 1 if you know it is necessary for your test suite. | [optional] 
**Seed** | **string** | A seed, in the form of a string, that will be used to randomly generate account and transaction data, if this data is not specified using the &#x60;override_accounts&#x60; argument. If no seed is specified, the randomly generated data will be different each time.  Note that transactions data is generated relative to the Item&#39;s creation date. Different Items created on different dates with the same seed for transactions data will have different dates for the transactions. The number of days between each transaction and the Item creation will remain constant. For example, an Item created on December 15 might show a transaction on December 14. An Item created on December 20, using the same seed, would show that same transaction occurring on December 19. | 
**OverrideAccounts** | [**[]OverrideAccounts**](OverrideAccounts.md) | An array of account overrides to configure the accounts for the Item. By default, if no override is specified, transactions and account data will be randomly generated based on the account type and subtype, and other products will have fixed or empty data. | 
**Mfa** | [**MFA**](MFA.md) |  | 
**Recaptcha** | **string** | You may trigger a reCAPTCHA in Plaid Link in the Sandbox environment by using the recaptcha field. Possible values are &#x60;good&#x60; or &#x60;bad&#x60;. A value of &#x60;good&#x60; will result in successful Item creation and &#x60;bad&#x60; will result in a &#x60;RECAPTCHA_BAD&#x60; error to simulate a failed reCAPTCHA. Both values require the reCAPTCHA to be manually solved within Plaid Link. | 
**ForceError** | **string** | An error code to force on Item creation. Possible values are:  &#x60;\&quot;INSTITUTION_NOT_RESPONDING\&quot;&#x60; &#x60;\&quot;INSTITUTION_NO_LONGER_SUPPORTED\&quot;&#x60; &#x60;\&quot;INVALID_CREDENTIALS\&quot;&#x60; &#x60;\&quot;INVALID_MFA\&quot;&#x60; &#x60;\&quot;ITEM_LOCKED\&quot;&#x60; &#x60;\&quot;ITEM_LOGIN_REQUIRED\&quot;&#x60; &#x60;\&quot;ITEM_NOT_SUPPORTED\&quot;&#x60; &#x60;\&quot;INVALID_LINK_TOKEN\&quot;&#x60; &#x60;\&quot;MFA_NOT_SUPPORTED\&quot;&#x60; &#x60;\&quot;NO_ACCOUNTS\&quot;&#x60; &#x60;\&quot;PLAID_ERROR\&quot;&#x60; &#x60;\&quot;PRODUCTS_NOT_SUPPORTED\&quot;&#x60; &#x60;\&quot;USER_SETUP_REQUIRED\&quot;&#x60; | 

## Methods

### NewUserCustomPassword

`func NewUserCustomPassword(seed string, overrideAccounts []OverrideAccounts, mfa MFA, recaptcha string, forceError string, ) *UserCustomPassword`

NewUserCustomPassword instantiates a new UserCustomPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCustomPasswordWithDefaults

`func NewUserCustomPasswordWithDefaults() *UserCustomPassword`

NewUserCustomPasswordWithDefaults instantiates a new UserCustomPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UserCustomPassword) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserCustomPassword) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserCustomPassword) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserCustomPassword) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *UserCustomPassword) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *UserCustomPassword) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetSeed

`func (o *UserCustomPassword) GetSeed() string`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *UserCustomPassword) GetSeedOk() (*string, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *UserCustomPassword) SetSeed(v string)`

SetSeed sets Seed field to given value.


### GetOverrideAccounts

`func (o *UserCustomPassword) GetOverrideAccounts() []OverrideAccounts`

GetOverrideAccounts returns the OverrideAccounts field if non-nil, zero value otherwise.

### GetOverrideAccountsOk

`func (o *UserCustomPassword) GetOverrideAccountsOk() (*[]OverrideAccounts, bool)`

GetOverrideAccountsOk returns a tuple with the OverrideAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAccounts

`func (o *UserCustomPassword) SetOverrideAccounts(v []OverrideAccounts)`

SetOverrideAccounts sets OverrideAccounts field to given value.


### GetMfa

`func (o *UserCustomPassword) GetMfa() MFA`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *UserCustomPassword) GetMfaOk() (*MFA, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *UserCustomPassword) SetMfa(v MFA)`

SetMfa sets Mfa field to given value.


### GetRecaptcha

`func (o *UserCustomPassword) GetRecaptcha() string`

GetRecaptcha returns the Recaptcha field if non-nil, zero value otherwise.

### GetRecaptchaOk

`func (o *UserCustomPassword) GetRecaptchaOk() (*string, bool)`

GetRecaptchaOk returns a tuple with the Recaptcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecaptcha

`func (o *UserCustomPassword) SetRecaptcha(v string)`

SetRecaptcha sets Recaptcha field to given value.


### GetForceError

`func (o *UserCustomPassword) GetForceError() string`

GetForceError returns the ForceError field if non-nil, zero value otherwise.

### GetForceErrorOk

`func (o *UserCustomPassword) GetForceErrorOk() (*string, bool)`

GetForceErrorOk returns a tuple with the ForceError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceError

`func (o *UserCustomPassword) SetForceError(v string)`

SetForceError sets ForceError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


