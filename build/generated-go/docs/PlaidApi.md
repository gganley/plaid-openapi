# \PlaidApi

All URIs are relative to *https://production.plaid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsBalanceGet**](PlaidApi.md#AccountsBalanceGet) | **Post** /accounts/balance/get | Retrieve real-time balance data
[**AccountsGet**](PlaidApi.md#AccountsGet) | **Post** /accounts/get | Retrieve accounts
[**AssetReportAuditCopyCreate**](PlaidApi.md#AssetReportAuditCopyCreate) | **Post** /asset_report/audit_copy/create | Create Asset Report Audit Copy
[**AssetReportAuditCopyGet**](PlaidApi.md#AssetReportAuditCopyGet) | **Post** /asset_report/audit_copy/get | Retrieve an Asset Report Audit Copy
[**AssetReportAuditCopyRemove**](PlaidApi.md#AssetReportAuditCopyRemove) | **Post** /asset_report/audit_copy/remove | Remove Asset Report Audit Copy
[**AssetReportCreate**](PlaidApi.md#AssetReportCreate) | **Post** /asset_report/create | Create an Asset Report
[**AssetReportFilter**](PlaidApi.md#AssetReportFilter) | **Post** /asset_report/filter | Filter Asset Report
[**AssetReportGet**](PlaidApi.md#AssetReportGet) | **Post** /asset_report/get | Retrieve an Asset Report
[**AssetReportPdfGet**](PlaidApi.md#AssetReportPdfGet) | **Post** /asset_report/pdf/get | Retrieve a PDF Asset Report
[**AssetReportRefresh**](PlaidApi.md#AssetReportRefresh) | **Post** /asset_report/refresh | Refresh an Asset Report
[**AssetReportRemove**](PlaidApi.md#AssetReportRemove) | **Post** /asset_report/remove | Delete an Asset Report
[**AuthGet**](PlaidApi.md#AuthGet) | **Post** /auth/get | Retrieve auth data
[**BankTransferBalanceGet**](PlaidApi.md#BankTransferBalanceGet) | **Post** /bank_transfer/balance/get | Get balance of your Bank Transfer account
[**BankTransferCancel**](PlaidApi.md#BankTransferCancel) | **Post** /bank_transfer/cancel | Cancel a bank transfer
[**BankTransferCreate**](PlaidApi.md#BankTransferCreate) | **Post** /bank_transfer/create | Create a bank transfer
[**BankTransferEventList**](PlaidApi.md#BankTransferEventList) | **Post** /bank_transfer/event/list | List bank transfer events
[**BankTransferEventSync**](PlaidApi.md#BankTransferEventSync) | **Post** /bank_transfer/event/sync | Sync bank transfer events
[**BankTransferGet**](PlaidApi.md#BankTransferGet) | **Post** /bank_transfer/get | Retrieve a bank transfer
[**BankTransferList**](PlaidApi.md#BankTransferList) | **Post** /bank_transfer/list | List bank transfers
[**BankTransferMigrateAccount**](PlaidApi.md#BankTransferMigrateAccount) | **Post** /bank_transfer/migrate_account | Migrate account into Bank Transfers
[**CategoriesGet**](PlaidApi.md#CategoriesGet) | **Post** /categories/get | Get Categories
[**CreatePaymentToken**](PlaidApi.md#CreatePaymentToken) | **Post** /payment_initiation/payment/token/create | Create payment token
[**DepositSwitchAltCreate**](PlaidApi.md#DepositSwitchAltCreate) | **Post** /deposit_switch/alt/create | Create a deposit switch when not using Plaid Exchange.&#39;
[**DepositSwitchCreate**](PlaidApi.md#DepositSwitchCreate) | **Post** /deposit_switch/create | Create a deposit switch
[**DepositSwitchGet**](PlaidApi.md#DepositSwitchGet) | **Post** /deposit_switch/get | Retrieve a deposit switch
[**DepositSwitchTokenCreate**](PlaidApi.md#DepositSwitchTokenCreate) | **Post** /deposit_switch/token/create | Create a deposit switch token
[**EmployersSearch**](PlaidApi.md#EmployersSearch) | **Post** /employers/search | Search employer database
[**IdentityGet**](PlaidApi.md#IdentityGet) | **Post** /identity/get | Retrieve identity data
[**IncomeVerificationDocumentsDownload**](PlaidApi.md#IncomeVerificationDocumentsDownload) | **Post** /income/verification/documents/download | Download the original documents used for income verification
[**IncomeVerificationPaystubGet**](PlaidApi.md#IncomeVerificationPaystubGet) | **Post** /income/verification/paystub/get | Retrieve information from the paystub used for income verification
[**IncomeVerificationSummaryGet**](PlaidApi.md#IncomeVerificationSummaryGet) | **Post** /income/verification/summary/get | Retrieve a summary of information derived from income verification
[**InstitutionsGet**](PlaidApi.md#InstitutionsGet) | **Post** /institutions/get | Get details of all supported institutions
[**InstitutionsGetById**](PlaidApi.md#InstitutionsGetById) | **Post** /institutions/get_by_id | Get details of an institution
[**InstitutionsSearch**](PlaidApi.md#InstitutionsSearch) | **Post** /institutions/search | Search institutions
[**InvestmentsHoldingsGet**](PlaidApi.md#InvestmentsHoldingsGet) | **Post** /investments/holdings/get | Get Investment holdings
[**InvestmentsTransactionsGet**](PlaidApi.md#InvestmentsTransactionsGet) | **Post** /investments/transactions/get | Get investment transactions
[**ItemAccessTokenInvalidate**](PlaidApi.md#ItemAccessTokenInvalidate) | **Post** /item/access_token/invalidate | Invalidate access_token
[**ItemCreatePublicToken**](PlaidApi.md#ItemCreatePublicToken) | **Post** /item/public_token/create | Create public token
[**ItemGet**](PlaidApi.md#ItemGet) | **Post** /item/get | Retrieve an Item
[**ItemImport**](PlaidApi.md#ItemImport) | **Post** /item/import | Import Item
[**ItemPublicTokenExchange**](PlaidApi.md#ItemPublicTokenExchange) | **Post** /item/public_token/exchange | Exchange public token for an access token
[**ItemRemove**](PlaidApi.md#ItemRemove) | **Post** /item/remove | Remove an Item
[**ItemWebhookUpdate**](PlaidApi.md#ItemWebhookUpdate) | **Post** /item/webhook/update | Update Webhook URL
[**LiabilitiesGet**](PlaidApi.md#LiabilitiesGet) | **Post** /liabilities/get | Retrieve Liabilities data
[**LinkTokenCreate**](PlaidApi.md#LinkTokenCreate) | **Post** /link/token/create | Create Link Token
[**LinkTokenGet**](PlaidApi.md#LinkTokenGet) | **Post** /link/token/get | Get Link Token
[**PaymentInitiationPaymentCreate**](PlaidApi.md#PaymentInitiationPaymentCreate) | **Post** /payment_initiation/payment/create | Create a payment
[**PaymentInitiationPaymentGet**](PlaidApi.md#PaymentInitiationPaymentGet) | **Post** /payment_initiation/payment/get | Get payment details
[**PaymentInitiationPaymentList**](PlaidApi.md#PaymentInitiationPaymentList) | **Post** /payment_initiation/payment/list | List payments
[**PaymentInitiationRecipientCreate**](PlaidApi.md#PaymentInitiationRecipientCreate) | **Post** /payment_initiation/recipient/create | Create payment recipient
[**PaymentInitiationRecipientGet**](PlaidApi.md#PaymentInitiationRecipientGet) | **Post** /payment_initiation/recipient/get | Get payment recipient
[**PaymentInitiationRecipientList**](PlaidApi.md#PaymentInitiationRecipientList) | **Post** /payment_initiation/recipient/list | List payment recipients
[**PostIncomeVerificationCreate**](PlaidApi.md#PostIncomeVerificationCreate) | **Post** /income/verification/create | Create an income verification instance
[**ProcessorApexProcessorTokenCreate**](PlaidApi.md#ProcessorApexProcessorTokenCreate) | **Post** /processor/apex/processor_token/create | Create Apex bank account token
[**ProcessorAuthGet**](PlaidApi.md#ProcessorAuthGet) | **Post** /processor/auth/get | Retrieve Auth data
[**ProcessorBalanceGet**](PlaidApi.md#ProcessorBalanceGet) | **Post** /processor/balance/get | Retrieve Balance data
[**ProcessorIdentityGet**](PlaidApi.md#ProcessorIdentityGet) | **Post** /processor/identity/get | Retrieve Identity data
[**ProcessorStripeBankAccountTokenCreate**](PlaidApi.md#ProcessorStripeBankAccountTokenCreate) | **Post** /processor/stripe/bank_account_token/create | Create Stripe bank account token
[**ProcessorTokenCreate**](PlaidApi.md#ProcessorTokenCreate) | **Post** /processor/token/create | Create processor token
[**SandboxBankTransferSimulate**](PlaidApi.md#SandboxBankTransferSimulate) | **Post** /sandbox/bank_transfer/simulate | Simulate a bank transfer event in Sandbox
[**SandboxItemFireWebhook**](PlaidApi.md#SandboxItemFireWebhook) | **Post** /sandbox/item/fire_webhook | Fire a test webhook
[**SandboxItemResetLogin**](PlaidApi.md#SandboxItemResetLogin) | **Post** /sandbox/item/reset_login | Force a Sandbox Item into an error state
[**SandboxItemSetVerificationStatus**](PlaidApi.md#SandboxItemSetVerificationStatus) | **Post** /sandbox/item/set_verification_status | Set verification status for Sandbox account
[**SandboxProcessorTokenCreate**](PlaidApi.md#SandboxProcessorTokenCreate) | **Post** /sandbox/processor_token/create | Create a test Item and processor token
[**SandboxPublicTokenCreate**](PlaidApi.md#SandboxPublicTokenCreate) | **Post** /sandbox/public_token/create | Create a test Item
[**TransactionsGet**](PlaidApi.md#TransactionsGet) | **Post** /transactions/get | Get transaction data
[**TransactionsRefresh**](PlaidApi.md#TransactionsRefresh) | **Post** /transactions/refresh | Refresh transaction data
[**WebhookVerificationKeyGet**](PlaidApi.md#WebhookVerificationKeyGet) | **Post** /webhook_verification_key/get | Get webhook verification key



## AccountsBalanceGet

> AccountsGetResponse AccountsBalanceGet(ctx).AccountsBalanceGetRequest(accountsBalanceGetRequest).Execute()

Retrieve real-time balance data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountsBalanceGetRequest := *openapiclient.NewAccountsBalanceGetRequest("AccessToken_example") // AccountsBalanceGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AccountsBalanceGet(context.Background()).AccountsBalanceGetRequest(accountsBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AccountsBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsBalanceGet`: AccountsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AccountsBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountsBalanceGetRequest** | [**AccountsBalanceGetRequest**](AccountsBalanceGetRequest.md) |  | 

### Return type

[**AccountsGetResponse**](AccountsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGet

> AccountsGetResponse AccountsGet(ctx).AccountsGetRequest(accountsGetRequest).Execute()

Retrieve accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountsGetRequest := *openapiclient.NewAccountsGetRequest("AccessToken_example") // AccountsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AccountsGet(context.Background()).AccountsGetRequest(accountsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGet`: AccountsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountsGetRequest** | [**AccountsGetRequest**](AccountsGetRequest.md) |  | 

### Return type

[**AccountsGetResponse**](AccountsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyCreate

> AssetReportAuditCopyCreateResponse AssetReportAuditCopyCreate(ctx).AssetReportAuditCopyCreateRequest(assetReportAuditCopyCreateRequest).Execute()

Create Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyCreateRequest := *openapiclient.NewAssetReportAuditCopyCreateRequest("AssetReportToken_example", "AuditorId_example") // AssetReportAuditCopyCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportAuditCopyCreate(context.Background()).AssetReportAuditCopyCreateRequest(assetReportAuditCopyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyCreate`: AssetReportAuditCopyCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyCreateRequest** | [**AssetReportAuditCopyCreateRequest**](AssetReportAuditCopyCreateRequest.md) |  | 

### Return type

[**AssetReportAuditCopyCreateResponse**](AssetReportAuditCopyCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyGet

> AssetReportGetResponse AssetReportAuditCopyGet(ctx).AssetReportAuditCopyGetRequest(assetReportAuditCopyGetRequest).Execute()

Retrieve an Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyGetRequest := *openapiclient.NewAssetReportAuditCopyGetRequest("AuditCopyToken_example") // AssetReportAuditCopyGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportAuditCopyGet(context.Background()).AssetReportAuditCopyGetRequest(assetReportAuditCopyGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyGet`: AssetReportGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyGetRequest** | [**AssetReportAuditCopyGetRequest**](AssetReportAuditCopyGetRequest.md) |  | 

### Return type

[**AssetReportGetResponse**](AssetReportGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyRemove

> AssetReportAuditCopyRemoveResponse AssetReportAuditCopyRemove(ctx).AssetReportAuditCopyRemoveRequest(assetReportAuditCopyRemoveRequest).Execute()

Remove Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyRemoveRequest := *openapiclient.NewAssetReportAuditCopyRemoveRequest("AuditCopyToken_example") // AssetReportAuditCopyRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportAuditCopyRemove(context.Background()).AssetReportAuditCopyRemoveRequest(assetReportAuditCopyRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyRemove`: AssetReportAuditCopyRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyRemoveRequest** | [**AssetReportAuditCopyRemoveRequest**](AssetReportAuditCopyRemoveRequest.md) |  | 

### Return type

[**AssetReportAuditCopyRemoveResponse**](AssetReportAuditCopyRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportCreate

> AssetReportCreateResponse AssetReportCreate(ctx).AssetReportCreateRequest(assetReportCreateRequest).Execute()

Create an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportCreateRequest := *openapiclient.NewAssetReportCreateRequest([]string{"AccessTokens_example"}, int32(123)) // AssetReportCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportCreate(context.Background()).AssetReportCreateRequest(assetReportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportCreate`: AssetReportCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportCreateRequest** | [**AssetReportCreateRequest**](AssetReportCreateRequest.md) |  | 

### Return type

[**AssetReportCreateResponse**](AssetReportCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportFilter

> AssetReportFilterResponse AssetReportFilter(ctx).AssetReportFilterRequest(assetReportFilterRequest).Execute()

Filter Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportFilterRequest := *openapiclient.NewAssetReportFilterRequest("AssetReportToken_example", []string{"AccountIdsToExclude_example"}) // AssetReportFilterRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportFilter(context.Background()).AssetReportFilterRequest(assetReportFilterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportFilter`: AssetReportFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportFilterRequest** | [**AssetReportFilterRequest**](AssetReportFilterRequest.md) |  | 

### Return type

[**AssetReportFilterResponse**](AssetReportFilterResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportGet

> AssetReportGetResponse AssetReportGet(ctx).AssetReportGetRequest(assetReportGetRequest).Execute()

Retrieve an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportGetRequest := *openapiclient.NewAssetReportGetRequest("AssetReportToken_example") // AssetReportGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportGet(context.Background()).AssetReportGetRequest(assetReportGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportGet`: AssetReportGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportGetRequest** | [**AssetReportGetRequest**](AssetReportGetRequest.md) |  | 

### Return type

[**AssetReportGetResponse**](AssetReportGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportPdfGet

> *os.File AssetReportPdfGet(ctx).AssetReportPDFGetRequest(assetReportPDFGetRequest).Execute()

Retrieve a PDF Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportPDFGetRequest := *openapiclient.NewAssetReportPDFGetRequest("AssetReportToken_example") // AssetReportPDFGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportPdfGet(context.Background()).AssetReportPDFGetRequest(assetReportPDFGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportPdfGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportPdfGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportPdfGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportPdfGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportPDFGetRequest** | [**AssetReportPDFGetRequest**](AssetReportPDFGetRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportRefresh

> AssetReportRefreshResponse AssetReportRefresh(ctx).AssetReportRefreshRequest(assetReportRefreshRequest).Execute()

Refresh an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportRefreshRequest := *openapiclient.NewAssetReportRefreshRequest("AssetReportToken_example") // AssetReportRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportRefresh(context.Background()).AssetReportRefreshRequest(assetReportRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportRefresh`: AssetReportRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportRefreshRequest** | [**AssetReportRefreshRequest**](AssetReportRefreshRequest.md) |  | 

### Return type

[**AssetReportRefreshResponse**](AssetReportRefreshResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportRemove

> AssetReportRemoveResponse AssetReportRemove(ctx).AssetReportRemoveRequest(assetReportRemoveRequest).Execute()

Delete an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportRemoveRequest := *openapiclient.NewAssetReportRemoveRequest("AssetReportToken_example") // AssetReportRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AssetReportRemove(context.Background()).AssetReportRemoveRequest(assetReportRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportRemove`: AssetReportRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportRemoveRequest** | [**AssetReportRemoveRequest**](AssetReportRemoveRequest.md) |  | 

### Return type

[**AssetReportRemoveResponse**](AssetReportRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthGet

> AuthGetResponse AuthGet(ctx).AuthGetRequest(authGetRequest).Execute()

Retrieve auth data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authGetRequest := *openapiclient.NewAuthGetRequest("AccessToken_example") // AuthGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.AuthGet(context.Background()).AuthGetRequest(authGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthGet`: AuthGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AuthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGetRequest** | [**AuthGetRequest**](AuthGetRequest.md) |  | 

### Return type

[**AuthGetResponse**](AuthGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferBalanceGet

> BankTransferBalanceGetResponse BankTransferBalanceGet(ctx).BankTransferBalanceGetRequest(bankTransferBalanceGetRequest).Execute()

Get balance of your Bank Transfer account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferBalanceGetRequest := *openapiclient.NewBankTransferBalanceGetRequest() // BankTransferBalanceGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferBalanceGet(context.Background()).BankTransferBalanceGetRequest(bankTransferBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferBalanceGet`: BankTransferBalanceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferBalanceGetRequest** | [**BankTransferBalanceGetRequest**](BankTransferBalanceGetRequest.md) |  | 

### Return type

[**BankTransferBalanceGetResponse**](BankTransferBalanceGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferCancel

> BankTransferCancelResponse BankTransferCancel(ctx).BankTransferCancelRequest(bankTransferCancelRequest).Execute()

Cancel a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferCancelRequest := *openapiclient.NewBankTransferCancelRequest("BankTransferId_example") // BankTransferCancelRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferCancel(context.Background()).BankTransferCancelRequest(bankTransferCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferCancel`: BankTransferCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferCancelRequest** | [**BankTransferCancelRequest**](BankTransferCancelRequest.md) |  | 

### Return type

[**BankTransferCancelResponse**](BankTransferCancelResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferCreate

> BankTransferCreateResponse BankTransferCreate(ctx).BankTransferCreateRequest(bankTransferCreateRequest).Execute()

Create a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferCreateRequest := *openapiclient.NewBankTransferCreateRequest("IdempotencyKey_example", "AccessToken_example", "AccountId_example", openapiclient.BankTransferType("debit"), openapiclient.BankTransferNetwork("ach"), "Amount_example", "IsoCurrencyCode_example", "Description_example", *openapiclient.NewBankTransferUser("LegalName_example")) // BankTransferCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferCreate(context.Background()).BankTransferCreateRequest(bankTransferCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferCreate`: BankTransferCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferCreateRequest** | [**BankTransferCreateRequest**](BankTransferCreateRequest.md) |  | 

### Return type

[**BankTransferCreateResponse**](BankTransferCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferEventList

> BankTransferEventListResponse BankTransferEventList(ctx).BankTransferEventListRequest(bankTransferEventListRequest).Execute()

List bank transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferEventListRequest := *openapiclient.NewBankTransferEventListRequest() // BankTransferEventListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferEventList(context.Background()).BankTransferEventListRequest(bankTransferEventListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferEventList`: BankTransferEventListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferEventListRequest** | [**BankTransferEventListRequest**](BankTransferEventListRequest.md) |  | 

### Return type

[**BankTransferEventListResponse**](BankTransferEventListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferEventSync

> BankTransferEventSyncResponse BankTransferEventSync(ctx).BankTransferEventSyncRequest(bankTransferEventSyncRequest).Execute()

Sync bank transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferEventSyncRequest := *openapiclient.NewBankTransferEventSyncRequest(int32(123)) // BankTransferEventSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferEventSync(context.Background()).BankTransferEventSyncRequest(bankTransferEventSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferEventSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferEventSync`: BankTransferEventSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferEventSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferEventSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferEventSyncRequest** | [**BankTransferEventSyncRequest**](BankTransferEventSyncRequest.md) |  | 

### Return type

[**BankTransferEventSyncResponse**](BankTransferEventSyncResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferGet

> BankTransferGetResponse BankTransferGet(ctx).BankTransferGetRequest(bankTransferGetRequest).Execute()

Retrieve a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferGetRequest := *openapiclient.NewBankTransferGetRequest("BankTransferId_example") // BankTransferGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferGet(context.Background()).BankTransferGetRequest(bankTransferGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferGet`: BankTransferGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferGetRequest** | [**BankTransferGetRequest**](BankTransferGetRequest.md) |  | 

### Return type

[**BankTransferGetResponse**](BankTransferGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferList

> BankTransferListResponse BankTransferList(ctx).BankTransferListRequest(bankTransferListRequest).Execute()

List bank transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferListRequest := *openapiclient.NewBankTransferListRequest() // BankTransferListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferList(context.Background()).BankTransferListRequest(bankTransferListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferList`: BankTransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferListRequest** | [**BankTransferListRequest**](BankTransferListRequest.md) |  | 

### Return type

[**BankTransferListResponse**](BankTransferListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferMigrateAccount

> BankTransferMigrateAccountResponse BankTransferMigrateAccount(ctx).BankTransferMigrateAccountRequest(bankTransferMigrateAccountRequest).Execute()

Migrate account into Bank Transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferMigrateAccountRequest := *openapiclient.NewBankTransferMigrateAccountRequest("AccountNumber_example", "RoutingNumber_example", "AccountType_example") // BankTransferMigrateAccountRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.BankTransferMigrateAccount(context.Background()).BankTransferMigrateAccountRequest(bankTransferMigrateAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferMigrateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferMigrateAccount`: BankTransferMigrateAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferMigrateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferMigrateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferMigrateAccountRequest** | [**BankTransferMigrateAccountRequest**](BankTransferMigrateAccountRequest.md) |  | 

### Return type

[**BankTransferMigrateAccountResponse**](BankTransferMigrateAccountResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesGet

> CategoriesGetResponse CategoriesGet(ctx).Body(body).Execute()

Get Categories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}(Object) // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.CategoriesGet(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.CategoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoriesGet`: CategoriesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.CategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CategoriesGetResponse**](CategoriesGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentToken

> PaymentInitiationPaymentTokenCreateResponse CreatePaymentToken(ctx).PaymentInitiationPaymentTokenCreateRequest(paymentInitiationPaymentTokenCreateRequest).Execute()

Create payment token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentTokenCreateRequest := *openapiclient.NewPaymentInitiationPaymentTokenCreateRequest("PaymentId_example") // PaymentInitiationPaymentTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.CreatePaymentToken(context.Background()).PaymentInitiationPaymentTokenCreateRequest(paymentInitiationPaymentTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.CreatePaymentToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentToken`: PaymentInitiationPaymentTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.CreatePaymentToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentTokenCreateRequest** | [**PaymentInitiationPaymentTokenCreateRequest**](PaymentInitiationPaymentTokenCreateRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentTokenCreateResponse**](PaymentInitiationPaymentTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchAltCreate

> DepositSwitchAltCreateResponse DepositSwitchAltCreate(ctx).DepositSwitchAltCreateRequest(depositSwitchAltCreateRequest).Execute()

Create a deposit switch when not using Plaid Exchange.'



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchAltCreateRequest := *openapiclient.NewDepositSwitchAltCreateRequest(*openapiclient.NewDepositSwitchTargetAccount("AccountNumber_example", "RoutingNumber_example", "AccountName_example", "AccountSubtype_example"), *openapiclient.NewDepositSwitchTargetUser("GivenName_example", "FamilyName_example", "Phone_example", "Email_example")) // DepositSwitchAltCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.DepositSwitchAltCreate(context.Background()).DepositSwitchAltCreateRequest(depositSwitchAltCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchAltCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchAltCreate`: DepositSwitchAltCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchAltCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchAltCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchAltCreateRequest** | [**DepositSwitchAltCreateRequest**](DepositSwitchAltCreateRequest.md) |  | 

### Return type

[**DepositSwitchAltCreateResponse**](DepositSwitchAltCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchCreate

> DepositSwitchCreateResponse DepositSwitchCreate(ctx).DepositSwitchCreateRequest(depositSwitchCreateRequest).Execute()

Create a deposit switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchCreateRequest := *openapiclient.NewDepositSwitchCreateRequest("TargetAccessToken_example", "TargetAccountId_example") // DepositSwitchCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.DepositSwitchCreate(context.Background()).DepositSwitchCreateRequest(depositSwitchCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchCreate`: DepositSwitchCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchCreateRequest** | [**DepositSwitchCreateRequest**](DepositSwitchCreateRequest.md) |  | 

### Return type

[**DepositSwitchCreateResponse**](DepositSwitchCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchGet

> DepositSwitchGetResponse DepositSwitchGet(ctx).DepositSwitchGetRequest(depositSwitchGetRequest).Execute()

Retrieve a deposit switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchGetRequest := *openapiclient.NewDepositSwitchGetRequest("DepositSwitchId_example") // DepositSwitchGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.DepositSwitchGet(context.Background()).DepositSwitchGetRequest(depositSwitchGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchGet`: DepositSwitchGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchGetRequest** | [**DepositSwitchGetRequest**](DepositSwitchGetRequest.md) |  | 

### Return type

[**DepositSwitchGetResponse**](DepositSwitchGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchTokenCreate

> DepositSwitchTokenCreateResponse DepositSwitchTokenCreate(ctx).DepositSwitchTokenCreateRequest(depositSwitchTokenCreateRequest).Execute()

Create a deposit switch token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchTokenCreateRequest := *openapiclient.NewDepositSwitchTokenCreateRequest("DepositSwitchId_example") // DepositSwitchTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.DepositSwitchTokenCreate(context.Background()).DepositSwitchTokenCreateRequest(depositSwitchTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchTokenCreate`: DepositSwitchTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchTokenCreateRequest** | [**DepositSwitchTokenCreateRequest**](DepositSwitchTokenCreateRequest.md) |  | 

### Return type

[**DepositSwitchTokenCreateResponse**](DepositSwitchTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployersSearch

> EmployersSearchResponse EmployersSearch(ctx).EmployersSearchRequest(employersSearchRequest).Execute()

Search employer database



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    employersSearchRequest := *openapiclient.NewEmployersSearchRequest("Query_example", []string{"Products_example"}) // EmployersSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.EmployersSearch(context.Background()).EmployersSearchRequest(employersSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.EmployersSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployersSearch`: EmployersSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.EmployersSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployersSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employersSearchRequest** | [**EmployersSearchRequest**](EmployersSearchRequest.md) |  | 

### Return type

[**EmployersSearchResponse**](EmployersSearchResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGet

> IdentityGetResponse IdentityGet(ctx).IdentityGetRequest(identityGetRequest).Execute()

Retrieve identity data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityGetRequest := *openapiclient.NewIdentityGetRequest("AccessToken_example") // IdentityGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.IdentityGet(context.Background()).IdentityGetRequest(identityGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGet`: IdentityGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IdentityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityGetRequest** | [**IdentityGetRequest**](IdentityGetRequest.md) |  | 

### Return type

[**IdentityGetResponse**](IdentityGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationDocumentsDownload

> *os.File IncomeVerificationDocumentsDownload(ctx).IncomeVerificationDocumentsDownloadRequest(incomeVerificationDocumentsDownloadRequest).Execute()

Download the original documents used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationDocumentsDownloadRequest := *openapiclient.NewIncomeVerificationDocumentsDownloadRequest("IncomeVerificationId_example") // IncomeVerificationDocumentsDownloadRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.IncomeVerificationDocumentsDownload(context.Background()).IncomeVerificationDocumentsDownloadRequest(incomeVerificationDocumentsDownloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationDocumentsDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationDocumentsDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationDocumentsDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationDocumentsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationDocumentsDownloadRequest** | [**IncomeVerificationDocumentsDownloadRequest**](IncomeVerificationDocumentsDownloadRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationPaystubGet

> IncomeVerificationPaystubGetResponse IncomeVerificationPaystubGet(ctx).IncomeVerificationPaystubGetRequest(incomeVerificationPaystubGetRequest).Execute()

Retrieve information from the paystub used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationPaystubGetRequest := *openapiclient.NewIncomeVerificationPaystubGetRequest("IncomeVerificationId_example") // IncomeVerificationPaystubGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.IncomeVerificationPaystubGet(context.Background()).IncomeVerificationPaystubGetRequest(incomeVerificationPaystubGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationPaystubGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationPaystubGet`: IncomeVerificationPaystubGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationPaystubGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationPaystubGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationPaystubGetRequest** | [**IncomeVerificationPaystubGetRequest**](IncomeVerificationPaystubGetRequest.md) |  | 

### Return type

[**IncomeVerificationPaystubGetResponse**](IncomeVerificationPaystubGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationSummaryGet

> IncomeVerificationSummaryGetResponse IncomeVerificationSummaryGet(ctx).IncomeVerificationSummaryGetRequest(incomeVerificationSummaryGetRequest).Execute()

Retrieve a summary of information derived from income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationSummaryGetRequest := *openapiclient.NewIncomeVerificationSummaryGetRequest("IncomeVerificationId_example") // IncomeVerificationSummaryGetRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.IncomeVerificationSummaryGet(context.Background()).IncomeVerificationSummaryGetRequest(incomeVerificationSummaryGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationSummaryGet`: IncomeVerificationSummaryGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationSummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationSummaryGetRequest** | [**IncomeVerificationSummaryGetRequest**](IncomeVerificationSummaryGetRequest.md) |  | 

### Return type

[**IncomeVerificationSummaryGetResponse**](IncomeVerificationSummaryGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsGet

> InstitutionsGetResponse InstitutionsGet(ctx).InstitutionsGetRequest(institutionsGetRequest).Execute()

Get details of all supported institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsGetRequest := *openapiclient.NewInstitutionsGetRequest(int32(123), int32(123), []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.InstitutionsGet(context.Background()).InstitutionsGetRequest(institutionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsGet`: InstitutionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsGetRequest** | [**InstitutionsGetRequest**](InstitutionsGetRequest.md) |  | 

### Return type

[**InstitutionsGetResponse**](InstitutionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsGetById

> InstitutionsGetByIdResponse InstitutionsGetById(ctx).InstitutionsGetByIdRequest(institutionsGetByIdRequest).Execute()

Get details of an institution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsGetByIdRequest := *openapiclient.NewInstitutionsGetByIdRequest("InstitutionId_example", []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsGetByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.InstitutionsGetById(context.Background()).InstitutionsGetByIdRequest(institutionsGetByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsGetById`: InstitutionsGetByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsGetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsGetByIdRequest** | [**InstitutionsGetByIdRequest**](InstitutionsGetByIdRequest.md) |  | 

### Return type

[**InstitutionsGetByIdResponse**](InstitutionsGetByIdResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsSearch

> InstitutionsSearchResponse InstitutionsSearch(ctx).InstitutionsSearchRequest(institutionsSearchRequest).Execute()

Search institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsSearchRequest := *openapiclient.NewInstitutionsSearchRequest("Query_example", []openapiclient.Products{openapiclient.Products("assets")}, []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.InstitutionsSearch(context.Background()).InstitutionsSearchRequest(institutionsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsSearch`: InstitutionsSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsSearchRequest** | [**InstitutionsSearchRequest**](InstitutionsSearchRequest.md) |  | 

### Return type

[**InstitutionsSearchResponse**](InstitutionsSearchResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestmentsHoldingsGet

> InvestmentsHoldingsGetResponse InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(investmentsHoldingsGetRequest).Execute()

Get Investment holdings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investmentsHoldingsGetRequest := *openapiclient.NewInvestmentsHoldingsGetRequest("AccessToken_example") // InvestmentsHoldingsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.InvestmentsHoldingsGet(context.Background()).InvestmentsHoldingsGetRequest(investmentsHoldingsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InvestmentsHoldingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestmentsHoldingsGet`: InvestmentsHoldingsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InvestmentsHoldingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestmentsHoldingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investmentsHoldingsGetRequest** | [**InvestmentsHoldingsGetRequest**](InvestmentsHoldingsGetRequest.md) |  | 

### Return type

[**InvestmentsHoldingsGetResponse**](InvestmentsHoldingsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestmentsTransactionsGet

> InvestmentsTransactionsGetResponse InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(investmentsTransactionsGetRequest).Execute()

Get investment transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    investmentsTransactionsGetRequest := *openapiclient.NewInvestmentsTransactionsGetRequest("AccessToken_example", time.Now(), time.Now()) // InvestmentsTransactionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.InvestmentsTransactionsGet(context.Background()).InvestmentsTransactionsGetRequest(investmentsTransactionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InvestmentsTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestmentsTransactionsGet`: InvestmentsTransactionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InvestmentsTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestmentsTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investmentsTransactionsGetRequest** | [**InvestmentsTransactionsGetRequest**](InvestmentsTransactionsGetRequest.md) |  | 

### Return type

[**InvestmentsTransactionsGetResponse**](InvestmentsTransactionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemAccessTokenInvalidate

> ItemAccessTokenInvalidateResponse ItemAccessTokenInvalidate(ctx).ItemAccessTokenInvalidateRequest(itemAccessTokenInvalidateRequest).Execute()

Invalidate access_token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemAccessTokenInvalidateRequest := *openapiclient.NewItemAccessTokenInvalidateRequest("AccessToken_example") // ItemAccessTokenInvalidateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemAccessTokenInvalidate(context.Background()).ItemAccessTokenInvalidateRequest(itemAccessTokenInvalidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemAccessTokenInvalidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemAccessTokenInvalidate`: ItemAccessTokenInvalidateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemAccessTokenInvalidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemAccessTokenInvalidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemAccessTokenInvalidateRequest** | [**ItemAccessTokenInvalidateRequest**](ItemAccessTokenInvalidateRequest.md) |  | 

### Return type

[**ItemAccessTokenInvalidateResponse**](ItemAccessTokenInvalidateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemCreatePublicToken

> ItemPublicTokenCreateResponse ItemCreatePublicToken(ctx).ItemPublicTokenCreateRequest(itemPublicTokenCreateRequest).Execute()

Create public token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemPublicTokenCreateRequest := *openapiclient.NewItemPublicTokenCreateRequest("AccessToken_example") // ItemPublicTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemCreatePublicToken(context.Background()).ItemPublicTokenCreateRequest(itemPublicTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemCreatePublicToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemCreatePublicToken`: ItemPublicTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemCreatePublicToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemCreatePublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemPublicTokenCreateRequest** | [**ItemPublicTokenCreateRequest**](ItemPublicTokenCreateRequest.md) |  | 

### Return type

[**ItemPublicTokenCreateResponse**](ItemPublicTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemGet

> ItemGetResponse ItemGet(ctx).ItemGetRequest(itemGetRequest).Execute()

Retrieve an Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemGetRequest := *openapiclient.NewItemGetRequest("AccessToken_example") // ItemGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemGet(context.Background()).ItemGetRequest(itemGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemGet`: ItemGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemGetRequest** | [**ItemGetRequest**](ItemGetRequest.md) |  | 

### Return type

[**ItemGetResponse**](ItemGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemImport

> ItemImportResponse ItemImport(ctx).ItemImportRequest(itemImportRequest).Execute()

Import Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemImportRequest := *openapiclient.NewItemImportRequest([]openapiclient.Products{openapiclient.Products("assets")}, *openapiclient.NewItemImportRequestUserAuth("UserId_example", "AuthToken_example")) // ItemImportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemImport(context.Background()).ItemImportRequest(itemImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemImport`: ItemImportResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemImportRequest** | [**ItemImportRequest**](ItemImportRequest.md) |  | 

### Return type

[**ItemImportResponse**](ItemImportResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemPublicTokenExchange

> ItemPublicTokenExchangeResponse ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(itemPublicTokenExchangeRequest).Execute()

Exchange public token for an access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemPublicTokenExchangeRequest := *openapiclient.NewItemPublicTokenExchangeRequest("PublicToken_example") // ItemPublicTokenExchangeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemPublicTokenExchange(context.Background()).ItemPublicTokenExchangeRequest(itemPublicTokenExchangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemPublicTokenExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemPublicTokenExchange`: ItemPublicTokenExchangeResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemPublicTokenExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemPublicTokenExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemPublicTokenExchangeRequest** | [**ItemPublicTokenExchangeRequest**](ItemPublicTokenExchangeRequest.md) |  | 

### Return type

[**ItemPublicTokenExchangeResponse**](ItemPublicTokenExchangeResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemRemove

> ItemRemoveResponse ItemRemove(ctx).ItemRemoveRequest(itemRemoveRequest).Execute()

Remove an Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemRemoveRequest := *openapiclient.NewItemRemoveRequest("AccessToken_example") // ItemRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemRemove(context.Background()).ItemRemoveRequest(itemRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemRemove`: ItemRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemRemoveRequest** | [**ItemRemoveRequest**](ItemRemoveRequest.md) |  | 

### Return type

[**ItemRemoveResponse**](ItemRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemWebhookUpdate

> ItemWebhookUpdateResponse ItemWebhookUpdate(ctx).ItemWebhookUpdateRequest(itemWebhookUpdateRequest).Execute()

Update Webhook URL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemWebhookUpdateRequest := *openapiclient.NewItemWebhookUpdateRequest("AccessToken_example", "Webhook_example") // ItemWebhookUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ItemWebhookUpdate(context.Background()).ItemWebhookUpdateRequest(itemWebhookUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemWebhookUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemWebhookUpdate`: ItemWebhookUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemWebhookUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemWebhookUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemWebhookUpdateRequest** | [**ItemWebhookUpdateRequest**](ItemWebhookUpdateRequest.md) |  | 

### Return type

[**ItemWebhookUpdateResponse**](ItemWebhookUpdateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiabilitiesGet

> LiabilitiesGetResponse LiabilitiesGet(ctx).LiabilitiesGetRequest(liabilitiesGetRequest).Execute()

Retrieve Liabilities data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    liabilitiesGetRequest := *openapiclient.NewLiabilitiesGetRequest("AccessToken_example") // LiabilitiesGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.LiabilitiesGet(context.Background()).LiabilitiesGetRequest(liabilitiesGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LiabilitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiabilitiesGet`: LiabilitiesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LiabilitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiabilitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liabilitiesGetRequest** | [**LiabilitiesGetRequest**](LiabilitiesGetRequest.md) |  | 

### Return type

[**LiabilitiesGetResponse**](LiabilitiesGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkTokenCreate

> LinkTokenCreateResponse LinkTokenCreate(ctx).LinkTokenCreateRequest(linkTokenCreateRequest).Execute()

Create Link Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    linkTokenCreateRequest := *openapiclient.NewLinkTokenCreateRequest("ClientName_example", "Language_example", []openapiclient.CountryCode{openapiclient.CountryCode("US")}, *openapiclient.NewLinkTokenCreateRequestUser("ClientUserId_example")) // LinkTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.LinkTokenCreate(context.Background()).LinkTokenCreateRequest(linkTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LinkTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkTokenCreate`: LinkTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LinkTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkTokenCreateRequest** | [**LinkTokenCreateRequest**](LinkTokenCreateRequest.md) |  | 

### Return type

[**LinkTokenCreateResponse**](LinkTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkTokenGet

> LinkTokenGetResponse LinkTokenGet(ctx).LinkTokenGetRequest(linkTokenGetRequest).Execute()

Get Link Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    linkTokenGetRequest := *openapiclient.NewLinkTokenGetRequest("LinkToken_example") // LinkTokenGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.LinkTokenGet(context.Background()).LinkTokenGetRequest(linkTokenGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LinkTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkTokenGet`: LinkTokenGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LinkTokenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkTokenGetRequest** | [**LinkTokenGetRequest**](LinkTokenGetRequest.md) |  | 

### Return type

[**LinkTokenGetResponse**](LinkTokenGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentCreate

> PaymentInitiationPaymentCreateResponse PaymentInitiationPaymentCreate(ctx).PaymentInitiationPaymentCreateRequest(paymentInitiationPaymentCreateRequest).Execute()

Create a payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentCreateRequest := *openapiclient.NewPaymentInitiationPaymentCreateRequest("RecipientId_example", "Reference_example", *openapiclient.NewAmount("Currency_example", float32(123))) // PaymentInitiationPaymentCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationPaymentCreate(context.Background()).PaymentInitiationPaymentCreateRequest(paymentInitiationPaymentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentCreate`: PaymentInitiationPaymentCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentCreateRequest** | [**PaymentInitiationPaymentCreateRequest**](PaymentInitiationPaymentCreateRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentCreateResponse**](PaymentInitiationPaymentCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentGet

> PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGet(ctx).PaymentInitiationPaymentGetRequest(paymentInitiationPaymentGetRequest).Execute()

Get payment details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentGetRequest := *openapiclient.NewPaymentInitiationPaymentGetRequest("PaymentId_example") // PaymentInitiationPaymentGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationPaymentGet(context.Background()).PaymentInitiationPaymentGetRequest(paymentInitiationPaymentGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentGet`: PaymentInitiationPaymentGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentGetRequest** | [**PaymentInitiationPaymentGetRequest**](PaymentInitiationPaymentGetRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentGetResponse**](PaymentInitiationPaymentGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentList

> PaymentInitiationPaymentListResponse PaymentInitiationPaymentList(ctx).PaymentInitiationPaymentListRequest(paymentInitiationPaymentListRequest).Execute()

List payments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentListRequest := *openapiclient.NewPaymentInitiationPaymentListRequest() // PaymentInitiationPaymentListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationPaymentList(context.Background()).PaymentInitiationPaymentListRequest(paymentInitiationPaymentListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentList`: PaymentInitiationPaymentListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentListRequest** | [**PaymentInitiationPaymentListRequest**](PaymentInitiationPaymentListRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentListResponse**](PaymentInitiationPaymentListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientCreate

> PaymentInitiationRecipientCreateResponse PaymentInitiationRecipientCreate(ctx).PaymentInitiationRecipientCreateRequest(paymentInitiationRecipientCreateRequest).Execute()

Create payment recipient



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientCreateRequest := *openapiclient.NewPaymentInitiationRecipientCreateRequest("Name_example") // PaymentInitiationRecipientCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationRecipientCreate(context.Background()).PaymentInitiationRecipientCreateRequest(paymentInitiationRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientCreate`: PaymentInitiationRecipientCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientCreateRequest** | [**PaymentInitiationRecipientCreateRequest**](PaymentInitiationRecipientCreateRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientCreateResponse**](PaymentInitiationRecipientCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientGet

> PaymentInitiationRecipientGetResponse PaymentInitiationRecipientGet(ctx).PaymentInitiationRecipientGetRequest(paymentInitiationRecipientGetRequest).Execute()

Get payment recipient



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientGetRequest := *openapiclient.NewPaymentInitiationRecipientGetRequest("RecipientId_example") // PaymentInitiationRecipientGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationRecipientGet(context.Background()).PaymentInitiationRecipientGetRequest(paymentInitiationRecipientGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientGet`: PaymentInitiationRecipientGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientGetRequest** | [**PaymentInitiationRecipientGetRequest**](PaymentInitiationRecipientGetRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientGetResponse**](PaymentInitiationRecipientGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientList

> PaymentInitiationRecipientListResponse PaymentInitiationRecipientList(ctx).PaymentInitiationRecipientListRequest(paymentInitiationRecipientListRequest).Execute()

List payment recipients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientListRequest := *openapiclient.NewPaymentInitiationRecipientListRequest() // PaymentInitiationRecipientListRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PaymentInitiationRecipientList(context.Background()).PaymentInitiationRecipientListRequest(paymentInitiationRecipientListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientList`: PaymentInitiationRecipientListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientListRequest** | [**PaymentInitiationRecipientListRequest**](PaymentInitiationRecipientListRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientListResponse**](PaymentInitiationRecipientListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIncomeVerificationCreate

> IncomeVerificationCreateResponse PostIncomeVerificationCreate(ctx).IncomeVerificationCreateRequest(incomeVerificationCreateRequest).Execute()

Create an income verification instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationCreateRequest := *openapiclient.NewIncomeVerificationCreateRequest("Webhook_example") // IncomeVerificationCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.PostIncomeVerificationCreate(context.Background()).IncomeVerificationCreateRequest(incomeVerificationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PostIncomeVerificationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostIncomeVerificationCreate`: IncomeVerificationCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PostIncomeVerificationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostIncomeVerificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationCreateRequest** | [**IncomeVerificationCreateRequest**](IncomeVerificationCreateRequest.md) |  | 

### Return type

[**IncomeVerificationCreateResponse**](IncomeVerificationCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorApexProcessorTokenCreate

> ProcessorTokenCreateResponse ProcessorApexProcessorTokenCreate(ctx).ProcessorApexProcessorTokenCreateRequest(processorApexProcessorTokenCreateRequest).Execute()

Create Apex bank account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorApexProcessorTokenCreateRequest := *openapiclient.NewProcessorApexProcessorTokenCreateRequest("AccessToken_example", "AccountId_example") // ProcessorApexProcessorTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorApexProcessorTokenCreate(context.Background()).ProcessorApexProcessorTokenCreateRequest(processorApexProcessorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorApexProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorApexProcessorTokenCreate`: ProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorApexProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorApexProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorApexProcessorTokenCreateRequest** | [**ProcessorApexProcessorTokenCreateRequest**](ProcessorApexProcessorTokenCreateRequest.md) |  | 

### Return type

[**ProcessorTokenCreateResponse**](ProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorAuthGet

> ProcessorAuthGetResponse ProcessorAuthGet(ctx).ProcessorAuthGetRequest(processorAuthGetRequest).Execute()

Retrieve Auth data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorAuthGetRequest := *openapiclient.NewProcessorAuthGetRequest("ProcessorToken_example") // ProcessorAuthGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorAuthGet(context.Background()).ProcessorAuthGetRequest(processorAuthGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorAuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorAuthGet`: ProcessorAuthGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorAuthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorAuthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorAuthGetRequest** | [**ProcessorAuthGetRequest**](ProcessorAuthGetRequest.md) |  | 

### Return type

[**ProcessorAuthGetResponse**](ProcessorAuthGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorBalanceGet

> ProcessorBalanceGetResponse ProcessorBalanceGet(ctx).ProcessorBalanceGetRequest(processorBalanceGetRequest).Execute()

Retrieve Balance data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorBalanceGetRequest := *openapiclient.NewProcessorBalanceGetRequest("ProcessorToken_example") // ProcessorBalanceGetRequest | The `/processor/balance/get` endpoint returns the real-time balance for the account associated with a given `processor_token`.  The current balance is the total amount of funds in the account. The available balance is the current balance less any outstanding holds or debits that have not yet posted to the account.  Note that not all institutions calculate the available balance. In the event that available balance is unavailable from the institution, Plaid will return an available balance value of `null`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorBalanceGet(context.Background()).ProcessorBalanceGetRequest(processorBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorBalanceGet`: ProcessorBalanceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorBalanceGetRequest** | [**ProcessorBalanceGetRequest**](ProcessorBalanceGetRequest.md) | The &#x60;/processor/balance/get&#x60; endpoint returns the real-time balance for the account associated with a given &#x60;processor_token&#x60;.  The current balance is the total amount of funds in the account. The available balance is the current balance less any outstanding holds or debits that have not yet posted to the account.  Note that not all institutions calculate the available balance. In the event that available balance is unavailable from the institution, Plaid will return an available balance value of &#x60;null&#x60;. | 

### Return type

[**ProcessorBalanceGetResponse**](ProcessorBalanceGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorIdentityGet

> ProcessorIdentityGetResponse ProcessorIdentityGet(ctx).ProcessorIdentityGetRequest(processorIdentityGetRequest).Execute()

Retrieve Identity data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorIdentityGetRequest := *openapiclient.NewProcessorIdentityGetRequest("ProcessorToken_example") // ProcessorIdentityGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorIdentityGet(context.Background()).ProcessorIdentityGetRequest(processorIdentityGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorIdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorIdentityGet`: ProcessorIdentityGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorIdentityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorIdentityGetRequest** | [**ProcessorIdentityGetRequest**](ProcessorIdentityGetRequest.md) |  | 

### Return type

[**ProcessorIdentityGetResponse**](ProcessorIdentityGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorStripeBankAccountTokenCreate

> ProcessorStripeBankAccountTokenCreateResponse ProcessorStripeBankAccountTokenCreate(ctx).ProcessorStripeBankAccountTokenCreateRequest(processorStripeBankAccountTokenCreateRequest).Execute()

Create Stripe bank account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorStripeBankAccountTokenCreateRequest := *openapiclient.NewProcessorStripeBankAccountTokenCreateRequest("AccessToken_example", "AccountId_example") // ProcessorStripeBankAccountTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorStripeBankAccountTokenCreate(context.Background()).ProcessorStripeBankAccountTokenCreateRequest(processorStripeBankAccountTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorStripeBankAccountTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorStripeBankAccountTokenCreate`: ProcessorStripeBankAccountTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorStripeBankAccountTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorStripeBankAccountTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorStripeBankAccountTokenCreateRequest** | [**ProcessorStripeBankAccountTokenCreateRequest**](ProcessorStripeBankAccountTokenCreateRequest.md) |  | 

### Return type

[**ProcessorStripeBankAccountTokenCreateResponse**](ProcessorStripeBankAccountTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorTokenCreate

> ProcessorTokenCreateResponse ProcessorTokenCreate(ctx).ProcessorTokenCreateRequest(processorTokenCreateRequest).Execute()

Create processor token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorTokenCreateRequest := *openapiclient.NewProcessorTokenCreateRequest("AccessToken_example", "AccountId_example", "Processor_example") // ProcessorTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.ProcessorTokenCreate(context.Background()).ProcessorTokenCreateRequest(processorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorTokenCreate`: ProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorTokenCreateRequest** | [**ProcessorTokenCreateRequest**](ProcessorTokenCreateRequest.md) |  | 

### Return type

[**ProcessorTokenCreateResponse**](ProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxBankTransferSimulate

> SandboxBankTransferSimulateResponse SandboxBankTransferSimulate(ctx).SandboxBankTransferSimulateRequest(sandboxBankTransferSimulateRequest).Execute()

Simulate a bank transfer event in Sandbox



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxBankTransferSimulateRequest := *openapiclient.NewSandboxBankTransferSimulateRequest("BankTransferId_example", "EventType_example") // SandboxBankTransferSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxBankTransferSimulate(context.Background()).SandboxBankTransferSimulateRequest(sandboxBankTransferSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxBankTransferSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxBankTransferSimulate`: SandboxBankTransferSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxBankTransferSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxBankTransferSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxBankTransferSimulateRequest** | [**SandboxBankTransferSimulateRequest**](SandboxBankTransferSimulateRequest.md) |  | 

### Return type

[**SandboxBankTransferSimulateResponse**](SandboxBankTransferSimulateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemFireWebhook

> SandboxItemFireWebhookResponse SandboxItemFireWebhook(ctx).SandboxItemFireWebhookRequest(sandboxItemFireWebhookRequest).Execute()

Fire a test webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemFireWebhookRequest := *openapiclient.NewSandboxItemFireWebhookRequest("AccessToken_example") // SandboxItemFireWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxItemFireWebhook(context.Background()).SandboxItemFireWebhookRequest(sandboxItemFireWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemFireWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemFireWebhook`: SandboxItemFireWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemFireWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemFireWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemFireWebhookRequest** | [**SandboxItemFireWebhookRequest**](SandboxItemFireWebhookRequest.md) |  | 

### Return type

[**SandboxItemFireWebhookResponse**](SandboxItemFireWebhookResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemResetLogin

> SandboxItemResetLoginResponse SandboxItemResetLogin(ctx).SandboxItemResetLoginRequest(sandboxItemResetLoginRequest).Execute()

Force a Sandbox Item into an error state



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemResetLoginRequest := *openapiclient.NewSandboxItemResetLoginRequest("AccessToken_example") // SandboxItemResetLoginRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxItemResetLogin(context.Background()).SandboxItemResetLoginRequest(sandboxItemResetLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemResetLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemResetLogin`: SandboxItemResetLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemResetLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemResetLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemResetLoginRequest** | [**SandboxItemResetLoginRequest**](SandboxItemResetLoginRequest.md) |  | 

### Return type

[**SandboxItemResetLoginResponse**](SandboxItemResetLoginResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemSetVerificationStatus

> SandboxItemSetVerificationStatusResponse SandboxItemSetVerificationStatus(ctx).SandboxItemSetVerificationStatusRequest(sandboxItemSetVerificationStatusRequest).Execute()

Set verification status for Sandbox account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemSetVerificationStatusRequest := *openapiclient.NewSandboxItemSetVerificationStatusRequest("AccessToken_example", "AccountId_example", "VerificationStatus_example") // SandboxItemSetVerificationStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxItemSetVerificationStatus(context.Background()).SandboxItemSetVerificationStatusRequest(sandboxItemSetVerificationStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemSetVerificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemSetVerificationStatus`: SandboxItemSetVerificationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemSetVerificationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemSetVerificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemSetVerificationStatusRequest** | [**SandboxItemSetVerificationStatusRequest**](SandboxItemSetVerificationStatusRequest.md) |  | 

### Return type

[**SandboxItemSetVerificationStatusResponse**](SandboxItemSetVerificationStatusResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxProcessorTokenCreate

> SandboxProcessorTokenCreateResponse SandboxProcessorTokenCreate(ctx).SandboxProcessorTokenCreateRequest(sandboxProcessorTokenCreateRequest).Execute()

Create a test Item and processor token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxProcessorTokenCreateRequest := *openapiclient.NewSandboxProcessorTokenCreateRequest("InstitutionId_example") // SandboxProcessorTokenCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxProcessorTokenCreate(context.Background()).SandboxProcessorTokenCreateRequest(sandboxProcessorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxProcessorTokenCreate`: SandboxProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxProcessorTokenCreateRequest** | [**SandboxProcessorTokenCreateRequest**](SandboxProcessorTokenCreateRequest.md) |  | 

### Return type

[**SandboxProcessorTokenCreateResponse**](SandboxProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxPublicTokenCreate

> SandboxPublicTokenCreateResponse SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(sandboxPublicTokenCreateRequest).Execute()

Create a test Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxPublicTokenCreateRequest := *openapiclient.NewSandboxPublicTokenCreateRequest("InstitutionId_example", []openapiclient.Products{openapiclient.Products("assets")}) // SandboxPublicTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.SandboxPublicTokenCreate(context.Background()).SandboxPublicTokenCreateRequest(sandboxPublicTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxPublicTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxPublicTokenCreate`: SandboxPublicTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxPublicTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxPublicTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxPublicTokenCreateRequest** | [**SandboxPublicTokenCreateRequest**](SandboxPublicTokenCreateRequest.md) |  | 

### Return type

[**SandboxPublicTokenCreateResponse**](SandboxPublicTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> TransactionsGetResponse TransactionsGet(ctx).TransactionsGetRequest(transactionsGetRequest).Execute()

Get transaction data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    transactionsGetRequest := *openapiclient.NewTransactionsGetRequest("AccessToken_example", time.Now(), time.Now()) // TransactionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.TransactionsGet(context.Background()).TransactionsGetRequest(transactionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsGet`: TransactionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsGetRequest** | [**TransactionsGetRequest**](TransactionsGetRequest.md) |  | 

### Return type

[**TransactionsGetResponse**](TransactionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsRefresh

> TransactionsRefreshResponse TransactionsRefresh(ctx).TransactionsRefreshRequest(transactionsRefreshRequest).Execute()

Refresh transaction data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionsRefreshRequest := *openapiclient.NewTransactionsRefreshRequest() // TransactionsRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.TransactionsRefresh(context.Background()).TransactionsRefreshRequest(transactionsRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsRefresh`: TransactionsRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsRefreshRequest** | [**TransactionsRefreshRequest**](TransactionsRefreshRequest.md) |  | 

### Return type

[**TransactionsRefreshResponse**](TransactionsRefreshResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookVerificationKeyGet

> WebhookVerificationKeyGetResponse WebhookVerificationKeyGet(ctx).WebhookVerificationKeyGetRequest(webhookVerificationKeyGetRequest).Execute()

Get webhook verification key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhookVerificationKeyGetRequest := *openapiclient.NewWebhookVerificationKeyGetRequest("KeyId_example") // WebhookVerificationKeyGetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PlaidApi.WebhookVerificationKeyGet(context.Background()).WebhookVerificationKeyGetRequest(webhookVerificationKeyGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.WebhookVerificationKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookVerificationKeyGet`: WebhookVerificationKeyGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.WebhookVerificationKeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookVerificationKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookVerificationKeyGetRequest** | [**WebhookVerificationKeyGetRequest**](WebhookVerificationKeyGetRequest.md) |  | 

### Return type

[**WebhookVerificationKeyGetResponse**](WebhookVerificationKeyGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

