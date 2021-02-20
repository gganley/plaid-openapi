# MFA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Possible values are &#x60;device&#x60;, &#x60;selections&#x60;, or &#x60;questions&#x60;.  If value is &#x60;device&#x60;, the MFA answer is &#x60;1234&#x60;.  If value is &#x60;selections&#x60;, the MFA answer is always the first option.  If value is &#x60;questions&#x60;, the MFA answer is  &#x60;answer_&lt;i&gt;_&lt;j&gt;&#x60; for the j-th question in the i-th round, starting from 0. For example, the answer to the first question in the second round is &#x60;answer_1_0&#x60;. | 
**QuestionRounds** | **float32** | Number of rounds of questions. Required if value of &#x60;type&#x60; is &#x60;questions&#x60;.  | 
**QuestionsPerRound** | **float32** | Number of questions per round. Required if value of &#x60;type&#x60; is &#x60;questions&#x60;. If value of type is &#x60;selections&#x60;, default value is 2. | 
**SelectionRounds** | **float32** | Number of rounds of selections, used if &#x60;type&#x60; is &#x60;selections&#x60;. Defaults to 1. | 
**SelectionsPerQuestion** | **float32** | Number of available answers per question, used if &#x60;type&#x60; is &#x60;selection&#x60;. Defaults to 2.  | 

## Methods

### NewMFA

`func NewMFA(type_ string, questionRounds float32, questionsPerRound float32, selectionRounds float32, selectionsPerQuestion float32, ) *MFA`

NewMFA instantiates a new MFA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAWithDefaults

`func NewMFAWithDefaults() *MFA`

NewMFAWithDefaults instantiates a new MFA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFA) SetType(v string)`

SetType sets Type field to given value.


### GetQuestionRounds

`func (o *MFA) GetQuestionRounds() float32`

GetQuestionRounds returns the QuestionRounds field if non-nil, zero value otherwise.

### GetQuestionRoundsOk

`func (o *MFA) GetQuestionRoundsOk() (*float32, bool)`

GetQuestionRoundsOk returns a tuple with the QuestionRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionRounds

`func (o *MFA) SetQuestionRounds(v float32)`

SetQuestionRounds sets QuestionRounds field to given value.


### GetQuestionsPerRound

`func (o *MFA) GetQuestionsPerRound() float32`

GetQuestionsPerRound returns the QuestionsPerRound field if non-nil, zero value otherwise.

### GetQuestionsPerRoundOk

`func (o *MFA) GetQuestionsPerRoundOk() (*float32, bool)`

GetQuestionsPerRoundOk returns a tuple with the QuestionsPerRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionsPerRound

`func (o *MFA) SetQuestionsPerRound(v float32)`

SetQuestionsPerRound sets QuestionsPerRound field to given value.


### GetSelectionRounds

`func (o *MFA) GetSelectionRounds() float32`

GetSelectionRounds returns the SelectionRounds field if non-nil, zero value otherwise.

### GetSelectionRoundsOk

`func (o *MFA) GetSelectionRoundsOk() (*float32, bool)`

GetSelectionRoundsOk returns a tuple with the SelectionRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRounds

`func (o *MFA) SetSelectionRounds(v float32)`

SetSelectionRounds sets SelectionRounds field to given value.


### GetSelectionsPerQuestion

`func (o *MFA) GetSelectionsPerQuestion() float32`

GetSelectionsPerQuestion returns the SelectionsPerQuestion field if non-nil, zero value otherwise.

### GetSelectionsPerQuestionOk

`func (o *MFA) GetSelectionsPerQuestionOk() (*float32, bool)`

GetSelectionsPerQuestionOk returns a tuple with the SelectionsPerQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionsPerQuestion

`func (o *MFA) SetSelectionsPerQuestion(v float32)`

SetSelectionsPerQuestion sets SelectionsPerQuestion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


