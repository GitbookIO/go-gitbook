# RequestImportContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL of the content to import. | 
**Source** | [**ImportContentSource**](ImportContentSource.md) |  | 

## Methods

### NewRequestImportContent

`func NewRequestImportContent(url string, source ImportContentSource, ) *RequestImportContent`

NewRequestImportContent instantiates a new RequestImportContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestImportContentWithDefaults

`func NewRequestImportContentWithDefaults() *RequestImportContent`

NewRequestImportContentWithDefaults instantiates a new RequestImportContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RequestImportContent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestImportContent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestImportContent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSource

`func (o *RequestImportContent) GetSource() ImportContentSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RequestImportContent) GetSourceOk() (*ImportContentSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RequestImportContent) SetSource(v ImportContentSource)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


