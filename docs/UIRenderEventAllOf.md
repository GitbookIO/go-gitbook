# UIRenderEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**FetchEventAllOfAuth**](FetchEventAllOfAuth.md) |  | [optional] 
**Type** | **string** |  | 
**ComponentId** | **string** |  | 
**Props** | **map[string]interface{}** | Properties to render the UI. | 
**State** | Pointer to **map[string]interface{}** | State of the UI. | [optional] 
**Context** | [**ContentKitContext**](ContentKitContext.md) |  | 
**Action** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUIRenderEventAllOf

`func NewUIRenderEventAllOf(type_ string, componentId string, props map[string]interface{}, context ContentKitContext, ) *UIRenderEventAllOf`

NewUIRenderEventAllOf instantiates a new UIRenderEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIRenderEventAllOfWithDefaults

`func NewUIRenderEventAllOfWithDefaults() *UIRenderEventAllOf`

NewUIRenderEventAllOfWithDefaults instantiates a new UIRenderEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *UIRenderEventAllOf) GetAuth() FetchEventAllOfAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UIRenderEventAllOf) GetAuthOk() (*FetchEventAllOfAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UIRenderEventAllOf) SetAuth(v FetchEventAllOfAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UIRenderEventAllOf) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetType

`func (o *UIRenderEventAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UIRenderEventAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UIRenderEventAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetComponentId

`func (o *UIRenderEventAllOf) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *UIRenderEventAllOf) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *UIRenderEventAllOf) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetProps

`func (o *UIRenderEventAllOf) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *UIRenderEventAllOf) GetPropsOk() (*map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *UIRenderEventAllOf) SetProps(v map[string]interface{})`

SetProps sets Props field to given value.


### GetState

`func (o *UIRenderEventAllOf) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UIRenderEventAllOf) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UIRenderEventAllOf) SetState(v map[string]interface{})`

SetState sets State field to given value.

### HasState

`func (o *UIRenderEventAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetContext

`func (o *UIRenderEventAllOf) GetContext() ContentKitContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UIRenderEventAllOf) GetContextOk() (*ContentKitContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UIRenderEventAllOf) SetContext(v ContentKitContext)`

SetContext sets Context field to given value.


### GetAction

`func (o *UIRenderEventAllOf) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UIRenderEventAllOf) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UIRenderEventAllOf) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *UIRenderEventAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


