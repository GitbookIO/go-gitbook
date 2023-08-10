# UIRenderEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**FetchEventAllOfAuth**](FetchEventAllOfAuth.md) |  | [optional] 
**Type** | **string** | Type of the event. | 
**ComponentId** | **string** |  | 
**Props** | **map[string]interface{}** | Properties to render the UI. | 
**State** | Pointer to **map[string]interface{}** | State of the UI. | [optional] 
**Context** | [**ContentKitContext**](ContentKitContext.md) |  | 
**Action** | Pointer to **map[string]interface{}** |  | [optional] 
**EventId** | **string** | Unique identifier for the event. | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 

## Methods

### NewUIRenderEvent

`func NewUIRenderEvent(type_ string, componentId string, props map[string]interface{}, context ContentKitContext, eventId string, installationId string, spaceId string, ) *UIRenderEvent`

NewUIRenderEvent instantiates a new UIRenderEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIRenderEventWithDefaults

`func NewUIRenderEventWithDefaults() *UIRenderEvent`

NewUIRenderEventWithDefaults instantiates a new UIRenderEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *UIRenderEvent) GetAuth() FetchEventAllOfAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UIRenderEvent) GetAuthOk() (*FetchEventAllOfAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UIRenderEvent) SetAuth(v FetchEventAllOfAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UIRenderEvent) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetType

`func (o *UIRenderEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UIRenderEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UIRenderEvent) SetType(v string)`

SetType sets Type field to given value.


### GetComponentId

`func (o *UIRenderEvent) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *UIRenderEvent) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *UIRenderEvent) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetProps

`func (o *UIRenderEvent) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *UIRenderEvent) GetPropsOk() (*map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *UIRenderEvent) SetProps(v map[string]interface{})`

SetProps sets Props field to given value.


### GetState

`func (o *UIRenderEvent) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UIRenderEvent) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UIRenderEvent) SetState(v map[string]interface{})`

SetState sets State field to given value.

### HasState

`func (o *UIRenderEvent) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContext

`func (o *UIRenderEvent) GetContext() ContentKitContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UIRenderEvent) GetContextOk() (*ContentKitContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UIRenderEvent) SetContext(v ContentKitContext)`

SetContext sets Context field to given value.


### GetAction

`func (o *UIRenderEvent) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UIRenderEvent) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UIRenderEvent) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UIRenderEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetEventId

`func (o *UIRenderEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *UIRenderEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *UIRenderEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetInstallationId

`func (o *UIRenderEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *UIRenderEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *UIRenderEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *UIRenderEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *UIRenderEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *UIRenderEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


