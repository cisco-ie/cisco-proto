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
// source: ipv4_dhcpd_disc_history.proto

package cisco_ios_xr_ipv4_dhcpd_oper_ipv4_dhcpd_nodes_node_proxy_disconnect_histories_disconnect_history

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

type Ipv4DhcpdDiscHistory_KEYS struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv4DhcpdDiscHistory_KEYS) Reset()         { *m = Ipv4DhcpdDiscHistory_KEYS{} }
func (m *Ipv4DhcpdDiscHistory_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdDiscHistory_KEYS) ProtoMessage()    {}
func (*Ipv4DhcpdDiscHistory_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d2eac51fb9ba1b4, []int{0}
}

func (m *Ipv4DhcpdDiscHistory_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS.Unmarshal(m, b)
}
func (m *Ipv4DhcpdDiscHistory_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdDiscHistory_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS.Merge(m, src)
}
func (m *Ipv4DhcpdDiscHistory_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS.Size(m)
}
func (m *Ipv4DhcpdDiscHistory_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdDiscHistory_KEYS proto.InternalMessageInfo

func (m *Ipv4DhcpdDiscHistory_KEYS) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

func (m *Ipv4DhcpdDiscHistory_KEYS) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type Ipv4DhcpdDiscHistory struct {
	SessionStartTimeEpoch uint64   `protobuf:"varint,50,opt,name=session_start_time_epoch,json=sessionStartTimeEpoch,proto3" json:"session_start_time_epoch,omitempty"`
	SessionEndTimeEpoch   uint64   `protobuf:"varint,51,opt,name=session_end_time_epoch,json=sessionEndTimeEpoch,proto3" json:"session_end_time_epoch,omitempty"`
	DiscReason            string   `protobuf:"bytes,52,opt,name=disc_reason,json=discReason,proto3" json:"disc_reason,omitempty"`
	SubLabel              uint32   `protobuf:"varint,53,opt,name=sub_label,json=subLabel,proto3" json:"sub_label,omitempty"`
	MacAddress            string   `protobuf:"bytes,54,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ipv4DhcpdDiscHistory) Reset()         { *m = Ipv4DhcpdDiscHistory{} }
func (m *Ipv4DhcpdDiscHistory) String() string { return proto.CompactTextString(m) }
func (*Ipv4DhcpdDiscHistory) ProtoMessage()    {}
func (*Ipv4DhcpdDiscHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d2eac51fb9ba1b4, []int{1}
}

func (m *Ipv4DhcpdDiscHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory.Unmarshal(m, b)
}
func (m *Ipv4DhcpdDiscHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory.Marshal(b, m, deterministic)
}
func (m *Ipv4DhcpdDiscHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv4DhcpdDiscHistory.Merge(m, src)
}
func (m *Ipv4DhcpdDiscHistory) XXX_Size() int {
	return xxx_messageInfo_Ipv4DhcpdDiscHistory.Size(m)
}
func (m *Ipv4DhcpdDiscHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv4DhcpdDiscHistory.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv4DhcpdDiscHistory proto.InternalMessageInfo

func (m *Ipv4DhcpdDiscHistory) GetSessionStartTimeEpoch() uint64 {
	if m != nil {
		return m.SessionStartTimeEpoch
	}
	return 0
}

func (m *Ipv4DhcpdDiscHistory) GetSessionEndTimeEpoch() uint64 {
	if m != nil {
		return m.SessionEndTimeEpoch
	}
	return 0
}

func (m *Ipv4DhcpdDiscHistory) GetDiscReason() string {
	if m != nil {
		return m.DiscReason
	}
	return ""
}

func (m *Ipv4DhcpdDiscHistory) GetSubLabel() uint32 {
	if m != nil {
		return m.SubLabel
	}
	return 0
}

func (m *Ipv4DhcpdDiscHistory) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv4DhcpdDiscHistory_KEYS)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.disconnect_histories.disconnect_history.ipv4_dhcpd_disc_history_KEYS")
	proto.RegisterType((*Ipv4DhcpdDiscHistory)(nil), "cisco_ios_xr_ipv4_dhcpd_oper.ipv4_dhcpd.nodes.node.proxy.disconnect_histories.disconnect_history.ipv4_dhcpd_disc_history")
}

func init() { proto.RegisterFile("ipv4_dhcpd_disc_history.proto", fileDescriptor_7d2eac51fb9ba1b4) }

var fileDescriptor_7d2eac51fb9ba1b4 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4b, 0x03, 0x31,
	0x18, 0xc5, 0x39, 0xd1, 0x62, 0x23, 0x2e, 0x51, 0x6b, 0x40, 0xc5, 0xd2, 0xa9, 0xd3, 0x0d, 0xb6,
	0xea, 0xec, 0xd0, 0xc9, 0x4e, 0x57, 0x17, 0xa7, 0xcf, 0x5c, 0xf2, 0xc1, 0x05, 0x7a, 0xc9, 0x91,
	0x2f, 0x95, 0xbb, 0x7f, 0xd8, 0xbf, 0x43, 0x92, 0x3b, 0x69, 0x41, 0xba, 0x04, 0xde, 0x7b, 0xf9,
	0xbd, 0x3c, 0x08, 0x7b, 0x30, 0xcd, 0xf7, 0x12, 0x74, 0xa5, 0x1a, 0x0d, 0xda, 0x90, 0x82, 0xca,
	0x50, 0x70, 0xbe, 0xcb, 0x1b, 0xef, 0x82, 0xe3, 0x5f, 0xca, 0x90, 0x72, 0x60, 0x1c, 0x41, 0xeb,
	0xe1, 0xe0, 0xae, 0x6b, 0xd0, 0xe7, 0x7b, 0x9d, 0x5b, 0xa7, 0x91, 0xd2, 0x19, 0xc9, 0xb6, 0xcb,
	0x63, 0x99, 0xb3, 0x16, 0x55, 0x18, 0x2a, 0x0d, 0xd2, 0x7f, 0xb3, 0x9b, 0xad, 0xd9, 0xfd, 0x91,
	0x09, 0xf0, 0xbe, 0xfa, 0xdc, 0xf0, 0x09, 0x1b, 0xc5, 0x56, 0xa3, 0x45, 0x36, 0xcd, 0xe6, 0xe3,
	0x62, 0x50, 0xfc, 0x9a, 0x9d, 0x19, 0xab, 0xb1, 0x15, 0x27, 0xc9, 0xee, 0xc5, 0xec, 0x27, 0x63,
	0xb7, 0x47, 0xea, 0xf8, 0x2b, 0x13, 0x84, 0x44, 0xc6, 0x59, 0xa0, 0x20, 0x7d, 0x80, 0x60, 0x6a,
	0x04, 0x6c, 0x9c, 0xaa, 0xc4, 0xd3, 0x34, 0x9b, 0x9f, 0x16, 0x37, 0x43, 0xbe, 0x89, 0xf1, 0x87,
	0xa9, 0x71, 0x15, 0x43, 0xbe, 0x60, 0x93, 0x3f, 0x10, 0xad, 0x3e, 0xc4, 0x16, 0x09, 0xbb, 0x1a,
	0xd2, 0x95, 0xd5, 0x7b, 0xe8, 0x91, 0x5d, 0xa4, 0xd7, 0x3d, 0x4a, 0x72, 0x56, 0x2c, 0xd3, 0x4a,
	0x16, 0xad, 0x22, 0x39, 0xfc, 0x8e, 0x8d, 0x69, 0x57, 0xc2, 0x56, 0x96, 0xb8, 0x15, 0xcf, 0xd3,
	0x6c, 0x7e, 0x59, 0x9c, 0xd3, 0xae, 0x5c, 0x47, 0x1d, 0xe9, 0x5a, 0x2a, 0x90, 0x5a, 0x7b, 0x24,
	0x12, 0x2f, 0x3d, 0x5d, 0x4b, 0xf5, 0xd6, 0x3b, 0xe5, 0x28, 0xfd, 0xcf, 0xe2, 0x37, 0x00, 0x00,
	0xff, 0xff, 0x41, 0x73, 0x83, 0x6c, 0xc0, 0x01, 0x00, 0x00,
}
