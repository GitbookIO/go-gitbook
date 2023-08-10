# IntegrationContentSecurityPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUri** | Pointer to **string** |  | [optional] 
**BlockAllMixedContent** | Pointer to **string** |  | [optional] 
**ChildSrc** | Pointer to **string** |  | [optional] 
**ConnectSrc** | Pointer to **string** |  | [optional] 
**DefaultSrc** | Pointer to **string** |  | [optional] 
**FontSrc** | Pointer to **string** |  | [optional] 
**FormAction** | Pointer to **string** |  | [optional] 
**FrameAncestors** | Pointer to **string** |  | [optional] 
**FrameSrc** | Pointer to **string** |  | [optional] 
**ImgSrc** | Pointer to **string** |  | [optional] 
**ManifestSrc** | Pointer to **string** |  | [optional] 
**MediaSrc** | Pointer to **string** |  | [optional] 
**NavigateTo** | Pointer to **string** |  | [optional] 
**ObjectSrc** | Pointer to **string** |  | [optional] 
**PluginTypes** | Pointer to **string** |  | [optional] 
**PrefetchSrc** | Pointer to **string** |  | [optional] 
**Referrer** | Pointer to **string** |  | [optional] 
**ReportTo** | Pointer to **string** |  | [optional] 
**ReportUri** | Pointer to **string** |  | [optional] 
**RequireSriFor** | Pointer to **string** |  | [optional] 
**RequireTrustedTypesFor** | Pointer to **string** |  | [optional] 
**Sandbox** | Pointer to **string** |  | [optional] 
**ScriptSrc** | Pointer to **string** |  | [optional] 
**ScriptSrcAttr** | Pointer to **string** |  | [optional] 
**ScriptSrcElem** | Pointer to **string** |  | [optional] 
**StyleSrc** | Pointer to **string** |  | [optional] 
**StyleSrcAttr** | Pointer to **string** |  | [optional] 
**StyleSrcElem** | Pointer to **string** |  | [optional] 
**TrustedTypes** | Pointer to **string** |  | [optional] 
**UpgradeInsecureRequests** | Pointer to **string** |  | [optional] 
**WorkerSrc** | Pointer to **string** |  | [optional] 

## Methods

### NewIntegrationContentSecurityPolicy

`func NewIntegrationContentSecurityPolicy() *IntegrationContentSecurityPolicy`

NewIntegrationContentSecurityPolicy instantiates a new IntegrationContentSecurityPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationContentSecurityPolicyWithDefaults

`func NewIntegrationContentSecurityPolicyWithDefaults() *IntegrationContentSecurityPolicy`

NewIntegrationContentSecurityPolicyWithDefaults instantiates a new IntegrationContentSecurityPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUri

`func (o *IntegrationContentSecurityPolicy) GetBaseUri() string`

GetBaseUri returns the BaseUri field if non-nil, zero value otherwise.

### GetBaseUriOk

`func (o *IntegrationContentSecurityPolicy) GetBaseUriOk() (*string, bool)`

GetBaseUriOk returns a tuple with the BaseUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUri

`func (o *IntegrationContentSecurityPolicy) SetBaseUri(v string)`

SetBaseUri sets BaseUri field to given value.

### HasBaseUri

`func (o *IntegrationContentSecurityPolicy) HasBaseUri() bool`

HasBaseUri returns a boolean if a field has been set.

### GetBlockAllMixedContent

`func (o *IntegrationContentSecurityPolicy) GetBlockAllMixedContent() string`

GetBlockAllMixedContent returns the BlockAllMixedContent field if non-nil, zero value otherwise.

### GetBlockAllMixedContentOk

`func (o *IntegrationContentSecurityPolicy) GetBlockAllMixedContentOk() (*string, bool)`

GetBlockAllMixedContentOk returns a tuple with the BlockAllMixedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockAllMixedContent

`func (o *IntegrationContentSecurityPolicy) SetBlockAllMixedContent(v string)`

SetBlockAllMixedContent sets BlockAllMixedContent field to given value.

### HasBlockAllMixedContent

`func (o *IntegrationContentSecurityPolicy) HasBlockAllMixedContent() bool`

