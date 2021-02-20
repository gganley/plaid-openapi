# PaymentStatusUpdateWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;PAYMENT_INITIATION&#x60; | 
**WebhookCode** | **string** | &#x60;PAYMENT_STATUS_UPDATE&#x60; | 
**PaymentId** | **string** | The &#x60;payment_id&#x60; for the payment being updated | 
**NewPaymentStatus** | **string** | The new status of the payment.  &#x60;PAYMENT_STATUS_INPUT_NEEDED&#x60;: This is the initial state of all payments. It indicates that the payment is waiting on user input to continue processing. A payment may re-enter this state later on if further input is needed.  &#x60;PAYMENT_STATUS_PROCESSING&#x60;: The payment is currently being processed. The payment will automatically exit this state when processing is complete.  &#x60;PAYMENT_STATUS_INITIATED&#x60;: The payment has been successfully initiated and is considered complete.  &#x60;PAYMENT_STATUS_COMPLETED&#x60;: Indicates that the standing order has been successfully established. This state is only used for standing orders.  &#x60;PAYMENT_STATUS_INSUFFICIENT_FUNDS&#x60;: The payment has failed due to insufficient funds.  &#x60;PAYMENT_STATUS_FAILED&#x60;: The payment has failed to be initiated. This error is retryable once the root cause is resolved.  &#x60;PAYMENT_STATUS_BLOCKED&#x60;: The payment has been blocked. This is a retryable error.  &#x60;PAYMENT_STATUS_UNKNOWN&#x60;: The payment status is unknown. | 
**OldPaymentStatus** | **string** | The previous status of the payment.  &#x60;PAYMENT_STATUS_INPUT_NEEDED&#x60;: This is the initial state of all payments. It indicates that the payment is waiting on user input to continue processing. A payment may re-enter this state later on if further input is needed.  &#x60;PAYMENT_STATUS_PROCESSING&#x60;: The payment is currently being processed. The payment will automatically exit this state when processing is complete.  &#x60;PAYMENT_STATUS_INITIATED&#x60;: The payment has been successfully initiated and is considered complete.  &#x60;PAYMENT_STATUS_COMPLETED&#x60;: Indicates that the standing order has been successfully established. This state is only used for standing orders.  &#x60;PAYMENT_STATUS_INSUFFICIENT_FUNDS&#x60;: The payment has failed due to insufficient funds.  &#x60;PAYMENT_STATUS_FAILED&#x60;: The payment has failed to be initiated. This error is retryable once the root cause is resolved.  &#x60;PAYMENT_STATUS_BLOCKED&#x60;: The payment has been blocked. This is a retryable error.  &#x60;PAYMENT_STATUS_UNKNOWN&#x60;: The payment status is unknown. | 
**Timestamp** | **string** | The timestamp of the update, in ISO 8601 format, e.g. &#x60;\&quot;2017-09-14T14:42:19.350Z\&quot;&#x60; | 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 

## Methods

### NewPaymentStatusUpdateWebhook

`func NewPaymentStatusUpdateWebhook(webhookType string, webhookCode string, paymentId string, newPaymentStatus string, oldPaymentStatus string, timestamp string, ) *PaymentStatusUpdateWebhook`

NewPaymentStatusUpdateWebhook instantiates a new PaymentStatusUpdateWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusUpdateWebhookWithDefaults

`func NewPaymentStatusUpdateWebhookWithDefaults() *PaymentStatusUpdateWebhook`

NewPaymentStatusUpdateWebhookWithDefaults instantiates a new PaymentStatusUpdateWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *PaymentStatusUpdateWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *PaymentStatusUpdateWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *PaymentStatusUpdateWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *PaymentStatusUpdateWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *PaymentStatusUpdateWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *PaymentStatusUpdateWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetPaymentId

`func (o *PaymentStatusUpdateWebhook) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentStatusUpdateWebhook) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentStatusUpdateWebhook) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetNewPaymentStatus

`func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatus() string`

GetNewPaymentStatus returns the NewPaymentStatus field if non-nil, zero value otherwise.

### GetNewPaymentStatusOk

`func (o *PaymentStatusUpdateWebhook) GetNewPaymentStatusOk() (*string, bool)`

GetNewPaymentStatusOk returns a tuple with the NewPaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPaymentStatus

`func (o *PaymentStatusUpdateWebhook) SetNewPaymentStatus(v string)`

SetNewPaymentStatus sets NewPaymentStatus field to given value.


### GetOldPaymentStatus

`func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatus() string`

GetOldPaymentStatus returns the OldPaymentStatus field if non-nil, zero value otherwise.

### GetOldPaymentStatusOk

`func (o *PaymentStatusUpdateWebhook) GetOldPaymentStatusOk() (*string, bool)`

GetOldPaymentStatusOk returns a tuple with the OldPaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPaymentStatus

`func (o *PaymentStatusUpdateWebhook) SetOldPaymentStatus(v string)`

SetOldPaymentStatus sets OldPaymentStatus field to given value.


### GetTimestamp

`func (o *PaymentStatusUpdateWebhook) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PaymentStatusUpdateWebhook) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PaymentStatusUpdateWebhook) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetError

`func (o *PaymentStatusUpdateWebhook) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PaymentStatusUpdateWebhook) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PaymentStatusUpdateWebhook) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *PaymentStatusUpdateWebhook) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *PaymentStatusUpdateWebhook) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *PaymentStatusUpdateWebhook) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


