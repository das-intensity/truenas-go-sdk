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

// SmbStatus3 struct for SmbStatus3
type SmbStatus3 struct {
	Verbose *bool `json:"verbose,omitempty"`
	Fast *bool `json:"fast,omitempty"`
	RestrictUser *string `json:"restrict_user,omitempty"`
}

// NewSmbStatus3 instantiates a new SmbStatus3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmbStatus3() *SmbStatus3 {
	this := SmbStatus3{}
	return &this
}

// NewSmbStatus3WithDefaults instantiates a new SmbStatus3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmbStatus3WithDefaults() *SmbStatus3 {
	this := SmbStatus3{}
	return &this
}

// GetVerbose returns the Verbose field value if set, zero value otherwise.
func (o *SmbStatus3) GetVerbose() bool {
	if o == nil || o.Verbose == nil {
		var ret bool
		return ret
	}
	return *o.Verbose
}

// GetVerboseOk returns a tuple with the Verbose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus3) GetVerboseOk() (*bool, bool) {
	if o == nil || o.Verbose == nil {
		return nil, false
	}
	return o.Verbose, true
}

// HasVerbose returns a boolean if a field has been set.
func (o *SmbStatus3) HasVerbose() bool {
	if o != nil && o.Verbose != nil {
		return true
	}

	return false
}

// SetVerbose gets a reference to the given bool and assigns it to the Verbose field.
func (o *SmbStatus3) SetVerbose(v bool) {
	o.Verbose = &v
}

// GetFast returns the Fast field value if set, zero value otherwise.
func (o *SmbStatus3) GetFast() bool {
	if o == nil || o.Fast == nil {
		var ret bool
		return ret
	}
	return *o.Fast
}

// GetFastOk returns a tuple with the Fast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus3) GetFastOk() (*bool, bool) {
	if o == nil || o.Fast == nil {
		return nil, false
	}
	return o.Fast, true
}

// HasFast returns a boolean if a field has been set.
func (o *SmbStatus3) HasFast() bool {
	if o != nil && o.Fast != nil {
		return true
	}

	return false
}

// SetFast gets a reference to the given bool and assigns it to the Fast field.
func (o *SmbStatus3) SetFast(v bool) {
	o.Fast = &v
}

// GetRestrictUser returns the RestrictUser field value if set, zero value otherwise.
func (o *SmbStatus3) GetRestrictUser() string {
	if o == nil || o.RestrictUser == nil {
		var ret string
		return ret
	}
	return *o.RestrictUser
}

// GetRestrictUserOk returns a tuple with the RestrictUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmbStatus3) GetRestrictUserOk() (*string, bool) {
	if o == nil || o.RestrictUser == nil {
		return nil, false
	}
	return o.RestrictUser, true
}

// HasRestrictUser returns a boolean if a field has been set.
func (o *SmbStatus3) HasRestrictUser() bool {
	if o != nil && o.RestrictUser != nil {
		return true
	}

	return false
}

// SetRestrictUser gets a reference to the given string and assigns it to the RestrictUser field.
func (o *SmbStatus3) SetRestrictUser(v string) {
	o.RestrictUser = &v
}

func (o SmbStatus3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Verbose != nil {
		toSerialize["verbose"] = o.Verbose
	}
	if o.Fast != nil {
		toSerialize["fast"] = o.Fast
	}
	if o.RestrictUser != nil {
		toSerialize["restrict_user"] = o.RestrictUser
	}
	return json.Marshal(toSerialize)
}

type NullableSmbStatus3 struct {
	value *SmbStatus3
	isSet bool
}

func (v NullableSmbStatus3) Get() *SmbStatus3 {
	return v.value
}

func (v *NullableSmbStatus3) Set(val *SmbStatus3) {
	v.value = val
	v.isSet = true
}

func (v NullableSmbStatus3) IsSet() bool {
	return v.isSet
}

func (v *NullableSmbStatus3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmbStatus3(val *SmbStatus3) *NullableSmbStatus3 {
	return &NullableSmbStatus3{value: val, isSet: true}
}

func (v NullableSmbStatus3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmbStatus3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

