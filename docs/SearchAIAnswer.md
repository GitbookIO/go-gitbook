# SearchAIAnswer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**FollowupQuestions** | **[]string** |  | 
**Pages** | [**[]SearchAIAnswerPagesInner**](SearchAIAnswerPagesInner.md) |  | 

## Methods

### NewSearchAIAnswer

`func NewSearchAIAnswer(text string, followupQuestions []string, pages []SearchAIAnswerPagesInner, ) *SearchAIAnswer`

NewSearchAIAnswer instantiates a new SearchAIAnswer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchAIAnswerWithDefaults

`func NewSearchAIAnswerWithDefaults() *SearchAIAnswer`

NewSearchAIAnswerWithDefaults instantiates a new SearchAIAnswer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *SearchAIAnswer) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SearchAIAnswer) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SearchAIAnswer) SetText(v string)`

SetText sets Text field to given value.


### GetFollowupQuestions

`func (o *SearchAIAnswer) GetFollowupQuestions() []string`

GetFollowupQuestions returns the FollowupQuestions field if non-nil, zero value otherwise.

### GetFollowupQuestionsOk

`func (o *SearchAIAnswer) GetFollowupQuestionsOk() (*[]string, bool)`

GetFollowupQuestionsOk returns a tuple with the FollowupQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupQuestions

`func (o *SearchAIAnswer) SetFollowupQuestions(v []string)`

SetFollowupQuestions sets FollowupQuestions field to given value.


### GetPages

`func (o *SearchAIAnswer) GetPages() []SearchAIAnswerPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *SearchAIAnswer) GetPagesOk() (*[]SearchAIAnswerPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *SearchAIAnswer) SetPages(v []SearchAIAnswerPagesInner)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


