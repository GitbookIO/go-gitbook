# RequestPublishIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Icon** | Pointer to **string** | Base64 content of the icon | [optional] 
**Title** | **string** | Title of the integration | 
**Description** | **string** | Description of the integration | 
**Summary** | Pointer to **string** | Long form markdown summary of the integration | [optional] 
**PreviewImages** | Pointer to **[]string** |  | [optional] 
**Visibility** | Pointer to [**IntegrationVisibility**](IntegrationVisibility.md) |  | [optional] 
**Scopes** | [**[]IntegrationScope**](IntegrationScope.md) | Permissions that should be granted to the integration | 
**Categories** | Pointer to [**[]IntegrationCategory**](IntegrationCategory.md) | Categories for which the integration is listed in the marketplace | [optional] 
**Blocks** | Pointer to [**[]IntegrationBlock**](IntegrationBlock.md) | Custom blocks defined by this integration. | [optional] 
**ExternalLinks** | Pointer to [**[]IntegrationExternalLinksInner**](IntegrationExternalLinksInner.md) | External urls configured by the developer of the integration | [optional] 
**Configurations** | Pointer to [**IntegrationConfigurations**](IntegrationConfigurations.md) |  | [optional] 
**Script** | **string** | Content of the script to use | 
**Organization** | Pointer to **string** | The ID or subdomain of the organization under which the integration should be published | [optional] 
**Secrets** | Pointer to **map[string]string** | Secrets stored on the integration and passed at runtime. | [optional] 
**ContentSecurityPolicy** | Pointer to [**IntegrationContentSecurityPolicy**](IntegrationContentSecurityPolicy.md) |  | [optional] 

## Methods

### NewRequestPublishIntegration

`func NewRequestPublishIntegration(title string, description string, scopes []IntegrationScope, script string, ) *RequestPublishIntegration`

NewRequestPublishIntegration instantiates a new RequestPublishIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestPublishIntegrationWithDefaults

`func NewRequestPublishIntegrationWithDefaults() *RequestPublishIntegration`

NewRequestPublishIntegrationWithDefaults instantiates a new RequestPublishIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIcon

`func (o *RequestPublishIntegration) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RequestPublishIntegration) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RequestPublishIntegration) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RequestPublishIntegration) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetTitle

`func (o *RequestPublishIntegration) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RequestPublishIntegration) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RequestPublishIntegration) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *RequestPublishIntegration) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestPublishIntegration) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestPublishIntegration) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSummary

`func (o *RequestPublishIntegration) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RequestPublishIntegration) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RequestPublishIntegration) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *RequestPublishIntegration) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetPreviewImages

`func (o *RequestPublishIntegration) GetPreviewImages() []string`

GetPreviewImages returns the PreviewImages field if non-nil, zero value otherwise.

### GetPreviewImagesOk

`func (o *RequestPublishIntegration) GetPreviewImagesOk() (*[]string, bool)`

GetPreviewImagesOk returns a tuple with the PreviewImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImages

`func (o *RequestPublishIntegration) SetPreviewImages(v []string)`

SetPreviewImages sets PreviewImages field to given value.

### HasPreviewImages

`func (o *RequestPublishIntegration) HasPreviewImages() bool`

HasPreviewImages returns a boolean if a field has been set.

### GetVisibility

`func (o *RequestPublishIntegration) GetVisibility() IntegrationVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *RequestPublishIntegration) GetVisibilityOk() (*IntegrationVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *RequestPublishIntegration) SetVisibility(v IntegrationVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *RequestPublishIntegration) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetScopes

`func (o *RequestPublishIntegration) GetScopes() []IntegrationScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *RequestPublishIntegration) GetScopesOk() (*[]IntegrationScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *RequestPublishIntegration) SetScopes(v []IntegrationScope)`

SetScopes sets Scopes field to given value.


### GetCategories

`func (o *RequestPublishIntegration) GetCategories() []IntegrationCategory`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *RequestPublishIntegration) GetCategoriesOk() (*[]IntegrationCategory, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *RequestPublishIntegration) SetCategories(v []IntegrationCategory)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *RequestPublishIntegration) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetBlocks

`func (o *RequestPublishIntegration) GetBlocks() []IntegrationBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *RequestPublishIntegration) GetBlocksOk() (*[]IntegrationBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *RequestPublishIntegration) SetBlocks(v []IntegrationBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *RequestPublishIntegration) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetExternalLinks

`func (o *RequestPublishIntegration) GetExternalLinks() []IntegrationExternalLinksInner`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *RequestPublishIntegration) GetExternalLinksOk() (*[]IntegrationExternalLinksInner, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *RequestPublishIntegration) SetExternalLinks(v []IntegrationExternalLinksInner)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *RequestPublishIntegration) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetConfigurations

`func (o *RequestPublishIntegration) GetConfigurations() IntegrationConfigurations`

GetConfigurations returns the Configurations field if non-nil, zero value otherwise.

### GetConfigurationsOk

`func (o *RequestPublishIntegration) GetConfigurationsOk() (*IntegrationConfigurations, bool)`

GetConfigurationsOk returns a tuple with the Configurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurations

`func (o *RequestPublishIntegration) SetConfigurations(v IntegrationConfigurations)`

SetConfigurations sets Configurations field to given value.

### HasConfigurations

`func (o *RequestPublishIntegration) HasConfigurations() bool`

HasConfigurations returns a boolean if a field has been set.

### GetScript

`func (o *RequestPublishIntegration) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *RequestPublishIntegration) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *RequestPublishIntegration) SetScript(v string)`

SetScript sets Script field to given value.


### GetOrganization

`func (o *RequestPublishIntegration) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *RequestPublishIntegration) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *RequestPublishIntegration) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *RequestPublishIntegration) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSecrets

`func (o *RequestPublishIntegration) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *RequestPublishIntegration) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *RequestPublishIntegration) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *RequestPublishIntegration) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetContentSecurityPolicy

`func (o *RequestPublishIntegration) GetContentSecurityPolicy() IntegrationContentSecurityPolicy`

GetContentSecurityPolicy returns the ContentSecurityPolicy field if non-nil, zero value otherwise.

### GetContentSecurityPolicyOk

`func (o *RequestPublishIntegration) GetContentSecurityPolicyOk() (*IntegrationContentSecurityPolicy, bool)`

GetContentSecurityPolicyOk returns a tuple with the ContentSecurityPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSecurityPolicy

`func (o *RequestPublishIntegration) SetContentSecurityPolicy(v IntegrationContentSecurityPolicy)`

SetContentSecurityPolicy sets ContentSecurityPolicy field to given value.

### HasContentSecurityPolicy

`func (o *RequestPublishIntegration) HasContentSecurityPolicy() bool`

HasContentSecurityPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


