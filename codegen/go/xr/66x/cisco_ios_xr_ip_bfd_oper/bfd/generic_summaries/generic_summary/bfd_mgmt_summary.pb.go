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
// source: bfd_mgmt_summary.proto

package cisco_ios_xr_ip_bfd_oper_bfd_generic_summaries_generic_summary

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

type BfdMgmtSummary_KEYS struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtSummary_KEYS) Reset()         { *m = BfdMgmtSummary_KEYS{} }
func (m *BfdMgmtSummary_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSummary_KEYS) ProtoMessage()    {}
func (*BfdMgmtSummary_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf0207967c2f965, []int{0}
}

func (m *BfdMgmtSummary_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSummary_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtSummary_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSummary_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSummary_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSummary_KEYS.Merge(m, src)
}
func (m *BfdMgmtSummary_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSummary_KEYS.Size(m)
}
func (m *BfdMgmtSummary_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSummary_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSummary_KEYS proto.InternalMessageInfo

func (m *BfdMgmtSummary_KEYS) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type BfdMgmtSummary struct {
	NodeId               string   `protobuf:"bytes,50,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	PpsAllocatedValue    uint32   `protobuf:"varint,51,opt,name=pps_allocated_value,json=ppsAllocatedValue,proto3" json:"pps_allocated_value,omitempty"`
	PpsmpAllocatedValue  uint32   `protobuf:"varint,52,opt,name=ppsmp_allocated_value,json=ppsmpAllocatedValue,proto3" json:"ppsmp_allocated_value,omitempty"`
	PpsMaxValue          uint32   `protobuf:"varint,53,opt,name=pps_max_value,json=ppsMaxValue,proto3" json:"pps_max_value,omitempty"`
	PpsmpMaxValue        uint32   `protobuf:"varint,54,opt,name=ppsmp_max_value,json=ppsmpMaxValue,proto3" json:"ppsmp_max_value,omitempty"`
	TotalSessionNumber   uint32   `protobuf:"varint,55,opt,name=total_session_number,json=totalSessionNumber,proto3" json:"total_session_number,omitempty"`
	MpSessionNumber      uint32   `protobuf:"varint,56,opt,name=mp_session_number,json=mpSessionNumber,proto3" json:"mp_session_number,omitempty"`
	MaxSessionNumber     uint32   `protobuf:"varint,57,opt,name=max_session_number,json=maxSessionNumber,proto3" json:"max_session_number,omitempty"`
	PpsAllPercentage     uint32   `protobuf:"varint,58,opt,name=pps_all_percentage,json=ppsAllPercentage,proto3" json:"pps_all_percentage,omitempty"`
	PpsmpPercentage      uint32   `protobuf:"varint,59,opt,name=ppsmp_percentage,json=ppsmpPercentage,proto3" json:"ppsmp_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtSummary) Reset()         { *m = BfdMgmtSummary{} }
func (m *BfdMgmtSummary) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtSummary) ProtoMessage()    {}
func (*BfdMgmtSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf0207967c2f965, []int{1}
}

func (m *BfdMgmtSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtSummary.Unmarshal(m, b)
}
func (m *BfdMgmtSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtSummary.Marshal(b, m, deterministic)
}
func (m *BfdMgmtSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtSummary.Merge(m, src)
}
func (m *BfdMgmtSummary) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtSummary.Size(m)
}
func (m *BfdMgmtSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtSummary.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtSummary proto.InternalMessageInfo

func (m *BfdMgmtSummary) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *BfdMgmtSummary) GetPpsAllocatedValue() uint32 {
	if m != nil {
		return m.PpsAllocatedValue
	}
	return 0
}

func (m *BfdMgmtSummary) GetPpsmpAllocatedValue() uint32 {
	if m != nil {
		return m.PpsmpAllocatedValue
	}
	return 0
}

func (m *BfdMgmtSummary) GetPpsMaxValue() uint32 {
	if m != nil {
		return m.PpsMaxValue
	}
	return 0
}

func (m *BfdMgmtSummary) GetPpsmpMaxValue() uint32 {
	if m != nil {
		return m.PpsmpMaxValue
	}
	return 0
}

func (m *BfdMgmtSummary) GetTotalSessionNumber() uint32 {
	if m != nil {
		return m.TotalSessionNumber
	}
	return 0
}

func (m *BfdMgmtSummary) GetMpSessionNumber() uint32 {
	if m != nil {
		return m.MpSessionNumber
	}
	return 0
}

func (m *BfdMgmtSummary) GetMaxSessionNumber() uint32 {
	if m != nil {
		return m.MaxSessionNumber
	}
	return 0
}

func (m *BfdMgmtSummary) GetPpsAllPercentage() uint32 {
	if m != nil {
		return m.PpsAllPercentage
	}
	return 0
}

func (m *BfdMgmtSummary) GetPpsmpPercentage() uint32 {
	if m != nil {
		return m.PpsmpPercentage
	}
	return 0
}

