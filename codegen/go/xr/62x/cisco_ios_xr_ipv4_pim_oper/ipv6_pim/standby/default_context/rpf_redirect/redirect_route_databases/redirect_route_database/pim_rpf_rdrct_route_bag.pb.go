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
// source: pim_rpf_rdrct_route_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_standby_default_context_rpf_redirect_redirect_route_databases_redirect_route_database

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

type PimRpfRdrctRouteBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         string   `protobuf:"bytes,2,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfRdrctRouteBag_KEYS) Reset()         { *m = PimRpfRdrctRouteBag_KEYS{} }
func (m *PimRpfRdrctRouteBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfRdrctRouteBag_KEYS) ProtoMessage()    {}
func (*PimRpfRdrctRouteBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b24f874596902a3, []int{0}
}

func (m *PimRpfRdrctRouteBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfRdrctRouteBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfRdrctRouteBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfRdrctRouteBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfRdrctRouteBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfRdrctRouteBag_KEYS.Merge(m, src)
}
func (m *PimRpfRdrctRouteBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfRdrctRouteBag_KEYS.Size(m)
}
func (m *PimRpfRdrctRouteBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfRdrctRouteBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfRdrctRouteBag_KEYS proto.InternalMessageInfo

func (m *PimRpfRdrctRouteBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimRpfRdrctRouteBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
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
	return fileDescriptor_6b24f874596902a3, []int{1}
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

type PimRpfRdrctRintfBag struct {
	InterfaceName        string       `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	RpfAddress           *PimAddrtype `protobuf:"bytes,2,opt,name=rpf_address,json=rpfAddress,proto3" json:"rpf_address,omitempty"`
	Uptime               uint64       `protobuf:"varint,3,opt,name=uptime,proto3" json:"uptime,omitempty"`
	Expiry               uint64       `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	IsRpfInterface       bool         `protobuf:"varint,5,opt,name=is_rpf_interface,json=isRpfInterface,proto3" json:"is_rpf_interface,omitempty"`
	IsOutgoingInterface  bool         `protobuf:"varint,6,opt,name=is_outgoing_interface,json=isOutgoingInterface,proto3" json:"is_outgoing_interface,omitempty"`
	IsSnoopInterface     bool         `protobuf:"varint,7,opt,name=is_snoop_interface,json=isSnoopInterface,proto3" json:"is_snoop_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimRpfRdrctRintfBag) Reset()         { *m = PimRpfRdrctRintfBag{} }
func (m *PimRpfRdrctRintfBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfRdrctRintfBag) ProtoMessage()    {}
func (*PimRpfRdrctRintfBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b24f874596902a3, []int{2}
}

func (m *PimRpfRdrctRintfBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfRdrctRintfBag.Unmarshal(m, b)
}
func (m *PimRpfRdrctRintfBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfRdrctRintfBag.Marshal(b, m, deterministic)
}
func (m *PimRpfRdrctRintfBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfRdrctRintfBag.Merge(m, src)
}
func (m *PimRpfRdrctRintfBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfRdrctRintfBag.Size(m)
}
func (m *PimRpfRdrctRintfBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfRdrctRintfBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfRdrctRintfBag proto.InternalMessageInfo

func (m *PimRpfRdrctRintfBag) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *PimRpfRdrctRintfBag) GetRpfAddress() *PimAddrtype {
	if m != nil {
		return m.RpfAddress
	}
	return nil
}

func (m *PimRpfRdrctRintfBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func (m *PimRpfRdrctRintfBag) GetExpiry() uint64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *PimRpfRdrctRintfBag) GetIsRpfInterface() bool {
	if m != nil {
		return m.IsRpfInterface
	}
	return false
}

func (m *PimRpfRdrctRintfBag) GetIsOutgoingInterface() bool {
	if m != nil {
		return m.IsOutgoingInterface
	}
	return false
}

func (m *PimRpfRdrctRintfBag) GetIsSnoopInterface() bool {
	if m != nil {
		return m.IsSnoopInterface
	}
	return false
}

