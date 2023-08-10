# Space

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;space\&quot; | 
**Id** | **string** | Unique identifier for the space | 
**Type** | [**SpaceType**](SpaceType.md) |  | 
**Title** | **string** | Title of the space | 
**Visibility** | [**ContentVisibility**](ContentVisibility.md) |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**Urls** | [**SpaceUrls**](SpaceUrls.md) |  | 
**Organization** | Pointer to **string** | ID of the organization owning this space | [optional] 
**User** | Pointer to **string** | ID of the user owning this space | [optional] 
**Private** | Pointer to **string** | If the space is part of the private content of a user, the property is equal to the ID of the user owning it | [optional] 
**Parent** | Pointer to **string** | ID of the parent collection. | [optional] 

## Methods

### NewSpace

`func NewSpace(object string, id string, type_ SpaceType, title string, visibility ContentVisibility, createdAt string, updatedAt string, urls SpaceUrls, ) *Space`

NewSpace instantiates a new Space object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpaceWithDefaults

`func NewSpaceWithDefaults() *Space`

NewSpaceWithDefaults instantiates a new Space object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Space) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Space) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Space) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Space) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Space) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Space) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Space) GetType() SpaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Space) GetTypeOk() (*SpaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Space) SetType(v SpaceType)`

SetType sets Type field to given value.


### GetTitle

`func (o *Space) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Space) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Space) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetVisibility

`func (o *Space) GetVisibility() ContentVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Space) GetVisibilityOk() (*ContentVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Space) SetVisibility(v ContentVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatedAt

`func (o *Space) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Space) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Space) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Space) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Space) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Space) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrls

`func (o *Space) GetUrls() SpaceUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Space) GetUrlsOk() (*SpaceUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Space) SetUrls(v SpaceUrls)`

SetUrls sets Urls field to given value.


### GetOrganization

`func (o *Space) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Space) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Space) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Space) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetUser

`func (o *Space) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Space) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Space) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Space) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPrivate

`func (o *Space) GetPrivate() string`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *Space) GetPrivateOk() (*string, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *Space) SetPrivate(v string)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *Space) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetParent

`func (o *Space) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Space) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Space) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Space) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


