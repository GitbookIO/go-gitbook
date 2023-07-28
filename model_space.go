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

// checks if the Space type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Space{}

// Space struct for Space
type Space struct {
	// Type of Object, always equals to \"space\"
	Object string `json:"object"`
	// Unique identifier for the space
	Id   string    `json:"id"`
	Type SpaceType `json:"type"`
	// Title of the space
	Title      string            `json:"title"`
	Visibility ContentVisibility `json:"visibility"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Urls       SpaceUrls         `json:"urls"`
	// ID of the organization owning this space
	Organization *string `json:"organization,omitempty"`
	// ID of the user owning this space
	// Deprecated
	User *string `json:"user,omitempty"`
	// If the space is part of the private content of a user, the property is equal to the ID of the user owning it
	// Deprecated
	Private *string `json:"private,omitempty"`
	// ID of the parent collection.
	Parent *string `json:"parent,omitempty"`
}

// NewSpace instantiates a new Space object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpace(object string, id string, type_ SpaceType, title string, visibility ContentVisibility, createdAt string, updatedAt string, urls SpaceUrls) *Space {
	this := Space{}
	this.Object = object
	this.Id = id
	this.Type = type_
	this.Title = title
	this.Visibility = visibility
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Urls = urls
	return &this
}

// NewSpaceWithDefaults instantiates a new Space object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceWithDefaults() *Space {
	this := Space{}
	return &this
}

// GetObject returns the Object field value
func (o *Space) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Space) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Space) SetObject(v string) {
	o.Object = v
}

// GetId returns the Id field value
func (o *Space) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Space) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Space) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *Space) GetType() SpaceType {
	if o == nil {
		var ret SpaceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Space) GetTypeOk() (*SpaceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Space) SetType(v SpaceType) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *Space) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Space) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Space) SetTitle(v string) {
	o.Title = v
}

// GetVisibility returns the Visibility field value
func (o *Space) GetVisibility() ContentVisibility {
	if o == nil {
		var ret ContentVisibility
		return ret
	}

	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value
// and a boolean to check if the value has been set.
func (o *Space) GetVisibilityOk() (*ContentVisibility, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visibility, true
}

// SetVisibility sets field value
func (o *Space) SetVisibility(v ContentVisibility) {
	o.Visibility = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Space) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Space) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Space) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Space) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Space) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Space) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUrls returns the Urls field value
func (o *Space) GetUrls() SpaceUrls {
	if o == nil {
		var ret SpaceUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *Space) GetUrlsOk() (*SpaceUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *Space) SetUrls(v SpaceUrls) {
	o.Urls = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Space) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Space) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *Space) SetOrganization(v string) {
	o.Organization = &v
}

// GetUser returns the User field value if set, zero value otherwise.
// Deprecated
func (o *Space) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Space) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Space) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
// Deprecated
func (o *Space) SetUser(v string) {
	o.User = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
// Deprecated
func (o *Space) GetPrivate() string {
	if o == nil || IsNil(o.Private) {
		var ret string
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Space) GetPrivateOk() (*string, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *Space) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given string and assigns it to the Private field.
// Deprecated
func (o *Space) SetPrivate(v string) {
	o.Private = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *Space) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Space) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *Space) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *Space) SetParent(v string) {
	o.Parent = &v
}

func (o Space) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Space) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["visibility"] = o.Visibility
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["urls"] = o.Urls
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	return toSerialize, nil
}

type NullableSpace struct {
	value *Space
	isSet bool
}

func (v NullableSpace) Get() *Space {
	return v.value
}

func (v *NullableSpace) Set(val *Space) {
	v.value = val
	v.isSet = true
}

func (v NullableSpace) IsSet() bool {
	return v.isSet
}

func (v *NullableSpace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpace(val *Space) *NullableSpace {
	return &NullableSpace{value: val, isSet: true}
}

func (v NullableSpace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
