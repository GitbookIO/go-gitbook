# SpaceVisibilityUpdatedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**PreviousVisibility** | [**ContentVisibility**](ContentVisibility.md) |  | 
**Visibility** | [**ContentVisibility**](ContentVisibility.md) |  | 

## Methods

### NewSpaceVisibilityUpdatedEvent

`func NewSpaceVisibilityUpdatedEvent(eventId string, type_ string, installationId string, spaceId string, previousVisibility ContentVisibility, visibility ContentVisibility, ) *SpaceVisibilityUpdatedEvent`

NewSpaceVisibilityUpdatedEvent instantiates a new SpaceVisibilityUpdatedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceVisibilityUpdatedEventWithDefaults

`func NewSpaceVisibilityUpdatedEventWithDefaults() *SpaceVisibilityUpdatedEvent`

NewSpaceVisibilityUpdatedEventWithDefaults instantiates a new SpaceVisibilityUpdatedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceVisibilityUpdatedEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceVisibilityUpdatedEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceVisibilityUpdatedEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceVisibilityUpdatedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceVisibilityUpdatedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceVisibilityUpdatedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceVisibilityUpdatedEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceVisibilityUpdatedEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceVisibilityUpdatedEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceVisibilityUpdatedEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceVisibilityUpdatedEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceVisibilityUpdatedEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetPreviousVisibility

`func (o *SpaceVisibilityUpdatedEvent) GetPreviousVisibility() ContentVisibility`

GetPreviousVisibility returns the PreviousVisibility field if non-nil, zero value otherwise.

### GetPreviousVisibilityOk

`func (o *SpaceVisibilityUpdatedEvent) GetPreviousVisibilityOk() (*ContentVisibility, bool)`

GetPreviousVisibilityOk returns a tuple with the PreviousVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVisibility

`func (o *SpaceVisibilityUpdatedEvent) SetPreviousVisibility(v ContentVisibility)`

SetPreviousVisibility sets PreviousVisibility field to given value.


### GetVisibility

`func (o *SpaceVisibilityUpdatedEvent) GetVisibility() ContentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *SpaceVisibilityUpdatedEvent) GetVisibilityOk() (*ContentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *SpaceVisibilityUpdatedEvent) SetVisibility(v ContentVisibility)`

SetVisibility sets Visibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


