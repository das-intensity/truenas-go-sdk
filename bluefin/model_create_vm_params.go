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

// checks if the CreateVMParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMParams{}

// CreateVMParams struct for CreateVMParams
type CreateVMParams struct {
	CommandLineArgs *string        `json:"command_line_args,omitempty"`
	CpuMode         *string        `json:"cpu_mode,omitempty"`
	CpuModel        NullableString `json:"cpu_model,omitempty"`
	Name            *string        `json:"name,omitempty"`
	Description     *string        `json:"description,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Vcpus *int32 `json:"vcpus,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Cores *int32 `json:"cores,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Threads           *int32         `json:"threads,omitempty"`
	Cpuset            NullableString `json:"cpuset,omitempty"`
	Nodeset           NullableString `json:"nodeset,omitempty"`
	PinVcpus          *bool          `json:"pin_vcpus,omitempty"`
	SuspendOnSnapshot *bool          `json:"suspend_on_snapshot,omitempty"`
	Memory            *int64         `json:"memory,omitempty"`
	MinMemory         NullableInt64  `json:"min_memory,omitempty"`
	// `hyperv_enlightenments` can be used to enable subset of predefined Hyper-V enlightenments for Windows guests. These enlightenments improve performance and enable otherwise missing features.
	HypervEnlightenments *bool   `json:"hyperv_enlightenments,omitempty"`
	Bootloader           *string `json:"bootloader,omitempty"`
	Autostart            *bool   `json:"autostart,omitempty"`
	// `hide_from_msr` is a boolean which when set will hide the KVM hypervisor from standard MSR based discovery and is useful to enable when doing GPU passthrough.
	HideFromMsr *bool `json:"hide_from_msr,omitempty"`
	// `ensure_display_device` when set ( the default ) will ensure that the guest always has access to a video device. For headless installations like ubuntu server this is required for the guest to operate properly. However for cases where consumer would like to use GPU passthrough and does not want a display device added should set this to `false`.
	EnsureDisplayDevice *bool   `json:"ensure_display_device,omitempty"`
	Time                *string `json:"time,omitempty"`
	// `shutdown_timeout` indicates the time in seconds the system waits for the VM to cleanly shutdown. During system shutdown, if the VM hasn't exited after a hardware shutdown signal has been sent by the system within `shutdown_timeout` seconds, system initiates poweroff for the VM to stop it.
	ShutdownTimeout *int32 `json:"shutdown_timeout,omitempty"`
	// `arch_type` refers to architecture type and can be specified for the guest. By default the value is `null` and system in this case will choose a reasonable default based on host. `machine_type` refers to machine type of the guest based on the architecture type selected with `arch_type`. By default the value is `null` and system in this case will choose a reasonable default based on `arch_type` configuration.
	ArchType NullableString `json:"arch_type,omitempty"`
	// `machine_type` refers to machine type of the guest based on the architecture type selected with `arch_type`. By default the value is `null` and system in this case will choose a reasonable default based on `arch_type` configuration.
	MachineType NullableString `json:"machine_type,omitempty"`
	Uuid        NullableString `json:"uuid,omitempty"`
}

// NewCreateVMParams instantiates a new CreateVMParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMParams() *CreateVMParams {
	this := CreateVMParams{}
	var commandLineArgs string = ""
	this.CommandLineArgs = &commandLineArgs
	var cpuMode string = "CUSTOM"
	this.CpuMode = &cpuMode
	var vcpus int32 = 1
	this.Vcpus = &vcpus
	var cores int32 = 1
	this.Cores = &cores
	var threads int32 = 1
	this.Threads = &threads
	var pinVcpus bool = false
	this.PinVcpus = &pinVcpus
	var suspendOnSnapshot bool = false
	this.SuspendOnSnapshot = &suspendOnSnapshot
	var hypervEnlightenments bool = false
	this.HypervEnlightenments = &hypervEnlightenments
	var bootloader string = "UEFI"
	this.Bootloader = &bootloader
	var autostart bool = true
	this.Autostart = &autostart
	var hideFromMsr bool = false
	this.HideFromMsr = &hideFromMsr
	var ensureDisplayDevice bool = true
	this.EnsureDisplayDevice = &ensureDisplayDevice
	var time string = "LOCAL"
	this.Time = &time
	var shutdownTimeout int32 = 90
	this.ShutdownTimeout = &shutdownTimeout
	return &this
}

