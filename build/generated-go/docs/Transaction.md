# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionType** | Pointer to **string** | Please use the &#x60;payment_channel&#x60; field, &#x60;transaction_type&#x60; will be deprecated in the future.  &#x60;digital:&#x60; transactions that took place online.  &#x60;place:&#x60; transactions that were made at a physical location.  &#x60;special:&#x60; transactions that relate to banks, e.g. fees or deposits.  &#x60;unresolved:&#x60; transactions that do not fit into the other three types.  | [optional] 
**TransactionId** | **string** | The unique ID of the transaction. Like all Plaid identifiers, the &#x60;transaction_id&#x60; is case sensitive. | 
**AccountOwner** | Pointer to **NullableString** | The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts. | [optional] 
**PendingTransactionId** | Pointer to **NullableString** | The ID of a posted transaction&#39;s associated pending transaction, where applicable. | [optional] 
**Pending** | **bool** | When &#x60;true&#x60;, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled. | 
**PaymentChannel** | Pointer to **string** | The channel used to make a payment. &#x60;online:&#x60; transactions that took place online.  &#x60;in store:&#x60; transactions that were made at a physical location.  &#x60;other:&#x60; transactions that relate to banks, e.g. fees or deposits.  This field replaces the &#x60;transaction_type&#x60; field.  | [optional] 
**PaymentMeta** | Pointer to [**PaymentMeta**](PaymentMeta.md) |  | [optional] 
**Name** | Pointer to **string** | The merchant name or transaction description.  If the &#x60;transaction&#x60; object was returned by a Transactions endpoint such as &#x60;/transactions/get&#x60;, this field will always appear. If the &#x60;transaction&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**MerchantName** | Pointer to **NullableString** | The merchant name, as extracted by Plaid from the &#x60;name&#x60; field. | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**AuthorizedDate** | Pointer to **NullableString** | The date that the transaction was authorized. Dates are returned in an ISO 8601 format ( &#x60;YYYY-MM-DD&#x60; ). | [optional] 
**Date** | **string** | For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an ISO 8601 format ( &#x60;YYYY-MM-DD&#x60; ). | 
**CategoryId** | Pointer to **string** | The ID of the category to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).  If the &#x60;transaction&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**Category** | Pointer to **[]string** | A hierarchical array of the categories to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).  If the &#x60;transaction&#x60; object was returned by an Assets endpoint such as &#x60;/asset_report/get/&#x60; or &#x60;/asset_report/pdf/get&#x60;, this field will only appear in an Asset Report with Insights. | [optional] 
**UnofficialCurrencyCode** | Pointer to **NullableString** | The unofficial currency code associated with the transaction. Always &#x60;null&#x60; if &#x60;iso_currency_code&#x60; is non-&#x60;null&#x60;. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](/docs/api/accounts#currency-code-schema) for a full listing of supported &#x60;iso_currency_code&#x60;s. | [optional] 
**IsoCurrencyCode** | Pointer to **NullableString** | The ISO-4217 currency code of the transaction. Always &#x60;null&#x60; if &#x60;unofficial_currency_code&#x60; is non-null. | [optional] 
**Amount** | **float32** | The settled value of the transaction, denominated in the account&#39;s currency, as stated in &#x60;iso_currency_code&#x60; or &#x60;unofficial_currency_code&#x60;. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative. | 
**AccountId** | **string** | The ID of the account in which this transaction occurred. | 
**TransactionCode** | Pointer to [**NullableTransactionCode**](TransactionCode.md) |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction(transactionId string, pending bool, date string, amount float32, accountId string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionType

`func (o *Transaction) GetTransactionType() string`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transaction) GetTransactionTypeOk() (*string, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transaction) SetTransactionType(v string)`

SetTransactionType sets TransactionType field to given value.

### HasTransactionType

`func (o *Transaction) HasTransactionType() bool`

HasTransactionType returns a boolean if a field has been set.

### GetTransactionId

`func (o *Transaction) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transaction) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transaction) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAccountOwner

`func (o *Transaction) GetAccountOwner() string`

GetAccountOwner returns the AccountOwner field if non-nil, zero value otherwise.

