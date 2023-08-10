# CreateOrganizationTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Members** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateOrganizationTeamRequest

`func NewCreateOrganizationTeamRequest(title string, ) *CreateOrganizationTeamRequest`

NewCreateOrganizationTeamRequest instantiates a new CreateOrganizationTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationTeamRequestWithDefaults

`func NewCreateOrganizationTeamRequestWithDefaults() *CreateOrganizationTeamRequest`

NewCreateOrganizationTeamRequestWithDefaults instantiates a new CreateOrganizationTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateOrganizationTeamRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateOrganizationTeamRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateOrganizationTeamRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMembers

`func (o *CreateOrganizationTeamRequest) GetMembers() []string`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreateOrganizationTeamRequest) GetMembersOk() (*[]string, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreateOrganizationTeamRequest) SetMembers(v []string)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *CreateOrganizationTeamRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


