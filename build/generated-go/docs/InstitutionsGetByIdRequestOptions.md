# InstitutionsGetByIdRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeOptionalMetadata** | Pointer to **bool** | When &#x60;true&#x60;, return an institution&#39;s logo, brand color, and URL. When available, the bank&#39;s logo is returned as a base64 encoded 152x152 PNG, the brand color is in hexadecimal format. The default value is &#x60;false&#x60;.  Note that Plaid does not own any of the logos shared by the API and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos. | [optional] [default to false]
**IncludeStatus** | Pointer to **bool** | If &#x60;true&#x60;, the response will include status information about the institution. Default value is &#x60;false&#x60;. | [optional] [default to false]

## Methods

### NewInstitutionsGetByIdRequestOptions

`func NewInstitutionsGetByIdRequestOptions() *InstitutionsGetByIdRequestOptions`

NewInstitutionsGetByIdRequestOptions instantiates a new InstitutionsGetByIdRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetByIdRequestOptionsWithDefaults

`func NewInstitutionsGetByIdRequestOptionsWithDefaults() *InstitutionsGetByIdRequestOptions`

NewInstitutionsGetByIdRequestOptionsWithDefaults instantiates a new InstitutionsGetByIdRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeOptionalMetadata

`func (o *InstitutionsGetByIdRequestOptions) GetIncludeOptionalMetadata() bool`

GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field if non-nil, zero value otherwise.

### GetIncludeOptionalMetadataOk

`func (o *InstitutionsGetByIdRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool)`

GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOptionalMetadata

`func (o *InstitutionsGetByIdRequestOptions) SetIncludeOptionalMetadata(v bool)`

SetIncludeOptionalMetadata sets IncludeOptionalMetadata field to given value.

### HasIncludeOptionalMetadata

`func (o *InstitutionsGetByIdRequestOptions) HasIncludeOptionalMetadata() bool`

HasIncludeOptionalMetadata returns a boolean if a field has been set.

### GetIncludeStatus

`func (o *InstitutionsGetByIdRequestOptions) GetIncludeStatus() bool`

GetIncludeStatus returns the IncludeStatus field if non-nil, zero value otherwise.

### GetIncludeStatusOk

`func (o *InstitutionsGetByIdRequestOptions) GetIncludeStatusOk() (*bool, bool)`

GetIncludeStatusOk returns a tuple with the IncludeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStatus

`func (o *InstitutionsGetByIdRequestOptions) SetIncludeStatus(v bool)`

SetIncludeStatus sets IncludeStatus field to given value.

### HasIncludeStatus

`func (o *InstitutionsGetByIdRequestOptions) HasIncludeStatus() bool`

HasIncludeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


