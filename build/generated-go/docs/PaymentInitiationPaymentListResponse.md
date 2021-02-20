# PaymentInitiationPaymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payments** | [**[]PaymentInitiationPaymentGetResponse**](PaymentInitiationPaymentGetResponse.md) | An array of payments that have been created, associated with the given &#x60;client_id&#x60;. | 
**NextCursor** | **string** | The value that, when used as the optional &#x60;cursor&#x60; parameter to &#x60;/payment_initiation/payment/list&#x60;, will return the next unreturned payment as its first payment. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationPaymentListResponse

`func NewPaymentInitiationPaymentListResponse(payments []PaymentInitiationPaymentGetResponse, nextCursor string, requestId string, ) *PaymentInitiationPaymentListResponse`

NewPaymentInitiationPaymentListResponse instantiates a new PaymentInitiationPaymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentListResponseWithDefaults

`func NewPaymentInitiationPaymentListResponseWithDefaults() *PaymentInitiationPaymentListResponse`

NewPaymentInitiationPaymentListResponseWithDefaults instantiates a new PaymentInitiationPaymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayments

`func (o *PaymentInitiationPaymentListResponse) GetPayments() []PaymentInitiationPaymentGetResponse`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PaymentInitiationPaymentListResponse) GetPaymentsOk() (*[]PaymentInitiationPaymentGetResponse, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PaymentInitiationPaymentListResponse) SetPayments(v []PaymentInitiationPaymentGetResponse)`

SetPayments sets Payments field to given value.


### GetNextCursor

`func (o *PaymentInitiationPaymentListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *PaymentInitiationPaymentListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *PaymentInitiationPaymentListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetRequestId

`func (o *PaymentInitiationPaymentListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationPaymentListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationPaymentListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


