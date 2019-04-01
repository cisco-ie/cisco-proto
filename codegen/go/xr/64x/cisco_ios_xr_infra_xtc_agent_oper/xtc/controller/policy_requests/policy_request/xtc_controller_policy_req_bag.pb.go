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
// source: xtc_controller_policy_req_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_controller_policy_requests_policy_request

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

type XtcControllerPolicyReqBag_KEYS struct {
	SourceAddress        string   `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	EndPointType         string   `protobuf:"bytes,2,opt,name=end_point_type,json=endPointType,proto3" json:"end_point_type,omitempty"`
	EndPointAddress      string   `protobuf:"bytes,3,opt,name=end_point_address,json=endPointAddress,proto3" json:"end_point_address,omitempty"`
	Color                uint32   `protobuf:"varint,4,opt,name=color,proto3" json:"color,omitempty"`
	RouteDistinguisher   uint32   `protobuf:"varint,5,opt,name=route_distinguisher,json=routeDistinguisher,proto3" json:"route_distinguisher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcControllerPolicyReqBag_KEYS) Reset()         { *m = XtcControllerPolicyReqBag_KEYS{} }
func (m *XtcControllerPolicyReqBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcControllerPolicyReqBag_KEYS) ProtoMessage()    {}
func (*XtcControllerPolicyReqBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{0}
}

func (m *XtcControllerPolicyReqBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcControllerPolicyReqBag_KEYS.Unmarshal(m, b)
}
func (m *XtcControllerPolicyReqBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcControllerPolicyReqBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcControllerPolicyReqBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcControllerPolicyReqBag_KEYS.Merge(m, src)
}
func (m *XtcControllerPolicyReqBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcControllerPolicyReqBag_KEYS.Size(m)
}
func (m *XtcControllerPolicyReqBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcControllerPolicyReqBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcControllerPolicyReqBag_KEYS proto.InternalMessageInfo

func (m *XtcControllerPolicyReqBag_KEYS) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *XtcControllerPolicyReqBag_KEYS) GetEndPointType() string {
	if m != nil {
		return m.EndPointType
	}
	return ""
}

func (m *XtcControllerPolicyReqBag_KEYS) GetEndPointAddress() string {
	if m != nil {
		return m.EndPointAddress
	}
	return ""
}

func (m *XtcControllerPolicyReqBag_KEYS) GetColor() uint32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *XtcControllerPolicyReqBag_KEYS) GetRouteDistinguisher() uint32 {
	if m != nil {
		return m.RouteDistinguisher
	}
	return 0
}

type XtcIpAddrType struct {
	AfName               string   `protobuf:"bytes,1,opt,name=af_name,json=afName,proto3" json:"af_name,omitempty"`
	Ipv4                 string   `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcIpAddrType) Reset()         { *m = XtcIpAddrType{} }
func (m *XtcIpAddrType) String() string { return proto.CompactTextString(m) }
func (*XtcIpAddrType) ProtoMessage()    {}
func (*XtcIpAddrType) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{1}
}

func (m *XtcIpAddrType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcIpAddrType.Unmarshal(m, b)
}
func (m *XtcIpAddrType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcIpAddrType.Marshal(b, m, deterministic)
}
func (m *XtcIpAddrType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcIpAddrType.Merge(m, src)
}
func (m *XtcIpAddrType) XXX_Size() int {
	return xxx_messageInfo_XtcIpAddrType.Size(m)
}
func (m *XtcIpAddrType) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcIpAddrType.DiscardUnknown(m)
}

var xxx_messageInfo_XtcIpAddrType proto.InternalMessageInfo

func (m *XtcIpAddrType) GetAfName() string {
	if m != nil {
		return m.AfName
	}
	return ""
}

func (m *XtcIpAddrType) GetIpv4() string {
	if m != nil {
		return m.Ipv4
	}
	return ""
}

func (m *XtcIpAddrType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type XtcPrefixSidType struct {
	SidType              string   `protobuf:"bytes,1,opt,name=sid_type,json=sidType,proto3" json:"sid_type,omitempty"`
	Label                uint32   `protobuf:"varint,2,opt,name=label,proto3" json:"label,omitempty"`
	Ipv6                 string   `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcPrefixSidType) Reset()         { *m = XtcPrefixSidType{} }
func (m *XtcPrefixSidType) String() string { return proto.CompactTextString(m) }
func (*XtcPrefixSidType) ProtoMessage()    {}
func (*XtcPrefixSidType) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{2}
}

