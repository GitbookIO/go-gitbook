# RecordingFileAddedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**Filename** | **string** |  | 
**FileSnapshot** | **string** |  | 

## Methods

### NewRecordingFileAddedEvent

`func NewRecordingFileAddedEvent(type_ string, timestamp time.Time, filename string, fileSnapshot string, ) *RecordingFileAddedEvent`

NewRecordingFileAddedEvent instantiates a new RecordingFileAddedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingFileAddedEventWithDefaults

`func NewRecordingFileAddedEventWithDefaults() *RecordingFileAddedEvent`

NewRecordingFileAddedEventWithDefaults instantiates a new RecordingFileAddedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingFileAddedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingFileAddedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingFileAddedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingFileAddedEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingFileAddedEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingFileAddedEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingFileAddedEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingFileAddedEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingFileAddedEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingFileAddedEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingFileAddedEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingFileAddedEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingFileAddedEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingFileAddedEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetFilename

`func (o *RecordingFileAddedEvent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RecordingFileAddedEvent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RecordingFileAddedEvent) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetFileSnapshot

`func (o *RecordingFileAddedEvent) GetFileSnapshot() string`

GetFileSnapshot returns the FileSnapshot field if non-nil, zero value otherwise.

### GetFileSnapshotOk

`func (o *RecordingFileAddedEvent) GetFileSnapshotOk() (*string, bool)`

GetFileSnapshotOk returns a tuple with the FileSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSnapshot

`func (o *RecordingFileAddedEvent) SetFileSnapshot(v string)`

SetFileSnapshot sets FileSnapshot field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


