# LinkTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**ClientName** | **string** | The name of your application, as it should be displayed in Link. | 
**Language** | **string** | The language that Link should be displayed in.  Supported languages are: - English (&#x60;&#39;en&#39;&#x60;) - French (&#x60;&#39;fr&#39;&#x60;) - Spanish (&#x60;&#39;es&#39;&#x60;) - Dutch (&#x60;&#39;nl&#39;&#x60;)  When using a Link customization, the language configured here must match the setting in the customization, or the customization will not be applied. | 
**CountryCodes** | [**[]CountryCode**](CountryCode.md) | Specify an array of Plaid-supported country codes using the ISO-3166-1 alpha-2 country code standard. Institutions from all listed countries will be shown.  Supported country codes are: &#x60;US&#x60;, &#x60;CA&#x60;, &#x60;ES&#x60;, &#x60;FR&#x60;, &#x60;GB&#x60;, &#x60;IE&#x60;, &#x60;NL&#x60;. Example value: &#x60;[&#39;US&#39;, &#39;CA&#39;]&#x60;.  If Link is launched with multiple country codes, only products that you are enabled for in all countries will be used by Link. Note that while all countries are enabled by default in Sandbox and Development, in Production only US and Canada are enabled by default. To gain access to European institutions in the Production environment, [file a product access Support ticket](https://dashboard.plaid.com/support/new/product-and-development/product-troubleshooting/request-product-access) via the Plaid dashboard. If you initialize with a European country code, your users will see the European consent panel during the Link flow.  If using a Link customization, make sure the country codes in the customization match those specified in &#x60;country_codes&#x60;. If both &#x60;country_codes&#x60; and a Link customization are used, the value in &#x60;country_codes&#x60; may override the value in the customization.  If using the Auth features Instant Match, Same-day Micro-deposits, or Automated Micro-deposits, &#x60;country_codes&#x60; must be set to &#x60;[&#39;US&#39;]&#x60;. | 
**User** | [**LinkTokenCreateRequestUser**](LinkTokenCreateRequestUser.md) |  | 
**Products** | Pointer to [**[]Products**](Products.md) | List of Plaid product(s) you wish to use. If launching Link in update mode, should be omitted; required otherwise. Valid products are:  &#x60;transactions&#x60;, &#x60;auth&#x60;, &#x60;identity&#x60;, &#x60;assets&#x60;, &#x60;investments&#x60;, &#x60;liabilities&#x60;, &#x60;payment_initiation&#x60;, &#x60;deposit_switch&#x60;  Example: &#x60;[&#39;auth&#39;, &#39;transactions&#39;]&#x60;  &#x60;balance&#x60; is *not* a valid value, the Balance product does not require explicit initalization and will automatically be initialized when any other product is initialized.  Only institutions that support *all* requested products will be shown in Link; to maximize the number of institutions listed, it is recommended to initialize Link with the minimal product set required for your use case. Additional products can be added after Link initialization by calling the relevant endpoints. For details and exceptions, see [Choosing when to initialize products](/docs/link/best-practices/#choosing-when-to-initialize-products).          In Production, you will be billed for each product that you specify when initializing Link. Note that a product cannot be removed from an Item once the Item has been initialized with that product. To stop billing on an Item for subscription-based products, such as Liabilities, Investments, and Transactions, remove the Item via &#x60;/item/remove&#x60;. | [optional] 
**Webhook** | Pointer to **string** | The destination URL to which any webhooks should be sent. | [optional] 
**AccessToken** | Pointer to **string** | The &#x60;access_token&#x60; associated with the Item to update, used when updating or modifying an existing &#x60;access_token&#x60;. Used when launching Link in update mode, when completing the Same-day (manual) Micro-deposit flow, or (optionally) when initializing Link as part of the Payment Initiation (UK and Europe) flow. | [optional] 
**LinkCustomizationName** | Pointer to **string** | The name of the Link customization from the Plaid Dashboard to be applied to Link. If not specified, the &#x60;default&#x60; customization will be used. When using a Link customization, the language in the customization must match the language selected via the &#x60;language&#x60; parameter, and the countries in the customization should match the country codes selected via &#x60;country_codes&#x60;. | [optional] 
**RedirectUri** | Pointer to **string** | A URI indicating the destination where a user should be forwarded after completing the Link flow; used to support OAuth authentication flows when launching Link in the browser or via a webview. The &#x60;redirect_uri&#x60; should not contain any query parameters. If &#x60;android_package_name&#x60; is specified, this field should be left blank. Any redirect URI specified here must also be added under the \&quot;Allowed redirect URIs\&quot; configuration on the [developer dashboard](https://dashboard.plaid.com/team/api). In non-Sandbox (Production and Development) environments, the &#x60;redirect_uri&#x60; must begin with https. | [optional] 
**AndroidPackageName** | Pointer to **string** | The name of your app&#39;s Android package. Required if using the &#x60;link_token&#x60; to initialize Link on Android. When creating a &#x60;link_token&#x60; for initializing Link on other platforms, this field must be left blank. Any package name specified here must also be added to the Allowed Android package names setting on the [developer dashboard](https://dashboard.plaid.com/team/api).  | [optional] 
**AccountFilters** | Pointer to [**LinkTokenAccountFilters**](LinkTokenAccountFilters.md) |  | [optional] 
**InstitutionId** | Pointer to **string** | Used for certain legacy use cases | [optional] 
**PaymentInitiation** | Pointer to [**LinkTokenCreateRequestPaymentInitiation**](LinkTokenCreateRequestPaymentInitiation.md) |  | [optional] 
**DepositSwitch** | Pointer to [**LinkTokenCreateRequestDepositSwitch**](LinkTokenCreateRequestDepositSwitch.md) |  | [optional] 

## Methods

### NewLinkTokenCreateRequest

`func NewLinkTokenCreateRequest(clientName string, language string, countryCodes []CountryCode, user LinkTokenCreateRequestUser, ) *LinkTokenCreateRequest`

NewLinkTokenCreateRequest instantiates a new LinkTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestWithDefaults

`func NewLinkTokenCreateRequestWithDefaults() *LinkTokenCreateRequest`

NewLinkTokenCreateRequestWithDefaults instantiates a new LinkTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *LinkTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *LinkTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *LinkTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *LinkTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *LinkTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *LinkTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *LinkTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *LinkTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetClientName

`func (o *LinkTokenCreateRequest) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *LinkTokenCreateRequest) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *LinkTokenCreateRequest) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### GetLanguage

`func (o *LinkTokenCreateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *LinkTokenCreateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *LinkTokenCreateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCountryCodes

`func (o *LinkTokenCreateRequest) GetCountryCodes() []CountryCode`

GetCountryCodes returns the CountryCodes field if non-nil, zero value otherwise.

### GetCountryCodesOk

`func (o *LinkTokenCreateRequest) GetCountryCodesOk() (*[]CountryCode, bool)`

GetCountryCodesOk returns a tuple with the CountryCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodes

`func (o *LinkTokenCreateRequest) SetCountryCodes(v []CountryCode)`

SetCountryCodes sets CountryCodes field to given value.


### GetUser

`func (o *LinkTokenCreateRequest) GetUser() LinkTokenCreateRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LinkTokenCreateRequest) GetUserOk() (*LinkTokenCreateRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LinkTokenCreateRequest) SetUser(v LinkTokenCreateRequestUser)`

SetUser sets User field to given value.


### GetProducts

`func (o *LinkTokenCreateRequest) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *LinkTokenCreateRequest) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *LinkTokenCreateRequest) SetProducts(v []Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *LinkTokenCreateRequest) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetWebhook

`func (o *LinkTokenCreateRequest) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *LinkTokenCreateRequest) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *LinkTokenCreateRequest) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *LinkTokenCreateRequest) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### GetAccessToken

