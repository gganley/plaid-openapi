# PaymentInitiationRecipientListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | [**[]PaymentInitiationRecipient**](PaymentInitiationRecipient.md) | An array of payment recipients created for Payment Initiation | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationRecipientListResponse

`func NewPaymentInitiationRecipientListResponse(recipients []PaymentInitiationRecipient, requestId string, ) *PaymentInitiationRecipientListResponse`

NewPaymentInitiationRecipientListResponse instantiates a new PaymentInitiationRecipientListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientListResponseWithDefaults

`func NewPaymentInitiationRecipientListResponseWithDefaults() *PaymentInitiationRecipientListResponse`

NewPaymentInitiationRecipientListResponseWithDefaults instantiates a new PaymentInitiationRecipientListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *PaymentInitiationRecipientListResponse) GetRecipients() []PaymentInitiationRecipient`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *PaymentInitiationRecipientListResponse) GetRecipientsOk() (*[]PaymentInitiationRecipient, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *PaymentInitiationRecipientListResponse) SetRecipients(v []PaymentInitiationRecipient)`

SetRecipients sets Recipients field to given value.


### GetRequestId

`func (o *PaymentInitiationRecipientListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationRecipientListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationRecipientListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


