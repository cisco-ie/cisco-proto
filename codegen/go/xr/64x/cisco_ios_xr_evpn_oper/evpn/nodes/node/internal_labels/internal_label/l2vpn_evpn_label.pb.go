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
// source: l2vpn_evpn_label.proto

package cisco_ios_xr_evpn_oper_evpn_nodes_node_internal_labels_internal_label

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

type L2VpnEvpnLabel_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Evi                  uint32   `protobuf:"varint,2,opt,name=evi,proto3" json:"evi,omitempty"`
	Esi1                 string   `protobuf:"bytes,3,opt,name=esi1,proto3" json:"esi1,omitempty"`
	Esi2                 string   `protobuf:"bytes,4,opt,name=esi2,proto3" json:"esi2,omitempty"`
	Esi3                 string   `protobuf:"bytes,5,opt,name=esi3,proto3" json:"esi3,omitempty"`
	Esi4                 string   `protobuf:"bytes,6,opt,name=esi4,proto3" json:"esi4,omitempty"`
	Esi5                 string   `protobuf:"bytes,7,opt,name=esi5,proto3" json:"esi5,omitempty"`
	EthernetTag          uint32   `protobuf:"varint,8,opt,name=ethernet_tag,json=ethernetTag,proto3" json:"ethernet_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnEvpnLabel_KEYS) Reset()         { *m = L2VpnEvpnLabel_KEYS{} }
func (m *L2VpnEvpnLabel_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnLabel_KEYS) ProtoMessage()    {}
func (*L2VpnEvpnLabel_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_da030edbf6bf8249, []int{0}
}

func (m *L2VpnEvpnLabel_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnLabel_KEYS.Unmarshal(m, b)
}
func (m *L2VpnEvpnLabel_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnLabel_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnLabel_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnLabel_KEYS.Merge(m, src)
}
func (m *L2VpnEvpnLabel_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnLabel_KEYS.Size(m)
}
func (m *L2VpnEvpnLabel_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnLabel_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnLabel_KEYS proto.InternalMessageInfo

func (m *L2VpnEvpnLabel_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEvi() uint32 {
	if m != nil {
		return m.Evi
	}
	return 0
}

func (m *L2VpnEvpnLabel_KEYS) GetEsi1() string {
	if m != nil {
		return m.Esi1
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEsi2() string {
	if m != nil {
		return m.Esi2
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEsi3() string {
	if m != nil {
		return m.Esi3
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEsi4() string {
	if m != nil {
		return m.Esi4
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEsi5() string {
	if m != nil {
		return m.Esi5
	}
	return ""
}

func (m *L2VpnEvpnLabel_KEYS) GetEthernetTag() uint32 {
	if m != nil {
		return m.EthernetTag
	}
	return 0
}

type L2VpnLabelPathBuffer struct {
	NextHop              string   `protobuf:"bytes,1,opt,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
	OutputLabel          uint32   `protobuf:"varint,2,opt,name=output_label,json=outputLabel,proto3" json:"output_label,omitempty"`
	SrteTunnel           string   `protobuf:"bytes,3,opt,name=srte_tunnel,json=srteTunnel,proto3" json:"srte_tunnel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnLabelPathBuffer) Reset()         { *m = L2VpnLabelPathBuffer{} }
func (m *L2VpnLabelPathBuffer) String() string { return proto.CompactTextString(m) }
func (*L2VpnLabelPathBuffer) ProtoMessage()    {}
func (*L2VpnLabelPathBuffer) Descriptor() ([]byte, []int) {
	return fileDescriptor_da030edbf6bf8249, []int{1}
}

func (m *L2VpnLabelPathBuffer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Unmarshal(m, b)
}
func (m *L2VpnLabelPathBuffer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Marshal(b, m, deterministic)
}
func (m *L2VpnLabelPathBuffer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnLabelPathBuffer.Merge(m, src)
}
func (m *L2VpnLabelPathBuffer) XXX_Size() int {
	return xxx_messageInfo_L2VpnLabelPathBuffer.Size(m)
}
func (m *L2VpnLabelPathBuffer) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnLabelPathBuffer.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnLabelPathBuffer proto.InternalMessageInfo

func (m *L2VpnLabelPathBuffer) GetNextHop() string {
	if m != nil {
		return m.NextHop
	}
	return ""
}

func (m *L2VpnLabelPathBuffer) GetOutputLabel() uint32 {
	if m != nil {
		return m.OutputLabel
	}
	return 0
}

func (m *L2VpnLabelPathBuffer) GetSrteTunnel() string {
	if m != nil {
		return m.SrteTunnel
	}
	return ""
}

type L2VpnEvpnLabel struct {
	EviXr                      uint32                  `protobuf:"varint,50,opt,name=evi_xr,json=eviXr,proto3" json:"evi_xr,omitempty"`
	Esi                        []uint32                `protobuf:"varint,51,rep,packed,name=esi,proto3" json:"esi,omitempty"`
	Tag                        uint32                  `protobuf:"varint,52,opt,name=tag,proto3" json:"tag,omitempty"`
	InternalLabel              uint32                  `protobuf:"varint,53,opt,name=internal_label,json=internalLabel,proto3" json:"internal_label,omitempty"`
	Encap                      uint32                  `protobuf:"varint,54,opt,name=encap,proto3" json:"encap,omitempty"`
	MacNumPaths                uint32                  `protobuf:"varint,55,opt,name=mac_num_paths,json=macNumPaths,proto3" json:"mac_num_paths,omitempty"`
	MacPathBuffer              []*L2VpnLabelPathBuffer `protobuf:"bytes,56,rep,name=mac_path_buffer,json=macPathBuffer,proto3" json:"mac_path_buffer,omitempty"`
	EadNumPaths                uint32                  `protobuf:"varint,57,opt,name=ead_num_paths,json=eadNumPaths,proto3" json:"ead_num_paths,omitempty"`
	EadPathBuffer              []*L2VpnLabelPathBuffer `protobuf:"bytes,58,rep,name=ead_path_buffer,json=eadPathBuffer,proto3" json:"ead_path_buffer,omitempty"`
	EviNumPaths                uint32                  `protobuf:"varint,59,opt,name=evi_num_paths,json=eviNumPaths,proto3" json:"evi_num_paths,omitempty"`
	EviPathBuffer              []*L2VpnLabelPathBuffer `protobuf:"bytes,60,rep,name=evi_path_buffer,json=eviPathBuffer,proto3" json:"evi_path_buffer,omitempty"`
	SumNumPaths                uint32                  `protobuf:"varint,61,opt,name=sum_num_paths,json=sumNumPaths,proto3" json:"sum_num_paths,omitempty"`
	SumNumActivePaths          uint32                  `protobuf:"varint,62,opt,name=sum_num_active_paths,json=sumNumActivePaths,proto3" json:"sum_num_active_paths,omitempty"`
	SummaryPathBuffer          []*L2VpnLabelPathBuffer `protobuf:"bytes,63,rep,name=summary_path_buffer,json=summaryPathBuffer,proto3" json:"summary_path_buffer,omitempty"`
	Resolved                   bool                    `protobuf:"varint,64,opt,name=resolved,proto3" json:"resolved,omitempty"`
	EcmpDisable                bool                    `protobuf:"varint,65,opt,name=ecmp_disable,json=ecmpDisable,proto3" json:"ecmp_disable,omitempty"`
	RedundancySingleActive     bool                    `protobuf:"varint,66,opt,name=redundancy_single_active,json=redundancySingleActive,proto3" json:"redundancy_single_active,omitempty"`
	RedundancySingleFlowActive bool                    `protobuf:"varint,67,opt,name=redundancy_single_flow_active,json=redundancySingleFlowActive,proto3" json:"redundancy_single_flow_active,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                `json:"-"`
	XXX_unrecognized           []byte                  `json:"-"`
	XXX_sizecache              int32                   `json:"-"`
}

