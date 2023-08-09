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

// checks if the SpaceRelationPair type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpaceRelationPair{}

// SpaceRelationPair struct for SpaceRelationPair
type SpaceRelationPair struct {
	Id        string            `json:"id"`
	Type      SpaceRelationType `json:"type"`
	Target    SpaceRelationEdge `json:"target"`
	CreatedAt string            `json:"createdAt"`
	Urls      SpaceRelationUrls `json:"urls"`
	Source    SpaceRelationEdge `json:"source"`
}

// NewSpaceRelationPair instantiates a new SpaceRelationPair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpaceRelationPair(id string, type_ SpaceRelationType, target SpaceRelationEdge, createdAt string, urls SpaceRelationUrls, source SpaceRelationEdge) *SpaceRelationPair {
	this := SpaceRelationPair{}
	this.Id = id
	this.Type = type_
	this.Target = target
	this.CreatedAt = createdAt
	this.Urls = urls
	this.Source = source
	return &this
}

// NewSpaceRelationPairWithDefaults instantiates a new SpaceRelationPair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpaceRelationPairWithDefaults() *SpaceRelationPair {
	this := SpaceRelationPair{}
	return &this
}

// GetId returns the Id field value
func (o *SpaceRelationPair) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpaceRelationPair) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *SpaceRelationPair) GetType() SpaceRelationType {
	if o == nil {
		var ret SpaceRelationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetTypeOk() (*SpaceRelationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpaceRelationPair) SetType(v SpaceRelationType) {
	o.Type = v
}

// GetTarget returns the Target field value
func (o *SpaceRelationPair) GetTarget() SpaceRelationEdge {
	if o == nil {
		var ret SpaceRelationEdge
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetTargetOk() (*SpaceRelationEdge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *SpaceRelationPair) SetTarget(v SpaceRelationEdge) {
	o.Target = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SpaceRelationPair) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SpaceRelationPair) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUrls returns the Urls field value
func (o *SpaceRelationPair) GetUrls() SpaceRelationUrls {
	if o == nil {
		var ret SpaceRelationUrls
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetUrlsOk() (*SpaceRelationUrls, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urls, true
}

// SetUrls sets field value
func (o *SpaceRelationPair) SetUrls(v SpaceRelationUrls) {
	o.Urls = v
}

// GetSource returns the Source field value
func (o *SpaceRelationPair) GetSource() SpaceRelationEdge {
	if o == nil {
		var ret SpaceRelationEdge
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *SpaceRelationPair) GetSourceOk() (*SpaceRelationEdge, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *SpaceRelationPair) SetSource(v SpaceRelationEdge) {
	o.Source = v
}

func (o SpaceRelationPair) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpaceRelationPair) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["target"] = o.Target
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["urls"] = o.Urls
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableSpaceRelationPair struct {
	value *SpaceRelationPair
	isSet bool
}

func (v NullableSpaceRelationPair) Get() *SpaceRelationPair {
	return v.value
}

func (v *NullableSpaceRelationPair) Set(val *SpaceRelationPair) {
	v.value = val
	v.isSet = true
}

func (v NullableSpaceRelationPair) IsSet() bool {
	return v.isSet
}

func (v *NullableSpaceRelationPair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpaceRelationPair(val *SpaceRelationPair) *NullableSpaceRelationPair {
	return &NullableSpaceRelationPair{value: val, isSet: true}
}

func (v NullableSpaceRelationPair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpaceRelationPair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}