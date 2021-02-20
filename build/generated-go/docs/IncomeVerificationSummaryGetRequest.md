# IncomeVerificationSummaryGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**IncomeVerificationId** | **string** | The ID of the verification. | 

## Methods

### NewIncomeVerificationSummaryGetRequest

`func NewIncomeVerificationSummaryGetRequest(incomeVerificationId string, ) *IncomeVerificationSummaryGetRequest`

NewIncomeVerificationSummaryGetRequest instantiates a new IncomeVerificationSummaryGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationSummaryGetRequestWithDefaults

`func NewIncomeVerificationSummaryGetRequestWithDefaults() *IncomeVerificationSummaryGetRequest`

NewIncomeVerificationSummaryGetRequestWithDefaults instantiates a new IncomeVerificationSummaryGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationSummaryGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationSummaryGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationSummaryGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationSummaryGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationSummaryGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationSummaryGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationSummaryGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationSummaryGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIncomeVerificationId

`func (o *IncomeVerificationSummaryGetRequest) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationSummaryGetRequest) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationSummaryGetRequest) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


