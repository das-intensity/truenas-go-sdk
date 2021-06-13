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

// ConfigReset0 struct for ConfigReset0
type ConfigReset0 struct {
	Reboot *bool `json:"reboot,omitempty"`
}

// NewConfigReset0 instantiates a new ConfigReset0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigReset0() *ConfigReset0 {
	this := ConfigReset0{}
	return &this
}

// NewConfigReset0WithDefaults instantiates a new ConfigReset0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigReset0WithDefaults() *ConfigReset0 {
	this := ConfigReset0{}
	return &this
}

// GetReboot returns the Reboot field value if set, zero value otherwise.
func (o *ConfigReset0) GetReboot() bool {
	if o == nil || o.Reboot == nil {
		var ret bool
		return ret
	}
	return *o.Reboot
}

// GetRebootOk returns a tuple with the Reboot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigReset0) GetRebootOk() (*bool, bool) {
	if o == nil || o.Reboot == nil {
		return nil, false
	}
	return o.Reboot, true
}

// HasReboot returns a boolean if a field has been set.
func (o *ConfigReset0) HasReboot() bool {
	if o != nil && o.Reboot != nil {
		return true
	}

	return false
}

// SetReboot gets a reference to the given bool and assigns it to the Reboot field.
func (o *ConfigReset0) SetReboot(v bool) {
	o.Reboot = &v
}

func (o ConfigReset0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reboot != nil {
		toSerialize["reboot"] = o.Reboot
	}
	return json.Marshal(toSerialize)
}

type NullableConfigReset0 struct {
	value *ConfigReset0
	isSet bool
}

func (v NullableConfigReset0) Get() *ConfigReset0 {
	return v.value
}

func (v *NullableConfigReset0) Set(val *ConfigReset0) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigReset0) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigReset0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigReset0(val *ConfigReset0) *NullableConfigReset0 {
	return &NullableConfigReset0{value: val, isSet: true}
}

func (v NullableConfigReset0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigReset0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

