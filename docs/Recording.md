# Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;recording\&quot; | 
**Id** | **string** | Unique identifier for the recording | 
**Title** | **string** | Optional title describing the recording | 
**Context** | [**RecordingContext**](RecordingContext.md) |  | 
**ExternalId** | Pointer to **string** | ID in the original source of the recording. | [optional] 
**ExternalURL** | Pointer to **string** | URL of the source from which the recording originated | [optional] 
**CreatedAt** | **time.Time** |  | 
**StoppedAt** | Pointer to **time.Time** |  | [optional] 
**Events** | [**RecordingEvents**](RecordingEvents.md) |  | 
**Contributors** | [**[]RecordingContributorsInner**](RecordingContributorsInner.md) |  | 
**Output** | Pointer to [**RecordingOutput**](RecordingOutput.md) |  | [optional] 
**Urls** | [**RecordingUrls**](RecordingUrls.md) |  | 

## Methods

### NewRecording

`func NewRecording(object string, id string, title string, context RecordingContext, createdAt time.Time, events RecordingEvents, contributors []RecordingContributorsInner, urls RecordingUrls, ) *Recording`

NewRecording instantiates a new Recording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingWithDefaults

`func NewRecordingWithDefaults() *Recording`

NewRecordingWithDefaults instantiates a new Recording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Recording) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Recording) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Recording) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Recording) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Recording) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Recording) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Recording) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Recording) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Recording) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetContext

`func (o *Recording) GetContext() RecordingContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Recording) GetContextOk() (*RecordingContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Recording) SetContext(v RecordingContext)`

SetContext sets Context field to given value.


### GetExternalId

`func (o *Recording) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Recording) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Recording) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Recording) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalURL

`func (o *Recording) GetExternalURL() string`

GetExternalURL returns the ExternalURL field if non-nil, zero value otherwise.

### GetExternalURLOk

`func (o *Recording) GetExternalURLOk() (*string, bool)`

GetExternalURLOk returns a tuple with the ExternalURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalURL

`func (o *Recording) SetExternalURL(v string)`

SetExternalURL sets ExternalURL field to given value.

### HasExternalURL

`func (o *Recording) HasExternalURL() bool`

HasExternalURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Recording) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Recording) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Recording) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStoppedAt

`func (o *Recording) GetStoppedAt() time.Time`

GetStoppedAt returns the StoppedAt field if non-nil, zero value otherwise.

### GetStoppedAtOk

`func (o *Recording) GetStoppedAtOk() (*time.Time, bool)`

GetStoppedAtOk returns a tuple with the StoppedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedAt

`func (o *Recording) SetStoppedAt(v time.Time)`

SetStoppedAt sets StoppedAt field to given value.

### HasStoppedAt

`func (o *Recording) HasStoppedAt() bool`

HasStoppedAt returns a boolean if a field has been set.

### GetEvents

`func (o *Recording) GetEvents() RecordingEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Recording) GetEventsOk() (*RecordingEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Recording) SetEvents(v RecordingEvents)`

SetEvents sets Events field to given value.


### GetContributors

`func (o *Recording) GetContributors() []RecordingContributorsInner`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *Recording) GetContributorsOk() (*[]RecordingContributorsInner, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *Recording) SetContributors(v []RecordingContributorsInner)`

SetContributors sets Contributors field to given value.


### GetOutput

`func (o *Recording) GetOutput() RecordingOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *Recording) GetOutputOk() (*RecordingOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *Recording) SetOutput(v RecordingOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *Recording) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetUrls

`func (o *Recording) GetUrls() RecordingUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Recording) GetUrlsOk() (*RecordingUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Recording) SetUrls(v RecordingUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