### GetAccountOwnerOk

`func (o *Transaction) GetAccountOwnerOk() (*string, bool)`

GetAccountOwnerOk returns a tuple with the AccountOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountOwner

`func (o *Transaction) SetAccountOwner(v string)`

SetAccountOwner sets AccountOwner field to given value.

### HasAccountOwner

`func (o *Transaction) HasAccountOwner() bool`

HasAccountOwner returns a boolean if a field has been set.

### SetAccountOwnerNil

`func (o *Transaction) SetAccountOwnerNil(b bool)`

 SetAccountOwnerNil sets the value for AccountOwner to be an explicit nil

### UnsetAccountOwner
`func (o *Transaction) UnsetAccountOwner()`

UnsetAccountOwner ensures that no value is present for AccountOwner, not even an explicit nil
### GetPendingTransactionId

`func (o *Transaction) GetPendingTransactionId() string`

GetPendingTransactionId returns the PendingTransactionId field if non-nil, zero value otherwise.

### GetPendingTransactionIdOk

`func (o *Transaction) GetPendingTransactionIdOk() (*string, bool)`

GetPendingTransactionIdOk returns a tuple with the PendingTransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingTransactionId

`func (o *Transaction) SetPendingTransactionId(v string)`

SetPendingTransactionId sets PendingTransactionId field to given value.

### HasPendingTransactionId

`func (o *Transaction) HasPendingTransactionId() bool`

HasPendingTransactionId returns a boolean if a field has been set.

### SetPendingTransactionIdNil

`func (o *Transaction) SetPendingTransactionIdNil(b bool)`

 SetPendingTransactionIdNil sets the value for PendingTransactionId to be an explicit nil

### UnsetPendingTransactionId
`func (o *Transaction) UnsetPendingTransactionId()`

UnsetPendingTransactionId ensures that no value is present for PendingTransactionId, not even an explicit nil
### GetPending

`func (o *Transaction) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *Transaction) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *Transaction) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetPaymentChannel

`func (o *Transaction) GetPaymentChannel() string`

GetPaymentChannel returns the PaymentChannel field if non-nil, zero value otherwise.

### GetPaymentChannelOk

`func (o *Transaction) GetPaymentChannelOk() (*string, bool)`

GetPaymentChannelOk returns a tuple with the PaymentChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentChannel

`func (o *Transaction) SetPaymentChannel(v string)`

SetPaymentChannel sets PaymentChannel field to given value.

### HasPaymentChannel

`func (o *Transaction) HasPaymentChannel() bool`

HasPaymentChannel returns a boolean if a field has been set.

### GetPaymentMeta

`func (o *Transaction) GetPaymentMeta() PaymentMeta`

GetPaymentMeta returns the PaymentMeta field if non-nil, zero value otherwise.

### GetPaymentMetaOk

`func (o *Transaction) GetPaymentMetaOk() (*PaymentMeta, bool)`

GetPaymentMetaOk returns a tuple with the PaymentMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMeta

`func (o *Transaction) SetPaymentMeta(v PaymentMeta)`

SetPaymentMeta sets PaymentMeta field to given value.

### HasPaymentMeta

`func (o *Transaction) HasPaymentMeta() bool`

HasPaymentMeta returns a boolean if a field has been set.

### GetName

`func (o *Transaction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Transaction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Transaction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Transaction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMerchantName

`func (o *Transaction) GetMerchantName() string`

GetMerchantName returns the MerchantName field if non-nil, zero value otherwise.

### GetMerchantNameOk

`func (o *Transaction) GetMerchantNameOk() (*string, bool)`

GetMerchantNameOk returns a tuple with the MerchantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantName

`func (o *Transaction) SetMerchantName(v string)`

SetMerchantName sets MerchantName field to given value.

### HasMerchantName

`func (o *Transaction) HasMerchantName() bool`

HasMerchantName returns a boolean if a field has been set.

### SetMerchantNameNil

`func (o *Transaction) SetMerchantNameNil(b bool)`

 SetMerchantNameNil sets the value for MerchantName to be an explicit nil

### UnsetMerchantName
`func (o *Transaction) UnsetMerchantName()`

UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
### GetLocation

`func (o *Transaction) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Transaction) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Transaction) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Transaction) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAuthorizedDate

`func (o *Transaction) GetAuthorizedDate() string`

GetAuthorizedDate returns the AuthorizedDate field if non-nil, zero value otherwise.

### GetAuthorizedDateOk

`func (o *Transaction) GetAuthorizedDateOk() (*string, bool)`

GetAuthorizedDateOk returns a tuple with the AuthorizedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedDate

`func (o *Transaction) SetAuthorizedDate(v string)`

SetAuthorizedDate sets AuthorizedDate field to given value.

### HasAuthorizedDate

`func (o *Transaction) HasAuthorizedDate() bool`

HasAuthorizedDate returns a boolean if a field has been set.

### SetAuthorizedDateNil

`func (o *Transaction) SetAuthorizedDateNil(b bool)`

 SetAuthorizedDateNil sets the value for AuthorizedDate to be an explicit nil

### UnsetAuthorizedDate
`func (o *Transaction) UnsetAuthorizedDate()`

UnsetAuthorizedDate ensures that no value is present for AuthorizedDate, not even an explicit nil
### GetDate

`func (o *Transaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Transaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Transaction) SetDate(v string)`

SetDate sets Date field to given value.


### GetCategoryId

`func (o *Transaction) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *Transaction) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *Transaction) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *Transaction) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategory

