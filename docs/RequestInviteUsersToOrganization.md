# RequestInviteUsersToOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | [**[]RequestInviteUsersToOrganizationEmailsInner**](RequestInviteUsersToOrganizationEmailsInner.md) |  | 
**Role** | Pointer to [**MemberRoleOrGuest**](MemberRoleOrGuest.md) |  | [optional] 
**Sso** | Pointer to **bool** | If true, invites the user as an SSO user of the organization. Defaults to false. | [optional] 

## Methods

### NewRequestInviteUsersToOrganization

`func NewRequestInviteUsersToOrganization(emails []RequestInviteUsersToOrganizationEmailsInner, ) *RequestInviteUsersToOrganization`

NewRequestInviteUsersToOrganization instantiates a new RequestInviteUsersToOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestInviteUsersToOrganizationWithDefaults

`func NewRequestInviteUsersToOrganizationWithDefaults() *RequestInviteUsersToOrganization`

NewRequestInviteUsersToOrganizationWithDefaults instantiates a new RequestInviteUsersToOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *RequestInviteUsersToOrganization) GetEmails() []RequestInviteUsersToOrganizationEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *RequestInviteUsersToOrganization) GetEmailsOk() (*[]RequestInviteUsersToOrganizationEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *RequestInviteUsersToOrganization) SetEmails(v []RequestInviteUsersToOrganizationEmailsInner)`

SetEmails sets Emails field to given value.


### GetRole

`func (o *RequestInviteUsersToOrganization) GetRole() MemberRoleOrGuest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RequestInviteUsersToOrganization) GetRoleOk() (*MemberRoleOrGuest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RequestInviteUsersToOrganization) SetRole(v MemberRoleOrGuest)`

SetRole sets Role field to given value.

### HasRole

`func (o *RequestInviteUsersToOrganization) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSso

`func (o *RequestInviteUsersToOrganization) GetSso() bool`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *RequestInviteUsersToOrganization) GetSsoOk() (*bool, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *RequestInviteUsersToOrganization) SetSso(v bool)`

SetSso sets Sso field to given value.

### HasSso

`func (o *RequestInviteUsersToOrganization) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


