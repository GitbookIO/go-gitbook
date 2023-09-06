# RecordingThreadMessageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timestamp** | **time.Time** | When the event happened | 
**Source** | Pointer to **string** | Optionally, provide the source of the event. GitBook may use this to improve the generated content. | [optional] 
**Actor** | Pointer to [**BaseRecordingEventActor**](BaseRecordingEventActor.md) |  | [optional] 
**IsFirst** | Pointer to **bool** |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewRecordingThreadMessageEvent

`func NewRecordingThreadMessageEvent(type_ string, timestamp time.Time, text string, ) *RecordingThreadMessageEvent`

NewRecordingThreadMessageEvent instantiates a new RecordingThreadMessageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingThreadMessageEventWithDefaults

`func NewRecordingThreadMessageEventWithDefaults() *RecordingThreadMessageEvent`

NewRecordingThreadMessageEventWithDefaults instantiates a new RecordingThreadMessageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RecordingThreadMessageEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordingThreadMessageEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordingThreadMessageEvent) SetType(v string)`

SetType sets Type field to given value.


### GetTimestamp

`func (o *RecordingThreadMessageEvent) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *RecordingThreadMessageEvent) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *RecordingThreadMessageEvent) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetSource

`func (o *RecordingThreadMessageEvent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecordingThreadMessageEvent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecordingThreadMessageEvent) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *RecordingThreadMessageEvent) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetActor

`func (o *RecordingThreadMessageEvent) GetActor() BaseRecordingEventActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RecordingThreadMessageEvent) GetActorOk() (*BaseRecordingEventActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RecordingThreadMessageEvent) SetActor(v BaseRecordingEventActor)`

SetActor sets Actor field to given value.

### HasActor

`func (o *RecordingThreadMessageEvent) HasActor() bool`

HasActor returns a boolean if a field has been set.

### GetIsFirst

`func (o *RecordingThreadMessageEvent) GetIsFirst() bool`

GetIsFirst returns the IsFirst field if non-nil, zero value otherwise.

### GetIsFirstOk

`func (o *RecordingThreadMessageEvent) GetIsFirstOk() (*bool, bool)`

GetIsFirstOk returns a tuple with the IsFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirst

`func (o *RecordingThreadMessageEvent) SetIsFirst(v bool)`

SetIsFirst sets IsFirst field to given value.

### HasIsFirst

`func (o *RecordingThreadMessageEvent) HasIsFirst() bool`

HasIsFirst returns a boolean if a field has been set.

### GetText

`func (o *RecordingThreadMessageEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RecordingThreadMessageEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RecordingThreadMessageEvent) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


