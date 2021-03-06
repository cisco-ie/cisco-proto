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
// source: ospf_sh_proto_area.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_process_information_protocol_areas_protocol_area

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

type OspfShProtoArea_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShProtoArea_KEYS) Reset()         { *m = OspfShProtoArea_KEYS{} }
func (m *OspfShProtoArea_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoArea_KEYS) ProtoMessage()    {}
func (*OspfShProtoArea_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2729fb732dd66d, []int{0}
}

func (m *OspfShProtoArea_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Unmarshal(m, b)
}
func (m *OspfShProtoArea_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShProtoArea_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoArea_KEYS.Merge(m, src)
}
func (m *OspfShProtoArea_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoArea_KEYS.Size(m)
}
func (m *OspfShProtoArea_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoArea_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoArea_KEYS proto.InternalMessageInfo

func (m *OspfShProtoArea_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShProtoArea_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShProtoArea_KEYS) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type OspfShProtoIntf struct {
	ProtocolInterfaceName       string   `protobuf:"bytes,1,opt,name=protocol_interface_name,json=protocolInterfaceName,proto3" json:"protocol_interface_name,omitempty"`
	ProtocolAuthenticationType  string   `protobuf:"bytes,2,opt,name=protocol_authentication_type,json=protocolAuthenticationType,proto3" json:"protocol_authentication_type,omitempty"`
	ProtocolInterfaceDistListIn string   `protobuf:"bytes,3,opt,name=protocol_interface_dist_list_in,json=protocolInterfaceDistListIn,proto3" json:"protocol_interface_dist_list_in,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *OspfShProtoIntf) Reset()         { *m = OspfShProtoIntf{} }
func (m *OspfShProtoIntf) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoIntf) ProtoMessage()    {}
func (*OspfShProtoIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2729fb732dd66d, []int{1}
}

func (m *OspfShProtoIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoIntf.Unmarshal(m, b)
}
func (m *OspfShProtoIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoIntf.Marshal(b, m, deterministic)
}
func (m *OspfShProtoIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoIntf.Merge(m, src)
}
func (m *OspfShProtoIntf) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoIntf.Size(m)
}
func (m *OspfShProtoIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoIntf.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoIntf proto.InternalMessageInfo

func (m *OspfShProtoIntf) GetProtocolInterfaceName() string {
	if m != nil {
		return m.ProtocolInterfaceName
	}
	return ""
}

func (m *OspfShProtoIntf) GetProtocolAuthenticationType() string {
	if m != nil {
		return m.ProtocolAuthenticationType
	}
	return ""
}

func (m *OspfShProtoIntf) GetProtocolInterfaceDistListIn() string {
	if m != nil {
		return m.ProtocolInterfaceDistListIn
	}
	return ""
}

type OspfShProtoArea struct {
	ProtcolArea            string             `protobuf:"bytes,50,opt,name=protcol_area,json=protcolArea,proto3" json:"protcol_area,omitempty"`
	ProtocolMpls           bool               `protobuf:"varint,51,opt,name=protocol_mpls,json=protocolMpls,proto3" json:"protocol_mpls,omitempty"`
	ProtocolAreaDistListIn string             `protobuf:"bytes,52,opt,name=protocol_area_dist_list_in,json=protocolAreaDistListIn,proto3" json:"protocol_area_dist_list_in,omitempty"`
	ProtocolInterface      []*OspfShProtoIntf `protobuf:"bytes,53,rep,name=protocol_interface,json=protocolInterface,proto3" json:"protocol_interface,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShProtoArea) Reset()         { *m = OspfShProtoArea{} }
func (m *OspfShProtoArea) String() string { return proto.CompactTextString(m) }
func (*OspfShProtoArea) ProtoMessage()    {}
func (*OspfShProtoArea) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f2729fb732dd66d, []int{2}
}

func (m *OspfShProtoArea) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShProtoArea.Unmarshal(m, b)
}
func (m *OspfShProtoArea) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShProtoArea.Marshal(b, m, deterministic)
}
func (m *OspfShProtoArea) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShProtoArea.Merge(m, src)
}
func (m *OspfShProtoArea) XXX_Size() int {
	return xxx_messageInfo_OspfShProtoArea.Size(m)
}
func (m *OspfShProtoArea) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShProtoArea.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShProtoArea proto.InternalMessageInfo

func (m *OspfShProtoArea) GetProtcolArea() string {
	if m != nil {
		return m.ProtcolArea
	}
	return ""
}

func (m *OspfShProtoArea) GetProtocolMpls() bool {
	if m != nil {
		return m.ProtocolMpls
	}
	return false
}

func (m *OspfShProtoArea) GetProtocolAreaDistListIn() string {
	if m != nil {
		return m.ProtocolAreaDistListIn
	}
	return ""
}

