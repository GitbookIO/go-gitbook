# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**Status** | [**IntegrationInstallationStatus**](IntegrationInstallationStatus.md) |  | 
**Previous** | [**SpaceInstallationDeletedEventAllOfPrevious**](SpaceInstallationDeletedEventAllOfPrevious.md) |  | 
**SpaceId** | **string** | ID of the space | 
**PageId** | Pointer to **string** | Unique identifier of the visited page. | [optional] 
**Visitor** | [**SpaceViewEventAllOfVisitor**](SpaceViewEventAllOfVisitor.md) |  | 
**Url** | **string** | The GitBook content&#39;s URL visited (including URL params). | 
**Referrer** | **string** | The URL of referrer that linked to the page. | 
**RevisionId** | **string** | Unique identifier of the new content revision | 
**State** | **map[string]interface{}** | State of the UI. | 
**CommitId** | **string** | Unique identifier for the commit (sha) | 
**PreviousVisibility** | [**ContentVisibility**](ContentVisibility.md) |  | 
**Visibility** | [**ContentVisibility**](ContentVisibility.md) |  | 
**Auth** | Pointer to [**FetchEventAllOfAuth**](FetchEventAllOfAuth.md) |  | [optional] 
**Request** | [**FetchRequest**](FetchRequest.md) |  | 
**ComponentId** | **string** |  | 
**Props** | **map[string]interface{}** | Properties to render the UI. | 
**Context** | [**ContentKitContext**](ContentKitContext.md) |  | 
**Action** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewEvent

`func NewEvent(eventId string, type_ string, installationId string, status IntegrationInstallationStatus, previous SpaceInstallationDeletedEventAllOfPrevious, spaceId string, visitor SpaceViewEventAllOfVisitor, url string, referrer string, revisionId string, state map[string]interface{}, commitId string, previousVisibility ContentVisibility, visibility ContentVisibility, request FetchRequest, componentId string, props map[string]interface{}, context ContentKitContext, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *Event) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Event) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Event) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *Event) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *Event) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *Event) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetStatus

`func (o *Event) GetStatus() IntegrationInstallationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*IntegrationInstallationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v IntegrationInstallationStatus)`

SetStatus sets Status field to given value.


### GetPrevious

`func (o *Event) GetPrevious() SpaceInstallationDeletedEventAllOfPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *Event) GetPreviousOk() (*SpaceInstallationDeletedEventAllOfPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *Event) SetPrevious(v SpaceInstallationDeletedEventAllOfPrevious)`

SetPrevious sets Previous field to given value.


### GetSpaceId

`func (o *Event) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *Event) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *Event) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetPageId

`func (o *Event) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *Event) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *Event) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *Event) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetVisitor

`func (o *Event) GetVisitor() SpaceViewEventAllOfVisitor`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *Event) GetVisitorOk() (*SpaceViewEventAllOfVisitor, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *Event) SetVisitor(v SpaceViewEventAllOfVisitor)`

SetVisitor sets Visitor field to given value.


### GetUrl

`func (o *Event) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Event) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Event) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReferrer

`func (o *Event) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *Event) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *Event) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.


### GetRevisionId

`func (o *Event) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *Event) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *Event) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.


### GetState

`func (o *Event) GetState() map[string]interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Event) GetStateOk() (*map[string]interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Event) SetState(v map[string]interface{})`

SetState sets State field to given value.


### GetCommitId

`func (o *Event) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *Event) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *Event) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### GetPreviousVisibility

`func (o *Event) GetPreviousVisibility() ContentVisibility`

GetPreviousVisibility returns the PreviousVisibility field if non-nil, zero value otherwise.

### GetPreviousVisibilityOk

`func (o *Event) GetPreviousVisibilityOk() (*ContentVisibility, bool)`

GetPreviousVisibilityOk returns a tuple with the PreviousVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousVisibility

`func (o *Event) SetPreviousVisibility(v ContentVisibility)`

SetPreviousVisibility sets PreviousVisibility field to given value.


### GetVisibility

`func (o *Event) GetVisibility() ContentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Event) GetVisibilityOk() (*ContentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Event) SetVisibility(v ContentVisibility)`

SetVisibility sets Visibility field to given value.


### GetAuth

`func (o *Event) GetAuth() FetchEventAllOfAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *Event) GetAuthOk() (*FetchEventAllOfAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *Event) SetAuth(v FetchEventAllOfAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *Event) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetRequest

`func (o *Event) GetRequest() FetchRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Event) GetRequestOk() (*FetchRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Event) SetRequest(v FetchRequest)`

SetRequest sets Request field to given value.


### GetComponentId

`func (o *Event) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *Event) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *Event) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.


### GetProps

`func (o *Event) GetProps() map[string]interface{}`

GetProps returns the Props field if non-nil, zero value otherwise.

### GetPropsOk

`func (o *Event) GetPropsOk() (*map[string]interface{}, bool)`

GetPropsOk returns a tuple with the Props field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProps

`func (o *Event) SetProps(v map[string]interface{})`

SetProps sets Props field to given value.


### GetContext

`func (o *Event) GetContext() ContentKitContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Event) GetContextOk() (*ContentKitContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Event) SetContext(v ContentKitContext)`

SetContext sets Context field to given value.


### GetAction

`func (o *Event) GetAction() map[string]interface{}`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Event) GetActionOk() (*map[string]interface{}, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Event) SetAction(v map[string]interface{})`

SetAction sets Action field to given value.

### HasAction

`func (o *Event) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


