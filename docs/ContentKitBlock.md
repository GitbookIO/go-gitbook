# ContentKitBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 
**Controls** | Pointer to [**[]ContentKitBlockControlsInner**](ContentKitBlockControlsInner.md) |  | [optional] 

## Methods

### NewContentKitBlock

`func NewContentKitBlock(type_ string, children []ContentKitDescendantElement, ) *ContentKitBlock`

NewContentKitBlock instantiates a new ContentKitBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitBlockWithDefaults

`func NewContentKitBlockWithDefaults() *ContentKitBlock`

NewContentKitBlockWithDefaults instantiates a new ContentKitBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitBlock) SetType(v string)`

SetType sets Type field to given value.


### GetChildren

`func (o *ContentKitBlock) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitBlock) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitBlock) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.


### GetControls

`func (o *ContentKitBlock) GetControls() []ContentKitBlockControlsInner`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *ContentKitBlock) GetControlsOk() (*[]ContentKitBlockControlsInner, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *ContentKitBlock) SetControls(v []ContentKitBlockControlsInner)`

SetControls sets Controls field to given value.

### HasControls

`func (o *ContentKitBlock) HasControls() bool`

HasControls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


