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
// source: mpls_lsd_lbl_range.proto

package cisco_ios_xr_mpls_lsd_oper_mpls_lsd_nodes_mpls_lsd_node_label_range

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

type MplsLsdLblRange_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdLblRange_KEYS) Reset()         { *m = MplsLsdLblRange_KEYS{} }
func (m *MplsLsdLblRange_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLsdLblRange_KEYS) ProtoMessage()    {}
func (*MplsLsdLblRange_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53d40e264282a28, []int{0}
}

func (m *MplsLsdLblRange_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdLblRange_KEYS.Unmarshal(m, b)
}
func (m *MplsLsdLblRange_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdLblRange_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLsdLblRange_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdLblRange_KEYS.Merge(m, src)
}
func (m *MplsLsdLblRange_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLsdLblRange_KEYS.Size(m)
}
func (m *MplsLsdLblRange_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdLblRange_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdLblRange_KEYS proto.InternalMessageInfo

func (m *MplsLsdLblRange_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type MplsLsdLblRange struct {
	MinDynamicLabelValue uint32   `protobuf:"varint,50,opt,name=min_dynamic_label_value,json=minDynamicLabelValue,proto3" json:"min_dynamic_label_value,omitempty"`
	MaxDynamicLabelValue uint32   `protobuf:"varint,51,opt,name=max_dynamic_label_value,json=maxDynamicLabelValue,proto3" json:"max_dynamic_label_value,omitempty"`
	MinStaticLabelValue  uint32   `protobuf:"varint,52,opt,name=min_static_label_value,json=minStaticLabelValue,proto3" json:"min_static_label_value,omitempty"`
	MaxStaticLabelValue  uint32   `protobuf:"varint,53,opt,name=max_static_label_value,json=maxStaticLabelValue,proto3" json:"max_static_label_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLsdLblRange) Reset()         { *m = MplsLsdLblRange{} }
func (m *MplsLsdLblRange) String() string { return proto.CompactTextString(m) }
func (*MplsLsdLblRange) ProtoMessage()    {}
func (*MplsLsdLblRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a53d40e264282a28, []int{1}
}

func (m *MplsLsdLblRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLsdLblRange.Unmarshal(m, b)
}
func (m *MplsLsdLblRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLsdLblRange.Marshal(b, m, deterministic)
}
func (m *MplsLsdLblRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLsdLblRange.Merge(m, src)
}
func (m *MplsLsdLblRange) XXX_Size() int {
	return xxx_messageInfo_MplsLsdLblRange.Size(m)
}
func (m *MplsLsdLblRange) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLsdLblRange.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLsdLblRange proto.InternalMessageInfo

func (m *MplsLsdLblRange) GetMinDynamicLabelValue() uint32 {
	if m != nil {
		return m.MinDynamicLabelValue
	}
	return 0
}

func (m *MplsLsdLblRange) GetMaxDynamicLabelValue() uint32 {
	if m != nil {
		return m.MaxDynamicLabelValue
	}
	return 0
}

func (m *MplsLsdLblRange) GetMinStaticLabelValue() uint32 {
	if m != nil {
		return m.MinStaticLabelValue
	}
	return 0
}

func (m *MplsLsdLblRange) GetMaxStaticLabelValue() uint32 {
	if m != nil {
		return m.MaxStaticLabelValue
	}
	return 0
}

func init() {
	proto.RegisterType((*MplsLsdLblRange_KEYS)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.label_range.mpls_lsd_lbl_range_KEYS")
	proto.RegisterType((*MplsLsdLblRange)(nil), "cisco_ios_xr_mpls_lsd_oper.mpls_lsd_nodes.mpls_lsd_node.label_range.mpls_lsd_lbl_range")
}

func init() { proto.RegisterFile("mpls_lsd_lbl_range.proto", fileDescriptor_a53d40e264282a28) }

var fileDescriptor_a53d40e264282a28 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0x3f, 0x4b, 0xc6, 0x30,
	0x10, 0x06, 0x70, 0xb2, 0x88, 0x6f, 0xc0, 0x25, 0x8a, 0x2d, 0xb8, 0x94, 0x4e, 0x9d, 0x3a, 0x58,
	0xeb, 0x17, 0x50, 0x27, 0xc5, 0xa1, 0x05, 0xc1, 0xe9, 0xb8, 0xb6, 0x41, 0x02, 0xf9, 0x53, 0x92,
	0x2a, 0xf1, 0x3b, 0xfb, 0x21, 0x24, 0x69, 0x29, 0xbc, 0x6d, 0xc7, 0xcb, 0x73, 0xbf, 0x3c, 0x70,
	0x34, 0x55, 0xa3, 0x74, 0x20, 0xdd, 0x00, 0xb2, 0x93, 0x60, 0x51, 0x7f, 0xf1, 0x72, 0xb4, 0x66,
	0x32, 0xec, 0xa9, 0x17, 0xae, 0x37, 0x20, 0x8c, 0x03, 0x6f, 0x61, 0x5d, 0x33, 0x23, 0xb7, 0xe5,
	0x3a, 0x69, 0x33, 0x70, 0x77, 0x3e, 0x96, 0x12, 0x3b, 0xbe, 0x7c, 0x95, 0x3f, 0xd2, 0x64, 0x5f,
	0x00, 0xaf, 0x2f, 0x9f, 0x2d, 0xbb, 0xa3, 0xa7, 0xb0, 0x0e, 0x1a, 0x15, 0x4f, 0x49, 0x46, 0x8a,
	0x53, 0x73, 0x19, 0x1e, 0xde, 0x51, 0xf1, 0xfc, 0x8f, 0x50, 0xb6, 0x87, 0xac, 0xa6, 0x89, 0x12,
	0x1a, 0x86, 0x5f, 0x8d, 0x4a, 0xf4, 0x30, 0x37, 0xfd, 0xa0, 0xfc, 0xe6, 0xe9, 0x7d, 0x46, 0x8a,
	0xab, 0xe6, 0x46, 0x09, 0xfd, 0x3c, 0xa7, 0x6f, 0x21, 0xfc, 0x08, 0x59, 0x64, 0xe8, 0x0f, 0x59,
	0xb5, 0x30, 0xf4, 0x7b, 0x56, 0xd1, 0xdb, 0xd0, 0xe6, 0x26, 0x9c, 0x36, 0xea, 0x21, 0xaa, 0x6b,
	0x25, 0x74, 0x1b, 0xc3, 0x0d, 0x42, 0x7f, 0x84, 0xea, 0x05, 0xa1, 0xdf, 0xa2, 0xee, 0x22, 0x9e,
	0xbc, 0xfa, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x48, 0xa3, 0x6d, 0xc5, 0x8e, 0x01, 0x00, 0x00,
}
