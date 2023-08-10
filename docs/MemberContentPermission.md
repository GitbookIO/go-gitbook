# MemberContentPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | [**MemberRole**](MemberRole.md) |  | 
**Space** | [**Space**](Space.md) |  | 

## Methods

### NewMemberContentPermission

`func NewMemberContentPermission(permission MemberRole, space Space, ) *MemberContentPermission`

NewMemberContentPermission instantiates a new MemberContentPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberContentPermissionWithDefaults

`func NewMemberContentPermissionWithDefaults() *MemberContentPermission`

NewMemberContentPermissionWithDefaults instantiates a new MemberContentPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *MemberContentPermission) GetPermission() MemberRole`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *MemberContentPermission) GetPermissionOk() (*MemberRole, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *MemberContentPermission) SetPermission(v MemberRole)`

SetPermission sets Permission field to given value.


### GetSpace

`func (o *MemberContentPermission) GetSpace() Space`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *MemberContentPermission) GetSpaceOk() (*Space, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *MemberContentPermission) SetSpace(v Space)`

SetSpace sets Space field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


