# IncomeVerificationPaystubGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**IncomeVerificationId** | **string** | The ID of the verification for which to get paystub information. | 

## Methods

### NewIncomeVerificationPaystubGetRequest

`func NewIncomeVerificationPaystubGetRequest(incomeVerificationId string, ) *IncomeVerificationPaystubGetRequest`

NewIncomeVerificationPaystubGetRequest instantiates a new IncomeVerificationPaystubGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPaystubGetRequestWithDefaults

`func NewIncomeVerificationPaystubGetRequestWithDefaults() *IncomeVerificationPaystubGetRequest`

NewIncomeVerificationPaystubGetRequestWithDefaults instantiates a new IncomeVerificationPaystubGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationPaystubGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationPaystubGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationPaystubGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationPaystubGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationPaystubGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationPaystubGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationPaystubGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationPaystubGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIncomeVerificationId

`func (o *IncomeVerificationPaystubGetRequest) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationPaystubGetRequest) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationPaystubGetRequest) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


