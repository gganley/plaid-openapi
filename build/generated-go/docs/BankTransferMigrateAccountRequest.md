# BankTransferMigrateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccountNumber** | **string** | The user&#39;s account number. | 
**RoutingNumber** | **string** | The user&#39;s routing number. | 
**AccountType** | **string** | The type of the bank account (&#x60;checking&#x60; or &#x60;savings&#x60;). | 

## Methods

### NewBankTransferMigrateAccountRequest

`func NewBankTransferMigrateAccountRequest(accountNumber string, routingNumber string, accountType string, ) *BankTransferMigrateAccountRequest`

NewBankTransferMigrateAccountRequest instantiates a new BankTransferMigrateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferMigrateAccountRequestWithDefaults

`func NewBankTransferMigrateAccountRequestWithDefaults() *BankTransferMigrateAccountRequest`

NewBankTransferMigrateAccountRequestWithDefaults instantiates a new BankTransferMigrateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *BankTransferMigrateAccountRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransferMigrateAccountRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransferMigrateAccountRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransferMigrateAccountRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *BankTransferMigrateAccountRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BankTransferMigrateAccountRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BankTransferMigrateAccountRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *BankTransferMigrateAccountRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccountNumber

`func (o *BankTransferMigrateAccountRequest) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *BankTransferMigrateAccountRequest) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *BankTransferMigrateAccountRequest) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetRoutingNumber

`func (o *BankTransferMigrateAccountRequest) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *BankTransferMigrateAccountRequest) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *BankTransferMigrateAccountRequest) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetAccountType

`func (o *BankTransferMigrateAccountRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankTransferMigrateAccountRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankTransferMigrateAccountRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