HasBlockAllMixedContent returns a boolean if a field has been set.

### GetChildSrc

`func (o *IntegrationContentSecurityPolicy) GetChildSrc() string`

GetChildSrc returns the ChildSrc field if non-nil, zero value otherwise.

### GetChildSrcOk

`func (o *IntegrationContentSecurityPolicy) GetChildSrcOk() (*string, bool)`

GetChildSrcOk returns a tuple with the ChildSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildSrc

`func (o *IntegrationContentSecurityPolicy) SetChildSrc(v string)`

SetChildSrc sets ChildSrc field to given value.

### HasChildSrc

`func (o *IntegrationContentSecurityPolicy) HasChildSrc() bool`

HasChildSrc returns a boolean if a field has been set.

### GetConnectSrc

`func (o *IntegrationContentSecurityPolicy) GetConnectSrc() string`

GetConnectSrc returns the ConnectSrc field if non-nil, zero value otherwise.

### GetConnectSrcOk

`func (o *IntegrationContentSecurityPolicy) GetConnectSrcOk() (*string, bool)`

GetConnectSrcOk returns a tuple with the ConnectSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectSrc

`func (o *IntegrationContentSecurityPolicy) SetConnectSrc(v string)`

SetConnectSrc sets ConnectSrc field to given value.

### HasConnectSrc

`func (o *IntegrationContentSecurityPolicy) HasConnectSrc() bool`

HasConnectSrc returns a boolean if a field has been set.

### GetDefaultSrc

`func (o *IntegrationContentSecurityPolicy) GetDefaultSrc() string`

GetDefaultSrc returns the DefaultSrc field if non-nil, zero value otherwise.

### GetDefaultSrcOk

`func (o *IntegrationContentSecurityPolicy) GetDefaultSrcOk() (*string, bool)`

GetDefaultSrcOk returns a tuple with the DefaultSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSrc

`func (o *IntegrationContentSecurityPolicy) SetDefaultSrc(v string)`

SetDefaultSrc sets DefaultSrc field to given value.

### HasDefaultSrc

`func (o *IntegrationContentSecurityPolicy) HasDefaultSrc() bool`

HasDefaultSrc returns a boolean if a field has been set.

### GetFontSrc

`func (o *IntegrationContentSecurityPolicy) GetFontSrc() string`

GetFontSrc returns the FontSrc field if non-nil, zero value otherwise.

### GetFontSrcOk

`func (o *IntegrationContentSecurityPolicy) GetFontSrcOk() (*string, bool)`

GetFontSrcOk returns a tuple with the FontSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFontSrc

`func (o *IntegrationContentSecurityPolicy) SetFontSrc(v string)`

SetFontSrc sets FontSrc field to given value.

### HasFontSrc

`func (o *IntegrationContentSecurityPolicy) HasFontSrc() bool`

HasFontSrc returns a boolean if a field has been set.

### GetFormAction

`func (o *IntegrationContentSecurityPolicy) GetFormAction() string`

GetFormAction returns the FormAction field if non-nil, zero value otherwise.

### GetFormActionOk

`func (o *IntegrationContentSecurityPolicy) GetFormActionOk() (*string, bool)`

GetFormActionOk returns a tuple with the FormAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAction

`func (o *IntegrationContentSecurityPolicy) SetFormAction(v string)`

SetFormAction sets FormAction field to given value.

### HasFormAction

`func (o *IntegrationContentSecurityPolicy) HasFormAction() bool`

HasFormAction returns a boolean if a field has been set.

### GetFrameAncestors

`func (o *IntegrationContentSecurityPolicy) GetFrameAncestors() string`

GetFrameAncestors returns the FrameAncestors field if non-nil, zero value otherwise.

### GetFrameAncestorsOk

`func (o *IntegrationContentSecurityPolicy) GetFrameAncestorsOk() (*string, bool)`

GetFrameAncestorsOk returns a tuple with the FrameAncestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameAncestors

`func (o *IntegrationContentSecurityPolicy) SetFrameAncestors(v string)`

SetFrameAncestors sets FrameAncestors field to given value.

### HasFrameAncestors

`func (o *IntegrationContentSecurityPolicy) HasFrameAncestors() bool`

