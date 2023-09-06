# RecordingSpeechEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**Audio** | **string** | WAV audio file, encoded as base64 | 
**Transcript** | **string** | Transcript of the speech | 

## Methods

### NewRecordingSpeechEvent

`func NewRecordingSpeechEvent(type_ string, timestamp time.Time, audio string, transcript string, ) *RecordingSpeechEvent`

NewRecordingSpeechEvent instantiates a new RecordingSpeechEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingSpeechEventWithDefaults

`func NewRecordingSpeechEventWithDefaults() *RecordingSpeechEvent`

NewRecordingSpeechEventWithDefaults instantiates a new RecordingSpeechEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingSpeechEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingSpeechEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingSpeechEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingSpeechEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingSpeechEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingSpeechEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingSpeechEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingSpeechEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingSpeechEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingSpeechEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingSpeechEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingSpeechEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingSpeechEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingSpeechEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetAudio

`func (o *RecordingSpeechEvent) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RecordingSpeechEvent) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RecordingSpeechEvent) SetAudio(v string)`

SetAudio sets Audio field to given value.


### GetTranscript

`func (o *RecordingSpeechEvent) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *RecordingSpeechEvent) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *RecordingSpeechEvent) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