// NewCreateVMParamsWithDefaults instantiates a new CreateVMParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMParamsWithDefaults() *CreateVMParams {
	this := CreateVMParams{}
	var commandLineArgs string = ""
	this.CommandLineArgs = &commandLineArgs
	var cpuMode string = "CUSTOM"
	this.CpuMode = &cpuMode
	var vcpus int32 = 1
	this.Vcpus = &vcpus
	var cores int32 = 1
	this.Cores = &cores
	var threads int32 = 1
	this.Threads = &threads
	var pinVcpus bool = false
	this.PinVcpus = &pinVcpus
	var suspendOnSnapshot bool = false
	this.SuspendOnSnapshot = &suspendOnSnapshot
	var hypervEnlightenments bool = false
	this.HypervEnlightenments = &hypervEnlightenments
	var bootloader string = "UEFI"
	this.Bootloader = &bootloader
	var autostart bool = true
	this.Autostart = &autostart
	var hideFromMsr bool = false
	this.HideFromMsr = &hideFromMsr
	var ensureDisplayDevice bool = true
	this.EnsureDisplayDevice = &ensureDisplayDevice
	var time string = "LOCAL"
	this.Time = &time
	var shutdownTimeout int32 = 90
	this.ShutdownTimeout = &shutdownTimeout
	return &this
}

// GetCommandLineArgs returns the CommandLineArgs field value if set, zero value otherwise.
func (o *CreateVMParams) GetCommandLineArgs() string {
	if o == nil || IsNil(o.CommandLineArgs) {
		var ret string
		return ret
	}
	return *o.CommandLineArgs
}

// GetCommandLineArgsOk returns a tuple with the CommandLineArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetCommandLineArgsOk() (*string, bool) {
	if o == nil || IsNil(o.CommandLineArgs) {
		return nil, false
	}
	return o.CommandLineArgs, true
}

// HasCommandLineArgs returns a boolean if a field has been set.
func (o *CreateVMParams) HasCommandLineArgs() bool {
	if o != nil && !IsNil(o.CommandLineArgs) {
		return true
	}

	return false
}

// SetCommandLineArgs gets a reference to the given string and assigns it to the CommandLineArgs field.
func (o *CreateVMParams) SetCommandLineArgs(v string) {
	o.CommandLineArgs = &v
}

// GetCpuMode returns the CpuMode field value if set, zero value otherwise.
func (o *CreateVMParams) GetCpuMode() string {
	if o == nil || IsNil(o.CpuMode) {
		var ret string
		return ret
	}
	return *o.CpuMode
}

// GetCpuModeOk returns a tuple with the CpuMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetCpuModeOk() (*string, bool) {
	if o == nil || IsNil(o.CpuMode) {
		return nil, false
	}
	return o.CpuMode, true
}

// HasCpuMode returns a boolean if a field has been set.
func (o *CreateVMParams) HasCpuMode() bool {
	if o != nil && !IsNil(o.CpuMode) {
		return true
	}

	return false
}

