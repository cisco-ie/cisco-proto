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
// source: l2vpn_pbb_bmac_sa.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_standby_l2vpn_pbb_bsa

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

type L2VpnPbbBmacSa_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPbbBmacSa_KEYS) Reset()         { *m = L2VpnPbbBmacSa_KEYS{} }
func (m *L2VpnPbbBmacSa_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnPbbBmacSa_KEYS) ProtoMessage()    {}
func (*L2VpnPbbBmacSa_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb8952c826bb85f, []int{0}
}

func (m *L2VpnPbbBmacSa_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPbbBmacSa_KEYS.Unmarshal(m, b)
}
func (m *L2VpnPbbBmacSa_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPbbBmacSa_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnPbbBmacSa_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPbbBmacSa_KEYS.Merge(m, src)
}
func (m *L2VpnPbbBmacSa_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnPbbBmacSa_KEYS.Size(m)
}
func (m *L2VpnPbbBmacSa_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPbbBmacSa_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPbbBmacSa_KEYS proto.InternalMessageInfo

type L2VpnPbbBmacSa struct {
	Provisioned          bool     `protobuf:"varint,50,opt,name=provisioned,proto3" json:"provisioned,omitempty"`
	ChassisIsProvisioned bool     `protobuf:"varint,51,opt,name=chassis_is_provisioned,json=chassisIsProvisioned,proto3" json:"chassis_is_provisioned,omitempty"`
	BmacSa               string   `protobuf:"bytes,52,opt,name=bmac_sa,json=bmacSa,proto3" json:"bmac_sa,omitempty"`
	ChassisMac           string   `protobuf:"bytes,53,opt,name=chassis_mac,json=chassisMac,proto3" json:"chassis_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnPbbBmacSa) Reset()         { *m = L2VpnPbbBmacSa{} }
func (m *L2VpnPbbBmacSa) String() string { return proto.CompactTextString(m) }
func (*L2VpnPbbBmacSa) ProtoMessage()    {}
func (*L2VpnPbbBmacSa) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb8952c826bb85f, []int{1}
}

func (m *L2VpnPbbBmacSa) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnPbbBmacSa.Unmarshal(m, b)
}
func (m *L2VpnPbbBmacSa) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnPbbBmacSa.Marshal(b, m, deterministic)
}
func (m *L2VpnPbbBmacSa) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnPbbBmacSa.Merge(m, src)
}
func (m *L2VpnPbbBmacSa) XXX_Size() int {
	return xxx_messageInfo_L2VpnPbbBmacSa.Size(m)
}
func (m *L2VpnPbbBmacSa) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnPbbBmacSa.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnPbbBmacSa proto.InternalMessageInfo

func (m *L2VpnPbbBmacSa) GetProvisioned() bool {
	if m != nil {
		return m.Provisioned
	}
	return false
}

func (m *L2VpnPbbBmacSa) GetChassisIsProvisioned() bool {
	if m != nil {
		return m.ChassisIsProvisioned
	}
	return false
}

func (m *L2VpnPbbBmacSa) GetBmacSa() string {
	if m != nil {
		return m.BmacSa
	}
	return ""
}

func (m *L2VpnPbbBmacSa) GetChassisMac() string {
	if m != nil {
		return m.ChassisMac
	}
	return ""
}

func init() {
	proto.RegisterType((*L2VpnPbbBmacSa_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.l2vpn_pbb_bsa.l2vpn_pbb_bmac_sa_KEYS")
	proto.RegisterType((*L2VpnPbbBmacSa)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.standby.l2vpn_pbb_bsa.l2vpn_pbb_bmac_sa")
}

func init() { proto.RegisterFile("l2vpn_pbb_bmac_sa.proto", fileDescriptor_3fb8952c826bb85f) }

var fileDescriptor_3fb8952c826bb85f = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcf, 0x31, 0x2a, 0x2b,
	0xc8, 0x8b, 0x2f, 0x48, 0x4a, 0x8a, 0x4f, 0xca, 0x4d, 0x4c, 0x8e, 0x2f, 0x4e, 0xd4, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x32, 0x4d, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf,
	0x28, 0x8a, 0x87, 0xa8, 0xca, 0x2f, 0x48, 0x2d, 0xd2, 0x03, 0x33, 0xcb, 0x8c, 0xf4, 0x8a, 0x4b,
	0x12, 0xf3, 0x52, 0x92, 0x2a, 0xf5, 0x90, 0x0c, 0x28, 0x4e, 0x54, 0x92, 0xe0, 0x12, 0xc3, 0x30,
	0x31, 0xde, 0xdb, 0x35, 0x32, 0x58, 0x69, 0x29, 0x23, 0x97, 0x20, 0x86, 0x94, 0x90, 0x02, 0x17,
	0x77, 0x41, 0x51, 0x7e, 0x59, 0x66, 0x71, 0x66, 0x7e, 0x5e, 0x6a, 0x8a, 0x84, 0x91, 0x02, 0xa3,
	0x06, 0x47, 0x10, 0xb2, 0x90, 0x90, 0x09, 0x97, 0x58, 0x72, 0x46, 0x62, 0x71, 0x71, 0x66, 0x71,
	0x7c, 0x66, 0x71, 0x3c, 0xb2, 0x62, 0x63, 0xb0, 0x62, 0x11, 0xa8, 0xac, 0x67, 0x71, 0x00, 0x92,
	0x2e, 0x71, 0x2e, 0x76, 0xa8, 0x15, 0x12, 0x26, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x6c, 0x20, 0x6e,
	0x70, 0xa2, 0x90, 0x3c, 0x17, 0x37, 0xcc, 0xb8, 0xdc, 0xc4, 0x64, 0x09, 0x53, 0xb0, 0x24, 0x17,
	0x54, 0xc8, 0x37, 0x31, 0x39, 0x89, 0x0d, 0xec, 0x7f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0xd8, 0x92, 0x17, 0x1a, 0x01, 0x00, 0x00,
}