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
// source: ospfv3_edm_retrans.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_retransmission_list_process_table_retransmission

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

type Ospfv3EdmRetrans_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmRetrans_KEYS) Reset()         { *m = Ospfv3EdmRetrans_KEYS{} }
func (m *Ospfv3EdmRetrans_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmRetrans_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a67392e7e502b23, []int{0}
}

func (m *Ospfv3EdmRetrans_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRetrans_KEYS.Size(m)
}
func (m *Ospfv3EdmRetrans_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRetrans_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRetrans_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmRetrans_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *Ospfv3EdmRetrans_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type Ospfv3EdmLsaSum struct {
	HeaderLsaType           string   `protobuf:"bytes,1,opt,name=header_lsa_type,json=headerLsaType,proto3" json:"header_lsa_type,omitempty"`
	HeaderLsaAge            uint32   `protobuf:"varint,2,opt,name=header_lsa_age,json=headerLsaAge,proto3" json:"header_lsa_age,omitempty"`
	HeaderLsaId             string   `protobuf:"bytes,3,opt,name=header_lsa_id,json=headerLsaId,proto3" json:"header_lsa_id,omitempty"`
	HeaderAdvertisingRouter string   `protobuf:"bytes,4,opt,name=header_advertising_router,json=headerAdvertisingRouter,proto3" json:"header_advertising_router,omitempty"`
	HeaderSequenceNumber    int32    `protobuf:"zigzag32,5,opt,name=header_sequence_number,json=headerSequenceNumber,proto3" json:"header_sequence_number,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Ospfv3EdmLsaSum) Reset()         { *m = Ospfv3EdmLsaSum{} }
func (m *Ospfv3EdmLsaSum) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmLsaSum) ProtoMessage()    {}
func (*Ospfv3EdmLsaSum) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a67392e7e502b23, []int{1}
}

func (m *Ospfv3EdmLsaSum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Unmarshal(m, b)
}
func (m *Ospfv3EdmLsaSum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmLsaSum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmLsaSum.Merge(m, src)
}
func (m *Ospfv3EdmLsaSum) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmLsaSum.Size(m)
}
func (m *Ospfv3EdmLsaSum) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmLsaSum.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmLsaSum proto.InternalMessageInfo

func (m *Ospfv3EdmLsaSum) GetHeaderLsaType() string {
	if m != nil {
		return m.HeaderLsaType
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaAge() uint32 {
	if m != nil {
		return m.HeaderLsaAge
	}
	return 0
}

func (m *Ospfv3EdmLsaSum) GetHeaderLsaId() string {
	if m != nil {
		return m.HeaderLsaId
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderAdvertisingRouter() string {
	if m != nil {
		return m.HeaderAdvertisingRouter
	}
	return ""
}

func (m *Ospfv3EdmLsaSum) GetHeaderSequenceNumber() int32 {
	if m != nil {
		return m.HeaderSequenceNumber
	}
	return 0
}

type Ospfv3EdmRetrans struct {
	RetransmissionNeighborAddress string             `protobuf:"bytes,50,opt,name=retransmission_neighbor_address,json=retransmissionNeighborAddress,proto3" json:"retransmission_neighbor_address,omitempty"`
	IsRetransmissionvirtualLink   bool               `protobuf:"varint,51,opt,name=is_retransmissionvirtual_link,json=isRetransmissionvirtualLink,proto3" json:"is_retransmissionvirtual_link,omitempty"`
	RetransmissionvirtualLinkId   uint32             `protobuf:"varint,52,opt,name=retransmissionvirtual_link_id,json=retransmissionvirtualLinkId,proto3" json:"retransmissionvirtual_link_id,omitempty"`
	IsRetransmissionShamLink      bool               `protobuf:"varint,53,opt,name=is_retransmission_sham_link,json=isRetransmissionShamLink,proto3" json:"is_retransmission_sham_link,omitempty"`
	RetransmissionShamLinkId      uint32             `protobuf:"varint,54,opt,name=retransmission_sham_link_id,json=retransmissionShamLinkId,proto3" json:"retransmission_sham_link_id,omitempty"`
	RetransmissionTimer           uint32             `protobuf:"varint,55,opt,name=retransmission_timer,json=retransmissionTimer,proto3" json:"retransmission_timer,omitempty"`
	RetransmissionLength          uint32             `protobuf:"varint,56,opt,name=retransmission_length,json=retransmissionLength,proto3" json:"retransmission_length,omitempty"`
	RetransmissionvirtualLinkDb   []*Ospfv3EdmLsaSum `protobuf:"bytes,57,rep,name=retransmissionvirtual_link_db,json=retransmissionvirtualLinkDb,proto3" json:"retransmissionvirtual_link_db,omitempty"`
	RetransmissionAreaDb          []*Ospfv3EdmLsaSum `protobuf:"bytes,58,rep,name=retransmission_area_db,json=retransmissionAreaDb,proto3" json:"retransmission_area_db,omitempty"`
	RetransmissionAsdb            []*Ospfv3EdmLsaSum `protobuf:"bytes,59,rep,name=retransmission_asdb,json=retransmissionAsdb,proto3" json:"retransmission_asdb,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}           `json:"-"`
	XXX_unrecognized              []byte             `json:"-"`
	XXX_sizecache                 int32              `json:"-"`
}

func (m *Ospfv3EdmRetrans) Reset()         { *m = Ospfv3EdmRetrans{} }
func (m *Ospfv3EdmRetrans) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmRetrans) ProtoMessage()    {}
func (*Ospfv3EdmRetrans) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a67392e7e502b23, []int{2}
}

