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
// source: ospfv3_edm_proto_stats.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_statistics_protocol_stats_protocol_stat

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

type Ospfv3EdmProtoStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmProtoStats_KEYS) Reset()         { *m = Ospfv3EdmProtoStats_KEYS{} }
func (m *Ospfv3EdmProtoStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmProtoStats_KEYS) ProtoMessage()    {}
func (*Ospfv3EdmProtoStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e053cd6e3e6cc3d, []int{0}
}

func (m *Ospfv3EdmProtoStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmProtoStats_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3EdmProtoStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmProtoStats_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmProtoStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmProtoStats_KEYS.Merge(m, src)
}
func (m *Ospfv3EdmProtoStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmProtoStats_KEYS.Size(m)
}
func (m *Ospfv3EdmProtoStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmProtoStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmProtoStats_KEYS proto.InternalMessageInfo

func (m *Ospfv3EdmProtoStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3EdmProtoStats_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type Ospfv3EdmProtoStats struct {
	TotalInPackets       uint64   `protobuf:"varint,50,opt,name=total_in_packets,json=totalInPackets,proto3" json:"total_in_packets,omitempty"`
	HelloInPackets       uint64   `protobuf:"varint,51,opt,name=hello_in_packets,json=helloInPackets,proto3" json:"hello_in_packets,omitempty"`
	DbdesInPackets       uint64   `protobuf:"varint,52,opt,name=dbdes_in_packets,json=dbdesInPackets,proto3" json:"dbdes_in_packets,omitempty"`
	DbdesInLsAs          uint64   `protobuf:"varint,53,opt,name=dbdes_in_ls_as,json=dbdesInLsAs,proto3" json:"dbdes_in_ls_as,omitempty"`
	LsReqInPackets       uint64   `protobuf:"varint,54,opt,name=ls_req_in_packets,json=lsReqInPackets,proto3" json:"ls_req_in_packets,omitempty"`
	LsReqInLsAs          uint64   `protobuf:"varint,55,opt,name=ls_req_in_ls_as,json=lsReqInLsAs,proto3" json:"ls_req_in_ls_as,omitempty"`
	LsUpdInPackets       uint64   `protobuf:"varint,56,opt,name=ls_upd_in_packets,json=lsUpdInPackets,proto3" json:"ls_upd_in_packets,omitempty"`
	LsUpdInIgnored       uint64   `protobuf:"varint,57,opt,name=ls_upd_in_ignored,json=lsUpdInIgnored,proto3" json:"ls_upd_in_ignored,omitempty"`
	LsUpdInLsAs          uint64   `protobuf:"varint,58,opt,name=ls_upd_in_ls_as,json=lsUpdInLsAs,proto3" json:"ls_upd_in_ls_as,omitempty"`
	LsAckInPackets       uint64   `protobuf:"varint,59,opt,name=ls_ack_in_packets,json=lsAckInPackets,proto3" json:"ls_ack_in_packets,omitempty"`
	LsAckInLsAs          uint64   `protobuf:"varint,60,opt,name=ls_ack_in_ls_as,json=lsAckInLsAs,proto3" json:"ls_ack_in_ls_as,omitempty"`
	TotalOutPackets      uint64   `protobuf:"varint,61,opt,name=total_out_packets,json=totalOutPackets,proto3" json:"total_out_packets,omitempty"`
	HelloOutPackets      uint64   `protobuf:"varint,62,opt,name=hello_out_packets,json=helloOutPackets,proto3" json:"hello_out_packets,omitempty"`
	DbdesOutPackets      uint64   `protobuf:"varint,63,opt,name=dbdes_out_packets,json=dbdesOutPackets,proto3" json:"dbdes_out_packets,omitempty"`
	DbdesOutLsAs         uint64   `protobuf:"varint,64,opt,name=dbdes_out_ls_as,json=dbdesOutLsAs,proto3" json:"dbdes_out_ls_as,omitempty"`
	LsReqOutPackets      uint64   `protobuf:"varint,65,opt,name=ls_req_out_packets,json=lsReqOutPackets,proto3" json:"ls_req_out_packets,omitempty"`
	LsReqOutLsAs         uint64   `protobuf:"varint,66,opt,name=ls_req_out_ls_as,json=lsReqOutLsAs,proto3" json:"ls_req_out_ls_as,omitempty"`
	LsUpdOutPackets      uint64   `protobuf:"varint,67,opt,name=ls_upd_out_packets,json=lsUpdOutPackets,proto3" json:"ls_upd_out_packets,omitempty"`
	LsUpdOutLsAs         uint64   `protobuf:"varint,68,opt,name=ls_upd_out_ls_as,json=lsUpdOutLsAs,proto3" json:"ls_upd_out_ls_as,omitempty"`
	LsAckOutPackets      uint64   `protobuf:"varint,69,opt,name=ls_ack_out_packets,json=lsAckOutPackets,proto3" json:"ls_ack_out_packets,omitempty"`
	LsAckOutLsAs         uint64   `protobuf:"varint,70,opt,name=ls_ack_out_ls_as,json=lsAckOutLsAs,proto3" json:"ls_ack_out_ls_as,omitempty"`
	DroppedInGs          uint64   `protobuf:"varint,71,opt,name=dropped_in_gs,json=droppedInGs,proto3" json:"dropped_in_gs,omitempty"`
	ChecksumErr          uint64   `protobuf:"varint,72,opt,name=checksum_err,json=checksumErr,proto3" json:"checksum_err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3EdmProtoStats) Reset()         { *m = Ospfv3EdmProtoStats{} }
func (m *Ospfv3EdmProtoStats) String() string { return proto.CompactTextString(m) }
func (*Ospfv3EdmProtoStats) ProtoMessage()    {}
func (*Ospfv3EdmProtoStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e053cd6e3e6cc3d, []int{1}
}

func (m *Ospfv3EdmProtoStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3EdmProtoStats.Unmarshal(m, b)
}
func (m *Ospfv3EdmProtoStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3EdmProtoStats.Marshal(b, m, deterministic)
}
func (m *Ospfv3EdmProtoStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3EdmProtoStats.Merge(m, src)
}
func (m *Ospfv3EdmProtoStats) XXX_Size() int {
	return xxx_messageInfo_Ospfv3EdmProtoStats.Size(m)
}
func (m *Ospfv3EdmProtoStats) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3EdmProtoStats.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3EdmProtoStats proto.InternalMessageInfo

func (m *Ospfv3EdmProtoStats) GetTotalInPackets() uint64 {
	if m != nil {
		return m.TotalInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetHelloInPackets() uint64 {
	if m != nil {
		return m.HelloInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetDbdesInPackets() uint64 {
	if m != nil {
		return m.DbdesInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetDbdesInLsAs() uint64 {
	if m != nil {
		return m.DbdesInLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsReqInPackets() uint64 {
	if m != nil {
		return m.LsReqInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsReqInLsAs() uint64 {
	if m != nil {
		return m.LsReqInLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsUpdInPackets() uint64 {
	if m != nil {
		return m.LsUpdInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsUpdInIgnored() uint64 {
	if m != nil {
		return m.LsUpdInIgnored
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsUpdInLsAs() uint64 {
	if m != nil {
		return m.LsUpdInLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsAckInPackets() uint64 {
	if m != nil {
		return m.LsAckInPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsAckInLsAs() uint64 {
	if m != nil {
		return m.LsAckInLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetTotalOutPackets() uint64 {
	if m != nil {
		return m.TotalOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetHelloOutPackets() uint64 {
	if m != nil {
		return m.HelloOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetDbdesOutPackets() uint64 {
	if m != nil {
		return m.DbdesOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetDbdesOutLsAs() uint64 {
	if m != nil {
		return m.DbdesOutLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsReqOutPackets() uint64 {
	if m != nil {
		return m.LsReqOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsReqOutLsAs() uint64 {
	if m != nil {
		return m.LsReqOutLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsUpdOutPackets() uint64 {
	if m != nil {
		return m.LsUpdOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsUpdOutLsAs() uint64 {
	if m != nil {
		return m.LsUpdOutLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsAckOutPackets() uint64 {
	if m != nil {
		return m.LsAckOutPackets
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetLsAckOutLsAs() uint64 {
	if m != nil {
		return m.LsAckOutLsAs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetDroppedInGs() uint64 {
	if m != nil {
		return m.DroppedInGs
	}
	return 0
}

func (m *Ospfv3EdmProtoStats) GetChecksumErr() uint64 {
	if m != nil {
		return m.ChecksumErr
	}
	return 0
}

func init() {
	proto.RegisterType((*Ospfv3EdmProtoStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.protocol_stats.protocol_stat.ospfv3_edm_proto_stats_KEYS")
	proto.RegisterType((*Ospfv3EdmProtoStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.statistics.protocol_stats.protocol_stat.ospfv3_edm_proto_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_proto_stats.proto", fileDescriptor_3e053cd6e3e6cc3d) }

var fileDescriptor_3e053cd6e3e6cc3d = []byte{
	// 498 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd3, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc0, 0x71, 0x15, 0x21, 0x10, 0xdb, 0x36, 0x69, 0x7c, 0x40, 0x96, 0xe0, 0x50, 0x02, 0x45,
	0xa1, 0x48, 0x39, 0x10, 0x28, 0xdf, 0x1f, 0x01, 0x42, 0x89, 0x40, 0x80, 0x82, 0x7a, 0xe0, 0xc2,
	0xca, 0x5d, 0x6f, 0x53, 0xcb, 0x8e, 0x77, 0xbb, 0xb3, 0xa9, 0x78, 0x46, 0x9e, 0x0a, 0x79, 0x66,
	0x37, 0x9e, 0xd0, 0xde, 0xdc, 0xf1, 0x4f, 0xff, 0x49, 0xbd, 0xb6, 0xb8, 0x6d, 0xc0, 0x9e, 0x9c,
	0x8f, 0xa4, 0xce, 0x17, 0xd2, 0x3a, 0xe3, 0x8d, 0x04, 0x9f, 0x79, 0x18, 0xe2, 0x75, 0xf2, 0x5b,
	0x15, 0xa0, 0x8c, 0x2c, 0x0c, 0xc8, 0x3f, 0x4e, 0x16, 0xf6, 0xfc, 0x40, 0x06, 0x6f, 0xac, 0x76,
	0x43, 0xba, 0x6e, 0xac, 0xd2, 0x00, 0x1a, 0xe2, 0xd5, 0xb0, 0x69, 0x14, 0xe0, 0x0b, 0x15, 0x42,
	0xca, 0x54, 0xbc, 0x1b, 0xff, 0xec, 0xcf, 0xc5, 0xad, 0xcb, 0xf7, 0xcb, 0x2f, 0x93, 0x5f, 0x3f,
	0x93, 0x3b, 0x62, 0x2b, 0x14, 0x65, 0x9d, 0x2d, 0x74, 0xba, 0xb1, 0xbb, 0x31, 0xb8, 0x31, 0xdb,
	0x0c, 0xb3, 0x6f, 0xd9, 0x42, 0x27, 0x7b, 0xa2, 0x53, 0xd4, 0x5e, 0xbb, 0x93, 0x4c, 0x69, 0x42,
	0x57, 0x10, 0x6d, 0xaf, 0xa6, 0x0d, 0xeb, 0xff, 0xbd, 0x2e, 0x6e, 0x5e, 0xbe, 0x29, 0x19, 0x88,
	0x1d, 0x6f, 0x7c, 0x56, 0xc9, 0xa2, 0x96, 0x36, 0x53, 0xa5, 0xf6, 0x90, 0x3e, 0xda, 0xdd, 0x18,
	0x5c, 0x9d, 0x75, 0x70, 0x3e, 0xad, 0x7f, 0xd0, 0xb4, 0x91, 0xa7, 0xba, 0xaa, 0x0c, 0x97, 0x23,
	0x92, 0x38, 0x5f, 0x93, 0xf9, 0x71, 0xae, 0x81, 0xcb, 0xc7, 0x24, 0x71, 0xde, 0xca, 0xbb, 0xa2,
	0xb3, 0x92, 0x15, 0xc8, 0x0c, 0xd2, 0x27, 0xe8, 0x36, 0x83, 0xfb, 0x0a, 0x63, 0x48, 0x1e, 0x88,
	0x5e, 0x05, 0xd2, 0xe9, 0x33, 0xde, 0x3b, 0xa0, 0x5e, 0x05, 0x33, 0x7d, 0xd6, 0xf6, 0xee, 0x89,
	0x6e, 0x4b, 0x29, 0xf8, 0x94, 0x82, 0x01, 0xb2, 0xe0, 0xd2, 0xe6, 0x3c, 0xf8, 0x2c, 0x06, 0x8f,
	0x6c, 0xde, 0x06, 0xd7, 0x68, 0x31, 0xaf, 0x8d, 0xd3, 0x79, 0xfa, 0x7c, 0x8d, 0x4e, 0x69, 0x1a,
	0x76, 0x07, 0x4a, 0xbb, 0x5f, 0xc4, 0xdd, 0x08, 0xd9, 0xee, 0x4c, 0x95, 0x7c, 0xf7, 0xcb, 0x18,
	0x1c, 0xab, 0xf2, 0xff, 0x7f, 0x26, 0x50, 0x0a, 0xbe, 0x8a, 0x41, 0x84, 0x18, 0xdc, 0x17, 0x3d,
	0x3a, 0x40, 0xb3, 0xf4, 0xab, 0xe0, 0x6b, 0x74, 0x5d, 0xbc, 0xf1, 0x7d, 0xe9, 0x63, 0x71, 0x5f,
	0xf4, 0xe8, 0x08, 0xb9, 0x7d, 0x43, 0x16, 0x6f, 0xac, 0x5b, 0x3a, 0x1a, 0x6e, 0xdf, 0x92, 0xc5,
	0x1b, 0xcc, 0xee, 0x89, 0x6e, 0x6b, 0xe9, 0x97, 0xbe, 0x43, 0xb9, 0x15, 0x25, 0xfe, 0xd4, 0x87,
	0x22, 0x09, 0xa7, 0xc3, 0x9b, 0x63, 0x6a, 0xe2, 0x01, 0xb1, 0xe6, 0x7d, 0xb1, 0xc3, 0x30, 0x45,
	0xdf, 0x53, 0x34, 0x52, 0x16, 0x6d, 0x1e, 0x3b, 0x8f, 0x7e, 0x88, 0xd1, 0x23, 0x9b, 0x5f, 0x88,
	0x46, 0x4c, 0xd1, 0x8f, 0x31, 0x4a, 0x94, 0x45, 0x9b, 0x47, 0xcf, 0xa3, 0x93, 0x18, 0x1d, 0xab,
	0xf2, 0x42, 0x34, 0x62, 0x8a, 0x7e, 0x8a, 0x51, 0xa2, 0x18, 0xed, 0x8b, 0xed, 0xdc, 0x19, 0x6b,
	0x35, 0xbe, 0x21, 0x73, 0x48, 0x0f, 0xc3, 0xbb, 0x4e, 0xc3, 0x69, 0x7d, 0x08, 0xcd, 0x37, 0xaf,
	0x4e, 0xb5, 0x2a, 0x61, 0xb9, 0x90, 0xda, 0xb9, 0xf4, 0x33, 0x91, 0x38, 0x9b, 0x38, 0x77, 0x7c,
	0x0d, 0x3f, 0xdf, 0xd1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x3a, 0x0a, 0x4d, 0xbc, 0x04,
	0x00, 0x00,
}
