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

package cisco_ios_xr_ipv4_igmp_oper_igmp_active_default_context_detail_groups_detail_group

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
	proto.RegisterType((*IgmpEdmGroupsDetailBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.default_context.detail_groups.detail_group.igmp_edm_groups_detail_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.default_context.detail_groups.detail_group.igmp_addrtype")
	proto.RegisterType((*IgmpEdmGroupsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.default_context.detail_groups.detail_group.igmp_edm_groups_bag")
	proto.RegisterType((*IgmpEdmGroupsSourceBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.default_context.detail_groups.detail_group.igmp_edm_groups_source_bag")
	proto.RegisterType((*IgmpEdmGroupsDetailBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.active.default_context.detail_groups.detail_group.igmp_edm_groups_detail_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_detail_bag.proto", fileDescriptor_9fe3409596f06866) }

var fileDescriptor_9fe3409596f06866 = []byte{
	// 794 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0x5b, 0x8f, 0xdb, 0x44,
	0x14, 0x96, 0x9b, 0x36, 0x97, 0xc9, 0x65, 0x9b, 0x69, 0xb5, 0xf5, 0x82, 0x10, 0x21, 0x08, 0x11,
	0x21, 0x14, 0x44, 0xb6, 0x54, 0x82, 0xb7, 0x3e, 0x6c, 0xc5, 0xa5, 0xe5, 0xc1, 0xa9, 0xa0, 0x3c,
	0x8d, 0x26, 0xf6, 0x71, 0x38, 0x60, 0x7b, 0xac, 0x99, 0xf1, 0x26, 0xfb, 0x07, 0x56, 0x48, 0xbc,
	0xa1, 0xfd, 0x13, 0xfc, 0x4b, 0x34, 0x67, 0x9c, 0xcd, 0x7a, 0x09, 0xe2, 0x05, 0xb6, 0x4f, 0xf1,
	0x9c, 0xf3, 0x9d, 0xcb, 0x7c, 0xe7, 0x7c, 0x76, 0xd8, 0x04, 0xd7, 0x79, 0x29, 0x20, 0xc9, 0xc5,
	0x5a, 0xab, 0xaa, 0x34, 0x22, 0x01, 0x2b, 0x31, 0x13, 0x2b, 0xb9, 0x9e, 0x97, 0x5a, 0x59, 0xc5,
	0xa3, 0x18, 0x4d, 0xac, 0x04, 0x2a, 0x23, 0xb6, 0x5a, 0x60, 0x79, 0xfe, 0x54, 0x50, 0x8c, 0x2a,
	0x41, 0xcf, 0xdd, 0xd3, 0x5c, 0xc6, 0x16, 0xcf, 0x61, 0x9e, 0x40, 0x2a, 0xab, 0xcc, 0x8a, 0x58,
	0x15, 0x16, 0xb6, 0x76, 0x5e, 0x67, 0xf2, 0x79, 0x1b, 0xa7, 0xe9, 0x55, 0xc0, 0xde, 0xff, 0xe7,
	0xc2, 0xe2, 0xbb, 0xb3, 0x9f, 0x96, 0xfc, 0x43, 0x36, 0x24, 0x8f, 0x90, 0x49, 0xa2, 0xc1, 0x98,
	0x30, 0x98, 0x04, 0xb3, 0x5e, 0x34, 0x20, 0xe3, 0x73, 0x6f, 0xe3, 0x1f, 0xb1, 0x11, 0x16, 0x16,
	0x74, 0x2a, 0x63, 0x10, 0x85, 0xcc, 0x21, 0xbc, 0x47, 0xa8, 0xe1, 0xb5, 0xf5, 0x7b, 0x99, 0x83,
	0x83, 0x19, 0x55, 0xe9, 0x18, 0xae, 0x93, 0xb5, 0x3c, 0xcc, 0x5b, 0xeb, 0x6c, 0xd3, 0x82, 0x0d,
	0xa9, 0x2b, 0x07, 0xb2, 0x17, 0x25, 0xf0, 0x27, 0xac, 0x23, 0x53, 0x9f, 0xd7, 0x57, 0x6f, 0xcb,
	0x94, 0x12, 0x7e, 0xc0, 0x06, 0xc4, 0xc4, 0x2e, 0x9d, 0xaf, 0xda, 0x77, 0xb6, 0x5d, 0x6b, 0x1e,
	0xf2, 0xec, 0x56, 0x45, 0x07, 0x79, 0xb6, 0xab, 0xf7, 0x67, 0x87, 0x3d, 0xba, 0x4d, 0xc3, 0x4a,
	0xae, 0xf9, 0xef, 0x01, 0x7b, 0xd8, 0xb8, 0xbb, 0xd8, 0x6a, 0x6a, 0xa0, 0xbf, 0x90, 0xf3, 0xff,
	0x7e, 0x1c, 0xf3, 0xc6, 0xa5, 0xa3, 0xd1, 0x4d, 0x86, 0xdf, 0x68, 0xfe, 0x09, 0x1b, 0x37, 0x39,
	0x76, 0xdd, 0xf8, 0x0b, 0x1f, 0x35, 0x68, 0x7e, 0xa3, 0xf9, 0x31, 0x6b, 0x57, 0xa5, 0xc5, 0x1c,
	0xe8, 0xba, 0xf7, 0xa3, 0xfa, 0xc4, 0x3f, 0x66, 0x47, 0xb0, 0x2d, 0x51, 0x4b, 0x8b, 0xaa, 0x10,
	0x04, 0xb8, 0x3f, 0x09, 0x66, 0xe3, 0x68, 0xb4, 0x37, 0xbf, 0x76, 0xc0, 0xcb, 0x80, 0x0d, 0x33,
	0x69, 0xac, 0xd0, 0x50, 0x2a, 0x6d, 0x41, 0x87, 0x0f, 0xee, 0xea, 0xde, 0x03, 0x57, 0x37, 0xaa,
	0xcb, 0xf2, 0xaf, 0xd8, 0x09, 0x6c, 0xcb, 0x0c, 0x63, 0xb4, 0xc2, 0x6a, 0x19, 0xff, 0x8a, 0xc5,
	0x5a, 0x40, 0x21, 0x57, 0x19, 0x24, 0x61, 0x7b, 0x12, 0xcc, 0xba, 0xd1, 0x93, 0x1d, 0xe0, 0x75,
	0xed, 0x3f, 0xf3, 0x6e, 0x3e, 0x61, 0x03, 0x34, 0xc2, 0x40, 0x96, 0x8a, 0x5f, 0x14, 0x16, 0x61,
	0x87, 0xe0, 0x0c, 0xcd, 0x12, 0xb2, 0xf4, 0x5b, 0x85, 0x05, 0x7f, 0x8f, 0x31, 0xad, 0x36, 0xc2,
	0x58, 0x69, 0x2b, 0x13, 0x76, 0x89, 0xcc, 0x9e, 0x56, 0x9b, 0x25, 0x19, 0xf8, 0x94, 0x0d, 0xd1,
	0x88, 0x4c, 0x6d, 0x44, 0x0e, 0xb9, 0xd2, 0x17, 0x61, 0x8f, 0x32, 0xf4, 0xd1, 0xbc, 0x54, 0x9b,
	0x57, 0x64, 0xe2, 0x9f, 0x32, 0xae, 0x55, 0x65, 0x41, 0x8b, 0x14, 0x33, 0xf7, 0x93, 0xab, 0x04,
	0x42, 0x36, 0x09, 0x66, 0xc3, 0xe8, 0xa1, 0xf7, 0xbc, 0x20, 0xc7, 0x2b, 0x95, 0x00, 0xff, 0x2d,
	0xf8, 0x9b, 0x04, 0xfa, 0x77, 0x45, 0x6c, 0x53, 0x65, 0xfc, 0x4b, 0x76, 0xa2, 0xb2, 0x04, 0xb4,
	0xf8, 0x59, 0x19, 0x2b, 0xce, 0x41, 0x1b, 0x54, 0xc5, 0xe7, 0xb4, 0x14, 0x3a, 0x1c, 0x50, 0xff,
	0xc7, 0x04, 0xf8, 0x5a, 0x19, 0xfb, 0x43, 0xed, 0x76, 0xcb, 0xa1, 0x0f, 0x87, 0x2e, 0xea, 0xd0,
	0xe1, 0xe1, 0xd0, 0x85, 0x0f, 0x3d, 0x61, 0x5d, 0x34, 0xae, 0x27, 0x48, 0xc2, 0x11, 0xb1, 0xd9,
	0x41, 0xf3, 0xdc, 0x1d, 0xdd, 0x9b, 0xc6, 0x8d, 0xab, 0x2a, 0x4b, 0xd7, 0x1f, 0x24, 0xe1, 0x11,
	0xf9, 0x07, 0x68, 0x96, 0xd7, 0xb6, 0xe9, 0x1f, 0x2d, 0xf6, 0xce, 0x6d, 0xad, 0xd6, 0x84, 0x3a,
	0xc9, 0x1e, 0xe0, 0x37, 0x78, 0x4b, 0xfc, 0xee, 0x35, 0x78, 0xef, 0xdf, 0x34, 0xd8, 0x3a, 0xa8,
	0x41, 0x4f, 0x55, 0xa6, 0x62, 0x99, 0x91, 0x4a, 0x89, 0xaa, 0x97, 0xee, 0xc8, 0xdf, 0x65, 0x3d,
	0x34, 0x42, 0x43, 0xae, 0x2c, 0x90, 0x32, 0xbb, 0x51, 0x17, 0x4d, 0x44, 0x67, 0xb7, 0xd4, 0x68,
	0x44, 0xaa, 0xf4, 0x46, 0xea, 0x9d, 0x46, 0x7a, 0x68, 0x5e, 0x78, 0x43, 0xad, 0x8a, 0x0d, 0xd4,
	0xd2, 0xde, 0xab, 0xe2, 0x47, 0xf0, 0xaa, 0xe3, 0x8f, 0xd9, 0x83, 0x34, 0x93, 0x6b, 0x2f, 0x88,
	0x71, 0xe4, 0x0f, 0x8d, 0xc9, 0xf5, 0x1a, 0x93, 0x9b, 0x5e, 0x1d, 0x18, 0xca, 0xfe, 0x3b, 0xc2,
	0x4f, 0xd9, 0xb1, 0xeb, 0xd6, 0xab, 0x04, 0xb6, 0x71, 0x56, 0x25, 0xe0, 0x65, 0xb2, 0xa0, 0x3c,
	0x8f, 0xd0, 0x44, 0xe4, 0x3c, 0xf3, 0x3e, 0x52, 0xca, 0x67, 0xec, 0x31, 0x1a, 0xbf, 0x60, 0x8d,
	0x90, 0x53, 0x0a, 0x19, 0xa3, 0x71, 0xbb, 0x75, 0x33, 0xe0, 0x32, 0x60, 0xcc, 0xbf, 0xad, 0xb1,
	0x48, 0x55, 0xf8, 0x94, 0xc6, 0xbe, 0xfe, 0xdf, 0xc6, 0xde, 0xfc, 0x56, 0x44, 0x3d, 0x7a, 0xfe,
	0xa6, 0x48, 0x95, 0x6b, 0xa4, 0xed, 0x57, 0x21, 0xfc, 0x62, 0xd2, 0x9a, 0xf5, 0x17, 0xc5, 0x5d,
	0x34, 0xb1, 0x17, 0x41, 0x54, 0x57, 0x5f, 0xb5, 0xe9, 0x9f, 0xc3, 0xe9, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x0d, 0x9c, 0x0f, 0x5d, 0x08, 0x00, 0x00,
}
