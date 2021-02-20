# NumbersACH

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The Plaid account ID associated with the account numbers | 
**Account** | **string** | The ACH account number for the account | 
**Routing** | **string** | The ACH routing number for the account | 
**WireRouting** | Pointer to **NullableString** | The wire transfer routing number for the account, if available | [optional] 

## Methods

### NewNumbersACH

`func NewNumbersACH(accountId string, account string, routing string, ) *NumbersACH`

NewNumbersACH instantiates a new NumbersACH object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumbersACHWithDefaults

`func NewNumbersACHWithDefaults() *NumbersACH`

NewNumbersACHWithDefaults instantiates a new NumbersACH object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NumbersACH) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NumbersACH) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NumbersACH) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAccount

`func (o *NumbersACH) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NumbersACH) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NumbersACH) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetRouting

`func (o *NumbersACH) GetRouting() string`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *NumbersACH) GetRoutingOk() (*string, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *NumbersACH) SetRouting(v string)`

SetRouting sets Routing field to given value.


### GetWireRouting

`func (o *NumbersACH) GetWireRouting() string`

GetWireRouting returns the WireRouting field if non-nil, zero value otherwise.

### GetWireRoutingOk

`func (o *NumbersACH) GetWireRoutingOk() (*string, bool)`

GetWireRoutingOk returns a tuple with the WireRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireRouting

`func (o *NumbersACH) SetWireRouting(v string)`

SetWireRouting sets WireRouting field to given value.

### HasWireRouting

`func (o *NumbersACH) HasWireRouting() bool`

HasWireRouting returns a boolean if a field has been set.

### SetWireRoutingNil

`func (o *NumbersACH) SetWireRoutingNil(b bool)`

 SetWireRoutingNil sets the value for WireRouting to be an explicit nil

### UnsetWireRouting
`func (o *NumbersACH) UnsetWireRouting()`

UnsetWireRouting ensures that no value is present for WireRouting, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


