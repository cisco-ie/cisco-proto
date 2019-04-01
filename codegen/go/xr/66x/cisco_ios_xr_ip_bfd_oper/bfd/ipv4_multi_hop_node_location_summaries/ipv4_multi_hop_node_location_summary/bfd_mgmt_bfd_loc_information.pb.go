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
// source: bfd_mgmt_bfd_loc_information.proto

package cisco_ios_xr_ip_bfd_oper_bfd_ipv4_multi_hop_node_location_summaries_ipv4_multi_hop_node_location_summary

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

type BfdMgmtBfdLocInformation_KEYS struct {
	Location             string   `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdMgmtBfdLocInformation_KEYS) Reset()         { *m = BfdMgmtBfdLocInformation_KEYS{} }
func (m *BfdMgmtBfdLocInformation_KEYS) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtBfdLocInformation_KEYS) ProtoMessage()    {}
func (*BfdMgmtBfdLocInformation_KEYS) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf90ce4d6b90762, []int{0}
}

func (m *BfdMgmtBfdLocInformation_KEYS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS.Unmarshal(m, b)
}
func (m *BfdMgmtBfdLocInformation_KEYS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS.Marshal(b, m, deterministic)
}
func (m *BfdMgmtBfdLocInformation_KEYS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS.Merge(m, src)
}
func (m *BfdMgmtBfdLocInformation_KEYS) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS.Size(m)
}
func (m *BfdMgmtBfdLocInformation_KEYS) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtBfdLocInformation_KEYS proto.InternalMessageInfo

func (m *BfdMgmtBfdLocInformation_KEYS) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

type BfdLocMgmtSessionStateInformation struct {
	TotalCount           uint32   `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	UpCount              uint32   `protobuf:"varint,2,opt,name=up_count,json=upCount,proto3" json:"up_count,omitempty"`
	DownCount            uint32   `protobuf:"varint,3,opt,name=down_count,json=downCount,proto3" json:"down_count,omitempty"`
	UnknownCount         uint32   `protobuf:"varint,4,opt,name=unknown_count,json=unknownCount,proto3" json:"unknown_count,omitempty"`
	RetryCount           uint32   `protobuf:"varint,5,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`
	StandbyCount         uint32   `protobuf:"varint,6,opt,name=standby_count,json=standbyCount,proto3" json:"standby_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BfdLocMgmtSessionStateInformation) Reset()         { *m = BfdLocMgmtSessionStateInformation{} }
func (m *BfdLocMgmtSessionStateInformation) String() string { return proto.CompactTextString(m) }
func (*BfdLocMgmtSessionStateInformation) ProtoMessage()    {}
func (*BfdLocMgmtSessionStateInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf90ce4d6b90762, []int{1}
}

func (m *BfdLocMgmtSessionStateInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdLocMgmtSessionStateInformation.Unmarshal(m, b)
}
func (m *BfdLocMgmtSessionStateInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdLocMgmtSessionStateInformation.Marshal(b, m, deterministic)
}
func (m *BfdLocMgmtSessionStateInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdLocMgmtSessionStateInformation.Merge(m, src)
}
func (m *BfdLocMgmtSessionStateInformation) XXX_Size() int {
	return xxx_messageInfo_BfdLocMgmtSessionStateInformation.Size(m)
}
func (m *BfdLocMgmtSessionStateInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdLocMgmtSessionStateInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdLocMgmtSessionStateInformation proto.InternalMessageInfo

func (m *BfdLocMgmtSessionStateInformation) GetTotalCount() uint32 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *BfdLocMgmtSessionStateInformation) GetUpCount() uint32 {
	if m != nil {
		return m.UpCount
	}
	return 0
}

func (m *BfdLocMgmtSessionStateInformation) GetDownCount() uint32 {
	if m != nil {
		return m.DownCount
	}
	return 0
}

func (m *BfdLocMgmtSessionStateInformation) GetUnknownCount() uint32 {
	if m != nil {
		return m.UnknownCount
	}
	return 0
}

func (m *BfdLocMgmtSessionStateInformation) GetRetryCount() uint32 {
	if m != nil {
		return m.RetryCount
	}
	return 0
}

func (m *BfdLocMgmtSessionStateInformation) GetStandbyCount() uint32 {
	if m != nil {
		return m.StandbyCount
	}
	return 0
}

type BfdMgmtBfdLocInformation struct {
	SessionState         *BfdLocMgmtSessionStateInformation `protobuf:"bytes,50,opt,name=session_state,json=sessionState,proto3" json:"session_state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *BfdMgmtBfdLocInformation) Reset()         { *m = BfdMgmtBfdLocInformation{} }
func (m *BfdMgmtBfdLocInformation) String() string { return proto.CompactTextString(m) }
func (*BfdMgmtBfdLocInformation) ProtoMessage()    {}
func (*BfdMgmtBfdLocInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_adf90ce4d6b90762, []int{2}
}

func (m *BfdMgmtBfdLocInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BfdMgmtBfdLocInformation.Unmarshal(m, b)
}
func (m *BfdMgmtBfdLocInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BfdMgmtBfdLocInformation.Marshal(b, m, deterministic)
}
func (m *BfdMgmtBfdLocInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BfdMgmtBfdLocInformation.Merge(m, src)
}
func (m *BfdMgmtBfdLocInformation) XXX_Size() int {
	return xxx_messageInfo_BfdMgmtBfdLocInformation.Size(m)
}
func (m *BfdMgmtBfdLocInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_BfdMgmtBfdLocInformation.DiscardUnknown(m)
}

var xxx_messageInfo_BfdMgmtBfdLocInformation proto.InternalMessageInfo

func (m *BfdMgmtBfdLocInformation) GetSessionState() *BfdLocMgmtSessionStateInformation {
	if m != nil {
		return m.SessionState
	}
	return nil
}

func init() {
	proto.RegisterType((*BfdMgmtBfdLocInformation_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_multi_hop_node_location_summaries.ipv4_multi_hop_node_location_summary.bfd_mgmt_bfd_loc_information_KEYS")
	proto.RegisterType((*BfdLocMgmtSessionStateInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_multi_hop_node_location_summaries.ipv4_multi_hop_node_location_summary.bfd_loc_mgmt_session_state_information")
	proto.RegisterType((*BfdMgmtBfdLocInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_multi_hop_node_location_summaries.ipv4_multi_hop_node_location_summary.bfd_mgmt_bfd_loc_information")
}

func init() { proto.RegisterFile("bfd_mgmt_bfd_loc_information.proto", fileDescriptor_adf90ce4d6b90762) }

var fileDescriptor_adf90ce4d6b90762 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x3d, 0x4e, 0xf3, 0x40,
	0x10, 0x86, 0xb5, 0xdf, 0x07, 0x21, 0x99, 0x24, 0x8d, 0xab, 0x80, 0x40, 0x09, 0x46, 0x42, 0xa9,
	0x5c, 0x04, 0x7a, 0x0a, 0x44, 0x45, 0x97, 0x54, 0x54, 0x23, 0xff, 0x85, 0xac, 0x88, 0x77, 0x56,
	0xfb, 0x03, 0xe4, 0x16, 0x9c, 0x86, 0xab, 0x70, 0x0b, 0xce, 0x80, 0x3c, 0x5e, 0x07, 0x68, 0xa2,
	0x54, 0x74, 0xf6, 0xfb, 0x3c, 0xfb, 0xae, 0x35, 0x63, 0x88, 0xb3, 0x65, 0x81, 0xd5, 0x63, 0xe5,
	0xb0, 0x7e, 0x58, 0x53, 0x8e, 0x52, 0x2d, 0xc9, 0x54, 0xa9, 0x93, 0xa4, 0x12, 0x6d, 0xc8, 0x51,
	0xb4, 0xca, 0xa5, 0xcd, 0x09, 0x25, 0x59, 0x7c, 0x35, 0x28, 0x35, 0xab, 0xa4, 0x4b, 0x93, 0x64,
	0xcb, 0x22, 0x91, 0xfa, 0xf9, 0x1a, 0x2b, 0xbf, 0x76, 0x12, 0x57, 0xa4, 0x51, 0x51, 0x51, 0xd6,
	0x3d, 0x5c, 0x80, 0xd6, 0x57, 0x55, 0x6a, 0x64, 0x69, 0xf7, 0xd1, 0x36, 0xf1, 0x0d, 0x9c, 0xef,
	0xfa, 0x1e, 0xbc, 0xbf, 0x7b, 0x58, 0x44, 0x27, 0xd0, 0x6d, 0x0f, 0x8e, 0xc4, 0x44, 0x4c, 0x7b,
	0xf3, 0xed, 0x7b, 0xfc, 0x29, 0xe0, 0xb2, 0x3d, 0xc8, 0x2d, 0xb6, 0xb4, 0x96, 0x6f, 0x70, 0xa9,
	0x2b, 0x7f, 0x76, 0x45, 0x63, 0xe8, 0x3b, 0x72, 0xe9, 0x1a, 0x73, 0xf2, 0xca, 0x71, 0xd3, 0x70,
	0x0e, 0x1c, 0xdd, 0xd6, 0x49, 0x74, 0x0c, 0x5d, 0xaf, 0x03, 0xfd, 0xc7, 0xf4, 0xc8, 0xeb, 0x06,
	0x9d, 0x01, 0x14, 0xf4, 0xa2, 0x02, 0xfc, 0xcf, 0xb0, 0x57, 0x27, 0x0d, 0xbe, 0x80, 0xa1, 0x57,
	0x4f, 0xea, 0xdb, 0x38, 0x60, 0x63, 0x10, 0xc2, 0x46, 0x1a, 0x43, 0xdf, 0x94, 0xce, 0x6c, 0x82,
	0x72, 0xd8, 0xdc, 0xcf, 0xd1, 0xb6, 0xc5, 0xba, 0x54, 0x15, 0x59, 0xab, 0x74, 0x9a, 0x96, 0x10,
	0xb2, 0x14, 0x7f, 0x08, 0x38, 0xdd, 0x35, 0xb2, 0xe8, 0x5d, 0xc0, 0xf0, 0xd7, 0x10, 0x46, 0xb3,
	0x89, 0x98, 0xf6, 0x67, 0x6f, 0x22, 0xf9, 0xab, 0xb5, 0x26, 0xfb, 0x6d, 0x64, 0x3e, 0x08, 0x68,
	0x51, 0x93, 0xac, 0xc3, 0x3f, 0xdf, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x43, 0x9d,
	0x03, 0xa2, 0x02, 0x00, 0x00,
}