func (m *OspfShProtoArea) GetProtocolInterface() []*OspfShProtoIntf {
	if m != nil {
		return m.ProtocolInterface
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShProtoArea_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_area_KEYS")
	proto.RegisterType((*OspfShProtoIntf)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_intf")
	proto.RegisterType((*OspfShProtoArea)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.process_information.protocol_areas.protocol_area.ospf_sh_proto_area")
}

func init() { proto.RegisterFile("ospf_sh_proto_area.proto", fileDescriptor_2f2729fb732dd66d) }

var fileDescriptor_2f2729fb732dd66d = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x52, 0x3d, 0x6f, 0xe2, 0x30,
	0x18, 0x56, 0x40, 0x82, 0xc3, 0xc0, 0x70, 0x96, 0xee, 0xb0, 0xb8, 0x93, 0x2e, 0xc7, 0x2d, 0x99,
	0x32, 0x00, 0x77, 0xc3, 0x4d, 0x87, 0xc4, 0x0d, 0xa8, 0x1f, 0x43, 0xda, 0xa5, 0x93, 0xe5, 0x26,
	0x8e, 0xb0, 0x94, 0xd8, 0x96, 0x5f, 0x83, 0xca, 0xdc, 0x9f, 0xd4, 0xff, 0xd1, 0xa9, 0x3f, 0xa8,
	0x8a, 0x8b, 0x43, 0xd3, 0x74, 0xee, 0x82, 0xde, 0x8f, 0x87, 0xf7, 0xf9, 0x70, 0x10, 0x51, 0xa0,
	0x73, 0x0a, 0x5b, 0xaa, 0x8d, 0xb2, 0x8a, 0x32, 0xc3, 0x59, 0xec, 0x4a, 0xac, 0x52, 0x01, 0xa9,
	0xa2, 0x42, 0x01, 0xbd, 0x33, 0x54, 0xe8, 0xfd, 0x92, 0x3a, 0xac, 0xd2, 0xdc, 0xc4, 0x55, 0x55,
	0xe1, 0x52, 0x0e, 0xc0, 0xc1, 0x57, 0x71, 0xc6, 0x73, 0xb6, 0x2b, 0x2c, 0xdd, 0x9b, 0x7a, 0x4b,
	0x85, 0xcc, 0x95, 0x29, 0x99, 0x15, 0x4a, 0xbe, 0x5c, 0x4e, 0x55, 0xe1, 0x78, 0xa0, 0xd9, 0xce,
	0x14, 0x9a, 0xb4, 0xc5, 0xd0, 0xb3, 0xff, 0x37, 0x57, 0xf8, 0x27, 0x1a, 0xf9, 0x73, 0x92, 0x95,
	0x9c, 0x04, 0x61, 0x10, 0x0d, 0x92, 0xe1, 0x71, 0x76, 0xc9, 0x4a, 0x8e, 0x27, 0xa8, 0xef, 0xf0,
	0x22, 0x23, 0x9d, 0x30, 0x88, 0xc6, 0x49, 0xaf, 0x6a, 0x37, 0x19, 0x26, 0xa8, 0xcf, 0xb2, 0xcc,
	0x70, 0x00, 0xd2, 0x75, 0x7f, 0xf3, 0xed, 0xec, 0x29, 0x40, 0xb8, 0xc9, 0x28, 0xa4, 0xcd, 0xf1,
	0x1f, 0x34, 0xa9, 0x85, 0x09, 0x69, 0xb9, 0xc9, 0x59, 0xca, 0x5f, 0xf3, 0x7e, 0xf1, 0xeb, 0x8d,
	0xdf, 0x3a, 0x05, 0xff, 0xd0, 0xf7, 0x93, 0xa1, 0x9d, 0xdd, 0x72, 0x69, 0x45, 0xea, 0x7c, 0x53,
	0x7b, 0xd0, 0xdc, 0xc9, 0x1a, 0x24, 0x53, 0x8f, 0x59, 0x35, 0x20, 0xd7, 0x07, 0xcd, 0xf1, 0x1a,
	0xfd, 0x78, 0x87, 0x39, 0x13, 0x60, 0x69, 0x51, 0xfd, 0x08, 0x79, 0xb4, 0xf0, 0xad, 0xa5, 0x60,
	0x2d, 0xc0, 0x9e, 0x0b, 0xb0, 0x1b, 0x39, 0x7b, 0xec, 0xbc, 0xb5, 0x55, 0x25, 0x71, 0xcc, 0xd0,
	0xfa, 0xb8, 0xc9, 0xbc, 0xce, 0xb0, 0x9a, 0xad, 0x2a, 0xc8, 0x2f, 0x34, 0xae, 0xf9, 0x4b, 0x5d,
	0x00, 0x59, 0x84, 0x41, 0xf4, 0x29, 0x19, 0xf9, 0xe1, 0x85, 0x2e, 0x00, 0xff, 0x45, 0xd3, 0xc6,
	0xbb, 0x35, 0xf5, 0x2d, 0xdd, 0xd5, 0xaf, 0xb5, 0x49, 0xc3, 0xd9, 0x49, 0x1a, 0x7e, 0x08, 0x10,
	0x6e, 0x3b, 0x24, 0xbf, 0xc3, 0x6e, 0x34, 0x9c, 0xdf, 0x07, 0xf1, 0x07, 0x7f, 0x72, 0x71, 0xfb,
	0xf5, 0x93, 0xcf, 0xad, 0x68, 0x6f, 0x7b, 0x6e, 0xb4, 0x78, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x02,
	0xb7, 0x6b, 0x7a, 0x2c, 0x03, 0x00, 0x00,
}
