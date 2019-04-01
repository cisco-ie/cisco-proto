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

package cisco_ios_xr_ipv4_pim_oper_ipv6_pim_active_vrfs_vrf_safs_saf_rpf_hash_sources_rpf_hash_source

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
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	SafName              string   `protobuf:"bytes,2,opt,name=saf_name,json=safName,proto3" json:"saf_name,omitempty"`
	TopologyName         string   `protobuf:"bytes,3,opt,name=topology_name,json=topologyName,proto3" json:"topology_name,omitempty"`
	SourceAddress        string   `protobuf:"bytes,4,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Mofrr                uint32   `protobuf:"varint,5,opt,name=mofrr,proto3" json:"mofrr,omitempty"`
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

func (m *PimRpfHashBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

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
	proto.RegisterType((*PimRpfHashBag_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.safs.saf.rpf_hash_sources.rpf_hash_source.pim_rpf_hash_bag_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.safs.saf.rpf_hash_sources.rpf_hash_source.pim_addrtype")
	proto.RegisterType((*PimRpfHashBag)(nil), "cisco_ios_xr_ipv4_pim_oper.ipv6_pim.active.vrfs.vrf.safs.saf.rpf_hash_sources.rpf_hash_source.pim_rpf_hash_bag")
}

func init() { proto.RegisterFile("pim_rpf_hash_bag.proto", fileDescriptor_9bbebdbd6b24e885) }

var fileDescriptor_9bbebdbd6b24e885 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x3f, 0xaf, 0xd3, 0x30,
	0x1c, 0x94, 0x79, 0xef, 0xb5, 0xc5, 0xfd, 0xa3, 0xca, 0x02, 0x9a, 0x56, 0x0c, 0xa1, 0x08, 0x29,
	0x03, 0xf2, 0xd0, 0x96, 0x2c, 0x0c, 0x88, 0xa1, 0x12, 0x08, 0xd1, 0x21, 0x4c, 0x0c, 0xc8, 0x72,
	0x13, 0xa7, 0xb1, 0x68, 0x62, 0xcb, 0x76, 0xa3, 0x76, 0xe4, 0x43, 0xf0, 0x1d, 0xd8, 0xf8, 0x8a,
	0xc8, 0x71, 0x12, 0x68, 0x60, 0xe5, 0x2d, 0x96, 0x7c, 0x77, 0xbf, 0xfc, 0xee, 0x2e, 0x32, 0x7c,
	0x22, 0x79, 0x4e, 0x94, 0x4c, 0x49, 0x46, 0x75, 0x46, 0xf6, 0xf4, 0x80, 0xa5, 0x12, 0x46, 0xa0,
	0x2f, 0x31, 0xd7, 0xb1, 0x20, 0x5c, 0x68, 0x72, 0x56, 0x84, 0xcb, 0x72, 0x43, 0xac, 0x52, 0x48,
	0xa6, 0x30, 0x97, 0x65, 0x68, 0x6f, 0x98, 0xc6, 0x86, 0x97, 0x0c, 0x97, 0x2a, 0xd5, 0xf6, 0xc0,
	0x9a, 0xa6, 0xda, 0x1e, 0xb8, 0xfd, 0xa0, 0x16, 0x27, 0x15, 0x33, 0xdd, 0x05, 0x96, 0x3f, 0x01,
	0x7c, 0xdc, 0xdd, 0x4c, 0x3e, 0x6c, 0x3f, 0x7f, 0x42, 0x73, 0x38, 0x28, 0x55, 0x4a, 0x0a, 0x9a,
	0x33, 0x0f, 0xf8, 0x20, 0x78, 0x18, 0xf5, 0x4b, 0x95, 0xee, 0x68, 0xce, 0x2c, 0xa5, 0x69, 0x4d,
	0x3d, 0x70, 0x94, 0xa6, 0x8e, 0x7a, 0x0e, 0xc7, 0x46, 0x48, 0x71, 0x14, 0x87, 0x8b, 0xe3, 0x6f,
	0x2a, 0x7e, 0xd4, 0x80, 0x95, 0xe8, 0x05, 0x9c, 0xb8, 0xf5, 0x84, 0x26, 0x89, 0x62, 0x5a, 0x7b,
	0xb7, 0x95, 0x6a, 0xec, 0xd0, 0xb7, 0x0e, 0x44, 0x8f, 0xe0, 0x5d, 0x2e, 0x52, 0xa5, 0xbc, 0x3b,
	0x1f, 0x04, 0xe3, 0xc8, 0x5d, 0x96, 0x39, 0x1c, 0x59, 0xc3, 0x76, 0xd2, 0x5c, 0x24, 0x43, 0x33,
	0xd8, 0xa7, 0x57, 0x36, 0x7b, 0xb5, 0x95, 0x67, 0x70, 0x54, 0xd5, 0xd5, 0xec, 0x70, 0x4e, 0x87,
	0x16, 0x6b, 0x36, 0x38, 0x49, 0xd8, 0x4a, 0x6e, 0x5a, 0x49, 0x58, 0x4b, 0x96, 0xdf, 0x6e, 0xe1,
	0xb4, 0x5b, 0x10, 0x7a, 0x0d, 0x17, 0x05, 0x3b, 0x1b, 0x92, 0x09, 0x49, 0xf2, 0xd3, 0xd1, 0x70,
	0x49, 0x4d, 0x46, 0x58, 0x41, 0xf7, 0x47, 0x96, 0x78, 0x2b, 0x1f, 0x04, 0x83, 0x68, 0x66, 0x15,
	0xef, 0x84, 0xfc, 0xd8, 0xf0, 0x5b, 0x47, 0xa3, 0xef, 0x00, 0x4e, 0xdb, 0xe9, 0x66, 0xf3, 0xda,
	0x07, 0xc1, 0x70, 0xf5, 0x15, 0xff, 0xd7, 0xbf, 0x8d, 0xff, 0x2c, 0x2e, 0x9a, 0xd4, 0x06, 0x9b,
	0x32, 0x5e, 0x42, 0xd4, 0xda, 0xe2, 0x85, 0x61, 0x2a, 0xa5, 0x31, 0xf3, 0x36, 0x55, 0x25, 0xd3,
	0x5a, 0xfb, 0xbe, 0xc1, 0xd1, 0x0f, 0x00, 0x17, 0x9a, 0xc5, 0xa2, 0x48, 0xa8, 0xba, 0x90, 0xbf,
	0xf2, 0xbc, 0xba, 0xff, 0x3c, 0xb3, 0xd6, 0xce, 0xee, 0x3a, 0xd8, 0x1b, 0xf8, 0xf4, 0x1f, 0x4e,
	0x7f, 0x47, 0x0c, 0xab, 0x88, 0xf3, 0xee, 0x78, 0x9b, 0x75, 0xdf, 0xab, 0x9e, 0xe2, 0xfa, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x92, 0x72, 0x20, 0xa4, 0x03, 0x00, 0x00,
}
