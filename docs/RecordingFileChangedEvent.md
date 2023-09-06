# RecordingFileChangedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**Filename** | **string** |  | 
**FileDiff** | **string** |  | 

## Methods

### NewRecordingFileChangedEvent

`func NewRecordingFileChangedEvent(type_ string, timestamp time.Time, filename string, fileDiff string, ) *RecordingFileChangedEvent`

NewRecordingFileChangedEvent instantiates a new RecordingFileChangedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingFileChangedEventWithDefaults

`func NewRecordingFileChangedEventWithDefaults() *RecordingFileChangedEvent`

NewRecordingFileChangedEventWithDefaults instantiates a new RecordingFileChangedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingFileChangedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingFileChangedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingFileChangedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingFileChangedEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingFileChangedEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingFileChangedEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingFileChangedEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingFileChangedEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingFileChangedEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingFileChangedEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingFileChangedEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingFileChangedEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingFileChangedEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingFileChangedEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetFilename

`func (o *RecordingFileChangedEvent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RecordingFileChangedEvent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RecordingFileChangedEvent) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetFileDiff

`func (o *RecordingFileChangedEvent) GetFileDiff() string`

GetFileDiff returns the FileDiff field if non-nil, zero value otherwise.

### GetFileDiffOk

`func (o *RecordingFileChangedEvent) GetFileDiffOk() (*string, bool)`

GetFileDiffOk returns a tuple with the FileDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiff

`func (o *RecordingFileChangedEvent) SetFileDiff(v string)`

SetFileDiff sets FileDiff field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


