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

package cisco_ios_xr_ipv4_bgp_oper_bgp_instances_instance_instance_standby_default_vrf_information

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
	proto.RegisterType((*BgpVrfInfoBag_KEYS)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.information.bgp_vrf_info_bag_KEYS")
	proto.RegisterType((*BgpVrfInfoBag)(nil), "cisco_ios_xr_ipv4_bgp_oper.bgp.instances.instance.instance_standby.default_vrf.information.bgp_vrf_info_bag")
}

func init() { proto.RegisterFile("bgp_vrf_info_bag.proto", fileDescriptor_40e4e0b5d33b768a) }

var fileDescriptor_40e4e0b5d33b768a = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x41, 0x8b, 0xc2, 0x30,
	0x10, 0x85, 0xe9, 0x65, 0x61, 0xc3, 0x2e, 0x2c, 0x59, 0x76, 0xe9, 0x51, 0xea, 0xc5, 0x53, 0x04,
	0xf5, 0xe8, 0x4d, 0x3d, 0x09, 0x1e, 0xf4, 0xa4, 0x97, 0x21, 0x69, 0xd3, 0x38, 0x60, 0x33, 0x21,
	0x49, 0x8b, 0xfe, 0x7b, 0x49, 0xc1, 0x22, 0x3d, 0xcd, 0x63, 0xde, 0xbc, 0x6f, 0x78, 0xec, 0x5f,
	0x19, 0x07, 0x9d, 0xaf, 0x01, 0x6d, 0x4d, 0xa0, 0xa4, 0x11, 0xce, 0x53, 0x24, 0x7e, 0x29, 0x31,
	0x94, 0x04, 0x48, 0x01, 0xee, 0x1e, 0xd0, 0x75, 0x2b, 0x48, 0x97, 0xe4, 0xb4, 0x17, 0xca, 0x38,
	0x81, 0x36, 0x44, 0x69, 0x4b, 0x1d, 0x06, 0x35, 0x08, 0x48, 0xa3, 0x52, 0x0f, 0x51, 0xe9, 0x5a,
	0xb6, 0xb7, 0x98, 0xf0, 0x22, 0xe1, 0x7d, 0x23, 0x23, 0x92, 0x2d, 0xd6, 0xec, 0x6f, 0xfc, 0x15,
	0xf6, 0xbb, 0xf3, 0x89, 0x4f, 0xd9, 0xf7, 0x00, 0xb1, 0xb2, 0xd1, 0x79, 0x36, 0xc9, 0x66, 0x9f,
	0xc7, 0xaf, 0xd7, 0xf2, 0x20, 0x1b, 0x5d, 0x6c, 0xd8, 0xcf, 0x38, 0xcd, 0xe7, 0xec, 0xd7, 0x53,
	0x1b, 0x35, 0x54, 0x18, 0x22, 0x5a, 0xd3, 0x62, 0xb8, 0x6a, 0x9f, 0x2f, 0xfa, 0x38, 0xef, 0xad,
	0xed, 0xbb, 0xa3, 0x3e, 0xfa, 0x96, 0xcb, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x89, 0x95,
	0xbd, 0xff, 0x00, 0x00, 0x00,
}
