# OrganizationDirectorySyncGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this group in WorkOS. Not the unique ID from GitBook. | 
**IdpId** | **string** | The identity provider&#39;s unique ID for this group, should be used to generate the team&#39;s unique ID when syncing the groups. | 
**DirectoryId** | **string** | The unique ID of the directory this group is owned by in WorkOS. Is not a unique ID from our database. | 
**Name** | **string** | The name of the group from the identity provider, it should always be set according to the WorkOS documentation. | 
**TeamKey** | Pointer to **string** | The unique ID of the GitBook team already synced to this group, if applicable. | [optional] 

## Methods

### NewOrganizationDirectorySyncGroup

`func NewOrganizationDirectorySyncGroup(id string, idpId string, directoryId string, name string, ) *OrganizationDirectorySyncGroup`

NewOrganizationDirectorySyncGroup instantiates a new OrganizationDirectorySyncGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDirectorySyncGroupWithDefaults

`func NewOrganizationDirectorySyncGroupWithDefaults() *OrganizationDirectorySyncGroup`

NewOrganizationDirectorySyncGroupWithDefaults instantiates a new OrganizationDirectorySyncGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationDirectorySyncGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationDirectorySyncGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationDirectorySyncGroup) SetId(v string)`

SetId sets Id field to given value.


### GetIdpId

`func (o *OrganizationDirectorySyncGroup) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *OrganizationDirectorySyncGroup) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *OrganizationDirectorySyncGroup) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetDirectoryId

`func (o *OrganizationDirectorySyncGroup) GetDirectoryId() string`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *OrganizationDirectorySyncGroup) GetDirectoryIdOk() (*string, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *OrganizationDirectorySyncGroup) SetDirectoryId(v string)`

SetDirectoryId sets DirectoryId field to given value.


### GetName

`func (o *OrganizationDirectorySyncGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDirectorySyncGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDirectorySyncGroup) SetName(v string)`

SetName sets Name field to given value.


### GetTeamKey

`func (o *OrganizationDirectorySyncGroup) GetTeamKey() string`

GetTeamKey returns the TeamKey field if non-nil, zero value otherwise.

### GetTeamKeyOk

`func (o *OrganizationDirectorySyncGroup) GetTeamKeyOk() (*string, bool)`

GetTeamKeyOk returns a tuple with the TeamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamKey

`func (o *OrganizationDirectorySyncGroup) SetTeamKey(v string)`

SetTeamKey sets TeamKey field to given value.

### HasTeamKey

`func (o *OrganizationDirectorySyncGroup) HasTeamKey() bool`

HasTeamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


