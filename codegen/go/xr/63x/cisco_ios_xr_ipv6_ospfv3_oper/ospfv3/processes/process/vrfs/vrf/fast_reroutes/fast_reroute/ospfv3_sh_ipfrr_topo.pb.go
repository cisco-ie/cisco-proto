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
// source: ospfv3_sh_ipfrr_topo.proto

package cisco_ios_xr_ipv6_ospfv3_oper_ospfv3_processes_process_vrfs_vrf_fast_reroutes_fast_reroute

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

type Ospfv3ShIpfrrTopo_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RouterId             string   `protobuf:"bytes,3,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	AreaId               uint32   `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ospfv3ShIpfrrTopo_KEYS) Reset()         { *m = Ospfv3ShIpfrrTopo_KEYS{} }
func (m *Ospfv3ShIpfrrTopo_KEYS) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShIpfrrTopo_KEYS) ProtoMessage()    {}
func (*Ospfv3ShIpfrrTopo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2837dbf1c4c39301, []int{0}
}

func (m *Ospfv3ShIpfrrTopo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS.Unmarshal(m, b)
}
func (m *Ospfv3ShIpfrrTopo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShIpfrrTopo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS.Merge(m, src)
}
func (m *Ospfv3ShIpfrrTopo_KEYS) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS.Size(m)
}
func (m *Ospfv3ShIpfrrTopo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShIpfrrTopo_KEYS proto.InternalMessageInfo

func (m *Ospfv3ShIpfrrTopo_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *Ospfv3ShIpfrrTopo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *Ospfv3ShIpfrrTopo_KEYS) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *Ospfv3ShIpfrrTopo_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type OspfShIpfrrTopoEntry struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Lsaid                uint32   `protobuf:"varint,2,opt,name=lsaid,proto3" json:"lsaid,omitempty"`
	Distance             uint32   `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	Type4                bool     `protobuf:"varint,4,opt,name=type4,proto3" json:"type4,omitempty"`
	Revision             uint32   `protobuf:"varint,5,opt,name=revision,proto3" json:"revision,omitempty"`
	NeighborSourced      bool     `protobuf:"varint,6,opt,name=neighbor_sourced,json=neighborSourced,proto3" json:"neighbor_sourced,omitempty"`
	Dr                   bool     `protobuf:"varint,7,opt,name=dr,proto3" json:"dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIpfrrTopoEntry) Reset()         { *m = OspfShIpfrrTopoEntry{} }
func (m *OspfShIpfrrTopoEntry) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopoEntry) ProtoMessage()    {}
func (*OspfShIpfrrTopoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2837dbf1c4c39301, []int{1}
}

func (m *OspfShIpfrrTopoEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopoEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Marshal(b, m, deterministic)
}
func (m *OspfShIpfrrTopoEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopoEntry.Merge(m, src)
}
func (m *OspfShIpfrrTopoEntry) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopoEntry.Size(m)
}
func (m *OspfShIpfrrTopoEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopoEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopoEntry proto.InternalMessageInfo

func (m *OspfShIpfrrTopoEntry) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *OspfShIpfrrTopoEntry) GetLsaid() uint32 {
	if m != nil {
		return m.Lsaid
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetType4() bool {
	if m != nil {
		return m.Type4
	}
	return false
}

func (m *OspfShIpfrrTopoEntry) GetRevision() uint32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetNeighborSourced() bool {
	if m != nil {
		return m.NeighborSourced
	}
	return false
}

func (m *OspfShIpfrrTopoEntry) GetDr() bool {
	if m != nil {
		return m.Dr
	}
	return false
}

type Ospfv3ShIpfrrTopo struct {
	IpfrrTopoAreaId      string                  `protobuf:"bytes,50,opt,name=ipfrr_topo_area_id,json=ipfrrTopoAreaId,proto3" json:"ipfrr_topo_area_id,omitempty"`
	IpfrrRouterId        string                  `protobuf:"bytes,51,opt,name=ipfrr_router_id,json=ipfrrRouterId,proto3" json:"ipfrr_router_id,omitempty"`
	IpfrrAreaRevision    uint32                  `protobuf:"varint,52,opt,name=ipfrr_area_revision,json=ipfrrAreaRevision,proto3" json:"ipfrr_area_revision,omitempty"`
	IpfrrTopo            []*OspfShIpfrrTopoEntry `protobuf:"bytes,53,rep,name=ipfrr_topo,json=ipfrrTopo,proto3" json:"ipfrr_topo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Ospfv3ShIpfrrTopo) Reset()         { *m = Ospfv3ShIpfrrTopo{} }
func (m *Ospfv3ShIpfrrTopo) String() string { return proto.CompactTextString(m) }
func (*Ospfv3ShIpfrrTopo) ProtoMessage()    {}
func (*Ospfv3ShIpfrrTopo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2837dbf1c4c39301, []int{2}
}

func (m *Ospfv3ShIpfrrTopo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo.Unmarshal(m, b)
}
func (m *Ospfv3ShIpfrrTopo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo.Marshal(b, m, deterministic)
}
func (m *Ospfv3ShIpfrrTopo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ospfv3ShIpfrrTopo.Merge(m, src)
}
func (m *Ospfv3ShIpfrrTopo) XXX_Size() int {
	return xxx_messageInfo_Ospfv3ShIpfrrTopo.Size(m)
}
func (m *Ospfv3ShIpfrrTopo) XXX_DiscardUnknown() {
	xxx_messageInfo_Ospfv3ShIpfrrTopo.DiscardUnknown(m)
}

var xxx_messageInfo_Ospfv3ShIpfrrTopo proto.InternalMessageInfo

func (m *Ospfv3ShIpfrrTopo) GetIpfrrTopoAreaId() string {
	if m != nil {
		return m.IpfrrTopoAreaId
	}
	return ""
}

func (m *Ospfv3ShIpfrrTopo) GetIpfrrRouterId() string {
	if m != nil {
		return m.IpfrrRouterId
	}
	return ""
}

func (m *Ospfv3ShIpfrrTopo) GetIpfrrAreaRevision() uint32 {
	if m != nil {
		return m.IpfrrAreaRevision
	}
	return 0
}

func (m *Ospfv3ShIpfrrTopo) GetIpfrrTopo() []*OspfShIpfrrTopoEntry {
	if m != nil {
		return m.IpfrrTopo
	}
	return nil
}

func init() {
	proto.RegisterType((*Ospfv3ShIpfrrTopo_KEYS)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.fast_reroutes.fast_reroute.ospfv3_sh_ipfrr_topo_KEYS")
	proto.RegisterType((*OspfShIpfrrTopoEntry)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.fast_reroutes.fast_reroute.ospf_sh_ipfrr_topo_entry")
	proto.RegisterType((*Ospfv3ShIpfrrTopo)(nil), "cisco_ios_xr_ipv6_ospfv3_oper.ospfv3.processes.process.vrfs.vrf.fast_reroutes.fast_reroute.ospfv3_sh_ipfrr_topo")
}

func init() { proto.RegisterFile("ospfv3_sh_ipfrr_topo.proto", fileDescriptor_2837dbf1c4c39301) }

var fileDescriptor_2837dbf1c4c39301 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xb2, 0x6c, 0x9a, 0xce, 0x12, 0x16, 0xcc, 0x4a, 0x64, 0x97, 0x4b, 0xe9, 0x01, 0x15,
	0x21, 0xe5, 0xb0, 0x5d, 0xb8, 0x73, 0xe0, 0x50, 0x21, 0x71, 0x70, 0xb9, 0xc0, 0xc5, 0x4a, 0x63,
	0x87, 0x5a, 0xa2, 0xb1, 0x35, 0x76, 0x23, 0xfa, 0x15, 0x1c, 0xf8, 0x12, 0x7e, 0x85, 0x2f, 0x42,
	0x1e, 0x37, 0x2d, 0xa0, 0x72, 0xe4, 0x12, 0xf9, 0xcd, 0x9b, 0xe7, 0x79, 0x6f, 0x62, 0xb8, 0x31,
	0xce, 0xb6, 0xfd, 0x5c, 0xb8, 0xb5, 0xd0, 0xb6, 0x45, 0x14, 0xde, 0x58, 0x53, 0x59, 0x34, 0xde,
	0xb0, 0x4f, 0x8d, 0x76, 0x8d, 0x11, 0xda, 0x38, 0xf1, 0x15, 0x85, 0xb6, 0xfd, 0x6b, 0xb1, 0xef,
	0x36, 0x56, 0x61, 0x15, 0xcf, 0xa1, 0xb7, 0x51, 0xce, 0x29, 0x37, 0x9c, 0xaa, 0x1e, 0x5b, 0xfa,
	0x54, 0x6d, 0xed, 0xbc, 0x40, 0x85, 0x66, 0xeb, 0x95, 0xfb, 0x03, 0x4d, 0xbf, 0x25, 0x70, 0x7d,
	0x6a, 0xb4, 0x78, 0xf7, 0xf6, 0xe3, 0x92, 0x3d, 0x83, 0xfb, 0xfb, 0xcb, 0x44, 0x57, 0x6f, 0x54,
	0x99, 0x4c, 0x92, 0xd9, 0x98, 0x5f, 0xec, 0x6b, 0xef, 0xeb, 0x8d, 0x62, 0xd7, 0x90, 0xf7, 0xd8,
	0x46, 0x3a, 0x25, 0x7a, 0xd4, 0x63, 0x4b, 0xd4, 0x53, 0x18, 0xd3, 0x10, 0x14, 0x5a, 0x96, 0x67,
	0xc4, 0xe5, 0xb1, 0xb0, 0x90, 0xec, 0x09, 0x8c, 0x6a, 0x54, 0x75, 0xa0, 0xee, 0x4d, 0x92, 0x59,
	0xc1, 0xb3, 0x00, 0x17, 0x72, 0xfa, 0x33, 0x81, 0x32, 0x38, 0xfa, 0xcb, 0x8f, 0xea, 0x3c, 0xee,
	0x82, 0xaa, 0x33, 0x52, 0x05, 0x55, 0xf4, 0x92, 0x05, 0xb8, 0x90, 0xec, 0x0a, 0xce, 0xbf, 0xb8,
	0x5a, 0x4b, 0xf2, 0x50, 0xf0, 0x08, 0xd8, 0x0d, 0xe4, 0x52, 0x3b, 0x5f, 0x77, 0x8d, 0x22, 0x03,
	0x05, 0x3f, 0xe0, 0xa0, 0xf0, 0x3b, 0xab, 0xee, 0x68, 0x7c, 0xce, 0x23, 0x08, 0x0a, 0x54, 0xbd,
	0x76, 0xda, 0x74, 0xe5, 0x79, 0x54, 0x0c, 0x98, 0xbd, 0x80, 0x87, 0x9d, 0xd2, 0x9f, 0xd7, 0x2b,
	0x83, 0xc2, 0x99, 0x2d, 0x36, 0x4a, 0x96, 0x19, 0x89, 0x2f, 0x87, 0xfa, 0x32, 0x96, 0xd9, 0x03,
	0x48, 0x25, 0x96, 0x23, 0x22, 0x53, 0x89, 0xd3, 0x1f, 0x29, 0x5c, 0x9d, 0x5a, 0x33, 0x7b, 0x09,
	0xec, 0xb7, 0x90, 0xc3, 0x46, 0x6e, 0x29, 0xdb, 0x25, 0x31, 0x1f, 0x8c, 0x35, 0x6f, 0x68, 0x35,
	0xec, 0x39, 0xc4, 0x92, 0x38, 0xae, 0x75, 0x4e, 0x9d, 0x05, 0x95, 0xf9, 0xb0, 0xdb, 0x0a, 0x1e,
	0xc7, 0x3e, 0xba, 0xef, 0x90, 0xe7, 0x8e, 0xf2, 0x3c, 0x22, 0x2a, 0xdc, 0xc8, 0x87, 0x60, 0xdf,
	0x13, 0x80, 0xa3, 0x8b, 0xf2, 0xd5, 0xe4, 0x6c, 0x76, 0x71, 0xeb, 0xab, 0xff, 0xf7, 0xec, 0xaa,
	0x7f, 0xfd, 0x60, 0x3e, 0x3e, 0x64, 0x5e, 0x65, 0xf4, 0xfa, 0xe7, 0xbf, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xb1, 0x02, 0x2e, 0x0f, 0x1b, 0x03, 0x00, 0x00,
}