type PimRpfRdrctRouteBag struct {
	GroupAddressXr       *PimAddrtype           `protobuf:"bytes,50,opt,name=group_address_xr,json=groupAddressXr,proto3" json:"group_address_xr,omitempty"`
	SourceAddressXr      *PimAddrtype           `protobuf:"bytes,51,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	Interface            []*PimRpfRdrctRintfBag `protobuf:"bytes,52,rep,name=interface,proto3" json:"interface,omitempty"`
	Bandwidth            uint32                 `protobuf:"varint,53,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	Uptime               uint64                 `protobuf:"varint,54,opt,name=uptime,proto3" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *PimRpfRdrctRouteBag) Reset()         { *m = PimRpfRdrctRouteBag{} }
func (m *PimRpfRdrctRouteBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfRdrctRouteBag) ProtoMessage()    {}
func (*PimRpfRdrctRouteBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b24f874596902a3, []int{3}
}

func (m *PimRpfRdrctRouteBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfRdrctRouteBag.Unmarshal(m, b)
}
func (m *PimRpfRdrctRouteBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfRdrctRouteBag.Marshal(b, m, deterministic)
}
func (m *PimRpfRdrctRouteBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfRdrctRouteBag.Merge(m, src)
}
func (m *PimRpfRdrctRouteBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfRdrctRouteBag.Size(m)
}
func (m *PimRpfRdrctRouteBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfRdrctRouteBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfRdrctRouteBag proto.InternalMessageInfo

func (m *PimRpfRdrctRouteBag) GetGroupAddressXr() *PimAddrtype {
	if m != nil {
		return m.GroupAddressXr
	}
	return nil
}

func (m *PimRpfRdrctRouteBag) GetSourceAddressXr() *PimAddrtype {
	if m != nil {
		return m.SourceAddressXr
	}
	return nil
}

func (m *PimRpfRdrctRouteBag) GetInterface() []*PimRpfRdrctRintfBag {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *PimRpfRdrctRouteBag) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *PimRpfRdrctRouteBag) GetUptime() uint64 {
	if m != nil {
		return m.Uptime
	}
	return 0
}

func init() {
	proto.RegisterType((*PimRpfRdrctRouteBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.redirect_route_databases.redirect_route_database.pim_rpf_rdrct_route_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.redirect_route_databases.redirect_route_database.pim_addrtype")
	proto.RegisterType((*PimRpfRdrctRintfBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.redirect_route_databases.redirect_route_database.pim_rpf_rdrct_rintf_bag")
	proto.RegisterType((*PimRpfRdrctRouteBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.standby.default_context.rpf_redirect.redirect_route_databases.redirect_route_database.pim_rpf_rdrct_route_bag")
}

func init() { proto.RegisterFile("pim_rpf_rdrct_route_bag.proto", fileDescriptor_6b24f874596902a3) }

var fileDescriptor_6b24f874596902a3 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x68, 0xe9, 0xd2, 0xe9, 0x0f, 0xc5, 0x08, 0xb6, 0x87, 0x45, 0x2a, 0x45, 0x48, 0x39,
	0xa0, 0x1c, 0xba, 0x4b, 0xef, 0x1c, 0x38, 0x20, 0x24, 0x90, 0xb2, 0x97, 0xe5, 0x64, 0x39, 0xb1,
	0x53, 0x06, 0x91, 0xd8, 0xb2, 0x9d, 0xa5, 0x7d, 0x02, 0xee, 0x3c, 0x05, 0x1c, 0x41, 0xbc, 0x17,
	0xaf, 0x80, 0xec, 0xfc, 0x34, 0x85, 0xdd, 0xf3, 0xee, 0x2d, 0xfe, 0xe6, 0x9b, 0x99, 0x6f, 0xfe,
	0x02, 0x4f, 0x14, 0xe6, 0x54, 0xab, 0x8c, 0x6a, 0xae, 0x53, 0x4b, 0xb5, 0x2c, 0xad, 0xa0, 0x09,
	0xdb, 0x44, 0x4a, 0x4b, 0x2b, 0xc9, 0x2e, 0x45, 0x93, 0x4a, 0x8a, 0xd2, 0xd0, 0xad, 0xa6, 0xa8,
	0x2e, 0xcf, 0xa8, 0x73, 0x90, 0x4a, 0xe8, 0x08, 0xd5, 0xe5, 0xda, 0xbd, 0x22, 0x63, 0x59, 0xc1,
	0x93, 0x5d, 0xc4, 0x45, 0xc6, 0xca, 0xcf, 0x96, 0xa6, 0xb2, 0xb0, 0x62, 0x6b, 0x23, 0x1f, 0x56,
	0x70, 0xd4, 0x22, 0xb5, 0x51, 0xf3, 0x51, 0xa7, 0xe0, 0xcc, 0xb2, 0x84, 0x19, 0x61, 0xae, 0x33,
	0x2c, 0x3f, 0xc1, 0xc9, 0x35, 0xda, 0xe8, 0xdb, 0xd7, 0x1f, 0xce, 0xc9, 0x73, 0x98, 0x1a, 0x59,
	0xea, 0x54, 0x50, 0xc6, 0xb9, 0x16, 0xc6, 0xcc, 0x83, 0x45, 0x10, 0x0e, 0xe3, 0x49, 0x85, 0xbe,
	0xaa, 0x40, 0xf2, 0x0c, 0x26, 0x1b, 0x2d, 0x4b, 0xd5, 0xb2, 0xee, 0x78, 0xd6, 0xd8, 0x83, 0x35,
	0x69, 0x99, 0xc3, 0xd8, 0xe5, 0x72, 0x14, 0xbb, 0x53, 0x82, 0x1c, 0xc3, 0x11, 0xcb, 0x68, 0xc1,
	0x72, 0x51, 0x07, 0x1d, 0xb0, 0xec, 0x1d, 0xcb, 0x05, 0x79, 0x0a, 0x63, 0xdf, 0x84, 0xc3, 0x60,
	0x23, 0x87, 0x35, 0x09, 0x2b, 0xca, 0xba, 0xa5, 0xf4, 0x5a, 0xca, 0xba, 0x49, 0xf7, 0xbd, 0x07,
	0xc7, 0xff, 0xd4, 0x86, 0x85, 0xcd, 0x5c, 0x6d, 0xae, 0x2c, 0x2c, 0xac, 0xd0, 0x19, 0x4b, 0x45,
	0x57, 0xc1, 0xa4, 0x45, 0xbd, 0x90, 0x1f, 0x01, 0x8c, 0x9c, 0x7b, 0x57, 0xc8, 0x68, 0xf5, 0x35,
	0x88, 0x6e, 0x6c, 0x60, 0x51, 0xb7, 0x83, 0x31, 0x68, 0x95, 0x35, 0x1d, 0x79, 0x0c, 0x83, 0x52,
	0x59, 0xcc, 0x85, 0xef, 0x45, 0x3f, 0xae, 0x5f, 0x0e, 0x17, 0x5b, 0x85, 0x7a, 0x37, 0xef, 0x57,
	0x78, 0xf5, 0x22, 0x21, 0xcc, 0xd0, 0xf8, 0xe6, 0xb4, 0x35, 0xcf, 0xef, 0x2e, 0x82, 0xf0, 0x5e,
	0x3c, 0x45, 0x13, 0xab, 0xec, 0x4d, 0x83, 0x92, 0x15, 0x3c, 0x42, 0x43, 0x65, 0x69, 0x37, 0x12,
	0x8b, 0x4d, 0x87, 0x3e, 0xf0, 0xf4, 0x87, 0x68, 0xde, 0xd7, 0xb6, 0xbd, 0xcf, 0x0b, 0x20, 0x68,
	0xa8, 0x29, 0xa4, 0x54, 0x1d, 0x87, 0x23, 0xef, 0x30, 0x43, 0x73, 0xee, 0x0c, 0x2d, 0x7b, 0xf9,
	0xa7, 0xff, 0xdf, 0xa8, 0x9a, 0x35, 0x24, 0xbf, 0x02, 0x98, 0x1d, 0xec, 0x16, 0xdd, 0xea, 0xf9,
	0xea, 0x96, 0x0d, 0x62, 0xda, 0x5d, 0xf4, 0x0b, 0x4d, 0x7e, 0x07, 0xf0, 0xe0, 0xf0, 0x6e, 0x9c,
	0xea, 0xd3, 0x5b, 0xa6, 0xfa, 0xfe, 0xc1, 0x11, 0x5f, 0x68, 0xf2, 0x33, 0x80, 0xe1, 0x7e, 0x5a,
	0x67, 0x8b, 0x5e, 0x38, 0x5a, 0x7d, 0xbb, 0x69, 0xb9, 0x57, 0xdc, 0x6f, 0xbc, 0x57, 0x49, 0x4e,
	0x60, 0x98, 0xb0, 0x82, 0x7f, 0x41, 0x6e, 0x3f, 0xce, 0x5f, 0x2e, 0x82, 0x70, 0x12, 0xef, 0x81,
	0xce, 0x55, 0xac, 0xbb, 0x57, 0x91, 0x0c, 0xfc, 0x9f, 0xf7, 0xf4, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0f, 0xce, 0x61, 0xe4, 0x9a, 0x05, 0x00, 0x00,
}