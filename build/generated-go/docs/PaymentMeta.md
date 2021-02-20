# PaymentMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceNumber** | Pointer to **NullableString** | The transaction reference number supplied by the financial institution. | [optional] 
**PpdId** | Pointer to **NullableString** | The ACH PPD ID for the payer. | [optional] 
**Payee** | Pointer to **NullableString** | For transfers, the party that is receiving the transaction. | [optional] 
**ByOrderOf** | Pointer to **NullableString** | The party initiating a wire transfer. Will be &#x60;null&#x60; if the transaction is not a wire transfer. | [optional] 
**Payer** | Pointer to **NullableString** | For transfers, the party that is paying the transaction. | [optional] 
**PaymentMethod** | Pointer to **NullableString** | The type of transfer, e.g. &#39;ACH&#39; | [optional] 
**PaymentProcessor** | Pointer to **NullableString** | The name of the payment processor | [optional] 
**Reason** | Pointer to **NullableString** | The payer-supplied description of the transfer. | [optional] 

## Methods

### NewPaymentMeta

`func NewPaymentMeta() *PaymentMeta`

NewPaymentMeta instantiates a new PaymentMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMetaWithDefaults

`func NewPaymentMetaWithDefaults() *PaymentMeta`

NewPaymentMetaWithDefaults instantiates a new PaymentMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceNumber

`func (o *PaymentMeta) GetReferenceNumber() string`

GetReferenceNumber returns the ReferenceNumber field if non-nil, zero value otherwise.

### GetReferenceNumberOk

`func (o *PaymentMeta) GetReferenceNumberOk() (*string, bool)`

GetReferenceNumberOk returns a tuple with the ReferenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNumber

`func (o *PaymentMeta) SetReferenceNumber(v string)`

SetReferenceNumber sets ReferenceNumber field to given value.

### HasReferenceNumber

`func (o *PaymentMeta) HasReferenceNumber() bool`

HasReferenceNumber returns a boolean if a field has been set.

### SetReferenceNumberNil

`func (o *PaymentMeta) SetReferenceNumberNil(b bool)`

 SetReferenceNumberNil sets the value for ReferenceNumber to be an explicit nil

### UnsetReferenceNumber
`func (o *PaymentMeta) UnsetReferenceNumber()`

UnsetReferenceNumber ensures that no value is present for ReferenceNumber, not even an explicit nil
### GetPpdId

`func (o *PaymentMeta) GetPpdId() string`

GetPpdId returns the PpdId field if non-nil, zero value otherwise.

### GetPpdIdOk

`func (o *PaymentMeta) GetPpdIdOk() (*string, bool)`

GetPpdIdOk returns a tuple with the PpdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpdId

`func (o *PaymentMeta) SetPpdId(v string)`

SetPpdId sets PpdId field to given value.

### HasPpdId

`func (o *PaymentMeta) HasPpdId() bool`

HasPpdId returns a boolean if a field has been set.

### SetPpdIdNil

`func (o *PaymentMeta) SetPpdIdNil(b bool)`

 SetPpdIdNil sets the value for PpdId to be an explicit nil

### UnsetPpdId
`func (o *PaymentMeta) UnsetPpdId()`

UnsetPpdId ensures that no value is present for PpdId, not even an explicit nil
### GetPayee

`func (o *PaymentMeta) GetPayee() string`

GetPayee returns the Payee field if non-nil, zero value otherwise.

### GetPayeeOk

`func (o *PaymentMeta) GetPayeeOk() (*string, bool)`

GetPayeeOk returns a tuple with the Payee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayee

`func (o *PaymentMeta) SetPayee(v string)`

SetPayee sets Payee field to given value.

### HasPayee

`func (o *PaymentMeta) HasPayee() bool`

HasPayee returns a boolean if a field has been set.

### SetPayeeNil

