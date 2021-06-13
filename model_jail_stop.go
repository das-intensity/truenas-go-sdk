/*
 * TrueNAS RESTful API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// JailStop struct for JailStop
type JailStop struct {
	Jail *string `json:"jail,omitempty"`
	Force *bool `json:"force,omitempty"`
}

// NewJailStop instantiates a new JailStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJailStop() *JailStop {
	this := JailStop{}
	var force bool = false
	this.Force = &force
	return &this
}

// NewJailStopWithDefaults instantiates a new JailStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJailStopWithDefaults() *JailStop {
	this := JailStop{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetJail returns the Jail field value if set, zero value otherwise.
func (o *JailStop) GetJail() string {
	if o == nil || o.Jail == nil {
		var ret string
		return ret
	}
	return *o.Jail
}

// GetJailOk returns a tuple with the Jail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailStop) GetJailOk() (*string, bool) {
	if o == nil || o.Jail == nil {
		return nil, false
	}
	return o.Jail, true
}

// HasJail returns a boolean if a field has been set.
func (o *JailStop) HasJail() bool {
	if o != nil && o.Jail != nil {
		return true
	}

	return false
}

// SetJail gets a reference to the given string and assigns it to the Jail field.
func (o *JailStop) SetJail(v string) {
	o.Jail = &v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *JailStop) GetForce() bool {
	if o == nil || o.Force == nil {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JailStop) GetForceOk() (*bool, bool) {
	if o == nil || o.Force == nil {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *JailStop) HasForce() bool {
	if o != nil && o.Force != nil {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *JailStop) SetForce(v bool) {
	o.Force = &v
}

func (o JailStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Jail != nil {
		toSerialize["jail"] = o.Jail
	}
	if o.Force != nil {
		toSerialize["force"] = o.Force
	}
	return json.Marshal(toSerialize)
}

type NullableJailStop struct {
	value *JailStop
	isSet bool
}

func (v NullableJailStop) Get() *JailStop {
	return v.value
}

func (v *NullableJailStop) Set(val *JailStop) {
	v.value = val
	v.isSet = true
}

func (v NullableJailStop) IsSet() bool {
	return v.isSet
}

func (v *NullableJailStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJailStop(val *JailStop) *NullableJailStop {
	return &NullableJailStop{value: val, isSet: true}
}

func (v NullableJailStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJailStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

