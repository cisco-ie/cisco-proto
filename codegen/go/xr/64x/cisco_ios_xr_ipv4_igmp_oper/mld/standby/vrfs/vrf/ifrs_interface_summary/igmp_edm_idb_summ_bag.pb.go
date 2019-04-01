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
// source: igmp_edm_idb_summ_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_standby_vrfs_vrf_ifrs_interface_summary

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

type IgmpEdmIdbSummBag_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmIdbSummBag_KEYS) Reset()         { *m = IgmpEdmIdbSummBag_KEYS{} }
func (m *IgmpEdmIdbSummBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbSummBag_KEYS) ProtoMessage()    {}
func (*IgmpEdmIdbSummBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edaa2fdc71633cc, []int{0}
}

func (m *IgmpEdmIdbSummBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbSummBag_KEYS.Unmarshal(m, b)
}
func (m *IgmpEdmIdbSummBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbSummBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbSummBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbSummBag_KEYS.Merge(m, src)
}
func (m *IgmpEdmIdbSummBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbSummBag_KEYS.Size(m)
}
func (m *IgmpEdmIdbSummBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbSummBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbSummBag_KEYS proto.InternalMessageInfo

func (m *IgmpEdmIdbSummBag_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IgmpEdmIdbSummBag struct {
	InterfaceCount       uint32   `protobuf:"varint,50,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	ConfigurationCount   uint32   `protobuf:"varint,51,opt,name=configuration_count,json=configurationCount,proto3" json:"configuration_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IgmpEdmIdbSummBag) Reset()         { *m = IgmpEdmIdbSummBag{} }
func (m *IgmpEdmIdbSummBag) String() string { return proto.CompactTextString(m) }
func (*IgmpEdmIdbSummBag) ProtoMessage()    {}
func (*IgmpEdmIdbSummBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_3edaa2fdc71633cc, []int{1}
}

func (m *IgmpEdmIdbSummBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IgmpEdmIdbSummBag.Unmarshal(m, b)
}
func (m *IgmpEdmIdbSummBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IgmpEdmIdbSummBag.Marshal(b, m, deterministic)
}
func (m *IgmpEdmIdbSummBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IgmpEdmIdbSummBag.Merge(m, src)
}
func (m *IgmpEdmIdbSummBag) XXX_Size() int {
	return xxx_messageInfo_IgmpEdmIdbSummBag.Size(m)
}
func (m *IgmpEdmIdbSummBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IgmpEdmIdbSummBag.DiscardUnknown(m)
}

var xxx_messageInfo_IgmpEdmIdbSummBag proto.InternalMessageInfo

func (m *IgmpEdmIdbSummBag) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *IgmpEdmIdbSummBag) GetConfigurationCount() uint32 {
	if m != nil {
		return m.ConfigurationCount
	}
	return 0
}

func init() {
	proto.RegisterType((*IgmpEdmIdbSummBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.ifrs_interface_summary.igmp_edm_idb_summ_bag_KEYS")
	proto.RegisterType((*IgmpEdmIdbSummBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.vrfs.vrf.ifrs_interface_summary.igmp_edm_idb_summ_bag")
}

func init() { proto.RegisterFile("igmp_edm_idb_summ_bag.proto", fileDescriptor_3edaa2fdc71633cc) }

var fileDescriptor_3edaa2fdc71633cc = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3d, 0x4b, 0x04, 0x31,
	0x10, 0x86, 0xd9, 0xc6, 0x8f, 0x80, 0x0a, 0x11, 0xe1, 0xd4, 0xe6, 0xb8, 0xc6, 0xab, 0x22, 0x78,
	0x82, 0x3f, 0x40, 0xc4, 0x42, 0xb0, 0x38, 0x2b, 0xab, 0x21, 0x9f, 0xcb, 0x80, 0xc9, 0xc4, 0x49,
	0x36, 0x78, 0xff, 0x5e, 0x2e, 0x88, 0x22, 0x6c, 0x33, 0xc5, 0x3b, 0xcf, 0x3b, 0x0f, 0x23, 0xae,
	0x71, 0x8c, 0x19, 0xbc, 0x8b, 0x80, 0xce, 0x40, 0x99, 0x62, 0x04, 0xa3, 0x47, 0x95, 0x99, 0x2a,
	0xc9, 0x67, 0x8b, 0xc5, 0x12, 0x20, 0x15, 0xf8, 0x62, 0xc0, 0xdc, 0xee, 0xa1, 0xe3, 0x94, 0x3d,
	0xab, 0xf8, 0xe1, 0x54, 0xa9, 0x3a, 0x39, 0xb3, 0x53, 0x8d, 0x43, 0xd9, 0x0f, 0x85, 0x81, 0x0b,
	0x60, 0xaa, 0x9e, 0x83, 0xb6, 0xbe, 0xdf, 0xd3, 0xbc, 0x5b, 0x3d, 0x88, 0xab, 0x59, 0x0f, 0xbc,
	0x3c, 0xbd, 0xbf, 0xc9, 0x4b, 0x71, 0xd4, 0x38, 0x40, 0xd2, 0xd1, 0x2f, 0x86, 0xe5, 0xb0, 0x3e,
	0xde, 0x1e, 0x36, 0x0e, 0xaf, 0x3a, 0xfa, 0xd5, 0xa7, 0xb8, 0x98, 0x2d, 0xca, 0x1b, 0x71, 0xf6,
	0xa7, 0xb1, 0x34, 0xa5, 0xba, 0xb8, 0x5b, 0x0e, 0xeb, 0x93, 0xed, 0xe9, 0x6f, 0xfc, 0xb8, 0x4f,
	0xe5, 0xad, 0x38, 0xb7, 0x94, 0x02, 0x8e, 0x13, 0xeb, 0x8a, 0x94, 0x7e, 0xe0, 0x4d, 0x87, 0xe5,
	0xbf, 0x55, 0x2f, 0x98, 0x83, 0xfe, 0xfb, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x52, 0xcb, 0x1c,
	0x5a, 0x1a, 0x01, 0x00, 0x00,
}
