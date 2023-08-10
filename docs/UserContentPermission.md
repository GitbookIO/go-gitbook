# UserContentPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | [**MemberRole**](MemberRole.md) |  | 
**User** | [**User**](User.md) |  | 

## Methods

### NewUserContentPermission

`func NewUserContentPermission(permission MemberRole, user User, ) *UserContentPermission`

NewUserContentPermission instantiates a new UserContentPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserContentPermissionWithDefaults

`func NewUserContentPermissionWithDefaults() *UserContentPermission`

NewUserContentPermissionWithDefaults instantiates a new UserContentPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *UserContentPermission) GetPermission() MemberRole`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *UserContentPermission) GetPermissionOk() (*MemberRole, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *UserContentPermission) SetPermission(v MemberRole)`

SetPermission sets Permission field to given value.


### GetUser

`func (o *UserContentPermission) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserContentPermission) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserContentPermission) SetUser(v User)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


