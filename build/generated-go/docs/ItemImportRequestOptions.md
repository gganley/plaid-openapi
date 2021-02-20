# ItemImportRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhook** | Pointer to **string** | Specifies a webhook URL to associate with an Item. Plaid fires a webhook if credentials fail.  | [optional] 

## Methods

### NewItemImportRequestOptions

`func NewItemImportRequestOptions() *ItemImportRequestOptions`

NewItemImportRequestOptions instantiates a new ItemImportRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemImportRequestOptionsWithDefaults

`func NewItemImportRequestOptionsWithDefaults() *ItemImportRequestOptions`

NewItemImportRequestOptionsWithDefaults instantiates a new ItemImportRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhook

`func (o *ItemImportRequestOptions) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *ItemImportRequestOptions) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *ItemImportRequestOptions) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *ItemImportRequestOptions) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


