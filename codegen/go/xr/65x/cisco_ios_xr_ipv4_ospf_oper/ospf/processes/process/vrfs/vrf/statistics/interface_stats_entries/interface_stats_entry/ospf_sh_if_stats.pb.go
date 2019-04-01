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

package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_vrfs_vrf_statistics_interface_stats_entries_interface_stats_entry

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
	VrfName              string   `protobuf:"bytes,2,opt,name=vrf_name,json=vrfName,proto3" json:"vrf_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
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

func (m *OspfShIfStats_KEYS) GetVrfName() string {
	if m != nil {
		return m.VrfName
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
	proto.RegisterType((*OspfShIfStats_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats_KEYS")
	proto.RegisterType((*OspfShIfStatsEntry)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats_entry")
	proto.RegisterType((*OspfShIfStats)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.vrfs.vrf.statistics.interface_stats_entries.interface_stats_entry.ospf_sh_if_stats")
}

func init() { proto.RegisterFile("ospf_sh_if_stats.proto", fileDescriptor_2ceab17d939ff790) }

var fileDescriptor_2ceab17d939ff790 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xc7, 0xd9, 0xa7, 0x0f, 0xed, 0x36, 0xb5, 0x22, 0x01, 0xeb, 0x8a, 0x20, 0xb5, 0x20, 0xf4,
	0x94, 0x43, 0xdb, 0xaf, 0xd0, 0x43, 0x11, 0x3c, 0x6c, 0x4f, 0x9e, 0x86, 0xb8, 0x9b, 0xd0, 0x11,
	0xf7, 0x85, 0x64, 0x58, 0x14, 0xbc, 0x7b, 0xf1, 0x83, 0xf8, 0x05, 0xbd, 0x4b, 0x5e, 0x54, 0xa8,
	0xbd, 0x7b, 0x59, 0x26, 0xf3, 0xff, 0x67, 0x7e, 0x33, 0xb3, 0x61, 0x93, 0xc6, 0xb6, 0x1a, 0xec,
	0x0e, 0x50, 0x83, 0x25, 0x49, 0x56, 0xb4, 0xa6, 0xa1, 0x86, 0x53, 0x81, 0xb6, 0x68, 0x00, 0x1b,
	0x0b, 0x4f, 0x06, 0xb0, 0xed, 0x56, 0xe0, 0x9d, 0x4d, 0xab, 0x8c, 0x70, 0x91, 0xf3, 0x15, 0xca,
	0x5a, 0x65, 0xbf, 0x22, 0xd1, 0x19, 0xed, 0x3f, 0xc2, 0x15, 0x42, 0x4b, 0x58, 0x58, 0x81, 0x35,
	0x29, 0xa3, 0x65, 0xa1, 0x42, 0x75, 0x50, 0x35, 0x19, 0x54, 0x87, 0xf3, 0xcf, 0xb3, 0x17, 0x76,
	0xba, 0xdf, 0x0f, 0xdc, 0xac, 0xef, 0xb6, 0xfc, 0x8a, 0x1d, 0x45, 0x0a, 0xd4, 0xb2, 0x52, 0x59,
	0x32, 0x4d, 0xe6, 0xc3, 0x7c, 0x14, 0x73, 0xb7, 0xb2, 0x52, 0xfc, 0x9c, 0xa5, 0x9d, 0xd1, 0x41,
	0xfe, 0xe7, 0xe5, 0x41, 0x67, 0xb4, 0x97, 0xae, 0xd9, 0xf1, 0x0f, 0xcf, 0x1b, 0x7a, 0xde, 0x30,
	0xfe, 0xce, 0x3a, 0xdb, 0xec, 0x35, 0xf9, 0xbd, 0x8e, 0xd0, 0x18, 0xbf, 0x60, 0xc3, 0x4a, 0x96,
	0x0f, 0x80, 0x35, 0x69, 0x0f, 0x4f, 0xf3, 0xd4, 0x25, 0x36, 0x35, 0x69, 0x7e, 0xc6, 0x06, 0xd2,
	0x28, 0x09, 0x58, 0x7a, 0xf0, 0x38, 0xef, 0xbb, 0xe3, 0xa6, 0xe4, 0x97, 0x6c, 0x14, 0x05, 0xb0,
	0x64, 0x22, 0x74, 0x18, 0xc4, 0x2d, 0x19, 0x77, 0x31, 0x72, 0xb2, 0xff, 0xd3, 0x9e, 0xbb, 0x88,
	0x7a, 0x4b, 0x92, 0x66, 0x1f, 0x09, 0x3b, 0xd9, 0xef, 0x84, 0x4f, 0x58, 0x7f, 0x27, 0xeb, 0xf2,
	0x51, 0x65, 0x0b, 0x5f, 0x28, 0x9e, 0x1c, 0x05, 0xc3, 0xdc, 0x9e, 0xb2, 0x0c, 0x14, 0xf4, 0xa3,
	0x3b, 0xca, 0x7b, 0xc2, 0x52, 0xd4, 0x61, 0x90, 0x6c, 0x35, 0xed, 0xcd, 0x47, 0x8b, 0xb7, 0x44,
	0xfc, 0xc5, 0xff, 0x15, 0x87, 0xb7, 0x9b, 0x0f, 0x50, 0xaf, 0x5d, 0x70, 0xdf, 0xf7, 0x8f, 0x6f,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0x28, 0x79, 0x45, 0x32, 0x96, 0x02, 0x00, 0x00,
}
