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

package cisco_ios_xr_ip_bfd_oper_bfd_ipv4_single_hop_location_summaries_ipv4_single_hop_location_summary

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
	LocationName         string   `protobuf:"bytes,1,opt,name=location_name,json=locationName,proto3" json:"location_name,omitempty"`
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

func (m *BfdMgmtBfdLocInformation_KEYS) GetLocationName() string {
	if m != nil {
		return m.LocationName
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
	proto.RegisterType((*BfdMgmtBfdLocInformation_KEYS)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_single_hop_location_summaries.ipv4_single_hop_location_summary.bfd_mgmt_bfd_loc_information_KEYS")
	proto.RegisterType((*BfdLocMgmtSessionStateInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_single_hop_location_summaries.ipv4_single_hop_location_summary.bfd_loc_mgmt_session_state_information")
	proto.RegisterType((*BfdMgmtBfdLocInformation)(nil), "cisco_ios_xr_ip_bfd_oper.bfd.ipv4_single_hop_location_summaries.ipv4_single_hop_location_summary.bfd_mgmt_bfd_loc_information")
}

func init() { proto.RegisterFile("bfd_mgmt_bfd_loc_information.proto", fileDescriptor_adf90ce4d6b90762) }

var fileDescriptor_adf90ce4d6b90762 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x65, 0xfe, 0x14, 0xfa, 0xda, 0x2c, 0x99, 0x8a, 0x04, 0x6a, 0x49, 0x25, 0xd4, 0x29,
	0x43, 0xe1, 0x06, 0x08, 0x09, 0x09, 0x89, 0xa1, 0x9d, 0x98, 0x8c, 0x93, 0xb8, 0xc5, 0xa2, 0xf6,
	0xb3, 0x6c, 0x07, 0xe8, 0x0d, 0x38, 0x0d, 0x57, 0xe1, 0x26, 0x9c, 0x01, 0xe5, 0x25, 0x29, 0xb0,
	0x14, 0x16, 0xb6, 0xe8, 0xfb, 0xfd, 0xfc, 0x39, 0x7a, 0xcf, 0x90, 0x64, 0x8b, 0x82, 0xeb, 0xa5,
	0x0e, 0xbc, 0xfa, 0x58, 0x61, 0xce, 0x95, 0x59, 0xa0, 0xd3, 0x22, 0x28, 0x34, 0xa9, 0x75, 0x18,
	0x30, 0xbe, 0xcf, 0x95, 0xcf, 0x91, 0x2b, 0xf4, 0xfc, 0xc5, 0x71, 0x65, 0x49, 0x45, 0x2b, 0x5d,
	0x9a, 0x2d, 0x8a, 0x54, 0xd9, 0xa7, 0x0b, 0xee, 0x95, 0x59, 0xae, 0x24, 0x7f, 0x40, 0x5b, 0x75,
	0xd0, 0x61, 0xee, 0x4b, 0xad, 0x85, 0x53, 0xd2, 0xff, 0xa6, 0xac, 0x93, 0x6b, 0x38, 0xdd, 0xf6,
	0x1f, 0xfc, 0xe6, 0xea, 0x6e, 0x1e, 0x8f, 0x21, 0xda, 0x1c, 0x34, 0x42, 0xcb, 0x01, 0x1b, 0xb1,
	0x49, 0x77, 0xd6, 0x6f, 0xc3, 0x5b, 0xa1, 0x65, 0xf2, 0xc1, 0xe0, 0xac, 0x6d, 0xa0, 0x3a, 0x2f,
	0xbd, 0xa7, 0xab, 0x82, 0x08, 0xf2, 0x7b, 0x69, 0x3c, 0x84, 0x5e, 0xc0, 0x20, 0x56, 0x3c, 0xc7,
	0xd2, 0x04, 0x6a, 0x8b, 0x66, 0x40, 0xd1, 0x65, 0x95, 0xc4, 0x47, 0x70, 0x58, 0xda, 0x86, 0xee,
	0x10, 0x3d, 0x28, 0x6d, 0x8d, 0x4e, 0x00, 0x0a, 0x7c, 0x36, 0x0d, 0xdc, 0x25, 0xd8, 0xad, 0x92,
	0x1a, 0x8f, 0x21, 0x2a, 0xcd, 0xa3, 0xf9, 0x32, 0xf6, 0xc8, 0xe8, 0x37, 0x61, 0x2d, 0x0d, 0xa1,
	0xe7, 0x64, 0x70, 0xeb, 0x46, 0xd9, 0xaf, 0xef, 0xa7, 0x68, 0xd3, 0xe2, 0x83, 0x30, 0x45, 0xd6,
	0x2a, 0x9d, 0xba, 0xa5, 0x09, 0x49, 0x4a, 0xde, 0x19, 0x1c, 0x6f, 0x9b, 0x5d, 0xfc, 0xc6, 0x20,
	0xfa, 0x31, 0x84, 0xc1, 0x74, 0xc4, 0x26, 0xbd, 0xe9, 0x2b, 0x4b, 0xff, 0x7b, 0xaf, 0xe9, 0xdf,
	0x36, 0x31, 0xeb, 0x37, 0x68, 0x5e, 0x91, 0xac, 0x43, 0xaf, 0xee, 0xfc, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x12, 0xc5, 0x5d, 0x9f, 0x9b, 0x02, 0x00, 0x00,
}
