# ApiInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Current release of GitBook | 
**Build** | **string** | Date of the latest release in ISO format | 

## Methods

### NewApiInformation

`func NewApiInformation(version string, build string, ) *ApiInformation`

NewApiInformation instantiates a new ApiInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInformationWithDefaults

`func NewApiInformationWithDefaults() *ApiInformation`

NewApiInformationWithDefaults instantiates a new ApiInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ApiInformation) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiInformation) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiInformation) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetBuild

`func (o *ApiInformation) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ApiInformation) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ApiInformation) SetBuild(v string)`

SetBuild sets Build field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


