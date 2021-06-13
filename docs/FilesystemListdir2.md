# FilesystemListdir2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Relationships** | Pointer to **bool** |  | [optional] 
**Extend** | Pointer to **NullableString** |  | [optional] 
**ExtendContext** | Pointer to **NullableString** |  | [optional] 
**Prefix** | Pointer to **NullableString** |  | [optional] 
**Extra** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**OrderBy** | Pointer to **[]interface{}** |  | [optional] 
**Select** | Pointer to **[]interface{}** |  | [optional] 
**Count** | Pointer to **bool** |  | [optional] 
**Get** | Pointer to **bool** |  | [optional] 
**Offset** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewFilesystemListdir2

`func NewFilesystemListdir2() *FilesystemListdir2`

NewFilesystemListdir2 instantiates a new FilesystemListdir2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilesystemListdir2WithDefaults

`func NewFilesystemListdir2WithDefaults() *FilesystemListdir2`

NewFilesystemListdir2WithDefaults instantiates a new FilesystemListdir2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelationships

`func (o *FilesystemListdir2) GetRelationships() bool`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *FilesystemListdir2) GetRelationshipsOk() (*bool, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *FilesystemListdir2) SetRelationships(v bool)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *FilesystemListdir2) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetExtend

`func (o *FilesystemListdir2) GetExtend() string`

GetExtend returns the Extend field if non-nil, zero value otherwise.

### GetExtendOk

`func (o *FilesystemListdir2) GetExtendOk() (*string, bool)`

GetExtendOk returns a tuple with the Extend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtend

`func (o *FilesystemListdir2) SetExtend(v string)`

SetExtend sets Extend field to given value.

### HasExtend

`func (o *FilesystemListdir2) HasExtend() bool`

HasExtend returns a boolean if a field has been set.

### SetExtendNil

`func (o *FilesystemListdir2) SetExtendNil(b bool)`

 SetExtendNil sets the value for Extend to be an explicit nil

### UnsetExtend
`func (o *FilesystemListdir2) UnsetExtend()`

UnsetExtend ensures that no value is present for Extend, not even an explicit nil
### GetExtendContext

`func (o *FilesystemListdir2) GetExtendContext() string`

GetExtendContext returns the ExtendContext field if non-nil, zero value otherwise.

### GetExtendContextOk

`func (o *FilesystemListdir2) GetExtendContextOk() (*string, bool)`

GetExtendContextOk returns a tuple with the ExtendContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendContext

`func (o *FilesystemListdir2) SetExtendContext(v string)`

SetExtendContext sets ExtendContext field to given value.

### HasExtendContext

`func (o *FilesystemListdir2) HasExtendContext() bool`

HasExtendContext returns a boolean if a field has been set.

### SetExtendContextNil

`func (o *FilesystemListdir2) SetExtendContextNil(b bool)`

 SetExtendContextNil sets the value for ExtendContext to be an explicit nil

### UnsetExtendContext
`func (o *FilesystemListdir2) UnsetExtendContext()`

UnsetExtendContext ensures that no value is present for ExtendContext, not even an explicit nil
### GetPrefix

`func (o *FilesystemListdir2) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FilesystemListdir2) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FilesystemListdir2) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *FilesystemListdir2) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### SetPrefixNil

`func (o *FilesystemListdir2) SetPrefixNil(b bool)`

 SetPrefixNil sets the value for Prefix to be an explicit nil

### UnsetPrefix
`func (o *FilesystemListdir2) UnsetPrefix()`

UnsetPrefix ensures that no value is present for Prefix, not even an explicit nil
### GetExtra

`func (o *FilesystemListdir2) GetExtra() map[string]map[string]interface{}`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *FilesystemListdir2) GetExtraOk() (*map[string]map[string]interface{}, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *FilesystemListdir2) SetExtra(v map[string]map[string]interface{})`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *FilesystemListdir2) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetOrderBy

`func (o *FilesystemListdir2) GetOrderBy() []interface{}`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *FilesystemListdir2) GetOrderByOk() (*[]interface{}, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *FilesystemListdir2) SetOrderBy(v []interface{})`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *FilesystemListdir2) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetSelect

`func (o *FilesystemListdir2) GetSelect() []interface{}`

GetSelect returns the Select field if non-nil, zero value otherwise.

### GetSelectOk

`func (o *FilesystemListdir2) GetSelectOk() (*[]interface{}, bool)`

GetSelectOk returns a tuple with the Select field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelect

`func (o *FilesystemListdir2) SetSelect(v []interface{})`

SetSelect sets Select field to given value.

### HasSelect

`func (o *FilesystemListdir2) HasSelect() bool`

HasSelect returns a boolean if a field has been set.

### GetCount

`func (o *FilesystemListdir2) GetCount() bool`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FilesystemListdir2) GetCountOk() (*bool, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FilesystemListdir2) SetCount(v bool)`

SetCount sets Count field to given value.

### HasCount

`func (o *FilesystemListdir2) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetGet

`func (o *FilesystemListdir2) GetGet() bool`

GetGet returns the Get field if non-nil, zero value otherwise.

### GetGetOk

`func (o *FilesystemListdir2) GetGetOk() (*bool, bool)`

GetGetOk returns a tuple with the Get field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGet

`func (o *FilesystemListdir2) SetGet(v bool)`

SetGet sets Get field to given value.

### HasGet

`func (o *FilesystemListdir2) HasGet() bool`

HasGet returns a boolean if a field has been set.

### GetOffset

`func (o *FilesystemListdir2) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *FilesystemListdir2) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *FilesystemListdir2) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *FilesystemListdir2) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *FilesystemListdir2) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FilesystemListdir2) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FilesystemListdir2) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *FilesystemListdir2) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

