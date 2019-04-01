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
// source: bgp_oc_rib_common_bag.proto

package cisco_ios_xr_ipv4_bgp_oc_oper_oc_bgp_bgp_rib_afi_safi_table_ipv4_unicast_open_config_neighbors_open_config_neighbor_adj_rib_in_pre_num_routes

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

type BgpOcRibCommonBag_KEYS struct {
	NeighborAddress      string   `protobuf:"bytes,1,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcRibCommonBag_KEYS) Reset()         { *m = BgpOcRibCommonBag_KEYS{} }
func (m *BgpOcRibCommonBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpOcRibCommonBag_KEYS) ProtoMessage()    {}
func (*BgpOcRibCommonBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8440dd453503ba48, []int{0}
}

func (m *BgpOcRibCommonBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcRibCommonBag_KEYS.Unmarshal(m, b)
}
func (m *BgpOcRibCommonBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcRibCommonBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpOcRibCommonBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcRibCommonBag_KEYS.Merge(m, src)
}
func (m *BgpOcRibCommonBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpOcRibCommonBag_KEYS.Size(m)
}
func (m *BgpOcRibCommonBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcRibCommonBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcRibCommonBag_KEYS proto.InternalMessageInfo

func (m *BgpOcRibCommonBag_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type BgpOcRibCommonBag struct {
	NumRoutes            uint64   `protobuf:"varint,50,opt,name=num_routes,json=numRoutes,proto3" json:"num_routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpOcRibCommonBag) Reset()         { *m = BgpOcRibCommonBag{} }
func (m *BgpOcRibCommonBag) String() string { return proto.CompactTextString(m) }
func (*BgpOcRibCommonBag) ProtoMessage()    {}
func (*BgpOcRibCommonBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_8440dd453503ba48, []int{1}
}

func (m *BgpOcRibCommonBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpOcRibCommonBag.Unmarshal(m, b)
}
func (m *BgpOcRibCommonBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpOcRibCommonBag.Marshal(b, m, deterministic)
}
func (m *BgpOcRibCommonBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpOcRibCommonBag.Merge(m, src)
}
func (m *BgpOcRibCommonBag) XXX_Size() int {
	return xxx_messageInfo_BgpOcRibCommonBag.Size(m)
}
func (m *BgpOcRibCommonBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpOcRibCommonBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpOcRibCommonBag proto.InternalMessageInfo

func (m *BgpOcRibCommonBag) GetNumRoutes() uint64 {
	if m != nil {
		return m.NumRoutes
	}
	return 0
}

func init() {
	proto.RegisterType((*BgpOcRibCommonBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.num_routes.bgp_oc_rib_common_bag_KEYS")
	proto.RegisterType((*BgpOcRibCommonBag)(nil), "cisco_ios_xr_ipv4_bgp_oc_oper.oc_bgp.bgp_rib.afi_safi_table.ipv4_unicast.open_config_neighbors.open_config_neighbor.adj_rib_in_pre.num_routes.bgp_oc_rib_common_bag")
}

func init() { proto.RegisterFile("bgp_oc_rib_common_bag.proto", fileDescriptor_8440dd453503ba48) }

var fileDescriptor_8440dd453503ba48 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0x04, 0x41,
	0x0c, 0x86, 0x59, 0x10, 0xe1, 0xa6, 0x51, 0x16, 0x84, 0x43, 0x11, 0x8e, 0xab, 0xce, 0x26, 0x85,
	0x8a, 0xbd, 0x85, 0x58, 0xd8, 0xad, 0x95, 0x55, 0x98, 0x99, 0x9d, 0x5b, 0x23, 0x6e, 0x32, 0x24,
	0x3b, 0xe2, 0x4b, 0xf8, 0xce, 0xb2, 0xa3, 0xc7, 0x35, 0xdb, 0x04, 0xf2, 0xf1, 0xf3, 0xe7, 0x23,
	0xee, 0x2a, 0x0c, 0x19, 0x25, 0xa2, 0x52, 0xc0, 0x28, 0xe3, 0x28, 0x8c, 0xc1, 0x0f, 0x90, 0x55,
	0x26, 0x69, 0x7f, 0x9a, 0x48, 0x16, 0x05, 0x49, 0x0c, 0xbf, 0x15, 0x29, 0x7f, 0xdd, 0xe3, 0x7f,
	0x5e, 0x72, 0x52, 0x90, 0x38, 0xaf, 0x30, 0x23, 0xa5, 0x00, 0x7e, 0x4f, 0x68, 0xf3, 0x98, 0x7c,
	0xf8, 0x4c, 0x50, 0xe3, 0x85, 0x29, 0x7a, 0x9b, 0x40, 0x72, 0x62, 0x8c, 0xc2, 0x7b, 0x1a, 0x90,
	0x13, 0x0d, 0xef, 0x41, 0xd4, 0x16, 0x29, 0xf8, 0xfe, 0xa3, 0xea, 0x10, 0x63, 0xd6, 0x04, 0x5c,
	0x46, 0x54, 0x29, 0x53, 0xb2, 0xed, 0xb3, 0xbb, 0x5c, 0xd4, 0xc5, 0x97, 0xa7, 0xb7, 0xd7, 0xf6,
	0xc6, 0x9d, 0x1f, 0x3a, 0xd0, 0xf7, 0xbd, 0x26, 0xb3, 0x75, 0xb3, 0x69, 0x76, 0xab, 0xee, 0xec,
	0xc0, 0x1f, 0xff, 0xf0, 0xf6, 0xc1, 0x5d, 0x2c, 0x16, 0xb5, 0xd7, 0xce, 0x1d, 0xef, 0xad, 0x6f,
	0x37, 0xcd, 0xee, 0xa4, 0x5b, 0x71, 0x19, 0xbb, 0x0a, 0xc2, 0x69, 0xfd, 0xcb, 0xdd, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0xc6, 0x53, 0x8b, 0x36, 0x01, 0x00, 0x00,
}
