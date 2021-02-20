# BankTransferMigrateAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The Plaid &#x60;access_token&#x60; for the newly created Item. | 
**AccountId** | **string** | The Plaid &#x60;account_id&#x60; for the newly created Item. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewBankTransferMigrateAccountResponse

`func NewBankTransferMigrateAccountResponse(accessToken string, accountId string, requestId string, ) *BankTransferMigrateAccountResponse`

NewBankTransferMigrateAccountResponse instantiates a new BankTransferMigrateAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransferMigrateAccountResponseWithDefaults

`func NewBankTransferMigrateAccountResponseWithDefaults() *BankTransferMigrateAccountResponse`

NewBankTransferMigrateAccountResponseWithDefaults instantiates a new BankTransferMigrateAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *BankTransferMigrateAccountResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BankTransferMigrateAccountResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BankTransferMigrateAccountResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *BankTransferMigrateAccountResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BankTransferMigrateAccountResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BankTransferMigrateAccountResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetRequestId

`func (o *BankTransferMigrateAccountResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *BankTransferMigrateAccountResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *BankTransferMigrateAccountResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


