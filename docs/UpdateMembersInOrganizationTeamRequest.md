# UpdateMembersInOrganizationTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Add** | Pointer to **[]string** |  | [optional] 
**Memberships** | Pointer to [**map[string]TeamMember**](TeamMember.md) |  | [optional] 
**Remove** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateMembersInOrganizationTeamRequest

`func NewUpdateMembersInOrganizationTeamRequest() *UpdateMembersInOrganizationTeamRequest`

NewUpdateMembersInOrganizationTeamRequest instantiates a new UpdateMembersInOrganizationTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMembersInOrganizationTeamRequestWithDefaults

`func NewUpdateMembersInOrganizationTeamRequestWithDefaults() *UpdateMembersInOrganizationTeamRequest`

NewUpdateMembersInOrganizationTeamRequestWithDefaults instantiates a new UpdateMembersInOrganizationTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdd

`func (o *UpdateMembersInOrganizationTeamRequest) GetAdd() []string`

GetAdd returns the Add field if non-nil, zero value otherwise.

### GetAddOk

`func (o *UpdateMembersInOrganizationTeamRequest) GetAddOk() (*[]string, bool)`

GetAddOk returns a tuple with the Add field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdd

`func (o *UpdateMembersInOrganizationTeamRequest) SetAdd(v []string)`

SetAdd sets Add field to given value.

### HasAdd

`func (o *UpdateMembersInOrganizationTeamRequest) HasAdd() bool`

HasAdd returns a boolean if a field has been set.

### GetMemberships

`func (o *UpdateMembersInOrganizationTeamRequest) GetMemberships() map[string]TeamMember`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *UpdateMembersInOrganizationTeamRequest) GetMembershipsOk() (*map[string]TeamMember, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *UpdateMembersInOrganizationTeamRequest) SetMemberships(v map[string]TeamMember)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *UpdateMembersInOrganizationTeamRequest) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetRemove

`func (o *UpdateMembersInOrganizationTeamRequest) GetRemove() []string`

GetRemove returns the Remove field if non-nil, zero value otherwise.

### GetRemoveOk

`func (o *UpdateMembersInOrganizationTeamRequest) GetRemoveOk() (*[]string, bool)`

GetRemoveOk returns a tuple with the Remove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemove

`func (o *UpdateMembersInOrganizationTeamRequest) SetRemove(v []string)`

SetRemove sets Remove field to given value.

### HasRemove

`func (o *UpdateMembersInOrganizationTeamRequest) HasRemove() bool`

HasRemove returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


