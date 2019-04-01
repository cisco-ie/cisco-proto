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
// source: dnx_qosea_show_class_st.proto

package cisco_ios_xr_ncs5500_qos_oper_platform_qos_nodes_node_interfaces_interface_classes_class

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

type DnxQoseaShowClassSt_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	QosDirection         string   `protobuf:"bytes,3,opt,name=qos_direction,json=qosDirection,proto3" json:"qos_direction,omitempty"`
	LevelOneClassName    string   `protobuf:"bytes,4,opt,name=level_one_class_name,json=levelOneClassName,proto3" json:"level_one_class_name,omitempty"`
	LevelTwoClassName    string   `protobuf:"bytes,5,opt,name=level_two_class_name,json=levelTwoClassName,proto3" json:"level_two_class_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnxQoseaShowClassSt_KEYS) Reset()         { *m = DnxQoseaShowClassSt_KEYS{} }
func (m *DnxQoseaShowClassSt_KEYS) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowClassSt_KEYS) ProtoMessage()    {}
func (*DnxQoseaShowClassSt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{0}
}

func (m *DnxQoseaShowClassSt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowClassSt_KEYS.Unmarshal(m, b)
}
func (m *DnxQoseaShowClassSt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowClassSt_KEYS.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowClassSt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowClassSt_KEYS.Merge(m, src)
}
func (m *DnxQoseaShowClassSt_KEYS) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowClassSt_KEYS.Size(m)
}
func (m *DnxQoseaShowClassSt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowClassSt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowClassSt_KEYS proto.InternalMessageInfo

func (m *DnxQoseaShowClassSt_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *DnxQoseaShowClassSt_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *DnxQoseaShowClassSt_KEYS) GetQosDirection() string {
	if m != nil {
		return m.QosDirection
	}
	return ""
}

func (m *DnxQoseaShowClassSt_KEYS) GetLevelOneClassName() string {
	if m != nil {
		return m.LevelOneClassName
	}
	return ""
}

func (m *DnxQoseaShowClassSt_KEYS) GetLevelTwoClassName() string {
	if m != nil {
		return m.LevelTwoClassName
	}
	return ""
}

type DnxQoseaShowMarkSt_ struct {
	MarkType             string   `protobuf:"bytes,1,opt,name=mark_type,json=markType,proto3" json:"mark_type,omitempty"`
	MarkValue            uint32   `protobuf:"varint,2,opt,name=mark_value,json=markValue,proto3" json:"mark_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnxQoseaShowMarkSt_) Reset()         { *m = DnxQoseaShowMarkSt_{} }
func (m *DnxQoseaShowMarkSt_) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowMarkSt_) ProtoMessage()    {}
func (*DnxQoseaShowMarkSt_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{1}
}

func (m *DnxQoseaShowMarkSt_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowMarkSt_.Unmarshal(m, b)
}
func (m *DnxQoseaShowMarkSt_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowMarkSt_.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowMarkSt_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowMarkSt_.Merge(m, src)
}
func (m *DnxQoseaShowMarkSt_) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowMarkSt_.Size(m)
}
func (m *DnxQoseaShowMarkSt_) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowMarkSt_.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowMarkSt_ proto.InternalMessageInfo

func (m *DnxQoseaShowMarkSt_) GetMarkType() string {
	if m != nil {
		return m.MarkType
	}
	return ""
}

func (m *DnxQoseaShowMarkSt_) GetMarkValue() uint32 {
	if m != nil {
		return m.MarkValue
	}
	return 0
}

type DnxQoseaShowPolicyParamSt_ struct {
	PolicyValue          uint32   `protobuf:"varint,1,opt,name=policy_value,json=policyValue,proto3" json:"policy_value,omitempty"`
	PolicyUnit           string   `protobuf:"bytes,2,opt,name=policy_unit,json=policyUnit,proto3" json:"policy_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnxQoseaShowPolicyParamSt_) Reset()         { *m = DnxQoseaShowPolicyParamSt_{} }
func (m *DnxQoseaShowPolicyParamSt_) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowPolicyParamSt_) ProtoMessage()    {}
func (*DnxQoseaShowPolicyParamSt_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{2}
}

func (m *DnxQoseaShowPolicyParamSt_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowPolicyParamSt_.Unmarshal(m, b)
}
func (m *DnxQoseaShowPolicyParamSt_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowPolicyParamSt_.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowPolicyParamSt_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowPolicyParamSt_.Merge(m, src)
}
func (m *DnxQoseaShowPolicyParamSt_) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowPolicyParamSt_.Size(m)
}
func (m *DnxQoseaShowPolicyParamSt_) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowPolicyParamSt_.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowPolicyParamSt_ proto.InternalMessageInfo

func (m *DnxQoseaShowPolicyParamSt_) GetPolicyValue() uint32 {
	if m != nil {
		return m.PolicyValue
	}
	return 0
}

func (m *DnxQoseaShowPolicyParamSt_) GetPolicyUnit() string {
	if m != nil {
		return m.PolicyUnit
	}
	return ""
}

type DnxQoseaShowPolActionSt_ struct {
	ActionType           string                 `protobuf:"bytes,1,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	Mark                 []*DnxQoseaShowMarkSt_ `protobuf:"bytes,2,rep,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *DnxQoseaShowPolActionSt_) Reset()         { *m = DnxQoseaShowPolActionSt_{} }
