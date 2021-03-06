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
// source: igmp_edm_groups_detail_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_default_context_detail_groups_detail_group

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

type IgmpEdmGroupsDetailBag_KEYS struct {
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SourceAddress        string   `protobuf:"bytes,3,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmGroupsDetailBag_KEYS) Reset()         { *m = IgmpEdmGroupsDetailBag_KEYS{} }
func (m *IgmpEdmGroupsDetailBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsDetailBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmGroupsDetailBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe3409596f06866, []int{0}
}

func (m *IgmpEdmGroupsDetailBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsDetailBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsDetailBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmGroupsDetailBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS.Size(m)
}
func (m *IgmpEdmGroupsDetailBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsDetailBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmGroupsDetailBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *IgmpEdmGroupsDetailBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IgmpEdmGroupsDetailBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

type IgmpAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpAddrtype) Reset()         { *m = IgmpAddrtype{} }
func (m *IgmpAddrtype) String() string { return proto.CompactTextString(m) }
func (*IgmpAddrtype) ProtoMessage()    {}
func (*IgmpAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe3409596f06866, []int{1}
}

func (m *IgmpAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpAddrtype.Unmarshal(m, b)
}
func (m *IgmpAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpAddrtype.Marshal(b, m, deterministic)
}
func (m *IgmpAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpAddrtype.Merge(m, src)
}
func (m *IgmpAddrtype) XXX_Size() int {
	return xxx_messageInfo_IgmpAddrtype.Size(m)
}
func (m *IgmpAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpAddrtype proto.InternalMessageInfo

func (m *IgmpAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IgmpAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type IgmpEdmGroupsBag struct {
	GroupAddressXr          *IgmpAddrtype `protobuf:"bytes,1,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	InterfaceNameXr         string        `protobuf:"bytes,2,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	Uptime                  uint64        `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	ExpirationTime          int32         `protobuf:"zigzag32,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	LastReporter            *IgmpAddrtype `protobuf:"bytes,5,opt,name=last_reporter,json=lastReporter,proto3" json:"last_reporter,omitempty"`
	ExplicitTrackingEnabled bool          `protobuf:"varint,6,opt,name=explicit_tracking_enabled,json=explicitTrackingEnabled,proto3" json:"explicit_tracking_enabled,omitempty"`
	IsSelfJoin              bool          `protobuf:"varint,7,opt,name=is_self_join,json=isSelfJoin,proto3" json:"is_self_join,omitempty"`
	RowStatus               string        `protobuf:"bytes,8,opt,name=row_status,json=rowStatus,proto3" json:"row_status,omitempty"`
	IsLowMemory             bool          `protobuf:"varint,9,opt,name=is_low_memory,json=isLowMemory,proto3" json:"is_low_memory,omitempty"`
	RouterFilterMode        uint32        `protobuf:"varint,10,opt,name=router_filter_mode,json=routerFilterMode,proto3" json:"router_filter_mode,omitempty"`
	SourceAddress           *IgmpAddrtype `protobuf:"bytes,11,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	OlderHostVersion1Timer  uint32        `protobuf:"varint,12,opt,name=older_host_version1_timer,json=olderHostVersion1Timer,proto3" json:"older_host_version1_timer,omitempty"`
	OlderHostVersion2Timer  uint32        `protobuf:"varint,13,opt,name=older_host_version2_timer,json=olderHostVersion2Timer,proto3" json:"older_host_version2_timer,omitempty"`
	IsAdded                 bool          `protobuf:"varint,14,opt,name=is_added,json=isAdded,proto3" json:"is_added,omitempty"`
	IsSuppressed            bool          `protobuf:"varint,15,opt,name=is_suppressed,json=isSuppressed,proto3" json:"is_suppressed,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}      `json:"-"`
	XXX_unrecognized        []byte        `json:"-"`
	XXX_sizecache           int32         `json:"-"`
}

func (m *IgmpEdmGroupsBag) Reset()         { *m = IgmpEdmGroupsBag{} }
func (m *IgmpEdmGroupsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsBag) ProtoMessage()    {}
func (*IgmpEdmGroupsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe3409596f06866, []int{2}
}

func (m *IgmpEdmGroupsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsBag.Merge(m, src)
}
func (m *IgmpEdmGroupsBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsBag.Size(m)
}
func (m *IgmpEdmGroupsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsBag) GetGroupAddressXr() *IgmpAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *IgmpEdmGroupsBag) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *IgmpEdmGroupsBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *IgmpEdmGroupsBag) GetExpirationTime() int32 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *IgmpEdmGroupsBag) GetLastReporter() *IgmpAddrtype {
	if m != nil {
		return m.LastReporter
	}
	return nil
}

