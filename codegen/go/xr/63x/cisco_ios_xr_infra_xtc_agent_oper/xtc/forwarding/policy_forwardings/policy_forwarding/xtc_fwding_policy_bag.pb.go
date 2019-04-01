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

package cisco_ios_xr_infra_xtc_agent_oper_xtc_forwarding_policy_forwardings_policy_forwarding

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
	return fileDescriptor_c0bea5c4e410cb41, []int{1}
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
	IsLocalLabelValid    bool                      `protobuf:"varint,51,opt,name=is_local_label_valid,json=isLocalLabelValid,proto3" json:"is_local_label_valid,omitempty"`
	LocalLabel           uint32                    `protobuf:"varint,52,opt,name=local_label,json=localLabel,proto3" json:"local_label,omitempty"`
	AreStatsValid        bool                      `protobuf:"varint,53,opt,name=are_stats_valid,json=areStatsValid,proto3" json:"are_stats_valid,omitempty"`
	ForwardingStatsPkts  uint64                    `protobuf:"varint,54,opt,name=forwarding_stats_pkts,json=forwardingStatsPkts,proto3" json:"forwarding_stats_pkts,omitempty"`
	ForwardingStatsBytes uint64                    `protobuf:"varint,55,opt,name=forwarding_stats_bytes,json=forwardingStatsBytes,proto3" json:"forwarding_stats_bytes,omitempty"`
	Paths                []*XtcFwdingPolicyPathBag `protobuf:"bytes,56,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *XtcFwdingPolicyBag) Reset()         { *m = XtcFwdingPolicyBag{} }
func (m *XtcFwdingPolicyBag) String() string { return proto.CompactTextString(m) }
func (*XtcFwdingPolicyBag) ProtoMessage()    {}
func (*XtcFwdingPolicyBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0bea5c4e410cb41, []int{2}
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
	proto.RegisterType((*XtcFwdingPolicyBag_KEYS)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.forwarding.policy_forwardings.policy_forwarding.xtc_fwding_policy_bag_KEYS")
	proto.RegisterType((*XtcFwdingPolicyPathBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.forwarding.policy_forwardings.policy_forwarding.xtc_fwding_policy_path_bag")
	proto.RegisterType((*XtcFwdingPolicyBag)(nil), "cisco_ios_xr_infra_xtc_agent_oper.xtc.forwarding.policy_forwardings.policy_forwarding.xtc_fwding_policy_bag")
}

func init() { proto.RegisterFile("xtc_fwding_policy_bag.proto", fileDescriptor_c0bea5c4e410cb41) }

var fileDescriptor_c0bea5c4e410cb41 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0x13, 0x3d,
	0x14, 0xd5, 0x7c, 0x49, 0xd3, 0xd6, 0x69, 0xbe, 0x2a, 0xa6, 0x05, 0x0b, 0x16, 0x0c, 0x59, 0xa0,
	0x80, 0xc4, 0x80, 0xd2, 0xf0, 0xb3, 0xae, 0x84, 0x44, 0x44, 0x40, 0x51, 0x02, 0x48, 0xac, 0xae,
	0x3c, 0x1e, 0x27, 0xb5, 0x32, 0x19, 0x1b, 0xdb, 0x49, 0xd3, 0x27, 0xe0, 0x1d, 0x78, 0x00, 0x9e,
	0x13, 0x5d, 0x4f, 0xa6, 0x91, 0x68, 0xba, 0x00, 0x89, 0x5d, 0xe6, 0x9c, 0x73, 0x8f, 0x75, 0xce,
	0xb5, 0x43, 0x1e, 0xac, 0xbd, 0x80, 0xe9, 0x65, 0xa6, 0x8a, 0x19, 0x18, 0x9d, 0x2b, 0x71, 0x05,
	0x29, 0x9f, 0x25, 0xc6, 0x6a, 0xaf, 0xe9, 0x67, 0xa1, 0x9c, 0xd0, 0xa0, 0xb4, 0x83, 0xb5, 0x05,
	0x55, 0x4c, 0x2d, 0x07, 0xd4, 0xf3, 0x99, 0x2c, 0x3c, 0x68, 0x23, 0x6d, 0xb2, 0xf6, 0x22, 0x99,
	0x6a, 0x7b, 0xc9, 0x2d, 0x5a, 0x24, 0x1b, 0x8b, 0x2d, 0xe2, 0x6e, 0x42, 0x9d, 0x17, 0xe4, 0xfe,
	0xce, 0x53, 0xe1, 0xfd, 0xdb, 0xaf, 0x13, 0x4a, 0x49, 0xbd, 0xe0, 0x0b, 0xc9, 0xa2, 0x38, 0xea,
	0x1e, 0x8e, 0xc3, 0xef, 0xce, 0x8f, 0xfa, 0xae, 0x11, 0xc3, 0xfd, 0x05, 0xce, 0xd1, 0x67, 0x84,
	0xea, 0xa5, 0x9f, 0x69, 0xe4, 0x54, 0xe1, 0xa5, 0x9d, 0x72, 0x51, 0x19, 0xb4, 0x2b, 0x66, 0x50,
	0x11, 0xb4, 0x43, 0x5a, 0x85, 0x5c, 0x7b, 0xb8, 0xd0, 0x06, 0x94, 0x59, 0xf5, 0xd9, 0x7f, 0x41,
	0xd9, 0x44, 0xf0, 0x9d, 0x36, 0x03, 0xb3, 0xea, 0xd3, 0x27, 0xa4, 0x7d, 0xad, 0xf1, 0x3c, 0xcd,
	0x25, 0xa8, 0x8c, 0xd5, 0xe2, 0xa8, 0xdb, 0x1a, 0xff, 0xbf, 0xd1, 0x7d, 0x42, 0x78, 0x90, 0xd1,
	0x47, 0xe4, 0x48, 0x39, 0xc0, 0xc6, 0xa4, 0xf0, 0x32, 0x63, 0xf5, 0x38, 0xea, 0x1e, 0x8c, 0x9b,
	0xca, 0x8d, 0x2a, 0x88, 0xc6, 0xa5, 0x64, 0x69, 0x25, 0xa4, 0xf3, 0xa5, 0x61, 0x7b, 0x41, 0x42,
	0x94, 0x1b, 0x2d, 0xad, 0x3c, 0x9f, 0x2f, 0x0d, 0x7d, 0x48, 0x9a, 0xb9, 0xe6, 0x19, 0x2c, 0xa4,
	0xb7, 0x4a, 0xb0, 0x46, 0x38, 0x89, 0x20, 0xf4, 0x21, 0x20, 0xf4, 0x1e, 0xd9, 0x0f, 0x79, 0x55,
	0xc6, 0xf6, 0x03, 0xd9, 0xc0, 0xcf, 0x41, 0xf0, 0x46, 0x4f, 0xa8, 0xd8, 0x83, 0x72, 0x14, 0xb1,
	0x51, 0xa9, 0x40, 0x6f, 0x9e, 0xca, 0x1c, 0x9c, 0xe7, 0x62, 0xce, 0x0e, 0xe3, 0x5a, 0xf0, 0x46,
	0x68, 0x82, 0x08, 0x7d, 0x4a, 0xda, 0x4e, 0xce, 0x16, 0xb8, 0xd4, 0x5c, 0x39, 0x0f, 0xa1, 0x7f,
	0x12, 0x4a, 0x39, 0xde, 0x10, 0x43, 0xe5, 0xfc, 0x47, 0xbe, 0x90, 0xf4, 0x31, 0x39, 0xe6, 0x56,
	0xa2, 0x95, 0x77, 0xb0, 0xe2, 0xb9, 0xca, 0x58, 0x33, 0xa4, 0x69, 0x71, 0x2b, 0x27, 0x88, 0x7e,
	0x41, 0x90, 0xf6, 0xc8, 0xe9, 0x76, 0xe5, 0x1b, 0xb9, 0x99, 0x7b, 0xc7, 0x8e, 0xe2, 0xa8, 0x5b,
	0x1f, 0xdf, 0xd9, 0x92, 0x61, 0x68, 0x34, 0xf7, 0x8e, 0xf6, 0xc9, 0xdd, 0x1b, 0x33, 0xe9, 0x95,
	0x97, 0x8e, 0xb5, 0xc2, 0xd0, 0xc9, 0x6f, 0x43, 0xe7, 0xc8, 0x75, 0x7e, 0xd6, 0xc8, 0xe9, 0xce,
	0xfb, 0x84, 0xc1, 0x37, 0x5f, 0x21, 0x51, 0x2f, 0x24, 0x22, 0x25, 0x14, 0xc2, 0x3c, 0x27, 0x27,
	0xca, 0x41, 0xae, 0x05, 0xcf, 0xa1, 0xac, 0xa8, 0x4c, 0x74, 0x16, 0x12, 0xb5, 0x95, 0x1b, 0x22,
	0x35, 0x44, 0xa6, 0x4c, 0x15, 0xd6, 0x74, 0xad, 0x66, 0xfd, 0x6a, 0x4d, 0x95, 0x6a, 0x57, 0x3d,
	0x2f, 0xff, 0xa8, 0x9e, 0x57, 0x7f, 0x53, 0xcf, 0xeb, 0xdb, 0xeb, 0xa1, 0xdf, 0x23, 0xb2, 0x87,
	0x77, 0xc3, 0xb1, 0x37, 0x71, 0xad, 0xdb, 0xec, 0x7d, 0x4b, 0xfe, 0xc9, 0xab, 0x4e, 0x6e, 0x7f,
	0x9f, 0xe3, 0xf2, 0xfc, 0xb4, 0x11, 0xfe, 0x55, 0xce, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0xf7,
	0xac, 0x2b, 0xd9, 0x74, 0x04, 0x00, 0x00,
}