`func (o *LinkTokenCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *LinkTokenCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *LinkTokenCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *LinkTokenCreateRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetLinkCustomizationName

`func (o *LinkTokenCreateRequest) GetLinkCustomizationName() string`

GetLinkCustomizationName returns the LinkCustomizationName field if non-nil, zero value otherwise.

### GetLinkCustomizationNameOk

`func (o *LinkTokenCreateRequest) GetLinkCustomizationNameOk() (*string, bool)`

GetLinkCustomizationNameOk returns a tuple with the LinkCustomizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCustomizationName

`func (o *LinkTokenCreateRequest) SetLinkCustomizationName(v string)`

SetLinkCustomizationName sets LinkCustomizationName field to given value.

### HasLinkCustomizationName

`func (o *LinkTokenCreateRequest) HasLinkCustomizationName() bool`

HasLinkCustomizationName returns a boolean if a field has been set.

### GetRedirectUri

`func (o *LinkTokenCreateRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *LinkTokenCreateRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *LinkTokenCreateRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *LinkTokenCreateRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetAndroidPackageName

`func (o *LinkTokenCreateRequest) GetAndroidPackageName() string`

GetAndroidPackageName returns the AndroidPackageName field if non-nil, zero value otherwise.

### GetAndroidPackageNameOk

`func (o *LinkTokenCreateRequest) GetAndroidPackageNameOk() (*string, bool)`

GetAndroidPackageNameOk returns a tuple with the AndroidPackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAndroidPackageName

`func (o *LinkTokenCreateRequest) SetAndroidPackageName(v string)`

SetAndroidPackageName sets AndroidPackageName field to given value.

### HasAndroidPackageName

`func (o *LinkTokenCreateRequest) HasAndroidPackageName() bool`

HasAndroidPackageName returns a boolean if a field has been set.

### GetAccountFilters

`func (o *LinkTokenCreateRequest) GetAccountFilters() LinkTokenAccountFilters`

GetAccountFilters returns the AccountFilters field if non-nil, zero value otherwise.

### GetAccountFiltersOk

`func (o *LinkTokenCreateRequest) GetAccountFiltersOk() (*LinkTokenAccountFilters, bool)`

GetAccountFiltersOk returns a tuple with the AccountFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountFilters

`func (o *LinkTokenCreateRequest) SetAccountFilters(v LinkTokenAccountFilters)`

SetAccountFilters sets AccountFilters field to given value.

### HasAccountFilters

`func (o *LinkTokenCreateRequest) HasAccountFilters() bool`

HasAccountFilters returns a boolean if a field has been set.

### GetInstitutionId

`func (o *LinkTokenCreateRequest) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *LinkTokenCreateRequest) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *LinkTokenCreateRequest) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *LinkTokenCreateRequest) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### GetPaymentInitiation

`func (o *LinkTokenCreateRequest) GetPaymentInitiation() LinkTokenCreateRequestPaymentInitiation`

GetPaymentInitiation returns the PaymentInitiation field if non-nil, zero value otherwise.

### GetPaymentInitiationOk

`func (o *LinkTokenCreateRequest) GetPaymentInitiationOk() (*LinkTokenCreateRequestPaymentInitiation, bool)`

GetPaymentInitiationOk returns a tuple with the PaymentInitiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInitiation

`func (o *LinkTokenCreateRequest) SetPaymentInitiation(v LinkTokenCreateRequestPaymentInitiation)`

SetPaymentInitiation sets PaymentInitiation field to given value.

### HasPaymentInitiation

`func (o *LinkTokenCreateRequest) HasPaymentInitiation() bool`

HasPaymentInitiation returns a boolean if a field has been set.

### GetDepositSwitch

`func (o *LinkTokenCreateRequest) GetDepositSwitch() LinkTokenCreateRequestDepositSwitch`

GetDepositSwitch returns the DepositSwitch field if non-nil, zero value otherwise.

### GetDepositSwitchOk

`func (o *LinkTokenCreateRequest) GetDepositSwitchOk() (*LinkTokenCreateRequestDepositSwitch, bool)`

GetDepositSwitchOk returns a tuple with the DepositSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitch

`func (o *LinkTokenCreateRequest) SetDepositSwitch(v LinkTokenCreateRequestDepositSwitch)`

SetDepositSwitch sets DepositSwitch field to given value.

### HasDepositSwitch

`func (o *LinkTokenCreateRequest) HasDepositSwitch() bool`

HasDepositSwitch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


