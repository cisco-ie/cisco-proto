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
// source: qos_show_policy_st.proto

package cisco_ios_xr_skp_qos_oper_platform_qos_nodes_node_interfaces_interface_input

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

type QosShowPolicySt_KEYS struct {
	NodeName             string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	InterfaceName        string   `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosShowPolicySt_KEYS) Reset()         { *m = QosShowPolicySt_KEYS{} }
func (m *QosShowPolicySt_KEYS) String() string { return proto.CompactTextString(m) }
func (*QosShowPolicySt_KEYS) ProtoMessage()    {}
func (*QosShowPolicySt_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{0}
}

func (m *QosShowPolicySt_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowPolicySt_KEYS.Unmarshal(m, b)
}
func (m *QosShowPolicySt_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowPolicySt_KEYS.Marshal(b, m, deterministic)
}
func (m *QosShowPolicySt_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowPolicySt_KEYS.Merge(m, src)
}
func (m *QosShowPolicySt_KEYS) XXX_Size() int {
	return xxx_messageInfo_QosShowPolicySt_KEYS.Size(m)
}
func (m *QosShowPolicySt_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowPolicySt_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowPolicySt_KEYS proto.InternalMessageInfo

func (m *QosShowPolicySt_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *QosShowPolicySt_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type QosShowEaHeaderSt struct {
	InterfaceName        string   `protobuf:"bytes,1,opt,name=interface_name,json=interfaceName,proto3" json:"interface_name,omitempty"`
	PolicyName           string   `protobuf:"bytes,2,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Direction            string   `protobuf:"bytes,3,opt,name=direction,proto3" json:"direction,omitempty"`
	Classes              uint32   `protobuf:"varint,4,opt,name=classes,proto3" json:"classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosShowEaHeaderSt) Reset()         { *m = QosShowEaHeaderSt{} }
func (m *QosShowEaHeaderSt) String() string { return proto.CompactTextString(m) }
func (*QosShowEaHeaderSt) ProtoMessage()    {}
func (*QosShowEaHeaderSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{1}
}

func (m *QosShowEaHeaderSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowEaHeaderSt.Unmarshal(m, b)
}
func (m *QosShowEaHeaderSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowEaHeaderSt.Marshal(b, m, deterministic)
}
func (m *QosShowEaHeaderSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowEaHeaderSt.Merge(m, src)
}
func (m *QosShowEaHeaderSt) XXX_Size() int {
	return xxx_messageInfo_QosShowEaHeaderSt.Size(m)
}
func (m *QosShowEaHeaderSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowEaHeaderSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowEaHeaderSt proto.InternalMessageInfo

func (m *QosShowEaHeaderSt) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *QosShowEaHeaderSt) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *QosShowEaHeaderSt) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *QosShowEaHeaderSt) GetClasses() uint32 {
	if m != nil {
		return m.Classes
	}
	return 0
}

type QosParam struct {
	Value                uint32   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	Unit                 string   `protobuf:"bytes,2,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosParam) Reset()         { *m = QosParam{} }
func (m *QosParam) String() string { return proto.CompactTextString(m) }
func (*QosParam) ProtoMessage()    {}
func (*QosParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{2}
}

func (m *QosParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosParam.Unmarshal(m, b)
}
func (m *QosParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosParam.Marshal(b, m, deterministic)
}
func (m *QosParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosParam.Merge(m, src)
}
func (m *QosParam) XXX_Size() int {
	return xxx_messageInfo_QosParam.Size(m)
}
func (m *QosParam) XXX_DiscardUnknown() {
	xxx_messageInfo_QosParam.DiscardUnknown(m)
}

var xxx_messageInfo_QosParam proto.InternalMessageInfo

func (m *QosParam) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *QosParam) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

type QosShowIfQosParamSt struct {
	InterfaceConfigRate  *QosParam `protobuf:"bytes,1,opt,name=interface_config_rate,json=interfaceConfigRate,proto3" json:"interface_config_rate,omitempty"`
	InterfaceProgramRate *QosParam `protobuf:"bytes,2,opt,name=interface_program_rate,json=interfaceProgramRate,proto3" json:"interface_program_rate,omitempty"`
	PortShaperRate       *QosParam `protobuf:"bytes,3,opt,name=port_shaper_rate,json=portShaperRate,proto3" json:"port_shaper_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QosShowIfQosParamSt) Reset()         { *m = QosShowIfQosParamSt{} }
func (m *QosShowIfQosParamSt) String() string { return proto.CompactTextString(m) }
func (*QosShowIfQosParamSt) ProtoMessage()    {}
func (*QosShowIfQosParamSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{3}
}

func (m *QosShowIfQosParamSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowIfQosParamSt.Unmarshal(m, b)
}
func (m *QosShowIfQosParamSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowIfQosParamSt.Marshal(b, m, deterministic)
}
func (m *QosShowIfQosParamSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowIfQosParamSt.Merge(m, src)
}
func (m *QosShowIfQosParamSt) XXX_Size() int {
	return xxx_messageInfo_QosShowIfQosParamSt.Size(m)
}
func (m *QosShowIfQosParamSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowIfQosParamSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowIfQosParamSt proto.InternalMessageInfo

func (m *QosShowIfQosParamSt) GetInterfaceConfigRate() *QosParam {
	if m != nil {
		return m.InterfaceConfigRate
	}
	return nil
}

func (m *QosShowIfQosParamSt) GetInterfaceProgramRate() *QosParam {
	if m != nil {
		return m.InterfaceProgramRate
	}
	return nil
}

func (m *QosShowIfQosParamSt) GetPortShaperRate() *QosParam {
	if m != nil {
		return m.PortShaperRate
	}
	return nil
}

type QosShowQueueParamsSt struct {
	QueueId              uint32   `protobuf:"varint,1,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	QueueType            string   `protobuf:"bytes,2,opt,name=queue_type,json=queueType,proto3" json:"queue_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QosShowQueueParamsSt) Reset()         { *m = QosShowQueueParamsSt{} }
