# SpaceContentUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**RevisionId** | **string** | Unique identifier of the new content revision | 

## Methods

### NewSpaceContentUpdatedEvent

`func NewSpaceContentUpdatedEvent(eventId string, type_ string, installationId string, spaceId string, revisionId string, ) *SpaceContentUpdatedEvent`

NewSpaceContentUpdatedEvent instantiates a new SpaceContentUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceContentUpdatedEventWithDefaults

`func NewSpaceContentUpdatedEventWithDefaults() *SpaceContentUpdatedEvent`

NewSpaceContentUpdatedEventWithDefaults instantiates a new SpaceContentUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceContentUpdatedEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceContentUpdatedEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceContentUpdatedEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceContentUpdatedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceContentUpdatedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceContentUpdatedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceContentUpdatedEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceContentUpdatedEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceContentUpdatedEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceContentUpdatedEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceContentUpdatedEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceContentUpdatedEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetRevisionId

`func (o *SpaceContentUpdatedEvent) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *SpaceContentUpdatedEvent) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *SpaceContentUpdatedEvent) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


