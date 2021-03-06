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
// source: pim_ma_idb.proto

package cisco_ios_xr_ipv4_pim_oper_pim6_ma_active_interface_table_interface_by_handles_interface_by_handle

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

type PimMaIdb_KEYS struct {
	InterfaceHandle      uint32   `protobuf:"varint,1,opt,name=interface_handle,json=interfaceHandle,proto3" json:"interface_handle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PimMaIdb_KEYS) Reset()         { *m = PimMaIdb_KEYS{} }
func (m *PimMaIdb_KEYS) String() string { return proto.CompactTextString(m) }
func (*PimMaIdb_KEYS) ProtoMessage()    {}
func (*PimMaIdb_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3d84b92a7bf4682, []int{0}
}

func (m *PimMaIdb_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMaIdb_KEYS.Unmarshal(m, b)
}
func (m *PimMaIdb_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMaIdb_KEYS.Marshal(b, m, deterministic)
}
func (m *PimMaIdb_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMaIdb_KEYS.Merge(m, src)
}
func (m *PimMaIdb_KEYS) XXX_Size() int {
	return xxx_messageInfo_PimMaIdb_KEYS.Size(m)
}
func (m *PimMaIdb_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMaIdb_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PimMaIdb_KEYS proto.InternalMessageInfo

func (m *PimMaIdb_KEYS) GetInterfaceHandle() uint32 {
	if m != nil {
		return m.InterfaceHandle
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
	return fileDescriptor_f3d84b92a7bf4682, []int{1}
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

type PimMaIdb struct {
	InterfaceNameXr      string       `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	VrfName              string       `protobuf:"bytes,51,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceType        string       `protobuf:"bytes,52,opt,name=interface_type,json=interfaceType,proto3" json:"interface_type,omitempty"`
	RpAddress            *PimAddrtype `protobuf:"bytes,53,opt,name=rp_address,json=rpAddress,proto3" json:"rp_address,omitempty"`
	SourceAddress        *PimAddrtype `protobuf:"bytes,54,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PimMaIdb) Reset()         { *m = PimMaIdb{} }
func (m *PimMaIdb) String() string { return proto.CompactTextString(m) }
func (*PimMaIdb) ProtoMessage()    {}
func (*PimMaIdb) Descriptor() ([]byte, []int) {
	return fileDescriptor_f3d84b92a7bf4682, []int{2}
}

func (m *PimMaIdb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PimMaIdb.Unmarshal(m, b)
}
func (m *PimMaIdb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PimMaIdb.Marshal(b, m, deterministic)
}
func (m *PimMaIdb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PimMaIdb.Merge(m, src)
}
func (m *PimMaIdb) XXX_Size() int {
	return xxx_messageInfo_PimMaIdb.Size(m)
}
func (m *PimMaIdb) XXX_DiscardUnknown() {
	xxx_messageInfo_PimMaIdb.DiscardUnknown(m)
}

var xxx_messageInfo_PimMaIdb proto.InternalMessageInfo

func (m *PimMaIdb) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *PimMaIdb) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *PimMaIdb) GetInterfaceType() string {
	if m != nil {
		return m.InterfaceType
	}
	return ""
}

func (m *PimMaIdb) GetRpAddress() *PimAddrtype {
	if m != nil {
		return m.RpAddress
	}
	return nil
}

func (m *PimMaIdb) GetSourceAddress() *PimAddrtype {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*PimMaIdb_KEYS)(nil), "cisco_ios_xr_ipv4_pim_oper.pim6_ma.active.interface_table.interface_by_handles.interface_by_handle.pim_ma_idb_KEYS")
	proto.RegisterType((*PimAddrtype)(nil), "cisco_ios_xr_ipv4_pim_oper.pim6_ma.active.interface_table.interface_by_handles.interface_by_handle.pim_addrtype")
	proto.RegisterType((*PimMaIdb)(nil), "cisco_ios_xr_ipv4_pim_oper.pim6_ma.active.interface_table.interface_by_handles.interface_by_handle.pim_ma_idb")
}

func init() { proto.RegisterFile("pim_ma_idb.proto", fileDescriptor_f3d84b92a7bf4682) }

