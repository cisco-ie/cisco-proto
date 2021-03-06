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
// source: ptp_summary_info.proto

package cisco_ios_xr_ptp_oper_ptp_nodes_node_summary

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

type PtpSummaryInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PtpSummaryInfo_KEYS) Reset()         { *m = PtpSummaryInfo_KEYS{} }
func (m *PtpSummaryInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*PtpSummaryInfo_KEYS) ProtoMessage()    {}
func (*PtpSummaryInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b38f0035fb65df8a, []int{0}
}

func (m *PtpSummaryInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpSummaryInfo_KEYS.Unmarshal(m, b)
}
func (m *PtpSummaryInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpSummaryInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *PtpSummaryInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpSummaryInfo_KEYS.Merge(m, src)
}
func (m *PtpSummaryInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_PtpSummaryInfo_KEYS.Size(m)
}
func (m *PtpSummaryInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpSummaryInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_PtpSummaryInfo_KEYS proto.InternalMessageInfo

func (m *PtpSummaryInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type PtpSummaryInfo struct {
	PortStateInitCount          uint32   `protobuf:"varint,50,opt,name=port_state_init_count,json=portStateInitCount,proto3" json:"port_state_init_count,omitempty"`
	PortStateListeningCount     uint32   `protobuf:"varint,51,opt,name=port_state_listening_count,json=portStateListeningCount,proto3" json:"port_state_listening_count,omitempty"`
	PortStatePassiveCount       uint32   `protobuf:"varint,52,opt,name=port_state_passive_count,json=portStatePassiveCount,proto3" json:"port_state_passive_count,omitempty"`
	PortStatePreMasterCount     uint32   `protobuf:"varint,53,opt,name=port_state_pre_master_count,json=portStatePreMasterCount,proto3" json:"port_state_pre_master_count,omitempty"`
	PortStateMasterCount        uint32   `protobuf:"varint,54,opt,name=port_state_master_count,json=portStateMasterCount,proto3" json:"port_state_master_count,omitempty"`
	PortStateSlaveCount         uint32   `protobuf:"varint,55,opt,name=port_state_slave_count,json=portStateSlaveCount,proto3" json:"port_state_slave_count,omitempty"`
	PortStateUncalibratedCount  uint32   `protobuf:"varint,56,opt,name=port_state_uncalibrated_count,json=portStateUncalibratedCount,proto3" json:"port_state_uncalibrated_count,omitempty"`
	PortStateFaultyCount        uint32   `protobuf:"varint,57,opt,name=port_state_faulty_count,json=portStateFaultyCount,proto3" json:"port_state_faulty_count,omitempty"`
	TotalInterfaces             uint32   `protobuf:"varint,58,opt,name=total_interfaces,json=totalInterfaces,proto3" json:"total_interfaces,omitempty"`
	TotalInterfacesValidPortNum uint32   `protobuf:"varint,59,opt,name=total_interfaces_valid_port_num,json=totalInterfacesValidPortNum,proto3" json:"total_interfaces_valid_port_num,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *PtpSummaryInfo) Reset()         { *m = PtpSummaryInfo{} }
func (m *PtpSummaryInfo) String() string { return proto.CompactTextString(m) }
func (*PtpSummaryInfo) ProtoMessage()    {}
func (*PtpSummaryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b38f0035fb65df8a, []int{1}
}

func (m *PtpSummaryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PtpSummaryInfo.Unmarshal(m, b)
}
func (m *PtpSummaryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PtpSummaryInfo.Marshal(b, m, deterministic)
}
func (m *PtpSummaryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PtpSummaryInfo.Merge(m, src)
}
func (m *PtpSummaryInfo) XXX_Size() int {
	return xxx_messageInfo_PtpSummaryInfo.Size(m)
}
func (m *PtpSummaryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PtpSummaryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PtpSummaryInfo proto.InternalMessageInfo

func (m *PtpSummaryInfo) GetPortStateInitCount() uint32 {
	if m != nil {
		return m.PortStateInitCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStateListeningCount() uint32 {
	if m != nil {
		return m.PortStateListeningCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStatePassiveCount() uint32 {
	if m != nil {
		return m.PortStatePassiveCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStatePreMasterCount() uint32 {
	if m != nil {
		return m.PortStatePreMasterCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStateMasterCount() uint32 {
	if m != nil {
		return m.PortStateMasterCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStateSlaveCount() uint32 {
	if m != nil {
		return m.PortStateSlaveCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStateUncalibratedCount() uint32 {
	if m != nil {
		return m.PortStateUncalibratedCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetPortStateFaultyCount() uint32 {
	if m != nil {
		return m.PortStateFaultyCount
	}
	return 0
}

func (m *PtpSummaryInfo) GetTotalInterfaces() uint32 {
	if m != nil {
		return m.TotalInterfaces
	}
	return 0
}

func (m *PtpSummaryInfo) GetTotalInterfacesValidPortNum() uint32 {
	if m != nil {
		return m.TotalInterfacesValidPortNum
	}
	return 0
}

func init() {
	proto.RegisterType((*PtpSummaryInfo_KEYS)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.summary.ptp_summary_info_KEYS")
	proto.RegisterType((*PtpSummaryInfo)(nil), "cisco_ios_xr_ptp_oper.ptp.nodes.node.summary.ptp_summary_info")
}

func init() { proto.RegisterFile("ptp_summary_info.proto", fileDescriptor_b38f0035fb65df8a) }

var fileDescriptor_b38f0035fb65df8a = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd2, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x06, 0x22, 0x2e, 0x20, 0x8e, 0xea, 0xb6, 0xb2, 0x21, 0x8e, 0x9d, 0x26, 0x48,
	0x41, 0xb7, 0x39, 0x75, 0x5e, 0xc4, 0x1f, 0x30, 0xd4, 0x31, 0x36, 0x14, 0x3c, 0x85, 0xac, 0xcb,
	0x24, 0xd0, 0x26, 0x21, 0x79, 0x1d, 0xee, 0xbf, 0xf4, 0x4f, 0x92, 0xa4, 0x35, 0xc4, 0xe2, 0xa5,
	0x87, 0xf7, 0xfd, 0x7e, 0xfa, 0xde, 0x21, 0xa8, 0x21, 0x41, 0x62, 0x9d, 0xa5, 0x29, 0x51, 0x5b,
	0xcc, 0xf8, 0x5a, 0x44, 0x52, 0x09, 0x10, 0xc1, 0x59, 0xcc, 0x74, 0x2c, 0x30, 0x13, 0x1a, 0x7f,
	0x29, 0x6c, 0x4a, 0x42, 0x52, 0x15, 0x49, 0x90, 0x11, 0x17, 0x2b, 0xaa, 0xed, 0x37, 0x2a, 0x60,
	0x77, 0x80, 0xea, 0xe5, 0xff, 0xe0, 0xe7, 0xc7, 0x8f, 0x45, 0xd0, 0x46, 0x55, 0x53, 0xc4, 0x9c,
	0xa4, 0x34, 0xac, 0x74, 0x2a, 0xbd, 0xea, 0x7c, 0xcf, 0x0c, 0xa6, 0x24, 0xa5, 0xdd, 0xef, 0x1d,
	0x54, 0x2b, 0xb3, 0xe0, 0x1c, 0xd5, 0xa5, 0x50, 0x80, 0x35, 0x10, 0xa0, 0x98, 0x71, 0x06, 0x38,
	0x16, 0x19, 0x87, 0xf0, 0xa2, 0x53, 0xe9, 0xed, 0xcf, 0x03, 0x13, 0x2e, 0x4c, 0x36, 0xe1, 0x0c,
	0xee, 0x4d, 0x12, 0x8c, 0x51, 0xcb, 0x23, 0x09, 0xd3, 0x40, 0x39, 0xe3, 0x9f, 0x85, 0xeb, 0x5b,
	0xd7, 0x74, 0xee, 0xe5, 0x37, 0xcf, 0xf1, 0x08, 0x85, 0x1e, 0x96, 0x44, 0x6b, 0xb6, 0xa1, 0x05,
	0x1d, 0x58, 0x5a, 0x77, 0x74, 0x96, 0xa7, 0x39, 0xbc, 0x45, 0x6d, 0x1f, 0x2a, 0x8a, 0x53, 0xa2,
	0x81, 0xaa, 0xc2, 0x0e, 0x4b, 0x6b, 0x67, 0x8a, 0xbe, 0xda, 0x3c, 0xd7, 0x43, 0xd4, 0xf4, 0xf4,
	0x1f, 0x79, 0x69, 0xe5, 0x91, 0x93, 0x3e, 0xeb, 0xa3, 0x86, 0xc7, 0x74, 0x42, 0xdc, 0xad, 0x23,
	0xab, 0x0e, 0x9d, 0x5a, 0x98, 0x2c, 0x47, 0x77, 0xe8, 0xd8, 0x43, 0x19, 0x8f, 0x49, 0xc2, 0x96,
	0x8a, 0x00, 0x5d, 0x15, 0xf6, 0xca, 0xda, 0x96, 0xb3, 0x6f, 0x5e, 0xe5, 0xbf, 0x73, 0xd7, 0x24,
	0x4b, 0x60, 0x5b, 0xe0, 0xeb, 0xd2, 0xb9, 0x4f, 0x36, 0xcc, 0xd9, 0x29, 0xaa, 0x81, 0x00, 0x92,
	0x60, 0xc6, 0x81, 0xaa, 0x35, 0x89, 0xa9, 0x0e, 0x6f, 0x6c, 0xff, 0xc0, 0xce, 0x27, 0x6e, 0x1c,
	0x3c, 0xa0, 0x93, 0x72, 0x15, 0x6f, 0x48, 0xc2, 0x56, 0xd8, 0x2e, 0xe6, 0x59, 0x1a, 0x8e, 0xad,
	0x6c, 0x97, 0xe4, 0xbb, 0x29, 0xcd, 0x84, 0x82, 0x69, 0x96, 0x2e, 0x77, 0xed, 0xeb, 0xed, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x34, 0xc8, 0xf1, 0xd7, 0x02, 0x00, 0x00,
}
