# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Name** | **string** | Unique named identifier for the integration | 
**Title** | **string** | Title of the integration | 
**Description** | Pointer to **string** | Description of the integration | [optional] 
**Summary** | Pointer to **string** | Long form markdown summary of the integration | [optional] 
**PreviewImages** | **[]string** | URLs of images to showcase the integration | 
**Visibility** | [**IntegrationVisibility**](IntegrationVisibility.md) |  | 
**Scopes** | [**[]IntegrationScope**](IntegrationScope.md) | Permissions that should be granted to the integration | 
**Categories** | [**[]IntegrationCategory**](IntegrationCategory.md) | Categories for which the integration is listed in the marketplace | 
**Blocks** | Pointer to [**[]IntegrationBlock**](IntegrationBlock.md) | Custom blocks defined by this integration. | [optional] 
**Configurations** | Pointer to [**IntegrationConfigurations**](IntegrationConfigurations.md) |  | [optional] 
**ExternalLinks** | [**[]IntegrationExternalLinksInner**](IntegrationExternalLinksInner.md) | External urls configured by the developer of the integration | 
**Urls** | [**IntegrationUrls**](IntegrationUrls.md) |  | 
**ContentSecurityPolicy** | Pointer to [**IntegrationContentSecurityPolicy**](IntegrationContentSecurityPolicy.md) |  | [optional] 

## Methods

### NewIntegration

`func NewIntegration(object string, name string, title string, previewImages []string, visibility IntegrationVisibility, scopes []IntegrationScope, categories []IntegrationCategory, externalLinks []IntegrationExternalLinksInner, urls IntegrationUrls, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Integration) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Integration) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Integration) SetObject(v string)`

SetObject sets Object field to given value.


### GetName

`func (o *Integration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Integration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Integration) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Integration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Integration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Integration) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Integration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Integration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Integration) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Integration) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSummary

`func (o *Integration) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Integration) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Integration) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Integration) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPreviewImages

`func (o *Integration) GetPreviewImages() []string`

GetPreviewImages returns the PreviewImages field if non-nil, zero value otherwise.

### GetPreviewImagesOk

`func (o *Integration) GetPreviewImagesOk() (*[]string, bool)`

GetPreviewImagesOk returns a tuple with the PreviewImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImages

`func (o *Integration) SetPreviewImages(v []string)`

SetPreviewImages sets PreviewImages field to given value.


### GetVisibility

`func (o *Integration) GetVisibility() IntegrationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Integration) GetVisibilityOk() (*IntegrationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Integration) SetVisibility(v IntegrationVisibility)`

SetVisibility sets Visibility field to given value.


### GetScopes

`func (o *Integration) GetScopes() []IntegrationScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *Integration) GetScopesOk() (*[]IntegrationScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *Integration) SetScopes(v []IntegrationScope)`

SetScopes sets Scopes field to given value.


### GetCategories

`func (o *Integration) GetCategories() []IntegrationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Integration) GetCategoriesOk() (*[]IntegrationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Integration) SetCategories(v []IntegrationCategory)`

SetCategories sets Categories field to given value.


### GetBlocks

`func (o *Integration) GetBlocks() []IntegrationBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *Integration) GetBlocksOk() (*[]IntegrationBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *Integration) SetBlocks(v []IntegrationBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *Integration) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetConfigurations

`func (o *Integration) GetConfigurations() IntegrationConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *Integration) GetConfigurationsOk() (*IntegrationConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *Integration) SetConfigurations(v IntegrationConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *Integration) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetExternalLinks

`func (o *Integration) GetExternalLinks() []IntegrationExternalLinksInner`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *Integration) GetExternalLinksOk() (*[]IntegrationExternalLinksInner, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *Integration) SetExternalLinks(v []IntegrationExternalLinksInner)`

SetExternalLinks sets ExternalLinks field to given value.


### GetUrls

`func (o *Integration) GetUrls() IntegrationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *Integration) GetUrlsOk() (*IntegrationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *Integration) SetUrls(v IntegrationUrls)`

SetUrls sets Urls field to given value.


### GetContentSecurityPolicy

`func (o *Integration) GetContentSecurityPolicy() IntegrationContentSecurityPolicy`

GetContentSecurityPolicy returns the ContentSecurityPolicy field if non-nil, zero value otherwise.

### GetContentSecurityPolicyOk

`func (o *Integration) GetContentSecurityPolicyOk() (*IntegrationContentSecurityPolicy, bool)`

GetContentSecurityPolicyOk returns a tuple with the ContentSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSecurityPolicy

`func (o *Integration) SetContentSecurityPolicy(v IntegrationContentSecurityPolicy)`

SetContentSecurityPolicy sets ContentSecurityPolicy field to given value.

### HasContentSecurityPolicy

`func (o *Integration) HasContentSecurityPolicy() bool`

HasContentSecurityPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


