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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_srms_policy_policy_ipv4_policy_ipv4_backup_policy_mi

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
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
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

func (m *SrmsMiTB_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
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
	proto.RegisterType((*SrmsMiTB_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b_KEYS")
	proto.RegisterType((*Addr)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.addr")
	proto.RegisterType((*SrmsMiTB)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.srms.policy.policy_ipv4.policy_ipv4_backup.policy_mi.srms_mi_t_b")
}

func init() { proto.RegisterFile("srms_mi_t_b.proto", fileDescriptor_bb21d79d36617d23) }

var fileDescriptor_bb21d79d36617d23 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x8b, 0x1a, 0x31,
	0x14, 0xc7, 0x19, 0xb5, 0x52, 0xdf, 0xa8, 0xd8, 0x14, 0x24, 0xd0, 0x43, 0xad, 0xed, 0xc1, 0xd3,
	0x1c, 0xd4, 0xda, 0x5f, 0x50, 0x28, 0xa5, 0x07, 0x5b, 0x28, 0x45, 0x4f, 0x3d, 0x85, 0x98, 0x64,
	0xdc, 0xb0, 0xc6, 0x0c, 0x49, 0x5c, 0xdc, 0xe3, 0xfe, 0x01, 0xfb, 0x3f, 0x2f, 0x2f, 0x93, 0x11,
	0xff, 0x82, 0x3d, 0xf9, 0x7d, 0x9f, 0xf7, 0xf2, 0x79, 0x31, 0x0c, 0xbc, 0xf2, 0xce, 0x78, 0x66,
	0x34, 0x0b, 0x6c, 0x57, 0x54, 0xce, 0x06, 0x4b, 0xf6, 0x42, 0x7b, 0x61, 0x99, 0xb6, 0x9e, 0x9d,
	0x1d, 0xd3, 0xd5, 0xdd, 0x92, 0x59, 0x5f, 0x95, 0xcc, 0x56, 0xca, 0x15, 0x98, 0x70, 0x4e, 0x28,
	0xef, 0x95, 0x6f, 0x52, 0x81, 0x96, 0xa2, 0xb2, 0x07, 0x2d, 0xee, 0xd3, 0x4f, 0x3c, 0x78, 0x9d,
	0xd9, 0x8e, 0x8b, 0xdb, 0x53, 0xd5, 0x20, 0xa3, 0xa7, 0xbf, 0x61, 0x74, 0xb5, 0x9d, 0xfd, 0xf9,
	0xf5, 0x7f, 0x4b, 0xde, 0x41, 0x3f, 0x39, 0xd9, 0x91, 0x1b, 0x45, 0xb3, 0x49, 0x36, 0xeb, 0x6d,
	0xf2, 0xc4, 0xfe, 0x72, 0xa3, 0xc8, 0x6b, 0x78, 0x61, 0x34, 0xd3, 0x92, 0xb6, 0x62, 0xaf, 0x63,
	0xf4, 0x5a, 0x4e, 0xbf, 0x43, 0x87, 0x4b, 0xe9, 0xc8, 0x10, 0x5a, 0xbc, 0x4c, 0xa7, 0x5a, 0xbc,
	0x24, 0x04, 0x3a, 0xb8, 0xbc, 0x99, 0xc5, 0x9c, 0xd8, 0x8a, 0xb6, 0x2f, 0x6c, 0x35, 0x7d, 0x6c,
	0x43, 0x7e, 0x75, 0x19, 0x32, 0x82, 0xb6, 0x77, 0x82, 0xce, 0xe3, 0x08, 0x46, 0x32, 0x86, 0xae,
	0xb3, 0xa7, 0xa0, 0x1c, 0x5d, 0x44, 0x98, 0x2a, 0xb4, 0x71, 0xa7, 0x38, 0x5d, 0xd6, 0x36, 0xcc,
	0xe4, 0x21, 0xab, 0xaf, 0x43, 0x3f, 0x4e, 0xb2, 0x59, 0x3e, 0x37, 0xc5, 0x33, 0x3d, 0x69, 0x81,
	0x4b, 0x37, 0xf5, 0x4b, 0x8c, 0xa1, 0x5b, 0x39, 0x55, 0xea, 0x33, 0x5d, 0x4d, 0xb2, 0xd9, 0x60,
	0x93, 0x2a, 0xf2, 0x06, 0x7a, 0x5e, 0x4b, 0xe6, 0x03, 0x77, 0x81, 0x7e, 0x8a, 0xad, 0x97, 0x5e,
	0xcb, 0x2d, 0xd6, 0x4d, 0x53, 0xd8, 0xd3, 0x31, 0xd0, 0xcf, 0x97, 0xe6, 0x4f, 0xac, 0xc9, 0x5b,
	0xc8, 0x0f, 0xdc, 0x07, 0x96, 0xb4, 0x5f, 0xe2, 0x1f, 0x06, 0x44, 0xff, 0x6a, 0xf5, 0x07, 0x18,
	0xc6, 0x01, 0x54, 0xe8, 0xa3, 0x54, 0x67, 0xfa, 0x35, 0x2a, 0xfa, 0x48, 0xb7, 0x5a, 0xae, 0x91,
	0x91, 0xf7, 0x30, 0x28, 0x0f, 0x7c, 0xcf, 0x78, 0x08, 0x5c, 0xdc, 0x28, 0x49, 0xbf, 0x45, 0x51,
	0x1f, 0xe1, 0x8f, 0xc4, 0x76, 0xdd, 0xf8, 0x2d, 0x2e, 0x9e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5e,
	0xf9, 0x7d, 0x5f, 0xa0, 0x02, 0x00, 0x00,
}
