# RequestInviteUsersToOrganizationEmailsInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email address of the user to invite as a member | 
**Role** | [**MemberRoleOrGuest**](MemberRoleOrGuest.md) |  | 

## Methods

### NewRequestInviteUsersToOrganizationEmailsInnerOneOf

`func NewRequestInviteUsersToOrganizationEmailsInnerOneOf(email string, role MemberRoleOrGuest, ) *RequestInviteUsersToOrganizationEmailsInnerOneOf`

NewRequestInviteUsersToOrganizationEmailsInnerOneOf instantiates a new RequestInviteUsersToOrganizationEmailsInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestInviteUsersToOrganizationEmailsInnerOneOfWithDefaults

`func NewRequestInviteUsersToOrganizationEmailsInnerOneOfWithDefaults() *RequestInviteUsersToOrganizationEmailsInnerOneOf`

NewRequestInviteUsersToOrganizationEmailsInnerOneOfWithDefaults instantiates a new RequestInviteUsersToOrganizationEmailsInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetRole

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) GetRole() MemberRoleOrGuest`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) GetRoleOk() (*MemberRoleOrGuest, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RequestInviteUsersToOrganizationEmailsInnerOneOf) SetRole(v MemberRoleOrGuest)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


