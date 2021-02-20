# PaymentInitiationRecipientListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 

## Methods

### NewPaymentInitiationRecipientListRequest

`func NewPaymentInitiationRecipientListRequest() *PaymentInitiationRecipientListRequest`

NewPaymentInitiationRecipientListRequest instantiates a new PaymentInitiationRecipientListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientListRequestWithDefaults

`func NewPaymentInitiationRecipientListRequestWithDefaults() *PaymentInitiationRecipientListRequest`

NewPaymentInitiationRecipientListRequestWithDefaults instantiates a new PaymentInitiationRecipientListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationRecipientListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationRecipientListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationRecipientListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationRecipientListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationRecipientListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationRecipientListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationRecipientListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationRecipientListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