func (m *XtcPrefixSidType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPrefixSidType.Unmarshal(m, b)
}
func (m *XtcPrefixSidType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPrefixSidType.Marshal(b, m, deterministic)
}
func (m *XtcPrefixSidType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPrefixSidType.Merge(m, src)
}
func (m *XtcPrefixSidType) XXX_Size() int {
	return xxx_messageInfo_XtcPrefixSidType.Size(m)
}
func (m *XtcPrefixSidType) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPrefixSidType.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPrefixSidType proto.InternalMessageInfo

func (m *XtcPrefixSidType) GetSidType() string {
	if m != nil {
		return m.SidType
	}
	return ""
}

func (m *XtcPrefixSidType) GetLabel() uint32 {
	if m != nil {
		return m.Label
	}
	return 0
}

func (m *XtcPrefixSidType) GetIpv6() string {
	if m != nil {
		return m.Ipv6
	}
	return ""
}

type XtcSrHopType struct {
	SidType              string            `protobuf:"bytes,1,opt,name=sid_type,json=sidType,proto3" json:"sid_type,omitempty"`
	Sid                  *XtcPrefixSidType `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`
	LocalAddress         *XtcIpAddrType    `protobuf:"bytes,3,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
	RemoteAddress        *XtcIpAddrType    `protobuf:"bytes,4,opt,name=remote_address,json=remoteAddress,proto3" json:"remote_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *XtcSrHopType) Reset()         { *m = XtcSrHopType{} }
func (m *XtcSrHopType) String() string { return proto.CompactTextString(m) }
func (*XtcSrHopType) ProtoMessage()    {}
func (*XtcSrHopType) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{3}
}

func (m *XtcSrHopType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcSrHopType.Unmarshal(m, b)
}
func (m *XtcSrHopType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcSrHopType.Marshal(b, m, deterministic)
}
func (m *XtcSrHopType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcSrHopType.Merge(m, src)
}
func (m *XtcSrHopType) XXX_Size() int {
	return xxx_messageInfo_XtcSrHopType.Size(m)
}
func (m *XtcSrHopType) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcSrHopType.DiscardUnknown(m)
}

var xxx_messageInfo_XtcSrHopType proto.InternalMessageInfo

func (m *XtcSrHopType) GetSidType() string {
	if m != nil {
		return m.SidType
	}
	return ""
}

func (m *XtcSrHopType) GetSid() *XtcPrefixSidType {
	if m != nil {
		return m.Sid
	}
	return nil
}

func (m *XtcSrHopType) GetLocalAddress() *XtcIpAddrType {
	if m != nil {
		return m.LocalAddress
	}
	return nil
}

func (m *XtcSrHopType) GetRemoteAddress() *XtcIpAddrType {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

type XtcAffinityColor struct {
	ColorXr              string   `protobuf:"bytes,1,opt,name=color_xr,json=colorXr,proto3" json:"color_xr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcAffinityColor) Reset()         { *m = XtcAffinityColor{} }
func (m *XtcAffinityColor) String() string { return proto.CompactTextString(m) }
func (*XtcAffinityColor) ProtoMessage()    {}
func (*XtcAffinityColor) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{4}
}

func (m *XtcAffinityColor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcAffinityColor.Unmarshal(m, b)
}
func (m *XtcAffinityColor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcAffinityColor.Marshal(b, m, deterministic)
}
func (m *XtcAffinityColor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcAffinityColor.Merge(m, src)
}
func (m *XtcAffinityColor) XXX_Size() int {
	return xxx_messageInfo_XtcAffinityColor.Size(m)
}
func (m *XtcAffinityColor) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcAffinityColor.DiscardUnknown(m)
}

var xxx_messageInfo_XtcAffinityColor proto.InternalMessageInfo

func (m *XtcAffinityColor) GetColorXr() string {
	if m != nil {
		return m.ColorXr
	}
	return ""
}

type XtcAffinityConstraint struct {
	Type                 uint32              `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                uint32              `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Color                []*XtcAffinityColor `protobuf:"bytes,3,rep,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *XtcAffinityConstraint) Reset()         { *m = XtcAffinityConstraint{} }
func (m *XtcAffinityConstraint) String() string { return proto.CompactTextString(m) }
func (*XtcAffinityConstraint) ProtoMessage()    {}
func (*XtcAffinityConstraint) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{5}
}

func (m *XtcAffinityConstraint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcAffinityConstraint.Unmarshal(m, b)
}
func (m *XtcAffinityConstraint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcAffinityConstraint.Marshal(b, m, deterministic)
}
func (m *XtcAffinityConstraint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcAffinityConstraint.Merge(m, src)
}
func (m *XtcAffinityConstraint) XXX_Size() int {
	return xxx_messageInfo_XtcAffinityConstraint.Size(m)
}
func (m *XtcAffinityConstraint) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcAffinityConstraint.DiscardUnknown(m)
}

