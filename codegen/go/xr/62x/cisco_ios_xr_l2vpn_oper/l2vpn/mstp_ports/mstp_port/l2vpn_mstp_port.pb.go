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
// source: l2vpn_mstp_port.proto

package cisco_ios_xr_l2vpn_oper_l2vpn_mstp_ports_mstp_port

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

type L2VpnMstpPort_KEYS struct {
	Interface            string   `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMstpPort_KEYS) Reset()         { *m = L2VpnMstpPort_KEYS{} }
func (m *L2VpnMstpPort_KEYS) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstpPort_KEYS) ProtoMessage()    {}
func (*L2VpnMstpPort_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_04b64d9f77acb1fa, []int{0}
}

func (m *L2VpnMstpPort_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstpPort_KEYS.Unmarshal(m, b)
}
func (m *L2VpnMstpPort_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstpPort_KEYS.Marshal(b, m, deterministic)
}
func (m *L2VpnMstpPort_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstpPort_KEYS.Merge(m, src)
}
func (m *L2VpnMstpPort_KEYS) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstpPort_KEYS.Size(m)
}
func (m *L2VpnMstpPort_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstpPort_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstpPort_KEYS proto.InternalMessageInfo

func (m *L2VpnMstpPort_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

type L2VpnMstiBdEntry struct {
	Bdid                 uint32   `protobuf:"varint,1,opt,name=bdid,proto3" json:"bdid,omitempty"`
	BdifCount            uint32   `protobuf:"varint,2,opt,name=bdif_count,json=bdifCount,proto3" json:"bdif_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L2VpnMstiBdEntry) Reset()         { *m = L2VpnMstiBdEntry{} }
func (m *L2VpnMstiBdEntry) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstiBdEntry) ProtoMessage()    {}
func (*L2VpnMstiBdEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_04b64d9f77acb1fa, []int{1}
}

