# StopRecording200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recording** | [**RecordingOutput**](RecordingOutput.md) |  | 
**FollowupQuestions** | Pointer to **[]string** | Example questions that would be answered by the content of this recording. | [optional] 

## Methods

### NewStopRecording200Response

`func NewStopRecording200Response(recording RecordingOutput, ) *StopRecording200Response`

NewStopRecording200Response instantiates a new StopRecording200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStopRecording200ResponseWithDefaults

`func NewStopRecording200ResponseWithDefaults() *StopRecording200Response`

NewStopRecording200ResponseWithDefaults instantiates a new StopRecording200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecording

`func (o *StopRecording200Response) GetRecording() RecordingOutput`

GetRecording returns the Recording field if non-nil, zero value otherwise.

### GetRecordingOk

`func (o *StopRecording200Response) GetRecordingOk() (*RecordingOutput, bool)`

GetRecordingOk returns a tuple with the Recording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecording

`func (o *StopRecording200Response) SetRecording(v RecordingOutput)`

SetRecording sets Recording field to given value.


### GetFollowupQuestions

`func (o *StopRecording200Response) GetFollowupQuestions() []string`

GetFollowupQuestions returns the FollowupQuestions field if non-nil, zero value otherwise.

### GetFollowupQuestionsOk

`func (o *StopRecording200Response) GetFollowupQuestionsOk() (*[]string, bool)`

GetFollowupQuestionsOk returns a tuple with the FollowupQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupQuestions

`func (o *StopRecording200Response) SetFollowupQuestions(v []string)`

SetFollowupQuestions sets FollowupQuestions field to given value.

### HasFollowupQuestions

`func (o *StopRecording200Response) HasFollowupQuestions() bool`

HasFollowupQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


