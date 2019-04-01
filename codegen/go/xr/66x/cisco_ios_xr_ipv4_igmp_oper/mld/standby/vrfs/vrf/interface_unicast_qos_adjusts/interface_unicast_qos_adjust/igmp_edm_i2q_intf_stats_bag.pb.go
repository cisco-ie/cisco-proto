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
// source: igmp_edm_i2q_intf_stats_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_standby_vrfs_vrf_interface_unicast_qos_adjusts_interface_unicast_qos_adjust

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

type IgmpEdmI2QIntfStatsBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmI2QIntfStatsBag_KEYS) Reset()         { *m = IgmpEdmI2QIntfStatsBag_KEYS{} }
func (m *IgmpEdmI2QIntfStatsBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmI2QIntfStatsBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmI2QIntfStatsBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4d11de3db8bf66c, []int{0}
}

func (m *IgmpEdmI2QIntfStatsBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmI2QIntfStatsBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmI2QIntfStatsBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmI2QIntfStatsBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS.Size(m)
}
func (m *IgmpEdmI2QIntfStatsBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmI2QIntfStatsBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmI2QIntfStatsBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *IgmpEdmI2QIntfStatsBag_KEYS) GetInterfaceName() string {
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
	return fileDescriptor_b4d11de3db8bf66c, []int{1}
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

type IgmpEdmI2QIntfRateBag struct {
	IsAdd                bool          `protobuf:"varint,1,opt,name=is_add,json=isAdd,proto3" json:"is_add,omitempty"`
	SourceAddress        *IgmpAddrtype `protobuf:"bytes,2,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         *IgmpAddrtype `protobuf:"bytes,3,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	Weight               uint32        `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	ReceivedTime         uint64        `protobuf:"varint,5,opt,name=received_time,json=receivedTime,proto3" json:"received_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *IgmpEdmI2QIntfRateBag) Reset()         { *m = IgmpEdmI2QIntfRateBag{} }
func (m *IgmpEdmI2QIntfRateBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmI2QIntfRateBag) ProtoMessage()    {}
func (*IgmpEdmI2QIntfRateBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4d11de3db8bf66c, []int{2}
}

func (m *IgmpEdmI2QIntfRateBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmI2QIntfRateBag.Unmarshal(m, b)
}
func (m *IgmpEdmI2QIntfRateBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmI2QIntfRateBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmI2QIntfRateBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmI2QIntfRateBag.Merge(m, src)
}
func (m *IgmpEdmI2QIntfRateBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmI2QIntfRateBag.Size(m)
}
func (m *IgmpEdmI2QIntfRateBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmI2QIntfRateBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmI2QIntfRateBag proto.InternalMessageInfo

func (m *IgmpEdmI2QIntfRateBag) GetIsAdd() bool {
	if m != nil {
		return m.IsAdd
	}
	return false
}

func (m *IgmpEdmI2QIntfRateBag) GetSourceAddress() *IgmpAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func (m *IgmpEdmI2QIntfRateBag) GetGroupAddress() *IgmpAddrtype {
	if m != nil {
		return m.GroupAddress
	}
	return nil
}

func (m *IgmpEdmI2QIntfRateBag) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *IgmpEdmI2QIntfRateBag) GetReceivedTime() uint64 {
	if m != nil {
		return m.ReceivedTime
	}
	return 0
}

type IgmpEdmI2QIntfStatsBag struct {
	IsVirtualAccess      bool                     `protobuf:"varint,50,opt,name=is_virtual_access,json=isVirtualAccess,proto3" json:"is_virtual_access,omitempty"`
	Rate                 uint32                   `protobuf:"varint,51,opt,name=rate,proto3" json:"rate,omitempty"`
	RateIncrements       uint32                   `protobuf:"varint,52,opt,name=rate_increments,json=rateIncrements,proto3" json:"rate_increments,omitempty"`
	RateDecrements       uint32                   `protobuf:"varint,53,opt,name=rate_decrements,json=rateDecrements,proto3" json:"rate_decrements,omitempty"`
	Update               []*IgmpEdmI2QIntfRateBag `protobuf:"bytes,54,rep,name=update,proto3" json:"update,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *IgmpEdmI2QIntfStatsBag) Reset()         { *m = IgmpEdmI2QIntfStatsBag{} }
func (m *IgmpEdmI2QIntfStatsBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmI2QIntfStatsBag) ProtoMessage()    {}
func (*IgmpEdmI2QIntfStatsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4d11de3db8bf66c, []int{3}
}

func (m *IgmpEdmI2QIntfStatsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag.Unmarshal(m, b)
}
func (m *IgmpEdmI2QIntfStatsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmI2QIntfStatsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmI2QIntfStatsBag.Merge(m, src)
}
func (m *IgmpEdmI2QIntfStatsBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmI2QIntfStatsBag.Size(m)
}
func (m *IgmpEdmI2QIntfStatsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmI2QIntfStatsBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmI2QIntfStatsBag proto.InternalMessageInfo

func (m *IgmpEdmI2QIntfStatsBag) GetIsVirtualAccess() bool {
	if m != nil {
		return m.IsVirtualAccess
	}
	return false
}

func (m *IgmpEdmI2QIntfStatsBag) GetRate() uint32 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *IgmpEdmI2QIntfStatsBag) GetRateIncrements() uint32 {
	if m != nil {
		return m.RateIncrements
	}
	return 0
}

func (m *IgmpEdmI2QIntfStatsBag) GetRateDecrements() uint32 {
	if m != nil {
		return m.RateDecrements
	}
	return 0
}

func (m *IgmpEdmI2QIntfStatsBag) GetUpdate() []*IgmpEdmI2QIntfRateBag {
	if m != nil {
		return m.Update
	}
	return nil
}

func init() {
	proto.RegisterType((*IgmpEdmI2QIntfStatsBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.interface_unicast_qos_adjusts.interface_unicast_qos_adjust.igmp_edm_i2q_intf_stats_bag_KEYS")
	proto.RegisterType((*IgmpAddrtype)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.interface_unicast_qos_adjusts.interface_unicast_qos_adjust.igmp_addrtype")
	proto.RegisterType((*IgmpEdmI2QIntfRateBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.interface_unicast_qos_adjusts.interface_unicast_qos_adjust.igmp_edm_i2q_intf_rate_bag")
	proto.RegisterType((*IgmpEdmI2QIntfStatsBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.interface_unicast_qos_adjusts.interface_unicast_qos_adjust.igmp_edm_i2q_intf_stats_bag")
}

func init() { proto.RegisterFile("igmp_edm_i2q_intf_stats_bag.proto", fileDescriptor_b4d11de3db8bf66c) }

var fileDescriptor_b4d11de3db8bf66c = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe5, 0x26, 0x75, 0xcb, 0x24, 0x4e, 0xc5, 0x4a, 0x80, 0x81, 0x8b, 0x1b, 0x84, 0x88,
	0x38, 0xf8, 0x90, 0x96, 0xdc, 0x23, 0xc1, 0x01, 0x21, 0x71, 0x30, 0x08, 0x89, 0xd3, 0x6a, 0xb3,
	0x3b, 0x0e, 0x03, 0xf5, 0x47, 0x77, 0xd7, 0x86, 0xf2, 0x02, 0x3c, 0x06, 0xf0, 0x12, 0x3c, 0x08,
	0x4f, 0x84, 0x76, 0xe3, 0xb8, 0x8a, 0x40, 0x39, 0x96, 0x8b, 0x65, 0xff, 0xf7, 0xe7, 0x99, 0xff,
	0x7c, 0x68, 0xe1, 0x94, 0xd6, 0x45, 0xcd, 0x51, 0x15, 0x9c, 0xe6, 0x97, 0x9c, 0x4a, 0x9b, 0x73,
	0x63, 0x85, 0x35, 0x7c, 0x25, 0xd6, 0x69, 0xad, 0x2b, 0x5b, 0xb1, 0x4f, 0x92, 0x8c, 0xac, 0x38,
	0x55, 0x86, 0x7f, 0xd1, 0x9c, 0xea, 0xf6, 0x9c, 0xfb, 0x9f, 0xaa, 0x1a, 0x75, 0x5a, 0x5c, 0xa8,
	0xd4, 0x58, 0x51, 0xaa, 0xd5, 0x55, 0xda, 0xea, 0xdc, 0xb8, 0x47, 0x4a, 0xa5, 0x45, 0x9d, 0x0b,
	0x89, 0xbc, 0x29, 0x49, 0x0a, 0x63, 0xf9, 0x65, 0x65, 0xb8, 0x50, 0x1f, 0x1b, 0x63, 0xcd, 0xde,
	0xd3, 0xa9, 0x82, 0x64, 0x8f, 0x23, 0xfe, 0xea, 0xc5, 0xfb, 0x37, 0xec, 0x3e, 0x1c, 0xb7, 0x3a,
	0xe7, 0xa5, 0x28, 0x30, 0x0e, 0x92, 0x60, 0x76, 0x2b, 0x3b, 0x6a, 0x75, 0xfe, 0x5a, 0x14, 0xc8,
	0x1e, 0xc3, 0xe4, 0x3a, 0xbc, 0x07, 0x0e, 0x3c, 0x10, 0xf5, 0xaa, 0xc3, 0xa6, 0x25, 0x44, 0x3e,
	0x8b, 0x50, 0x4a, 0xdb, 0xab, 0x1a, 0xd9, 0x3d, 0x38, 0x12, 0x3b, 0x11, 0x43, 0xb1, 0x09, 0x78,
	0x0a, 0x63, 0x5f, 0xb1, 0x23, 0xd1, 0x98, 0x2e, 0xdc, 0xc8, 0x69, 0xcb, 0x8d, 0xd4, 0x21, 0x8b,
	0x1e, 0x19, 0xf4, 0xc8, 0xa2, 0x43, 0xa6, 0xbf, 0x06, 0xf0, 0xe0, 0xef, 0xb2, 0xb4, 0xb0, 0xe8,
	0xaa, 0x62, 0x77, 0x20, 0x24, 0xd7, 0x01, 0xe5, 0x93, 0x1f, 0x67, 0x87, 0x64, 0x96, 0x4a, 0xb1,
	0x9f, 0x01, 0x4c, 0x4c, 0xd5, 0x68, 0x89, 0x3b, 0xe9, 0x47, 0xf3, 0xaf, 0xe9, 0x0d, 0x8e, 0x24,
	0xdd, 0xe9, 0x54, 0x16, 0x6d, 0x1c, 0x6d, 0x8b, 0xff, 0x1e, 0x40, 0xb4, 0xd6, 0x55, 0x53, 0xef,
	0x94, 0xff, 0x7f, 0x2d, 0x8e, 0xbd, 0xa1, 0xad, 0xc3, 0xbb, 0x10, 0x7e, 0x46, 0x5a, 0x7f, 0xb0,
	0xf1, 0x30, 0x09, 0x66, 0x51, 0xd6, 0x7d, 0xb1, 0x47, 0x10, 0x69, 0x94, 0x48, 0x2d, 0x2a, 0x6e,
	0xa9, 0xc0, 0xf8, 0x30, 0x09, 0x66, 0xc3, 0x6c, 0xbc, 0x15, 0xdf, 0x52, 0x81, 0xd3, 0xdf, 0x07,
	0xf0, 0x70, 0xcf, 0x3e, 0xb2, 0xa7, 0x70, 0x9b, 0x0c, 0x6f, 0x49, 0xdb, 0x46, 0x5c, 0x70, 0x21,
	0xa5, 0xeb, 0xc0, 0xdc, 0x0f, 0xf1, 0x84, 0xcc, 0xbb, 0x8d, 0xbe, 0xf4, 0x32, 0x63, 0x30, 0x74,
	0x13, 0x8f, 0xcf, 0xbc, 0x0d, 0xff, 0xce, 0x9e, 0xc0, 0x89, 0xdf, 0x02, 0x2a, 0xa5, 0xc6, 0x02,
	0x4b, 0x6b, 0xe2, 0x73, 0x7f, 0x3c, 0x71, 0xf2, 0xcb, 0x5e, 0xed, 0x41, 0x85, 0x3d, 0xf8, 0xec,
	0x1a, 0x7c, 0xde, 0xab, 0xec, 0x47, 0x00, 0x61, 0x53, 0x2b, 0x97, 0x68, 0x91, 0x0c, 0x66, 0xa3,
	0xf9, 0xb7, 0xe0, 0xe6, 0x47, 0xf1, 0xcf, 0x35, 0xcf, 0x3a, 0x5f, 0xab, 0xd0, 0xdf, 0x2b, 0x67,
	0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x51, 0xcf, 0xff, 0x23, 0x7c, 0x04, 0x00, 0x00,
}
