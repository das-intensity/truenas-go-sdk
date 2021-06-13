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

// PoolUnlock1 struct for PoolUnlock1
type PoolUnlock1 struct {
	Passphrase *string `json:"passphrase,omitempty"`
	Recoverykey *bool `json:"recoverykey,omitempty"`
	ServicesRestart *[]interface{} `json:"services_restart,omitempty"`
}

// NewPoolUnlock1 instantiates a new PoolUnlock1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolUnlock1() *PoolUnlock1 {
	this := PoolUnlock1{}
	return &this
}

// NewPoolUnlock1WithDefaults instantiates a new PoolUnlock1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolUnlock1WithDefaults() *PoolUnlock1 {
	this := PoolUnlock1{}
	return &this
}

// GetPassphrase returns the Passphrase field value if set, zero value otherwise.
func (o *PoolUnlock1) GetPassphrase() string {
	if o == nil || o.Passphrase == nil {
		var ret string
		return ret
	}
	return *o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUnlock1) GetPassphraseOk() (*string, bool) {
	if o == nil || o.Passphrase == nil {
		return nil, false
	}
	return o.Passphrase, true
}

// HasPassphrase returns a boolean if a field has been set.
func (o *PoolUnlock1) HasPassphrase() bool {
	if o != nil && o.Passphrase != nil {
		return true
	}

	return false
}

// SetPassphrase gets a reference to the given string and assigns it to the Passphrase field.
func (o *PoolUnlock1) SetPassphrase(v string) {
	o.Passphrase = &v
}

// GetRecoverykey returns the Recoverykey field value if set, zero value otherwise.
func (o *PoolUnlock1) GetRecoverykey() bool {
	if o == nil || o.Recoverykey == nil {
		var ret bool
		return ret
	}
	return *o.Recoverykey
}

// GetRecoverykeyOk returns a tuple with the Recoverykey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUnlock1) GetRecoverykeyOk() (*bool, bool) {
	if o == nil || o.Recoverykey == nil {
		return nil, false
	}
	return o.Recoverykey, true
}

// HasRecoverykey returns a boolean if a field has been set.
func (o *PoolUnlock1) HasRecoverykey() bool {
	if o != nil && o.Recoverykey != nil {
		return true
	}

	return false
}

// SetRecoverykey gets a reference to the given bool and assigns it to the Recoverykey field.
func (o *PoolUnlock1) SetRecoverykey(v bool) {
	o.Recoverykey = &v
}

// GetServicesRestart returns the ServicesRestart field value if set, zero value otherwise.
func (o *PoolUnlock1) GetServicesRestart() []interface{} {
	if o == nil || o.ServicesRestart == nil {
		var ret []interface{}
		return ret
	}
	return *o.ServicesRestart
}

// GetServicesRestartOk returns a tuple with the ServicesRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolUnlock1) GetServicesRestartOk() (*[]interface{}, bool) {
	if o == nil || o.ServicesRestart == nil {
		return nil, false
	}
	return o.ServicesRestart, true
}

// HasServicesRestart returns a boolean if a field has been set.
func (o *PoolUnlock1) HasServicesRestart() bool {
	if o != nil && o.ServicesRestart != nil {
		return true
	}

	return false
}

// SetServicesRestart gets a reference to the given []interface{} and assigns it to the ServicesRestart field.
func (o *PoolUnlock1) SetServicesRestart(v []interface{}) {
	o.ServicesRestart = &v
}

func (o PoolUnlock1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Passphrase != nil {
		toSerialize["passphrase"] = o.Passphrase
	}
	if o.Recoverykey != nil {
		toSerialize["recoverykey"] = o.Recoverykey
	}
	if o.ServicesRestart != nil {
		toSerialize["services_restart"] = o.ServicesRestart
	}
	return json.Marshal(toSerialize)
}

type NullablePoolUnlock1 struct {
	value *PoolUnlock1
	isSet bool
}

func (v NullablePoolUnlock1) Get() *PoolUnlock1 {
	return v.value
}

func (v *NullablePoolUnlock1) Set(val *PoolUnlock1) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolUnlock1) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolUnlock1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolUnlock1(val *PoolUnlock1) *NullablePoolUnlock1 {
	return &NullablePoolUnlock1{value: val, isSet: true}
}

func (v NullablePoolUnlock1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolUnlock1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

