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
// source: ospf_sh_neighbor.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_adjacency_information_neighbors_neighbor

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

type OspfShNeighbor_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighborAddress      string   `protobuf:"bytes,3,opt,name=neighbor_address,json=neighborAddress,proto3" json:"neighbor_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighbor_KEYS) Reset()         { *m = OspfShNeighbor_KEYS{} }
func (m *OspfShNeighbor_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor_KEYS) ProtoMessage()    {}
func (*OspfShNeighbor_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_776818f2d81e6901, []int{0}
}

func (m *OspfShNeighbor_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Unmarshal(m, b)
}
func (m *OspfShNeighbor_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShNeighbor_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor_KEYS.Merge(m, src)
}
func (m *OspfShNeighbor_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor_KEYS.Size(m)
}
func (m *OspfShNeighbor_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor_KEYS proto.InternalMessageInfo

func (m *OspfShNeighbor_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShNeighbor_KEYS) GetNeighborAddress() string {
	if m != nil {
		return m.NeighborAddress
	}
	return ""
}

type OspfShNeighborBfd struct {
	BfdIntfEnableMode    uint32   `protobuf:"varint,1,opt,name=bfd_intf_enable_mode,json=bfdIntfEnableMode,proto3" json:"bfd_intf_enable_mode,omitempty"`
	BfdStatusFlag        uint32   `protobuf:"varint,2,opt,name=bfd_status_flag,json=bfdStatusFlag,proto3" json:"bfd_status_flag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShNeighborBfd) Reset()         { *m = OspfShNeighborBfd{} }
func (m *OspfShNeighborBfd) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighborBfd) ProtoMessage()    {}
func (*OspfShNeighborBfd) Descriptor() ([]byte, []int) {
	return fileDescriptor_776818f2d81e6901, []int{1}
}

func (m *OspfShNeighborBfd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighborBfd.Unmarshal(m, b)
}
func (m *OspfShNeighborBfd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighborBfd.Marshal(b, m, deterministic)
}
func (m *OspfShNeighborBfd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighborBfd.Merge(m, src)
}
func (m *OspfShNeighborBfd) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighborBfd.Size(m)
}
func (m *OspfShNeighborBfd) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighborBfd.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighborBfd proto.InternalMessageInfo

func (m *OspfShNeighborBfd) GetBfdIntfEnableMode() uint32 {
	if m != nil {
		return m.BfdIntfEnableMode
	}
	return 0
}

func (m *OspfShNeighborBfd) GetBfdStatusFlag() uint32 {
	if m != nil {
		return m.BfdStatusFlag
	}
	return 0
}

