# PaymentInitiationRecipientCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Name** | **string** | The name of the recipient | 
**Iban** | Pointer to **string** | The International Bank Account Number (IBAN) for the recipient. If BACS data is not provided, an IBAN is required. | [optional] 
**Bacs** | Pointer to [**NullableRecipientBACS**](RecipientBACS.md) |  | [optional] 
**Address** | Pointer to [**NullablePaymentInitiationAddress**](PaymentInitiationAddress.md) |  | [optional] 

## Methods

### NewPaymentInitiationRecipientCreateRequest

`func NewPaymentInitiationRecipientCreateRequest(name string, ) *PaymentInitiationRecipientCreateRequest`

NewPaymentInitiationRecipientCreateRequest instantiates a new PaymentInitiationRecipientCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationRecipientCreateRequestWithDefaults

`func NewPaymentInitiationRecipientCreateRequestWithDefaults() *PaymentInitiationRecipientCreateRequest`

NewPaymentInitiationRecipientCreateRequestWithDefaults instantiates a new PaymentInitiationRecipientCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationRecipientCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationRecipientCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationRecipientCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationRecipientCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationRecipientCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationRecipientCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationRecipientCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationRecipientCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetName

`func (o *PaymentInitiationRecipientCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentInitiationRecipientCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentInitiationRecipientCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetIban

`func (o *PaymentInitiationRecipientCreateRequest) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PaymentInitiationRecipientCreateRequest) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PaymentInitiationRecipientCreateRequest) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PaymentInitiationRecipientCreateRequest) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetBacs

`func (o *PaymentInitiationRecipientCreateRequest) GetBacs() RecipientBACS`

GetBacs returns the Bacs field if non-nil, zero value otherwise.

### GetBacsOk

`func (o *PaymentInitiationRecipientCreateRequest) GetBacsOk() (*RecipientBACS, bool)`

GetBacsOk returns a tuple with the Bacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacs

`func (o *PaymentInitiationRecipientCreateRequest) SetBacs(v RecipientBACS)`

SetBacs sets Bacs field to given value.

### HasBacs

`func (o *PaymentInitiationRecipientCreateRequest) HasBacs() bool`

HasBacs returns a boolean if a field has been set.

### SetBacsNil

`func (o *PaymentInitiationRecipientCreateRequest) SetBacsNil(b bool)`

 SetBacsNil sets the value for Bacs to be an explicit nil

### UnsetBacs
`func (o *PaymentInitiationRecipientCreateRequest) UnsetBacs()`

UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
### GetAddress

`func (o *PaymentInitiationRecipientCreateRequest) GetAddress() PaymentInitiationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaymentInitiationRecipientCreateRequest) GetAddressOk() (*PaymentInitiationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaymentInitiationRecipientCreateRequest) SetAddress(v PaymentInitiationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaymentInitiationRecipientCreateRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *PaymentInitiationRecipientCreateRequest) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *PaymentInitiationRecipientCreateRequest) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


