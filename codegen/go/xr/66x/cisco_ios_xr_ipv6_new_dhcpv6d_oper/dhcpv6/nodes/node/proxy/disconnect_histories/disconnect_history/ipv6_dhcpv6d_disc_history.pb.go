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
// source: ipv6_dhcpv6d_disc_history.proto

package cisco_ios_xr_ipv6_new_dhcpv6d_oper_dhcpv6_nodes_node_proxy_disconnect_histories_disconnect_history

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

type Ipv6Dhcpv6DDiscHistory_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DDiscHistory_KEYS) Reset()         { *m = Ipv6Dhcpv6DDiscHistory_KEYS{} }
func (m *Ipv6Dhcpv6DDiscHistory_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DDiscHistory_KEYS) ProtoMessage()    {}
func (*Ipv6Dhcpv6DDiscHistory_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b47d63561a9d6, []int{0}
}

func (m *Ipv6Dhcpv6DDiscHistory_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DDiscHistory_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DDiscHistory_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS.Merge(m, src)
}
func (m *Ipv6Dhcpv6DDiscHistory_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS.Size(m)
}
func (m *Ipv6Dhcpv6DDiscHistory_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DDiscHistory_KEYS proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DDiscHistory_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Ipv6Dhcpv6DDiscHistory_KEYS) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type Ipv6Dhcpv6DDiscHistory struct {
	SessionStartTimeEpoch uint64   `protobuf:"varint,50,opt,name=session_start_time_epoch,json=sessionStartTimeEpoch,proto3" json:"session_start_time_epoch,omitempty"`
	SessionEndTimeEpoch   uint64   `protobuf:"varint,51,opt,name=session_end_time_epoch,json=sessionEndTimeEpoch,proto3" json:"session_end_time_epoch,omitempty"`
	DiscReason            string   `protobuf:"bytes,52,opt,name=disc_reason,json=discReason,proto3" json:"disc_reason,omitempty"`
	SubLabel              uint32   `protobuf:"varint,53,opt,name=sub_label,json=subLabel,proto3" json:"sub_label,omitempty"`
	Duid                  string   `protobuf:"bytes,54,opt,name=duid,proto3" json:"duid,omitempty"`
	IaType                string   `protobuf:"bytes,55,opt,name=ia_type,json=iaType,proto3" json:"ia_type,omitempty"`
	IaId                  uint32   `protobuf:"varint,56,opt,name=ia_id,json=iaId,proto3" json:"ia_id,omitempty"`
	MacAddress            string   `protobuf:"bytes,57,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *Ipv6Dhcpv6DDiscHistory) Reset()         { *m = Ipv6Dhcpv6DDiscHistory{} }
func (m *Ipv6Dhcpv6DDiscHistory) String() string { return proto.CompactTextString(m) }
func (*Ipv6Dhcpv6DDiscHistory) ProtoMessage()    {}
func (*Ipv6Dhcpv6DDiscHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_937b47d63561a9d6, []int{1}
}

func (m *Ipv6Dhcpv6DDiscHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory.Unmarshal(m, b)
}
func (m *Ipv6Dhcpv6DDiscHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory.Marshal(b, m, deterministic)
}
func (m *Ipv6Dhcpv6DDiscHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipv6Dhcpv6DDiscHistory.Merge(m, src)
}
func (m *Ipv6Dhcpv6DDiscHistory) XXX_Size() int {
	return xxx_messageInfo_Ipv6Dhcpv6DDiscHistory.Size(m)
}
func (m *Ipv6Dhcpv6DDiscHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipv6Dhcpv6DDiscHistory.DiscardUnknown(m)
}

var xxx_messageInfo_Ipv6Dhcpv6DDiscHistory proto.InternalMessageInfo

func (m *Ipv6Dhcpv6DDiscHistory) GetSessionStartTimeEpoch() uint64 {
	if m != nil {
		return m.SessionStartTimeEpoch
	}
	return 0
}

func (m *Ipv6Dhcpv6DDiscHistory) GetSessionEndTimeEpoch() uint64 {
	if m != nil {
		return m.SessionEndTimeEpoch
	}
	return 0
}

func (m *Ipv6Dhcpv6DDiscHistory) GetDiscReason() string {
	if m != nil {
		return m.DiscReason
	}
	return ""
}

func (m *Ipv6Dhcpv6DDiscHistory) GetSubLabel() uint32 {
	if m != nil {
		return m.SubLabel
	}
	return 0
}

func (m *Ipv6Dhcpv6DDiscHistory) GetDuid() string {
	if m != nil {
		return m.Duid
	}
	return ""
}

func (m *Ipv6Dhcpv6DDiscHistory) GetIaType() string {
	if m != nil {
		return m.IaType
	}
	return ""
}

func (m *Ipv6Dhcpv6DDiscHistory) GetIaId() uint32 {
	if m != nil {
		return m.IaId
	}
	return 0
}

func (m *Ipv6Dhcpv6DDiscHistory) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Ipv6Dhcpv6DDiscHistory_KEYS)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.disconnect_histories.disconnect_history.ipv6_dhcpv6d_disc_history_KEYS")
	proto.RegisterType((*Ipv6Dhcpv6DDiscHistory)(nil), "cisco_ios_xr_ipv6_new_dhcpv6d_oper.dhcpv6.nodes.node.proxy.disconnect_histories.disconnect_history.ipv6_dhcpv6d_disc_history")
}

func init() { proto.RegisterFile("ipv6_dhcpv6d_disc_history.proto", fileDescriptor_937b47d63561a9d6) }

var fileDescriptor_937b47d63561a9d6 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x6f, 0xe2, 0x30,
	0x10, 0xc5, 0x15, 0x04, 0x2c, 0x78, 0xb5, 0x17, 0xb3, 0x7f, 0xbc, 0x5a, 0x69, 0x41, 0x9c, 0x38,
	0xe5, 0x50, 0x5a, 0x68, 0x8f, 0x3d, 0x70, 0xa8, 0x5a, 0xf5, 0x10, 0xb8, 0xf4, 0x34, 0x72, 0xec,
	0x91, 0x18, 0x89, 0xd8, 0x51, 0x26, 0x50, 0xf2, 0x95, 0xfa, 0x29, 0x2b, 0x3b, 0xf4, 0x8f, 0x54,
	0x71, 0x89, 0x32, 0xef, 0xf9, 0xf7, 0xde, 0x58, 0x16, 0x63, 0x2a, 0x0f, 0x0b, 0xb0, 0x5b, 0x53,
	0x1e, 0x16, 0x16, 0x2c, 0xb1, 0x81, 0x2d, 0x71, 0xed, 0xab, 0x26, 0x2d, 0x2b, 0x5f, 0x7b, 0x99,
	0x1b, 0x62, 0xe3, 0x81, 0x3c, 0xc3, 0xb1, 0x82, 0x78, 0xda, 0xe1, 0xf3, 0x3b, 0xe1, 0x4b, 0xac,
	0xd2, 0x76, 0x48, 0x9d, 0xb7, 0xc8, 0xf1, 0x1b, 0xd8, 0x63, 0x93, 0x86, 0x38, 0xef, 0x1c, 0x9a,
	0xfa, 0x14, 0x4a, 0xc8, 0x5f, 0xc5, 0x66, 0xba, 0x16, 0xff, 0xcf, 0xae, 0x01, 0xf7, 0xab, 0xa7,
	0xb5, 0xfc, 0x27, 0x86, 0x21, 0x17, 0x9c, 0x2e, 0x50, 0x25, 0x93, 0x64, 0x36, 0xcc, 0x06, 0x41,
	0x78, 0xd4, 0x05, 0xca, 0x9f, 0xa2, 0x47, 0xce, 0xe2, 0x51, 0x75, 0xa2, 0xd1, 0x0e, 0xd3, 0x97,
	0x8e, 0xf8, 0x7b, 0x36, 0x55, 0x2e, 0x85, 0x62, 0x64, 0x26, 0xef, 0x80, 0x6b, 0x5d, 0xd5, 0x50,
	0x53, 0x81, 0x80, 0xa5, 0x37, 0x5b, 0x75, 0x31, 0x49, 0x66, 0xdd, 0xec, 0xd7, 0xc9, 0x5f, 0x07,
	0x7b, 0x43, 0x05, 0xae, 0x82, 0x29, 0xe7, 0xe2, 0xf7, 0x1b, 0x88, 0xce, 0x7e, 0xc6, 0xe6, 0x11,
	0x1b, 0x9d, 0xdc, 0x95, 0xb3, 0x1f, 0xd0, 0x58, 0x7c, 0x8f, 0xed, 0x15, 0x6a, 0xf6, 0x4e, 0x5d,
	0xc6, 0x3d, 0x45, 0x90, 0xb2, 0xa8, 0x84, 0xfb, 0xf1, 0x3e, 0x87, 0x9d, 0xce, 0x71, 0xa7, 0xae,
	0x26, 0xc9, 0xec, 0x47, 0x36, 0xe0, 0x7d, 0xfe, 0x10, 0x66, 0x29, 0x45, 0xd7, 0xee, 0xc9, 0xaa,
	0x45, 0xc4, 0xe2, 0xbf, 0xfc, 0x23, 0xbe, 0x91, 0x86, 0xba, 0x29, 0x51, 0x2d, 0xa3, 0xdc, 0x27,
	0xbd, 0x69, 0x4a, 0x94, 0x23, 0xd1, 0x23, 0x0d, 0x64, 0xd5, 0x75, 0x4c, 0xe9, 0x92, 0xbe, 0xb3,
	0xa1, 0xbf, 0xd0, 0x06, 0xb4, 0xb5, 0x15, 0x32, 0xab, 0x9b, 0xb6, 0xbf, 0xd0, 0xe6, 0xb6, 0x55,
	0xf2, 0x7e, 0x7c, 0xec, 0xf9, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xf4, 0x3c, 0x33, 0x0f,
	0x02, 0x00, 0x00,
}
