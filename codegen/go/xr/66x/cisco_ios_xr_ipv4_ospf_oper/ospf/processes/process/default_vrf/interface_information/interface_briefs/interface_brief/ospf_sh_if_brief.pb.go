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
// source: ospf_sh_if_brief.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_interface_information_interface_briefs_interface_brief

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

type OspfShIfBrief_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIfBrief_KEYS) Reset()         { *m = OspfShIfBrief_KEYS{} }
func (m *OspfShIfBrief_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShIfBrief_KEYS) ProtoMessage()    {}
func (*OspfShIfBrief_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0466f349213f248, []int{0}
}

func (m *OspfShIfBrief_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Unmarshal(m, b)
}
func (m *OspfShIfBrief_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShIfBrief_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfBrief_KEYS.Merge(m, src)
}
func (m *OspfShIfBrief_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShIfBrief_KEYS.Size(m)
}
func (m *OspfShIfBrief_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfBrief_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfBrief_KEYS proto.InternalMessageInfo

func (m *OspfShIfBrief_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShInterfaceMadj struct {
	InterfaceArea             string   `protobuf:"bytes,1,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	MadjAreaId                uint32   `protobuf:"varint,2,opt,name=madj_area_id,json=madjAreaId,proto3" json:"madj_area_id,omitempty"`
	InterfaceNeighborCount    uint32   `protobuf:"varint,3,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	InterfaceAdjNeighborCount uint32   `protobuf:"varint,4,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	InterfaceLinkCost         uint32   `protobuf:"varint,5,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	OspfInterfaceState        string   `protobuf:"bytes,6,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *OspfShInterfaceMadj) Reset()         { *m = OspfShInterfaceMadj{} }
func (m *OspfShInterfaceMadj) String() string { return proto.CompactTextString(m) }
func (*OspfShInterfaceMadj) ProtoMessage()    {}
func (*OspfShInterfaceMadj) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0466f349213f248, []int{1}
}

func (m *OspfShInterfaceMadj) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShInterfaceMadj.Unmarshal(m, b)
}
func (m *OspfShInterfaceMadj) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShInterfaceMadj.Marshal(b, m, deterministic)
}
func (m *OspfShInterfaceMadj) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShInterfaceMadj.Merge(m, src)
}
func (m *OspfShInterfaceMadj) XXX_Size() int {
	return xxx_messageInfo_OspfShInterfaceMadj.Size(m)
}
func (m *OspfShInterfaceMadj) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShInterfaceMadj.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShInterfaceMadj proto.InternalMessageInfo

