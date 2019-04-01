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
// source: ospf_sh_ipfrr_topo.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_fast_reroute_topologies_topology

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

type OspfShIpfrrTopo_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	RouterId             string   `protobuf:"bytes,3,opt,name=router_id,json=routerId,proto3" json:"router_id,omitempty"`
	AreaId               uint32   `protobuf:"varint,4,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIpfrrTopo_KEYS) Reset()         { *m = OspfShIpfrrTopo_KEYS{} }
func (m *OspfShIpfrrTopo_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo_KEYS) ProtoMessage()    {}
func (*OspfShIpfrrTopo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b517c7b924eb010c, []int{0}
}

func (m *OspfShIpfrrTopo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopo_KEYS.Merge(m, src)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopo_KEYS.Size(m)
}
func (m *OspfShIpfrrTopo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopo_KEYS proto.InternalMessageInfo

func (m *OspfShIpfrrTopo_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetRouterId() string {
	if m != nil {
		return m.RouterId
	}
	return ""
}

func (m *OspfShIpfrrTopo_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

type OspfShIpfrrTopoEntry struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Distance             uint32   `protobuf:"varint,2,opt,name=distance,proto3" json:"distance,omitempty"`
	DistanceReverse      uint32   `protobuf:"varint,3,opt,name=distance_reverse,json=distanceReverse,proto3" json:"distance_reverse,omitempty"`
	Type4                bool     `protobuf:"varint,4,opt,name=type4,proto3" json:"type4,omitempty"`
	Revision             uint32   `protobuf:"varint,5,opt,name=revision,proto3" json:"revision,omitempty"`
	NeighborSourced      bool     `protobuf:"varint,6,opt,name=neighbor_sourced,json=neighborSourced,proto3" json:"neighbor_sourced,omitempty"`
	Dr                   bool     `protobuf:"varint,7,opt,name=dr,proto3" json:"dr,omitempty"`
	Poison               bool     `protobuf:"varint,8,opt,name=poison,proto3" json:"poison,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIpfrrTopoEntry) Reset()         { *m = OspfShIpfrrTopoEntry{} }
