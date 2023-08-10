# ContentKitVStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Align** | Pointer to **string** |  | [optional] [default to "start"]
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 

## Methods

### NewContentKitVStack

`func NewContentKitVStack(type_ string, children []ContentKitDescendantElement, ) *ContentKitVStack`

NewContentKitVStack instantiates a new ContentKitVStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitVStackWithDefaults

`func NewContentKitVStackWithDefaults() *ContentKitVStack`

NewContentKitVStackWithDefaults instantiates a new ContentKitVStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitVStack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitVStack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitVStack) SetType(v string)`

SetType sets Type field to given value.


### GetAlign

`func (o *ContentKitVStack) GetAlign() string`

GetAlign returns the Align field if non-nil, zero value otherwise.

### GetAlignOk

`func (o *ContentKitVStack) GetAlignOk() (*string, bool)`

GetAlignOk returns a tuple with the Align field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlign

`func (o *ContentKitVStack) SetAlign(v string)`

SetAlign sets Align field to given value.

### HasAlign

`func (o *ContentKitVStack) HasAlign() bool`

HasAlign returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitVStack) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitVStack) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitVStack) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


