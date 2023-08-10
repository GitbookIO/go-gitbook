# SpaceInstallationSetupEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Previous** | Pointer to [**SpaceInstallationSetupEventAllOfPrevious**](SpaceInstallationSetupEventAllOfPrevious.md) |  | [optional] 

## Methods

### NewSpaceInstallationSetupEvent

`func NewSpaceInstallationSetupEvent(eventId string, type_ string, installationId string, spaceId string, status IntegrationInstallationStatus, ) *SpaceInstallationSetupEvent`

NewSpaceInstallationSetupEvent instantiates a new SpaceInstallationSetupEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceInstallationSetupEventWithDefaults

`func NewSpaceInstallationSetupEventWithDefaults() *SpaceInstallationSetupEvent`

NewSpaceInstallationSetupEventWithDefaults instantiates a new SpaceInstallationSetupEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceInstallationSetupEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceInstallationSetupEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceInstallationSetupEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceInstallationSetupEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceInstallationSetupEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceInstallationSetupEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceInstallationSetupEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceInstallationSetupEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceInstallationSetupEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceInstallationSetupEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceInstallationSetupEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceInstallationSetupEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetStatus

`func (o *SpaceInstallationSetupEvent) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpaceInstallationSetupEvent) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpaceInstallationSetupEvent) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetPrevious

`func (o *SpaceInstallationSetupEvent) GetPrevious() SpaceInstallationSetupEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *SpaceInstallationSetupEvent) GetPreviousOk() (*SpaceInstallationSetupEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *SpaceInstallationSetupEvent) SetPrevious(v SpaceInstallationSetupEventAllOfPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *SpaceInstallationSetupEvent) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


