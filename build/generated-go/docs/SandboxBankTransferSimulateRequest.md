# SandboxBankTransferSimulateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**BankTransferId** | **string** | Plaid’s unique identifier for a bank transfer. | 
**EventType** | **string** | The asynchronous event to be simulated. May be: &#x60;posted&#x60;, &#x60;failed&#x60;, or &#x60;reversed&#x60;.  An error will be returned if the event type is incompatible with the current transfer status. Compatible status --&gt; event type transitions include:  &#x60;pending&#x60; --&gt; &#x60;failed&#x60;  &#x60;pending&#x60; --&gt; &#x60;posted&#x60;  &#x60;posted&#x60; --&gt; &#x60;reversed&#x60;  | 
**FailureReason** | Pointer to [**NullableBankTransferFailure**](BankTransferFailure.md) |  | [optional] 

## Methods

### NewSandboxBankTransferSimulateRequest

`func NewSandboxBankTransferSimulateRequest(bankTransferId string, eventType string, ) *SandboxBankTransferSimulateRequest`

NewSandboxBankTransferSimulateRequest instantiates a new SandboxBankTransferSimulateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxBankTransferSimulateRequestWithDefaults

`func NewSandboxBankTransferSimulateRequestWithDefaults() *SandboxBankTransferSimulateRequest`

NewSandboxBankTransferSimulateRequestWithDefaults instantiates a new SandboxBankTransferSimulateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *SandboxBankTransferSimulateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SandboxBankTransferSimulateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SandboxBankTransferSimulateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SandboxBankTransferSimulateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *SandboxBankTransferSimulateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SandboxBankTransferSimulateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SandboxBankTransferSimulateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *SandboxBankTransferSimulateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetBankTransferId

`func (o *SandboxBankTransferSimulateRequest) GetBankTransferId() string`

GetBankTransferId returns the BankTransferId field if non-nil, zero value otherwise.

### GetBankTransferIdOk

`func (o *SandboxBankTransferSimulateRequest) GetBankTransferIdOk() (*string, bool)`

GetBankTransferIdOk returns a tuple with the BankTransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransferId

`func (o *SandboxBankTransferSimulateRequest) SetBankTransferId(v string)`

SetBankTransferId sets BankTransferId field to given value.


### GetEventType

`func (o *SandboxBankTransferSimulateRequest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *SandboxBankTransferSimulateRequest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *SandboxBankTransferSimulateRequest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFailureReason

`func (o *SandboxBankTransferSimulateRequest) GetFailureReason() BankTransferFailure`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *SandboxBankTransferSimulateRequest) GetFailureReasonOk() (*BankTransferFailure, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *SandboxBankTransferSimulateRequest) SetFailureReason(v BankTransferFailure)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *SandboxBankTransferSimulateRequest) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *SandboxBankTransferSimulateRequest) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *SandboxBankTransferSimulateRequest) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


