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
// source: pim_rpf_hash_bag.proto

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_safs_saf_rpf_hash_source_groups_rpf_hash_source_group

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

type PimRpfHashBag_KEYS struct {
	SafName              string   `protobuf:"bytes,1,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,2,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	SourceAddress        string   `protobuf:"bytes,3,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	GroupAddress         string   `protobuf:"bytes,4,opt,name=group_address,json=groupAddress,proto3" json:"group_address,omitempty"`
	MaskLength           uint32   `protobuf:"varint,5,opt,name=mask_length,json=maskLength,proto3" json:"mask_length,omitempty"`
	Mofrr                uint32   `protobuf:"varint,6,opt,name=mofrr,proto3" json:"mofrr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimRpfHashBag_KEYS) Reset()         { *m = PimRpfHashBag_KEYS{} }
func (m *PimRpfHashBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimRpfHashBag_KEYS) ProtoMessage()    {}
func (*PimRpfHashBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bbebdbd6b24e885, []int{0}
}

func (m *PimRpfHashBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfHashBag_KEYS.Unmarshal(m, b)
}
func (m *PimRpfHashBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfHashBag_KEYS.Marshal(b, m, deterministic)
}
func (m *PimRpfHashBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfHashBag_KEYS.Merge(m, src)
}
func (m *PimRpfHashBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimRpfHashBag_KEYS.Size(m)
}
func (m *PimRpfHashBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfHashBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfHashBag_KEYS proto.InternalMessageInfo

func (m *PimRpfHashBag_KEYS) GetSafName() string {
	if m != nil {
		return m.SafName
	}
	return ""
}

func (m *PimRpfHashBag_KEYS) GetTopologyName() string {
	if m != nil {
		return m.TopologyName
	}
	return ""
}

func (m *PimRpfHashBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *PimRpfHashBag_KEYS) GetGroupAddress() string {
	if m != nil {
		return m.GroupAddress
	}
	return ""
}

func (m *PimRpfHashBag_KEYS) GetMaskLength() uint32 {
	if m != nil {
		return m.MaskLength
	}
	return 0
}

func (m *PimRpfHashBag_KEYS) GetMofrr() uint32 {
	if m != nil {
		return m.Mofrr
	}
	return 0
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
	return fileDescriptor_9bbebdbd6b24e885, []int{1}
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

type PimRpfHashBag struct {
	NextHopMultipathEnabled   bool         `protobuf:"varint,50,opt,name=next_hop_multipath_enabled,json=nextHopMultipathEnabled,proto3" json:"next_hop_multipath_enabled,omitempty"`
	NextHopAddress            *PimAddrtype `protobuf:"bytes,51,opt,name=next_hop_address,json=nextHopAddress,proto3" json:"next_hop_address,omitempty"`
	NextHopInterface          string       `protobuf:"bytes,52,opt,name=next_hop_interface,json=nextHopInterface,proto3" json:"next_hop_interface,omitempty"`
	SecondaryNextHopAddress   *PimAddrtype `protobuf:"bytes,53,opt,name=secondary_next_hop_address,json=secondaryNextHopAddress,proto3" json:"secondary_next_hop_address,omitempty"`
	SecondaryNextHopInterface string       `protobuf:"bytes,54,opt,name=secondary_next_hop_interface,json=secondaryNextHopInterface,proto3" json:"secondary_next_hop_interface,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}     `json:"-"`
	XXX_unrecognized          []byte       `json:"-"`
	XXX_sizecache             int32        `json:"-"`
}

func (m *PimRpfHashBag) Reset()         { *m = PimRpfHashBag{} }
func (m *PimRpfHashBag) String() string { return proto.CompactTextString(m) }
func (*PimRpfHashBag) ProtoMessage()    {}
func (*PimRpfHashBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9bbebdbd6b24e885, []int{2}
}

