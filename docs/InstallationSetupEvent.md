# InstallationSetupEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Previous** | Pointer to [**InstallationSetupEventAllOfPrevious**](InstallationSetupEventAllOfPrevious.md) |  | [optional] 

## Methods

### NewInstallationSetupEvent

`func NewInstallationSetupEvent(eventId string, type_ string, installationId string, status IntegrationInstallationStatus, ) *InstallationSetupEvent`

NewInstallationSetupEvent instantiates a new InstallationSetupEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstallationSetupEventWithDefaults

`func NewInstallationSetupEventWithDefaults() *InstallationSetupEvent`

NewInstallationSetupEventWithDefaults instantiates a new InstallationSetupEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *InstallationSetupEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *InstallationSetupEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *InstallationSetupEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *InstallationSetupEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstallationSetupEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstallationSetupEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *InstallationSetupEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *InstallationSetupEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *InstallationSetupEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetStatus

`func (o *InstallationSetupEvent) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstallationSetupEvent) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstallationSetupEvent) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetPrevious

`func (o *InstallationSetupEvent) GetPrevious() InstallationSetupEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *InstallationSetupEvent) GetPreviousOk() (*InstallationSetupEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *InstallationSetupEvent) SetPrevious(v InstallationSetupEventAllOfPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *InstallationSetupEvent) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


