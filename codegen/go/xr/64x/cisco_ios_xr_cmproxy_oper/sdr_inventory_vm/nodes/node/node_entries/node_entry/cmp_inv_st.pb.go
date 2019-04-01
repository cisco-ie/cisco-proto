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
// source: cmp_inv_st.proto

package cisco_ios_xr_cmproxy_oper_sdr_inventory_vm_nodes_node_node_entries_node_entry

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

type CmpInvSt_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Name_1               string   `protobuf:"bytes,2,opt,name=name_1,json=name1,proto3" json:"name_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmpInvSt_KEYS) Reset()         { *m = CmpInvSt_KEYS{} }
func (m *CmpInvSt_KEYS) String() string { return proto.CompactTextString(m) }
func (*CmpInvSt_KEYS) ProtoMessage()    {}
func (*CmpInvSt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7a33142286073de, []int{0}
}

func (m *CmpInvSt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmpInvSt_KEYS.Unmarshal(m, b)
}
func (m *CmpInvSt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmpInvSt_KEYS.Marshal(b, m, deterministic)
}
func (m *CmpInvSt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmpInvSt_KEYS.Merge(m, src)
}
func (m *CmpInvSt_KEYS) XXX_Size() int {
	return xxx_messageInfo_CmpInvSt_KEYS.Size(m)
}
func (m *CmpInvSt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_CmpInvSt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_CmpInvSt_KEYS proto.InternalMessageInfo

func (m *CmpInvSt_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CmpInvSt_KEYS) GetName_1() string {
	if m != nil {
		return m.Name_1
	}
	return ""
}

type CmpInvSt struct {
	Valid                uint32   `protobuf:"varint,50,opt,name=valid,proto3" json:"valid,omitempty"`
	CardType             uint32   `protobuf:"varint,51,opt,name=card_type,json=cardType,proto3" json:"card_type,omitempty"`
	CardTypeString       []string `protobuf:"bytes,52,rep,name=card_type_string,json=cardTypeString,proto3" json:"card_type_string,omitempty"`
	Nodeid               int32    `protobuf:"zigzag32,53,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	NodeName             []string `protobuf:"bytes,54,rep,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	PartnerId            int32    `protobuf:"zigzag32,55,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"`
	PartnerName          []string `protobuf:"bytes,56,rep,name=partner_name,json=partnerName,proto3" json:"partner_name,omitempty"`
	RedState             uint32   `protobuf:"varint,57,opt,name=red_state,json=redState,proto3" json:"red_state,omitempty"`
	RedStateString       []string `protobuf:"bytes,58,rep,name=red_state_string,json=redStateString,proto3" json:"red_state_string,omitempty"`
	NodeSwState          uint32   `protobuf:"varint,59,opt,name=node_sw_state,json=nodeSwState,proto3" json:"node_sw_state,omitempty"`
	NodeSwStateString    []string `protobuf:"bytes,60,rep,name=node_sw_state_string,json=nodeSwStateString,proto3" json:"node_sw_state_string,omitempty"`
	PrevSwState          uint32   `protobuf:"varint,61,opt,name=prev_sw_state,json=prevSwState,proto3" json:"prev_sw_state,omitempty"`
	PrevSwStateString    []string `protobuf:"bytes,62,rep,name=prev_sw_state_string,json=prevSwStateString,proto3" json:"prev_sw_state_string,omitempty"`
	NodeIp               uint32   `protobuf:"varint,63,opt,name=node_ip,json=nodeIp,proto3" json:"node_ip,omitempty"`
	NodeIpv4String       []string `protobuf:"bytes,64,rep,name=node_ipv4_string,json=nodeIpv4String,proto3" json:"node_ipv4_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmpInvSt) Reset()         { *m = CmpInvSt{} }
func (m *CmpInvSt) String() string { return proto.CompactTextString(m) }
func (*CmpInvSt) ProtoMessage()    {}
func (*CmpInvSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7a33142286073de, []int{1}
}

func (m *CmpInvSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmpInvSt.Unmarshal(m, b)
}
func (m *CmpInvSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmpInvSt.Marshal(b, m, deterministic)
}
func (m *CmpInvSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmpInvSt.Merge(m, src)
}
func (m *CmpInvSt) XXX_Size() int {
	return xxx_messageInfo_CmpInvSt.Size(m)
}
func (m *CmpInvSt) XXX_DiscardUnknown() {
	xxx_messageInfo_CmpInvSt.DiscardUnknown(m)
}

var xxx_messageInfo_CmpInvSt proto.InternalMessageInfo

func (m *CmpInvSt) GetValid() uint32 {
	if m != nil {
		return m.Valid
	}
	return 0
}

func (m *CmpInvSt) GetCardType() uint32 {
	if m != nil {
		return m.CardType
	}
	return 0
}

func (m *CmpInvSt) GetCardTypeString() []string {
	if m != nil {
		return m.CardTypeString
	}
	return nil
}

func (m *CmpInvSt) GetNodeid() int32 {
	if m != nil {
		return m.Nodeid
	}
	return 0
}

func (m *CmpInvSt) GetNodeName() []string {
	if m != nil {
		return m.NodeName
	}
	return nil
}

func (m *CmpInvSt) GetPartnerId() int32 {
	if m != nil {
		return m.PartnerId
	}
	return 0
}

func (m *CmpInvSt) GetPartnerName() []string {
	if m != nil {
		return m.PartnerName
	}
	return nil
}

func (m *CmpInvSt) GetRedState() uint32 {
	if m != nil {
		return m.RedState
	}
	return 0
}

func (m *CmpInvSt) GetRedStateString() []string {
	if m != nil {
		return m.RedStateString
	}
	return nil
}

func (m *CmpInvSt) GetNodeSwState() uint32 {
	if m != nil {
		return m.NodeSwState
	}
	return 0
}

func (m *CmpInvSt) GetNodeSwStateString() []string {
	if m != nil {
		return m.NodeSwStateString
	}
	return nil
}

func (m *CmpInvSt) GetPrevSwState() uint32 {
	if m != nil {
		return m.PrevSwState
	}
	return 0
}

func (m *CmpInvSt) GetPrevSwStateString() []string {
	if m != nil {
		return m.PrevSwStateString
	}
	return nil
}

func (m *CmpInvSt) GetNodeIp() uint32 {
	if m != nil {
		return m.NodeIp
	}
	return 0
}

func (m *CmpInvSt) GetNodeIpv4String() []string {
	if m != nil {
		return m.NodeIpv4String
	}
	return nil
}

func init() {
	proto.RegisterType((*CmpInvSt_KEYS)(nil), "cisco_ios_xr_cmproxy_oper.sdr_inventory_vm.nodes.node.node_entries.node_entry.cmp_inv_st_KEYS")
	proto.RegisterType((*CmpInvSt)(nil), "cisco_ios_xr_cmproxy_oper.sdr_inventory_vm.nodes.node.node_entries.node_entry.cmp_inv_st")
}

func init() { proto.RegisterFile("cmp_inv_st.proto", fileDescriptor_b7a33142286073de) }

var fileDescriptor_b7a33142286073de = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x15, 0xd8, 0x2d, 0x9b, 0x59, 0x16, 0xb6, 0x56, 0x01, 0x4b, 0x15, 0x52, 0xe8, 0x29,
	0xa7, 0xa2, 0xd2, 0xf2, 0xbf, 0xfc, 0xb9, 0x70, 0xa8, 0x10, 0x1c, 0x12, 0x2e, 0x9c, 0xac, 0x10,
	0x5b, 0xc8, 0x12, 0x89, 0x2d, 0xdb, 0x4a, 0x9b, 0xef, 0xc1, 0x07, 0x5e, 0x79, 0xec, 0x34, 0xed,
	0xc5, 0xc9, 0xbc, 0x79, 0xf3, 0xcb, 0x9b, 0xc8, 0x70, 0x5b, 0x37, 0x9a, 0xc9, 0xb6, 0x63, 0xd6,
	0x2d, 0xb5, 0x51, 0x4e, 0x91, 0x1f, 0xb5, 0xb4, 0xb5, 0x62, 0x52, 0x59, 0x76, 0x30, 0xac, 0x6e,
	0xb4, 0x51, 0x87, 0x9e, 0x29, 0x2d, 0xcc, 0xd2, 0x72, 0xe3, 0xbd, 0xa2, 0x75, 0xca, 0xf4, 0xac,
	0x6b, 0x96, 0xad, 0xe2, 0xc2, 0xe2, 0x89, 0x07, 0x13, 0xad, 0x33, 0x32, 0x2a, 0x58, 0xf4, 0x8b,
	0x2d, 0x3c, 0x1e, 0x3f, 0xc1, 0xbe, 0x7f, 0xfb, 0x5d, 0x12, 0x02, 0x17, 0x6d, 0xd5, 0x08, 0x9a,
	0x64, 0x49, 0x9e, 0x16, 0xf8, 0x4e, 0x9e, 0xc0, 0xc4, 0x3f, 0xd9, 0x8a, 0xde, 0x43, 0xf5, 0xd2,
	0x57, 0xab, 0xc5, 0xff, 0x0b, 0x80, 0x71, 0x9c, 0xcc, 0xe0, 0xb2, 0xab, 0xfe, 0x49, 0x4e, 0x5f,
	0x65, 0x49, 0x7e, 0x53, 0x84, 0x82, 0xcc, 0x21, 0xad, 0x2b, 0xc3, 0x99, 0xeb, 0xb5, 0xa0, 0x6b,
	0xec, 0x5c, 0x79, 0xe1, 0x57, 0xaf, 0x05, 0xc9, 0xe1, 0xf6, 0xd8, 0x64, 0xd6, 0x19, 0xd9, 0xfe,
	0xa5, 0x9b, 0xec, 0x7e, 0x9e, 0x16, 0x8f, 0x06, 0x4f, 0x89, 0x2a, 0x79, 0x0a, 0x13, 0x9f, 0x5b,
	0x72, 0xfa, 0x3a, 0x4b, 0xf2, 0x69, 0x11, 0x2b, 0x8f, 0xc7, 0x7d, 0x30, 0xf3, 0x1b, 0x1c, 0xbd,
	0xf2, 0xc2, 0x4f, 0x9f, 0xfb, 0x39, 0x80, 0xae, 0x8c, 0x6b, 0x85, 0x61, 0x92, 0xd3, 0xb7, 0x38,
	0x98, 0x46, 0x65, 0xc7, 0xc9, 0x0b, 0x78, 0x38, 0xb4, 0x71, 0xfc, 0x1d, 0x8e, 0x5f, 0x47, 0x0d,
	0x09, 0x73, 0x48, 0x8d, 0xe0, 0xcc, 0xba, 0xca, 0x09, 0xfa, 0x3e, 0xa4, 0x37, 0x82, 0x97, 0xbe,
	0xf6, 0xe9, 0x8f, 0xcd, 0x21, 0xfd, 0x87, 0x90, 0x7e, 0xf0, 0xc4, 0xf4, 0x0b, 0xb8, 0xc1, 0x94,
	0x76, 0x1f, 0x51, 0x1f, 0x11, 0x75, 0xed, 0xc5, 0x72, 0x1f, 0x68, 0x2f, 0x61, 0x76, 0xe6, 0x19,
	0x88, 0x5b, 0x24, 0x4e, 0x4f, 0xac, 0x23, 0x54, 0x1b, 0xd1, 0x8d, 0xd0, 0x4f, 0x01, 0xea, 0xc5,
	0x13, 0xe8, 0x99, 0x67, 0x80, 0x7e, 0x0e, 0xd0, 0x13, 0x6b, 0x84, 0x3e, 0x83, 0x07, 0x98, 0x42,
	0x6a, 0xfa, 0x05, 0x71, 0xf8, 0xa3, 0x77, 0xda, 0x2f, 0x1b, 0x1b, 0xdd, 0x66, 0xa0, 0x7c, 0x0d,
	0xcb, 0x06, 0x47, 0xb7, 0x09, 0x88, 0x3f, 0x13, 0xbc, 0xaa, 0xeb, 0xbb, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc9, 0x77, 0x7a, 0xf2, 0xbe, 0x02, 0x00, 0x00,
}
