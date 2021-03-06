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
// source: l2vpn_globals.proto

package cisco_ios_xr_l2vpn_oper_l2vpnv2_active_global_settings

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

type L2VpnGlobals_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnGlobals_KEYS) Reset()         { *m = L2VpnGlobals_KEYS{} }
func (m *L2VpnGlobals_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnGlobals_KEYS) ProtoMessage()    {}
func (*L2VpnGlobals_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_74028bda9454b4c7, []int{0}
}

func (m *L2VpnGlobals_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGlobals_KEYS.Unmarshal(m, b)
}
func (m *L2VpnGlobals_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGlobals_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnGlobals_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGlobals_KEYS.Merge(m, src)
}
func (m *L2VpnGlobals_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnGlobals_KEYS.Size(m)
}
func (m *L2VpnGlobals_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGlobals_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGlobals_KEYS proto.InternalMessageInfo

type L2VpnGlobals struct {
	PwGroupingEnabled        bool     `protobuf:"varint,50,opt,name=pw_grouping_enabled,json=pwGroupingEnabled,proto3" json:"pw_grouping_enabled,omitempty"`
	PwStatusEnabled          bool     `protobuf:"varint,51,opt,name=pw_status_enabled,json=pwStatusEnabled,proto3" json:"pw_status_enabled,omitempty"`
	LoggingPwEnabled         bool     `protobuf:"varint,52,opt,name=logging_pw_enabled,json=loggingPwEnabled,proto3" json:"logging_pw_enabled,omitempty"`
	LoggingBdEnabled         bool     `protobuf:"varint,53,opt,name=logging_bd_enabled,json=loggingBdEnabled,proto3" json:"logging_bd_enabled,omitempty"`
	LoggingVfiEnabled        bool     `protobuf:"varint,54,opt,name=logging_vfi_enabled,json=loggingVfiEnabled,proto3" json:"logging_vfi_enabled,omitempty"`
	LoggingNsrEnabled        bool     `protobuf:"varint,55,opt,name=logging_nsr_enabled,json=loggingNsrEnabled,proto3" json:"logging_nsr_enabled,omitempty"`
	LoggingDfElectionEnabled bool     `protobuf:"varint,56,opt,name=logging_df_election_enabled,json=loggingDfElectionEnabled,proto3" json:"logging_df_election_enabled,omitempty"`
	TcnPropagationEnabled    bool     `protobuf:"varint,57,opt,name=tcn_propagation_enabled,json=tcnPropagationEnabled,proto3" json:"tcn_propagation_enabled,omitempty"`
	PwOamRefreshTransmitTime uint32   `protobuf:"varint,58,opt,name=pw_oam_refresh_transmit_time,json=pwOamRefreshTransmitTime,proto3" json:"pw_oam_refresh_transmit_time,omitempty"`
	HaRole                   string   `protobuf:"bytes,59,opt,name=ha_role,json=haRole,proto3" json:"ha_role,omitempty"`
	IssuRole                 string   `protobuf:"bytes,60,opt,name=issu_role,json=issuRole,proto3" json:"issu_role,omitempty"`
	ProcessFsm               string   `protobuf:"bytes,61,opt,name=process_fsm,json=processFsm,proto3" json:"process_fsm,omitempty"`
	GoingActive              bool     `protobuf:"varint,62,opt,name=going_active,json=goingActive,proto3" json:"going_active,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *L2VpnGlobals) Reset()         { *m = L2VpnGlobals{} }
func (m *L2VpnGlobals) String() string { return proto.CompactTextString(m) }
func (*L2VpnGlobals) ProtoMessage()    {}
func (*L2VpnGlobals) Descriptor() ([]byte, []int) {
	return fileDescriptor_74028bda9454b4c7, []int{1}
}

func (m *L2VpnGlobals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnGlobals.Unmarshal(m, b)
}
func (m *L2VpnGlobals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnGlobals.Marshal(b, m, deterministic)
}
func (m *L2VpnGlobals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnGlobals.Merge(m, src)
}
func (m *L2VpnGlobals) XXX_Size() int {
	return xxx_messageInfo_L2VpnGlobals.Size(m)
}
func (m *L2VpnGlobals) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnGlobals.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnGlobals proto.InternalMessageInfo

func (m *L2VpnGlobals) GetPwGroupingEnabled() bool {
	if m != nil {
		return m.PwGroupingEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetPwStatusEnabled() bool {
	if m != nil {
		return m.PwStatusEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetLoggingPwEnabled() bool {
	if m != nil {
		return m.LoggingPwEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetLoggingBdEnabled() bool {
	if m != nil {
		return m.LoggingBdEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetLoggingVfiEnabled() bool {
	if m != nil {
		return m.LoggingVfiEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetLoggingNsrEnabled() bool {
	if m != nil {
		return m.LoggingNsrEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetLoggingDfElectionEnabled() bool {
	if m != nil {
		return m.LoggingDfElectionEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetTcnPropagationEnabled() bool {
	if m != nil {
		return m.TcnPropagationEnabled
	}
	return false
}

func (m *L2VpnGlobals) GetPwOamRefreshTransmitTime() uint32 {
	if m != nil {
		return m.PwOamRefreshTransmitTime
	}
	return 0
}

func (m *L2VpnGlobals) GetHaRole() string {
	if m != nil {
		return m.HaRole
	}
	return ""
}

func (m *L2VpnGlobals) GetIssuRole() string {
	if m != nil {
		return m.IssuRole
	}
	return ""
}

func (m *L2VpnGlobals) GetProcessFsm() string {
	if m != nil {
		return m.ProcessFsm
	}
	return ""
}

func (m *L2VpnGlobals) GetGoingActive() bool {
	if m != nil {
		return m.GoingActive
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnGlobals_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.global_settings.l2vpn_globals_KEYS")
	proto.RegisterType((*L2VpnGlobals)(nil), "cisco_ios_xr_l2vpn_oper.l2vpnv2.active.global_settings.l2vpn_globals")
}

func init() { proto.RegisterFile("l2vpn_globals.proto", fileDescriptor_74028bda9454b4c7) }

var fileDescriptor_74028bda9454b4c7 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0x5d, 0x8b, 0xd4, 0x30,
	0x18, 0x05, 0x60, 0x06, 0x64, 0xdd, 0xcd, 0xba, 0xe8, 0x66, 0x94, 0x2d, 0xac, 0xe0, 0xb8, 0x57,
	0x83, 0x48, 0x2f, 0x66, 0x75, 0xfc, 0xdc, 0x05, 0xc5, 0xd1, 0x0b, 0x41, 0x97, 0xce, 0x20, 0x78,
	0xf5, 0x92, 0xb6, 0x69, 0x26, 0x90, 0x26, 0x21, 0x6f, 0xa6, 0xf1, 0xa7, 0xfa, 0x73, 0x64, 0xd2,
	0x0f, 0x29, 0xcc, 0x5d, 0x39, 0xe7, 0x39, 0xa1, 0x2d, 0x21, 0x53, 0xb5, 0x68, 0xac, 0x06, 0xa1,
	0x4c, 0xce, 0x14, 0xa6, 0xd6, 0x19, 0x6f, 0xe8, 0xb2, 0x90, 0x58, 0x18, 0x90, 0x06, 0xe1, 0x8f,
	0x83, 0x56, 0x18, 0xcb, 0x5d, 0x1a, 0x1f, 0x9b, 0x45, 0xca, 0x0a, 0x2f, 0x1b, 0x9e, 0xb6, 0x2b,
	0x40, 0xee, 0xbd, 0xd4, 0x02, 0xaf, 0x1e, 0x13, 0x3a, 0x3a, 0x0e, 0xbe, 0xaf, 0x7e, 0xaf, 0xaf,
	0xfe, 0xde, 0x23, 0x67, 0xa3, 0x98, 0xa6, 0x64, 0x6a, 0x03, 0x08, 0x67, 0x76, 0x56, 0x6a, 0x01,
	0x5c, 0xb3, 0x5c, 0xf1, 0x32, 0x59, 0xcc, 0x26, 0xf3, 0xe3, 0xec, 0xdc, 0x86, 0x6f, 0x5d, 0xb3,
	0x6a, 0x0b, 0xfa, 0x82, 0x9c, 0xdb, 0x00, 0xe8, 0x99, 0xdf, 0xe1, 0xa0, 0xaf, 0xa3, 0x7e, 0x68,
	0xc3, 0x3a, 0xe6, 0xbd, 0x7d, 0x49, 0xa8, 0x32, 0x42, 0xec, 0xcf, 0xb5, 0x61, 0xc0, 0xaf, 0x22,
	0x7e, 0xd4, 0x35, 0x77, 0xe1, 0x80, 0xce, 0xcb, 0x41, 0xbf, 0x1e, 0xe9, 0xcf, 0x65, 0xaf, 0x53,
	0x32, 0xed, 0x75, 0x53, 0xc9, 0x81, 0x2f, 0xdb, 0xf7, 0xee, 0xaa, 0x5f, 0x95, 0x3c, 0xe0, 0x35,
	0xba, 0xc1, 0xbf, 0x19, 0xf9, 0x1f, 0xe8, 0x7a, 0x7f, 0x43, 0x2e, 0x7b, 0x5f, 0x56, 0xc0, 0x15,
	0x2f, 0xbc, 0x34, 0x7a, 0xd8, 0xbd, 0x8d, 0xbb, 0xa4, 0x23, 0x5f, 0xaa, 0x55, 0x07, 0xfa, 0xf9,
	0x92, 0x5c, 0xf8, 0x42, 0x83, 0x75, 0xc6, 0x32, 0xc1, 0x46, 0xd3, 0x77, 0x71, 0xfa, 0xc4, 0x17,
	0xfa, 0xee, 0x7f, 0xdb, 0xef, 0x6e, 0xc9, 0x53, 0x1b, 0xc0, 0xb0, 0x1a, 0x1c, 0xaf, 0x1c, 0xc7,
	0x2d, 0x78, 0xc7, 0x34, 0xd6, 0xd2, 0x83, 0x97, 0x35, 0x4f, 0xde, 0xcf, 0x26, 0xf3, 0xb3, 0x2c,
	0xb1, 0xe1, 0x27, 0xab, 0xb3, 0x56, 0x6c, 0x3a, 0xb0, 0x91, 0x35, 0xa7, 0x17, 0xe4, 0xfe, 0x96,
	0x81, 0x33, 0x8a, 0x27, 0x1f, 0x66, 0x93, 0xf9, 0x49, 0x76, 0xb4, 0x65, 0x99, 0x51, 0x9c, 0x5e,
	0x92, 0x13, 0x89, 0xb8, 0x6b, 0xab, 0x8f, 0xb1, 0x3a, 0xde, 0x07, 0xb1, 0x7c, 0x46, 0x4e, 0xad,
	0x33, 0x05, 0x47, 0x84, 0x0a, 0xeb, 0xe4, 0x26, 0xd6, 0xa4, 0x8b, 0xbe, 0x62, 0x4d, 0x9f, 0x93,
	0x07, 0xc2, 0xec, 0xff, 0x45, 0x7b, 0xdb, 0x92, 0xdb, 0xf8, 0x0d, 0xa7, 0x31, 0xfb, 0x14, 0xa3,
	0xfc, 0x28, 0xde, 0xd7, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0xa0, 0x7b, 0xc8, 0xc6,
	0x02, 0x00, 0x00,
}
