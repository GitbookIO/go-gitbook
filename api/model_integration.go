// Copyright 2023 GitBook, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the Integration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Integration{}

// Integration struct for Integration
type Integration struct {
	Object string `json:"object"`
	// Unique named identifier for the integration
	Name string `json:"name"`
	// Version of the integration
	Version float32 `json:"version"`
	// Title of the integration
	Title string `json:"title"`
	// Description of the integration
	Description *string `json:"description,omitempty"`
	// Long form markdown summary of the integration
	Summary *string `json:"summary,omitempty"`
	// URLs of images to showcase the integration
	PreviewImages []string              `json:"previewImages"`
	Visibility    IntegrationVisibility `json:"visibility"`
	// Permissions that should be granted to the integration
	Scopes []IntegrationScope `json:"scopes"`
	// Categories for which the integration is listed in the marketplace
	Categories []IntegrationCategory `json:"categories"`
	// Custom blocks defined by this integration.
	Blocks         []IntegrationBlock         `json:"blocks,omitempty"`
	Configurations *IntegrationConfigurations `json:"configurations,omitempty"`
	// External urls configured by the developer of the integration
	ExternalLinks         []IntegrationExternalLinksInner   `json:"externalLinks"`
	Urls                  IntegrationUrls                   `json:"urls"`
	ContentSecurityPolicy *IntegrationContentSecurityPolicy `json:"contentSecurityPolicy,omitempty"`
}

// NewIntegration instantiates a new Integration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegration(object string, name string, version float32, title string, previewImages []string, visibility IntegrationVisibility, scopes []IntegrationScope, categories []IntegrationCategory, externalLinks []IntegrationExternalLinksInner, urls IntegrationUrls) *Integration {
	this := Integration{}
	this.Object = object
	this.Name = name
	this.Version = version
	this.Title = title
	this.PreviewImages = previewImages
	this.Visibility = visibility
	this.Scopes = scopes
	this.Categories = categories
	this.ExternalLinks = externalLinks
	this.Urls = urls
	return &this
}

// NewIntegrationWithDefaults instantiates a new Integration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationWithDefaults() *Integration {
	this := Integration{}
	return &this
}

// GetObject returns the Object field value
func (o *Integration) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Integration) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Integration) SetObject(v string) {
	o.Object = v
}

// GetName returns the Name field value
func (o *Integration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Integration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Integration) SetName(v string) {
	o.Name = v
}

// GetVersion returns the Version field value
func (o *Integration) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *Integration) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *Integration) SetVersion(v float32) {
	o.Version = v
}

// GetTitle returns the Title field value
func (o *Integration) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Integration) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Integration) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Integration) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Integration) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Integration) SetDescription(v string) {
	o.Description = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Integration) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Integration) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Integration) SetSummary(v string) {
	o.Summary = &v
}

// GetPreviewImages returns the PreviewImages field value
func (o *Integration) GetPreviewImages() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PreviewImages
}

// GetPreviewImagesOk returns a tuple with the PreviewImages field value
// and a boolean to check if the value has been set.
func (o *Integration) GetPreviewImagesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviewImages, true
}

// SetPreviewImages sets field value
func (o *Integration) SetPreviewImages(v []string) {
	o.PreviewImages = v
}

