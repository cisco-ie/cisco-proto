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
// source: vpdn_redundancy.proto

package cisco_ios_xr_tunnel_vpdn_oper_vpdn_nodes_node_vpdn_redundancy

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

type VpdnRedundancy_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpdnRedundancy_KEYS) Reset()         { *m = VpdnRedundancy_KEYS{} }
func (m *VpdnRedundancy_KEYS) String() string { return proto.CompactTextString(m) }
func (*VpdnRedundancy_KEYS) ProtoMessage()    {}
func (*VpdnRedundancy_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_52420718010483c6, []int{0}
}

func (m *VpdnRedundancy_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpdnRedundancy_KEYS.Unmarshal(m, b)
}
func (m *VpdnRedundancy_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpdnRedundancy_KEYS.Marshal(b, m, deterministic)
}
func (m *VpdnRedundancy_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpdnRedundancy_KEYS.Merge(m, src)
}
func (m *VpdnRedundancy_KEYS) XXX_Size() int {
	return xxx_messageInfo_VpdnRedundancy_KEYS.Size(m)
}
func (m *VpdnRedundancy_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VpdnRedundancy_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VpdnRedundancy_KEYS proto.InternalMessageInfo

func (m *VpdnRedundancy_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type VpdnRedundancy struct {
	SessionTotal         uint32   `protobuf:"varint,50,opt,name=session_total,json=sessionTotal,proto3" json:"session_total,omitempty"`
	SessionSynced        uint32   `protobuf:"varint,51,opt,name=session_synced,json=sessionSynced,proto3" json:"session_synced,omitempty"`
	State                string   `protobuf:"bytes,52,opt,name=state,proto3" json:"state,omitempty"`
	StartTime            uint64   `protobuf:"varint,53,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	FinishTime           uint64   `protobuf:"varint,54,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	AbortTime            uint64   `protobuf:"varint,55,opt,name=abort_time,json=abortTime,proto3" json:"abort_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VpdnRedundancy) Reset()         { *m = VpdnRedundancy{} }
func (m *VpdnRedundancy) String() string { return proto.CompactTextString(m) }
func (*VpdnRedundancy) ProtoMessage()    {}
func (*VpdnRedundancy) Descriptor() ([]byte, []int) {
	return fileDescriptor_52420718010483c6, []int{1}
}

func (m *VpdnRedundancy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VpdnRedundancy.Unmarshal(m, b)
}
func (m *VpdnRedundancy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VpdnRedundancy.Marshal(b, m, deterministic)
}
func (m *VpdnRedundancy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VpdnRedundancy.Merge(m, src)
}
func (m *VpdnRedundancy) XXX_Size() int {
	return xxx_messageInfo_VpdnRedundancy.Size(m)
}
func (m *VpdnRedundancy) XXX_DiscardUnknown() {
	xxx_messageInfo_VpdnRedundancy.DiscardUnknown(m)
}

var xxx_messageInfo_VpdnRedundancy proto.InternalMessageInfo

func (m *VpdnRedundancy) GetSessionTotal() uint32 {
	if m != nil {
		return m.SessionTotal
	}
	return 0
}

func (m *VpdnRedundancy) GetSessionSynced() uint32 {
	if m != nil {
		return m.SessionSynced
	}
	return 0
}

func (m *VpdnRedundancy) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *VpdnRedundancy) GetStartTime() uint64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *VpdnRedundancy) GetFinishTime() uint64 {
	if m != nil {
		return m.FinishTime
	}
	return 0
}

func (m *VpdnRedundancy) GetAbortTime() uint64 {
	if m != nil {
		return m.AbortTime
	}
	return 0
}

func init() {
	proto.RegisterType((*VpdnRedundancy_KEYS)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_redundancy.vpdn_redundancy_KEYS")
	proto.RegisterType((*VpdnRedundancy)(nil), "cisco_ios_xr_tunnel_vpdn_oper.vpdn.nodes.node.vpdn_redundancy.vpdn_redundancy")
}

func init() { proto.RegisterFile("vpdn_redundancy.proto", fileDescriptor_52420718010483c6) }

var fileDescriptor_52420718010483c6 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0x29, 0xa8, 0xb8, 0xa3, 0xab, 0x10, 0x56, 0x28, 0xc8, 0x62, 0x59, 0x11, 0x7a, 0xea,
	0xc1, 0xfa, 0x71, 0xf2, 0xe8, 0x49, 0xf0, 0xd0, 0xdd, 0x8b, 0xa7, 0x90, 0x6d, 0x46, 0x0c, 0x6c,
	0x27, 0x25, 0x93, 0x15, 0xf7, 0x6f, 0xfa, 0x8b, 0x24, 0x53, 0x7b, 0xe9, 0x25, 0x24, 0xcf, 0xfb,
	0xbc, 0x19, 0x18, 0xb8, 0xfa, 0xee, 0x2d, 0xe9, 0x80, 0x76, 0x4f, 0xd6, 0x50, 0x7b, 0xa8, 0xfa,
	0xe0, 0xa3, 0x57, 0x2f, 0xad, 0xe3, 0xd6, 0x6b, 0xe7, 0x59, 0xff, 0x04, 0x1d, 0xf7, 0x44, 0xb8,
	0xd3, 0xa2, 0xfa, 0x1e, 0x43, 0x95, 0x6e, 0x15, 0x79, 0x8b, 0x2c, 0x67, 0x35, 0xf9, 0x64, 0x55,
	0xc3, 0x62, 0x82, 0xf4, 0xdb, 0xeb, 0xc7, 0x5a, 0x5d, 0xc3, 0x2c, 0xf9, 0x9a, 0x4c, 0x87, 0x79,
	0x56, 0x64, 0xe5, 0xac, 0x39, 0x4d, 0xe0, 0xdd, 0x74, 0xb8, 0xfa, 0xcd, 0xe0, 0x72, 0xd2, 0x52,
	0xb7, 0x30, 0x67, 0x64, 0x76, 0x9e, 0x74, 0xf4, 0xd1, 0xec, 0xf2, 0xfb, 0x22, 0x2b, 0xe7, 0xcd,
	0xf9, 0x3f, 0xdc, 0x24, 0xa6, 0xee, 0xe0, 0x62, 0x94, 0xf8, 0x40, 0x2d, 0xda, 0xbc, 0x16, 0x6b,
	0xac, 0xae, 0x05, 0xaa, 0x05, 0x1c, 0x73, 0x34, 0x11, 0xf3, 0x07, 0x19, 0x3c, 0x3c, 0xd4, 0x12,
	0x80, 0xa3, 0x09, 0x51, 0x47, 0xd7, 0x61, 0xfe, 0x58, 0x64, 0xe5, 0x51, 0x33, 0x13, 0xb2, 0x71,
	0x1d, 0xaa, 0x1b, 0x38, 0xfb, 0x74, 0xe4, 0xf8, 0x6b, 0xc8, 0x9f, 0x24, 0x87, 0x01, 0x89, 0xb0,
	0x04, 0x30, 0x5b, 0x3f, 0xf6, 0x9f, 0x87, 0xbe, 0x90, 0x14, 0x6f, 0x4f, 0x64, 0x9f, 0xf5, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xe5, 0x70, 0x0f, 0x68, 0x01, 0x00, 0x00,
}
