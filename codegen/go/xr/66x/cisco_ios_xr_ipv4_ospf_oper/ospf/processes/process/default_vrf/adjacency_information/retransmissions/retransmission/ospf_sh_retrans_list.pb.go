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
// source: ospf_sh_retrans_list.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_retransmissions_retransmission

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

type OspfShRetransList_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShRetransList_KEYS) Reset()         { *m = OspfShRetransList_KEYS{} }
func (m *OspfShRetransList_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShRetransList_KEYS) ProtoMessage()    {}
func (*OspfShRetransList_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_513be076a78114a4, []int{0}
}

func (m *OspfShRetransList_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList_KEYS.Unmarshal(m, b)
}
func (m *OspfShRetransList_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShRetransList_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList_KEYS.Merge(m, src)
}
func (m *OspfShRetransList_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShRetransList_KEYS.Size(m)
}
func (m *OspfShRetransList_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRetransList_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRetransList_KEYS proto.InternalMessageInfo

func (m *OspfShRetransList_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShRetransList_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShRetransList_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShLsaSum struct {
	HeaderLsaType           string   `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	HeaderLsaAge            uint32   `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	HeaderLsId              string   `protobuf:"bytes,3,opt,name=header_ls_id,json=headerLsId,proto3" json:"header_ls_id,omitempty"`
	HeaderAdvertisingRouter string   `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	HeaderSequenceNumber    uint32   `protobuf:"varint,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	HeaderLsaChecksum       uint32   `protobuf:"varint,6,opt,name=header_lsa_checksum,json=headerLsaChecksum,proto3" json:"header_lsa_checksum,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *OspfShLsaSum) Reset()         { *m = OspfShLsaSum{} }
func (m *OspfShLsaSum) String() string { return proto.CompactTextString(m) }
func (*OspfShLsaSum) ProtoMessage()    {}
func (*OspfShLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_513be076a78114a4, []int{1}
}

func (m *OspfShLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShLsaSum.Unmarshal(m, b)
}
func (m *OspfShLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShLsaSum.Marshal(b, m, deterministic)
}
func (m *OspfShLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShLsaSum.Merge(m, src)
}
func (m *OspfShLsaSum) XXX_Size() int {
	return xxx_messageInfo_OspfShLsaSum.Size(m)
}
func (m *OspfShLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShLsaSum proto.InternalMessageInfo

func (m *OspfShLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsId() string {
	if m != nil {
		return m.HeaderLsId
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *OspfShLsaSum) GetHeaderSequenceNumber() uint32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

func (m *OspfShLsaSum) GetHeaderLsaChecksum() uint32 {
	if m != nil {
		return m.HeaderLsaChecksum
	}
	return 0
}

type OspfShRetransList struct {
	RetransmissionNeighborId        string          `protobuf:"bytes,50,opt,name=retransmission_neighbor_id,json=retransmissionNeighborId,proto3" json:"retransmission_neighbor_id,omitempty"`
	RetransmissionNeighborIpAddress string          `protobuf:"bytes,51,opt,name=retransmission_neighbor_ip_address,json=retransmissionNeighborIpAddress,proto3" json:"retransmission_neighbor_ip_address,omitempty"`
	RetransmissionInterfaceName     string          `protobuf:"bytes,52,opt,name=retransmission_interface_name,json=retransmissionInterfaceName,proto3" json:"retransmission_interface_name,omitempty"`
	RetransmissionTimer             uint32          `protobuf:"varint,53,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	RetransmissionCount             uint32          `protobuf:"varint,54,opt,name=retransmission_count,json=retransmissionCount,proto3" json:"retransmission_count,omitempty"`
	RetransmissionAreaDb            []*OspfShLsaSum `protobuf:"bytes,55,rep,name=retransmission_area_db,json=retransmissionAreaDb,proto3" json:"retransmission_area_db,omitempty"`
	RetransmissionAsdb              []*OspfShLsaSum `protobuf:"bytes,56,rep,name=retransmission_asdb,json=retransmissionAsdb,proto3" json:"retransmission_asdb,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}        `json:"-"`
	XXX_unrecognized                []byte          `json:"-"`
	XXX_sizecache                   int32           `json:"-"`
}

func (m *OspfShRetransList) Reset()         { *m = OspfShRetransList{} }
func (m *OspfShRetransList) String() string { return proto.CompactTextString(m) }
func (*OspfShRetransList) ProtoMessage()    {}
func (*OspfShRetransList) Descriptor() ([]byte, []int) {
	return fileDescriptor_513be076a78114a4, []int{2}
}

func (m *OspfShRetransList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShRetransList.Unmarshal(m, b)
}
func (m *OspfShRetransList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShRetransList.Marshal(b, m, deterministic)
}
func (m *OspfShRetransList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShRetransList.Merge(m, src)
}
func (m *OspfShRetransList) XXX_Size() int {
	return xxx_messageInfo_OspfShRetransList.Size(m)
}
func (m *OspfShRetransList) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShRetransList.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShRetransList proto.InternalMessageInfo

func (m *OspfShRetransList) GetRetransmissionNeighborId() string {
	if m != nil {
		return m.RetransmissionNeighborId
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionNeighborIpAddress() string {
	if m != nil {
		return m.RetransmissionNeighborIpAddress
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionInterfaceName() string {
	if m != nil {
		return m.RetransmissionInterfaceName
	}
	return ""
}

func (m *OspfShRetransList) GetRetransmissionTimer() uint32 {
	if m != nil {
		return m.RetransmissionTimer
	}
	return 0
}

func (m *OspfShRetransList) GetRetransmissionCount() uint32 {
	if m != nil {
		return m.RetransmissionCount
	}
	return 0
}

func (m *OspfShRetransList) GetRetransmissionAreaDb() []*OspfShLsaSum {
	if m != nil {
		return m.RetransmissionAreaDb
	}
	return nil
}

func (m *OspfShRetransList) GetRetransmissionAsdb() []*OspfShLsaSum {
	if m != nil {
		return m.RetransmissionAsdb
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShRetransList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list_KEYS")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum")
	proto.RegisterType((*OspfShRetransList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list")
}

func init() { proto.RegisterFile("ospf_sh_retrans_list.proto", fileDescriptor_513be076a78114a4) }

var fileDescriptor_513be076a78114a4 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0xb5, 0x2d, 0xad, 0x84, 0x9b, 0x34, 0xe0, 0x44, 0xc5, 0x2d, 0x42, 0x84, 0x08, 0x50,
	0xb9, 0xac, 0x44, 0x1b, 0xfe, 0x08, 0x71, 0x09, 0x85, 0x43, 0x54, 0x94, 0x43, 0xda, 0x0b, 0x27,
	0xcb, 0xbb, 0x9e, 0x24, 0x86, 0xac, 0xbd, 0x78, 0xbc, 0x11, 0xb9, 0xf3, 0x00, 0xbc, 0x03, 0xef,
	0xc0, 0x85, 0x1b, 0x4f, 0x86, 0xe2, 0xdd, 0x2c, 0xdd, 0x55, 0x7a, 0xa6, 0x37, 0x67, 0xbe, 0xdf,
	0x37, 0xfe, 0xe2, 0xb1, 0x97, 0x1c, 0x19, 0x4c, 0x27, 0x1c, 0x67, 0xdc, 0x82, 0xb3, 0x42, 0x23,
	0x9f, 0x2b, 0x74, 0x61, 0x6a, 0x8d, 0x33, 0x14, 0x63, 0x85, 0xb1, 0xe1, 0xca, 0x20, 0xff, 0x66,
	0xb9, 0x4a, 0x17, 0x7d, 0xee, 0x69, 0x93, 0x82, 0x0d, 0x57, 0xab, 0x15, 0x17, 0x03, 0x22, 0xe0,
	0x7a, 0x15, 0x4a, 0x98, 0x88, 0x6c, 0xee, 0xf8, 0xc2, 0x4e, 0x42, 0x21, 0x3f, 0x8b, 0x18, 0x74,
	0xbc, 0xe4, 0x4a, 0x4f, 0x8c, 0x4d, 0x84, 0x53, 0x46, 0x87, 0xc5, 0x46, 0x89, 0x42, 0x54, 0x46,
	0x63, 0xed, 0x77, 0xef, 0x47, 0x40, 0x0e, 0x37, 0x65, 0xe2, 0xe7, 0x1f, 0x3e, 0x5d, 0xd0, 0x47,
	0xa4, 0x51, 0xec, 0xc4, 0xb5, 0x48, 0x80, 0x05, 0xdd, 0xe0, 0xf8, 0xf6, 0x78, 0xaf, 0xa8, 0x8d,
	0x44, 0x02, 0xf4, 0x09, 0xd9, 0x57, 0xda, 0x81, 0x9d, 0x88, 0x18, 0x72, 0x68, 0xcb, 0x43, 0xcd,
	0xb2, 0xea, 0xb1, 0x67, 0xe4, 0x8e, 0x06, 0x35, 0x9d, 0x45, 0xc6, 0x72, 0x21, 0xa5, 0x05, 0x44,
	0xb6, 0xed, 0xc1, 0xd6, 0xba, 0x3e, 0xc8, 0xcb, 0xbd, 0x9f, 0x5b, 0xa4, 0xb5, 0x8e, 0x34, 0x47,
	0xc1, 0x31, 0x4b, 0xe8, 0x53, 0xd2, 0x9a, 0x81, 0x90, 0x60, 0x7d, 0xc5, 0x2d, 0xd3, 0x75, 0x96,
	0x66, 0x5e, 0xfe, 0x88, 0xe2, 0x72, 0x99, 0x02, 0x7d, 0x4c, 0xf6, 0xaf, 0x70, 0x62, 0x9a, 0xa7,
	0x69, 0x8e, 0x1b, 0x25, 0x36, 0x98, 0x02, 0xed, 0x92, 0x46, 0x49, 0x71, 0x25, 0x8b, 0x20, 0x64,
	0xcd, 0x0c, 0x25, 0x7d, 0x43, 0x0e, 0x0b, 0x42, 0xc8, 0x05, 0x58, 0xa7, 0x50, 0xe9, 0x29, 0xb7,
	0x26, 0x73, 0x60, 0xd9, 0x2d, 0x8f, 0xdf, 0xcb, 0x81, 0xc1, 0x3f, 0x7d, 0xec, 0x65, 0xda, 0x27,
	0x07, 0x85, 0x17, 0xe1, 0x6b, 0x06, 0x7a, 0x75, 0x2e, 0x59, 0x12, 0x81, 0x65, 0x3b, 0x3e, 0x4b,
	0x27, 0x57, 0x2f, 0x0a, 0x71, 0xe4, 0x35, 0x1a, 0x92, 0xf6, 0x95, 0xe4, 0xf1, 0x0c, 0xe2, 0x2f,
	0x98, 0x25, 0x6c, 0xd7, 0x5b, 0xee, 0x96, 0xf1, 0xcf, 0x0a, 0xa1, 0xf7, 0x67, 0x87, 0x74, 0x36,
	0x0d, 0x8e, 0xbe, 0x25, 0x47, 0xd5, 0x19, 0xf3, 0xf2, 0xe0, 0x95, 0x64, 0x27, 0x3e, 0x3b, 0xab,
	0x12, 0xa3, 0x02, 0x18, 0x4a, 0x7a, 0x4e, 0x7a, 0xd7, 0xba, 0xd3, 0x72, 0x72, 0xa7, 0xbe, 0xcb,
	0xc3, 0x6b, 0xba, 0xa4, 0xc5, 0x24, 0xe9, 0x3b, 0xf2, 0xa0, 0xd6, 0xac, 0x76, 0x55, 0xfa, 0xbe,
	0xcf, 0xfd, 0x2a, 0x34, 0xac, 0x5c, 0x9c, 0xe7, 0xa4, 0x53, 0xeb, 0xe1, 0x54, 0x02, 0x96, 0xbd,
	0xf0, 0x07, 0xd3, 0xae, 0x6a, 0x97, 0x2b, 0x69, 0x83, 0x25, 0x36, 0x99, 0x76, 0xec, 0xe5, 0x26,
	0xcb, 0xd9, 0x4a, 0xa2, 0xbf, 0x03, 0x72, 0x50, 0xf3, 0x08, 0x0b, 0x82, 0xcb, 0x88, 0xbd, 0xea,
	0x6e, 0x1f, 0xef, 0x9d, 0x7c, 0x0f, 0xc2, 0xff, 0xf0, 0x3c, 0xc3, 0xda, 0x3b, 0x18, 0xd7, 0xfe,
	0xd8, 0xc0, 0x82, 0x78, 0x1f, 0xd1, 0x5f, 0x01, 0x69, 0xd7, 0xd3, 0xa3, 0x8c, 0xd8, 0xeb, 0x9b,
	0x14, 0x9d, 0xd6, 0xa2, 0xa3, 0x8c, 0xa2, 0x5d, 0xff, 0xe5, 0x3b, 0xfd, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0x81, 0xa8, 0xd1, 0x17, 0x05, 0x00, 0x00,
}