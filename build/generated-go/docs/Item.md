# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | The Plaid Item ID. The &#x60;item_id&#x60; is always unique; linking the same account at the same institution twice will result in two Items with different &#x60;item_id&#x60; values. Like all Plaid identifiers, the &#x60;item_id&#x60; is case-sensitive. | 
**InstitutionId** | Pointer to **NullableString** | The Plaid Institution ID associated with the Item. Field is &#x60;null&#x60; for Items created via Same Day Micro-deposits. | [optional] 
**Webhook** | Pointer to **NullableString** | The URL registered to receive webhooks for the Item. | [optional] 
**Error** | Pointer to [**NullableError**](Error.md) |  | [optional] 
**AvailableProducts** | [**[]Products**](Products.md) | A list of products available for the Item that have not yet been accessed. | 
**BilledProducts** | [**[]Products**](Products.md) | A list of products that have been billed for the Item. Note - &#x60;billed_products&#x60; is populated in all environments but only requests in Production are billed.  | 
**ConsentExpirationTime** | Pointer to **NullableString** | The RFC 3339 timestamp after which the consent provided by the end user will expire. Upon consent expiration, the item will enter the &#x60;ITEM_LOGIN_REQUIRED&#x60; error state. To circumvent the &#x60;ITEM_LOGIN_REQUIRED&#x60; error and maintain continuous consent, the end user can reauthenticate via Link’s update mode in advance of the consent expiration time.  Note - This is only relevant for European institutions subject to PSD2 regulations mandating a 90-day consent window. For all other institutions, this field will be null.  | [optional] 
**UpdateType** | Pointer to **string** |  | [optional] 

## Methods

### NewItem

`func NewItem(itemId string, availableProducts []Products, billedProducts []Products, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *Item) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Item) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Item) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetInstitutionId

`func (o *Item) GetInstitutionId() string`

GetInstitutionId returns the InstitutionId field if non-nil, zero value otherwise.

### GetInstitutionIdOk

`func (o *Item) GetInstitutionIdOk() (*string, bool)`

GetInstitutionIdOk returns a tuple with the InstitutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstitutionId

`func (o *Item) SetInstitutionId(v string)`

SetInstitutionId sets InstitutionId field to given value.

### HasInstitutionId

`func (o *Item) HasInstitutionId() bool`

HasInstitutionId returns a boolean if a field has been set.

### SetInstitutionIdNil

`func (o *Item) SetInstitutionIdNil(b bool)`

 SetInstitutionIdNil sets the value for InstitutionId to be an explicit nil

### UnsetInstitutionId
`func (o *Item) UnsetInstitutionId()`

UnsetInstitutionId ensures that no value is present for InstitutionId, not even an explicit nil
### GetWebhook

`func (o *Item) GetWebhook() string`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *Item) GetWebhookOk() (*string, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *Item) SetWebhook(v string)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *Item) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.

### SetWebhookNil

`func (o *Item) SetWebhookNil(b bool)`

 SetWebhookNil sets the value for Webhook to be an explicit nil

### UnsetWebhook
`func (o *Item) UnsetWebhook()`

UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
### GetError

`func (o *Item) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *Item) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *Item) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *Item) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *Item) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *Item) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetAvailableProducts

`func (o *Item) GetAvailableProducts() []Products`

GetAvailableProducts returns the AvailableProducts field if non-nil, zero value otherwise.

### GetAvailableProductsOk

`func (o *Item) GetAvailableProductsOk() (*[]Products, bool)`

GetAvailableProductsOk returns a tuple with the AvailableProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableProducts

`func (o *Item) SetAvailableProducts(v []Products)`

SetAvailableProducts sets AvailableProducts field to given value.


### GetBilledProducts

`func (o *Item) GetBilledProducts() []Products`

GetBilledProducts returns the BilledProducts field if non-nil, zero value otherwise.

### GetBilledProductsOk

`func (o *Item) GetBilledProductsOk() (*[]Products, bool)`

GetBilledProductsOk returns a tuple with the BilledProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledProducts

`func (o *Item) SetBilledProducts(v []Products)`

SetBilledProducts sets BilledProducts field to given value.


### GetConsentExpirationTime

`func (o *Item) GetConsentExpirationTime() string`

GetConsentExpirationTime returns the ConsentExpirationTime field if non-nil, zero value otherwise.

### GetConsentExpirationTimeOk

`func (o *Item) GetConsentExpirationTimeOk() (*string, bool)`

GetConsentExpirationTimeOk returns a tuple with the ConsentExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentExpirationTime

`func (o *Item) SetConsentExpirationTime(v string)`

SetConsentExpirationTime sets ConsentExpirationTime field to given value.

### HasConsentExpirationTime

`func (o *Item) HasConsentExpirationTime() bool`

HasConsentExpirationTime returns a boolean if a field has been set.

### SetConsentExpirationTimeNil

`func (o *Item) SetConsentExpirationTimeNil(b bool)`

 SetConsentExpirationTimeNil sets the value for ConsentExpirationTime to be an explicit nil

### UnsetConsentExpirationTime
`func (o *Item) UnsetConsentExpirationTime()`

UnsetConsentExpirationTime ensures that no value is present for ConsentExpirationTime, not even an explicit nil
### GetUpdateType

`func (o *Item) GetUpdateType() string`

GetUpdateType returns the UpdateType field if non-nil, zero value otherwise.

### GetUpdateTypeOk

`func (o *Item) GetUpdateTypeOk() (*string, bool)`

GetUpdateTypeOk returns a tuple with the UpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateType

`func (o *Item) SetUpdateType(v string)`

SetUpdateType sets UpdateType field to given value.

### HasUpdateType

`func (o *Item) HasUpdateType() bool`

HasUpdateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


