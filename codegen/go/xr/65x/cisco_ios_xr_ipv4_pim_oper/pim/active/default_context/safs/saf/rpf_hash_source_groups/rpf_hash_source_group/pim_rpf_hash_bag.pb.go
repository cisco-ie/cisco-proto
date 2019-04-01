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

package cisco_ios_xr_ipv4_pim_oper_pim_active_default_context_safs_saf_rpf_hash_source_groups_rpf_hash_source_group

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
	proto.RegisterType((*PimRpfHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_rpf_hash_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_addrtype")
	proto.RegisterType((*PimRpfHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.pim.active.default_context.safs.saf.rpf_hash_source_groups.rpf_hash_source_group.pim_rpf_hash_bag")
}

func init() { proto.RegisterFile("pim_rpf_hash_bag.proto", fileDescriptor_9bbebdbd6b24e885) }

var fileDescriptor_9bbebdbd6b24e885 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x65, 0x68, 0xd3, 0x32, 0x49, 0xaa, 0x68, 0x05, 0xc4, 0xad, 0x90, 0x08, 0x41, 0x48,
	0x39, 0x20, 0x1f, 0xda, 0x92, 0x0b, 0x07, 0xc4, 0xa1, 0x12, 0x08, 0xe8, 0x21, 0x9c, 0x38, 0x8d,
	0x36, 0xf6, 0x38, 0xb6, 0x6a, 0x7b, 0x57, 0xbb, 0x9b, 0x2a, 0x79, 0x18, 0x4e, 0xbc, 0x03, 0x0f,
	0xc3, 0xd3, 0xa0, 0xdd, 0xf5, 0x1a, 0x1a, 0x72, 0xa6, 0x17, 0x4b, 0xfe, 0xe7, 0xdf, 0xd9, 0xff,
	0xfb, 0x65, 0xc3, 0x53, 0x59, 0xd6, 0xa8, 0x64, 0x8e, 0x05, 0xd7, 0x05, 0x2e, 0xf9, 0x2a, 0x91,
	0x4a, 0x18, 0xc1, 0x6e, 0xd2, 0x52, 0xa7, 0x02, 0x4b, 0xa1, 0x71, 0xa3, 0xb0, 0x94, 0xb7, 0x97,
	0x68, 0x9d, 0x42, 0x92, 0x4a, 0x64, 0x59, 0x27, 0x3c, 0x35, 0xe5, 0x2d, 0x25, 0x19, 0xe5, 0x7c,
	0x5d, 0x19, 0x4c, 0x45, 0x63, 0x68, 0x63, 0x12, 0xcd, 0x73, 0x6d, 0x1f, 0x49, 0xb7, 0x52, 0x8b,
	0xb5, 0x4a, 0x09, 0x57, 0x4a, 0xac, 0xa5, 0xde, 0x2f, 0x4f, 0x7f, 0x45, 0xf0, 0x64, 0x37, 0x07,
	0x7e, 0xba, 0xfa, 0xf6, 0x95, 0x9d, 0xc2, 0xb1, 0xe6, 0x39, 0x36, 0xbc, 0xa6, 0x38, 0x9a, 0x44,
	0xb3, 0x47, 0x8b, 0x23, 0xcd, 0xf3, 0x6b, 0x5e, 0x13, 0x7b, 0x09, 0x43, 0x23, 0xa4, 0xa8, 0xc4,
	0x6a, 0xeb, 0xe7, 0x0f, 0xdc, 0x7c, 0x10, 0x44, 0x67, 0x7a, 0x05, 0x27, 0xed, 0x4d, 0x3c, 0xcb,
	0x14, 0x69, 0x1d, 0x3f, 0x74, 0xae, 0xa1, 0x57, 0xdf, 0x7b, 0xd1, 0xee, 0x72, 0x49, 0x3a, 0xd7,
	0x81, 0xdf, 0xe5, 0xc4, 0x60, 0x7a, 0x0e, 0xfd, 0x9a, 0xeb, 0x1b, 0xac, 0xa8, 0x59, 0x99, 0x22,
	0x3e, 0x9c, 0x44, 0xb3, 0xe1, 0x02, 0xac, 0xf4, 0xd9, 0x29, 0xec, 0x31, 0x1c, 0xd6, 0x22, 0x57,
	0x2a, 0xee, 0xb9, 0x91, 0x7f, 0x99, 0xd6, 0x30, 0xb0, 0x6c, 0x76, 0xb3, 0xd9, 0x4a, 0x62, 0x63,
	0x38, 0xba, 0x4b, 0xd4, 0x6b, 0x81, 0x5e, 0xc0, 0xc0, 0xf5, 0x1c, 0x32, 0x78, 0x9e, 0xbe, 0xd5,
	0x42, 0x04, 0x6f, 0x99, 0xef, 0xc0, 0x58, 0xcb, 0xbc, 0xb5, 0x4c, 0xbf, 0x1f, 0xc0, 0x68, 0xb7,
	0x4b, 0xf6, 0x16, 0xce, 0x1a, 0xda, 0x18, 0x2c, 0x84, 0xc4, 0x7a, 0x5d, 0x99, 0x52, 0x72, 0x53,
	0x20, 0x35, 0x7c, 0x59, 0x51, 0x16, 0x9f, 0x4f, 0xa2, 0xd9, 0xf1, 0x62, 0x6c, 0x1d, 0x1f, 0x84,
	0xfc, 0x12, 0xe6, 0x57, 0x7e, 0xcc, 0x7e, 0x44, 0x30, 0xea, 0x4e, 0x87, 0x9b, 0x2f, 0x26, 0xd1,
	0xac, 0x7f, 0xbe, 0x4d, 0xfe, 0xe3, 0x67, 0x92, 0xfc, 0x5d, 0xe3, 0xe2, 0xa4, 0x8d, 0x1b, 0xaa,
	0x79, 0x0d, 0xac, 0x0b, 0x59, 0x36, 0x86, 0x54, 0xce, 0x53, 0x8a, 0x2f, 0x5d, 0x41, 0xa3, 0xd6,
	0xfb, 0x31, 0xe8, 0xec, 0x67, 0x04, 0x67, 0x9a, 0x52, 0xd1, 0x64, 0x5c, 0x6d, 0xf1, 0x1f, 0xba,
	0x37, 0xf7, 0x4d, 0x37, 0xee, 0xc2, 0x5d, 0xdf, 0xc5, 0x7c, 0x07, 0xcf, 0xf6, 0xe4, 0xfe, 0x03,
	0x3c, 0x77, 0xc0, 0xa7, 0xbb, 0xc7, 0x3b, 0xf2, 0x65, 0xcf, 0xfd, 0xdf, 0x17, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x66, 0xf7, 0x06, 0xf1, 0xf9, 0x03, 0x00, 0x00,
}