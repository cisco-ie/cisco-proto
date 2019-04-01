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
// source: ikev2_edm_policy.proto

package cisco_ios_xr_ikev2_oper_ik_ev2_nodes_node_policies_policy

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

type Ikev2EdmPolicy_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ikev2EdmPolicy_KEYS) Reset()         { *m = Ikev2EdmPolicy_KEYS{} }
func (m *Ikev2EdmPolicy_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ikev2EdmPolicy_KEYS) ProtoMessage()    {}
func (*Ikev2EdmPolicy_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b42b3dd28a2be5, []int{0}
}

func (m *Ikev2EdmPolicy_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ikev2EdmPolicy_KEYS.Unmarshal(m, b)
}
func (m *Ikev2EdmPolicy_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ikev2EdmPolicy_KEYS.Marshal(b, m, deterministic)
}
func (m *Ikev2EdmPolicy_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ikev2EdmPolicy_KEYS.Merge(m, src)
}
func (m *Ikev2EdmPolicy_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ikev2EdmPolicy_KEYS.Size(m)
}
func (m *Ikev2EdmPolicy_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ikev2EdmPolicy_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ikev2EdmPolicy_KEYS proto.InternalMessageInfo

func (m *Ikev2EdmPolicy_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ikev2EdmPolicy_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Ikev2EdmPolicy struct {
	PolicyName           string   `protobuf:"bytes,50,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Addr                 []string `protobuf:"bytes,51,rep,name=addr,proto3" json:"addr,omitempty"`
	Proposal             []string `protobuf:"bytes,52,rep,name=proposal,proto3" json:"proposal,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ikev2EdmPolicy) Reset()         { *m = Ikev2EdmPolicy{} }
func (m *Ikev2EdmPolicy) String() string { return proto.CompactTextString(m) }
func (*Ikev2EdmPolicy) ProtoMessage()    {}
func (*Ikev2EdmPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_65b42b3dd28a2be5, []int{1}
}

func (m *Ikev2EdmPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ikev2EdmPolicy.Unmarshal(m, b)
}
func (m *Ikev2EdmPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ikev2EdmPolicy.Marshal(b, m, deterministic)
}
func (m *Ikev2EdmPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ikev2EdmPolicy.Merge(m, src)
}
func (m *Ikev2EdmPolicy) XXX_Size() int {
	return xxx_messageInfo_Ikev2EdmPolicy.Size(m)
}
func (m *Ikev2EdmPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_Ikev2EdmPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_Ikev2EdmPolicy proto.InternalMessageInfo

func (m *Ikev2EdmPolicy) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *Ikev2EdmPolicy) GetAddr() []string {
	if m != nil {
		return m.Addr
	}
	return nil
}

func (m *Ikev2EdmPolicy) GetProposal() []string {
	if m != nil {
		return m.Proposal
	}
	return nil
}

func init() {
	proto.RegisterType((*Ikev2EdmPolicy_KEYS)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.policies.policy.ikev2_edm_policy_KEYS")
	proto.RegisterType((*Ikev2EdmPolicy)(nil), "cisco_ios_xr_ikev2_oper.ik_ev2.nodes.node.policies.policy.ikev2_edm_policy")
}

func init() { proto.RegisterFile("ikev2_edm_policy.proto", fileDescriptor_65b42b3dd28a2be5) }

var fileDescriptor_65b42b3dd28a2be5 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcc, 0x4e, 0x2d,
	0x33, 0x8a, 0x4f, 0x4d, 0xc9, 0x8d, 0x2f, 0xc8, 0xcf, 0xc9, 0x4c, 0xae, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xb2, 0x4c, 0xce, 0x2c, 0x4e, 0xce, 0x8f, 0xcf, 0xcc, 0x2f, 0x8e, 0xaf, 0x28,
	0x8a, 0x87, 0x28, 0xca, 0x2f, 0x48, 0x2d, 0xd2, 0xcb, 0xcc, 0x8e, 0x4f, 0x2d, 0x33, 0xd2, 0xcb,
	0xcb, 0x4f, 0x49, 0x2d, 0x06, 0x93, 0x7a, 0x60, 0x7d, 0x99, 0xa9, 0xc5, 0x10, 0x46, 0xa5, 0x92,
	0x07, 0x97, 0x28, 0xba, 0xa1, 0xf1, 0xde, 0xae, 0x91, 0xc1, 0x42, 0xd2, 0x5c, 0x9c, 0x20, 0x0d,
	0xf1, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0x20, 0x01, 0xbf,
	0xc4, 0xdc, 0x54, 0x21, 0x21, 0x2e, 0x16, 0xb0, 0x38, 0x13, 0x58, 0x1c, 0xcc, 0x56, 0x4a, 0xe6,
	0x12, 0x40, 0x37, 0x49, 0x48, 0x9e, 0x8b, 0x1b, 0x6a, 0x26, 0x58, 0xb9, 0x11, 0x58, 0x39, 0x17,
	0x44, 0x08, 0x66, 0x50, 0x62, 0x4a, 0x4a, 0x91, 0x84, 0xb1, 0x02, 0x33, 0xc8, 0x20, 0x10, 0x5b,
	0x48, 0x8a, 0x8b, 0xa3, 0xa0, 0x28, 0xbf, 0x20, 0xbf, 0x38, 0x31, 0x47, 0xc2, 0x04, 0x2c, 0x0e,
	0xe7, 0x27, 0xb1, 0x81, 0x3d, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x82, 0x3d, 0x9e, 0xdd,
	0x0a, 0x01, 0x00, 0x00,
}
