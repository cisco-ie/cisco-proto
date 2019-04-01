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
// source: fib_sh_rsrc_avail_state.proto

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_resource_resource_summary_info

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

type FibShRsrcAvailState_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	ProtocolName         string   `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShRsrcAvailState_KEYS) Reset()         { *m = FibShRsrcAvailState_KEYS{} }
func (m *FibShRsrcAvailState_KEYS) String() string { return proto.CompactTextString(m) }
func (*FibShRsrcAvailState_KEYS) ProtoMessage()    {}
func (*FibShRsrcAvailState_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9e169109da4431, []int{0}
}

func (m *FibShRsrcAvailState_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShRsrcAvailState_KEYS.Unmarshal(m, b)
}
func (m *FibShRsrcAvailState_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShRsrcAvailState_KEYS.Marshal(b, m, deterministic)
}
func (m *FibShRsrcAvailState_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShRsrcAvailState_KEYS.Merge(m, src)
}
func (m *FibShRsrcAvailState_KEYS) XXX_Size() int {
	return xxx_messageInfo_FibShRsrcAvailState_KEYS.Size(m)
}
func (m *FibShRsrcAvailState_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShRsrcAvailState_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_FibShRsrcAvailState_KEYS proto.InternalMessageInfo

func (m *FibShRsrcAvailState_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibShRsrcAvailState_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

type FibShRsrcShm struct {
	SrsCurrMode          uint32   `protobuf:"varint,1,opt,name=srs_curr_mode,json=srsCurrMode,proto3" json:"srs_curr_mode,omitempty"`
	SrsAvgAvail          uint64   `protobuf:"varint,2,opt,name=srs_avg_avail,json=srsAvgAvail,proto3" json:"srs_avg_avail,omitempty"`
	SrsMaxAvail          uint64   `protobuf:"varint,3,opt,name=srs_max_avail,json=srsMaxAvail,proto3" json:"srs_max_avail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FibShRsrcShm) Reset()         { *m = FibShRsrcShm{} }
func (m *FibShRsrcShm) String() string { return proto.CompactTextString(m) }
func (*FibShRsrcShm) ProtoMessage()    {}
func (*FibShRsrcShm) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9e169109da4431, []int{1}
}

func (m *FibShRsrcShm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShRsrcShm.Unmarshal(m, b)
}
func (m *FibShRsrcShm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShRsrcShm.Marshal(b, m, deterministic)
}
func (m *FibShRsrcShm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShRsrcShm.Merge(m, src)
}
func (m *FibShRsrcShm) XXX_Size() int {
	return xxx_messageInfo_FibShRsrcShm.Size(m)
}
func (m *FibShRsrcShm) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShRsrcShm.DiscardUnknown(m)
}

var xxx_messageInfo_FibShRsrcShm proto.InternalMessageInfo

func (m *FibShRsrcShm) GetSrsCurrMode() uint32 {
	if m != nil {
		return m.SrsCurrMode
	}
	return 0
}

func (m *FibShRsrcShm) GetSrsAvgAvail() uint64 {
	if m != nil {
		return m.SrsAvgAvail
	}
	return 0
}

func (m *FibShRsrcShm) GetSrsMaxAvail() uint64 {
	if m != nil {
		return m.SrsMaxAvail
	}
	return 0
}

