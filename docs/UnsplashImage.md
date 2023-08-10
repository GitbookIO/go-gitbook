# UnsplashImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Id** | **string** |  | 
**Description** | **string** |  | 
**DownloadLocation** | **string** |  | 
**Urls** | [**UnsplashImageUrls**](UnsplashImageUrls.md) |  | 
**Author** | [**UnsplashImageAuthor**](UnsplashImageAuthor.md) |  | 

## Methods

### NewUnsplashImage

`func NewUnsplashImage(kind string, id string, description string, downloadLocation string, urls UnsplashImageUrls, author UnsplashImageAuthor, ) *UnsplashImage`

NewUnsplashImage instantiates a new UnsplashImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsplashImageWithDefaults

`func NewUnsplashImageWithDefaults() *UnsplashImage`

NewUnsplashImageWithDefaults instantiates a new UnsplashImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *UnsplashImage) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UnsplashImage) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UnsplashImage) SetKind(v string)`

SetKind sets Kind field to given value.


### GetId

`func (o *UnsplashImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnsplashImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnsplashImage) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *UnsplashImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UnsplashImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UnsplashImage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDownloadLocation

`func (o *UnsplashImage) GetDownloadLocation() string`

GetDownloadLocation returns the DownloadLocation field if non-nil, zero value otherwise.

### GetDownloadLocationOk

`func (o *UnsplashImage) GetDownloadLocationOk() (*string, bool)`

GetDownloadLocationOk returns a tuple with the DownloadLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLocation

`func (o *UnsplashImage) SetDownloadLocation(v string)`

SetDownloadLocation sets DownloadLocation field to given value.


### GetUrls

`func (o *UnsplashImage) GetUrls() UnsplashImageUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *UnsplashImage) GetUrlsOk() (*UnsplashImageUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *UnsplashImage) SetUrls(v UnsplashImageUrls)`

SetUrls sets Urls field to given value.


### GetAuthor

`func (o *UnsplashImage) GetAuthor() UnsplashImageAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *UnsplashImage) GetAuthorOk() (*UnsplashImageAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *UnsplashImage) SetAuthor(v UnsplashImageAuthor)`

SetAuthor sets Author field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


