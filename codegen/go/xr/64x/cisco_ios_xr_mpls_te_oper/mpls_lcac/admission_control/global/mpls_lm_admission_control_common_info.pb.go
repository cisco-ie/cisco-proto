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
// source: mpls_lm_admission_control_common_info.proto

package cisco_ios_xr_mpls_te_oper_mpls_lcac_admission_control_global

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

type MplsLmAdmissionControlCommonInfo_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmAdmissionControlCommonInfo_KEYS) Reset()         { *m = MplsLmAdmissionControlCommonInfo_KEYS{} }
func (m *MplsLmAdmissionControlCommonInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*MplsLmAdmissionControlCommonInfo_KEYS) ProtoMessage()    {}
func (*MplsLmAdmissionControlCommonInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b285e0c45d15e3, []int{0}
}

func (m *MplsLmAdmissionControlCommonInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS.Unmarshal(m, b)
}
func (m *MplsLmAdmissionControlCommonInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *MplsLmAdmissionControlCommonInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS.Merge(m, src)
}
func (m *MplsLmAdmissionControlCommonInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS.Size(m)
}
func (m *MplsLmAdmissionControlCommonInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmAdmissionControlCommonInfo_KEYS proto.InternalMessageInfo

type MplsLmAdmissionControlCommonInfo struct {
	IsRoleStandby        bool     `protobuf:"varint,50,opt,name=is_role_standby,json=isRoleStandby,proto3" json:"is_role_standby,omitempty"`
	TotalTunnels         uint32   `protobuf:"varint,51,opt,name=total_tunnels,json=totalTunnels,proto3" json:"total_tunnels,omitempty"`
	TotalP2MpTunnels     uint32   `protobuf:"varint,52,opt,name=total_p2mp_tunnels,json=totalP2mpTunnels,proto3" json:"total_p2mp_tunnels,omitempty"`
	SelectedTunnels      uint32   `protobuf:"varint,53,opt,name=selected_tunnels,json=selectedTunnels,proto3" json:"selected_tunnels,omitempty"`
	BandwidthUnits       string   `protobuf:"bytes,54,opt,name=bandwidth_units,json=bandwidthUnits,proto3" json:"bandwidth_units,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MplsLmAdmissionControlCommonInfo) Reset()         { *m = MplsLmAdmissionControlCommonInfo{} }
func (m *MplsLmAdmissionControlCommonInfo) String() string { return proto.CompactTextString(m) }
func (*MplsLmAdmissionControlCommonInfo) ProtoMessage()    {}
func (*MplsLmAdmissionControlCommonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8b285e0c45d15e3, []int{1}
}

func (m *MplsLmAdmissionControlCommonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo.Unmarshal(m, b)
}
func (m *MplsLmAdmissionControlCommonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo.Marshal(b, m, deterministic)
}
func (m *MplsLmAdmissionControlCommonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MplsLmAdmissionControlCommonInfo.Merge(m, src)
}
func (m *MplsLmAdmissionControlCommonInfo) XXX_Size() int {
	return xxx_messageInfo_MplsLmAdmissionControlCommonInfo.Size(m)
}
func (m *MplsLmAdmissionControlCommonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MplsLmAdmissionControlCommonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MplsLmAdmissionControlCommonInfo proto.InternalMessageInfo

func (m *MplsLmAdmissionControlCommonInfo) GetIsRoleStandby() bool {
	if m != nil {
		return m.IsRoleStandby
	}
	return false
}

func (m *MplsLmAdmissionControlCommonInfo) GetTotalTunnels() uint32 {
	if m != nil {
		return m.TotalTunnels
	}
	return 0
}

func (m *MplsLmAdmissionControlCommonInfo) GetTotalP2MpTunnels() uint32 {
	if m != nil {
		return m.TotalP2MpTunnels
	}
	return 0
}

func (m *MplsLmAdmissionControlCommonInfo) GetSelectedTunnels() uint32 {
	if m != nil {
		return m.SelectedTunnels
	}
	return 0
}

func (m *MplsLmAdmissionControlCommonInfo) GetBandwidthUnits() string {
	if m != nil {
		return m.BandwidthUnits
	}
	return ""
}

func init() {
	proto.RegisterType((*MplsLmAdmissionControlCommonInfo_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.admission_control.global.mpls_lm_admission_control_common_info_KEYS")
	proto.RegisterType((*MplsLmAdmissionControlCommonInfo)(nil), "cisco_ios_xr_mpls_te_oper.mpls_lcac.admission_control.global.mpls_lm_admission_control_common_info")
}

func init() {
	proto.RegisterFile("mpls_lm_admission_control_common_info.proto", fileDescriptor_d8b285e0c45d15e3)
}

var fileDescriptor_d8b285e0c45d15e3 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x71, 0xba, 0x88, 0x06, 0xcf, 0x1e, 0x9d, 0x3a, 0x96, 0x13, 0xb5, 0xea, 0xd1, 0xa1,
	0xa7, 0x4e, 0xae, 0x4e, 0x2e, 0xd2, 0xd3, 0xc1, 0xe9, 0x91, 0xa6, 0x51, 0x03, 0x49, 0x5e, 0xc8,
	0x7b, 0x87, 0xfa, 0x8f, 0x3b, 0x8b, 0x29, 0xed, 0xe2, 0x72, 0xeb, 0x37, 0x9f, 0xfc, 0x02, 0x11,
	0xd7, 0x2e, 0x58, 0x02, 0xeb, 0x40, 0x0e, 0xce, 0x10, 0x19, 0xf4, 0xa0, 0xd0, 0x73, 0x44, 0x0b,
	0x0a, 0x9d, 0x43, 0x0f, 0xc6, 0xbf, 0x61, 0x13, 0x22, 0x32, 0x16, 0xf7, 0xca, 0x90, 0x42, 0x30,
	0x48, 0xf0, 0x15, 0x21, 0xdd, 0x64, 0x0d, 0x18, 0x74, 0x6c, 0xc6, 0x19, 0x25, 0x55, 0xf3, 0x6f,
	0xa8, 0x79, 0xb7, 0xd8, 0x4b, 0xbb, 0x5a, 0x8b, 0xab, 0xbd, 0x1e, 0x83, 0xc7, 0x87, 0xd7, 0xed,
	0xea, 0x27, 0x13, 0x67, 0x7b, 0xf1, 0xe2, 0x5c, 0xe4, 0x86, 0x20, 0xa2, 0xd5, 0x40, 0x2c, 0xfd,
	0xd0, 0x7f, 0x97, 0x6d, 0x95, 0xd5, 0x87, 0xdd, 0xc2, 0x50, 0x87, 0x56, 0x6f, 0xc7, 0x58, 0x9c,
	0x8a, 0x05, 0x23, 0x4b, 0x0b, 0xbc, 0xf3, 0x5e, 0x5b, 0x2a, 0x37, 0x55, 0x56, 0x2f, 0xba, 0xe3,
	0x14, 0x9f, 0xc7, 0x56, 0xac, 0x45, 0x31, 0xa2, 0xd0, 0xba, 0x30, 0xcb, 0x9b, 0x24, 0x97, 0xe9,
	0xe4, 0xa9, 0x75, 0x61, 0xd2, 0x97, 0x62, 0x49, 0xda, 0x6a, 0xc5, 0x7a, 0x98, 0xed, 0x6d, 0xb2,
	0xf9, 0xd4, 0x27, 0x7a, 0x21, 0xf2, 0x5e, 0xfa, 0xe1, 0xd3, 0x0c, 0xfc, 0x01, 0x3b, 0x6f, 0x98,
	0xca, 0xbb, 0x2a, 0xab, 0x8f, 0xba, 0x93, 0x39, 0xbf, 0xfc, 0xd5, 0xfe, 0x20, 0xfd, 0xf5, 0xe6,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x2a, 0x99, 0x08, 0x9a, 0x01, 0x00, 0x00,
}
