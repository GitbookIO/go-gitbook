# ContentKitHStack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Align** | Pointer to **string** |  | [optional] [default to "start"]
**Children** | [**[]ContentKitDescendantElement**](ContentKitDescendantElement.md) |  | 

## Methods

### NewContentKitHStack

`func NewContentKitHStack(type_ string, children []ContentKitDescendantElement, ) *ContentKitHStack`

NewContentKitHStack instantiates a new ContentKitHStack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitHStackWithDefaults

`func NewContentKitHStackWithDefaults() *ContentKitHStack`

NewContentKitHStackWithDefaults instantiates a new ContentKitHStack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitHStack) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitHStack) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitHStack) SetType(v string)`

SetType sets Type field to given value.


### GetAlign

`func (o *ContentKitHStack) GetAlign() string`

GetAlign returns the Align field if non-nil, zero value otherwise.

### GetAlignOk

`func (o *ContentKitHStack) GetAlignOk() (*string, bool)`

GetAlignOk returns a tuple with the Align field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlign

`func (o *ContentKitHStack) SetAlign(v string)`

SetAlign sets Align field to given value.

### HasAlign

`func (o *ContentKitHStack) HasAlign() bool`

HasAlign returns a boolean if a field has been set.

### GetChildren

`func (o *ContentKitHStack) GetChildren() []ContentKitDescendantElement`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitHStack) GetChildrenOk() (*[]ContentKitDescendantElement, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitHStack) SetChildren(v []ContentKitDescendantElement)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


