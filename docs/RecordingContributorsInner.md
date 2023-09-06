# RecordingContributorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** | Type of Object, always equals to \&quot;user\&quot; | 
**Name** | **string** | Unique named identifier for the integration | 
**Version** | **float32** | Version of the integration | 
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
**Id** | **string** | Unique identifier for the user | 
**DisplayName** | **string** | Full name for the user | 
**Email** | Pointer to **string** | Email address of the user | [optional] 
**PhotoURL** | Pointer to **string** | URL of the user&#39;s profile picture | [optional] 

## Methods

### NewRecordingContributorsInner

`func NewRecordingContributorsInner(object string, name string, version float32, title string, previewImages []string, visibility IntegrationVisibility, scopes []IntegrationScope, categories []IntegrationCategory, externalLinks []IntegrationExternalLinksInner, urls IntegrationUrls, id string, displayName string, ) *RecordingContributorsInner`

NewRecordingContributorsInner instantiates a new RecordingContributorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingContributorsInnerWithDefaults

`func NewRecordingContributorsInnerWithDefaults() *RecordingContributorsInner`

NewRecordingContributorsInnerWithDefaults instantiates a new RecordingContributorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *RecordingContributorsInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RecordingContributorsInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RecordingContributorsInner) SetObject(v string)`

SetObject sets Object field to given value.


### GetName

`func (o *RecordingContributorsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordingContributorsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordingContributorsInner) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *RecordingContributorsInner) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RecordingContributorsInner) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RecordingContributorsInner) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetTitle

`func (o *RecordingContributorsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RecordingContributorsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RecordingContributorsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RecordingContributorsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RecordingContributorsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RecordingContributorsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RecordingContributorsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSummary

`func (o *RecordingContributorsInner) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RecordingContributorsInner) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RecordingContributorsInner) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RecordingContributorsInner) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPreviewImages

`func (o *RecordingContributorsInner) GetPreviewImages() []string`

GetPreviewImages returns the PreviewImages field if non-nil, zero value otherwise.

### GetPreviewImagesOk

`func (o *RecordingContributorsInner) GetPreviewImagesOk() (*[]string, bool)`

GetPreviewImagesOk returns a tuple with the PreviewImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImages

`func (o *RecordingContributorsInner) SetPreviewImages(v []string)`

SetPreviewImages sets PreviewImages field to given value.


### GetVisibility

`func (o *RecordingContributorsInner) GetVisibility() IntegrationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RecordingContributorsInner) GetVisibilityOk() (*IntegrationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RecordingContributorsInner) SetVisibility(v IntegrationVisibility)`

SetVisibility sets Visibility field to given value.


### GetScopes

`func (o *RecordingContributorsInner) GetScopes() []IntegrationScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RecordingContributorsInner) GetScopesOk() (*[]IntegrationScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RecordingContributorsInner) SetScopes(v []IntegrationScope)`

SetScopes sets Scopes field to given value.


### GetCategories

`func (o *RecordingContributorsInner) GetCategories() []IntegrationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *RecordingContributorsInner) GetCategoriesOk() (*[]IntegrationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *RecordingContributorsInner) SetCategories(v []IntegrationCategory)`

SetCategories sets Categories field to given value.


### GetBlocks

`func (o *RecordingContributorsInner) GetBlocks() []IntegrationBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *RecordingContributorsInner) GetBlocksOk() (*[]IntegrationBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *RecordingContributorsInner) SetBlocks(v []IntegrationBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *RecordingContributorsInner) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetConfigurations

`func (o *RecordingContributorsInner) GetConfigurations() IntegrationConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *RecordingContributorsInner) GetConfigurationsOk() (*IntegrationConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *RecordingContributorsInner) SetConfigurations(v IntegrationConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *RecordingContributorsInner) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetExternalLinks

`func (o *RecordingContributorsInner) GetExternalLinks() []IntegrationExternalLinksInner`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *RecordingContributorsInner) GetExternalLinksOk() (*[]IntegrationExternalLinksInner, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *RecordingContributorsInner) SetExternalLinks(v []IntegrationExternalLinksInner)`

SetExternalLinks sets ExternalLinks field to given value.


### GetUrls

`func (o *RecordingContributorsInner) GetUrls() IntegrationUrls`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *RecordingContributorsInner) GetUrlsOk() (*IntegrationUrls, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *RecordingContributorsInner) SetUrls(v IntegrationUrls)`

SetUrls sets Urls field to given value.


### GetContentSecurityPolicy

`func (o *RecordingContributorsInner) GetContentSecurityPolicy() IntegrationContentSecurityPolicy`

GetContentSecurityPolicy returns the ContentSecurityPolicy field if non-nil, zero value otherwise.

### GetContentSecurityPolicyOk

`func (o *RecordingContributorsInner) GetContentSecurityPolicyOk() (*IntegrationContentSecurityPolicy, bool)`

GetContentSecurityPolicyOk returns a tuple with the ContentSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSecurityPolicy

`func (o *RecordingContributorsInner) SetContentSecurityPolicy(v IntegrationContentSecurityPolicy)`

SetContentSecurityPolicy sets ContentSecurityPolicy field to given value.

### HasContentSecurityPolicy

`func (o *RecordingContributorsInner) HasContentSecurityPolicy() bool`

HasContentSecurityPolicy returns a boolean if a field has been set.

### GetId

`func (o *RecordingContributorsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordingContributorsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordingContributorsInner) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *RecordingContributorsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RecordingContributorsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RecordingContributorsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEmail

`func (o *RecordingContributorsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RecordingContributorsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RecordingContributorsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *RecordingContributorsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhotoURL

`func (o *RecordingContributorsInner) GetPhotoURL() string`

GetPhotoURL returns the PhotoURL field if non-nil, zero value otherwise.

### GetPhotoURLOk

`func (o *RecordingContributorsInner) GetPhotoURLOk() (*string, bool)`

GetPhotoURLOk returns a tuple with the PhotoURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoURL

`func (o *RecordingContributorsInner) SetPhotoURL(v string)`

SetPhotoURL sets PhotoURL field to given value.

### HasPhotoURL

`func (o *RecordingContributorsInner) HasPhotoURL() bool`

HasPhotoURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


