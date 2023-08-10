# ContentKitImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Source** | [**ContentKitImageSource**](ContentKitImageSource.md) |  | 
**AspectRatio** | **float32** |  | 

## Methods

### NewContentKitImage

`func NewContentKitImage(type_ string, source ContentKitImageSource, aspectRatio float32, ) *ContentKitImage`

NewContentKitImage instantiates a new ContentKitImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitImageWithDefaults

`func NewContentKitImageWithDefaults() *ContentKitImage`

NewContentKitImageWithDefaults instantiates a new ContentKitImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitImage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitImage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitImage) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *ContentKitImage) GetSource() ContentKitImageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContentKitImage) GetSourceOk() (*ContentKitImageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContentKitImage) SetSource(v ContentKitImageSource)`

SetSource sets Source field to given value.


### GetAspectRatio

`func (o *ContentKitImage) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ContentKitImage) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ContentKitImage) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


