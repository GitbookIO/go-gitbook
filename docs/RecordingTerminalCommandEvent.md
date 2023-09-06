# RecordingTerminalCommandEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**Command** | **string** |  | 
**Stdout** | **string** |  | 

## Methods

### NewRecordingTerminalCommandEvent

`func NewRecordingTerminalCommandEvent(type_ string, timestamp time.Time, command string, stdout string, ) *RecordingTerminalCommandEvent`

NewRecordingTerminalCommandEvent instantiates a new RecordingTerminalCommandEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingTerminalCommandEventWithDefaults

`func NewRecordingTerminalCommandEventWithDefaults() *RecordingTerminalCommandEvent`

NewRecordingTerminalCommandEventWithDefaults instantiates a new RecordingTerminalCommandEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingTerminalCommandEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingTerminalCommandEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingTerminalCommandEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingTerminalCommandEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingTerminalCommandEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingTerminalCommandEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingTerminalCommandEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingTerminalCommandEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingTerminalCommandEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingTerminalCommandEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingTerminalCommandEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingTerminalCommandEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingTerminalCommandEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingTerminalCommandEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetCommand

`func (o *RecordingTerminalCommandEvent) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *RecordingTerminalCommandEvent) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *RecordingTerminalCommandEvent) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetStdout

`func (o *RecordingTerminalCommandEvent) GetStdout() string`

GetStdout returns the Stdout field if non-nil, zero value otherwise.

### GetStdoutOk

`func (o *RecordingTerminalCommandEvent) GetStdoutOk() (*string, bool)`

GetStdoutOk returns a tuple with the Stdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdout

`func (o *RecordingTerminalCommandEvent) SetStdout(v string)`

SetStdout sets Stdout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