func (m *DnxQoseaShowPolActionSt_) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowPolActionSt_) ProtoMessage()    {}
func (*DnxQoseaShowPolActionSt_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{3}
}

func (m *DnxQoseaShowPolActionSt_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowPolActionSt_.Unmarshal(m, b)
}
func (m *DnxQoseaShowPolActionSt_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowPolActionSt_.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowPolActionSt_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowPolActionSt_.Merge(m, src)
}
func (m *DnxQoseaShowPolActionSt_) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowPolActionSt_.Size(m)
}
func (m *DnxQoseaShowPolActionSt_) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowPolActionSt_.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowPolActionSt_ proto.InternalMessageInfo

func (m *DnxQoseaShowPolActionSt_) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

func (m *DnxQoseaShowPolActionSt_) GetMark() []*DnxQoseaShowMarkSt_ {
	if m != nil {
		return m.Mark
	}
	return nil
}

type DnxQoseaShowUint8RngSt_ struct {
	RangeStart           uint32   `protobuf:"varint,1,opt,name=range_start,json=rangeStart,proto3" json:"range_start,omitempty"`
	RangeEnd             uint32   `protobuf:"varint,2,opt,name=range_end,json=rangeEnd,proto3" json:"range_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DnxQoseaShowUint8RngSt_) Reset()         { *m = DnxQoseaShowUint8RngSt_{} }
func (m *DnxQoseaShowUint8RngSt_) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowUint8RngSt_) ProtoMessage()    {}
func (*DnxQoseaShowUint8RngSt_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{4}
}

func (m *DnxQoseaShowUint8RngSt_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowUint8RngSt_.Unmarshal(m, b)
}
func (m *DnxQoseaShowUint8RngSt_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowUint8RngSt_.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowUint8RngSt_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowUint8RngSt_.Merge(m, src)
}
func (m *DnxQoseaShowUint8RngSt_) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowUint8RngSt_.Size(m)
}
func (m *DnxQoseaShowUint8RngSt_) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowUint8RngSt_.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowUint8RngSt_ proto.InternalMessageInfo

func (m *DnxQoseaShowUint8RngSt_) GetRangeStart() uint32 {
	if m != nil {
		return m.RangeStart
	}
	return 0
}

func (m *DnxQoseaShowUint8RngSt_) GetRangeEnd() uint32 {
	if m != nil {
		return m.RangeEnd
	}
	return 0
}

type DnxQoseaShowWredParamsSt_ struct {
	WredMatchType             string                      `protobuf:"bytes,1,opt,name=wred_match_type,json=wredMatchType,proto3" json:"wred_match_type,omitempty"`
	WredMatchValue            []*DnxQoseaShowUint8RngSt_  `protobuf:"bytes,2,rep,name=wred_match_value,json=wredMatchValue,proto3" json:"wred_match_value,omitempty"`
	ConfigMinThreshold        *DnxQoseaShowPolicyParamSt_ `protobuf:"bytes,3,opt,name=config_min_threshold,json=configMinThreshold,proto3" json:"config_min_threshold,omitempty"`
	HardwareMinThresholdBytes uint32                      `protobuf:"varint,4,opt,name=hardware_min_threshold_bytes,json=hardwareMinThresholdBytes,proto3" json:"hardware_min_threshold_bytes,omitempty"`
	ConfigMaxThreshold        *DnxQoseaShowPolicyParamSt_ `protobuf:"bytes,5,opt,name=config_max_threshold,json=configMaxThreshold,proto3" json:"config_max_threshold,omitempty"`
	HardwareMaxThresholdBytes uint32                      `protobuf:"varint,6,opt,name=hardware_max_threshold_bytes,json=hardwareMaxThresholdBytes,proto3" json:"hardware_max_threshold_bytes,omitempty"`
	FirstSegment              uint32                      `protobuf:"varint,7,opt,name=first_segment,json=firstSegment,proto3" json:"first_segment,omitempty"`
	SegmentSize               uint32                      `protobuf:"varint,8,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                    `json:"-"`
	XXX_unrecognized          []byte                      `json:"-"`
	XXX_sizecache             int32                       `json:"-"`
}

func (m *DnxQoseaShowWredParamsSt_) Reset()         { *m = DnxQoseaShowWredParamsSt_{} }
func (m *DnxQoseaShowWredParamsSt_) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowWredParamsSt_) ProtoMessage()    {}
func (*DnxQoseaShowWredParamsSt_) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{5}
}

func (m *DnxQoseaShowWredParamsSt_) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowWredParamsSt_.Unmarshal(m, b)
}
func (m *DnxQoseaShowWredParamsSt_) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowWredParamsSt_.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowWredParamsSt_) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowWredParamsSt_.Merge(m, src)
}
func (m *DnxQoseaShowWredParamsSt_) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowWredParamsSt_.Size(m)
}
func (m *DnxQoseaShowWredParamsSt_) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowWredParamsSt_.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowWredParamsSt_ proto.InternalMessageInfo

func (m *DnxQoseaShowWredParamsSt_) GetWredMatchType() string {
	if m != nil {
		return m.WredMatchType
	}
	return ""
}