HasFrameAncestors returns a boolean if a field has been set.

### GetFrameSrc

`func (o *IntegrationContentSecurityPolicy) GetFrameSrc() string`

GetFrameSrc returns the FrameSrc field if non-nil, zero value otherwise.

### GetFrameSrcOk

`func (o *IntegrationContentSecurityPolicy) GetFrameSrcOk() (*string, bool)`

GetFrameSrcOk returns a tuple with the FrameSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameSrc

`func (o *IntegrationContentSecurityPolicy) SetFrameSrc(v string)`

SetFrameSrc sets FrameSrc field to given value.

### HasFrameSrc

`func (o *IntegrationContentSecurityPolicy) HasFrameSrc() bool`

HasFrameSrc returns a boolean if a field has been set.

### GetImgSrc

`func (o *IntegrationContentSecurityPolicy) GetImgSrc() string`

GetImgSrc returns the ImgSrc field if non-nil, zero value otherwise.

### GetImgSrcOk

`func (o *IntegrationContentSecurityPolicy) GetImgSrcOk() (*string, bool)`

GetImgSrcOk returns a tuple with the ImgSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgSrc

`func (o *IntegrationContentSecurityPolicy) SetImgSrc(v string)`

SetImgSrc sets ImgSrc field to given value.

### HasImgSrc

`func (o *IntegrationContentSecurityPolicy) HasImgSrc() bool`

HasImgSrc returns a boolean if a field has been set.

### GetManifestSrc

`func (o *IntegrationContentSecurityPolicy) GetManifestSrc() string`

GetManifestSrc returns the ManifestSrc field if non-nil, zero value otherwise.

### GetManifestSrcOk

`func (o *IntegrationContentSecurityPolicy) GetManifestSrcOk() (*string, bool)`

GetManifestSrcOk returns a tuple with the ManifestSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestSrc

`func (o *IntegrationContentSecurityPolicy) SetManifestSrc(v string)`

SetManifestSrc sets ManifestSrc field to given value.

### HasManifestSrc

`func (o *IntegrationContentSecurityPolicy) HasManifestSrc() bool`

HasManifestSrc returns a boolean if a field has been set.

### GetMediaSrc

`func (o *IntegrationContentSecurityPolicy) GetMediaSrc() string`

GetMediaSrc returns the MediaSrc field if non-nil, zero value otherwise.

### GetMediaSrcOk

`func (o *IntegrationContentSecurityPolicy) GetMediaSrcOk() (*string, bool)`

GetMediaSrcOk returns a tuple with the MediaSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaSrc

`func (o *IntegrationContentSecurityPolicy) SetMediaSrc(v string)`

SetMediaSrc sets MediaSrc field to given value.

### HasMediaSrc

`func (o *IntegrationContentSecurityPolicy) HasMediaSrc() bool`

HasMediaSrc returns a boolean if a field has been set.

### GetNavigateTo

`func (o *IntegrationContentSecurityPolicy) GetNavigateTo() string`

GetNavigateTo returns the NavigateTo field if non-nil, zero value otherwise.

### GetNavigateToOk

`func (o *IntegrationContentSecurityPolicy) GetNavigateToOk() (*string, bool)`

GetNavigateToOk returns a tuple with the NavigateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigateTo

`func (o *IntegrationContentSecurityPolicy) SetNavigateTo(v string)`

SetNavigateTo sets NavigateTo field to given value.

### HasNavigateTo

`func (o *IntegrationContentSecurityPolicy) HasNavigateTo() bool`

HasNavigateTo returns a boolean if a field has been set.

### GetObjectSrc

`func (o *IntegrationContentSecurityPolicy) GetObjectSrc() string`

GetObjectSrc returns the ObjectSrc field if non-nil, zero value otherwise.

### GetObjectSrcOk

`func (o *IntegrationContentSecurityPolicy) GetObjectSrcOk() (*string, bool)`

GetObjectSrcOk returns a tuple with the ObjectSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSrc

`func (o *IntegrationContentSecurityPolicy) SetObjectSrc(v string)`

SetObjectSrc sets ObjectSrc field to given value.

### HasObjectSrc

