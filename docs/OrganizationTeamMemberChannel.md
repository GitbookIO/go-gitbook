# OrganizationTeamMemberChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | **string** |  | 
**Organization** | **string** |  | 
**Team** | Pointer to **string** |  | [optional] 
**Member** | **string** |  | 

## Methods

### NewOrganizationTeamMemberChannel

`func NewOrganizationTeamMemberChannel(channel string, organization string, member string, ) *OrganizationTeamMemberChannel`

NewOrganizationTeamMemberChannel instantiates a new OrganizationTeamMemberChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamMemberChannelWithDefaults

`func NewOrganizationTeamMemberChannelWithDefaults() *OrganizationTeamMemberChannel`

NewOrganizationTeamMemberChannelWithDefaults instantiates a new OrganizationTeamMemberChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *OrganizationTeamMemberChannel) GetChannel() string`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *OrganizationTeamMemberChannel) GetChannelOk() (*string, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *OrganizationTeamMemberChannel) SetChannel(v string)`

SetChannel sets Channel field to given value.


### GetOrganization

`func (o *OrganizationTeamMemberChannel) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrganizationTeamMemberChannel) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrganizationTeamMemberChannel) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetTeam

`func (o *OrganizationTeamMemberChannel) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *OrganizationTeamMemberChannel) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *OrganizationTeamMemberChannel) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *OrganizationTeamMemberChannel) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetMember

`func (o *OrganizationTeamMemberChannel) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OrganizationTeamMemberChannel) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OrganizationTeamMemberChannel) SetMember(v string)`

SetMember sets Member field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