func (m *OspfShInterfaceMadj) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShInterfaceMadj) GetMadjAreaId() uint32 {
	if m != nil {
		return m.MadjAreaId
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

type OspfShIfBrief struct {
	InterfaceNameXr             string                 `protobuf:"bytes,50,opt,name=interface_name_xr,json=interfaceNameXr,proto3" json:"interface_name_xr,omitempty"`
	InterfaceArea               string                 `protobuf:"bytes,51,opt,name=interface_area,json=interfaceArea,proto3" json:"interface_area,omitempty"`
	InterfaceAddress            string                 `protobuf:"bytes,52,opt,name=interface_address,json=interfaceAddress,proto3" json:"interface_address,omitempty"`
	InterfaceMask               uint32                 `protobuf:"varint,53,opt,name=interface_mask,json=interfaceMask,proto3" json:"interface_mask,omitempty"`
	InterfaceLinkCost           uint32                 `protobuf:"varint,54,opt,name=interface_link_cost,json=interfaceLinkCost,proto3" json:"interface_link_cost,omitempty"`
	OspfInterfaceState          string                 `protobuf:"bytes,55,opt,name=ospf_interface_state,json=ospfInterfaceState,proto3" json:"ospf_interface_state,omitempty"`
	InterfaceFastDetectHoldDown bool                   `protobuf:"varint,56,opt,name=interface_fast_detect_hold_down,json=interfaceFastDetectHoldDown,proto3" json:"interface_fast_detect_hold_down,omitempty"`
	InterfaceNeighborCount      uint32                 `protobuf:"varint,57,opt,name=interface_neighbor_count,json=interfaceNeighborCount,proto3" json:"interface_neighbor_count,omitempty"`
	InterfaceAdjNeighborCount   uint32                 `protobuf:"varint,58,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount,proto3" json:"interface_adj_neighbor_count,omitempty"`
	InterfaceisMadj             bool                   `protobuf:"varint,59,opt,name=interfaceis_madj,json=interfaceisMadj,proto3" json:"interfaceis_madj,omitempty"`
	InterfaceMadjCount          uint32                 `protobuf:"varint,60,opt,name=interface_madj_count,json=interfaceMadjCount,proto3" json:"interface_madj_count,omitempty"`
	InterfaceMadj               []*OspfShInterfaceMadj `protobuf:"bytes,61,rep,name=interface_madj,json=interfaceMadj,proto3" json:"interface_madj,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}               `json:"-"`
	XXX_unrecognized            []byte                 `json:"-"`
	XXX_sizecache               int32                  `json:"-"`
}

func (m *OspfShIfBrief) Reset()         { *m = OspfShIfBrief{} }
func (m *OspfShIfBrief) String() string { return proto.CompactTextString(m) }
func (*OspfShIfBrief) ProtoMessage()    {}
func (*OspfShIfBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0466f349213f248, []int{2}
}

func (m *OspfShIfBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfBrief.Unmarshal(m, b)
}
func (m *OspfShIfBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfBrief.Marshal(b, m, deterministic)
}
func (m *OspfShIfBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfBrief.Merge(m, src)
}
func (m *OspfShIfBrief) XXX_Size() int {
	return xxx_messageInfo_OspfShIfBrief.Size(m)
}
func (m *OspfShIfBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfBrief.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfBrief proto.InternalMessageInfo

func (m *OspfShIfBrief) GetInterfaceNameXr() string {
	if m != nil {
		return m.InterfaceNameXr
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceMask() uint32 {
	if m != nil {
		return m.InterfaceMask
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShIfBrief) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceFastDetectHoldDown() bool {
	if m != nil {
		return m.InterfaceFastDetectHoldDown
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceisMadj() bool {
	if m != nil {
		return m.InterfaceisMadj
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceMadjCount() uint32 {
	if m != nil {
		return m.InterfaceMadjCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceMadj() []*OspfShInterfaceMadj {
	if m != nil {
		return m.InterfaceMadj
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShIfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS")
	proto.RegisterType((*OspfShInterfaceMadj)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_interface_madj")
	proto.RegisterType((*OspfShIfBrief)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief")
}

func init() { proto.RegisterFile("ospf_sh_if_brief.proto", fileDescriptor_d0466f349213f248) }

var fileDescriptor_d0466f349213f248 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0xd5, 0xb6, 0x5f, 0xab, 0x0f, 0xb7, 0xa5, 0xad, 0x29, 0x95, 0x11, 0x48, 0x84, 0x48, 0x48,
	0x01, 0xa4, 0x55, 0xd5, 0x16, 0x28, 0x7f, 0x42, 0xa8, 0x01, 0x51, 0x41, 0xb9, 0x48, 0x6f, 0xe0,
	0x6a, 0xe4, 0xac, 0xbd, 0xc4, 0x9b, 0xec, 0x3a, 0xf2, 0x38, 0x6d, 0x5f, 0x82, 0xc7, 0xe1, 0x96,
	0x77, 0xe2, 0x0d, 0x90, 0x87, 0x6d, 0x76, 0x37, 0x4a, 0x41, 0x88, 0x0b, 0xee, 0xac, 0x39, 0x67,
	0xce, 0xcc, 0x1c, 0x8f, 0xcd, 0xb6, 0x2d, 0x8e, 0x53, 0xc0, 0x01, 0x98, 0x14, 0xfa, 0xce, 0xe8,
	0x34, 0x1e, 0x3b, 0xeb, 0x2d, 0x9f, 0x24, 0x06, 0x13, 0x0b, 0xc6, 0x22, 0x9c, 0x3b, 0x30, 0xe3,
	0xd3, 0x7d, 0x20, 0xa6, 0x1d, 0x6b, 0x17, 0x87, 0x53, 0xe0, 0x25, 0x1a, 0x51, 0xe3, 0xc5, 0x29,
	0x56, 0x3a, 0x95, 0x93, 0x91, 0x87, 0x53, 0x97, 0xc6, 0xa6, 0xf0, 0xda, 0xa5, 0x32, 0xd1, 0x60,
	0x8a, 0xd4, 0xba, 0x5c, 0x7a, 0x63, 0x8b, 0x5a, 0x94, 0x6a, 0xe1, 0x6c, 0xa0, 0x2d, 0xd9, 0xf5,
	0xd9, 0x86, 0xe0, 0xdd, 0xeb, 0x4f, 0x27, 0xfc, 0x0e, 0x5b, 0x2d, 0xcb, 0x40, 0x21, 0x73, 0x2d,
	0xa2, 0x56, 0xd4, 0xb9, 0xd2, 0x5b, 0x29, 0x63, 0x1f, 0x64, 0xae, 0xf9, 0x5d, 0x76, 0xb5, 0x92,
	0x23, 0xd2, 0x02, 0x91, 0xd6, 0xa6, 0xd1, 0x40, 0x6b, 0x7f, 0x5b, 0xa8, 0x0d, 0x3d, 0xe5, 0xe7,
	0x52, 0x65, 0x4d, 0x05, 0xe9, 0xb4, 0x2c, 0xcb, 0x54, 0x0a, 0xaf, 0x9c, 0x96, 0xbc, 0xc5, 0x56,
	0x03, 0x9d, 0x18, 0x60, 0x14, 0x95, 0x59, 0xeb, 0xb1, 0x10, 0x0b, 0xf8, 0x91, 0xe2, 0x07, 0x4c,
	0xd4, 0x5a, 0xd1, 0xe6, 0xf3, 0xa0, 0x6f, 0x1d, 0x24, 0x76, 0x52, 0x78, 0xb1, 0x48, 0xec, 0xed,
	0xaa, 0xa9, 0x12, 0x3e, 0x0c, 0x28, 0x7f, 0xc9, 0x6e, 0xd5, 0x5a, 0x50, 0xd9, 0x6c, 0xf6, 0x7f,
	0x94, 0x7d, 0xa3, 0x6a, 0x48, 0x65, 0x4d, 0x81, 0x98, 0x5d, 0xab, 0x04, 0x46, 0xa6, 0x18, 0x42,
	0x62, 0xd1, 0x8b, 0x25, 0xca, 0xdb, 0x9c, 0x42, 0xef, 0x4d, 0x31, 0x3c, 0xb4, 0xe8, 0xf9, 0x0e,
	0xdb, 0x22, 0x37, 0xaa, 0x24, 0xf4, 0xd2, 0x6b, 0xb1, 0x4c, 0x93, 0xf3, 0x80, 0x1d, 0x5d, 0x40,
	0x27, 0x01, 0x69, 0x7f, 0x5f, 0x62, 0x1b, 0xb3, 0x97, 0xc4, 0xef, 0xb3, 0xcd, 0xa6, 0xf9, 0x70,
	0xee, 0xc4, 0x2e, 0x69, 0xac, 0x37, 0xfc, 0xff, 0xe8, 0xe6, 0xd8, 0xbc, 0x37, 0xcf, 0xe6, 0x07,
	0x75, 0x49, 0xa9, 0x94, 0xd3, 0x88, 0x62, 0x9f, 0x98, 0x1b, 0xb5, 0xf9, 0x29, 0xde, 0xd4, 0xcc,
	0x25, 0x0e, 0xc5, 0x43, 0x9a, 0xb8, 0xd2, 0x3c, 0x96, 0x38, 0xbc, 0xcc, 0x9d, 0x47, 0x7f, 0xea,
	0xce, 0xe3, 0xcb, 0xdc, 0xe1, 0x5d, 0x76, 0xbb, 0x22, 0xa7, 0x12, 0x3d, 0x28, 0xed, 0x75, 0xe2,
	0x61, 0x60, 0x47, 0x0a, 0x94, 0x3d, 0x2b, 0xc4, 0x41, 0x2b, 0xea, 0xfc, 0xdf, 0xbb, 0x39, 0xa5,
	0xbd, 0x91, 0xe8, 0xbb, 0x44, 0x7a, 0x6b, 0x47, 0xaa, 0x6b, 0xcf, 0x8a, 0x5f, 0x2e, 0xd0, 0x93,
	0xbf, 0x5a, 0xa0, 0xa7, 0xbf, 0x5b, 0xa0, 0x7b, 0xac, 0x72, 0xd7, 0x20, 0x3d, 0x0c, 0xf1, 0x8c,
	0x3a, 0x5e, 0xaf, 0xc5, 0x8f, 0xc3, 0x7b, 0xd9, 0x61, 0x5b, 0xcd, 0x17, 0x54, 0xd6, 0x78, 0x4e,
	0x35, 0x78, 0xcd, 0x7a, 0x95, 0xfd, 0x14, 0xff, 0x1a, 0x35, 0xef, 0x49, 0x65, 0xe2, 0x45, 0x6b,
	0xb1, 0xb3, 0xb2, 0xfb, 0x25, 0x8a, 0xff, 0xc9, 0x8f, 0x13, 0xcf, 0xff, 0x0a, 0x1a, 0x7b, 0xa3,
	0xb2, 0xfe, 0x32, 0xfd, 0x8a, 0x7b, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xd5, 0x7b, 0x04,
	0x2f, 0x05, 0x00, 0x00,
}
