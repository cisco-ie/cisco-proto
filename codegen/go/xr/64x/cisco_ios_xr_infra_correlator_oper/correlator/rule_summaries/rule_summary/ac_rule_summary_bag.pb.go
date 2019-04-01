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
// source: ac_rule_summary_bag.proto

package cisco_ios_xr_infra_correlator_oper_correlator_rule_summaries_rule_summary

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

type AcRuleSummaryBag_KEYS struct {
	RuleName             string   `protobuf:"bytes,1,opt,name=rule_name,json=ruleName,proto3" json:"rule_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcRuleSummaryBag_KEYS) Reset()         { *m = AcRuleSummaryBag_KEYS{} }
func (m *AcRuleSummaryBag_KEYS) String() string { return proto.CompactTextString(m) }
func (*AcRuleSummaryBag_KEYS) ProtoMessage()    {}
func (*AcRuleSummaryBag_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_9613d5f4bdff62c9, []int{0}
}

func (m *AcRuleSummaryBag_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcRuleSummaryBag_KEYS.Unmarshal(m, b)
}
func (m *AcRuleSummaryBag_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcRuleSummaryBag_KEYS.Marshal(b, m, deterministic)
}
func (m *AcRuleSummaryBag_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcRuleSummaryBag_KEYS.Merge(m, src)
}
func (m *AcRuleSummaryBag_KEYS) XXX_Size() int {
	return xxx_messageInfo_AcRuleSummaryBag_KEYS.Size(m)
}
func (m *AcRuleSummaryBag_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_AcRuleSummaryBag_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_AcRuleSummaryBag_KEYS proto.InternalMessageInfo

func (m *AcRuleSummaryBag_KEYS) GetRuleName() string {
	if m != nil {
		return m.RuleName
	}
	return ""
}

type AcRuleSummaryBag struct {
	RuleNameXr           string   `protobuf:"bytes,50,opt,name=rule_name_xr,json=ruleNameXr,proto3" json:"rule_name_xr,omitempty"`
	Stateful             bool     `protobuf:"varint,51,opt,name=stateful,proto3" json:"stateful,omitempty"`
	RuleState            string   `protobuf:"bytes,52,opt,name=rule_state,json=ruleState,proto3" json:"rule_state,omitempty"`
	BufferedAlarmsCount  uint32   `protobuf:"varint,53,opt,name=buffered_alarms_count,json=bufferedAlarmsCount,proto3" json:"buffered_alarms_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcRuleSummaryBag) Reset()         { *m = AcRuleSummaryBag{} }
func (m *AcRuleSummaryBag) String() string { return proto.CompactTextString(m) }
func (*AcRuleSummaryBag) ProtoMessage()    {}
func (*AcRuleSummaryBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_9613d5f4bdff62c9, []int{1}
}

func (m *AcRuleSummaryBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcRuleSummaryBag.Unmarshal(m, b)
}
func (m *AcRuleSummaryBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcRuleSummaryBag.Marshal(b, m, deterministic)
}
func (m *AcRuleSummaryBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcRuleSummaryBag.Merge(m, src)
}
func (m *AcRuleSummaryBag) XXX_Size() int {
	return xxx_messageInfo_AcRuleSummaryBag.Size(m)
}
func (m *AcRuleSummaryBag) XXX_DiscardUnknown() {
	xxx_messageInfo_AcRuleSummaryBag.DiscardUnknown(m)
}

var xxx_messageInfo_AcRuleSummaryBag proto.InternalMessageInfo

func (m *AcRuleSummaryBag) GetRuleNameXr() string {
	if m != nil {
		return m.RuleNameXr
	}
	return ""
}

func (m *AcRuleSummaryBag) GetStateful() bool {
	if m != nil {
		return m.Stateful
	}
	return false
}

func (m *AcRuleSummaryBag) GetRuleState() string {
	if m != nil {
		return m.RuleState
	}
	return ""
}

func (m *AcRuleSummaryBag) GetBufferedAlarmsCount() uint32 {
	if m != nil {
		return m.BufferedAlarmsCount
	}
	return 0
}

func init() {
	proto.RegisterType((*AcRuleSummaryBag_KEYS)(nil), "cisco_ios_xr_infra_correlator_oper.correlator.rule_summaries.rule_summary.ac_rule_summary_bag_KEYS")
	proto.RegisterType((*AcRuleSummaryBag)(nil), "cisco_ios_xr_infra_correlator_oper.correlator.rule_summaries.rule_summary.ac_rule_summary_bag")
}

func init() { proto.RegisterFile("ac_rule_summary_bag.proto", fileDescriptor_9613d5f4bdff62c9) }

var fileDescriptor_9613d5f4bdff62c9 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0x49, 0x23, 0x7b, 0x83, 0x36, 0x39, 0x84, 0xa8, 0x08, 0xcb, 0x55, 0x5b, 0x6d, 0x71,
	0xa7, 0x58, 0x8b, 0x58, 0x88, 0x60, 0xb1, 0xd7, 0x68, 0x35, 0xcc, 0xc6, 0x59, 0x59, 0xd8, 0x6c,
	0x8e, 0x49, 0x02, 0x77, 0x5f, 0xc8, 0xcf, 0x29, 0x1b, 0xb9, 0x3f, 0xc5, 0x95, 0xef, 0xfd, 0xf8,
	0xbd, 0x84, 0x81, 0x1b, 0xb2, 0x28, 0x69, 0x60, 0x0c, 0xc9, 0x39, 0x92, 0x1d, 0xb6, 0xf4, 0x53,
	0x6f, 0xc4, 0x47, 0xaf, 0xdf, 0x6c, 0x1f, 0xac, 0xc7, 0xde, 0x07, 0xdc, 0x0a, 0xf6, 0x63, 0x27,
	0x84, 0xd6, 0x8b, 0xf0, 0x40, 0xd1, 0x0b, 0xfa, 0x0d, 0x4b, 0x7d, 0xcc, 0xf5, 0xc9, 0x4a, 0xcf,
	0xe1, 0x34, 0xee, 0x16, 0x4f, 0x60, 0xce, 0xbc, 0x83, 0xef, 0xaf, 0x5f, 0x6b, 0x7d, 0x07, 0xb3,
	0x0c, 0x46, 0x72, 0x6c, 0x54, 0xa9, 0xaa, 0x59, 0x53, 0x4c, 0xc5, 0x07, 0x39, 0x5e, 0xfc, 0x2a,
	0x98, 0x9f, 0x31, 0x75, 0x09, 0x97, 0x07, 0x09, 0xb7, 0x62, 0x96, 0xd9, 0x83, 0xbd, 0xf7, 0x29,
	0xfa, 0x16, 0x8a, 0x10, 0x29, 0x72, 0x97, 0x06, 0xb3, 0x2a, 0x55, 0x55, 0x34, 0x87, 0xac, 0xef,
	0x01, 0xfe, 0x17, 0xa7, 0xc2, 0x3c, 0x64, 0x37, 0x7f, 0x62, 0x3d, 0x15, 0x7a, 0x09, 0xd7, 0x6d,
	0xea, 0x3a, 0x16, 0xfe, 0x46, 0x1a, 0x48, 0x5c, 0x40, 0xeb, 0xd3, 0x18, 0xcd, 0x63, 0xa9, 0xaa,
	0xab, 0x66, 0xbe, 0x87, 0xcf, 0x99, 0xbd, 0x4c, 0xa8, 0xbd, 0xc8, 0x37, 0x5b, 0xfd, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0xe5, 0x56, 0x11, 0x50, 0x01, 0x00, 0x00,
}