// SetCpuMode gets a reference to the given string and assigns it to the CpuMode field.
func (o *CreateVMParams) SetCpuMode(v string) {
	o.CpuMode = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetCpuModel() string {
	if o == nil || IsNil(o.CpuModel.Get()) {
		var ret string
		return ret
	}
	return *o.CpuModel.Get()
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetCpuModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuModel.Get(), o.CpuModel.IsSet()
}

// HasCpuModel returns a boolean if a field has been set.
func (o *CreateVMParams) HasCpuModel() bool {
	if o != nil && o.CpuModel.IsSet() {
		return true
	}

	return false
}

// SetCpuModel gets a reference to the given NullableString and assigns it to the CpuModel field.
func (o *CreateVMParams) SetCpuModel(v string) {
	o.CpuModel.Set(&v)
}

// SetCpuModelNil sets the value for CpuModel to be an explicit nil
func (o *CreateVMParams) SetCpuModelNil() {
	o.CpuModel.Set(nil)
}

// UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
func (o *CreateVMParams) UnsetCpuModel() {
	o.CpuModel.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateVMParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateVMParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateVMParams) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVMParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVMParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVMParams) SetDescription(v string) {
	o.Description = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise.
func (o *CreateVMParams) GetVcpus() int32 {
	if o == nil || IsNil(o.Vcpus) {
		var ret int32
		return ret
	}
	return *o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetVcpusOk() (*int32, bool) {
	if o == nil || IsNil(o.Vcpus) {
		return nil, false
	}
	return o.Vcpus, true
}

// HasVcpus returns a boolean if a field has been set.
func (o *CreateVMParams) HasVcpus() bool {
	if o != nil && !IsNil(o.Vcpus) {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given int32 and assigns it to the Vcpus field.
func (o *CreateVMParams) SetVcpus(v int32) {
	o.Vcpus = &v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *CreateVMParams) GetCores() int32 {
	if o == nil || IsNil(o.Cores) {
		var ret int32
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetCoresOk() (*int32, bool) {
	if o == nil || IsNil(o.Cores) {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *CreateVMParams) HasCores() bool {
	if o != nil && !IsNil(o.Cores) {
		return true
	}

	return false
}

// SetCores gets a reference to the given int32 and assigns it to the Cores field.
func (o *CreateVMParams) SetCores(v int32) {
	o.Cores = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *CreateVMParams) GetThreads() int32 {
	if o == nil || IsNil(o.Threads) {
		var ret int32
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetThreadsOk() (*int32, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *CreateVMParams) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given int32 and assigns it to the Threads field.
func (o *CreateVMParams) SetThreads(v int32) {
	o.Threads = &v
}

// GetCpuset returns the Cpuset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetCpuset() string {
	if o == nil || IsNil(o.Cpuset.Get()) {
		var ret string
		return ret
	}
	return *o.Cpuset.Get()
}

// GetCpusetOk returns a tuple with the Cpuset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetCpusetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpuset.Get(), o.Cpuset.IsSet()
}

// HasCpuset returns a boolean if a field has been set.
func (o *CreateVMParams) HasCpuset() bool {
	if o != nil && o.Cpuset.IsSet() {
		return true
	}

	return false
}

// SetCpuset gets a reference to the given NullableString and assigns it to the Cpuset field.
func (o *CreateVMParams) SetCpuset(v string) {
	o.Cpuset.Set(&v)
}

// SetCpusetNil sets the value for Cpuset to be an explicit nil
func (o *CreateVMParams) SetCpusetNil() {
	o.Cpuset.Set(nil)
}

// UnsetCpuset ensures that no value is present for Cpuset, not even an explicit nil
func (o *CreateVMParams) UnsetCpuset() {
	o.Cpuset.Unset()
}

// GetNodeset returns the Nodeset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetNodeset() string {
	if o == nil || IsNil(o.Nodeset.Get()) {
		var ret string
		return ret
	}
	return *o.Nodeset.Get()
}

// GetNodesetOk returns a tuple with the Nodeset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetNodesetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodeset.Get(), o.Nodeset.IsSet()
}

// HasNodeset returns a boolean if a field has been set.
func (o *CreateVMParams) HasNodeset() bool {
	if o != nil && o.Nodeset.IsSet() {
		return true
	}

	return false
}

// SetNodeset gets a reference to the given NullableString and assigns it to the Nodeset field.
func (o *CreateVMParams) SetNodeset(v string) {
	o.Nodeset.Set(&v)
}

// SetNodesetNil sets the value for Nodeset to be an explicit nil
func (o *CreateVMParams) SetNodesetNil() {
	o.Nodeset.Set(nil)
}

// UnsetNodeset ensures that no value is present for Nodeset, not even an explicit nil
func (o *CreateVMParams) UnsetNodeset() {
	o.Nodeset.Unset()
}

// GetPinVcpus returns the PinVcpus field value if set, zero value otherwise.
func (o *CreateVMParams) GetPinVcpus() bool {
	if o == nil || IsNil(o.PinVcpus) {
		var ret bool
		return ret
	}
	return *o.PinVcpus
}

// GetPinVcpusOk returns a tuple with the PinVcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetPinVcpusOk() (*bool, bool) {
	if o == nil || IsNil(o.PinVcpus) {
		return nil, false
	}
	return o.PinVcpus, true
}

// HasPinVcpus returns a boolean if a field has been set.
func (o *CreateVMParams) HasPinVcpus() bool {
	if o != nil && !IsNil(o.PinVcpus) {
		return true
	}

	return false
}

// SetPinVcpus gets a reference to the given bool and assigns it to the PinVcpus field.
func (o *CreateVMParams) SetPinVcpus(v bool) {
	o.PinVcpus = &v
}

// GetSuspendOnSnapshot returns the SuspendOnSnapshot field value if set, zero value otherwise.
func (o *CreateVMParams) GetSuspendOnSnapshot() bool {
	if o == nil || IsNil(o.SuspendOnSnapshot) {
		var ret bool
		return ret
	}
	return *o.SuspendOnSnapshot
}

// GetSuspendOnSnapshotOk returns a tuple with the SuspendOnSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetSuspendOnSnapshotOk() (*bool, bool) {
	if o == nil || IsNil(o.SuspendOnSnapshot) {
		return nil, false
	}
	return o.SuspendOnSnapshot, true
}

// HasSuspendOnSnapshot returns a boolean if a field has been set.
func (o *CreateVMParams) HasSuspendOnSnapshot() bool {
	if o != nil && !IsNil(o.SuspendOnSnapshot) {
		return true
	}

	return false
}

// SetSuspendOnSnapshot gets a reference to the given bool and assigns it to the SuspendOnSnapshot field.
func (o *CreateVMParams) SetSuspendOnSnapshot(v bool) {
	o.SuspendOnSnapshot = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CreateVMParams) GetMemory() int64 {
	if o == nil || IsNil(o.Memory) {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetMemoryOk() (*int64, bool) {
	if o == nil || IsNil(o.Memory) {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CreateVMParams) HasMemory() bool {
	if o != nil && !IsNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *CreateVMParams) SetMemory(v int64) {
	o.Memory = &v
}

// GetMinMemory returns the MinMemory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetMinMemory() int64 {
	if o == nil || IsNil(o.MinMemory.Get()) {
		var ret int64
		return ret
	}
	return *o.MinMemory.Get()
}

// GetMinMemoryOk returns a tuple with the MinMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetMinMemoryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinMemory.Get(), o.MinMemory.IsSet()
}

// HasMinMemory returns a boolean if a field has been set.
func (o *CreateVMParams) HasMinMemory() bool {
	if o != nil && o.MinMemory.IsSet() {
		return true
	}

	return false
}

// SetMinMemory gets a reference to the given NullableInt64 and assigns it to the MinMemory field.
func (o *CreateVMParams) SetMinMemory(v int64) {
	o.MinMemory.Set(&v)
}

// SetMinMemoryNil sets the value for MinMemory to be an explicit nil
func (o *CreateVMParams) SetMinMemoryNil() {
	o.MinMemory.Set(nil)
}

// UnsetMinMemory ensures that no value is present for MinMemory, not even an explicit nil
func (o *CreateVMParams) UnsetMinMemory() {
	o.MinMemory.Unset()
}

// GetHypervEnlightenments returns the HypervEnlightenments field value if set, zero value otherwise.
func (o *CreateVMParams) GetHypervEnlightenments() bool {
	if o == nil || IsNil(o.HypervEnlightenments) {
		var ret bool
		return ret
	}
	return *o.HypervEnlightenments
}

// GetHypervEnlightenmentsOk returns a tuple with the HypervEnlightenments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetHypervEnlightenmentsOk() (*bool, bool) {
	if o == nil || IsNil(o.HypervEnlightenments) {
		return nil, false
	}
	return o.HypervEnlightenments, true
}

// HasHypervEnlightenments returns a boolean if a field has been set.
func (o *CreateVMParams) HasHypervEnlightenments() bool {
	if o != nil && !IsNil(o.HypervEnlightenments) {
		return true
	}

	return false
}

// SetHypervEnlightenments gets a reference to the given bool and assigns it to the HypervEnlightenments field.
func (o *CreateVMParams) SetHypervEnlightenments(v bool) {
	o.HypervEnlightenments = &v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise.
func (o *CreateVMParams) GetBootloader() string {
	if o == nil || IsNil(o.Bootloader) {
		var ret string
		return ret
	}
	return *o.Bootloader
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetBootloaderOk() (*string, bool) {
	if o == nil || IsNil(o.Bootloader) {
		return nil, false
	}
	return o.Bootloader, true
}

// HasBootloader returns a boolean if a field has been set.
func (o *CreateVMParams) HasBootloader() bool {
	if o != nil && !IsNil(o.Bootloader) {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given string and assigns it to the Bootloader field.
func (o *CreateVMParams) SetBootloader(v string) {
	o.Bootloader = &v
}

// GetAutostart returns the Autostart field value if set, zero value otherwise.
func (o *CreateVMParams) GetAutostart() bool {
	if o == nil || IsNil(o.Autostart) {
		var ret bool
		return ret
	}
	return *o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetAutostartOk() (*bool, bool) {
	if o == nil || IsNil(o.Autostart) {
		return nil, false
	}
	return o.Autostart, true
}

// HasAutostart returns a boolean if a field has been set.
func (o *CreateVMParams) HasAutostart() bool {
	if o != nil && !IsNil(o.Autostart) {
		return true
	}

	return false
}

// SetAutostart gets a reference to the given bool and assigns it to the Autostart field.
func (o *CreateVMParams) SetAutostart(v bool) {
	o.Autostart = &v
}

// GetHideFromMsr returns the HideFromMsr field value if set, zero value otherwise.
func (o *CreateVMParams) GetHideFromMsr() bool {
	if o == nil || IsNil(o.HideFromMsr) {
		var ret bool
		return ret
	}
	return *o.HideFromMsr
}

// GetHideFromMsrOk returns a tuple with the HideFromMsr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetHideFromMsrOk() (*bool, bool) {
	if o == nil || IsNil(o.HideFromMsr) {
		return nil, false
	}
	return o.HideFromMsr, true
}

// HasHideFromMsr returns a boolean if a field has been set.
func (o *CreateVMParams) HasHideFromMsr() bool {
	if o != nil && !IsNil(o.HideFromMsr) {
		return true
	}

	return false
}

// SetHideFromMsr gets a reference to the given bool and assigns it to the HideFromMsr field.
func (o *CreateVMParams) SetHideFromMsr(v bool) {
	o.HideFromMsr = &v
}

// GetEnsureDisplayDevice returns the EnsureDisplayDevice field value if set, zero value otherwise.
func (o *CreateVMParams) GetEnsureDisplayDevice() bool {
	if o == nil || IsNil(o.EnsureDisplayDevice) {
		var ret bool
		return ret
	}
	return *o.EnsureDisplayDevice
}

// GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetEnsureDisplayDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.EnsureDisplayDevice) {
		return nil, false
	}
	return o.EnsureDisplayDevice, true
}

// HasEnsureDisplayDevice returns a boolean if a field has been set.
func (o *CreateVMParams) HasEnsureDisplayDevice() bool {
	if o != nil && !IsNil(o.EnsureDisplayDevice) {
		return true
	}

	return false
}

// SetEnsureDisplayDevice gets a reference to the given bool and assigns it to the EnsureDisplayDevice field.
func (o *CreateVMParams) SetEnsureDisplayDevice(v bool) {
	o.EnsureDisplayDevice = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *CreateVMParams) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *CreateVMParams) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *CreateVMParams) SetTime(v string) {
	o.Time = &v
}

// GetShutdownTimeout returns the ShutdownTimeout field value if set, zero value otherwise.
func (o *CreateVMParams) GetShutdownTimeout() int32 {
	if o == nil || IsNil(o.ShutdownTimeout) {
		var ret int32
		return ret
	}
	return *o.ShutdownTimeout
}

// GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMParams) GetShutdownTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.ShutdownTimeout) {
		return nil, false
	}
	return o.ShutdownTimeout, true
}

// HasShutdownTimeout returns a boolean if a field has been set.
func (o *CreateVMParams) HasShutdownTimeout() bool {
	if o != nil && !IsNil(o.ShutdownTimeout) {
		return true
	}

	return false
}

// SetShutdownTimeout gets a reference to the given int32 and assigns it to the ShutdownTimeout field.
func (o *CreateVMParams) SetShutdownTimeout(v int32) {
	o.ShutdownTimeout = &v
}

// GetArchType returns the ArchType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetArchType() string {
	if o == nil || IsNil(o.ArchType.Get()) {
		var ret string
		return ret
	}
	return *o.ArchType.Get()
}

