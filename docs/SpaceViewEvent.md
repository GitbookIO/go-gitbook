# SpaceViewEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for the event. | 
**Type** | **string** |  | 
**InstallationId** | **string** | ID of the integration installation | 
**SpaceId** | **string** | ID of the space | 
**PageId** | Pointer to **string** | Unique identifier of the visited page. | [optional] 
**Visitor** | [**SpaceViewEventAllOfVisitor**](SpaceViewEventAllOfVisitor.md) |  | 
**Url** | **string** | The GitBook content&#39;s URL visited (including URL params). | 
**Referrer** | **string** | The URL of referrer that linked to the page. | 

## Methods

### NewSpaceViewEvent

`func NewSpaceViewEvent(eventId string, type_ string, installationId string, spaceId string, visitor SpaceViewEventAllOfVisitor, url string, referrer string, ) *SpaceViewEvent`

NewSpaceViewEvent instantiates a new SpaceViewEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceViewEventWithDefaults

`func NewSpaceViewEventWithDefaults() *SpaceViewEvent`

NewSpaceViewEventWithDefaults instantiates a new SpaceViewEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *SpaceViewEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SpaceViewEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SpaceViewEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetType

`func (o *SpaceViewEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpaceViewEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpaceViewEvent) SetType(v string)`

SetType sets Type field to given value.


### GetInstallationId

`func (o *SpaceViewEvent) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *SpaceViewEvent) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *SpaceViewEvent) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.


### GetSpaceId

`func (o *SpaceViewEvent) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *SpaceViewEvent) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *SpaceViewEvent) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.


### GetPageId

`func (o *SpaceViewEvent) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *SpaceViewEvent) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *SpaceViewEvent) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *SpaceViewEvent) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetVisitor

`func (o *SpaceViewEvent) GetVisitor() SpaceViewEventAllOfVisitor`

GetVisitor returns the Visitor field if non-nil, zero value otherwise.

### GetVisitorOk

`func (o *SpaceViewEvent) GetVisitorOk() (*SpaceViewEventAllOfVisitor, bool)`

GetVisitorOk returns a tuple with the Visitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitor

`func (o *SpaceViewEvent) SetVisitor(v SpaceViewEventAllOfVisitor)`

SetVisitor sets Visitor field to given value.


### GetUrl

`func (o *SpaceViewEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SpaceViewEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SpaceViewEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetReferrer

`func (o *SpaceViewEvent) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *SpaceViewEvent) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *SpaceViewEvent) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