var xxx_messageInfo_XtcAffinityConstraint proto.InternalMessageInfo

func (m *XtcAffinityConstraint) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *XtcAffinityConstraint) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *XtcAffinityConstraint) GetColor() []*XtcAffinityColor {
	if m != nil {
		return m.Color
	}
	return nil
}

type XtcPathMetricsBag struct {
	MarginRelative        uint32   `protobuf:"varint,1,opt,name=margin_relative,json=marginRelative,proto3" json:"margin_relative,omitempty"`
	MarginAbsolute        uint32   `protobuf:"varint,2,opt,name=margin_absolute,json=marginAbsolute,proto3" json:"margin_absolute,omitempty"`
	MaximumSegments       uint32   `protobuf:"varint,3,opt,name=maximum_segments,json=maximumSegments,proto3" json:"maximum_segments,omitempty"`
	AccumulativeTeMetric  uint32   `protobuf:"varint,4,opt,name=accumulative_te_metric,json=accumulativeTeMetric,proto3" json:"accumulative_te_metric,omitempty"`
	AccumulativeIgpMetric uint32   `protobuf:"varint,5,opt,name=accumulative_igp_metric,json=accumulativeIgpMetric,proto3" json:"accumulative_igp_metric,omitempty"`
	AccumulativeDelay     uint32   `protobuf:"varint,6,opt,name=accumulative_delay,json=accumulativeDelay,proto3" json:"accumulative_delay,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *XtcPathMetricsBag) Reset()         { *m = XtcPathMetricsBag{} }
func (m *XtcPathMetricsBag) String() string { return proto.CompactTextString(m) }
func (*XtcPathMetricsBag) ProtoMessage()    {}
func (*XtcPathMetricsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{6}
}

func (m *XtcPathMetricsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPathMetricsBag.Unmarshal(m, b)
}
func (m *XtcPathMetricsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPathMetricsBag.Marshal(b, m, deterministic)
}
func (m *XtcPathMetricsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPathMetricsBag.Merge(m, src)
}
func (m *XtcPathMetricsBag) XXX_Size() int {
	return xxx_messageInfo_XtcPathMetricsBag.Size(m)
}
func (m *XtcPathMetricsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPathMetricsBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPathMetricsBag proto.InternalMessageInfo

func (m *XtcPathMetricsBag) GetMarginRelative() uint32 {
	if m != nil {
		return m.MarginRelative
	}
	return 0
}

func (m *XtcPathMetricsBag) GetMarginAbsolute() uint32 {
	if m != nil {
		return m.MarginAbsolute
	}
	return 0
}

func (m *XtcPathMetricsBag) GetMaximumSegments() uint32 {
	if m != nil {
		return m.MaximumSegments
	}
	return 0
}

func (m *XtcPathMetricsBag) GetAccumulativeTeMetric() uint32 {
	if m != nil {
		return m.AccumulativeTeMetric
	}
	return 0
}

func (m *XtcPathMetricsBag) GetAccumulativeIgpMetric() uint32 {
	if m != nil {
		return m.AccumulativeIgpMetric
	}
	return 0
}

func (m *XtcPathMetricsBag) GetAccumulativeDelay() uint32 {
	if m != nil {
		return m.AccumulativeDelay
	}
	return 0
}

type XtcPathConstraintsBag struct {
	AffinityConstraint   []*XtcAffinityConstraint `protobuf:"bytes,1,rep,name=affinity_constraint,json=affinityConstraint,proto3" json:"affinity_constraint,omitempty"`
	PathMetrics          *XtcPathMetricsBag       `protobuf:"bytes,2,opt,name=path_metrics,json=pathMetrics,proto3" json:"path_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *XtcPathConstraintsBag) Reset()         { *m = XtcPathConstraintsBag{} }
func (m *XtcPathConstraintsBag) String() string { return proto.CompactTextString(m) }
func (*XtcPathConstraintsBag) ProtoMessage()    {}
func (*XtcPathConstraintsBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{7}
}

func (m *XtcPathConstraintsBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPathConstraintsBag.Unmarshal(m, b)
}
func (m *XtcPathConstraintsBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPathConstraintsBag.Marshal(b, m, deterministic)
}
func (m *XtcPathConstraintsBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPathConstraintsBag.Merge(m, src)
}
func (m *XtcPathConstraintsBag) XXX_Size() int {
	return xxx_messageInfo_XtcPathConstraintsBag.Size(m)
}
func (m *XtcPathConstraintsBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPathConstraintsBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPathConstraintsBag proto.InternalMessageInfo

func (m *XtcPathConstraintsBag) GetAffinityConstraint() []*XtcAffinityConstraint {
	if m != nil {
		return m.AffinityConstraint
	}
	return nil
}

func (m *XtcPathConstraintsBag) GetPathMetrics() *XtcPathMetricsBag {
	if m != nil {
		return m.PathMetrics
	}
	return nil
}

type XtcPolicyPathBag struct {
	Index                uint32                 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Type                 string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Active               bool                   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	Weight               uint32                 `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	MetricType           uint32                 `protobuf:"varint,6,opt,name=metric_type,json=metricType,proto3" json:"metric_type,omitempty"`
	MetricValue          uint32                 `protobuf:"varint,7,opt,name=metric_value,json=metricValue,proto3" json:"metric_value,omitempty"`
	Hops                 []*XtcSrHopType        `protobuf:"bytes,8,rep,name=hops,proto3" json:"hops,omitempty"`
	IsValid              bool                   `protobuf:"varint,9,opt,name=is_valid,json=isValid,proto3" json:"is_valid,omitempty"`
	PceBasedPath         bool                   `protobuf:"varint,10,opt,name=pce_based_path,json=pceBasedPath,proto3" json:"pce_based_path,omitempty"`
	PceAddress           string                 `protobuf:"bytes,11,opt,name=pce_address,json=pceAddress,proto3" json:"pce_address,omitempty"`
	Error                string                 `protobuf:"bytes,12,opt,name=error,proto3" json:"error,omitempty"`
	SrPathConstraints    *XtcPathConstraintsBag `protobuf:"bytes,13,opt,name=sr_path_constraints,json=srPathConstraints,proto3" json:"sr_path_constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *XtcPolicyPathBag) Reset()         { *m = XtcPolicyPathBag{} }
func (m *XtcPolicyPathBag) String() string { return proto.CompactTextString(m) }
func (*XtcPolicyPathBag) ProtoMessage()    {}
func (*XtcPolicyPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{8}
}

func (m *XtcPolicyPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcPolicyPathBag.Unmarshal(m, b)
}
func (m *XtcPolicyPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcPolicyPathBag.Marshal(b, m, deterministic)
}
func (m *XtcPolicyPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcPolicyPathBag.Merge(m, src)
}
func (m *XtcPolicyPathBag) XXX_Size() int {
	return xxx_messageInfo_XtcPolicyPathBag.Size(m)
}
func (m *XtcPolicyPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcPolicyPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcPolicyPathBag proto.InternalMessageInfo

func (m *XtcPolicyPathBag) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *XtcPolicyPathBag) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *XtcPolicyPathBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *XtcPolicyPathBag) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *XtcPolicyPathBag) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *XtcPolicyPathBag) GetMetricType() uint32 {
	if m != nil {
		return m.MetricType
	}
	return 0
}

