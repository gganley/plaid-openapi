# PaymentInitiationPaymentTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentToken** | **string** | A &#x60;payment_token&#x60; that can be provided to Link initialization to enter the payment initiation flow | 
**PaymentTokenExpirationTime** | **string** | The date and time at which the token will expire, in ISO 8601 format. A &#x60;payment_token&#x60; expires after 15 minutes. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationPaymentTokenCreateResponse

`func NewPaymentInitiationPaymentTokenCreateResponse(paymentToken string, paymentTokenExpirationTime string, requestId string, ) *PaymentInitiationPaymentTokenCreateResponse`

NewPaymentInitiationPaymentTokenCreateResponse instantiates a new PaymentInitiationPaymentTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentTokenCreateResponseWithDefaults

`func NewPaymentInitiationPaymentTokenCreateResponseWithDefaults() *PaymentInitiationPaymentTokenCreateResponse`

NewPaymentInitiationPaymentTokenCreateResponseWithDefaults instantiates a new PaymentInitiationPaymentTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentToken

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentToken() string`

GetPaymentToken returns the PaymentToken field if non-nil, zero value otherwise.

### GetPaymentTokenOk

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenOk() (*string, bool)`

GetPaymentTokenOk returns a tuple with the PaymentToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentToken

`func (o *PaymentInitiationPaymentTokenCreateResponse) SetPaymentToken(v string)`

SetPaymentToken sets PaymentToken field to given value.


### GetPaymentTokenExpirationTime

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenExpirationTime() string`

GetPaymentTokenExpirationTime returns the PaymentTokenExpirationTime field if non-nil, zero value otherwise.

### GetPaymentTokenExpirationTimeOk

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetPaymentTokenExpirationTimeOk() (*string, bool)`

GetPaymentTokenExpirationTimeOk returns a tuple with the PaymentTokenExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTokenExpirationTime

`func (o *PaymentInitiationPaymentTokenCreateResponse) SetPaymentTokenExpirationTime(v string)`

SetPaymentTokenExpirationTime sets PaymentTokenExpirationTime field to given value.


### GetRequestId

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