`func (o *IntegrationContentSecurityPolicy) HasObjectSrc() bool`

HasObjectSrc returns a boolean if a field has been set.

### GetPluginTypes

`func (o *IntegrationContentSecurityPolicy) GetPluginTypes() string`

GetPluginTypes returns the PluginTypes field if non-nil, zero value otherwise.

### GetPluginTypesOk

`func (o *IntegrationContentSecurityPolicy) GetPluginTypesOk() (*string, bool)`

GetPluginTypesOk returns a tuple with the PluginTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginTypes

`func (o *IntegrationContentSecurityPolicy) SetPluginTypes(v string)`

SetPluginTypes sets PluginTypes field to given value.

### HasPluginTypes

`func (o *IntegrationContentSecurityPolicy) HasPluginTypes() bool`

HasPluginTypes returns a boolean if a field has been set.

### GetPrefetchSrc

`func (o *IntegrationContentSecurityPolicy) GetPrefetchSrc() string`

GetPrefetchSrc returns the PrefetchSrc field if non-nil, zero value otherwise.

### GetPrefetchSrcOk

`func (o *IntegrationContentSecurityPolicy) GetPrefetchSrcOk() (*string, bool)`

GetPrefetchSrcOk returns a tuple with the PrefetchSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefetchSrc

`func (o *IntegrationContentSecurityPolicy) SetPrefetchSrc(v string)`

SetPrefetchSrc sets PrefetchSrc field to given value.

### HasPrefetchSrc

`func (o *IntegrationContentSecurityPolicy) HasPrefetchSrc() bool`

HasPrefetchSrc returns a boolean if a field has been set.

### GetReferrer

`func (o *IntegrationContentSecurityPolicy) GetReferrer() string`

GetReferrer returns the Referrer field if non-nil, zero value otherwise.

### GetReferrerOk

`func (o *IntegrationContentSecurityPolicy) GetReferrerOk() (*string, bool)`

GetReferrerOk returns a tuple with the Referrer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrer

`func (o *IntegrationContentSecurityPolicy) SetReferrer(v string)`

SetReferrer sets Referrer field to given value.

### HasReferrer

`func (o *IntegrationContentSecurityPolicy) HasReferrer() bool`

HasReferrer returns a boolean if a field has been set.

### GetReportTo

`func (o *IntegrationContentSecurityPolicy) GetReportTo() string`

GetReportTo returns the ReportTo field if non-nil, zero value otherwise.

### GetReportToOk

`func (o *IntegrationContentSecurityPolicy) GetReportToOk() (*string, bool)`

GetReportToOk returns a tuple with the ReportTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTo

`func (o *IntegrationContentSecurityPolicy) SetReportTo(v string)`

SetReportTo sets ReportTo field to given value.

### HasReportTo

`func (o *IntegrationContentSecurityPolicy) HasReportTo() bool`

HasReportTo returns a boolean if a field has been set.

### GetReportUri

`func (o *IntegrationContentSecurityPolicy) GetReportUri() string`

GetReportUri returns the ReportUri field if non-nil, zero value otherwise.

### GetReportUriOk

`func (o *IntegrationContentSecurityPolicy) GetReportUriOk() (*string, bool)`

GetReportUriOk returns a tuple with the ReportUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUri

`func (o *IntegrationContentSecurityPolicy) SetReportUri(v string)`

SetReportUri sets ReportUri field to given value.

### HasReportUri

`func (o *IntegrationContentSecurityPolicy) HasReportUri() bool`

HasReportUri returns a boolean if a field has been set.

### GetRequireSriFor

`func (o *IntegrationContentSecurityPolicy) GetRequireSriFor() string`

GetRequireSriFor returns the RequireSriFor field if non-nil, zero value otherwise.

### GetRequireSriForOk

`func (o *IntegrationContentSecurityPolicy) GetRequireSriForOk() (*string, bool)`

GetRequireSriForOk returns a tuple with the RequireSriFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireSriFor

`func (o *IntegrationContentSecurityPolicy) SetRequireSriFor(v string)`

SetRequireSriFor sets RequireSriFor field to given value.

### HasRequireSriFor

`func (o *IntegrationContentSecurityPolicy) HasRequireSriFor() bool`

