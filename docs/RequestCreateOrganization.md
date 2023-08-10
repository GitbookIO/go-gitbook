# RequestCreateOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Name of the organization | 
**EmailDomains** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to [**OrganizationType**](OrganizationType.md) |  | [optional] 
**UseCase** | Pointer to [**OrganizationUseCase**](OrganizationUseCase.md) |  | [optional] 

## Methods

### NewRequestCreateOrganization

`func NewRequestCreateOrganization(title string, ) *RequestCreateOrganization`

NewRequestCreateOrganization instantiates a new RequestCreateOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreateOrganizationWithDefaults

`func NewRequestCreateOrganizationWithDefaults() *RequestCreateOrganization`

NewRequestCreateOrganizationWithDefaults instantiates a new RequestCreateOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *RequestCreateOrganization) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestCreateOrganization) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestCreateOrganization) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEmailDomains

`func (o *RequestCreateOrganization) GetEmailDomains() []string`

GetEmailDomains returns the EmailDomains field if non-nil, zero value otherwise.

### GetEmailDomainsOk

`func (o *RequestCreateOrganization) GetEmailDomainsOk() (*[]string, bool)`

GetEmailDomainsOk returns a tuple with the EmailDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailDomains

`func (o *RequestCreateOrganization) SetEmailDomains(v []string)`

SetEmailDomains sets EmailDomains field to given value.

### HasEmailDomains

`func (o *RequestCreateOrganization) HasEmailDomains() bool`

HasEmailDomains returns a boolean if a field has been set.

### GetType

`func (o *RequestCreateOrganization) GetType() OrganizationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestCreateOrganization) GetTypeOk() (*OrganizationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestCreateOrganization) SetType(v OrganizationType)`

SetType sets Type field to given value.

### HasType

`func (o *RequestCreateOrganization) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUseCase

`func (o *RequestCreateOrganization) GetUseCase() OrganizationUseCase`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *RequestCreateOrganization) GetUseCaseOk() (*OrganizationUseCase, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *RequestCreateOrganization) SetUseCase(v OrganizationUseCase)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *RequestCreateOrganization) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


