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
// source: rcmd_show_ipfrr_event_info.proto

package cisco_ios_xr_infra_rcmd_oper_rcmd_isis_instances_instance_ipfrr_event_offlines_ipfrr_event_offline

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

type RcmdShowIpfrrEventInfo_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	EventId              uint32   `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowIpfrrEventInfo_KEYS) Reset()         { *m = RcmdShowIpfrrEventInfo_KEYS{} }
func (m *RcmdShowIpfrrEventInfo_KEYS) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrEventInfo_KEYS) ProtoMessage()    {}
func (*RcmdShowIpfrrEventInfo_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c887fad56114e97, []int{0}
}

func (m *RcmdShowIpfrrEventInfo_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrEventInfo_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrEventInfo_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS.Merge(m, src)
}
func (m *RcmdShowIpfrrEventInfo_KEYS) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS.Size(m)
}
func (m *RcmdShowIpfrrEventInfo_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrEventInfo_KEYS proto.InternalMessageInfo

func (m *RcmdShowIpfrrEventInfo_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *RcmdShowIpfrrEventInfo_KEYS) GetEventId() uint32 {
	if m != nil {
		return m.EventId
	}
	return 0
}

type RcmdShowIpfrrStats struct {
	Priority                 string   `protobuf:"bytes,1,opt,name=priority,proto3" json:"priority,omitempty"`
	TotalRoutes              uint32   `protobuf:"varint,2,opt,name=total_routes,json=totalRoutes,proto3" json:"total_routes,omitempty"`
	FullyProtectedRoutes     uint32   `protobuf:"varint,3,opt,name=fully_protected_routes,json=fullyProtectedRoutes,proto3" json:"fully_protected_routes,omitempty"`
	PartiallyProtectedRoutes uint32   `protobuf:"varint,4,opt,name=partially_protected_routes,json=partiallyProtectedRoutes,proto3" json:"partially_protected_routes,omitempty"`
	Coverage                 string   `protobuf:"bytes,5,opt,name=coverage,proto3" json:"coverage,omitempty"`
	LocalLfaCoverage         string   `protobuf:"bytes,6,opt,name=local_lfa_coverage,json=localLfaCoverage,proto3" json:"local_lfa_coverage,omitempty"`
	RemoteLfaCoverage        string   `protobuf:"bytes,7,opt,name=remote_lfa_coverage,json=remoteLfaCoverage,proto3" json:"remote_lfa_coverage,omitempty"`
	BelowThreshold           bool     `protobuf:"varint,8,opt,name=below_threshold,json=belowThreshold,proto3" json:"below_threshold,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *RcmdShowIpfrrStats) Reset()         { *m = RcmdShowIpfrrStats{} }
func (m *RcmdShowIpfrrStats) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrStats) ProtoMessage()    {}
func (*RcmdShowIpfrrStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c887fad56114e97, []int{1}
}

func (m *RcmdShowIpfrrStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrStats.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrStats.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrStats.Merge(m, src)
}
func (m *RcmdShowIpfrrStats) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrStats.Size(m)
}
func (m *RcmdShowIpfrrStats) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrStats.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrStats proto.InternalMessageInfo

