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
	AaMetricType             string   `protobuf:"bytes,58,opt,name=aa_metric_type,json=aaMetricType,proto3" json:"aa_metric_type,omitempty"`
	AaMetric                 int32    `protobuf:"zigzag32,59,opt,name=aa_metric,json=aaMetric,proto3" json:"aa_metric,omitempty"`
	ShortcutEligible         bool     `protobuf:"varint,60,opt,name=shortcut_eligible,json=shortcutEligible,proto3" json:"shortcut_eligible,omitempty"`
	ProtectionStyleRequested string   `protobuf:"bytes,61,opt,name=protection_style_requested,json=protectionStyleRequested,proto3" json:"protection_style_requested,omitempty"`
	ReoptimizeTimer          uint32   `protobuf:"varint,62,opt,name=reoptimize_timer,json=reoptimizeTimer,proto3" json:"reoptimize_timer,omitempty"`
	Role                     string   `protobuf:"bytes,63,opt,name=role,proto3" json:"role,omitempty"`
	SetupPriority            uint32   `protobuf:"varint,64,opt,name=setup_priority,json=setupPriority,proto3" json:"setup_priority,omitempty"`
	HoldPriority             uint32   `protobuf:"varint,65,opt,name=hold_priority,json=holdPriority,proto3" json:"hold_priority,omitempty"`
	SoftPreemption           bool     `protobuf:"varint,66,opt,name=soft_preemption,json=softPreemption,proto3" json:"soft_preemption,omitempty"`
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

func (m *OcMplsTeLspState) GetAaMetricType() string {
	if m != nil {
		return m.AaMetricType
	}
	return ""
}

func (m *OcMplsTeLspState) GetAaMetric() int32 {
	if m != nil {
		return m.AaMetric
	}
	return 0
}

