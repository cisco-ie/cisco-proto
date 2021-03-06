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
// source: ptp_if_interop_info.proto

package cisco_ios_xr_ptp_oper_ptp_interface_interops_interface_interop

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

type PtpIfInteropInfo_KEYS struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpIfInteropInfo_KEYS) Reset()         { *m = PtpIfInteropInfo_KEYS{} }
func (m *PtpIfInteropInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpIfInteropInfo_KEYS) ProtoMessage()    {}
func (*PtpIfInteropInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{0}
}

func (m *PtpIfInteropInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfInteropInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpIfInteropInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfInteropInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpIfInteropInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfInteropInfo_KEYS.Merge(m, src)
}
func (m *PtpIfInteropInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpIfInteropInfo_KEYS.Size(m)
}
func (m *PtpIfInteropInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfInteropInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfInteropInfo_KEYS proto.InternalMessageInfo

func (m *PtpIfInteropInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type PtpBagMacAddrType struct {
	Macaddr              string   `protobuf:"bytes,1,opt,name=macaddr,proto3" json:"macaddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagMacAddrType) Reset()         { *m = PtpBagMacAddrType{} }
func (m *PtpBagMacAddrType) String() string { return proto.CompactTextString(m) }
func (*PtpBagMacAddrType) ProtoMessage()    {}
func (*PtpBagMacAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{1}
}

func (m *PtpBagMacAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagMacAddrType.Unmarshal(m, b)
}
func (m *PtpBagMacAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagMacAddrType.Marshal(b, m, deterministic)
}
func (m *PtpBagMacAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagMacAddrType.Merge(m, src)
}
func (m *PtpBagMacAddrType) XXX_Size() int {
	return xxx_messageInfo_PtpBagMacAddrType.Size(m)
}
func (m *PtpBagMacAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagMacAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagMacAddrType proto.InternalMessageInfo

func (m *PtpBagMacAddrType) GetMacaddr() string {
	if m != nil {
		return m.Macaddr
	}
	return ""
}

type PtpBagIpv6AddrType struct {
	Ipv6Address          string   `protobuf:"bytes,1,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpBagIpv6AddrType) Reset()         { *m = PtpBagIpv6AddrType{} }
func (m *PtpBagIpv6AddrType) String() string { return proto.CompactTextString(m) }
func (*PtpBagIpv6AddrType) ProtoMessage()    {}
func (*PtpBagIpv6AddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{2}
}

func (m *PtpBagIpv6AddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagIpv6AddrType.Unmarshal(m, b)
}
func (m *PtpBagIpv6AddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagIpv6AddrType.Marshal(b, m, deterministic)
}
func (m *PtpBagIpv6AddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagIpv6AddrType.Merge(m, src)
}
func (m *PtpBagIpv6AddrType) XXX_Size() int {
	return xxx_messageInfo_PtpBagIpv6AddrType.Size(m)
}
func (m *PtpBagIpv6AddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagIpv6AddrType.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagIpv6AddrType proto.InternalMessageInfo

func (m *PtpBagIpv6AddrType) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type PtpBagAddress struct {
	Encapsulation        string              `protobuf:"bytes,1,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	AddressUnknown       bool                `protobuf:"varint,2,opt,name=address_unknown,json=addressUnknown,proto3" json:"address_unknown,omitempty"`
	MacAddress           *PtpBagMacAddrType  `protobuf:"bytes,3,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Ipv4Address          string              `protobuf:"bytes,4,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          *PtpBagIpv6AddrType `protobuf:"bytes,5,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PtpBagAddress) Reset()         { *m = PtpBagAddress{} }
func (m *PtpBagAddress) String() string { return proto.CompactTextString(m) }
func (*PtpBagAddress) ProtoMessage()    {}
func (*PtpBagAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{3}
}

func (m *PtpBagAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagAddress.Unmarshal(m, b)
}
func (m *PtpBagAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagAddress.Marshal(b, m, deterministic)
}
func (m *PtpBagAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagAddress.Merge(m, src)
}
func (m *PtpBagAddress) XXX_Size() int {
	return xxx_messageInfo_PtpBagAddress.Size(m)
}
func (m *PtpBagAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagAddress.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagAddress proto.InternalMessageInfo

func (m *PtpBagAddress) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *PtpBagAddress) GetAddressUnknown() bool {
	if m != nil {
		return m.AddressUnknown
	}
	return false
}

func (m *PtpBagAddress) GetMacAddress() *PtpBagMacAddrType {
	if m != nil {
		return m.MacAddress
	}
	return nil
}

func (m *PtpBagAddress) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *PtpBagAddress) GetIpv6Address() *PtpBagIpv6AddrType {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

type PtpBagProfileInterop struct {
	FromPriority1         uint32   `protobuf:"varint,1,opt,name=from_priority1,json=fromPriority1,proto3" json:"from_priority1,omitempty"`
	ToPriority1           uint32   `protobuf:"varint,2,opt,name=to_priority1,json=toPriority1,proto3" json:"to_priority1,omitempty"`
	FromPriority2         uint32   `protobuf:"varint,3,opt,name=from_priority2,json=fromPriority2,proto3" json:"from_priority2,omitempty"`
	ToPriority2           uint32   `protobuf:"varint,4,opt,name=to_priority2,json=toPriority2,proto3" json:"to_priority2,omitempty"`
	FromAccuracy          uint32   `protobuf:"varint,5,opt,name=from_accuracy,json=fromAccuracy,proto3" json:"from_accuracy,omitempty"`
	ToAccuracy            uint32   `protobuf:"varint,6,opt,name=to_accuracy,json=toAccuracy,proto3" json:"to_accuracy,omitempty"`
	FromClockClass        uint32   `protobuf:"varint,7,opt,name=from_clock_class,json=fromClockClass,proto3" json:"from_clock_class,omitempty"`
	ToClockClass          uint32   `protobuf:"varint,8,opt,name=to_clock_class,json=toClockClass,proto3" json:"to_clock_class,omitempty"`
	FromOffsetLogVariance uint32   `protobuf:"varint,9,opt,name=from_offset_log_variance,json=fromOffsetLogVariance,proto3" json:"from_offset_log_variance,omitempty"`
	ToOffsetLogVariance   uint32   `protobuf:"varint,10,opt,name=to_offset_log_variance,json=toOffsetLogVariance,proto3" json:"to_offset_log_variance,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *PtpBagProfileInterop) Reset()         { *m = PtpBagProfileInterop{} }
func (m *PtpBagProfileInterop) String() string { return proto.CompactTextString(m) }
func (*PtpBagProfileInterop) ProtoMessage()    {}
func (*PtpBagProfileInterop) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{4}
}

func (m *PtpBagProfileInterop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagProfileInterop.Unmarshal(m, b)
}
func (m *PtpBagProfileInterop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagProfileInterop.Marshal(b, m, deterministic)
}
func (m *PtpBagProfileInterop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagProfileInterop.Merge(m, src)
}
func (m *PtpBagProfileInterop) XXX_Size() int {
	return xxx_messageInfo_PtpBagProfileInterop.Size(m)
}
func (m *PtpBagProfileInterop) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagProfileInterop.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagProfileInterop proto.InternalMessageInfo

func (m *PtpBagProfileInterop) GetFromPriority1() uint32 {
	if m != nil {
		return m.FromPriority1
	}
	return 0
}

func (m *PtpBagProfileInterop) GetToPriority1() uint32 {
	if m != nil {
		return m.ToPriority1
	}
	return 0
}

func (m *PtpBagProfileInterop) GetFromPriority2() uint32 {
	if m != nil {
		return m.FromPriority2
	}
	return 0
}

func (m *PtpBagProfileInterop) GetToPriority2() uint32 {
	if m != nil {
		return m.ToPriority2
	}
	return 0
}

func (m *PtpBagProfileInterop) GetFromAccuracy() uint32 {
	if m != nil {
		return m.FromAccuracy
	}
	return 0
}

func (m *PtpBagProfileInterop) GetToAccuracy() uint32 {
	if m != nil {
		return m.ToAccuracy
	}
	return 0
}

func (m *PtpBagProfileInterop) GetFromClockClass() uint32 {
	if m != nil {
		return m.FromClockClass
	}
	return 0
}

func (m *PtpBagProfileInterop) GetToClockClass() uint32 {
	if m != nil {
		return m.ToClockClass
	}
	return 0
}

func (m *PtpBagProfileInterop) GetFromOffsetLogVariance() uint32 {
	if m != nil {
		return m.FromOffsetLogVariance
	}
	return 0
}

func (m *PtpBagProfileInterop) GetToOffsetLogVariance() uint32 {
	if m != nil {
		return m.ToOffsetLogVariance
	}
	return 0
}

type PtpBagProfileInteropPeer struct {
	Address              *PtpBagAddress        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Interop              *PtpBagProfileInterop `protobuf:"bytes,2,opt,name=interop,proto3" json:"interop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PtpBagProfileInteropPeer) Reset()         { *m = PtpBagProfileInteropPeer{} }
func (m *PtpBagProfileInteropPeer) String() string { return proto.CompactTextString(m) }
func (*PtpBagProfileInteropPeer) ProtoMessage()    {}
func (*PtpBagProfileInteropPeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{5}
}

func (m *PtpBagProfileInteropPeer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpBagProfileInteropPeer.Unmarshal(m, b)
}
func (m *PtpBagProfileInteropPeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpBagProfileInteropPeer.Marshal(b, m, deterministic)
}
func (m *PtpBagProfileInteropPeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpBagProfileInteropPeer.Merge(m, src)
}
func (m *PtpBagProfileInteropPeer) XXX_Size() int {
	return xxx_messageInfo_PtpBagProfileInteropPeer.Size(m)
}
func (m *PtpBagProfileInteropPeer) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpBagProfileInteropPeer.DiscardUnknown(m)
}

var xxx_messageInfo_PtpBagProfileInteropPeer proto.InternalMessageInfo

func (m *PtpBagProfileInteropPeer) GetAddress() *PtpBagAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *PtpBagProfileInteropPeer) GetInterop() *PtpBagProfileInterop {
	if m != nil {
		return m.Interop
	}
	return nil
}

type PtpIfInteropInfo struct {
	LocalDomain          uint32                      `protobuf:"varint,50,opt,name=local_domain,json=localDomain,proto3" json:"local_domain,omitempty"`
	InteropDomain        uint32                      `protobuf:"varint,51,opt,name=interop_domain,json=interopDomain,proto3" json:"interop_domain,omitempty"`
	LocalProfile         string                      `protobuf:"bytes,52,opt,name=local_profile,json=localProfile,proto3" json:"local_profile,omitempty"`
	InteropProfile       string                      `protobuf:"bytes,53,opt,name=interop_profile,json=interopProfile,proto3" json:"interop_profile,omitempty"`
	IngressInterop       []*PtpBagProfileInteropPeer `protobuf:"bytes,54,rep,name=ingress_interop,json=ingressInterop,proto3" json:"ingress_interop,omitempty"`
	EgressInterop        *PtpBagProfileInterop       `protobuf:"bytes,55,opt,name=egress_interop,json=egressInterop,proto3" json:"egress_interop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *PtpIfInteropInfo) Reset()         { *m = PtpIfInteropInfo{} }
func (m *PtpIfInteropInfo) String() string { return proto.CompactTextString(m) }
func (*PtpIfInteropInfo) ProtoMessage()    {}
func (*PtpIfInteropInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_92348307e2b17687, []int{6}
}

func (m *PtpIfInteropInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpIfInteropInfo.Unmarshal(m, b)
}
func (m *PtpIfInteropInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpIfInteropInfo.Marshal(b, m, deterministic)
}
func (m *PtpIfInteropInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpIfInteropInfo.Merge(m, src)
}
func (m *PtpIfInteropInfo) XXX_Size() int {
	return xxx_messageInfo_PtpIfInteropInfo.Size(m)
}
func (m *PtpIfInteropInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpIfInteropInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpIfInteropInfo proto.InternalMessageInfo

func (m *PtpIfInteropInfo) GetLocalDomain() uint32 {
	if m != nil {
		return m.LocalDomain
	}
	return 0
}

func (m *PtpIfInteropInfo) GetInteropDomain() uint32 {
	if m != nil {
		return m.InteropDomain
	}
	return 0
}

func (m *PtpIfInteropInfo) GetLocalProfile() string {
	if m != nil {
		return m.LocalProfile
	}
	return ""
}

func (m *PtpIfInteropInfo) GetInteropProfile() string {
	if m != nil {
		return m.InteropProfile
	}
	return ""
}

func (m *PtpIfInteropInfo) GetIngressInterop() []*PtpBagProfileInteropPeer {
	if m != nil {
		return m.IngressInterop
	}
	return nil
}

func (m *PtpIfInteropInfo) GetEgressInterop() *PtpBagProfileInterop {
	if m != nil {
		return m.EgressInterop
	}
	return nil
}

func init() {
	proto.RegisterType((*PtpIfInteropInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_if_interop_info_KEYS")
	proto.RegisterType((*PtpBagMacAddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_bag_mac_addr_type")
	proto.RegisterType((*PtpBagIpv6AddrType)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_bag_ipv6_addr_type")
	proto.RegisterType((*PtpBagAddress)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_bag_address")
	proto.RegisterType((*PtpBagProfileInterop)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_bag_profile_interop")
	proto.RegisterType((*PtpBagProfileInteropPeer)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_bag_profile_interop_peer")
	proto.RegisterType((*PtpIfInteropInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.interface_interops.interface_interop.ptp_if_interop_info")
}

func init() { proto.RegisterFile("ptp_if_interop_info.proto", fileDescriptor_92348307e2b17687) }

var fileDescriptor_92348307e2b17687 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4f, 0x4f, 0xd4, 0x4e,
	0x18, 0xce, 0xb2, 0xbf, 0x1f, 0x0b, 0x6f, 0xe9, 0x62, 0x4a, 0xc0, 0x9a, 0x98, 0xb8, 0x2e, 0x18,
	0xf7, 0xb4, 0x09, 0x05, 0xe1, 0x60, 0x62, 0x42, 0xd0, 0x83, 0xd1, 0x08, 0xa9, 0x01, 0x63, 0x62,
	0x32, 0x19, 0x86, 0xe9, 0xa6, 0xa1, 0xed, 0x3b, 0xb6, 0x03, 0xba, 0x17, 0x6f, 0x7e, 0x20, 0x0f,
	0x7e, 0x35, 0x4f, 0x1e, 0xcc, 0xbc, 0x9d, 0x29, 0xbb, 0xcb, 0x7a, 0x12, 0x8e, 0x7d, 0xe6, 0x79,
	0x9e, 0xf7, 0xef, 0x4c, 0xe1, 0x81, 0xd2, 0x8a, 0xa5, 0x09, 0x4b, 0x0b, 0x2d, 0x4b, 0x54, 0x2c,
	0x2d, 0x12, 0x1c, 0xaa, 0x12, 0x35, 0x06, 0x2f, 0x44, 0x5a, 0x09, 0x64, 0x29, 0x56, 0xec, 0x6b,
	0xc9, 0x0c, 0x0f, 0x95, 0x2c, 0x87, 0x4a, 0xab, 0x21, 0xb1, 0x13, 0x2e, 0xa4, 0xd3, 0x55, 0x37,
	0xa1, 0xfe, 0x01, 0x84, 0x73, 0xcc, 0xd9, 0x9b, 0x57, 0x1f, 0xdf, 0x07, 0x4f, 0xa0, 0x7b, 0x2d,
	0x28, 0x78, 0x2e, 0xc3, 0x56, 0xaf, 0x35, 0x58, 0x8e, 0xfd, 0x06, 0x7d, 0xc7, 0x73, 0xd9, 0xdf,
	0x86, 0x75, 0x63, 0x71, 0xc6, 0x47, 0x2c, 0xe7, 0x82, 0xf1, 0xf3, 0xf3, 0x92, 0xe9, 0xb1, 0x92,
	0x41, 0x08, 0x9d, 0x9c, 0x0b, 0xf3, 0x6d, 0x85, 0xee, 0xb3, 0xff, 0x1c, 0x36, 0x9c, 0x24, 0x55,
	0x57, 0x7b, 0x13, 0x9a, 0xc7, 0xb0, 0xd2, 0x20, 0xb2, 0xaa, 0xac, 0xd0, 0x33, 0xd8, 0x41, 0x0d,
	0xf5, 0x7f, 0x2d, 0xc0, 0xaa, 0x53, 0x5b, 0x5a, 0xb0, 0x05, 0xbe, 0x2c, 0x04, 0x57, 0xd5, 0x65,
	0xc6, 0x75, 0x8a, 0x85, 0xcb, 0x74, 0x0a, 0x0c, 0x9e, 0xc2, 0xaa, 0x15, 0xb0, 0xcb, 0xe2, 0xa2,
	0xc0, 0x2f, 0x45, 0xb8, 0xd0, 0x6b, 0x0d, 0x96, 0xe2, 0xae, 0x85, 0x4f, 0x6a, 0x34, 0xb8, 0x02,
	0xcf, 0x95, 0x62, 0x92, 0x68, 0xf7, 0x5a, 0x03, 0x2f, 0x3a, 0x19, 0xfe, 0x5b, 0xaf, 0x87, 0x73,
	0xbb, 0x14, 0x43, 0xce, 0x85, 0x2d, 0xcd, 0x56, 0xbf, 0xdb, 0x04, 0xfe, 0xaf, 0xa9, 0x7e, 0xd7,
	0x51, 0xc6, 0x33, 0x0d, 0xfa, 0x9f, 0x72, 0x3b, 0xbd, 0xad, 0xdc, 0xa6, 0xc7, 0x31, 0xdd, 0xf8,
	0x1f, 0x6d, 0xb8, 0xef, 0x78, 0xaa, 0xc4, 0x24, 0xcd, 0x1a, 0xbd, 0xd9, 0x95, 0xa4, 0xc4, 0x9c,
	0xa9, 0x32, 0xc5, 0x32, 0xd5, 0xe3, 0x6d, 0x9a, 0x80, 0x1f, 0xfb, 0x06, 0x3d, 0x76, 0xa0, 0x29,
	0x50, 0xe3, 0x04, 0x69, 0x81, 0x48, 0x9e, 0xc6, 0x6b, 0xca, 0xac, 0x53, 0x44, 0xed, 0x9f, 0x71,
	0x8a, 0x66, 0x9c, 0x22, 0x6a, 0xd5, 0x94, 0x53, 0x14, 0x6c, 0x02, 0x69, 0x18, 0x17, 0xe2, 0xb2,
	0xe4, 0x62, 0x4c, 0xbd, 0xf2, 0xe3, 0x15, 0x03, 0x1e, 0x58, 0x2c, 0x78, 0x04, 0x9e, 0xc6, 0x6b,
	0xca, 0x22, 0x51, 0x40, 0x63, 0x43, 0x18, 0xc0, 0x3d, 0x72, 0x11, 0x19, 0x8a, 0x0b, 0x26, 0x32,
	0x5e, 0x55, 0x61, 0x87, 0x58, 0x94, 0xe7, 0xa1, 0x81, 0x0f, 0x0d, 0x1a, 0x6c, 0x41, 0x57, 0xe3,
	0x14, 0x6f, 0xa9, 0x0e, 0xa8, 0x71, 0x82, 0xb5, 0x0f, 0x21, 0xf9, 0x61, 0x92, 0x54, 0x52, 0xb3,
	0x0c, 0x47, 0xec, 0x8a, 0x97, 0x29, 0x2f, 0x84, 0x0c, 0x97, 0x89, 0xbf, 0x6e, 0xce, 0x8f, 0xe8,
	0xf8, 0x2d, 0x8e, 0x4e, 0xed, 0x61, 0xb0, 0x03, 0x1b, 0x1a, 0xe7, 0xca, 0x80, 0x64, 0x6b, 0x1a,
	0x6f, 0x88, 0xfa, 0xbf, 0x5b, 0xf0, 0xf0, 0x2f, 0x33, 0x63, 0x4a, 0xca, 0x32, 0x48, 0xa1, 0x33,
	0x79, 0xd7, 0xbc, 0xe8, 0xe8, 0xb6, 0x56, 0xc9, 0xda, 0xc6, 0xce, 0x3f, 0xf8, 0x0c, 0x1d, 0xcb,
	0xa1, 0xb9, 0x7b, 0xd1, 0x87, 0xdb, 0x0a, 0x35, 0x53, 0x59, 0xec, 0xe2, 0xf4, 0x7f, 0xb6, 0x61,
	0x6d, 0xce, 0xfb, 0x66, 0xb6, 0x27, 0x43, 0xc1, 0x33, 0x76, 0x8e, 0x39, 0x4f, 0x8b, 0x30, 0xaa,
	0xb7, 0x87, 0xb0, 0x97, 0x04, 0x35, 0xaf, 0x1f, 0x2a, 0x47, 0xda, 0xa9, 0xf7, 0xd0, 0xa2, 0x96,
	0xb6, 0x09, 0x7e, 0xed, 0x64, 0x73, 0x08, 0x77, 0xe9, 0xce, 0xd6, 0xf6, 0xc7, 0x35, 0x66, 0x1e,
	0x9e, 0xa6, 0xe9, 0x96, 0xf6, 0x8c, 0x68, 0x2e, 0x84, 0x23, 0x7e, 0x6f, 0x19, 0xe6, 0x88, 0x9e,
	0x28, 0xd7, 0xab, 0xbd, 0x5e, 0x7b, 0xe0, 0x45, 0x9f, 0xee, 0xa8, 0x57, 0xb4, 0x05, 0x26, 0x0f,
	0x0a, 0xfa, 0xda, 0x5e, 0xe7, 0x6f, 0xd0, 0x95, 0xd3, 0x59, 0xec, 0xdf, 0xed, 0xc4, 0x7c, 0x39,
	0x19, 0xff, 0x6c, 0x91, 0xfe, 0x6e, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x49, 0x51,
	0xbf, 0xfa, 0x06, 0x00, 0x00,
}
