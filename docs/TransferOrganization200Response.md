# TransferOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | **string** | The unique id of the collection created in the target organization containing the content of the source organization. | 
**NewSourceHostname** | Pointer to **string** | The new hostname if the source organization needed to change hostname. | [optional] 

## Methods

### NewTransferOrganization200Response

`func NewTransferOrganization200Response(collection string, ) *TransferOrganization200Response`

NewTransferOrganization200Response instantiates a new TransferOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferOrganization200ResponseWithDefaults

`func NewTransferOrganization200ResponseWithDefaults() *TransferOrganization200Response`

NewTransferOrganization200ResponseWithDefaults instantiates a new TransferOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *TransferOrganization200Response) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *TransferOrganization200Response) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *TransferOrganization200Response) SetCollection(v string)`

SetCollection sets Collection field to given value.


### GetNewSourceHostname

`func (o *TransferOrganization200Response) GetNewSourceHostname() string`

GetNewSourceHostname returns the NewSourceHostname field if non-nil, zero value otherwise.

### GetNewSourceHostnameOk

`func (o *TransferOrganization200Response) GetNewSourceHostnameOk() (*string, bool)`

GetNewSourceHostnameOk returns a tuple with the NewSourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceHostname

`func (o *TransferOrganization200Response) SetNewSourceHostname(v string)`

SetNewSourceHostname sets NewSourceHostname field to given value.

### HasNewSourceHostname

`func (o *TransferOrganization200Response) HasNewSourceHostname() bool`

HasNewSourceHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


