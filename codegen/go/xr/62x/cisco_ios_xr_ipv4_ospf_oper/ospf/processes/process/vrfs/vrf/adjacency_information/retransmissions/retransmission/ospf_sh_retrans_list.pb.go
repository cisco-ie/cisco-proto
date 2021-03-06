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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_adjacency_information_retransmissions_retransmission

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,4,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
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

func (m *OspfShRetransList_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*OspfShRetransList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list_KEYS")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_lsa_sum")
	proto.RegisterType((*OspfShRetransList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.adjacency_information.retransmissions.retransmission.ospf_sh_retrans_list")
}

func init() { proto.RegisterFile("ospf_sh_retrans_list.proto", fileDescriptor_513be076a78114a4) }

var fileDescriptor_513be076a78114a4 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xe5, 0xf6, 0x6d, 0x5f, 0x98, 0x26, 0x0d, 0x6c, 0xa2, 0xe2, 0x14, 0x21, 0x42, 0x04,
	0xa8, 0x5c, 0x2c, 0xd1, 0x86, 0x3f, 0x42, 0x5c, 0x42, 0xe1, 0x10, 0x15, 0xe5, 0x90, 0xf6, 0xc2,
	0x69, 0xb5, 0xf6, 0x8e, 0x93, 0x85, 0xda, 0x6b, 0x76, 0xec, 0x88, 0x1c, 0xf9, 0x2c, 0x1c, 0x39,
	0x72, 0xe1, 0xc0, 0x87, 0x43, 0x59, 0x6f, 0x4c, 0x63, 0xa5, 0x57, 0xc4, 0xc5, 0x5a, 0xcf, 0xf3,
	0x9b, 0xd9, 0xc7, 0x33, 0xbb, 0x86, 0x43, 0x4d, 0x59, 0xcc, 0x69, 0xc6, 0x0d, 0xe6, 0x46, 0xa4,
	0xc4, 0x2f, 0x15, 0xe5, 0x41, 0x66, 0x74, 0xae, 0x59, 0x16, 0x29, 0x8a, 0x34, 0x57, 0x9a, 0xf8,
	0x17, 0xc3, 0x55, 0x36, 0x1f, 0x70, 0x4b, 0xeb, 0x0c, 0x4d, 0xb0, 0x5c, 0x2d, 0xb9, 0x08, 0x89,
	0x90, 0x56, 0xab, 0x60, 0x6e, 0x62, 0xfb, 0x08, 0x84, 0xfc, 0x28, 0x22, 0x4c, 0xa3, 0x05, 0x57,
	0x69, 0xac, 0x4d, 0x22, 0x72, 0xa5, 0xd3, 0xc0, 0xed, 0x92, 0x28, 0x22, 0xa5, 0x53, 0xaa, 0xbd,
	0xf7, 0xbf, 0x7b, 0xd0, 0xdd, 0x64, 0x88, 0x9f, 0xbd, 0xfb, 0x70, 0xce, 0x1e, 0x40, 0xc3, 0x6d,
	0xc3, 0x53, 0x91, 0xa0, 0xef, 0xf5, 0xbc, 0xa3, 0x9b, 0x93, 0x3d, 0x17, 0x1b, 0x8b, 0x04, 0x59,
	0x17, 0x6e, 0xcc, 0x4d, 0x5c, 0xca, 0x5b, 0x56, 0xfe, 0x7f, 0x6e, 0x62, 0x2b, 0x3d, 0x82, 0x7d,
	0x95, 0xe6, 0x68, 0x62, 0x11, 0x61, 0x09, 0x6c, 0x5b, 0xa0, 0x59, 0x45, 0x2d, 0xf6, 0x04, 0x6e,
	0xa5, 0xa8, 0xa6, 0xb3, 0x50, 0x1b, 0x2e, 0xa4, 0x34, 0x48, 0xe4, 0xff, 0x67, 0xc1, 0xd6, 0x2a,
	0x3e, 0x2c, 0xc3, 0xfd, 0x6f, 0x5b, 0xd0, 0x5a, 0xb9, 0xbd, 0x24, 0xc1, 0xa9, 0x48, 0xd8, 0x63,
	0x68, 0xcd, 0x50, 0x48, 0x34, 0x36, 0x92, 0x2f, 0xb2, 0x95, 0xcd, 0x66, 0x19, 0x7e, 0x4f, 0xe2,
	0x62, 0x91, 0x21, 0x7b, 0x08, 0xfb, 0x57, 0x38, 0x31, 0x2d, 0xed, 0x36, 0x27, 0x8d, 0x0a, 0x1b,
	0x4e, 0x91, 0xf5, 0xa0, 0x51, 0x51, 0x5c, 0x49, 0xe7, 0x18, 0x56, 0xcc, 0x48, 0xb2, 0x57, 0xd0,
	0x75, 0x84, 0x90, 0x73, 0x34, 0xb9, 0x22, 0x95, 0x4e, 0xb9, 0xd1, 0x45, 0x8e, 0xc6, 0xf9, 0xbe,
	0x53, 0x02, 0xc3, 0x3f, 0xfa, 0xc4, 0xca, 0x6c, 0x00, 0x07, 0x2e, 0x97, 0xf0, 0x73, 0x81, 0xe9,
	0xb2, 0x2f, 0x45, 0x12, 0xa2, 0xf1, 0x77, 0xac, 0x97, 0x4e, 0xa9, 0x9e, 0x3b, 0x71, 0x6c, 0x35,
	0x16, 0x40, 0xfb, 0x8a, 0xf3, 0x68, 0x86, 0xd1, 0x27, 0x2a, 0x12, 0x7f, 0xd7, 0xa6, 0xdc, 0xae,
	0xec, 0x9f, 0x3a, 0xa1, 0xff, 0x6b, 0x07, 0x3a, 0x9b, 0x66, 0xca, 0x5e, 0xc3, 0xe1, 0xfa, 0xf8,
	0x79, 0xd5, 0x78, 0x25, 0xfd, 0x63, 0xeb, 0xdd, 0x5f, 0x27, 0xc6, 0x0e, 0x18, 0x49, 0x76, 0x06,
	0xfd, 0x6b, 0xb3, 0xb3, 0x6a, 0x72, 0x27, 0xb6, 0xca, 0xfd, 0x6b, 0xaa, 0x64, 0x6e, 0x92, 0xec,
	0x0d, 0xdc, 0xab, 0x15, 0xab, 0x1d, 0x95, 0x81, 0xad, 0x73, 0x77, 0x1d, 0x1a, 0xad, 0x1d, 0x9c,
	0xa7, 0xd0, 0xa9, 0xd5, 0xc8, 0x55, 0x82, 0xc6, 0x7f, 0x66, 0x1b, 0xd3, 0x5e, 0xd7, 0x2e, 0x96,
	0xd2, 0x86, 0x94, 0x48, 0x17, 0x69, 0xee, 0x3f, 0xdf, 0x94, 0x72, 0xba, 0x94, 0xd8, 0x4f, 0x0f,
	0x0e, 0x6a, 0x39, 0xc2, 0xa0, 0xe0, 0x32, 0xf4, 0x5f, 0xf4, 0xb6, 0x8f, 0xf6, 0x8e, 0xbf, 0x7a,
	0xc1, 0xdf, 0xbe, 0xb6, 0x41, 0xed, 0x12, 0x4c, 0x6a, 0x5f, 0x35, 0x34, 0x28, 0xde, 0x86, 0xec,
	0x87, 0x07, 0xed, 0xba, 0x75, 0x92, 0xa1, 0xff, 0xf2, 0x9f, 0xf1, 0xcd, 0x6a, 0xbe, 0x49, 0x86,
	0xe1, 0xae, 0xfd, 0x17, 0x9e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x26, 0x71, 0xab, 0x29,
	0x05, 0x00, 0x00,
}