HasRequireSriFor returns a boolean if a field has been set.

### GetRequireTrustedTypesFor

`func (o *IntegrationContentSecurityPolicy) GetRequireTrustedTypesFor() string`

GetRequireTrustedTypesFor returns the RequireTrustedTypesFor field if non-nil, zero value otherwise.

### GetRequireTrustedTypesForOk

`func (o *IntegrationContentSecurityPolicy) GetRequireTrustedTypesForOk() (*string, bool)`

GetRequireTrustedTypesForOk returns a tuple with the RequireTrustedTypesFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireTrustedTypesFor

`func (o *IntegrationContentSecurityPolicy) SetRequireTrustedTypesFor(v string)`

SetRequireTrustedTypesFor sets RequireTrustedTypesFor field to given value.

### HasRequireTrustedTypesFor

`func (o *IntegrationContentSecurityPolicy) HasRequireTrustedTypesFor() bool`

HasRequireTrustedTypesFor returns a boolean if a field has been set.

### GetSandbox

`func (o *IntegrationContentSecurityPolicy) GetSandbox() string`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *IntegrationContentSecurityPolicy) GetSandboxOk() (*string, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *IntegrationContentSecurityPolicy) SetSandbox(v string)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *IntegrationContentSecurityPolicy) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetScriptSrc

`func (o *IntegrationContentSecurityPolicy) GetScriptSrc() string`

GetScriptSrc returns the ScriptSrc field if non-nil, zero value otherwise.

### GetScriptSrcOk

`func (o *IntegrationContentSecurityPolicy) GetScriptSrcOk() (*string, bool)`

GetScriptSrcOk returns a tuple with the ScriptSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSrc

`func (o *IntegrationContentSecurityPolicy) SetScriptSrc(v string)`

SetScriptSrc sets ScriptSrc field to given value.

### HasScriptSrc

`func (o *IntegrationContentSecurityPolicy) HasScriptSrc() bool`

HasScriptSrc returns a boolean if a field has been set.

### GetScriptSrcAttr

`func (o *IntegrationContentSecurityPolicy) GetScriptSrcAttr() string`

GetScriptSrcAttr returns the ScriptSrcAttr field if non-nil, zero value otherwise.

### GetScriptSrcAttrOk

`func (o *IntegrationContentSecurityPolicy) GetScriptSrcAttrOk() (*string, bool)`

GetScriptSrcAttrOk returns a tuple with the ScriptSrcAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSrcAttr

`func (o *IntegrationContentSecurityPolicy) SetScriptSrcAttr(v string)`

SetScriptSrcAttr sets ScriptSrcAttr field to given value.

### HasScriptSrcAttr

`func (o *IntegrationContentSecurityPolicy) HasScriptSrcAttr() bool`

HasScriptSrcAttr returns a boolean if a field has been set.

### GetScriptSrcElem

`func (o *IntegrationContentSecurityPolicy) GetScriptSrcElem() string`

GetScriptSrcElem returns the ScriptSrcElem field if non-nil, zero value otherwise.

### GetScriptSrcElemOk

`func (o *IntegrationContentSecurityPolicy) GetScriptSrcElemOk() (*string, bool)`

GetScriptSrcElemOk returns a tuple with the ScriptSrcElem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptSrcElem

`func (o *IntegrationContentSecurityPolicy) SetScriptSrcElem(v string)`

SetScriptSrcElem sets ScriptSrcElem field to given value.

### HasScriptSrcElem

`func (o *IntegrationContentSecurityPolicy) HasScriptSrcElem() bool`

HasScriptSrcElem returns a boolean if a field has been set.

### GetStyleSrc

`func (o *IntegrationContentSecurityPolicy) GetStyleSrc() string`

GetStyleSrc returns the StyleSrc field if non-nil, zero value otherwise.

### GetStyleSrcOk

`func (o *IntegrationContentSecurityPolicy) GetStyleSrcOk() (*string, bool)`

GetStyleSrcOk returns a tuple with the StyleSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleSrc

`func (o *IntegrationContentSecurityPolicy) SetStyleSrc(v string)`

SetStyleSrc sets StyleSrc field to given value.

