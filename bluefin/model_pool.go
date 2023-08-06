/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// checks if the Pool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pool{}

// Pool struct for Pool
type Pool struct {
	Id                   int32   `json:"id"`
	Name                 string  `json:"name"`
	Guid                 *string `json:"guid,omitempty"`
	Path                 string  `json:"path"`
	Status               *string `json:"status,omitempty"`
	Healthy              *bool   `json:"healthy,omitempty"`
	IsDecrypted          *bool   `json:"is_decrypted,omitempty"`
	EncryptkeyPath       *string `json:"encryptkey_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Pool Pool

// NewPool instantiates a new Pool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPool(id int32, name string, path string) *Pool {
	this := Pool{}
	this.Id = id
	this.Name = name
	this.Path = path
	return &this
}

// NewPoolWithDefaults instantiates a new Pool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolWithDefaults() *Pool {
	this := Pool{}
	return &this
}

// GetId returns the Id field value
func (o *Pool) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pool) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pool) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Pool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Pool) SetName(v string) {
	o.Name = v
}

// GetGuid returns the Guid field value if set, zero value otherwise.
func (o *Pool) GetGuid() string {
	if o == nil || IsNil(o.Guid) {
		var ret string
		return ret
	}
	return *o.Guid
}

// GetGuidOk returns a tuple with the Guid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetGuidOk() (*string, bool) {
	if o == nil || IsNil(o.Guid) {
		return nil, false
	}
	return o.Guid, true
}

// HasGuid returns a boolean if a field has been set.
func (o *Pool) HasGuid() bool {
	if o != nil && !IsNil(o.Guid) {
		return true
	}

	return false
}

// SetGuid gets a reference to the given string and assigns it to the Guid field.
func (o *Pool) SetGuid(v string) {
	o.Guid = &v
}

// GetPath returns the Path field value
func (o *Pool) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Pool) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Pool) SetPath(v string) {
	o.Path = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Pool) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Pool) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Pool) SetStatus(v string) {
	o.Status = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *Pool) GetHealthy() bool {
	if o == nil || IsNil(o.Healthy) {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetHealthyOk() (*bool, bool) {
	if o == nil || IsNil(o.Healthy) {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *Pool) HasHealthy() bool {
	if o != nil && !IsNil(o.Healthy) {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *Pool) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetIsDecrypted returns the IsDecrypted field value if set, zero value otherwise.
func (o *Pool) GetIsDecrypted() bool {
	if o == nil || IsNil(o.IsDecrypted) {
		var ret bool
		return ret
	}
	return *o.IsDecrypted
}

// GetIsDecryptedOk returns a tuple with the IsDecrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetIsDecryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDecrypted) {
		return nil, false
	}
	return o.IsDecrypted, true
}

// HasIsDecrypted returns a boolean if a field has been set.
func (o *Pool) HasIsDecrypted() bool {
	if o != nil && !IsNil(o.IsDecrypted) {
		return true
	}

	return false
}

// SetIsDecrypted gets a reference to the given bool and assigns it to the IsDecrypted field.
func (o *Pool) SetIsDecrypted(v bool) {
	o.IsDecrypted = &v
}

// GetEncryptkeyPath returns the EncryptkeyPath field value if set, zero value otherwise.
func (o *Pool) GetEncryptkeyPath() string {
	if o == nil || IsNil(o.EncryptkeyPath) {
		var ret string
		return ret
	}
	return *o.EncryptkeyPath
}

// GetEncryptkeyPathOk returns a tuple with the EncryptkeyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pool) GetEncryptkeyPathOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptkeyPath) {
		return nil, false
	}
	return o.EncryptkeyPath, true
}

// HasEncryptkeyPath returns a boolean if a field has been set.
func (o *Pool) HasEncryptkeyPath() bool {
	if o != nil && !IsNil(o.EncryptkeyPath) {
		return true
	}

	return false
}

// SetEncryptkeyPath gets a reference to the given string and assigns it to the EncryptkeyPath field.
func (o *Pool) SetEncryptkeyPath(v string) {
	o.EncryptkeyPath = &v
}

func (o Pool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Guid) {
		toSerialize["guid"] = o.Guid
	}
	toSerialize["path"] = o.Path
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Healthy) {
		toSerialize["healthy"] = o.Healthy
	}
	if !IsNil(o.IsDecrypted) {
		toSerialize["is_decrypted"] = o.IsDecrypted
	}
	if !IsNil(o.EncryptkeyPath) {
		toSerialize["encryptkey_path"] = o.EncryptkeyPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Pool) UnmarshalJSON(bytes []byte) (err error) {
	varPool := _Pool{}

	if err = json.Unmarshal(bytes, &varPool); err == nil {
		*o = Pool(varPool)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "guid")
		delete(additionalProperties, "path")
		delete(additionalProperties, "status")
		delete(additionalProperties, "healthy")
		delete(additionalProperties, "is_decrypted")
		delete(additionalProperties, "encryptkey_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePool struct {
	value *Pool
	isSet bool
}

func (v NullablePool) Get() *Pool {
	return v.value
}

func (v *NullablePool) Set(val *Pool) {
	v.value = val
	v.isSet = true
}

func (v NullablePool) IsSet() bool {
	return v.isSet
}

func (v *NullablePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePool(val *Pool) *NullablePool {
	return &NullablePool{value: val, isSet: true}
}

func (v NullablePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
