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
// source: ipv6_nd_bl_node_entry.proto

package cisco_ios_xr_ipv6_nd_oper_ipv6_node_discovery_nodes_node_bundle_nodes_bundle_node

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

type Ipv6NdBlNodeEntry_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	NodeName_1           string   `protobuf:"bytes,2,opt,name=node_name_1,json=nodeName1,proto3" json:"node_name_1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6NdBlNodeEntry_KEYS) Reset()         { *m = Ipv6NdBlNodeEntry_KEYS{} }
func (m *Ipv6NdBlNodeEntry_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdBlNodeEntry_KEYS) ProtoMessage()    {}
func (*Ipv6NdBlNodeEntry_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca7327ff191973f, []int{0}
}

func (m *Ipv6NdBlNodeEntry_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS.Unmarshal(m, b)
}
func (m *Ipv6NdBlNodeEntry_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6NdBlNodeEntry_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS.Merge(m, src)
}
func (m *Ipv6NdBlNodeEntry_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS.Size(m)
}
func (m *Ipv6NdBlNodeEntry_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdBlNodeEntry_KEYS proto.InternalMessageInfo

func (m *Ipv6NdBlNodeEntry_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6NdBlNodeEntry_KEYS) GetNodeName_1() string {
	if m != nil {
		return m.NodeName_1
	}
	return ""
}

type BagTimespec struct {
	Seconds              uint32   `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BagTimespec) Reset()         { *m = BagTimespec{} }
func (m *BagTimespec) String() string { return proto.CompactTextString(m) }
func (*BagTimespec) ProtoMessage()    {}
func (*BagTimespec) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca7327ff191973f, []int{1}
}

func (m *BagTimespec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BagTimespec.Unmarshal(m, b)
}
func (m *BagTimespec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BagTimespec.Marshal(b, m, deterministic)
}
func (m *BagTimespec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BagTimespec.Merge(m, src)
}
func (m *BagTimespec) XXX_Size() int {
	return xxx_messageInfo_BagTimespec.Size(m)
}
func (m *BagTimespec) XXX_DiscardUnknown() {
	xxx_messageInfo_BagTimespec.DiscardUnknown(m)
}

var xxx_messageInfo_BagTimespec proto.InternalMessageInfo

func (m *BagTimespec) GetSeconds() uint32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

