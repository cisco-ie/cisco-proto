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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_areas_area_retransmissions_retransmission

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
	AreaId               uint32   `protobuf:"varint,3,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	InterfaceName        string   `protobuf:"bytes,4,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,5,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
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

func (m *OspfShRetransList_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
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
	proto.RegisterType((*OspfShRetransList_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.retransmissions.retransmission.ospf_sh_retrans_list_KEYS")
	proto.RegisterType((*OspfShLsaSum)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.retransmissions.retransmission.ospf_sh_lsa_sum")
	proto.RegisterType((*OspfShRetransList)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.areas.area.retransmissions.retransmission.ospf_sh_retrans_list")
}

func init() { proto.RegisterFile("ospf_sh_retrans_list.proto", fileDescriptor_513be076a78114a4) }

var fileDescriptor_513be076a78114a4 = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x96, 0x5b, 0x9a, 0xd2, 0x69, 0xd2, 0xc0, 0x26, 0x6a, 0x9d, 0x22, 0x44, 0x88, 0x00, 0x85,
	0x8b, 0x25, 0xda, 0xf0, 0x10, 0xe2, 0x12, 0x0a, 0x87, 0xa8, 0x28, 0x07, 0xb7, 0x17, 0x4e, 0xab,
	0xb5, 0x77, 0x92, 0xac, 0xa8, 0x1f, 0xec, 0xd8, 0x16, 0xfd, 0x3d, 0x20, 0x71, 0xe1, 0x67, 0xf0,
	0xc3, 0x90, 0xd7, 0x0f, 0x1a, 0x2b, 0xbd, 0xc2, 0x65, 0xb5, 0x3b, 0xdf, 0x43, 0xdf, 0xec, 0x8e,
	0x0d, 0xc7, 0x11, 0xc5, 0x0b, 0x4e, 0x2b, 0xae, 0x31, 0xd1, 0x22, 0x24, 0x7e, 0xa5, 0x28, 0x71,
	0x62, 0x1d, 0x25, 0x11, 0x43, 0x5f, 0x91, 0x1f, 0x71, 0x15, 0x11, 0xff, 0xa6, 0xb9, 0x8a, 0xb3,
	0x09, 0x37, 0xec, 0x28, 0x46, 0xed, 0xe4, 0xbb, 0x9c, 0xe7, 0x23, 0x11, 0x52, 0xb5, 0x73, 0x32,
	0xbd, 0x30, 0x8b, 0x23, 0x34, 0x0a, 0x32, 0xab, 0x53, 0x5a, 0x07, 0x8a, 0x48, 0x45, 0x21, 0x35,
	0xce, 0xa3, 0xdf, 0x16, 0x0c, 0x36, 0xa5, 0xe0, 0xe7, 0x1f, 0x3f, 0x5f, 0xb0, 0xc7, 0xd0, 0x2e,
	0xbd, 0x79, 0x28, 0x02, 0xb4, 0xad, 0xa1, 0x35, 0xde, 0x73, 0xf7, 0xcb, 0xda, 0x5c, 0x04, 0xc8,
	0x06, 0x70, 0x37, 0xd3, 0x8b, 0x02, 0xde, 0x32, 0xf0, 0x6e, 0xa6, 0x17, 0x06, 0x3a, 0x82, 0xdd,
	0x3c, 0x02, 0x57, 0xd2, 0xde, 0x1e, 0x5a, 0xe3, 0x8e, 0xdb, 0xca, 0x8f, 0x33, 0xc9, 0x9e, 0xc2,
	0x81, 0x0a, 0x13, 0xd4, 0x0b, 0xe1, 0x63, 0xa1, 0xbc, 0x63, 0x94, 0x9d, 0xba, 0x6a, 0xf4, 0xcf,
	0xe1, 0x5e, 0x88, 0x6a, 0xb9, 0xf2, 0x22, 0xcd, 0x85, 0x94, 0x1a, 0x89, 0xec, 0x1d, 0x43, 0xec,
	0x56, 0xf5, 0x69, 0x51, 0x1e, 0x7d, 0xdf, 0x82, 0x6e, 0xd5, 0xc6, 0x15, 0x09, 0x4e, 0x69, 0xc0,
	0x9e, 0x41, 0x77, 0x85, 0x42, 0xa2, 0x36, 0x95, 0xe4, 0x3a, 0xae, 0xf2, 0x77, 0x8a, 0xf2, 0x27,
	0x12, 0x97, 0xd7, 0x31, 0xb2, 0x27, 0x70, 0x70, 0x83, 0x27, 0x96, 0x45, 0x1f, 0x1d, 0xb7, 0x5d,
	0xd3, 0xa6, 0x4b, 0x64, 0x43, 0x68, 0xd7, 0xac, 0xaa, 0xa3, 0x3d, 0x17, 0x2a, 0xce, 0x4c, 0xb2,
	0xb7, 0x30, 0x28, 0x19, 0x42, 0x66, 0xa8, 0x13, 0x45, 0x2a, 0x5c, 0x72, 0x1d, 0xa5, 0x09, 0xea,
	0xb2, 0xc1, 0xa3, 0x82, 0x30, 0xfd, 0x8b, 0xbb, 0x06, 0x66, 0x13, 0x38, 0x2c, 0xb5, 0x84, 0x5f,
	0x53, 0x0c, 0xf3, 0x7b, 0x49, 0x03, 0x0f, 0xb5, 0x69, 0xb8, 0xe3, 0xf6, 0x0b, 0xf4, 0xa2, 0x04,
	0xe7, 0x06, 0x63, 0x0e, 0xf4, 0x6e, 0x24, 0xf7, 0x57, 0xe8, 0x7f, 0xa1, 0x34, 0xb0, 0x5b, 0x46,
	0x72, 0xbf, 0x8e, 0x7f, 0x56, 0x02, 0xa3, 0x1f, 0x3b, 0xd0, 0xdf, 0xf4, 0xd8, 0xec, 0x1d, 0x1c,
	0xaf, 0xcf, 0x05, 0xaf, 0x2f, 0x5e, 0x49, 0xfb, 0xc4, 0x64, 0xb7, 0xd7, 0x19, 0xf3, 0x92, 0x30,
	0x93, 0xec, 0x1c, 0x46, 0xb7, 0xaa, 0xe3, 0xfa, 0xe5, 0x4e, 0x8d, 0xcb, 0xa3, 0x5b, 0x5c, 0xe2,
	0xf2, 0x25, 0xd9, 0x7b, 0x78, 0xd8, 0x30, 0x6b, 0x8c, 0xca, 0xc4, 0xf8, 0x3c, 0x58, 0x27, 0xcd,
	0xd6, 0x06, 0xe7, 0x05, 0xf4, 0x1b, 0x1e, 0x89, 0x0a, 0x50, 0xdb, 0x2f, 0xcd, 0xc5, 0xf4, 0xd6,
	0xb1, 0xcb, 0x1c, 0xda, 0x20, 0xf1, 0xa3, 0x34, 0x4c, 0xec, 0x57, 0x9b, 0x24, 0x67, 0x39, 0xc4,
	0x7e, 0x59, 0x70, 0xd8, 0xd0, 0x98, 0x71, 0x97, 0x9e, 0xfd, 0x7a, 0xb8, 0x3d, 0xde, 0x3f, 0xc9,
	0x9c, 0x7f, 0xf2, 0x0d, 0x3b, 0x8d, 0xc1, 0x77, 0x1b, 0x9d, 0x4c, 0x35, 0x8a, 0x0f, 0x1e, 0xfb,
	0x69, 0x41, 0xaf, 0x19, 0x97, 0xa4, 0x67, 0xbf, 0xf9, 0xaf, 0x59, 0x59, 0x23, 0x2b, 0x49, 0xcf,
	0x6b, 0x99, 0x3f, 0xe0, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0x9f, 0xac, 0xd9, 0x1f,
	0x05, 0x00, 0x00,
}