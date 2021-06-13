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

// AcmeDnsAuthenticatorCreate0 struct for AcmeDnsAuthenticatorCreate0
type AcmeDnsAuthenticatorCreate0 struct {
	Authenticator *string `json:"authenticator,omitempty"`
	Name *string `json:"name,omitempty"`
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
}

// NewAcmeDnsAuthenticatorCreate0 instantiates a new AcmeDnsAuthenticatorCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcmeDnsAuthenticatorCreate0() *AcmeDnsAuthenticatorCreate0 {
	this := AcmeDnsAuthenticatorCreate0{}
	return &this
}

// NewAcmeDnsAuthenticatorCreate0WithDefaults instantiates a new AcmeDnsAuthenticatorCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcmeDnsAuthenticatorCreate0WithDefaults() *AcmeDnsAuthenticatorCreate0 {
	this := AcmeDnsAuthenticatorCreate0{}
	return &this
}

// GetAuthenticator returns the Authenticator field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetAuthenticator() string {
	if o == nil || o.Authenticator == nil {
		var ret string
		return ret
	}
	return *o.Authenticator
}

// GetAuthenticatorOk returns a tuple with the Authenticator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetAuthenticatorOk() (*string, bool) {
	if o == nil || o.Authenticator == nil {
		return nil, false
	}
	return o.Authenticator, true
}

// HasAuthenticator returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasAuthenticator() bool {
	if o != nil && o.Authenticator != nil {
		return true
	}

	return false
}

// SetAuthenticator gets a reference to the given string and assigns it to the Authenticator field.
func (o *AcmeDnsAuthenticatorCreate0) SetAuthenticator(v string) {
	o.Authenticator = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AcmeDnsAuthenticatorCreate0) SetName(v string) {
	o.Name = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AcmeDnsAuthenticatorCreate0) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcmeDnsAuthenticatorCreate0) GetAttributesOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AcmeDnsAuthenticatorCreate0) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *AcmeDnsAuthenticatorCreate0) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

func (o AcmeDnsAuthenticatorCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Authenticator != nil {
		toSerialize["authenticator"] = o.Authenticator
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableAcmeDnsAuthenticatorCreate0 struct {
	value *AcmeDnsAuthenticatorCreate0
	isSet bool
}

func (v NullableAcmeDnsAuthenticatorCreate0) Get() *AcmeDnsAuthenticatorCreate0 {
	return v.value
}

func (v *NullableAcmeDnsAuthenticatorCreate0) Set(val *AcmeDnsAuthenticatorCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableAcmeDnsAuthenticatorCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableAcmeDnsAuthenticatorCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcmeDnsAuthenticatorCreate0(val *AcmeDnsAuthenticatorCreate0) *NullableAcmeDnsAuthenticatorCreate0 {
	return &NullableAcmeDnsAuthenticatorCreate0{value: val, isSet: true}
}

func (v NullableAcmeDnsAuthenticatorCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcmeDnsAuthenticatorCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

