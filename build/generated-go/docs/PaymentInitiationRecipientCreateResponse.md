# PaymentInitiationRecipientCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientId** | **string** | A unique ID identifying the recipient | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewPaymentInitiationRecipientCreateResponse

`func NewPaymentInitiationRecipientCreateResponse(recipientId string, requestId string, ) *PaymentInitiationRecipientCreateResponse`

NewPaymentInitiationRecipientCreateResponse instantiates a new PaymentInitiationRecipientCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientCreateResponseWithDefaults

`func NewPaymentInitiationRecipientCreateResponseWithDefaults() *PaymentInitiationRecipientCreateResponse`

NewPaymentInitiationRecipientCreateResponseWithDefaults instantiates a new PaymentInitiationRecipientCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientId

`func (o *PaymentInitiationRecipientCreateResponse) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationRecipientCreateResponse) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationRecipientCreateResponse) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetRequestId

`func (o *PaymentInitiationRecipientCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *PaymentInitiationRecipientCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *PaymentInitiationRecipientCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


