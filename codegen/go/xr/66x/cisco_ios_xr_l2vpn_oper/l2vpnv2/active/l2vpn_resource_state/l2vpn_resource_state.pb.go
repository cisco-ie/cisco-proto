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
// source: l2vpn_resource_state.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_l2vpn_resource_state

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

type L2VpnResourceState_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnResourceState_KEYS) Reset()         { *m = L2VpnResourceState_KEYS{} }
func (m *L2VpnResourceState_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnResourceState_KEYS) ProtoMessage()    {}
func (*L2VpnResourceState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c34676918da1de, []int{0}
}

func (m *L2VpnResourceState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Unmarshal(m, b)
}
func (m *L2VpnResourceState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnResourceState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnResourceState_KEYS.Merge(m, src)
}
func (m *L2VpnResourceState_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnResourceState_KEYS.Size(m)
}
func (m *L2VpnResourceState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnResourceState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnResourceState_KEYS proto.InternalMessageInfo

type L2VpnResourceState struct {
	ResourceOutOfMemoryState uint32   `protobuf:"varint,50,opt,name=resource_out_of_memory_state,json=resourceOutOfMemoryState,proto3" json:"resource_out_of_memory_state,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *L2VpnResourceState) Reset()         { *m = L2VpnResourceState{} }
func (m *L2VpnResourceState) String() string { return proto.CompactTextString(m) }
func (*L2VpnResourceState) ProtoMessage()    {}
func (*L2VpnResourceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b2c34676918da1de, []int{1}
}

func (m *L2VpnResourceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnResourceState.Unmarshal(m, b)
}
func (m *L2VpnResourceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnResourceState.Marshal(b, m, deterministic)
}
func (m *L2VpnResourceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnResourceState.Merge(m, src)
}
func (m *L2VpnResourceState) XXX_Size() int {
	return xxx_messageInfo_L2VpnResourceState.Size(m)
}
func (m *L2VpnResourceState) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnResourceState.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnResourceState proto.InternalMessageInfo

func (m *L2VpnResourceState) GetResourceOutOfMemoryState() uint32 {
	if m != nil {
		return m.ResourceOutOfMemoryState
	}
	return 0
}

func init() {
	proto.RegisterType((*L2VpnResourceState_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.l2vpn_resource_state.l2vpn_resource_state_KEYS")
	proto.RegisterType((*L2VpnResourceState)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.l2vpn_resource_state.l2vpn_resource_state")
}

func init() { proto.RegisterFile("l2vpn_resource_state.proto", fileDescriptor_b2c34676918da1de) }

var fileDescriptor_b2c34676918da1de = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x31, 0x2a, 0x2b,
	0xc8, 0x8b, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0x2f, 0x2e, 0x49, 0x2c, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xb2, 0x4e, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc,
	0x2f, 0x8e, 0xaf, 0x28, 0x8a, 0x87, 0x28, 0xcc, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0x33, 0xcb, 0x8c,
	0xf4, 0x12, 0x93, 0x4b, 0x32, 0xcb, 0x52, 0xf5, 0xb0, 0x19, 0xa1, 0x24, 0xcd, 0x25, 0x89, 0x4d,
	0x3c, 0xde, 0xdb, 0x35, 0x32, 0x58, 0x29, 0x8c, 0x4b, 0x04, 0x9b, 0xa4, 0x90, 0x1d, 0x97, 0x0c,
	0x5c, 0x24, 0xbf, 0xb4, 0x24, 0x3e, 0x3f, 0x2d, 0x3e, 0x37, 0x35, 0x37, 0xbf, 0xa8, 0x12, 0x22,
	0x2f, 0x61, 0xa4, 0xc0, 0xa8, 0xc1, 0x1b, 0x24, 0x01, 0x53, 0xe3, 0x5f, 0x5a, 0xe2, 0x9f, 0xe6,
	0x0b, 0x56, 0x10, 0x0c, 0x92, 0x4f, 0x62, 0x03, 0x3b, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xa0, 0x01, 0x9d, 0x6a, 0xd6, 0x00, 0x00, 0x00,
}
