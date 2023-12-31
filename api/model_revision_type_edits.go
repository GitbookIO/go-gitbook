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

// checks if the RevisionTypeEdits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionTypeEdits{}

// RevisionTypeEdits struct for RevisionTypeEdits
type RevisionTypeEdits struct {
	// Type of Object, always equals to \"revision\"
	Object string `json:"object"`
	// Unique identifier for the revision
	Id string `json:"id"`
	// IDs of the parent revisions
	Parents []string         `json:"parents"`
	Pages   []RevisionPage   `json:"pages"`
	Files   []RevisionFile   `json:"files"`
	Git     *RevisionBaseGit `json:"git,omitempty"`
	Urls    RevisionBaseUrls `json:"urls"`
	// Revision created by editing the content.
	Type string `json:"type"`
}

// NewRevisionTypeEdits instantiates a new RevisionTypeEdits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionTypeEdits(object string, id string, parents []string, pages []RevisionPage, files []RevisionFile, urls RevisionBaseUrls, type_ string) *RevisionTypeEdits {
	this := RevisionTypeEdits{}
	this.Object = object
	this.Id = id
	this.Parents = parents
	this.Pages = pages
	this.Files = files
	this.Urls = urls
	this.Type = type_
	return &this
}

// NewRevisionTypeEditsWithDefaults instantiates a new RevisionTypeEdits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionTypeEditsWithDefaults() *RevisionTypeEdits {
	this := RevisionTypeEdits{}
	return &this
}

// GetObject returns the Object field value
func (o *RevisionTypeEdits) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *RevisionTypeEdits) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *RevisionTypeEdits) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RevisionTypeEdits) SetId(v string) {
	o.Id = v
}

// GetParents returns the Parents field value
func (o *RevisionTypeEdits) GetParents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetParentsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parents, true
}

// SetParents sets field value
func (o *RevisionTypeEdits) SetParents(v []string) {
	o.Parents = v
}

// GetPages returns the Pages field value
func (o *RevisionTypeEdits) GetPages() []RevisionPage {
	if o == nil {
		var ret []RevisionPage
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetPagesOk() ([]RevisionPage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *RevisionTypeEdits) SetPages(v []RevisionPage) {
	o.Pages = v
}

// GetFiles returns the Files field value
func (o *RevisionTypeEdits) GetFiles() []RevisionFile {
	if o == nil {
		var ret []RevisionFile
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetFilesOk() ([]RevisionFile, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *RevisionTypeEdits) SetFiles(v []RevisionFile) {
	o.Files = v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *RevisionTypeEdits) GetGit() RevisionBaseGit {
	if o == nil || IsNil(o.Git) {
		var ret RevisionBaseGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetGitOk() (*RevisionBaseGit, bool) {
	if o == nil || IsNil(o.Git) {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *RevisionTypeEdits) HasGit() bool {
	if o != nil && !IsNil(o.Git) {
		return true
	}

	return false
}

// SetGit gets a reference to the given RevisionBaseGit and assigns it to the Git field.
func (o *RevisionTypeEdits) SetGit(v RevisionBaseGit) {
	o.Git = &v
}

// GetUrls returns the Urls field value
func (o *RevisionTypeEdits) GetUrls() RevisionBaseUrls {
	if o == nil {
		var ret RevisionBaseUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetUrlsOk() (*RevisionBaseUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *RevisionTypeEdits) SetUrls(v RevisionBaseUrls) {
	o.Urls = v
}

// GetType returns the Type field value
func (o *RevisionTypeEdits) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RevisionTypeEdits) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RevisionTypeEdits) SetType(v string) {
	o.Type = v
}

func (o RevisionTypeEdits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionTypeEdits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["parents"] = o.Parents
	toSerialize["pages"] = o.Pages
	toSerialize["files"] = o.Files
	if !IsNil(o.Git) {
		toSerialize["git"] = o.Git
	}
	toSerialize["urls"] = o.Urls
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableRevisionTypeEdits struct {
	value *RevisionTypeEdits
	isSet bool
}

func (v NullableRevisionTypeEdits) Get() *RevisionTypeEdits {
	return v.value
}

func (v *NullableRevisionTypeEdits) Set(val *RevisionTypeEdits) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionTypeEdits) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionTypeEdits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionTypeEdits(val *RevisionTypeEdits) *NullableRevisionTypeEdits {
	return &NullableRevisionTypeEdits{value: val, isSet: true}
}

func (v NullableRevisionTypeEdits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionTypeEdits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
