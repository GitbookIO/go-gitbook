# TransferOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | The unique id of the organization to transfer into the current one. | 
**DefaultOrgRole** | Pointer to [**TransferOrganizationRequestDefaultOrgRole**](TransferOrganizationRequestDefaultOrgRole.md) |  | [optional] 

## Methods

### NewTransferOrganizationRequest

`func NewTransferOrganizationRequest(source string, ) *TransferOrganizationRequest`

NewTransferOrganizationRequest instantiates a new TransferOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferOrganizationRequestWithDefaults

`func NewTransferOrganizationRequestWithDefaults() *TransferOrganizationRequest`

NewTransferOrganizationRequestWithDefaults instantiates a new TransferOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *TransferOrganizationRequest) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TransferOrganizationRequest) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TransferOrganizationRequest) SetSource(v string)`

SetSource sets Source field to given value.


### GetDefaultOrgRole

`func (o *TransferOrganizationRequest) GetDefaultOrgRole() TransferOrganizationRequestDefaultOrgRole`

GetDefaultOrgRole returns the DefaultOrgRole field if non-nil, zero value otherwise.

### GetDefaultOrgRoleOk

`func (o *TransferOrganizationRequest) GetDefaultOrgRoleOk() (*TransferOrganizationRequestDefaultOrgRole, bool)`

GetDefaultOrgRoleOk returns a tuple with the DefaultOrgRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrgRole

`func (o *TransferOrganizationRequest) SetDefaultOrgRole(v TransferOrganizationRequestDefaultOrgRole)`

SetDefaultOrgRole sets DefaultOrgRole field to given value.

### HasDefaultOrgRole

`func (o *TransferOrganizationRequest) HasDefaultOrgRole() bool`

HasDefaultOrgRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


