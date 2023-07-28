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

// checks if the ContentKitCard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitCard{}

// ContentKitCard struct for ContentKitCard
type ContentKitCard struct {
	Type     string                        `json:"type"`
	Title    *string                       `json:"title,omitempty"`
	Hint     *ContentKitCardHint           `json:"hint,omitempty"`
	Icon     *ContentKitCardIcon           `json:"icon,omitempty"`
	OnPress  *ContentKitAction             `json:"onPress,omitempty"`
	Children []ContentKitDescendantElement `json:"children,omitempty"`
	// Buttons displayed in the top right corner of the card.
	Buttons []ContentKitButton `json:"buttons,omitempty"`
}

// NewContentKitCard instantiates a new ContentKitCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitCard(type_ string) *ContentKitCard {
	this := ContentKitCard{}
	this.Type = type_
	return &this
}

// NewContentKitCardWithDefaults instantiates a new ContentKitCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitCardWithDefaults() *ContentKitCard {
	this := ContentKitCard{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitCard) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitCard) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ContentKitCard) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ContentKitCard) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ContentKitCard) SetTitle(v string) {
	o.Title = &v
}

// GetHint returns the Hint field value if set, zero value otherwise.
func (o *ContentKitCard) GetHint() ContentKitCardHint {
	if o == nil || IsNil(o.Hint) {
		var ret ContentKitCardHint
		return ret
	}
	return *o.Hint
}

// GetHintOk returns a tuple with the Hint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetHintOk() (*ContentKitCardHint, bool) {
	if o == nil || IsNil(o.Hint) {
		return nil, false
	}
	return o.Hint, true
}

// HasHint returns a boolean if a field has been set.
func (o *ContentKitCard) HasHint() bool {
	if o != nil && !IsNil(o.Hint) {
		return true
	}

	return false
}

// SetHint gets a reference to the given ContentKitCardHint and assigns it to the Hint field.
func (o *ContentKitCard) SetHint(v ContentKitCardHint) {
	o.Hint = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *ContentKitCard) GetIcon() ContentKitCardIcon {
	if o == nil || IsNil(o.Icon) {
		var ret ContentKitCardIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetIconOk() (*ContentKitCardIcon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *ContentKitCard) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ContentKitCardIcon and assigns it to the Icon field.
func (o *ContentKitCard) SetIcon(v ContentKitCardIcon) {
	o.Icon = &v
}

// GetOnPress returns the OnPress field value if set, zero value otherwise.
func (o *ContentKitCard) GetOnPress() ContentKitAction {
	if o == nil || IsNil(o.OnPress) {
		var ret ContentKitAction
		return ret
	}
	return *o.OnPress
}

// GetOnPressOk returns a tuple with the OnPress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetOnPressOk() (*ContentKitAction, bool) {
	if o == nil || IsNil(o.OnPress) {
		return nil, false
	}
	return o.OnPress, true
}

// HasOnPress returns a boolean if a field has been set.
func (o *ContentKitCard) HasOnPress() bool {
	if o != nil && !IsNil(o.OnPress) {
		return true
	}

	return false
}

// SetOnPress gets a reference to the given ContentKitAction and assigns it to the OnPress field.
func (o *ContentKitCard) SetOnPress(v ContentKitAction) {
	o.OnPress = &v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ContentKitCard) GetChildren() []ContentKitDescendantElement {
	if o == nil || IsNil(o.Children) {
		var ret []ContentKitDescendantElement
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetChildrenOk() ([]ContentKitDescendantElement, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ContentKitCard) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ContentKitDescendantElement and assigns it to the Children field.
func (o *ContentKitCard) SetChildren(v []ContentKitDescendantElement) {
	o.Children = v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *ContentKitCard) GetButtons() []ContentKitButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []ContentKitButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCard) GetButtonsOk() ([]ContentKitButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *ContentKitCard) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []ContentKitButton and assigns it to the Buttons field.
func (o *ContentKitCard) SetButtons(v []ContentKitButton) {
	o.Buttons = v
}

func (o ContentKitCard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Hint) {
		toSerialize["hint"] = o.Hint
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.OnPress) {
		toSerialize["onPress"] = o.OnPress
	}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	return toSerialize, nil
}

type NullableContentKitCard struct {
	value *ContentKitCard
	isSet bool
}

func (v NullableContentKitCard) Get() *ContentKitCard {
	return v.value
}

func (v *NullableContentKitCard) Set(val *ContentKitCard) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitCard) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitCard(val *ContentKitCard) *NullableContentKitCard {
	return &NullableContentKitCard{value: val, isSet: true}
}

func (v NullableContentKitCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
