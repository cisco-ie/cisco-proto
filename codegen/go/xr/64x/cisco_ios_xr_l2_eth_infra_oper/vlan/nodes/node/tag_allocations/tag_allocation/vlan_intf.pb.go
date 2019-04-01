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
// source: vlan_intf.proto

package cisco_ios_xr_l2_eth_infra_oper_vlan_nodes_node_tag_allocations_tag_allocation

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

type VlanIntf_KEYS struct {
	NodeId               string   `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	FirstTag             uint32   `protobuf:"varint,3,opt,name=first_tag,json=firstTag,proto3" json:"first_tag,omitempty"`
	SecondTag            string   `protobuf:"bytes,4,opt,name=second_tag,json=secondTag,proto3" json:"second_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlanIntf_KEYS) Reset()         { *m = VlanIntf_KEYS{} }
func (m *VlanIntf_KEYS) String() string { return proto.CompactTextString(m) }
func (*VlanIntf_KEYS) ProtoMessage()    {}
func (*VlanIntf_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{0}
}

func (m *VlanIntf_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanIntf_KEYS.Unmarshal(m, b)
}
func (m *VlanIntf_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanIntf_KEYS.Marshal(b, m, deterministic)
}
func (m *VlanIntf_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanIntf_KEYS.Merge(m, src)
}
func (m *VlanIntf_KEYS) XXX_Size() int {
	return xxx_messageInfo_VlanIntf_KEYS.Size(m)
}
func (m *VlanIntf_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanIntf_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_VlanIntf_KEYS proto.InternalMessageInfo

func (m *VlanIntf_KEYS) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *VlanIntf_KEYS) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *VlanIntf_KEYS) GetFirstTag() uint32 {
	if m != nil {
		return m.FirstTag
	}
	return 0
}

func (m *VlanIntf_KEYS) GetSecondTag() string {
	if m != nil {
		return m.SecondTag
	}
	return ""
}

