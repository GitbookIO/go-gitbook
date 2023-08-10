# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;organization\&quot; | 
**Id** | **string** | Unique identifier for the organization | 
**Title** | **string** | Name of the organization | 
**EmailDomains** | **[]string** |  | 
**Type** | [**OrganizationType**](OrganizationType.md) |  | 
**UseCase** | Pointer to [**OrganizationUseCase**](OrganizationUseCase.md) |  | [optional] 
**CommunityType** | Pointer to [**OrganizationCommunityType**](OrganizationCommunityType.md) |  | [optional] 
**Urls** | [**OrganizationUrls**](OrganizationUrls.md) |  | 

## Methods

### NewOrganization

`func NewOrganization(object string, id string, title string, emailDomains []string, type_ OrganizationType, urls OrganizationUrls, ) *Organization`

NewOrganization instantiates a new Organization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Organization) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Organization) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Organization) SetObject(v string)`

SetObject sets Object field to given value.


### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Organization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *Organization) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Organization) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Organization) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEmailDomains

`func (o *Organization) GetEmailDomains() []string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *Organization) GetEmailDomainsOk() (*[]string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *Organization) SetEmailDomains(v []string)`

SetEmailDomains sets EmailDomains field to given value.


### GetType

`func (o *Organization) GetType() OrganizationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Organization) GetTypeOk() (*OrganizationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Organization) SetType(v OrganizationType)`

SetType sets Type field to given value.


### GetUseCase

`func (o *Organization) GetUseCase() OrganizationUseCase`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *Organization) GetUseCaseOk() (*OrganizationUseCase, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *Organization) SetUseCase(v OrganizationUseCase)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *Organization) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### GetCommunityType

`func (o *Organization) GetCommunityType() OrganizationCommunityType`

GetCommunityType returns the CommunityType field if non-nil, zero value otherwise.

### GetCommunityTypeOk

`func (o *Organization) GetCommunityTypeOk() (*OrganizationCommunityType, bool)`

GetCommunityTypeOk returns a tuple with the CommunityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityType

`func (o *Organization) SetCommunityType(v OrganizationCommunityType)`

SetCommunityType sets CommunityType field to given value.

### HasCommunityType

`func (o *Organization) HasCommunityType() bool`

HasCommunityType returns a boolean if a field has been set.

### GetUrls

`func (o *Organization) GetUrls() OrganizationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Organization) GetUrlsOk() (*OrganizationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Organization) SetUrls(v OrganizationUrls)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


