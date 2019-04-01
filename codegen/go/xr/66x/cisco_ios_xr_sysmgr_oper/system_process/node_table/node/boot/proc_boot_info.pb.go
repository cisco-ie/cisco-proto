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
// source: proc_boot_info.proto

package cisco_ios_xr_sysmgr_oper_system_process_node_table_node_boot

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

type ProcBootInfo_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcBootInfo_KEYS) Reset()         { *m = ProcBootInfo_KEYS{} }
func (m *ProcBootInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*ProcBootInfo_KEYS) ProtoMessage()    {}
func (*ProcBootInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_079c659ca0c40559, []int{0}
}

func (m *ProcBootInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBootInfo_KEYS.Unmarshal(m, b)
}
func (m *ProcBootInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBootInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *ProcBootInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBootInfo_KEYS.Merge(m, src)
}
func (m *ProcBootInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_ProcBootInfo_KEYS.Size(m)
}
func (m *ProcBootInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBootInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBootInfo_KEYS proto.InternalMessageInfo

func (m *ProcBootInfo_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type BandStatsInfo struct {
	Level                string   `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty"`
	BandName             string   `protobuf:"bytes,2,opt,name=band_name,json=bandName,proto3" json:"band_name,omitempty"`
	BandFinishTime       string   `protobuf:"bytes,3,opt,name=band_finish_time,json=bandFinishTime,proto3" json:"band_finish_time,omitempty"`
	BandTime             string   `protobuf:"bytes,4,opt,name=band_time,json=bandTime,proto3" json:"band_time,omitempty"`
	FinishTime           string   `protobuf:"bytes,5,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	IdlePercentage       string   `protobuf:"bytes,6,opt,name=idle_percentage,json=idlePercentage,proto3" json:"idle_percentage,omitempty"`
	Jid                  uint32   `protobuf:"varint,7,opt,name=jid,proto3" json:"jid,omitempty"`
	ReadyTime            string   `protobuf:"bytes,8,opt,name=ready_time,json=readyTime,proto3" json:"ready_time,omitempty"`
	LastProcess          string   `protobuf:"bytes,9,opt,name=last_process,json=lastProcess,proto3" json:"last_process,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BandStatsInfo) Reset()         { *m = BandStatsInfo{} }
func (m *BandStatsInfo) String() string { return proto.CompactTextString(m) }
func (*BandStatsInfo) ProtoMessage()    {}
func (*BandStatsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_079c659ca0c40559, []int{1}
}

func (m *BandStatsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BandStatsInfo.Unmarshal(m, b)
}
func (m *BandStatsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BandStatsInfo.Marshal(b, m, deterministic)
}
func (m *BandStatsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BandStatsInfo.Merge(m, src)
}
func (m *BandStatsInfo) XXX_Size() int {
	return xxx_messageInfo_BandStatsInfo.Size(m)
}
func (m *BandStatsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BandStatsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BandStatsInfo proto.InternalMessageInfo

func (m *BandStatsInfo) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *BandStatsInfo) GetBandName() string {
	if m != nil {
		return m.BandName
	}
	return ""
}

func (m *BandStatsInfo) GetBandFinishTime() string {
	if m != nil {
		return m.BandFinishTime
	}
	return ""
}

func (m *BandStatsInfo) GetBandTime() string {
	if m != nil {
		return m.BandTime
	}
	return ""
}

func (m *BandStatsInfo) GetFinishTime() string {
	if m != nil {
		return m.FinishTime
	}
	return ""
}

func (m *BandStatsInfo) GetIdlePercentage() string {
	if m != nil {
		return m.IdlePercentage
	}
	return ""
}

func (m *BandStatsInfo) GetJid() uint32 {
	if m != nil {
		return m.Jid
	}
	return 0
}

func (m *BandStatsInfo) GetReadyTime() string {
	if m != nil {
		return m.ReadyTime
	}
	return ""
}

func (m *BandStatsInfo) GetLastProcess() string {
	if m != nil {
		return m.LastProcess
	}
	return ""
}

type BootedProcessInfo struct {
	StartTimeStamp       string   `protobuf:"bytes,1,opt,name=start_time_stamp,json=startTimeStamp,proto3" json:"start_time_stamp,omitempty"`
	Started              string   `protobuf:"bytes,2,opt,name=started,proto3" json:"started,omitempty"`
	Level                string   `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	Jid                  int32    `protobuf:"zigzag32,4,opt,name=jid,proto3" json:"jid,omitempty"`
	InstanceId           int32    `protobuf:"zigzag32,5,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	ReadyTimeStamp       string   `protobuf:"bytes,6,opt,name=ready_time_stamp,json=readyTimeStamp,proto3" json:"ready_time_stamp,omitempty"`
	Ready                string   `protobuf:"bytes,7,opt,name=ready,proto3" json:"ready,omitempty"`
	IsEoiTimeout         bool     `protobuf:"varint,8,opt,name=is_eoi_timeout,json=isEoiTimeout,proto3" json:"is_eoi_timeout,omitempty"`
	ProcessName          string   `protobuf:"bytes,9,opt,name=process_name,json=processName,proto3" json:"process_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BootedProcessInfo) Reset()         { *m = BootedProcessInfo{} }
func (m *BootedProcessInfo) String() string { return proto.CompactTextString(m) }
func (*BootedProcessInfo) ProtoMessage()    {}
func (*BootedProcessInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_079c659ca0c40559, []int{2}
}

func (m *BootedProcessInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BootedProcessInfo.Unmarshal(m, b)
}
func (m *BootedProcessInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BootedProcessInfo.Marshal(b, m, deterministic)
}
func (m *BootedProcessInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootedProcessInfo.Merge(m, src)
}
func (m *BootedProcessInfo) XXX_Size() int {
	return xxx_messageInfo_BootedProcessInfo.Size(m)
}
func (m *BootedProcessInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BootedProcessInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BootedProcessInfo proto.InternalMessageInfo

func (m *BootedProcessInfo) GetStartTimeStamp() string {
	if m != nil {
		return m.StartTimeStamp
	}
	return ""
}

func (m *BootedProcessInfo) GetStarted() string {
	if m != nil {
		return m.Started
	}
	return ""
}

func (m *BootedProcessInfo) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *BootedProcessInfo) GetJid() int32 {
	if m != nil {
		return m.Jid
	}
	return 0
}

func (m *BootedProcessInfo) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

func (m *BootedProcessInfo) GetReadyTimeStamp() string {
	if m != nil {
		return m.ReadyTimeStamp
	}
	return ""
}

func (m *BootedProcessInfo) GetReady() string {
	if m != nil {
		return m.Ready
	}
	return ""
}

func (m *BootedProcessInfo) GetIsEoiTimeout() bool {
	if m != nil {
		return m.IsEoiTimeout
	}
	return false
}

func (m *BootedProcessInfo) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

type ProcBootInfo struct {
	StandbyBandStatistic []*BandStatsInfo     `protobuf:"bytes,50,rep,name=standby_band_statistic,json=standbyBandStatistic,proto3" json:"standby_band_statistic,omitempty"`
	ActiveBandStatistic  []*BandStatsInfo     `protobuf:"bytes,51,rep,name=active_band_statistic,json=activeBandStatistic,proto3" json:"active_band_statistic,omitempty"`
	BootedProcess        []*BootedProcessInfo `protobuf:"bytes,52,rep,name=booted_process,json=bootedProcess,proto3" json:"booted_process,omitempty"`
	LastProcessStarted   string               `protobuf:"bytes,53,opt,name=last_process_started,json=lastProcessStarted,proto3" json:"last_process_started,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProcBootInfo) Reset()         { *m = ProcBootInfo{} }
func (m *ProcBootInfo) String() string { return proto.CompactTextString(m) }
func (*ProcBootInfo) ProtoMessage()    {}
func (*ProcBootInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_079c659ca0c40559, []int{3}
}

func (m *ProcBootInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcBootInfo.Unmarshal(m, b)
}
func (m *ProcBootInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcBootInfo.Marshal(b, m, deterministic)
}
func (m *ProcBootInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcBootInfo.Merge(m, src)
}
func (m *ProcBootInfo) XXX_Size() int {
	return xxx_messageInfo_ProcBootInfo.Size(m)
}
func (m *ProcBootInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcBootInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProcBootInfo proto.InternalMessageInfo

func (m *ProcBootInfo) GetStandbyBandStatistic() []*BandStatsInfo {
	if m != nil {
		return m.StandbyBandStatistic
	}
	return nil
}

func (m *ProcBootInfo) GetActiveBandStatistic() []*BandStatsInfo {
	if m != nil {
		return m.ActiveBandStatistic
	}
	return nil
}

func (m *ProcBootInfo) GetBootedProcess() []*BootedProcessInfo {
	if m != nil {
		return m.BootedProcess
	}
	return nil
}

func (m *ProcBootInfo) GetLastProcessStarted() string {
	if m != nil {
		return m.LastProcessStarted
	}
	return ""
}

func init() {
	proto.RegisterType((*ProcBootInfo_KEYS)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot.proc_boot_info_KEYS")
	proto.RegisterType((*BandStatsInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot.band_stats_info")
	proto.RegisterType((*BootedProcessInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot.booted_process_info")
	proto.RegisterType((*ProcBootInfo)(nil), "cisco_ios_xr_sysmgr_oper.system_process.node_table.node.boot.proc_boot_info")
}

func init() { proto.RegisterFile("proc_boot_info.proto", fileDescriptor_079c659ca0c40559) }

var fileDescriptor_079c659ca0c40559 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0xc6, 0x95, 0x84, 0xb4, 0xc9, 0xa4, 0x4d, 0x53, 0x27, 0xa0, 0x95, 0x10, 0x6a, 0x88, 0x90,
	0xc8, 0x69, 0x85, 0x52, 0xb8, 0x71, 0x42, 0x2a, 0x12, 0x42, 0xa0, 0xb2, 0xe9, 0x85, 0x93, 0xe5,
	0x5d, 0x4f, 0x8a, 0xd1, 0xee, 0x7a, 0xb5, 0x36, 0x55, 0x73, 0x44, 0x3c, 0x0d, 0xe2, 0x3d, 0x78,
	0x2e, 0xe4, 0xb1, 0x37, 0x7f, 0xca, 0xb1, 0xea, 0xcd, 0xfe, 0x79, 0xfc, 0xcd, 0xf8, 0xcb, 0x97,
	0x85, 0x49, 0x55, 0xeb, 0x8c, 0xa7, 0x5a, 0x5b, 0xae, 0xca, 0x95, 0x8e, 0xab, 0x5a, 0x5b, 0xcd,
	0xde, 0x66, 0xca, 0x64, 0x9a, 0x2b, 0x6d, 0xf8, 0x6d, 0xcd, 0xcd, 0xda, 0x14, 0xd7, 0x35, 0xd7,
	0x15, 0xd6, 0xb1, 0x59, 0x1b, 0x8b, 0x05, 0x77, 0xb7, 0xd0, 0x98, 0xb8, 0xd4, 0x12, 0xb9, 0x15,
	0x69, 0x8e, 0xb4, 0x8c, 0x9d, 0xd0, 0x6c, 0x01, 0xe3, 0x7d, 0x55, 0xfe, 0xf1, 0xe2, 0xeb, 0x92,
	0x3d, 0x85, 0x3e, 0x95, 0x97, 0xa2, 0xc0, 0xa8, 0x35, 0x6d, 0xcd, 0xfb, 0x49, 0xcf, 0x81, 0xcf,
	0xa2, 0xc0, 0xd9, 0xef, 0x36, 0x9c, 0xa4, 0xa2, 0x94, 0xdc, 0x58, 0x61, 0x0d, 0xdd, 0x62, 0x13,
	0xe8, 0xe6, 0x78, 0x83, 0x79, 0x28, 0xf6, 0x1b, 0x27, 0x43, 0x85, 0x24, 0xd3, 0xf6, 0x32, 0x0e,
	0x38, 0x19, 0x36, 0x87, 0x11, 0x1d, 0xae, 0x54, 0xa9, 0xcc, 0x37, 0x6e, 0x55, 0x81, 0x51, 0x87,
	0x6a, 0x86, 0x8e, 0xbf, 0x27, 0x7c, 0xa5, 0x0a, 0xdc, 0xc8, 0x50, 0xc9, 0xa3, 0xad, 0x0c, 0x1d,
	0x9e, 0xc1, 0x60, 0x57, 0xa1, 0x4b, 0xc7, 0xb0, 0xda, 0xde, 0x7e, 0x09, 0x27, 0x4a, 0xe6, 0xc8,
	0x2b, 0xac, 0x33, 0x2c, 0xad, 0xb8, 0xc6, 0xe8, 0xc0, 0xb7, 0x71, 0xf8, 0x72, 0x43, 0xd9, 0x08,
	0x3a, 0xdf, 0x95, 0x8c, 0x0e, 0xa7, 0xad, 0xf9, 0x71, 0xe2, 0x96, 0xec, 0x19, 0x40, 0x8d, 0x42,
	0xae, 0xbd, 0x74, 0x8f, 0x6e, 0xf5, 0x89, 0x90, 0xf2, 0x73, 0x38, 0xca, 0x85, 0xb1, 0x8d, 0xc3,
	0x51, 0x9f, 0x0a, 0x06, 0x8e, 0x5d, 0x7a, 0x34, 0xfb, 0xd3, 0x86, 0xb1, 0xf3, 0x16, 0x65, 0x53,
	0xe5, 0xfd, 0x9a, 0xc3, 0xc8, 0x58, 0x51, 0x5b, 0x52, 0x76, 0x46, 0x16, 0x55, 0xb0, 0x6e, 0x48,
	0xdc, 0xe9, 0x2f, 0x1d, 0x65, 0x11, 0x1c, 0x12, 0x41, 0x19, 0x1c, 0x6c, 0xb6, 0x5b, 0xcf, 0x3b,
	0xbb, 0x9e, 0x87, 0x57, 0x38, 0x9b, 0x4e, 0xfd, 0x2b, 0xce, 0x60, 0xa0, 0x4a, 0x63, 0x45, 0x99,
	0x21, 0x57, 0x92, 0x1c, 0x3a, 0x4d, 0xa0, 0x41, 0x1f, 0xa4, 0x1b, 0x66, 0xfb, 0xcc, 0x30, 0x4c,
	0xb0, 0x68, 0xf3, 0x58, 0x3f, 0xcc, 0x04, 0xba, 0x44, 0xc8, 0xa4, 0x7e, 0xe2, 0x37, 0xec, 0x05,
	0x0c, 0x95, 0xe1, 0xa8, 0x15, 0x09, 0xe8, 0x1f, 0x96, 0xac, 0xea, 0x25, 0x47, 0xca, 0x5c, 0x68,
	0x75, 0xe5, 0x99, 0x73, 0xab, 0xb1, 0x80, 0xf2, 0x10, 0xdc, 0x0a, 0x8c, 0x92, 0xf5, 0xb7, 0x03,
	0xc3, 0xfd, 0x38, 0xb2, 0x5f, 0x2d, 0x78, 0xe2, 0x06, 0x95, 0xe9, 0x9a, 0x6f, 0x42, 0xa7, 0x8c,
	0x55, 0x59, 0xb4, 0x98, 0x76, 0xe6, 0x83, 0xc5, 0xa7, 0xf8, 0x3e, 0x7f, 0x80, 0xf8, 0x4e, 0x90,
	0x93, 0x49, 0x68, 0xf6, 0x4e, 0x94, 0x72, 0xd9, 0xb4, 0x62, 0x3f, 0x5b, 0xf0, 0x58, 0x64, 0x56,
	0xdd, 0xe0, 0xdd, 0x21, 0xce, 0x1f, 0x62, 0x88, 0xb1, 0xef, 0xb5, 0x3f, 0xc3, 0x2d, 0x0c, 0xf7,
	0x93, 0x14, 0xbd, 0xa6, 0xde, 0x5f, 0xee, 0xd9, 0xfb, 0xff, 0x74, 0x26, 0xc7, 0x1e, 0x86, 0x10,
	0xb3, 0x57, 0x30, 0xd9, 0xcd, 0x39, 0x6f, 0xf2, 0xf8, 0x86, 0x7e, 0x41, 0xb6, 0x93, 0xf7, 0xa5,
	0x3f, 0x49, 0x0f, 0xe8, 0xdb, 0x74, 0xfe, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x67, 0xdd, 0x42,
	0xb3, 0x04, 0x00, 0x00,
}
