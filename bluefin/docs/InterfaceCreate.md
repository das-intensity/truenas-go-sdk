# InterfaceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [default to ""]
**Type** | Pointer to **string** | For BRIDGE &#x60;type&#x60; the following attribute is required: bridge_members. For LINK_AGGREGATION &#x60;type&#x60; the following attributes are required: lag_ports, lag_protocol. For VLAN &#x60;type&#x60; the following attributes are required: vlan_parent_interface, vlan_tag and vlan_pcp. | [optional] 
**Ipv4Dhcp** | Pointer to **bool** |  | [optional] [default to false]
**Ipv6Auto** | Pointer to **bool** |  | [optional] [default to false]
**Aliases** | Pointer to [**[]InterfaceAlias**](InterfaceAlias.md) |  | [optional] [default to []]
**FailoverCritical** | Pointer to **bool** |  | [optional] [default to false]
**FailoverGroup** | Pointer to **NullableInt32** |  | [optional] 
**FailoverVhid** | Pointer to **NullableInt32** |  | [optional] 
**FailoverAliases** | Pointer to [**[]InterfaceFailoverAlias**](InterfaceFailoverAlias.md) |  | [optional] [default to []]
**FailoverVirtualAliases** | Pointer to [**[]InterfaceVirtualAlias**](InterfaceVirtualAlias.md) |  | [optional] [default to []]
**BridgeMembers** | Pointer to **[]interface{}** |  | [optional] [default to []]
**Stp** | Pointer to **bool** |  | [optional] [default to true]
**LagProtocol** | Pointer to **string** |  | [optional] 
**XmitHashPolicy** | Pointer to **NullableString** |  | [optional] 
**LacpduRate** | Pointer to **NullableString** |  | [optional] 
**LagPorts** | Pointer to **[]string** |  | [optional] [default to []]
**VlanParentInterface** | Pointer to **string** |  | [optional] 
**VlanTag** | Pointer to **int32** |  | [optional] 
**VlanPcp** | Pointer to **NullableInt32** |  | [optional] 
**Mtu** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewInterfaceCreate

`func NewInterfaceCreate() *InterfaceCreate`

NewInterfaceCreate instantiates a new InterfaceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceCreateWithDefaults

`func NewInterfaceCreateWithDefaults() *InterfaceCreate`

NewInterfaceCreateWithDefaults instantiates a new InterfaceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InterfaceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterfaceCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InterfaceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InterfaceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InterfaceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InterfaceCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *InterfaceCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InterfaceCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InterfaceCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InterfaceCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpv4Dhcp

`func (o *InterfaceCreate) GetIpv4Dhcp() bool`

GetIpv4Dhcp returns the Ipv4Dhcp field if non-nil, zero value otherwise.

### GetIpv4DhcpOk

`func (o *InterfaceCreate) GetIpv4DhcpOk() (*bool, bool)`

GetIpv4DhcpOk returns a tuple with the Ipv4Dhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4Dhcp

`func (o *InterfaceCreate) SetIpv4Dhcp(v bool)`

SetIpv4Dhcp sets Ipv4Dhcp field to given value.

### HasIpv4Dhcp

`func (o *InterfaceCreate) HasIpv4Dhcp() bool`

HasIpv4Dhcp returns a boolean if a field has been set.

### GetIpv6Auto

`func (o *InterfaceCreate) GetIpv6Auto() bool`

GetIpv6Auto returns the Ipv6Auto field if non-nil, zero value otherwise.

### GetIpv6AutoOk

`func (o *InterfaceCreate) GetIpv6AutoOk() (*bool, bool)`

GetIpv6AutoOk returns a tuple with the Ipv6Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Auto

`func (o *InterfaceCreate) SetIpv6Auto(v bool)`

SetIpv6Auto sets Ipv6Auto field to given value.

### HasIpv6Auto

`func (o *InterfaceCreate) HasIpv6Auto() bool`

HasIpv6Auto returns a boolean if a field has been set.

