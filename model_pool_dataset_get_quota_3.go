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

// PoolDatasetGetQuota3 struct for PoolDatasetGetQuota3
type PoolDatasetGetQuota3 struct {
	Relationships *bool `json:"relationships,omitempty"`
	Extend NullableString `json:"extend,omitempty"`
	ExtendContext NullableString `json:"extend_context,omitempty"`
	Prefix NullableString `json:"prefix,omitempty"`
	Extra *map[string]map[string]interface{} `json:"extra,omitempty"`
	OrderBy *[]interface{} `json:"order_by,omitempty"`
	Select *[]interface{} `json:"select,omitempty"`
	Count *bool `json:"count,omitempty"`
	Get *bool `json:"get,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
}

// NewPoolDatasetGetQuota3 instantiates a new PoolDatasetGetQuota3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolDatasetGetQuota3() *PoolDatasetGetQuota3 {
	this := PoolDatasetGetQuota3{}
	return &this
}

// NewPoolDatasetGetQuota3WithDefaults instantiates a new PoolDatasetGetQuota3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolDatasetGetQuota3WithDefaults() *PoolDatasetGetQuota3 {
	this := PoolDatasetGetQuota3{}
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetRelationships() bool {
	if o == nil || o.Relationships == nil {
		var ret bool
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetRelationshipsOk() (*bool, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given bool and assigns it to the Relationships field.
func (o *PoolDatasetGetQuota3) SetRelationships(v bool) {
	o.Relationships = &v
}

// GetExtend returns the Extend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolDatasetGetQuota3) GetExtend() string {
	if o == nil || o.Extend.Get() == nil {
		var ret string
		return ret
	}
	return *o.Extend.Get()
}

// GetExtendOk returns a tuple with the Extend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolDatasetGetQuota3) GetExtendOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Extend.Get(), o.Extend.IsSet()
}

// HasExtend returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasExtend() bool {
	if o != nil && o.Extend.IsSet() {
		return true
	}

	return false
}

// SetExtend gets a reference to the given NullableString and assigns it to the Extend field.
func (o *PoolDatasetGetQuota3) SetExtend(v string) {
	o.Extend.Set(&v)
}
// SetExtendNil sets the value for Extend to be an explicit nil
func (o *PoolDatasetGetQuota3) SetExtendNil() {
	o.Extend.Set(nil)
}

// UnsetExtend ensures that no value is present for Extend, not even an explicit nil
func (o *PoolDatasetGetQuota3) UnsetExtend() {
	o.Extend.Unset()
}

// GetExtendContext returns the ExtendContext field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolDatasetGetQuota3) GetExtendContext() string {
	if o == nil || o.ExtendContext.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExtendContext.Get()
}

// GetExtendContextOk returns a tuple with the ExtendContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolDatasetGetQuota3) GetExtendContextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExtendContext.Get(), o.ExtendContext.IsSet()
}

// HasExtendContext returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasExtendContext() bool {
	if o != nil && o.ExtendContext.IsSet() {
		return true
	}

	return false
}

// SetExtendContext gets a reference to the given NullableString and assigns it to the ExtendContext field.
func (o *PoolDatasetGetQuota3) SetExtendContext(v string) {
	o.ExtendContext.Set(&v)
}
// SetExtendContextNil sets the value for ExtendContext to be an explicit nil
func (o *PoolDatasetGetQuota3) SetExtendContextNil() {
	o.ExtendContext.Set(nil)
}

// UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
func (o *PoolDatasetGetQuota3) UnsetExtendContext() {
	o.ExtendContext.Unset()
}

// GetPrefix returns the Prefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PoolDatasetGetQuota3) GetPrefix() string {
	if o == nil || o.Prefix.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prefix.Get()
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PoolDatasetGetQuota3) GetPrefixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prefix.Get(), o.Prefix.IsSet()
}

// HasPrefix returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasPrefix() bool {
	if o != nil && o.Prefix.IsSet() {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given NullableString and assigns it to the Prefix field.
func (o *PoolDatasetGetQuota3) SetPrefix(v string) {
	o.Prefix.Set(&v)
}
// SetPrefixNil sets the value for Prefix to be an explicit nil
func (o *PoolDatasetGetQuota3) SetPrefixNil() {
	o.Prefix.Set(nil)
}

// UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
func (o *PoolDatasetGetQuota3) UnsetPrefix() {
	o.Prefix.Unset()
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetExtra() map[string]map[string]interface{} {
	if o == nil || o.Extra == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetExtraOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given map[string]map[string]interface{} and assigns it to the Extra field.
func (o *PoolDatasetGetQuota3) SetExtra(v map[string]map[string]interface{}) {
	o.Extra = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetOrderBy() []interface{} {
	if o == nil || o.OrderBy == nil {
		var ret []interface{}
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetOrderByOk() (*[]interface{}, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []interface{} and assigns it to the OrderBy field.
func (o *PoolDatasetGetQuota3) SetOrderBy(v []interface{}) {
	o.OrderBy = &v
}

// GetSelect returns the Select field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetSelect() []interface{} {
	if o == nil || o.Select == nil {
		var ret []interface{}
		return ret
	}
	return *o.Select
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetSelectOk() (*[]interface{}, bool) {
	if o == nil || o.Select == nil {
		return nil, false
	}
	return o.Select, true
}

// HasSelect returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasSelect() bool {
	if o != nil && o.Select != nil {
		return true
	}

	return false
}

// SetSelect gets a reference to the given []interface{} and assigns it to the Select field.
func (o *PoolDatasetGetQuota3) SetSelect(v []interface{}) {
	o.Select = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetCount() bool {
	if o == nil || o.Count == nil {
		var ret bool
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetCountOk() (*bool, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given bool and assigns it to the Count field.
func (o *PoolDatasetGetQuota3) SetCount(v bool) {
	o.Count = &v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetGet() bool {
	if o == nil || o.Get == nil {
		var ret bool
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetGetOk() (*bool, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given bool and assigns it to the Get field.
func (o *PoolDatasetGetQuota3) SetGet(v bool) {
	o.Get = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PoolDatasetGetQuota3) SetOffset(v int32) {
	o.Offset = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PoolDatasetGetQuota3) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PoolDatasetGetQuota3) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PoolDatasetGetQuota3) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PoolDatasetGetQuota3) SetLimit(v int32) {
	o.Limit = &v
}

func (o PoolDatasetGetQuota3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Extend.IsSet() {
		toSerialize["extend"] = o.Extend.Get()
	}
	if o.ExtendContext.IsSet() {
		toSerialize["extend_context"] = o.ExtendContext.Get()
	}
	if o.Prefix.IsSet() {
		toSerialize["prefix"] = o.Prefix.Get()
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.OrderBy != nil {
		toSerialize["order_by"] = o.OrderBy
	}
	if o.Select != nil {
		toSerialize["select"] = o.Select
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	return json.Marshal(toSerialize)
}

type NullablePoolDatasetGetQuota3 struct {
	value *PoolDatasetGetQuota3
	isSet bool
}

func (v NullablePoolDatasetGetQuota3) Get() *PoolDatasetGetQuota3 {
	return v.value
}

func (v *NullablePoolDatasetGetQuota3) Set(val *PoolDatasetGetQuota3) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolDatasetGetQuota3) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolDatasetGetQuota3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolDatasetGetQuota3(val *PoolDatasetGetQuota3) *NullablePoolDatasetGetQuota3 {
	return &NullablePoolDatasetGetQuota3{value: val, isSet: true}
}

func (v NullablePoolDatasetGetQuota3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolDatasetGetQuota3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

