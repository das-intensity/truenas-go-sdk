# InterfaceAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "INET"]
**Address** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **int32** |  | [optional] 

## Methods

### NewInterfaceAlias

`func NewInterfaceAlias() *InterfaceAlias`

NewInterfaceAlias instantiates a new InterfaceAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceAliasWithDefaults

`func NewInterfaceAliasWithDefaults() *InterfaceAlias`

NewInterfaceAliasWithDefaults instantiates a new InterfaceAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InterfaceAlias) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceAlias) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceAlias) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InterfaceAlias) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAddress

`func (o *InterfaceAlias) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *InterfaceAlias) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *InterfaceAlias) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *InterfaceAlias) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNetmask

`func (o *InterfaceAlias) GetNetmask() int32`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *InterfaceAlias) GetNetmaskOk() (*int32, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *InterfaceAlias) SetNetmask(v int32)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *InterfaceAlias) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


