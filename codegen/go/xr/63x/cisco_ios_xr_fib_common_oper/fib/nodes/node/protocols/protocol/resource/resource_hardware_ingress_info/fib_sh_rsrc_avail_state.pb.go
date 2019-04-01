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

package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_resource_resource_hardware_ingress_info

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
	proto.RegisterType((*FibShRsrcAvailState_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_hardware_ingress_info.fib_sh_rsrc_avail_state_KEYS")
	proto.RegisterType((*FibShRsrcShm)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_hardware_ingress_info.fib_sh_rsrc_shm")
	proto.RegisterType((*FibShRsrcAvailState)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.resource.resource_hardware_ingress_info.fib_sh_rsrc_avail_state")
}

func init() { proto.RegisterFile("fib_sh_rsrc_avail_state.proto", fileDescriptor_cb9e169109da4431) }

var fileDescriptor_cb9e169109da4431 = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x91, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x15, 0x82, 0x10, 0xd9, 0x26, 0x14, 0x99, 0x03, 0x96, 0x00, 0x61, 0x02, 0x48, 0xb9,
	0xe0, 0x43, 0xcb, 0xdf, 0x63, 0x55, 0x21, 0x51, 0xa1, 0x52, 0x29, 0xe1, 0xc2, 0x69, 0xd8, 0xac,
	0xd7, 0xf1, 0x4a, 0xdd, 0x9d, 0x6a, 0xc6, 0x8e, 0x73, 0xe1, 0x41, 0x78, 0x18, 0xde, 0x0d, 0xed,
	0xb8, 0x36, 0x01, 0xa9, 0x57, 0x2e, 0xd6, 0xea, 0xf3, 0x6f, 0xbe, 0xf9, 0x66, 0x46, 0x3d, 0x29,
	0xdd, 0x1a, 0xb8, 0x02, 0x62, 0x32, 0xa0, 0xb7, 0xda, 0x5d, 0x02, 0xd7, 0xba, 0xb6, 0xf9, 0x15,
	0x61, 0x8d, 0x49, 0x69, 0x1c, 0x1b, 0x04, 0x87, 0x0c, 0x3b, 0x82, 0xc8, 0x1a, 0xf4, 0x1e, 0x03,
	0xe0, 0x95, 0xa5, 0xbc, 0x74, 0xeb, 0x3c, 0x60, 0x61, 0x59, 0xbe, 0x5d, 0x89, 0xc1, 0x4b, 0x1e,
	0x5e, 0x39, 0x59, 0xc6, 0x86, 0x8c, 0x1d, 0x1e, 0x50, 0x69, 0x2a, 0x5a, 0x4d, 0x16, 0x5c, 0xd8,
	0x90, 0x65, 0x06, 0x17, 0x4a, 0x9c, 0x7f, 0x57, 0x8f, 0x6f, 0x08, 0x02, 0x9f, 0x3f, 0x7e, 0x5b,
	0x25, 0x8f, 0xd4, 0x24, 0xb6, 0x80, 0xa0, 0xbd, 0x4d, 0x47, 0xd9, 0x68, 0x31, 0x59, 0xde, 0x8d,
	0xc2, 0x17, 0xed, 0x6d, 0xf2, 0x5c, 0xcd, 0xfa, 0x86, 0x1d, 0x70, 0x4b, 0x80, 0x69, 0x2f, 0x46,
	0x68, 0xfe, 0x43, 0x1d, 0xee, 0x77, 0xe0, 0xca, 0x27, 0x73, 0x35, 0x63, 0x62, 0x30, 0x0d, 0x11,
	0x78, 0x2c, 0x3a, 0xe3, 0xd9, 0xf2, 0x80, 0x89, 0x4f, 0x1b, 0xa2, 0x73, 0x2c, 0x6c, 0xcf, 0xe8,
	0xed, 0xa6, 0x0b, 0x25, 0xde, 0xb7, 0x85, 0x39, 0xd9, 0x6e, 0x4e, 0xa2, 0xd4, 0x33, 0x5e, 0xef,
	0xae, 0x99, 0xf1, 0xc0, 0x9c, 0xeb, 0x9d, 0x30, 0xf3, 0x5f, 0x63, 0xf5, 0xf0, 0x86, 0x09, 0x93,
	0x4c, 0x4d, 0x99, 0xf6, 0x62, 0x1c, 0x49, 0x0c, 0xc5, 0x34, 0xa4, 0xf8, 0x39, 0x12, 0x84, 0x2b,
	0xdf, 0x95, 0xa4, 0xc7, 0xd9, 0x78, 0x71, 0x70, 0xd4, 0xe6, 0xff, 0xe7, 0x3c, 0xf9, 0x3f, 0x9b,
	0x8b, 0xd9, 0x56, 0x95, 0x5f, 0x49, 0xfa, 0x17, 0xea, 0x1e, 0x13, 0x54, 0xad, 0xfc, 0x94, 0xfc,
	0xaf, 0xb3, 0xf1, 0x62, 0xb6, 0x9c, 0x32, 0x7d, 0x12, 0x51, 0x26, 0xf8, 0x8b, 0x8a, 0x9e, 0xe9,
	0x9b, 0x6c, 0xb4, 0x98, 0xfe, 0xa1, 0xce, 0x42, 0x89, 0xc9, 0x2b, 0xf5, 0xa0, 0x1b, 0xb3, 0x75,
	0x01, 0x10, 0x09, 0x0c, 0x36, 0xa1, 0x4e, 0xdf, 0xca, 0x42, 0xee, 0x4b, 0xd3, 0xd6, 0x85, 0x0b,
	0xa4, 0xd3, 0xa8, 0x27, 0x2f, 0xd5, 0xa1, 0x98, 0xee, 0xa1, 0xef, 0x04, 0x15, 0xd7, 0x01, 0x7b,
	0x2a, 0xcb, 0xbb, 0x76, 0xac, 0x39, 0x7d, 0x2f, 0xe7, 0x99, 0x30, 0x9d, 0x45, 0xab, 0xaf, 0x9c,
	0x3c, 0x8b, 0x07, 0x04, 0x6c, 0xea, 0x9e, 0xf8, 0x20, 0x84, 0x62, 0xba, 0x68, 0x6a, 0x41, 0xd6,
	0x77, 0x64, 0x67, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x16, 0x3b, 0x4f, 0xcd, 0x30, 0x03,
	0x00, 0x00,
}