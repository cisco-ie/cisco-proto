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
// source: radius_dynamic_server.proto

package cisco_ios_xr_aaa_protocol_radius_oper_radius_nodes_node_dynamic_authorization

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

type RadiusDynamicServer_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RadiusDynamicServer_KEYS) Reset()         { *m = RadiusDynamicServer_KEYS{} }
func (m *RadiusDynamicServer_KEYS) String() string { return proto.CompactTextString(m) }
func (*RadiusDynamicServer_KEYS) ProtoMessage()    {}
func (*RadiusDynamicServer_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ee1e85f5e5506, []int{0}
}

func (m *RadiusDynamicServer_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDynamicServer_KEYS.Unmarshal(m, b)
}
func (m *RadiusDynamicServer_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDynamicServer_KEYS.Marshal(b, m, deterministic)
}
func (m *RadiusDynamicServer_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDynamicServer_KEYS.Merge(m, src)
}
func (m *RadiusDynamicServer_KEYS) XXX_Size() int {
	return xxx_messageInfo_RadiusDynamicServer_KEYS.Size(m)
}
func (m *RadiusDynamicServer_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDynamicServer_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDynamicServer_KEYS proto.InternalMessageInfo

func (m *RadiusDynamicServer_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RadiusDynamicServer struct {
	DisconnectedInvalidRequests uint32   `protobuf:"varint,50,opt,name=disconnected_invalid_requests,json=disconnectedInvalidRequests,proto3" json:"disconnected_invalid_requests,omitempty"`
	InvalidCoaRequests          uint32   `protobuf:"varint,51,opt,name=invalid_coa_requests,json=invalidCoaRequests,proto3" json:"invalid_coa_requests,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *RadiusDynamicServer) Reset()         { *m = RadiusDynamicServer{} }
func (m *RadiusDynamicServer) String() string { return proto.CompactTextString(m) }
func (*RadiusDynamicServer) ProtoMessage()    {}
func (*RadiusDynamicServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ee1e85f5e5506, []int{1}
}

func (m *RadiusDynamicServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RadiusDynamicServer.Unmarshal(m, b)
}
func (m *RadiusDynamicServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RadiusDynamicServer.Marshal(b, m, deterministic)
}
func (m *RadiusDynamicServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RadiusDynamicServer.Merge(m, src)
}
func (m *RadiusDynamicServer) XXX_Size() int {
	return xxx_messageInfo_RadiusDynamicServer.Size(m)
}
func (m *RadiusDynamicServer) XXX_DiscardUnknown() {
	xxx_messageInfo_RadiusDynamicServer.DiscardUnknown(m)
}

var xxx_messageInfo_RadiusDynamicServer proto.InternalMessageInfo

func (m *RadiusDynamicServer) GetDisconnectedInvalidRequests() uint32 {
	if m != nil {
		return m.DisconnectedInvalidRequests
	}
	return 0
}

func (m *RadiusDynamicServer) GetInvalidCoaRequests() uint32 {
	if m != nil {
		return m.InvalidCoaRequests
	}
	return 0
}

func init() {
	proto.RegisterType((*RadiusDynamicServer_KEYS)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dynamic_authorization.radius_dynamic_server_KEYS")
	proto.RegisterType((*RadiusDynamicServer)(nil), "cisco_ios_xr_aaa_protocol_radius_oper.radius.nodes.node.dynamic_authorization.radius_dynamic_server")
}

func init() { proto.RegisterFile("radius_dynamic_server.proto", fileDescriptor_576ee1e85f5e5506) }

var fileDescriptor_576ee1e85f5e5506 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd0, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0x06, 0x70, 0xd2, 0x88, 0x17, 0xb0, 0x09, 0x0a, 0x87, 0x8b, 0x70, 0x5c, 0x75, 0x55, 0x10,
	0xaf, 0xb2, 0x55, 0x2c, 0x44, 0xb4, 0x58, 0x2b, 0xab, 0x61, 0x4c, 0x06, 0x1c, 0xb8, 0xcd, 0x9c,
	0x93, 0xec, 0xa1, 0xbe, 0x83, 0xef, 0x2c, 0xfb, 0x0f, 0xaf, 0xd8, 0x26, 0x84, 0x2f, 0xdf, 0x2f,
	0x09, 0x63, 0x2b, 0xc5, 0xc8, 0x6d, 0x86, 0xf8, 0x9d, 0xb0, 0xe1, 0x00, 0x99, 0xf4, 0x40, 0xea,
	0xf7, 0x2a, 0x45, 0xdc, 0x73, 0xe0, 0x1c, 0x04, 0x58, 0x32, 0x7c, 0x29, 0x20, 0x22, 0xf4, 0x79,
	0x90, 0x1d, 0x8c, 0x4c, 0xf6, 0xa4, 0x7e, 0xd8, 0xfb, 0x24, 0x91, 0x86, 0xd5, 0x4f, 0xb7, 0x61,
	0x5b, 0x3e, 0x44, 0xf9, 0x07, 0x0b, 0x4b, 0x5a, 0xdf, 0xda, 0xcb, 0xd9, 0xd7, 0xe0, 0xe9, 0xe1,
	0xed, 0xd5, 0x55, 0x76, 0xd1, 0x59, 0x48, 0xd8, 0xd0, 0xd2, 0xac, 0xcc, 0x66, 0x51, 0x9f, 0x76,
	0xc1, 0x0b, 0x36, 0xb4, 0xfe, 0x35, 0xf6, 0x62, 0xd6, 0xba, 0x3b, 0x7b, 0x15, 0xbb, 0x5f, 0xa6,
	0x44, 0xa1, 0x50, 0x04, 0x4e, 0x07, 0xdc, 0x71, 0x04, 0xa5, 0xcf, 0x96, 0x72, 0xc9, 0xcb, 0x9b,
	0x95, 0xd9, 0x9c, 0xd5, 0xd5, 0x71, 0xe9, 0x71, 0xe8, 0xd4, 0x63, 0xc5, 0x5d, 0xdb, 0xf3, 0x89,
	0x05, 0xc1, 0x7f, 0xba, 0xed, 0xa9, 0x1b, 0xcf, 0xee, 0x05, 0x27, 0xf1, 0x7e, 0xd2, 0x0f, 0x62,
	0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x20, 0xc3, 0x1f, 0x6a, 0x3f, 0x01, 0x00, 0x00,
}
