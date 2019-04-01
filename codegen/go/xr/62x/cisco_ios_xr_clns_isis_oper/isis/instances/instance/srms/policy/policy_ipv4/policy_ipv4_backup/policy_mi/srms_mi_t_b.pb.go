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
// source: srms_mi_t_b.proto

package cisco_ios_xr_clns_isis_oper_isis_instances_instance_srms_policy_policy_ipv4_policy_ipv4_backup_policy_mi

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

type SrmsMiTB_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	MiId                 string   `protobuf:"bytes,2,opt,name=mi_id,json=miId,proto3" json:"mi_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB_KEYS) Reset()         { *m = SrmsMiTB_KEYS{} }
func (m *SrmsMiTB_KEYS) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB_KEYS) ProtoMessage()    {}
func (*SrmsMiTB_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{0}
}

func (m *SrmsMiTB_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB_KEYS.Unmarshal(m, b)
}
func (m *SrmsMiTB_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB_KEYS.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB_KEYS.Merge(m, src)
}
func (m *SrmsMiTB_KEYS) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB_KEYS.Size(m)
}
func (m *SrmsMiTB_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB_KEYS proto.InternalMessageInfo

func (m *SrmsMiTB_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *SrmsMiTB_KEYS) GetMiId() string {
	if m != nil {
		return m.MiId
	}
	return ""
}

type Addr struct {
	Af                   string   `protobuf:"bytes,1,opt,name=af,proto3" json:"af,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{1}
}

func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetAf() string {
	if m != nil {
		return m.Af
	}
	return ""
}

