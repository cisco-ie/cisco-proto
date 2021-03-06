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
// source: te_stats_autobackup_t.proto

package cisco_ios_xr_mpls_te_oper_mpls_te_standby_auto_tunnel_backup_statistics

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

type TeStatsAutobackupT_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsAutobackupT_KEYS) Reset()         { *m = TeStatsAutobackupT_KEYS{} }
func (m *TeStatsAutobackupT_KEYS) String() string { return proto.CompactTextString(m) }
func (*TeStatsAutobackupT_KEYS) ProtoMessage()    {}
func (*TeStatsAutobackupT_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3473e6b107a6f9, []int{0}
}

func (m *TeStatsAutobackupT_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsAutobackupT_KEYS.Unmarshal(m, b)
}
func (m *TeStatsAutobackupT_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsAutobackupT_KEYS.Marshal(b, m, deterministic)
}
func (m *TeStatsAutobackupT_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsAutobackupT_KEYS.Merge(m, src)
}
func (m *TeStatsAutobackupT_KEYS) XXX_Size() int {
	return xxx_messageInfo_TeStatsAutobackupT_KEYS.Size(m)
}
func (m *TeStatsAutobackupT_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsAutobackupT_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsAutobackupT_KEYS proto.InternalMessageInfo

type TeStatsAutobackupCountersT struct {
	TotalTunnels         uint32   `protobuf:"varint,1,opt,name=total_tunnels,json=totalTunnels,proto3" json:"total_tunnels,omitempty"`
	NextHopTunnels       uint32   `protobuf:"varint,2,opt,name=next_hop_tunnels,json=nextHopTunnels,proto3" json:"next_hop_tunnels,omitempty"`
	NextNextHopTunnels   uint32   `protobuf:"varint,3,opt,name=next_next_hop_tunnels,json=nextNextHopTunnels,proto3" json:"next_next_hop_tunnels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeStatsAutobackupCountersT) Reset()         { *m = TeStatsAutobackupCountersT{} }
func (m *TeStatsAutobackupCountersT) String() string { return proto.CompactTextString(m) }
func (*TeStatsAutobackupCountersT) ProtoMessage()    {}
func (*TeStatsAutobackupCountersT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3473e6b107a6f9, []int{1}
}

func (m *TeStatsAutobackupCountersT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsAutobackupCountersT.Unmarshal(m, b)
}
func (m *TeStatsAutobackupCountersT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsAutobackupCountersT.Marshal(b, m, deterministic)
}
func (m *TeStatsAutobackupCountersT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsAutobackupCountersT.Merge(m, src)
}
func (m *TeStatsAutobackupCountersT) XXX_Size() int {
	return xxx_messageInfo_TeStatsAutobackupCountersT.Size(m)
}
func (m *TeStatsAutobackupCountersT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsAutobackupCountersT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsAutobackupCountersT proto.InternalMessageInfo

func (m *TeStatsAutobackupCountersT) GetTotalTunnels() uint32 {
	if m != nil {
		return m.TotalTunnels
	}
	return 0
}

func (m *TeStatsAutobackupCountersT) GetNextHopTunnels() uint32 {
	if m != nil {
		return m.NextHopTunnels
	}
	return 0
}

func (m *TeStatsAutobackupCountersT) GetNextNextHopTunnels() uint32 {
	if m != nil {
		return m.NextNextHopTunnels
	}
	return 0
}

type TeStatsAutobackupT struct {
	Created              *TeStatsAutobackupCountersT `protobuf:"bytes,50,opt,name=created,proto3" json:"created,omitempty"`
	Connected            *TeStatsAutobackupCountersT `protobuf:"bytes,51,opt,name=connected,proto3" json:"connected,omitempty"`
	RemovedDown          *TeStatsAutobackupCountersT `protobuf:"bytes,52,opt,name=removed_down,json=removedDown,proto3" json:"removed_down,omitempty"`
	RemovedUnused        *TeStatsAutobackupCountersT `protobuf:"bytes,53,opt,name=removed_unused,json=removedUnused,proto3" json:"removed_unused,omitempty"`
	RemovedInUse         *TeStatsAutobackupCountersT `protobuf:"bytes,54,opt,name=removed_in_use,json=removedInUse,proto3" json:"removed_in_use,omitempty"`
	RemovedRangeExceeded *TeStatsAutobackupCountersT `protobuf:"bytes,55,opt,name=removed_range_exceeded,json=removedRangeExceeded,proto3" json:"removed_range_exceeded,omitempty"`
	LastClearedTime      uint32                      `protobuf:"varint,56,opt,name=last_cleared_time,json=lastClearedTime,proto3" json:"last_cleared_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TeStatsAutobackupT) Reset()         { *m = TeStatsAutobackupT{} }
func (m *TeStatsAutobackupT) String() string { return proto.CompactTextString(m) }
func (*TeStatsAutobackupT) ProtoMessage()    {}
func (*TeStatsAutobackupT) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d3473e6b107a6f9, []int{2}
}

func (m *TeStatsAutobackupT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeStatsAutobackupT.Unmarshal(m, b)
}
func (m *TeStatsAutobackupT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeStatsAutobackupT.Marshal(b, m, deterministic)
}
func (m *TeStatsAutobackupT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeStatsAutobackupT.Merge(m, src)
}
func (m *TeStatsAutobackupT) XXX_Size() int {
	return xxx_messageInfo_TeStatsAutobackupT.Size(m)
}
func (m *TeStatsAutobackupT) XXX_DiscardUnknown() {
	xxx_messageInfo_TeStatsAutobackupT.DiscardUnknown(m)
}

var xxx_messageInfo_TeStatsAutobackupT proto.InternalMessageInfo

func (m *TeStatsAutobackupT) GetCreated() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *TeStatsAutobackupT) GetConnected() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.Connected
	}
	return nil
}

