# EmployersSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. | [optional] 
**Query** | **string** | The employer name to be searched for. | 
**Products** | **[]string** | The Plaid products the returned employers should support. Currently, this field must be set to &#x60;\&quot;deposit_switch\&quot;&#x60;. | 

## Methods

### NewEmployersSearchRequest

`func NewEmployersSearchRequest(query string, products []string, ) *EmployersSearchRequest`

NewEmployersSearchRequest instantiates a new EmployersSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersSearchRequestWithDefaults

`func NewEmployersSearchRequestWithDefaults() *EmployersSearchRequest`

NewEmployersSearchRequestWithDefaults instantiates a new EmployersSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *EmployersSearchRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EmployersSearchRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EmployersSearchRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *EmployersSearchRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *EmployersSearchRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *EmployersSearchRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *EmployersSearchRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *EmployersSearchRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetQuery

`func (o *EmployersSearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EmployersSearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EmployersSearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetProducts

`func (o *EmployersSearchRequest) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EmployersSearchRequest) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EmployersSearchRequest) SetProducts(v []string)`

SetProducts sets Products field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


