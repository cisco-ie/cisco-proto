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
// source: igmp_edm_groups_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_igmp_standby_default_context_groups_group

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

type IgmpEdmGroupsBag_KEYS struct {
	GroupAddress         string   `protobuf:"bytes,1,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmGroupsBag_KEYS) Reset()         { *m = IgmpEdmGroupsBag_KEYS{} }
func (m *IgmpEdmGroupsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmGroupsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6312508847e15206, []int{0}
}

func (m *IgmpEdmGroupsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmGroupsBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmGroupsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmGroupsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmGroupsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmGroupsBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmGroupsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmGroupsBag_KEYS.Size(m)
}
func (m *IgmpEdmGroupsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmGroupsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmGroupsBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmGroupsBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *IgmpEdmGroupsBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
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
	return fileDescriptor_6312508847e15206, []int{1}
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
	GroupAddressXr          *IgmpAddrtype `protobuf:"bytes,50,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	InterfaceNameXr         string        `protobuf:"bytes,51,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	Uptime                  uint64        `protobuf:"varint,52,opt,name=uptime,proto3" json:"uptime,omitempty"`
	ExpirationTime          int32         `protobuf:"zigzag32,53,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	LastReporter            *IgmpAddrtype `protobuf:"bytes,54,opt,name=last_reporter,json=lastReporter,proto3" json:"last_reporter,omitempty"`
	ExplicitTrackingEnabled bool          `protobuf:"varint,55,opt,name=explicit_tracking_enabled,json=explicitTrackingEnabled,proto3" json:"explicit_tracking_enabled,omitempty"`
	IsSelfJoin              bool          `protobuf:"varint,56,opt,name=is_self_join,json=isSelfJoin,proto3" json:"is_self_join,omitempty"`
	RowStatus               string        `protobuf:"bytes,57,opt,name=row_status,json=rowStatus,proto3" json:"row_status,omitempty"`
	IsLowMemory             bool          `protobuf:"varint,58,opt,name=is_low_memory,json=isLowMemory,proto3" json:"is_low_memory,omitempty"`
	RouterFilterMode        uint32        `protobuf:"varint,59,opt,name=router_filter_mode,json=routerFilterMode,proto3" json:"router_filter_mode,omitempty"`
	SourceAddress           *IgmpAddrtype `protobuf:"bytes,60,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	OlderHostVersion1Timer  uint32        `protobuf:"varint,61,opt,name=older_host_version1_timer,json=olderHostVersion1Timer,proto3" json:"older_host_version1_timer,omitempty"`
	OlderHostVersion2Timer  uint32        `protobuf:"varint,62,opt,name=older_host_version2_timer,json=olderHostVersion2Timer,proto3" json:"older_host_version2_timer,omitempty"`
	IsAdded                 bool          `protobuf:"varint,63,opt,name=is_added,json=isAdded,proto3" json:"is_added,omitempty"`
	IsSuppressed            bool          `protobuf:"varint,64,opt,name=is_suppressed,json=isSuppressed,proto3" json:"is_suppressed,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}      `json:"-"`
	XXX_unrecognized        []byte        `json:"-"`
	XXX_sizecache           int32         `json:"-"`
}

func (m *IgmpEdmGroupsBag) Reset()         { *m = IgmpEdmGroupsBag{} }
func (m *IgmpEdmGroupsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmGroupsBag) ProtoMessage()    {}
func (*IgmpEdmGroupsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6312508847e15206, []int{2}
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

func init() {
	proto.RegisterType((*IgmpEdmGroupsBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.groups.group.igmp_edm_groups_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.groups.group.igmp_addrtype")
	proto.RegisterType((*IgmpEdmGroupsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.igmp.standby.default_context.groups.group.igmp_edm_groups_bag")
}

func init() { proto.RegisterFile("igmp_edm_groups_bag.proto", fileDescriptor_6312508847e15206) }

var fileDescriptor_6312508847e15206 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x6f, 0xd3, 0x3c,
	0x14, 0x55, 0xbe, 0x0f, 0x75, 0x9b, 0xd7, 0x74, 0x9b, 0x91, 0x36, 0xf7, 0x01, 0x29, 0x14, 0x21,
	0x2a, 0x84, 0x22, 0xd1, 0x8d, 0xc1, 0xc6, 0xcf, 0x3d, 0x0c, 0x21, 0x60, 0x3c, 0xa4, 0x13, 0x1a,
	0x4f, 0x96, 0x9b, 0xdc, 0x14, 0x43, 0x12, 0x5b, 0xd7, 0xce, 0xda, 0x0a, 0x89, 0x7f, 0x1d, 0x64,
	0xa7, 0x2d, 0x54, 0xea, 0x23, 0x7b, 0x4a, 0x7c, 0xee, 0xb9, 0xf7, 0xdc, 0x7b, 0x7c, 0x65, 0xd2,
	0x95, 0xe3, 0x52, 0x73, 0xc8, 0x4a, 0x3e, 0x46, 0x55, 0x6b, 0xc3, 0x47, 0x62, 0x1c, 0x6b, 0x54,
	0x56, 0xd1, 0xf3, 0x54, 0x9a, 0x54, 0x71, 0xa9, 0x0c, 0x9f, 0x22, 0x97, 0xfa, 0xfa, 0x88, 0x7b,
	0xb2, 0xd2, 0x80, 0xb1, 0xfb, 0x8b, 0x8d, 0x15, 0x55, 0x36, 0x9a, 0xc5, 0x19, 0xe4, 0xa2, 0x2e,
	0x2c, 0x4f, 0x55, 0x65, 0x61, 0x6a, 0xe3, 0xa6, 0x54, 0xf3, 0xe9, 0xe5, 0x84, 0xad, 0xd1, 0xe0,
	0x1f, 0xce, 0xbf, 0x0c, 0xe9, 0x3d, 0x12, 0x7a, 0x88, 0x8b, 0x2c, 0x43, 0x30, 0x86, 0x05, 0x51,
	0xd0, 0xdf, 0x4a, 0xda, 0x1e, 0x3c, 0x6b, 0x30, 0x7a, 0x9f, 0x74, 0x64, 0x65, 0x01, 0x73, 0x91,
	0x02, 0xaf, 0x44, 0x09, 0xec, 0x3f, 0xcf, 0x0a, 0x97, 0xe8, 0x27, 0x51, 0x42, 0xaf, 0x22, 0xa1,
	0xd7, 0x71, 0xa5, 0xec, 0x4c, 0x03, 0x3d, 0x20, 0x1b, 0x22, 0x6f, 0x12, 0x9a, 0xb2, 0x2d, 0x91,
	0x3b, 0x26, 0xbd, 0x4b, 0xda, 0x7e, 0x9a, 0x85, 0x68, 0x53, 0x6e, 0xdb, 0x61, 0x0b, 0xcd, 0x86,
	0x72, 0xbc, 0xa4, 0xfc, 0xbf, 0xa4, 0x1c, 0xcf, 0x29, 0xbd, 0x5f, 0x2d, 0x72, 0x7b, 0xcd, 0x60,
	0xf4, 0x27, 0xd9, 0x5d, 0x99, 0x89, 0x4f, 0x91, 0x0d, 0xa2, 0xa0, 0xbf, 0x3d, 0xb8, 0x8c, 0xff,
	0x89, 0xa3, 0xf1, 0xca, 0x98, 0x49, 0xe7, 0x6f, 0xb3, 0xae, 0x90, 0x3e, 0x24, 0x7b, 0xab, 0x76,
	0xb9, 0x06, 0x0e, 0x7d, 0xff, 0x3b, 0x2b, 0x8e, 0x5d, 0x21, 0xdd, 0x27, 0xad, 0x5a, 0x5b, 0x59,
	0x02, 0x3b, 0x8a, 0x82, 0xfe, 0xad, 0x64, 0x7e, 0xa2, 0x0f, 0xc8, 0x0e, 0x4c, 0xb5, 0x44, 0x61,
	0xa5, 0xaa, 0xb8, 0x27, 0x3c, 0x89, 0x82, 0xfe, 0x5e, 0xd2, 0xf9, 0x03, 0x5f, 0x3a, 0xe2, 0x8c,
	0x84, 0x85, 0x30, 0x96, 0x23, 0x68, 0x85, 0x16, 0x90, 0x1d, 0xdf, 0xe0, 0xa4, 0x6d, 0x27, 0x95,
	0xcc, 0x95, 0xe8, 0x29, 0xe9, 0xc2, 0x54, 0x17, 0x32, 0x95, 0x96, 0x5b, 0x14, 0xe9, 0x77, 0x59,
	0x8d, 0x39, 0x54, 0x62, 0x54, 0x40, 0xc6, 0x9e, 0x46, 0x41, 0x7f, 0x33, 0x39, 0x58, 0x10, 0x2e,
	0xe7, 0xf1, 0xf3, 0x26, 0x4c, 0x23, 0xd2, 0x96, 0x86, 0x1b, 0x28, 0x72, 0xfe, 0x4d, 0xc9, 0x8a,
	0x3d, 0xf3, 0x74, 0x22, 0xcd, 0x10, 0x8a, 0xfc, 0xbd, 0x92, 0x15, 0xbd, 0x43, 0x08, 0xaa, 0x09,
	0x37, 0x56, 0xd8, 0xda, 0xb0, 0x13, 0x6f, 0xdf, 0x16, 0xaa, 0xc9, 0xd0, 0x03, 0xb4, 0x47, 0x42,
	0x69, 0x78, 0xa1, 0x26, 0xbc, 0x84, 0x52, 0xe1, 0x8c, 0x9d, 0xfa, 0x0a, 0xdb, 0xd2, 0x7c, 0x54,
	0x93, 0x0b, 0x0f, 0xd1, 0x47, 0x84, 0xa2, 0xaa, 0x2d, 0x20, 0xcf, 0x65, 0xe1, 0x3e, 0xa5, 0xca,
	0x80, 0x3d, 0x8f, 0x82, 0x7e, 0x98, 0xec, 0x36, 0x91, 0xb7, 0x3e, 0x70, 0xa1, 0x32, 0xa0, 0x3f,
	0x48, 0xc7, 0xa8, 0x1a, 0x53, 0x58, 0xee, 0xdc, 0x8b, 0x1b, 0xb4, 0x32, 0x6c, 0xb4, 0x16, 0xeb,
	0x7e, 0x42, 0xba, 0xaa, 0xc8, 0x00, 0xf9, 0x57, 0x65, 0x2c, 0xbf, 0x06, 0x34, 0x52, 0x55, 0x8f,
	0xfd, 0xc5, 0x23, 0x7b, 0xe9, 0x3b, 0xde, 0xf7, 0x84, 0x77, 0xca, 0xd8, 0xcf, 0xf3, 0xb0, 0x5b,
	0x00, 0x5c, 0x9f, 0x3a, 0x98, 0xa7, 0xbe, 0x5a, 0x9f, 0x3a, 0x68, 0x52, 0xbb, 0x64, 0x53, 0x1a,
	0xd7, 0x13, 0x64, 0xec, 0xb5, 0xf7, 0x6f, 0x43, 0x9a, 0x33, 0x77, 0x74, 0x0f, 0x83, 0xbb, 0xa0,
	0x5a, 0x6b, 0xd7, 0x1f, 0x64, 0xec, 0x8d, 0x8f, 0xb7, 0xa5, 0x19, 0x2e, 0xb1, 0x51, 0xcb, 0xbf,
	0x53, 0x87, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x9f, 0x88, 0x50, 0xc4, 0x04, 0x00, 0x00,
}