type Ipv6NdBlNodeEntry struct {
	GroupId                uint32       `protobuf:"varint,50,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ProcessName            string       `protobuf:"bytes,51,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	Age                    *BagTimespec `protobuf:"bytes,52,opt,name=age,proto3" json:"age,omitempty"`
	SentSequenceNumber     uint32       `protobuf:"varint,53,opt,name=sent_sequence_number,json=sentSequenceNumber,proto3" json:"sent_sequence_number,omitempty"`
	ReceivedSequenceNumber uint32       `protobuf:"varint,54,opt,name=received_sequence_number,json=receivedSequenceNumber,proto3" json:"received_sequence_number,omitempty"`
	State                  string       `protobuf:"bytes,55,opt,name=state,proto3" json:"state,omitempty"`
	StateChanges           uint32       `protobuf:"varint,56,opt,name=state_changes,json=stateChanges,proto3" json:"state_changes,omitempty"`
	SentPackets            uint32       `protobuf:"varint,57,opt,name=sent_packets,json=sentPackets,proto3" json:"sent_packets,omitempty"`
	ReceivedPackets        uint32       `protobuf:"varint,58,opt,name=received_packets,json=receivedPackets,proto3" json:"received_packets,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}     `json:"-"`
	XXX_unrecognized       []byte       `json:"-"`
	XXX_sizecache          int32        `json:"-"`
}

func (m *Ipv6NdBlNodeEntry) Reset()         { *m = Ipv6NdBlNodeEntry{} }
func (m *Ipv6NdBlNodeEntry) String() string { return proto.CompactTextString(m) }
func (*Ipv6NdBlNodeEntry) ProtoMessage()    {}
func (*Ipv6NdBlNodeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bca7327ff191973f, []int{2}
}

func (m *Ipv6NdBlNodeEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6NdBlNodeEntry.Unmarshal(m, b)
}
func (m *Ipv6NdBlNodeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6NdBlNodeEntry.Marshal(b, m, deterministic)
}
func (m *Ipv6NdBlNodeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6NdBlNodeEntry.Merge(m, src)
}
func (m *Ipv6NdBlNodeEntry) XXX_Size() int {
	return xxx_messageInfo_Ipv6NdBlNodeEntry.Size(m)
}
func (m *Ipv6NdBlNodeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6NdBlNodeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6NdBlNodeEntry proto.InternalMessageInfo

func (m *Ipv6NdBlNodeEntry) GetGroupId() uint32 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *Ipv6NdBlNodeEntry) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ipv6NdBlNodeEntry) GetAge() *BagTimespec {
	if m != nil {
		return m.Age
	}
	return nil
}

func (m *Ipv6NdBlNodeEntry) GetSentSequenceNumber() uint32 {
	if m != nil {
		return m.SentSequenceNumber
	}
	return 0
}

func (m *Ipv6NdBlNodeEntry) GetReceivedSequenceNumber() uint32 {
	if m != nil {
		return m.ReceivedSequenceNumber
	}
	return 0
}

func (m *Ipv6NdBlNodeEntry) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Ipv6NdBlNodeEntry) GetStateChanges() uint32 {
	if m != nil {
		return m.StateChanges
	}
	return 0
}

func (m *Ipv6NdBlNodeEntry) GetSentPackets() uint32 {
	if m != nil {
		return m.SentPackets
	}
	return 0
}

func (m *Ipv6NdBlNodeEntry) GetReceivedPackets() uint32 {
	if m != nil {
		return m.ReceivedPackets
	}
	return 0
}

func init() {
	proto.RegisterType((*Ipv6NdBlNodeEntry_KEYS)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_nodes.bundle_node.ipv6_nd_bl_node_entry_KEYS")
	proto.RegisterType((*BagTimespec)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_nodes.bundle_node.bag_timespec")
	proto.RegisterType((*Ipv6NdBlNodeEntry)(nil), "cisco_ios_xr_ipv6_nd_oper.ipv6_node_discovery.nodes.node.bundle_nodes.bundle_node.ipv6_nd_bl_node_entry")
}

func init() { proto.RegisterFile("ipv6_nd_bl_node_entry.proto", fileDescriptor_bca7327ff191973f) }

var fileDescriptor_bca7327ff191973f = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4f, 0x6f, 0xda, 0x30,
	0x14, 0x17, 0x43, 0xdb, 0xc0, 0x01, 0x6d, 0xb2, 0xd8, 0xe4, 0x0d, 0x69, 0x62, 0xec, 0xc2, 0x2e,
	0xd1, 0x80, 0x8d, 0xb1, 0x5d, 0xab, 0x1e, 0xaa, 0x4a, 0xa8, 0x85, 0x13, 0xa7, 0xa7, 0xc4, 0x7e,
	0x4a, 0xa3, 0x12, 0x3b, 0xd8, 0x09, 0x2a, 0x1f, 0xae, 0xdf, 0xad, 0xf2, 0x0b, 0x41, 0xa8, 0xe5,
	0xd8, 0x4b, 0x94, 0xdf, 0x3f, 0xbf, 0x9f, 0x9f, 0xcc, 0xfa, 0x69, 0xbe, 0x9b, 0x81, 0x56, 0x10,
	0x6f, 0x40, 0x1b, 0x85, 0x80, 0xba, 0xb0, 0xfb, 0x30, 0xb7, 0xa6, 0x30, 0xfc, 0x56, 0xa6, 0x4e,
	0x1a, 0x48, 0x8d, 0x83, 0x07, 0x0b, 0xb5, 0xd3, 0xe4, 0x68, 0xc3, 0x0a, 0xf8, 0x80, 0xf2, 0x9e,
	0x1d, 0xda, 0x7d, 0xe8, 0xa1, 0xa3, 0x6f, 0x18, 0x97, 0x5a, 0x6d, 0x10, 0x2a, 0xe6, 0x04, 0x0c,
	0xd7, 0xec, 0xeb, 0xd9, 0x89, 0x70, 0x7d, 0xb9, 0x5e, 0xf1, 0x3e, 0x6b, 0x13, 0xa5, 0xa3, 0x0c,
	0x45, 0x63, 0xd0, 0x18, 0xb5, 0x97, 0x2d, 0x4f, 0x2c, 0xa2, 0x0c, 0xf9, 0x37, 0x16, 0x1c, 0x45,
	0x18, 0x8b, 0x37, 0x24, 0xb7, 0x6b, 0x79, 0x3c, 0x1c, 0xb1, 0x4e, 0x1c, 0x25, 0x50, 0xa4, 0x19,
	0xba, 0x1c, 0x25, 0x17, 0xec, 0xbd, 0x43, 0x69, 0xb4, 0x72, 0x74, 0x54, 0x77, 0x59, 0xc3, 0xe1,
	0x63, 0x93, 0x7d, 0x3a, 0xdb, 0x82, 0x7f, 0x61, 0xad, 0xc4, 0x9a, 0x32, 0x87, 0x54, 0x89, 0x49,
	0x15, 0x22, 0x7c, 0xa5, 0xf8, 0x77, 0xd6, 0xc9, 0xad, 0x91, 0xe8, 0x5c, 0x55, 0x6f, 0x4a, 0xf3,
	0x83, 0x03, 0x47, 0x0d, 0xb7, 0xac, 0x19, 0x25, 0x28, 0x7e, 0x0f, 0x1a, 0xa3, 0x60, 0x02, 0xe1,
	0xab, 0x6f, 0x2f, 0x3c, 0xbd, 0xdf, 0xd2, 0xcf, 0xe2, 0xbf, 0x58, 0xcf, 0xa1, 0x2e, 0xc0, 0xe1,
	0xb6, 0x44, 0x2d, 0x11, 0x74, 0x99, 0xc5, 0x68, 0xc5, 0x1f, 0x2a, 0xcf, 0xbd, 0xb6, 0x3a, 0x48,
	0x0b, 0x52, 0xf8, 0x9c, 0x09, 0x8b, 0x12, 0xd3, 0x1d, 0xaa, 0x17, 0xa9, 0x19, 0xa5, 0x3e, 0xd7,
	0xfa, 0xb3, 0x64, 0x8f, 0xbd, 0x75, 0x45, 0x54, 0xa0, 0xf8, 0x4b, 0x57, 0xaf, 0x00, 0xff, 0xc1,
	0xba, 0xf4, 0x03, 0xf2, 0x2e, 0xd2, 0x09, 0x3a, 0x31, 0xa7, 0x43, 0x3a, 0x44, 0x5e, 0x54, 0x9c,
	0x5f, 0x1e, 0xd5, 0xcc, 0x23, 0x79, 0x8f, 0x85, 0x13, 0xff, 0xc8, 0x13, 0x78, 0xee, 0xa6, 0xa2,
	0xf8, 0x4f, 0xf6, 0xf1, 0xd8, 0xab, 0xb6, 0xfd, 0x27, 0xdb, 0x87, 0x9a, 0x3f, 0x58, 0xe3, 0x77,
	0xf4, 0x3c, 0xa7, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x71, 0x15, 0x73, 0xbd, 0x02, 0x00,
	0x00,
}
