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
// source: bgp_vrf_info_bag.proto

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_active_vrfs_vrf_information

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

type BgpVrfInfoBag_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpVrfInfoBag_KEYS) Reset()         { *m = BgpVrfInfoBag_KEYS{} }
func (m *BgpVrfInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*BgpVrfInfoBag_KEYS) ProtoMessage()    {}
func (*BgpVrfInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_40e4e0b5d33b768a, []int{0}
}

func (m *BgpVrfInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpVrfInfoBag_KEYS.Unmarshal(m, b)
}
func (m *BgpVrfInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpVrfInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *BgpVrfInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpVrfInfoBag_KEYS.Merge(m, src)
}
func (m *BgpVrfInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_BgpVrfInfoBag_KEYS.Size(m)
}
func (m *BgpVrfInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpVrfInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BgpVrfInfoBag_KEYS proto.InternalMessageInfo

func (m *BgpVrfInfoBag_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *BgpVrfInfoBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type BgpVrfInfoBag struct {
	RouteDistinguisher   string   `protobuf:"bytes,50,opt,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BgpVrfInfoBag) Reset()         { *m = BgpVrfInfoBag{} }
func (m *BgpVrfInfoBag) String() string { return proto.CompactTextString(m) }
func (*BgpVrfInfoBag) ProtoMessage()    {}
func (*BgpVrfInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_40e4e0b5d33b768a, []int{1}
}

func (m *BgpVrfInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BgpVrfInfoBag.Unmarshal(m, b)
}
func (m *BgpVrfInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BgpVrfInfoBag.Marshal(b, m, deterministic)
}
func (m *BgpVrfInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BgpVrfInfoBag.Merge(m, src)
}
func (m *BgpVrfInfoBag) XXX_Size() int {
	return xxx_messageInfo_BgpVrfInfoBag.Size(m)
}
func (m *BgpVrfInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_BgpVrfInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_BgpVrfInfoBag proto.InternalMessageInfo

func (m *BgpVrfInfoBag) GetRouteDistinguisher() string {
	if m != nil {
		return m.RouteDistinguisher
	}
	return ""
}

func init() {
	proto.RegisterType((*BgpVrfInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.information.bgp_vrf_info_bag_KEYS")
	proto.RegisterType((*BgpVrfInfoBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_active.vrfs.vrf.information.bgp_vrf_info_bag")
}

func init() { proto.RegisterFile("bgp_vrf_info_bag.proto", fileDescriptor_40e4e0b5d33b768a) }

var fileDescriptor_40e4e0b5d33b768a = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4b, 0x07, 0x21,
	0x10, 0x86, 0xd9, 0x0e, 0xfd, 0x91, 0x82, 0x30, 0x8a, 0xed, 0x16, 0xdb, 0xa5, 0x93, 0x41, 0xf5,
	0x0d, 0xaa, 0x53, 0xd0, 0xa1, 0xa0, 0xe8, 0x34, 0xa8, 0xa9, 0xcd, 0x61, 0x1d, 0x19, 0x5d, 0xe9,
	0xe3, 0x87, 0x42, 0xcb, 0x8f, 0xbd, 0xc8, 0x8b, 0xcf, 0xbc, 0xcf, 0x30, 0xe2, 0xc2, 0x84, 0x04,
	0x95, 0x3d, 0x60, 0xf4, 0x04, 0x46, 0x07, 0x95, 0x98, 0x0a, 0xc9, 0x0f, 0x8b, 0xd9, 0x12, 0x20,
	0x65, 0xf8, 0x65, 0xc0, 0x54, 0x1f, 0xa0, 0x4d, 0x52, 0x72, 0xac, 0x4c, 0x48, 0x0a, 0x63, 0x2e,
	0x3a, 0x5a, 0x97, 0xd7, 0xb4, 0x06, 0xd0, 0xb6, 0x60, 0x75, 0xaa, 0xb2, 0xcf, 0xed, 0x51, 0x4d,
	0xcd, 0xb3, 0x2e, 0x48, 0x71, 0xfa, 0x14, 0xe7, 0xdb, 0x8d, 0xf0, 0xf2, 0xfc, 0xf5, 0x2e, 0xaf,
	0xc5, 0xc9, 0x2a, 0x88, 0x7a, 0x76, 0xe3, 0x70, 0x35, 0xdc, 0x1c, 0xbd, 0x1d, 0xff, 0x7f, 0xbe,
	0xea, 0xd9, 0xc9, 0x4b, 0x71, 0xd8, 0x9a, 0x9d, 0xef, 0x75, 0x7e, 0x50, 0xd9, 0x37, 0x34, 0x3d,
	0x8a, 0xd3, 0xad, 0x58, 0xde, 0x8a, 0x33, 0xa6, 0xa5, 0x38, 0xf8, 0xc6, 0x5c, 0x30, 0x86, 0x05,
	0xf3, 0x8f, 0xe3, 0xf1, 0xae, 0x37, 0x65, 0x47, 0x4f, 0xbb, 0xc4, 0xec, 0xf7, 0xe3, 0xef, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x58, 0xbe, 0xb3, 0x69, 0x16, 0x01, 0x00, 0x00,
}
