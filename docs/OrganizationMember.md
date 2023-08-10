# OrganizationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;member\&quot; | 
**Id** | **string** | Unique identifier for the user. | 
**Role** | [**MemberRoleOrGuest**](MemberRoleOrGuest.md) |  | 
**User** | [**User**](User.md) |  | 
**Disabled** | **bool** | Whatever the membership of this user is disabled and prevent them from accessing content. | 
**JoinedAt** | **string** |  | 
**LastSeenAt** | Pointer to **string** |  | [optional] 
**Sso** | **bool** | Whether the user can login with SSO. | 
**Spaces** | **float32** |  | 
**Teams** | **float32** |  | 

## Methods

### NewOrganizationMember

`func NewOrganizationMember(object string, id string, role MemberRoleOrGuest, user User, disabled bool, joinedAt string, sso bool, spaces float32, teams float32, ) *OrganizationMember`

NewOrganizationMember instantiates a new OrganizationMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberWithDefaults

`func NewOrganizationMemberWithDefaults() *OrganizationMember`

NewOrganizationMemberWithDefaults instantiates a new OrganizationMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *OrganizationMember) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrganizationMember) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrganizationMember) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *OrganizationMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationMember) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *OrganizationMember) GetRole() MemberRoleOrGuest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMember) GetRoleOk() (*MemberRoleOrGuest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMember) SetRole(v MemberRoleOrGuest)`

SetRole sets Role field to given value.


### GetUser

`func (o *OrganizationMember) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrganizationMember) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrganizationMember) SetUser(v User)`

SetUser sets User field to given value.


### GetDisabled

`func (o *OrganizationMember) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *OrganizationMember) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *OrganizationMember) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetJoinedAt

`func (o *OrganizationMember) GetJoinedAt() string`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *OrganizationMember) GetJoinedAtOk() (*string, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *OrganizationMember) SetJoinedAt(v string)`

SetJoinedAt sets JoinedAt field to given value.


### GetLastSeenAt

`func (o *OrganizationMember) GetLastSeenAt() string`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *OrganizationMember) GetLastSeenAtOk() (*string, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *OrganizationMember) SetLastSeenAt(v string)`

SetLastSeenAt sets LastSeenAt field to given value.

### HasLastSeenAt

`func (o *OrganizationMember) HasLastSeenAt() bool`

HasLastSeenAt returns a boolean if a field has been set.

### GetSso

`func (o *OrganizationMember) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *OrganizationMember) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *OrganizationMember) SetSso(v bool)`

SetSso sets Sso field to given value.


### GetSpaces

`func (o *OrganizationMember) GetSpaces() float32`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *OrganizationMember) GetSpacesOk() (*float32, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *OrganizationMember) SetSpaces(v float32)`

SetSpaces sets Spaces field to given value.


### GetTeams

`func (o *OrganizationMember) GetTeams() float32`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *OrganizationMember) GetTeamsOk() (*float32, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *OrganizationMember) SetTeams(v float32)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


