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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_interface_information_interface_briefs_interface_brief

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *OspfShIfBrief_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*OspfShIfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS")
	proto.RegisterType((*OspfShInterfaceMadj)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_interface_madj")
	proto.RegisterType((*OspfShIfBrief)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.interface_information.interface_briefs.interface_brief.ospf_sh_if_brief")
}

func init() { proto.RegisterFile("ospf_sh_if_brief.proto", fileDescriptor_d0466f349213f248) }

var fileDescriptor_d0466f349213f248 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xdd, 0x4e, 0x14, 0x31,
	0x14, 0xce, 0x80, 0x20, 0x16, 0x56, 0xa0, 0x22, 0x29, 0xd1, 0xc4, 0x75, 0x13, 0x93, 0x55, 0x93,
	0x09, 0x01, 0x54, 0xfc, 0x8b, 0x31, 0xac, 0x46, 0xa2, 0x78, 0xb1, 0xdc, 0xe8, 0x55, 0xd3, 0x9d,
	0x76, 0xdc, 0xce, 0xee, 0x4c, 0x37, 0x3d, 0x75, 0xe1, 0xc2, 0x27, 0xf0, 0x71, 0x7c, 0x00, 0x1f,
	0xc9, 0x67, 0x30, 0x3d, 0x0c, 0xdb, 0x99, 0xcd, 0xa2, 0x31, 0x26, 0xdc, 0x4c, 0x26, 0xe7, 0xfb,
	0xfa, 0x9d, 0xaf, 0x5f, 0x4f, 0x4b, 0x36, 0x0d, 0x8c, 0x52, 0x0e, 0x7d, 0xae, 0x53, 0xde, 0xb3,
	0x5a, 0xa5, 0xf1, 0xc8, 0x1a, 0x67, 0xa8, 0x4d, 0x34, 0x24, 0x86, 0x6b, 0x03, 0xfc, 0xd4, 0x72,
	0x3d, 0x1a, 0xef, 0x71, 0x64, 0x9a, 0x91, 0xb2, 0xb1, 0xff, 0xf3, 0xbc, 0x44, 0x01, 0x28, 0x38,
	0xff, 0x8b, 0xc7, 0x36, 0xc5, 0x4f, 0xac, 0x0b, 0xa7, 0x6c, 0x2a, 0x12, 0xc5, 0x75, 0x91, 0x1a,
	0x9b, 0x0b, 0xa7, 0x4d, 0x51, 0xa9, 0x62, 0x23, 0x98, 0x2e, 0xb4, 0xbe, 0x91, 0x9b, 0xd3, 0x6e,
	0xf8, 0xfb, 0x37, 0x9f, 0x8f, 0xe9, 0x5d, 0xb2, 0x52, 0xf6, 0xe0, 0x85, 0xc8, 0x15, 0x8b, 0x9a,
	0x51, 0xfb, 0x5a, 0x77, 0xb9, 0xac, 0x7d, 0x14, 0xb9, 0xa2, 0x5b, 0x64, 0x69, 0x6c, 0xd3, 0x33,
	0x78, 0x0e, 0xe1, 0xab, 0x63, 0x9b, 0x22, 0x74, 0x8f, 0x5c, 0x0f, 0x9d, 0x90, 0x30, 0x8f, 0x84,
	0xc6, 0xa4, 0xea, 0x69, 0xad, 0x9f, 0x73, 0x95, 0x30, 0x26, 0xfc, 0x5c, 0xc8, 0xac, 0xae, 0x20,
	0xac, 0x12, 0xa5, 0x83, 0xa0, 0xf0, 0xda, 0x2a, 0x41, 0x9b, 0x64, 0xc5, 0xd3, 0x91, 0xc1, 0xb5,
	0x44, 0x1f, 0x8d, 0x2e, 0xf1, 0x35, 0x8f, 0x1f, 0x4a, 0xba, 0x4f, 0x58, 0xc5, 0x8a, 0xd2, 0x5f,
	0xfa, 0x3d, 0x63, 0x79, 0x62, 0xbe, 0x16, 0x0e, 0x4d, 0x35, 0xba, 0x9b, 0xc1, 0x54, 0x09, 0x1f,
	0x78, 0x94, 0xbe, 0x22, 0xb7, 0x2b, 0x16, 0x64, 0x36, 0xbd, 0xfa, 0x0a, 0xae, 0xde, 0x0a, 0x86,
	0x64, 0x56, 0x17, 0x88, 0xc9, 0x8d, 0x20, 0x30, 0xd4, 0xc5, 0x80, 0x27, 0x06, 0x1c, 0x5b, 0xc0,
	0x75, 0xeb, 0x13, 0xe8, 0x83, 0x2e, 0x06, 0x07, 0x06, 0x1c, 0xdd, 0x26, 0x1b, 0x98, 0x46, 0x58,
	0x04, 0x4e, 0x38, 0xc5, 0x16, 0x71, 0xe7, 0xd4, 0x63, 0x87, 0xe7, 0xd0, 0xb1, 0x47, 0x5a, 0xbf,
	0x16, 0xc8, 0xda, 0xf4, 0xf9, 0xd1, 0x07, 0x64, 0xbd, 0x1e, 0x3e, 0x3f, 0xb5, 0x6c, 0x07, 0x35,
	0x56, 0x6b, 0xf9, 0x7f, 0xb2, 0x33, 0x62, 0xde, 0x9d, 0x15, 0xf3, 0xc3, 0xaa, 0xa4, 0x90, 0xd2,
	0x2a, 0x00, 0xb6, 0x87, 0xcc, 0xb5, 0xca, 0xfe, 0xb1, 0x5e, 0xd7, 0xcc, 0x05, 0x0c, 0xd8, 0x23,
	0xdc, 0x71, 0xd0, 0x3c, 0x12, 0x30, 0xb8, 0x28, 0x9d, 0xc7, 0xff, 0x9a, 0xce, 0x93, 0x8b, 0xd2,
	0xa1, 0x1d, 0x72, 0x27, 0x90, 0x53, 0x01, 0x8e, 0x4b, 0xe5, 0x54, 0xe2, 0x78, 0xdf, 0x0c, 0x25,
	0x97, 0xe6, 0xa4, 0x60, 0xfb, 0xcd, 0xa8, 0xbd, 0xd4, 0xbd, 0x35, 0xa1, 0xbd, 0x15, 0xe0, 0x3a,
	0x48, 0x7a, 0x67, 0x86, 0xb2, 0x63, 0x4e, 0x8a, 0x3f, 0x0e, 0xd0, 0xd3, 0xff, 0x1a, 0xa0, 0x67,
	0x7f, 0x1b, 0xa0, 0xfb, 0x24, 0xa4, 0xab, 0x01, 0x2f, 0x06, 0x7b, 0x8e, 0x8e, 0x57, 0x2b, 0xf5,
	0x23, 0x7f, 0x5f, 0xb6, 0xc9, 0x46, 0xfd, 0x06, 0x95, 0x3d, 0x5e, 0x60, 0x0f, 0x5a, 0x89, 0x5e,
	0x66, 0x67, 0xe2, 0x3f, 0xa2, 0xfa, 0x39, 0xc9, 0x8c, 0xbd, 0x6c, 0xce, 0xb7, 0x97, 0x77, 0xbe,
	0x47, 0xf1, 0xe5, 0xbf, 0x44, 0xf1, 0xec, 0x77, 0xa0, 0x36, 0x34, 0x32, 0xeb, 0x2d, 0xe2, 0x53,
	0xb9, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x07, 0x1a, 0x43, 0x44, 0x05, 0x00, 0x00,
}
