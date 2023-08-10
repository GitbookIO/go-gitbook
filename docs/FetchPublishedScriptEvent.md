# FetchPublishedScriptEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 

## Methods

### NewFetchPublishedScriptEvent

`func NewFetchPublishedScriptEvent(eventId string, type_ string, installationId string, spaceId string, ) *FetchPublishedScriptEvent`

NewFetchPublishedScriptEvent instantiates a new FetchPublishedScriptEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchPublishedScriptEventWithDefaults

`func NewFetchPublishedScriptEventWithDefaults() *FetchPublishedScriptEvent`

NewFetchPublishedScriptEventWithDefaults instantiates a new FetchPublishedScriptEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *FetchPublishedScriptEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *FetchPublishedScriptEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *FetchPublishedScriptEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *FetchPublishedScriptEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchPublishedScriptEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchPublishedScriptEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *FetchPublishedScriptEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *FetchPublishedScriptEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *FetchPublishedScriptEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *FetchPublishedScriptEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *FetchPublishedScriptEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *FetchPublishedScriptEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


