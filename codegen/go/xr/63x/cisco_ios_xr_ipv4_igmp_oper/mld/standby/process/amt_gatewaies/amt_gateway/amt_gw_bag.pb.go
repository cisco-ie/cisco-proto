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
// source: amt_gw_bag.proto

package cisco_ios_xr_ipv4_igmp_oper_mld_standby_process_amt_gatewaies_amt_gateway

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

type AmtGwBag_KEYS struct {
	GatewayAddress       string   `protobuf:"bytes,1,opt,name=gateway_address,json=gatewayAddress,proto3" json:"gateway_address,omitempty"`
	Port                 uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmtGwBag_KEYS) Reset()         { *m = AmtGwBag_KEYS{} }
func (m *AmtGwBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AmtGwBag_KEYS) ProtoMessage()    {}
func (*AmtGwBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff04cd119f53e4f3, []int{0}
}

func (m *AmtGwBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmtGwBag_KEYS.Unmarshal(m, b)
}
func (m *AmtGwBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmtGwBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AmtGwBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmtGwBag_KEYS.Merge(m, src)
}
func (m *AmtGwBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AmtGwBag_KEYS.Size(m)
}
func (m *AmtGwBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AmtGwBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AmtGwBag_KEYS proto.InternalMessageInfo

func (m *AmtGwBag_KEYS) GetGatewayAddress() string {
	if m != nil {
		return m.GatewayAddress
	}
	return ""
}

func (m *AmtGwBag_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type AmtGwBag struct {
	Amtgw                string   `protobuf:"bytes,50,opt,name=amtgw,proto3" json:"amtgw,omitempty"`
	AmtPort              uint32   `protobuf:"varint,51,opt,name=amt_port,json=amtPort,proto3" json:"amt_port,omitempty"`
	KeyLen               uint32   `protobuf:"varint,52,opt,name=key_len,json=keyLen,proto3" json:"key_len,omitempty"`
	Amtnh                uint32   `protobuf:"varint,53,opt,name=amtnh,proto3" json:"amtnh,omitempty"`
	AmtNonce             uint32   `protobuf:"varint,54,opt,name=amt_nonce,json=amtNonce,proto3" json:"amt_nonce,omitempty"`
	Idb                  uint64   `protobuf:"varint,55,opt,name=idb,proto3" json:"idb,omitempty"`
	MemUpdIn             uint32   `protobuf:"varint,56,opt,name=mem_upd_in,json=memUpdIn,proto3" json:"mem_upd_in,omitempty"`
	MemUpdOut            uint32   `protobuf:"varint,57,opt,name=mem_upd_out,json=memUpdOut,proto3" json:"mem_upd_out,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AmtGwBag) Reset()         { *m = AmtGwBag{} }
func (m *AmtGwBag) String() string { return proto.CompactTextString(m) }
func (*AmtGwBag) ProtoMessage()    {}
func (*AmtGwBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff04cd119f53e4f3, []int{1}
}

func (m *AmtGwBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AmtGwBag.Unmarshal(m, b)
}
func (m *AmtGwBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AmtGwBag.Marshal(b, m, deterministic)
}
func (m *AmtGwBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AmtGwBag.Merge(m, src)
}
func (m *AmtGwBag) XXX_Size() int {
	return xxx_messageInfo_AmtGwBag.Size(m)
}
func (m *AmtGwBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AmtGwBag.DiscardUnknown(m)
}

var xxx_messageInfo_AmtGwBag proto.InternalMessageInfo

func (m *AmtGwBag) GetAmtgw() string {
	if m != nil {
		return m.Amtgw
	}
	return ""
}

func (m *AmtGwBag) GetAmtPort() uint32 {
	if m != nil {
		return m.AmtPort
	}
	return 0
}

func (m *AmtGwBag) GetKeyLen() uint32 {
	if m != nil {
		return m.KeyLen
	}
	return 0
}

func (m *AmtGwBag) GetAmtnh() uint32 {
	if m != nil {
		return m.Amtnh
	}
	return 0
}

func (m *AmtGwBag) GetAmtNonce() uint32 {
	if m != nil {
		return m.AmtNonce
	}
	return 0
}

func (m *AmtGwBag) GetIdb() uint64 {
	if m != nil {
		return m.Idb
	}
	return 0
}

func (m *AmtGwBag) GetMemUpdIn() uint32 {
	if m != nil {
		return m.MemUpdIn
	}
	return 0
}

func (m *AmtGwBag) GetMemUpdOut() uint32 {
	if m != nil {
		return m.MemUpdOut
	}
	return 0
}

func init() {
	proto.RegisterType((*AmtGwBag_KEYS)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.process.amt_gatewaies.amt_gateway.amt_gw_bag_KEYS")
	proto.RegisterType((*AmtGwBag)(nil), "cisco_ios_xr_ipv4_igmp_oper.mld.standby.process.amt_gatewaies.amt_gateway.amt_gw_bag")
}

func init() { proto.RegisterFile("amt_gw_bag.proto", fileDescriptor_ff04cd119f53e4f3) }

var fileDescriptor_ff04cd119f53e4f3 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xcd, 0x4e, 0x02, 0x31,
	0x14, 0x85, 0x33, 0x8a, 0x20, 0xd7, 0x28, 0xa4, 0x31, 0xb1, 0x46, 0x63, 0x08, 0x1b, 0x59, 0xcd,
	0x42, 0xf0, 0x6f, 0xe9, 0xc2, 0x05, 0xd1, 0xa0, 0xc1, 0xb8, 0x70, 0x75, 0xd3, 0x99, 0xde, 0x8c,
	0x0d, 0xf4, 0x27, 0xd3, 0x22, 0xce, 0xa3, 0xfa, 0x36, 0x66, 0x3a, 0x18, 0xdc, 0xf5, 0x9c, 0xef,
	0xdc, 0x6f, 0x51, 0xe8, 0x0b, 0x1d, 0xb0, 0x58, 0x63, 0x26, 0x8a, 0xd4, 0x95, 0x36, 0x58, 0x36,
	0xcd, 0x95, 0xcf, 0x2d, 0x2a, 0xeb, 0xf1, 0xbb, 0x44, 0xe5, 0xbe, 0x26, 0xa8, 0x0a, 0xed, 0xd0,
	0x3a, 0x2a, 0x53, 0xbd, 0x94, 0xa9, 0x0f, 0xc2, 0xc8, 0xac, 0xaa, 0xe7, 0x39, 0x79, 0x9f, 0x46,
	0x83, 0x08, 0xb4, 0x16, 0x8a, 0xfe, 0xa7, 0x6a, 0x38, 0x83, 0xde, 0x56, 0x8f, 0x4f, 0x8f, 0x1f,
	0x6f, 0xec, 0x12, 0x7a, 0x1b, 0x8a, 0x42, 0xca, 0x92, 0xbc, 0xe7, 0xc9, 0x20, 0x19, 0x75, 0xe7,
	0x47, 0x9b, 0xfa, 0xa1, 0x69, 0x19, 0x83, 0x96, 0xb3, 0x65, 0xe0, 0x3b, 0x83, 0x64, 0x74, 0x38,
	0x8f, 0xef, 0xe1, 0x4f, 0x02, 0xb0, 0x15, 0xb2, 0x63, 0xd8, 0x13, 0x3a, 0x14, 0x6b, 0x7e, 0x15,
	0x0d, 0x4d, 0x60, 0xa7, 0xb0, 0x5f, 0x6f, 0xe2, 0xf1, 0x38, 0x1e, 0x77, 0x84, 0x0e, 0xaf, 0xb6,
	0x0c, 0xec, 0x04, 0x3a, 0x0b, 0xaa, 0x70, 0x49, 0x86, 0x4f, 0x22, 0x69, 0x2f, 0xa8, 0x7a, 0x26,
	0xb3, 0x31, 0x99, 0x4f, 0x7e, 0x1d, 0xeb, 0x26, 0xb0, 0x33, 0xe8, 0xd6, 0x26, 0x63, 0x4d, 0x4e,
	0xfc, 0x26, 0x92, 0x5a, 0x3d, 0xab, 0x33, 0xeb, 0xc3, 0xae, 0x92, 0x19, 0xbf, 0x1d, 0x24, 0xa3,
	0xd6, 0xbc, 0x7e, 0xb2, 0x73, 0x00, 0x4d, 0x1a, 0x57, 0x4e, 0xa2, 0x32, 0xfc, 0xae, 0xd9, 0x6b,
	0xd2, 0xef, 0x4e, 0x4e, 0x0d, 0xbb, 0x80, 0x83, 0x3f, 0x6a, 0x57, 0x81, 0xdf, 0x47, 0xdc, 0x6d,
	0xf0, 0xcb, 0x2a, 0x64, 0xed, 0xf8, 0xfb, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x16,
	0xa0, 0x5b, 0x91, 0x01, 0x00, 0x00,
}
