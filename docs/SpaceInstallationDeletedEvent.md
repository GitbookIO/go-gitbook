# SpaceInstallationDeletedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**Previous** | [**SpaceInstallationDeletedEventAllOfPrevious**](SpaceInstallationDeletedEventAllOfPrevious.md) |  | 

## Methods

### NewSpaceInstallationDeletedEvent

`func NewSpaceInstallationDeletedEvent(eventId string, type_ string, installationId string, spaceId string, previous SpaceInstallationDeletedEventAllOfPrevious, ) *SpaceInstallationDeletedEvent`

NewSpaceInstallationDeletedEvent instantiates a new SpaceInstallationDeletedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceInstallationDeletedEventWithDefaults

`func NewSpaceInstallationDeletedEventWithDefaults() *SpaceInstallationDeletedEvent`

NewSpaceInstallationDeletedEventWithDefaults instantiates a new SpaceInstallationDeletedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceInstallationDeletedEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceInstallationDeletedEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceInstallationDeletedEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceInstallationDeletedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceInstallationDeletedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceInstallationDeletedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceInstallationDeletedEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceInstallationDeletedEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceInstallationDeletedEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceInstallationDeletedEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceInstallationDeletedEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceInstallationDeletedEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetPrevious

`func (o *SpaceInstallationDeletedEvent) GetPrevious() SpaceInstallationDeletedEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *SpaceInstallationDeletedEvent) GetPreviousOk() (*SpaceInstallationDeletedEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *SpaceInstallationDeletedEvent) SetPrevious(v SpaceInstallationDeletedEventAllOfPrevious)`

SetPrevious sets Previous field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


