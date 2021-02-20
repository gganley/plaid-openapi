# IncomeVerificationSummaryGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomeSummaries** | [**[]IncomeSummary**](IncomeSummary.md) | A list of income summaries. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewIncomeVerificationSummaryGetResponse

`func NewIncomeVerificationSummaryGetResponse(incomeSummaries []IncomeSummary, requestId string, ) *IncomeVerificationSummaryGetResponse`

NewIncomeVerificationSummaryGetResponse instantiates a new IncomeVerificationSummaryGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationSummaryGetResponseWithDefaults

`func NewIncomeVerificationSummaryGetResponseWithDefaults() *IncomeVerificationSummaryGetResponse`

NewIncomeVerificationSummaryGetResponseWithDefaults instantiates a new IncomeVerificationSummaryGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomeSummaries

`func (o *IncomeVerificationSummaryGetResponse) GetIncomeSummaries() []IncomeSummary`

GetIncomeSummaries returns the IncomeSummaries field if non-nil, zero value otherwise.

### GetIncomeSummariesOk

`func (o *IncomeVerificationSummaryGetResponse) GetIncomeSummariesOk() (*[]IncomeSummary, bool)`

GetIncomeSummariesOk returns a tuple with the IncomeSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeSummaries

`func (o *IncomeVerificationSummaryGetResponse) SetIncomeSummaries(v []IncomeSummary)`

SetIncomeSummaries sets IncomeSummaries field to given value.


### GetRequestId

`func (o *IncomeVerificationSummaryGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationSummaryGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationSummaryGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


