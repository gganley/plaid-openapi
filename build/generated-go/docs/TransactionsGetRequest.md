# TransactionsGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Options** | Pointer to [**TransactionsGetRequestOptions**](TransactionsGetRequestOptions.md) |  | [optional] 
**AccessToken** | **string** | The access token associated with the Item data is being requested for. | 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**StartDate** | **string** | The earliest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 
**EndDate** | **string** | The latest date for which data should be returned. Dates should be formatted as YYYY-MM-DD. | 

## Methods

### NewTransactionsGetRequest

`func NewTransactionsGetRequest(accessToken string, startDate string, endDate string, ) *TransactionsGetRequest`

NewTransactionsGetRequest instantiates a new TransactionsGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsGetRequestWithDefaults

`func NewTransactionsGetRequestWithDefaults() *TransactionsGetRequest`

NewTransactionsGetRequestWithDefaults instantiates a new TransactionsGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransactionsGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransactionsGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransactionsGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransactionsGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOptions

`func (o *TransactionsGetRequest) GetOptions() TransactionsGetRequestOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *TransactionsGetRequest) GetOptionsOk() (*TransactionsGetRequestOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *TransactionsGetRequest) SetOptions(v TransactionsGetRequestOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *TransactionsGetRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransactionsGetRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransactionsGetRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransactionsGetRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetSecret

`func (o *TransactionsGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransactionsGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransactionsGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransactionsGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStartDate

`func (o *TransactionsGetRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TransactionsGetRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TransactionsGetRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *TransactionsGetRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TransactionsGetRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TransactionsGetRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


