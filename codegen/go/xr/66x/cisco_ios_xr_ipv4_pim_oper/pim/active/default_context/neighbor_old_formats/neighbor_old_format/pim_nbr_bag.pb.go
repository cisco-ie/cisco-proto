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
// source: pim_nbr_bag.proto

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_neighbor_old_formats_neighbor_old_format

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

type PimNbrBag_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,2,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimNbrBag_KEYS) Reset()         { *m = PimNbrBag_KEYS{} }
func (m *PimNbrBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimNbrBag_KEYS) ProtoMessage()    {}
func (*PimNbrBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f976d211d835a3f, []int{0}
}

func (m *PimNbrBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNbrBag_KEYS.Unmarshal(m, b)
}
func (m *PimNbrBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNbrBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimNbrBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNbrBag_KEYS.Merge(m, src)
}
func (m *PimNbrBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimNbrBag_KEYS.Size(m)
}
func (m *PimNbrBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNbrBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimNbrBag_KEYS proto.InternalMessageInfo

func (m *PimNbrBag_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PimNbrBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type PimAddrtype struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimAddrtype) Reset()         { *m = PimAddrtype{} }
func (m *PimAddrtype) String() string { return proto.CompactTextString(m) }
func (*PimAddrtype) ProtoMessage()    {}
func (*PimAddrtype) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f976d211d835a3f, []int{1}
}

func (m *PimAddrtype) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimAddrtype.Unmarshal(m, b)
}
func (m *PimAddrtype) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimAddrtype.Marshal(b, m, deterministic)
}
func (m *PimAddrtype) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimAddrtype.Merge(m, src)
}
func (m *PimAddrtype) XXX_Size() int {
	return xxx_messageInfo_PimAddrtype.Size(m)
}
func (m *PimAddrtype) XXX_DiscardUnknown() {
	xxx_messageInfo_PimAddrtype.DiscardUnknown(m)
}

var xxx_messageInfo_PimAddrtype proto.InternalMessageInfo