`func (o *PaymentMeta) SetPayeeNil(b bool)`

 SetPayeeNil sets the value for Payee to be an explicit nil

### UnsetPayee
`func (o *PaymentMeta) UnsetPayee()`

UnsetPayee ensures that no value is present for Payee, not even an explicit nil
### GetByOrderOf

`func (o *PaymentMeta) GetByOrderOf() string`

GetByOrderOf returns the ByOrderOf field if non-nil, zero value otherwise.

### GetByOrderOfOk

`func (o *PaymentMeta) GetByOrderOfOk() (*string, bool)`

GetByOrderOfOk returns a tuple with the ByOrderOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByOrderOf

`func (o *PaymentMeta) SetByOrderOf(v string)`

SetByOrderOf sets ByOrderOf field to given value.

### HasByOrderOf

`func (o *PaymentMeta) HasByOrderOf() bool`

HasByOrderOf returns a boolean if a field has been set.

### SetByOrderOfNil

`func (o *PaymentMeta) SetByOrderOfNil(b bool)`

 SetByOrderOfNil sets the value for ByOrderOf to be an explicit nil

### UnsetByOrderOf
`func (o *PaymentMeta) UnsetByOrderOf()`

UnsetByOrderOf ensures that no value is present for ByOrderOf, not even an explicit nil
### GetPayer

`func (o *PaymentMeta) GetPayer() string`

GetPayer returns the Payer field if non-nil, zero value otherwise.

### GetPayerOk

`func (o *PaymentMeta) GetPayerOk() (*string, bool)`

GetPayerOk returns a tuple with the Payer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayer

`func (o *PaymentMeta) SetPayer(v string)`

SetPayer sets Payer field to given value.

### HasPayer

`func (o *PaymentMeta) HasPayer() bool`

HasPayer returns a boolean if a field has been set.

### SetPayerNil

`func (o *PaymentMeta) SetPayerNil(b bool)`

 SetPayerNil sets the value for Payer to be an explicit nil

### UnsetPayer
`func (o *PaymentMeta) UnsetPayer()`

UnsetPayer ensures that no value is present for Payer, not even an explicit nil
### GetPaymentMethod

`func (o *PaymentMeta) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *PaymentMeta) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *PaymentMeta) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *PaymentMeta) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### SetPaymentMethodNil

`func (o *PaymentMeta) SetPaymentMethodNil(b bool)`

 SetPaymentMethodNil sets the value for PaymentMethod to be an explicit nil

### UnsetPaymentMethod
`func (o *PaymentMeta) UnsetPaymentMethod()`

UnsetPaymentMethod ensures that no value is present for PaymentMethod, not even an explicit nil
### GetPaymentProcessor

`func (o *PaymentMeta) GetPaymentProcessor() string`

GetPaymentProcessor returns the PaymentProcessor field if non-nil, zero value otherwise.

### GetPaymentProcessorOk

`func (o *PaymentMeta) GetPaymentProcessorOk() (*string, bool)`

GetPaymentProcessorOk returns a tuple with the PaymentProcessor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentProcessor

`func (o *PaymentMeta) SetPaymentProcessor(v string)`

SetPaymentProcessor sets PaymentProcessor field to given value.

### HasPaymentProcessor

`func (o *PaymentMeta) HasPaymentProcessor() bool`

HasPaymentProcessor returns a boolean if a field has been set.

### SetPaymentProcessorNil

`func (o *PaymentMeta) SetPaymentProcessorNil(b bool)`

 SetPaymentProcessorNil sets the value for PaymentProcessor to be an explicit nil

### UnsetPaymentProcessor
`func (o *PaymentMeta) UnsetPaymentProcessor()`

UnsetPaymentProcessor ensures that no value is present for PaymentProcessor, not even an explicit nil
### GetReason

`func (o *PaymentMeta) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PaymentMeta) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PaymentMeta) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PaymentMeta) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *PaymentMeta) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PaymentMeta) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