func (m *PimRpfHashBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimRpfHashBag.Unmarshal(m, b)
}
func (m *PimRpfHashBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimRpfHashBag.Marshal(b, m, deterministic)
}
func (m *PimRpfHashBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimRpfHashBag.Merge(m, src)
}
func (m *PimRpfHashBag) XXX_Size() int {
	return xxx_messageInfo_PimRpfHashBag.Size(m)
}
func (m *PimRpfHashBag) XXX_DiscardUnknown() {
	xxx_messageInfo_PimRpfHashBag.DiscardUnknown(m)
}

var xxx_messageInfo_PimRpfHashBag proto.InternalMessageInfo

func (m *PimRpfHashBag) GetNextHopMultipathEnabled() bool {
	if m != nil {
		return m.NextHopMultipathEnabled
	}
	return false
}

func (m *PimRpfHashBag) GetNextHopAddress() *PimAddrtype {
	if m != nil {
		return m.NextHopAddress
	}
	return nil
}

func (m *PimRpfHashBag) GetNextHopInterface() string {
	if m != nil {
		return m.NextHopInterface
	}
	return ""
}

func (m *PimRpfHashBag) GetSecondaryNextHopAddress() *PimAddrtype {
	if m != nil {
		return m.SecondaryNextHopAddress
	}
	return nil
}

func (m *PimRpfHashBag) GetSecondaryNextHopInterface() string {
	if m != nil {
		return m.SecondaryNextHopInterface
	}
	return ""
}

