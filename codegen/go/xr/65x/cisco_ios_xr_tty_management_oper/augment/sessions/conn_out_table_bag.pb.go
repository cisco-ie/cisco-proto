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
// source: conn_out_table_bag.proto

package cisco_ios_xr_tty_management_oper_augment_sessions

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

type ConnOutTableBag_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnOutTableBag_KEYS) Reset()         { *m = ConnOutTableBag_KEYS{} }
func (m *ConnOutTableBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*ConnOutTableBag_KEYS) ProtoMessage()    {}
func (*ConnOutTableBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_de32bcf87d153fa3, []int{0}
}

func (m *ConnOutTableBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnOutTableBag_KEYS.Unmarshal(m, b)
}
func (m *ConnOutTableBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnOutTableBag_KEYS.Marshal(b, m, deterministic)
}
func (m *ConnOutTableBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnOutTableBag_KEYS.Merge(m, src)
}
func (m *ConnOutTableBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_ConnOutTableBag_KEYS.Size(m)
}
func (m *ConnOutTableBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnOutTableBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ConnOutTableBag_KEYS proto.InternalMessageInfo

type IpAddrTUnion struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4Address          string   `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address,proto3" json:"ipv4_address,omitempty"`
	Ipv6Address          string   `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address,proto3" json:"ipv6_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpAddrTUnion) Reset()         { *m = IpAddrTUnion{} }
func (m *IpAddrTUnion) String() string { return proto.CompactTextString(m) }
func (*IpAddrTUnion) ProtoMessage()    {}
func (*IpAddrTUnion) Descriptor() ([]byte, []int) {
	return fileDescriptor_de32bcf87d153fa3, []int{1}
}

func (m *IpAddrTUnion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpAddrTUnion.Unmarshal(m, b)
}
func (m *IpAddrTUnion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpAddrTUnion.Marshal(b, m, deterministic)
}
func (m *IpAddrTUnion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpAddrTUnion.Merge(m, src)
}
func (m *IpAddrTUnion) XXX_Size() int {
	return xxx_messageInfo_IpAddrTUnion.Size(m)
}
func (m *IpAddrTUnion) XXX_DiscardUnknown() {
	xxx_messageInfo_IpAddrTUnion.DiscardUnknown(m)
}

var xxx_messageInfo_IpAddrTUnion proto.InternalMessageInfo

func (m *IpAddrTUnion) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *IpAddrTUnion) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *IpAddrTUnion) GetIpv6Address() string {
	if m != nil {
		return m.Ipv6Address
	}
	return ""
}

type SuspendedSessionsTd struct {
	ConnectionId         uint32        `protobuf:"varint,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	HostName             string        `protobuf:"bytes,2,opt,name=host_name,json=hostName,proto3" json:"host_name,omitempty"`
	HostAddress          *IpAddrTUnion `protobuf:"bytes,3,opt,name=host_address,json=hostAddress,proto3" json:"host_address,omitempty"`
	TransportProtocol    string        `protobuf:"bytes,4,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	IsLastActiveSession  bool          `protobuf:"varint,5,opt,name=is_last_active_session,json=isLastActiveSession,proto3" json:"is_last_active_session,omitempty"`
	IdleTime             uint32        `protobuf:"varint,6,opt,name=idle_time,json=idleTime,proto3" json:"idle_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SuspendedSessionsTd) Reset()         { *m = SuspendedSessionsTd{} }
func (m *SuspendedSessionsTd) String() string { return proto.CompactTextString(m) }
func (*SuspendedSessionsTd) ProtoMessage()    {}
func (*SuspendedSessionsTd) Descriptor() ([]byte, []int) {
	return fileDescriptor_de32bcf87d153fa3, []int{2}
}

func (m *SuspendedSessionsTd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuspendedSessionsTd.Unmarshal(m, b)
}
func (m *SuspendedSessionsTd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuspendedSessionsTd.Marshal(b, m, deterministic)
}
func (m *SuspendedSessionsTd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuspendedSessionsTd.Merge(m, src)
}
func (m *SuspendedSessionsTd) XXX_Size() int {
	return xxx_messageInfo_SuspendedSessionsTd.Size(m)
}
func (m *SuspendedSessionsTd) XXX_DiscardUnknown() {
	xxx_messageInfo_SuspendedSessionsTd.DiscardUnknown(m)
}

var xxx_messageInfo_SuspendedSessionsTd proto.InternalMessageInfo

func (m *SuspendedSessionsTd) GetConnectionId() uint32 {
	if m != nil {
		return m.ConnectionId
	}
	return 0
}

func (m *SuspendedSessionsTd) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *SuspendedSessionsTd) GetHostAddress() *IpAddrTUnion {
	if m != nil {
		return m.HostAddress
	}
	return nil
}

