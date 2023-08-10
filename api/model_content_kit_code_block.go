/*
Copyright 2023 GitBook, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gitbook

import (
	"encoding/json"
)

// checks if the ContentKitCodeBlock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentKitCodeBlock{}

// ContentKitCodeBlock Code block with syntax highlighting
type ContentKitCodeBlock struct {
	Type    string                     `json:"type"`
	Content ContentKitCodeBlockContent `json:"content"`
	// Syntax to use for highlighting (ex: javascript, python)
	Syntax      *string                         `json:"syntax,omitempty"`
	LineNumbers *ContentKitCodeBlockLineNumbers `json:"lineNumbers,omitempty"`
	// Controls button shown as an overlay in a corner of the code block.
	Buttons []ContentKitButton `json:"buttons,omitempty"`
	// State binding when editable. The value of the input will be stored as a property in the state named after this ID.
	State           *string           `json:"state,omitempty"`
	OnContentChange *ContentKitAction `json:"onContentChange,omitempty"`
	// Header displayed before the code lines
	Header []ContentKitDescendantElement `json:"header,omitempty"`
	// Footer displayed after the code lines
	Footer []ContentKitDescendantElement `json:"footer,omitempty"`
}

// NewContentKitCodeBlock instantiates a new ContentKitCodeBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentKitCodeBlock(type_ string, content ContentKitCodeBlockContent) *ContentKitCodeBlock {
	this := ContentKitCodeBlock{}
	this.Type = type_
	this.Content = content
	return &this
}

// NewContentKitCodeBlockWithDefaults instantiates a new ContentKitCodeBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentKitCodeBlockWithDefaults() *ContentKitCodeBlock {
	this := ContentKitCodeBlock{}
	return &this
}

// GetType returns the Type field value
func (o *ContentKitCodeBlock) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContentKitCodeBlock) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *ContentKitCodeBlock) GetContent() ContentKitCodeBlockContent {
	if o == nil {
		var ret ContentKitCodeBlockContent
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetContentOk() (*ContentKitCodeBlockContent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *ContentKitCodeBlock) SetContent(v ContentKitCodeBlockContent) {
	o.Content = v
}

// GetSyntax returns the Syntax field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetSyntax() string {
	if o == nil || IsNil(o.Syntax) {
		var ret string
		return ret
	}
	return *o.Syntax
}

// GetSyntaxOk returns a tuple with the Syntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetSyntaxOk() (*string, bool) {
	if o == nil || IsNil(o.Syntax) {
		return nil, false
	}
	return o.Syntax, true
}

// HasSyntax returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasSyntax() bool {
	if o != nil && !IsNil(o.Syntax) {
		return true
	}

	return false
}

// SetSyntax gets a reference to the given string and assigns it to the Syntax field.
func (o *ContentKitCodeBlock) SetSyntax(v string) {
	o.Syntax = &v
}

// GetLineNumbers returns the LineNumbers field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetLineNumbers() ContentKitCodeBlockLineNumbers {
	if o == nil || IsNil(o.LineNumbers) {
		var ret ContentKitCodeBlockLineNumbers
		return ret
	}
	return *o.LineNumbers
}

// GetLineNumbersOk returns a tuple with the LineNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetLineNumbersOk() (*ContentKitCodeBlockLineNumbers, bool) {
	if o == nil || IsNil(o.LineNumbers) {
		return nil, false
	}
	return o.LineNumbers, true
}

// HasLineNumbers returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasLineNumbers() bool {
	if o != nil && !IsNil(o.LineNumbers) {
		return true
	}

	return false
}

// SetLineNumbers gets a reference to the given ContentKitCodeBlockLineNumbers and assigns it to the LineNumbers field.
func (o *ContentKitCodeBlock) SetLineNumbers(v ContentKitCodeBlockLineNumbers) {
	o.LineNumbers = &v
}

// GetButtons returns the Buttons field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetButtons() []ContentKitButton {
	if o == nil || IsNil(o.Buttons) {
		var ret []ContentKitButton
		return ret
	}
	return o.Buttons
}

// GetButtonsOk returns a tuple with the Buttons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetButtonsOk() ([]ContentKitButton, bool) {
	if o == nil || IsNil(o.Buttons) {
		return nil, false
	}
	return o.Buttons, true
}

// HasButtons returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasButtons() bool {
	if o != nil && !IsNil(o.Buttons) {
		return true
	}

	return false
}

// SetButtons gets a reference to the given []ContentKitButton and assigns it to the Buttons field.
func (o *ContentKitCodeBlock) SetButtons(v []ContentKitButton) {
	o.Buttons = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ContentKitCodeBlock) SetState(v string) {
	o.State = &v
}

// GetOnContentChange returns the OnContentChange field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetOnContentChange() ContentKitAction {
	if o == nil || IsNil(o.OnContentChange) {
		var ret ContentKitAction
		return ret
	}
	return *o.OnContentChange
}

// GetOnContentChangeOk returns a tuple with the OnContentChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetOnContentChangeOk() (*ContentKitAction, bool) {
	if o == nil || IsNil(o.OnContentChange) {
		return nil, false
	}
	return o.OnContentChange, true
}

// HasOnContentChange returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasOnContentChange() bool {
	if o != nil && !IsNil(o.OnContentChange) {
		return true
	}

	return false
}

// SetOnContentChange gets a reference to the given ContentKitAction and assigns it to the OnContentChange field.
func (o *ContentKitCodeBlock) SetOnContentChange(v ContentKitAction) {
	o.OnContentChange = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetHeader() []ContentKitDescendantElement {
	if o == nil || IsNil(o.Header) {
		var ret []ContentKitDescendantElement
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetHeaderOk() ([]ContentKitDescendantElement, bool) {
	if o == nil || IsNil(o.Header) {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given []ContentKitDescendantElement and assigns it to the Header field.
func (o *ContentKitCodeBlock) SetHeader(v []ContentKitDescendantElement) {
	o.Header = v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *ContentKitCodeBlock) GetFooter() []ContentKitDescendantElement {
	if o == nil || IsNil(o.Footer) {
		var ret []ContentKitDescendantElement
		return ret
	}
	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentKitCodeBlock) GetFooterOk() ([]ContentKitDescendantElement, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *ContentKitCodeBlock) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given []ContentKitDescendantElement and assigns it to the Footer field.
func (o *ContentKitCodeBlock) SetFooter(v []ContentKitDescendantElement) {
	o.Footer = v
}

func (o ContentKitCodeBlock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentKitCodeBlock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["content"] = o.Content
	if !IsNil(o.Syntax) {
		toSerialize["syntax"] = o.Syntax
	}
	if !IsNil(o.LineNumbers) {
		toSerialize["lineNumbers"] = o.LineNumbers
	}
	if !IsNil(o.Buttons) {
		toSerialize["buttons"] = o.Buttons
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.OnContentChange) {
		toSerialize["onContentChange"] = o.OnContentChange
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	return toSerialize, nil
}

type NullableContentKitCodeBlock struct {
	value *ContentKitCodeBlock
	isSet bool
}

func (v NullableContentKitCodeBlock) Get() *ContentKitCodeBlock {
	return v.value
}

func (v *NullableContentKitCodeBlock) Set(val *ContentKitCodeBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableContentKitCodeBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableContentKitCodeBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentKitCodeBlock(val *ContentKitCodeBlock) *NullableContentKitCodeBlock {
	return &NullableContentKitCodeBlock{value: val, isSet: true}
}

func (v NullableContentKitCodeBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentKitCodeBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
