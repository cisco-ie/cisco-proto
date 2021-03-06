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

package cisco_ios_xr_mpls_te_oper_mpls_te_open_config_lsp_states_lsp_state

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
	proto.RegisterType((*OcMplsTeLspState_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.open_config.lsp_states.lsp_state.oc_mpls_te_lsp_state_KEYS")
	proto.RegisterType((*OcMplsTeLspState)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te.open_config.lsp_states.lsp_state.oc_mpls_te_lsp_state")
}

func init() { proto.RegisterFile("oc_mpls_te_lsp_state.proto", fileDescriptor_c6d718505ddbfb2f) }

var fileDescriptor_c6d718505ddbfb2f = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0x86, 0x15, 0x54, 0x95, 0x74, 0x9a, 0xa4, 0x8d, 0x85, 0x90, 0x5b, 0x0e, 0x2c, 0x05, 0x44,
	0x10, 0x22, 0x07, 0xca, 0x37, 0xe5, 0xab, 0x52, 0x0f, 0x08, 0x81, 0xa2, 0x4d, 0x2f, 0x1c, 0x90,
	0xe5, 0x7a, 0xa7, 0xa9, 0x25, 0xef, 0xda, 0xd8, 0x5e, 0x89, 0xf0, 0x8b, 0xf9, 0x19, 0x95, 0x67,
	0x37, 0x9b, 0x1c, 0x7a, 0x1b, 0x3f, 0xf3, 0x8c, 0xbd, 0x7e, 0xbd, 0x70, 0x68, 0x95, 0x28, 0x9d,
	0x09, 0x22, 0xa2, 0x30, 0xc1, 0x89, 0x10, 0x65, 0xc4, 0xa9, 0xf3, 0x36, 0x5a, 0x76, 0xaa, 0x74,
	0x50, 0x56, 0x68, 0x1b, 0xc4, 0x5f, 0xdf, 0x59, 0xd6, 0xa1, 0x9f, 0xb6, 0x8b, 0xa9, 0x75, 0x58,
	0x09, 0x65, 0xab, 0x4b, 0xbd, 0x98, 0x76, 0xe3, 0x61, 0x5d, 0x1e, 0xfd, 0x86, 0x83, 0x9b, 0x4e,
	0x10, 0xdf, 0xcf, 0x7e, 0xcd, 0xd9, 0x7d, 0xd8, 0x8d, 0x75, 0x55, 0xa1, 0x11, 0x95, 0x2c, 0x91,
	0xf7, 0xb2, 0xde, 0x64, 0x27, 0x87, 0x06, 0xfd, 0x94, 0x25, 0x6e, 0x08, 0x71, 0xe9, 0x90, 0xdf,
	0xda, 0x14, 0xce, 0x97, 0x0e, 0x8f, 0xfe, 0x6f, 0xc1, 0x9d, 0x9b, 0xf6, 0x67, 0x0c, 0xb6, 0x68,
	0xcf, 0x17, 0x34, 0x42, 0x75, 0x62, 0xb4, 0xcd, 0x71, 0xc3, 0x52, 0xcd, 0x9e, 0x03, 0x0b, 0x7a,
	0x51, 0x49, 0xa3, 0xab, 0x85, 0xa0, 0x6b, 0x2b, 0x6b, 0xf8, 0x4b, 0x32, 0xc6, 0x5d, 0x67, 0xd6,
	0x36, 0xd8, 0x01, 0xf4, 0x8d, 0x55, 0xd2, 0x08, 0x5d, 0xf0, 0x57, 0x59, 0x6f, 0x32, 0xcc, 0x6f,
	0xd3, 0xfa, 0x5b, 0xc1, 0xee, 0xc2, 0x76, 0xb0, 0xb5, 0x57, 0xc8, 0x5f, 0xd3, 0x74, 0xbb, 0x62,
	0x19, 0xec, 0x16, 0x18, 0x94, 0xd7, 0x2e, 0x6a, 0x5b, 0xf1, 0x37, 0xd4, 0xdc, 0x44, 0xec, 0x01,
	0x0c, 0x64, 0x51, 0xea, 0x8a, 0x3e, 0xbd, 0x0e, 0xfc, 0x6d, 0xa3, 0x10, 0x9b, 0x13, 0x4a, 0x41,
	0xa4, 0xdc, 0x57, 0xc6, 0xbb, 0x26, 0x88, 0x84, 0x5a, 0xe1, 0x11, 0x8c, 0xa4, 0x14, 0x25, 0x46,
	0xaf, 0x55, 0x13, 0xd6, 0x7b, 0x72, 0x06, 0x52, 0xfe, 0x20, 0x98, 0xe2, 0x62, 0xf7, 0x60, 0xa7,
	0xb3, 0xf8, 0x87, 0xac, 0x37, 0x19, 0xe7, 0xfd, 0x95, 0xc0, 0x9e, 0xc1, 0x38, 0x5c, 0x59, 0x1f,
	0x55, 0x1d, 0x05, 0x1a, 0xbd, 0xd0, 0x17, 0x06, 0xf9, 0x49, 0xd6, 0x9b, 0xf4, 0xf3, 0xfd, 0x55,
	0xe3, 0xac, 0xe5, 0xec, 0x04, 0x0e, 0x53, 0x5a, 0xa8, 0xd2, 0x0d, 0x44, 0x88, 0x4b, 0x83, 0xc2,
	0xe3, 0x9f, 0x1a, 0x43, 0xc4, 0x82, 0x7f, 0xa4, 0xb3, 0xf9, 0xda, 0x98, 0x27, 0x21, 0x5f, 0xf5,
	0xd9, 0x53, 0xd8, 0xf7, 0x68, 0x5d, 0xd4, 0xa5, 0xfe, 0x87, 0x22, 0xea, 0x12, 0x3d, 0xff, 0x44,
	0x71, 0xee, 0xad, 0xf9, 0x79, 0xc2, 0xe9, 0xd1, 0xbc, 0x35, 0xc8, 0x3f, 0x37, 0x8f, 0x96, 0x6a,
	0xf6, 0x18, 0x46, 0x01, 0x63, 0xed, 0x84, 0xf3, 0xda, 0x7a, 0x1d, 0x97, 0xfc, 0x0b, 0x0d, 0x0f,
	0x89, 0xce, 0x5a, 0xc8, 0x1e, 0xc2, 0xf0, 0xca, 0x9a, 0x62, 0x6d, 0x7d, 0x25, 0x6b, 0x90, 0x60,
	0x27, 0x3d, 0x81, 0xbd, 0x60, 0x2f, 0xa3, 0x70, 0x1e, 0xb1, 0x6c, 0x9e, 0xe8, 0x94, 0xee, 0x3c,
	0x4a, 0x78, 0xd6, 0xd1, 0x8b, 0x6d, 0xfa, 0x3b, 0x8e, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x55, 0xa4, 0x54, 0x32, 0x03, 0x00, 0x00,
}