func (m *Addr) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *Addr) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type SrmsMiTB struct {
	Src                  string   `protobuf:"bytes,50,opt,name=src,proto3" json:"src,omitempty"`
	Router               string   `protobuf:"bytes,51,opt,name=router,proto3" json:"router,omitempty"`
	Area                 string   `protobuf:"bytes,52,opt,name=area,proto3" json:"area,omitempty"`
	Addr                 *Addr    `protobuf:"bytes,53,opt,name=addr,proto3" json:"addr,omitempty"`
	Prefix               uint32   `protobuf:"varint,54,opt,name=prefix,proto3" json:"prefix,omitempty"`
	SidStart             uint32   `protobuf:"varint,55,opt,name=sid_start,json=sidStart,proto3" json:"sid_start,omitempty"`
	SidCount             uint32   `protobuf:"varint,56,opt,name=sid_count,json=sidCount,proto3" json:"sid_count,omitempty"`
	LastPrefix           string   `protobuf:"bytes,57,opt,name=last_prefix,json=lastPrefix,proto3" json:"last_prefix,omitempty"`
	LastSidIndex         uint32   `protobuf:"varint,58,opt,name=last_sid_index,json=lastSidIndex,proto3" json:"last_sid_index,omitempty"`
	FlagAttached         string   `protobuf:"bytes,59,opt,name=flag_attached,json=flagAttached,proto3" json:"flag_attached,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SrmsMiTB) Reset()         { *m = SrmsMiTB{} }
func (m *SrmsMiTB) String() string { return proto.CompactTextString(m) }
func (*SrmsMiTB) ProtoMessage()    {}
func (*SrmsMiTB) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb21d79d36617d23, []int{2}
}

func (m *SrmsMiTB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SrmsMiTB.Unmarshal(m, b)
}
func (m *SrmsMiTB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SrmsMiTB.Marshal(b, m, deterministic)
}
func (m *SrmsMiTB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SrmsMiTB.Merge(m, src)
}
func (m *SrmsMiTB) XXX_Size() int {
	return xxx_messageInfo_SrmsMiTB.Size(m)
}
func (m *SrmsMiTB) XXX_DiscardUnknown() {
	xxx_messageInfo_SrmsMiTB.DiscardUnknown(m)
}

var xxx_messageInfo_SrmsMiTB proto.InternalMessageInfo

func (m *SrmsMiTB) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *SrmsMiTB) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *SrmsMiTB) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *SrmsMiTB) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *SrmsMiTB) GetPrefix() uint32 {
	if m != nil {
		return m.Prefix
	}
	return 0
}

func (m *SrmsMiTB) GetSidStart() uint32 {
	if m != nil {
		return m.SidStart
	}
	return 0
}

func (m *SrmsMiTB) GetSidCount() uint32 {
	if m != nil {
		return m.SidCount
	}
	return 0
}

func (m *SrmsMiTB) GetLastPrefix() string {
	if m != nil {
		return m.LastPrefix
	}
	return ""
}

func (m *SrmsMiTB) GetLastSidIndex() uint32 {
	if m != nil {
		return m.LastSidIndex
	}
	return 0
}

func (m *SrmsMiTB) GetFlagAttached() string {
	if m != nil {
		return m.FlagAttached
	}
	return ""
}

func init() {
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.addr")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor_bb21d79d36617d23) }

var fileDescriptor_bb21d79d36617d23 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0xae, 0xd3, 0x30,
	0x10, 0x85, 0x95, 0xb6, 0x54, 0xd4, 0xfd, 0x51, 0x31, 0x52, 0x65, 0x89, 0x05, 0x55, 0x61, 0xd1,
	0x55, 0x16, 0x6d, 0x29, 0x7f, 0x12, 0x12, 0x42, 0x2c, 0x2a, 0x10, 0x42, 0xed, 0x8a, 0xd5, 0xc8,
	0xb5, 0x1d, 0x3a, 0x22, 0x89, 0x23, 0xdb, 0x45, 0x65, 0xcb, 0x03, 0xdc, 0x67, 0xbe, 0x1a, 0x27,
	0xe9, 0xed, 0x13, 0xdc, 0x55, 0xce, 0x7c, 0x67, 0x72, 0xc6, 0xf1, 0x84, 0x3d, 0xf3, 0xae, 0xf0,
	0x50, 0x20, 0x04, 0x38, 0xa6, 0x95, 0xb3, 0xc1, 0xf2, 0x93, 0x42, 0xaf, 0x2c, 0xa0, 0xf5, 0x70,
	0x71, 0xa0, 0xf2, 0xd2, 0x03, 0x7a, 0xf4, 0x60, 0x2b, 0xe3, 0x52, 0x52, 0x29, 0x96, 0x3e, 0xc8,
	0x52, 0x99, 0x07, 0x95, 0x52, 0x4c, 0x5a, 0xd9, 0x1c, 0xd5, 0xbf, 0xe6, 0x01, 0x58, 0xfd, 0xdd,
	0xdc, 0x6a, 0x38, 0x4a, 0xf5, 0xe7, 0x5c, 0xb5, 0xa8, 0xc0, 0xc5, 0x77, 0x36, 0xbd, 0x19, 0x0f,
	0xdf, 0xbe, 0xfe, 0x3a, 0xf0, 0x57, 0x6c, 0xdc, 0x86, 0x42, 0x29, 0x0b, 0x23, 0x92, 0x79, 0xb2,
	0x1c, 0xec, 0x47, 0x2d, 0xfc, 0x21, 0x0b, 0xc3, 0x9f, 0xb3, 0x27, 0x05, 0x02, 0x6a, 0xd1, 0x89,
	0x66, 0xaf, 0xc0, 0x9d, 0x5e, 0x7c, 0x62, 0x3d, 0xa9, 0xb5, 0xe3, 0x13, 0xd6, 0x91, 0x59, 0xf3,
	0x5a, 0x47, 0x66, 0x9c, 0xb3, 0x1e, 0x8d, 0x6f, 0x7b, 0x49, 0x37, 0x6c, 0x2b, 0xba, 0x57, 0xb6,
	0x5d, 0xdc, 0x75, 0xd9, 0xf0, 0xe6, 0x38, 0x7c, 0xca, 0xba, 0xde, 0x29, 0xb1, 0x8a, 0x2d, 0x24,
	0xf9, 0x8c, 0xf5, 0x9d, 0x3d, 0x07, 0xe3, 0xc4, 0x3a, 0xc2, 0xa6, 0xa2, 0x34, 0xe9, 0x8c, 0x14,
	0x9b, 0x3a, 0x8d, 0x34, 0xff, 0x9f, 0xd4, 0xc7, 0x11, 0x6f, 0xe6, 0xc9, 0x72, 0xb8, 0x2a, 0xd3,
	0xc7, 0xba, 0xd5, 0x94, 0xa6, 0xee, 0xeb, 0xab, 0x98, 0xb1, 0x7e, 0xe5, 0x4c, 0x86, 0x17, 0xb1,
	0x9d, 0x27, 0xcb, 0xf1, 0xbe, 0xa9, 0xf8, 0x0b, 0x36, 0xf0, 0xa8, 0xc1, 0x07, 0xe9, 0x82, 0x78,
	0x1b, 0xad, 0xa7, 0x1e, 0xf5, 0x81, 0xea, 0xd6, 0x54, 0xf6, 0x5c, 0x06, 0xf1, 0xee, 0x6a, 0x7e,
	0xa1, 0x9a, 0xbf, 0x64, 0xc3, 0x5c, 0xfa, 0x00, 0x4d, 0xec, 0xfb, 0xf8, 0xc5, 0x8c, 0xd0, 0xcf,
	0x3a, 0xfa, 0x35, 0x9b, 0xc4, 0x06, 0x8a, 0xc0, 0x52, 0x9b, 0x8b, 0xf8, 0x10, 0x23, 0x46, 0x44,
	0x0f, 0xa8, 0x77, 0xc4, 0x68, 0xcb, 0x59, 0x2e, 0x7f, 0x83, 0x0c, 0x41, 0xaa, 0x93, 0xd1, 0xe2,
	0x63, 0xbd, 0x65, 0x82, 0x9f, 0x1b, 0x76, 0xec, 0xc7, 0xff, 0x71, 0x7d, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0x78, 0x66, 0x38, 0xd2, 0xa4, 0x02, 0x00, 0x00,
}
