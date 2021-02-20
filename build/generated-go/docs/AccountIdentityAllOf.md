# AccountIdentityAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owners** | [**[]Owner**](Owner.md) | Data returned by the financial institution about the account owner or owners. Only returned by Identity or Assets endpoints. Multiple owners on a single account will be represented in the same &#x60;owner&#x60; object, not in multiple owner objects within the array. | 

## Methods

### NewAccountIdentityAllOf

`func NewAccountIdentityAllOf(owners []Owner, ) *AccountIdentityAllOf`

NewAccountIdentityAllOf instantiates a new AccountIdentityAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdentityAllOfWithDefaults

`func NewAccountIdentityAllOfWithDefaults() *AccountIdentityAllOf`

NewAccountIdentityAllOfWithDefaults instantiates a new AccountIdentityAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwners

`func (o *AccountIdentityAllOf) GetOwners() []Owner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *AccountIdentityAllOf) GetOwnersOk() (*[]Owner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *AccountIdentityAllOf) SetOwners(v []Owner)`

SetOwners sets Owners field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