func (m *PimAddrtype) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *PimAddrtype) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PimAddrtype) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PimNbrBag struct {
	InterfaceNameXr        string         `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	NeighborAddressXr      []*PimAddrtype `protobuf:"bytes,51,rep,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	Uptime                 uint64         `protobuf:"varint,52,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Expires                uint64         `protobuf:"varint,53,opt,name=expires,proto3" json:"expires,omitempty"`
	ExpiryTimer            uint64         `protobuf:"varint,54,opt,name=expiry_timer,json=expiryTimer,proto3" json:"expiry_timer,omitempty"`
	IsThisNeighborUs       bool           `protobuf:"varint,55,opt,name=is_this_neighbor_us,json=isThisNeighborUs,proto3" json:"is_this_neighbor_us,omitempty"`
	IsThisNeighborDr       bool           `protobuf:"varint,56,opt,name=is_this_neighbor_dr,json=isThisNeighborDr,proto3" json:"is_this_neighbor_dr,omitempty"`
	IsDrPriorityCapable    bool           `protobuf:"varint,57,opt,name=is_dr_priority_capable,json=isDrPriorityCapable,proto3" json:"is_dr_priority_capable,omitempty"`
	DrPriority             uint32         `protobuf:"varint,58,opt,name=dr_priority,json=drPriority,proto3" json:"dr_priority,omitempty"`
	IsBidirectionalCapable bool           `protobuf:"varint,59,opt,name=is_bidirectional_capable,json=isBidirectionalCapable,proto3" json:"is_bidirectional_capable,omitempty"`
	IsProxyCapable         bool           `protobuf:"varint,60,opt,name=is_proxy_capable,json=isProxyCapable,proto3" json:"is_proxy_capable,omitempty"`
	IsBatchAssertsCapable  bool           `protobuf:"varint,61,opt,name=is_batch_asserts_capable,json=isBatchAssertsCapable,proto3" json:"is_batch_asserts_capable,omitempty"`
	IsEcmpRedirectCapable  bool           `protobuf:"varint,62,opt,name=is_ecmp_redirect_capable,json=isEcmpRedirectCapable,proto3" json:"is_ecmp_redirect_capable,omitempty"`
	IsBfdState             bool           `protobuf:"varint,63,opt,name=is_bfd_state,json=isBfdState,proto3" json:"is_bfd_state,omitempty"`
	PropagationDelay       uint32         `protobuf:"varint,64,opt,name=propagation_delay,json=propagationDelay,proto3" json:"propagation_delay,omitempty"`
	OverrideInterval       uint32         `protobuf:"varint,65,opt,name=override_interval,json=overrideInterval,proto3" json:"override_interval,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}       `json:"-"`
	XXX_unrecognized       []byte         `json:"-"`
	XXX_sizecache          int32          `json:"-"`
}

func (m *PimNbrBag) Reset()         { *m = PimNbrBag{} }
func (m *PimNbrBag) String() string { return proto.CompactTextString(m) }
func (*PimNbrBag) ProtoMessage()    {}
func (*PimNbrBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f976d211d835a3f, []int{2}
}

func (m *PimNbrBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimNbrBag.Unmarshal(m, b)
}
func (m *PimNbrBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimNbrBag.Marshal(b, m, deterministic)
}
func (m *PimNbrBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimNbrBag.Merge(m, src)
}
func (m *PimNbrBag) XXX_Size() int {
	return xxx_messageInfo_PimNbrBag.Size(m)
}
func (m *PimNbrBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimNbrBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimNbrBag proto.InternalMessageInfo

func (m *PimNbrBag) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *PimNbrBag) GetNeighborAddressXr() []*PimAddrtype {
	if m != nil {
		return m.NeighborAddressXr
	}
	return nil
}

func (m *PimNbrBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *PimNbrBag) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func (m *PimNbrBag) GetExpiryTimer() uint64 {
	if m != nil {
		return m.ExpiryTimer
	}
	return 0
}

func (m *PimNbrBag) GetIsThisNeighborUs() bool {
	if m != nil {
		return m.IsThisNeighborUs
	}
	return false
}

func (m *PimNbrBag) GetIsThisNeighborDr() bool {
	if m != nil {
		return m.IsThisNeighborDr
	}
	return false
}

func (m *PimNbrBag) GetIsDrPriorityCapable() bool {
	if m != nil {
		return m.IsDrPriorityCapable
	}
	return false
}

func (m *PimNbrBag) GetDrPriority() uint32 {
	if m != nil {
		return m.DrPriority
	}
	return 0
}

func (m *PimNbrBag) GetIsBidirectionalCapable() bool {
	if m != nil {
		return m.IsBidirectionalCapable
	}
	return false
}

func (m *PimNbrBag) GetIsProxyCapable() bool {
	if m != nil {
		return m.IsProxyCapable
	}
	return false
}

func (m *PimNbrBag) GetIsBatchAssertsCapable() bool {
	if m != nil {
		return m.IsBatchAssertsCapable
	}
	return false
}

func (m *PimNbrBag) GetIsEcmpRedirectCapable() bool {
	if m != nil {
		return m.IsEcmpRedirectCapable
	}
	return false
}

func (m *PimNbrBag) GetIsBfdState() bool {
	if m != nil {
		return m.IsBfdState
	}
	return false
}

func (m *PimNbrBag) GetPropagationDelay() uint32 {
	if m != nil {
		return m.PropagationDelay
	}
	return 0
}

func (m *PimNbrBag) GetOverrideInterval() uint32 {
	if m != nil {
		return m.OverrideInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*PimNbrBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.neighbor_old_formats.neighbor_old_format.pim_nbr_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.neighbor_old_formats.neighbor_old_format.pim_addrtype")
	proto.RegisterType((*PimNbrBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.neighbor_old_formats.neighbor_old_format.pim_nbr_bag")
}

func init() { proto.RegisterFile("pim_nbr_bag.proto", fileDescriptor_0f976d211d835a3f) }

var fileDescriptor_0f976d211d835a3f = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x5f, 0x4f, 0x53, 0x31,
	0x18, 0xc6, 0x33, 0x31, 0x43, 0x3b, 0xc0, 0xad, 0x44, 0xec, 0x9d, 0x93, 0xc4, 0x64, 0x6a, 0xdc,
	0x05, 0x20, 0xe0, 0x7f, 0x41, 0xb8, 0x30, 0x26, 0x84, 0x0c, 0x4c, 0xf0, 0xc6, 0x37, 0xdd, 0x69,
	0x0f, 0x7b, 0x93, 0x73, 0xd6, 0xe6, 0x6d, 0x59, 0xb6, 0xef, 0xe2, 0xb7, 0xf0, 0x0b, 0x9a, 0xf6,
	0xfc, 0xf1, 0x8c, 0x70, 0xeb, 0xdd, 0xfa, 0x3c, 0xbf, 0xe7, 0x79, 0xfb, 0x76, 0xcb, 0x58, 0xcf,
	0x62, 0x0e, 0xd3, 0x31, 0xc1, 0x58, 0x5e, 0x0f, 0x2d, 0x19, 0x6f, 0xf8, 0xaf, 0x04, 0x5d, 0x62,
	0x00, 0x8d, 0x83, 0x39, 0x01, 0xda, 0xd9, 0x1e, 0x04, 0xc8, 0x58, 0x4d, 0x43, 0x8b, 0xf9, 0x50,
	0x26, 0x1e, 0x67, 0x7a, 0xa8, 0x74, 0x2a, 0x6f, 0x32, 0x0f, 0x89, 0x99, 0x7a, 0x3d, 0xf7, 0xc3,
	0xa9, 0xc6, 0xeb, 0xc9, 0xd8, 0x10, 0x98, 0x4c, 0x41, 0x6a, 0x28, 0x97, 0xde, 0xdd, 0x25, 0x6e,
	0x2b, 0xd6, 0x6d, 0x0c, 0x85, 0xef, 0xa7, 0x3f, 0x2f, 0xf8, 0x73, 0xb6, 0x81, 0x53, 0xaf, 0x29,
	0x95, 0x89, 0x86, 0xa9, 0xcc, 0xb5, 0x68, 0xf5, 0x5b, 0x83, 0x87, 0xa3, 0xf5, 0x5a, 0x3d, 0x93,
	0xb9, 0xe6, 0x2f, 0x58, 0xb7, 0x6e, 0x94, 0x4a, 0x91, 0x76, 0x4e, 0xdc, 0x8b, 0xe0, 0xa3, 0x4a,
	0x3f, 0x2a, 0xe4, 0xed, 0x9c, 0xad, 0x85, 0x29, 0x81, 0xf2, 0x0b, 0xab, 0xf9, 0x13, 0xb6, 0x2a,
	0xd3, 0x66, 0x75, 0x5b, 0xa6, 0xb1, 0xf3, 0x19, 0x5b, 0x8b, 0x3b, 0x2e, 0xf7, 0x75, 0x82, 0x56,
	0x76, 0x95, 0xc8, 0x7e, 0x8d, 0xac, 0xd4, 0xc8, 0x7e, 0x35, 0xee, 0x4f, 0x9b, 0x75, 0x1a, 0x5b,
	0xf1, 0x97, 0xac, 0xb7, 0xbc, 0x10, 0xcc, 0x49, 0xec, 0x14, 0x57, 0x5d, 0xda, 0xe9, 0x8a, 0xf8,
	0xef, 0x16, 0xdb, 0xbc, 0xbd, 0x56, 0xc0, 0x77, 0xfb, 0x2b, 0x83, 0xce, 0x4e, 0x36, 0xfc, 0xbf,
	0xdf, 0xc7, 0xb0, 0xf9, 0x4c, 0xa3, 0xde, 0xad, 0x77, 0xbc, 0x22, 0xbe, 0xc5, 0xda, 0x37, 0xd6,
	0x63, 0xae, 0xc5, 0x5e, 0xbf, 0x35, 0xb8, 0x3f, 0x2a, 0x4f, 0x5c, 0xb0, 0x55, 0x3d, 0xb7, 0x48,
	0xda, 0x89, 0x37, 0xd1, 0xa8, 0x8e, 0xe1, 0xbd, 0xe2, 0xc7, 0x05, 0x04, 0x90, 0xc4, 0x7e, 0xb4,
	0x3b, 0x85, 0x76, 0x19, 0x24, 0xfe, 0x9a, 0x6d, 0xa2, 0x03, 0x3f, 0x41, 0x07, 0xf5, 0x9d, 0x6e,
	0x9c, 0x38, 0xe8, 0xb7, 0x06, 0x0f, 0x46, 0x5d, 0x74, 0x97, 0x13, 0x74, 0x67, 0xa5, 0xf1, 0xc3,
	0xdd, 0x89, 0x2b, 0x12, 0x87, 0x77, 0xe1, 0x27, 0xc4, 0x77, 0xd9, 0x16, 0x3a, 0x50, 0x04, 0x96,
	0xd0, 0x10, 0xfa, 0x05, 0x24, 0xd2, 0xca, 0x71, 0xa6, 0xc5, 0xdb, 0x98, 0xd8, 0x44, 0x77, 0x42,
	0xe7, 0xa5, 0xf7, 0xb5, 0xb0, 0xf8, 0x53, 0xd6, 0x69, 0x24, 0xc4, 0xbb, 0x7e, 0x6b, 0xb0, 0x3e,
	0x62, 0xaa, 0xe6, 0xf8, 0x21, 0x13, 0xe8, 0x60, 0x8c, 0x0a, 0x49, 0x27, 0x1e, 0xcd, 0x54, 0x66,
	0x75, 0xef, 0xfb, 0xd8, 0xbb, 0x85, 0xee, 0xb8, 0x69, 0x57, 0xd5, 0x03, 0xd6, 0x45, 0x07, 0x96,
	0xcc, 0xfc, 0xdf, 0x4d, 0x3e, 0xc4, 0xc4, 0x06, 0xba, 0xf3, 0x20, 0x57, 0xe4, 0x41, 0x31, 0x43,
	0xfa, 0x64, 0x02, 0xd2, 0x39, 0x4d, 0xde, 0xd5, 0x89, 0x8f, 0x31, 0xf1, 0x18, 0xdd, 0x71, 0xb0,
	0x8f, 0x0a, 0x77, 0x39, 0xa8, 0x93, 0xdc, 0x02, 0xe9, 0xe2, 0x0a, 0x75, 0xf0, 0x53, 0x15, 0x3c,
	0x4d, 0x72, 0x3b, 0x2a, 0xdd, 0x2a, 0xd8, 0x67, 0x6b, 0x61, 0x62, 0xaa, 0xc0, 0x79, 0xe9, 0xb5,
	0xf8, 0x1c, 0x61, 0x86, 0xee, 0x38, 0x55, 0x17, 0x41, 0xe1, 0xaf, 0x58, 0xcf, 0x92, 0xb1, 0xf2,
	0x5a, 0x86, 0xad, 0x40, 0xe9, 0x4c, 0x2e, 0xc4, 0x97, 0xf8, 0x3c, 0xdd, 0x86, 0x71, 0x12, 0xf4,
	0x00, 0x9b, 0x99, 0x26, 0x42, 0xa5, 0x21, 0xfe, 0xd0, 0x67, 0x32, 0x13, 0x47, 0x05, 0x5c, 0x19,
	0xdf, 0x4a, 0x7d, 0xdc, 0x8e, 0xff, 0x38, 0xbb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xe1,
	0x9f, 0x5f, 0x86, 0x04, 0x00, 0x00,
}