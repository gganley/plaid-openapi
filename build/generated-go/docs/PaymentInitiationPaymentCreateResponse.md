# PaymentInitiationPaymentCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | A unique ID identifying the payment | 
**Status** | **string** | For a payment returned by this endpoint, there is only one possible value:  &#x60;PAYMENT_STATUS_INPUT_NEEDED&#x60;: The initial phase of the payment | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationPaymentCreateResponse

`func NewPaymentInitiationPaymentCreateResponse(paymentId string, status string, requestId string, ) *PaymentInitiationPaymentCreateResponse`

NewPaymentInitiationPaymentCreateResponse instantiates a new PaymentInitiationPaymentCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentCreateResponseWithDefaults

`func NewPaymentInitiationPaymentCreateResponseWithDefaults() *PaymentInitiationPaymentCreateResponse`

NewPaymentInitiationPaymentCreateResponseWithDefaults instantiates a new PaymentInitiationPaymentCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentInitiationPaymentCreateResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPaymentCreateResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPaymentCreateResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetStatus

`func (o *PaymentInitiationPaymentCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationPaymentCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationPaymentCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestId

`func (o *PaymentInitiationPaymentCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