func (m *RcmdShowIpfrrStats) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *RcmdShowIpfrrStats) GetTotalRoutes() uint32 {
	if m != nil {
		return m.TotalRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrStats) GetFullyProtectedRoutes() uint32 {
	if m != nil {
		return m.FullyProtectedRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrStats) GetPartiallyProtectedRoutes() uint32 {
	if m != nil {
		return m.PartiallyProtectedRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrStats) GetCoverage() string {
	if m != nil {
		return m.Coverage
	}
	return ""
}

func (m *RcmdShowIpfrrStats) GetLocalLfaCoverage() string {
	if m != nil {
		return m.LocalLfaCoverage
	}
	return ""
}

func (m *RcmdShowIpfrrStats) GetRemoteLfaCoverage() string {
	if m != nil {
		return m.RemoteLfaCoverage
	}
	return ""
}

func (m *RcmdShowIpfrrStats) GetBelowThreshold() bool {
	if m != nil {
		return m.BelowThreshold
	}
	return false
}

type RcmdShowIpfrrPath struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighbourAddress     string   `protobuf:"bytes,2,opt,name=neighbour_address,json=neighbourAddress,proto3" json:"neighbour_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowIpfrrPath) Reset()         { *m = RcmdShowIpfrrPath{} }
func (m *RcmdShowIpfrrPath) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrPath) ProtoMessage()    {}
func (*RcmdShowIpfrrPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c887fad56114e97, []int{2}
}

func (m *RcmdShowIpfrrPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrPath.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrPath.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrPath.Merge(m, src)
}
func (m *RcmdShowIpfrrPath) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrPath.Size(m)
}
func (m *RcmdShowIpfrrPath) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrPath.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrPath proto.InternalMessageInfo

func (m *RcmdShowIpfrrPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RcmdShowIpfrrPath) GetNeighbourAddress() string {
	if m != nil {
		return m.NeighbourAddress
	}
	return ""
}

type RcmdShowIpfrrRlfaPqNode struct {
	RemoteNodeId         string               `protobuf:"bytes,1,opt,name=remote_node_id,json=remoteNodeId,proto3" json:"remote_node_id,omitempty"`
	InterfaceName        string               `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighbourAddress     string               `protobuf:"bytes,3,opt,name=neighbour_address,json=neighbourAddress,proto3" json:"neighbour_address,omitempty"`
	PathCount            uint32               `protobuf:"varint,4,opt,name=path_count,json=pathCount,proto3" json:"path_count,omitempty"`
	InUseTime            string               `protobuf:"bytes,5,opt,name=in_use_time,json=inUseTime,proto3" json:"in_use_time,omitempty"`
	PrimaryPath          []*RcmdShowIpfrrPath `protobuf:"bytes,6,rep,name=primary_path,json=primaryPath,proto3" json:"primary_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RcmdShowIpfrrRlfaPqNode) Reset()         { *m = RcmdShowIpfrrRlfaPqNode{} }
func (m *RcmdShowIpfrrRlfaPqNode) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrRlfaPqNode) ProtoMessage()    {}
func (*RcmdShowIpfrrRlfaPqNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c887fad56114e97, []int{3}
}

func (m *RcmdShowIpfrrRlfaPqNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrRlfaPqNode.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrRlfaPqNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrRlfaPqNode.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrRlfaPqNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrRlfaPqNode.Merge(m, src)
}
func (m *RcmdShowIpfrrRlfaPqNode) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrRlfaPqNode.Size(m)
}
func (m *RcmdShowIpfrrRlfaPqNode) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrRlfaPqNode.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrRlfaPqNode proto.InternalMessageInfo

func (m *RcmdShowIpfrrRlfaPqNode) GetRemoteNodeId() string {
	if m != nil {
		return m.RemoteNodeId
	}
	return ""
}

func (m *RcmdShowIpfrrRlfaPqNode) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RcmdShowIpfrrRlfaPqNode) GetNeighbourAddress() string {
	if m != nil {
		return m.NeighbourAddress
	}
	return ""
}

func (m *RcmdShowIpfrrRlfaPqNode) GetPathCount() uint32 {
	if m != nil {
		return m.PathCount
	}
	return 0
}

func (m *RcmdShowIpfrrRlfaPqNode) GetInUseTime() string {
	if m != nil {
		return m.InUseTime
	}
	return ""
}

func (m *RcmdShowIpfrrRlfaPqNode) GetPrimaryPath() []*RcmdShowIpfrrPath {
	if m != nil {
		return m.PrimaryPath
	}
	return nil
}

type RcmdShowIpfrrEventInfo struct {
	EventIdXr                uint32                     `protobuf:"varint,50,opt,name=event_id_xr,json=eventIdXr,proto3" json:"event_id_xr,omitempty"`
	TriggerTime              string                     `protobuf:"bytes,51,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	TriggerSpfRun            uint32                     `protobuf:"varint,52,opt,name=trigger_spf_run,json=triggerSpfRun,proto3" json:"trigger_spf_run,omitempty"`
	WaitTime                 uint32                     `protobuf:"varint,53,opt,name=wait_time,json=waitTime,proto3" json:"wait_time,omitempty"`
	StartTimeOffset          string                     `protobuf:"bytes,54,opt,name=start_time_offset,json=startTimeOffset,proto3" json:"start_time_offset,omitempty"`
	Duration                 string                     `protobuf:"bytes,55,opt,name=duration,proto3" json:"duration,omitempty"`
	CompletedSpfRun          uint32                     `protobuf:"varint,56,opt,name=completed_spf_run,json=completedSpfRun,proto3" json:"completed_spf_run,omitempty"`
	TotalRoutes              uint32                     `protobuf:"varint,57,opt,name=total_routes,json=totalRoutes,proto3" json:"total_routes,omitempty"`
	FullyProtectedRoutes     uint32                     `protobuf:"varint,58,opt,name=fully_protected_routes,json=fullyProtectedRoutes,proto3" json:"fully_protected_routes,omitempty"`
	PartiallyProtectedRoutes uint32                     `protobuf:"varint,59,opt,name=partially_protected_routes,json=partiallyProtectedRoutes,proto3" json:"partially_protected_routes,omitempty"`
	Coverage                 string                     `protobuf:"bytes,60,opt,name=coverage,proto3" json:"coverage,omitempty"`
	IpfrrStatistic           []*RcmdShowIpfrrStats      `protobuf:"bytes,61,rep,name=ipfrr_statistic,json=ipfrrStatistic,proto3" json:"ipfrr_statistic,omitempty"`
	RemoteNode               []*RcmdShowIpfrrRlfaPqNode `protobuf:"bytes,62,rep,name=remote_node,json=remoteNode,proto3" json:"remote_node,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *RcmdShowIpfrrEventInfo) Reset()         { *m = RcmdShowIpfrrEventInfo{} }
func (m *RcmdShowIpfrrEventInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrEventInfo) ProtoMessage()    {}
func (*RcmdShowIpfrrEventInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c887fad56114e97, []int{4}
}

func (m *RcmdShowIpfrrEventInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrEventInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrEventInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrEventInfo.Merge(m, src)
}
func (m *RcmdShowIpfrrEventInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrEventInfo.Size(m)
}
func (m *RcmdShowIpfrrEventInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrEventInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrEventInfo proto.InternalMessageInfo

func (m *RcmdShowIpfrrEventInfo) GetEventIdXr() uint32 {
	if m != nil {
		return m.EventIdXr
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetTriggerTime() string {
	if m != nil {
		return m.TriggerTime
	}
	return ""
}

func (m *RcmdShowIpfrrEventInfo) GetTriggerSpfRun() uint32 {
	if m != nil {
		return m.TriggerSpfRun
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetWaitTime() uint32 {
	if m != nil {
		return m.WaitTime
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetStartTimeOffset() string {
	if m != nil {
		return m.StartTimeOffset
	}
	return ""
}

func (m *RcmdShowIpfrrEventInfo) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *RcmdShowIpfrrEventInfo) GetCompletedSpfRun() uint32 {
	if m != nil {
		return m.CompletedSpfRun
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetTotalRoutes() uint32 {
	if m != nil {
		return m.TotalRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetFullyProtectedRoutes() uint32 {
	if m != nil {
		return m.FullyProtectedRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetPartiallyProtectedRoutes() uint32 {
	if m != nil {
		return m.PartiallyProtectedRoutes
	}
	return 0
}

func (m *RcmdShowIpfrrEventInfo) GetCoverage() string {
	if m != nil {
		return m.Coverage
	}
	return ""
}

func (m *RcmdShowIpfrrEventInfo) GetIpfrrStatistic() []*RcmdShowIpfrrStats {
	if m != nil {
		return m.IpfrrStatistic
	}
	return nil
}

func (m *RcmdShowIpfrrEventInfo) GetRemoteNode() []*RcmdShowIpfrrRlfaPqNode {
	if m != nil {
		return m.RemoteNode
	}
	return nil
}

func init() {
	proto.RegisterType((*RcmdShowIpfrrEventInfo_KEYS)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.isis.instances.instance.ipfrr_event_offlines.ipfrr_event_offline.rcmd_show_ipfrr_event_info_KEYS")
	proto.RegisterType((*RcmdShowIpfrrStats)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.isis.instances.instance.ipfrr_event_offlines.ipfrr_event_offline.rcmd_show_ipfrr_stats")
	proto.RegisterType((*RcmdShowIpfrrPath)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.isis.instances.instance.ipfrr_event_offlines.ipfrr_event_offline.rcmd_show_ipfrr_path")
	proto.RegisterType((*RcmdShowIpfrrRlfaPqNode)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.isis.instances.instance.ipfrr_event_offlines.ipfrr_event_offline.rcmd_show_ipfrr_rlfa_pq_node")
	proto.RegisterType((*RcmdShowIpfrrEventInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.isis.instances.instance.ipfrr_event_offlines.ipfrr_event_offline.rcmd_show_ipfrr_event_info")
}

func init() { proto.RegisterFile("rcmd_show_ipfrr_event_info.proto", fileDescriptor_9c887fad56114e97) }

var fileDescriptor_9c887fad56114e97 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x96, 0xd3, 0xf7, 0x4d, 0x93, 0x4d, 0x93, 0xb4, 0xfb, 0xf6, 0x45, 0x26, 0x7c, 0x85, 0xf0,
	0x15, 0x01, 0xf2, 0xa1, 0x2d, 0xdf, 0x05, 0x09, 0x55, 0x1c, 0x2a, 0x50, 0xa9, 0xdc, 0x22, 0xc1,
	0x69, 0xb5, 0xb1, 0xd7, 0xc9, 0x22, 0x7b, 0xd7, 0xec, 0xae, 0xdb, 0xe6, 0xc6, 0x9f, 0xe0, 0xc4,
	0x8d, 0xff, 0xc7, 0x91, 0x3b, 0xda, 0xb1, 0x9d, 0xa6, 0x21, 0x95, 0x28, 0x12, 0xea, 0xcd, 0xf3,
	0xcc, 0x33, 0xb3, 0xcf, 0xec, 0xec, 0x8c, 0x51, 0x57, 0x05, 0x49, 0x48, 0xf4, 0x48, 0x1e, 0x12,
	0x9e, 0x46, 0x4a, 0x11, 0x76, 0xc0, 0x84, 0x21, 0x5c, 0x44, 0xd2, 0x4b, 0x95, 0x34, 0x12, 0x0f,
	0x02, 0xae, 0x03, 0x49, 0xb8, 0xd4, 0xe4, 0x48, 0x59, 0x87, 0xa2, 0x04, 0x82, 0x64, 0xca, 0x94,
	0x67, 0xbf, 0x3c, 0xae, 0xb9, 0xf6, 0xb8, 0xd0, 0x86, 0x8a, 0x80, 0x1d, 0x7f, 0x79, 0xd3, 0x19,
	0x65, 0x14, 0xc5, 0x5c, 0x58, 0xef, 0xaf, 0x60, 0x8f, 0xa2, 0x6b, 0xa7, 0xeb, 0x20, 0xaf, 0x5f,
	0x7d, 0xd8, 0xc3, 0x37, 0x50, 0xb3, 0xcc, 0x4b, 0x04, 0x4d, 0x98, 0xeb, 0x74, 0x9d, 0x7e, 0xdd,
	0x5f, 0x2a, 0xc1, 0x1d, 0x9a, 0x30, 0x7c, 0x11, 0xd5, 0x8a, 0xb8, 0xd0, 0xad, 0x74, 0x9d, 0x7e,
	0xd3, 0x5f, 0x04, 0x7b, 0x3b, 0xec, 0x7d, 0xaf, 0xa0, 0xff, 0x67, 0xcf, 0xd0, 0x86, 0x1a, 0x8d,
	0x3b, 0xa8, 0x96, 0x2a, 0x2e, 0x15, 0x37, 0xe3, 0x22, 0xe9, 0xc4, 0xc6, 0xd7, 0xd1, 0x92, 0x91,
	0x86, 0xc6, 0x44, 0xc9, 0xcc, 0x30, 0x5d, 0x24, 0x6d, 0x00, 0xe6, 0x03, 0x84, 0x37, 0xd0, 0x85,
	0x28, 0x8b, 0xe3, 0x31, 0xb1, 0xd7, 0xc5, 0x02, 0xc3, 0xc2, 0x92, 0xbc, 0x00, 0xe4, 0x55, 0xf0,
	0xee, 0x96, 0xce, 0x22, 0x6a, 0x13, 0x75, 0x52, 0xaa, 0x0c, 0xa7, 0x73, 0x23, 0xff, 0x81, 0x48,
	0x77, 0xc2, 0x98, 0x8d, 0xee, 0xa0, 0x5a, 0x20, 0x0f, 0x98, 0xa2, 0x43, 0xe6, 0xfe, 0x9b, 0x4b,
	0x2e, 0x6d, 0x7c, 0x1f, 0xe1, 0x58, 0x06, 0x34, 0x26, 0x71, 0x44, 0xc9, 0x84, 0x55, 0x05, 0xd6,
	0x32, 0x78, 0xde, 0x44, 0x74, 0xab, 0x64, 0x7b, 0xe8, 0x3f, 0xc5, 0x12, 0x69, 0xd8, 0x49, 0xfa,
	0x22, 0xd0, 0x57, 0x72, 0xd7, 0x34, 0xff, 0x0e, 0x6a, 0x0f, 0x58, 0x2c, 0x0f, 0x89, 0x19, 0x29,
	0xa6, 0x47, 0x32, 0x0e, 0xdd, 0x5a, 0xd7, 0xe9, 0xd7, 0xfc, 0x16, 0xc0, 0xfb, 0x25, 0xda, 0xfb,
	0x88, 0x56, 0x67, 0xaf, 0x3b, 0xa5, 0x66, 0x84, 0x6f, 0xa1, 0x16, 0x17, 0x86, 0xa9, 0x88, 0x9e,
	0x6c, 0x64, 0x73, 0x82, 0x42, 0x27, 0xef, 0xa1, 0x15, 0xc1, 0xf8, 0x70, 0x34, 0x90, 0x99, 0x22,
	0x34, 0x0c, 0x15, 0xd3, 0xf9, 0xed, 0xd7, 0xfd, 0xe5, 0x89, 0xe3, 0x65, 0x8e, 0xf7, 0x7e, 0x54,
	0xd0, 0xe5, 0xd9, 0xc3, 0x94, 0xad, 0x27, 0xfd, 0x44, 0x84, 0x0c, 0x19, 0xbe, 0x89, 0x5a, 0x45,
	0x95, 0xd6, 0xb4, 0xaf, 0xa3, 0x78, 0x3d, 0x39, 0xba, 0x23, 0x43, 0xb6, 0x1d, 0xce, 0x91, 0x56,
	0xf9, 0x6d, 0x69, 0x0b, 0xf3, 0xa5, 0xe1, 0x2b, 0x08, 0xd9, 0xb2, 0x49, 0x20, 0x33, 0x61, 0x8a,
	0xbe, 0xd6, 0x2d, 0xb2, 0x65, 0x01, 0x7c, 0x15, 0x35, 0xb8, 0x20, 0x99, 0x66, 0xc4, 0xf0, 0xa4,
	0xec, 0x65, 0x9d, 0x8b, 0x77, 0x9a, 0xed, 0xf3, 0x84, 0xe1, 0xaf, 0x0e, 0x5a, 0x4a, 0x15, 0x4f,
	0xa8, 0x1a, 0xc3, 0xf5, 0xb9, 0xd5, 0xee, 0x42, 0xbf, 0xb1, 0x76, 0xe4, 0xfd, 0xfd, 0xa1, 0xf4,
	0xe6, 0xb5, 0xcf, 0x6f, 0x14, 0x6a, 0x76, 0xa9, 0x19, 0xf5, 0xbe, 0x54, 0x51, 0xe7, 0xf4, 0xb9,
	0xb5, 0xc5, 0x95, 0xd3, 0x48, 0x8e, 0x94, 0xbb, 0x96, 0x17, 0x5f, 0x0c, 0xe4, 0x7b, 0x05, 0xc3,
	0xa5, 0xf8, 0x70, 0xc8, 0x54, 0x5e, 0xfd, 0x3a, 0x54, 0xdf, 0x28, 0x30, 0xa8, 0xff, 0x36, 0x6a,
	0x97, 0x14, 0x9d, 0x46, 0x44, 0x65, 0xc2, 0xdd, 0x80, 0x34, 0xcd, 0x02, 0xde, 0x4b, 0x23, 0x3f,
	0x13, 0xf8, 0x12, 0xaa, 0x1f, 0x52, 0x6e, 0xf2, 0x3c, 0x0f, 0x80, 0x51, 0xb3, 0x00, 0x24, 0xb9,
	0x8b, 0x56, 0xb4, 0xa1, 0x2a, 0xf7, 0xda, 0xf2, 0x34, 0x33, 0xee, 0x43, 0x38, 0xac, 0x0d, 0x0e,
	0xcb, 0x7a, 0x0b, 0xb0, 0x9d, 0xac, 0x30, 0x53, 0xd4, 0x70, 0x29, 0xdc, 0x47, 0xf9, 0x64, 0x95,
	0xb6, 0xcd, 0x13, 0xc8, 0x24, 0x8d, 0x99, 0x9d, 0xd4, 0x52, 0xce, 0x63, 0x38, 0xac, 0x3d, 0x71,
	0x14, 0x82, 0x66, 0x17, 0xc7, 0x93, 0xb3, 0x2c, 0x8e, 0xa7, 0x7f, 0xbc, 0x38, 0x9e, 0x9d, 0x61,
	0x71, 0x6c, 0xce, 0x2c, 0x8e, 0x6f, 0x0e, 0x6a, 0x1f, 0xef, 0x45, 0xae, 0x0d, 0x0f, 0xdc, 0xe7,
	0xf0, 0xdc, 0xc6, 0xe7, 0xf1, 0xdc, 0x60, 0x39, 0xfb, 0x2d, 0x30, 0xf6, 0x4a, 0x41, 0x56, 0x64,
	0x63, 0x6a, 0x94, 0xdd, 0x17, 0x20, 0xf0, 0xb3, 0x73, 0x1e, 0x0a, 0xa7, 0x57, 0x8c, 0x8f, 0x8e,
	0x57, 0xc9, 0xa0, 0x0a, 0x7f, 0xce, 0xf5, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5c, 0x59, 0x2f,
	0xae, 0x5d, 0x07, 0x00, 0x00,
}