### GetAliases

`func (o *InterfaceCreate) GetAliases() []InterfaceAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *InterfaceCreate) GetAliasesOk() (*[]InterfaceAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *InterfaceCreate) SetAliases(v []InterfaceAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *InterfaceCreate) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetFailoverCritical

`func (o *InterfaceCreate) GetFailoverCritical() bool`

GetFailoverCritical returns the FailoverCritical field if non-nil, zero value otherwise.

### GetFailoverCriticalOk

`func (o *InterfaceCreate) GetFailoverCriticalOk() (*bool, bool)`

GetFailoverCriticalOk returns a tuple with the FailoverCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverCritical

`func (o *InterfaceCreate) SetFailoverCritical(v bool)`

SetFailoverCritical sets FailoverCritical field to given value.

### HasFailoverCritical

`func (o *InterfaceCreate) HasFailoverCritical() bool`

HasFailoverCritical returns a boolean if a field has been set.

### GetFailoverGroup

`func (o *InterfaceCreate) GetFailoverGroup() int32`

GetFailoverGroup returns the FailoverGroup field if non-nil, zero value otherwise.

### GetFailoverGroupOk

`func (o *InterfaceCreate) GetFailoverGroupOk() (*int32, bool)`

GetFailoverGroupOk returns a tuple with the FailoverGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverGroup

`func (o *InterfaceCreate) SetFailoverGroup(v int32)`

SetFailoverGroup sets FailoverGroup field to given value.

### HasFailoverGroup

`func (o *InterfaceCreate) HasFailoverGroup() bool`

HasFailoverGroup returns a boolean if a field has been set.

### SetFailoverGroupNil

`func (o *InterfaceCreate) SetFailoverGroupNil(b bool)`

 SetFailoverGroupNil sets the value for FailoverGroup to be an explicit nil

### UnsetFailoverGroup
`func (o *InterfaceCreate) UnsetFailoverGroup()`

UnsetFailoverGroup ensures that no value is present for FailoverGroup, not even an explicit nil
### GetFailoverVhid

`func (o *InterfaceCreate) GetFailoverVhid() int32`

GetFailoverVhid returns the FailoverVhid field if non-nil, zero value otherwise.

### GetFailoverVhidOk

`func (o *InterfaceCreate) GetFailoverVhidOk() (*int32, bool)`

GetFailoverVhidOk returns a tuple with the FailoverVhid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVhid

`func (o *InterfaceCreate) SetFailoverVhid(v int32)`

SetFailoverVhid sets FailoverVhid field to given value.

### HasFailoverVhid

`func (o *InterfaceCreate) HasFailoverVhid() bool`

HasFailoverVhid returns a boolean if a field has been set.

### SetFailoverVhidNil

`func (o *InterfaceCreate) SetFailoverVhidNil(b bool)`

 SetFailoverVhidNil sets the value for FailoverVhid to be an explicit nil

### UnsetFailoverVhid
`func (o *InterfaceCreate) UnsetFailoverVhid()`

UnsetFailoverVhid ensures that no value is present for FailoverVhid, not even an explicit nil
### GetFailoverAliases

`func (o *InterfaceCreate) GetFailoverAliases() []InterfaceFailoverAlias`

GetFailoverAliases returns the FailoverAliases field if non-nil, zero value otherwise.

### GetFailoverAliasesOk

`func (o *InterfaceCreate) GetFailoverAliasesOk() (*[]InterfaceFailoverAlias, bool)`

GetFailoverAliasesOk returns a tuple with the FailoverAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAliases

`func (o *InterfaceCreate) SetFailoverAliases(v []InterfaceFailoverAlias)`

SetFailoverAliases sets FailoverAliases field to given value.

### HasFailoverAliases

`func (o *InterfaceCreate) HasFailoverAliases() bool`

HasFailoverAliases returns a boolean if a field has been set.

### GetFailoverVirtualAliases

`func (o *InterfaceCreate) GetFailoverVirtualAliases() []InterfaceVirtualAlias`

GetFailoverVirtualAliases returns the FailoverVirtualAliases field if non-nil, zero value otherwise.

### GetFailoverVirtualAliasesOk

`func (o *InterfaceCreate) GetFailoverVirtualAliasesOk() (*[]InterfaceVirtualAlias, bool)`

GetFailoverVirtualAliasesOk returns a tuple with the FailoverVirtualAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverVirtualAliases

`func (o *InterfaceCreate) SetFailoverVirtualAliases(v []InterfaceVirtualAlias)`

SetFailoverVirtualAliases sets FailoverVirtualAliases field to given value.

### HasFailoverVirtualAliases

`func (o *InterfaceCreate) HasFailoverVirtualAliases() bool`

HasFailoverVirtualAliases returns a boolean if a field has been set.

### GetBridgeMembers

`func (o *InterfaceCreate) GetBridgeMembers() []interface{}`

GetBridgeMembers returns the BridgeMembers field if non-nil, zero value otherwise.

### GetBridgeMembersOk

`func (o *InterfaceCreate) GetBridgeMembersOk() (*[]interface{}, bool)`

GetBridgeMembersOk returns a tuple with the BridgeMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeMembers

`func (o *InterfaceCreate) SetBridgeMembers(v []interface{})`

SetBridgeMembers sets BridgeMembers field to given value.

### HasBridgeMembers

`func (o *InterfaceCreate) HasBridgeMembers() bool`

HasBridgeMembers returns a boolean if a field has been set.

### GetStp

`func (o *InterfaceCreate) GetStp() bool`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *InterfaceCreate) GetStpOk() (*bool, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *InterfaceCreate) SetStp(v bool)`

