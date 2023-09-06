# RecordingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**Command** | **string** |  | 
**Stdout** | **string** |  | 
**Audio** | **string** | WAV audio file, encoded as base64 | 
**Transcript** | **string** | Transcript of the speech | 
**IsFirst** | Pointer to **bool** |  | [optional] 
**Text** | **string** |  | 
**Filename** | **string** |  | 
**FileSnapshot** | **string** |  | 
**FileDiff** | **string** |  | 

## Methods

### NewRecordingEvent

`func NewRecordingEvent(type_ string, timestamp time.Time, command string, stdout string, audio string, transcript string, text string, filename string, fileSnapshot string, fileDiff string, ) *RecordingEvent`

NewRecordingEvent instantiates a new RecordingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingEventWithDefaults

`func NewRecordingEventWithDefaults() *RecordingEvent`

NewRecordingEventWithDefaults instantiates a new RecordingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCommand

`func (o *RecordingEvent) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RecordingEvent) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RecordingEvent) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetStdout

`func (o *RecordingEvent) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *RecordingEvent) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *RecordingEvent) SetStdout(v string)`

SetStdout sets Stdout field to given value.


### GetAudio

`func (o *RecordingEvent) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RecordingEvent) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RecordingEvent) SetAudio(v string)`

SetAudio sets Audio field to given value.


### GetTranscript

`func (o *RecordingEvent) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *RecordingEvent) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *RecordingEvent) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.


### GetIsFirst

`func (o *RecordingEvent) GetIsFirst() bool`

GetIsFirst returns the IsFirst field if non-nil, zero value otherwise.

### GetIsFirstOk

`func (o *RecordingEvent) GetIsFirstOk() (*bool, bool)`

GetIsFirstOk returns a tuple with the IsFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirst

`func (o *RecordingEvent) SetIsFirst(v bool)`

SetIsFirst sets IsFirst field to given value.

### HasIsFirst

`func (o *RecordingEvent) HasIsFirst() bool`

HasIsFirst returns a boolean if a field has been set.

### GetText

`func (o *RecordingEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RecordingEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RecordingEvent) SetText(v string)`

SetText sets Text field to given value.


### GetFilename

`func (o *RecordingEvent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RecordingEvent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RecordingEvent) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetFileSnapshot

`func (o *RecordingEvent) GetFileSnapshot() string`

GetFileSnapshot returns the FileSnapshot field if non-nil, zero value otherwise.

### GetFileSnapshotOk

`func (o *RecordingEvent) GetFileSnapshotOk() (*string, bool)`

GetFileSnapshotOk returns a tuple with the FileSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSnapshot

`func (o *RecordingEvent) SetFileSnapshot(v string)`

SetFileSnapshot sets FileSnapshot field to given value.


### GetFileDiff

`func (o *RecordingEvent) GetFileDiff() string`

GetFileDiff returns the FileDiff field if non-nil, zero value otherwise.

### GetFileDiffOk

`func (o *RecordingEvent) GetFileDiffOk() (*string, bool)`

GetFileDiffOk returns a tuple with the FileDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiff

`func (o *RecordingEvent) SetFileDiff(v string)`

SetFileDiff sets FileDiff field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