func (m *DnxQoseaShowWredParamsSt_) GetWredMatchValue() []*DnxQoseaShowUint8RngSt_ {
	if m != nil {
		return m.WredMatchValue
	}
	return nil
}

func (m *DnxQoseaShowWredParamsSt_) GetConfigMinThreshold() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigMinThreshold
	}
	return nil
}

func (m *DnxQoseaShowWredParamsSt_) GetHardwareMinThresholdBytes() uint32 {
	if m != nil {
		return m.HardwareMinThresholdBytes
	}
	return 0
}

func (m *DnxQoseaShowWredParamsSt_) GetConfigMaxThreshold() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigMaxThreshold
	}
	return nil
}

func (m *DnxQoseaShowWredParamsSt_) GetHardwareMaxThresholdBytes() uint32 {
	if m != nil {
		return m.HardwareMaxThresholdBytes
	}
	return 0
}

func (m *DnxQoseaShowWredParamsSt_) GetFirstSegment() uint32 {
	if m != nil {
		return m.FirstSegment
	}
	return 0
}

func (m *DnxQoseaShowWredParamsSt_) GetSegmentSize() uint32 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

type DnxQoseaShowClassSt struct {
	ClassLevel                       string                       `protobuf:"bytes,50,opt,name=class_level,json=classLevel,proto3" json:"class_level,omitempty"`
	IpMark                           []*DnxQoseaShowMarkSt_       `protobuf:"bytes,51,rep,name=ip_mark,json=ipMark,proto3" json:"ip_mark,omitempty"`
	CommonMark                       []*DnxQoseaShowMarkSt_       `protobuf:"bytes,52,rep,name=common_mark,json=commonMark,proto3" json:"common_mark,omitempty"`
	MplsMark                         []*DnxQoseaShowMarkSt_       `protobuf:"bytes,53,rep,name=mpls_mark,json=mplsMark,proto3" json:"mpls_mark,omitempty"`
	EgressQueueId                    int32                        `protobuf:"zigzag32,54,opt,name=egress_queue_id,json=egressQueueId,proto3" json:"egress_queue_id,omitempty"`
	QueueType                        string                       `protobuf:"bytes,55,opt,name=queue_type,json=queueType,proto3" json:"queue_type,omitempty"`
	PriorityLevel                    string                       `protobuf:"bytes,56,opt,name=priority_level,json=priorityLevel,proto3" json:"priority_level,omitempty"`
	ConfigMaxRate                    *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,57,opt,name=config_max_rate,json=configMaxRate,proto3" json:"config_max_rate,omitempty"`
	HardwareMaxRateKbps              uint32                       `protobuf:"varint,58,opt,name=hardware_max_rate_kbps,json=hardwareMaxRateKbps,proto3" json:"hardware_max_rate_kbps,omitempty"`
	ConfigMinRate                    *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,59,opt,name=config_min_rate,json=configMinRate,proto3" json:"config_min_rate,omitempty"`
	HardwareMinRateKbps              uint32                       `protobuf:"varint,60,opt,name=hardware_min_rate_kbps,json=hardwareMinRateKbps,proto3" json:"hardware_min_rate_kbps,omitempty"`
	ConfigExcessBandwidthPercent     uint32                       `protobuf:"varint,61,opt,name=config_excess_bandwidth_percent,json=configExcessBandwidthPercent,proto3" json:"config_excess_bandwidth_percent,omitempty"`
	ConfigExcessBandwidthUnit        uint32                       `protobuf:"varint,62,opt,name=config_excess_bandwidth_unit,json=configExcessBandwidthUnit,proto3" json:"config_excess_bandwidth_unit,omitempty"`
	HardwareExcessBandwidthWeight    uint32                       `protobuf:"varint,63,opt,name=hardware_excess_bandwidth_weight,json=hardwareExcessBandwidthWeight,proto3" json:"hardware_excess_bandwidth_weight,omitempty"`
	NetworkMinBandwidthKbps          uint32                       `protobuf:"varint,64,opt,name=network_min_bandwidth_kbps,json=networkMinBandwidthKbps,proto3" json:"network_min_bandwidth_kbps,omitempty"`
	ConfigQueueLimit                 *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,65,opt,name=config_queue_limit,json=configQueueLimit,proto3" json:"config_queue_limit,omitempty"`
	HardwareQueueLimitBytes          uint64                       `protobuf:"varint,66,opt,name=hardware_queue_limit_bytes,json=hardwareQueueLimitBytes,proto3" json:"hardware_queue_limit_bytes,omitempty"`
	HardwareQueueLimitMicroseconds   uint64                       `protobuf:"varint,67,opt,name=hardware_queue_limit_microseconds,json=hardwareQueueLimitMicroseconds,proto3" json:"hardware_queue_limit_microseconds,omitempty"`
	PolicerBucketId                  uint32                       `protobuf:"varint,68,opt,name=policer_bucket_id,json=policerBucketId,proto3" json:"policer_bucket_id,omitempty"`
	PolicerStatsHandle               uint64                       `protobuf:"varint,69,opt,name=policer_stats_handle,json=policerStatsHandle,proto3" json:"policer_stats_handle,omitempty"`
	ConfigPolicerAverageRate         *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,70,opt,name=config_policer_average_rate,json=configPolicerAverageRate,proto3" json:"config_policer_average_rate,omitempty"`
	HardwarePolicerAverageRateKbps   uint32                       `protobuf:"varint,71,opt,name=hardware_policer_average_rate_kbps,json=hardwarePolicerAverageRateKbps,proto3" json:"hardware_policer_average_rate_kbps,omitempty"`
	ConfigPolicerPeakRate            *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,72,opt,name=config_policer_peak_rate,json=configPolicerPeakRate,proto3" json:"config_policer_peak_rate,omitempty"`
	HardwarePolicerPeakRateKbps      uint32                       `protobuf:"varint,73,opt,name=hardware_policer_peak_rate_kbps,json=hardwarePolicerPeakRateKbps,proto3" json:"hardware_policer_peak_rate_kbps,omitempty"`
	ConfigPolicerConformBurst        *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,74,opt,name=config_policer_conform_burst,json=configPolicerConformBurst,proto3" json:"config_policer_conform_burst,omitempty"`
	HardwarePolicerConformBurstBytes uint32                       `protobuf:"varint,75,opt,name=hardware_policer_conform_burst_bytes,json=hardwarePolicerConformBurstBytes,proto3" json:"hardware_policer_conform_burst_bytes,omitempty"`
	ConfigPolicerExcessBurst         *DnxQoseaShowPolicyParamSt_  `protobuf:"bytes,76,opt,name=config_policer_excess_burst,json=configPolicerExcessBurst,proto3" json:"config_policer_excess_burst,omitempty"`
	HardwarePolicerExcessBurstBytes  uint32                       `protobuf:"varint,77,opt,name=hardware_policer_excess_burst_bytes,json=hardwarePolicerExcessBurstBytes,proto3" json:"hardware_policer_excess_burst_bytes,omitempty"`
	ConformAction                    *DnxQoseaShowPolActionSt_    `protobuf:"bytes,78,opt,name=conform_action,json=conformAction,proto3" json:"conform_action,omitempty"`
	ExceedAction                     *DnxQoseaShowPolActionSt_    `protobuf:"bytes,79,opt,name=exceed_action,json=exceedAction,proto3" json:"exceed_action,omitempty"`
	ViolateAction                    *DnxQoseaShowPolActionSt_    `protobuf:"bytes,80,opt,name=violate_action,json=violateAction,proto3" json:"violate_action,omitempty"`
	Wred                             []*DnxQoseaShowWredParamsSt_ `protobuf:"bytes,81,rep,name=wred,proto3" json:"wred,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                     `json:"-"`
	XXX_unrecognized                 []byte                       `json:"-"`
	XXX_sizecache                    int32                        `json:"-"`
}

func (m *DnxQoseaShowClassSt) Reset()         { *m = DnxQoseaShowClassSt{} }
func (m *DnxQoseaShowClassSt) String() string { return proto.CompactTextString(m) }
func (*DnxQoseaShowClassSt) ProtoMessage()    {}
func (*DnxQoseaShowClassSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a790e8a59d1ee85, []int{6}
}

func (m *DnxQoseaShowClassSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnxQoseaShowClassSt.Unmarshal(m, b)
}
func (m *DnxQoseaShowClassSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnxQoseaShowClassSt.Marshal(b, m, deterministic)
}
func (m *DnxQoseaShowClassSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnxQoseaShowClassSt.Merge(m, src)
}
func (m *DnxQoseaShowClassSt) XXX_Size() int {
	return xxx_messageInfo_DnxQoseaShowClassSt.Size(m)
}
func (m *DnxQoseaShowClassSt) XXX_DiscardUnknown() {
	xxx_messageInfo_DnxQoseaShowClassSt.DiscardUnknown(m)
}

var xxx_messageInfo_DnxQoseaShowClassSt proto.InternalMessageInfo

func (m *DnxQoseaShowClassSt) GetClassLevel() string {
	if m != nil {
		return m.ClassLevel
	}
	return ""
}

func (m *DnxQoseaShowClassSt) GetIpMark() []*DnxQoseaShowMarkSt_ {
	if m != nil {
		return m.IpMark
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetCommonMark() []*DnxQoseaShowMarkSt_ {
	if m != nil {
		return m.CommonMark
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetMplsMark() []*DnxQoseaShowMarkSt_ {
	if m != nil {
		return m.MplsMark
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetEgressQueueId() int32 {
	if m != nil {
		return m.EgressQueueId
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetQueueType() string {
	if m != nil {
		return m.QueueType
	}
	return ""
}

func (m *DnxQoseaShowClassSt) GetPriorityLevel() string {
	if m != nil {
		return m.PriorityLevel
	}
	return ""
}

func (m *DnxQoseaShowClassSt) GetConfigMaxRate() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigMaxRate
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwareMaxRateKbps() uint32 {
	if m != nil {
		return m.HardwareMaxRateKbps
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigMinRate() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigMinRate
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwareMinRateKbps() uint32 {
	if m != nil {
		return m.HardwareMinRateKbps
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigExcessBandwidthPercent() uint32 {
	if m != nil {
		return m.ConfigExcessBandwidthPercent
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigExcessBandwidthUnit() uint32 {
	if m != nil {
		return m.ConfigExcessBandwidthUnit
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetHardwareExcessBandwidthWeight() uint32 {
	if m != nil {
		return m.HardwareExcessBandwidthWeight
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetNetworkMinBandwidthKbps() uint32 {
	if m != nil {
		return m.NetworkMinBandwidthKbps
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigQueueLimit() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigQueueLimit
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwareQueueLimitBytes() uint64 {
	if m != nil {
		return m.HardwareQueueLimitBytes
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetHardwareQueueLimitMicroseconds() uint64 {
	if m != nil {
		return m.HardwareQueueLimitMicroseconds
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetPolicerBucketId() uint32 {
	if m != nil {
		return m.PolicerBucketId
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetPolicerStatsHandle() uint64 {
	if m != nil {
		return m.PolicerStatsHandle
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigPolicerAverageRate() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigPolicerAverageRate
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwarePolicerAverageRateKbps() uint32 {
	if m != nil {
		return m.HardwarePolicerAverageRateKbps
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigPolicerPeakRate() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigPolicerPeakRate
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwarePolicerPeakRateKbps() uint32 {
	if m != nil {
		return m.HardwarePolicerPeakRateKbps
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigPolicerConformBurst() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigPolicerConformBurst
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwarePolicerConformBurstBytes() uint32 {
	if m != nil {
		return m.HardwarePolicerConformBurstBytes
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConfigPolicerExcessBurst() *DnxQoseaShowPolicyParamSt_ {
	if m != nil {
		return m.ConfigPolicerExcessBurst
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetHardwarePolicerExcessBurstBytes() uint32 {
	if m != nil {
		return m.HardwarePolicerExcessBurstBytes
	}
	return 0
}

func (m *DnxQoseaShowClassSt) GetConformAction() *DnxQoseaShowPolActionSt_ {
	if m != nil {
		return m.ConformAction
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetExceedAction() *DnxQoseaShowPolActionSt_ {
	if m != nil {
		return m.ExceedAction
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetViolateAction() *DnxQoseaShowPolActionSt_ {
	if m != nil {
		return m.ViolateAction
	}
	return nil
}

func (m *DnxQoseaShowClassSt) GetWred() []*DnxQoseaShowWredParamsSt_ {
	if m != nil {
		return m.Wred
	}
	return nil
}

func init() {
	proto.RegisterType((*DnxQoseaShowClassSt_KEYS)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_class_st_KEYS")
	proto.RegisterType((*DnxQoseaShowMarkSt_)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_mark_st_")
	proto.RegisterType((*DnxQoseaShowPolicyParamSt_)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_policy_param_st_")
	proto.RegisterType((*DnxQoseaShowPolActionSt_)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_pol_action_st_")
	proto.RegisterType((*DnxQoseaShowUint8RngSt_)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_uint8_rng_st_")
	proto.RegisterType((*DnxQoseaShowWredParamsSt_)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_wred_params_st_")
	proto.RegisterType((*DnxQoseaShowClassSt)(nil), "cisco_ios_xr_ncs5500_qos_oper.platform_qos.nodes.node.interfaces.interface.classes.class.dnx_qosea_show_class_st")
}

func init() { proto.RegisterFile("dnx_qosea_show_class_st.proto", fileDescriptor_4a790e8a59d1ee85) }

var fileDescriptor_4a790e8a59d1ee85 = []byte{
	// 1299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x58, 0x4f, 0x6f, 0x1b, 0xc5,
	0x1b, 0xd6, 0xb6, 0x69, 0x9a, 0x4c, 0xe2, 0xa4, 0xdd, 0x5f, 0x7f, 0x74, 0xfb, 0xd7, 0xae, 0x4b,
	0xab, 0x88, 0x83, 0xa9, 0x5a, 0x0a, 0x85, 0x02, 0xa5, 0x69, 0x43, 0x9b, 0x36, 0x6e, 0x53, 0xa7,
	0xe5, 0x8f, 0x84, 0x34, 0x1a, 0xef, 0x4e, 0xec, 0x91, 0x77, 0x67, 0x36, 0x33, 0xe3, 0xd8, 0xee,
	0x9d, 0x0b, 0x42, 0x48, 0x70, 0x46, 0x08, 0x4e, 0x1c, 0xb8, 0xc1, 0x05, 0xf1, 0x09, 0xf8, 0x2a,
	0x7c, 0x0b, 0x34, 0xef, 0xcc, 0x6e, 0xd6, 0xeb, 0xe4, 0x48, 0x9c, 0x4b, 0x62, 0x3f, 0xf3, 0xbe,
	0xef, 0x3e, 0xcf, 0x33, 0xf3, 0xbe, 0xeb, 0x5d, 0x74, 0x29, 0xe2, 0x43, 0xbc, 0x23, 0x14, 0x25,
	0x58, 0x75, 0xc5, 0x00, 0x87, 0x31, 0x51, 0x0a, 0x2b, 0xdd, 0x48, 0xa5, 0xd0, 0xc2, 0xff, 0x22,
	0x64, 0x2a, 0x14, 0x98, 0x09, 0x85, 0x87, 0x12, 0xf3, 0x50, 0xdd, 0xbe, 0x7d, 0xe3, 0x86, 0x89,
	0xc7, 0x22, 0xa5, 0xb2, 0x91, 0xc6, 0x44, 0x6f, 0x0b, 0x99, 0x18, 0xa4, 0xc1, 0x45, 0x44, 0xed,
	0xdf, 0x06, 0xe3, 0x9a, 0xca, 0x6d, 0x12, 0x52, 0xb5, 0xf7, 0xb1, 0x01, 0xa5, 0xa9, 0xb2, 0xff,
	0xeb, 0xff, 0x78, 0xe8, 0xe2, 0x01, 0xd7, 0xc6, 0x4f, 0xd7, 0xbe, 0xdc, 0xf2, 0x2f, 0xa0, 0x79,
	0x53, 0x0b, 0x73, 0x92, 0xd0, 0xc0, 0xab, 0x79, 0x2b, 0xf3, 0xad, 0x39, 0x03, 0x3c, 0x23, 0x09,
	0xf5, 0xaf, 0xa1, 0xa5, 0xbc, 0xb0, 0x8d, 0x38, 0x06, 0x11, 0x95, 0x1c, 0x85, 0xb0, 0xab, 0xa8,
	0x62, 0xb8, 0x46, 0x4c, 0xd2, 0x50, 0x33, 0xc1, 0x83, 0xe3, 0x10, 0xb5, 0xb8, 0x23, 0xd4, 0xc3,
	0x0c, 0xf3, 0xdf, 0x46, 0x67, 0x62, 0xba, 0x4b, 0x63, 0x2c, 0x38, 0x75, 0x1c, 0xa0, 0xe2, 0x0c,
	0xc4, 0x9e, 0x86, 0xb5, 0xe7, 0x9c, 0x3e, 0x30, 0x2b, 0x50, 0x35, 0x4f, 0xd0, 0x03, 0x51, 0x4c,
	0x38, 0x51, 0x48, 0x78, 0x39, 0x10, 0x79, 0x42, 0xfd, 0x15, 0x3a, 0x5b, 0x92, 0x9a, 0x10, 0xd9,
	0x33, 0x4a, 0x8d, 0x4a, 0xf8, 0xac, 0x47, 0x69, 0xae, 0xd2, 0x00, 0x2f, 0x47, 0x29, 0xf5, 0x2f,
	0x21, 0x04, 0x8b, 0xbb, 0x24, 0xee, 0x5b, 0x85, 0x95, 0x16, 0x84, 0x7f, 0x66, 0x80, 0x3a, 0x45,
	0xd5, 0x52, 0xd9, 0x54, 0xc4, 0x2c, 0x1c, 0xe1, 0x94, 0x48, 0x92, 0x40, 0xf9, 0x2b, 0x68, 0xd1,
	0x61, 0xb6, 0x86, 0x07, 0x35, 0x16, 0x2c, 0x06, 0x55, 0xfc, 0x2a, 0x72, 0x5f, 0x71, 0x9f, 0x33,
	0xed, 0x7c, 0x44, 0x16, 0x7a, 0xc5, 0x99, 0xae, 0xff, 0xed, 0x4d, 0x9c, 0x92, 0x54, 0xc4, 0x98,
	0x80, 0x7b, 0x70, 0x95, 0x2a, 0x5a, 0x70, 0xdf, 0x0a, 0x32, 0x90, 0x85, 0x40, 0xc8, 0xd7, 0x1e,
	0x9a, 0x31, 0xbc, 0x83, 0x63, 0xb5, 0xe3, 0x2b, 0x0b, 0x37, 0x77, 0x1a, 0xff, 0xd5, 0xb1, 0x6a,
	0x1c, 0xe0, 0x73, 0x0b, 0x2e, 0x5f, 0xff, 0x6a, 0xe2, 0xcc, 0xf5, 0x19, 0xd7, 0x77, 0xb0, 0xe4,
	0x9d, 0x4c, 0x88, 0x24, 0xbc, 0x43, 0xb1, 0xd2, 0x44, 0x6a, 0xe7, 0x16, 0x02, 0x68, 0xcb, 0x20,
	0x66, 0xbb, 0x6c, 0x00, 0xe5, 0x91, 0xdb, 0x90, 0x39, 0x00, 0xd6, 0x78, 0x54, 0xff, 0x7e, 0x16,
	0x5d, 0x2e, 0x95, 0x1f, 0x48, 0x1a, 0xd9, 0xed, 0x80, 0x83, 0xed, 0x5f, 0x47, 0xcb, 0x00, 0x25,
	0x44, 0x87, 0xdd, 0xa2, 0x5b, 0x15, 0x03, 0x37, 0x0d, 0x0a, 0x86, 0xfd, 0xec, 0xa1, 0x53, 0x85,
	0xc0, 0xec, 0x00, 0x18, 0xf3, 0x76, 0x0f, 0xcd, 0xbc, 0x31, 0x6f, 0x5a, 0x4b, 0x39, 0x43, 0x7b,
	0x6e, 0x7e, 0xf3, 0xd0, 0x99, 0x50, 0xf0, 0x6d, 0xd6, 0xc1, 0x09, 0xe3, 0x58, 0x77, 0x25, 0x55,
	0x5d, 0x11, 0x47, 0xd0, 0x63, 0x0b, 0x37, 0x47, 0x87, 0x46, 0xb3, 0x7c, 0xe8, 0x5b, 0xbe, 0xa5,
	0xd5, 0x64, 0xfc, 0x65, 0x46, 0xca, 0xbf, 0x87, 0x2e, 0x76, 0x89, 0x8c, 0x06, 0x44, 0xd2, 0x71,
	0xba, 0xb8, 0x3d, 0xd2, 0x54, 0x41, 0xb3, 0x57, 0x5a, 0xe7, 0xb2, 0x98, 0x62, 0xee, 0xaa, 0x09,
	0x18, 0x93, 0x4b, 0x86, 0x05, 0xb9, 0x27, 0x8e, 0x8a, 0x5c, 0x32, 0x3c, 0x40, 0x6e, 0x91, 0xae,
	0x93, 0x3b, 0x5b, 0x92, 0x5b, 0xc8, 0xb5, 0x72, 0xaf, 0xa2, 0xca, 0x36, 0x93, 0x4a, 0x63, 0x45,
	0x3b, 0x09, 0xe5, 0x3a, 0x38, 0x09, 0x19, 0x8b, 0x00, 0x6e, 0x59, 0xcc, 0x4c, 0x17, 0xb7, 0x8c,
	0x15, 0x7b, 0x4d, 0x83, 0x39, 0x3b, 0x5d, 0x1c, 0xb6, 0xc5, 0x5e, 0xd3, 0xfa, 0x1f, 0xe7, 0x27,
	0x66, 0x5f, 0x36, 0xe6, 0x4d, 0xb7, 0xd9, 0xcf, 0x30, 0x31, 0x83, 0x9b, 0x76, 0x6c, 0x00, 0xb4,
	0x61, 0x10, 0xff, 0x1b, 0x0f, 0x9d, 0x64, 0x29, 0x34, 0x71, 0x70, 0x6b, 0x5a, 0x93, 0x63, 0x96,
	0xa5, 0x4d, 0x22, 0x7b, 0xfe, 0x0f, 0x1e, 0x5a, 0x08, 0x45, 0x92, 0x08, 0x6e, 0x09, 0xbd, 0x33,
	0x2d, 0x42, 0xc8, 0xb2, 0x00, 0x52, 0xdf, 0x79, 0x68, 0x3e, 0x49, 0x63, 0x65, 0x29, 0xdd, 0x9e,
	0x16, 0xa5, 0x39, 0xc3, 0x01, 0x08, 0x5d, 0x47, 0xcb, 0xb4, 0x23, 0xa9, 0x52, 0x78, 0xa7, 0x4f,
	0xfb, 0x14, 0xb3, 0x28, 0x78, 0xb7, 0xe6, 0xad, 0x9c, 0x6e, 0x55, 0x2c, 0xfc, 0xc2, 0xa0, 0xeb,
	0x91, 0xb9, 0xb5, 0xd9, 0x00, 0x98, 0x81, 0xef, 0xc1, 0xd6, 0xcf, 0x03, 0x02, 0xf3, 0xef, 0x1a,
	0x5a, 0x4a, 0x25, 0x13, 0x92, 0xe9, 0x91, 0x3b, 0x1d, 0x77, 0xec, 0x98, 0xcc, 0x50, 0x7b, 0x40,
	0x7e, 0xf1, 0xd0, 0x72, 0xa1, 0x29, 0x25, 0xd1, 0x34, 0x78, 0x7f, 0xda, 0xfd, 0x58, 0xc9, 0xfb,
	0xb1, 0x45, 0x34, 0xf5, 0x6f, 0xa1, 0x37, 0xc6, 0x5a, 0xd1, 0x90, 0xc4, 0xbd, 0x76, 0xaa, 0x82,
	0x0f, 0xa0, 0x5d, 0xfe, 0x57, 0x68, 0x42, 0x93, 0xf0, 0xb4, 0x9d, 0xaa, 0x31, 0x61, 0x8c, 0x5b,
	0x61, 0x77, 0x8f, 0x8a, 0x30, 0xc6, 0x27, 0x85, 0x39, 0x92, 0x56, 0xd8, 0x87, 0x25, 0x61, 0x36,
	0x01, 0x84, 0xad, 0xa1, 0xaa, 0xd3, 0x45, 0x87, 0xa1, 0x39, 0x26, 0x6d, 0xc2, 0xa3, 0x01, 0x8b,
	0x74, 0x17, 0xa7, 0x54, 0x86, 0x66, 0xd2, 0x7c, 0x04, 0xd9, 0x17, 0x6d, 0xd8, 0x1a, 0x44, 0xad,
	0x66, 0x41, 0x9b, 0x36, 0xc6, 0xcc, 0xb7, 0x83, 0xca, 0xc0, 0xaf, 0x98, 0x8f, 0xed, 0x7c, 0xdb,
	0xb7, 0x86, 0xf9, 0x51, 0xe3, 0x3f, 0x42, 0xb5, 0x9c, 0xfc, 0x44, 0x89, 0x01, 0x65, 0x9d, 0xae,
	0x0e, 0xee, 0x41, 0x91, 0x4b, 0x59, 0x5c, 0xa9, 0xcc, 0xe7, 0x10, 0xe4, 0xdf, 0x45, 0xe7, 0x39,
	0xd5, 0x03, 0x21, 0x7b, 0x60, 0xc2, 0x5e, 0x11, 0x70, 0xe2, 0x13, 0x28, 0x71, 0xd6, 0x45, 0x34,
	0x19, 0xcf, 0xd3, 0xc1, 0x8d, 0x5f, 0x3d, 0xe4, 0xa6, 0xb7, 0x6b, 0x97, 0x98, 0x25, 0x4c, 0x07,
	0xf7, 0xa7, 0xbd, 0xd3, 0xa7, 0x2c, 0x29, 0x68, 0xd6, 0x0d, 0x43, 0xc9, 0xc8, 0xcc, 0xfd, 0x2a,
	0x50, 0x75, 0xb7, 0x93, 0xd5, 0x9a, 0xb7, 0x32, 0xd3, 0x3a, 0x9b, 0x45, 0xec, 0xe5, 0xd9, 0x9b,
	0xc9, 0x3a, 0xba, 0xb2, 0x6f, 0x72, 0xc2, 0x42, 0x29, 0x14, 0x0d, 0x05, 0x8f, 0x54, 0xf0, 0x00,
	0x6a, 0x5c, 0x9e, 0xac, 0xd1, 0x2c, 0x44, 0xf9, 0x6f, 0xa1, 0xd3, 0xc0, 0x96, 0x4a, 0xdc, 0xee,
	0x87, 0x3d, 0xaa, 0xcd, 0x84, 0x79, 0x08, 0x2e, 0x2f, 0xbb, 0x85, 0x55, 0xc0, 0xd7, 0x23, 0xff,
	0x06, 0x3a, 0x93, 0xc5, 0x2a, 0x4d, 0xb4, 0xc2, 0x5d, 0xc2, 0xa3, 0x98, 0x06, 0x6b, 0x70, 0x25,
	0xdf, 0xad, 0x6d, 0x99, 0xa5, 0xc7, 0xb0, 0xe2, 0xff, 0xe9, 0xa1, 0x0b, 0x6e, 0x3f, 0xb2, 0x4c,
	0xb2, 0x4b, 0x25, 0xe9, 0x50, 0xdb, 0x82, 0x9f, 0x4e, 0x7b, 0x63, 0x02, 0xcb, 0x6e, 0xd3, 0x92,
	0xbb, 0x6f, 0xb9, 0x41, 0x37, 0x3e, 0x41, 0xf5, 0xdc, 0xe3, 0xfd, 0xb8, 0xdb, 0xf3, 0xf8, 0x08,
	0x9c, 0xca, 0x4d, 0x9e, 0xac, 0x03, 0xc7, 0xf2, 0x77, 0x0f, 0x05, 0x25, 0x1b, 0x52, 0x4a, 0x7a,
	0xd6, 0x83, 0xc7, 0xd3, 0xf6, 0xe0, 0xff, 0x63, 0x1e, 0x6c, 0x52, 0xd2, 0x03, 0x03, 0x1e, 0xa2,
	0xea, 0x84, 0x01, 0x39, 0x6b, 0xab, 0x7e, 0x1d, 0xd4, 0x5f, 0x28, 0xa9, 0xcf, 0x2a, 0x80, 0xf4,
	0xbf, 0xbc, 0x7c, 0xb2, 0x64, 0x45, 0xcc, 0x57, 0xa3, 0xa4, 0xdd, 0x97, 0x4a, 0x07, 0x4f, 0xa6,
	0x2d, 0xff, 0xdc, 0x98, 0xfc, 0x07, 0x96, 0xdc, 0xaa, 0xe1, 0xe6, 0x3f, 0x43, 0x6f, 0x4e, 0x58,
	0x30, 0xc6, 0xde, 0xb5, 0xeb, 0x53, 0xf0, 0xa1, 0x56, 0xf2, 0xa1, 0x58, 0xca, 0xf6, 0xed, 0x3e,
	0xed, 0x90, 0xcd, 0x4a, 0xf0, 0x62, 0xe3, 0x68, 0xb5, 0x83, 0x1b, 0xd0, 0x60, 0xc5, 0x06, 0xba,
	0x3a, 0x61, 0x45, 0x91, 0xbb, 0x73, 0xa2, 0x09, 0x4e, 0x54, 0x4b, 0x4e, 0x14, 0x0a, 0x59, 0x23,
	0x7e, 0xf2, 0xd0, 0x52, 0x66, 0xa4, 0x7d, 0xac, 0x0d, 0x9e, 0x81, 0xf6, 0xc1, 0x61, 0x6a, 0x2f,
	0x3c, 0x72, 0xdb, 0x7b, 0xb1, 0x90, 0xc9, 0x7d, 0xfb, 0x0e, 0xe3, 0x47, 0x0f, 0x55, 0x8c, 0x3c,
	0x1a, 0x65, 0xfc, 0x9e, 0x4f, 0x97, 0xdf, 0xa2, 0x65, 0xe3, 0xe8, 0x19, 0xff, 0x76, 0x99, 0x88,
	0x4d, 0x27, 0x3a, 0x7e, 0x9b, 0x53, 0xf6, 0xcf, 0xd1, 0x71, 0x04, 0xbf, 0xf5, 0xd0, 0x8c, 0x79,
	0xbe, 0x0d, 0x5e, 0xc0, 0x4f, 0xe8, 0xe1, 0xa1, 0xd1, 0x2a, 0xbd, 0x20, 0x68, 0x01, 0x8b, 0xf6,
	0x2c, 0xbc, 0x7d, 0xbb, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x60, 0x1c, 0x29, 0x9e,
	0x13, 0x00, 0x00,
}