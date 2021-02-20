# ItemWebhookUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**Item**](Item.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewItemWebhookUpdateResponse

`func NewItemWebhookUpdateResponse(item Item, requestId string, ) *ItemWebhookUpdateResponse`

NewItemWebhookUpdateResponse instantiates a new ItemWebhookUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWebhookUpdateResponseWithDefaults

`func NewItemWebhookUpdateResponseWithDefaults() *ItemWebhookUpdateResponse`

NewItemWebhookUpdateResponseWithDefaults instantiates a new ItemWebhookUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *ItemWebhookUpdateResponse) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ItemWebhookUpdateResponse) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ItemWebhookUpdateResponse) SetItem(v Item)`

SetItem sets Item field to given value.


### GetRequestId

`func (o *ItemWebhookUpdateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemWebhookUpdateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemWebhookUpdateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