func (m *Ospfv3EdmRetrans) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmRetrans.Unmarshal(m, b)
}
func (m *Ospfv3EdmRetrans) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmRetrans.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmRetrans) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmRetrans.Merge(m, src)
}
func (m *Ospfv3EdmRetrans) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmRetrans.Size(m)
}
func (m *Ospfv3EdmRetrans) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmRetrans.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmRetrans proto.InternalMessageInfo

func (m *Ospfv3EdmRetrans) GetRetransmissionNeighborAddress() string {
	if m != nil {
		return m.RetransmissionNeighborAddress
	}
	return ""
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionvirtualLink() bool {
	if m != nil {
		return m.IsRetransmissionvirtualLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionvirtualLinkId() uint32 {
	if m != nil {
		return m.RetransmissionvirtualLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetIsRetransmissionShamLink() bool {
	if m != nil {
		return m.IsRetransmissionShamLink
	}
	return false
}

func (m *Ospfv3EdmRetrans) GetRetransmissionShamLinkId() uint32 {
	if m != nil {
		return m.RetransmissionShamLinkId
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionTimer() uint32 {
	if m != nil {
		return m.RetransmissionTimer
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionLength() uint32 {
	if m != nil {
		return m.RetransmissionLength
	}
	return 0
}

func (m *Ospfv3EdmRetrans) GetRetransmissionvirtualLinkDb() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionvirtualLinkDb
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAreaDb() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAreaDb
	}
	return nil
}

func (m *Ospfv3EdmRetrans) GetRetransmissionAsdb() []*Ospfv3EdmLsaSum {
	if m != nil {
		return m.RetransmissionAsdb
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3EdmRetrans_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.retransmission_list_process_table.retransmission.ospfv3_edm_retrans_KEYS")
	proto.RegisterType((*Ospfv3EdmLsaSum)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.retransmission_list_process_table.retransmission.ospfv3_edm_lsa_sum")
	proto.RegisterType((*Ospfv3EdmRetrans)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.retransmission_list_process_table.retransmission.ospfv3_edm_retrans")
}

func init() { proto.RegisterFile("ospfv3_edm_retrans.proto", fileDescriptor_7a67392e7e502b23) }

var fileDescriptor_7a67392e7e502b23 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0x87, 0xb5, 0x14, 0x50, 0x71, 0x9b, 0x16, 0xdc, 0xd2, 0x1a, 0x45, 0x15, 0x21, 0x02, 0x14,
	0x2e, 0x2b, 0xd1, 0x94, 0x02, 0x45, 0x3d, 0x04, 0x15, 0xa4, 0x88, 0x28, 0x87, 0x4d, 0x2f, 0x9c,
	0x2c, 0x6f, 0x3c, 0x49, 0xac, 0xee, 0x3f, 0x3c, 0xde, 0x88, 0x3e, 0x01, 0x37, 0x8e, 0xbc, 0x0b,
	0x17, 0x9e, 0x80, 0x07, 0xe2, 0x88, 0xd6, 0xbb, 0x09, 0xd9, 0x4d, 0xd3, 0x2b, 0xbd, 0xad, 0xe6,
	0xf7, 0xcd, 0xe4, 0xf3, 0x58, 0x56, 0x08, 0x8b, 0x31, 0x19, 0x4d, 0xdb, 0x1c, 0x64, 0xc8, 0x35,
	0x18, 0x2d, 0x22, 0x74, 0x13, 0x1d, 0x9b, 0x98, 0xe2, 0x50, 0xe1, 0x30, 0xe6, 0x2a, 0x46, 0xfe,
	0x55, 0x73, 0x95, 0x4c, 0x8f, 0x79, 0xc1, 0xc6, 0x09, 0x68, 0x37, 0xff, 0xce, 0xd8, 0x21, 0x20,
	0x02, 0xce, 0xbe, 0x5c, 0x09, 0x23, 0x91, 0x06, 0x86, 0x4f, 0xf5, 0xc8, 0x2d, 0x26, 0x86, 0x0a,
	0x51, 0xc5, 0x11, 0x0f, 0x14, 0x1a, 0x5e, 0x70, 0xdc, 0x08, 0x3f, 0x80, 0x0a, 0xd1, 0xfc, 0xee,
	0x90, 0xfd, 0x65, 0x23, 0xfe, 0xe9, 0xc3, 0xe7, 0x01, 0x7d, 0x42, 0x36, 0x67, 0xbd, 0x91, 0x08,
	0x81, 0x39, 0x0d, 0xa7, 0x75, 0xcf, 0xdb, 0x28, 0x6a, 0x7d, 0x11, 0x02, 0x7d, 0x46, 0xb6, 0x54,
	0x64, 0x40, 0x8f, 0xc4, 0x10, 0x72, 0xe8, 0x96, 0x85, 0x6a, 0xf3, 0xaa, 0xc5, 0x5e, 0x90, 0xfb,
	0x11, 0xa8, 0xf1, 0xc4, 0x8f, 0x35, 0x17, 0x52, 0x6a, 0x40, 0x64, 0x6b, 0x16, 0xdc, 0x9e, 0xd5,
	0x3b, 0x79, 0xb9, 0xf9, 0xc7, 0x21, 0x74, 0x41, 0x28, 0x40, 0xc1, 0x31, 0x0d, 0xe9, 0x73, 0xb2,
	0x3d, 0x01, 0x21, 0x41, 0xdb, 0x8a, 0xb9, 0x4c, 0x66, 0x3a, 0xb5, 0xbc, 0xdc, 0x43, 0x71, 0x7e,
	0x99, 0x00, 0x7d, 0x4a, 0xb6, 0x16, 0x38, 0x31, 0xce, 0x85, 0x6a, 0xde, 0xe6, 0x1c, 0xeb, 0x8c,
	0x81, 0x36, 0x49, 0x6d, 0x81, 0x52, 0xb2, 0x90, 0xd9, 0x98, 0x43, 0x5d, 0x49, 0x4f, 0xc8, 0xa3,
	0x82, 0x11, 0x72, 0x0a, 0xda, 0x28, 0x54, 0xd1, 0x98, 0xeb, 0x38, 0x35, 0xa0, 0xd9, 0x6d, 0xcb,
	0xef, 0xe7, 0x40, 0xe7, 0x5f, 0xee, 0xd9, 0x98, 0x1e, 0x91, 0xbd, 0xa2, 0x17, 0xe1, 0x4b, 0x0a,
	0x51, 0xb6, 0x9c, 0x34, 0xf4, 0x41, 0xb3, 0x3b, 0x0d, 0xa7, 0xf5, 0xc0, 0xdb, 0xcd, 0xd3, 0x41,
	0x11, 0xf6, 0x6d, 0xd6, 0xfc, 0xb1, 0x5e, 0x3a, 0x7a, 0x71, 0x17, 0xf4, 0x23, 0x79, 0x5c, 0xb9,
	0xd6, 0xa5, 0x5d, 0x1e, 0x5a, 0x9d, 0x83, 0x32, 0xd6, 0x2f, 0x6f, 0x96, 0xbe, 0x27, 0x07, 0x0a,
	0x79, 0x99, 0x99, 0x2a, 0x6d, 0x52, 0x11, 0xf0, 0x40, 0x45, 0x17, 0xac, 0xdd, 0x70, 0x5a, 0xeb,
	0x5e, 0x5d, 0xa1, 0x77, 0x15, 0xd3, 0x53, 0xd1, 0x45, 0x36, 0x63, 0xf5, 0x80, 0x6c, 0x91, 0x47,
	0x76, 0xdb, 0x75, 0xbd, 0x6a, 0x42, 0x57, 0xd2, 0x53, 0x52, 0x5f, 0xf2, 0xe0, 0x38, 0x11, 0x61,
	0x6e, 0xf1, 0xca, 0x5a, 0xb0, 0xaa, 0xc5, 0x60, 0x22, 0x42, 0xab, 0x70, 0x4a, 0xea, 0xab, 0x7a,
	0x33, 0x81, 0x63, 0x2b, 0xc0, 0xf4, 0x95, 0xcd, 0x5d, 0x49, 0x5f, 0x92, 0xdd, 0x4a, 0xbb, 0x51,
	0x21, 0x68, 0xf6, 0xda, 0xf6, 0xed, 0x94, 0xb3, 0xf3, 0x2c, 0xa2, 0x6d, 0xf2, 0xb0, 0xfa, 0xae,
	0x20, 0x1a, 0x9b, 0x09, 0x7b, 0x63, 0x7b, 0x2a, 0xf3, 0x7a, 0x36, 0xa3, 0xbf, 0x9d, 0x6b, 0x57,
	0x25, 0x7d, 0xf6, 0xb6, 0xb1, 0xd6, 0xda, 0x38, 0xfc, 0xe6, 0xb8, 0xff, 0xe1, 0xdd, 0xbb, 0xcb,
	0x4f, 0xec, 0x9a, 0x4b, 0x3b, 0xf3, 0xe9, 0x2f, 0x87, 0xec, 0x55, 0x7e, 0x44, 0x68, 0x10, 0xd9,
	0x39, 0x4e, 0x6e, 0xd8, 0x39, 0x2a, 0xf7, 0xd1, 0xd1, 0x20, 0xce, 0x7c, 0xfa, 0xd3, 0x21, 0x3b,
	0xd5, 0x03, 0xa0, 0xf4, 0xd9, 0xbb, 0x1b, 0x66, 0x4f, 0x2b, 0xf6, 0x28, 0x7d, 0xff, 0xae, 0xfd,
	0x83, 0x68, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x09, 0x9c, 0xc0, 0x3c, 0x06, 0x00, 0x00,
}