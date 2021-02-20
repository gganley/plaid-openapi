# LinkTokenGetMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialProducts** | Pointer to [**[]Products**](Products.md) | The &#x60;products&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 
**Webhook** | Pointer to **NullableString** | The &#x60;webhook&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 
**CountryCodes** | Pointer to [**[]CountryCode**](CountryCode.md) | The &#x60;country_codes&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 
**Language** | Pointer to **NullableString** | The &#x60;language&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 
**AccountFilters** | Pointer to [**AccountFiltersResponse**](AccountFiltersResponse.md) |  | [optional] 
**RedirectUri** | Pointer to **NullableString** | The &#x60;redirect_uri&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 
**ClientName** | Pointer to **NullableString** | The &#x60;client_name&#x60; specified in the &#x60;/link/token/create&#x60; call. | [optional] 

## Methods

### NewLinkTokenGetMetadataResponse

`func NewLinkTokenGetMetadataResponse() *LinkTokenGetMetadataResponse`

NewLinkTokenGetMetadataResponse instantiates a new LinkTokenGetMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenGetMetadataResponseWithDefaults

`func NewLinkTokenGetMetadataResponseWithDefaults() *LinkTokenGetMetadataResponse`

NewLinkTokenGetMetadataResponseWithDefaults instantiates a new LinkTokenGetMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialProducts

`func (o *LinkTokenGetMetadataResponse) GetInitialProducts() []Products`

GetInitialProducts returns the InitialProducts field if non-nil, zero value otherwise.

### GetInitialProductsOk

`func (o *LinkTokenGetMetadataResponse) GetInitialProductsOk() (*[]Products, bool)`

GetInitialProductsOk returns a tuple with the InitialProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialProducts

`func (o *LinkTokenGetMetadataResponse) SetInitialProducts(v []Products)`

SetInitialProducts sets InitialProducts field to given value.

### HasInitialProducts

`func (o *LinkTokenGetMetadataResponse) HasInitialProducts() bool`

HasInitialProducts returns a boolean if a field has been set.

### GetWebhook

`func (o *LinkTokenGetMetadataResponse) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *LinkTokenGetMetadataResponse) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *LinkTokenGetMetadataResponse) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *LinkTokenGetMetadataResponse) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *LinkTokenGetMetadataResponse) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *LinkTokenGetMetadataResponse) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetCountryCodes

`func (o *LinkTokenGetMetadataResponse) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *LinkTokenGetMetadataResponse) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *LinkTokenGetMetadataResponse) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.

### HasCountryCodes

`func (o *LinkTokenGetMetadataResponse) HasCountryCodes() bool`

HasCountryCodes returns a boolean if a field has been set.

### GetLanguage

`func (o *LinkTokenGetMetadataResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LinkTokenGetMetadataResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LinkTokenGetMetadataResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *LinkTokenGetMetadataResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *LinkTokenGetMetadataResponse) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *LinkTokenGetMetadataResponse) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetAccountFilters

`func (o *LinkTokenGetMetadataResponse) GetAccountFilters() AccountFiltersResponse`

GetAccountFilters returns the AccountFilters field if non-nil, zero value otherwise.

### GetAccountFiltersOk

`func (o *LinkTokenGetMetadataResponse) GetAccountFiltersOk() (*AccountFiltersResponse, bool)`

GetAccountFiltersOk returns a tuple with the AccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilters

`func (o *LinkTokenGetMetadataResponse) SetAccountFilters(v AccountFiltersResponse)`

SetAccountFilters sets AccountFilters field to given value.

### HasAccountFilters

`func (o *LinkTokenGetMetadataResponse) HasAccountFilters() bool`

HasAccountFilters returns a boolean if a field has been set.

### GetRedirectUri

`func (o *LinkTokenGetMetadataResponse) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *LinkTokenGetMetadataResponse) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *LinkTokenGetMetadataResponse) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *LinkTokenGetMetadataResponse) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *LinkTokenGetMetadataResponse) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *LinkTokenGetMetadataResponse) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil
### GetClientName

`func (o *LinkTokenGetMetadataResponse) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *LinkTokenGetMetadataResponse) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *LinkTokenGetMetadataResponse) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *LinkTokenGetMetadataResponse) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### SetClientNameNil

`func (o *LinkTokenGetMetadataResponse) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *LinkTokenGetMetadataResponse) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