func init() {
	proto.RegisterType((*BfdMgmtSummary_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.generic_summaries.generic_summary.bfd_mgmt_summary_KEYS")
	proto.RegisterType((*BfdMgmtSummary)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.generic_summaries.generic_summary.bfd_mgmt_summary")
}

func init() { proto.RegisterFile("bfd_mgmt_summary.proto", fileDescriptor_8bf0207967c2f965) }

var fileDescriptor_8bf0207967c2f965 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd2, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0x07, 0x70, 0x86, 0x30, 0x35, 0x32, 0xb6, 0x65, 0x4e, 0x8b, 0xa7, 0xb1, 0x83, 0x4c, 0x91,
	0x22, 0x9b, 0xbf, 0x05, 0xc1, 0x83, 0x07, 0x11, 0x45, 0x36, 0x10, 0x3c, 0x3d, 0xd2, 0x36, 0x1b,
	0x81, 0xa6, 0x79, 0x24, 0x9d, 0x74, 0xff, 0x9d, 0x7f, 0x9a, 0xf4, 0x65, 0x6e, 0xda, 0x1d, 0xf3,
	0xbe, 0x9f, 0x6f, 0xd2, 0x86, 0xb0, 0x83, 0x68, 0x9a, 0x80, 0x9e, 0xe9, 0x1c, 0xdc, 0x5c, 0x6b,
	0x61, 0x17, 0x21, 0x5a, 0x93, 0x1b, 0xfe, 0x10, 0x2b, 0x17, 0x1b, 0x50, 0xc6, 0x41, 0x61, 0x41,
	0x21, 0x94, 0xce, 0xa0, 0xb4, 0x61, 0x34, 0x4d, 0xc2, 0x99, 0xcc, 0xa4, 0x55, 0xf1, 0xb2, 0xa3,
	0xa4, 0xab, 0x4c, 0x16, 0xfd, 0x11, 0xeb, 0x56, 0x77, 0x86, 0x97, 0xa7, 0xcf, 0x09, 0x3f, 0x62,
	0x3b, 0xa9, 0x89, 0x45, 0xae, 0x4c, 0x16, 0xd4, 0x7a, 0xb5, 0xc1, 0xee, 0x78, 0xb5, 0xee, 0x7f,
	0x6f, 0xb1, 0x56, 0xb5, 0xc5, 0x0f, 0xd9, 0x76, 0x66, 0x12, 0x09, 0x2a, 0x09, 0x86, 0xe4, 0xeb,
	0xe5, 0xf2, 0x39, 0xe1, 0x21, 0xeb, 0x20, 0x3a, 0x10, 0x29, 0xf5, 0x65, 0x02, 0x5f, 0x22, 0x9d,
	0xcb, 0x60, 0xd4, 0xab, 0x0d, 0x1a, 0xe3, 0x36, 0xa2, 0x7b, 0xfc, 0x4d, 0x3e, 0xca, 0x80, 0x0f,
	0x59, 0x17, 0xd1, 0x69, 0xdc, 0x68, 0x5c, 0x50, 0xa3, 0x43, 0x61, 0xa5, 0xd3, 0x67, 0x8d, 0xf2,
	0x0c, 0x2d, 0x8a, 0xa5, 0xbd, 0x24, 0xbb, 0x87, 0xe8, 0x5e, 0x45, 0xe1, 0xcd, 0x31, 0x6b, 0xfa,
	0x7d, 0xd7, 0xea, 0x8a, 0x54, 0x83, 0xc6, 0x2b, 0x77, 0xce, 0xf6, 0x73, 0x93, 0x8b, 0x14, 0x9c,
	0x74, 0x4e, 0x99, 0x0c, 0xb2, 0xb9, 0x8e, 0xa4, 0x0d, 0xae, 0x09, 0x73, 0xca, 0x26, 0x3e, 0x7a,
	0xa3, 0x84, 0x9f, 0xb2, 0xb6, 0xc6, 0x2a, 0xbf, 0x21, 0xde, 0xd4, 0xf8, 0xdf, 0x9e, 0x31, 0x5e,
	0x9e, 0x5f, 0xc1, 0xb7, 0x84, 0x5b, 0x5a, 0x14, 0x1b, 0x7a, 0x79, 0x77, 0x80, 0xd2, 0xc6, 0x32,
	0xcb, 0xc5, 0x4c, 0x06, 0x77, 0x5e, 0xfb, 0xab, 0x7b, 0x5f, 0xcd, 0xf9, 0x09, 0x6b, 0xf9, 0x3f,
	0xfc, 0x63, 0xef, 0xfd, 0x67, 0xd0, 0x7c, 0x4d, 0xa3, 0x3a, 0x3d, 0x9f, 0xd1, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x70, 0x6a, 0x46, 0x06, 0x58, 0x02, 0x00, 0x00,
}
