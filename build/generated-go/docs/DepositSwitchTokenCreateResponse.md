# DepositSwitchTokenCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepositSwitchToken** | **string** | Deposit switch token, used to initialize Link for the Deposit Switch product | 
**DepositSwitchTokenExpirationTime** | **string** | Expiration time of the token, in ISO8601 format | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewDepositSwitchTokenCreateResponse

`func NewDepositSwitchTokenCreateResponse(depositSwitchToken string, depositSwitchTokenExpirationTime string, requestId string, ) *DepositSwitchTokenCreateResponse`

NewDepositSwitchTokenCreateResponse instantiates a new DepositSwitchTokenCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositSwitchTokenCreateResponseWithDefaults

`func NewDepositSwitchTokenCreateResponseWithDefaults() *DepositSwitchTokenCreateResponse`

NewDepositSwitchTokenCreateResponseWithDefaults instantiates a new DepositSwitchTokenCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepositSwitchToken

`func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchToken() string`

GetDepositSwitchToken returns the DepositSwitchToken field if non-nil, zero value otherwise.

### GetDepositSwitchTokenOk

`func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenOk() (*string, bool)`

GetDepositSwitchTokenOk returns a tuple with the DepositSwitchToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchToken

`func (o *DepositSwitchTokenCreateResponse) SetDepositSwitchToken(v string)`

SetDepositSwitchToken sets DepositSwitchToken field to given value.


### GetDepositSwitchTokenExpirationTime

`func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenExpirationTime() string`

GetDepositSwitchTokenExpirationTime returns the DepositSwitchTokenExpirationTime field if non-nil, zero value otherwise.

### GetDepositSwitchTokenExpirationTimeOk

`func (o *DepositSwitchTokenCreateResponse) GetDepositSwitchTokenExpirationTimeOk() (*string, bool)`

GetDepositSwitchTokenExpirationTimeOk returns a tuple with the DepositSwitchTokenExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepositSwitchTokenExpirationTime

`func (o *DepositSwitchTokenCreateResponse) SetDepositSwitchTokenExpirationTime(v string)`

SetDepositSwitchTokenExpirationTime sets DepositSwitchTokenExpirationTime field to given value.


### GetRequestId

`func (o *DepositSwitchTokenCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *DepositSwitchTokenCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *DepositSwitchTokenCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


