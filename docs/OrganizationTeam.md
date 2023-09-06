# OrganizationTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;team\&quot; | 
**Id** | **string** | Unique identifier for the team. | 
**Title** | **string** | Title of the team. | 
**Members** | **int32** | Count of members in this team. | 
**Spaces** | **float32** | Count of spaces this team has access to. | 
**CreatedAt** | **string** |  | 

## Methods

### NewOrganizationTeam

`func NewOrganizationTeam(object string, id string, title string, members int32, spaces float32, createdAt string, ) *OrganizationTeam`

NewOrganizationTeam instantiates a new OrganizationTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationTeamWithDefaults

`func NewOrganizationTeamWithDefaults() *OrganizationTeam`

NewOrganizationTeamWithDefaults instantiates a new OrganizationTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *OrganizationTeam) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OrganizationTeam) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OrganizationTeam) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *OrganizationTeam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationTeam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationTeam) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *OrganizationTeam) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OrganizationTeam) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OrganizationTeam) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetMembers

`func (o *OrganizationTeam) GetMembers() int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationTeam) GetMembersOk() (*int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationTeam) SetMembers(v int32)`

SetMembers sets Members field to given value.


### GetSpaces

`func (o *OrganizationTeam) GetSpaces() float32`

GetSpaces returns the Spaces field if non-nil, zero value otherwise.

### GetSpacesOk

`func (o *OrganizationTeam) GetSpacesOk() (*float32, bool)`

GetSpacesOk returns a tuple with the Spaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaces

`func (o *OrganizationTeam) SetSpaces(v float32)`

SetSpaces sets Spaces field to given value.


### GetCreatedAt

`func (o *OrganizationTeam) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationTeam) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationTeam) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


