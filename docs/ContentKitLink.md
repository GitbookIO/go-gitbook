# ContentKitLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Target** | [**ContentKitLinkTarget**](ContentKitLinkTarget.md) |  | 
**Children** | [**ContentKitLinkChildren**](ContentKitLinkChildren.md) |  | 

## Methods

### NewContentKitLink

`func NewContentKitLink(type_ string, target ContentKitLinkTarget, children ContentKitLinkChildren, ) *ContentKitLink`

NewContentKitLink instantiates a new ContentKitLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentKitLinkWithDefaults

`func NewContentKitLinkWithDefaults() *ContentKitLink`

NewContentKitLinkWithDefaults instantiates a new ContentKitLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentKitLink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentKitLink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentKitLink) SetType(v string)`

SetType sets Type field to given value.


### GetTarget

`func (o *ContentKitLink) GetTarget() ContentKitLinkTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ContentKitLink) GetTargetOk() (*ContentKitLinkTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ContentKitLink) SetTarget(v ContentKitLinkTarget)`

SetTarget sets Target field to given value.


### GetChildren

`func (o *ContentKitLink) GetChildren() ContentKitLinkChildren`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ContentKitLink) GetChildrenOk() (*ContentKitLinkChildren, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ContentKitLink) SetChildren(v ContentKitLinkChildren)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


