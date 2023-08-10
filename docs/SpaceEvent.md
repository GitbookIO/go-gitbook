# SpaceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** | Type of the event. | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 

## Methods

### NewSpaceEvent

`func NewSpaceEvent(eventId string, type_ string, installationId string, spaceId string, ) *SpaceEvent`

NewSpaceEvent instantiates a new SpaceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceEventWithDefaults

`func NewSpaceEventWithDefaults() *SpaceEvent`

NewSpaceEventWithDefaults instantiates a new SpaceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


