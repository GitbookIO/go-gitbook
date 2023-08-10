# FetchEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**SpaceId** | Pointer to **string** | The space ID, if requests are specific to a single space | [optional] 
**InstallationId** | Pointer to **string** | The installation ID, if requests are specific to a single installation | [optional] 
**Auth** | Pointer to [**FetchEventAllOfAuth**](FetchEventAllOfAuth.md) |  | [optional] 
**Request** | [**FetchRequest**](FetchRequest.md) |  | 

## Methods

### NewFetchEvent

`func NewFetchEvent(eventId string, type_ string, request FetchRequest, ) *FetchEvent`

NewFetchEvent instantiates a new FetchEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchEventWithDefaults

`func NewFetchEventWithDefaults() *FetchEvent`

NewFetchEventWithDefaults instantiates a new FetchEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *FetchEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *FetchEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *FetchEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *FetchEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchEvent) SetType(v string)`

SetType sets Type field to given value.


### GetSpaceId

`func (o *FetchEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *FetchEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *FetchEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *FetchEvent) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetInstallationId

`func (o *FetchEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *FetchEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *FetchEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *FetchEvent) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetAuth

`func (o *FetchEvent) GetAuth() FetchEventAllOfAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *FetchEvent) GetAuthOk() (*FetchEventAllOfAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *FetchEvent) SetAuth(v FetchEventAllOfAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *FetchEvent) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetRequest

`func (o *FetchEvent) GetRequest() FetchRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *FetchEvent) GetRequestOk() (*FetchRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *FetchEvent) SetRequest(v FetchRequest)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


