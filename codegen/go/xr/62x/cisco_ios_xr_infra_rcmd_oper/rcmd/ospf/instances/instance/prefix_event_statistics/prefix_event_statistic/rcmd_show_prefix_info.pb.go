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
// source: rcmd_show_prefix_info.proto

package cisco_ios_xr_infra_rcmd_oper_rcmd_ospf_instances_instance_prefix_event_statistics_prefix_event_statistic

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

type RcmdShowPrefixInfo_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	PrefixInfo           string   `protobuf:"bytes,2,opt,name=prefix_info,json=prefixInfo,proto3" json:"prefix_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowPrefixInfo_KEYS) Reset()         { *m = RcmdShowPrefixInfo_KEYS{} }
func (m *RcmdShowPrefixInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixInfo_KEYS) ProtoMessage()    {}
func (*RcmdShowPrefixInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4675c9eec193e1a, []int{0}
}

func (m *RcmdShowPrefixInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixInfo_KEYS.Unmarshal(m, b)
}
func (m *RcmdShowPrefixInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixInfo_KEYS.Merge(m, src)
}
func (m *RcmdShowPrefixInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixInfo_KEYS.Size(m)
}
func (m *RcmdShowPrefixInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixInfo_KEYS proto.InternalMessageInfo

func (m *RcmdShowPrefixInfo_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *RcmdShowPrefixInfo_KEYS) GetPrefixInfo() string {
	if m != nil {
		return m.PrefixInfo
	}
	return ""
}

type RcmdShowPrefixInfo struct {
	Prefix               string   `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLenth          uint32   `protobuf:"varint,51,opt,name=prefix_lenth,json=prefixLenth,proto3" json:"prefix_lenth,omitempty"`
	LastEventTime        string   `protobuf:"bytes,52,opt,name=last_event_time,json=lastEventTime,proto3" json:"last_event_time,omitempty"`
	LastPriority         string   `protobuf:"bytes,53,opt,name=last_priority,json=lastPriority,proto3" json:"last_priority,omitempty"`
	LastRouteType        string   `protobuf:"bytes,54,opt,name=last_route_type,json=lastRouteType,proto3" json:"last_route_type,omitempty"`
	LastChangeType       string   `protobuf:"bytes,55,opt,name=last_change_type,json=lastChangeType,proto3" json:"last_change_type,omitempty"`
	LastCost             uint32   `protobuf:"varint,56,opt,name=last_cost,json=lastCost,proto3" json:"last_cost,omitempty"`
	CriticalPriority     uint32   `protobuf:"varint,57,opt,name=critical_priority,json=criticalPriority,proto3" json:"critical_priority,omitempty"`
	HighPriority         uint32   `protobuf:"varint,58,opt,name=high_priority,json=highPriority,proto3" json:"high_priority,omitempty"`
	MediumPriority       uint32   `protobuf:"varint,59,opt,name=medium_priority,json=mediumPriority,proto3" json:"medium_priority,omitempty"`
	LowPriority          uint32   `protobuf:"varint,60,opt,name=low_priority,json=lowPriority,proto3" json:"low_priority,omitempty"`
	AddCount             uint32   `protobuf:"varint,61,opt,name=add_count,json=addCount,proto3" json:"add_count,omitempty"`
	ModifyCount          uint32   `protobuf:"varint,62,opt,name=modify_count,json=modifyCount,proto3" json:"modify_count,omitempty"`
	DeleteCount          uint32   `protobuf:"varint,63,opt,name=delete_count,json=deleteCount,proto3" json:"delete_count,omitempty"`
	ThresholdExceedCount uint32   `protobuf:"varint,64,opt,name=threshold_exceed_count,json=thresholdExceedCount,proto3" json:"threshold_exceed_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowPrefixInfo) Reset()         { *m = RcmdShowPrefixInfo{} }
func (m *RcmdShowPrefixInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixInfo) ProtoMessage()    {}
func (*RcmdShowPrefixInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4675c9eec193e1a, []int{1}
}

func (m *RcmdShowPrefixInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixInfo.Unmarshal(m, b)
}
func (m *RcmdShowPrefixInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixInfo.Merge(m, src)
}
func (m *RcmdShowPrefixInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixInfo.Size(m)
}
func (m *RcmdShowPrefixInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixInfo proto.InternalMessageInfo

func (m *RcmdShowPrefixInfo) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RcmdShowPrefixInfo) GetPrefixLenth() uint32 {
	if m != nil {
		return m.PrefixLenth
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetLastEventTime() string {
	if m != nil {
		return m.LastEventTime
	}
	return ""
}

func (m *RcmdShowPrefixInfo) GetLastPriority() string {
	if m != nil {
		return m.LastPriority
	}
	return ""
}

func (m *RcmdShowPrefixInfo) GetLastRouteType() string {
	if m != nil {
		return m.LastRouteType
	}
	return ""
}

func (m *RcmdShowPrefixInfo) GetLastChangeType() string {
	if m != nil {
		return m.LastChangeType
	}
	return ""
}

func (m *RcmdShowPrefixInfo) GetLastCost() uint32 {
	if m != nil {
		return m.LastCost
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetCriticalPriority() uint32 {
	if m != nil {
		return m.CriticalPriority
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetHighPriority() uint32 {
	if m != nil {
		return m.HighPriority
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetMediumPriority() uint32 {
	if m != nil {
		return m.MediumPriority
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetLowPriority() uint32 {
	if m != nil {
		return m.LowPriority
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetAddCount() uint32 {
	if m != nil {
		return m.AddCount
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetModifyCount() uint32 {
	if m != nil {
		return m.ModifyCount
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetDeleteCount() uint32 {
	if m != nil {
		return m.DeleteCount
	}
	return 0
}

func (m *RcmdShowPrefixInfo) GetThresholdExceedCount() uint32 {
	if m != nil {
		return m.ThresholdExceedCount
	}
	return 0
}

func init() {
	proto.RegisterType((*RcmdShowPrefixInfo_KEYS)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.prefix_event_statistics.prefix_event_statistic.rcmd_show_prefix_info_KEYS")
	proto.RegisterType((*RcmdShowPrefixInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.prefix_event_statistics.prefix_event_statistic.rcmd_show_prefix_info")
}

func init() { proto.RegisterFile("rcmd_show_prefix_info.proto", fileDescriptor_d4675c9eec193e1a) }

var fileDescriptor_d4675c9eec193e1a = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x4b, 0x6b, 0x1b, 0x31,
	0x10, 0x80, 0x71, 0x29, 0xa1, 0x56, 0xed, 0x24, 0x5d, 0xda, 0x20, 0xea, 0x43, 0xf3, 0x80, 0xd6,
	0x50, 0xd8, 0x43, 0x93, 0xbe, 0x9f, 0x10, 0x7c, 0x28, 0x2d, 0xa5, 0xb8, 0xb9, 0xf4, 0x24, 0x14,
	0x69, 0x36, 0x2b, 0xd8, 0x95, 0x16, 0x69, 0xdc, 0xd8, 0x3f, 0xb0, 0xff, 0xab, 0x68, 0x24, 0xef,
	0xf6, 0x90, 0xdc, 0x34, 0xdf, 0x7c, 0xf3, 0xb2, 0x59, 0x36, 0xf3, 0xaa, 0xd5, 0x22, 0xd4, 0xee,
	0x5a, 0x74, 0x1e, 0x2a, 0xb3, 0x16, 0xc6, 0x56, 0xae, 0xec, 0xbc, 0x43, 0x57, 0xd4, 0xca, 0x04,
	0xe5, 0x84, 0x71, 0x41, 0xac, 0x7d, 0x4c, 0x78, 0x29, 0xc8, 0x77, 0x1d, 0xf8, 0x32, 0xbe, 0x4a,
	0x17, 0xba, 0xaa, 0x34, 0x36, 0xa0, 0xb4, 0x0a, 0x42, 0xff, 0x2a, 0x73, 0x33, 0xf8, 0x03, 0x16,
	0x45, 0x40, 0x89, 0x26, 0xa0, 0x51, 0xe1, 0x16, 0x7e, 0x7c, 0xc9, 0x1e, 0xdf, 0xb8, 0x88, 0xf8,
	0xb6, 0xf8, 0xfd, 0xab, 0x38, 0x61, 0xd3, 0x6d, 0x63, 0x61, 0x65, 0x0b, 0x7c, 0x74, 0x38, 0x9a,
	0x8f, 0x97, 0x93, 0x2d, 0xfc, 0x21, 0x5b, 0x28, 0x9e, 0xb0, 0xfb, 0xff, 0x15, 0xf2, 0x3b, 0xa4,
	0xb0, 0x84, 0xbe, 0xda, 0xca, 0x1d, 0xff, 0xbd, 0xcb, 0x1e, 0xdd, 0x38, 0xa4, 0x38, 0x60, 0x3b,
	0x29, 0xe4, 0x2f, 0xa8, 0x2a, 0x47, 0xc5, 0x11, 0x9b, 0x64, 0xad, 0x01, 0x8b, 0x35, 0x3f, 0x3d,
	0x1c, 0xcd, 0xa7, 0xcb, 0x3c, 0xe6, 0x7b, 0x44, 0xc5, 0x53, 0xb6, 0xd7, 0xc8, 0x80, 0xf9, 0x20,
	0x34, 0x2d, 0xf0, 0x33, 0xea, 0x31, 0x8d, 0x78, 0x11, 0xe9, 0x85, 0x69, 0x21, 0x9e, 0x40, 0x5e,
	0xe7, 0x8d, 0xf3, 0x06, 0x37, 0xfc, 0x65, 0x3a, 0x21, 0xc2, 0x9f, 0x99, 0xf5, 0xcd, 0xbc, 0x5b,
	0x21, 0x08, 0xdc, 0x74, 0xc0, 0x5f, 0x0d, 0xcd, 0x96, 0x91, 0x5e, 0x6c, 0x3a, 0x28, 0xe6, 0x6c,
	0x9f, 0x3c, 0x55, 0x4b, 0x7b, 0x95, 0xc5, 0xd7, 0x24, 0xee, 0x46, 0x7e, 0x4e, 0x98, 0xcc, 0x19,
	0x1b, 0x27, 0xd3, 0x05, 0xe4, 0x6f, 0x68, 0xfd, 0x7b, 0xa4, 0xb8, 0x80, 0xc5, 0x73, 0xf6, 0x40,
	0x79, 0x83, 0x46, 0xc9, 0x66, 0xd8, 0xeb, 0x2d, 0x49, 0xfb, 0xdb, 0x44, 0xbf, 0xdb, 0x09, 0x9b,
	0xd6, 0xe6, 0xaa, 0x1e, 0xc4, 0x77, 0x24, 0x4e, 0x22, 0xec, 0xa5, 0x67, 0x6c, 0xaf, 0x05, 0x6d,
	0x56, 0xed, 0xa0, 0xbd, 0x27, 0x6d, 0x37, 0xe1, 0x5e, 0x3c, 0x62, 0x93, 0x86, 0xfe, 0x84, 0x6c,
	0x7d, 0x48, 0xbf, 0x6c, 0xe3, 0xae, 0x7b, 0x65, 0xc6, 0xc6, 0x52, 0x6b, 0xa1, 0xdc, 0xca, 0x22,
	0xff, 0x98, 0x56, 0x97, 0x5a, 0x9f, 0xc7, 0x38, 0xd6, 0xb7, 0x4e, 0x9b, 0x6a, 0x93, 0xf3, 0x9f,
	0x52, 0x7d, 0x62, 0xbd, 0xa2, 0xa1, 0x01, 0x84, 0xac, 0x7c, 0x4e, 0x4a, 0x62, 0x49, 0x39, 0x63,
	0x07, 0x58, 0x7b, 0x08, 0xb5, 0x6b, 0xb4, 0x80, 0xb5, 0x02, 0xd8, 0xce, 0xfb, 0x42, 0xf2, 0xc3,
	0x3e, 0xbb, 0xa0, 0x24, 0x55, 0x5d, 0xee, 0xd0, 0xc7, 0x71, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0x59, 0xec, 0x38, 0xe8, 0x3b, 0x03, 0x00, 0x00,
}