var fileDescriptor_f3d84b92a7bf4682 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x15, 0x2a, 0xb5, 0xf4, 0xfa, 0x97, 0x2c, 0x84, 0xad, 0x54, 0x42, 0x2a, 0x0c, 0x19,
	0xda, 0x92, 0x89, 0x85, 0x01, 0x09, 0x09, 0x89, 0xa1, 0x30, 0xc0, 0x74, 0x72, 0x12, 0x57, 0x58,
	0xaa, 0x1b, 0xeb, 0x1c, 0xa2, 0xf6, 0x29, 0x18, 0x79, 0x03, 0x9e, 0x13, 0xd9, 0x0e, 0x4e, 0x07,
	0x56, 0xc4, 0x98, 0x2f, 0xbf, 0x4f, 0xf7, 0xf3, 0xd9, 0x30, 0x56, 0x42, 0xa2, 0x64, 0x28, 0xf2,
	0x34, 0x56, 0x54, 0x94, 0x45, 0x98, 0x66, 0x42, 0x67, 0x05, 0x8a, 0x42, 0xe3, 0x8e, 0x50, 0xa8,
	0x6a, 0x89, 0x86, 0x29, 0x14, 0xa7, 0x58, 0x09, 0x99, 0xa0, 0x64, 0x31, 0xcb, 0x4a, 0x51, 0xf1,
	0x58, 0x6c, 0x4b, 0x4e, 0x6b, 0x96, 0x71, 0x2c, 0x59, 0xba, 0x39, 0xfc, 0x4e, 0xf7, 0xf8, 0xc6,
	0xb6, 0xf9, 0x86, 0xeb, 0xdf, 0xc2, 0xe9, 0x0d, 0x8c, 0x9a, 0xb9, 0xf8, 0x70, 0xf7, 0xfa, 0x14,
	0x5e, 0xc2, 0xb8, 0x21, 0x1d, 0x16, 0x05, 0x93, 0x60, 0x36, 0x58, 0x8d, 0x7c, 0x7e, 0xef, 0xda,
	0x12, 0xfa, 0xa6, 0xcd, 0xf2, 0x9c, 0xca, 0xbd, 0xe2, 0xe1, 0x29, 0x74, 0xd8, 0x1a, 0xb7, 0x4c,
	0xba, 0x46, 0x77, 0xd5, 0x66, 0xeb, 0x47, 0x26, 0x79, 0x78, 0x0e, 0x7d, 0xeb, 0x6f, 0x48, 0xae,
	0x75, 0x74, 0x64, 0xff, 0xf6, 0x4c, 0x76, 0xeb, 0xa2, 0x1a, 0x49, 0x3c, 0xd2, 0xf2, 0x48, 0x52,
	0x23, 0xd3, 0xaf, 0x16, 0x40, 0x63, 0x1b, 0x5e, 0xc1, 0x49, 0x23, 0x6a, 0x86, 0xe2, 0x8e, 0xa2,
	0xb9, 0xad, 0x35, 0xa6, 0x66, 0xfc, 0x0b, 0x85, 0x67, 0x70, 0x5c, 0x51, 0xad, 0xb6, 0xb0, 0x48,
	0xa7, 0x22, 0xe7, 0x76, 0x01, 0xc3, 0x83, 0xf5, 0xed, 0x15, 0x8f, 0x96, 0x16, 0x18, 0xf8, 0xf4,
	0xd9, 0x9c, 0xed, 0x23, 0x00, 0x20, 0xe5, 0xf5, 0xae, 0x27, 0xc1, 0xac, 0x37, 0x57, 0xf1, 0xdf,
	0xdf, 0x51, 0x7c, 0xb8, 0xe2, 0x55, 0x97, 0xd4, 0xcf, 0xc6, 0x3e, 0x03, 0x18, 0xea, 0xe2, 0x9d,
	0x32, 0xee, 0xad, 0x92, 0x7f, 0xb2, 0x1a, 0x38, 0x8f, 0xda, 0x2c, 0x6d, 0xdb, 0x07, 0xbc, 0xf8,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xed, 0x7e, 0x5a, 0xd4, 0x02, 0x00, 0x00,
}