func (m *IgmpEdmGroupsBag) GetExplicitTrackingEnabled() bool {
	if m != nil {
		return m.ExplicitTrackingEnabled
	}
	return false
}

func (m *IgmpEdmGroupsBag) GetIsSelfJoin() bool {
	if m != nil {
		return m.IsSelfJoin
	}
	return false
}

func (m *IgmpEdmGroupsBag) GetRowStatus() string {
	if m != nil {
		return m.RowStatus
	}
	return ""
}

func (m *IgmpEdmGroupsBag) GetIsLowMemory() bool {
	if m != nil {
		return m.IsLowMemory
	}
	return false
}

func (m *IgmpEdmGroupsBag) GetRouterFilterMode() uint32 {
	if m != nil {
		return m.RouterFilterMode
	}
	return 0
}

func (m *IgmpEdmGroupsBag) GetSourceAddress() *IgmpAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *IgmpEdmGroupsBag) GetOlderHostVersion1Timer() uint32 {
	if m != nil {
		return m.OlderHostVersion1Timer
	}
	return 0
}

func (m *IgmpEdmGroupsBag) GetOlderHostVersion2Timer() uint32 {
	if m != nil {
		return m.OlderHostVersion2Timer
	}
	return 0
}

func (m *IgmpEdmGroupsBag) GetIsAdded() bool {
	if m != nil {
		return m.IsAdded
	}
	return false
}

func (m *IgmpEdmGroupsBag) GetIsSuppressed() bool {
	if m != nil {
		return m.IsSuppressed
	}
	return false
}

type IgmpEdmGroupsSourceBag struct {
	SourceAddress        *IgmpAddrtype `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Uptime               uint64        `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	ExpirationTime       int32         `protobuf:"zigzag32,3,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	IsLocal              bool          `protobuf:"varint,4,opt,name=is_local,json=isLocal,proto3" json:"is_local,omitempty"`
	IsRemote             bool          `protobuf:"varint,5,opt,name=is_remote,json=isRemote,proto3" json:"is_remote,omitempty"`
	IsForward            bool          `protobuf:"varint,6,opt,name=is_forward,json=isForward,proto3" json:"is_forward,omitempty"`
	IsWeReport           bool          `protobuf:"varint,7,opt,name=is_we_report,json=isWeReport,proto3" json:"is_we_report,omitempty"`
	Flags                int32         `protobuf:"zigzag32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	IsAdded              bool          `protobuf:"varint,9,opt,name=is_added,json=isAdded,proto3" json:"is_added,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpEdmGroupsSourceBag) Reset()         { *m = IgmpEdmGroupsSourceBag{} }
func (m *IgmpEdmGroupsSourceBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsSourceBag) ProtoMessage()    {}
func (*IgmpEdmGroupsSourceBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe3409596f06866, []int{3}
}

func (m *IgmpEdmGroupsSourceBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsSourceBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsSourceBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsSourceBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsSourceBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsSourceBag.Merge(m, src)
}
func (m *IgmpEdmGroupsSourceBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsSourceBag.Size(m)
}
func (m *IgmpEdmGroupsSourceBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsSourceBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsSourceBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsSourceBag) GetSourceAddress() *IgmpAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *IgmpEdmGroupsSourceBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *IgmpEdmGroupsSourceBag) GetExpirationTime() int32 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *IgmpEdmGroupsSourceBag) GetIsLocal() bool {
	if m != nil {
		return m.IsLocal
	}
	return false
}

func (m *IgmpEdmGroupsSourceBag) GetIsRemote() bool {
	if m != nil {
		return m.IsRemote
	}
	return false
}

func (m *IgmpEdmGroupsSourceBag) GetIsForward() bool {
	if m != nil {
		return m.IsForward
	}
	return false
}

func (m *IgmpEdmGroupsSourceBag) GetIsWeReport() bool {
	if m != nil {
		return m.IsWeReport
	}
	return false
}

func (m *IgmpEdmGroupsSourceBag) GetFlags() int32 {
	if m != nil {
		return m.Flags
	}
	return 0
}

func (m *IgmpEdmGroupsSourceBag) GetIsAdded() bool {
	if m != nil {
		return m.IsAdded
	}
	return false
}

