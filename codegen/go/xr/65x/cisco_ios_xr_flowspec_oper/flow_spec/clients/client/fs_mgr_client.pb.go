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
// source: fs_mgr_client.proto

package cisco_ios_xr_flowspec_oper_flow_spec_clients_client

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

type FsMgrClient_KEYS struct {
	ClientType           string   `protobuf:"bytes,1,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	ClientId             uint32   `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrClient_KEYS) Reset()         { *m = FsMgrClient_KEYS{} }
func (m *FsMgrClient_KEYS) String() string { return proto.CompactTextString(m) }
func (*FsMgrClient_KEYS) ProtoMessage()    {}
func (*FsMgrClient_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_80490240ce1ec692, []int{0}
}

func (m *FsMgrClient_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrClient_KEYS.Unmarshal(m, b)
}
func (m *FsMgrClient_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrClient_KEYS.Marshal(b, m, deterministic)
}
func (m *FsMgrClient_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrClient_KEYS.Merge(m, src)
}
func (m *FsMgrClient_KEYS) XXX_Size() int {
	return xxx_messageInfo_FsMgrClient_KEYS.Size(m)
}
func (m *FsMgrClient_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrClient_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrClient_KEYS proto.InternalMessageInfo

func (m *FsMgrClient_KEYS) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *FsMgrClient_KEYS) GetClientId() uint32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type FsMgrClient struct {
	ClientState          string   `protobuf:"bytes,50,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty"`
	ClientFlows          uint32   `protobuf:"varint,51,opt,name=client_flows,json=clientFlows,proto3" json:"client_flows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FsMgrClient) Reset()         { *m = FsMgrClient{} }
func (m *FsMgrClient) String() string { return proto.CompactTextString(m) }
func (*FsMgrClient) ProtoMessage()    {}
func (*FsMgrClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_80490240ce1ec692, []int{1}
}

func (m *FsMgrClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FsMgrClient.Unmarshal(m, b)
}
func (m *FsMgrClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FsMgrClient.Marshal(b, m, deterministic)
}
func (m *FsMgrClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FsMgrClient.Merge(m, src)
}
func (m *FsMgrClient) XXX_Size() int {
	return xxx_messageInfo_FsMgrClient.Size(m)
}
func (m *FsMgrClient) XXX_DiscardUnknown() {
	xxx_messageInfo_FsMgrClient.DiscardUnknown(m)
}

var xxx_messageInfo_FsMgrClient proto.InternalMessageInfo

func (m *FsMgrClient) GetClientState() string {
	if m != nil {
		return m.ClientState
	}
	return ""
}

func (m *FsMgrClient) GetClientFlows() uint32 {
	if m != nil {
		return m.ClientFlows
	}
	return 0
}

func init() {
	proto.RegisterType((*FsMgrClient_KEYS)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.clients.client.fs_mgr_client_KEYS")
	proto.RegisterType((*FsMgrClient)(nil), "cisco_ios_xr_flowspec_oper.flow_spec.clients.client.fs_mgr_client")
}

func init() { proto.RegisterFile("fs_mgr_client.proto", fileDescriptor_80490240ce1ec692) }

var fileDescriptor_80490240ce1ec692 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x2b, 0x8e, 0xcf,
	0x4d, 0x2f, 0x8a, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x32, 0x4e, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x4f, 0xcb,
	0xc9, 0x2f, 0x2f, 0x2e, 0x48, 0x4d, 0x8e, 0xcf, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0xf1, 0xe2, 0x41,
	0x5c, 0x3d, 0x88, 0x96, 0x62, 0x28, 0xad, 0x14, 0xc4, 0x25, 0x84, 0x62, 0x56, 0xbc, 0xb7, 0x6b,
	0x64, 0xb0, 0x90, 0x3c, 0x17, 0x37, 0x94, 0x5b, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x05, 0x11, 0x0a, 0xa9, 0x2c, 0x48, 0x15, 0x92, 0xe6, 0xe2, 0x84, 0x2a, 0xc8,
	0x4c, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0d, 0xe2, 0x80, 0x08, 0x78, 0xa6, 0x28, 0x85, 0x72,
	0xf1, 0xa2, 0x98, 0x29, 0xa4, 0xc8, 0xc5, 0x03, 0x55, 0x5d, 0x5c, 0x92, 0x58, 0x92, 0x2a, 0x61,
	0x04, 0x36, 0x0f, 0x6a, 0x45, 0x30, 0x48, 0x08, 0x49, 0x09, 0xd8, 0xe1, 0x12, 0xc6, 0x60, 0x33,
	0xa1, 0x4a, 0xdc, 0x40, 0x42, 0x49, 0x6c, 0x60, 0x6f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x33, 0x91, 0x04, 0xfd, 0x00, 0x00, 0x00,
}