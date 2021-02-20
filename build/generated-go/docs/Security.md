# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityId** | **string** | A unique, Plaid-specific identifier for the security, used to associate securities with holdings. Like all Plaid identifiers, the &#x60;security_id&#x60; is case sensitive. | 
**Isin** | Pointer to **NullableString** | 12-character ISIN, a globally unique securities identifier. | [optional] 
**Cusip** | Pointer to **NullableString** | 9-character CUSIP, an identifier assigned to North American securities. | [optional] 
**Sedol** | Pointer to **NullableString** | 7-character SEDOL, an identifier assigned to securities in the UK. | [optional] 
**InstitutionSecurityId** | Pointer to **NullableString** | An identifier given to the security by the institution | [optional] 
**InstitutionId** | Pointer to **NullableString** | If &#x60;institution_security_id&#x60; is present, this field indicates the Plaid &#x60;institution_id&#x60; of the institution to whom the identifier belongs. | [optional] 
**ProxySecurityId** | Pointer to **NullableString** | In certain cases, Plaid will provide the ID of another security whose performance resembles this security, typically when the original security has low volume, or when a private security can be modeled with a publicly traded security. | [optional] 
**Name** | Pointer to **NullableString** | A descriptive name for the security, suitable for display. | [optional] 
**TickerSymbol** | Pointer to **NullableString** | The security’s trading symbol for publicly traded securities, and otherwise a short identifier if available. | [optional] 
**IsCashEquivalent** | Pointer to **bool** | Indicates that a security is a highly liquid asset and can be treated like cash. | [optional] 
**Type** | **string** | The security type of the holding. Valid security types are:  &#x60;cash&#x60;: Cash, currency, and money market funds  &#x60;derivative&#x60;: Options, warrants, and other derivative instruments  &#x60;equity&#x60;: Domestic and foreign equities  &#x60;etf&#x60;: Multi-asset exchange-traded investment funds  &#x60;fixed income&#x60;: Bonds and certificates of deposit (CDs)  &#x60;loan&#x60;: Loans and loan receivables.  &#x60;mutual fund&#x60;: Open- and closed-end vehicles pooling funds of multiple investors.  &#x60;other&#x60;: Unknown or other investment types | 
**ClosePrice** | Pointer to **NullableFloat32** | Price of the security at the close of the previous trading session. &#x60;null&#x60; for non-public securities. | [optional] 
**ClosePriceAsOf** | Pointer to **NullableString** | Date for which &#x60;close_price&#x60; is accurate. Always &#x60;null&#x60; if &#x60;close_price&#x60; is &#x60;null&#x60;. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the price given. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-&#x60;null&#x60;. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the security. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 

## Methods

### NewSecurity

`func NewSecurity(securityId string, type_ string, ) *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityId

`func (o *Security) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *Security) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *Security) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.


### GetIsin

`func (o *Security) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *Security) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *Security) SetIsin(v string)`

SetIsin sets Isin field to given value.

### HasIsin

`func (o *Security) HasIsin() bool`

HasIsin returns a boolean if a field has been set.

### SetIsinNil

`func (o *Security) SetIsinNil(b bool)`

 SetIsinNil sets the value for Isin to be an explicit nil

### UnsetIsin
`func (o *Security) UnsetIsin()`

UnsetIsin ensures that no value is present for Isin, not even an explicit nil
### GetCusip

`func (o *Security) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *Security) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *Security) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *Security) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### SetCusipNil

`func (o *Security) SetCusipNil(b bool)`

 SetCusipNil sets the value for Cusip to be an explicit nil

### UnsetCusip
`func (o *Security) UnsetCusip()`

UnsetCusip ensures that no value is present for Cusip, not even an explicit nil
### GetSedol

`func (o *Security) GetSedol() string`

GetSedol returns the Sedol field if non-nil, zero value otherwise.

### GetSedolOk

`func (o *Security) GetSedolOk() (*string, bool)`

GetSedolOk returns a tuple with the Sedol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSedol

`func (o *Security) SetSedol(v string)`

SetSedol sets Sedol field to given value.

### HasSedol

`func (o *Security) HasSedol() bool`

HasSedol returns a boolean if a field has been set.

### SetSedolNil

`func (o *Security) SetSedolNil(b bool)`

 SetSedolNil sets the value for Sedol to be an explicit nil

### UnsetSedol
`func (o *Security) UnsetSedol()`

UnsetSedol ensures that no value is present for Sedol, not even an explicit nil
### GetInstitutionSecurityId

`func (o *Security) GetInstitutionSecurityId() string`

GetInstitutionSecurityId returns the InstitutionSecurityId field if non-nil, zero value otherwise.

### GetInstitutionSecurityIdOk

`func (o *Security) GetInstitutionSecurityIdOk() (*string, bool)`

GetInstitutionSecurityIdOk returns a tuple with the InstitutionSecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionSecurityId

`func (o *Security) SetInstitutionSecurityId(v string)`

SetInstitutionSecurityId sets InstitutionSecurityId field to given value.

### HasInstitutionSecurityId

`func (o *Security) HasInstitutionSecurityId() bool`

HasInstitutionSecurityId returns a boolean if a field has been set.

### SetInstitutionSecurityIdNil

`func (o *Security) SetInstitutionSecurityIdNil(b bool)`

 SetInstitutionSecurityIdNil sets the value for InstitutionSecurityId to be an explicit nil

### UnsetInstitutionSecurityId
`func (o *Security) UnsetInstitutionSecurityId()`

UnsetInstitutionSecurityId ensures that no value is present for InstitutionSecurityId, not even an explicit nil
### GetInstitutionId

`func (o *Security) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *Security) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *Security) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *Security) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *Security) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *Security) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetProxySecurityId

