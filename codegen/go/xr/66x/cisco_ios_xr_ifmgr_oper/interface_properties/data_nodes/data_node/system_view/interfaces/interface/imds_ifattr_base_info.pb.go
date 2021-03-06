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
// source: imds_ifattr_base_info.proto

package cisco_ios_xr_ifmgr_oper_interface_properties_data_nodes_data_node_system_view_interfaces_interface

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

type ImdsIfattrBaseInfo_KEYS struct {
	DataNodeName         string   `protobuf:"bytes,1,opt,name=data_node_name,json=dataNodeName,proto3" json:"data_node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImdsIfattrBaseInfo_KEYS) Reset()         { *m = ImdsIfattrBaseInfo_KEYS{} }
func (m *ImdsIfattrBaseInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ImdsIfattrBaseInfo_KEYS) ProtoMessage()    {}
func (*ImdsIfattrBaseInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_558a1bcb3b685904, []int{0}
}

func (m *ImdsIfattrBaseInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImdsIfattrBaseInfo_KEYS.Unmarshal(m, b)
}
func (m *ImdsIfattrBaseInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImdsIfattrBaseInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ImdsIfattrBaseInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImdsIfattrBaseInfo_KEYS.Merge(m, src)
}
func (m *ImdsIfattrBaseInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ImdsIfattrBaseInfo_KEYS.Size(m)
}
func (m *ImdsIfattrBaseInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ImdsIfattrBaseInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ImdsIfattrBaseInfo_KEYS proto.InternalMessageInfo

func (m *ImdsIfattrBaseInfo_KEYS) GetDataNodeName() string {
	if m != nil {
		return m.DataNodeName
	}
	return ""
}

func (m *ImdsIfattrBaseInfo_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type ImdsIfattrBaseInfo struct {
	Interface               string   `protobuf:"bytes,50,opt,name=interface,proto3" json:"interface,omitempty"`
	ParentInterface         string   `protobuf:"bytes,51,opt,name=parent_interface,json=parentInterface,proto3" json:"parent_interface,omitempty"`
	Type                    string   `protobuf:"bytes,52,opt,name=type,proto3" json:"type,omitempty"`
	State                   string   `protobuf:"bytes,53,opt,name=state,proto3" json:"state,omitempty"`
	ActualState             string   `protobuf:"bytes,54,opt,name=actual_state,json=actualState,proto3" json:"actual_state,omitempty"`
	LineState               string   `protobuf:"bytes,55,opt,name=line_state,json=lineState,proto3" json:"line_state,omitempty"`
	ActualLineState         string   `protobuf:"bytes,56,opt,name=actual_line_state,json=actualLineState,proto3" json:"actual_line_state,omitempty"`
	Encapsulation           string   `protobuf:"bytes,57,opt,name=encapsulation,proto3" json:"encapsulation,omitempty"`
	EncapsulationTypeString string   `protobuf:"bytes,58,opt,name=encapsulation_type_string,json=encapsulationTypeString,proto3" json:"encapsulation_type_string,omitempty"`
	Mtu                     uint32   `protobuf:"varint,59,opt,name=mtu,proto3" json:"mtu,omitempty"`
	SubInterfaceMtuOverhead uint32   `protobuf:"varint,60,opt,name=sub_interface_mtu_overhead,json=subInterfaceMtuOverhead,proto3" json:"sub_interface_mtu_overhead,omitempty"`
	L2Transport             bool     `protobuf:"varint,61,opt,name=l2_transport,json=l2Transport,proto3" json:"l2_transport,omitempty"`
	Bandwidth               uint32   `protobuf:"varint,62,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ImdsIfattrBaseInfo) Reset()         { *m = ImdsIfattrBaseInfo{} }
func (m *ImdsIfattrBaseInfo) String() string { return proto.CompactTextString(m) }
func (*ImdsIfattrBaseInfo) ProtoMessage()    {}
func (*ImdsIfattrBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_558a1bcb3b685904, []int{1}
}

func (m *ImdsIfattrBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImdsIfattrBaseInfo.Unmarshal(m, b)
}
func (m *ImdsIfattrBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImdsIfattrBaseInfo.Marshal(b, m, deterministic)
}
func (m *ImdsIfattrBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImdsIfattrBaseInfo.Merge(m, src)
}
func (m *ImdsIfattrBaseInfo) XXX_Size() int {
	return xxx_messageInfo_ImdsIfattrBaseInfo.Size(m)
}
func (m *ImdsIfattrBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ImdsIfattrBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ImdsIfattrBaseInfo proto.InternalMessageInfo

func (m *ImdsIfattrBaseInfo) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetParentInterface() string {
	if m != nil {
		return m.ParentInterface
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetActualState() string {
	if m != nil {
		return m.ActualState
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetLineState() string {
	if m != nil {
		return m.LineState
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetActualLineState() string {
	if m != nil {
		return m.ActualLineState
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetEncapsulation() string {
	if m != nil {
		return m.Encapsulation
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetEncapsulationTypeString() string {
	if m != nil {
		return m.EncapsulationTypeString
	}
	return ""
}

func (m *ImdsIfattrBaseInfo) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *ImdsIfattrBaseInfo) GetSubInterfaceMtuOverhead() uint32 {
	if m != nil {
		return m.SubInterfaceMtuOverhead
	}
	return 0
}

func (m *ImdsIfattrBaseInfo) GetL2Transport() bool {
	if m != nil {
		return m.L2Transport
	}
	return false
}

func (m *ImdsIfattrBaseInfo) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func init() {
	proto.RegisterType((*ImdsIfattrBaseInfo_KEYS)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.system_view.interfaces.interface.imds_ifattr_base_info_KEYS")
	proto.RegisterType((*ImdsIfattrBaseInfo)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.system_view.interfaces.interface.imds_ifattr_base_info")
}

func init() { proto.RegisterFile("imds_ifattr_base_info.proto", fileDescriptor_558a1bcb3b685904) }

var fileDescriptor_558a1bcb3b685904 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0x15, 0x0a, 0x88, 0x4e, 0x9b, 0x52, 0x56, 0xa0, 0x2e, 0x05, 0xa4, 0x50, 0x15, 0x29,
	0x70, 0xc8, 0x21, 0xe5, 0x6f, 0x0b, 0xdc, 0x38, 0x20, 0xa0, 0x48, 0x49, 0x2f, 0x9c, 0x46, 0x63,
	0x7b, 0xd2, 0xae, 0x64, 0xef, 0x5a, 0xbb, 0xe3, 0x96, 0xbc, 0x32, 0x4f, 0x81, 0xbc, 0x9b, 0xda,
	0x8d, 0xd4, 0xdb, 0xf8, 0xfb, 0x7d, 0x33, 0x1e, 0x8d, 0x0d, 0xcf, 0x4c, 0x55, 0x04, 0x34, 0x0b,
	0x12, 0xf1, 0x98, 0x51, 0x60, 0x34, 0x76, 0xe1, 0x26, 0xb5, 0x77, 0xe2, 0x54, 0x96, 0x9b, 0x90,
	0x3b, 0x34, 0x2e, 0xe0, 0x5f, 0x8f, 0x66, 0x51, 0x9d, 0x7b, 0x74, 0x35, 0xfb, 0x89, 0xb1, 0xc2,
	0x7e, 0x41, 0x39, 0x63, 0xed, 0x5b, 0x20, 0x86, 0xc3, 0xa4, 0x20, 0x21, 0xb4, 0xae, 0xb8, 0x59,
	0x4e, 0xc2, 0x32, 0x08, 0x57, 0x78, 0x69, 0xf8, 0xaa, 0xef, 0x0a, 0x7d, 0x79, 0x60, 0x60, 0xff,
	0xd6, 0x15, 0xf0, 0xc7, 0xb7, 0x3f, 0x73, 0x75, 0x08, 0x3b, 0xdd, 0x2c, 0xb4, 0x54, 0xb1, 0x1e,
	0x8c, 0x06, 0xe3, 0xcd, 0xd9, 0x76, 0x4b, 0x4f, 0x5d, 0xc1, 0xa7, 0x54, 0xb1, 0x7a, 0x05, 0x3b,
	0xfd, 0x46, 0xd1, 0xba, 0x13, 0xad, 0x61, 0x47, 0x5b, 0xed, 0xe0, 0xdf, 0x06, 0x3c, 0xb9, 0xf5,
	0x5d, 0xea, 0x39, 0x6c, 0x76, 0xaa, 0x9e, 0xc6, 0xde, 0x1e, 0xa8, 0xd7, 0xb0, 0x5b, 0x93, 0x67,
	0x2b, 0xd8, 0x4b, 0x47, 0x51, 0x7a, 0x98, 0xf8, 0xf7, 0x4e, 0x55, 0x70, 0x57, 0x96, 0x35, 0xeb,
	0xb7, 0x31, 0x8e, 0xb5, 0x7a, 0x0c, 0xf7, 0x82, 0x90, 0xb0, 0x7e, 0x17, 0x61, 0x7a, 0x50, 0x2f,
	0x61, 0x9b, 0x72, 0x69, 0xa8, 0xc4, 0x14, 0xbe, 0x8f, 0xe1, 0x56, 0x62, 0xf3, 0xa8, 0xbc, 0x00,
	0x28, 0x8d, 0xe5, 0x95, 0xf0, 0x21, 0xad, 0xd5, 0x92, 0x14, 0xbf, 0x81, 0x47, 0xab, 0x09, 0x37,
	0xac, 0x8f, 0x69, 0xaf, 0x14, 0xfc, 0xec, 0xdc, 0x43, 0x18, 0xb2, 0xcd, 0xa9, 0x0e, 0x4d, 0x49,
	0x62, 0x9c, 0xd5, 0x9f, 0xd2, 0x81, 0xd6, 0xa0, 0x3a, 0x86, 0xa7, 0x6b, 0x00, 0xdb, 0xfd, 0x31,
	0x88, 0x37, 0xf6, 0x5c, 0x1f, 0xc7, 0x8e, 0xbd, 0x35, 0xe1, 0x6c, 0x59, 0xf3, 0x3c, 0xc6, 0x6a,
	0x17, 0x36, 0x2a, 0x69, 0xf4, 0xc9, 0x68, 0x30, 0x1e, 0xce, 0xda, 0x52, 0x9d, 0xc0, 0x7e, 0x68,
	0xb2, 0xfe, 0x66, 0x58, 0x49, 0x83, 0xee, 0x92, 0xfd, 0x05, 0x53, 0xa1, 0x3f, 0x47, 0x71, 0x2f,
	0x34, 0x59, 0x77, 0xbd, 0x5f, 0xd2, 0xfc, 0x5e, 0xc5, 0xed, 0x79, 0xca, 0x29, 0x8a, 0x27, 0x1b,
	0x6a, 0xe7, 0x45, 0x7f, 0x19, 0x0d, 0xc6, 0x0f, 0x66, 0x5b, 0xe5, 0xf4, 0xec, 0x1a, 0xb5, 0x1f,
	0x2d, 0x23, 0x5b, 0x5c, 0x99, 0x42, 0x2e, 0xf4, 0xd7, 0x38, 0xae, 0x07, 0xd9, 0xfd, 0xf8, 0x0b,
	0x1f, 0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x88, 0xc0, 0x49, 0xaa, 0xe1, 0x02, 0x00, 0x00,
}
