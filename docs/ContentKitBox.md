# ContentKitBox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Grow** | Pointer to **float32** | specifies how much of the remaining space in the container should be assigned to the element | [optional] 
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 

## Methods

### NewContentKitBox

`func NewContentKitBox(type_ string, children []ContentKitDescendantElement, ) *ContentKitBox`

NewContentKitBox instantiates a new ContentKitBox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitBoxWithDefaults

`func NewContentKitBoxWithDefaults() *ContentKitBox`

NewContentKitBoxWithDefaults instantiates a new ContentKitBox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitBox) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitBox) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitBox) SetType(v string)`

SetType sets Type field to given value.


### GetGrow

`func (o *ContentKitBox) GetGrow() float32`

GetGrow returns the Grow field if non-nil, zero value otherwise.

### GetGrowOk

`func (o *ContentKitBox) GetGrowOk() (*float32, bool)`

GetGrowOk returns a tuple with the Grow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrow

`func (o *ContentKitBox) SetGrow(v float32)`

SetGrow sets Grow field to given value.

### HasGrow

`func (o *ContentKitBox) HasGrow() bool`

HasGrow returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitBox) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitBox) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitBox) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


