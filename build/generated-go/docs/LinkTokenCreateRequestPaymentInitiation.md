# LinkTokenCreateRequestPaymentInitiation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The &#x60;payment_id&#x60; provided by the &#x60;/payment_initiation/payment/create&#x60; endpoint. | 

## Methods

### NewLinkTokenCreateRequestPaymentInitiation

`func NewLinkTokenCreateRequestPaymentInitiation(paymentId string, ) *LinkTokenCreateRequestPaymentInitiation`

NewLinkTokenCreateRequestPaymentInitiation instantiates a new LinkTokenCreateRequestPaymentInitiation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkTokenCreateRequestPaymentInitiationWithDefaults

`func NewLinkTokenCreateRequestPaymentInitiationWithDefaults() *LinkTokenCreateRequestPaymentInitiation`

NewLinkTokenCreateRequestPaymentInitiationWithDefaults instantiates a new LinkTokenCreateRequestPaymentInitiation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *LinkTokenCreateRequestPaymentInitiation) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *LinkTokenCreateRequestPaymentInitiation) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *LinkTokenCreateRequestPaymentInitiation) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