### HasStyleSrc

`func (o *IntegrationContentSecurityPolicy) HasStyleSrc() bool`

HasStyleSrc returns a boolean if a field has been set.

### GetStyleSrcAttr

`func (o *IntegrationContentSecurityPolicy) GetStyleSrcAttr() string`

GetStyleSrcAttr returns the StyleSrcAttr field if non-nil, zero value otherwise.

### GetStyleSrcAttrOk

`func (o *IntegrationContentSecurityPolicy) GetStyleSrcAttrOk() (*string, bool)`

GetStyleSrcAttrOk returns a tuple with the StyleSrcAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleSrcAttr

`func (o *IntegrationContentSecurityPolicy) SetStyleSrcAttr(v string)`

SetStyleSrcAttr sets StyleSrcAttr field to given value.

### HasStyleSrcAttr

`func (o *IntegrationContentSecurityPolicy) HasStyleSrcAttr() bool`

HasStyleSrcAttr returns a boolean if a field has been set.

### GetStyleSrcElem

`func (o *IntegrationContentSecurityPolicy) GetStyleSrcElem() string`

GetStyleSrcElem returns the StyleSrcElem field if non-nil, zero value otherwise.

### GetStyleSrcElemOk

`func (o *IntegrationContentSecurityPolicy) GetStyleSrcElemOk() (*string, bool)`

GetStyleSrcElemOk returns a tuple with the StyleSrcElem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleSrcElem

`func (o *IntegrationContentSecurityPolicy) SetStyleSrcElem(v string)`

SetStyleSrcElem sets StyleSrcElem field to given value.

### HasStyleSrcElem

`func (o *IntegrationContentSecurityPolicy) HasStyleSrcElem() bool`

HasStyleSrcElem returns a boolean if a field has been set.

### GetTrustedTypes

`func (o *IntegrationContentSecurityPolicy) GetTrustedTypes() string`

GetTrustedTypes returns the TrustedTypes field if non-nil, zero value otherwise.

### GetTrustedTypesOk

`func (o *IntegrationContentSecurityPolicy) GetTrustedTypesOk() (*string, bool)`

GetTrustedTypesOk returns a tuple with the TrustedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedTypes

`func (o *IntegrationContentSecurityPolicy) SetTrustedTypes(v string)`

SetTrustedTypes sets TrustedTypes field to given value.

### HasTrustedTypes

`func (o *IntegrationContentSecurityPolicy) HasTrustedTypes() bool`

HasTrustedTypes returns a boolean if a field has been set.

### GetUpgradeInsecureRequests

`func (o *IntegrationContentSecurityPolicy) GetUpgradeInsecureRequests() string`

GetUpgradeInsecureRequests returns the UpgradeInsecureRequests field if non-nil, zero value otherwise.

### GetUpgradeInsecureRequestsOk

`func (o *IntegrationContentSecurityPolicy) GetUpgradeInsecureRequestsOk() (*string, bool)`

GetUpgradeInsecureRequestsOk returns a tuple with the UpgradeInsecureRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeInsecureRequests

`func (o *IntegrationContentSecurityPolicy) SetUpgradeInsecureRequests(v string)`

SetUpgradeInsecureRequests sets UpgradeInsecureRequests field to given value.

### HasUpgradeInsecureRequests

`func (o *IntegrationContentSecurityPolicy) HasUpgradeInsecureRequests() bool`

HasUpgradeInsecureRequests returns a boolean if a field has been set.

### GetWorkerSrc

`func (o *IntegrationContentSecurityPolicy) GetWorkerSrc() string`

GetWorkerSrc returns the WorkerSrc field if non-nil, zero value otherwise.

### GetWorkerSrcOk

`func (o *IntegrationContentSecurityPolicy) GetWorkerSrcOk() (*string, bool)`

GetWorkerSrcOk returns a tuple with the WorkerSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerSrc

`func (o *IntegrationContentSecurityPolicy) SetWorkerSrc(v string)`

SetWorkerSrc sets WorkerSrc field to given value.

### HasWorkerSrc

`func (o *IntegrationContentSecurityPolicy) HasWorkerSrc() bool`

HasWorkerSrc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