// GetArchTypeOk returns a tuple with the ArchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetArchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchType.Get(), o.ArchType.IsSet()
}

// HasArchType returns a boolean if a field has been set.
func (o *CreateVMParams) HasArchType() bool {
	if o != nil && o.ArchType.IsSet() {
		return true
	}

	return false
}

// SetArchType gets a reference to the given NullableString and assigns it to the ArchType field.
func (o *CreateVMParams) SetArchType(v string) {
	o.ArchType.Set(&v)
}

// SetArchTypeNil sets the value for ArchType to be an explicit nil
func (o *CreateVMParams) SetArchTypeNil() {
	o.ArchType.Set(nil)
}

// UnsetArchType ensures that no value is present for ArchType, not even an explicit nil
func (o *CreateVMParams) UnsetArchType() {
	o.ArchType.Unset()
}

// GetMachineType returns the MachineType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetMachineType() string {
	if o == nil || IsNil(o.MachineType.Get()) {
		var ret string
		return ret
	}
	return *o.MachineType.Get()
}

// GetMachineTypeOk returns a tuple with the MachineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetMachineTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineType.Get(), o.MachineType.IsSet()
}

// HasMachineType returns a boolean if a field has been set.
func (o *CreateVMParams) HasMachineType() bool {
	if o != nil && o.MachineType.IsSet() {
		return true
	}

	return false
}

