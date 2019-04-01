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
// source: ospf_sh_if_stats.proto

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_area_statistics_interface_stats_entries_interface_stats_entry

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

type OspfShIfStats_KEYS struct {
	ProcessName          string   `protobuf:"bytes,1,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIfStats_KEYS) Reset()         { *m = OspfShIfStats_KEYS{} }
func (m *OspfShIfStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*OspfShIfStats_KEYS) ProtoMessage()    {}
func (*OspfShIfStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ceab17d939ff790, []int{0}
}

func (m *OspfShIfStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfStats_KEYS.Unmarshal(m, b)
}
func (m *OspfShIfStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfStats_KEYS.Marshal(b, m, deterministic)
}
func (m *OspfShIfStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfStats_KEYS.Merge(m, src)
}
func (m *OspfShIfStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_OspfShIfStats_KEYS.Size(m)
}
func (m *OspfShIfStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfStats_KEYS proto.InternalMessageInfo

func (m *OspfShIfStats_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIfStats_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShIfStatsEntry struct {
	MadjIntf             bool     `protobuf:"varint,1,opt,name=madj_intf,json=madjIntf,proto3" json:"madj_intf,omitempty"`
	AreaId               uint32   `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	AreaIdStr            string   `protobuf:"bytes,3,opt,name=area_id_str,json=areaIdStr,proto3" json:"area_id_str,omitempty"`
	IfStat               []uint32 `protobuf:"varint,4,rep,packed,name=if_stat,json=ifStat,proto3" json:"if_stat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OspfShIfStatsEntry) Reset()         { *m = OspfShIfStatsEntry{} }
func (m *OspfShIfStatsEntry) String() string { return proto.CompactTextString(m) }
func (*OspfShIfStatsEntry) ProtoMessage()    {}
func (*OspfShIfStatsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ceab17d939ff790, []int{1}
}

func (m *OspfShIfStatsEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfStatsEntry.Unmarshal(m, b)
}
func (m *OspfShIfStatsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfStatsEntry.Marshal(b, m, deterministic)
}
func (m *OspfShIfStatsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfStatsEntry.Merge(m, src)
}
func (m *OspfShIfStatsEntry) XXX_Size() int {
	return xxx_messageInfo_OspfShIfStatsEntry.Size(m)
}
func (m *OspfShIfStatsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfStatsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfStatsEntry proto.InternalMessageInfo

func (m *OspfShIfStatsEntry) GetMadjIntf() bool {
	if m != nil {
		return m.MadjIntf
	}
	return false
}

func (m *OspfShIfStatsEntry) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShIfStatsEntry) GetAreaIdStr() string {
	if m != nil {
		return m.AreaIdStr
	}
	return ""
}

func (m *OspfShIfStatsEntry) GetIfStat() []uint32 {
	if m != nil {
		return m.IfStat
	}
	return nil
}

type OspfShIfStats struct {
	Handle               string                `protobuf:"bytes,50,opt,name=handle,proto3" json:"handle,omitempty"`
	IfNameStr            string                `protobuf:"bytes,51,opt,name=if_name_str,json=ifNameStr,proto3" json:"if_name_str,omitempty"`
	IfEntry              []*OspfShIfStatsEntry `protobuf:"bytes,52,rep,name=if_entry,json=ifEntry,proto3" json:"if_entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *OspfShIfStats) Reset()         { *m = OspfShIfStats{} }
func (m *OspfShIfStats) String() string { return proto.CompactTextString(m) }
func (*OspfShIfStats) ProtoMessage()    {}
func (*OspfShIfStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ceab17d939ff790, []int{2}
}

func (m *OspfShIfStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OspfShIfStats.Unmarshal(m, b)
}
func (m *OspfShIfStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OspfShIfStats.Marshal(b, m, deterministic)
}
func (m *OspfShIfStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OspfShIfStats.Merge(m, src)
}
func (m *OspfShIfStats) XXX_Size() int {
	return xxx_messageInfo_OspfShIfStats.Size(m)
}
func (m *OspfShIfStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OspfShIfStats.DiscardUnknown(m)
}

var xxx_messageInfo_OspfShIfStats proto.InternalMessageInfo

func (m *OspfShIfStats) GetHandle() string {
	if m != nil {
		return m.Handle
	}
	return ""
}

func (m *OspfShIfStats) GetIfNameStr() string {
	if m != nil {
		return m.IfNameStr
	}
	return ""
}

func (m *OspfShIfStats) GetIfEntry() []*OspfShIfStatsEntry {
	if m != nil {
		return m.IfEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*OspfShIfStats_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.area_statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats_KEYS")
	proto.RegisterType((*OspfShIfStatsEntry)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.area_statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats_entry")
	proto.RegisterType((*OspfShIfStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.area_statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats")
}

func init() { proto.RegisterFile("ospf_sh_if_stats.proto", fileDescriptor_2ceab17d939ff790) }

var fileDescriptor_2ceab17d939ff790 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x19, 0x2b, 0xfd, 0x49, 0xad, 0x48, 0xc0, 0x3a, 0x20, 0x48, 0x2d, 0x08, 0x5d, 0x65,
	0xd1, 0xf6, 0x15, 0xba, 0x28, 0x82, 0x8b, 0xe9, 0xca, 0xd5, 0x25, 0xce, 0x24, 0xf4, 0x4a, 0x3b,
	0x19, 0x92, 0x6b, 0x51, 0x70, 0xef, 0x0b, 0xf8, 0x20, 0xbe, 0xa0, 0x20, 0xb9, 0x33, 0x2a, 0xd4,
	0xee, 0xdd, 0xdd, 0x9c, 0x9c, 0x9c, 0x2f, 0x67, 0x26, 0x62, 0xe8, 0x42, 0x65, 0x21, 0xac, 0x01,
	0x2d, 0x04, 0xd2, 0x14, 0x54, 0xe5, 0x1d, 0x39, 0xf9, 0x9a, 0x63, 0xc8, 0x1d, 0xa0, 0x0b, 0xf0,
	0xec, 0x01, 0xab, 0xdd, 0x1c, 0xd8, 0xe9, 0x2a, 0xe3, 0x55, 0x9c, 0xa2, 0x2f, 0x37, 0x21, 0x98,
	0xf0, 0x3d, 0xa9, 0xc2, 0x58, 0xfd, 0xb4, 0x21, 0xd8, 0x79, 0xab, 0xb4, 0x37, 0x9a, 0x03, 0x31,
	0x10, 0xe6, 0x41, 0x61, 0x49, 0xc6, 0x5b, 0x9d, 0x9b, 0x9a, 0x02, 0xa6, 0x24, 0x8f, 0xe6, 0xb0,
	0xfe, 0x32, 0xd6, 0xe2, 0x7c, 0xff, 0x5e, 0x70, 0xbb, 0xb8, 0x5f, 0xc9, 0x6b, 0x71, 0xd2, 0xd0,
	0xa0, 0xd4, 0x5b, 0x93, 0x26, 0xa3, 0x64, 0xd2, 0xcb, 0xfa, 0x8d, 0x76, 0xa7, 0xb7, 0x46, 0xde,
	0x88, 0xd3, 0xdf, 0x50, 0x36, 0x1d, 0xb1, 0x69, 0xf0, 0xa3, 0x46, 0xdb, 0xf8, 0x2d, 0xf9, 0xdb,
	0xbd, 0xa6, 0xcb, 0x4b, 0xd1, 0xdb, 0xea, 0xe2, 0x11, 0xb0, 0x24, 0xcb, 0x84, 0x6e, 0xd6, 0x8d,
	0xc2, 0xb2, 0x24, 0x2b, 0x2f, 0x44, 0x87, 0xbb, 0x61, 0xc1, 0xb9, 0x83, 0xac, 0x1d, 0x97, 0xcb,
	0x42, 0x5e, 0x89, 0x7e, 0xb3, 0x01, 0x81, 0x7c, 0xda, 0x62, 0x68, 0xaf, 0xde, 0x5c, 0x91, 0x8f,
	0x07, 0x1b, 0x4e, 0x7a, 0x3c, 0x6a, 0xc5, 0x83, 0x68, 0x57, 0xa4, 0x69, 0xfc, 0x99, 0x88, 0xb3,
	0xfd, 0x9b, 0xc8, 0xa1, 0x68, 0xaf, 0x75, 0x59, 0x6c, 0x4c, 0x3a, 0xe5, 0xa0, 0x66, 0x15, 0x29,
	0x68, 0xb9, 0x16, 0x53, 0x66, 0x35, 0x05, 0x6d, 0xec, 0x14, 0x29, 0x1f, 0x89, 0xe8, 0xa2, 0xad,
	0x8b, 0xa4, 0xf3, 0x51, 0x6b, 0xd2, 0x9f, 0xbe, 0x27, 0xea, 0x3f, 0x7f, 0xa6, 0x3a, 0xfc, 0x95,
	0xb3, 0x0e, 0xda, 0x45, 0x1c, 0x1e, 0xda, 0xfc, 0xe2, 0x66, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x18, 0x9f, 0xb8, 0xbb, 0x8b, 0x02, 0x00, 0x00,
}