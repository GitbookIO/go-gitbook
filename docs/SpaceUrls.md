# SpaceUrls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | URL of the space in the API | 
**App** | **string** | URL of the space in the application | 
**Published** | Pointer to **string** | URL of the published version of the space. Only defined when visibility is not \&quot;private.\&quot; | [optional] 
**Public** | Pointer to **string** | URL of the public version of the space. Only defined when visibility is \&quot;public\&quot;. | [optional] 

## Methods

### NewSpaceUrls

`func NewSpaceUrls(location string, app string, ) *SpaceUrls`

NewSpaceUrls instantiates a new SpaceUrls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceUrlsWithDefaults

`func NewSpaceUrlsWithDefaults() *SpaceUrls`

NewSpaceUrlsWithDefaults instantiates a new SpaceUrls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *SpaceUrls) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SpaceUrls) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SpaceUrls) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetApp

`func (o *SpaceUrls) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *SpaceUrls) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *SpaceUrls) SetApp(v string)`

SetApp sets App field to given value.


### GetPublished

`func (o *SpaceUrls) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SpaceUrls) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SpaceUrls) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SpaceUrls) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetPublic

`func (o *SpaceUrls) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *SpaceUrls) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *SpaceUrls) SetPublic(v string)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *SpaceUrls) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