`func (o *Transaction) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Transaction) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *Transaction) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Transaction) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetUnofficialCurrencyCode

`func (o *Transaction) GetUnofficialCurrencyCode() string`

GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field if non-nil, zero value otherwise.

### GetUnofficialCurrencyCodeOk

`func (o *Transaction) GetUnofficialCurrencyCodeOk() (*string, bool)`

GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnofficialCurrencyCode

`func (o *Transaction) SetUnofficialCurrencyCode(v string)`

SetUnofficialCurrencyCode sets UnofficialCurrencyCode field to given value.

### HasUnofficialCurrencyCode

`func (o *Transaction) HasUnofficialCurrencyCode() bool`

HasUnofficialCurrencyCode returns a boolean if a field has been set.

### SetUnofficialCurrencyCodeNil

`func (o *Transaction) SetUnofficialCurrencyCodeNil(b bool)`

 SetUnofficialCurrencyCodeNil sets the value for UnofficialCurrencyCode to be an explicit nil

### UnsetUnofficialCurrencyCode
`func (o *Transaction) UnsetUnofficialCurrencyCode()`

UnsetUnofficialCurrencyCode ensures that no value is present for UnofficialCurrencyCode, not even an explicit nil
### GetIsoCurrencyCode

`func (o *Transaction) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *Transaction) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *Transaction) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *Transaction) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.

### SetIsoCurrencyCodeNil

`func (o *Transaction) SetIsoCurrencyCodeNil(b bool)`

 SetIsoCurrencyCodeNil sets the value for IsoCurrencyCode to be an explicit nil

### UnsetIsoCurrencyCode
`func (o *Transaction) UnsetIsoCurrencyCode()`

UnsetIsoCurrencyCode ensures that no value is present for IsoCurrencyCode, not even an explicit nil
### GetAmount

`func (o *Transaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetAccountId

`func (o *Transaction) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTransactionCode

`func (o *Transaction) GetTransactionCode() TransactionCode`

GetTransactionCode returns the TransactionCode field if non-nil, zero value otherwise.

### GetTransactionCodeOk

`func (o *Transaction) GetTransactionCodeOk() (*TransactionCode, bool)`

GetTransactionCodeOk returns a tuple with the TransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionCode

`func (o *Transaction) SetTransactionCode(v TransactionCode)`

SetTransactionCode sets TransactionCode field to given value.

### HasTransactionCode

`func (o *Transaction) HasTransactionCode() bool`

HasTransactionCode returns a boolean if a field has been set.

### SetTransactionCodeNil

`func (o *Transaction) SetTransactionCodeNil(b bool)`

 SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil

### UnsetTransactionCode
`func (o *Transaction) UnsetTransactionCode()`

UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