func (m *SuspendedSessionsTd) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *SuspendedSessionsTd) GetIsLastActiveSession() bool {
	if m != nil {
		return m.IsLastActiveSession
	}
	return false
}

func (m *SuspendedSessionsTd) GetIdleTime() uint32 {
	if m != nil {
		return m.IdleTime
	}
	return 0
}

type ConnOutTableBag struct {
	OutgoingConnection   []*SuspendedSessionsTd `protobuf:"bytes,50,rep,name=outgoing_connection,json=outgoingConnection,proto3" json:"outgoing_connection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnOutTableBag) Reset()         { *m = ConnOutTableBag{} }
func (m *ConnOutTableBag) String() string { return proto.CompactTextString(m) }
func (*ConnOutTableBag) ProtoMessage()    {}
func (*ConnOutTableBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_de32bcf87d153fa3, []int{3}
}

func (m *ConnOutTableBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnOutTableBag.Unmarshal(m, b)
}
func (m *ConnOutTableBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnOutTableBag.Marshal(b, m, deterministic)
}
func (m *ConnOutTableBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnOutTableBag.Merge(m, src)
}
func (m *ConnOutTableBag) XXX_Size() int {
	return xxx_messageInfo_ConnOutTableBag.Size(m)
}
func (m *ConnOutTableBag) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnOutTableBag.DiscardUnknown(m)
}

var xxx_messageInfo_ConnOutTableBag proto.InternalMessageInfo

func (m *ConnOutTableBag) GetOutgoingConnection() []*SuspendedSessionsTd {
	if m != nil {
		return m.OutgoingConnection
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnOutTableBag_KEYS)(nil), "cisco_ios_xr_tty_management_oper.augment.sessions.conn_out_table_bag_KEYS")
	proto.RegisterType((*IpAddrTUnion)(nil), "cisco_ios_xr_tty_management_oper.augment.sessions.ip_addr_t_union")
	proto.RegisterType((*SuspendedSessionsTd)(nil), "cisco_ios_xr_tty_management_oper.augment.sessions.suspended_sessions_td")
	proto.RegisterType((*ConnOutTableBag)(nil), "cisco_ios_xr_tty_management_oper.augment.sessions.conn_out_table_bag")
}

func init() { proto.RegisterFile("conn_out_table_bag.proto", fileDescriptor_de32bcf87d153fa3) }

var fileDescriptor_de32bcf87d153fa3 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0xef, 0xd2, 0x40,
	0x10, 0xc5, 0x53, 0x50, 0x84, 0x2d, 0xc4, 0xb8, 0x44, 0xa9, 0xf1, 0x82, 0xf5, 0xc2, 0xc5, 0x26,
	0x82, 0xe1, 0x8e, 0xc6, 0x44, 0xa3, 0x31, 0xa6, 0x78, 0xf1, 0x34, 0x59, 0xba, 0x43, 0xdd, 0xa4,
	0xdd, 0x6d, 0xba, 0x53, 0x22, 0x9f, 0xc2, 0xcf, 0xe4, 0x37, 0x33, 0xbb, 0xa5, 0xc5, 0x88, 0x17,
	0xff, 0xc7, 0x7d, 0xef, 0x65, 0xe7, 0x37, 0x6f, 0x58, 0x94, 0x19, 0xad, 0xc1, 0x34, 0x04, 0x24,
	0x0e, 0x05, 0xc2, 0x41, 0xe4, 0x49, 0x55, 0x1b, 0x32, 0xfc, 0x55, 0xa6, 0x6c, 0x66, 0x40, 0x19,
	0x0b, 0x3f, 0x6a, 0x20, 0x3a, 0x43, 0x29, 0xb4, 0xc8, 0xb1, 0x44, 0x4d, 0x60, 0x2a, 0xac, 0x13,
	0xd1, 0xe4, 0xee, 0x91, 0x58, 0xb4, 0x56, 0x19, 0x6d, 0xe3, 0xa7, 0x6c, 0x71, 0xfb, 0x1d, 0x7c,
	0x7c, 0xf7, 0x6d, 0x1f, 0x57, 0xec, 0xa1, 0xaa, 0x40, 0x48, 0x59, 0x03, 0x41, 0xa3, 0x95, 0xd1,
	0x7c, 0xc1, 0x1e, 0x88, 0x23, 0x68, 0x51, 0x62, 0x14, 0x2c, 0x83, 0xd5, 0x24, 0x1d, 0x89, 0xe3,
	0x67, 0x51, 0x22, 0x7f, 0xce, 0xa6, 0xaa, 0x3a, 0xbd, 0xf6, 0x69, 0xb4, 0x36, 0x1a, 0x78, 0x37,
	0x74, 0xda, 0xae, 0x95, 0x2e, 0x91, 0x6d, 0x1f, 0x19, 0xf6, 0x91, 0xed, 0x25, 0x12, 0xff, 0x1a,
	0xb0, 0xc7, 0xb6, 0xb1, 0x15, 0x6a, 0x89, 0x12, 0x3a, 0x46, 0x20, 0xc9, 0x5f, 0xb0, 0x99, 0xc3,
	0xc4, 0x8c, 0x94, 0xd1, 0xa0, 0xa4, 0x1f, 0x3f, 0x4b, 0xa7, 0x57, 0xf1, 0x83, 0xe4, 0xcf, 0xd8,
	0xe4, 0xbb, 0xb1, 0xd4, 0xf2, 0xb5, 0x04, 0x63, 0x27, 0x78, 0x42, 0x64, 0x53, 0x6f, 0xfe, 0x39,
	0x3e, 0x5c, 0xbf, 0x49, 0xfe, 0xbb, 0xb2, 0xe4, 0xaf, 0x52, 0xd2, 0xd0, 0xfd, 0xdb, 0x6d, 0xf9,
	0x92, 0x71, 0xaa, 0x85, 0xb6, 0x95, 0xa9, 0x09, 0xfc, 0x55, 0x32, 0x53, 0x44, 0xf7, 0x3c, 0xcc,
	0xa3, 0xde, 0xf9, 0x72, 0x31, 0xf8, 0x86, 0x3d, 0x51, 0x16, 0x0a, 0xe1, 0xc0, 0x32, 0x52, 0x27,
	0xec, 0xb6, 0x8e, 0xee, 0x2f, 0x83, 0xd5, 0x38, 0x9d, 0x2b, 0xfb, 0x49, 0x58, 0xda, 0x79, 0x6f,
	0xdf, 0x5a, 0x6e, 0x4f, 0x25, 0x0b, 0x04, 0x52, 0x25, 0x46, 0x23, 0x5f, 0xc4, 0xd8, 0x09, 0x5f,
	0x55, 0x89, 0xf1, 0xcf, 0x80, 0xf1, 0xdb, 0x8b, 0xf2, 0x33, 0x9b, 0x9b, 0x86, 0x72, 0xa3, 0x74,
	0x0e, 0xd7, 0xd2, 0xa2, 0xf5, 0x72, 0xb8, 0x0a, 0xd7, 0xef, 0xef, 0xd0, 0xc2, 0x3f, 0xef, 0x94,
	0xf2, 0x6e, 0xc8, 0xdb, 0x7e, 0xc6, 0x61, 0xe4, 0x6b, 0xd8, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x28, 0x41, 0x89, 0xe9, 0xb8, 0x02, 0x00, 0x00,
}