type VlanDoubleTagStack struct {
	OuterTag             uint32   `protobuf:"varint,1,opt,name=outer_tag,json=outerTag,proto3" json:"outer_tag,omitempty"`
	SecondTag            uint32   `protobuf:"varint,2,opt,name=second_tag,json=secondTag,proto3" json:"second_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VlanDoubleTagStack) Reset()         { *m = VlanDoubleTagStack{} }
func (m *VlanDoubleTagStack) String() string { return proto.CompactTextString(m) }
func (*VlanDoubleTagStack) ProtoMessage()    {}
func (*VlanDoubleTagStack) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{1}
}

func (m *VlanDoubleTagStack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanDoubleTagStack.Unmarshal(m, b)
}
func (m *VlanDoubleTagStack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanDoubleTagStack.Marshal(b, m, deterministic)
}
func (m *VlanDoubleTagStack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanDoubleTagStack.Merge(m, src)
}
func (m *VlanDoubleTagStack) XXX_Size() int {
	return xxx_messageInfo_VlanDoubleTagStack.Size(m)
}
func (m *VlanDoubleTagStack) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanDoubleTagStack.DiscardUnknown(m)
}

var xxx_messageInfo_VlanDoubleTagStack proto.InternalMessageInfo

func (m *VlanDoubleTagStack) GetOuterTag() uint32 {
	if m != nil {
		return m.OuterTag
	}
	return 0
}

func (m *VlanDoubleTagStack) GetSecondTag() uint32 {
	if m != nil {
		return m.SecondTag
	}
	return 0
}

type EfpRange struct {
	VlanIdLow            uint32   `protobuf:"varint,1,opt,name=vlan_id_low,json=vlanIdLow,proto3" json:"vlan_id_low,omitempty"`
	VlanIdHigh           uint32   `protobuf:"varint,2,opt,name=vlan_id_high,json=vlanIdHigh,proto3" json:"vlan_id_high,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EfpRange) Reset()         { *m = EfpRange{} }
func (m *EfpRange) String() string { return proto.CompactTextString(m) }
func (*EfpRange) ProtoMessage()    {}
func (*EfpRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{2}
}

func (m *EfpRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EfpRange.Unmarshal(m, b)
}
func (m *EfpRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EfpRange.Marshal(b, m, deterministic)
}
func (m *EfpRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EfpRange.Merge(m, src)
}
func (m *EfpRange) XXX_Size() int {
	return xxx_messageInfo_EfpRange.Size(m)
}
func (m *EfpRange) XXX_DiscardUnknown() {
	xxx_messageInfo_EfpRange.DiscardUnknown(m)
}

var xxx_messageInfo_EfpRange proto.InternalMessageInfo

func (m *EfpRange) GetVlanIdLow() uint32 {
	if m != nil {
		return m.VlanIdLow
	}
	return 0
}

func (m *EfpRange) GetVlanIdHigh() uint32 {
	if m != nil {
		return m.VlanIdHigh
	}
	return 0
}

type EfpTagMatch struct {
	Ethertype            string      `protobuf:"bytes,1,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	VlanRange            []*EfpRange `protobuf:"bytes,2,rep,name=vlan_range,json=vlanRange,proto3" json:"vlan_range,omitempty"`
	Priority             string      `protobuf:"bytes,3,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *EfpTagMatch) Reset()         { *m = EfpTagMatch{} }
func (m *EfpTagMatch) String() string { return proto.CompactTextString(m) }
func (*EfpTagMatch) ProtoMessage()    {}
func (*EfpTagMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{3}
}

func (m *EfpTagMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EfpTagMatch.Unmarshal(m, b)
}
func (m *EfpTagMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EfpTagMatch.Marshal(b, m, deterministic)
}
func (m *EfpTagMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EfpTagMatch.Merge(m, src)
}
func (m *EfpTagMatch) XXX_Size() int {
	return xxx_messageInfo_EfpTagMatch.Size(m)
}
func (m *EfpTagMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_EfpTagMatch.DiscardUnknown(m)
}

var xxx_messageInfo_EfpTagMatch proto.InternalMessageInfo

func (m *EfpTagMatch) GetEthertype() string {
	if m != nil {
		return m.Ethertype
	}
	return ""
}

func (m *EfpTagMatch) GetVlanRange() []*EfpRange {
	if m != nil {
		return m.VlanRange
	}
	return nil
}

func (m *EfpTagMatch) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

type EfpVlanTag struct {
	Ethertype            string   `protobuf:"bytes,1,opt,name=ethertype,proto3" json:"ethertype,omitempty"`
	VlanId               uint32   `protobuf:"varint,2,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EfpVlanTag) Reset()         { *m = EfpVlanTag{} }
func (m *EfpVlanTag) String() string { return proto.CompactTextString(m) }
func (*EfpVlanTag) ProtoMessage()    {}
func (*EfpVlanTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{4}
}

func (m *EfpVlanTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EfpVlanTag.Unmarshal(m, b)
}
func (m *EfpVlanTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EfpVlanTag.Marshal(b, m, deterministic)
}
func (m *EfpVlanTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EfpVlanTag.Merge(m, src)
}
func (m *EfpVlanTag) XXX_Size() int {
	return xxx_messageInfo_EfpVlanTag.Size(m)
}
func (m *EfpVlanTag) XXX_DiscardUnknown() {
	xxx_messageInfo_EfpVlanTag.DiscardUnknown(m)
}

var xxx_messageInfo_EfpVlanTag proto.InternalMessageInfo

func (m *EfpVlanTag) GetEthertype() string {
	if m != nil {
		return m.Ethertype
	}
	return ""
}

func (m *EfpVlanTag) GetVlanId() uint32 {
	if m != nil {
		return m.VlanId
	}
	return 0
}

type EfpTagLocalTrafficStack struct {
	LocalTrafficTag      []*EfpVlanTag `protobuf:"bytes,1,rep,name=local_traffic_tag,json=localTrafficTag,proto3" json:"local_traffic_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EfpTagLocalTrafficStack) Reset()         { *m = EfpTagLocalTrafficStack{} }
func (m *EfpTagLocalTrafficStack) String() string { return proto.CompactTextString(m) }
func (*EfpTagLocalTrafficStack) ProtoMessage()    {}
func (*EfpTagLocalTrafficStack) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{5}
}

func (m *EfpTagLocalTrafficStack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EfpTagLocalTrafficStack.Unmarshal(m, b)
}
func (m *EfpTagLocalTrafficStack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EfpTagLocalTrafficStack.Marshal(b, m, deterministic)
}
func (m *EfpTagLocalTrafficStack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EfpTagLocalTrafficStack.Merge(m, src)
}
func (m *EfpTagLocalTrafficStack) XXX_Size() int {
	return xxx_messageInfo_EfpTagLocalTrafficStack.Size(m)
}
func (m *EfpTagLocalTrafficStack) XXX_DiscardUnknown() {
	xxx_messageInfo_EfpTagLocalTrafficStack.DiscardUnknown(m)
}

var xxx_messageInfo_EfpTagLocalTrafficStack proto.InternalMessageInfo

func (m *EfpTagLocalTrafficStack) GetLocalTrafficTag() []*EfpVlanTag {
	if m != nil {
		return m.LocalTrafficTag
	}
	return nil
}

type EfpDescription struct {
	TagsToMatch          []*EfpTagMatch           `protobuf:"bytes,1,rep,name=tags_to_match,json=tagsToMatch,proto3" json:"tags_to_match,omitempty"`
	PayloadEthertype     string                   `protobuf:"bytes,2,opt,name=payload_ethertype,json=payloadEthertype,proto3" json:"payload_ethertype,omitempty"`
	TagsPopped           uint32                   `protobuf:"varint,3,opt,name=tags_popped,json=tagsPopped,proto3" json:"tags_popped,omitempty"`
	Pushe                []*EfpVlanTag            `protobuf:"bytes,4,rep,name=pushe,proto3" json:"pushe,omitempty"`
	LocalTrafficStack    *EfpTagLocalTrafficStack `protobuf:"bytes,5,opt,name=local_traffic_stack,json=localTrafficStack,proto3" json:"local_traffic_stack,omitempty"`
	IsExactMatch         bool                     `protobuf:"varint,6,opt,name=is_exact_match,json=isExactMatch,proto3" json:"is_exact_match,omitempty"`
	IsNativeVlan         bool                     `protobuf:"varint,7,opt,name=is_native_vlan,json=isNativeVlan,proto3" json:"is_native_vlan,omitempty"`
	IsNativePreserving   bool                     `protobuf:"varint,8,opt,name=is_native_preserving,json=isNativePreserving,proto3" json:"is_native_preserving,omitempty"`
	SourceMacMatch       string                   `protobuf:"bytes,9,opt,name=source_mac_match,json=sourceMacMatch,proto3" json:"source_mac_match,omitempty"`
	DestinationMacMatch  string                   `protobuf:"bytes,10,opt,name=destination_mac_match,json=destinationMacMatch,proto3" json:"destination_mac_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *EfpDescription) Reset()         { *m = EfpDescription{} }
func (m *EfpDescription) String() string { return proto.CompactTextString(m) }
func (*EfpDescription) ProtoMessage()    {}
func (*EfpDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{6}
}

func (m *EfpDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EfpDescription.Unmarshal(m, b)
}
func (m *EfpDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EfpDescription.Marshal(b, m, deterministic)
}
func (m *EfpDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EfpDescription.Merge(m, src)
}
func (m *EfpDescription) XXX_Size() int {
	return xxx_messageInfo_EfpDescription.Size(m)
}
func (m *EfpDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_EfpDescription.DiscardUnknown(m)
}

var xxx_messageInfo_EfpDescription proto.InternalMessageInfo

func (m *EfpDescription) GetTagsToMatch() []*EfpTagMatch {
	if m != nil {
		return m.TagsToMatch
	}
	return nil
}

func (m *EfpDescription) GetPayloadEthertype() string {
	if m != nil {
		return m.PayloadEthertype
	}
	return ""
}

func (m *EfpDescription) GetTagsPopped() uint32 {
	if m != nil {
		return m.TagsPopped
	}
	return 0
}

func (m *EfpDescription) GetPushe() []*EfpVlanTag {
	if m != nil {
		return m.Pushe
	}
	return nil
}

func (m *EfpDescription) GetLocalTrafficStack() *EfpTagLocalTrafficStack {
	if m != nil {
		return m.LocalTrafficStack
	}
	return nil
}

func (m *EfpDescription) GetIsExactMatch() bool {
	if m != nil {
		return m.IsExactMatch
	}
	return false
}

func (m *EfpDescription) GetIsNativeVlan() bool {
	if m != nil {
		return m.IsNativeVlan
	}
	return false
}

func (m *EfpDescription) GetIsNativePreserving() bool {
	if m != nil {
		return m.IsNativePreserving
	}
	return false
}

func (m *EfpDescription) GetSourceMacMatch() string {
	if m != nil {
		return m.SourceMacMatch
	}
	return ""
}

func (m *EfpDescription) GetDestinationMacMatch() string {
	if m != nil {
		return m.DestinationMacMatch
	}
	return ""
}

type VlanTagStackType struct {
	VlanEncapsulation      string              `protobuf:"bytes,1,opt,name=vlan_encapsulation,json=vlanEncapsulation,proto3" json:"vlan_encapsulation,omitempty"`
	Tag                    uint32              `protobuf:"varint,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Stack                  *VlanDoubleTagStack `protobuf:"bytes,3,opt,name=stack,proto3" json:"stack,omitempty"`
	OuterTag               uint32              `protobuf:"varint,4,opt,name=outer_tag,json=outerTag,proto3" json:"outer_tag,omitempty"`
	NativeTag              uint32              `protobuf:"varint,5,opt,name=native_tag,json=nativeTag,proto3" json:"native_tag,omitempty"`
	Dot1AdTag              uint32              `protobuf:"varint,6,opt,name=dot1ad_tag,json=dot1adTag,proto3" json:"dot1ad_tag,omitempty"`
	Dot1AdNativeTag        uint32              `protobuf:"varint,7,opt,name=dot1ad_native_tag,json=dot1adNativeTag,proto3" json:"dot1ad_native_tag,omitempty"`
	ServiceInstanceDetails *EfpDescription     `protobuf:"bytes,8,opt,name=service_instance_details,json=serviceInstanceDetails,proto3" json:"service_instance_details,omitempty"`
	Dot1AdDot1QStack       *VlanDoubleTagStack `protobuf:"bytes,9,opt,name=dot1ad_dot1q_stack,json=dot1adDot1qStack,proto3" json:"dot1ad_dot1q_stack,omitempty"`
	Dot1AdOuterTag         uint32              `protobuf:"varint,10,opt,name=dot1ad_outer_tag,json=dot1adOuterTag,proto3" json:"dot1ad_outer_tag,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *VlanTagStackType) Reset()         { *m = VlanTagStackType{} }
func (m *VlanTagStackType) String() string { return proto.CompactTextString(m) }
func (*VlanTagStackType) ProtoMessage()    {}
func (*VlanTagStackType) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{7}
}

func (m *VlanTagStackType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanTagStackType.Unmarshal(m, b)
}
func (m *VlanTagStackType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanTagStackType.Marshal(b, m, deterministic)
}
func (m *VlanTagStackType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanTagStackType.Merge(m, src)
}
func (m *VlanTagStackType) XXX_Size() int {
	return xxx_messageInfo_VlanTagStackType.Size(m)
}
func (m *VlanTagStackType) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanTagStackType.DiscardUnknown(m)
}

var xxx_messageInfo_VlanTagStackType proto.InternalMessageInfo

func (m *VlanTagStackType) GetVlanEncapsulation() string {
	if m != nil {
		return m.VlanEncapsulation
	}
	return ""
}

func (m *VlanTagStackType) GetTag() uint32 {
	if m != nil {
		return m.Tag
	}
	return 0
}

func (m *VlanTagStackType) GetStack() *VlanDoubleTagStack {
	if m != nil {
		return m.Stack
	}
	return nil
}

func (m *VlanTagStackType) GetOuterTag() uint32 {
	if m != nil {
		return m.OuterTag
	}
	return 0
}

func (m *VlanTagStackType) GetNativeTag() uint32 {
	if m != nil {
		return m.NativeTag
	}
	return 0
}

func (m *VlanTagStackType) GetDot1AdTag() uint32 {
	if m != nil {
		return m.Dot1AdTag
	}
	return 0
}

func (m *VlanTagStackType) GetDot1AdNativeTag() uint32 {
	if m != nil {
		return m.Dot1AdNativeTag
	}
	return 0
}

func (m *VlanTagStackType) GetServiceInstanceDetails() *EfpDescription {
	if m != nil {
		return m.ServiceInstanceDetails
	}
	return nil
}

func (m *VlanTagStackType) GetDot1AdDot1QStack() *VlanDoubleTagStack {
	if m != nil {
		return m.Dot1AdDot1QStack
	}
	return nil
}

func (m *VlanTagStackType) GetDot1AdOuterTag() uint32 {
	if m != nil {
		return m.Dot1AdOuterTag
	}
	return 0
}

type VlanIntf struct {
	InterfaceXr          string            `protobuf:"bytes,50,opt,name=interface_xr,json=interfaceXr,proto3" json:"interface_xr,omitempty"`
	ParentInterface      string            `protobuf:"bytes,51,opt,name=parent_interface,json=parentInterface,proto3" json:"parent_interface,omitempty"`
	EncapsulationDetails *VlanTagStackType `protobuf:"bytes,52,opt,name=encapsulation_details,json=encapsulationDetails,proto3" json:"encapsulation_details,omitempty"`
	Service              string            `protobuf:"bytes,53,opt,name=service,proto3" json:"service,omitempty"`
	State                string            `protobuf:"bytes,54,opt,name=state,proto3" json:"state,omitempty"`
	Mtu                  uint32            `protobuf:"varint,55,opt,name=mtu,proto3" json:"mtu,omitempty"`
	SwitchedMtu          uint32            `protobuf:"varint,56,opt,name=switched_mtu,json=switchedMtu,proto3" json:"switched_mtu,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VlanIntf) Reset()         { *m = VlanIntf{} }
func (m *VlanIntf) String() string { return proto.CompactTextString(m) }
func (*VlanIntf) ProtoMessage()    {}
func (*VlanIntf) Descriptor() ([]byte, []int) {
	return fileDescriptor_a8ef138bb8a02877, []int{8}
}

func (m *VlanIntf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VlanIntf.Unmarshal(m, b)
}
func (m *VlanIntf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VlanIntf.Marshal(b, m, deterministic)
}
func (m *VlanIntf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VlanIntf.Merge(m, src)
}
func (m *VlanIntf) XXX_Size() int {
	return xxx_messageInfo_VlanIntf.Size(m)
}
func (m *VlanIntf) XXX_DiscardUnknown() {
	xxx_messageInfo_VlanIntf.DiscardUnknown(m)
}

var xxx_messageInfo_VlanIntf proto.InternalMessageInfo

func (m *VlanIntf) GetInterfaceXr() string {
	if m != nil {
		return m.InterfaceXr
	}
	return ""
}

func (m *VlanIntf) GetParentInterface() string {
	if m != nil {
		return m.ParentInterface
	}
	return ""
}

func (m *VlanIntf) GetEncapsulationDetails() *VlanTagStackType {
	if m != nil {
		return m.EncapsulationDetails
	}
	return nil
}

func (m *VlanIntf) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *VlanIntf) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *VlanIntf) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *VlanIntf) GetSwitchedMtu() uint32 {
	if m != nil {
		return m.SwitchedMtu
	}
	return 0
}

