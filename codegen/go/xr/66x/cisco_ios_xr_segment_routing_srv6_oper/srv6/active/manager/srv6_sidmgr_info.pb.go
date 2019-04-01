/*
Copyright 2019 Cisco Systems

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv6_sidmgr_info.proto

package cisco_ios_xr_segment_routing_srv6_oper_srv6_active_manager

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Srv6SidmgrInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Srv6SidmgrInfo_KEYS) Reset()         { *m = Srv6SidmgrInfo_KEYS{} }
func (m *Srv6SidmgrInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*Srv6SidmgrInfo_KEYS) ProtoMessage()    {}
func (*Srv6SidmgrInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{0}
}

func (m *Srv6SidmgrInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6SidmgrInfo_KEYS.Unmarshal(m, b)
}
func (m *Srv6SidmgrInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6SidmgrInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *Srv6SidmgrInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6SidmgrInfo_KEYS.Merge(m, src)
}
func (m *Srv6SidmgrInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_Srv6SidmgrInfo_KEYS.Size(m)
}
func (m *Srv6SidmgrInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6SidmgrInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6SidmgrInfo_KEYS proto.InternalMessageInfo

type Srv6SidmgrParamsEncapHoplimit struct {
	UseDefault           bool     `protobuf:"varint,1,opt,name=use_default,json=useDefault,proto3" json:"use_default,omitempty"`
	DoPropagate          bool     `protobuf:"varint,2,opt,name=do_propagate,json=doPropagate,proto3" json:"do_propagate,omitempty"`
	Value                uint32   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Srv6SidmgrParamsEncapHoplimit) Reset()         { *m = Srv6SidmgrParamsEncapHoplimit{} }
func (m *Srv6SidmgrParamsEncapHoplimit) String() string { return proto.CompactTextString(m) }
func (*Srv6SidmgrParamsEncapHoplimit) ProtoMessage()    {}
func (*Srv6SidmgrParamsEncapHoplimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{1}
}

func (m *Srv6SidmgrParamsEncapHoplimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit.Unmarshal(m, b)
}
func (m *Srv6SidmgrParamsEncapHoplimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit.Marshal(b, m, deterministic)
}
func (m *Srv6SidmgrParamsEncapHoplimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit.Merge(m, src)
}
func (m *Srv6SidmgrParamsEncapHoplimit) XXX_Size() int {
	return xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit.Size(m)
}
func (m *Srv6SidmgrParamsEncapHoplimit) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6SidmgrParamsEncapHoplimit proto.InternalMessageInfo

func (m *Srv6SidmgrParamsEncapHoplimit) GetUseDefault() bool {
	if m != nil {
		return m.UseDefault
	}
	return false
}

func (m *Srv6SidmgrParamsEncapHoplimit) GetDoPropagate() bool {
	if m != nil {
		return m.DoPropagate
	}
	return false
}

func (m *Srv6SidmgrParamsEncapHoplimit) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Srv6SidmgrParams struct {
	Srv6Enabled                  bool                           `protobuf:"varint,1,opt,name=srv6_enabled,json=srv6Enabled,proto3" json:"srv6_enabled,omitempty"`
	ConfiguredEncapSourceAddress string                         `protobuf:"bytes,2,opt,name=configured_encap_source_address,json=configuredEncapSourceAddress,proto3" json:"configured_encap_source_address,omitempty"`
	DefaultEncapSourceAddress    string                         `protobuf:"bytes,3,opt,name=default_encap_source_address,json=defaultEncapSourceAddress,proto3" json:"default_encap_source_address,omitempty"`
	EncapHopLimit                *Srv6SidmgrParamsEncapHoplimit `protobuf:"bytes,4,opt,name=encap_hop_limit,json=encapHopLimit,proto3" json:"encap_hop_limit,omitempty"`
	EncapTtlPropagate            bool                           `protobuf:"varint,5,opt,name=encap_ttl_propagate,json=encapTtlPropagate,proto3" json:"encap_ttl_propagate,omitempty"`
	IsSidHoldtimeConfigured      bool                           `protobuf:"varint,6,opt,name=is_sid_holdtime_configured,json=isSidHoldtimeConfigured,proto3" json:"is_sid_holdtime_configured,omitempty"`
	SidHoldtimeMinsConfigured    uint32                         `protobuf:"varint,7,opt,name=sid_holdtime_mins_configured,json=sidHoldtimeMinsConfigured,proto3" json:"sid_holdtime_mins_configured,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                       `json:"-"`
	XXX_unrecognized             []byte                         `json:"-"`
	XXX_sizecache                int32                          `json:"-"`
}

func (m *Srv6SidmgrParams) Reset()         { *m = Srv6SidmgrParams{} }
func (m *Srv6SidmgrParams) String() string { return proto.CompactTextString(m) }
func (*Srv6SidmgrParams) ProtoMessage()    {}
func (*Srv6SidmgrParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{2}
}

func (m *Srv6SidmgrParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6SidmgrParams.Unmarshal(m, b)
}
func (m *Srv6SidmgrParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6SidmgrParams.Marshal(b, m, deterministic)
}
func (m *Srv6SidmgrParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6SidmgrParams.Merge(m, src)
}
func (m *Srv6SidmgrParams) XXX_Size() int {
	return xxx_messageInfo_Srv6SidmgrParams.Size(m)
}
func (m *Srv6SidmgrParams) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6SidmgrParams.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6SidmgrParams proto.InternalMessageInfo

func (m *Srv6SidmgrParams) GetSrv6Enabled() bool {
	if m != nil {
		return m.Srv6Enabled
	}
	return false
}

func (m *Srv6SidmgrParams) GetConfiguredEncapSourceAddress() string {
	if m != nil {
		return m.ConfiguredEncapSourceAddress
	}
	return ""
}

func (m *Srv6SidmgrParams) GetDefaultEncapSourceAddress() string {
	if m != nil {
		return m.DefaultEncapSourceAddress
	}
	return ""
}

func (m *Srv6SidmgrParams) GetEncapHopLimit() *Srv6SidmgrParamsEncapHoplimit {
	if m != nil {
		return m.EncapHopLimit
	}
	return nil
}

func (m *Srv6SidmgrParams) GetEncapTtlPropagate() bool {
	if m != nil {
		return m.EncapTtlPropagate
	}
	return false
}

func (m *Srv6SidmgrParams) GetIsSidHoldtimeConfigured() bool {
	if m != nil {
		return m.IsSidHoldtimeConfigured
	}
	return false
}

func (m *Srv6SidmgrParams) GetSidHoldtimeMinsConfigured() uint32 {
	if m != nil {
		return m.SidHoldtimeMinsConfigured
	}
	return 0
}

type Srv6OorSummary struct {
	OutOfResourcesState       string   `protobuf:"bytes,1,opt,name=out_of_resources_state,json=outOfResourcesState,proto3" json:"out_of_resources_state,omitempty"`
	OorYellowFreeSidThreshold uint32   `protobuf:"varint,2,opt,name=oor_yellow_free_sid_threshold,json=oorYellowFreeSidThreshold,proto3" json:"oor_yellow_free_sid_threshold,omitempty"`
	OorGreenFreeSidThreshold  uint32   `protobuf:"varint,3,opt,name=oor_green_free_sid_threshold,json=oorGreenFreeSidThreshold,proto3" json:"oor_green_free_sid_threshold,omitempty"`
	OorGreenCount             uint32   `protobuf:"varint,4,opt,name=oor_green_count,json=oorGreenCount,proto3" json:"oor_green_count,omitempty"`
	OorYellowCount            uint32   `protobuf:"varint,5,opt,name=oor_yellow_count,json=oorYellowCount,proto3" json:"oor_yellow_count,omitempty"`
	OorRedCount               uint32   `protobuf:"varint,6,opt,name=oor_red_count,json=oorRedCount,proto3" json:"oor_red_count,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *Srv6OorSummary) Reset()         { *m = Srv6OorSummary{} }
func (m *Srv6OorSummary) String() string { return proto.CompactTextString(m) }
func (*Srv6OorSummary) ProtoMessage()    {}
func (*Srv6OorSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{3}
}

func (m *Srv6OorSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6OorSummary.Unmarshal(m, b)
}
func (m *Srv6OorSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6OorSummary.Marshal(b, m, deterministic)
}
func (m *Srv6OorSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6OorSummary.Merge(m, src)
}
func (m *Srv6OorSummary) XXX_Size() int {
	return xxx_messageInfo_Srv6OorSummary.Size(m)
}
func (m *Srv6OorSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6OorSummary.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6OorSummary proto.InternalMessageInfo

func (m *Srv6OorSummary) GetOutOfResourcesState() string {
	if m != nil {
		return m.OutOfResourcesState
	}
	return ""
}

func (m *Srv6OorSummary) GetOorYellowFreeSidThreshold() uint32 {
	if m != nil {
		return m.OorYellowFreeSidThreshold
	}
	return 0
}

func (m *Srv6OorSummary) GetOorGreenFreeSidThreshold() uint32 {
	if m != nil {
		return m.OorGreenFreeSidThreshold
	}
	return 0
}

func (m *Srv6OorSummary) GetOorGreenCount() uint32 {
	if m != nil {
		return m.OorGreenCount
	}
	return 0
}

func (m *Srv6OorSummary) GetOorYellowCount() uint32 {
	if m != nil {
		return m.OorYellowCount
	}
	return 0
}

func (m *Srv6OorSummary) GetOorRedCount() uint32 {
	if m != nil {
		return m.OorRedCount
	}
	return 0
}

type Srv6SidmgrSummary struct {
	LocatorsCount            uint32          `protobuf:"varint,1,opt,name=locators_count,json=locatorsCount,proto3" json:"locators_count,omitempty"`
	OperLocatorsCount        uint32          `protobuf:"varint,2,opt,name=oper_locators_count,json=operLocatorsCount,proto3" json:"oper_locators_count,omitempty"`
	SidsCount                uint32          `protobuf:"varint,3,opt,name=sids_count,json=sidsCount,proto3" json:"sids_count,omitempty"`
	StaleSidsCount           uint32          `protobuf:"varint,4,opt,name=stale_sids_count,json=staleSidsCount,proto3" json:"stale_sids_count,omitempty"`
	MaximumSidsCount         uint32          `protobuf:"varint,5,opt,name=maximum_sids_count,json=maximumSidsCount,proto3" json:"maximum_sids_count,omitempty"`
	SidsOutOfResourceSummary *Srv6OorSummary `protobuf:"bytes,6,opt,name=sids_out_of_resource_summary,json=sidsOutOfResourceSummary,proto3" json:"sids_out_of_resource_summary,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}        `json:"-"`
	XXX_unrecognized         []byte          `json:"-"`
	XXX_sizecache            int32           `json:"-"`
}

func (m *Srv6SidmgrSummary) Reset()         { *m = Srv6SidmgrSummary{} }
func (m *Srv6SidmgrSummary) String() string { return proto.CompactTextString(m) }
func (*Srv6SidmgrSummary) ProtoMessage()    {}
func (*Srv6SidmgrSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{4}
}

func (m *Srv6SidmgrSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6SidmgrSummary.Unmarshal(m, b)
}
func (m *Srv6SidmgrSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6SidmgrSummary.Marshal(b, m, deterministic)
}
func (m *Srv6SidmgrSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6SidmgrSummary.Merge(m, src)
}
func (m *Srv6SidmgrSummary) XXX_Size() int {
	return xxx_messageInfo_Srv6SidmgrSummary.Size(m)
}
func (m *Srv6SidmgrSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6SidmgrSummary.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6SidmgrSummary proto.InternalMessageInfo

func (m *Srv6SidmgrSummary) GetLocatorsCount() uint32 {
	if m != nil {
		return m.LocatorsCount
	}
	return 0
}

func (m *Srv6SidmgrSummary) GetOperLocatorsCount() uint32 {
	if m != nil {
		return m.OperLocatorsCount
	}
	return 0
}

func (m *Srv6SidmgrSummary) GetSidsCount() uint32 {
	if m != nil {
		return m.SidsCount
	}
	return 0
}

func (m *Srv6SidmgrSummary) GetStaleSidsCount() uint32 {
	if m != nil {
		return m.StaleSidsCount
	}
	return 0
}

func (m *Srv6SidmgrSummary) GetMaximumSidsCount() uint32 {
	if m != nil {
		return m.MaximumSidsCount
	}
	return 0
}

func (m *Srv6SidmgrSummary) GetSidsOutOfResourceSummary() *Srv6OorSummary {
	if m != nil {
		return m.SidsOutOfResourceSummary
	}
	return nil
}

type MgmtSidmgrString struct {
	String_              string   `protobuf:"bytes,1,opt,name=string,proto3" json:"string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MgmtSidmgrString) Reset()         { *m = MgmtSidmgrString{} }
func (m *MgmtSidmgrString) String() string { return proto.CompactTextString(m) }
func (*MgmtSidmgrString) ProtoMessage()    {}
func (*MgmtSidmgrString) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{5}
}

func (m *MgmtSidmgrString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MgmtSidmgrString.Unmarshal(m, b)
}
func (m *MgmtSidmgrString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MgmtSidmgrString.Marshal(b, m, deterministic)
}
func (m *MgmtSidmgrString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MgmtSidmgrString.Merge(m, src)
}
func (m *MgmtSidmgrString) XXX_Size() int {
	return xxx_messageInfo_MgmtSidmgrString.Size(m)
}
func (m *MgmtSidmgrString) XXX_DiscardUnknown() {
	xxx_messageInfo_MgmtSidmgrString.DiscardUnknown(m)
}

var xxx_messageInfo_MgmtSidmgrString proto.InternalMessageInfo

func (m *MgmtSidmgrString) GetString_() string {
	if m != nil {
		return m.String_
	}
	return ""
}

type MgmtSrv6PlatformSigParams struct {
	MaxSl                uint32   `protobuf:"varint,1,opt,name=max_sl,json=maxSl,proto3" json:"max_sl,omitempty"`
	MaxEndPopSrh         uint32   `protobuf:"varint,2,opt,name=max_end_pop_srh,json=maxEndPopSrh,proto3" json:"max_end_pop_srh,omitempty"`
	MaxTInsert           uint32   `protobuf:"varint,3,opt,name=max_t_insert,json=maxTInsert,proto3" json:"max_t_insert,omitempty"`
	MaxTEncap            uint32   `protobuf:"varint,4,opt,name=max_t_encap,json=maxTEncap,proto3" json:"max_t_encap,omitempty"`
	MaxEndD              uint32   `protobuf:"varint,5,opt,name=max_end_d,json=maxEndD,proto3" json:"max_end_d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MgmtSrv6PlatformSigParams) Reset()         { *m = MgmtSrv6PlatformSigParams{} }
func (m *MgmtSrv6PlatformSigParams) String() string { return proto.CompactTextString(m) }
func (*MgmtSrv6PlatformSigParams) ProtoMessage()    {}
func (*MgmtSrv6PlatformSigParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{6}
}

func (m *MgmtSrv6PlatformSigParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MgmtSrv6PlatformSigParams.Unmarshal(m, b)
}
func (m *MgmtSrv6PlatformSigParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MgmtSrv6PlatformSigParams.Marshal(b, m, deterministic)
}
func (m *MgmtSrv6PlatformSigParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MgmtSrv6PlatformSigParams.Merge(m, src)
}
func (m *MgmtSrv6PlatformSigParams) XXX_Size() int {
	return xxx_messageInfo_MgmtSrv6PlatformSigParams.Size(m)
}
func (m *MgmtSrv6PlatformSigParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MgmtSrv6PlatformSigParams.DiscardUnknown(m)
}

var xxx_messageInfo_MgmtSrv6PlatformSigParams proto.InternalMessageInfo

func (m *MgmtSrv6PlatformSigParams) GetMaxSl() uint32 {
	if m != nil {
		return m.MaxSl
	}
	return 0
}

func (m *MgmtSrv6PlatformSigParams) GetMaxEndPopSrh() uint32 {
	if m != nil {
		return m.MaxEndPopSrh
	}
	return 0
}

func (m *MgmtSrv6PlatformSigParams) GetMaxTInsert() uint32 {
	if m != nil {
		return m.MaxTInsert
	}
	return 0
}

func (m *MgmtSrv6PlatformSigParams) GetMaxTEncap() uint32 {
	if m != nil {
		return m.MaxTEncap
	}
	return 0
}

func (m *MgmtSrv6PlatformSigParams) GetMaxEndD() uint32 {
	if m != nil {
		return m.MaxEndD
	}
	return 0
}

type Srv6PlatformSupport struct {
	Srv6                 bool                       `protobuf:"varint,1,opt,name=srv6,proto3" json:"srv6,omitempty"`
	Tilfa                bool                       `protobuf:"varint,2,opt,name=tilfa,proto3" json:"tilfa,omitempty"`
	MicroloopAvoidance   bool                       `protobuf:"varint,3,opt,name=microloop_avoidance,json=microloopAvoidance,proto3" json:"microloop_avoidance,omitempty"`
	EndFunc              []*MgmtSidmgrString        `protobuf:"bytes,4,rep,name=end_func,json=endFunc,proto3" json:"end_func,omitempty"`
	TransitFunc          []*MgmtSidmgrString        `protobuf:"bytes,5,rep,name=transit_func,json=transitFunc,proto3" json:"transit_func,omitempty"`
	SecurityRule         []*MgmtSidmgrString        `protobuf:"bytes,6,rep,name=security_rule,json=securityRule,proto3" json:"security_rule,omitempty"`
	Counter              []*MgmtSidmgrString        `protobuf:"bytes,7,rep,name=counter,proto3" json:"counter,omitempty"`
	SignaledParameters   *MgmtSrv6PlatformSigParams `protobuf:"bytes,8,opt,name=signaled_parameters,json=signaledParameters,proto3" json:"signaled_parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Srv6PlatformSupport) Reset()         { *m = Srv6PlatformSupport{} }
func (m *Srv6PlatformSupport) String() string { return proto.CompactTextString(m) }
func (*Srv6PlatformSupport) ProtoMessage()    {}
func (*Srv6PlatformSupport) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{7}
}

func (m *Srv6PlatformSupport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6PlatformSupport.Unmarshal(m, b)
}
func (m *Srv6PlatformSupport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6PlatformSupport.Marshal(b, m, deterministic)
}
func (m *Srv6PlatformSupport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6PlatformSupport.Merge(m, src)
}
func (m *Srv6PlatformSupport) XXX_Size() int {
	return xxx_messageInfo_Srv6PlatformSupport.Size(m)
}
func (m *Srv6PlatformSupport) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6PlatformSupport.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6PlatformSupport proto.InternalMessageInfo

func (m *Srv6PlatformSupport) GetSrv6() bool {
	if m != nil {
		return m.Srv6
	}
	return false
}

func (m *Srv6PlatformSupport) GetTilfa() bool {
	if m != nil {
		return m.Tilfa
	}
	return false
}

func (m *Srv6PlatformSupport) GetMicroloopAvoidance() bool {
	if m != nil {
		return m.MicroloopAvoidance
	}
	return false
}

func (m *Srv6PlatformSupport) GetEndFunc() []*MgmtSidmgrString {
	if m != nil {
		return m.EndFunc
	}
	return nil
}

func (m *Srv6PlatformSupport) GetTransitFunc() []*MgmtSidmgrString {
	if m != nil {
		return m.TransitFunc
	}
	return nil
}

func (m *Srv6PlatformSupport) GetSecurityRule() []*MgmtSidmgrString {
	if m != nil {
		return m.SecurityRule
	}
	return nil
}

func (m *Srv6PlatformSupport) GetCounter() []*MgmtSidmgrString {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *Srv6PlatformSupport) GetSignaledParameters() *MgmtSrv6PlatformSigParams {
	if m != nil {
		return m.SignaledParameters
	}
	return nil
}

type Srv6PlatformCapability struct {
	Support              *Srv6PlatformSupport `protobuf:"bytes,1,opt,name=support,proto3" json:"support,omitempty"`
	MaxSid               uint32               `protobuf:"varint,2,opt,name=max_sid,json=maxSid,proto3" json:"max_sid,omitempty"`
	SidHoldtimeMins      uint32               `protobuf:"varint,3,opt,name=sid_holdtime_mins,json=sidHoldtimeMins,proto3" json:"sid_holdtime_mins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Srv6PlatformCapability) Reset()         { *m = Srv6PlatformCapability{} }
func (m *Srv6PlatformCapability) String() string { return proto.CompactTextString(m) }
func (*Srv6PlatformCapability) ProtoMessage()    {}
func (*Srv6PlatformCapability) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{8}
}

func (m *Srv6PlatformCapability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6PlatformCapability.Unmarshal(m, b)
}
func (m *Srv6PlatformCapability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6PlatformCapability.Marshal(b, m, deterministic)
}
func (m *Srv6PlatformCapability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6PlatformCapability.Merge(m, src)
}
func (m *Srv6PlatformCapability) XXX_Size() int {
	return xxx_messageInfo_Srv6PlatformCapability.Size(m)
}
func (m *Srv6PlatformCapability) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6PlatformCapability.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6PlatformCapability proto.InternalMessageInfo

func (m *Srv6PlatformCapability) GetSupport() *Srv6PlatformSupport {
	if m != nil {
		return m.Support
	}
	return nil
}

func (m *Srv6PlatformCapability) GetMaxSid() uint32 {
	if m != nil {
		return m.MaxSid
	}
	return 0
}

func (m *Srv6PlatformCapability) GetSidHoldtimeMins() uint32 {
	if m != nil {
		return m.SidHoldtimeMins
	}
	return 0
}

type Srv6SidmgrInfo struct {
	SidMgrParams         *Srv6SidmgrParams       `protobuf:"bytes,50,opt,name=sid_mgr_params,json=sidMgrParams,proto3" json:"sid_mgr_params,omitempty"`
	SidMgrSummary        *Srv6SidmgrSummary      `protobuf:"bytes,51,opt,name=sid_mgr_summary,json=sidMgrSummary,proto3" json:"sid_mgr_summary,omitempty"`
	PlatformCapabilities *Srv6PlatformCapability `protobuf:"bytes,52,opt,name=platform_capabilities,json=platformCapabilities,proto3" json:"platform_capabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Srv6SidmgrInfo) Reset()         { *m = Srv6SidmgrInfo{} }
func (m *Srv6SidmgrInfo) String() string { return proto.CompactTextString(m) }
func (*Srv6SidmgrInfo) ProtoMessage()    {}
func (*Srv6SidmgrInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_63439f60bd628eba, []int{9}
}

func (m *Srv6SidmgrInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Srv6SidmgrInfo.Unmarshal(m, b)
}
func (m *Srv6SidmgrInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Srv6SidmgrInfo.Marshal(b, m, deterministic)
}
func (m *Srv6SidmgrInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Srv6SidmgrInfo.Merge(m, src)
}
func (m *Srv6SidmgrInfo) XXX_Size() int {
	return xxx_messageInfo_Srv6SidmgrInfo.Size(m)
}
func (m *Srv6SidmgrInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Srv6SidmgrInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Srv6SidmgrInfo proto.InternalMessageInfo

func (m *Srv6SidmgrInfo) GetSidMgrParams() *Srv6SidmgrParams {
	if m != nil {
		return m.SidMgrParams
	}
	return nil
}

func (m *Srv6SidmgrInfo) GetSidMgrSummary() *Srv6SidmgrSummary {
	if m != nil {
		return m.SidMgrSummary
	}
	return nil
}

func (m *Srv6SidmgrInfo) GetPlatformCapabilities() *Srv6PlatformCapability {
	if m != nil {
		return m.PlatformCapabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*Srv6SidmgrInfo_KEYS)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_sidmgr_info_KEYS")
	proto.RegisterType((*Srv6SidmgrParamsEncapHoplimit)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_sidmgr_params_encap_hoplimit")
	proto.RegisterType((*Srv6SidmgrParams)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_sidmgr_params")
	proto.RegisterType((*Srv6OorSummary)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_oor_summary")
	proto.RegisterType((*Srv6SidmgrSummary)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_sidmgr_summary")
	proto.RegisterType((*MgmtSidmgrString)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.mgmt_sidmgr_string")
	proto.RegisterType((*MgmtSrv6PlatformSigParams)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.mgmt_srv6_platform_sig_params")
	proto.RegisterType((*Srv6PlatformSupport)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_platform_support")
	proto.RegisterType((*Srv6PlatformCapability)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_platform_capability")
	proto.RegisterType((*Srv6SidmgrInfo)(nil), "cisco_ios_xr_segment_routing_srv6_oper.srv6.active.manager.srv6_sidmgr_info")
}

func init() { proto.RegisterFile("srv6_sidmgr_info.proto", fileDescriptor_63439f60bd628eba) }

var fileDescriptor_63439f60bd628eba = []byte{
	// 1093 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5b, 0x4f, 0x24, 0x45,
	0x14, 0x4e, 0x03, 0x33, 0x03, 0x67, 0x18, 0x2e, 0x35, 0x0b, 0x34, 0x1b, 0x70, 0xd9, 0x4e, 0xd6,
	0x10, 0x63, 0xc6, 0x04, 0x8c, 0x0f, 0x9a, 0xa8, 0x1b, 0x96, 0x75, 0x8d, 0xec, 0x82, 0x3d, 0xbc,
	0xf0, 0x60, 0x2a, 0x45, 0x77, 0x4d, 0x4f, 0xc5, 0xee, 0xaa, 0xb6, 0xaa, 0x9a, 0x1d, 0x1e, 0x7c,
	0xf4, 0x41, 0xa3, 0xbf, 0x48, 0x13, 0xff, 0x80, 0xbf, 0xc6, 0xc4, 0x77, 0x53, 0x97, 0x9e, 0x19,
	0x2e, 0xc6, 0x64, 0x81, 0xb7, 0xae, 0x73, 0xbe, 0x73, 0xbe, 0x53, 0xe7, 0x56, 0x0d, 0xeb, 0x4a,
	0x5e, 0x7c, 0x82, 0x15, 0x4b, 0x8b, 0x4c, 0x62, 0xc6, 0x07, 0xa2, 0x57, 0x4a, 0xa1, 0x05, 0xfa,
	0x34, 0x61, 0x2a, 0x11, 0x98, 0x09, 0x85, 0x47, 0x12, 0x2b, 0x9a, 0x15, 0x94, 0x6b, 0x2c, 0x45,
	0xa5, 0x19, 0xcf, 0xb0, 0x35, 0x12, 0x25, 0x95, 0x3d, 0xf3, 0xd5, 0x23, 0x89, 0x66, 0x17, 0xb4,
	0x57, 0x10, 0x4e, 0x32, 0x2a, 0xa3, 0x0d, 0x58, 0xbb, 0xee, 0x15, 0x7f, 0x73, 0x78, 0xd6, 0x8f,
	0x7e, 0x84, 0xa7, 0xd3, 0x8a, 0x92, 0x48, 0x52, 0x28, 0x4c, 0x79, 0x42, 0x4a, 0x3c, 0x14, 0x65,
	0xce, 0x0a, 0xa6, 0xd1, 0x13, 0x68, 0x57, 0x8a, 0xe2, 0x94, 0x0e, 0x48, 0x95, 0xeb, 0x30, 0xd8,
	0x09, 0x76, 0xe7, 0x63, 0xa8, 0x14, 0x7d, 0xe1, 0x24, 0xe8, 0x29, 0x2c, 0xa6, 0x02, 0x97, 0x52,
	0x94, 0x24, 0x23, 0x9a, 0x86, 0x33, 0x16, 0xd1, 0x4e, 0xc5, 0x49, 0x2d, 0x42, 0x8f, 0xa0, 0x71,
	0x41, 0xf2, 0x8a, 0x86, 0xb3, 0x3b, 0xc1, 0x6e, 0x27, 0x76, 0x87, 0xe8, 0x9f, 0x59, 0x40, 0x37,
	0xf9, 0x8d, 0x3f, 0x2b, 0xa5, 0x9c, 0x9c, 0xe7, 0x34, 0xf5, 0x8c, 0x6d, 0x23, 0x3b, 0x74, 0x22,
	0x74, 0x08, 0x4f, 0x12, 0xc1, 0x07, 0x2c, 0xab, 0x24, 0x4d, 0x7d, 0xc0, 0x4a, 0x54, 0x32, 0xa1,
	0x98, 0xa4, 0xa9, 0xa4, 0x4a, 0xd9, 0x28, 0x16, 0xe2, 0xad, 0x09, 0xec, 0xd0, 0xa0, 0xfa, 0x16,
	0xf4, 0xdc, 0x61, 0xd0, 0x17, 0xb0, 0xe5, 0xaf, 0x75, 0xbb, 0x8f, 0x59, 0xeb, 0x63, 0xd3, 0x63,
	0x6e, 0x71, 0xf0, 0x53, 0x00, 0xcb, 0xe3, 0x74, 0x61, 0x9b, 0xaf, 0x70, 0x6e, 0x27, 0xd8, 0x6d,
	0xef, 0x7d, 0xd7, 0x7b, 0xf7, 0x82, 0xf5, 0xfe, 0xb7, 0x28, 0x71, 0xc7, 0x9e, 0x5f, 0x89, 0xf2,
	0xc8, 0xd6, 0xa8, 0x07, 0x5d, 0x07, 0xd0, 0x3a, 0x9f, 0xaa, 0x44, 0xc3, 0x66, 0x6e, 0xd5, 0xaa,
	0x4e, 0x75, 0x3e, 0xa9, 0xc7, 0x67, 0xf0, 0x98, 0x29, 0xc3, 0x80, 0x87, 0x22, 0x4f, 0x35, 0x2b,
	0x28, 0x9e, 0x24, 0x2a, 0x6c, 0x5a, 0xb3, 0x0d, 0xa6, 0xfa, 0x2c, 0x7d, 0xe5, 0xf5, 0x07, 0x63,
	0xb5, 0xc9, 0xda, 0x15, 0xcb, 0x82, 0x71, 0x35, 0x6d, 0xde, 0xb2, 0x35, 0xde, 0x54, 0x13, 0xe3,
	0xd7, 0x8c, 0xab, 0x89, 0x83, 0xe8, 0xcf, 0x19, 0x58, 0x71, 0x09, 0x10, 0x12, 0xab, 0xaa, 0x28,
	0x88, 0xbc, 0x44, 0xfb, 0xb0, 0x2e, 0x2a, 0x8d, 0xc5, 0x00, 0x4b, 0xea, 0xca, 0xa0, 0xb0, 0xd2,
	0xe6, 0x16, 0x81, 0xad, 0x42, 0x57, 0x54, 0xfa, 0x78, 0x10, 0xd7, 0xba, 0xbe, 0x51, 0xa1, 0x2f,
	0x61, 0xdb, 0xf8, 0xb8, 0xa4, 0x79, 0x2e, 0xde, 0xe2, 0x81, 0xa4, 0xd4, 0x5e, 0x4a, 0x0f, 0x25,
	0x55, 0x26, 0x3e, 0xdb, 0x05, 0x9d, 0x78, 0x53, 0x08, 0x79, 0x66, 0x31, 0x2f, 0x25, 0xa5, 0x7d,
	0x96, 0x9e, 0xd6, 0x00, 0xf4, 0x39, 0x6c, 0x19, 0x0f, 0x99, 0xa4, 0x94, 0xdf, 0xe6, 0xc0, 0x35,
	0x6c, 0x28, 0x84, 0xfc, 0xca, 0x40, 0x6e, 0xd8, 0xbf, 0x0f, 0xcb, 0x13, 0xfb, 0x44, 0x54, 0xdc,
	0x35, 0x40, 0x27, 0xee, 0xd4, 0x26, 0x07, 0x46, 0x88, 0x76, 0x61, 0x65, 0x2a, 0x52, 0x07, 0x6c,
	0x58, 0xe0, 0xd2, 0x38, 0x38, 0x87, 0x8c, 0xc0, 0x98, 0x62, 0xd3, 0xd8, 0x0e, 0xd6, 0xb4, 0xb0,
	0xb6, 0x10, 0x32, 0xa6, 0xa9, 0xc5, 0x44, 0x7f, 0xcf, 0x40, 0x77, 0xba, 0x49, 0xea, 0x24, 0x3e,
	0x83, 0xa5, 0x5c, 0x24, 0x44, 0x0b, 0xa9, 0xbc, 0x71, 0xe0, 0x82, 0xa9, 0xa5, 0x8e, 0xa2, 0x07,
	0x5d, 0xd3, 0x7b, 0xf8, 0x1a, 0xd6, 0x25, 0x6b, 0xd5, 0xa8, 0x8e, 0xae, 0xe0, 0xb7, 0x01, 0x14,
	0x4b, 0x6b, 0x98, 0x4b, 0xc9, 0x82, 0x91, 0x8c, 0xef, 0xa6, 0x34, 0xc9, 0x6d, 0xea, 0xd4, 0x95,
	0x24, 0x2c, 0x59, 0x79, 0x7f, 0x8c, 0xfc, 0x10, 0x50, 0x41, 0x46, 0xac, 0xa8, 0x8a, 0x69, 0xac,
	0xcb, 0xc3, 0x8a, 0xd7, 0x4c, 0xd0, 0xbf, 0x06, 0xb6, 0xd3, 0x14, 0xbe, 0xd6, 0x18, 0xf5, 0x75,
	0x6d, 0x66, 0xda, 0x7b, 0x47, 0x77, 0x1e, 0xb5, 0xa9, 0x3e, 0x8c, 0x43, 0xc3, 0x78, 0x3c, 0xdd,
	0x6b, 0x7d, 0xa7, 0x89, 0x4c, 0xf0, 0x59, 0xa1, 0xc7, 0x39, 0xd7, 0x92, 0xf1, 0x0c, 0xad, 0x43,
	0xd3, 0x7d, 0xf9, 0x3e, 0xf5, 0xa7, 0xe8, 0xf7, 0x00, 0xb6, 0x1d, 0xdc, 0x30, 0x94, 0x39, 0xd1,
	0x03, 0x21, 0xcd, 0xb5, 0xb3, 0x7a, 0xcf, 0xad, 0x41, 0xb3, 0x20, 0x23, 0xac, 0x72, 0x5f, 0xa4,
	0x46, 0x41, 0x46, 0xfd, 0x1c, 0x3d, 0x83, 0x65, 0x23, 0xa6, 0x3c, 0xc5, 0xa5, 0x28, 0xb1, 0x92,
	0x43, 0x5f, 0x98, 0xc5, 0x82, 0x8c, 0x0e, 0x79, 0x7a, 0x22, 0xca, 0xbe, 0x1c, 0xa2, 0x1d, 0x30,
	0x67, 0xac, 0x31, 0xe3, 0x8a, 0xca, 0xba, 0x2a, 0x50, 0x90, 0xd1, 0xe9, 0xd7, 0x56, 0x82, 0xde,
	0x83, 0xb6, 0x43, 0xd8, 0xf9, 0xf7, 0x15, 0x59, 0x30, 0x00, 0xbb, 0xc9, 0xd0, 0x63, 0x58, 0xa8,
	0x89, 0x52, 0x5f, 0x83, 0x96, 0xa3, 0x78, 0x11, 0xfd, 0xd1, 0xf0, 0x6f, 0xc6, 0x24, 0xf0, 0xaa,
	0x2c, 0x85, 0xd4, 0x08, 0xc1, 0x9c, 0x51, 0xf8, 0xad, 0x6c, 0xbf, 0xcd, 0x7a, 0xd7, 0x2c, 0x1f,
	0x10, 0xbf, 0xfa, 0xdd, 0x01, 0x7d, 0x04, 0xdd, 0x82, 0x25, 0x52, 0xe4, 0x42, 0x94, 0x98, 0x5c,
	0x08, 0x96, 0x12, 0x9e, 0xb8, 0x27, 0x60, 0x3e, 0x46, 0x63, 0xd5, 0xf3, 0x5a, 0x83, 0x18, 0xcc,
	0x9b, 0x60, 0x06, 0x15, 0x4f, 0xc2, 0xb9, 0x9d, 0xd9, 0xdd, 0xf6, 0xde, 0x9b, 0xbb, 0x94, 0xf6,
	0x66, 0xb1, 0xe2, 0x16, 0xe5, 0xe9, 0xcb, 0x8a, 0x27, 0xe8, 0x07, 0x58, 0xd4, 0x92, 0x70, 0xc5,
	0xb4, 0xa3, 0x6b, 0x3c, 0x08, 0x5d, 0xdb, 0x73, 0x58, 0x4a, 0x05, 0x1d, 0x45, 0x93, 0x4a, 0x32,
	0x7d, 0x89, 0x65, 0x95, 0xd3, 0xb0, 0xf9, 0x20, 0x9c, 0x8b, 0x35, 0x49, 0x5c, 0xe5, 0x14, 0x0d,
	0xa1, 0x65, 0x67, 0x8c, 0xca, 0xb0, 0xf5, 0x30, 0x19, 0xf5, 0xee, 0xd1, 0x2f, 0x01, 0x74, 0x15,
	0xcb, 0x38, 0xc9, 0x69, 0xea, 0x3a, 0x9c, 0x6a, 0x2a, 0x55, 0x38, 0x6f, 0x67, 0xf4, 0xec, 0xee,
	0xb4, 0xff, 0x31, 0x46, 0x31, 0xaa, 0x59, 0x4f, 0xc6, 0xa4, 0xd1, 0x5f, 0x01, 0x84, 0x57, 0x0d,
	0x12, 0x52, 0x92, 0x73, 0x96, 0x33, 0x7d, 0x89, 0xbe, 0x87, 0x96, 0x6f, 0x66, 0xdb, 0xc4, 0xed,
	0xbd, 0x6f, 0xef, 0xbc, 0x40, 0xae, 0x4f, 0x49, 0x5c, 0x33, 0xa0, 0x0d, 0x68, 0xd9, 0x21, 0x67,
	0xf5, 0x5b, 0x64, 0x66, 0xbe, 0xcf, 0x52, 0xf4, 0x01, 0xac, 0xde, 0x78, 0x45, 0xfd, 0x10, 0x2f,
	0x5f, 0x7b, 0x3a, 0xa3, 0xdf, 0x66, 0xfd, 0x83, 0x39, 0xf5, 0x07, 0x87, 0x34, 0x2c, 0x19, 0x07,
	0x93, 0x7f, 0x84, 0x70, 0xcf, 0xde, 0xe6, 0xcd, 0xfd, 0xfe, 0x79, 0xc4, 0x8b, 0x8a, 0xa5, 0xaf,
	0x33, 0x79, 0xe2, 0x96, 0xd6, 0x5b, 0x58, 0xae, 0x59, 0xeb, 0x2d, 0xbc, 0x6f, 0x69, 0x8f, 0xef,
	0x8b, 0xb6, 0x5e, 0xc4, 0x1d, 0xc7, 0xeb, 0xb7, 0x2f, 0xfa, 0x39, 0x80, 0xb5, 0x9b, 0xd5, 0x64,
	0x54, 0x85, 0x1f, 0x5b, 0xfe, 0xd3, 0xfb, 0x2b, 0xe2, 0xa4, 0x57, 0xe2, 0x47, 0xb5, 0xf0, 0x60,
	0x8a, 0xf1, 0xbc, 0x69, 0xff, 0xc9, 0xf7, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8c, 0x03,
	0x44, 0xad, 0x0b, 0x00, 0x00,
}