func (m *L2VpnMstiBdEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstiBdEntry.Unmarshal(m, b)
}
func (m *L2VpnMstiBdEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstiBdEntry.Marshal(b, m, deterministic)
}
func (m *L2VpnMstiBdEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstiBdEntry.Merge(m, src)
}
func (m *L2VpnMstiBdEntry) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstiBdEntry.Size(m)
}
func (m *L2VpnMstiBdEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstiBdEntry.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstiBdEntry proto.InternalMessageInfo

func (m *L2VpnMstiBdEntry) GetBdid() uint32 {
	if m != nil {
		return m.Bdid
	}
	return 0
}

func (m *L2VpnMstiBdEntry) GetBdifCount() uint32 {
	if m != nil {
		return m.BdifCount
	}
	return 0
}

type L2VpnMstiEntry struct {
	CfgMsTi              uint32              `protobuf:"varint,1,opt,name=cfg_ms_ti,json=cfgMsTi,proto3" json:"cfg_ms_ti,omitempty"`
	RcvCount             uint32              `protobuf:"varint,2,opt,name=rcv_count,json=rcvCount,proto3" json:"rcv_count,omitempty"`
	AckCount             uint32              `protobuf:"varint,3,opt,name=ack_count,json=ackCount,proto3" json:"ack_count,omitempty"`
	NackCount            uint32              `protobuf:"varint,4,opt,name=nack_count,json=nackCount,proto3" json:"nack_count,omitempty"`
	FlushCount           uint32              `protobuf:"varint,5,opt,name=flush_count,json=flushCount,proto3" json:"flush_count,omitempty"`
	InterfaceCount       uint32              `protobuf:"varint,6,opt,name=interface_count,json=interfaceCount,proto3" json:"interface_count,omitempty"`
	BdCount              uint32              `protobuf:"varint,7,opt,name=bd_count,json=bdCount,proto3" json:"bd_count,omitempty"`
	MstiFlags            uint32              `protobuf:"varint,8,opt,name=msti_flags,json=mstiFlags,proto3" json:"msti_flags,omitempty"`
	MstiState            string              `protobuf:"bytes,9,opt,name=msti_state,json=mstiState,proto3" json:"msti_state,omitempty"`
	BdEntry              []*L2VpnMstiBdEntry `protobuf:"bytes,10,rep,name=bd_entry,json=bdEntry,proto3" json:"bd_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *L2VpnMstiEntry) Reset()         { *m = L2VpnMstiEntry{} }
func (m *L2VpnMstiEntry) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstiEntry) ProtoMessage()    {}
func (*L2VpnMstiEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_04b64d9f77acb1fa, []int{2}
}

func (m *L2VpnMstiEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstiEntry.Unmarshal(m, b)
}
func (m *L2VpnMstiEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstiEntry.Marshal(b, m, deterministic)
}
func (m *L2VpnMstiEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstiEntry.Merge(m, src)
}
func (m *L2VpnMstiEntry) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstiEntry.Size(m)
}
func (m *L2VpnMstiEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstiEntry.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstiEntry proto.InternalMessageInfo

func (m *L2VpnMstiEntry) GetCfgMsTi() uint32 {
	if m != nil {
		return m.CfgMsTi
	}
	return 0
}

func (m *L2VpnMstiEntry) GetRcvCount() uint32 {
	if m != nil {
		return m.RcvCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetAckCount() uint32 {
	if m != nil {
		return m.AckCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetNackCount() uint32 {
	if m != nil {
		return m.NackCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetFlushCount() uint32 {
	if m != nil {
		return m.FlushCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetInterfaceCount() uint32 {
	if m != nil {
		return m.InterfaceCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetBdCount() uint32 {
	if m != nil {
		return m.BdCount
	}
	return 0
}

func (m *L2VpnMstiEntry) GetMstiFlags() uint32 {
	if m != nil {
		return m.MstiFlags
	}
	return 0
}

func (m *L2VpnMstiEntry) GetMstiState() string {
	if m != nil {
		return m.MstiState
	}
	return ""
}

func (m *L2VpnMstiEntry) GetBdEntry() []*L2VpnMstiBdEntry {
	if m != nil {
		return m.BdEntry
	}
	return nil
}

type L2VpnMstpPort struct {
	MstpInterfaceHandle  string            `protobuf:"bytes,50,opt,name=mstp_interface_handle,json=mstpInterfaceHandle,proto3" json:"mstp_interface_handle,omitempty"`
	InterfaceName        string            `protobuf:"bytes,51,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	Protected            bool              `protobuf:"varint,52,opt,name=protected,proto3" json:"protected,omitempty"`
	ReferenceCount       uint32            `protobuf:"varint,53,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	MstiEntry            []*L2VpnMstiEntry `protobuf:"bytes,54,rep,name=msti_entry,json=mstiEntry,proto3" json:"msti_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *L2VpnMstpPort) Reset()         { *m = L2VpnMstpPort{} }
func (m *L2VpnMstpPort) String() string { return proto.CompactTextString(m) }
func (*L2VpnMstpPort) ProtoMessage()    {}
func (*L2VpnMstpPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_04b64d9f77acb1fa, []int{3}
}

func (m *L2VpnMstpPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L2VpnMstpPort.Unmarshal(m, b)
}
func (m *L2VpnMstpPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L2VpnMstpPort.Marshal(b, m, deterministic)
}
func (m *L2VpnMstpPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L2VpnMstpPort.Merge(m, src)
}
func (m *L2VpnMstpPort) XXX_Size() int {
	return xxx_messageInfo_L2VpnMstpPort.Size(m)
}
func (m *L2VpnMstpPort) XXX_DiscardUnknown() {
	xxx_messageInfo_L2VpnMstpPort.DiscardUnknown(m)
}

var xxx_messageInfo_L2VpnMstpPort proto.InternalMessageInfo

func (m *L2VpnMstpPort) GetMstpInterfaceHandle() string {
	if m != nil {
		return m.MstpInterfaceHandle
	}
	return ""
}

func (m *L2VpnMstpPort) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *L2VpnMstpPort) GetProtected() bool {
	if m != nil {
		return m.Protected
	}
	return false
}

func (m *L2VpnMstpPort) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}