func (m *OcMplsTeLspState) GetShortcutEligible() bool {
	if m != nil {
		return m.ShortcutEligible
	}
	return false
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
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4b, 0x6f, 0x13, 0x3d,
	0x14, 0x86, 0x95, 0x4f, 0x55, 0xbf, 0xd4, 0x4d, 0xd2, 0xc6, 0x42, 0xc8, 0x2d, 0x0b, 0x86, 0x02,
	0x22, 0x08, 0x91, 0x05, 0xe5, 0x4e, 0xb9, 0x4a, 0x5d, 0x00, 0x02, 0x45, 0x93, 0x6e, 0x58, 0x20,
	0xcb, 0xf1, 0x9c, 0xa6, 0x96, 0x3c, 0xb6, 0xb1, 0x3d, 0x12, 0xc3, 0x2f, 0xe6, 0x67, 0x20, 0x9f,
	0xb9, 0x24, 0x8b, 0xee, 0x8e, 0x9f, 0xf3, 0xf8, 0x78, 0xfc, 0x7a, 0xc8, 0xb1, 0x95, 0xbc, 0x74,
	0x3a, 0xf0, 0x08, 0x5c, 0x07, 0xc7, 0x43, 0x14, 0x11, 0xe6, 0xce, 0xdb, 0x68, 0xe9, 0x17, 0xa9,
	0x82, 0xb4, 0x5c, 0xd9, 0xc0, 0x7f, 0xfb, 0xde, 0xb2, 0x0e, 0xfc, 0xbc, 0x5b, 0x84, 0x28, 0x4c,
	0xb1, 0xaa, 0xe7, 0xd6, 0x81, 0xe1, 0xd2, 0x9a, 0x4b, 0xb5, 0x9e, 0xf7, 0x63, 0xc2, 0xa6, 0x3c,
	0xf9, 0x49, 0x8e, 0xae, 0x3b, 0x89, 0x7f, 0x3d, 0xff, 0xb1, 0xa4, 0xb7, 0xc9, 0x7e, 0xac, 0x8c,
	0x01, 0xcd, 0x8d, 0x28, 0x81, 0x0d, 0xb2, 0xc1, 0x6c, 0x2f, 0x27, 0x0d, 0xfa, 0x2e, 0x4a, 0xd8,
	0x12, 0x62, 0xed, 0x80, 0xfd, 0xb7, 0x2d, 0x5c, 0xd4, 0x0e, 0x4e, 0xfe, 0xee, 0x90, 0x1b, 0xd7,
	0xcd, 0xa7, 0x94, 0xec, 0xe0, 0xcc, 0x27, 0xb8, 0x05, 0xeb, 0xc4, 0x70, 0xcc, 0x69, 0xc3, 0x52,
	0x4d, 0x1f, 0x13, 0x1a, 0xd4, 0xda, 0x08, 0xad, 0xcc, 0x9a, 0xe3, 0xf5, 0xa5, 0xd5, 0xec, 0x29,
	0x1a, 0xd3, 0xbe, 0xb3, 0x68, 0x1b, 0xf4, 0x88, 0x0c, 0xb5, 0x95, 0x42, 0x73, 0x55, 0xb0, 0x67,
	0xd9, 0x60, 0x36, 0xce, 0xff, 0xc7, 0xf5, 0xe7, 0x82, 0xde, 0x24, 0xbb, 0xc1, 0x56, 0x5e, 0x02,
	0x7b, 0x8e, 0xbb, 0xdb, 0x15, 0xcd, 0xc8, 0x7e, 0x01, 0x41, 0x7a, 0xe5, 0xa2, 0xb2, 0x86, 0xbd,
	0xc0, 0xe6, 0x36, 0xa2, 0x77, 0xc8, 0x48, 0x14, 0xa5, 0x32, 0xf8, 0xe9, 0x55, 0x60, 0x2f, 0x1b,
	0x05, 0xd9, 0x12, 0x51, 0x0a, 0x22, 0xe5, 0xdf, 0x19, 0xaf, 0x9a, 0x20, 0x12, 0x6a, 0x85, 0x7b,
	0x64, 0x22, 0x04, 0x2f, 0x21, 0x7a, 0x25, 0x9b, 0xb0, 0x5e, 0xa3, 0x33, 0x12, 0xe2, 0x1b, 0xc2,
	0x14, 0x17, 0xbd, 0x45, 0xf6, 0x7a, 0x8b, 0xbd, 0xc9, 0x06, 0xb3, 0x69, 0x3e, 0xec, 0x04, 0xfa,
	0x88, 0x4c, 0xc3, 0x95, 0xf5, 0x51, 0x56, 0x91, 0x83, 0x56, 0x6b, 0xb5, 0xd2, 0xc0, 0xce, 0xb2,
	0xc1, 0x6c, 0x98, 0x1f, 0x76, 0x8d, 0xf3, 0x96, 0xd3, 0x33, 0x72, 0x9c, 0xd2, 0x02, 0x99, 0x6e,
	0xc0, 0x43, 0xac, 0x35, 0x70, 0x0f, 0xbf, 0x2a, 0x08, 0x11, 0x0a, 0xf6, 0x16, 0xcf, 0x66, 0x1b,
	0x63, 0x99, 0x84, 0xbc, 0xeb, 0xd3, 0x87, 0xe4, 0xd0, 0x83, 0x75, 0x51, 0x95, 0xea, 0x0f, 0xf0,
	0xa8, 0x4a, 0xf0, 0xec, 0x1d, 0xc6, 0x79, 0xb0, 0xe1, 0x17, 0x09, 0xa7, 0x47, 0xf3, 0x56, 0x03,
	0x7b, 0xdf, 0x3c, 0x5a, 0xaa, 0xe9, 0x7d, 0x32, 0x09, 0x10, 0x2b, 0xc7, 0x9d, 0x57, 0xd6, 0xab,
	0x58, 0xb3, 0x0f, 0xb8, 0x79, 0x8c, 0x74, 0xd1, 0x42, 0x7a, 0x97, 0x8c, 0xaf, 0xac, 0x2e, 0x36,
	0xd6, 0x47, 0xb4, 0x46, 0x09, 0xf6, 0xd2, 0x03, 0x72, 0x10, 0xec, 0x65, 0xe4, 0xce, 0x03, 0x94,
	0xcd, 0x13, 0x7d, 0xc2, 0x3b, 0x4f, 0x12, 0x5e, 0xf4, 0x74, 0xb5, 0x8b, 0x7f, 0xc7, 0xe9, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xed, 0x25, 0xa4, 0x3a, 0x03, 0x00, 0x00,
}