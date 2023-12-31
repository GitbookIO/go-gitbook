# RevisionBaseUrls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | **string** | URL in the application for the revision | 
**Published** | Pointer to **string** | URL of the published version of the revision. Only defined when the space visibility is not \&quot;private.\&quot; | [optional] 
**Public** | Pointer to **string** | URL of the public version of the revision. Only defined when the space visibility is \&quot;public\&quot;. | [optional] 

## Methods

### NewRevisionBaseUrls

`func NewRevisionBaseUrls(app string, ) *RevisionBaseUrls`

NewRevisionBaseUrls instantiates a new RevisionBaseUrls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevisionBaseUrlsWithDefaults

`func NewRevisionBaseUrlsWithDefaults() *RevisionBaseUrls`

NewRevisionBaseUrlsWithDefaults instantiates a new RevisionBaseUrls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *RevisionBaseUrls) GetApp() string`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *RevisionBaseUrls) GetAppOk() (*string, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *RevisionBaseUrls) SetApp(v string)`

SetApp sets App field to given value.


### GetPublished

`func (o *RevisionBaseUrls) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *RevisionBaseUrls) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *RevisionBaseUrls) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *RevisionBaseUrls) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetPublic

`func (o *RevisionBaseUrls) GetPublic() string`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RevisionBaseUrls) GetPublicOk() (*string, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RevisionBaseUrls) SetPublic(v string)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RevisionBaseUrls) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


