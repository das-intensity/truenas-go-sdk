/*
 * TrueNAS RESTful API
 *
 * Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: v2.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package truenas

import (
	"encoding/json"
)

// ShareSMB struct for ShareSMB
type ShareSMB struct {
	Id                   int32     `json:"id"`
	Path                 string    `json:"path"`
	PathSuffix           *string   `json:"path_suffix,omitempty"`
	Purpose              *string   `json:"purpose,omitempty"`
	Home                 *bool     `json:"home,omitempty"`
	Timemachine          *bool     `json:"timemachine,omitempty"`
	Name                 *string   `json:"name,omitempty"`
	Comment              *string   `json:"comment,omitempty"`
	Ro                   *bool     `json:"ro,omitempty"`
	Browsable            *bool     `json:"browsable,omitempty"`
	Recyclebin           *bool     `json:"recyclebin,omitempty"`
	Shadowcopy           *bool     `json:"shadowcopy,omitempty"`
	Guestok              *bool     `json:"guestok,omitempty"`
	Abe                  *bool     `json:"abe,omitempty"`
	Hostsallow           *[]string `json:"hostsallow,omitempty"`
	Hostsdeny            *[]string `json:"hostsdeny,omitempty"`
	AaplNameMangling     *bool     `json:"aapl_name_mangling,omitempty"`
	Acl                  *bool     `json:"acl,omitempty"`
	Durablehandle        *bool     `json:"durablehandle,omitempty"`
	Streams              *bool     `json:"streams,omitempty"`
	Fsrvp                *bool     `json:"fsrvp,omitempty"`
	Auxsmbconf           *string   `json:"auxsmbconf,omitempty"`
	Enabled              *bool     `json:"enabled,omitempty"`
	Locked               *bool     `json:"locked,omitempty"`
	Vuid                 *string   `json:"vuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ShareSMB ShareSMB

// NewShareSMB instantiates a new ShareSMB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareSMB(id int32, path string) *ShareSMB {
	this := ShareSMB{}
	this.Id = id
	this.Path = path
	return &this
}

// NewShareSMBWithDefaults instantiates a new ShareSMB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareSMBWithDefaults() *ShareSMB {
	this := ShareSMB{}
	return &this
}

// GetId returns the Id field value
func (o *ShareSMB) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShareSMB) SetId(v int32) {
	o.Id = v
}

// GetPath returns the Path field value
func (o *ShareSMB) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ShareSMB) SetPath(v string) {
	o.Path = v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *ShareSMB) GetPathSuffix() string {
	if o == nil || o.PathSuffix == nil {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetPathSuffixOk() (*string, bool) {
	if o == nil || o.PathSuffix == nil {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *ShareSMB) HasPathSuffix() bool {
	if o != nil && o.PathSuffix != nil {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *ShareSMB) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *ShareSMB) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *ShareSMB) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *ShareSMB) SetPurpose(v string) {
	o.Purpose = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *ShareSMB) GetHome() bool {
	if o == nil || o.Home == nil {
		var ret bool
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetHomeOk() (*bool, bool) {
	if o == nil || o.Home == nil {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *ShareSMB) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given bool and assigns it to the Home field.
func (o *ShareSMB) SetHome(v bool) {
	o.Home = &v
}

// GetTimemachine returns the Timemachine field value if set, zero value otherwise.
func (o *ShareSMB) GetTimemachine() bool {
	if o == nil || o.Timemachine == nil {
		var ret bool
		return ret
	}
	return *o.Timemachine
}

// GetTimemachineOk returns a tuple with the Timemachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetTimemachineOk() (*bool, bool) {
	if o == nil || o.Timemachine == nil {
		return nil, false
	}
	return o.Timemachine, true
}

// HasTimemachine returns a boolean if a field has been set.
func (o *ShareSMB) HasTimemachine() bool {
	if o != nil && o.Timemachine != nil {
		return true
	}

	return false
}

// SetTimemachine gets a reference to the given bool and assigns it to the Timemachine field.
func (o *ShareSMB) SetTimemachine(v bool) {
	o.Timemachine = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ShareSMB) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ShareSMB) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ShareSMB) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ShareSMB) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ShareSMB) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ShareSMB) SetComment(v string) {
	o.Comment = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *ShareSMB) GetRo() bool {
	if o == nil || o.Ro == nil {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetRoOk() (*bool, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *ShareSMB) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *ShareSMB) SetRo(v bool) {
	o.Ro = &v
}

// GetBrowsable returns the Browsable field value if set, zero value otherwise.
func (o *ShareSMB) GetBrowsable() bool {
	if o == nil || o.Browsable == nil {
		var ret bool
		return ret
	}
	return *o.Browsable
}

// GetBrowsableOk returns a tuple with the Browsable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetBrowsableOk() (*bool, bool) {
	if o == nil || o.Browsable == nil {
		return nil, false
	}
	return o.Browsable, true
}

// HasBrowsable returns a boolean if a field has been set.
func (o *ShareSMB) HasBrowsable() bool {
	if o != nil && o.Browsable != nil {
		return true
	}

	return false
}

// SetBrowsable gets a reference to the given bool and assigns it to the Browsable field.
func (o *ShareSMB) SetBrowsable(v bool) {
	o.Browsable = &v
}

// GetRecyclebin returns the Recyclebin field value if set, zero value otherwise.
func (o *ShareSMB) GetRecyclebin() bool {
	if o == nil || o.Recyclebin == nil {
		var ret bool
		return ret
	}
	return *o.Recyclebin
}

// GetRecyclebinOk returns a tuple with the Recyclebin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetRecyclebinOk() (*bool, bool) {
	if o == nil || o.Recyclebin == nil {
		return nil, false
	}
	return o.Recyclebin, true
}

// HasRecyclebin returns a boolean if a field has been set.
func (o *ShareSMB) HasRecyclebin() bool {
	if o != nil && o.Recyclebin != nil {
		return true
	}

	return false
}

// SetRecyclebin gets a reference to the given bool and assigns it to the Recyclebin field.
func (o *ShareSMB) SetRecyclebin(v bool) {
	o.Recyclebin = &v
}

// GetShadowcopy returns the Shadowcopy field value if set, zero value otherwise.
func (o *ShareSMB) GetShadowcopy() bool {
	if o == nil || o.Shadowcopy == nil {
		var ret bool
		return ret
	}
	return *o.Shadowcopy
}

// GetShadowcopyOk returns a tuple with the Shadowcopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetShadowcopyOk() (*bool, bool) {
	if o == nil || o.Shadowcopy == nil {
		return nil, false
	}
	return o.Shadowcopy, true
}

// HasShadowcopy returns a boolean if a field has been set.
func (o *ShareSMB) HasShadowcopy() bool {
	if o != nil && o.Shadowcopy != nil {
		return true
	}

	return false
}

// SetShadowcopy gets a reference to the given bool and assigns it to the Shadowcopy field.
func (o *ShareSMB) SetShadowcopy(v bool) {
	o.Shadowcopy = &v
}

// GetGuestok returns the Guestok field value if set, zero value otherwise.
func (o *ShareSMB) GetGuestok() bool {
	if o == nil || o.Guestok == nil {
		var ret bool
		return ret
	}
	return *o.Guestok
}

// GetGuestokOk returns a tuple with the Guestok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetGuestokOk() (*bool, bool) {
	if o == nil || o.Guestok == nil {
		return nil, false
	}
	return o.Guestok, true
}

// HasGuestok returns a boolean if a field has been set.
func (o *ShareSMB) HasGuestok() bool {
	if o != nil && o.Guestok != nil {
		return true
	}

	return false
}

// SetGuestok gets a reference to the given bool and assigns it to the Guestok field.
func (o *ShareSMB) SetGuestok(v bool) {
	o.Guestok = &v
}

// GetAbe returns the Abe field value if set, zero value otherwise.
func (o *ShareSMB) GetAbe() bool {
	if o == nil || o.Abe == nil {
		var ret bool
		return ret
	}
	return *o.Abe
}

// GetAbeOk returns a tuple with the Abe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetAbeOk() (*bool, bool) {
	if o == nil || o.Abe == nil {
		return nil, false
	}
	return o.Abe, true
}

// HasAbe returns a boolean if a field has been set.
func (o *ShareSMB) HasAbe() bool {
	if o != nil && o.Abe != nil {
		return true
	}

	return false
}

// SetAbe gets a reference to the given bool and assigns it to the Abe field.
func (o *ShareSMB) SetAbe(v bool) {
	o.Abe = &v
}

// GetHostsallow returns the Hostsallow field value if set, zero value otherwise.
func (o *ShareSMB) GetHostsallow() []string {
	if o == nil || o.Hostsallow == nil {
		var ret []string
		return ret
	}
	return *o.Hostsallow
}

// GetHostsallowOk returns a tuple with the Hostsallow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetHostsallowOk() (*[]string, bool) {
	if o == nil || o.Hostsallow == nil {
		return nil, false
	}
	return o.Hostsallow, true
}

// HasHostsallow returns a boolean if a field has been set.
func (o *ShareSMB) HasHostsallow() bool {
	if o != nil && o.Hostsallow != nil {
		return true
	}

	return false
}

// SetHostsallow gets a reference to the given []string and assigns it to the Hostsallow field.
func (o *ShareSMB) SetHostsallow(v []string) {
	o.Hostsallow = &v
}

// GetHostsdeny returns the Hostsdeny field value if set, zero value otherwise.
func (o *ShareSMB) GetHostsdeny() []string {
	if o == nil || o.Hostsdeny == nil {
		var ret []string
		return ret
	}
	return *o.Hostsdeny
}

// GetHostsdenyOk returns a tuple with the Hostsdeny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetHostsdenyOk() (*[]string, bool) {
	if o == nil || o.Hostsdeny == nil {
		return nil, false
	}
	return o.Hostsdeny, true
}

// HasHostsdeny returns a boolean if a field has been set.
func (o *ShareSMB) HasHostsdeny() bool {
	if o != nil && o.Hostsdeny != nil {
		return true
	}

	return false
}

// SetHostsdeny gets a reference to the given []string and assigns it to the Hostsdeny field.
func (o *ShareSMB) SetHostsdeny(v []string) {
	o.Hostsdeny = &v
}

// GetAaplNameMangling returns the AaplNameMangling field value if set, zero value otherwise.
func (o *ShareSMB) GetAaplNameMangling() bool {
	if o == nil || o.AaplNameMangling == nil {
		var ret bool
		return ret
	}
	return *o.AaplNameMangling
}

// GetAaplNameManglingOk returns a tuple with the AaplNameMangling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetAaplNameManglingOk() (*bool, bool) {
	if o == nil || o.AaplNameMangling == nil {
		return nil, false
	}
	return o.AaplNameMangling, true
}

// HasAaplNameMangling returns a boolean if a field has been set.
func (o *ShareSMB) HasAaplNameMangling() bool {
	if o != nil && o.AaplNameMangling != nil {
		return true
	}

	return false
}

// SetAaplNameMangling gets a reference to the given bool and assigns it to the AaplNameMangling field.
func (o *ShareSMB) SetAaplNameMangling(v bool) {
	o.AaplNameMangling = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *ShareSMB) GetAcl() bool {
	if o == nil || o.Acl == nil {
		var ret bool
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetAclOk() (*bool, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *ShareSMB) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given bool and assigns it to the Acl field.
func (o *ShareSMB) SetAcl(v bool) {
	o.Acl = &v
}

// GetDurablehandle returns the Durablehandle field value if set, zero value otherwise.
func (o *ShareSMB) GetDurablehandle() bool {
	if o == nil || o.Durablehandle == nil {
		var ret bool
		return ret
	}
	return *o.Durablehandle
}

// GetDurablehandleOk returns a tuple with the Durablehandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetDurablehandleOk() (*bool, bool) {
	if o == nil || o.Durablehandle == nil {
		return nil, false
	}
	return o.Durablehandle, true
}

// HasDurablehandle returns a boolean if a field has been set.
func (o *ShareSMB) HasDurablehandle() bool {
	if o != nil && o.Durablehandle != nil {
		return true
	}

	return false
}

// SetDurablehandle gets a reference to the given bool and assigns it to the Durablehandle field.
func (o *ShareSMB) SetDurablehandle(v bool) {
	o.Durablehandle = &v
}

// GetStreams returns the Streams field value if set, zero value otherwise.
func (o *ShareSMB) GetStreams() bool {
	if o == nil || o.Streams == nil {
		var ret bool
		return ret
	}
	return *o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetStreamsOk() (*bool, bool) {
	if o == nil || o.Streams == nil {
		return nil, false
	}
	return o.Streams, true
}

// HasStreams returns a boolean if a field has been set.
func (o *ShareSMB) HasStreams() bool {
	if o != nil && o.Streams != nil {
		return true
	}

	return false
}

// SetStreams gets a reference to the given bool and assigns it to the Streams field.
func (o *ShareSMB) SetStreams(v bool) {
	o.Streams = &v
}

// GetFsrvp returns the Fsrvp field value if set, zero value otherwise.
func (o *ShareSMB) GetFsrvp() bool {
	if o == nil || o.Fsrvp == nil {
		var ret bool
		return ret
	}
	return *o.Fsrvp
}

// GetFsrvpOk returns a tuple with the Fsrvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetFsrvpOk() (*bool, bool) {
	if o == nil || o.Fsrvp == nil {
		return nil, false
	}
	return o.Fsrvp, true
}

// HasFsrvp returns a boolean if a field has been set.
func (o *ShareSMB) HasFsrvp() bool {
	if o != nil && o.Fsrvp != nil {
		return true
	}

	return false
}

// SetFsrvp gets a reference to the given bool and assigns it to the Fsrvp field.
func (o *ShareSMB) SetFsrvp(v bool) {
	o.Fsrvp = &v
}

// GetAuxsmbconf returns the Auxsmbconf field value if set, zero value otherwise.
func (o *ShareSMB) GetAuxsmbconf() string {
	if o == nil || o.Auxsmbconf == nil {
		var ret string
		return ret
	}
	return *o.Auxsmbconf
}

// GetAuxsmbconfOk returns a tuple with the Auxsmbconf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetAuxsmbconfOk() (*string, bool) {
	if o == nil || o.Auxsmbconf == nil {
		return nil, false
	}
	return o.Auxsmbconf, true
}

// HasAuxsmbconf returns a boolean if a field has been set.
func (o *ShareSMB) HasAuxsmbconf() bool {
	if o != nil && o.Auxsmbconf != nil {
		return true
	}

	return false
}

// SetAuxsmbconf gets a reference to the given string and assigns it to the Auxsmbconf field.
func (o *ShareSMB) SetAuxsmbconf(v string) {
	o.Auxsmbconf = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ShareSMB) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ShareSMB) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ShareSMB) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ShareSMB) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ShareSMB) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *ShareSMB) SetLocked(v bool) {
	o.Locked = &v
}

// GetVuid returns the Vuid field value if set, zero value otherwise.
func (o *ShareSMB) GetVuid() string {
	if o == nil || o.Vuid == nil {
		var ret string
		return ret
	}
	return *o.Vuid
}

// GetVuidOk returns a tuple with the Vuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShareSMB) GetVuidOk() (*string, bool) {
	if o == nil || o.Vuid == nil {
		return nil, false
	}
	return o.Vuid, true
}

// HasVuid returns a boolean if a field has been set.
func (o *ShareSMB) HasVuid() bool {
	if o != nil && o.Vuid != nil {
		return true
	}

	return false
}

// SetVuid gets a reference to the given string and assigns it to the Vuid field.
func (o *ShareSMB) SetVuid(v string) {
	o.Vuid = &v
}

func (o ShareSMB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.PathSuffix != nil {
		toSerialize["path_suffix"] = o.PathSuffix
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.Home != nil {
		toSerialize["home"] = o.Home
	}
	if o.Timemachine != nil {
		toSerialize["timemachine"] = o.Timemachine
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Browsable != nil {
		toSerialize["browsable"] = o.Browsable
	}
	if o.Recyclebin != nil {
		toSerialize["recyclebin"] = o.Recyclebin
	}
	if o.Shadowcopy != nil {
		toSerialize["shadowcopy"] = o.Shadowcopy
	}
	if o.Guestok != nil {
		toSerialize["guestok"] = o.Guestok
	}
	if o.Abe != nil {
		toSerialize["abe"] = o.Abe
	}
	if o.Hostsallow != nil {
		toSerialize["hostsallow"] = o.Hostsallow
	}
	if o.Hostsdeny != nil {
		toSerialize["hostsdeny"] = o.Hostsdeny
	}
	if o.AaplNameMangling != nil {
		toSerialize["aapl_name_mangling"] = o.AaplNameMangling
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Durablehandle != nil {
		toSerialize["durablehandle"] = o.Durablehandle
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	if o.Fsrvp != nil {
		toSerialize["fsrvp"] = o.Fsrvp
	}
	if o.Auxsmbconf != nil {
		toSerialize["auxsmbconf"] = o.Auxsmbconf
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.Vuid != nil {
		toSerialize["vuid"] = o.Vuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ShareSMB) UnmarshalJSON(bytes []byte) (err error) {
	varShareSMB := _ShareSMB{}

	if err = json.Unmarshal(bytes, &varShareSMB); err == nil {
		*o = ShareSMB(varShareSMB)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "path")
		delete(additionalProperties, "path_suffix")
		delete(additionalProperties, "purpose")
		delete(additionalProperties, "home")
		delete(additionalProperties, "timemachine")
		delete(additionalProperties, "name")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "ro")
		delete(additionalProperties, "browsable")
		delete(additionalProperties, "recyclebin")
		delete(additionalProperties, "shadowcopy")
		delete(additionalProperties, "guestok")
		delete(additionalProperties, "abe")
		delete(additionalProperties, "hostsallow")
		delete(additionalProperties, "hostsdeny")
		delete(additionalProperties, "aapl_name_mangling")
		delete(additionalProperties, "acl")
		delete(additionalProperties, "durablehandle")
		delete(additionalProperties, "streams")
		delete(additionalProperties, "fsrvp")
		delete(additionalProperties, "auxsmbconf")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "vuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShareSMB struct {
	value *ShareSMB
	isSet bool
}

func (v NullableShareSMB) Get() *ShareSMB {
	return v.value
}

func (v *NullableShareSMB) Set(val *ShareSMB) {
	v.value = val
	v.isSet = true
}

func (v NullableShareSMB) IsSet() bool {
	return v.isSet
}

func (v *NullableShareSMB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareSMB(val *ShareSMB) *NullableShareSMB {
	return &NullableShareSMB{value: val, isSet: true}
}

func (v NullableShareSMB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareSMB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