`func (o *Security) GetProxySecurityId() string`

GetProxySecurityId returns the ProxySecurityId field if non-nil, zero value otherwise.

### GetProxySecurityIdOk

`func (o *Security) GetProxySecurityIdOk() (*string, bool)`

GetProxySecurityIdOk returns a tuple with the ProxySecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxySecurityId

`func (o *Security) SetProxySecurityId(v string)`

SetProxySecurityId sets ProxySecurityId field to given value.

### HasProxySecurityId

`func (o *Security) HasProxySecurityId() bool`

HasProxySecurityId returns a boolean if a field has been set.

### SetProxySecurityIdNil

`func (o *Security) SetProxySecurityIdNil(b bool)`

 SetProxySecurityIdNil sets the value for ProxySecurityId to be an explicit nil

### UnsetProxySecurityId
`func (o *Security) UnsetProxySecurityId()`

UnsetProxySecurityId ensures that no value is present for ProxySecurityId, not even an explicit nil
### GetName

`func (o *Security) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Security) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Security) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Security) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Security) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Security) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTickerSymbol

`func (o *Security) GetTickerSymbol() string`

GetTickerSymbol returns the TickerSymbol field if non-nil, zero value otherwise.

### GetTickerSymbolOk

`func (o *Security) GetTickerSymbolOk() (*string, bool)`

GetTickerSymbolOk returns a tuple with the TickerSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTickerSymbol

`func (o *Security) SetTickerSymbol(v string)`

SetTickerSymbol sets TickerSymbol field to given value.

### HasTickerSymbol

`func (o *Security) HasTickerSymbol() bool`

HasTickerSymbol returns a boolean if a field has been set.

### SetTickerSymbolNil

`func (o *Security) SetTickerSymbolNil(b bool)`

 SetTickerSymbolNil sets the value for TickerSymbol to be an explicit nil

### UnsetTickerSymbol
`func (o *Security) UnsetTickerSymbol()`

UnsetTickerSymbol ensures that no value is present for TickerSymbol, not even an explicit nil
### GetIsCashEquivalent

`func (o *Security) GetIsCashEquivalent() bool`

GetIsCashEquivalent returns the IsCashEquivalent field if non-nil, zero value otherwise.

### GetIsCashEquivalentOk

`func (o *Security) GetIsCashEquivalentOk() (*bool, bool)`

GetIsCashEquivalentOk returns a tuple with the IsCashEquivalent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCashEquivalent

`func (o *Security) SetIsCashEquivalent(v bool)`

SetIsCashEquivalent sets IsCashEquivalent field to given value.

### HasIsCashEquivalent

`func (o *Security) HasIsCashEquivalent() bool`

HasIsCashEquivalent returns a boolean if a field has been set.

### GetType

`func (o *Security) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Security) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Security) SetType(v string)`

SetType sets Type field to given value.


### GetClosePrice

`func (o *Security) GetClosePrice() float32`

GetClosePrice returns the ClosePrice field if non-nil, zero value otherwise.

### GetClosePriceOk

`func (o *Security) GetClosePriceOk() (*float32, bool)`

GetClosePriceOk returns a tuple with the ClosePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrice

`func (o *Security) SetClosePrice(v float32)`

SetClosePrice sets ClosePrice field to given value.

### HasClosePrice

`func (o *Security) HasClosePrice() bool`

HasClosePrice returns a boolean if a field has been set.

### SetClosePriceNil

`func (o *Security) SetClosePriceNil(b bool)`

 SetClosePriceNil sets the value for ClosePrice to be an explicit nil

### UnsetClosePrice
`func (o *Security) UnsetClosePrice()`

UnsetClosePrice ensures that no value is present for ClosePrice, not even an explicit nil
### GetClosePriceAsOf

`func (o *Security) GetClosePriceAsOf() string`

GetClosePriceAsOf returns the ClosePriceAsOf field if non-nil, zero value otherwise.

### GetClosePriceAsOfOk

`func (o *Security) GetClosePriceAsOfOk() (*string, bool)`

GetClosePriceAsOfOk returns a tuple with the ClosePriceAsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePriceAsOf

`func (o *Security) SetClosePriceAsOf(v string)`

SetClosePriceAsOf sets ClosePriceAsOf field to given value.

### HasClosePriceAsOf

`func (o *Security) HasClosePriceAsOf() bool`

HasClosePriceAsOf returns a boolean if a field has been set.

### SetClosePriceAsOfNil

`func (o *Security) SetClosePriceAsOfNil(b bool)`

 SetClosePriceAsOfNil sets the value for ClosePriceAsOf to be an explicit nil

### UnsetClosePriceAsOf
`func (o *Security) UnsetClosePriceAsOf()`

UnsetClosePriceAsOf ensures that no value is present for ClosePriceAsOf, not even an explicit nil
### GetIsoCurrencyCode

`func (o *Security) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *Security) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *Security) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *Security) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *Security) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *Security) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *Security) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *Security) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *Security) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *Security) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *Security) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *Security) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


