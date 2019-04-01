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
// source: alarm_mgr_show_alarm_stats.proto

package cisco_ios_xr_alarmgr_server_oper_alarms_detail_detail_card_detail_locations_detail_location_stats

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

type AlarmMgrShowAlarmStats_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmStats_KEYS) Reset()         { *m = AlarmMgrShowAlarmStats_KEYS{} }
func (m *AlarmMgrShowAlarmStats_KEYS) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmStats_KEYS) ProtoMessage()    {}
func (*AlarmMgrShowAlarmStats_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a57311bc831b075f, []int{0}
}

func (m *AlarmMgrShowAlarmStats_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmStats_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmStats_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS.Merge(m, src)
}
func (m *AlarmMgrShowAlarmStats_KEYS) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS.Size(m)
}
func (m *AlarmMgrShowAlarmStats_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmStats_KEYS proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmStats_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type AlarmMgrShowAlarmStats struct {
	Reported               uint64   `protobuf:"varint,50,opt,name=reported,proto3" json:"reported,omitempty"`
	Dropped                uint64   `protobuf:"varint,51,opt,name=dropped,proto3" json:"dropped,omitempty"`
	Active                 uint64   `protobuf:"varint,52,opt,name=active,proto3" json:"active,omitempty"`
	History                uint64   `protobuf:"varint,53,opt,name=history,proto3" json:"history,omitempty"`
	Suppressed             uint64   `protobuf:"varint,54,opt,name=suppressed,proto3" json:"suppressed,omitempty"`
	SysadminActive         uint64   `protobuf:"varint,55,opt,name=sysadmin_active,json=sysadminActive,proto3" json:"sysadmin_active,omitempty"`
	SysadminHistory        uint64   `protobuf:"varint,56,opt,name=sysadmin_history,json=sysadminHistory,proto3" json:"sysadmin_history,omitempty"`
	SysadminSuppressed     uint64   `protobuf:"varint,57,opt,name=sysadmin_suppressed,json=sysadminSuppressed,proto3" json:"sysadmin_suppressed,omitempty"`
	DroppedInvalidAid      uint32   `protobuf:"varint,58,opt,name=dropped_invalid_aid,json=droppedInvalidAid,proto3" json:"dropped_invalid_aid,omitempty"`
	DroppedInsuffMem       uint32   `protobuf:"varint,59,opt,name=dropped_insuff_mem,json=droppedInsuffMem,proto3" json:"dropped_insuff_mem,omitempty"`
	DroppedDbError         uint32   `protobuf:"varint,60,opt,name=dropped_db_error,json=droppedDbError,proto3" json:"dropped_db_error,omitempty"`
	DroppedClearWithoutSet uint32   `protobuf:"varint,61,opt,name=dropped_clear_without_set,json=droppedClearWithoutSet,proto3" json:"dropped_clear_without_set,omitempty"`
	DroppedDuplicate       uint32   `protobuf:"varint,62,opt,name=dropped_duplicate,json=droppedDuplicate,proto3" json:"dropped_duplicate,omitempty"`
	CacheHit               uint32   `protobuf:"varint,63,opt,name=cache_hit,json=cacheHit,proto3" json:"cache_hit,omitempty"`
	CacheMiss              uint32   `protobuf:"varint,64,opt,name=cache_miss,json=cacheMiss,proto3" json:"cache_miss,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *AlarmMgrShowAlarmStats) Reset()         { *m = AlarmMgrShowAlarmStats{} }
func (m *AlarmMgrShowAlarmStats) String() string { return proto.CompactTextString(m) }
func (*AlarmMgrShowAlarmStats) ProtoMessage()    {}
func (*AlarmMgrShowAlarmStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a57311bc831b075f, []int{1}
}

func (m *AlarmMgrShowAlarmStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlarmMgrShowAlarmStats.Unmarshal(m, b)
}
func (m *AlarmMgrShowAlarmStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlarmMgrShowAlarmStats.Marshal(b, m, deterministic)
}
func (m *AlarmMgrShowAlarmStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlarmMgrShowAlarmStats.Merge(m, src)
}
func (m *AlarmMgrShowAlarmStats) XXX_Size() int {
	return xxx_messageInfo_AlarmMgrShowAlarmStats.Size(m)
}
func (m *AlarmMgrShowAlarmStats) XXX_DiscardUnknown() {
	xxx_messageInfo_AlarmMgrShowAlarmStats.DiscardUnknown(m)
}

var xxx_messageInfo_AlarmMgrShowAlarmStats proto.InternalMessageInfo

func (m *AlarmMgrShowAlarmStats) GetReported() uint64 {
	if m != nil {
		return m.Reported
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDropped() uint64 {
	if m != nil {
		return m.Dropped
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetActive() uint64 {
	if m != nil {
		return m.Active
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetHistory() uint64 {
	if m != nil {
		return m.History
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetSuppressed() uint64 {
	if m != nil {
		return m.Suppressed
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetSysadminActive() uint64 {
	if m != nil {
		return m.SysadminActive
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetSysadminHistory() uint64 {
	if m != nil {
		return m.SysadminHistory
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetSysadminSuppressed() uint64 {
	if m != nil {
		return m.SysadminSuppressed
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDroppedInvalidAid() uint32 {
	if m != nil {
		return m.DroppedInvalidAid
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDroppedInsuffMem() uint32 {
	if m != nil {
		return m.DroppedInsuffMem
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDroppedDbError() uint32 {
	if m != nil {
		return m.DroppedDbError
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDroppedClearWithoutSet() uint32 {
	if m != nil {
		return m.DroppedClearWithoutSet
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetDroppedDuplicate() uint32 {
	if m != nil {
		return m.DroppedDuplicate
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetCacheHit() uint32 {
	if m != nil {
		return m.CacheHit
	}
	return 0
}

func (m *AlarmMgrShowAlarmStats) GetCacheMiss() uint32 {
	if m != nil {
		return m.CacheMiss
	}
	return 0
}

func init() {
	proto.RegisterType((*AlarmMgrShowAlarmStats_KEYS)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.stats.alarm_mgr_show_alarm_stats_KEYS")
	proto.RegisterType((*AlarmMgrShowAlarmStats)(nil), "cisco_ios_xr_alarmgr_server_oper.alarms.detail.detail_card.detail_locations.detail_location.stats.alarm_mgr_show_alarm_stats")
}

func init() { proto.RegisterFile("alarm_mgr_show_alarm_stats.proto", fileDescriptor_a57311bc831b075f) }

var fileDescriptor_a57311bc831b075f = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5d, 0x4b, 0xdc, 0x40,
	0x14, 0x86, 0x59, 0x90, 0x55, 0x0f, 0xd4, 0xea, 0x08, 0x76, 0x6a, 0x69, 0xbb, 0x78, 0xd3, 0x2d,
	0x2d, 0x29, 0xd4, 0x7e, 0x69, 0x3f, 0xa5, 0x0a, 0x4a, 0xf1, 0x66, 0xf7, 0xa2, 0xf4, 0xea, 0x30,
	0x9b, 0x39, 0x36, 0x07, 0x92, 0x9d, 0x30, 0x67, 0x76, 0xad, 0x3f, 0xb4, 0xff, 0xa7, 0x64, 0x32,
	0x89, 0x56, 0xf0, 0x2a, 0x39, 0xef, 0xf3, 0x9c, 0x79, 0x27, 0x10, 0x18, 0x99, 0xd2, 0xf8, 0x0a,
	0xab, 0xdf, 0x1e, 0xa5, 0x70, 0x97, 0xd8, 0x8e, 0x12, 0x4c, 0x90, 0xac, 0xf6, 0x2e, 0x38, 0x65,
	0x72, 0x96, 0xdc, 0x21, 0x3b, 0xc1, 0x3f, 0xbe, 0xe5, 0x8d, 0x4c, 0x7e, 0x49, 0x1e, 0x5d, 0x4d,
	0x3e, 0x8b, 0x99, 0x64, 0x96, 0x82, 0xe1, 0x32, 0x3d, 0x30, 0x37, 0xde, 0x76, 0xef, 0xa5, 0xcb,
	0x4d, 0x60, 0x37, 0x97, 0xdb, 0x41, 0x16, 0x8b, 0xf6, 0x0e, 0xe1, 0xe9, 0xdd, 0xd7, 0xc0, 0x1f,
	0x27, 0xbf, 0xa6, 0xea, 0x01, 0xac, 0xce, 0x9d, 0x25, 0x64, 0xab, 0x07, 0xa3, 0xc1, 0x78, 0x7d,
	0x32, 0x6c, 0xc6, 0x33, 0xbb, 0xf7, 0x77, 0x05, 0x76, 0xef, 0x5e, 0x56, 0xbb, 0xb0, 0xe6, 0xa9,
	0x76, 0x3e, 0x90, 0xd5, 0xaf, 0x47, 0x83, 0xf1, 0xca, 0xa4, 0x9f, 0x95, 0x86, 0x55, 0xeb, 0x5d,
	0x5d, 0x93, 0xd5, 0xfb, 0x11, 0x75, 0xa3, 0xda, 0x81, 0xa1, 0xc9, 0x03, 0x2f, 0x49, 0xbf, 0x89,
	0x20, 0x4d, 0xcd, 0x46, 0xc1, 0x12, 0x9c, 0xbf, 0xd2, 0x6f, 0xdb, 0x8d, 0x34, 0xaa, 0x27, 0x00,
	0xb2, 0xa8, 0x6b, 0x4f, 0x22, 0x64, 0xf5, 0xbb, 0x08, 0x6f, 0x24, 0xea, 0x19, 0xdc, 0x97, 0x2b,
	0x31, 0xb6, 0xe2, 0x39, 0xa6, 0xa3, 0xdf, 0x47, 0x69, 0xa3, 0x8b, 0x8f, 0xda, 0x8a, 0xe7, 0xb0,
	0xd9, 0x8b, 0x5d, 0xd7, 0x87, 0x68, 0xf6, 0x07, 0x9c, 0xa6, 0xce, 0x57, 0xb0, 0xdd, 0xab, 0x37,
	0xca, 0x0f, 0xa2, 0xad, 0x3a, 0x34, 0xbd, 0xbe, 0x44, 0x06, 0xdb, 0xe9, 0x0b, 0x91, 0xe7, 0x4b,
	0x53, 0xb2, 0x45, 0xc3, 0x56, 0x1f, 0x8e, 0x06, 0xe3, 0x7b, 0x93, 0xad, 0x84, 0xce, 0x5a, 0x72,
	0xc4, 0x56, 0xbd, 0x04, 0x75, 0xed, 0xcb, 0xe2, 0xe2, 0x02, 0x2b, 0xaa, 0xf4, 0xc7, 0xa8, 0x6f,
	0xf6, 0x7a, 0x03, 0xce, 0xa9, 0x52, 0x63, 0xe8, 0x32, 0xb4, 0x33, 0x24, 0xef, 0x9d, 0xd7, 0x9f,
	0xa2, 0xbb, 0x91, 0xf2, 0xe3, 0xd9, 0x49, 0x93, 0xaa, 0x03, 0x78, 0xd8, 0x99, 0x79, 0x49, 0xc6,
	0xe3, 0x25, 0x87, 0xc2, 0x2d, 0x02, 0x0a, 0x05, 0xfd, 0x39, 0xae, 0xec, 0x24, 0xe1, 0x7b, 0xc3,
	0x7f, 0xb6, 0x78, 0x4a, 0x41, 0xbd, 0x80, 0xad, 0xbe, 0x64, 0x51, 0x97, 0x9c, 0x9b, 0x40, 0xfa,
	0xcb, 0x7f, 0x37, 0x3a, 0xee, 0x72, 0xf5, 0x08, 0xd6, 0x73, 0x93, 0x17, 0x84, 0x05, 0x07, 0xfd,
	0x35, 0x4a, 0x6b, 0x31, 0x38, 0xe5, 0xa0, 0x1e, 0x03, 0xb4, 0xb0, 0x62, 0x11, 0xfd, 0x2d, 0xd2,
	0x56, 0x3f, 0x67, 0x91, 0xd9, 0x30, 0xfe, 0xfd, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02,
	0xd1, 0x3c, 0xb1, 0x21, 0x03, 0x00, 0x00,
}
