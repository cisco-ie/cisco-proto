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
// source: ipsla_resp_port_info_bag.proto

package cisco_ios_xr_man_ipsla_oper_ipsla_responder_ports_port

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

type IpslaRespPortInfoBag_KEYS struct {
	Port                 uint32   `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaRespPortInfoBag_KEYS) Reset()         { *m = IpslaRespPortInfoBag_KEYS{} }
func (m *IpslaRespPortInfoBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*IpslaRespPortInfoBag_KEYS) ProtoMessage()    {}
func (*IpslaRespPortInfoBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_6019f5a2240c08c5, []int{0}
}

func (m *IpslaRespPortInfoBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaRespPortInfoBag_KEYS.Unmarshal(m, b)
}
func (m *IpslaRespPortInfoBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaRespPortInfoBag_KEYS.Marshal(b, m, deterministic)
}
func (m *IpslaRespPortInfoBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaRespPortInfoBag_KEYS.Merge(m, src)
}
func (m *IpslaRespPortInfoBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_IpslaRespPortInfoBag_KEYS.Size(m)
}
func (m *IpslaRespPortInfoBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaRespPortInfoBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaRespPortInfoBag_KEYS proto.InternalMessageInfo

func (m *IpslaRespPortInfoBag_KEYS) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type IpslaRespSenderBag struct {
	IpAddress            uint32   `protobuf:"varint,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PortXr               uint32   `protobuf:"varint,2,opt,name=port_xr,json=portXr,proto3" json:"port_xr,omitempty"`
	LastRecvTime         uint64   `protobuf:"varint,3,opt,name=last_recv_time,json=lastRecvTime,proto3" json:"last_recv_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpslaRespSenderBag) Reset()         { *m = IpslaRespSenderBag{} }
func (m *IpslaRespSenderBag) String() string { return proto.CompactTextString(m) }
func (*IpslaRespSenderBag) ProtoMessage()    {}
func (*IpslaRespSenderBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6019f5a2240c08c5, []int{1}
}

func (m *IpslaRespSenderBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaRespSenderBag.Unmarshal(m, b)
}
func (m *IpslaRespSenderBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaRespSenderBag.Marshal(b, m, deterministic)
}
func (m *IpslaRespSenderBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaRespSenderBag.Merge(m, src)
}
func (m *IpslaRespSenderBag) XXX_Size() int {
	return xxx_messageInfo_IpslaRespSenderBag.Size(m)
}
func (m *IpslaRespSenderBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaRespSenderBag.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaRespSenderBag proto.InternalMessageInfo

func (m *IpslaRespSenderBag) GetIpAddress() uint32 {
	if m != nil {
		return m.IpAddress
	}
	return 0
}

func (m *IpslaRespSenderBag) GetPortXr() uint32 {
	if m != nil {
		return m.PortXr
	}
	return 0
}

func (m *IpslaRespSenderBag) GetLastRecvTime() uint64 {
	if m != nil {
		return m.LastRecvTime
	}
	return 0
}

type IpslaRespPortInfoBag struct {
	PortXr               uint32                `protobuf:"varint,50,opt,name=port_xr,json=portXr,proto3" json:"port_xr,omitempty"`
	LocalAddress         string                `protobuf:"bytes,51,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	NumProbes            uint32                `protobuf:"varint,52,opt,name=num_probes,json=numProbes,proto3" json:"num_probes,omitempty"`
	CtrlProbes           uint32                `protobuf:"varint,53,opt,name=ctrl_probes,json=ctrlProbes,proto3" json:"ctrl_probes,omitempty"`
	Permanent            bool                  `protobuf:"varint,54,opt,name=permanent,proto3" json:"permanent,omitempty"`
	DiscardOn            bool                  `protobuf:"varint,55,opt,name=discard_on,json=discardOn,proto3" json:"discard_on,omitempty"`
	PdTimeStampFailed    bool                  `protobuf:"varint,56,opt,name=pd_time_stamp_failed,json=pdTimeStampFailed,proto3" json:"pd_time_stamp_failed,omitempty"`
	Sender               []*IpslaRespSenderBag `protobuf:"bytes,57,rep,name=sender,proto3" json:"sender,omitempty"`
	IsIpsla              bool                  `protobuf:"varint,58,opt,name=is_ipsla,json=isIpsla,proto3" json:"is_ipsla,omitempty"`
	DropCounter          uint32                `protobuf:"varint,59,opt,name=drop_counter,json=dropCounter,proto3" json:"drop_counter,omitempty"`
	Socket               int32                 `protobuf:"zigzag32,60,opt,name=socket,proto3" json:"socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IpslaRespPortInfoBag) Reset()         { *m = IpslaRespPortInfoBag{} }
func (m *IpslaRespPortInfoBag) String() string { return proto.CompactTextString(m) }
func (*IpslaRespPortInfoBag) ProtoMessage()    {}
func (*IpslaRespPortInfoBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_6019f5a2240c08c5, []int{2}
}

func (m *IpslaRespPortInfoBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpslaRespPortInfoBag.Unmarshal(m, b)
}
func (m *IpslaRespPortInfoBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpslaRespPortInfoBag.Marshal(b, m, deterministic)
}
func (m *IpslaRespPortInfoBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpslaRespPortInfoBag.Merge(m, src)
}
func (m *IpslaRespPortInfoBag) XXX_Size() int {
	return xxx_messageInfo_IpslaRespPortInfoBag.Size(m)
}
func (m *IpslaRespPortInfoBag) XXX_DiscardUnknown() {
	xxx_messageInfo_IpslaRespPortInfoBag.DiscardUnknown(m)
}

var xxx_messageInfo_IpslaRespPortInfoBag proto.InternalMessageInfo

func (m *IpslaRespPortInfoBag) GetPortXr() uint32 {
	if m != nil {
		return m.PortXr
	}
	return 0
}

func (m *IpslaRespPortInfoBag) GetLocalAddress() string {
	if m != nil {
		return m.LocalAddress
	}
	return ""
}

func (m *IpslaRespPortInfoBag) GetNumProbes() uint32 {
	if m != nil {
		return m.NumProbes
	}
	return 0
}

func (m *IpslaRespPortInfoBag) GetCtrlProbes() uint32 {
	if m != nil {
		return m.CtrlProbes
	}
	return 0
}

func (m *IpslaRespPortInfoBag) GetPermanent() bool {
	if m != nil {
		return m.Permanent
	}
	return false
}

func (m *IpslaRespPortInfoBag) GetDiscardOn() bool {
	if m != nil {
		return m.DiscardOn
	}
	return false
}

func (m *IpslaRespPortInfoBag) GetPdTimeStampFailed() bool {
	if m != nil {
		return m.PdTimeStampFailed
	}
	return false
}

func (m *IpslaRespPortInfoBag) GetSender() []*IpslaRespSenderBag {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *IpslaRespPortInfoBag) GetIsIpsla() bool {
	if m != nil {
		return m.IsIpsla
	}
	return false
}

func (m *IpslaRespPortInfoBag) GetDropCounter() uint32 {
	if m != nil {
		return m.DropCounter
	}
	return 0
}

func (m *IpslaRespPortInfoBag) GetSocket() int32 {
	if m != nil {
		return m.Socket
	}
	return 0
}

func init() {
	proto.RegisterType((*IpslaRespPortInfoBag_KEYS)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.responder.ports.port.ipsla_resp_port_info_bag_KEYS")
	proto.RegisterType((*IpslaRespSenderBag)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.responder.ports.port.ipsla_resp_sender_bag")
	proto.RegisterType((*IpslaRespPortInfoBag)(nil), "cisco_ios_xr_man_ipsla_oper.ipsla.responder.ports.port.ipsla_resp_port_info_bag")
}

func init() { proto.RegisterFile("ipsla_resp_port_info_bag.proto", fileDescriptor_6019f5a2240c08c5) }

var fileDescriptor_6019f5a2240c08c5 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x5f, 0x6b, 0x14, 0x31,
	0x14, 0xc5, 0x19, 0xb7, 0x4c, 0xbb, 0x77, 0xb7, 0x42, 0x83, 0x7f, 0x22, 0xb8, 0x3a, 0xae, 0x3e,
	0xcc, 0xd3, 0x08, 0x5d, 0xad, 0x7f, 0x5f, 0x44, 0x14, 0x44, 0x44, 0x99, 0xfa, 0xa0, 0x4f, 0x97,
	0x6c, 0x92, 0x4a, 0x70, 0x26, 0x09, 0x37, 0x99, 0xd2, 0xef, 0xe7, 0x17, 0x93, 0x64, 0xa6, 0x5a,
	0xc1, 0xbe, 0xf4, 0x65, 0xc9, 0xfd, 0x9d, 0x9c, 0xb3, 0x27, 0xc9, 0xc0, 0x3d, 0xe3, 0x43, 0x27,
	0x90, 0x74, 0xf0, 0xe8, 0x1d, 0x45, 0x34, 0xf6, 0xc4, 0xe1, 0x56, 0xfc, 0x68, 0x3c, 0xb9, 0xe8,
	0xd8, 0x91, 0x34, 0x41, 0x3a, 0x34, 0x2e, 0xe0, 0x19, 0x61, 0x2f, 0x2c, 0x8e, 0x06, 0xe7, 0x35,
	0x35, 0x79, 0xd9, 0x24, 0xaf, 0xb3, 0x4a, 0x53, 0x93, 0x02, 0x42, 0xfe, 0x5d, 0x6f, 0x60, 0x75,
	0x59, 0x32, 0x7e, 0x7c, 0xf7, 0xfd, 0x98, 0x31, 0xd8, 0x49, 0x94, 0x17, 0x55, 0x51, 0xef, 0xb7,
	0x79, 0xbd, 0x1e, 0xe0, 0xe6, 0x05, 0x53, 0xd0, 0x29, 0x34, 0x39, 0xd8, 0x0a, 0xc0, 0x78, 0x14,
	0x4a, 0x91, 0x0e, 0x61, 0xb2, 0xcc, 0x8d, 0x7f, 0x33, 0x02, 0x76, 0x1b, 0x76, 0xf3, 0x3f, 0x9c,
	0x11, 0xbf, 0x96, 0xb5, 0x32, 0x8d, 0xdf, 0x88, 0x3d, 0x82, 0xeb, 0x9d, 0x08, 0x11, 0x49, 0xcb,
	0x53, 0x8c, 0xa6, 0xd7, 0x7c, 0x56, 0x15, 0xf5, 0x4e, 0xbb, 0x4c, 0xb4, 0xd5, 0xf2, 0xf4, 0xab,
	0xe9, 0xf5, 0xfa, 0xd7, 0x0c, 0xf8, 0x65, 0x65, 0x2f, 0x66, 0x1f, 0xfe, 0x93, 0xfd, 0x10, 0xf6,
	0x3b, 0x27, 0x45, 0xf7, 0xa7, 0xd6, 0xa6, 0x2a, 0xea, 0x79, 0xbb, 0xcc, 0xf0, 0xbc, 0xd9, 0x0a,
	0xc0, 0x0e, 0x3d, 0x7a, 0x72, 0x5b, 0x1d, 0xf8, 0x93, 0xb1, 0xb8, 0x1d, 0xfa, 0x2f, 0x19, 0xb0,
	0xfb, 0xb0, 0x90, 0x91, 0xba, 0x73, 0xfd, 0x69, 0xd6, 0x21, 0xa1, 0x69, 0xc3, 0x5d, 0x98, 0x7b,
	0x4d, 0xbd, 0xb0, 0xda, 0x46, 0x7e, 0x54, 0x15, 0xf5, 0x5e, 0xfb, 0x17, 0xa4, 0x74, 0x65, 0x82,
	0x14, 0xa4, 0xd0, 0x59, 0xfe, 0x6c, 0x94, 0x27, 0xf2, 0xd9, 0xb2, 0xc7, 0x70, 0xc3, 0xab, 0x7c,
	0x6c, 0x0c, 0x51, 0xf4, 0x1e, 0x4f, 0x84, 0xe9, 0xb4, 0xe2, 0xcf, 0xf3, 0xc6, 0x03, 0xaf, 0xd2,
	0xe9, 0x8f, 0x93, 0xf2, 0x3e, 0x0b, 0x4c, 0x43, 0x39, 0x5e, 0x3a, 0x7f, 0x51, 0xcd, 0xea, 0xc5,
	0xe1, 0xa7, 0xe6, 0x6a, 0xaf, 0xdf, 0xfc, 0xf7, 0x15, 0xdb, 0x29, 0x9c, 0xdd, 0x81, 0x3d, 0x13,
	0xc6, 0x24, 0xfe, 0x32, 0x77, 0xd9, 0x35, 0xe1, 0x43, 0x1a, 0xd9, 0x03, 0x58, 0x2a, 0x72, 0x1e,
	0xa5, 0x1b, 0x6c, 0xd4, 0xc4, 0x5f, 0xe5, 0x1b, 0x59, 0x24, 0xf6, 0x76, 0x44, 0xec, 0x16, 0x94,
	0xc1, 0xc9, 0x9f, 0x3a, 0xf2, 0xd7, 0x55, 0x51, 0x1f, 0xb4, 0xd3, 0xb4, 0x2d, 0xf3, 0x07, 0xbb,
	0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x43, 0x40, 0x74, 0x9c, 0xd2, 0x02, 0x00, 0x00,
}
