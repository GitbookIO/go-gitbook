/*
GitBook API

The GitBook API

API version: 0.0.1-beta
Contact: support@gitbook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the RevisionPageDocumentAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevisionPageDocumentAllOf{}

// RevisionPageDocumentAllOf struct for RevisionPageDocumentAllOf
type RevisionPageDocumentAllOf struct {
	// Deprecated
	Kind string `json:"kind"`
	Type string `json:"type"`
	// Page's slug in its direct parent
	Slug string `json:"slug"`
	// Complete path to access the page in the revision.
	Path        string                                `json:"path"`
	Description *string                               `json:"description,omitempty"`
	Pages       []RevisionPageDocumentAllOfPagesInner `json:"pages"`
}

// NewRevisionPageDocumentAllOf instantiates a new RevisionPageDocumentAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevisionPageDocumentAllOf(kind string, type_ string, slug string, path string, pages []RevisionPageDocumentAllOfPagesInner) *RevisionPageDocumentAllOf {
	this := RevisionPageDocumentAllOf{}
	this.Kind = kind
	this.Type = type_
	this.Slug = slug
	this.Path = path
	this.Pages = pages
	return &this
}

// NewRevisionPageDocumentAllOfWithDefaults instantiates a new RevisionPageDocumentAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevisionPageDocumentAllOfWithDefaults() *RevisionPageDocumentAllOf {
	this := RevisionPageDocumentAllOf{}
	return &this
}

// GetKind returns the Kind field value
// Deprecated
func (o *RevisionPageDocumentAllOf) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *RevisionPageDocumentAllOf) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
// Deprecated
func (o *RevisionPageDocumentAllOf) SetKind(v string) {
	o.Kind = v
}

// GetType returns the Type field value
func (o *RevisionPageDocumentAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RevisionPageDocumentAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RevisionPageDocumentAllOf) SetType(v string) {
	o.Type = v
}

// GetSlug returns the Slug field value
func (o *RevisionPageDocumentAllOf) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *RevisionPageDocumentAllOf) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *RevisionPageDocumentAllOf) SetSlug(v string) {
	o.Slug = v
}

// GetPath returns the Path field value
func (o *RevisionPageDocumentAllOf) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *RevisionPageDocumentAllOf) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *RevisionPageDocumentAllOf) SetPath(v string) {
	o.Path = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RevisionPageDocumentAllOf) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RevisionPageDocumentAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RevisionPageDocumentAllOf) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RevisionPageDocumentAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetPages returns the Pages field value
func (o *RevisionPageDocumentAllOf) GetPages() []RevisionPageDocumentAllOfPagesInner {
	if o == nil {
		var ret []RevisionPageDocumentAllOfPagesInner
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *RevisionPageDocumentAllOf) GetPagesOk() ([]RevisionPageDocumentAllOfPagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *RevisionPageDocumentAllOf) SetPages(v []RevisionPageDocumentAllOfPagesInner) {
	o.Pages = v
}

func (o RevisionPageDocumentAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevisionPageDocumentAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["kind"] = o.Kind
	toSerialize["type"] = o.Type
	toSerialize["slug"] = o.Slug
	toSerialize["path"] = o.Path
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["pages"] = o.Pages
	return toSerialize, nil
}

type NullableRevisionPageDocumentAllOf struct {
	value *RevisionPageDocumentAllOf
	isSet bool
}

func (v NullableRevisionPageDocumentAllOf) Get() *RevisionPageDocumentAllOf {
	return v.value
}

func (v *NullableRevisionPageDocumentAllOf) Set(val *RevisionPageDocumentAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevisionPageDocumentAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevisionPageDocumentAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevisionPageDocumentAllOf(val *RevisionPageDocumentAllOf) *NullableRevisionPageDocumentAllOf {
	return &NullableRevisionPageDocumentAllOf{value: val, isSet: true}
}

func (v NullableRevisionPageDocumentAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevisionPageDocumentAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
