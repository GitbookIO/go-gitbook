# ContentKitInlineElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Style** | Pointer to **string** |  | [optional] 
**Children** | [**ContentKitLinkChildren**](ContentKitLinkChildren.md) |  | 
**Source** | [**ContentKitImageSource**](ContentKitImageSource.md) |  | 
**AspectRatio** | **float32** |  | 
**Target** | [**ContentKitLinkTarget**](ContentKitLinkTarget.md) |  | 

## Methods

### NewContentKitInlineElement

`func NewContentKitInlineElement(type_ string, children ContentKitLinkChildren, source ContentKitImageSource, aspectRatio float32, target ContentKitLinkTarget, ) *ContentKitInlineElement`

NewContentKitInlineElement instantiates a new ContentKitInlineElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitInlineElementWithDefaults

`func NewContentKitInlineElementWithDefaults() *ContentKitInlineElement`

NewContentKitInlineElementWithDefaults instantiates a new ContentKitInlineElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitInlineElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitInlineElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitInlineElement) SetType(v string)`

SetType sets Type field to given value.


### GetStyle

`func (o *ContentKitInlineElement) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ContentKitInlineElement) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ContentKitInlineElement) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ContentKitInlineElement) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitInlineElement) GetChildren() ContentKitLinkChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitInlineElement) GetChildrenOk() (*ContentKitLinkChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitInlineElement) SetChildren(v ContentKitLinkChildren)`

SetChildren sets Children field to given value.


### GetSource

`func (o *ContentKitInlineElement) GetSource() ContentKitImageSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ContentKitInlineElement) GetSourceOk() (*ContentKitImageSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ContentKitInlineElement) SetSource(v ContentKitImageSource)`

SetSource sets Source field to given value.


### GetAspectRatio

`func (o *ContentKitInlineElement) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *ContentKitInlineElement) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *ContentKitInlineElement) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.


### GetTarget

`func (o *ContentKitInlineElement) GetTarget() ContentKitLinkTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContentKitInlineElement) GetTargetOk() (*ContentKitLinkTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContentKitInlineElement) SetTarget(v ContentKitLinkTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


