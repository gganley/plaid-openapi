# InvestmentsTransactionsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**StartDate** | **string** | The earliest date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD. | 
**EndDate** | **string** | The most recent date for which to fetch transaction history. Dates should be formatted as YYYY-MM-DD. | 
**Options** | Pointer to [**InvestmentsTransactionsGetRequestOptions**](InvestmentsTransactionsGetRequestOptions.md) |  | [optional] 

## Methods

### NewInvestmentsTransactionsGetRequest

`func NewInvestmentsTransactionsGetRequest(accessToken string, startDate string, endDate string, ) *InvestmentsTransactionsGetRequest`

NewInvestmentsTransactionsGetRequest instantiates a new InvestmentsTransactionsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentsTransactionsGetRequestWithDefaults

`func NewInvestmentsTransactionsGetRequestWithDefaults() *InvestmentsTransactionsGetRequest`

NewInvestmentsTransactionsGetRequestWithDefaults instantiates a new InvestmentsTransactionsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *InvestmentsTransactionsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *InvestmentsTransactionsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *InvestmentsTransactionsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *InvestmentsTransactionsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *InvestmentsTransactionsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *InvestmentsTransactionsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *InvestmentsTransactionsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *InvestmentsTransactionsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *InvestmentsTransactionsGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *InvestmentsTransactionsGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *InvestmentsTransactionsGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetStartDate

`func (o *InvestmentsTransactionsGetRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InvestmentsTransactionsGetRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *InvestmentsTransactionsGetRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *InvestmentsTransactionsGetRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *InvestmentsTransactionsGetRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *InvestmentsTransactionsGetRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetOptions

`func (o *InvestmentsTransactionsGetRequest) GetOptions() InvestmentsTransactionsGetRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *InvestmentsTransactionsGetRequest) GetOptionsOk() (*InvestmentsTransactionsGetRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *InvestmentsTransactionsGetRequest) SetOptions(v InvestmentsTransactionsGetRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *InvestmentsTransactionsGetRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


