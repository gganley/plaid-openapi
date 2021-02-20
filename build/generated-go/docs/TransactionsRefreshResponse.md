# TransactionsRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransactionsRefreshResponse

`func NewTransactionsRefreshResponse(requestId string, ) *TransactionsRefreshResponse`

NewTransactionsRefreshResponse instantiates a new TransactionsRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsRefreshResponseWithDefaults

`func NewTransactionsRefreshResponseWithDefaults() *TransactionsRefreshResponse`

NewTransactionsRefreshResponseWithDefaults instantiates a new TransactionsRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *TransactionsRefreshResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionsRefreshResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionsRefreshResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


