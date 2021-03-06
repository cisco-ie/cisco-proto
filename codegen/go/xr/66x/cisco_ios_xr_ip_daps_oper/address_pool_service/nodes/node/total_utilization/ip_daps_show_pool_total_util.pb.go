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
// source: ip_daps_show_pool_total_util.proto

package cisco_ios_xr_ip_daps_oper_address_pool_service_nodes_node_total_utilization

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

type IpDapsShowPoolTotalUtil_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpDapsShowPoolTotalUtil_KEYS) Reset()         { *m = IpDapsShowPoolTotalUtil_KEYS{} }
func (m *IpDapsShowPoolTotalUtil_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpDapsShowPoolTotalUtil_KEYS) ProtoMessage()    {}
func (*IpDapsShowPoolTotalUtil_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e1d15e256dc8f, []int{0}
}

func (m *IpDapsShowPoolTotalUtil_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS.Unmarshal(m, b)
}
func (m *IpDapsShowPoolTotalUtil_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS.Marshal(b, m, deterministic)
}
func (m *IpDapsShowPoolTotalUtil_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS.Merge(m, src)
}
func (m *IpDapsShowPoolTotalUtil_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS.Size(m)
}
func (m *IpDapsShowPoolTotalUtil_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpDapsShowPoolTotalUtil_KEYS proto.InternalMessageInfo

func (m *IpDapsShowPoolTotalUtil_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type IpDapsShowPoolTotalUtil struct {
	TotalUtilizationHighMark uint32   `protobuf:"varint,50,opt,name=total_utilization_high_mark,json=totalUtilizationHighMark,proto3" json:"total_utilization_high_mark,omitempty"`
	TotalUtilizationLowMark  uint32   `protobuf:"varint,51,opt,name=total_utilization_low_mark,json=totalUtilizationLowMark,proto3" json:"total_utilization_low_mark,omitempty"`
	CurrentTotalUtilization  uint32   `protobuf:"varint,52,opt,name=current_total_utilization,json=currentTotalUtilization,proto3" json:"current_total_utilization,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *IpDapsShowPoolTotalUtil) Reset()         { *m = IpDapsShowPoolTotalUtil{} }
func (m *IpDapsShowPoolTotalUtil) String() string { return proto.CompactTextString(m) }
func (*IpDapsShowPoolTotalUtil) ProtoMessage()    {}
func (*IpDapsShowPoolTotalUtil) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6e1d15e256dc8f, []int{1}
}

func (m *IpDapsShowPoolTotalUtil) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil.Unmarshal(m, b)
}
func (m *IpDapsShowPoolTotalUtil) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil.Marshal(b, m, deterministic)
}
func (m *IpDapsShowPoolTotalUtil) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpDapsShowPoolTotalUtil.Merge(m, src)
}
func (m *IpDapsShowPoolTotalUtil) XXX_Size() int {
	return xxx_messageInfo_IpDapsShowPoolTotalUtil.Size(m)
}
func (m *IpDapsShowPoolTotalUtil) XXX_DiscardUnknown() {
	xxx_messageInfo_IpDapsShowPoolTotalUtil.DiscardUnknown(m)
}

var xxx_messageInfo_IpDapsShowPoolTotalUtil proto.InternalMessageInfo

func (m *IpDapsShowPoolTotalUtil) GetTotalUtilizationHighMark() uint32 {
	if m != nil {
		return m.TotalUtilizationHighMark
	}
	return 0
}

func (m *IpDapsShowPoolTotalUtil) GetTotalUtilizationLowMark() uint32 {
	if m != nil {
		return m.TotalUtilizationLowMark
	}
	return 0
}

func (m *IpDapsShowPoolTotalUtil) GetCurrentTotalUtilization() uint32 {
	if m != nil {
		return m.CurrentTotalUtilization
	}
	return 0
}

func init() {
	proto.RegisterType((*IpDapsShowPoolTotalUtil_KEYS)(nil), "cisco_ios_xr_ip_daps_oper.address_pool_service.nodes.node.total_utilization.ip_daps_show_pool_total_util_KEYS")
	proto.RegisterType((*IpDapsShowPoolTotalUtil)(nil), "cisco_ios_xr_ip_daps_oper.address_pool_service.nodes.node.total_utilization.ip_daps_show_pool_total_util")
}

func init() { proto.RegisterFile("ip_daps_show_pool_total_util.proto", fileDescriptor_4d6e1d15e256dc8f) }

var fileDescriptor_4d6e1d15e256dc8f = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xb1, 0x4b, 0xc3, 0x40,
	0x14, 0x06, 0x70, 0xb2, 0x88, 0x3d, 0x70, 0xc9, 0x62, 0xb4, 0x0e, 0x35, 0x53, 0xa7, 0x0c, 0xd6,
	0x49, 0x11, 0x5c, 0x04, 0xa1, 0xea, 0x50, 0x75, 0x70, 0x7a, 0x9c, 0xc9, 0xa3, 0x79, 0x34, 0xb9,
	0x77, 0xbc, 0x77, 0x35, 0xe2, 0x1f, 0xea, 0xdf, 0x23, 0xbd, 0x58, 0x0a, 0x0d, 0x74, 0xb9, 0xe1,
	0xf8, 0x7e, 0xdf, 0x07, 0xcf, 0xe4, 0xe4, 0xa1, 0xb2, 0x5e, 0x41, 0x6b, 0xee, 0xc0, 0x33, 0x37,
	0x10, 0x38, 0xd8, 0x06, 0xd6, 0x81, 0x9a, 0xc2, 0x0b, 0x07, 0x4e, 0xe7, 0x25, 0x69, 0xc9, 0x40,
	0xac, 0xf0, 0x2d, 0xb0, 0x05, 0xec, 0x51, 0x0a, 0x5b, 0x55, 0x82, 0xaa, 0x3d, 0x54, 0x94, 0x2f,
	0x2a, 0xb1, 0x70, 0x5c, 0xa1, 0xc6, 0xb7, 0xd8, 0x75, 0xd1, 0x8f, 0x0d, 0xc4, 0x2e, 0xbf, 0x37,
	0x97, 0x87, 0x26, 0x61, 0xfe, 0xf0, 0xf1, 0x9a, 0x8e, 0xcd, 0x68, 0xc3, 0xc1, 0xd9, 0x16, 0xb3,
	0x64, 0x92, 0x4c, 0x47, 0x8b, 0xe3, 0xcd, 0xc7, 0x8b, 0x6d, 0x31, 0xff, 0x4d, 0xcc, 0xc5, 0xa1,
	0x8a, 0xf4, 0xce, 0x8c, 0x07, 0xbb, 0x50, 0xd3, 0xb2, 0x86, 0xd6, 0xca, 0x2a, 0xbb, 0x9a, 0x24,
	0xd3, 0x93, 0x45, 0x16, 0x23, 0xef, 0xbb, 0xc4, 0x23, 0x2d, 0xeb, 0x67, 0x2b, 0xab, 0xf4, 0xd6,
	0x9c, 0x0f, 0x79, 0xc3, 0x5d, 0xaf, 0x67, 0x51, 0x9f, 0xee, 0xeb, 0x27, 0xee, 0x22, 0xbe, 0x31,
	0x67, 0xe5, 0x5a, 0x04, 0x5d, 0x80, 0x41, 0x49, 0x76, 0xdd, 0xdb, 0xff, 0xc0, 0xdb, 0x5e, 0xc5,
	0xe7, 0x51, 0x3c, 0xf7, 0xec, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xac, 0x9b, 0x72, 0xba, 0x94, 0x01,
	0x00, 0x00,
}
