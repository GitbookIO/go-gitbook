# CreateOrganizationCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | [**CustomFieldType**](CustomFieldType.md) |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateOrganizationCustomFieldRequest

`func NewCreateOrganizationCustomFieldRequest(name string, type_ CustomFieldType, ) *CreateOrganizationCustomFieldRequest`

NewCreateOrganizationCustomFieldRequest instantiates a new CreateOrganizationCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationCustomFieldRequestWithDefaults

`func NewCreateOrganizationCustomFieldRequestWithDefaults() *CreateOrganizationCustomFieldRequest`

NewCreateOrganizationCustomFieldRequestWithDefaults instantiates a new CreateOrganizationCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrganizationCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrganizationCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrganizationCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateOrganizationCustomFieldRequest) GetType() CustomFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateOrganizationCustomFieldRequest) GetTypeOk() (*CustomFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateOrganizationCustomFieldRequest) SetType(v CustomFieldType)`

SetType sets Type field to given value.


### GetTitle

`func (o *CreateOrganizationCustomFieldRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateOrganizationCustomFieldRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateOrganizationCustomFieldRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateOrganizationCustomFieldRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CreateOrganizationCustomFieldRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrganizationCustomFieldRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrganizationCustomFieldRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateOrganizationCustomFieldRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlaceholder

`func (o *CreateOrganizationCustomFieldRequest) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *CreateOrganizationCustomFieldRequest) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *CreateOrganizationCustomFieldRequest) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *CreateOrganizationCustomFieldRequest) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetOptions

`func (o *CreateOrganizationCustomFieldRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateOrganizationCustomFieldRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateOrganizationCustomFieldRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateOrganizationCustomFieldRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


