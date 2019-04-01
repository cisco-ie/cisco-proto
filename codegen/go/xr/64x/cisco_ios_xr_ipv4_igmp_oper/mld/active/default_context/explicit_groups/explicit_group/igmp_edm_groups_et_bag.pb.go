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
// source: igmp_edm_groups_et_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_active_default_context_explicit_groups_explicit_group

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

type IgmpEdmGroupsEtBag_KEYS struct {
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	SourceAddress        string   `protobuf:"bytes,3,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmGroupsEtBag_KEYS) Reset()         { *m = IgmpEdmGroupsEtBag_KEYS{} }
func (m *IgmpEdmGroupsEtBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsEtBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmGroupsEtBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cc168d6bd20e9e, []int{0}
}

func (m *IgmpEdmGroupsEtBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsEtBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsEtBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmGroupsEtBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS.Size(m)
}
func (m *IgmpEdmGroupsEtBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsEtBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmGroupsEtBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *IgmpEdmGroupsEtBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *IgmpEdmGroupsEtBag_KEYS) GetSourceAddress() string {
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
	return fileDescriptor_63cc168d6bd20e9e, []int{1}
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
	return fileDescriptor_63cc168d6bd20e9e, []int{2}
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

type IgmpEdmGroupsHostBag struct {
	Address              *IgmpAddrtype   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Uptime               uint32          `protobuf:"varint,2,opt,name=uptime,proto3" json:"uptime,omitempty"`
	IsExclude            bool            `protobuf:"varint,3,opt,name=is_exclude,json=isExclude,proto3" json:"is_exclude,omitempty"`
	ExpirationTime       uint32          `protobuf:"varint,4,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	SourceCount          uint32          `protobuf:"varint,5,opt,name=source_count,json=sourceCount,proto3" json:"source_count,omitempty"`
	SourceAddress        []*IgmpAddrtype `protobuf:"bytes,6,rep,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IgmpEdmGroupsHostBag) Reset()         { *m = IgmpEdmGroupsHostBag{} }
func (m *IgmpEdmGroupsHostBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsHostBag) ProtoMessage()    {}
func (*IgmpEdmGroupsHostBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cc168d6bd20e9e, []int{3}
}

func (m *IgmpEdmGroupsHostBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsHostBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsHostBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsHostBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsHostBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsHostBag.Merge(m, src)
}
func (m *IgmpEdmGroupsHostBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsHostBag.Size(m)
}
func (m *IgmpEdmGroupsHostBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsHostBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsHostBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsHostBag) GetAddress() *IgmpAddrtype {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *IgmpEdmGroupsHostBag) GetUptime() uint32 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *IgmpEdmGroupsHostBag) GetIsExclude() bool {
	if m != nil {
		return m.IsExclude
	}
	return false
}

func (m *IgmpEdmGroupsHostBag) GetExpirationTime() uint32 {
	if m != nil {
		return m.ExpirationTime
	}
	return 0
}

func (m *IgmpEdmGroupsHostBag) GetSourceCount() uint32 {
	if m != nil {
		return m.SourceCount
	}
	return 0
}

func (m *IgmpEdmGroupsHostBag) GetSourceAddress() []*IgmpAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

type IgmpEdmGroupsEtBag struct {
	IncludeHosts         uint32                  `protobuf:"varint,50,opt,name=include_hosts,json=includeHosts,proto3" json:"include_hosts,omitempty"`
	ExcludeHosts         uint32                  `protobuf:"varint,51,opt,name=exclude_hosts,json=excludeHosts,proto3" json:"exclude_hosts,omitempty"`
	GroupInfo            *IgmpEdmGroupsBag       `protobuf:"bytes,52,opt,name=group_info,json=groupInfo,proto3" json:"group_info,omitempty"`
	Host                 []*IgmpEdmGroupsHostBag `protobuf:"bytes,53,rep,name=host,proto3" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *IgmpEdmGroupsEtBag) Reset()         { *m = IgmpEdmGroupsEtBag{} }
func (m *IgmpEdmGroupsEtBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsEtBag) ProtoMessage()    {}
func (*IgmpEdmGroupsEtBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_63cc168d6bd20e9e, []int{4}
}

func (m *IgmpEdmGroupsEtBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsEtBag.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsEtBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsEtBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsEtBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsEtBag.Merge(m, src)
}
func (m *IgmpEdmGroupsEtBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsEtBag.Size(m)
}
func (m *IgmpEdmGroupsEtBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsEtBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsEtBag proto.InternalMessageInfo

func (m *IgmpEdmGroupsEtBag) GetIncludeHosts() uint32 {
	if m != nil {
		return m.IncludeHosts
	}
	return 0
}

func (m *IgmpEdmGroupsEtBag) GetExcludeHosts() uint32 {
	if m != nil {
		return m.ExcludeHosts
	}
	return 0
}

func (m *IgmpEdmGroupsEtBag) GetGroupInfo() *IgmpEdmGroupsBag {
	if m != nil {
		return m.GroupInfo
	}
	return nil
}

func (m *IgmpEdmGroupsEtBag) GetHost() []*IgmpEdmGroupsHostBag {
	if m != nil {
		return m.Host
	}
	return nil
}

func init() {
	proto.RegisterType((*IgmpEdmGroupsEtBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.explicit_groups.explicit_group.igmp_edm_groups_et_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.explicit_groups.explicit_group.igmp_addrtype")
	proto.RegisterType((*IgmpEdmGroupsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.explicit_groups.explicit_group.igmp_edm_groups_bag")
	proto.RegisterType((*IgmpEdmGroupsHostBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.explicit_groups.explicit_group.igmp_edm_groups_host_bag")
	proto.RegisterType((*IgmpEdmGroupsEtBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.active.default_context.explicit_groups.explicit_group.igmp_edm_groups_et_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_et_bag.proto", fileDescriptor_63cc168d6bd20e9e) }

var fileDescriptor_63cc168d6bd20e9e = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x05, 0x6d, 0x57, 0x8f, 0x2b, 0x52, 0xb6, 0xa7, 0x80, 0x4d, 0xa3, 0x35, 0x20, 0xab, 0x28,
	0x2a, 0x14, 0x85, 0x80, 0xca, 0xae, 0x81, 0x76, 0x67, 0x14, 0x2e, 0xfa, 0x72, 0x16, 0x94, 0x13,
	0x38, 0xab, 0x01, 0x45, 0x0e, 0x95, 0x71, 0x48, 0x0e, 0x31, 0x33, 0xb4, 0xe4, 0x4d, 0x36, 0x59,
	0x19, 0x01, 0x82, 0xfc, 0x47, 0x76, 0xf9, 0xb6, 0x7c, 0x40, 0x30, 0x77, 0x28, 0xc5, 0x92, 0x95,
	0xa5, 0x95, 0x95, 0xcd, 0x73, 0xcf, 0xdc, 0xc7, 0x99, 0x7b, 0x48, 0xc1, 0xf7, 0x7c, 0x9c, 0x15,
	0x94, 0xc5, 0x19, 0x1d, 0x4b, 0x51, 0x16, 0x8a, 0x32, 0x4d, 0x47, 0xe1, 0xb8, 0x5f, 0x48, 0xa1,
	0x05, 0x79, 0x1a, 0x71, 0x15, 0x09, 0xca, 0x85, 0xa2, 0x53, 0x49, 0x79, 0x71, 0x73, 0x42, 0x91,
	0x2f, 0x0a, 0x26, 0xfb, 0x59, 0x1a, 0xf7, 0xc3, 0x48, 0xf3, 0x1b, 0xd6, 0x8f, 0x59, 0x12, 0x96,
	0xa9, 0xa6, 0x91, 0xc8, 0x35, 0x9b, 0xea, 0x3e, 0x9b, 0x16, 0x29, 0x8f, 0xb8, 0xae, 0x92, 0x2e,
	0x3d, 0x77, 0xdf, 0x39, 0xf0, 0xdd, 0xea, 0xba, 0xf4, 0xbf, 0xf3, 0xe7, 0x43, 0xf2, 0x03, 0x78,
	0x88, 0xd2, 0x30, 0x8e, 0x25, 0x53, 0xca, 0x77, 0x3a, 0x4e, 0xaf, 0x19, 0xb8, 0x08, 0x9e, 0x59,
	0x8c, 0xfc, 0x08, 0x6d, 0x9e, 0x6b, 0x26, 0x93, 0x30, 0x62, 0x34, 0x0f, 0x33, 0xe6, 0x6f, 0x20,
	0xcb, 0x9b, 0xa3, 0x4f, 0xc2, 0x8c, 0x19, 0x9a, 0x12, 0xa5, 0x8c, 0xd8, 0x3c, 0xd9, 0xa6, 0xa5,
	0x59, 0xb4, 0xca, 0xd6, 0xcd, 0xc1, 0xc3, 0x8e, 0x0c, 0x49, 0xdf, 0x16, 0x8c, 0xec, 0x43, 0x3d,
	0x4c, 0x6c, 0x5e, 0x5b, 0xbd, 0x16, 0x26, 0x98, 0xf0, 0x08, 0x5c, 0x14, 0x62, 0x96, 0xce, 0x56,
	0x6d, 0x19, 0x6c, 0xd6, 0x9a, 0xa5, 0x9c, 0x2e, 0x55, 0x34, 0x94, 0xd3, 0x59, 0xbd, 0x0f, 0x75,
	0xf8, 0x76, 0x59, 0x82, 0x51, 0x38, 0x26, 0x6f, 0x1d, 0xd8, 0x59, 0x98, 0x9d, 0x4e, 0x25, 0x36,
	0xd0, 0x1a, 0xc4, 0xfd, 0x47, 0xb9, 0x8d, 0xfe, 0xc2, 0xdc, 0x41, 0xfb, 0xbe, 0xc8, 0x57, 0x92,
	0xfc, 0x0c, 0xbb, 0x8b, 0x32, 0x9b, 0x86, 0xec, 0xcc, 0xdb, 0x0b, 0x4a, 0x5f, 0x49, 0xb2, 0x07,
	0xb5, 0xb2, 0xd0, 0x3c, 0x63, 0x38, 0xf1, 0x56, 0x50, 0x3d, 0x91, 0x9f, 0x60, 0x9b, 0x4d, 0x0b,
	0x2e, 0x43, 0xcd, 0x45, 0x4e, 0x91, 0xb0, 0xd5, 0x71, 0x7a, 0xbb, 0x41, 0xfb, 0x33, 0x7c, 0x69,
	0x88, 0x77, 0x0e, 0x78, 0x69, 0xa8, 0x34, 0x95, 0xac, 0x10, 0x52, 0x33, 0xe9, 0x7f, 0xb3, 0xc6,
	0xd1, 0x5d, 0x53, 0x3a, 0xa8, 0x2a, 0x93, 0x3f, 0xe0, 0x60, 0xce, 0xd6, 0x32, 0x8c, 0x5e, 0xf2,
	0x7c, 0x4c, 0x59, 0x1e, 0x8e, 0x52, 0x16, 0xfb, 0xb5, 0x8e, 0xd3, 0x6b, 0x04, 0xfb, 0x33, 0xc2,
	0x65, 0x15, 0x3f, 0xb7, 0x61, 0xd2, 0x01, 0x97, 0x2b, 0xaa, 0x58, 0x9a, 0xd0, 0x6b, 0xc1, 0x73,
	0xbf, 0x8e, 0x74, 0xe0, 0x6a, 0xc8, 0xd2, 0xe4, 0x5f, 0xc1, 0x73, 0x72, 0x08, 0x20, 0xc5, 0x84,
	0x2a, 0x1d, 0xea, 0x52, 0xf9, 0x0d, 0xd4, 0xb3, 0x29, 0xc5, 0x64, 0x88, 0x00, 0xe9, 0x82, 0xc7,
	0x15, 0x4d, 0xc5, 0x84, 0x66, 0x2c, 0x13, 0xf2, 0xd6, 0x6f, 0x62, 0x86, 0x16, 0x57, 0xff, 0x8b,
	0xc9, 0x05, 0x42, 0xe4, 0x17, 0x20, 0x52, 0x94, 0x9a, 0x49, 0x9a, 0xf0, 0xd4, 0xfc, 0xc9, 0x44,
	0xcc, 0x7c, 0xe8, 0x38, 0x3d, 0x2f, 0xd8, 0xb1, 0x91, 0xbf, 0x30, 0x70, 0x21, 0x62, 0x46, 0xde,
	0x38, 0x0f, 0x8c, 0xd0, 0x5a, 0xa3, 0xb6, 0x8b, 0x76, 0x23, 0xbf, 0xc3, 0x81, 0x48, 0x63, 0x26,
	0xe9, 0x0b, 0xa1, 0x34, 0xbd, 0x61, 0x52, 0x71, 0x91, 0xff, 0x8a, 0xab, 0x21, 0x7d, 0x17, 0x47,
	0xd8, 0x43, 0xc2, 0xdf, 0x42, 0xe9, 0x67, 0x55, 0xd8, 0xac, 0x88, 0x5c, 0x7d, 0x74, 0x50, 0x1d,
	0xf5, 0x56, 0x1f, 0x1d, 0xd8, 0xa3, 0x07, 0xd0, 0xe0, 0xca, 0xf4, 0xc4, 0x62, 0xbf, 0x8d, 0x82,
	0xd6, 0xb9, 0x3a, 0x33, 0x8f, 0xe6, 0x95, 0x63, 0x6e, 0xac, 0x2c, 0x0a, 0xd3, 0x1f, 0x8b, 0xfd,
	0x6d, 0x8c, 0xbb, 0x5c, 0x0d, 0xe7, 0x58, 0xf7, 0xfd, 0x26, 0xf8, 0xcb, 0xa6, 0xc5, 0x2e, 0x8c,
	0x73, 0x5f, 0x41, 0xfd, 0xfe, 0xeb, 0x6a, 0x5d, 0xc2, 0xce, 0x8a, 0xde, 0x33, 0xdf, 0x06, 0x8a,
	0x30, 0x33, 0xdf, 0x21, 0x00, 0x57, 0x94, 0x4d, 0xa3, 0xb4, 0x8c, 0xad, 0x31, 0x1b, 0x41, 0x93,
	0xab, 0x73, 0x0b, 0x7c, 0xc9, 0x9b, 0xde, 0x03, 0x6f, 0x1e, 0x81, 0x5b, 0xed, 0x4f, 0x24, 0xca,
	0x5c, 0xa3, 0x33, 0xbd, 0xa0, 0x65, 0xb1, 0x3f, 0x0d, 0xb4, 0x6a, 0xc7, 0x6a, 0x9d, 0xcd, 0xaf,
	0xb4, 0x63, 0xdd, 0x8f, 0x1b, 0xb0, 0xb7, 0xfa, 0x2b, 0x83, 0xb7, 0x9d, 0xe3, 0xfc, 0x78, 0x7f,
	0xca, 0x1f, 0xe0, 0x30, 0x6e, 0x05, 0x9a, 0xcd, 0x51, 0x86, 0x54, 0xa9, 0x56, 0x91, 0x8e, 0x2d,
	0xa9, 0x02, 0x2d, 0xe9, 0xce, 0x01, 0xb0, 0xef, 0x6b, 0x9e, 0x27, 0xc2, 0x3f, 0xc1, 0x9b, 0xbf,
	0x7e, 0xcc, 0x71, 0x17, 0x3f, 0x18, 0x41, 0x13, 0xff, 0xff, 0x27, 0x4f, 0x04, 0x79, 0xed, 0xc0,
	0x96, 0xe9, 0xd4, 0xff, 0x0d, 0x45, 0x17, 0x6b, 0xea, 0x62, 0xe6, 0x80, 0x00, 0x8b, 0x8f, 0x6a,
	0xf8, 0xd3, 0xe1, 0xf8, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0xec, 0xfd, 0x57, 0x5a, 0x08,
	0x00, 0x00,
}