func init() {
	proto.RegisterType((*VlanIntf_KEYS)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.vlan_intf_KEYS")
	proto.RegisterType((*VlanDoubleTagStack)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.vlan_double_tag_stack")
	proto.RegisterType((*EfpRange)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.efp_range")
	proto.RegisterType((*EfpTagMatch)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.efp_tag_match")
	proto.RegisterType((*EfpVlanTag)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.efp_vlan_tag")
	proto.RegisterType((*EfpTagLocalTrafficStack)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.efp_tag_local_traffic_stack")
	proto.RegisterType((*EfpDescription)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.efp_description")
	proto.RegisterType((*VlanTagStackType)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.vlan_tag_stack_type")
	proto.RegisterType((*VlanIntf)(nil), "cisco_ios_xr_l2_eth_infra_oper.vlan.nodes.node.tag_allocations.tag_allocation.vlan_intf")
}

func init() { proto.RegisterFile("vlan_intf.proto", fileDescriptor_a8ef138bb8a02877) }

var fileDescriptor_a8ef138bb8a02877 = []byte{
	// 926 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdb, 0x6e, 0x23, 0x35,
	0x18, 0xd6, 0xb4, 0x4d, 0xd2, 0xfc, 0x69, 0x9b, 0xc4, 0x6d, 0xd9, 0x11, 0xcb, 0x21, 0x44, 0x5c,
	0x04, 0x10, 0x11, 0x64, 0x39, 0x3d, 0xc0, 0x46, 0xa2, 0x82, 0x2c, 0xab, 0x69, 0x85, 0x16, 0x81,
	0xb0, 0x5c, 0x8f, 0x33, 0x31, 0x4c, 0xc7, 0xb3, 0xb6, 0xd3, 0x03, 0x57, 0x48, 0x5c, 0x70, 0xcf,
	0x0d, 0x88, 0x17, 0xe0, 0x01, 0x78, 0x03, 0x1e, 0x80, 0x6b, 0x1e, 0x07, 0xf9, 0xb7, 0x67, 0x92,
	0xb2, 0x95, 0xb8, 0xd9, 0xec, 0x4d, 0xdb, 0xf9, 0xfe, 0xf3, 0xf7, 0xf9, 0x77, 0x0d, 0xdd, 0xcb,
	0x9c, 0x15, 0x54, 0x16, 0x76, 0x3e, 0x2e, 0xb5, 0xb2, 0x8a, 0xcc, 0xb8, 0x34, 0x5c, 0x51, 0xa9,
	0x0c, 0xbd, 0xd6, 0x34, 0x9f, 0x50, 0x61, 0x17, 0x54, 0x16, 0x73, 0xcd, 0xa8, 0x2a, 0x85, 0x1e,
	0x3b, 0xff, 0x71, 0xa1, 0x52, 0x61, 0xf0, 0xe7, 0xd8, 0xb2, 0x8c, 0xb2, 0x3c, 0x57, 0x9c, 0x59,
	0xa9, 0x0a, 0xf3, 0x9f, 0xef, 0xe1, 0x4f, 0x11, 0x1c, 0xd4, 0x25, 0xe8, 0x67, 0xd3, 0xaf, 0x4e,
	0xc9, 0x3d, 0x68, 0xb9, 0x48, 0x2a, 0xd3, 0x38, 0x1a, 0x44, 0xa3, 0x76, 0xd2, 0x74, 0x9f, 0x27,
	0x29, 0x79, 0x05, 0xda, 0xb2, 0xb0, 0x42, 0xcf, 0x19, 0x17, 0xf1, 0x16, 0x9a, 0x56, 0x00, 0xb9,
	0x0f, 0xed, 0xb9, 0xd4, 0xc6, 0x52, 0xcb, 0xb2, 0x78, 0x7b, 0x10, 0x8d, 0xf6, 0x93, 0x5d, 0x04,
	0xce, 0x58, 0x46, 0x5e, 0x05, 0x30, 0x82, 0xab, 0x22, 0x45, 0xeb, 0x8e, 0x8f, 0xf5, 0xc8, 0x19,
	0xcb, 0x86, 0xa7, 0x70, 0x8c, 0x4d, 0xa4, 0x6a, 0x79, 0x9e, 0x0b, 0xe7, 0x43, 0x8d, 0x65, 0xfc,
	0x7b, 0x97, 0x54, 0x2d, 0xad, 0xd0, 0x18, 0x16, 0xf9, 0xa4, 0x08, 0x3c, 0x9b, 0x74, 0x0b, 0xad,
	0x6b, 0x49, 0x67, 0xd0, 0x16, 0xf3, 0x92, 0x6a, 0x56, 0x64, 0x82, 0xbc, 0x06, 0x1d, 0x3f, 0x66,
	0x4a, 0x73, 0x75, 0x15, 0x52, 0xb5, 0x1d, 0x74, 0x92, 0x7e, 0xae, 0xae, 0xc8, 0x00, 0xf6, 0x2a,
	0xfb, 0x42, 0x66, 0x8b, 0x90, 0x0d, 0xbc, 0xc3, 0xa7, 0x32, 0x5b, 0x0c, 0xff, 0x8a, 0x60, 0xdf,
	0xe5, 0x73, 0xcd, 0x5d, 0x30, 0xcb, 0x17, 0x8e, 0x0f, 0x61, 0x17, 0x42, 0xdb, 0x9b, 0x52, 0x04,
	0xaa, 0x56, 0x00, 0xb9, 0x02, 0x8c, 0xf6, 0xf5, 0xe3, 0xad, 0xc1, 0xf6, 0xa8, 0x33, 0x79, 0x32,
	0x7e, 0xae, 0xea, 0x8d, 0xeb, 0xf9, 0xfc, 0x28, 0x09, 0x8e, 0xfa, 0x32, 0xec, 0x96, 0x5a, 0x2a,
	0x2d, 0xed, 0x0d, 0xea, 0xd0, 0x4e, 0xea, 0xef, 0xe1, 0x14, 0xf6, 0x5c, 0x0c, 0x36, 0x66, 0x59,
	0xf6, 0x3f, 0x23, 0xdc, 0x83, 0x56, 0x20, 0x25, 0xf0, 0xd1, 0xf4, 0x7c, 0x0c, 0xff, 0x88, 0xe0,
	0x7e, 0xc5, 0x85, 0x6b, 0x26, 0xa7, 0x56, 0xb3, 0xf9, 0x5c, 0xf2, 0x20, 0xdb, 0xcf, 0x11, 0xf4,
	0x6f, 0xe3, 0x5e, 0x3f, 0xc7, 0xc1, 0xd7, 0x1b, 0xe0, 0xa0, 0x9a, 0x27, 0xe9, 0x62, 0xd5, 0x33,
	0x5f, 0xd4, 0x1d, 0x82, 0xbf, 0x1b, 0xd0, 0x75, 0x1e, 0xa9, 0x30, 0x5c, 0xcb, 0xd2, 0x45, 0x90,
	0x1f, 0x23, 0xd8, 0xb7, 0x2c, 0x33, 0xd4, 0x2a, 0xaf, 0x64, 0xe8, 0xec, 0x9b, 0x0d, 0x74, 0x56,
	0x9f, 0x96, 0xa4, 0xe3, 0x4a, 0x9e, 0xa9, 0x19, 0x1e, 0x9d, 0x77, 0xa0, 0x5f, 0xb2, 0x9b, 0x5c,
	0xb1, 0x94, 0xae, 0xf8, 0xf7, 0x2b, 0xd5, 0x0b, 0x86, 0x69, 0x2d, 0xc3, 0xeb, 0x80, 0xb1, 0xb4,
	0x54, 0x65, 0x29, 0xd2, 0xb0, 0x5b, 0xe0, 0xa0, 0xc7, 0x88, 0x90, 0xa7, 0xd0, 0x28, 0x97, 0x66,
	0x21, 0xe2, 0x9d, 0xcd, 0x33, 0xec, 0x2b, 0x91, 0xdf, 0x23, 0x38, 0xbc, 0x43, 0xf9, 0xb8, 0x31,
	0x88, 0x46, 0x9d, 0xc9, 0x77, 0x1b, 0x62, 0xf2, 0x8e, 0x8a, 0x49, 0x7f, 0x5d, 0xf2, 0x53, 0x3c,
	0x7e, 0x6f, 0xc2, 0x81, 0x34, 0x54, 0x5c, 0x33, 0x6e, 0x83, 0xc0, 0xcd, 0x41, 0x34, 0xda, 0x4d,
	0xf6, 0xa4, 0x99, 0x3a, 0xd0, 0x6b, 0xe0, 0xbd, 0x0a, 0x66, 0xe5, 0xa5, 0xc0, 0xf9, 0xe2, 0x56,
	0xe5, 0xf5, 0x08, 0xc1, 0x2f, 0x73, 0x56, 0x90, 0xf7, 0xe0, 0x68, 0xe5, 0x55, 0x6a, 0x61, 0x84,
	0xbe, 0x94, 0x45, 0x16, 0xef, 0xa2, 0x2f, 0xa9, 0x7c, 0x1f, 0xd7, 0x16, 0x32, 0x82, 0x9e, 0x51,
	0x4b, 0xcd, 0x05, 0xbd, 0x60, 0x3c, 0xd4, 0x6f, 0xa3, 0xb4, 0x07, 0x1e, 0x9f, 0x31, 0xee, 0x3b,
	0x98, 0xc0, 0x71, 0x2a, 0x8c, 0x95, 0x05, 0x4e, 0xb9, 0xe6, 0x0e, 0xe8, 0x7e, 0xb8, 0x66, 0xac,
	0x62, 0x86, 0x7f, 0x36, 0xe0, 0xb0, 0x12, 0xc3, 0x33, 0x40, 0xf1, 0x90, 0xbc, 0x0b, 0x04, 0x61,
	0x51, 0x70, 0x56, 0x9a, 0x65, 0x8e, 0x51, 0x61, 0xa5, 0xfb, 0xce, 0x32, 0x5d, 0x37, 0x90, 0x1e,
	0x6c, 0xaf, 0x2e, 0x4d, 0xf7, 0x27, 0xf9, 0x01, 0x1a, 0x5e, 0xc2, 0x6d, 0x94, 0x30, 0x7d, 0xce,
	0x12, 0xde, 0x79, 0xbf, 0x27, 0x8d, 0x3b, 0xae, 0xf9, 0x9d, 0x67, 0xaf, 0xf9, 0x40, 0xbf, 0xb3,
	0x36, 0xfc, 0xcd, 0xed, 0x91, 0x60, 0x4e, 0x95, 0x7d, 0x9f, 0xf9, 0xff, 0x02, 0x4d, 0x6f, 0xf6,
	0x88, 0x33, 0xbf, 0x0d, 0xfd, 0x60, 0x5e, 0x4b, 0xd2, 0x42, 0xaf, 0xae, 0x37, 0x3c, 0xaa, 0x53,
	0xfd, 0x16, 0x41, 0x8c, 0x2a, 0x72, 0x41, 0x65, 0x61, 0x2c, 0x2b, 0xb8, 0xa0, 0xa9, 0xb0, 0x4c,
	0xe6, 0x06, 0x05, 0xef, 0x4c, 0xbe, 0xdd, 0xc0, 0xc9, 0x5e, 0xbb, 0x9b, 0x92, 0x97, 0x42, 0xfd,
	0x93, 0x50, 0xfe, 0xa1, 0xaf, 0x4e, 0x7e, 0x89, 0x80, 0x84, 0x39, 0xdc, 0xaf, 0xa7, 0x61, 0xdd,
	0xda, 0x2f, 0x50, 0xab, 0x9e, 0xaf, 0xff, 0xd0, 0x95, 0xf7, 0x7b, 0x36, 0x82, 0x80, 0xd1, 0x95,
	0x7a, 0x80, 0xd4, 0x1e, 0x78, 0xfc, 0x8b, 0xa0, 0xe1, 0xf0, 0x9f, 0x2d, 0x68, 0xd7, 0xcf, 0x0c,
	0xf2, 0x06, 0xec, 0xd5, 0xef, 0x06, 0x7a, 0xad, 0xe3, 0x09, 0x9e, 0xd2, 0x4e, 0x8d, 0x3d, 0xd1,
	0xe4, 0x2d, 0xe8, 0x95, 0x4c, 0x8b, 0xc2, 0xd2, 0xd5, 0x93, 0xe3, 0x01, 0xba, 0x75, 0x3d, 0x7e,
	0x52, 0x3f, 0x3c, 0x7e, 0x8d, 0xe0, 0xf8, 0xd6, 0xa9, 0xaf, 0x25, 0xfb, 0x00, 0xd9, 0x39, 0xdf,
	0x04, 0x3b, 0xb7, 0xb7, 0x2f, 0x39, 0xba, 0xd5, 0x40, 0x25, 0x5a, 0x0c, 0xad, 0x20, 0x67, 0xfc,
	0x21, 0xf6, 0x5e, 0x7d, 0x92, 0x23, 0x5c, 0x36, 0x2b, 0xe2, 0x8f, 0x10, 0xf7, 0x1f, 0x6e, 0x29,
	0x2f, 0xec, 0x32, 0xfe, 0xd8, 0x2f, 0xe5, 0x85, 0x5d, 0x3a, 0xa6, 0xcc, 0x95, 0xb4, 0x7c, 0x21,
	0x52, 0xea, 0x4c, 0x9f, 0xa0, 0xa9, 0x53, 0x61, 0x33, 0xbb, 0x3c, 0x6f, 0xe2, 0xbb, 0xf0, 0xc1,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x90, 0x50, 0x7e, 0x2a, 0x0a, 0x00, 0x00,
}
