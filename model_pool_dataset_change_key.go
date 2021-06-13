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

// PoolDatasetChangeKey struct for PoolDatasetChangeKey
type PoolDatasetChangeKey struct {
	Id *string `json:"id,omitempty"`
	ChangeKeyOptions *PoolDatasetChangeKey1 `json:"change_key_options,omitempty"`
}

// NewPoolDatasetChangeKey instantiates a new PoolDatasetChangeKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetChangeKey() *PoolDatasetChangeKey {
	this := PoolDatasetChangeKey{}
	var changeKeyOptions PoolDatasetChangeKey1 = {}
	this.ChangeKeyOptions = &changeKeyOptions
	return &this
}

// NewPoolDatasetChangeKeyWithDefaults instantiates a new PoolDatasetChangeKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetChangeKeyWithDefaults() *PoolDatasetChangeKey {
	this := PoolDatasetChangeKey{}
	var changeKeyOptions PoolDatasetChangeKey1 = {}
	this.ChangeKeyOptions = &changeKeyOptions
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PoolDatasetChangeKey) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetChangeKey) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PoolDatasetChangeKey) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PoolDatasetChangeKey) SetId(v string) {
	o.Id = &v
}

// GetChangeKeyOptions returns the ChangeKeyOptions field value if set, zero value otherwise.
func (o *PoolDatasetChangeKey) GetChangeKeyOptions() PoolDatasetChangeKey1 {
	if o == nil || o.ChangeKeyOptions == nil {
		var ret PoolDatasetChangeKey1
		return ret
	}
	return *o.ChangeKeyOptions
}

// GetChangeKeyOptionsOk returns a tuple with the ChangeKeyOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetChangeKey) GetChangeKeyOptionsOk() (*PoolDatasetChangeKey1, bool) {
	if o == nil || o.ChangeKeyOptions == nil {
		return nil, false
	}
	return o.ChangeKeyOptions, true
}

// HasChangeKeyOptions returns a boolean if a field has been set.
func (o *PoolDatasetChangeKey) HasChangeKeyOptions() bool {
	if o != nil && o.ChangeKeyOptions != nil {
		return true
	}

	return false
}

// SetChangeKeyOptions gets a reference to the given PoolDatasetChangeKey1 and assigns it to the ChangeKeyOptions field.
func (o *PoolDatasetChangeKey) SetChangeKeyOptions(v PoolDatasetChangeKey1) {
	o.ChangeKeyOptions = &v
}

func (o PoolDatasetChangeKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ChangeKeyOptions != nil {
		toSerialize["change_key_options"] = o.ChangeKeyOptions
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetChangeKey struct {
	value *PoolDatasetChangeKey
	isSet bool
}

func (v NullablePoolDatasetChangeKey) Get() *PoolDatasetChangeKey {
	return v.value
}

func (v *NullablePoolDatasetChangeKey) Set(val *PoolDatasetChangeKey) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetChangeKey) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetChangeKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetChangeKey(val *PoolDatasetChangeKey) *NullablePoolDatasetChangeKey {
	return &NullablePoolDatasetChangeKey{value: val, isSet: true}
}

func (v NullablePoolDatasetChangeKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetChangeKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