func (m *XtcPolicyPathBag) GetMetricValue() uint32 {
	if m != nil {
		return m.MetricValue
	}
	return 0
}

func (m *XtcPolicyPathBag) GetHops() []*XtcSrHopType {
	if m != nil {
		return m.Hops
	}
	return nil
}

func (m *XtcPolicyPathBag) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *XtcPolicyPathBag) GetPceBasedPath() bool {
	if m != nil {
		return m.PceBasedPath
	}
	return false
}

func (m *XtcPolicyPathBag) GetPceAddress() string {
	if m != nil {
		return m.PceAddress
	}
	return ""
}

func (m *XtcPolicyPathBag) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *XtcPolicyPathBag) GetSrPathConstraints() *XtcPathConstraintsBag {
	if m != nil {
		return m.SrPathConstraints
	}
	return nil
}

type XtcControllerPolicyReqBag struct {
	SourceAddressXr      string              `protobuf:"bytes,50,opt,name=source_address_xr,json=sourceAddressXr,proto3" json:"source_address_xr,omitempty"`
	EndPoint             *XtcIpAddrType      `protobuf:"bytes,51,opt,name=end_point,json=endPoint,proto3" json:"end_point,omitempty"`
	BindingSid           uint32              `protobuf:"varint,52,opt,name=binding_sid,json=bindingSid,proto3" json:"binding_sid,omitempty"`
	Preference           uint32              `protobuf:"varint,53,opt,name=preference,proto3" json:"preference,omitempty"`
	ColorXr              uint32              `protobuf:"varint,54,opt,name=color_xr,json=colorXr,proto3" json:"color_xr,omitempty"`
	RouteDistinguisherXr uint32              `protobuf:"varint,55,opt,name=route_distinguisher_xr,json=routeDistinguisherXr,proto3" json:"route_distinguisher_xr,omitempty"`
	Paths                []*XtcPolicyPathBag `protobuf:"bytes,56,rep,name=paths,proto3" json:"paths,omitempty"`
	CreationTime         uint32              `protobuf:"varint,57,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	LastUpdatedTime      uint32              `protobuf:"varint,58,opt,name=last_updated_time,json=lastUpdatedTime,proto3" json:"last_updated_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *XtcControllerPolicyReqBag) Reset()         { *m = XtcControllerPolicyReqBag{} }
func (m *XtcControllerPolicyReqBag) String() string { return proto.CompactTextString(m) }
func (*XtcControllerPolicyReqBag) ProtoMessage()    {}
func (*XtcControllerPolicyReqBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fb96e9d4a4b50d4, []int{9}
}

func (m *XtcControllerPolicyReqBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcControllerPolicyReqBag.Unmarshal(m, b)
}
func (m *XtcControllerPolicyReqBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcControllerPolicyReqBag.Marshal(b, m, deterministic)
}
func (m *XtcControllerPolicyReqBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcControllerPolicyReqBag.Merge(m, src)
}
func (m *XtcControllerPolicyReqBag) XXX_Size() int {
	return xxx_messageInfo_XtcControllerPolicyReqBag.Size(m)
}
func (m *XtcControllerPolicyReqBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcControllerPolicyReqBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcControllerPolicyReqBag proto.InternalMessageInfo

func (m *XtcControllerPolicyReqBag) GetSourceAddressXr() string {
	if m != nil {
		return m.SourceAddressXr
	}
	return ""
}

func (m *XtcControllerPolicyReqBag) GetEndPoint() *XtcIpAddrType {
	if m != nil {
		return m.EndPoint
	}
	return nil
}

func (m *XtcControllerPolicyReqBag) GetBindingSid() uint32 {
	if m != nil {
		return m.BindingSid
	}
	return 0
}

func (m *XtcControllerPolicyReqBag) GetPreference() uint32 {
	if m != nil {
		return m.Preference
	}
	return 0
}

func (m *XtcControllerPolicyReqBag) GetColorXr() uint32 {
	if m != nil {
		return m.ColorXr
	}
	return 0
}

func (m *XtcControllerPolicyReqBag) GetRouteDistinguisherXr() uint32 {
	if m != nil {
		return m.RouteDistinguisherXr
	}
	return 0
}

func (m *XtcControllerPolicyReqBag) GetPaths() []*XtcPolicyPathBag {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *XtcControllerPolicyReqBag) GetCreationTime() uint32 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *XtcControllerPolicyReqBag) GetLastUpdatedTime() uint32 {
	if m != nil {
		return m.LastUpdatedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*XtcControllerPolicyReqBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_controller_policy_req_bag_KEYS")
	proto.RegisterType((*XtcIpAddrType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_ip_addr_type")
	proto.RegisterType((*XtcPrefixSidType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_prefix_sid_type")
	proto.RegisterType((*XtcSrHopType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_sr_hop_type")
	proto.RegisterType((*XtcAffinityColor)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_affinity_color")
	proto.RegisterType((*XtcAffinityConstraint)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_affinity_constraint")
	proto.RegisterType((*XtcPathMetricsBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_path_metrics_bag")
	proto.RegisterType((*XtcPathConstraintsBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_path_constraints_bag")
	proto.RegisterType((*XtcPolicyPathBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_policy_path_bag")
	proto.RegisterType((*XtcControllerPolicyReqBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.controller.policy_requests.policy_request.xtc_controller_policy_req_bag")
}

func init() {
	proto.RegisterFile("xtc_controller_policy_req_bag.proto", fileDescriptor_7fb96e9d4a4b50d4)
}

var fileDescriptor_7fb96e9d4a4b50d4 = []byte{
	// 1024 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x72, 0x23, 0x35,
	0x10, 0x2e, 0x27, 0x76, 0xe2, 0xb4, 0xed, 0x64, 0xa3, 0x98, 0xcd, 0xec, 0x01, 0x08, 0x5e, 0x28,
	0x02, 0x55, 0x78, 0xab, 0xb2, 0x21, 0xfc, 0xdc, 0x16, 0x96, 0x03, 0x45, 0x01, 0x5b, 0x93, 0xb0,
	0x15, 0xb8, 0x08, 0x79, 0xa6, 0x6d, 0xab, 0x6a, 0x66, 0x34, 0x48, 0x9a, 0x60, 0x73, 0xe0, 0x48,
	0xed, 0x2b, 0xf0, 0x0a, 0x9c, 0x38, 0x72, 0xa2, 0x8a, 0x47, 0xe0, 0x1d, 0x78, 0x10, 0x4a, 0x2d,
	0x4d, 0x6c, 0x67, 0x53, 0xcb, 0x05, 0xef, 0x4d, 0xfd, 0x75, 0xab, 0xd5, 0xad, 0xfe, 0xd4, 0x6a,
	0xb8, 0x3f, 0xb3, 0x09, 0x4f, 0x54, 0x61, 0xb5, 0xca, 0x32, 0xd4, 0xbc, 0x54, 0x99, 0x4c, 0xe6,
	0x5c, 0xe3, 0x0f, 0x7c, 0x24, 0x26, 0xc3, 0x52, 0x2b, 0xab, 0xd8, 0xd7, 0x89, 0x34, 0x89, 0xe2,
	0x52, 0x19, 0x3e, 0xd3, 0x5c, 0x16, 0x63, 0x2d, 0xb8, 0xdb, 0x27, 0x26, 0x58, 0x58, 0xae, 0x4a,
	0xd4, 0xc3, 0x99, 0x4d, 0x86, 0x0b, 0x37, 0xc3, 0x85, 0x9b, 0x0a, 0x8d, 0x35, 0x37, 0xe4, 0xc1,
	0x3f, 0x0d, 0x18, 0xbc, 0xf0, 0x60, 0xfe, 0xc5, 0x67, 0xdf, 0x9e, 0xb3, 0xb7, 0x60, 0xd7, 0xa8,
	0x4a, 0x27, 0xc8, 0x45, 0x9a, 0x6a, 0x34, 0x26, 0x6a, 0x1c, 0x35, 0x8e, 0x77, 0xe2, 0x9e, 0x47,
	0x1f, 0x79, 0x90, 0xbd, 0x09, 0xbb, 0x58, 0xa4, 0xbc, 0x54, 0xb2, 0xb0, 0xdc, 0xce, 0x4b, 0x8c,
	0x36, 0xc8, 0xac, 0x8b, 0x45, 0xfa, 0xc4, 0x81, 0x17, 0xf3, 0x12, 0xd9, 0xbb, 0xb0, 0xbf, 0xb0,
	0xaa, 0xfd, 0x6d, 0x92, 0xe1, 0x5e, 0x6d, 0x58, 0x7b, 0xec, 0x43, 0x2b, 0x51, 0x99, 0xd2, 0x51,
	0xf3, 0xa8, 0x71, 0xdc, 0x8b, 0xbd, 0xc0, 0x1e, 0xc0, 0x81, 0x56, 0x95, 0x45, 0x9e, 0x4a, 0x63,
	0x65, 0x31, 0xa9, 0xa4, 0x99, 0xa2, 0x8e, 0x5a, 0x64, 0xc3, 0x48, 0xf5, 0x78, 0x59, 0x33, 0x38,
	0x87, 0x3b, 0x2e, 0x4b, 0x59, 0xd2, 0x79, 0x14, 0x1a, 0x3b, 0x84, 0x6d, 0x31, 0xe6, 0x85, 0xc8,
	0x31, 0x24, 0xb3, 0x25, 0xc6, 0x5f, 0x89, 0x1c, 0x19, 0x83, 0xa6, 0x2c, 0xaf, 0x4e, 0x43, 0xec,
	0xb4, 0x0e, 0xd8, 0x59, 0x08, 0x93, 0xd6, 0x83, 0xef, 0xe0, 0xc0, 0x39, 0x2d, 0x35, 0x8e, 0xe5,
	0x8c, 0x1b, 0x99, 0x7a, 0xbf, 0xf7, 0xa0, 0x5d, 0xaf, 0x83, 0xe3, 0x6d, 0x23, 0x53, 0xca, 0xbc,
	0x0f, 0xad, 0x4c, 0x8c, 0x30, 0x23, 0xd7, 0xbd, 0xd8, 0x0b, 0xb7, 0xfa, 0xfe, 0x63, 0x13, 0xf6,
	0x9c, 0x73, 0xa3, 0xf9, 0x54, 0x95, 0xff, 0xe9, 0xf8, 0x0a, 0x36, 0x8d, 0x4c, 0xc9, 0x6d, 0xe7,
	0x24, 0x1d, 0xfe, 0xcf, 0x2c, 0x19, 0xde, 0x92, 0x66, 0xec, 0x0e, 0x64, 0xbf, 0x34, 0xa0, 0x97,
	0xa9, 0x44, 0x64, 0x2b, 0x75, 0xec, 0x9c, 0x88, 0xb5, 0x84, 0xb0, 0x5c, 0xbe, 0xb8, 0x4b, 0xe7,
	0xd6, 0x3c, 0x79, 0xd6, 0x80, 0x5d, 0x8d, 0xb9, 0xb2, 0x0b, 0x86, 0x36, 0x5f, 0x56, 0x24, 0x3d,
	0x7f, 0x70, 0x08, 0x65, 0xf0, 0x00, 0x18, 0xf9, 0x1f, 0x8f, 0x65, 0x21, 0xed, 0x9c, 0x7b, 0xca,
	0xde, 0x83, 0x36, 0x2d, 0xf8, 0x4c, 0xd7, 0xc5, 0x23, 0xf9, 0x52, 0x0f, 0xfe, 0x6c, 0xc0, 0xe1,
	0x8d, 0x1d, 0x85, 0xb1, 0x5a, 0xc8, 0xc2, 0x3a, 0x6e, 0x5c, 0xd7, 0xbb, 0x17, 0xd3, 0xda, 0xb1,
	0xe8, 0x4a, 0x64, 0x15, 0xd6, 0x2c, 0x22, 0x81, 0xcd, 0xeb, 0x97, 0xb2, 0x79, 0xb4, 0x79, 0xdc,
	0x39, 0x49, 0xd6, 0x92, 0xf7, 0x6a, 0x52, 0xe1, 0x39, 0x0e, 0x7e, 0xdf, 0x80, 0x3e, 0x51, 0x44,
	0xd8, 0x29, 0xcf, 0xd1, 0x6a, 0x99, 0x18, 0xd7, 0x3b, 0xd8, 0xdb, 0xb0, 0x97, 0x0b, 0x3d, 0x91,
	0x05, 0xd7, 0x98, 0x09, 0x2b, 0xaf, 0xea, 0x44, 0x76, 0x3d, 0x1c, 0x07, 0x74, 0xc9, 0x50, 0x8c,
	0x8c, 0xca, 0x2a, 0x5b, 0x27, 0x17, 0x0c, 0x1f, 0x05, 0x94, 0xbd, 0x03, 0x77, 0x72, 0x31, 0x93,
	0x79, 0x95, 0x73, 0x83, 0x93, 0x1c, 0x0b, 0xeb, 0x29, 0xd7, 0x8b, 0xf7, 0x02, 0x7e, 0x1e, 0x60,
	0x76, 0x0a, 0x77, 0x45, 0x92, 0x54, 0x79, 0xe5, 0xcf, 0xe0, 0x16, 0x43, 0x6c, 0xa1, 0x97, 0xf4,
	0x97, 0xb5, 0x17, 0xf8, 0x25, 0xe9, 0xd8, 0x19, 0x1c, 0xae, 0xec, 0x92, 0x93, 0xb2, 0xde, 0xe6,
	0xdb, 0xcb, 0x2b, 0xcb, 0xea, 0xcf, 0x27, 0x65, 0xd8, 0xf7, 0x1e, 0xb0, 0x95, 0x7d, 0x29, 0x66,
	0x62, 0x1e, 0x6d, 0xd1, 0x96, 0xfd, 0x65, 0xcd, 0x63, 0xa7, 0x18, 0xfc, 0xb5, 0x01, 0xd1, 0xf5,
	0x95, 0x2d, 0xea, 0xed, 0xaf, 0xed, 0xd7, 0x06, 0x1c, 0xdc, 0x42, 0x86, 0xa8, 0x41, 0x95, 0x9d,
	0xae, 0xbb, 0xb2, 0xf5, 0x79, 0x31, 0xab, 0xc1, 0x4f, 0x17, 0x84, 0x7c, 0xd6, 0x80, 0xee, 0x72,
	0x9d, 0x43, 0xcf, 0xc1, 0xf5, 0xf4, 0x9c, 0x1b, 0x84, 0x8a, 0x3b, 0x0e, 0xf1, 0x37, 0x6e, 0x06,
	0x7f, 0x37, 0x43, 0x03, 0xf6, 0x1b, 0xc9, 0xd8, 0x5d, 0x5f, 0x1f, 0x5a, 0xb2, 0x48, 0x71, 0x16,
	0xb8, 0xe6, 0x85, 0xeb, 0x97, 0x14, 0xba, 0x3a, 0xbd, 0x24, 0x06, 0x4d, 0xea, 0xff, 0xa1, 0xf3,
	0xba, 0x35, 0xbb, 0x0b, 0x5b, 0x22, 0x21, 0xaa, 0x3a, 0x9a, 0xb4, 0xe3, 0x20, 0x39, 0xfc, 0x47,
	0x94, 0x93, 0xa9, 0x0d, 0x3c, 0x08, 0x12, 0x7b, 0x1d, 0x3a, 0x3e, 0x42, 0xdf, 0x98, 0x7d, 0xc5,
	0xc1, 0x43, 0xd4, 0x9b, 0xdf, 0x80, 0x6e, 0x30, 0xf0, 0xaf, 0x76, 0x9b, 0x2c, 0xc2, 0xa6, 0xa7,
	0xf4, 0x76, 0x2d, 0x34, 0xa7, 0xaa, 0x34, 0x51, 0x9b, 0x0a, 0xfc, 0xfd, 0x5a, 0xee, 0x72, 0xe9,
	0x27, 0x89, 0xe9, 0x34, 0xd7, 0x92, 0xa4, 0x71, 0x41, 0xc9, 0x34, 0xda, 0xa1, 0x5c, 0xb7, 0xa5,
	0x79, 0xea, 0x44, 0xf7, 0x91, 0x97, 0x09, 0xf2, 0x91, 0x30, 0x98, 0xd2, 0xc5, 0x46, 0x40, 0x06,
	0xdd, 0x32, 0xc1, 0x4f, 0x1c, 0xf8, 0x44, 0xd8, 0xa9, 0x4b, 0xbd, 0x5c, 0x1a, 0x09, 0x3a, 0x74,
	0x8b, 0x50, 0x2e, 0xe6, 0x81, 0x3e, 0xb4, 0x50, 0x6b, 0xa5, 0xa3, 0x2e, 0xa9, 0xbc, 0x40, 0xf4,
	0x36, 0xfa, 0x39, 0xea, 0x47, 0x3d, 0x62, 0x92, 0x5c, 0x1f, 0x93, 0x6e, 0xbc, 0xb3, 0x78, 0xdf,
	0x68, 0x97, 0xc7, 0x82, 0xdd, 0x66, 0xf0, 0x5b, 0x13, 0x5e, 0x7d, 0xe1, 0x3c, 0xe4, 0xa6, 0x97,
	0xd5, 0x51, 0xc8, 0x75, 0xf4, 0x13, 0x3f, 0xbd, 0xac, 0x4c, 0x43, 0x97, 0x9a, 0xfd, 0x0c, 0x3b,
	0xd7, 0x93, 0x4e, 0xf4, 0xf0, 0x65, 0xfd, 0x47, 0xed, 0x7a, 0x88, 0x72, 0x05, 0x1a, 0xc9, 0x22,
	0x95, 0xc5, 0xc4, 0xfd, 0xdb, 0xd1, 0xa9, 0xe7, 0x66, 0x80, 0xce, 0x65, 0xca, 0x5e, 0x03, 0x70,
	0xff, 0x3a, 0x6a, 0x2c, 0x12, 0x8c, 0xde, 0xf7, 0xfa, 0x05, 0xb2, 0xf2, 0x6b, 0x9d, 0x91, 0xb6,
	0xfe, 0xb5, 0x5c, 0x7b, 0xbd, 0x65, 0x06, 0x73, 0x86, 0x1f, 0xf8, 0xf6, 0xfa, 0xfc, 0x18, 0x76,
	0xa9, 0xd9, 0x4f, 0xd0, 0x72, 0xa5, 0x30, 0xd1, 0x87, 0x44, 0xf5, 0x35, 0x8d, 0x2a, 0xab, 0x0d,
	0x21, 0xf6, 0x47, 0xb2, 0xfb, 0xd0, 0x4b, 0x34, 0x0a, 0x2b, 0x55, 0xc1, 0xad, 0xcc, 0x31, 0xfa,
	0x88, 0x02, 0xed, 0xd6, 0xe0, 0x85, 0xcc, 0x69, 0x38, 0xcd, 0x84, 0xb1, 0xbc, 0x2a, 0x53, 0x61,
	0x31, 0xf5, 0x86, 0x1f, 0xfb, 0x1f, 0xc6, 0x29, 0xbe, 0xf1, 0xb8, 0xb3, 0x1d, 0x6d, 0xd1, 0x50,
	0xfe, 0xf0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xeb, 0xee, 0x01, 0xbb, 0x0b, 0x00, 0x00,
}
