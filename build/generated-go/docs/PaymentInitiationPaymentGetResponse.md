# PaymentInitiationPaymentGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The ID of the payment. Like all Plaid identifiers, the &#x60;payment_id&#x60; is case sensitive. | 
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | [optional] 
**Amount** | [**PaymentAmount**](PaymentAmount.md) |  | 
**Status** | **string** | The status of the payment.  &#x60;PAYMENT_STATUS_INPUT_NEEDED&#x60;: This is the initial state of all payments. It indicates that the payment is waiting on user input to continue processing. A payment may re-enter this state later on if further input is needed.  &#x60;PAYMENT_STATUS_PROCESSING&#x60;: The payment is currently being processed. The payment will automatically exit this state when processing is complete.  &#x60;PAYMENT_STATUS_INITIATED&#x60;: The payment has been successfully initiated and is considered complete.  &#x60;PAYMENT_STATUS_COMPLETED&#x60;: Indicates that the standing order has been successfully established. This state is only used for standing orders.  &#x60;PAYMENT_STATUS_INSUFFICIENT_FUNDS&#x60;: The payment has failed due to insufficient funds.  &#x60;PAYMENT_STATUS_FAILED&#x60;: The payment has failed to be initiated. This error is retryable once the root cause is resolved.  &#x60;PAYMENT_STATUS_BLOCKED&#x60;: The payment has been blocked. This is a retryable error.  &#x60;PAYMENT_STATUS_UNKNOWN&#x60;: The payment status is unknown. | 
**RecipientId** | **string** | The ID of the recipient | 
**Reference** | **string** | A reference for the payment. | 
**LastStatusUpdate** | **string** | The date and time of the last time the &#x60;status&#x60; was updated, in IS0 8601 format | 
**Schedule** | Pointer to [**NullableExternalPaymentSchedule**](ExternalPaymentSchedule.md) |  | [optional] 
**AdjustedReference** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaymentInitiationPaymentGetResponse

`func NewPaymentInitiationPaymentGetResponse(paymentId string, amount PaymentAmount, status string, recipientId string, reference string, lastStatusUpdate string, ) *PaymentInitiationPaymentGetResponse`

NewPaymentInitiationPaymentGetResponse instantiates a new PaymentInitiationPaymentGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentGetResponseWithDefaults

`func NewPaymentInitiationPaymentGetResponseWithDefaults() *PaymentInitiationPaymentGetResponse`

NewPaymentInitiationPaymentGetResponseWithDefaults instantiates a new PaymentInitiationPaymentGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *PaymentInitiationPaymentGetResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentInitiationPaymentGetResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetRequestId

`func (o *PaymentInitiationPaymentGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *PaymentInitiationPaymentGetResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentInitiationPaymentGetResponse) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitiationPaymentGetResponse) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitiationPaymentGetResponse) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.


### GetStatus

`func (o *PaymentInitiationPaymentGetResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentInitiationPaymentGetResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentInitiationPaymentGetResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRecipientId

`func (o *PaymentInitiationPaymentGetResponse) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationPaymentGetResponse) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationPaymentGetResponse) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetReference

`func (o *PaymentInitiationPaymentGetResponse) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInitiationPaymentGetResponse) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInitiationPaymentGetResponse) SetReference(v string)`

SetReference sets Reference field to given value.


### GetLastStatusUpdate

`func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdate() string`

GetLastStatusUpdate returns the LastStatusUpdate field if non-nil, zero value otherwise.

### GetLastStatusUpdateOk

`func (o *PaymentInitiationPaymentGetResponse) GetLastStatusUpdateOk() (*string, bool)`

GetLastStatusUpdateOk returns a tuple with the LastStatusUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusUpdate

`func (o *PaymentInitiationPaymentGetResponse) SetLastStatusUpdate(v string)`

SetLastStatusUpdate sets LastStatusUpdate field to given value.


### GetSchedule

`func (o *PaymentInitiationPaymentGetResponse) GetSchedule() ExternalPaymentSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentInitiationPaymentGetResponse) GetScheduleOk() (*ExternalPaymentSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentInitiationPaymentGetResponse) SetSchedule(v ExternalPaymentSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PaymentInitiationPaymentGetResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *PaymentInitiationPaymentGetResponse) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *PaymentInitiationPaymentGetResponse) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReference() string`

GetAdjustedReference returns the AdjustedReference field if non-nil, zero value otherwise.

### GetAdjustedReferenceOk

`func (o *PaymentInitiationPaymentGetResponse) GetAdjustedReferenceOk() (*string, bool)`

GetAdjustedReferenceOk returns a tuple with the AdjustedReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReference(v string)`

SetAdjustedReference sets AdjustedReference field to given value.

### HasAdjustedReference

`func (o *PaymentInitiationPaymentGetResponse) HasAdjustedReference() bool`

HasAdjustedReference returns a boolean if a field has been set.

### SetAdjustedReferenceNil

`func (o *PaymentInitiationPaymentGetResponse) SetAdjustedReferenceNil(b bool)`

 SetAdjustedReferenceNil sets the value for AdjustedReference to be an explicit nil

### UnsetAdjustedReference
`func (o *PaymentInitiationPaymentGetResponse) UnsetAdjustedReference()`

UnsetAdjustedReference ensures that no value is present for AdjustedReference, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


