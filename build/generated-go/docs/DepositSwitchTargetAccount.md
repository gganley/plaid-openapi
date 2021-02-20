# DepositSwitchTargetAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | **string** | Account number for deposit switch destination | 
**RoutingNumber** | **string** | Routing number for deposit switch destination | 
**AccountName** | **string** | The name of the deposit switch destination account, as it will be displayed to the end user in the Deposit Switch interface. It is not required to match the name used in online banking. | 
**AccountSubtype** | **string** | The account subtype of the account, either &#x60;checking&#x60; or &#x60;savings&#x60;. | 

## Methods

### NewDepositSwitchTargetAccount

`func NewDepositSwitchTargetAccount(accountNumber string, routingNumber string, accountName string, accountSubtype string, ) *DepositSwitchTargetAccount`

NewDepositSwitchTargetAccount instantiates a new DepositSwitchTargetAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchTargetAccountWithDefaults

`func NewDepositSwitchTargetAccountWithDefaults() *DepositSwitchTargetAccount`

NewDepositSwitchTargetAccountWithDefaults instantiates a new DepositSwitchTargetAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *DepositSwitchTargetAccount) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *DepositSwitchTargetAccount) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *DepositSwitchTargetAccount) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetRoutingNumber

`func (o *DepositSwitchTargetAccount) GetRoutingNumber() string`

GetRoutingNumber returns the RoutingNumber field if non-nil, zero value otherwise.

### GetRoutingNumberOk

`func (o *DepositSwitchTargetAccount) GetRoutingNumberOk() (*string, bool)`

GetRoutingNumberOk returns a tuple with the RoutingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumber

`func (o *DepositSwitchTargetAccount) SetRoutingNumber(v string)`

SetRoutingNumber sets RoutingNumber field to given value.


### GetAccountName

`func (o *DepositSwitchTargetAccount) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *DepositSwitchTargetAccount) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *DepositSwitchTargetAccount) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetAccountSubtype

`func (o *DepositSwitchTargetAccount) GetAccountSubtype() string`

GetAccountSubtype returns the AccountSubtype field if non-nil, zero value otherwise.

### GetAccountSubtypeOk

`func (o *DepositSwitchTargetAccount) GetAccountSubtypeOk() (*string, bool)`

GetAccountSubtypeOk returns a tuple with the AccountSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSubtype

`func (o *DepositSwitchTargetAccount) SetAccountSubtype(v string)`

SetAccountSubtype sets AccountSubtype field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


