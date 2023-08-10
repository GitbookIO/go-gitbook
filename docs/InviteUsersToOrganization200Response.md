# InviteUsersToOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** |  | 
**Invited** | **float32** | The number of users who were added to the organization | 
**FailedSSOEmails** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInviteUsersToOrganization200Response

`func NewInviteUsersToOrganization200Response(users []string, invited float32, ) *InviteUsersToOrganization200Response`

NewInviteUsersToOrganization200Response instantiates a new InviteUsersToOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInviteUsersToOrganization200ResponseWithDefaults

`func NewInviteUsersToOrganization200ResponseWithDefaults() *InviteUsersToOrganization200Response`

NewInviteUsersToOrganization200ResponseWithDefaults instantiates a new InviteUsersToOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *InviteUsersToOrganization200Response) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *InviteUsersToOrganization200Response) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *InviteUsersToOrganization200Response) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetInvited

`func (o *InviteUsersToOrganization200Response) GetInvited() float32`

GetInvited returns the Invited field if non-nil, zero value otherwise.

### GetInvitedOk

`func (o *InviteUsersToOrganization200Response) GetInvitedOk() (*float32, bool)`

GetInvitedOk returns a tuple with the Invited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvited

`func (o *InviteUsersToOrganization200Response) SetInvited(v float32)`

SetInvited sets Invited field to given value.


### GetFailedSSOEmails

`func (o *InviteUsersToOrganization200Response) GetFailedSSOEmails() []string`

GetFailedSSOEmails returns the FailedSSOEmails field if non-nil, zero value otherwise.

### GetFailedSSOEmailsOk

`func (o *InviteUsersToOrganization200Response) GetFailedSSOEmailsOk() (*[]string, bool)`

GetFailedSSOEmailsOk returns a tuple with the FailedSSOEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedSSOEmails

`func (o *InviteUsersToOrganization200Response) SetFailedSSOEmails(v []string)`

SetFailedSSOEmails sets FailedSSOEmails field to given value.

### HasFailedSSOEmails

`func (o *InviteUsersToOrganization200Response) HasFailedSSOEmails() bool`

HasFailedSSOEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


