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
// source: rcmd_show_prefix_event.proto

package cisco_ios_xr_infra_rcmd_oper_rcmd_ospf_instances_instance_summary_external_event_offlines_summary_external_event_offline

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

type RcmdShowPrefixEvent_KEYS struct {
	InstanceName         string   `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	EventId              uint32   `protobuf:"varint,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowPrefixEvent_KEYS) Reset()         { *m = RcmdShowPrefixEvent_KEYS{} }
func (m *RcmdShowPrefixEvent_KEYS) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixEvent_KEYS) ProtoMessage()    {}
func (*RcmdShowPrefixEvent_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{0}
}

func (m *RcmdShowPrefixEvent_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixEvent_KEYS.Unmarshal(m, b)
}
func (m *RcmdShowPrefixEvent_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixEvent_KEYS.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixEvent_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixEvent_KEYS.Merge(m, src)
}
func (m *RcmdShowPrefixEvent_KEYS) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixEvent_KEYS.Size(m)
}
func (m *RcmdShowPrefixEvent_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixEvent_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixEvent_KEYS proto.InternalMessageInfo

func (m *RcmdShowPrefixEvent_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *RcmdShowPrefixEvent_KEYS) GetEventId() uint32 {
	if m != nil {
		return m.EventId
	}
	return 0
}

type RcmdShowIpfrrLfaPath struct {
	LfaType              string   `protobuf:"bytes,1,opt,name=lfa_type,json=lfaType,proto3" json:"lfa_type,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighbourAddress     string   `protobuf:"bytes,3,opt,name=neighbour_address,json=neighbourAddress,proto3" json:"neighbour_address,omitempty"`
	ChangeType           string   `protobuf:"bytes,4,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	PathMetric           uint32   `protobuf:"varint,5,opt,name=path_metric,json=pathMetric,proto3" json:"path_metric,omitempty"`
	RemoteNodeId         string   `protobuf:"bytes,6,opt,name=remote_node_id,json=remoteNodeId,proto3" json:"remote_node_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowIpfrrLfaPath) Reset()         { *m = RcmdShowIpfrrLfaPath{} }
func (m *RcmdShowIpfrrLfaPath) String() string { return proto.CompactTextString(m) }
func (*RcmdShowIpfrrLfaPath) ProtoMessage()    {}
func (*RcmdShowIpfrrLfaPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{1}
}

func (m *RcmdShowIpfrrLfaPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowIpfrrLfaPath.Unmarshal(m, b)
}
func (m *RcmdShowIpfrrLfaPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowIpfrrLfaPath.Marshal(b, m, deterministic)
}
func (m *RcmdShowIpfrrLfaPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowIpfrrLfaPath.Merge(m, src)
}
func (m *RcmdShowIpfrrLfaPath) XXX_Size() int {
	return xxx_messageInfo_RcmdShowIpfrrLfaPath.Size(m)
}
func (m *RcmdShowIpfrrLfaPath) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowIpfrrLfaPath.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowIpfrrLfaPath proto.InternalMessageInfo

func (m *RcmdShowIpfrrLfaPath) GetLfaType() string {
	if m != nil {
		return m.LfaType
	}
	return ""
}

func (m *RcmdShowIpfrrLfaPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RcmdShowIpfrrLfaPath) GetNeighbourAddress() string {
	if m != nil {
		return m.NeighbourAddress
	}
	return ""
}

func (m *RcmdShowIpfrrLfaPath) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *RcmdShowIpfrrLfaPath) GetPathMetric() uint32 {
	if m != nil {
		return m.PathMetric
	}
	return 0
}

func (m *RcmdShowIpfrrLfaPath) GetRemoteNodeId() string {
	if m != nil {
		return m.RemoteNodeId
	}
	return ""
}

type RcmdShowPrefixPath struct {
	InterfaceName        string                  `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	NeighbourAddress     string                  `protobuf:"bytes,2,opt,name=neighbour_address,json=neighbourAddress,proto3" json:"neighbour_address,omitempty"`
	ChangeType           string                  `protobuf:"bytes,3,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	PathMetric           uint32                  `protobuf:"varint,4,opt,name=path_metric,json=pathMetric,proto3" json:"path_metric,omitempty"`
	LfaPath              []*RcmdShowIpfrrLfaPath `protobuf:"bytes,5,rep,name=lfa_path,json=lfaPath,proto3" json:"lfa_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RcmdShowPrefixPath) Reset()         { *m = RcmdShowPrefixPath{} }
func (m *RcmdShowPrefixPath) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixPath) ProtoMessage()    {}
func (*RcmdShowPrefixPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{2}
}

func (m *RcmdShowPrefixPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixPath.Unmarshal(m, b)
}
func (m *RcmdShowPrefixPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixPath.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixPath.Merge(m, src)
}
func (m *RcmdShowPrefixPath) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixPath.Size(m)
}
func (m *RcmdShowPrefixPath) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixPath.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixPath proto.InternalMessageInfo

