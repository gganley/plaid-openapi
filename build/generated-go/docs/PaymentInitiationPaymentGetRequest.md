# PaymentInitiationPaymentGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**PaymentId** | **string** | The &#x60;payment_id&#x60; returned from &#x60;/payment_initiation/payment/create&#x60;. | 

## Methods

### NewPaymentInitiationPaymentGetRequest

`func NewPaymentInitiationPaymentGetRequest(paymentId string, ) *PaymentInitiationPaymentGetRequest`

NewPaymentInitiationPaymentGetRequest instantiates a new PaymentInitiationPaymentGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentGetRequestWithDefaults

`func NewPaymentInitiationPaymentGetRequestWithDefaults() *PaymentInitiationPaymentGetRequest`

NewPaymentInitiationPaymentGetRequestWithDefaults instantiates a new PaymentInitiationPaymentGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationPaymentGetRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationPaymentGetRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationPaymentGetRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationPaymentGetRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationPaymentGetRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationPaymentGetRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationPaymentGetRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationPaymentGetRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPaymentId

`func (o *PaymentInitiationPaymentGetRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPaymentGetRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPaymentGetRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


