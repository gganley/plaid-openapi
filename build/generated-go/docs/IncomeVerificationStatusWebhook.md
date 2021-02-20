# IncomeVerificationStatusWebhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookType** | **string** | &#x60;\&quot;INCOME\&quot;&#x60; | 
**WebhookCode** | **string** | &#x60;income_verification&#x60; | 
**IncomeVerificationId** | **string** | The &#x60;income_verification_id&#x60; of the verification instance whose status is being reported. | 
**VerificationStatus** | **string** | &#x60;VERIFICATION_STATUS_PROCESSING_COMPLETE&#x60;: The income verification status processing has completed.  &#x60;VERIFICATION_STATUS_UPLOAD_ERROR&#x60;: An upload error occurred when the end user attempted to upload their verification documentation.  &#x60;VERIFICATION_STATUS_INVALID_TYPE&#x60;: The end user attempted to upload verification documentation in an unsupported file format.  &#x60;VERIFICATION_STATUS_DOCUMENT_REJECTED&#x60;: The documentation uploaded by the end user was recognized as a supported file format, but not recognized as a valid paystub.  &#x60;VERIFICATION_STATUS_PROCESSING_FAILED&#x60;: A failure occurred when attempting to process the verification documentation. | 

## Methods

### NewIncomeVerificationStatusWebhook

`func NewIncomeVerificationStatusWebhook(webhookType string, webhookCode string, incomeVerificationId string, verificationStatus string, ) *IncomeVerificationStatusWebhook`

NewIncomeVerificationStatusWebhook instantiates a new IncomeVerificationStatusWebhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationStatusWebhookWithDefaults

`func NewIncomeVerificationStatusWebhookWithDefaults() *IncomeVerificationStatusWebhook`

NewIncomeVerificationStatusWebhookWithDefaults instantiates a new IncomeVerificationStatusWebhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookType

`func (o *IncomeVerificationStatusWebhook) GetWebhookType() string`

GetWebhookType returns the WebhookType field if non-nil, zero value otherwise.

### GetWebhookTypeOk

`func (o *IncomeVerificationStatusWebhook) GetWebhookTypeOk() (*string, bool)`

GetWebhookTypeOk returns a tuple with the WebhookType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookType

`func (o *IncomeVerificationStatusWebhook) SetWebhookType(v string)`

SetWebhookType sets WebhookType field to given value.


### GetWebhookCode

`func (o *IncomeVerificationStatusWebhook) GetWebhookCode() string`

GetWebhookCode returns the WebhookCode field if non-nil, zero value otherwise.

### GetWebhookCodeOk

`func (o *IncomeVerificationStatusWebhook) GetWebhookCodeOk() (*string, bool)`

GetWebhookCodeOk returns a tuple with the WebhookCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookCode

`func (o *IncomeVerificationStatusWebhook) SetWebhookCode(v string)`

SetWebhookCode sets WebhookCode field to given value.


### GetIncomeVerificationId

`func (o *IncomeVerificationStatusWebhook) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationStatusWebhook) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationStatusWebhook) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.


### GetVerificationStatus

`func (o *IncomeVerificationStatusWebhook) GetVerificationStatus() string`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *IncomeVerificationStatusWebhook) GetVerificationStatusOk() (*string, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *IncomeVerificationStatusWebhook) SetVerificationStatus(v string)`

SetVerificationStatus sets VerificationStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


