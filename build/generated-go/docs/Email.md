# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The email address. | 
**Primary** | **bool** | When &#x60;true&#x60;, identifies the email address as the primary email on an account. | 
**Type** | **string** | The type of email account as described by the financial institution. | 

## Methods

### NewEmail

`func NewEmail(data string, primary bool, type_ string, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Email) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Email) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Email) SetData(v string)`

SetData sets Data field to given value.


### GetPrimary

`func (o *Email) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Email) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Email) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetType

`func (o *Email) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Email) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Email) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


