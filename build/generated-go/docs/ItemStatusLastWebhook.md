# ItemStatusLastWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SentAt** | Pointer to **NullableString** | ISO 8601 timestamp of when the webhook was fired. | [optional] 
**CodeSent** | Pointer to **NullableString** | The last webhook code sent. | [optional] 

## Methods

### NewItemStatusLastWebhook

`func NewItemStatusLastWebhook() *ItemStatusLastWebhook`

NewItemStatusLastWebhook instantiates a new ItemStatusLastWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemStatusLastWebhookWithDefaults

`func NewItemStatusLastWebhookWithDefaults() *ItemStatusLastWebhook`

NewItemStatusLastWebhookWithDefaults instantiates a new ItemStatusLastWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentAt

`func (o *ItemStatusLastWebhook) GetSentAt() string`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *ItemStatusLastWebhook) GetSentAtOk() (*string, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *ItemStatusLastWebhook) SetSentAt(v string)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *ItemStatusLastWebhook) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *ItemStatusLastWebhook) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *ItemStatusLastWebhook) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetCodeSent

`func (o *ItemStatusLastWebhook) GetCodeSent() string`

GetCodeSent returns the CodeSent field if non-nil, zero value otherwise.

### GetCodeSentOk

`func (o *ItemStatusLastWebhook) GetCodeSentOk() (*string, bool)`

GetCodeSentOk returns a tuple with the CodeSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSent

`func (o *ItemStatusLastWebhook) SetCodeSent(v string)`

SetCodeSent sets CodeSent field to given value.

### HasCodeSent

`func (o *ItemStatusLastWebhook) HasCodeSent() bool`

HasCodeSent returns a boolean if a field has been set.

### SetCodeSentNil

`func (o *ItemStatusLastWebhook) SetCodeSentNil(b bool)`

 SetCodeSentNil sets the value for CodeSent to be an explicit nil

### UnsetCodeSent
`func (o *ItemStatusLastWebhook) UnsetCodeSent()`

UnsetCodeSent ensures that no value is present for CodeSent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