func (m *OspfShIpfrrTopoEntry) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopoEntry) ProtoMessage()    {}
func (*OspfShIpfrrTopoEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_b517c7b924eb010c, []int{1}
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

func (m *OspfShIpfrrTopoEntry) GetDistance() uint32 {
	if m != nil {
		return m.Distance
	}
	return 0
}

func (m *OspfShIpfrrTopoEntry) GetDistanceReverse() uint32 {
	if m != nil {
		return m.DistanceReverse
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

func (m *OspfShIpfrrTopoEntry) GetPoison() bool {
	if m != nil {
		return m.Poison
	}
	return false
}

type OspfShIpfrrTopo struct {
	IpfrrTopoAreaId      string                  `protobuf:"bytes,50,opt,name=ipfrr_topo_area_id,json=ipfrrTopoAreaId,proto3" json:"ipfrr_topo_area_id,omitempty"`
	IpfrrRouterId        string                  `protobuf:"bytes,51,opt,name=ipfrr_router_id,json=ipfrrRouterId,proto3" json:"ipfrr_router_id,omitempty"`
	IpfrrAreaRevision    uint32                  `protobuf:"varint,52,opt,name=ipfrr_area_revision,json=ipfrrAreaRevision,proto3" json:"ipfrr_area_revision,omitempty"`
	IpfrrTopo            []*OspfShIpfrrTopoEntry `protobuf:"bytes,53,rep,name=ipfrr_topo,json=ipfrrTopo,proto3" json:"ipfrr_topo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *OspfShIpfrrTopo) Reset()         { *m = OspfShIpfrrTopo{} }
func (m *OspfShIpfrrTopo) String() string { return proto.CompactTextString(m) }
func (*OspfShIpfrrTopo) ProtoMessage()    {}
func (*OspfShIpfrrTopo) Descriptor() ([]byte, []int) {
	return fileDescriptor_b517c7b924eb010c, []int{2}
}

func (m *OspfShIpfrrTopo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIpfrrTopo.Unmarshal(m, b)
}
func (m *OspfShIpfrrTopo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIpfrrTopo.Marshal(b, m, deterministic)
}
func (m *OspfShIpfrrTopo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIpfrrTopo.Merge(m, src)
}
func (m *OspfShIpfrrTopo) XXX_Size() int {
	return xxx_messageInfo_OspfShIpfrrTopo.Size(m)
}
func (m *OspfShIpfrrTopo) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIpfrrTopo.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIpfrrTopo proto.InternalMessageInfo

func (m *OspfShIpfrrTopo) GetIpfrrTopoAreaId() string {
	if m != nil {
		return m.IpfrrTopoAreaId
	}
	return ""
}

func (m *OspfShIpfrrTopo) GetIpfrrRouterId() string {
	if m != nil {
		return m.IpfrrRouterId
	}
	return ""
}

func (m *OspfShIpfrrTopo) GetIpfrrAreaRevision() uint32 {
	if m != nil {
		return m.IpfrrAreaRevision
	}
	return 0
}

func (m *OspfShIpfrrTopo) GetIpfrrTopo() []*OspfShIpfrrTopoEntry {
	if m != nil {
		return m.IpfrrTopo
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShIpfrrTopo_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_KEYS")
	proto.RegisterType((*OspfShIpfrrTopoEntry)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo_entry")
	proto.RegisterType((*OspfShIpfrrTopo)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.fast_reroute.topologies.topology.ospf_sh_ipfrr_topo")
}

func init() { proto.RegisterFile("ospf_sh_ipfrr_topo.proto", fileDescriptor_b517c7b924eb010c) }

var fileDescriptor_b517c7b924eb010c = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x55, 0xb6, 0x34, 0x9b, 0x75, 0x09, 0x0b, 0x06, 0xd1, 0x00, 0x97, 0x65, 0x0f, 0x68, 0x11,
	0x52, 0x0e, 0xed, 0xf2, 0x01, 0x1c, 0x38, 0xac, 0x90, 0x38, 0xb8, 0x5c, 0x90, 0x90, 0xac, 0x34,
	0x9e, 0xb4, 0x96, 0x68, 0xc6, 0x1a, 0xa7, 0x11, 0xfb, 0x11, 0x1c, 0xf9, 0x10, 0xbe, 0x10, 0xe4,
	0xf1, 0x7a, 0x8b, 0xb4, 0x70, 0xe4, 0x12, 0xcd, 0xbc, 0x37, 0xf3, 0x3c, 0xf3, 0xec, 0x88, 0x0a,
	0xbd, 0xeb, 0xb4, 0xbf, 0xd6, 0xd6, 0x75, 0x44, 0x7a, 0x40, 0x87, 0xb5, 0x23, 0x1c, 0x50, 0x7e,
	0x69, 0xad, 0x6f, 0x51, 0x5b, 0xf4, 0xfa, 0x1b, 0x69, 0xeb, 0xc6, 0xb5, 0xe6, 0x5a, 0x74, 0x40,
	0x75, 0x88, 0x42, 0x5d, 0x0b, 0xde, 0x83, 0x4f, 0x51, 0x3d, 0x52, 0xc7, 0x9f, 0xba, 0x6b, 0xfc,
	0xa0, 0x09, 0x08, 0x6f, 0x07, 0xa8, 0x83, 0xe8, 0x57, 0xbc, 0xb2, 0xe0, 0x53, 0xb8, 0x5d, 0x7e,
	0xcf, 0xc4, 0xe9, 0xe1, 0xd1, 0xfa, 0xc3, 0xfb, 0xcf, 0x17, 0xf2, 0xa5, 0xb8, 0xbf, 0x13, 0xd4,
	0x7d, 0x73, 0x03, 0x55, 0xb6, 0xc8, 0x56, 0x33, 0x75, 0xb2, 0xc3, 0x3e, 0x36, 0x37, 0x20, 0x9f,
	0x89, 0x62, 0xa4, 0x2e, 0xd2, 0x13, 0xa6, 0xa7, 0x23, 0x75, 0x4c, 0xbd, 0x10, 0x33, 0x3e, 0x96,
	0xb4, 0x35, 0xd5, 0x11, 0x73, 0x45, 0x04, 0x36, 0x46, 0x9e, 0x8a, 0x69, 0x43, 0xd0, 0x04, 0xea,
	0xde, 0x22, 0x5b, 0x95, 0x2a, 0x0f, 0xe9, 0xc6, 0x2c, 0x7f, 0x65, 0x7f, 0xb3, 0x42, 0x43, 0x3f,
	0xd0, 0x36, 0x74, 0xf5, 0x68, 0x20, 0x74, 0xc5, 0x59, 0xf2, 0x90, 0x6e, 0x8c, 0x7c, 0x2e, 0x0a,
	0x63, 0xfd, 0xd0, 0xf4, 0x6d, 0x1c, 0xa3, 0x54, 0xfb, 0x5c, 0xbe, 0x16, 0x0f, 0x53, 0xac, 0x09,
	0x46, 0x20, 0x0f, 0x3c, 0x4e, 0xa9, 0xe6, 0x09, 0x57, 0x11, 0x96, 0x4f, 0xc4, 0xf1, 0xb0, 0x75,
	0xb0, 0xe6, 0x99, 0x0a, 0x15, 0x93, 0x20, 0x4e, 0x30, 0x5a, 0x6f, 0xb1, 0xaf, 0x8e, 0xa3, 0x78,
	0xca, 0x83, 0x78, 0x0f, 0xf6, 0xea, 0xfa, 0x12, 0x49, 0x7b, 0xbc, 0xa5, 0x16, 0x4c, 0x95, 0x73,
	0xf3, 0x3c, 0xe1, 0x17, 0x11, 0x96, 0x0f, 0xc4, 0xc4, 0x50, 0x35, 0x65, 0x72, 0x62, 0x48, 0x3e,
	0x15, 0xb9, 0x43, 0xeb, 0xb1, 0xaf, 0x0a, 0xc6, 0x76, 0xd9, 0xf2, 0xe7, 0x44, 0xc8, 0x43, 0x07,
	0xe4, 0x1b, 0x21, 0xff, 0xf0, 0x23, 0x99, 0x77, 0xc6, 0x36, 0xcc, 0x99, 0xf9, 0x84, 0x0e, 0xdf,
	0xb1, 0x8b, 0xf2, 0x95, 0x88, 0x90, 0xbe, 0xbb, 0x81, 0x73, 0xae, 0x2c, 0x19, 0x56, 0xe9, 0x1a,
	0x6a, 0xf1, 0x38, 0xd6, 0xb1, 0xde, 0x7e, 0xcb, 0x35, 0x6f, 0xf9, 0x88, 0xa9, 0xa0, 0xa8, 0xd2,
	0xba, 0x3f, 0x32, 0x21, 0xee, 0xa6, 0xa8, 0xde, 0x2e, 0x8e, 0x56, 0x27, 0x67, 0x63, 0xfd, 0x3f,
	0x5f, 0x68, 0xfd, 0xaf, 0xd7, 0xa0, 0x66, 0xfb, 0xad, 0x2f, 0x73, 0xfe, 0x55, 0xce, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x5d, 0xc2, 0xcc, 0x46, 0x03, 0x00, 0x00,
}