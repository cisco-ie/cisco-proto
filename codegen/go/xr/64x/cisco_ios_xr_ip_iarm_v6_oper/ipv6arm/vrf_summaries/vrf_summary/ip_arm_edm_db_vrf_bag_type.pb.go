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
// source: ip_arm_edm_db_vrf_bag_type.proto

package cisco_ios_xr_ip_iarm_v6_oper_ipv6arm_vrf_summaries_vrf_summary

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

type IpArmEdmDbVrfBagType_KEYS struct {
	VrfName              string   `protobuf:"bytes,1,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArmEdmDbVrfBagType_KEYS) Reset()         { *m = IpArmEdmDbVrfBagType_KEYS{} }
func (m *IpArmEdmDbVrfBagType_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpArmEdmDbVrfBagType_KEYS) ProtoMessage()    {}
func (*IpArmEdmDbVrfBagType_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9532332d7e724399, []int{0}
}

func (m *IpArmEdmDbVrfBagType_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS.Unmarshal(m, b)
}
func (m *IpArmEdmDbVrfBagType_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS.Marshal(b, m, deterministic)
}
func (m *IpArmEdmDbVrfBagType_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS.Merge(m, src)
}
func (m *IpArmEdmDbVrfBagType_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS.Size(m)
}
func (m *IpArmEdmDbVrfBagType_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpArmEdmDbVrfBagType_KEYS proto.InternalMessageInfo

func (m *IpArmEdmDbVrfBagType_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

type IpArmEdmDbVrfBagType struct {
	VrfId                uint32   `protobuf:"varint,50,opt,name=vrf_id,json=vrfId,proto3" json:"vrf_id,omitempty"`
	VrfNameXr            string   `protobuf:"bytes,51,opt,name=vrf_name_xr,json=vrfNameXr,proto3" json:"vrf_name_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpArmEdmDbVrfBagType) Reset()         { *m = IpArmEdmDbVrfBagType{} }
func (m *IpArmEdmDbVrfBagType) String() string { return proto.CompactTextString(m) }
func (*IpArmEdmDbVrfBagType) ProtoMessage()    {}
func (*IpArmEdmDbVrfBagType) Descriptor() ([]byte, []int) {
	return fileDescriptor_9532332d7e724399, []int{1}
}

func (m *IpArmEdmDbVrfBagType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpArmEdmDbVrfBagType.Unmarshal(m, b)
}
func (m *IpArmEdmDbVrfBagType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpArmEdmDbVrfBagType.Marshal(b, m, deterministic)
}
func (m *IpArmEdmDbVrfBagType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpArmEdmDbVrfBagType.Merge(m, src)
}
func (m *IpArmEdmDbVrfBagType) XXX_Size() int {
	return xxx_messageInfo_IpArmEdmDbVrfBagType.Size(m)
}
func (m *IpArmEdmDbVrfBagType) XXX_DiscardUnknown() {
	xxx_messageInfo_IpArmEdmDbVrfBagType.DiscardUnknown(m)
}

var xxx_messageInfo_IpArmEdmDbVrfBagType proto.InternalMessageInfo

func (m *IpArmEdmDbVrfBagType) GetVrfId() uint32 {
	if m != nil {
		return m.VrfId
	}
	return 0
}

func (m *IpArmEdmDbVrfBagType) GetVrfNameXr() string {
	if m != nil {
		return m.VrfNameXr
	}
	return ""
}

func init() {
	proto.RegisterType((*IpArmEdmDbVrfBagType_KEYS)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.vrf_summaries.vrf_summary.ip_arm_edm_db_vrf_bag_type_KEYS")
	proto.RegisterType((*IpArmEdmDbVrfBagType)(nil), "cisco_ios_xr_ip_iarm_v6_oper.ipv6arm.vrf_summaries.vrf_summary.ip_arm_edm_db_vrf_bag_type")
}

func init() { proto.RegisterFile("ip_arm_edm_db_vrf_bag_type.proto", fileDescriptor_9532332d7e724399) }

var fileDescriptor_9532332d7e724399 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x2c, 0x88, 0x4f,
	0x2c, 0xca, 0x8d, 0x4f, 0x4d, 0xc9, 0x8d, 0x4f, 0x49, 0x8a, 0x2f, 0x2b, 0x4a, 0x8b, 0x4f, 0x4a,
	0x4c, 0x8f, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4b, 0xce,
	0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0xcf, 0x2c, 0x88, 0xcf, 0x04,
	0xa9, 0x2f, 0x33, 0x8b, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0xcb, 0x2c, 0x28, 0x33, 0x4b, 0x2c, 0xca,
	0xd5, 0x03, 0x69, 0x2c, 0x2e, 0xcd, 0xcd, 0x4d, 0x2c, 0xca, 0x4c, 0x2d, 0x46, 0xe2, 0x55, 0x2a,
	0xd9, 0x70, 0xc9, 0xe3, 0xb6, 0x23, 0xde, 0xdb, 0x35, 0x32, 0x58, 0x48, 0x92, 0x8b, 0x03, 0x24,
	0x98, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x5e, 0x56, 0x94, 0xe6,
	0x97, 0x98, 0x9b, 0xaa, 0x14, 0xcc, 0x25, 0x85, 0x5b, 0xb7, 0x90, 0x28, 0x17, 0x1b, 0x88, 0x9f,
	0x99, 0x22, 0x61, 0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0xc4, 0x5a, 0x56, 0x94, 0xe6, 0x99, 0x22, 0x24,
	0xc7, 0xc5, 0x0d, 0x33, 0x2f, 0xbe, 0xa2, 0x48, 0xc2, 0x18, 0x6c, 0x24, 0x27, 0xd4, 0xc8, 0x88,
	0xa2, 0x24, 0x36, 0xb0, 0xcf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0x62, 0xf2, 0xd9,
	0xfd, 0x00, 0x00, 0x00,
}