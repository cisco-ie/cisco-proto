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
// source: rsi_agent_state.proto

package cisco_ios_xr_rsi_agent_oper_rsi_agent_nodes_node_state

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

type RsiAgentState_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiAgentState_KEYS) Reset()         { *m = RsiAgentState_KEYS{} }
func (m *RsiAgentState_KEYS) String() string { return proto.CompactTextString(m) }
func (*RsiAgentState_KEYS) ProtoMessage()    {}
func (*RsiAgentState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a94dec556de376e, []int{0}
}

func (m *RsiAgentState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiAgentState_KEYS.Unmarshal(m, b)
}
func (m *RsiAgentState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiAgentState_KEYS.Marshal(b, m, deterministic)
}
func (m *RsiAgentState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiAgentState_KEYS.Merge(m, src)
}
func (m *RsiAgentState_KEYS) XXX_Size() int {
	return xxx_messageInfo_RsiAgentState_KEYS.Size(m)
}
func (m *RsiAgentState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiAgentState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RsiAgentState_KEYS proto.InternalMessageInfo

func (m *RsiAgentState_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type RsiAgentState struct {
	StateId              uint32   `protobuf:"varint,50,opt,name=state_id,json=stateId,proto3" json:"state_id,omitempty"`
	StateVal             uint32   `protobuf:"varint,51,opt,name=state_val,json=stateVal,proto3" json:"state_val,omitempty"`
	StateCtx             []uint32 `protobuf:"varint,52,rep,packed,name=state_ctx,json=stateCtx,proto3" json:"state_ctx,omitempty"`
	SvdIsEnabled         uint32   `protobuf:"varint,53,opt,name=svd_is_enabled,json=svdIsEnabled,proto3" json:"svd_is_enabled,omitempty"`
	IntfCapsV4Count      uint32   `protobuf:"varint,54,opt,name=intf_caps_v4_count,json=intfCapsV4Count,proto3" json:"intf_caps_v4_count,omitempty"`
	IntfCapsV6Count      uint32   `protobuf:"varint,55,opt,name=intf_caps_v6_count,json=intfCapsV6Count,proto3" json:"intf_caps_v6_count,omitempty"`
	VrfIntfCapsV4Count   uint32   `protobuf:"varint,56,opt,name=vrf_intf_caps_v4_count,json=vrfIntfCapsV4Count,proto3" json:"vrf_intf_caps_v4_count,omitempty"`
	VrfIntfCapsV6Count   uint32   `protobuf:"varint,57,opt,name=vrf_intf_caps_v6_count,json=vrfIntfCapsV6Count,proto3" json:"vrf_intf_caps_v6_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RsiAgentState) Reset()         { *m = RsiAgentState{} }
func (m *RsiAgentState) String() string { return proto.CompactTextString(m) }
func (*RsiAgentState) ProtoMessage()    {}
func (*RsiAgentState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a94dec556de376e, []int{1}
}

func (m *RsiAgentState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RsiAgentState.Unmarshal(m, b)
}
func (m *RsiAgentState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RsiAgentState.Marshal(b, m, deterministic)
}
func (m *RsiAgentState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RsiAgentState.Merge(m, src)
}
func (m *RsiAgentState) XXX_Size() int {
	return xxx_messageInfo_RsiAgentState.Size(m)
}
func (m *RsiAgentState) XXX_DiscardUnknown() {
	xxx_messageInfo_RsiAgentState.DiscardUnknown(m)
}

var xxx_messageInfo_RsiAgentState proto.InternalMessageInfo

func (m *RsiAgentState) GetStateId() uint32 {
	if m != nil {
		return m.StateId
	}
	return 0
}

func (m *RsiAgentState) GetStateVal() uint32 {
	if m != nil {
		return m.StateVal
	}
	return 0
}

func (m *RsiAgentState) GetStateCtx() []uint32 {
	if m != nil {
		return m.StateCtx
	}
	return nil
}

func (m *RsiAgentState) GetSvdIsEnabled() uint32 {
	if m != nil {
		return m.SvdIsEnabled
	}
	return 0
}

func (m *RsiAgentState) GetIntfCapsV4Count() uint32 {
	if m != nil {
		return m.IntfCapsV4Count
	}
	return 0
}

func (m *RsiAgentState) GetIntfCapsV6Count() uint32 {
	if m != nil {
		return m.IntfCapsV6Count
	}
	return 0
}

func (m *RsiAgentState) GetVrfIntfCapsV4Count() uint32 {
	if m != nil {
		return m.VrfIntfCapsV4Count
	}
	return 0
}

func (m *RsiAgentState) GetVrfIntfCapsV6Count() uint32 {
	if m != nil {
		return m.VrfIntfCapsV6Count
	}
	return 0
}

func init() {
	proto.RegisterType((*RsiAgentState_KEYS)(nil), "cisco_ios_xr_rsi_agent_oper.rsi_agent.nodes.node.state.rsi_agent_state_KEYS")
	proto.RegisterType((*RsiAgentState)(nil), "cisco_ios_xr_rsi_agent_oper.rsi_agent.nodes.node.state.rsi_agent_state")
}

func init() { proto.RegisterFile("rsi_agent_state.proto", fileDescriptor_8a94dec556de376e) }

var fileDescriptor_8a94dec556de376e = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0xd9, 0xef, 0x07, 0xba, 0x05, 0x75, 0x10, 0x54, 0x22, 0x5e, 0xc6, 0xf0, 0x30, 0x10,
	0x7a, 0x58, 0x6b, 0xd5, 0x73, 0xd9, 0xa1, 0x08, 0x1e, 0x26, 0x0c, 0x3c, 0x3d, 0xb2, 0x26, 0x95,
	0x40, 0x97, 0x94, 0xbc, 0x18, 0xfa, 0x57, 0xfa, 0x37, 0x49, 0x13, 0xdd, 0xa0, 0xf3, 0x12, 0x78,
	0xef, 0xfb, 0xf9, 0xbc, 0xef, 0x21, 0xe4, 0xca, 0xa2, 0x02, 0xfe, 0x21, 0xb5, 0x03, 0x74, 0xdc,
	0xc9, 0xa4, 0xb5, 0xc6, 0x19, 0x9a, 0x57, 0x0a, 0x2b, 0x03, 0xca, 0x20, 0x74, 0x16, 0x0e, 0x8c,
	0x69, 0xa5, 0x4d, 0xf6, 0x63, 0xa2, 0x8d, 0x90, 0x18, 0xde, 0x24, 0xd8, 0xf3, 0x94, 0x5c, 0x0e,
	0x0e, 0xc2, 0xcb, 0xea, 0xfd, 0x8d, 0xde, 0x92, 0x49, 0x4f, 0x81, 0xe6, 0x3b, 0xc9, 0x46, 0xb3,
	0xd1, 0x62, 0xb2, 0x1e, 0xf7, 0x8b, 0x57, 0xbe, 0x93, 0xf3, 0xaf, 0x7f, 0x64, 0x3a, 0xb0, 0xe8,
	0x0d, 0x19, 0x47, 0x5d, 0x09, 0xb6, 0x9c, 0x8d, 0x16, 0xe7, 0xeb, 0xd3, 0x30, 0x97, 0xa2, 0xbf,
	0x15, 0x23, 0xcf, 0x1b, 0x96, 0x86, 0x2c, 0xb2, 0x1b, 0xde, 0x1c, 0xc2, 0xca, 0x75, 0x2c, 0x9b,
	0xfd, 0xdf, 0x87, 0x85, 0xeb, 0xe8, 0x1d, 0xb9, 0x40, 0x2f, 0x40, 0x21, 0x48, 0xcd, 0xb7, 0x8d,
	0x14, 0xec, 0x21, 0xe8, 0x67, 0xe8, 0x45, 0x89, 0xab, 0xb8, 0xa3, 0xf7, 0x84, 0x2a, 0xed, 0x6a,
	0xa8, 0x78, 0x8b, 0xe0, 0x33, 0xa8, 0xcc, 0xa7, 0x76, 0x2c, 0x0f, 0xe4, 0xb4, 0x4f, 0x0a, 0xde,
	0xe2, 0x26, 0x2b, 0xfa, 0xf5, 0x00, 0xce, 0x7f, 0xe0, 0xc7, 0x01, 0x9c, 0x47, 0x78, 0x49, 0xae,
	0xbd, 0xad, 0xe1, 0x8f, 0xeb, 0x4f, 0x41, 0xa0, 0xde, 0xd6, 0xe5, 0xa0, 0xe0, 0xd8, 0xf9, 0x2d,
	0x79, 0x3e, 0x76, 0x62, 0xcf, 0xf6, 0x24, 0x7c, 0x62, 0xfa, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x23,
	0x8f, 0xdf, 0x79, 0xdd, 0x01, 0x00, 0x00,
}