func (m *L2VpnMstpPort) GetMstiEntry() []*L2VpnMstiEntry {
	if m != nil {
		return m.MstiEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*L2VpnMstpPort_KEYS)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mstp_ports.mstp_port.l2vpn_mstp_port_KEYS")
	proto.RegisterType((*L2VpnMstiBdEntry)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mstp_ports.mstp_port.l2vpn_msti_bd_entry")
	proto.RegisterType((*L2VpnMstiEntry)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mstp_ports.mstp_port.l2vpn_msti_entry")
	proto.RegisterType((*L2VpnMstpPort)(nil), "cisco_ios_xr_l2vpn_oper.l2vpn.mstp_ports.mstp_port.l2vpn_mstp_port")
}

func init() { proto.RegisterFile("l2vpn_mstp_port.proto", fileDescriptor_04b64d9f77acb1fa) }

var fileDescriptor_04b64d9f77acb1fa = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcd, 0x6e, 0xd4, 0x30,
	0x14, 0x85, 0x35, 0x9d, 0xd2, 0x26, 0xb7, 0x2a, 0x45, 0x2e, 0x95, 0xcc, 0x4f, 0xc5, 0x28, 0x12,
	0x62, 0x56, 0x59, 0xa4, 0x85, 0x17, 0x80, 0x42, 0x11, 0x82, 0xc5, 0x94, 0x0d, 0x2b, 0xcb, 0xb1,
	0x9d, 0xa9, 0xd5, 0xc4, 0x89, 0x6c, 0xcf, 0x08, 0x5e, 0x82, 0x35, 0x8f, 0x8b, 0x7c, 0x1d, 0x1c,
	0x18, 0xb1, 0x41, 0xec, 0xee, 0x9c, 0x73, 0xe6, 0xe8, 0xfa, 0xb3, 0x03, 0x67, 0x6d, 0xb5, 0x1d,
	0x0c, 0xeb, 0x9c, 0x1f, 0xd8, 0xd0, 0x5b, 0x5f, 0x0e, 0xb6, 0xf7, 0x3d, 0xa9, 0x84, 0x76, 0xa2,
	0x67, 0xba, 0x77, 0xec, 0xab, 0x65, 0x31, 0xd3, 0x0f, 0xca, 0x96, 0x38, 0x96, 0x29, 0xee, 0xa6,
	0xb1, 0xb8, 0x84, 0x87, 0x3b, 0x65, 0xec, 0xc3, 0xd5, 0x97, 0x1b, 0xf2, 0x14, 0x72, 0x6d, 0xbc,
	0xb2, 0x0d, 0x17, 0x8a, 0xce, 0x16, 0xb3, 0x65, 0xbe, 0x9a, 0x84, 0xe2, 0x1a, 0x4e, 0xd3, 0xbf,
	0x34, 0xab, 0x25, 0x53, 0xc6, 0xdb, 0x6f, 0x84, 0xc0, 0x7e, 0x2d, 0xb5, 0xc4, 0xfc, 0xf1, 0x0a,
	0x67, 0x72, 0x0e, 0x50, 0x4b, 0xdd, 0x30, 0xd1, 0x6f, 0x8c, 0xa7, 0x7b, 0xe8, 0xe4, 0x41, 0x79,
	0x1d, 0x84, 0xe2, 0xfb, 0x1c, 0x1e, 0xfc, 0x56, 0x15, 0x7b, 0x1e, 0x43, 0x2e, 0x9a, 0x35, 0xeb,
	0x1c, 0xf3, 0x7a, 0x2c, 0x3b, 0x14, 0xcd, 0xfa, 0xa3, 0xfb, 0xac, 0xc9, 0x13, 0xc8, 0xad, 0xd8,
	0xfe, 0x51, 0x97, 0x59, 0xb1, 0xc5, 0xb6, 0x60, 0x72, 0x71, 0x37, 0x9a, 0xf3, 0x68, 0x72, 0x71,
	0x17, 0xcd, 0x73, 0x00, 0x33, 0xb9, 0xfb, 0x71, 0x13, 0x93, 0xec, 0x67, 0x70, 0xd4, 0xb4, 0x1b,
	0x77, 0x3b, 0xfa, 0xf7, 0xd0, 0x07, 0x94, 0x62, 0xe0, 0x05, 0x9c, 0x24, 0x02, 0x63, 0xe8, 0x00,
	0x43, 0xf7, 0x93, 0x1c, 0x83, 0x8f, 0x20, 0xab, 0xe5, 0x98, 0x38, 0x8c, 0xdb, 0xd7, 0x32, 0xed,
	0x80, 0xe7, 0x6c, 0x5a, 0xbe, 0x76, 0x34, 0x8b, 0x3b, 0x04, 0xe5, 0x6d, 0x10, 0x92, 0xed, 0x3c,
	0xf7, 0x8a, 0xe6, 0x11, 0x7b, 0x50, 0x6e, 0x82, 0x40, 0x6a, 0x2c, 0x46, 0x46, 0x14, 0x16, 0xf3,
	0xe5, 0x51, 0xf5, 0xae, 0xfc, 0xf7, 0x3b, 0x2f, 0xff, 0x72, 0x75, 0x61, 0xc3, 0xab, 0x30, 0x14,
	0x3f, 0xf6, 0xe0, 0x64, 0xe7, 0x45, 0x90, 0x0a, 0xce, 0xf0, 0xc7, 0x74, 0xfc, 0x5b, 0x6e, 0x64,
	0xab, 0x68, 0x85, 0x1b, 0x9e, 0x06, 0xf3, 0xfd, 0x2f, 0xef, 0x1a, 0x2d, 0xf2, 0x1c, 0x26, 0x2c,
	0xcc, 0xf0, 0x4e, 0xd1, 0x0b, 0x0c, 0x1f, 0x27, 0xf5, 0x13, 0xef, 0x54, 0x78, 0x67, 0xe1, 0xf1,
	0x2a, 0xe1, 0x95, 0xa4, 0x97, 0x8b, 0xd9, 0x32, 0x5b, 0x4d, 0x42, 0x40, 0x6e, 0x55, 0xa3, 0xac,
	0x32, 0x09, 0xf9, 0xcb, 0x88, 0x3c, 0xc9, 0x91, 0xab, 0x18, 0xc1, 0x45, 0x36, 0xaf, 0x90, 0xcd,
	0x9b, 0xff, 0x64, 0x13, 0xc1, 0x20, 0x7e, 0x44, 0x53, 0x1f, 0xe0, 0x67, 0x76, 0xf1, 0x33, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x6d, 0x8b, 0x59, 0x7f, 0x03, 0x00, 0x00,
}
