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
// source: rsi_srlg_intf_brief_show.proto

package cisco_ios_xr_infra_rsi_oper_srlg_nodes_node_interfaces_interface

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

type RsiSrlgIntfBriefShow_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgIntfBriefShow_KEYS) Reset()         { *m = RsiSrlgIntfBriefShow_KEYS{} }
func (m *RsiSrlgIntfBriefShow_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgIntfBriefShow_KEYS) ProtoMessage()    {}
func (*RsiSrlgIntfBriefShow_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20df4511db4384c, []int{0}
}

func (m *RsiSrlgIntfBriefShow_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS.Unmarshal(m, b)
}
func (m *RsiSrlgIntfBriefShow_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS.Marshal(b, m, deterministic)
}
func (m *RsiSrlgIntfBriefShow_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS.Merge(m, src)
}
func (m *RsiSrlgIntfBriefShow_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS.Size(m)
}
func (m *RsiSrlgIntfBriefShow_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgIntfBriefShow_KEYS proto.InternalMessageInfo

func (m *RsiSrlgIntfBriefShow_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RsiSrlgIntfBriefShow_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type RsiSrlgIntfBriefShow struct {
	InterfaceNameXr      string   `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	ValueCount           uint32   `protobuf:"varint,51,opt,name=value_count,json=valueCount,proto3" json:"value_count,omitempty"`
	Registrations        uint32   `protobuf:"varint,52,opt,name=registrations,proto3" json:"registrations,omitempty"`
	SrlgValue            []uint32 `protobuf:"varint,53,rep,packed,name=srlg_value,json=srlgValue,proto3" json:"srlg_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiSrlgIntfBriefShow) Reset()         { *m = RsiSrlgIntfBriefShow{} }
func (m *RsiSrlgIntfBriefShow) String() string { return proto.CompactTextString(m) }
func (*RsiSrlgIntfBriefShow) ProtoMessage()    {}
func (*RsiSrlgIntfBriefShow) Descriptor() ([]byte, []int) {
	return fileDescriptor_a20df4511db4384c, []int{1}
}

func (m *RsiSrlgIntfBriefShow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiSrlgIntfBriefShow.Unmarshal(m, b)
}
func (m *RsiSrlgIntfBriefShow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiSrlgIntfBriefShow.Marshal(b, m, deterministic)
}
func (m *RsiSrlgIntfBriefShow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiSrlgIntfBriefShow.Merge(m, src)
}
func (m *RsiSrlgIntfBriefShow) XXX_Size() int {
	return xxx_messageInfo_RsiSrlgIntfBriefShow.Size(m)
}
func (m *RsiSrlgIntfBriefShow) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiSrlgIntfBriefShow.DiscardUnknown(m)
}

var xxx_messageInfo_RsiSrlgIntfBriefShow proto.InternalMessageInfo

func (m *RsiSrlgIntfBriefShow) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *RsiSrlgIntfBriefShow) GetValueCount() uint32 {
	if m != nil {
		return m.ValueCount
	}
	return 0
}

func (m *RsiSrlgIntfBriefShow) GetRegistrations() uint32 {
	if m != nil {
		return m.Registrations
	}
	return 0
}

func (m *RsiSrlgIntfBriefShow) GetSrlgValue() []uint32 {
	if m != nil {
		return m.SrlgValue
	}
	return nil
}

func init() {
	proto.RegisterType((*RsiSrlgIntfBriefShow_KEYS)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.nodes.node.interfaces.interface.rsi_srlg_intf_brief_show_KEYS")
	proto.RegisterType((*RsiSrlgIntfBriefShow)(nil), "cisco_ios_xr_infra_rsi_oper.srlg.nodes.node.interfaces.interface.rsi_srlg_intf_brief_show")
}

func init() { proto.RegisterFile("rsi_srlg_intf_brief_show.proto", fileDescriptor_a20df4511db4384c) }

var fileDescriptor_a20df4511db4384c = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x59, 0x05, 0x71, 0x47, 0x56, 0x31, 0xa7, 0x80, 0x54, 0x97, 0xa2, 0xb0, 0x78, 0xd8,
	0x83, 0xd5, 0xbb, 0x20, 0x9e, 0x04, 0x0f, 0x2b, 0x88, 0x9e, 0x86, 0x34, 0x66, 0x6b, 0xa0, 0xcd,
	0x94, 0x99, 0x54, 0xfb, 0xa3, 0xfc, 0x91, 0x92, 0x08, 0x95, 0x3d, 0xec, 0x25, 0x84, 0x8f, 0xf7,
	0xbe, 0x17, 0x02, 0xe7, 0x2c, 0x1e, 0x85, 0x97, 0x0b, 0xf4, 0x21, 0xf6, 0x38, 0x67, 0xef, 0x7a,
	0x94, 0x4f, 0xfa, 0x6e, 0xd7, 0x4c, 0x91, 0xd4, 0xbd, 0xf5, 0x62, 0x09, 0x3d, 0x09, 0x6e, 0x19,
	0x7d, 0xe8, 0xd9, 0x60, 0xaa, 0xd0, 0xda, 0x71, 0x9b, 0x7a, 0x6d, 0xa0, 0x0f, 0x27, 0xf9, 0x6c,
	0x7d, 0x88, 0x8e, 0x7b, 0x63, 0x9d, 0xfc, 0x5f, 0xa7, 0x16, 0x26, 0x63, 0x1b, 0xf8, 0xf4, 0xf8,
	0xfe, 0xa2, 0xce, 0xa0, 0x4c, 0x6d, 0x0c, 0x66, 0xe5, 0x74, 0x51, 0x17, 0x4d, 0xd9, 0x1d, 0x26,
	0xf0, 0x6c, 0x56, 0x4e, 0x5d, 0xc1, 0xf1, 0x4e, 0xf5, 0x97, 0xd8, 0xcb, 0x89, 0x6a, 0x47, 0x53,
	0x6c, 0xfa, 0x53, 0x80, 0x1e, 0x5b, 0x51, 0xd7, 0x70, 0x3a, 0x74, 0xe0, 0x96, 0xf5, 0x4d, 0xd6,
	0x9c, 0x0c, 0x34, 0x6f, 0xac, 0x2e, 0xe0, 0xe8, 0xcb, 0x2c, 0x37, 0x0e, 0x2d, 0x6d, 0x42, 0xd4,
	0xb3, 0xba, 0x68, 0xaa, 0x0e, 0x32, 0x7a, 0x48, 0x44, 0x5d, 0x42, 0xc5, 0x6e, 0xe1, 0x25, 0xb2,
	0x89, 0x9e, 0x82, 0xe8, 0xdb, 0x1c, 0x19, 0x42, 0x35, 0x01, 0xc8, 0x4f, 0xc9, 0x45, 0x7d, 0x57,
	0xef, 0x37, 0x55, 0x57, 0x26, 0xf2, 0x9a, 0xc0, 0xfc, 0x20, 0x7f, 0xee, 0xec, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x4a, 0x12, 0x35, 0x7e, 0x01, 0x00, 0x00,
}