func (m *TeStatsAutobackupT) GetRemovedDown() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.RemovedDown
	}
	return nil
}

func (m *TeStatsAutobackupT) GetRemovedUnused() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.RemovedUnused
	}
	return nil
}

func (m *TeStatsAutobackupT) GetRemovedInUse() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.RemovedInUse
	}
	return nil
}

func (m *TeStatsAutobackupT) GetRemovedRangeExceeded() *TeStatsAutobackupCountersT {
	if m != nil {
		return m.RemovedRangeExceeded
	}
	return nil
}

func (m *TeStatsAutobackupT) GetLastClearedTime() uint32 {
	if m != nil {
		return m.LastClearedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*TeStatsAutobackupT_KEYS)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.backup.statistics.te_stats_autobackup_t_KEYS")
	proto.RegisterType((*TeStatsAutobackupCountersT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.backup.statistics.te_stats_autobackup_counters_t")
	proto.RegisterType((*TeStatsAutobackupT)(nil), "cisco_ios_xr_mpls_te_oper.mpls_te_standby.auto_tunnel.backup.statistics.te_stats_autobackup_t")
}

func init() { proto.RegisterFile("te_stats_autobackup_t.proto", fileDescriptor_0d3473e6b107a6f9) }

var fileDescriptor_0d3473e6b107a6f9 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xcb, 0x6b, 0xdb, 0x40,
	0x10, 0xc6, 0xd9, 0x16, 0x6a, 0xba, 0x7e, 0xb4, 0x5d, 0xea, 0x22, 0xda, 0x52, 0x8c, 0x7b, 0x31,
	0x3d, 0x08, 0x6a, 0xf7, 0x75, 0x6f, 0x4d, 0x13, 0x02, 0x39, 0x38, 0xf6, 0x21, 0xa7, 0x61, 0xbd,
	0x3b, 0x38, 0x22, 0xd2, 0xae, 0xd8, 0x1d, 0xc5, 0xce, 0x31, 0x87, 0x5c, 0x02, 0xc9, 0x7f, 0xe0,
	0x4b, 0xfe, 0xd2, 0x20, 0x59, 0xca, 0x83, 0x98, 0x9c, 0x82, 0x8e, 0xfa, 0xe6, 0x9b, 0xf9, 0x7e,
	0x9a, 0x65, 0x97, 0x7f, 0x22, 0x04, 0x4f, 0x92, 0x3c, 0xc8, 0x8c, 0xec, 0x5c, 0xaa, 0xe3, 0x2c,
	0x05, 0x0a, 0x53, 0x67, 0xc9, 0x8a, 0xff, 0x2a, 0xf2, 0xca, 0x42, 0x64, 0x3d, 0xac, 0x1c, 0x24,
	0x69, 0xec, 0x81, 0x10, 0x6c, 0x8a, 0x2e, 0xac, 0x3e, 0x3c, 0x49, 0xa3, 0xe7, 0xa7, 0x61, 0xde,
	0x0d, 0x94, 0x19, 0x83, 0x71, 0xb8, 0x99, 0x12, 0xe6, 0x63, 0x23, 0x4f, 0x91, 0xf2, 0xfd, 0xcf,
	0xfc, 0xe3, 0xd6, 0x1c, 0xd8, 0x1b, 0x1f, 0x1e, 0xf4, 0xaf, 0x19, 0xff, 0xb2, 0xad, 0xac, 0x6c,
	0x66, 0x08, 0x9d, 0x07, 0x12, 0x5f, 0x79, 0x9b, 0x2c, 0xc9, 0xb8, 0x8c, 0xf0, 0x01, 0xeb, 0xb1,
	0x41, 0x7b, 0xd2, 0x2a, 0xc4, 0xe9, 0x46, 0x13, 0x03, 0xfe, 0xd6, 0xe0, 0x8a, 0xe0, 0xc8, 0xa6,
	0xb7, 0xbe, 0x17, 0x85, 0xaf, 0x93, 0xeb, 0x3b, 0x36, 0xad, 0x9c, 0xdf, 0x79, 0xb7, 0x70, 0x3e,
	0xb2, 0xbf, 0x2c, 0xec, 0x22, 0xd7, 0xf7, 0x1f, 0xb4, 0xf4, 0xd7, 0x0d, 0xde, 0xdd, 0xfa, 0x0f,
	0xe2, 0x8c, 0xf1, 0x86, 0x72, 0x28, 0x09, 0x75, 0x30, 0xec, 0xb1, 0x41, 0x73, 0xb8, 0x08, 0x9f,
	0x69, 0x71, 0xe1, 0xd3, 0x6b, 0x99, 0x54, 0xb9, 0xe2, 0x9c, 0xf1, 0xd7, 0xca, 0x1a, 0x83, 0x2a,
	0xa7, 0x18, 0xd5, 0x4b, 0x71, 0x97, 0x2c, 0x2e, 0x18, 0x6f, 0x39, 0x4c, 0xec, 0x09, 0x6a, 0xd0,
	0x76, 0x69, 0x82, 0x1f, 0xf5, 0xa2, 0x34, 0xcb, 0xf0, 0x7f, 0x76, 0x69, 0xc4, 0x15, 0xe3, 0x9d,
	0x0a, 0x26, 0x33, 0x99, 0x47, 0x1d, 0xfc, 0xac, 0x17, 0xa7, 0x5d, 0xc6, 0xcf, 0x8a, 0x74, 0x71,
	0x79, 0x0f, 0x28, 0x32, 0x90, 0x79, 0x0c, 0x7e, 0xd5, 0x0b, 0x54, 0x1d, 0xce, 0xae, 0x99, 0x79,
	0x14, 0x6b, 0xc6, 0x3f, 0x54, 0x3c, 0x4e, 0x9a, 0x05, 0x02, 0xae, 0x14, 0xa2, 0x46, 0x1d, 0xfc,
	0xae, 0x97, 0xeb, 0x7d, 0x89, 0x31, 0xc9, 0x29, 0xc6, 0x25, 0x84, 0xf8, 0xc6, 0xdf, 0xc5, 0xd2,
	0x13, 0xa8, 0x18, 0xa5, 0x43, 0x0d, 0x14, 0x25, 0x18, 0xfc, 0x29, 0xae, 0xe8, 0x9b, 0xbc, 0xf0,
	0x77, 0xa3, 0x4f, 0xa3, 0x04, 0xe7, 0xaf, 0x8a, 0x27, 0x6b, 0x74, 0x13, 0x00, 0x00, 0xff, 0xff,
	0x06, 0x37, 0xbc, 0xd5, 0xd1, 0x04, 0x00, 0x00,
}
