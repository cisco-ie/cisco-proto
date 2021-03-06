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
// source: xtc_fwding_policy_bag.proto

package cisco_ios_xr_infra_xtc_agent_oper_xtc_policy_forwardings_policy_forwarding

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

type XtcFwdingPolicyBag_KEYS struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcFwdingPolicyBag_KEYS) Reset()         { *m = XtcFwdingPolicyBag_KEYS{} }
func (m *XtcFwdingPolicyBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*XtcFwdingPolicyBag_KEYS) ProtoMessage()    {}
func (*XtcFwdingPolicyBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0bea5c4e410cb41, []int{0}
}

func (m *XtcFwdingPolicyBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcFwdingPolicyBag_KEYS.Unmarshal(m, b)
}
func (m *XtcFwdingPolicyBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcFwdingPolicyBag_KEYS.Marshal(b, m, deterministic)
}
func (m *XtcFwdingPolicyBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcFwdingPolicyBag_KEYS.Merge(m, src)
}
func (m *XtcFwdingPolicyBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_XtcFwdingPolicyBag_KEYS.Size(m)
}
func (m *XtcFwdingPolicyBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcFwdingPolicyBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_XtcFwdingPolicyBag_KEYS proto.InternalMessageInfo

func (m *XtcFwdingPolicyBag_KEYS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
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
	return fileDescriptor_c0bea5c4e410cb41, []int{1}
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

type XtcFwdingPolicyPathBag struct {
	OutgoingInterface    string   `protobuf:"bytes,1,opt,name=outgoing_interface,json=outgoingInterface,proto3" json:"outgoing_interface,omitempty"`
	NextHopIpv4          string   `protobuf:"bytes,2,opt,name=next_hop_ipv4,json=nextHopIpv4,proto3" json:"next_hop_ipv4,omitempty"`
	NextHopTableId       uint32   `protobuf:"varint,3,opt,name=next_hop_table_id,json=nextHopTableId,proto3" json:"next_hop_table_id,omitempty"`
	IsProtected          bool     `protobuf:"varint,4,opt,name=is_protected,json=isProtected,proto3" json:"is_protected,omitempty"`
	IsPureBkup           bool     `protobuf:"varint,5,opt,name=is_pure_bkup,json=isPureBkup,proto3" json:"is_pure_bkup,omitempty"`
	LoadMetric           uint32   `protobuf:"varint,6,opt,name=load_metric,json=loadMetric,proto3" json:"load_metric,omitempty"`
	PathId               uint32   `protobuf:"varint,7,opt,name=path_id,json=pathId,proto3" json:"path_id,omitempty"`
	BkupPathId           uint32   `protobuf:"varint,8,opt,name=bkup_path_id,json=bkupPathId,proto3" json:"bkup_path_id,omitempty"`
	LabelStack           []uint32 `protobuf:"varint,9,rep,packed,name=label_stack,json=labelStack,proto3" json:"label_stack,omitempty"`
	SegmentListName      string   `protobuf:"bytes,10,opt,name=segment_list_name,json=segmentListName,proto3" json:"segment_list_name,omitempty"`
	AreStatsValid        bool     `protobuf:"varint,11,opt,name=are_stats_valid,json=areStatsValid,proto3" json:"are_stats_valid,omitempty"`
	ForwardingStatsPkts  uint64   `protobuf:"varint,12,opt,name=forwarding_stats_pkts,json=forwardingStatsPkts,proto3" json:"forwarding_stats_pkts,omitempty"`
	ForwardingStatsBytes uint64   `protobuf:"varint,13,opt,name=forwarding_stats_bytes,json=forwardingStatsBytes,proto3" json:"forwarding_stats_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XtcFwdingPolicyPathBag) Reset()         { *m = XtcFwdingPolicyPathBag{} }
func (m *XtcFwdingPolicyPathBag) String() string { return proto.CompactTextString(m) }
func (*XtcFwdingPolicyPathBag) ProtoMessage()    {}
func (*XtcFwdingPolicyPathBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0bea5c4e410cb41, []int{2}
}

func (m *XtcFwdingPolicyPathBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcFwdingPolicyPathBag.Unmarshal(m, b)
}
func (m *XtcFwdingPolicyPathBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcFwdingPolicyPathBag.Marshal(b, m, deterministic)
}
func (m *XtcFwdingPolicyPathBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcFwdingPolicyPathBag.Merge(m, src)
}
func (m *XtcFwdingPolicyPathBag) XXX_Size() int {
	return xxx_messageInfo_XtcFwdingPolicyPathBag.Size(m)
}
func (m *XtcFwdingPolicyPathBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcFwdingPolicyPathBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcFwdingPolicyPathBag proto.InternalMessageInfo

func (m *XtcFwdingPolicyPathBag) GetOutgoingInterface() string {
	if m != nil {
		return m.OutgoingInterface
	}
	return ""
}

func (m *XtcFwdingPolicyPathBag) GetNextHopIpv4() string {
	if m != nil {
		return m.NextHopIpv4
	}
	return ""
}

func (m *XtcFwdingPolicyPathBag) GetNextHopTableId() uint32 {
	if m != nil {
		return m.NextHopTableId
	}
	return 0
}

func (m *XtcFwdingPolicyPathBag) GetIsProtected() bool {
	if m != nil {
		return m.IsProtected
	}
	return false
}

func (m *XtcFwdingPolicyPathBag) GetIsPureBkup() bool {
	if m != nil {
		return m.IsPureBkup
	}
	return false
}

func (m *XtcFwdingPolicyPathBag) GetLoadMetric() uint32 {
	if m != nil {
		return m.LoadMetric
	}
	return 0
}

func (m *XtcFwdingPolicyPathBag) GetPathId() uint32 {
	if m != nil {
		return m.PathId
	}
	return 0
}

func (m *XtcFwdingPolicyPathBag) GetBkupPathId() uint32 {
	if m != nil {
		return m.BkupPathId
	}
	return 0
}

func (m *XtcFwdingPolicyPathBag) GetLabelStack() []uint32 {
	if m != nil {
		return m.LabelStack
	}
	return nil
}

func (m *XtcFwdingPolicyPathBag) GetSegmentListName() string {
	if m != nil {
		return m.SegmentListName
	}
	return ""
}

func (m *XtcFwdingPolicyPathBag) GetAreStatsValid() bool {
	if m != nil {
		return m.AreStatsValid
	}
	return false
}

func (m *XtcFwdingPolicyPathBag) GetForwardingStatsPkts() uint64 {
	if m != nil {
		return m.ForwardingStatsPkts
	}
	return 0
}

func (m *XtcFwdingPolicyPathBag) GetForwardingStatsBytes() uint64 {
	if m != nil {
		return m.ForwardingStatsBytes
	}
	return 0
}

type XtcFwdingPolicyBag struct {
	PolicyName           string                    `protobuf:"bytes,50,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	Color                uint32                    `protobuf:"varint,51,opt,name=color,proto3" json:"color,omitempty"`
	EndpointAddress      *XtcIpAddrType            `protobuf:"bytes,52,opt,name=endpoint_address,json=endpointAddress,proto3" json:"endpoint_address,omitempty"`
	IsLocalLabelValid    bool                      `protobuf:"varint,53,opt,name=is_local_label_valid,json=isLocalLabelValid,proto3" json:"is_local_label_valid,omitempty"`
	LocalLabel           uint32                    `protobuf:"varint,54,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	AreStatsValid        bool                      `protobuf:"varint,55,opt,name=are_stats_valid,json=areStatsValid,proto3" json:"are_stats_valid,omitempty"`
	ForwardingStatsPkts  uint64                    `protobuf:"varint,56,opt,name=forwarding_stats_pkts,json=forwardingStatsPkts,proto3" json:"forwarding_stats_pkts,omitempty"`
	ForwardingStatsBytes uint64                    `protobuf:"varint,57,opt,name=forwarding_stats_bytes,json=forwardingStatsBytes,proto3" json:"forwarding_stats_bytes,omitempty"`
	Paths                []*XtcFwdingPolicyPathBag `protobuf:"bytes,58,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *XtcFwdingPolicyBag) Reset()         { *m = XtcFwdingPolicyBag{} }
func (m *XtcFwdingPolicyBag) String() string { return proto.CompactTextString(m) }
func (*XtcFwdingPolicyBag) ProtoMessage()    {}
func (*XtcFwdingPolicyBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0bea5c4e410cb41, []int{3}
}

func (m *XtcFwdingPolicyBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XtcFwdingPolicyBag.Unmarshal(m, b)
}
func (m *XtcFwdingPolicyBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XtcFwdingPolicyBag.Marshal(b, m, deterministic)
}
func (m *XtcFwdingPolicyBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XtcFwdingPolicyBag.Merge(m, src)
}
func (m *XtcFwdingPolicyBag) XXX_Size() int {
	return xxx_messageInfo_XtcFwdingPolicyBag.Size(m)
}
func (m *XtcFwdingPolicyBag) XXX_DiscardUnknown() {
	xxx_messageInfo_XtcFwdingPolicyBag.DiscardUnknown(m)
}

var xxx_messageInfo_XtcFwdingPolicyBag proto.InternalMessageInfo

func (m *XtcFwdingPolicyBag) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *XtcFwdingPolicyBag) GetColor() uint32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *XtcFwdingPolicyBag) GetEndpointAddress() *XtcIpAddrType {
	if m != nil {
		return m.EndpointAddress
	}
	return nil
}

func (m *XtcFwdingPolicyBag) GetIsLocalLabelValid() bool {
	if m != nil {
		return m.IsLocalLabelValid
	}
	return false
}

func (m *XtcFwdingPolicyBag) GetLocalLabel() uint32 {
	if m != nil {
		return m.LocalLabel
	}
	return 0
}

func (m *XtcFwdingPolicyBag) GetAreStatsValid() bool {
	if m != nil {
		return m.AreStatsValid
	}
	return false
}

func (m *XtcFwdingPolicyBag) GetForwardingStatsPkts() uint64 {
	if m != nil {
		return m.ForwardingStatsPkts
	}
	return 0
}

func (m *XtcFwdingPolicyBag) GetForwardingStatsBytes() uint64 {
	if m != nil {
		return m.ForwardingStatsBytes
	}
	return 0
}

func (m *XtcFwdingPolicyBag) GetPaths() []*XtcFwdingPolicyPathBag {
	if m != nil {
		return m.Paths
	}
	return nil
}

func init() {
	proto.RegisterType((*XtcFwdingPolicyBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.policy_forwardings.policy_forwarding.xtc_fwding_policy_bag_KEYS")
	proto.RegisterType((*XtcIpAddrType)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.policy_forwardings.policy_forwarding.xtc_ip_addr_type")
	proto.RegisterType((*XtcFwdingPolicyPathBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.policy_forwardings.policy_forwarding.xtc_fwding_policy_path_bag")
	proto.RegisterType((*XtcFwdingPolicyBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.policy_forwardings.policy_forwarding.xtc_fwding_policy_bag")
}

func init() { proto.RegisterFile("xtc_fwding_policy_bag.proto", fileDescriptor_c0bea5c4e410cb41) }

var fileDescriptor_c0bea5c4e410cb41 = []byte{
	// 614 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x55, 0x7e, 0xeb, 0xba, 0xcd, 0x59, 0x7f, 0x5b, 0xcd, 0x06, 0x16, 0x3c, 0x10, 0xfa, 0x80,
	0x0a, 0x12, 0x05, 0x75, 0x63, 0xfc, 0x79, 0x63, 0x12, 0x12, 0x85, 0x81, 0xaa, 0x14, 0x21, 0x21,
	0x21, 0x59, 0x4e, 0xe2, 0x74, 0x56, 0xd3, 0xd8, 0xb2, 0xdd, 0xad, 0x95, 0x78, 0xe7, 0x3b, 0xf0,
	0x51, 0xf8, 0x74, 0xe8, 0xda, 0x49, 0x37, 0x6d, 0xe5, 0x01, 0xb4, 0xb7, 0xf8, 0x9c, 0x73, 0xef,
	0xf5, 0x3d, 0xf7, 0xc6, 0xe8, 0xde, 0xdc, 0xa6, 0x34, 0x3f, 0xcf, 0x44, 0x39, 0xa6, 0x4a, 0x16,
	0x22, 0x5d, 0xd0, 0x84, 0x8d, 0x7b, 0x4a, 0x4b, 0x2b, 0xf1, 0xfb, 0x54, 0x98, 0x54, 0x52, 0x21,
	0x0d, 0x9d, 0x6b, 0x2a, 0xca, 0x5c, 0x33, 0x0a, 0x7a, 0x36, 0xe6, 0xa5, 0xa5, 0x52, 0x71, 0xdd,
	0x9b, 0xdb, 0xb4, 0x57, 0xc5, 0xe5, 0x52, 0x9f, 0x33, 0x0d, 0x99, 0xcc, 0x75, 0xa8, 0xf3, 0x0c,
	0xdd, 0x5d, 0x59, 0x8a, 0x7e, 0x78, 0xfb, 0x75, 0x84, 0x31, 0x6a, 0x94, 0x6c, 0xca, 0x49, 0x10,
	0x05, 0xdd, 0xad, 0xd8, 0x7d, 0x77, 0x46, 0x68, 0x17, 0x22, 0x84, 0xa2, 0x2c, 0xcb, 0x34, 0xb5,
	0x0b, 0xc5, 0xf1, 0x1d, 0xb4, 0xc1, 0x72, 0x7a, 0x49, 0xda, 0x64, 0xf9, 0x27, 0x36, 0xe5, 0x90,
	0x40, 0xa8, 0xb3, 0x43, 0xf2, 0x9f, 0x4f, 0x00, 0xdf, 0x15, 0x76, 0x44, 0xd6, 0x96, 0xd8, 0x51,
	0xe7, 0x67, 0x63, 0xd5, 0x3d, 0x14, 0xb3, 0xa7, 0x70, 0x19, 0xfc, 0x04, 0x61, 0x39, 0xb3, 0x63,
	0x09, 0x9c, 0x28, 0x2d, 0xd7, 0x39, 0x4b, 0xeb, 0x52, 0xed, 0x9a, 0x19, 0xd4, 0x04, 0xee, 0xa0,
	0x56, 0xc9, 0xe7, 0x96, 0x9e, 0x4a, 0x45, 0x2f, 0x95, 0x0f, 0x01, 0x7c, 0x27, 0xd5, 0x00, 0x6e,
	0xf1, 0x08, 0xb5, 0x97, 0x1a, 0xcb, 0x92, 0x82, 0x53, 0x91, 0xb9, 0x2b, 0xb5, 0xe2, 0xff, 0x2b,
	0xdd, 0x67, 0x80, 0x07, 0x19, 0x7e, 0x80, 0xb6, 0x85, 0xa1, 0xe0, 0x3d, 0x4f, 0x2d, 0xcf, 0x48,
	0x23, 0x0a, 0xba, 0x9b, 0x71, 0x28, 0xcc, 0xb0, 0x86, 0x70, 0xe4, 0x25, 0x33, 0xcd, 0x69, 0x32,
	0x99, 0x29, 0xb2, 0xee, 0x24, 0x48, 0x98, 0xe1, 0x4c, 0xf3, 0xe3, 0xc9, 0x4c, 0xe1, 0xfb, 0x28,
	0x2c, 0x24, 0xcb, 0xe8, 0x94, 0x5b, 0x2d, 0x52, 0xd2, 0x74, 0x95, 0x10, 0x40, 0x1f, 0x1d, 0x02,
	0x1e, 0xba, 0x7e, 0x45, 0x46, 0x36, 0x1c, 0xd9, 0x84, 0xe3, 0xc0, 0xe5, 0x86, 0x9c, 0xb4, 0x66,
	0x37, 0x7d, 0x28, 0x60, 0x43, 0xaf, 0x80, 0xdc, 0x2c, 0xe1, 0x05, 0x35, 0x96, 0xa5, 0x13, 0xb2,
	0x15, 0xad, 0xb9, 0xdc, 0x00, 0x8d, 0x00, 0xc1, 0x8f, 0x51, 0xdb, 0xf0, 0xf1, 0x14, 0xd6, 0xa3,
	0x10, 0xc6, 0xfa, 0x49, 0x21, 0x67, 0xca, 0x4e, 0x45, 0x9c, 0x08, 0x63, 0xdd, 0xc8, 0x1e, 0xa2,
	0x1d, 0xa6, 0x39, 0xa4, 0xb2, 0x86, 0x9e, 0xb1, 0x42, 0x64, 0x24, 0x74, 0xdd, 0xb4, 0x98, 0xe6,
	0x23, 0x40, 0xbf, 0x00, 0x88, 0xfb, 0x68, 0xff, 0x62, 0x8f, 0x2a, 0xb9, 0x9a, 0x58, 0x43, 0xb6,
	0xa3, 0xa0, 0xdb, 0x88, 0x6f, 0x5d, 0x90, 0x2e, 0x68, 0x38, 0xb1, 0x06, 0x1f, 0xa2, 0xdb, 0xd7,
	0x62, 0x92, 0x85, 0xe5, 0x86, 0xb4, 0x5c, 0xd0, 0xde, 0x95, 0xa0, 0x63, 0xe0, 0x3a, 0xbf, 0x1a,
	0x68, 0x7f, 0xe5, 0x92, 0x42, 0xe3, 0xd5, 0xc9, 0x75, 0xd4, 0x77, 0x1d, 0x21, 0x0f, 0xb9, 0x66,
	0xf6, 0xd0, 0x7a, 0x2a, 0x0b, 0xa9, 0xc9, 0x81, 0x33, 0xcd, 0x1f, 0xf0, 0x8f, 0x00, 0xed, 0xf2,
	0x32, 0x53, 0x52, 0x94, 0xd6, 0x6d, 0x31, 0x37, 0x86, 0x1c, 0x46, 0x41, 0x37, 0xec, 0x7f, 0xeb,
	0xdd, 0xdc, 0xcf, 0xd5, 0xbb, 0xfa, 0x9f, 0xc4, 0x3b, 0x75, 0xd5, 0x37, 0xbe, 0x28, 0x7e, 0x8a,
	0xf6, 0x84, 0xa1, 0x85, 0x4c, 0x59, 0x41, 0xfd, 0x08, 0xbd, 0xe3, 0xcf, 0x9d, 0xe3, 0x6d, 0x61,
	0x4e, 0x80, 0x3a, 0x01, 0xc6, 0xbb, 0xee, 0xd6, 0x68, 0xa9, 0x26, 0x47, 0xf5, 0x1a, 0xd5, 0xaa,
	0x55, 0xe3, 0x7b, 0xf1, 0x57, 0xe3, 0x7b, 0xf9, 0x2f, 0xe3, 0x7b, 0xf5, 0xe7, 0xf1, 0xe1, 0xef,
	0x68, 0x1d, 0x56, 0xd7, 0x90, 0xd7, 0xd1, 0x5a, 0x37, 0xec, 0xe7, 0x37, 0xed, 0xf0, 0xea, 0x37,
	0x23, 0xf6, 0x45, 0x93, 0xa6, 0x7b, 0x33, 0x0f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xae,
	0x68, 0xba, 0x52, 0x05, 0x00, 0x00,
}
