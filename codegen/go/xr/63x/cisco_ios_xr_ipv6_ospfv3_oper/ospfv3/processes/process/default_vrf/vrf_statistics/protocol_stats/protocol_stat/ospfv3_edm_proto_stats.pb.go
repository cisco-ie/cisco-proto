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

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_default_vrf_vrf_statistics_protocol_stats_protocol_stat

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
	proto.RegisterType((*Ospfv3EdmProtoStats_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.vrf_statistics.protocol_stats.protocol_stat.ospfv3_edm_proto_stats_KEYS")
	proto.RegisterType((*Ospfv3EdmProtoStats)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.default_vrf.vrf_statistics.protocol_stats.protocol_stat.ospfv3_edm_proto_stats")
}

func init() { proto.RegisterFile("ospfv3_edm_proto_stats.proto", fileDescriptor_3e053cd6e3e6cc3d) }

var fileDescriptor_3e053cd6e3e6cc3d = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd3, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc0, 0x71, 0x15, 0x21, 0x10, 0xde, 0xd6, 0xae, 0x39, 0xa0, 0x48, 0x70, 0x18, 0x85, 0xa1,
	0x32, 0xa4, 0x1e, 0x28, 0x8c, 0xdf, 0x3f, 0x0a, 0x94, 0x51, 0x81, 0x00, 0x15, 0xed, 0xc0, 0xc9,
	0xca, 0x1c, 0xa7, 0x8b, 0x92, 0xc6, 0x9e, 0x9f, 0x53, 0xf1, 0x37, 0xf2, 0x57, 0xa1, 0xbc, 0x67,
	0x37, 0xee, 0xd6, 0x5b, 0xf6, 0xf2, 0xd1, 0xf7, 0x65, 0x75, 0xc2, 0xee, 0x2a, 0xd0, 0xd9, 0x6a,
	0xcc, 0x65, 0xba, 0xe4, 0xda, 0x28, 0xab, 0x38, 0xd8, 0xc4, 0xc2, 0x08, 0xaf, 0xa3, 0x4a, 0xe4,
	0x20, 0x14, 0xcf, 0x15, 0xf0, 0xbf, 0x86, 0xe7, 0x7a, 0x75, 0xcc, 0x9d, 0x57, 0x5a, 0x9a, 0x11,
	0x5d, 0x37, 0x56, 0x48, 0x00, 0x09, 0xfe, 0x6a, 0x94, 0xca, 0x2c, 0xa9, 0x4b, 0xcb, 0x57, 0x26,
	0x1b, 0xad, 0x4c, 0x86, 0xcd, 0x1c, 0x6c, 0x2e, 0x5c, 0x58, 0xa8, 0x32, 0xdc, 0xe3, 0xff, 0x1c,
	0x2c, 0xd8, 0x9d, 0xed, 0xcf, 0xc3, 0xbf, 0x4d, 0xff, 0xfc, 0x8e, 0xee, 0xb1, 0x5d, 0xb7, 0x81,
	0x57, 0xc9, 0x52, 0xc6, 0x9d, 0x83, 0xce, 0xf0, 0xd6, 0x7c, 0xc7, 0xcd, 0x7e, 0x24, 0x4b, 0x19,
	0x1d, 0xb2, 0x6e, 0x5e, 0x59, 0x69, 0xb2, 0x44, 0x48, 0x42, 0xd7, 0x10, 0xed, 0xad, 0xa7, 0x0d,
	0x1b, 0xfc, 0xbb, 0xc9, 0x6e, 0x6f, 0xdf, 0x14, 0x0d, 0xd9, 0xbe, 0x55, 0x36, 0x29, 0x79, 0x5e,
	0x71, 0x9d, 0x88, 0x42, 0x5a, 0x88, 0x9f, 0x1c, 0x74, 0x86, 0xd7, 0xe7, 0x5d, 0x9c, 0xcf, 0xaa,
	0x5f, 0x34, 0x6d, 0xe4, 0xb9, 0x2c, 0x4b, 0x15, 0xca, 0x31, 0x49, 0x9c, 0x6f, 0xc8, 0xf4, 0x2c,
	0x95, 0x10, 0xca, 0xa7, 0x24, 0x71, 0xde, 0xca, 0xfb, 0xac, 0xbb, 0x96, 0x25, 0xf0, 0x04, 0xe2,
	0x67, 0xe8, 0x76, 0x9c, 0xfb, 0x0e, 0x13, 0x88, 0x1e, 0xb1, 0x7e, 0x09, 0xdc, 0xc8, 0x8b, 0xb0,
	0x77, 0x4c, 0xbd, 0x12, 0xe6, 0xf2, 0xa2, 0xed, 0x3d, 0x60, 0xbd, 0x96, 0x52, 0xf0, 0x39, 0x05,
	0x1d, 0x0c, 0x82, 0xb5, 0x4e, 0xc3, 0xe0, 0x0b, 0x1f, 0x3c, 0xd5, 0x69, 0x1b, 0xdc, 0xa0, 0xf9,
	0xa2, 0x52, 0x46, 0xa6, 0xf1, 0xcb, 0x0d, 0x3a, 0xa3, 0xa9, 0xdb, 0xed, 0x28, 0xed, 0x7e, 0xe5,
	0x77, 0x23, 0x0c, 0x76, 0x27, 0xa2, 0x08, 0x77, 0xbf, 0xf6, 0xc1, 0x89, 0x28, 0x2e, 0xff, 0x33,
	0x8e, 0x52, 0xf0, 0x8d, 0x0f, 0x22, 0xc4, 0xe0, 0x11, 0xeb, 0xd3, 0x01, 0xaa, 0xda, 0xae, 0x83,
	0x6f, 0xd1, 0xf5, 0xf0, 0xc6, 0xcf, 0xda, 0xfa, 0xe2, 0x11, 0xeb, 0xd3, 0x11, 0x86, 0xf6, 0x1d,
	0x59, 0xbc, 0xb1, 0x69, 0xe9, 0x68, 0x42, 0xfb, 0x9e, 0x2c, 0xde, 0x08, 0xec, 0x21, 0xeb, 0xb5,
	0x96, 0x9e, 0xf4, 0x03, 0xca, 0x5d, 0x2f, 0xf1, 0x51, 0x1f, 0xb3, 0xc8, 0x9d, 0x4e, 0xd8, 0x9c,
	0x50, 0x13, 0x0f, 0x28, 0x68, 0x3e, 0x64, 0xfb, 0x01, 0xa6, 0xe8, 0x47, 0x8a, 0x7a, 0x1a, 0x44,
	0x9b, 0x9f, 0x3d, 0x8c, 0x7e, 0xf2, 0xd1, 0x53, 0x9d, 0x5e, 0x89, 0x7a, 0x4c, 0xd1, 0xcf, 0x3e,
	0x4a, 0x34, 0x88, 0x36, 0x3f, 0x7d, 0x18, 0x9d, 0xfa, 0xe8, 0x44, 0x14, 0x57, 0xa2, 0x1e, 0x53,
	0xf4, 0x8b, 0x8f, 0x12, 0xc5, 0xe8, 0x80, 0xed, 0xa5, 0x46, 0x69, 0x2d, 0xf1, 0x0d, 0x59, 0x40,
	0x7c, 0xe2, 0xde, 0x75, 0x1a, 0xce, 0xaa, 0x13, 0x68, 0xbe, 0x79, 0x71, 0x2e, 0x45, 0x01, 0xf5,
	0x92, 0x4b, 0x63, 0xe2, 0xaf, 0x44, 0xfc, 0x6c, 0x6a, 0xcc, 0xd9, 0x0d, 0xfc, 0x7c, 0xc7, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x32, 0x28, 0x55, 0xcc, 0x04, 0x00, 0x00,
}