func init() {
	proto.RegisterType((*PimRpfHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_rpf_hash_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_addrtype")
	proto.RegisterType((*PimRpfHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_rpf_hash_bag")
}

func init() { proto.RegisterFile("pim_rpf_hash_bag.proto", fileDescriptor_9bbebdbd6b24e885) }

var fileDescriptor_9bbebdbd6b24e885 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x15, 0xd8, 0xba, 0xf1, 0xda, 0x4e, 0x95, 0x05, 0x34, 0x9b, 0x90, 0x28, 0x45, 0x48,
	0x3d, 0xa0, 0x1c, 0xb6, 0xd1, 0x0b, 0x07, 0xc4, 0x61, 0x12, 0x08, 0xd8, 0xa1, 0x9c, 0x38, 0x3d,
	0xb9, 0xc9, 0x4b, 0x13, 0x91, 0xc4, 0x96, 0xed, 0x56, 0xed, 0x85, 0xaf, 0x03, 0xdf, 0x81, 0x6f,
	0xc2, 0xa7, 0x41, 0xb6, 0xe3, 0xc0, 0xca, 0xce, 0x68, 0x97, 0x4a, 0xfe, 0xbf, 0xbf, 0x9f, 0xff,
	0xbf, 0xbf, 0x1a, 0x78, 0x2c, 0xcb, 0x1a, 0x95, 0xcc, 0xb1, 0xe0, 0xba, 0xc0, 0x25, 0x5f, 0x25,
	0x52, 0x09, 0x23, 0x98, 0x4c, 0x4b, 0x9d, 0x0a, 0x2c, 0x85, 0xc6, 0xad, 0xc2, 0x52, 0x6e, 0x2e,
	0xd1, 0x3a, 0x85, 0x24, 0x95, 0x94, 0x72, 0x33, 0xb7, 0xa7, 0x84, 0xa7, 0xa6, 0xdc, 0x50, 0x92,
	0x51, 0xce, 0xd7, 0x95, 0xc1, 0x54, 0x34, 0x86, 0xb6, 0x26, 0xd1, 0x3c, 0xd7, 0xf6, 0x27, 0xe9,
	0xf6, 0x6a, 0xb1, 0x56, 0x29, 0xe1, 0x4a, 0x89, 0xb5, 0xd4, 0xb7, 0xcb, 0xd3, 0x5f, 0x11, 0x3c,
	0xda, 0x0f, 0x83, 0x1f, 0xae, 0xbe, 0x7c, 0x66, 0xa7, 0x70, 0xac, 0x79, 0x8e, 0x0d, 0xaf, 0x29,
	0x8e, 0x26, 0xd1, 0xec, 0xc1, 0xe2, 0x48, 0xf3, 0xfc, 0x9a, 0xd7, 0xc4, 0x9e, 0xc3, 0xd0, 0x08,
	0x29, 0x2a, 0xb1, 0xda, 0xf9, 0xf9, 0x3d, 0x37, 0x1f, 0x04, 0xd1, 0x99, 0x5e, 0xc0, 0x49, 0xfb,
	0x12, 0xcf, 0x32, 0x45, 0x5a, 0xc7, 0xf7, 0x9d, 0x6b, 0xe8, 0xd5, 0xb7, 0x5e, 0xb4, 0xbb, 0x5c,
	0x92, 0xce, 0x75, 0xe0, 0x77, 0x39, 0x31, 0x98, 0x9e, 0x42, 0xbf, 0xe6, 0xfa, 0x2b, 0x56, 0xd4,
	0xac, 0x4c, 0x11, 0x1f, 0x4e, 0xa2, 0xd9, 0x70, 0x01, 0x56, 0xfa, 0xe8, 0x14, 0xf6, 0x10, 0x0e,
	0x6b, 0x91, 0x2b, 0x15, 0xf7, 0xdc, 0xc8, 0x1f, 0xa6, 0x35, 0x0c, 0x2c, 0x9b, 0xdd, 0x6c, 0x76,
	0x92, 0xd8, 0x18, 0x8e, 0x6e, 0x12, 0xf5, 0x5a, 0xa0, 0x67, 0x30, 0x70, 0x65, 0x87, 0x0c, 0x9e,
	0xa7, 0x6f, 0xb5, 0x10, 0xc1, 0x5b, 0xe6, 0x7b, 0x30, 0xd6, 0x32, 0x6f, 0x2d, 0xd3, 0xef, 0x07,
	0x30, 0xda, 0xef, 0x92, 0xbd, 0x86, 0xb3, 0x86, 0xb6, 0x06, 0x0b, 0x21, 0xb1, 0x5e, 0x57, 0xa6,
	0x94, 0xdc, 0x14, 0x48, 0x0d, 0x5f, 0x56, 0x94, 0xc5, 0xe7, 0x93, 0x68, 0x76, 0xbc, 0x18, 0x5b,
	0xc7, 0x3b, 0x21, 0x3f, 0x85, 0xf9, 0x95, 0x1f, 0xb3, 0x1f, 0x11, 0x8c, 0xba, 0xdb, 0xe1, 0xe5,
	0x8b, 0x49, 0x34, 0xeb, 0x9f, 0x7f, 0x4b, 0xfe, 0xf7, 0x7f, 0x25, 0xf9, 0xbb, 0xcb, 0xc5, 0x49,
	0x9b, 0x39, 0xf4, 0xf3, 0x12, 0x58, 0x97, 0xb4, 0x6c, 0x0c, 0xa9, 0x9c, 0xa7, 0x14, 0x5f, 0xba,
	0x96, 0x46, 0xad, 0xf7, 0x7d, 0xd0, 0xd9, 0xcf, 0x08, 0xce, 0x34, 0xa5, 0xa2, 0xc9, 0xb8, 0xda,
	0xe1, 0x3f, 0x88, 0xaf, 0xee, 0x04, 0xe2, 0xb8, 0x4b, 0x78, 0x7d, 0x93, 0xf5, 0x0d, 0x3c, 0xb9,
	0x25, 0xfc, 0x1f, 0xea, 0xb9, 0xa3, 0x3e, 0xdd, 0xbf, 0xde, 0xe1, 0x2f, 0x7b, 0xee, 0x73, 0xbf,
	0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x43, 0x2a, 0x7c, 0x67, 0x08, 0x04, 0x00, 0x00,
}