func (m *L2VpnEvpnLabel) Reset()         { *m = L2VpnEvpnLabel{} }
func (m *L2VpnEvpnLabel) String() string { return proto.CompactTextString(m) }
func (*L2VpnEvpnLabel) ProtoMessage()    {}
func (*L2VpnEvpnLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_da030edbf6bf8249, []int{2}
}

func (m *L2VpnEvpnLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnEvpnLabel.Unmarshal(m, b)
}
func (m *L2VpnEvpnLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnEvpnLabel.Marshal(b, m, deterministic)
}
func (m *L2VpnEvpnLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnEvpnLabel.Merge(m, src)
}
func (m *L2VpnEvpnLabel) XXX_Size() int {
	return xxx_messageInfo_L2VpnEvpnLabel.Size(m)
}
func (m *L2VpnEvpnLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnEvpnLabel.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnEvpnLabel proto.InternalMessageInfo

func (m *L2VpnEvpnLabel) GetEviXr() uint32 {
	if m != nil {
		return m.EviXr
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetEsi() []uint32 {
	if m != nil {
		return m.Esi
	}
	return nil
}

func (m *L2VpnEvpnLabel) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetInternalLabel() uint32 {
	if m != nil {
		return m.InternalLabel
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetEncap() uint32 {
	if m != nil {
		return m.Encap
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetMacNumPaths() uint32 {
	if m != nil {
		return m.MacNumPaths
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetMacPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.MacPathBuffer
	}
	return nil
}

func (m *L2VpnEvpnLabel) GetEadNumPaths() uint32 {
	if m != nil {
		return m.EadNumPaths
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetEadPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.EadPathBuffer
	}
	return nil
}

func (m *L2VpnEvpnLabel) GetEviNumPaths() uint32 {
	if m != nil {
		return m.EviNumPaths
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetEviPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.EviPathBuffer
	}
	return nil
}

func (m *L2VpnEvpnLabel) GetSumNumPaths() uint32 {
	if m != nil {
		return m.SumNumPaths
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetSumNumActivePaths() uint32 {
	if m != nil {
		return m.SumNumActivePaths
	}
	return 0
}

func (m *L2VpnEvpnLabel) GetSummaryPathBuffer() []*L2VpnLabelPathBuffer {
	if m != nil {
		return m.SummaryPathBuffer
	}
	return nil
}

func (m *L2VpnEvpnLabel) GetResolved() bool {
	if m != nil {
		return m.Resolved
	}
	return false
}

func (m *L2VpnEvpnLabel) GetEcmpDisable() bool {
	if m != nil {
		return m.EcmpDisable
	}
	return false
}

func (m *L2VpnEvpnLabel) GetRedundancySingleActive() bool {
	if m != nil {
		return m.RedundancySingleActive
	}
	return false
}

func (m *L2VpnEvpnLabel) GetRedundancySingleFlowActive() bool {
	if m != nil {
		return m.RedundancySingleFlowActive
	}
	return false
}

func init() {
	proto.RegisterType((*L2VpnEvpnLabel_KEYS)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.internal_labels.internal_label.l2vpn_evpn_label_KEYS")
	proto.RegisterType((*L2VpnLabelPathBuffer)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.internal_labels.internal_label.l2vpn_label_path_buffer")
	proto.RegisterType((*L2VpnEvpnLabel)(nil), "cisco_ios_xr_evpn_oper.evpn.nodes.node.internal_labels.internal_label.l2vpn_evpn_label")
}

func init() { proto.RegisterFile("l2vpn_evpn_label.proto", fileDescriptor_da030edbf6bf8249) }

var fileDescriptor_da030edbf6bf8249 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xc7, 0xb5, 0xa4, 0x4d, 0x83, 0x43, 0xa0, 0x35, 0xfd, 0x30, 0x95, 0x10, 0x21, 0x12, 0x52,
	0x4e, 0x41, 0x24, 0x2d, 0x94, 0x6f, 0x5a, 0x28, 0x02, 0x81, 0x10, 0x4a, 0x7b, 0x80, 0x0b, 0x96,
	0xb3, 0x3b, 0x4d, 0x2d, 0xed, 0xda, 0x2b, 0xdb, 0xbb, 0x4d, 0x5f, 0x80, 0x23, 0x2f, 0xc3, 0x53,
	0xf0, 0x56, 0xc8, 0x1f, 0x09, 0xd9, 0x20, 0x6e, 0x28, 0x97, 0x95, 0xe7, 0x37, 0x33, 0xf6, 0xff,
	0x3f, 0x5e, 0x19, 0x6d, 0xa7, 0xfd, 0x32, 0x17, 0x14, 0xec, 0x27, 0x65, 0x23, 0x48, 0x7b, 0xb9,
	0x92, 0x46, 0xe2, 0xe3, 0x98, 0xeb, 0x58, 0x52, 0x2e, 0x35, 0x9d, 0x28, 0x9f, 0x96, 0x39, 0xa8,
	0x9e, 0x5d, 0xf5, 0x84, 0x4c, 0x40, 0xbb, 0x6f, 0x8f, 0x0b, 0x03, 0x4a, 0xb0, 0xd4, 0x37, 0xeb,
	0x85, 0xb8, 0xf3, 0x2b, 0x42, 0x5b, 0x8b, 0x27, 0xd0, 0x0f, 0xc7, 0x5f, 0x4f, 0xf0, 0x0e, 0x5a,
	0xb3, 0x1b, 0x50, 0x9e, 0x90, 0xa8, 0x1d, 0x75, 0xaf, 0x0e, 0xeb, 0x36, 0x7c, 0x9f, 0xe0, 0x75,
	0x54, 0x83, 0x92, 0x93, 0x2b, 0xed, 0xa8, 0xdb, 0x1a, 0xda, 0x25, 0xc6, 0x68, 0x05, 0x34, 0x7f,
	0x40, 0x6a, 0xae, 0xce, 0xad, 0x03, 0xeb, 0x93, 0x95, 0x19, 0xeb, 0x07, 0x36, 0x20, 0xab, 0x33,
	0x36, 0x08, 0x6c, 0x8f, 0xd4, 0x67, 0x6c, 0x2f, 0xb0, 0x7d, 0xb2, 0x36, 0x63, 0xfb, 0xf8, 0x2e,
	0xba, 0x06, 0xe6, 0x1c, 0x94, 0x00, 0x43, 0x0d, 0x1b, 0x93, 0x86, 0x3b, 0xbe, 0x39, 0x65, 0xa7,
	0x6c, 0xdc, 0x99, 0xa0, 0x1d, 0x6f, 0xc5, 0xbb, 0xc8, 0x99, 0x39, 0xa7, 0xa3, 0xe2, 0xec, 0x0c,
	0x14, 0xbe, 0x85, 0x1a, 0x02, 0x26, 0x86, 0x9e, 0xcb, 0x3c, 0xb8, 0x59, 0xb3, 0xf1, 0x3b, 0x99,
	0xdb, 0x8d, 0x65, 0x61, 0xf2, 0xc2, 0xf8, 0xb6, 0xe0, 0xab, 0xe9, 0xd9, 0x47, 0x8b, 0xf0, 0x1d,
	0xd4, 0xd4, 0xca, 0x00, 0x35, 0x85, 0x10, 0x90, 0x06, 0x9b, 0xc8, 0xa2, 0x53, 0x47, 0x3a, 0x3f,
	0x1b, 0x68, 0x7d, 0x71, 0x8a, 0x78, 0x0b, 0xd5, 0xa1, 0xe4, 0x74, 0xa2, 0x48, 0xdf, 0x6d, 0xb9,
	0x0a, 0x25, 0xff, 0xa2, 0xdc, 0xf8, 0x34, 0x27, 0x83, 0x76, 0xcd, 0x8d, 0x4f, 0x73, 0x4b, 0xac,
	0xa3, 0x3d, 0x3f, 0x50, 0xc3, 0xc6, 0xf8, 0x1e, 0xba, 0x5e, 0xbd, 0x27, 0xb2, 0xef, 0x92, 0xad,
	0x29, 0xf5, 0xba, 0x36, 0xd1, 0x2a, 0x88, 0x98, 0xe5, 0xe4, 0x61, 0x38, 0xc0, 0x06, 0xb8, 0x83,
	0x5a, 0x19, 0x8b, 0xa9, 0x28, 0x32, 0x37, 0x02, 0x4d, 0x1e, 0x79, 0x47, 0x19, 0x8b, 0x3f, 0x15,
	0xd9, 0x67, 0x8b, 0xf0, 0xf7, 0x08, 0xdd, 0xb0, 0x45, 0x73, 0x33, 0x22, 0x07, 0xed, 0x5a, 0xb7,
	0xd9, 0xff, 0xd6, 0xfb, 0x2f, 0x3f, 0x56, 0xef, 0x1f, 0x37, 0x31, 0xb4, 0xda, 0xac, 0x86, 0x23,
	0x7f, 0x31, 0x1d, 0xd4, 0x02, 0x96, 0xcc, 0x89, 0x7d, 0x1c, 0xee, 0x95, 0x25, 0x15, 0xb1, 0xb6,
	0x68, 0x5e, 0xec, 0x93, 0xe5, 0x88, 0x05, 0x96, 0x2c, 0x88, 0x2d, 0xf9, 0x9c, 0xd8, 0xa7, 0x41,
	0x6c, 0xc9, 0xab, 0x62, 0x4b, 0x5e, 0x11, 0xfb, 0x6c, 0x49, 0x62, 0x4b, 0x5e, 0x15, 0xab, 0x8b,
	0x6c, 0x4e, 0xec, 0x73, 0x2f, 0x56, 0x17, 0xd9, 0x4c, 0xec, 0x7d, 0xb4, 0x39, 0xad, 0x61, 0xb1,
	0xe1, 0x25, 0x84, 0xd2, 0x17, 0xae, 0x74, 0xc3, 0x97, 0x1e, 0xba, 0x8c, 0x6f, 0xf8, 0x11, 0xa1,
	0x9b, 0xba, 0xc8, 0x32, 0xa6, 0x2e, 0x2b, 0x0e, 0x5f, 0x2e, 0xc5, 0xe1, 0x46, 0x38, 0x7a, 0xce,
	0xe5, 0x2e, 0x6a, 0x28, 0xd0, 0x32, 0x2d, 0x21, 0x21, 0xaf, 0xda, 0x51, 0xb7, 0x31, 0x9c, 0xc5,
	0xee, 0xc9, 0x88, 0xb3, 0x9c, 0x26, 0x5c, 0xb3, 0x51, 0x0a, 0xe4, 0xd0, 0xe5, 0x9b, 0x96, 0xbd,
	0xf1, 0x08, 0x1f, 0x20, 0xa2, 0x20, 0x29, 0x44, 0xc2, 0x44, 0x7c, 0x49, 0x35, 0x17, 0xe3, 0x14,
	0xc2, 0x28, 0xc8, 0x91, 0x2b, 0xdf, 0xfe, 0x93, 0x3f, 0x71, 0x69, 0x3f, 0x0e, 0x7c, 0x88, 0x6e,
	0xff, 0xdd, 0x79, 0x96, 0xca, 0x8b, 0x69, 0xfb, 0x6b, 0xd7, 0xbe, 0xbb, 0xd8, 0xfe, 0x36, 0x95,
	0x17, 0x7e, 0x8b, 0x51, 0xdd, 0xbd, 0xe4, 0x83, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xe2,
	0xef, 0x7e, 0xe3, 0x05, 0x00, 0x00,
}