// SetMachineType gets a reference to the given NullableString and assigns it to the MachineType field.
func (o *CreateVMParams) SetMachineType(v string) {
	o.MachineType.Set(&v)
}

// SetMachineTypeNil sets the value for MachineType to be an explicit nil
func (o *CreateVMParams) SetMachineTypeNil() {
	o.MachineType.Set(nil)
}

// UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
func (o *CreateVMParams) UnsetMachineType() {
	o.MachineType.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVMParams) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVMParams) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *CreateVMParams) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *CreateVMParams) SetUuid(v string) {
	o.Uuid.Set(&v)
}

// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *CreateVMParams) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *CreateVMParams) UnsetUuid() {
	o.Uuid.Unset()
}

func (o CreateVMParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommandLineArgs) {
		toSerialize["command_line_args"] = o.CommandLineArgs
	}
	if !IsNil(o.CpuMode) {
		toSerialize["cpu_mode"] = o.CpuMode
	}
	if o.CpuModel.IsSet() {
		toSerialize["cpu_model"] = o.CpuModel.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Vcpus) {
		toSerialize["vcpus"] = o.Vcpus
	}
	if !IsNil(o.Cores) {
		toSerialize["cores"] = o.Cores
	}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}
	if o.Cpuset.IsSet() {
		toSerialize["cpuset"] = o.Cpuset.Get()
	}
	if o.Nodeset.IsSet() {
		toSerialize["nodeset"] = o.Nodeset.Get()
	}
	if !IsNil(o.PinVcpus) {
		toSerialize["pin_vcpus"] = o.PinVcpus
	}
	if !IsNil(o.SuspendOnSnapshot) {
		toSerialize["suspend_on_snapshot"] = o.SuspendOnSnapshot
	}
	if !IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if o.MinMemory.IsSet() {
		toSerialize["min_memory"] = o.MinMemory.Get()
	}
	if !IsNil(o.HypervEnlightenments) {
		toSerialize["hyperv_enlightenments"] = o.HypervEnlightenments
	}
	if !IsNil(o.Bootloader) {
		toSerialize["bootloader"] = o.Bootloader
	}
	if !IsNil(o.Autostart) {
		toSerialize["autostart"] = o.Autostart
	}
	if !IsNil(o.HideFromMsr) {
		toSerialize["hide_from_msr"] = o.HideFromMsr
	}
	if !IsNil(o.EnsureDisplayDevice) {
		toSerialize["ensure_display_device"] = o.EnsureDisplayDevice
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.ShutdownTimeout) {
		toSerialize["shutdown_timeout"] = o.ShutdownTimeout
	}
	if o.ArchType.IsSet() {
		toSerialize["arch_type"] = o.ArchType.Get()
	}
	if o.MachineType.IsSet() {
		toSerialize["machine_type"] = o.MachineType.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	return toSerialize, nil
}

type NullableCreateVMParams struct {
	value *CreateVMParams
	isSet bool
}

func (v NullableCreateVMParams) Get() *CreateVMParams {
	return v.value
}

func (v *NullableCreateVMParams) Set(val *CreateVMParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMParams(val *CreateVMParams) *NullableCreateVMParams {
	return &NullableCreateVMParams{value: val, isSet: true}
}

func (v NullableCreateVMParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
