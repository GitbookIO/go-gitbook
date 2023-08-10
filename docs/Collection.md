# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;collection\&quot; | 
**Id** | **string** | Unique identifier for the collection | 
**Title** | **string** | Title of the collection | 
**Path** | Pointer to **string** | Path in the published URL | [optional] 
**Visibility** | [**ContentVisibility**](ContentVisibility.md) |  | 
**PublishingType** | Pointer to **string** |  | [optional] 
**PrimarySpace** | Pointer to **string** | ID of the primary space for this collection | [optional] 
**Collection** | Pointer to **string** | ID of the parent collection, if any | [optional] 

## Methods

### NewCollection

`func NewCollection(object string, id string, title string, visibility ContentVisibility, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Collection) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Collection) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Collection) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Collection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Collection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Collection) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Collection) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Collection) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Collection) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPath

`func (o *Collection) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Collection) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Collection) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Collection) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVisibility

`func (o *Collection) GetVisibility() ContentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Collection) GetVisibilityOk() (*ContentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Collection) SetVisibility(v ContentVisibility)`

SetVisibility sets Visibility field to given value.


### GetPublishingType

`func (o *Collection) GetPublishingType() string`

GetPublishingType returns the PublishingType field if non-nil, zero value otherwise.

### GetPublishingTypeOk

`func (o *Collection) GetPublishingTypeOk() (*string, bool)`

GetPublishingTypeOk returns a tuple with the PublishingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishingType

`func (o *Collection) SetPublishingType(v string)`

SetPublishingType sets PublishingType field to given value.

### HasPublishingType

`func (o *Collection) HasPublishingType() bool`

HasPublishingType returns a boolean if a field has been set.

### GetPrimarySpace

`func (o *Collection) GetPrimarySpace() string`

GetPrimarySpace returns the PrimarySpace field if non-nil, zero value otherwise.

### GetPrimarySpaceOk

`func (o *Collection) GetPrimarySpaceOk() (*string, bool)`

GetPrimarySpaceOk returns a tuple with the PrimarySpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimarySpace

`func (o *Collection) SetPrimarySpace(v string)`

SetPrimarySpace sets PrimarySpace field to given value.

### HasPrimarySpace

`func (o *Collection) HasPrimarySpace() bool`

HasPrimarySpace returns a boolean if a field has been set.

### GetCollection

`func (o *Collection) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *Collection) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *Collection) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *Collection) HasCollection() bool`

HasCollection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


