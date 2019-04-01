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

package cisco_ios_xr_ifmgr_oper_interface_properties_data_nodes_data_node_pq_node_locations_pq_node_location_interfaces_interface

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
	PqNodeName           string   `protobuf:"bytes,2,opt,name=pq_node_name,json=pqNodeName,proto3" json:"pq_node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *ImdsIfattrBaseInfo_KEYS) GetPqNodeName() string {
	if m != nil {
		return m.PqNodeName
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
	proto.RegisterType((*ImdsIfattrBaseInfo_KEYS)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.pq_node_locations.pq_node_location.interfaces.interface.imds_ifattr_base_info_KEYS")
	proto.RegisterType((*ImdsIfattrBaseInfo)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.pq_node_locations.pq_node_location.interfaces.interface.imds_ifattr_base_info")
}

func init() { proto.RegisterFile("imds_ifattr_base_info.proto", fileDescriptor_558a1bcb3b685904) }

var fileDescriptor_558a1bcb3b685904 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0x87, 0x15, 0x0a, 0x88, 0x4e, 0x77, 0x4b, 0xb1, 0x40, 0x35, 0x05, 0xa4, 0xa5, 0x2a, 0xd2,
	0xc2, 0x61, 0x0f, 0x5b, 0xfe, 0xb6, 0xc0, 0x8d, 0x03, 0x02, 0x8a, 0xb4, 0xdb, 0x0b, 0xa7, 0x91,
	0x93, 0x78, 0x5b, 0x4b, 0x89, 0xed, 0xda, 0x13, 0x60, 0x9f, 0x81, 0x37, 0xe5, 0x29, 0x50, 0xc6,
	0xdb, 0xa4, 0x2b, 0x7a, 0xca, 0xe4, 0xfb, 0x7d, 0xb6, 0x27, 0xe3, 0xc0, 0x23, 0x53, 0x97, 0x11,
	0xcd, 0x42, 0x11, 0x05, 0xcc, 0x55, 0xd4, 0x68, 0xec, 0xc2, 0x4d, 0x7c, 0x70, 0xe4, 0xc4, 0xb2,
	0x30, 0xb1, 0x70, 0x68, 0x5c, 0xc4, 0xdf, 0x01, 0xcd, 0xa2, 0x3e, 0x0b, 0xe8, 0xbc, 0x0e, 0x13,
	0x63, 0x49, 0x87, 0x85, 0x2a, 0x34, 0xfa, 0xd0, 0x02, 0x32, 0x3a, 0x4e, 0x4a, 0x45, 0x0a, 0xad,
	0x2b, 0xaf, 0x96, 0x13, 0x7f, 0xc1, 0x4f, 0xac, 0x5c, 0xa1, 0xc8, 0x38, 0x1b, 0xff, 0x23, 0xfd,
	0x66, 0xb1, 0x2f, 0xf7, 0xff, 0x64, 0xb0, 0x77, 0x6d, 0x6b, 0xf8, 0xe5, 0xd3, 0x8f, 0xb9, 0x38,
	0x80, 0xed, 0xee, 0x0c, 0xb4, 0xaa, 0xd6, 0x32, 0x1b, 0x65, 0xe3, 0xcd, 0xd9, 0xa0, 0xa5, 0x27,
	0xae, 0xd4, 0x27, 0xaa, 0xd6, 0x62, 0x04, 0x83, 0xcb, 0xd3, 0xd8, 0xb9, 0xc1, 0x0e, 0xf8, 0x8b,
	0xce, 0x78, 0x06, 0xdb, 0xfd, 0xb7, 0xb0, 0xb3, 0xc1, 0xce, 0xb0, 0xa3, 0xad, 0xb6, 0xff, 0x77,
	0x03, 0x1e, 0x5c, 0xdb, 0x8d, 0x78, 0x0c, 0x9b, 0x9d, 0x2a, 0xa7, 0xbc, 0xb6, 0x07, 0xe2, 0x39,
	0xec, 0x78, 0x15, 0xb4, 0x25, 0xec, 0xa5, 0x43, 0x96, 0xee, 0x26, 0xfe, 0xb9, 0x53, 0x05, 0xdc,
	0xa4, 0xa5, 0xd7, 0xf2, 0x25, 0xc7, 0x5c, 0x8b, 0xfb, 0x70, 0x2b, 0x92, 0x22, 0x2d, 0x5f, 0x31,
	0x4c, 0x2f, 0xe2, 0x29, 0x0c, 0x54, 0x41, 0x8d, 0xaa, 0x30, 0x85, 0xaf, 0x39, 0xdc, 0x4a, 0x6c,
	0xce, 0xca, 0x13, 0x80, 0xca, 0x58, 0xbd, 0x12, 0xde, 0xa4, 0xb6, 0x5a, 0x92, 0xe2, 0x17, 0x70,
	0x6f, 0xb5, 0xc3, 0x15, 0xeb, 0x6d, 0xea, 0x2b, 0x05, 0x5f, 0x3b, 0xf7, 0x00, 0x86, 0xda, 0x16,
	0xca, 0xc7, 0xa6, 0xe2, 0xeb, 0x92, 0xef, 0xd2, 0x80, 0xd6, 0xa0, 0x38, 0x82, 0x87, 0x6b, 0x00,
	0xdb, 0xfe, 0x31, 0x52, 0x30, 0xf6, 0x4c, 0x1e, 0xf1, 0x8a, 0xdd, 0x35, 0xe1, 0x74, 0xe9, 0xf5,
	0x9c, 0x63, 0xb1, 0x03, 0x1b, 0x35, 0x35, 0xf2, 0x78, 0x94, 0x8d, 0x87, 0xb3, 0xb6, 0x14, 0xc7,
	0xb0, 0x17, 0x9b, 0xbc, 0x9f, 0x19, 0xd6, 0xd4, 0xa0, 0xfb, 0xa9, 0xc3, 0xb9, 0x56, 0xa5, 0x7c,
	0xcf, 0xe2, 0x6e, 0x6c, 0xf2, 0x6e, 0x7a, 0xdf, 0xa8, 0xf9, 0xbe, 0x8a, 0xdb, 0xf1, 0x54, 0x53,
	0xa4, 0xa0, 0x6c, 0xf4, 0x2e, 0x90, 0xfc, 0x30, 0xca, 0xc6, 0x77, 0x66, 0x5b, 0xd5, 0xf4, 0xf4,
	0x12, 0xb5, 0x97, 0x96, 0x2b, 0x5b, 0xfe, 0x32, 0x25, 0x9d, 0xcb, 0x8f, 0xbc, 0x5d, 0x0f, 0xf2,
	0xdb, 0xfc, 0xf3, 0x1f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x81, 0xbe, 0x6b, 0x1b, 0x03,
	0x00, 0x00,
}
