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
// source: red_summary.proto

package cisco_ios_xr_infra_rmf_oper_redundancy_summary

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

type RedSummary_KEYS struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedSummary_KEYS) Reset()         { *m = RedSummary_KEYS{} }
func (m *RedSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*RedSummary_KEYS) ProtoMessage()    {}
func (*RedSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a5d6dca6cb046eb, []int{0}
}

func (m *RedSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedSummary_KEYS.Unmarshal(m, b)
}
func (m *RedSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *RedSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedSummary_KEYS.Merge(m, src)
}
func (m *RedSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_RedSummary_KEYS.Size(m)
}
func (m *RedSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_RedSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_RedSummary_KEYS proto.InternalMessageInfo

type RedGroupInfo struct {
	Active               string   `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	Standby              string   `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	HaState              string   `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	NsrState             string   `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedGroupInfo) Reset()         { *m = RedGroupInfo{} }
func (m *RedGroupInfo) String() string { return proto.CompactTextString(m) }
func (*RedGroupInfo) ProtoMessage()    {}
func (*RedGroupInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a5d6dca6cb046eb, []int{1}
}

func (m *RedGroupInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedGroupInfo.Unmarshal(m, b)
}
func (m *RedGroupInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedGroupInfo.Marshal(b, m, deterministic)
}
func (m *RedGroupInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedGroupInfo.Merge(m, src)
}
func (m *RedGroupInfo) XXX_Size() int {
	return xxx_messageInfo_RedGroupInfo.Size(m)
}
func (m *RedGroupInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RedGroupInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RedGroupInfo proto.InternalMessageInfo

func (m *RedGroupInfo) GetActive() string {
	if m != nil {
		return m.Active
	}
	return ""
}

func (m *RedGroupInfo) GetStandby() string {
	if m != nil {
		return m.Standby
	}
	return ""
}

func (m *RedGroupInfo) GetHaState() string {
	if m != nil {
		return m.HaState
	}
	return ""
}

func (m *RedGroupInfo) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

type RedSummaryPair struct {
	Active               string          `protobuf:"bytes,1,opt,name=active,proto3" json:"active,omitempty"`
	Standby              string          `protobuf:"bytes,2,opt,name=standby,proto3" json:"standby,omitempty"`
	HaState              string          `protobuf:"bytes,3,opt,name=ha_state,json=haState,proto3" json:"ha_state,omitempty"`
	NsrState             string          `protobuf:"bytes,4,opt,name=nsr_state,json=nsrState,proto3" json:"nsr_state,omitempty"`
	Groupinfo            []*RedGroupInfo `protobuf:"bytes,5,rep,name=groupinfo,proto3" json:"groupinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RedSummaryPair) Reset()         { *m = RedSummaryPair{} }
func (m *RedSummaryPair) String() string { return proto.CompactTextString(m) }
func (*RedSummaryPair) ProtoMessage()    {}
func (*RedSummaryPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a5d6dca6cb046eb, []int{2}
}

func (m *RedSummaryPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedSummaryPair.Unmarshal(m, b)
}
func (m *RedSummaryPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedSummaryPair.Marshal(b, m, deterministic)
}
func (m *RedSummaryPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedSummaryPair.Merge(m, src)
}
func (m *RedSummaryPair) XXX_Size() int {
	return xxx_messageInfo_RedSummaryPair.Size(m)
}
func (m *RedSummaryPair) XXX_DiscardUnknown() {
	xxx_messageInfo_RedSummaryPair.DiscardUnknown(m)
}

var xxx_messageInfo_RedSummaryPair proto.InternalMessageInfo

func (m *RedSummaryPair) GetActive() string {
	if m != nil {
		return m.Active
	}
	return ""
}

func (m *RedSummaryPair) GetStandby() string {
	if m != nil {
		return m.Standby
	}
	return ""
}

func (m *RedSummaryPair) GetHaState() string {
	if m != nil {
		return m.HaState
	}
	return ""
}

func (m *RedSummaryPair) GetNsrState() string {
	if m != nil {
		return m.NsrState
	}
	return ""
}

func (m *RedSummaryPair) GetGroupinfo() []*RedGroupInfo {
	if m != nil {
		return m.Groupinfo
	}
	return nil
}

type RedSummary struct {
	RedPair              []*RedSummaryPair `protobuf:"bytes,50,rep,name=red_pair,json=redPair,proto3" json:"red_pair,omitempty"`
	ErrLog               string            `protobuf:"bytes,51,opt,name=err_log,json=errLog,proto3" json:"err_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RedSummary) Reset()         { *m = RedSummary{} }
func (m *RedSummary) String() string { return proto.CompactTextString(m) }
func (*RedSummary) ProtoMessage()    {}
func (*RedSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a5d6dca6cb046eb, []int{3}
}

func (m *RedSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedSummary.Unmarshal(m, b)
}
func (m *RedSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedSummary.Marshal(b, m, deterministic)
}
func (m *RedSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedSummary.Merge(m, src)
}
func (m *RedSummary) XXX_Size() int {
	return xxx_messageInfo_RedSummary.Size(m)
}
func (m *RedSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_RedSummary.DiscardUnknown(m)
}

var xxx_messageInfo_RedSummary proto.InternalMessageInfo

func (m *RedSummary) GetRedPair() []*RedSummaryPair {
	if m != nil {
		return m.RedPair
	}
	return nil
}

func (m *RedSummary) GetErrLog() string {
	if m != nil {
		return m.ErrLog
	}
	return ""
}

func init() {
	proto.RegisterType((*RedSummary_KEYS)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_KEYS")
	proto.RegisterType((*RedGroupInfo)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_group_info")
	proto.RegisterType((*RedSummaryPair)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary_pair")
	proto.RegisterType((*RedSummary)(nil), "cisco_ios_xr_infra_rmf_oper.redundancy.summary.red_summary")
}

func init() { proto.RegisterFile("red_summary.proto", fileDescriptor_5a5d6dca6cb046eb) }

var fileDescriptor_5a5d6dca6cb046eb = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x31, 0x4f, 0xf3, 0x30,
	0x18, 0x84, 0x95, 0xaf, 0x1f, 0x4d, 0xf2, 0x56, 0x42, 0xe0, 0x01, 0x8c, 0x58, 0xaa, 0x4c, 0x9d,
	0x3c, 0xb4, 0x3b, 0x62, 0x61, 0x82, 0x01, 0xa5, 0x13, 0x02, 0xc9, 0x72, 0x63, 0x37, 0xb5, 0x44,
	0xec, 0xe8, 0xb5, 0x83, 0x08, 0x2b, 0xff, 0x92, 0x5f, 0x83, 0xec, 0x06, 0x91, 0x8e, 0x30, 0x30,
	0x9e, 0x4f, 0xbe, 0xf3, 0x3d, 0x32, 0x9c, 0xa2, 0x92, 0xdc, 0x75, 0x4d, 0x23, 0xb0, 0x67, 0x2d,
	0x5a, 0x6f, 0x09, 0xab, 0xb4, 0xab, 0x2c, 0xd7, 0xd6, 0xf1, 0x57, 0xe4, 0xda, 0x6c, 0x51, 0x70,
	0x6c, 0xb6, 0xdc, 0xb6, 0x0a, 0x19, 0x2a, 0xd9, 0x19, 0x29, 0x4c, 0xd5, 0xb3, 0xe1, 0x56, 0x41,
	0xe0, 0x64, 0x14, 0xc2, 0x6f, 0x6f, 0x1e, 0xd6, 0xc5, 0x1b, 0x1c, 0x87, 0xb3, 0x1a, 0x6d, 0xd7,
	0x86, 0x08, 0x4b, 0xce, 0x60, 0x2a, 0x2a, 0xaf, 0x5f, 0x14, 0x4d, 0xe6, 0xc9, 0x22, 0x2f, 0x07,
	0x45, 0x28, 0xa4, 0xce, 0x0b, 0x23, 0x37, 0x3d, 0xfd, 0x17, 0x8d, 0x2f, 0x49, 0x2e, 0x20, 0xdb,
	0x09, 0xee, 0xbc, 0xf0, 0x8a, 0x4e, 0xf6, 0xd6, 0x4e, 0xac, 0x83, 0x24, 0x97, 0x90, 0x1b, 0x87,
	0x83, 0xf7, 0x3f, 0x7a, 0x99, 0x71, 0x18, 0xcd, 0xe2, 0x23, 0x39, 0x7c, 0x50, 0x2b, 0x34, 0xfe,
	0x5d, 0x3d, 0x79, 0x82, 0x3c, 0xce, 0x0e, 0xab, 0xe9, 0xd1, 0x7c, 0xb2, 0x98, 0x2d, 0xaf, 0x7e,
	0x88, 0x94, 0x1d, 0xb2, 0x2b, 0xbf, 0x03, 0x8b, 0xf7, 0x04, 0x66, 0xa3, 0x71, 0xe4, 0x11, 0xb2,
	0x20, 0xc3, 0x46, 0xba, 0x8c, 0x65, 0xd7, 0xbf, 0x29, 0x1b, 0xb3, 0x2a, 0x53, 0x54, 0xf2, 0x3e,
	0x40, 0x3b, 0x87, 0x54, 0x21, 0xf2, 0x67, 0x5b, 0xd3, 0xd5, 0x9e, 0x9a, 0x42, 0xbc, 0xb3, 0xf5,
	0x66, 0x1a, 0x7f, 0xca, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0x24, 0x5f, 0x79, 0xb6, 0x3e, 0x02,
	0x00, 0x00,
}
