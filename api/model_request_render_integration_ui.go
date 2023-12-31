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

// checks if the RequestRenderIntegrationUI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestRenderIntegrationUI{}

// RequestRenderIntegrationUI struct for RequestRenderIntegrationUI
type RequestRenderIntegrationUI struct {
	// ID of the component to render in the integration.
	ComponentId string `json:"componentId"`
	// Current properties of the UI.
	Props map[string]interface{} `json:"props"`
	// Current local state of the UI.
	State   map[string]interface{} `json:"state,omitempty"`
	Context ContentKitContext      `json:"context"`
	Action  *ContentKitAction      `json:"action,omitempty"`
}

// NewRequestRenderIntegrationUI instantiates a new RequestRenderIntegrationUI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestRenderIntegrationUI(componentId string, props map[string]interface{}, context ContentKitContext) *RequestRenderIntegrationUI {
	this := RequestRenderIntegrationUI{}
	this.ComponentId = componentId
	this.Props = props
	this.Context = context
	return &this
}

// NewRequestRenderIntegrationUIWithDefaults instantiates a new RequestRenderIntegrationUI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestRenderIntegrationUIWithDefaults() *RequestRenderIntegrationUI {
	this := RequestRenderIntegrationUI{}
	return &this
}

// GetComponentId returns the ComponentId field value
func (o *RequestRenderIntegrationUI) GetComponentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentId
}

// GetComponentIdOk returns a tuple with the ComponentId field value
// and a boolean to check if the value has been set.
func (o *RequestRenderIntegrationUI) GetComponentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentId, true
}

// SetComponentId sets field value
func (o *RequestRenderIntegrationUI) SetComponentId(v string) {
	o.ComponentId = v
}

// GetProps returns the Props field value
func (o *RequestRenderIntegrationUI) GetProps() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Props
}

// GetPropsOk returns a tuple with the Props field value
// and a boolean to check if the value has been set.
func (o *RequestRenderIntegrationUI) GetPropsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Props, true
}

// SetProps sets field value
func (o *RequestRenderIntegrationUI) SetProps(v map[string]interface{}) {
	o.Props = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RequestRenderIntegrationUI) GetState() map[string]interface{} {
	if o == nil || IsNil(o.State) {
		var ret map[string]interface{}
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestRenderIntegrationUI) GetStateOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.State) {
		return map[string]interface{}{}, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RequestRenderIntegrationUI) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given map[string]interface{} and assigns it to the State field.
func (o *RequestRenderIntegrationUI) SetState(v map[string]interface{}) {
	o.State = v
}

// GetContext returns the Context field value
func (o *RequestRenderIntegrationUI) GetContext() ContentKitContext {
	if o == nil {
		var ret ContentKitContext
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *RequestRenderIntegrationUI) GetContextOk() (*ContentKitContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *RequestRenderIntegrationUI) SetContext(v ContentKitContext) {
	o.Context = v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RequestRenderIntegrationUI) GetAction() ContentKitAction {
	if o == nil || IsNil(o.Action) {
		var ret ContentKitAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestRenderIntegrationUI) GetActionOk() (*ContentKitAction, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RequestRenderIntegrationUI) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given ContentKitAction and assigns it to the Action field.
func (o *RequestRenderIntegrationUI) SetAction(v ContentKitAction) {
	o.Action = &v
}

func (o RequestRenderIntegrationUI) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestRenderIntegrationUI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["componentId"] = o.ComponentId
	toSerialize["props"] = o.Props
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["context"] = o.Context
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableRequestRenderIntegrationUI struct {
	value *RequestRenderIntegrationUI
	isSet bool
}

func (v NullableRequestRenderIntegrationUI) Get() *RequestRenderIntegrationUI {
	return v.value
}

func (v *NullableRequestRenderIntegrationUI) Set(val *RequestRenderIntegrationUI) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestRenderIntegrationUI) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestRenderIntegrationUI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestRenderIntegrationUI(val *RequestRenderIntegrationUI) *NullableRequestRenderIntegrationUI {
	return &NullableRequestRenderIntegrationUI{value: val, isSet: true}
}

func (v NullableRequestRenderIntegrationUI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestRenderIntegrationUI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