func (m *QosShowQueueParamsSt) String() string { return proto.CompactTextString(m) }
func (*QosShowQueueParamsSt) ProtoMessage()    {}
func (*QosShowQueueParamsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{4}
}

func (m *QosShowQueueParamsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowQueueParamsSt.Unmarshal(m, b)
}
func (m *QosShowQueueParamsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowQueueParamsSt.Marshal(b, m, deterministic)
}
func (m *QosShowQueueParamsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowQueueParamsSt.Merge(m, src)
}
func (m *QosShowQueueParamsSt) XXX_Size() int {
	return xxx_messageInfo_QosShowQueueParamsSt.Size(m)
}
func (m *QosShowQueueParamsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowQueueParamsSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowQueueParamsSt proto.InternalMessageInfo

func (m *QosShowQueueParamsSt) GetQueueId() uint32 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *QosShowQueueParamsSt) GetQueueType() string {
	if m != nil {
		return m.QueueType
	}
	return ""
}

type QosShowEaShaperParamsSt struct {
	Pir                  *QosParam `protobuf:"bytes,1,opt,name=pir,proto3" json:"pir,omitempty"`
	Pbs                  *QosParam `protobuf:"bytes,2,opt,name=pbs,proto3" json:"pbs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QosShowEaShaperParamsSt) Reset()         { *m = QosShowEaShaperParamsSt{} }
func (m *QosShowEaShaperParamsSt) String() string { return proto.CompactTextString(m) }
func (*QosShowEaShaperParamsSt) ProtoMessage()    {}
func (*QosShowEaShaperParamsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{5}
}

func (m *QosShowEaShaperParamsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowEaShaperParamsSt.Unmarshal(m, b)
}
func (m *QosShowEaShaperParamsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowEaShaperParamsSt.Marshal(b, m, deterministic)
}
func (m *QosShowEaShaperParamsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowEaShaperParamsSt.Merge(m, src)
}
func (m *QosShowEaShaperParamsSt) XXX_Size() int {
	return xxx_messageInfo_QosShowEaShaperParamsSt.Size(m)
}
func (m *QosShowEaShaperParamsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowEaShaperParamsSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowEaShaperParamsSt proto.InternalMessageInfo

func (m *QosShowEaShaperParamsSt) GetPir() *QosParam {
	if m != nil {
		return m.Pir
	}
	return nil
}

func (m *QosShowEaShaperParamsSt) GetPbs() *QosParam {
	if m != nil {
		return m.Pbs
	}
	return nil
}

type QosShowEaWfqParamsSt struct {
	Bandwidth            *QosParam `protobuf:"bytes,1,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	SumOfBandwidth       *QosParam `protobuf:"bytes,2,opt,name=sum_of_bandwidth,json=sumOfBandwidth,proto3" json:"sum_of_bandwidth,omitempty"`
	ExcessRatio          uint32    `protobuf:"varint,3,opt,name=excess_ratio,json=excessRatio,proto3" json:"excess_ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QosShowEaWfqParamsSt) Reset()         { *m = QosShowEaWfqParamsSt{} }
func (m *QosShowEaWfqParamsSt) String() string { return proto.CompactTextString(m) }
func (*QosShowEaWfqParamsSt) ProtoMessage()    {}
func (*QosShowEaWfqParamsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{6}
}

func (m *QosShowEaWfqParamsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowEaWfqParamsSt.Unmarshal(m, b)
}
func (m *QosShowEaWfqParamsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowEaWfqParamsSt.Marshal(b, m, deterministic)
}
func (m *QosShowEaWfqParamsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowEaWfqParamsSt.Merge(m, src)
}
func (m *QosShowEaWfqParamsSt) XXX_Size() int {
	return xxx_messageInfo_QosShowEaWfqParamsSt.Size(m)
}
func (m *QosShowEaWfqParamsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowEaWfqParamsSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowEaWfqParamsSt proto.InternalMessageInfo

func (m *QosShowEaWfqParamsSt) GetBandwidth() *QosParam {
	if m != nil {
		return m.Bandwidth
	}
	return nil
}

func (m *QosShowEaWfqParamsSt) GetSumOfBandwidth() *QosParam {
	if m != nil {
		return m.SumOfBandwidth
	}
	return nil
}

func (m *QosShowEaWfqParamsSt) GetExcessRatio() uint32 {
	if m != nil {
		return m.ExcessRatio
	}
	return 0
}

type QosShowWfqParamsSt struct {
	CommittedWeight      *QosParam             `protobuf:"bytes,1,opt,name=committed_weight,json=committedWeight,proto3" json:"committed_weight,omitempty"`
	ExcessWeight         uint32                `protobuf:"varint,2,opt,name=excess_weight,json=excessWeight,proto3" json:"excess_weight,omitempty"`
	ProgrammedWfq        *QosShowEaWfqParamsSt `protobuf:"bytes,3,opt,name=programmed_wfq,json=programmedWfq,proto3" json:"programmed_wfq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QosShowWfqParamsSt) Reset()         { *m = QosShowWfqParamsSt{} }
func (m *QosShowWfqParamsSt) String() string { return proto.CompactTextString(m) }
func (*QosShowWfqParamsSt) ProtoMessage()    {}
func (*QosShowWfqParamsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{7}
}

func (m *QosShowWfqParamsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowWfqParamsSt.Unmarshal(m, b)
}
func (m *QosShowWfqParamsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowWfqParamsSt.Marshal(b, m, deterministic)
}
func (m *QosShowWfqParamsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowWfqParamsSt.Merge(m, src)
}
func (m *QosShowWfqParamsSt) XXX_Size() int {
	return xxx_messageInfo_QosShowWfqParamsSt.Size(m)
}
func (m *QosShowWfqParamsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowWfqParamsSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowWfqParamsSt proto.InternalMessageInfo

func (m *QosShowWfqParamsSt) GetCommittedWeight() *QosParam {
	if m != nil {
		return m.CommittedWeight
	}
	return nil
}

func (m *QosShowWfqParamsSt) GetExcessWeight() uint32 {
	if m != nil {
		return m.ExcessWeight
	}
	return 0
}

func (m *QosShowWfqParamsSt) GetProgrammedWfq() *QosShowEaWfqParamsSt {
	if m != nil {
		return m.ProgrammedWfq
	}
	return nil
}

type QosShowPoliceParamsSt struct {
	PolicerId            uint32    `protobuf:"varint,1,opt,name=policer_id,json=policerId,proto3" json:"policer_id,omitempty"`
	PolicerType          string    `protobuf:"bytes,2,opt,name=policer_type,json=policerType,proto3" json:"policer_type,omitempty"`
	Cir                  *QosParam `protobuf:"bytes,3,opt,name=cir,proto3" json:"cir,omitempty"`
	Cbs                  *QosParam `protobuf:"bytes,4,opt,name=cbs,proto3" json:"cbs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QosShowPoliceParamsSt) Reset()         { *m = QosShowPoliceParamsSt{} }
func (m *QosShowPoliceParamsSt) String() string { return proto.CompactTextString(m) }
func (*QosShowPoliceParamsSt) ProtoMessage()    {}
func (*QosShowPoliceParamsSt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{8}
}

func (m *QosShowPoliceParamsSt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowPoliceParamsSt.Unmarshal(m, b)
}
func (m *QosShowPoliceParamsSt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowPoliceParamsSt.Marshal(b, m, deterministic)
}
func (m *QosShowPoliceParamsSt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowPoliceParamsSt.Merge(m, src)
}
func (m *QosShowPoliceParamsSt) XXX_Size() int {
	return xxx_messageInfo_QosShowPoliceParamsSt.Size(m)
}
func (m *QosShowPoliceParamsSt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowPoliceParamsSt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowPoliceParamsSt proto.InternalMessageInfo

func (m *QosShowPoliceParamsSt) GetPolicerId() uint32 {
	if m != nil {
		return m.PolicerId
	}
	return 0
}

func (m *QosShowPoliceParamsSt) GetPolicerType() string {
	if m != nil {
		return m.PolicerType
	}
	return ""
}

func (m *QosShowPoliceParamsSt) GetCir() *QosParam {
	if m != nil {
		return m.Cir
	}
	return nil
}

func (m *QosShowPoliceParamsSt) GetCbs() *QosParam {
	if m != nil {
		return m.Cbs
	}
	return nil
}

type MarkAction struct {
	MarkValue            uint32   `protobuf:"varint,1,opt,name=mark_value,json=markValue,proto3" json:"mark_value,omitempty"`
	ActionOpcode         string   `protobuf:"bytes,2,opt,name=action_opcode,json=actionOpcode,proto3" json:"action_opcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarkAction) Reset()         { *m = MarkAction{} }
func (m *MarkAction) String() string { return proto.CompactTextString(m) }
func (*MarkAction) ProtoMessage()    {}
func (*MarkAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{9}
}

func (m *MarkAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkAction.Unmarshal(m, b)
}
func (m *MarkAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkAction.Marshal(b, m, deterministic)
}
func (m *MarkAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkAction.Merge(m, src)
}
func (m *MarkAction) XXX_Size() int {
	return xxx_messageInfo_MarkAction.Size(m)
}
func (m *MarkAction) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkAction.DiscardUnknown(m)
}

var xxx_messageInfo_MarkAction proto.InternalMessageInfo

func (m *MarkAction) GetMarkValue() uint32 {
	if m != nil {
		return m.MarkValue
	}
	return 0
}

func (m *MarkAction) GetActionOpcode() string {
	if m != nil {
		return m.ActionOpcode
	}
	return ""
}

type MarkActionPerCategory struct {
	MarkDetail           []*MarkAction `protobuf:"bytes,1,rep,name=mark_detail,json=markDetail,proto3" json:"mark_detail,omitempty"`
	ActionType           string        `protobuf:"bytes,2,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MarkActionPerCategory) Reset()         { *m = MarkActionPerCategory{} }
func (m *MarkActionPerCategory) String() string { return proto.CompactTextString(m) }
func (*MarkActionPerCategory) ProtoMessage()    {}
func (*MarkActionPerCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{10}
}

func (m *MarkActionPerCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarkActionPerCategory.Unmarshal(m, b)
}
func (m *MarkActionPerCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarkActionPerCategory.Marshal(b, m, deterministic)
}
func (m *MarkActionPerCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarkActionPerCategory.Merge(m, src)
}
func (m *MarkActionPerCategory) XXX_Size() int {
	return xxx_messageInfo_MarkActionPerCategory.Size(m)
}
func (m *MarkActionPerCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_MarkActionPerCategory.DiscardUnknown(m)
}

var xxx_messageInfo_MarkActionPerCategory proto.InternalMessageInfo

func (m *MarkActionPerCategory) GetMarkDetail() []*MarkAction {
	if m != nil {
		return m.MarkDetail
	}
	return nil
}

func (m *MarkActionPerCategory) GetActionType() string {
	if m != nil {
		return m.ActionType
	}
	return ""
}

type QosClassMarkActions struct {
	MarkOnly             *MarkActionPerCategory `protobuf:"bytes,1,opt,name=mark_only,json=markOnly,proto3" json:"mark_only,omitempty"`
	PoliceConform        *MarkActionPerCategory `protobuf:"bytes,2,opt,name=police_conform,json=policeConform,proto3" json:"police_conform,omitempty"`
	PoliceExceed         *MarkActionPerCategory `protobuf:"bytes,3,opt,name=police_exceed,json=policeExceed,proto3" json:"police_exceed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *QosClassMarkActions) Reset()         { *m = QosClassMarkActions{} }
func (m *QosClassMarkActions) String() string { return proto.CompactTextString(m) }
func (*QosClassMarkActions) ProtoMessage()    {}
func (*QosClassMarkActions) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{11}
}

func (m *QosClassMarkActions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosClassMarkActions.Unmarshal(m, b)
}
func (m *QosClassMarkActions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosClassMarkActions.Marshal(b, m, deterministic)
}
func (m *QosClassMarkActions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosClassMarkActions.Merge(m, src)
}
func (m *QosClassMarkActions) XXX_Size() int {
	return xxx_messageInfo_QosClassMarkActions.Size(m)
}
func (m *QosClassMarkActions) XXX_DiscardUnknown() {
	xxx_messageInfo_QosClassMarkActions.DiscardUnknown(m)
}

var xxx_messageInfo_QosClassMarkActions proto.InternalMessageInfo

func (m *QosClassMarkActions) GetMarkOnly() *MarkActionPerCategory {
	if m != nil {
		return m.MarkOnly
	}
	return nil
}

func (m *QosClassMarkActions) GetPoliceConform() *MarkActionPerCategory {
	if m != nil {
		return m.PoliceConform
	}
	return nil
}

func (m *QosClassMarkActions) GetPoliceExceed() *MarkActionPerCategory {
	if m != nil {
		return m.PoliceExceed
	}
	return nil
}

type QosShowPclassStItem struct {
	ClassLevel           uint32                   `protobuf:"varint,1,opt,name=class_level,json=classLevel,proto3" json:"class_level,omitempty"`
	ClassName            string                   `protobuf:"bytes,2,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	Queue                *QosShowQueueParamsSt    `protobuf:"bytes,3,opt,name=queue,proto3" json:"queue,omitempty"`
	Shape                *QosShowEaShaperParamsSt `protobuf:"bytes,4,opt,name=shape,proto3" json:"shape,omitempty"`
	Wfq                  *QosShowWfqParamsSt      `protobuf:"bytes,5,opt,name=wfq,proto3" json:"wfq,omitempty"`
	Police               *QosShowPoliceParamsSt   `protobuf:"bytes,6,opt,name=police,proto3" json:"police,omitempty"`
	Marking              *QosClassMarkActions     `protobuf:"bytes,7,opt,name=marking,proto3" json:"marking,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *QosShowPclassStItem) Reset()         { *m = QosShowPclassStItem{} }
func (m *QosShowPclassStItem) String() string { return proto.CompactTextString(m) }
func (*QosShowPclassStItem) ProtoMessage()    {}
func (*QosShowPclassStItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{12}
}

func (m *QosShowPclassStItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowPclassStItem.Unmarshal(m, b)
}
func (m *QosShowPclassStItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowPclassStItem.Marshal(b, m, deterministic)
}
func (m *QosShowPclassStItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowPclassStItem.Merge(m, src)
}
func (m *QosShowPclassStItem) XXX_Size() int {
	return xxx_messageInfo_QosShowPclassStItem.Size(m)
}
func (m *QosShowPclassStItem) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowPclassStItem.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowPclassStItem proto.InternalMessageInfo

func (m *QosShowPclassStItem) GetClassLevel() uint32 {
	if m != nil {
		return m.ClassLevel
	}
	return 0
}

func (m *QosShowPclassStItem) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *QosShowPclassStItem) GetQueue() *QosShowQueueParamsSt {
	if m != nil {
		return m.Queue
	}
	return nil
}

func (m *QosShowPclassStItem) GetShape() *QosShowEaShaperParamsSt {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *QosShowPclassStItem) GetWfq() *QosShowWfqParamsSt {
	if m != nil {
		return m.Wfq
	}
	return nil
}

func (m *QosShowPclassStItem) GetPolice() *QosShowPoliceParamsSt {
	if m != nil {
		return m.Police
	}
	return nil
}

func (m *QosShowPclassStItem) GetMarking() *QosClassMarkActions {
	if m != nil {
		return m.Marking
	}
	return nil
}

type QosShowPclassStEntry struct {
	QosShowPclassSt      []*QosShowPclassStItem `protobuf:"bytes,8,rep,name=qos_show_pclass_st,json=qosShowPclassSt,proto3" json:"qos_show_pclass_st,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *QosShowPclassStEntry) Reset()         { *m = QosShowPclassStEntry{} }
func (m *QosShowPclassStEntry) String() string { return proto.CompactTextString(m) }
func (*QosShowPclassStEntry) ProtoMessage()    {}
func (*QosShowPclassStEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{13}
}

func (m *QosShowPclassStEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowPclassStEntry.Unmarshal(m, b)
}
func (m *QosShowPclassStEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowPclassStEntry.Marshal(b, m, deterministic)
}
func (m *QosShowPclassStEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowPclassStEntry.Merge(m, src)
}
func (m *QosShowPclassStEntry) XXX_Size() int {
	return xxx_messageInfo_QosShowPclassStEntry.Size(m)
}
func (m *QosShowPclassStEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowPclassStEntry.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowPclassStEntry proto.InternalMessageInfo

func (m *QosShowPclassStEntry) GetQosShowPclassSt() []*QosShowPclassStItem {
	if m != nil {
		return m.QosShowPclassSt
	}
	return nil
}

type QosShowPolicySt struct {
	Header                *QosShowEaHeaderSt    `protobuf:"bytes,50,opt,name=header,proto3" json:"header,omitempty"`
	InterfaceParameters   *QosShowIfQosParamSt  `protobuf:"bytes,51,opt,name=interface_parameters,json=interfaceParameters,proto3" json:"interface_parameters,omitempty"`
	SkywarpQosPolicyClass *QosShowPclassStEntry `protobuf:"bytes,52,opt,name=skywarp_qos_policy_class,json=skywarpQosPolicyClass,proto3" json:"skywarp_qos_policy_class,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *QosShowPolicySt) Reset()         { *m = QosShowPolicySt{} }
func (m *QosShowPolicySt) String() string { return proto.CompactTextString(m) }
func (*QosShowPolicySt) ProtoMessage()    {}
func (*QosShowPolicySt) Descriptor() ([]byte, []int) {
	return fileDescriptor_88520c2e06938e72, []int{14}
}

func (m *QosShowPolicySt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QosShowPolicySt.Unmarshal(m, b)
}
func (m *QosShowPolicySt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QosShowPolicySt.Marshal(b, m, deterministic)
}
func (m *QosShowPolicySt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QosShowPolicySt.Merge(m, src)
}
func (m *QosShowPolicySt) XXX_Size() int {
	return xxx_messageInfo_QosShowPolicySt.Size(m)
}
func (m *QosShowPolicySt) XXX_DiscardUnknown() {
	xxx_messageInfo_QosShowPolicySt.DiscardUnknown(m)
}

var xxx_messageInfo_QosShowPolicySt proto.InternalMessageInfo

func (m *QosShowPolicySt) GetHeader() *QosShowEaHeaderSt {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *QosShowPolicySt) GetInterfaceParameters() *QosShowIfQosParamSt {
	if m != nil {
		return m.InterfaceParameters
	}
	return nil
}

func (m *QosShowPolicySt) GetSkywarpQosPolicyClass() *QosShowPclassStEntry {
	if m != nil {
		return m.SkywarpQosPolicyClass
	}
	return nil
}

func init() {
	proto.RegisterType((*QosShowPolicySt_KEYS)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_policy_st_KEYS")
	proto.RegisterType((*QosShowEaHeaderSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_ea_header_st")
	proto.RegisterType((*QosParam)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_param")
	proto.RegisterType((*QosShowIfQosParamSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_if_qos_param_st")
	proto.RegisterType((*QosShowQueueParamsSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_queue_params_st")
	proto.RegisterType((*QosShowEaShaperParamsSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_ea_shaper_params_st")
	proto.RegisterType((*QosShowEaWfqParamsSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_ea_wfq_params_st")
	proto.RegisterType((*QosShowWfqParamsSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_wfq_params_st")
	proto.RegisterType((*QosShowPoliceParamsSt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_police_params_st")
	proto.RegisterType((*MarkAction)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.mark_action")
	proto.RegisterType((*MarkActionPerCategory)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.mark_action_per_category")
	proto.RegisterType((*QosClassMarkActions)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_class_mark_actions")
	proto.RegisterType((*QosShowPclassStItem)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_pclass_st_item")
	proto.RegisterType((*QosShowPclassStEntry)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_pclass_st_entry")
	proto.RegisterType((*QosShowPolicySt)(nil), "cisco_ios_xr_skp_qos_oper.platform_qos.nodes.node.interfaces.interface.input.qos_show_policy_st")
}

func init() { proto.RegisterFile("qos_show_policy_st.proto", fileDescriptor_88520c2e06938e72) }

var fileDescriptor_88520c2e06938e72 = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x58, 0xcb, 0x6e, 0x1b, 0x37,
	0x17, 0xc6, 0x48, 0x91, 0x6d, 0x1d, 0x45, 0x4a, 0xc0, 0xdf, 0xce, 0x3f, 0x41, 0x63, 0xd4, 0x9d,
	0xa0, 0x80, 0x57, 0x5a, 0x38, 0xed, 0x0b, 0xd4, 0xcd, 0x22, 0x68, 0x50, 0x3b, 0xe3, 0xa0, 0x46,
	0x16, 0x05, 0x31, 0x9a, 0xe1, 0x48, 0xac, 0x35, 0xc3, 0x11, 0x49, 0x59, 0x51, 0x5b, 0xf4, 0xba,
	0x2c, 0xda, 0x45, 0x37, 0x45, 0x57, 0x7d, 0x84, 0x2e, 0xba, 0x28, 0x0a, 0xe4, 0x1d, 0xba, 0xc8,
	0xab, 0xf4, 0x01, 0x0a, 0x1e, 0x52, 0x23, 0xfa, 0xd2, 0x9d, 0x27, 0x9b, 0x40, 0xf3, 0x91, 0xe1,
	0xf7, 0x9d, 0x2b, 0x0f, 0x0d, 0xe1, 0x4c, 0x28, 0xaa, 0x26, 0x62, 0x41, 0x2b, 0x31, 0xe5, 0xe9,
	0x92, 0x2a, 0x3d, 0xac, 0xa4, 0xd0, 0x82, 0x3c, 0x4d, 0xb9, 0x4a, 0x05, 0xe5, 0x42, 0xd1, 0x97,
	0x92, 0xaa, 0xb3, 0x8a, 0x9a, 0xad, 0xa2, 0x62, 0x72, 0x58, 0x4d, 0x13, 0x9d, 0x0b, 0x59, 0x18,
	0x64, 0x58, 0x8a, 0x8c, 0xd9, 0x7f, 0x87, 0xbc, 0xd4, 0x4c, 0xe6, 0x49, 0xca, 0xd4, 0xfa, 0xe7,
	0x90, 0x97, 0xd5, 0x5c, 0x47, 0x9f, 0xc2, 0xff, 0xaf, 0x32, 0xd1, 0x8f, 0x1e, 0xbf, 0x38, 0x21,
	0x6f, 0x41, 0xd7, 0xfc, 0x6f, 0x5a, 0x26, 0x05, 0x0b, 0x83, 0xbd, 0x60, 0xbf, 0x1b, 0x6f, 0x19,
	0xe0, 0xe3, 0xa4, 0x60, 0xe4, 0x5d, 0x18, 0xd4, 0x47, 0xd9, 0x1d, 0x2d, 0xdc, 0xd1, 0xaf, 0x51,
	0xb3, 0x2d, 0xfa, 0x25, 0x80, 0x9d, 0xfa, 0x7c, 0x96, 0xd0, 0x09, 0x4b, 0x32, 0x26, 0xa9, 0xd2,
	0xd7, 0x1c, 0x10, 0x5c, 0x73, 0x00, 0x79, 0x1b, 0x7a, 0x4e, 0x96, 0x47, 0x02, 0x16, 0xc2, 0x0d,
	0x0f, 0xa0, 0x9b, 0x71, 0xc9, 0x52, 0xcd, 0x45, 0x19, 0xb6, 0x71, 0x79, 0x0d, 0x90, 0x10, 0x36,
	0xd3, 0x69, 0xa2, 0x14, 0x53, 0xe1, 0xad, 0xbd, 0x60, 0xbf, 0x1f, 0xaf, 0x3e, 0xa3, 0xf7, 0xa1,
	0x6b, 0x84, 0x55, 0x89, 0x4c, 0x0a, 0xb2, 0x0d, 0x9d, 0xf3, 0x64, 0x3a, 0xb7, 0x1a, 0xfa, 0xb1,
	0xfd, 0x20, 0x04, 0x6e, 0xcd, 0x4b, 0xae, 0x1d, 0x29, 0xfe, 0x8e, 0x5e, 0xb5, 0xbd, 0xd0, 0xf0,
	0x9c, 0xd6, 0x67, 0x18, 0x9b, 0x7e, 0x08, 0x60, 0x67, 0x6d, 0x54, 0x2a, 0xca, 0x9c, 0x8f, 0xa9,
	0x4c, 0xb4, 0x3d, 0xb7, 0x77, 0x70, 0x3a, 0xbc, 0xc9, 0xd8, 0x0d, 0x6b, 0xee, 0xf8, 0x7f, 0xf5,
	0xd2, 0x21, 0x92, 0xc6, 0x89, 0x66, 0xe4, 0xc7, 0x00, 0xee, 0xad, 0xd5, 0x54, 0x52, 0x8c, 0x8d,
	0x4c, 0x94, 0xd3, 0x6a, 0x56, 0xce, 0x76, 0xbd, 0x74, 0x6c, 0x59, 0x51, 0xcf, 0xb7, 0x01, 0xdc,
	0xad, 0x84, 0xd4, 0x54, 0x4d, 0x92, 0x8a, 0x49, 0xab, 0xa4, 0xdd, 0xac, 0x92, 0x81, 0x21, 0x3c,
	0x41, 0x3e, 0xa3, 0x21, 0x7a, 0xee, 0x45, 0x6f, 0x36, 0x67, 0x73, 0x66, 0xf7, 0x29, 0x13, 0xbd,
	0xfb, 0xb0, 0x65, 0x21, 0x9e, 0xb9, 0x3c, 0xd8, 0xc4, 0xef, 0x27, 0x19, 0xd9, 0x05, 0xb0, 0x4b,
	0x7a, 0x59, 0xad, 0x92, 0xb0, 0x8b, 0xc8, 0xf3, 0x65, 0xc5, 0xa2, 0x7f, 0x02, 0x78, 0xe0, 0x67,
	0xb9, 0x33, 0x70, 0x7d, 0x34, 0x87, 0x76, 0xc5, 0x65, 0xd3, 0x59, 0x60, 0x38, 0x90, 0x6a, 0xa4,
	0x9a, 0x8e, 0xb0, 0xe1, 0x88, 0xfe, 0x6a, 0xc1, 0x7d, 0xdf, 0xec, 0x45, 0x3e, 0xf3, 0x6c, 0x9e,
	0x43, 0x77, 0x94, 0x94, 0xd9, 0x82, 0x67, 0x7a, 0xd2, 0xb4, 0xe5, 0x6b, 0x26, 0xcc, 0x32, 0x35,
	0x2f, 0xa8, 0xc8, 0xe9, 0x9a, 0xbe, 0x61, 0x6f, 0x0c, 0xd4, 0xbc, 0x38, 0xca, 0x3f, 0xa8, 0x35,
	0xbc, 0x03, 0xb7, 0xd9, 0xcb, 0x94, 0x29, 0x65, 0x72, 0x9c, 0x0b, 0x4c, 0xf2, 0x7e, 0xdc, 0xb3,
	0x58, 0x6c, 0xa0, 0xe8, 0x75, 0x0b, 0xee, 0xd5, 0xbe, 0xbb, 0xe8, 0xb8, 0xef, 0x02, 0xb8, 0x9b,
	0x8a, 0xa2, 0xe0, 0x5a, 0xb3, 0x8c, 0x2e, 0x18, 0x1f, 0x4f, 0x74, 0xd3, 0x0e, 0xbc, 0x53, 0x13,
	0x9e, 0x22, 0x1f, 0x79, 0x08, 0x7d, 0x67, 0x82, 0x13, 0xd0, 0x42, 0x1b, 0x9c, 0x5d, 0x6e, 0xd3,
	0x4f, 0x01, 0x0c, 0x5c, 0x5f, 0x29, 0x8c, 0xd4, 0x7c, 0xe6, 0xea, 0x79, 0x7c, 0xf3, 0x3a, 0xaf,
	0x4d, 0xb2, 0xb8, 0xbf, 0xa6, 0x3f, 0xcd, 0x67, 0xd1, 0x1f, 0x7e, 0x46, 0xe2, 0x25, 0xe1, 0x17,
	0xf8, 0x2e, 0xd8, 0x8b, 0x83, 0xc9, 0x75, 0x89, 0x77, 0x1d, 0xf2, 0x24, 0x33, 0x51, 0x5b, 0x2d,
	0x7b, 0x65, 0xde, 0x73, 0x98, 0x29, 0x74, 0x53, 0x5c, 0x29, 0x97, 0x4d, 0x37, 0x2d, 0xc3, 0x81,
	0x54, 0x23, 0x7b, 0x6b, 0x35, 0x4a, 0x35, 0x52, 0xd1, 0x33, 0xe8, 0x15, 0x89, 0x3c, 0xa3, 0x89,
	0xbd, 0x33, 0x77, 0x01, 0xf0, 0xd3, 0xbf, 0x11, 0xbb, 0x06, 0xf9, 0x04, 0x6f, 0xc5, 0x87, 0xd0,
	0xb7, 0x1b, 0xa9, 0xa8, 0x52, 0x91, 0xad, 0xfc, 0x74, 0xdb, 0x82, 0x47, 0x88, 0x45, 0x7f, 0x06,
	0x10, 0x7a, 0x67, 0x52, 0xd3, 0x0e, 0xd3, 0x44, 0xb3, 0xb1, 0x90, 0x4b, 0xf2, 0xb9, 0xe3, 0xcb,
	0x98, 0x4e, 0xf8, 0x34, 0x0c, 0xf6, 0xda, 0xfb, 0xbd, 0x83, 0x17, 0x37, 0x6b, 0xa2, 0x47, 0x1e,
	0xa3, 0x39, 0x1f, 0x22, 0x99, 0x99, 0x27, 0x9c, 0x24, 0x2f, 0xc6, 0x60, 0x21, 0xec, 0xe5, 0xaf,
	0xda, 0xb6, 0x30, 0x71, 0x4e, 0xa0, 0xde, 0x31, 0x8a, 0x7c, 0x1f, 0x00, 0xfa, 0x81, 0x8a, 0x72,
	0xba, 0x74, 0x15, 0x99, 0x37, 0x26, 0xfb, 0x82, 0xcf, 0xe2, 0x2d, 0xb3, 0x72, 0x54, 0x4e, 0x97,
	0xe6, 0x5a, 0x1f, 0xb8, 0xd4, 0x36, 0x13, 0x86, 0x90, 0x85, 0x6b, 0x6f, 0x6f, 0x4a, 0x4a, 0xdf,
	0xb2, 0x1f, 0x5a, 0x72, 0x33, 0xf4, 0x38, 0x84, 0x9a, 0xe6, 0xc0, 0x32, 0x57, 0x1e, 0x6f, 0x4a,
	0x8e, 0x2b, 0xda, 0xc7, 0xc8, 0x1d, 0xbd, 0xee, 0xf8, 0x03, 0xad, 0x0d, 0xa2, 0xd2, 0x94, 0x6b,
	0x56, 0x98, 0xd8, 0x5b, 0x60, 0xca, 0xce, 0xd9, 0xd4, 0x65, 0x36, 0x20, 0xf4, 0xd4, 0x20, 0x26,
	0xf3, 0xed, 0x06, 0x6f, 0xd6, 0xec, 0x22, 0x82, 0xa3, 0xe6, 0x97, 0xd0, 0xc1, 0x3b, 0xbf, 0x19,
	0x03, 0xff, 0x6b, 0x2e, 0x89, 0x2d, 0x29, 0xf9, 0x26, 0x80, 0x0e, 0x0e, 0x16, 0xae, 0x27, 0x7c,
	0xd6, 0x5c, 0x8f, 0xbd, 0x3c, 0xbf, 0xc4, 0x96, 0x98, 0x9c, 0x43, 0xdb, 0xf4, 0xf8, 0x0e, 0xf2,
	0x67, 0x0d, 0xf1, 0x5f, 0x6c, 0xf0, 0x86, 0x90, 0x7c, 0x0d, 0x1b, 0x36, 0xc8, 0xe1, 0x46, 0xa3,
	0xd7, 0xcb, 0xe5, 0x1b, 0x23, 0x76, 0xb4, 0xe4, 0x2b, 0xd8, 0x34, 0xf9, 0xc7, 0xcb, 0x71, 0xb8,
	0xd9, 0x94, 0xf1, 0x57, 0x1b, 0x4e, 0xbc, 0x22, 0x8d, 0x7e, 0x0f, 0xfc, 0x07, 0x61, 0x9d, 0xd5,
	0xac, 0xd4, 0x72, 0x49, 0x7e, 0x0e, 0x80, 0x5c, 0x5d, 0x0c, 0xb7, 0xb0, 0xad, 0xb2, 0xa6, 0x5c,
	0x75, 0xa1, 0xb4, 0xe2, 0x3b, 0x33, 0xa1, 0x4e, 0x26, 0x62, 0x71, 0x8c, 0xf0, 0x89, 0x8e, 0xfe,
	0x6e, 0xfb, 0xa2, 0x56, 0x0f, 0x4b, 0xf2, 0x05, 0x6c, 0xd8, 0x27, 0x60, 0x78, 0x80, 0x7e, 0x4c,
	0x9b, 0x4b, 0xe2, 0xfa, 0xa9, 0x19, 0x3b, 0x4a, 0xf2, 0x6b, 0x00, 0xdb, 0xde, 0x83, 0xc8, 0x44,
	0x99, 0x69, 0x26, 0x55, 0xf8, 0xa8, 0xd1, 0x7a, 0xbe, 0xf4, 0x4a, 0xf4, 0x1e, 0x6b, 0xc7, 0xb5,
	0x04, 0xf2, 0x5b, 0x00, 0xa1, 0x3a, 0x5b, 0x2e, 0x12, 0x69, 0x59, 0x9d, 0xcb, 0xd0, 0x9d, 0xe1,
	0x7b, 0x8d, 0xea, 0xbb, 0x94, 0x4f, 0xf1, 0x8e, 0xd3, 0xf1, 0x4c, 0xa8, 0x63, 0x54, 0x71, 0x68,
	0xd6, 0x47, 0x1b, 0xf8, 0xe7, 0x87, 0x47, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x99, 0x28, 0x55,
	0xf0, 0x9a, 0x10, 0x00, 0x00,
}