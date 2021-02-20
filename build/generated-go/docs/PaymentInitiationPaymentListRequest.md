# PaymentInitiationPaymentListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Count** | Pointer to **int32** | The maximum number of payments to return. If &#x60;count&#x60; is not specified, a maximum of 10 payments will be returned, beginning with the most recent payment before the cursor (if specified). | [optional] [default to 10]
**Cursor** | Pointer to **string** | A string in RFC 3339 format (i.e. \&quot;2019-12-06T22:35:49Z\&quot;). Only payments created before the cursor will be returned. | [optional] 

## Methods

### NewPaymentInitiationPaymentListRequest

`func NewPaymentInitiationPaymentListRequest() *PaymentInitiationPaymentListRequest`

NewPaymentInitiationPaymentListRequest instantiates a new PaymentInitiationPaymentListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentListRequestWithDefaults

`func NewPaymentInitiationPaymentListRequestWithDefaults() *PaymentInitiationPaymentListRequest`

NewPaymentInitiationPaymentListRequestWithDefaults instantiates a new PaymentInitiationPaymentListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationPaymentListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationPaymentListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationPaymentListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationPaymentListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationPaymentListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationPaymentListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationPaymentListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationPaymentListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCount

`func (o *PaymentInitiationPaymentListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentInitiationPaymentListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentInitiationPaymentListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaymentInitiationPaymentListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCursor

`func (o *PaymentInitiationPaymentListRequest) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PaymentInitiationPaymentListRequest) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PaymentInitiationPaymentListRequest) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PaymentInitiationPaymentListRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


