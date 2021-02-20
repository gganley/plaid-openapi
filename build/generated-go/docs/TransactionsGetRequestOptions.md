# TransactionsGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIds** | Pointer to **[]string** | A list of &#x60;account_ids&#x60; to retrieve for the Item  Note: An error will be returned if a provided &#x60;account_id&#x60; is not associated with the Item. | [optional] 
**Count** | Pointer to **int32** | The number of transactions to fetch. | [optional] [default to 100]
**Offset** | Pointer to **int32** | The number of transactions to skip. The default value is 0. | [optional] [default to 0]

## Methods

### NewTransactionsGetRequestOptions

`func NewTransactionsGetRequestOptions() *TransactionsGetRequestOptions`

NewTransactionsGetRequestOptions instantiates a new TransactionsGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsGetRequestOptionsWithDefaults

`func NewTransactionsGetRequestOptionsWithDefaults() *TransactionsGetRequestOptions`

NewTransactionsGetRequestOptionsWithDefaults instantiates a new TransactionsGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIds

`func (o *TransactionsGetRequestOptions) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *TransactionsGetRequestOptions) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *TransactionsGetRequestOptions) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *TransactionsGetRequestOptions) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.

### GetCount

`func (o *TransactionsGetRequestOptions) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransactionsGetRequestOptions) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransactionsGetRequestOptions) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TransactionsGetRequestOptions) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetOffset

`func (o *TransactionsGetRequestOptions) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TransactionsGetRequestOptions) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TransactionsGetRequestOptions) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TransactionsGetRequestOptions) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