type OspfShNeighbor struct {
	NeighborId             string             `protobuf:"bytes,50,opt,name=neighbor_id,json=neighborId,proto3" json:"neighbor_id,omitempty"`
	NeighborAddressXr      string             `protobuf:"bytes,51,opt,name=neighbor_address_xr,json=neighborAddressXr,proto3" json:"neighbor_address_xr,omitempty"`
	NeighborInterfaceName  string             `protobuf:"bytes,52,opt,name=neighbor_interface_name,json=neighborInterfaceName,proto3" json:"neighbor_interface_name,omitempty"`
	NeighborDrPriority     uint32             `protobuf:"varint,53,opt,name=neighbor_dr_priority,json=neighborDrPriority,proto3" json:"neighbor_dr_priority,omitempty"`
	NeighborState          string             `protobuf:"bytes,54,opt,name=neighbor_state,json=neighborState,proto3" json:"neighbor_state,omitempty"`
	DrBdrState             string             `protobuf:"bytes,55,opt,name=dr_bdr_state,json=drBdrState,proto3" json:"dr_bdr_state,omitempty"`
	NeighborDeadTimer      uint32             `protobuf:"varint,56,opt,name=neighbor_dead_timer,json=neighborDeadTimer,proto3" json:"neighbor_dead_timer,omitempty"`
	NeighborUpTime         uint32             `protobuf:"varint,57,opt,name=neighbor_up_time,json=neighborUpTime,proto3" json:"neighbor_up_time,omitempty"`
	NeighborMadjInterface  bool               `protobuf:"varint,58,opt,name=neighbor_madj_interface,json=neighborMadjInterface,proto3" json:"neighbor_madj_interface,omitempty"`
	NeighborBfdInformation *OspfShNeighborBfd `protobuf:"bytes,59,opt,name=neighbor_bfd_information,json=neighborBfdInformation,proto3" json:"neighbor_bfd_information,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *OspfShNeighbor) Reset()         { *m = OspfShNeighbor{} }
func (m *OspfShNeighbor) String() string { return proto.CompactTextString(m) }
func (*OspfShNeighbor) ProtoMessage()    {}
func (*OspfShNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_776818f2d81e6901, []int{2}
}

func (m *OspfShNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShNeighbor.Unmarshal(m, b)
}
func (m *OspfShNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShNeighbor.Marshal(b, m, deterministic)
}
func (m *OspfShNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShNeighbor.Merge(m, src)
}
func (m *OspfShNeighbor) XXX_Size() int {
	return xxx_messageInfo_OspfShNeighbor.Size(m)
}
func (m *OspfShNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShNeighbor proto.InternalMessageInfo

func (m *OspfShNeighbor) GetNeighborId() string {
	if m != nil {
		return m.NeighborId
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborAddressXr() string {
	if m != nil {
		return m.NeighborAddressXr
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborInterfaceName() string {
	if m != nil {
		return m.NeighborInterfaceName
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDrPriority() uint32 {
	if m != nil {
		return m.NeighborDrPriority
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborState() string {
	if m != nil {
		return m.NeighborState
	}
	return ""
}

func (m *OspfShNeighbor) GetDrBdrState() string {
	if m != nil {
		return m.DrBdrState
	}
	return ""
}

func (m *OspfShNeighbor) GetNeighborDeadTimer() uint32 {
	if m != nil {
		return m.NeighborDeadTimer
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborUpTime() uint32 {
	if m != nil {
		return m.NeighborUpTime
	}
	return 0
}

func (m *OspfShNeighbor) GetNeighborMadjInterface() bool {
	if m != nil {
		return m.NeighborMadjInterface
	}
	return false
}

func (m *OspfShNeighbor) GetNeighborBfdInformation() *OspfShNeighborBfd {
	if m != nil {
		return m.NeighborBfdInformation
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShNeighbor_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor_KEYS")
	proto.RegisterType((*OspfShNeighborBfd)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor_bfd")
	proto.RegisterType((*OspfShNeighbor)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.adjacency_information.neighbors.neighbor.ospf_sh_neighbor")
}

func init() { proto.RegisterFile("ospf_sh_neighbor.proto", fileDescriptor_776818f2d81e6901) }

var fileDescriptor_776818f2d81e6901 = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0x02, 0x42, 0xe0, 0x34, 0xb4, 0x35, 0x69, 0xf1, 0x8d, 0x10, 0x09, 0x14, 0x2e, 0x0b,
	0x6a, 0x4b, 0xf9, 0x3b, 0x51, 0x51, 0xa4, 0x0a, 0x15, 0xa1, 0x14, 0x24, 0x38, 0x8d, 0xbc, 0x3b,
	0xe3, 0xd4, 0x51, 0x76, 0xbd, 0xb2, 0x9d, 0xaa, 0x39, 0xf0, 0x06, 0xbc, 0x0f, 0xef, 0xc3, 0x93,
	0xa0, 0xf5, 0x66, 0x9d, 0x1f, 0x38, 0x73, 0xf3, 0x7e, 0xf3, 0xcd, 0xcf, 0xf7, 0xcd, 0x2c, 0xdb,
	0x37, 0xae, 0x52, 0xe0, 0x2e, 0xa1, 0x24, 0x3d, 0xbe, 0xcc, 0x8c, 0x4d, 0x2b, 0x6b, 0xbc, 0xe1,
	0xe3, 0x5c, 0xbb, 0xdc, 0x80, 0x36, 0x0e, 0xae, 0x2d, 0xe8, 0xea, 0xea, 0x08, 0x02, 0xd3, 0x54,
	0x64, 0xd3, 0xfa, 0x55, 0xf3, 0x72, 0x72, 0x8e, 0x5c, 0xfb, 0x4a, 0x91, 0x94, 0x9c, 0x4d, 0x3d,
	0x5c, 0x59, 0x95, 0x4a, 0x9c, 0xc8, 0x9c, 0xca, 0x7c, 0x0e, 0xba, 0x54, 0xc6, 0x16, 0xd2, 0x6b,
	0x53, 0xa6, 0x6d, 0x13, 0x17, 0x5f, 0x83, 0x9f, 0x09, 0xdb, 0xdb, 0x9c, 0x01, 0x3e, 0x9e, 0x7e,
	0xbf, 0xe0, 0x8f, 0xd8, 0xd6, 0xa2, 0x32, 0x94, 0xb2, 0x20, 0x91, 0xf4, 0x93, 0xe1, 0xdd, 0x51,
	0x67, 0x81, 0x7d, 0x92, 0x05, 0xf1, 0xc7, 0xec, 0x9e, 0x2e, 0x3d, 0x59, 0x25, 0x73, 0x6a, 0x48,
	0x37, 0x02, 0xa9, 0x1b, 0xd1, 0x40, 0x7b, 0xca, 0x76, 0x62, 0x69, 0x89, 0x68, 0xc9, 0x39, 0x71,
	0x33, 0x10, 0xb7, 0x5b, 0xfc, 0x5d, 0x03, 0x0f, 0x0c, 0xeb, 0xfd, 0x35, 0x4d, 0xa6, 0x90, 0x3f,
	0x63, 0xbd, 0x4c, 0x21, 0xe8, 0xd2, 0x2b, 0xa0, 0x52, 0x66, 0x53, 0x82, 0xc2, 0x60, 0x33, 0x54,
	0x77, 0xb4, 0x9b, 0x29, 0x3c, 0x2b, 0xbd, 0x3a, 0x0d, 0x91, 0x73, 0x83, 0xc4, 0x9f, 0xb0, 0xed,
	0x3a, 0xc1, 0x79, 0xe9, 0x67, 0x0e, 0xd4, 0x54, 0x8e, 0xc3, 0x6c, 0xdd, 0x51, 0x37, 0x53, 0x78,
	0x11, 0xd0, 0x0f, 0x53, 0x39, 0x1e, 0xfc, 0xbe, 0xc5, 0x76, 0x36, 0x3b, 0xf2, 0x87, 0xac, 0x13,
	0xbb, 0x6b, 0x14, 0x07, 0x61, 0x56, 0xd6, 0x42, 0x67, 0xc8, 0x53, 0x76, 0x7f, 0x53, 0x11, 0x5c,
	0x5b, 0x71, 0x18, 0x88, 0xbb, 0x1b, 0xa2, 0xbe, 0x59, 0x7e, 0xcc, 0x1e, 0x2c, 0x0b, 0xae, 0x3b,
	0x76, 0x14, 0x72, 0xf6, 0x62, 0xf1, 0x35, 0xe7, 0x9e, 0xb3, 0x5e, 0xcc, 0x43, 0x0b, 0x95, 0xd5,
	0xc6, 0x6a, 0x3f, 0x17, 0x2f, 0x82, 0x14, 0xde, 0xc6, 0xde, 0xdb, 0xcf, 0x8b, 0x48, 0xbd, 0x92,
	0x98, 0x51, 0x8b, 0x27, 0x71, 0xdc, 0xac, 0xa4, 0x45, 0x6b, 0xed, 0xc4, 0xfb, 0x6c, 0x0b, 0x2d,
	0x64, 0xd8, 0x92, 0x5e, 0x36, 0x12, 0xd1, 0x9e, 0xe0, 0x82, 0xb1, 0x2a, 0x11, 0x49, 0x22, 0x78,
	0x5d, 0x90, 0x15, 0xaf, 0x1a, 0xc3, 0x63, 0x67, 0x92, 0xf8, 0xa5, 0x0e, 0xf0, 0xe1, 0xca, 0x92,
	0x67, 0x55, 0x60, 0x8b, 0xd7, 0x81, 0x1c, 0x07, 0xfa, 0x5a, 0xd5, 0xd4, 0x35, 0x33, 0x0a, 0x89,
	0x93, 0xa5, 0x23, 0xe2, 0x4d, 0x3f, 0x19, 0xde, 0x59, 0x9a, 0x71, 0x2e, 0x71, 0x12, 0x0d, 0xe1,
	0xbf, 0x12, 0x26, 0x56, 0x8f, 0x62, 0xf5, 0xb4, 0xc5, 0xdb, 0x7e, 0x32, 0xec, 0x1c, 0xfc, 0x48,
	0xff, 0xd3, 0x7f, 0x93, 0xfe, 0xeb, 0x4a, 0x47, 0xfb, 0xed, 0xd7, 0x49, 0x7d, 0x8f, 0x31, 0x3f,
	0xbb, 0x1d, 0x7e, 0xea, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x41, 0x12, 0xcb, 0xee,
	0x03, 0x00, 0x00,
}
