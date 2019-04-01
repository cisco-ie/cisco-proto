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
// source: ospf_sh_protocol.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_process_information_protocol_summary

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

type OspfShProtocol_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShProtocol_KEYS) Reset()         { *m = OspfShProtocol_KEYS{} }
func (m *OspfShProtocol_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShProtocol_KEYS) ProtoMessage()    {}
func (*OspfShProtocol_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4559d5f25e5c5134, []int{0}
}

func (m *OspfShProtocol_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtocol_KEYS.Unmarshal(m, b)
}
func (m *OspfShProtocol_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtocol_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShProtocol_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtocol_KEYS.Merge(m, src)
}
func (m *OspfShProtocol_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShProtocol_KEYS.Size(m)
}
func (m *OspfShProtocol_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtocol_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtocol_KEYS proto.InternalMessageInfo

func (m *OspfShProtocol_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type OspfShProtocol struct {
	ProtocolRouterId                string   `protobuf:"bytes,50,opt,name=protocol_router_id,json=protocolRouterId,proto3" json:"protocol_router_id,omitempty"`
	ProtocolDistance                uint32   `protobuf:"varint,51,opt,name=protocol_distance,json=protocolDistance,proto3" json:"protocol_distance,omitempty"`
	AdministrativeDistanceInterArea uint32   `protobuf:"varint,52,opt,name=administrative_distance_inter_area,json=administrativeDistanceInterArea,proto3" json:"administrative_distance_inter_area,omitempty"`
	AdministrativeDistanceExternal  uint32   `protobuf:"varint,53,opt,name=administrative_distance_external,json=administrativeDistanceExternal,proto3" json:"administrative_distance_external,omitempty"`
	ProtocolNsf                     bool     `protobuf:"varint,54,opt,name=protocol_nsf,json=protocolNsf,proto3" json:"protocol_nsf,omitempty"`
	DistListIn                      string   `protobuf:"bytes,55,opt,name=dist_list_in,json=distListIn,proto3" json:"dist_list_in,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *OspfShProtocol) Reset()         { *m = OspfShProtocol{} }
func (m *OspfShProtocol) String() string { return proto.CompactTextString(m) }
func (*OspfShProtocol) ProtoMessage()    {}
func (*OspfShProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_4559d5f25e5c5134, []int{1}
}

func (m *OspfShProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtocol.Unmarshal(m, b)
}
func (m *OspfShProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtocol.Marshal(b, m, deterministic)
}
func (m *OspfShProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtocol.Merge(m, src)
}
func (m *OspfShProtocol) XXX_Size() int {
	return xxx_messageInfo_OspfShProtocol.Size(m)
}
func (m *OspfShProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtocol proto.InternalMessageInfo

func (m *OspfShProtocol) GetProtocolRouterId() string {
	if m != nil {
		return m.ProtocolRouterId
	}
	return ""
}

func (m *OspfShProtocol) GetProtocolDistance() uint32 {
	if m != nil {
		return m.ProtocolDistance
	}
	return 0
}

func (m *OspfShProtocol) GetAdministrativeDistanceInterArea() uint32 {
	if m != nil {
		return m.AdministrativeDistanceInterArea
	}
	return 0
}

func (m *OspfShProtocol) GetAdministrativeDistanceExternal() uint32 {
	if m != nil {
		return m.AdministrativeDistanceExternal
	}
	return 0
}

func (m *OspfShProtocol) GetProtocolNsf() bool {
	if m != nil {
		return m.ProtocolNsf
	}
	return false
}

func (m *OspfShProtocol) GetDistListIn() string {
	if m != nil {
		return m.DistListIn
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShProtocol_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_summary.ospf_sh_protocol_KEYS")
	proto.RegisterType((*OspfShProtocol)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_summary.ospf_sh_protocol")
}

func init() { proto.RegisterFile("ospf_sh_protocol.proto", fileDescriptor_4559d5f25e5c5134) }

var fileDescriptor_4559d5f25e5c5134 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xc9, 0x77, 0xf8, 0xd0, 0xb5, 0x42, 0x5d, 0x50, 0xf6, 0xa4, 0xb1, 0xa7, 0x82, 0x92,
	0x83, 0xad, 0x0a, 0xde, 0x04, 0x0b, 0x96, 0x4a, 0x0f, 0xf5, 0xe4, 0x69, 0x58, 0x93, 0x09, 0x0e,
	0x24, 0xbb, 0x61, 0x67, 0x5b, 0xea, 0x4f, 0xf4, 0x5f, 0xc9, 0xae, 0x49, 0xc4, 0xa2, 0x97, 0x65,
	0x98, 0x79, 0xde, 0x67, 0xd8, 0x11, 0x27, 0x96, 0x9b, 0x12, 0xf8, 0x0d, 0x1a, 0x67, 0xbd, 0xcd,
	0x6d, 0x95, 0xc5, 0x42, 0xe6, 0x39, 0x71, 0x6e, 0x81, 0x2c, 0xc3, 0xd6, 0x01, 0x35, 0x9b, 0x29,
	0x44, 0xd2, 0x36, 0xe8, 0xb2, 0x50, 0x05, 0x2e, 0x47, 0x66, 0xe4, 0xae, 0xca, 0x0a, 0x2c, 0xf5,
	0xba, 0xf2, 0xb0, 0x71, 0xfd, 0x14, 0xc8, 0x94, 0xd6, 0xd5, 0xda, 0x93, 0x35, 0x59, 0xb7, 0x02,
	0x78, 0x5d, 0xd7, 0xda, 0xbd, 0x8f, 0xee, 0xc4, 0xf1, 0xee, 0x7a, 0x58, 0xcc, 0x5e, 0x9e, 0xe5,
	0xb9, 0x18, 0x74, 0x02, 0xa3, 0x6b, 0x54, 0x49, 0x9a, 0x8c, 0xf7, 0x57, 0x07, 0x6d, 0x6f, 0xa9,
	0x6b, 0x1c, 0x7d, 0xfc, 0x13, 0xc3, 0xdd, 0xb0, 0xbc, 0x14, 0xb2, 0x17, 0x39, 0xbb, 0xf6, 0xe8,
	0x80, 0x0a, 0x75, 0x15, 0xd3, 0xc3, 0x6e, 0xb2, 0x8a, 0x83, 0x79, 0x21, 0x2f, 0xc4, 0x51, 0x4f,
	0x17, 0xc4, 0x5e, 0x9b, 0x1c, 0xd5, 0x24, 0x4d, 0xc6, 0x87, 0xdf, 0xf0, 0x43, 0xdb, 0x97, 0x0b,
	0x31, 0xd2, 0x45, 0x4d, 0x86, 0xd8, 0x3b, 0xed, 0x69, 0x83, 0x7d, 0x04, 0xc8, 0x84, 0x45, 0xda,
	0xa1, 0x56, 0xd3, 0x98, 0x3e, 0xfb, 0x49, 0x76, 0x8e, 0x79, 0xe0, 0xee, 0x1d, 0x6a, 0xf9, 0x28,
	0xd2, 0xbf, 0x64, 0xb8, 0xf5, 0xe8, 0x8c, 0xae, 0xd4, 0x75, 0x54, 0x9d, 0xfe, 0xae, 0x9a, 0xb5,
	0x54, 0x7b, 0xa9, 0xaf, 0x3f, 0x18, 0x2e, 0xd5, 0x4d, 0x9a, 0x8c, 0xf7, 0xe2, 0xa5, 0x62, 0x6f,
	0xc9, 0xa5, 0x4c, 0xc5, 0x20, 0xd8, 0xa1, 0x0a, 0x0f, 0x19, 0x75, 0x1b, 0xcf, 0x21, 0x42, 0xef,
	0x89, 0xd8, 0xcf, 0xcd, 0xeb, 0xff, 0x88, 0x4f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xab,
	0xa5, 0xb3, 0x0d, 0x02, 0x00, 0x00,
}