type IgmpEdmGroupsDetailBag struct {
	IsRouterExcludeMode  bool                      `protobuf:"varint,50,opt,name=is_router_exclude_mode,json=isRouterExcludeMode,proto3" json:"is_router_exclude_mode,omitempty"`
	IsHostExcludeMode    bool                      `protobuf:"varint,51,opt,name=is_host_exclude_mode,json=isHostExcludeMode,proto3" json:"is_host_exclude_mode,omitempty"`
	GroupInfo            *IgmpEdmGroupsBag         `protobuf:"bytes,52,opt,name=group_info,json=groupInfo,proto3" json:"group_info,omitempty"`
	Source               []*IgmpEdmGroupsSourceBag `protobuf:"bytes,53,rep,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *IgmpEdmGroupsDetailBag) Reset()         { *m = IgmpEdmGroupsDetailBag{} }
func (m *IgmpEdmGroupsDetailBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsDetailBag) ProtoMessage()    {}
func (*IgmpEdmGroupsDetailBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fe3409596f06866, []int{4}
}

func (m *IgmpEdmGroupsDetailBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsDetailBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsDetailBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsDetailBag.Merge(m, src)
}
func (m *IgmpEdmGroupsDetailBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsDetailBag.Size(m)
}
func (m *IgmpEdmGroupsDetailBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsDetailBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsDetailBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsDetailBag) GetIsRouterExcludeMode() bool {
	if m != nil {
		return m.IsRouterExcludeMode
	}
	return false
}

func (m *IgmpEdmGroupsDetailBag) GetIsHostExcludeMode() bool {
	if m != nil {
		return m.IsHostExcludeMode
	}
	return false
}

func (m *IgmpEdmGroupsDetailBag) GetGroupInfo() *IgmpEdmGroupsBag {
	if m != nil {
		return m.GroupInfo
	}
	return nil
}

func (m *IgmpEdmGroupsDetailBag) GetSource() []*IgmpEdmGroupsSourceBag {
	if m != nil {
		return m.Source
	}
	return nil
}

func init() {
	proto.RegisterType((*IgmpEdmGroupsDetailBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.detail_groups.detail_group.igmp_edm_groups_detail_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.detail_groups.detail_group.igmp_addrtype")
	proto.RegisterType((*IgmpEdmGroupsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.detail_groups.detail_group.igmp_edm_groups_bag")
	proto.RegisterType((*IgmpEdmGroupsSourceBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.detail_groups.detail_group.igmp_edm_groups_source_bag")
	proto.RegisterType((*IgmpEdmGroupsDetailBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.detail_groups.detail_group.igmp_edm_groups_detail_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_detail_bag.proto", fileDescriptor_9fe3409596f06866) }

var fileDescriptor_9fe3409596f06866 = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x4b, 0x6f, 0xc3, 0x44,
	0x10, 0x96, 0x9b, 0x36, 0x8f, 0xcd, 0xa3, 0xcd, 0xb6, 0x6a, 0x5d, 0x10, 0x22, 0x04, 0x21, 0x22,
	0x84, 0x82, 0x48, 0x4b, 0x25, 0xb8, 0xf5, 0xd0, 0x8a, 0x47, 0xcb, 0xc1, 0xa9, 0xa0, 0x9c, 0x56,
	0x1b, 0x7b, 0x9c, 0x2e, 0xd8, 0x5e, 0x6b, 0x77, 0xdd, 0xa4, 0xbf, 0xa0, 0x42, 0xe2, 0x06, 0xe2,
	0x57, 0xf0, 0x23, 0xd1, 0xce, 0x3a, 0x4d, 0x5d, 0x82, 0x38, 0xb5, 0x9c, 0xe2, 0x79, 0xcf, 0x7c,
	0x33, 0x9f, 0x1d, 0x32, 0x10, 0xf3, 0x34, 0x67, 0x10, 0xa5, 0x6c, 0xae, 0x64, 0x91, 0x6b, 0x16,
	0x81, 0xe1, 0x22, 0x61, 0x33, 0x3e, 0x1f, 0xe7, 0x4a, 0x1a, 0x49, 0xa7, 0xa1, 0xd0, 0xa1, 0x64,
	0x42, 0x6a, 0xb6, 0x54, 0x4c, 0xe4, 0xf7, 0xa7, 0x0c, 0x63, 0x64, 0x0e, 0x6a, 0x6c, 0x9f, 0xc6,
	0xda, 0xf0, 0x2c, 0x9a, 0x3d, 0x8c, 0x23, 0x88, 0x79, 0x91, 0x18, 0x16, 0xca, 0xcc, 0xc0, 0xd2,
	0x8c, 0xcb, 0x54, 0x2e, 0x71, 0x45, 0x1a, 0xfe, 0xe1, 0x91, 0xf7, 0xff, 0xbd, 0x32, 0xfb, 0xee,
	0xe2, 0xa7, 0x29, 0xfd, 0x90, 0x74, 0xd1, 0xc2, 0x78, 0x14, 0x29, 0xd0, 0xda, 0xf7, 0x06, 0xde,
	0xa8, 0x15, 0x74, 0x50, 0x79, 0xee, 0x74, 0xf4, 0x23, 0xd2, 0x13, 0x99, 0x01, 0x15, 0xf3, 0x10,
	0x58, 0xc6, 0x53, 0xf0, 0xb7, 0xd0, 0xab, 0xfb, 0xa4, 0xfd, 0x9e, 0xa7, 0x60, 0xdd, 0xb4, 0x2c,
	0x54, 0x08, 0x4f, 0xc9, 0x6a, 0xce, 0xcd, 0x69, 0xcb, 0x6c, 0xc3, 0x8c, 0x74, 0xb1, 0x2b, 0xeb,
	0x64, 0x1e, 0x72, 0xa0, 0x47, 0xa4, 0xc1, 0x63, 0x97, 0xd7, 0x55, 0xaf, 0xf3, 0x18, 0x13, 0x7e,
	0x40, 0x3a, 0x08, 0xc5, 0x2a, 0x9d, 0xab, 0xda, 0xb6, 0xba, 0x55, 0x6b, 0xce, 0xe5, 0xec, 0x45,
	0x45, 0xeb, 0x72, 0xb6, 0xaa, 0xf7, 0x57, 0x83, 0xec, 0xbf, 0x84, 0x61, 0xc6, 0xe7, 0xf4, 0x37,
	0x8f, 0xec, 0x55, 0x66, 0x67, 0x4b, 0x85, 0x0d, 0xb4, 0x27, 0xb3, 0xf1, 0x2b, 0xec, 0x63, 0x5c,
	0x99, 0x3a, 0xe8, 0x3d, 0x87, 0xf8, 0x56, 0xd1, 0x4f, 0x48, 0xbf, 0x0a, 0xb2, 0x6d, 0xc7, 0x4d,
	0xbc, 0x5b, 0xc1, 0xf9, 0x56, 0xd1, 0x43, 0x52, 0x2f, 0x72, 0x23, 0x52, 0xc0, 0x79, 0xb7, 0x83,
	0x52, 0xa2, 0x1f, 0x93, 0x5d, 0x58, 0xe6, 0x42, 0x71, 0x23, 0x64, 0xc6, 0xd0, 0x61, 0x7b, 0xe0,
	0x8d, 0xfa, 0x41, 0x6f, 0xad, 0xbe, 0xb1, 0x8e, 0x8f, 0x1e, 0xe9, 0x26, 0x5c, 0x1b, 0xa6, 0x20,
	0x97, 0xca, 0x80, 0xf2, 0x77, 0xde, 0x6c, 0xf0, 0x8e, 0x2d, 0x1c, 0x94, 0x75, 0xe9, 0x57, 0xe4,
	0x18, 0x96, 0x79, 0x22, 0x42, 0x61, 0x98, 0x51, 0x3c, 0xfc, 0x45, 0x64, 0x73, 0x06, 0x19, 0x9f,
	0x25, 0x10, 0xf9, 0xf5, 0x81, 0x37, 0x6a, 0x06, 0x47, 0x2b, 0x87, 0x9b, 0xd2, 0x7e, 0xe1, 0xcc,
	0x74, 0x40, 0x3a, 0x42, 0x33, 0x0d, 0x49, 0xcc, 0x7e, 0x96, 0x22, 0xf3, 0x1b, 0xe8, 0x4e, 0x84,
	0x9e, 0x42, 0x12, 0x7f, 0x2b, 0x45, 0x46, 0xdf, 0x23, 0x44, 0xc9, 0x05, 0xd3, 0x86, 0x9b, 0x42,
	0xfb, 0x4d, 0x44, 0xb3, 0xa5, 0xe4, 0x62, 0x8a, 0x0a, 0x3a, 0x24, 0x5d, 0xa1, 0x59, 0x22, 0x17,
	0x2c, 0x85, 0x54, 0xaa, 0x07, 0xbf, 0x85, 0x19, 0xda, 0x42, 0x5f, 0xc9, 0xc5, 0x35, 0xaa, 0xe8,
	0xa7, 0x84, 0x2a, 0x59, 0x18, 0x50, 0x2c, 0x16, 0x89, 0xfd, 0x49, 0x65, 0x04, 0x3e, 0x19, 0x78,
	0xa3, 0x6e, 0xb0, 0xe7, 0x2c, 0x97, 0x68, 0xb8, 0x96, 0x11, 0xd0, 0x5f, 0xbd, 0x7f, 0x90, 0xa0,
	0xfd, 0x66, 0xc8, 0x56, 0x89, 0x46, 0xbf, 0x24, 0xc7, 0x32, 0x89, 0x40, 0xb1, 0x3b, 0xa9, 0x0d,
	0xbb, 0x07, 0xa5, 0x85, 0xcc, 0x3e, 0xc7, 0xb3, 0x50, 0x7e, 0x07, 0x07, 0x38, 0x44, 0x87, 0xaf,
	0xa5, 0x36, 0x3f, 0x94, 0x66, 0x7b, 0x1e, 0x6a, 0x73, 0xe8, 0xa4, 0x0c, 0xed, 0x6e, 0x0e, 0x9d,
	0xb8, 0xd0, 0x63, 0xd2, 0x14, 0xda, 0xf6, 0x04, 0x91, 0xdf, 0x43, 0x38, 0x1b, 0x42, 0x9f, 0x5b,
	0xd1, 0xbe, 0x6c, 0xec, 0xbe, 0x8a, 0x3c, 0xb7, 0xfd, 0x41, 0xe4, 0xef, 0xa2, 0xbd, 0x23, 0xf4,
	0xf4, 0x49, 0x37, 0xfc, 0xbd, 0x46, 0xde, 0x79, 0x49, 0xd7, 0x12, 0x51, 0xcb, 0xda, 0x0d, 0x00,
	0x7b, 0xff, 0x17, 0xc0, 0x6b, 0x1a, 0x6e, 0xfd, 0x17, 0x0d, 0x6b, 0x1b, 0x69, 0xe8, 0xb0, 0x4a,
	0x64, 0xc8, 0x13, 0x24, 0x2a, 0x62, 0x75, 0x65, 0x45, 0xfa, 0x2e, 0x69, 0x09, 0xcd, 0x14, 0xa4,
	0xd2, 0x00, 0x92, 0xb3, 0x19, 0x34, 0x85, 0x0e, 0x50, 0xb6, 0x67, 0x2d, 0x34, 0x8b, 0xa5, 0x5a,
	0x70, 0xb5, 0x62, 0x49, 0x4b, 0xe8, 0x4b, 0xa7, 0x28, 0x79, 0xb1, 0x80, 0x92, 0xdd, 0x6b, 0x5e,
	0xfc, 0x08, 0x8e, 0x77, 0xf4, 0x80, 0xec, 0xc4, 0x09, 0x9f, 0x3b, 0x4a, 0xf4, 0x03, 0x27, 0x54,
	0x56, 0xd7, 0xaa, 0xac, 0x6e, 0xf8, 0xe7, 0x86, 0xad, 0xac, 0xbf, 0x25, 0xf4, 0x84, 0x1c, 0xda,
	0x6e, 0x1d, 0x4f, 0x60, 0x19, 0x26, 0x45, 0x04, 0x8e, 0x28, 0x13, 0xcc, 0xb3, 0x2f, 0x74, 0x80,
	0xc6, 0x0b, 0x67, 0x43, 0xae, 0x7c, 0x46, 0x0e, 0x84, 0x76, 0x17, 0x56, 0x09, 0x39, 0xc1, 0x90,
	0xbe, 0xd0, 0xf6, 0xb8, 0x9e, 0x07, 0x3c, 0x7a, 0x84, 0xb8, 0x37, 0xb6, 0xc8, 0x62, 0xe9, 0x9f,
	0xe2, 0xde, 0xef, 0x5e, 0x6f, 0xef, 0xd5, 0x0f, 0x46, 0xd0, 0xc2, 0xe7, 0x6f, 0xb2, 0x58, 0xda,
	0x4e, 0xea, 0xee, 0x16, 0xfc, 0x2f, 0x06, 0xb5, 0x51, 0x7b, 0x22, 0xdf, 0xa4, 0x8b, 0x35, 0x0f,
	0x82, 0xb2, 0xfc, 0xac, 0x8e, 0x7f, 0x20, 0x4e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0e,
	0x5f, 0xb2, 0x64, 0x08, 0x00, 0x00,
}
