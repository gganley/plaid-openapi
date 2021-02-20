# PaymentInitiationPaymentTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**PaymentId** | **string** | The &#x60;payment_id&#x60; returned from &#x60;/payment_initiation/payment/create&#x60;. | 

## Methods

### NewPaymentInitiationPaymentTokenCreateRequest

`func NewPaymentInitiationPaymentTokenCreateRequest(paymentId string, ) *PaymentInitiationPaymentTokenCreateRequest`

NewPaymentInitiationPaymentTokenCreateRequest instantiates a new PaymentInitiationPaymentTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentTokenCreateRequestWithDefaults

`func NewPaymentInitiationPaymentTokenCreateRequestWithDefaults() *PaymentInitiationPaymentTokenCreateRequest`

NewPaymentInitiationPaymentTokenCreateRequestWithDefaults instantiates a new PaymentInitiationPaymentTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationPaymentTokenCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationPaymentTokenCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationPaymentTokenCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationPaymentTokenCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPaymentTokenCreateRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPaymentTokenCreateRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


