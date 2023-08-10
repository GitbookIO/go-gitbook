# ContentKitText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Style** | Pointer to **string** |  | [optional] 
**Children** | [**[]ContentKitTextChildrenInner**](ContentKitTextChildrenInner.md) |  | 

## Methods

### NewContentKitText

`func NewContentKitText(type_ string, children []ContentKitTextChildrenInner, ) *ContentKitText`

NewContentKitText instantiates a new ContentKitText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitTextWithDefaults

`func NewContentKitTextWithDefaults() *ContentKitText`

NewContentKitTextWithDefaults instantiates a new ContentKitText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitText) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitText) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitText) SetType(v string)`

SetType sets Type field to given value.


### GetStyle

`func (o *ContentKitText) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ContentKitText) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ContentKitText) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ContentKitText) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitText) GetChildren() []ContentKitTextChildrenInner`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitText) GetChildrenOk() (*[]ContentKitTextChildrenInner, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitText) SetChildren(v []ContentKitTextChildrenInner)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


