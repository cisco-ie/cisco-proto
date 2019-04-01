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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_default_context_safs_saf_rpf_hash_sources_rpf_hash_source

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
	Mofrr                uint32   `protobuf:"varint,4,opt,name=mofrr,proto3" json:"mofrr,omitempty"`
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
	proto.RegisterType((*PimRpfHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_sources.rpf_hash_source.pim_rpf_hash_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_sources.rpf_hash_source.pim_addrtype")
	proto.RegisterType((*PimRpfHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.default_context.safs.saf.rpf_hash_sources.rpf_hash_source.pim_rpf_hash_bag")
}

func init() { proto.RegisterFile("pim_rpf_hash_bag.proto", fileDescriptor_9bbebdbd6b24e885) }

var fileDescriptor_9bbebdbd6b24e885 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xbf, 0xae, 0xd3, 0x30,
	0x18, 0xc5, 0x65, 0x28, 0x6d, 0x71, 0xff, 0xa8, 0xb2, 0x80, 0xa6, 0x15, 0x43, 0x28, 0x42, 0xca,
	0x80, 0x3c, 0xb4, 0x25, 0x0b, 0x03, 0x62, 0xa8, 0x04, 0x42, 0x74, 0x08, 0x13, 0x93, 0xe5, 0x26,
	0x4e, 0x63, 0x29, 0x89, 0x2d, 0xdb, 0xad, 0xda, 0xa7, 0x60, 0x61, 0xe4, 0x1d, 0x78, 0x45, 0xe4,
	0x38, 0x09, 0x34, 0xdc, 0xf9, 0xde, 0x25, 0x92, 0xcf, 0x77, 0x9c, 0xef, 0x77, 0x4e, 0x14, 0xf8,
	0x42, 0xf2, 0x82, 0x28, 0x99, 0x92, 0x8c, 0xea, 0x8c, 0x1c, 0xe8, 0x11, 0x4b, 0x25, 0x8c, 0x40,
	0x49, 0xcc, 0x75, 0x2c, 0x08, 0x17, 0x9a, 0x5c, 0x14, 0xe1, 0xf2, 0xbc, 0x25, 0xd6, 0x29, 0x24,
	0x53, 0x98, 0xcb, 0x73, 0x68, 0x4f, 0x98, 0xc6, 0x86, 0x9f, 0x19, 0x4e, 0x58, 0x4a, 0x4f, 0xb9,
	0x21, 0xb1, 0x28, 0x0d, 0xbb, 0x18, 0xac, 0x69, 0xaa, 0xed, 0x03, 0xb7, 0xef, 0xd5, 0xe2, 0xa4,
	0x62, 0xa6, 0xbb, 0xc2, 0xea, 0x27, 0x80, 0xcf, 0xbb, 0x00, 0xe4, 0xcb, 0xee, 0xfb, 0x37, 0xb4,
	0x80, 0x43, 0x4d, 0x53, 0x52, 0xd2, 0x82, 0x79, 0xc0, 0x07, 0xc1, 0xd3, 0x68, 0xa0, 0x69, 0xba,
	0xa7, 0x05, 0x43, 0xaf, 0xe1, 0xc4, 0x08, 0x29, 0x72, 0x71, 0xbc, 0xba, 0xf9, 0xa3, 0x6a, 0x3e,
	0x6e, 0xc4, 0xca, 0xf4, 0x06, 0x4e, 0xdd, 0x0e, 0x42, 0x93, 0x44, 0x31, 0xad, 0xbd, 0xc7, 0x95,
	0x6b, 0xe2, 0xd4, 0x8f, 0x4e, 0x44, 0xcf, 0xe0, 0x93, 0x42, 0xa4, 0x4a, 0x79, 0x3d, 0x1f, 0x04,
	0x93, 0xc8, 0x1d, 0x56, 0x05, 0x1c, 0x5b, 0x2a, 0x7b, 0xd3, 0x5c, 0x25, 0x43, 0x73, 0x38, 0xb8,
	0x65, 0xe9, 0xd7, 0x28, 0xaf, 0xe0, 0xb8, 0xaa, 0xa6, 0xd9, 0xe1, 0x48, 0x46, 0x56, 0x6b, 0x36,
	0x38, 0x4b, 0xd8, 0xc1, 0xb0, 0x96, 0xb0, 0xb6, 0xac, 0x7e, 0xf4, 0xe0, 0xac, 0xdb, 0x02, 0x7a,
	0x0f, 0x97, 0x25, 0xbb, 0x18, 0x92, 0x09, 0x49, 0x8a, 0x53, 0x6e, 0xb8, 0xa4, 0x26, 0x23, 0xac,
	0xa4, 0x87, 0x9c, 0x25, 0xde, 0xda, 0x07, 0xc1, 0x30, 0x9a, 0x5b, 0xc7, 0x27, 0x21, 0xbf, 0x36,
	0xf3, 0x9d, 0x1b, 0xa3, 0x5f, 0x00, 0xce, 0xda, 0xdb, 0xcd, 0xe6, 0x8d, 0x0f, 0x82, 0xd1, 0x5a,
	0xe1, 0xfb, 0xf8, 0xb2, 0xf8, 0xdf, 0xfe, 0xa2, 0x69, 0xcd, 0xd9, 0x74, 0xf2, 0x16, 0xa2, 0x96,
	0x8e, 0x97, 0x86, 0xa9, 0x94, 0xc6, 0xcc, 0xdb, 0x56, 0xcd, 0xcc, 0x6a, 0xef, 0xe7, 0x46, 0x47,
	0xbf, 0x01, 0x5c, 0x6a, 0x16, 0x8b, 0x32, 0xa1, 0xea, 0x4a, 0xfe, 0x8b, 0xf5, 0xee, 0xc1, 0x62,
	0xcd, 0x5b, 0xaa, 0xfd, 0x6d, 0xbe, 0x0f, 0xf0, 0xe5, 0x1d, 0xc0, 0x7f, 0x93, 0x86, 0x55, 0xd2,
	0x45, 0xf7, 0x7a, 0x1b, 0xf9, 0xd0, 0xaf, 0x7e, 0xc2, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4b, 0x3c, 0x19, 0x58, 0x9e, 0x03, 0x00, 0x00,
}