SetStp sets Stp field to given value.

### HasStp

`func (o *InterfaceCreate) HasStp() bool`

HasStp returns a boolean if a field has been set.

### GetLagProtocol

`func (o *InterfaceCreate) GetLagProtocol() string`

GetLagProtocol returns the LagProtocol field if non-nil, zero value otherwise.

### GetLagProtocolOk

`func (o *InterfaceCreate) GetLagProtocolOk() (*string, bool)`

GetLagProtocolOk returns a tuple with the LagProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagProtocol

`func (o *InterfaceCreate) SetLagProtocol(v string)`

SetLagProtocol sets LagProtocol field to given value.

### HasLagProtocol

`func (o *InterfaceCreate) HasLagProtocol() bool`

HasLagProtocol returns a boolean if a field has been set.

### GetXmitHashPolicy

`func (o *InterfaceCreate) GetXmitHashPolicy() string`

GetXmitHashPolicy returns the XmitHashPolicy field if non-nil, zero value otherwise.

### GetXmitHashPolicyOk

`func (o *InterfaceCreate) GetXmitHashPolicyOk() (*string, bool)`

GetXmitHashPolicyOk returns a tuple with the XmitHashPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmitHashPolicy

`func (o *InterfaceCreate) SetXmitHashPolicy(v string)`

SetXmitHashPolicy sets XmitHashPolicy field to given value.

### HasXmitHashPolicy

`func (o *InterfaceCreate) HasXmitHashPolicy() bool`

HasXmitHashPolicy returns a boolean if a field has been set.

### SetXmitHashPolicyNil

`func (o *InterfaceCreate) SetXmitHashPolicyNil(b bool)`

 SetXmitHashPolicyNil sets the value for XmitHashPolicy to be an explicit nil

### UnsetXmitHashPolicy
`func (o *InterfaceCreate) UnsetXmitHashPolicy()`

UnsetXmitHashPolicy ensures that no value is present for XmitHashPolicy, not even an explicit nil
### GetLacpduRate

`func (o *InterfaceCreate) GetLacpduRate() string`

GetLacpduRate returns the LacpduRate field if non-nil, zero value otherwise.