type FibShRsrcAvailState struct {
	SrCurrMode           uint32          `protobuf:"varint,50,opt,name=sr_curr_mode,json=srCurrMode,proto3" json:"sr_curr_mode,omitempty"`
	SrShmState           []*FibShRsrcShm `protobuf:"bytes,51,rep,name=sr_shm_state,json=srShmState,proto3" json:"sr_shm_state,omitempty"`
	SrHwrsrcMode         []uint32        `protobuf:"varint,52,rep,packed,name=sr_hwrsrc_mode,json=srHwrsrcMode,proto3" json:"sr_hwrsrc_mode,omitempty"`
	SrHwrsrcInfo         []byte          `protobuf:"bytes,53,opt,name=sr_hwrsrc_info,json=srHwrsrcInfo,proto3" json:"sr_hwrsrc_info,omitempty"`
	SrShmwinOorCount     uint32          `protobuf:"varint,54,opt,name=sr_shmwin_oor_count,json=srShmwinOorCount,proto3" json:"sr_shmwin_oor_count,omitempty"`
	SrHwOorCount         uint32          `protobuf:"varint,55,opt,name=sr_hw_oor_count,json=srHwOorCount,proto3" json:"sr_hw_oor_count,omitempty"`
	SrInOorTs            uint64          `protobuf:"varint,56,opt,name=sr_in_oor_ts,json=srInOorTs,proto3" json:"sr_in_oor_ts,omitempty"`
	SrOutOorTs           uint64          `protobuf:"varint,57,opt,name=sr_out_oor_ts,json=srOutOorTs,proto3" json:"sr_out_oor_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FibShRsrcAvailState) Reset()         { *m = FibShRsrcAvailState{} }
func (m *FibShRsrcAvailState) String() string { return proto.CompactTextString(m) }
func (*FibShRsrcAvailState) ProtoMessage()    {}
func (*FibShRsrcAvailState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb9e169109da4431, []int{2}
}

func (m *FibShRsrcAvailState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FibShRsrcAvailState.Unmarshal(m, b)
}
func (m *FibShRsrcAvailState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FibShRsrcAvailState.Marshal(b, m, deterministic)
}
func (m *FibShRsrcAvailState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FibShRsrcAvailState.Merge(m, src)
}
func (m *FibShRsrcAvailState) XXX_Size() int {
	return xxx_messageInfo_FibShRsrcAvailState.Size(m)
}
func (m *FibShRsrcAvailState) XXX_DiscardUnknown() {
	xxx_messageInfo_FibShRsrcAvailState.DiscardUnknown(m)
}

var xxx_messageInfo_FibShRsrcAvailState proto.InternalMessageInfo

func (m *FibShRsrcAvailState) GetSrCurrMode() uint32 {
	if m != nil {
		return m.SrCurrMode
	}
	return 0
}

func (m *FibShRsrcAvailState) GetSrShmState() []*FibShRsrcShm {
	if m != nil {
		return m.SrShmState
	}
	return nil
}

func (m *FibShRsrcAvailState) GetSrHwrsrcMode() []uint32 {
	if m != nil {
		return m.SrHwrsrcMode
	}
	return nil
}

func (m *FibShRsrcAvailState) GetSrHwrsrcInfo() []byte {
	if m != nil {
		return m.SrHwrsrcInfo
	}
	return nil
}

func (m *FibShRsrcAvailState) GetSrShmwinOorCount() uint32 {
	if m != nil {
		return m.SrShmwinOorCount
	}
	return 0
}

func (m *FibShRsrcAvailState) GetSrHwOorCount() uint32 {
	if m != nil {
		return m.SrHwOorCount
	}
	return 0
}

func (m *FibShRsrcAvailState) GetSrInOorTs() uint64 {
	if m != nil {
		return m.SrInOorTs
	}
	return 0
}

func (m *FibShRsrcAvailState) GetSrOutOorTs() uint64 {
	if m != nil {
		return m.SrOutOorTs
	}
	return 0
}

func init() {
	proto.RegisterType((*FibShRsrcAvailState_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_summary_info.fib_sh_rsrc_avail_state_KEYS")
	proto.RegisterType((*FibShRsrcShm)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_summary_info.fib_sh_rsrc_shm")
	proto.RegisterType((*FibShRsrcAvailState)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_summary_info.fib_sh_rsrc_avail_state")
}

func init() { proto.RegisterFile("fib_sh_rsrc_avail_state.proto", fileDescriptor_cb9e169109da4431) }

var fileDescriptor_cb9e169109da4431 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x4f, 0x8b, 0x13, 0x41,
	0x10, 0xc5, 0x89, 0x11, 0x31, 0xbd, 0x89, 0x2b, 0xed, 0xc1, 0x01, 0x15, 0xc7, 0xa8, 0x90, 0x8b,
	0x73, 0xd8, 0xf5, 0xef, 0x71, 0x59, 0x04, 0x17, 0x59, 0x17, 0x12, 0x2f, 0x1e, 0xa4, 0xec, 0x74,
	0x7a, 0x76, 0x1a, 0xb6, 0xbb, 0x96, 0xaa, 0x99, 0x24, 0x1e, 0xfc, 0x0e, 0x7e, 0x18, 0x3f, 0xa0,
	0x74, 0xcd, 0xce, 0x18, 0x85, 0x3d, 0x7a, 0x19, 0x9a, 0x37, 0xbf, 0x7e, 0xf5, 0xfa, 0x95, 0x7a,
	0x54, 0xfa, 0x25, 0x70, 0x05, 0xc4, 0x64, 0xc1, 0xac, 0x8d, 0xbf, 0x00, 0xae, 0x4d, 0xed, 0x8a,
	0x4b, 0xc2, 0x1a, 0xf5, 0x57, 0xeb, 0xd9, 0x22, 0x78, 0x64, 0xd8, 0x12, 0x24, 0xd6, 0x62, 0x08,
	0x18, 0x01, 0x2f, 0x1d, 0x15, 0xa5, 0x5f, 0x16, 0x11, 0x57, 0x8e, 0xe5, 0xdb, 0x5e, 0xb1, 0x78,
	0xc1, 0xfd, 0xa9, 0x20, 0xc7, 0xd8, 0x90, 0x75, 0xfd, 0x01, 0xb8, 0x09, 0xc1, 0xd0, 0x77, 0xf0,
	0xb1, 0xc4, 0xe9, 0x37, 0xf5, 0xf0, 0x9a, 0xf9, 0xf0, 0xf1, 0xfd, 0x97, 0x85, 0x7e, 0xa0, 0x46,
	0xc9, 0x19, 0xa2, 0x09, 0x2e, 0x1b, 0xe4, 0x83, 0xd9, 0x68, 0x7e, 0x3b, 0x09, 0x9f, 0x4c, 0x70,
	0xfa, 0xa9, 0x9a, 0x74, 0x73, 0x5a, 0xe0, 0x86, 0x00, 0xe3, 0x4e, 0x4c, 0xd0, 0xf4, 0x87, 0xda,
	0xdf, 0x9d, 0xc0, 0x55, 0xd0, 0x53, 0x35, 0x61, 0x62, 0xb0, 0x0d, 0x11, 0x04, 0x5c, 0xb5, 0xc6,
	0x93, 0xf9, 0x1e, 0x13, 0x1f, 0x37, 0x44, 0xa7, 0xb8, 0x72, 0x1d, 0x63, 0xd6, 0xe7, 0x6d, 0x28,
	0xf1, 0xbe, 0x29, 0xcc, 0xd1, 0xfa, 0xfc, 0x28, 0x49, 0x1d, 0x13, 0xcc, 0xf6, 0x8a, 0x19, 0xf6,
	0xcc, 0xa9, 0xd9, 0x0a, 0x33, 0xfd, 0x35, 0x54, 0xf7, 0xaf, 0x79, 0xa1, 0xce, 0xd5, 0x98, 0x69,
	0x27, 0xc6, 0x81, 0xc4, 0x50, 0x4c, 0x7d, 0x8a, 0x9f, 0x03, 0x41, 0xb8, 0x0a, 0xed, 0x95, 0xec,
	0x30, 0x1f, 0xce, 0xf6, 0x0e, 0x62, 0xf1, 0x5f, 0xb7, 0x52, 0xfc, 0x53, 0x58, 0x8a, 0xb4, 0xa8,
	0xc2, 0x42, 0x42, 0x3f, 0x53, 0x77, 0x98, 0xa0, 0xda, 0xc8, 0x4f, 0x89, 0xfd, 0x32, 0x1f, 0xce,
	0x26, 0xf3, 0x31, 0xd3, 0x07, 0x11, 0x25, 0xf8, 0x5f, 0x54, 0xf2, 0xcc, 0x5e, 0xe5, 0x83, 0xd9,
	0xf8, 0x0f, 0x75, 0x12, 0x4b, 0xd4, 0x2f, 0xd4, 0xbd, 0xf6, 0x75, 0x1b, 0x1f, 0x01, 0x91, 0xc0,
	0x62, 0x13, 0xeb, 0xec, 0xb5, 0xf4, 0x70, 0x57, 0x86, 0x6e, 0x7c, 0x3c, 0x43, 0x3a, 0x4e, 0xba,
	0x7e, 0xae, 0xf6, 0xc5, 0x74, 0x07, 0x7d, 0x23, 0xa8, 0xb8, 0xf6, 0xd8, 0x63, 0xe9, 0xec, 0xca,
	0xb1, 0xe6, 0xec, 0xad, 0x6c, 0x65, 0xc4, 0x74, 0x92, 0xac, 0x3e, 0xb3, 0x7e, 0x92, 0xf6, 0x06,
	0xd8, 0xd4, 0x1d, 0xf1, 0x4e, 0x08, 0xc5, 0x74, 0xd6, 0xd4, 0x82, 0x2c, 0x6f, 0x49, 0x55, 0x87,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x33, 0x39, 0xe7, 0x1e, 0x03, 0x00, 0x00,
}