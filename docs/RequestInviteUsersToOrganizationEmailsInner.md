# RequestInviteUsersToOrganizationEmailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the user to invite as a member | 
**Role** | [**MemberRoleOrGuest**](MemberRoleOrGuest.md) |  | 

## Methods

### NewRequestInviteUsersToOrganizationEmailsInner

`func NewRequestInviteUsersToOrganizationEmailsInner(email string, role MemberRoleOrGuest, ) *RequestInviteUsersToOrganizationEmailsInner`

NewRequestInviteUsersToOrganizationEmailsInner instantiates a new RequestInviteUsersToOrganizationEmailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestInviteUsersToOrganizationEmailsInnerWithDefaults

`func NewRequestInviteUsersToOrganizationEmailsInnerWithDefaults() *RequestInviteUsersToOrganizationEmailsInner`

NewRequestInviteUsersToOrganizationEmailsInnerWithDefaults instantiates a new RequestInviteUsersToOrganizationEmailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RequestInviteUsersToOrganizationEmailsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestInviteUsersToOrganizationEmailsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestInviteUsersToOrganizationEmailsInner) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *RequestInviteUsersToOrganizationEmailsInner) GetRole() MemberRoleOrGuest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RequestInviteUsersToOrganizationEmailsInner) GetRoleOk() (*MemberRoleOrGuest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RequestInviteUsersToOrganizationEmailsInner) SetRole(v MemberRoleOrGuest)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


