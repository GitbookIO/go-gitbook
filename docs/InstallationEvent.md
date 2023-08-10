# InstallationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** | Type of the event. | 
**InstallationId** | **string** | ID of the integration installation | 

## Methods

### NewInstallationEvent

`func NewInstallationEvent(eventId string, type_ string, installationId string, ) *InstallationEvent`

NewInstallationEvent instantiates a new InstallationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationEventWithDefaults

`func NewInstallationEventWithDefaults() *InstallationEvent`

NewInstallationEventWithDefaults instantiates a new InstallationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *InstallationEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *InstallationEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *InstallationEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *InstallationEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallationEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallationEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *InstallationEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *InstallationEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *InstallationEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


