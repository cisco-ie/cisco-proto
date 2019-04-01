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
// source: oc_mpls_te_lsp_state.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_open_config_lsp_states_lsp_state

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

type OcMplsTeLspState_KEYS struct {
	TunnelName           string   `protobuf:"bytes,1,opt,name=tunnel_name,json=tunnelName,proto3" json:"tunnel_name,omitempty"`
	TunnelType           string   `protobuf:"bytes,2,opt,name=tunnel_type,json=tunnelType,proto3" json:"tunnel_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OcMplsTeLspState_KEYS) Reset()         { *m = OcMplsTeLspState_KEYS{} }
func (m *OcMplsTeLspState_KEYS) String() string { return proto.CompactTextString(m) }
func (*OcMplsTeLspState_KEYS) ProtoMessage()    {}
func (*OcMplsTeLspState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d718505ddbfb2f, []int{0}
}

func (m *OcMplsTeLspState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OcMplsTeLspState_KEYS.Unmarshal(m, b)
}
func (m *OcMplsTeLspState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OcMplsTeLspState_KEYS.Marshal(b, m, deterministic)
}
func (m *OcMplsTeLspState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OcMplsTeLspState_KEYS.Merge(m, src)
}
func (m *OcMplsTeLspState_KEYS) XXX_Size() int {
	return xxx_messageInfo_OcMplsTeLspState_KEYS.Size(m)
}
func (m *OcMplsTeLspState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OcMplsTeLspState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OcMplsTeLspState_KEYS proto.InternalMessageInfo

func (m *OcMplsTeLspState_KEYS) GetTunnelName() string {
	if m != nil {
		return m.TunnelName
	}
	return ""
}

func (m *OcMplsTeLspState_KEYS) GetTunnelType() string {
	if m != nil {
		return m.TunnelType
	}
	return ""
}

type OcMplsTeLspState struct {
	Name                     string   `protobuf:"bytes,50,opt,name=name,proto3" json:"name,omitempty"`
	Type                     string   `protobuf:"bytes,51,opt,name=type,proto3" json:"type,omitempty"`
	SignalingProtocol        string   `protobuf:"bytes,52,opt,name=signaling_protocol,json=signalingProtocol,proto3" json:"signaling_protocol,omitempty"`
	LocalId                  uint32   `protobuf:"varint,53,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	Source                   string   `protobuf:"bytes,54,opt,name=source,proto3" json:"source,omitempty"`
	Description              string   `protobuf:"bytes,55,opt,name=description,proto3" json:"description,omitempty"`
	AdminStatus              string   `protobuf:"bytes,56,opt,name=admin_status,json=adminStatus,proto3" json:"admin_status,omitempty"`
	OperStatus               string   `protobuf:"bytes,57,opt,name=oper_status,json=operStatus,proto3" json:"oper_status,omitempty"`
	Metric                   string   `protobuf:"bytes,58,opt,name=metric,proto3" json:"metric,omitempty"`
	ProtectionStyleRequested string   `protobuf:"bytes,59,opt,name=protection_style_requested,json=protectionStyleRequested,proto3" json:"protection_style_requested,omitempty"`
	ReoptimizeTimer          uint32   `protobuf:"varint,60,opt,name=reoptimize_timer,json=reoptimizeTimer,proto3" json:"reoptimize_timer,omitempty"`
	Role                     string   `protobuf:"bytes,61,opt,name=role,proto3" json:"role,omitempty"`
	SetupPriority            uint32   `protobuf:"varint,62,opt,name=setup_priority,json=setupPriority,proto3" json:"setup_priority,omitempty"`
	HoldPriority             uint32   `protobuf:"varint,63,opt,name=hold_priority,json=holdPriority,proto3" json:"hold_priority,omitempty"`
	SoftPreemption           bool     `protobuf:"varint,64,opt,name=soft_preemption,json=softPreemption,proto3" json:"soft_preemption,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *OcMplsTeLspState) Reset()         { *m = OcMplsTeLspState{} }
func (m *OcMplsTeLspState) String() string { return proto.CompactTextString(m) }
func (*OcMplsTeLspState) ProtoMessage()    {}
func (*OcMplsTeLspState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6d718505ddbfb2f, []int{1}
}

func (m *OcMplsTeLspState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OcMplsTeLspState.Unmarshal(m, b)
}
func (m *OcMplsTeLspState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OcMplsTeLspState.Marshal(b, m, deterministic)
}
func (m *OcMplsTeLspState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OcMplsTeLspState.Merge(m, src)
}
func (m *OcMplsTeLspState) XXX_Size() int {
	return xxx_messageInfo_OcMplsTeLspState.Size(m)
}
func (m *OcMplsTeLspState) XXX_DiscardUnknown() {
	xxx_messageInfo_OcMplsTeLspState.DiscardUnknown(m)
}

var xxx_messageInfo_OcMplsTeLspState proto.InternalMessageInfo

func (m *OcMplsTeLspState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OcMplsTeLspState) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OcMplsTeLspState) GetSignalingProtocol() string {
	if m != nil {
		return m.SignalingProtocol
	}
	return ""
}

func (m *OcMplsTeLspState) GetLocalId() uint32 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *OcMplsTeLspState) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *OcMplsTeLspState) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *OcMplsTeLspState) GetAdminStatus() string {
	if m != nil {
		return m.AdminStatus
	}
	return ""
}

func (m *OcMplsTeLspState) GetOperStatus() string {
	if m != nil {
		return m.OperStatus
	}
	return ""
}

func (m *OcMplsTeLspState) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *OcMplsTeLspState) GetProtectionStyleRequested() string {
	if m != nil {
		return m.ProtectionStyleRequested
	}
	return ""
}

func (m *OcMplsTeLspState) GetReoptimizeTimer() uint32 {
	if m != nil {
		return m.ReoptimizeTimer
	}
	return 0
}

func (m *OcMplsTeLspState) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *OcMplsTeLspState) GetSetupPriority() uint32 {
	if m != nil {
		return m.SetupPriority
	}
	return 0
}

func (m *OcMplsTeLspState) GetHoldPriority() uint32 {
	if m != nil {
		return m.HoldPriority
	}
	return 0
}

func (m *OcMplsTeLspState) GetSoftPreemption() bool {
	if m != nil {
		return m.SoftPreemption
	}
	return false
}

func init() {
	proto.RegisterType((*OcMplsTeLspState_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.open_config.lsp_states.lsp_state.oc_mpls_te_lsp_state_KEYS")
	proto.RegisterType((*OcMplsTeLspState)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.open_config.lsp_states.lsp_state.oc_mpls_te_lsp_state")
}

func init() { proto.RegisterFile("oc_mpls_te_lsp_state.proto", fileDescriptor_c6d718505ddbfb2f) }

var fileDescriptor_c6d718505ddbfb2f = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x86, 0x15, 0x54, 0x85, 0x32, 0x6d, 0x5a, 0xb0, 0x50, 0xe5, 0xf6, 0x42, 0x28, 0x42, 0x84,
	0x03, 0x39, 0x50, 0xbe, 0x29, 0x1f, 0x17, 0x0e, 0x80, 0x84, 0xa2, 0xa4, 0x17, 0x0e, 0xc8, 0xda,
	0x7a, 0xa7, 0xc1, 0x92, 0xd7, 0x63, 0x6c, 0xaf, 0xc4, 0xf2, 0x2b, 0xf8, 0xc9, 0x95, 0xc7, 0xc9,
	0x26, 0x87, 0xde, 0x66, 0x9e, 0x79, 0x3c, 0x5e, 0xbf, 0x0b, 0x27, 0xa4, 0x55, 0xe3, 0x6d, 0x54,
	0x09, 0x95, 0x8d, 0x5e, 0xc5, 0x54, 0x25, 0x9c, 0xfa, 0x40, 0x89, 0xc4, 0x37, 0x6d, 0xa2, 0x26,
	0x65, 0x28, 0xaa, 0xbf, 0xa1, 0xb7, 0xc8, 0x63, 0x98, 0xae, 0x9b, 0x98, 0x2a, 0x57, 0x5f, 0x76,
	0x53, 0xf2, 0xe8, 0x94, 0x26, 0x77, 0x65, 0x96, 0xd3, 0x7e, 0x4d, 0xdc, 0x94, 0xa7, 0xbf, 0xe0,
	0xf8, 0xa6, 0x9b, 0xd4, 0xf7, 0x2f, 0x3f, 0x17, 0xe2, 0x01, 0xec, 0xa5, 0xd6, 0x39, 0xb4, 0xca,
	0x55, 0x0d, 0xca, 0xc1, 0x78, 0x30, 0xb9, 0x33, 0x87, 0x82, 0x7e, 0x54, 0x0d, 0x6e, 0x09, 0xa9,
	0xf3, 0x28, 0x6f, 0x6d, 0x0b, 0x17, 0x9d, 0xc7, 0xd3, 0xff, 0x3b, 0x70, 0xff, 0xa6, 0xfd, 0x42,
	0xc0, 0x0e, 0xef, 0x7c, 0xce, 0x47, 0xb8, 0xce, 0x8c, 0xd7, 0x9c, 0x15, 0x96, 0x6b, 0xf1, 0x0c,
	0x44, 0x34, 0x4b, 0x57, 0x59, 0xe3, 0x96, 0x8a, 0x9f, 0xaf, 0xc9, 0xca, 0x17, 0x6c, 0xdc, 0xeb,
	0x27, 0xb3, 0xd5, 0x40, 0x1c, 0xc3, 0xae, 0x25, 0x5d, 0x59, 0x65, 0x6a, 0xf9, 0x72, 0x3c, 0x98,
	0x8c, 0xe6, 0xb7, 0xb9, 0xff, 0x5a, 0x8b, 0x23, 0x18, 0x46, 0x6a, 0x83, 0x46, 0xf9, 0x8a, 0x4f,
	0xaf, 0x3a, 0x31, 0x86, 0xbd, 0x1a, 0xa3, 0x0e, 0xc6, 0x27, 0x43, 0x4e, 0xbe, 0xe6, 0xe1, 0x36,
	0x12, 0x0f, 0x61, 0xbf, 0xaa, 0x1b, 0xe3, 0xf8, 0xd3, 0xdb, 0x28, 0xdf, 0x14, 0x85, 0xd9, 0x82,
	0x51, 0x0e, 0x22, 0xe7, 0xbf, 0x36, 0xde, 0x96, 0x20, 0x32, 0x5a, 0x09, 0x47, 0x30, 0x6c, 0x30,
	0x05, 0xa3, 0xe5, 0xbb, 0x72, 0x7b, 0xe9, 0xc4, 0x39, 0x9c, 0xe4, 0x57, 0xa1, 0xce, 0x37, 0xa9,
	0x98, 0x3a, 0x8b, 0x2a, 0xe0, 0x9f, 0x16, 0x63, 0xc2, 0x5a, 0xbe, 0x67, 0x57, 0x6e, 0x8c, 0x45,
	0x16, 0xe6, 0xeb, 0xb9, 0x78, 0x0a, 0x77, 0x03, 0x92, 0x4f, 0xa6, 0x31, 0xff, 0x50, 0x25, 0xd3,
	0x60, 0x90, 0xe7, 0xfc, 0xec, 0xc3, 0x0d, 0xbf, 0xc8, 0x38, 0x87, 0x1b, 0xc8, 0xa2, 0xfc, 0x50,
	0xc2, 0xcd, 0xb5, 0x78, 0x0c, 0x07, 0x11, 0x53, 0xeb, 0x95, 0x0f, 0x86, 0x82, 0x49, 0x9d, 0xfc,
	0xc8, 0x87, 0x47, 0x4c, 0x67, 0x2b, 0x28, 0x1e, 0xc1, 0xe8, 0x37, 0xd9, 0x7a, 0x63, 0x7d, 0x62,
	0x6b, 0x3f, 0xc3, 0x5e, 0x7a, 0x02, 0x87, 0x91, 0xae, 0x92, 0xf2, 0x01, 0xb1, 0x29, 0x51, 0x7e,
	0x1e, 0x0f, 0x26, 0xbb, 0xf3, 0x83, 0x8c, 0x67, 0x3d, 0xbd, 0x1c, 0xf2, 0x5f, 0x3c, 0xbb, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x20, 0x39, 0x42, 0xdd, 0xe2, 0x02, 0x00, 0x00,
}