// GetVisibility returns the Visibility field value
func (o *Integration) GetVisibility() IntegrationVisibility {
	if o == nil {
		var ret IntegrationVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *Integration) GetVisibilityOk() (*IntegrationVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *Integration) SetVisibility(v IntegrationVisibility) {
	o.Visibility = v
}

// GetScopes returns the Scopes field value
func (o *Integration) GetScopes() []IntegrationScope {
	if o == nil {
		var ret []IntegrationScope
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *Integration) GetScopesOk() ([]IntegrationScope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *Integration) SetScopes(v []IntegrationScope) {
	o.Scopes = v
}

// GetCategories returns the Categories field value
func (o *Integration) GetCategories() []IntegrationCategory {
	if o == nil {
		var ret []IntegrationCategory
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *Integration) GetCategoriesOk() ([]IntegrationCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *Integration) SetCategories(v []IntegrationCategory) {
	o.Categories = v
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *Integration) GetBlocks() []IntegrationBlock {
	if o == nil || IsNil(o.Blocks) {
		var ret []IntegrationBlock
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetBlocksOk() ([]IntegrationBlock, bool) {
	if o == nil || IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *Integration) HasBlocks() bool {
	if o != nil && !IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []IntegrationBlock and assigns it to the Blocks field.
func (o *Integration) SetBlocks(v []IntegrationBlock) {
	o.Blocks = v
}

// GetConfigurations returns the Configurations field value if set, zero value otherwise.
func (o *Integration) GetConfigurations() IntegrationConfigurations {
	if o == nil || IsNil(o.Configurations) {
		var ret IntegrationConfigurations
		return ret
	}
	return *o.Configurations
}

// GetConfigurationsOk returns a tuple with the Configurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetConfigurationsOk() (*IntegrationConfigurations, bool) {
	if o == nil || IsNil(o.Configurations) {
		return nil, false
	}
	return o.Configurations, true
}

// HasConfigurations returns a boolean if a field has been set.
func (o *Integration) HasConfigurations() bool {
	if o != nil && !IsNil(o.Configurations) {
		return true
	}

	return false
}

// SetConfigurations gets a reference to the given IntegrationConfigurations and assigns it to the Configurations field.
func (o *Integration) SetConfigurations(v IntegrationConfigurations) {
	o.Configurations = &v
}

// GetExternalLinks returns the ExternalLinks field value
func (o *Integration) GetExternalLinks() []IntegrationExternalLinksInner {
	if o == nil {
		var ret []IntegrationExternalLinksInner
		return ret
	}

	return o.ExternalLinks
}

// GetExternalLinksOk returns a tuple with the ExternalLinks field value
// and a boolean to check if the value has been set.
func (o *Integration) GetExternalLinksOk() ([]IntegrationExternalLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalLinks, true
}

// SetExternalLinks sets field value
func (o *Integration) SetExternalLinks(v []IntegrationExternalLinksInner) {
	o.ExternalLinks = v
}

// GetUrls returns the Urls field value
func (o *Integration) GetUrls() IntegrationUrls {
	if o == nil {
		var ret IntegrationUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Integration) GetUrlsOk() (*IntegrationUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Integration) SetUrls(v IntegrationUrls) {
	o.Urls = v
}

// GetContentSecurityPolicy returns the ContentSecurityPolicy field value if set, zero value otherwise.
func (o *Integration) GetContentSecurityPolicy() IntegrationContentSecurityPolicy {
	if o == nil || IsNil(o.ContentSecurityPolicy) {
		var ret IntegrationContentSecurityPolicy
		return ret
	}
	return *o.ContentSecurityPolicy
}

// GetContentSecurityPolicyOk returns a tuple with the ContentSecurityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Integration) GetContentSecurityPolicyOk() (*IntegrationContentSecurityPolicy, bool) {
	if o == nil || IsNil(o.ContentSecurityPolicy) {
		return nil, false
	}
	return o.ContentSecurityPolicy, true
}

// HasContentSecurityPolicy returns a boolean if a field has been set.
func (o *Integration) HasContentSecurityPolicy() bool {
	if o != nil && !IsNil(o.ContentSecurityPolicy) {
		return true
	}

	return false
}

// SetContentSecurityPolicy gets a reference to the given IntegrationContentSecurityPolicy and assigns it to the ContentSecurityPolicy field.
func (o *Integration) SetContentSecurityPolicy(v IntegrationContentSecurityPolicy) {
	o.ContentSecurityPolicy = &v
}

func (o Integration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Integration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["name"] = o.Name
	toSerialize["version"] = o.Version
	toSerialize["title"] = o.Title
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	toSerialize["previewImages"] = o.PreviewImages
	toSerialize["visibility"] = o.Visibility
	toSerialize["scopes"] = o.Scopes
	toSerialize["categories"] = o.Categories
	if !IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	if !IsNil(o.Configurations) {
		toSerialize["configurations"] = o.Configurations
	}
	toSerialize["externalLinks"] = o.ExternalLinks
	toSerialize["urls"] = o.Urls
	if !IsNil(o.ContentSecurityPolicy) {
		toSerialize["contentSecurityPolicy"] = o.ContentSecurityPolicy
	}
	return toSerialize, nil
}

type NullableIntegration struct {
	value *Integration
	isSet bool
}

func (v NullableIntegration) Get() *Integration {
	return v.value
}

func (v *NullableIntegration) Set(val *Integration) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegration(val *Integration) *NullableIntegration {
	return &NullableIntegration{value: val, isSet: true}
}

func (v NullableIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
