# ContentKitWebFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AspectRatio** | Pointer to **float32** | Ratio between width and height. Used to size the webframe. | [optional] 
**Source** | [**ContentKitWebFrameSource**](ContentKitWebFrameSource.md) |  | 
**Buttons** | Pointer to [**[]ContentKitButton**](ContentKitButton.md) | Controls button shown as an overlay in a corner of the frame. | [optional] 
**Data** | Pointer to [**map[string]ContentKitWebFrameDataValue**](ContentKitWebFrameDataValue.md) | Data to communicated to the webframe&#39;s content. Each state update will cause the webframe to receive a message. | [optional] 

## Methods

### NewContentKitWebFrame

`func NewContentKitWebFrame(type_ string, source ContentKitWebFrameSource, ) *ContentKitWebFrame`

NewContentKitWebFrame instantiates a new ContentKitWebFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitWebFrameWithDefaults

`func NewContentKitWebFrameWithDefaults() *ContentKitWebFrame`

NewContentKitWebFrameWithDefaults instantiates a new ContentKitWebFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitWebFrame) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitWebFrame) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitWebFrame) SetType(v string)`

SetType sets Type field to given value.


### GetAspectRatio

`func (o *ContentKitWebFrame) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ContentKitWebFrame) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ContentKitWebFrame) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *ContentKitWebFrame) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### GetSource

`func (o *ContentKitWebFrame) GetSource() ContentKitWebFrameSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContentKitWebFrame) GetSourceOk() (*ContentKitWebFrameSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContentKitWebFrame) SetSource(v ContentKitWebFrameSource)`

SetSource sets Source field to given value.


### GetButtons

`func (o *ContentKitWebFrame) GetButtons() []ContentKitButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ContentKitWebFrame) GetButtonsOk() (*[]ContentKitButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ContentKitWebFrame) SetButtons(v []ContentKitButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ContentKitWebFrame) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetData

`func (o *ContentKitWebFrame) GetData() map[string]ContentKitWebFrameDataValue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentKitWebFrame) GetDataOk() (*map[string]ContentKitWebFrameDataValue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentKitWebFrame) SetData(v map[string]ContentKitWebFrameDataValue)`

SetData sets Data field to given value.

### HasData

`func (o *ContentKitWebFrame) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


