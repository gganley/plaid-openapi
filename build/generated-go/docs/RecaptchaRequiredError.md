# RecaptchaRequiredError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorType** | **string** | RECAPTCHA_ERROR | 
**ErrorCode** | **string** | RECAPTCHA_REQUIRED | 
**DisplayMessage** | **string** |  | 
**HttpCode** | **string** | 400 | 
**LinkUserExperience** | **string** | Your user will be prompted to solve a Google reCAPTCHA challenge in the Link Recaptcha pane. If they solve the challenge successfully, the user&#39;s request is resubmitted and they are directed to the next Item creation step. | 
**CommonCauses** | **string** | Plaid&#39;s fraud system detects abusive traffic and considers a variety of parameters throughout Item creation requests. When a request is considered risky or possibly fraudulent, Link presents a reCAPTCHA for the user to solve. | 
**TroubleshootingSteps** | **string** | Link will automatically guide your user through reCAPTCHA verification. As a general rule, we recommend instrumenting basic fraud monitoring to detect and protect your website from spam and abuse.  If your user cannot verify their session, please submit a Support ticket with the following identifiers: &#x60;link_session_id&#x60; or &#x60;request_id&#x60; | 

## Methods

### NewRecaptchaRequiredError

`func NewRecaptchaRequiredError(errorType string, errorCode string, displayMessage string, httpCode string, linkUserExperience string, commonCauses string, troubleshootingSteps string, ) *RecaptchaRequiredError`

NewRecaptchaRequiredError instantiates a new RecaptchaRequiredError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecaptchaRequiredErrorWithDefaults

`func NewRecaptchaRequiredErrorWithDefaults() *RecaptchaRequiredError`

NewRecaptchaRequiredErrorWithDefaults instantiates a new RecaptchaRequiredError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorType

`func (o *RecaptchaRequiredError) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *RecaptchaRequiredError) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *RecaptchaRequiredError) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.


### GetErrorCode

`func (o *RecaptchaRequiredError) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *RecaptchaRequiredError) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *RecaptchaRequiredError) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetDisplayMessage

`func (o *RecaptchaRequiredError) GetDisplayMessage() string`

GetDisplayMessage returns the DisplayMessage field if non-nil, zero value otherwise.

### GetDisplayMessageOk

`func (o *RecaptchaRequiredError) GetDisplayMessageOk() (*string, bool)`

GetDisplayMessageOk returns a tuple with the DisplayMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMessage

`func (o *RecaptchaRequiredError) SetDisplayMessage(v string)`

SetDisplayMessage sets DisplayMessage field to given value.


### GetHttpCode

`func (o *RecaptchaRequiredError) GetHttpCode() string`

GetHttpCode returns the HttpCode field if non-nil, zero value otherwise.

### GetHttpCodeOk

`func (o *RecaptchaRequiredError) GetHttpCodeOk() (*string, bool)`

GetHttpCodeOk returns a tuple with the HttpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpCode

`func (o *RecaptchaRequiredError) SetHttpCode(v string)`

SetHttpCode sets HttpCode field to given value.


### GetLinkUserExperience

`func (o *RecaptchaRequiredError) GetLinkUserExperience() string`

GetLinkUserExperience returns the LinkUserExperience field if non-nil, zero value otherwise.

### GetLinkUserExperienceOk

`func (o *RecaptchaRequiredError) GetLinkUserExperienceOk() (*string, bool)`

GetLinkUserExperienceOk returns a tuple with the LinkUserExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUserExperience

`func (o *RecaptchaRequiredError) SetLinkUserExperience(v string)`

SetLinkUserExperience sets LinkUserExperience field to given value.


### GetCommonCauses

`func (o *RecaptchaRequiredError) GetCommonCauses() string`

GetCommonCauses returns the CommonCauses field if non-nil, zero value otherwise.

### GetCommonCausesOk

`func (o *RecaptchaRequiredError) GetCommonCausesOk() (*string, bool)`

GetCommonCausesOk returns a tuple with the CommonCauses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonCauses

`func (o *RecaptchaRequiredError) SetCommonCauses(v string)`

SetCommonCauses sets CommonCauses field to given value.


### GetTroubleshootingSteps

`func (o *RecaptchaRequiredError) GetTroubleshootingSteps() string`

GetTroubleshootingSteps returns the TroubleshootingSteps field if non-nil, zero value otherwise.

### GetTroubleshootingStepsOk

`func (o *RecaptchaRequiredError) GetTroubleshootingStepsOk() (*string, bool)`

GetTroubleshootingStepsOk returns a tuple with the TroubleshootingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleshootingSteps

`func (o *RecaptchaRequiredError) SetTroubleshootingSteps(v string)`

SetTroubleshootingSteps sets TroubleshootingSteps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


