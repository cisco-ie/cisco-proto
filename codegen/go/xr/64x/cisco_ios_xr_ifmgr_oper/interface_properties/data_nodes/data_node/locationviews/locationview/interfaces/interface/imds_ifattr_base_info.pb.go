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

package cisco_ios_xr_ifmgr_oper_interface_properties_data_nodes_data_node_locationviews_locationview_interfaces_interface

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
	LocationviewName     string   `protobuf:"bytes,2,opt,name=locationview_name,json=locationviewName,proto3" json:"locationview_name,omitempty"`
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

func (m *ImdsIfattrBaseInfo_KEYS) GetLocationviewName() string {
	if m != nil {
		return m.LocationviewName
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
	proto.RegisterType((*ImdsIfattrBaseInfo_KEYS)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.locationviews.locationview.interfaces.interface.imds_ifattr_base_info_KEYS")
	proto.RegisterType((*ImdsIfattrBaseInfo)(nil), "cisco_ios_xr_ifmgr_oper.interface_properties.data_nodes.data_node.locationviews.locationview.interfaces.interface.imds_ifattr_base_info")
}

func init() { proto.RegisterFile("imds_ifattr_base_info.proto", fileDescriptor_558a1bcb3b685904) }

var fileDescriptor_558a1bcb3b685904 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0x65, 0x02, 0x88, 0x4e, 0x9b, 0x92, 0xae, 0x40, 0x5d, 0x0a, 0x48, 0xa1, 0x2a, 0x52,
	0x00, 0x29, 0x87, 0x94, 0xbf, 0x2d, 0x70, 0xe3, 0x80, 0x80, 0x22, 0x25, 0xbd, 0x70, 0x1a, 0xad,
	0xed, 0x4d, 0xbb, 0x92, 0xbd, 0x6b, 0x76, 0xc7, 0x2d, 0x7d, 0x11, 0x1e, 0x92, 0xa7, 0x40, 0x9e,
	0x4d, 0xed, 0x58, 0xea, 0x6d, 0xf6, 0xfb, 0x7d, 0x3b, 0x1e, 0x8f, 0x0d, 0x8f, 0x4d, 0x99, 0x07,
	0x34, 0x4b, 0x45, 0xe4, 0x31, 0x55, 0x41, 0xa3, 0xb1, 0x4b, 0x37, 0xad, 0xbc, 0x23, 0x27, 0x7e,
	0x67, 0x26, 0x64, 0x0e, 0x8d, 0x0b, 0xf8, 0xc7, 0xa3, 0x59, 0x96, 0x67, 0x1e, 0x5d, 0xa5, 0xfd,
	0xd4, 0x58, 0xd2, 0x7e, 0xa9, 0x32, 0x8d, 0x95, 0x6f, 0x00, 0x19, 0x1d, 0xa6, 0xb9, 0x22, 0x85,
	0xd6, 0xe5, 0xeb, 0xe5, 0xb4, 0x70, 0x99, 0x22, 0xe3, 0xec, 0x85, 0xd1, 0x97, 0xa1, 0x77, 0xea,
	0x9a, 0x84, 0xae, 0xdc, 0xff, 0x9b, 0xc0, 0xde, 0x8d, 0x23, 0xe1, 0xb7, 0x2f, 0xbf, 0x16, 0xe2,
	0x00, 0xb6, 0xdb, 0xde, 0x68, 0x55, 0xa9, 0x65, 0x32, 0x4e, 0x26, 0x1b, 0xf3, 0xad, 0x86, 0x9e,
	0xb8, 0x5c, 0x9f, 0xa8, 0x52, 0x8b, 0x57, 0xb0, 0xb3, 0xfe, 0xa4, 0x28, 0xde, 0x62, 0x71, 0xb4,
	0x1e, 0xb0, 0xfc, 0x1c, 0xb6, 0xbb, 0xd7, 0x61, 0x73, 0xc0, 0xe6, 0xb0, 0xa5, 0x8d, 0xb6, 0xff,
	0x6f, 0x00, 0x0f, 0x6f, 0x1c, 0x4c, 0x3c, 0x81, 0x8d, 0x56, 0x95, 0x33, 0xbe, 0xdb, 0x01, 0xf1,
	0x02, 0x46, 0x95, 0xf2, 0xda, 0x12, 0x76, 0xd2, 0x21, 0x4b, 0xf7, 0x23, 0xff, 0xda, 0xaa, 0x02,
	0x6e, 0xd3, 0x55, 0xa5, 0xe5, 0x6b, 0x8e, 0xb9, 0x16, 0x0f, 0xe0, 0x4e, 0x20, 0x45, 0x5a, 0xbe,
	0x61, 0x18, 0x0f, 0xe2, 0x19, 0x6c, 0xa9, 0x8c, 0x6a, 0x55, 0x60, 0x0c, 0xdf, 0x72, 0xb8, 0x19,
	0xd9, 0x82, 0x95, 0xa7, 0x00, 0x85, 0xb1, 0x7a, 0x25, 0xbc, 0x8b, 0x63, 0x35, 0x24, 0xc6, 0x2f,
	0x61, 0x67, 0xd5, 0x61, 0xcd, 0x7a, 0x1f, 0xe7, 0x8a, 0xc1, 0xf7, 0xd6, 0x3d, 0x80, 0xa1, 0xb6,
	0x99, 0xaa, 0x42, 0x5d, 0xf0, 0xea, 0xe4, 0x87, 0xb8, 0xa0, 0x1e, 0x14, 0x47, 0xf0, 0xa8, 0x07,
	0xb0, 0x99, 0x1f, 0x03, 0x79, 0x63, 0xcf, 0xe4, 0x11, 0xdf, 0xd8, 0xed, 0x09, 0xa7, 0x57, 0x95,
	0x5e, 0x70, 0x2c, 0x46, 0x30, 0x28, 0xa9, 0x96, 0xc7, 0xe3, 0x64, 0x32, 0x9c, 0x37, 0xa5, 0x38,
	0x86, 0xbd, 0x50, 0xa7, 0xdd, 0xce, 0xb0, 0xa4, 0x1a, 0xdd, 0x85, 0xf6, 0xe7, 0x5a, 0xe5, 0xf2,
	0x23, 0x8b, 0xbb, 0xa1, 0x4e, 0xdb, 0xed, 0xfd, 0xa0, 0xfa, 0xe7, 0x2a, 0x6e, 0xd6, 0x53, 0xcc,
	0x90, 0xbc, 0xb2, 0xa1, 0x72, 0x9e, 0xe4, 0xa7, 0x71, 0x32, 0xb9, 0x37, 0xdf, 0x2c, 0x66, 0xa7,
	0xd7, 0xa8, 0xf9, 0x68, 0xa9, 0xb2, 0xf9, 0xa5, 0xc9, 0xe9, 0x5c, 0x7e, 0xe6, 0x76, 0x1d, 0x48,
	0xef, 0xf2, 0xff, 0x7f, 0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xf8, 0xd4, 0x6c, 0x1e, 0x03,
	0x00, 0x00,
}