### GetLacpduRateOk

`func (o *InterfaceCreate) GetLacpduRateOk() (*string, bool)`

GetLacpduRateOk returns a tuple with the LacpduRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLacpduRate

`func (o *InterfaceCreate) SetLacpduRate(v string)`

SetLacpduRate sets LacpduRate field to given value.

### HasLacpduRate

`func (o *InterfaceCreate) HasLacpduRate() bool`

HasLacpduRate returns a boolean if a field has been set.

### SetLacpduRateNil

`func (o *InterfaceCreate) SetLacpduRateNil(b bool)`

 SetLacpduRateNil sets the value for LacpduRate to be an explicit nil

### UnsetLacpduRate
`func (o *InterfaceCreate) UnsetLacpduRate()`

UnsetLacpduRate ensures that no value is present for LacpduRate, not even an explicit nil
### GetLagPorts

`func (o *InterfaceCreate) GetLagPorts() []string`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *InterfaceCreate) GetLagPortsOk() (*[]string, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *InterfaceCreate) SetLagPorts(v []string)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *InterfaceCreate) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetVlanParentInterface

`func (o *InterfaceCreate) GetVlanParentInterface() string`

GetVlanParentInterface returns the VlanParentInterface field if non-nil, zero value otherwise.

### GetVlanParentInterfaceOk

`func (o *InterfaceCreate) GetVlanParentInterfaceOk() (*string, bool)`

GetVlanParentInterfaceOk returns a tuple with the VlanParentInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanParentInterface

`func (o *InterfaceCreate) SetVlanParentInterface(v string)`

SetVlanParentInterface sets VlanParentInterface field to given value.

### HasVlanParentInterface

`func (o *InterfaceCreate) HasVlanParentInterface() bool`

HasVlanParentInterface returns a boolean if a field has been set.

### GetVlanTag

`func (o *InterfaceCreate) GetVlanTag() int32`

GetVlanTag returns the VlanTag field if non-nil, zero value otherwise.

### GetVlanTagOk

`func (o *InterfaceCreate) GetVlanTagOk() (*int32, bool)`

GetVlanTagOk returns a tuple with the VlanTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanTag

`func (o *InterfaceCreate) SetVlanTag(v int32)`

SetVlanTag sets VlanTag field to given value.

### HasVlanTag

`func (o *InterfaceCreate) HasVlanTag() bool`

HasVlanTag returns a boolean if a field has been set.

### GetVlanPcp

`func (o *InterfaceCreate) GetVlanPcp() int32`

GetVlanPcp returns the VlanPcp field if non-nil, zero value otherwise.

### GetVlanPcpOk

`func (o *InterfaceCreate) GetVlanPcpOk() (*int32, bool)`

GetVlanPcpOk returns a tuple with the VlanPcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPcp

`func (o *InterfaceCreate) SetVlanPcp(v int32)`

SetVlanPcp sets VlanPcp field to given value.

### HasVlanPcp

`func (o *InterfaceCreate) HasVlanPcp() bool`

HasVlanPcp returns a boolean if a field has been set.

### SetVlanPcpNil

`func (o *InterfaceCreate) SetVlanPcpNil(b bool)`

 SetVlanPcpNil sets the value for VlanPcp to be an explicit nil

### UnsetVlanPcp
`func (o *InterfaceCreate) UnsetVlanPcp()`

UnsetVlanPcp ensures that no value is present for VlanPcp, not even an explicit nil
### GetMtu

`func (o *InterfaceCreate) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *InterfaceCreate) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *InterfaceCreate) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *InterfaceCreate) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### SetMtuNil

`func (o *InterfaceCreate) SetMtuNil(b bool)`

 SetMtuNil sets the value for Mtu to be an explicit nil

### UnsetMtu
`func (o *InterfaceCreate) UnsetMtu()`

UnsetMtu ensures that no value is present for Mtu, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