func (m *RcmdShowPrefixPath) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *RcmdShowPrefixPath) GetNeighbourAddress() string {
	if m != nil {
		return m.NeighbourAddress
	}
	return ""
}

func (m *RcmdShowPrefixPath) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *RcmdShowPrefixPath) GetPathMetric() uint32 {
	if m != nil {
		return m.PathMetric
	}
	return 0
}

func (m *RcmdShowPrefixPath) GetLfaPath() []*RcmdShowIpfrrLfaPath {
	if m != nil {
		return m.LfaPath
	}
	return nil
}

type RcmdLsaInfo struct {
	LsaId                string   `protobuf:"bytes,1,opt,name=lsa_id,json=lsaId,proto3" json:"lsa_id,omitempty"`
	SequenceNumber       string   `protobuf:"bytes,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	LsaType              string   `protobuf:"bytes,3,opt,name=lsa_type,json=lsaType,proto3" json:"lsa_type,omitempty"`
	OriginRouterId       string   `protobuf:"bytes,4,opt,name=origin_router_id,json=originRouterId,proto3" json:"origin_router_id,omitempty"`
	ChangeType           string   `protobuf:"bytes,5,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	ReceptionTime        string   `protobuf:"bytes,6,opt,name=reception_time,json=receptionTime,proto3" json:"reception_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdLsaInfo) Reset()         { *m = RcmdLsaInfo{} }
func (m *RcmdLsaInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdLsaInfo) ProtoMessage()    {}
func (*RcmdLsaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{3}
}

func (m *RcmdLsaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdLsaInfo.Unmarshal(m, b)
}
func (m *RcmdLsaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdLsaInfo.Marshal(b, m, deterministic)
}
func (m *RcmdLsaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdLsaInfo.Merge(m, src)
}
func (m *RcmdLsaInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdLsaInfo.Size(m)
}
func (m *RcmdLsaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdLsaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdLsaInfo proto.InternalMessageInfo

func (m *RcmdLsaInfo) GetLsaId() string {
	if m != nil {
		return m.LsaId
	}
	return ""
}

func (m *RcmdLsaInfo) GetSequenceNumber() string {
	if m != nil {
		return m.SequenceNumber
	}
	return ""
}

func (m *RcmdLsaInfo) GetLsaType() string {
	if m != nil {
		return m.LsaType
	}
	return ""
}

func (m *RcmdLsaInfo) GetOriginRouterId() string {
	if m != nil {
		return m.OriginRouterId
	}
	return ""
}

func (m *RcmdLsaInfo) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *RcmdLsaInfo) GetReceptionTime() string {
	if m != nil {
		return m.ReceptionTime
	}
	return ""
}

type RcmdTimestamp struct {
	MinimumTime          string   `protobuf:"bytes,1,opt,name=minimum_time,json=minimumTime,proto3" json:"minimum_time,omitempty"`
	MaximumTime          string   `protobuf:"bytes,2,opt,name=maximum_time,json=maximumTime,proto3" json:"maximum_time,omitempty"`
	SlowestNodeName      string   `protobuf:"bytes,3,opt,name=slowest_node_name,json=slowestNodeName,proto3" json:"slowest_node_name,omitempty"`
	FastestNodeName      string   `protobuf:"bytes,4,opt,name=fastest_node_name,json=fastestNodeName,proto3" json:"fastest_node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdTimestamp) Reset()         { *m = RcmdTimestamp{} }
func (m *RcmdTimestamp) String() string { return proto.CompactTextString(m) }
func (*RcmdTimestamp) ProtoMessage()    {}
func (*RcmdTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{4}
}

func (m *RcmdTimestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdTimestamp.Unmarshal(m, b)
}
func (m *RcmdTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdTimestamp.Marshal(b, m, deterministic)
}
func (m *RcmdTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdTimestamp.Merge(m, src)
}
func (m *RcmdTimestamp) XXX_Size() int {
	return xxx_messageInfo_RcmdTimestamp.Size(m)
}
func (m *RcmdTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdTimestamp proto.InternalMessageInfo

func (m *RcmdTimestamp) GetMinimumTime() string {
	if m != nil {
		return m.MinimumTime
	}
	return ""
}

func (m *RcmdTimestamp) GetMaximumTime() string {
	if m != nil {
		return m.MaximumTime
	}
	return ""
}

func (m *RcmdTimestamp) GetSlowestNodeName() string {
	if m != nil {
		return m.SlowestNodeName
	}
	return ""
}

func (m *RcmdTimestamp) GetFastestNodeName() string {
	if m != nil {
		return m.FastestNodeName
	}
	return ""
}

type RcmdShowPrefixLcInfo struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Speed                string   `protobuf:"bytes,2,opt,name=speed,proto3" json:"speed,omitempty"`
	FibComplete          string   `protobuf:"bytes,3,opt,name=fib_complete,json=fibComplete,proto3" json:"fib_complete,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RcmdShowPrefixLcInfo) Reset()         { *m = RcmdShowPrefixLcInfo{} }
func (m *RcmdShowPrefixLcInfo) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixLcInfo) ProtoMessage()    {}
func (*RcmdShowPrefixLcInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{5}
}

func (m *RcmdShowPrefixLcInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixLcInfo.Unmarshal(m, b)
}
func (m *RcmdShowPrefixLcInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixLcInfo.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixLcInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixLcInfo.Merge(m, src)
}
func (m *RcmdShowPrefixLcInfo) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixLcInfo.Size(m)
}
func (m *RcmdShowPrefixLcInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixLcInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixLcInfo proto.InternalMessageInfo

func (m *RcmdShowPrefixLcInfo) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RcmdShowPrefixLcInfo) GetSpeed() string {
	if m != nil {
		return m.Speed
	}
	return ""
}

func (m *RcmdShowPrefixLcInfo) GetFibComplete() string {
	if m != nil {
		return m.FibComplete
	}
	return ""
}

type RcmdShowPrefixTimeline struct {
	RouteOrigin          string                  `protobuf:"bytes,1,opt,name=route_origin,json=routeOrigin,proto3" json:"route_origin,omitempty"`
	RiBv4Enter           string                  `protobuf:"bytes,2,opt,name=ri_bv4_enter,json=riBv4Enter,proto3" json:"ri_bv4_enter,omitempty"`
	RiBv4Exit            string                  `protobuf:"bytes,3,opt,name=ri_bv4_exit,json=riBv4Exit,proto3" json:"ri_bv4_exit,omitempty"`
	RiBv4Redistribute    string                  `protobuf:"bytes,4,opt,name=ri_bv4_redistribute,json=riBv4Redistribute,proto3" json:"ri_bv4_redistribute,omitempty"`
	LdpEnter             string                  `protobuf:"bytes,5,opt,name=ldp_enter,json=ldpEnter,proto3" json:"ldp_enter,omitempty"`
	LdpExit              string                  `protobuf:"bytes,6,opt,name=ldp_exit,json=ldpExit,proto3" json:"ldp_exit,omitempty"`
	LsdEnter             string                  `protobuf:"bytes,7,opt,name=lsd_enter,json=lsdEnter,proto3" json:"lsd_enter,omitempty"`
	LsdExit              string                  `protobuf:"bytes,8,opt,name=lsd_exit,json=lsdExit,proto3" json:"lsd_exit,omitempty"`
	LcIp                 []*RcmdShowPrefixLcInfo `protobuf:"bytes,9,rep,name=lc_ip,json=lcIp,proto3" json:"lc_ip,omitempty"`
	LcMpls               []*RcmdShowPrefixLcInfo `protobuf:"bytes,10,rep,name=lc_mpls,json=lcMpls,proto3" json:"lc_mpls,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RcmdShowPrefixTimeline) Reset()         { *m = RcmdShowPrefixTimeline{} }
func (m *RcmdShowPrefixTimeline) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixTimeline) ProtoMessage()    {}
func (*RcmdShowPrefixTimeline) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{6}
}

func (m *RcmdShowPrefixTimeline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixTimeline.Unmarshal(m, b)
}
func (m *RcmdShowPrefixTimeline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixTimeline.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixTimeline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixTimeline.Merge(m, src)
}
func (m *RcmdShowPrefixTimeline) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixTimeline.Size(m)
}
func (m *RcmdShowPrefixTimeline) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixTimeline.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixTimeline proto.InternalMessageInfo

func (m *RcmdShowPrefixTimeline) GetRouteOrigin() string {
	if m != nil {
		return m.RouteOrigin
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetRiBv4Enter() string {
	if m != nil {
		return m.RiBv4Enter
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetRiBv4Exit() string {
	if m != nil {
		return m.RiBv4Exit
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetRiBv4Redistribute() string {
	if m != nil {
		return m.RiBv4Redistribute
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetLdpEnter() string {
	if m != nil {
		return m.LdpEnter
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetLdpExit() string {
	if m != nil {
		return m.LdpExit
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetLsdEnter() string {
	if m != nil {
		return m.LsdEnter
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetLsdExit() string {
	if m != nil {
		return m.LsdExit
	}
	return ""
}

func (m *RcmdShowPrefixTimeline) GetLcIp() []*RcmdShowPrefixLcInfo {
	if m != nil {
		return m.LcIp
	}
	return nil
}

func (m *RcmdShowPrefixTimeline) GetLcMpls() []*RcmdShowPrefixLcInfo {
	if m != nil {
		return m.LcMpls
	}
	return nil
}

type RcmdShowPrefixEvent struct {
	Prefix               string                    `protobuf:"bytes,50,opt,name=prefix,proto3" json:"prefix,omitempty"`
	PrefixLenth          uint32                    `protobuf:"varint,51,opt,name=prefix_lenth,json=prefixLenth,proto3" json:"prefix_lenth,omitempty"`
	SpfRunNo             uint32                    `protobuf:"varint,52,opt,name=spf_run_no,json=spfRunNo,proto3" json:"spf_run_no,omitempty"`
	IpfrrEventId         uint32                    `protobuf:"varint,53,opt,name=ipfrr_event_id,json=ipfrrEventId,proto3" json:"ipfrr_event_id,omitempty"`
	ThresholdExceeded    bool                      `protobuf:"varint,54,opt,name=threshold_exceeded,json=thresholdExceeded,proto3" json:"threshold_exceeded,omitempty"`
	Priority             string                    `protobuf:"bytes,55,opt,name=priority,proto3" json:"priority,omitempty"`
	ChangeType           string                    `protobuf:"bytes,56,opt,name=change_type,json=changeType,proto3" json:"change_type,omitempty"`
	RouteType            string                    `protobuf:"bytes,57,opt,name=route_type,json=routeType,proto3" json:"route_type,omitempty"`
	RoutePathChangeType  string                    `protobuf:"bytes,58,opt,name=route_path_change_type,json=routePathChangeType,proto3" json:"route_path_change_type,omitempty"`
	Cost                 uint32                    `protobuf:"varint,59,opt,name=cost,proto3" json:"cost,omitempty"`
	Path                 []*RcmdShowPrefixPath     `protobuf:"bytes,60,rep,name=path,proto3" json:"path,omitempty"`
	TriggerTime          string                    `protobuf:"bytes,61,opt,name=trigger_time,json=triggerTime,proto3" json:"trigger_time,omitempty"`
	TriggerLsa           []*RcmdLsaInfo            `protobuf:"bytes,62,rep,name=trigger_lsa,json=triggerLsa,proto3" json:"trigger_lsa,omitempty"`
	IpConvergenceTime    *RcmdTimestamp            `protobuf:"bytes,63,opt,name=ip_convergence_time,json=ipConvergenceTime,proto3" json:"ip_convergence_time,omitempty"`
	MplsConvergenceTime  *RcmdTimestamp            `protobuf:"bytes,64,opt,name=mpls_convergence_time,json=mplsConvergenceTime,proto3" json:"mpls_convergence_time,omitempty"`
	TimeLine             []*RcmdShowPrefixTimeline `protobuf:"bytes,65,rep,name=time_line,json=timeLine,proto3" json:"time_line,omitempty"`
	LsaProcessed         []*RcmdLsaInfo            `protobuf:"bytes,66,rep,name=lsa_processed,json=lsaProcessed,proto3" json:"lsa_processed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RcmdShowPrefixEvent) Reset()         { *m = RcmdShowPrefixEvent{} }
func (m *RcmdShowPrefixEvent) String() string { return proto.CompactTextString(m) }
func (*RcmdShowPrefixEvent) ProtoMessage()    {}
func (*RcmdShowPrefixEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3220bf186cf0ca5e, []int{7}
}

func (m *RcmdShowPrefixEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RcmdShowPrefixEvent.Unmarshal(m, b)
}
func (m *RcmdShowPrefixEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RcmdShowPrefixEvent.Marshal(b, m, deterministic)
}
func (m *RcmdShowPrefixEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RcmdShowPrefixEvent.Merge(m, src)
}
func (m *RcmdShowPrefixEvent) XXX_Size() int {
	return xxx_messageInfo_RcmdShowPrefixEvent.Size(m)
}
func (m *RcmdShowPrefixEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RcmdShowPrefixEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RcmdShowPrefixEvent proto.InternalMessageInfo

func (m *RcmdShowPrefixEvent) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetPrefixLenth() uint32 {
	if m != nil {
		return m.PrefixLenth
	}
	return 0
}

func (m *RcmdShowPrefixEvent) GetSpfRunNo() uint32 {
	if m != nil {
		return m.SpfRunNo
	}
	return 0
}

func (m *RcmdShowPrefixEvent) GetIpfrrEventId() uint32 {
	if m != nil {
		return m.IpfrrEventId
	}
	return 0
}

func (m *RcmdShowPrefixEvent) GetThresholdExceeded() bool {
	if m != nil {
		return m.ThresholdExceeded
	}
	return false
}

func (m *RcmdShowPrefixEvent) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetChangeType() string {
	if m != nil {
		return m.ChangeType
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetRouteType() string {
	if m != nil {
		return m.RouteType
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetRoutePathChangeType() string {
	if m != nil {
		return m.RoutePathChangeType
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetCost() uint32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *RcmdShowPrefixEvent) GetPath() []*RcmdShowPrefixPath {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *RcmdShowPrefixEvent) GetTriggerTime() string {
	if m != nil {
		return m.TriggerTime
	}
	return ""
}

func (m *RcmdShowPrefixEvent) GetTriggerLsa() []*RcmdLsaInfo {
	if m != nil {
		return m.TriggerLsa
	}
	return nil
}

func (m *RcmdShowPrefixEvent) GetIpConvergenceTime() *RcmdTimestamp {
	if m != nil {
		return m.IpConvergenceTime
	}
	return nil
}

func (m *RcmdShowPrefixEvent) GetMplsConvergenceTime() *RcmdTimestamp {
	if m != nil {
		return m.MplsConvergenceTime
	}
	return nil
}

func (m *RcmdShowPrefixEvent) GetTimeLine() []*RcmdShowPrefixTimeline {
	if m != nil {
		return m.TimeLine
	}
	return nil
}

func (m *RcmdShowPrefixEvent) GetLsaProcessed() []*RcmdLsaInfo {
	if m != nil {
		return m.LsaProcessed
	}
	return nil
}

func init() {
	proto.RegisterType((*RcmdShowPrefixEvent_KEYS)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_prefix_event_KEYS")
	proto.RegisterType((*RcmdShowIpfrrLfaPath)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_ipfrr_lfa_path")
	proto.RegisterType((*RcmdShowPrefixPath)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_prefix_path")
	proto.RegisterType((*RcmdLsaInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_lsa_info")
	proto.RegisterType((*RcmdTimestamp)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_timestamp")
	proto.RegisterType((*RcmdShowPrefixLcInfo)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_prefix_lc_info")
	proto.RegisterType((*RcmdShowPrefixTimeline)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_prefix_timeline")
	proto.RegisterType((*RcmdShowPrefixEvent)(nil), "cisco_ios_xr_infra_rcmd_oper.rcmd.ospf.instances.instance.summary_external_event_offlines.summary_external_event_offline.rcmd_show_prefix_event")
}

func init() { proto.RegisterFile("rcmd_show_prefix_event.proto", fileDescriptor_3220bf186cf0ca5e) }

var fileDescriptor_3220bf186cf0ca5e = []byte{
	// 1099 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0x26, 0xb6, 0xe3, 0x3c, 0xc7, 0x49, 0xbd, 0x21, 0xd1, 0x96, 0x16, 0x6a, 0x0c, 0x15,
	0x56, 0x11, 0x3e, 0x24, 0xe1, 0xff, 0xdf, 0x36, 0xca, 0x21, 0x22, 0x0d, 0xd5, 0xd2, 0x0b, 0x07,
	0x34, 0x5a, 0xef, 0xcc, 0xda, 0x23, 0xed, 0xee, 0x0c, 0x33, 0xe3, 0xd4, 0xf9, 0x04, 0x70, 0x43,
	0x82, 0xaf, 0xc0, 0x01, 0x24, 0xa8, 0xc4, 0x91, 0x03, 0x07, 0xbe, 0x07, 0xe2, 0x3b, 0x70, 0xe5,
	0x84, 0xe6, 0xcd, 0xac, 0x9d, 0xc6, 0x45, 0xed, 0x2d, 0xbe, 0xcd, 0xfc, 0xde, 0x7b, 0xf3, 0xde,
	0x6f, 0xe6, 0xb7, 0xef, 0x69, 0xe1, 0xa6, 0x4a, 0x0b, 0x4a, 0xf4, 0x58, 0x3c, 0x22, 0x52, 0xb1,
	0x8c, 0x4f, 0x09, 0x3b, 0x63, 0xa5, 0x19, 0x48, 0x25, 0x8c, 0x08, 0xa7, 0x29, 0xd7, 0xa9, 0x20,
	0x5c, 0x68, 0x32, 0x55, 0x84, 0x97, 0x99, 0x4a, 0x08, 0x06, 0x08, 0xc9, 0xd4, 0xc0, 0xae, 0x06,
	0x42, 0xcb, 0x6c, 0xc0, 0x4b, 0x6d, 0x92, 0x32, 0x65, 0x7a, 0xb6, 0x1a, 0xe8, 0x49, 0x51, 0x24,
	0xea, 0x9c, 0xb0, 0xa9, 0x61, 0xaa, 0x4c, 0x72, 0x77, 0x2e, 0x11, 0x59, 0x96, 0xf3, 0x92, 0xe9,
	0x67, 0xd8, 0x7b, 0x5f, 0xc1, 0x8d, 0xa7, 0x57, 0x46, 0x3e, 0x3b, 0xfa, 0xf2, 0x8b, 0xf0, 0x55,
	0x68, 0x57, 0x99, 0x48, 0x99, 0x14, 0x2c, 0x0a, 0xba, 0x41, 0x7f, 0x3d, 0xde, 0xa8, 0xc0, 0xd3,
	0xa4, 0x60, 0xe1, 0x75, 0x68, 0xba, 0x10, 0x4e, 0xa3, 0x95, 0x6e, 0xd0, 0x6f, 0xc7, 0x6b, 0xb8,
	0x3f, 0xa6, 0xbd, 0x7f, 0x02, 0x88, 0xe6, 0xe7, 0x73, 0x99, 0x29, 0x45, 0xf2, 0x2c, 0x21, 0x32,
	0x31, 0x63, 0x1b, 0x67, 0xd7, 0xe6, 0x5c, 0x56, 0xe7, 0xae, 0xe5, 0x59, 0xf2, 0xf0, 0x5c, 0xb2,
	0xf0, 0x36, 0x6c, 0xf2, 0xd2, 0x30, 0x95, 0x25, 0x55, 0xe2, 0x15, 0x74, 0x68, 0xcf, 0x50, 0xcc,
	0xfc, 0x06, 0x74, 0x4a, 0xc6, 0x47, 0xe3, 0xa1, 0x98, 0x28, 0x92, 0x50, 0xaa, 0x98, 0xd6, 0xd1,
	0x2a, 0x7a, 0x5e, 0x9b, 0x19, 0xee, 0x3a, 0x3c, 0xbc, 0x05, 0xad, 0x74, 0x9c, 0x94, 0x23, 0xe6,
	0x32, 0xd6, 0xd0, 0x0d, 0x1c, 0x84, 0x49, 0x6f, 0x41, 0xcb, 0xd6, 0x45, 0x0a, 0x66, 0x14, 0x4f,
	0xa3, 0x3a, 0x52, 0x01, 0x0b, 0xdd, 0x47, 0x24, 0x7c, 0x0d, 0x36, 0x15, 0x2b, 0x84, 0x61, 0xa4,
	0x14, 0x94, 0x59, 0xba, 0x0d, 0x77, 0x1d, 0x0e, 0x3d, 0x15, 0x94, 0x1d, 0xd3, 0xde, 0xdf, 0x2b,
	0xb0, 0xb3, 0x70, 0xa7, 0x48, 0x78, 0x91, 0x55, 0xf0, 0xdc, 0xac, 0x56, 0x9e, 0x8f, 0xd5, 0xea,
	0xb3, 0x58, 0xd5, 0x16, 0x58, 0x3d, 0x0e, 0xdc, 0x3b, 0x58, 0x28, 0xaa, 0x77, 0x57, 0xfb, 0xad,
	0xbd, 0xef, 0x83, 0xc1, 0x55, 0x29, 0x72, 0xf0, 0x7f, 0x72, 0x41, 0x71, 0x3c, 0x48, 0xcc, 0xb8,
	0xf7, 0x57, 0x00, 0x6d, 0xf4, 0xca, 0x75, 0x62, 0x4b, 0x13, 0xe1, 0x0e, 0x34, 0x70, 0x4d, 0xfd,
	0x85, 0xd6, 0x73, 0x9d, 0x1c, 0xd3, 0xf0, 0x75, 0xd8, 0xd2, 0xec, 0xeb, 0x09, 0x43, 0xf5, 0x4e,
	0x8a, 0x21, 0x53, 0xfe, 0x1a, 0x37, 0x2b, 0xf8, 0x14, 0x51, 0x54, 0xa2, 0x4e, 0x2e, 0xde, 0xe0,
	0x5a, 0xae, 0x9d, 0x12, 0xfb, 0x70, 0x4d, 0x28, 0x3e, 0xe2, 0x25, 0x51, 0x62, 0x62, 0x98, 0xb2,
	0x49, 0x9c, 0x74, 0x36, 0x1d, 0x1e, 0x23, 0x7c, 0x4c, 0x2f, 0xbf, 0x44, 0x7d, 0xe1, 0x25, 0x6e,
	0x5b, 0xf9, 0xa4, 0x4c, 0x1a, 0x2e, 0x4a, 0x62, 0x78, 0xc1, 0xbc, 0x7c, 0xda, 0x33, 0xf4, 0x21,
	0x2f, 0x58, 0xef, 0x71, 0x00, 0x9b, 0x48, 0xcf, 0xba, 0x68, 0x93, 0x14, 0x32, 0x7c, 0x05, 0x36,
	0x0a, 0x5e, 0xf2, 0x62, 0x52, 0xb8, 0x38, 0xc7, 0xb2, 0xe5, 0x31, 0x1b, 0x85, 0x2e, 0xc9, 0x74,
	0xee, 0xb2, 0xe2, 0x5d, 0x1c, 0x86, 0x2e, 0x77, 0xa0, 0xa3, 0x73, 0xf1, 0x88, 0x69, 0xe3, 0xf4,
	0x8b, 0x0a, 0x74, 0x74, 0xb7, 0xbc, 0xc1, 0x4a, 0x18, 0x35, 0x78, 0x07, 0x3a, 0x59, 0xa2, 0xcd,
	0x93, 0xbe, 0x8e, 0xf7, 0x96, 0x37, 0x54, 0xbe, 0x3d, 0x79, 0xf1, 0x1b, 0xf7, 0x7a, 0xcf, 0x53,
	0xf7, 0x32, 0x37, 0x60, 0x7d, 0x1e, 0xef, 0xca, 0x6e, 0x96, 0x55, 0x92, 0x17, 0xa0, 0xae, 0x25,
	0x63, 0xd4, 0x17, 0xeb, 0x36, 0x96, 0x49, 0xc6, 0x87, 0x24, 0x15, 0x85, 0xcc, 0x99, 0xa9, 0x2a,
	0x6c, 0x65, 0x7c, 0x78, 0xe8, 0xa1, 0xde, 0xbf, 0x35, 0xb8, 0xbe, 0x90, 0xd2, 0xd2, 0xb6, 0x0a,
	0xb2, 0x07, 0xe0, 0x5b, 0x11, 0xf7, 0x40, 0xd5, 0x6d, 0x21, 0xf6, 0x39, 0x42, 0x61, 0x17, 0x36,
	0x14, 0x27, 0xc3, 0xb3, 0x03, 0xc2, 0xec, 0xa7, 0xe7, 0x0b, 0x00, 0xc5, 0xef, 0x9d, 0x1d, 0x1c,
	0x59, 0x24, 0x7c, 0x19, 0x5a, 0x95, 0xc7, 0x94, 0x1b, 0x5f, 0xc4, 0xba, 0x73, 0x98, 0x72, 0x13,
	0x0e, 0x60, 0xdb, 0xdb, 0x15, 0xa3, 0x5c, 0x1b, 0xc5, 0x87, 0x13, 0x53, 0x5d, 0x51, 0x07, 0xfd,
	0xe2, 0x0b, 0x06, 0x7b, 0x11, 0x39, 0x95, 0x3e, 0x9d, 0xd3, 0x46, 0x33, 0xa7, 0xd2, 0x25, 0xb3,
	0xfa, 0xb3, 0x46, 0x9b, 0xa9, 0xe1, 0xf5, 0x47, 0x25, 0xe6, 0xb1, 0x71, 0x9a, 0xfa, 0xb8, 0x35,
	0x1f, 0xa7, 0xe9, 0x3c, 0xce, 0x1a, 0x6d, 0x5c, 0xb3, 0xd2, 0x2d, 0xc5, 0xb8, 0x9f, 0x02, 0xa8,
	0xdb, 0x47, 0x90, 0xd1, 0xfa, 0x12, 0x7d, 0xd2, 0x4f, 0xaa, 0x23, 0xae, 0xe5, 0xe9, 0xb1, 0x0c,
	0x7f, 0x09, 0x60, 0x2d, 0x4f, 0x49, 0x21, 0x73, 0x1d, 0xc1, 0xf2, 0x16, 0xdb, 0xc8, 0xd3, 0xfb,
	0x32, 0xd7, 0xbd, 0x3f, 0x5b, 0xb0, 0xfb, 0xf4, 0x99, 0x19, 0xee, 0x42, 0xc3, 0xed, 0xa3, 0x3d,
	0x7c, 0x0d, 0xbf, 0xb3, 0x8a, 0xac, 0x0e, 0x63, 0xa5, 0x19, 0x47, 0xfb, 0xd8, 0x84, 0x5b, 0x0e,
	0x3b, 0xb1, 0x50, 0x78, 0x13, 0x40, 0xcb, 0x8c, 0xa8, 0x49, 0x49, 0x4a, 0x11, 0x1d, 0xa0, 0x43,
	0x53, 0xcb, 0x2c, 0x9e, 0x94, 0xa7, 0xc2, 0x4e, 0x1e, 0xd7, 0x0d, 0x67, 0x83, 0xf6, 0x2d, 0xf4,
	0xd8, 0x40, 0xf4, 0xc8, 0x4d, 0xdb, 0xf0, 0x4d, 0x08, 0xcd, 0x58, 0x31, 0x3d, 0x16, 0xb9, 0x15,
	0x45, 0xca, 0x18, 0x65, 0x34, 0x7a, 0xbb, 0x1b, 0xf4, 0x9b, 0x71, 0x67, 0x66, 0x39, 0xf2, 0x86,
	0xf0, 0x45, 0x68, 0x4a, 0xc5, 0x85, 0xe2, 0xe6, 0x3c, 0x7a, 0xc7, 0x29, 0xab, 0xda, 0x5f, 0x6e,
	0x66, 0xef, 0x2e, 0x34, 0xb3, 0x97, 0x00, 0xdc, 0x47, 0x86, 0xf6, 0xf7, 0xfc, 0xe7, 0x61, 0x11,
	0x34, 0xef, 0xc3, 0xae, 0x33, 0xe3, 0xec, 0xb9, 0x78, 0xd4, 0xfb, 0xe8, 0xba, 0x8d, 0x56, 0xdb,
	0xce, 0x0f, 0xe7, 0x67, 0x86, 0x50, 0x4b, 0x85, 0x36, 0xd1, 0x07, 0xc8, 0x0d, 0xd7, 0xe1, 0x8f,
	0x01, 0xd4, 0x70, 0x32, 0x7d, 0x88, 0xca, 0xf8, 0x6e, 0x99, 0x94, 0x81, 0x63, 0x09, 0xab, 0xb3,
	0x2f, 0x6c, 0x14, 0x1f, 0x8d, 0x98, 0x72, 0xed, 0xf7, 0x23, 0xd7, 0x73, 0x3c, 0x86, 0xed, 0xf7,
	0xe7, 0x00, 0xaa, 0xbd, 0x9d, 0x5c, 0xd1, 0xc7, 0x48, 0xe8, 0x9b, 0xab, 0x26, 0x54, 0x0d, 0xd1,
	0x18, 0x7c, 0x71, 0x27, 0x3a, 0x09, 0x7f, 0x0f, 0x60, 0x9b, 0x4b, 0x92, 0x8a, 0xf2, 0x8c, 0xa9,
	0x11, 0x0e, 0x50, 0xa4, 0xf5, 0x49, 0x37, 0xe8, 0xb7, 0xf6, 0xbe, 0xbd, 0xea, 0x9a, 0x67, 0x93,
	0x31, 0xee, 0x70, 0x79, 0x38, 0x2f, 0x12, 0xef, 0xf9, 0x8f, 0x00, 0x76, 0x6c, 0x2f, 0x59, 0xac,
	0xfe, 0xd3, 0x65, 0xab, 0x7e, 0xdb, 0xd6, 0x79, 0xb9, 0xfe, 0xdf, 0x02, 0x58, 0xb7, 0x2e, 0xc4,
	0x46, 0x44, 0x77, 0x51, 0x25, 0x3f, 0x2c, 0x93, 0xec, 0xab, 0x41, 0x1b, 0x37, 0xed, 0xea, 0xc4,
	0x8e, 0xdc, 0x5f, 0x03, 0x68, 0x5b, 0x21, 0x49, 0x25, 0x52, 0xa6, 0x35, 0xa3, 0xd1, 0xbd, 0x25,
	0x53, 0xf7, 0x46, 0xae, 0x93, 0x07, 0x55, 0x75, 0xc3, 0x06, 0xfe, 0x77, 0xed, 0xff, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x68, 0x60, 0x43, 0xde, 0x97, 0x0d, 0x00, 0x00,
}
