# ImportContentResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revision** | **string** | ID of the newly created revision. | 
**ImportedResources** | **float32** | How many resources were imported | 
**TotalResources** | **float32** | How many resources were processed | 

## Methods

### NewImportContentResult

`func NewImportContentResult(revision string, importedResources float32, totalResources float32, ) *ImportContentResult`

NewImportContentResult instantiates a new ImportContentResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportContentResultWithDefaults

`func NewImportContentResultWithDefaults() *ImportContentResult`

NewImportContentResultWithDefaults instantiates a new ImportContentResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevision

`func (o *ImportContentResult) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ImportContentResult) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ImportContentResult) SetRevision(v string)`

SetRevision sets Revision field to given value.


### GetImportedResources

`func (o *ImportContentResult) GetImportedResources() float32`

GetImportedResources returns the ImportedResources field if non-nil, zero value otherwise.

### GetImportedResourcesOk

`func (o *ImportContentResult) GetImportedResourcesOk() (*float32, bool)`

GetImportedResourcesOk returns a tuple with the ImportedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedResources

`func (o *ImportContentResult) SetImportedResources(v float32)`

SetImportedResources sets ImportedResources field to given value.


### GetTotalResources

`func (o *ImportContentResult) GetTotalResources() float32`

GetTotalResources returns the TotalResources field if non-nil, zero value otherwise.

### GetTotalResourcesOk

`func (o *ImportContentResult) GetTotalResourcesOk() (*float32, bool)`

GetTotalResourcesOk returns a tuple with the TotalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResources

`func (o *ImportContentResult) SetTotalResources(v float32)`

SetTotalResources sets TotalResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


